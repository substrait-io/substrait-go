// SPDX-License-Identifier: Apache-2.0

package plan

import (
	"fmt"
	"strconv"
	"strings"

	substraitgo "github.com/substrait-io/substrait-go/v9"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// ExpressionConverter resolves extensions and subqueries as used in expressions.
// It extends the base ExtensionRegistry to handle subquery expressions
// that may appear within other expressions.
type ExpressionConverter struct {
	expr.ExtensionRegistry
}

// SubqueryFromProto creates a subquery expression from a protobuf message
func (r *ExpressionConverter) SubqueryFromProto(sub *proto.Expression_Subquery, baseSchema *types.RecordType, reg expr.ExtensionRegistry) (expr.Expression, error) {
	switch subType := sub.SubqueryType.(type) {
	case *proto.Expression_Subquery_Scalar_:
		rel, err := RelFromProto(subType.Scalar.Input, reg)
		if err != nil {
			return nil, err
		}
		return NewScalarSubquery(rel), nil

	case *proto.Expression_Subquery_InPredicate_:
		needles := make([]expr.Expression, len(subType.InPredicate.Needles))
		for i, needle := range subType.InPredicate.Needles {
			expr, err := expr.ExprFromProto(needle, baseSchema, reg)
			if err != nil {
				return nil, fmt.Errorf("error parsing needle %d in IN predicate: %w", i, err)
			}
			needles[i] = expr
		}

		rel, err := RelFromProto(subType.InPredicate.Haystack, reg)
		if err != nil {
			return nil, err
		}

		return NewInPredicateSubquery(needles, rel), nil

	case *proto.Expression_Subquery_SetPredicate_:
		tuples, err := RelFromProto(subType.SetPredicate.Tuples, reg)
		if err != nil {
			return nil, fmt.Errorf("error parsing tuples in set predicate: %w", err)
		}
		return NewSetPredicateSubquery(SetPredicateOp(subType.SetPredicate.PredicateOp), tuples), nil
	case *proto.Expression_Subquery_SetComparison_:
		left, err := expr.ExprFromProto(subType.SetComparison.Left, baseSchema, reg)
		if err != nil {
			return nil, fmt.Errorf("error parsing left expression in set comparison: %w", err)
		}

		right, err := RelFromProto(subType.SetComparison.Right, reg)
		if err != nil {
			return nil, fmt.Errorf("error parsing right relation in set comparison: %w", err)
		}

		return NewSetComparisonSubquery(
			SetComparisonReductionOp(subType.SetComparison.ReductionOp),
			SetComparisonOp(subType.SetComparison.ComparisonOp),
			left,
			right,
		), nil

	default:
		return nil, fmt.Errorf("%w: unknown subquery type: %T", substraitgo.ErrNotImplemented, subType)
	}
}

// ScalarSubquery is a subquery that returns one row and one column
type ScalarSubquery struct {
	Input Rel
	// Subqueries are expressions and so can be the "root" of a field reference, so we embed this marker interface to denote that.
	expr.RootRefType
}

func NewScalarSubquery(input Rel) *ScalarSubquery {
	return &ScalarSubquery{Input: input}
}

func (s *ScalarSubquery) GetSubqueryType() string { return "scalar" }

func (s *ScalarSubquery) String() string {
	return fmt.Sprintf("SCALAR_SUBQUERY(%s)", s.Input)
}

func (s *ScalarSubquery) IsScalar() bool { return true }

func (s *ScalarSubquery) GetType() types.Type {
	schema := s.Input.RecordType()
	schemaTypes := schema.Types()
	if len(schemaTypes) != 1 {
		panic("scalar subquery must return exactly one column")
	}
	return schemaTypes[0]
}

func (s *ScalarSubquery) ToProto() *proto.Expression {
	return &proto.Expression{
		RexType: &proto.Expression_Subquery_{
			Subquery: &proto.Expression_Subquery{
				SubqueryType: &proto.Expression_Subquery_Scalar_{
					Scalar: &proto.Expression_Subquery_Scalar{
						Input: s.Input.ToProto(),
					},
				},
			},
		},
	}
}

func (s *ScalarSubquery) Equals(other expr.Expression) bool {
	otherScalar, ok := other.(*ScalarSubquery)
	if !ok {
		return false
	}
	return isRelEqual(s.Input, otherScalar.Input)
}

func (s *ScalarSubquery) Visit(visit expr.VisitFunc) expr.Expression {
	// ScalarSubquery doesn't contain other expressions that need visiting
	return s
}

// InPredicateSubquery checks that the left expressions are contained in the right subquery
type InPredicateSubquery struct {
	Needles  []expr.Expression // Expressions whose existence will be checked
	Haystack Rel               // Subquery to check

	// Subqueries can be the "root" of a field reference, so we embed this marker interface to denote that.
	expr.RootRefType
}

func NewInPredicateSubquery(needles []expr.Expression, haystack Rel) *InPredicateSubquery {
	return &InPredicateSubquery{
		Needles:  needles,
		Haystack: haystack,
	}
}

func (s *InPredicateSubquery) String() string {
	var b strings.Builder
	if len(s.Needles) == 1 {
		b.WriteString(s.Needles[0].String())
	} else {
		b.WriteByte('(')
		for i, needle := range s.Needles {
			if i > 0 {
				b.WriteString(", ")
			}
			b.WriteString(needle.String())
		}
		b.WriteByte(')')
	}
	b.WriteString(" IN (")
	b.WriteString(fmt.Sprintf("%s", s.Haystack))
	b.WriteByte(')')
	return b.String()
}

func (s *InPredicateSubquery) IsScalar() bool {
	for _, needle := range s.Needles {
		if !needle.IsScalar() {
			return false
		}
	}
	return true
}

func (s *InPredicateSubquery) GetType() types.Type {
	return &types.BooleanType{Nullability: types.NullabilityRequired}
}

func (s *InPredicateSubquery) ToProto() *proto.Expression {
	needles := make([]*proto.Expression, len(s.Needles))
	for i, needle := range s.Needles {
		needles[i] = needle.ToProto()
	}

	return &proto.Expression{
		RexType: &proto.Expression_Subquery_{
			Subquery: &proto.Expression_Subquery{
				SubqueryType: &proto.Expression_Subquery_InPredicate_{
					InPredicate: &proto.Expression_Subquery_InPredicate{
						Needles:  needles,
						Haystack: s.Haystack.ToProto(),
					},
				},
			},
		},
	}
}

func (s *InPredicateSubquery) Equals(other expr.Expression) bool {
	otherInPredicate, ok := other.(*InPredicateSubquery)
	if !ok {
		return false
	}

	if len(s.Needles) != len(otherInPredicate.Needles) {
		return false
	}

	for i, needle := range s.Needles {
		if !needle.Equals(otherInPredicate.Needles[i]) {
			return false
		}
	}

	return isRelEqual(s.Haystack, otherInPredicate.Haystack)
}

func (s *InPredicateSubquery) Visit(visit expr.VisitFunc) expr.Expression {
	var out *InPredicateSubquery
	for i, needle := range s.Needles {
		afterNeedle := visit(needle)
		if out == nil && afterNeedle != needle {
			out = &InPredicateSubquery{
				Needles:  make([]expr.Expression, len(s.Needles)),
				Haystack: s.Haystack,
			}
			copy(out.Needles, s.Needles[:i])
		}
		if out != nil {
			out.Needles[i] = afterNeedle
		}
	}

	if out == nil {
		return s
	}
	return out
}

func (s *InPredicateSubquery) GetSubqueryType() string {
	return "in_predicate"
}

// SetPredicateOp indicates the kind of set predicate (EXISTS/UNIQUE) applied to a subquery.
type SetPredicateOp int32

const (
	SetPredicateOpUnspecified SetPredicateOp = 0
	SetPredicateOpExists      SetPredicateOp = 1
	SetPredicateOpUnique      SetPredicateOp = 2
)

// String returns the protobuf enum name for the set predicate operation.
func (o SetPredicateOp) String() string {
	switch o {
	case SetPredicateOpUnspecified:
		return "PREDICATE_OP_UNSPECIFIED"
	case SetPredicateOpExists:
		return "PREDICATE_OP_EXISTS"
	case SetPredicateOpUnique:
		return "PREDICATE_OP_UNIQUE"
	default:
		return strconv.Itoa(int(o))
	}
}

// SetPredicateSubquery is a predicate over a set of rows (EXISTS/UNIQUE)
type SetPredicateSubquery struct {
	Operation SetPredicateOp
	Tuples    Rel

	// Subqueries can be the "root" of a field reference, so we embed this marker interface to denote that.
	expr.RootRefType
}

func NewSetPredicateSubquery(op SetPredicateOp, tuples Rel) *SetPredicateSubquery {
	return &SetPredicateSubquery{
		Operation: op,
		Tuples:    tuples,
	}
}

func (s *SetPredicateSubquery) String() string {
	var opStr string
	switch s.Operation {
	case SetPredicateOpExists:
		opStr = "EXISTS"
	case SetPredicateOpUnique:
		opStr = "UNIQUE"
	default:
		opStr = "UNKNOWN"
	}
	return fmt.Sprintf("%s(%s)", opStr, s.Tuples)
}

func (s *SetPredicateSubquery) IsScalar() bool { return true }

func (s *SetPredicateSubquery) GetType() types.Type {
	return &types.BooleanType{Nullability: types.NullabilityRequired}
}

func (s *SetPredicateSubquery) ToProto() *proto.Expression {
	return &proto.Expression{
		RexType: &proto.Expression_Subquery_{
			Subquery: &proto.Expression_Subquery{
				SubqueryType: &proto.Expression_Subquery_SetPredicate_{
					SetPredicate: &proto.Expression_Subquery_SetPredicate{
						PredicateOp: proto.Expression_Subquery_SetPredicate_PredicateOp(s.Operation),
						Tuples:      s.Tuples.ToProto(),
					},
				},
			},
		},
	}
}

func (s *SetPredicateSubquery) Equals(other expr.Expression) bool {
	otherSetPredicate, ok := other.(*SetPredicateSubquery)
	if !ok {
		return false
	}
	return s.Operation == otherSetPredicate.Operation &&
		isRelEqual(s.Tuples, otherSetPredicate.Tuples)
}

func (s *SetPredicateSubquery) Visit(visit expr.VisitFunc) expr.Expression {
	// SetPredicateSubquery doesn't contain expressions that need visiting
	return s
}

func (s *SetPredicateSubquery) GetSubqueryType() string {
	return "set_predicate"
}

// SetComparisonReductionOp indicates how a set comparison reduces its results (ANY/ALL).
type SetComparisonReductionOp int32

const (
	SetComparisonReductionOpUnspecified SetComparisonReductionOp = 0
	SetComparisonReductionOpAny         SetComparisonReductionOp = 1
	SetComparisonReductionOpAll         SetComparisonReductionOp = 2
)

// String returns the protobuf enum name for the set comparison reduction operation.
func (o SetComparisonReductionOp) String() string {
	switch o {
	case SetComparisonReductionOpUnspecified:
		return "REDUCTION_OP_UNSPECIFIED"
	case SetComparisonReductionOpAny:
		return "REDUCTION_OP_ANY"
	case SetComparisonReductionOpAll:
		return "REDUCTION_OP_ALL"
	default:
		return strconv.Itoa(int(o))
	}
}

// SetComparisonOp indicates the comparison operator used in a set comparison.
type SetComparisonOp int32

const (
	SetComparisonOpUnspecified SetComparisonOp = 0
	SetComparisonOpEq          SetComparisonOp = 1
	SetComparisonOpNe          SetComparisonOp = 2
	SetComparisonOpLt          SetComparisonOp = 3
	SetComparisonOpGt          SetComparisonOp = 4
	SetComparisonOpLe          SetComparisonOp = 5
	SetComparisonOpGe          SetComparisonOp = 6
)

// String returns the protobuf enum name for the set comparison operation.
func (o SetComparisonOp) String() string {
	switch o {
	case SetComparisonOpUnspecified:
		return "COMPARISON_OP_UNSPECIFIED"
	case SetComparisonOpEq:
		return "COMPARISON_OP_EQ"
	case SetComparisonOpNe:
		return "COMPARISON_OP_NE"
	case SetComparisonOpLt:
		return "COMPARISON_OP_LT"
	case SetComparisonOpGt:
		return "COMPARISON_OP_GT"
	case SetComparisonOpLe:
		return "COMPARISON_OP_LE"
	case SetComparisonOpGe:
		return "COMPARISON_OP_GE"
	default:
		return strconv.Itoa(int(o))
	}
}

// SetComparisonSubquery is a subquery comparison using ANY or ALL operations
type SetComparisonSubquery struct {
	ReductionOp  SetComparisonReductionOp
	ComparisonOp SetComparisonOp
	Left         expr.Expression
	Right        Rel

	// Subqueries can be the "root" of a field reference, so we embed this marker interface to denote that.
	expr.RootRefType
}

func NewSetComparisonSubquery(
	reductionOp SetComparisonReductionOp,
	comparisonOp SetComparisonOp,
	left expr.Expression,
	right Rel,
) *SetComparisonSubquery {
	return &SetComparisonSubquery{
		ReductionOp:  reductionOp,
		ComparisonOp: comparisonOp,
		Left:         left,
		Right:        right,
	}
}

func (s *SetComparisonSubquery) String() string {
	var reductionStr, comparisonStr string

	switch s.ReductionOp {
	case SetComparisonReductionOpAny:
		reductionStr = "ANY"
	case SetComparisonReductionOpAll:
		reductionStr = "ALL"
	default:
		reductionStr = "UNKNOWN"
	}

	switch s.ComparisonOp {
	case SetComparisonOpEq:
		comparisonStr = "="
	case SetComparisonOpNe:
		comparisonStr = "!="
	case SetComparisonOpLt:
		comparisonStr = "<"
	case SetComparisonOpGt:
		comparisonStr = ">"
	case SetComparisonOpLe:
		comparisonStr = "<="
	case SetComparisonOpGe:
		comparisonStr = ">="
	default:
		comparisonStr = "?"
	}

	return fmt.Sprintf("%s %s %s(%s)", s.Left, comparisonStr, reductionStr, s.Right)
}

func (s *SetComparisonSubquery) IsScalar() bool {
	return s.Left.IsScalar()
}

func (s *SetComparisonSubquery) GetType() types.Type {
	return &types.BooleanType{Nullability: types.NullabilityRequired}
}

func (s *SetComparisonSubquery) ToProto() *proto.Expression {
	return &proto.Expression{
		RexType: &proto.Expression_Subquery_{
			Subquery: &proto.Expression_Subquery{
				SubqueryType: &proto.Expression_Subquery_SetComparison_{
					SetComparison: &proto.Expression_Subquery_SetComparison{
						ReductionOp:  proto.Expression_Subquery_SetComparison_ReductionOp(s.ReductionOp),
						ComparisonOp: proto.Expression_Subquery_SetComparison_ComparisonOp(s.ComparisonOp),
						Left:         s.Left.ToProto(),
						Right:        s.Right.ToProto(),
					},
				},
			},
		},
	}
}

func (s *SetComparisonSubquery) Equals(other expr.Expression) bool {
	otherSetComparison, ok := other.(*SetComparisonSubquery)
	if !ok {
		return false
	}

	if s.Left == nil || otherSetComparison.Left == nil {
		return s.Left == otherSetComparison.Left
	}
	if s.Right == nil || otherSetComparison.Right == nil {
		return s.Right == otherSetComparison.Right
	}

	return s.ReductionOp == otherSetComparison.ReductionOp &&
		s.ComparisonOp == otherSetComparison.ComparisonOp &&
		s.Left.Equals(otherSetComparison.Left) &&
		isRelEqual(s.Right, otherSetComparison.Right)
}

func (s *SetComparisonSubquery) Visit(visit expr.VisitFunc) expr.Expression {
	afterLeft := visit(s.Left)
	if afterLeft == s.Left {
		return s
	}

	return &SetComparisonSubquery{
		ReductionOp:  s.ReductionOp,
		ComparisonOp: s.ComparisonOp,
		Left:         afterLeft,
		Right:        s.Right,
	}
}

func (s *SetComparisonSubquery) GetSubqueryType() string {
	return "set_comparison"
}

// TODO: Implement proper relation equality comparison
// Currently using pointer equality as a temporary solution.
// This should be replaced with proper deep equality comparison
// that compares the actual structure and content of relations.
// Ideally, we should also add Equals to the Rel interface, instead of
// relying on this function.
func isRelEqual(a, b Rel) bool {
	return a == b
}
