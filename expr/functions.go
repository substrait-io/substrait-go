// SPDX-License-Identifier: Apache-2.0

package expr

import (
	"fmt"
	"strings"

	substraitgo "github.com/substrait-io/substrait-go"
	"github.com/substrait-io/substrait-go/extensions"
	"github.com/substrait-io/substrait-go/proto"
	"github.com/substrait-io/substrait-go/types"
	"golang.org/x/exp/slices"
	pb "google.golang.org/protobuf/proto"
)

func FuncArgsEqual(a, b types.FuncArg) bool {
	if a == b {
		return true
	}

	switch lhs := a.(type) {
	case Expression:
		rhs, ok := b.(Expression)
		if !ok {
			return false
		}

		return lhs.Equals(rhs)
	case types.Type:
		rhs, ok := b.(types.Type)
		if !ok {
			return false
		}

		return lhs.Equals(rhs)
	case types.Enum:
		rhs, ok := b.(types.Enum)
		if !ok {
			return false
		}

		return lhs == rhs
	}

	return false
}

type (
	SortField struct {
		Expr Expression
		Kind types.SortKind
	}

	Bound interface {
		ToProto() *proto.Expression_WindowFunction_Bound
	}

	PrecedingBound int64
	FollowingBound int64
	CurrentRow     struct{}
	Unbounded      struct{}
)

func (s *SortField) ToProto() *proto.SortField {
	ret := &proto.SortField{Expr: s.Expr.ToProto()}
	switch k := s.Kind.(type) {
	case types.SortDirection:
		ret.SortKind = &proto.SortField_Direction{
			Direction: proto.SortField_SortDirection(k)}
	case types.FunctionRef:
		ret.SortKind = &proto.SortField_ComparisonFunctionReference{
			ComparisonFunctionReference: uint32(k)}
	}

	return ret
}

func SortFieldFromProto(f *proto.SortField, baseSchema types.Type, ext ExtensionLookup, c *extensions.Collection) (sf SortField, err error) {
	sf.Expr, err = ExprFromProto(f.Expr, baseSchema, ext, c)
	if err != nil {
		return
	}

	switch k := f.SortKind.(type) {
	case *proto.SortField_Direction:
		sf.Kind = types.SortDirection(k.Direction)
	case *proto.SortField_ComparisonFunctionReference:
		sf.Kind = types.FunctionRef(k.ComparisonFunctionReference)
	default:
		err = substraitgo.ErrNotImplemented
	}
	return
}

func (fb PrecedingBound) ToProto() *proto.Expression_WindowFunction_Bound {
	return &proto.Expression_WindowFunction_Bound{
		Kind: &proto.Expression_WindowFunction_Bound_Preceding_{
			Preceding: &proto.Expression_WindowFunction_Bound_Preceding{Offset: int64(fb)},
		},
	}
}

func (fb FollowingBound) ToProto() *proto.Expression_WindowFunction_Bound {
	return &proto.Expression_WindowFunction_Bound{
		Kind: &proto.Expression_WindowFunction_Bound_Following_{
			Following: &proto.Expression_WindowFunction_Bound_Following{Offset: int64(fb)},
		},
	}
}

func (CurrentRow) ToProto() *proto.Expression_WindowFunction_Bound {
	return &proto.Expression_WindowFunction_Bound{
		Kind: &proto.Expression_WindowFunction_Bound_CurrentRow_{
			CurrentRow: &proto.Expression_WindowFunction_Bound_CurrentRow{},
		},
	}
}

func (Unbounded) ToProto() *proto.Expression_WindowFunction_Bound {
	return &proto.Expression_WindowFunction_Bound{
		Kind: &proto.Expression_WindowFunction_Bound_Unbounded_{
			Unbounded: &proto.Expression_WindowFunction_Bound_Unbounded{},
		}}
}

func BoundFromProto(b *proto.Expression_WindowFunction_Bound) Bound {
	switch t := b.Kind.(type) {
	case *proto.Expression_WindowFunction_Bound_Preceding_:
		return PrecedingBound(t.Preceding.Offset)
	case *proto.Expression_WindowFunction_Bound_CurrentRow_:
		return CurrentRow{}
	case *proto.Expression_WindowFunction_Bound_Following_:
		return FollowingBound(t.Following.Offset)
	case *proto.Expression_WindowFunction_Bound_Unbounded_:
		return Unbounded{}
	}

	// bound is optional
	return nil
}

type ScalarFunction struct {
	FuncRef     uint32
	ID          extensions.ID
	Declaration *extensions.ScalarFunctionVariant

	Args       []types.FuncArg
	Options    []*types.FunctionOption
	OutputType types.Type
}

func (s *ScalarFunction) IsScalar() bool {
	for _, arg := range s.Args {
		if ex, ok := arg.(Expression); ok {
			if !ex.IsScalar() {
				return false
			}
		}
	}
	return true
}

func (*ScalarFunction) isRootRef() {}

func (s *ScalarFunction) String() string {
	var b strings.Builder

	b.WriteString(s.ID.Name)
	b.WriteByte('(')

	for i, arg := range s.Args {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(arg.String())
	}

	b.WriteString(") => ")
	b.WriteString(s.OutputType.String())

	return b.String()
}

func (s *ScalarFunction) GetOption(name string) []string {
	for _, o := range s.Options {
		if name == o.Name {
			return o.GetPreference()
		}
	}
	return nil
}

func (s *ScalarFunction) GetType() types.Type { return s.OutputType }
func (s *ScalarFunction) ToProtoFuncArg() *proto.FunctionArgument {
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Value{
			Value: s.ToProto(),
		},
	}
}

func (s *ScalarFunction) ToProto() *proto.Expression {
	args := make([]*proto.FunctionArgument, len(s.Args))
	for i, a := range s.Args {
		args[i] = a.ToProtoFuncArg()
	}

	return &proto.Expression{
		RexType: &proto.Expression_ScalarFunction_{
			ScalarFunction: &proto.Expression_ScalarFunction{
				FunctionReference: s.FuncRef,
				Options:           s.Options,
				OutputType:        types.TypeToProto(s.OutputType),
				Arguments:         args,
			},
		},
	}
}

func (s *ScalarFunction) Equals(rhs Expression) bool {
	other, ok := rhs.(*ScalarFunction)
	if !ok {
		return false
	}

	switch {
	case s.FuncRef != other.FuncRef:
		return false
	case !s.OutputType.Equals(other.OutputType):
		return false
	}

	for i := range s.Args {
		if !FuncArgsEqual(s.Args[i], other.Args[i]) {
			return false
		}
	}

	return slices.EqualFunc(s.Options, other.Options, func(a, b *types.FunctionOption) bool {
		return pb.Equal(a, b)
	})
}

func (s *ScalarFunction) Visit(visit VisitFunc) Expression {
	var args []types.FuncArg
	for i, arg := range s.Args {
		var after types.FuncArg
		switch t := arg.(type) {
		case Expression:
			after = visit(t)
		default:
			after = arg
		}

		if args == nil && arg != after {
			args = make([]types.FuncArg, len(s.Args))
			for j := 0; j < i; j++ {
				args[j] = s.Args[i]
			}
		}

		if args != nil {
			args[i] = after
		}
	}

	if args == nil {
		return s
	}

	out := *s
	out.Args = args
	return &out
}

type WindowFunction struct {
	FuncRef     uint32
	ID          extensions.ID
	Declaration *extensions.WindowFunctionVariant

	Args       []types.FuncArg
	Options    []*types.FunctionOption
	OutputType types.Type

	Phase      types.AggregationPhase
	Sorts      []SortField
	Invocation types.AggregationInvocation
	Partitions []Expression

	LowerBound, UpperBound Bound
}

func (*WindowFunction) IsScalar() bool { return false }

func (*WindowFunction) isRootRef() {}

func (w *WindowFunction) String() string {
	var b strings.Builder

	b.WriteString(w.ID.Name)
	b.WriteByte('(')

	for i, arg := range w.Args {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(arg.String())
	}

	if len(w.Sorts) > 0 {
		b.WriteString("; sort: [")
		for i, s := range w.Sorts {
			if i != 0 {
				b.WriteString(",")
			}
			fmt.Fprintf(&b, "{expr: %s, %s}", s.Expr, s.Kind.String())
		}
		b.WriteString("]")
	}

	if len(w.Options) > 0 {
		b.WriteString("; [options: {")
		for i, opt := range w.Options {
			if i != 0 {
				b.WriteString(", ")
			}
			fmt.Fprintf(&b, "%s => %v", opt.Name, opt.Preference)
		}
		b.WriteString("}]")
	}

	if len(w.Partitions) > 0 {
		b.WriteString("; partitions: [")
		for i, part := range w.Partitions {
			if i != 0 {
				b.WriteString(", ")
			}
			b.WriteString(part.String())
		}
		b.WriteString("]")
	}

	fmt.Fprintf(&b, "; phase: %s, invocation: %s) => %s",
		w.Phase, w.Invocation, w.OutputType)

	return b.String()
}

func (w *WindowFunction) GetType() types.Type { return w.OutputType }
func (w *WindowFunction) Equals(other Expression) bool {
	rhs, ok := other.(*WindowFunction)
	if !ok {
		return false
	}

	switch {
	case w.FuncRef != rhs.FuncRef:
		return false
	case !w.OutputType.Equals(rhs.OutputType):
		return false
	case w.Phase != rhs.Phase || w.Invocation != rhs.Invocation:
		return false
	case w.LowerBound != rhs.LowerBound || w.UpperBound != rhs.UpperBound:
		return false
	case !slices.EqualFunc(w.Options, rhs.Options, func(l, r *types.FunctionOption) bool {
		return l.Name == r.Name && slices.Equal(l.Preference, r.Preference)
	}):
		return false
	case !slices.EqualFunc(w.Partitions, rhs.Partitions, exprEqual):
		return false
	case !slices.EqualFunc(w.Sorts, rhs.Sorts, func(l, r SortField) bool {
		return l.Expr.Equals(r.Expr) && l.Kind == r.Kind
	}):
		return false
	case !slices.EqualFunc(w.Args, rhs.Args, FuncArgsEqual):
		return false
	}

	return true
}

func (w *WindowFunction) ToProto() *proto.Expression {
	var (
		args       []*proto.FunctionArgument
		sorts      []*proto.SortField
		parts      []*proto.Expression
		upperBound *proto.Expression_WindowFunction_Bound
		lowerBound *proto.Expression_WindowFunction_Bound
	)

	if len(w.Args) > 0 {
		args = make([]*proto.FunctionArgument, len(w.Args))
		for i, a := range w.Args {
			args[i] = a.ToProtoFuncArg()
		}
	}

	if len(w.Sorts) > 0 {
		sorts = make([]*proto.SortField, len(w.Sorts))
		for i, s := range w.Sorts {
			sorts[i] = s.ToProto()
		}
	}

	if len(w.Partitions) > 0 {
		parts = make([]*proto.Expression, len(w.Partitions))
		for i, p := range w.Partitions {
			parts[i] = p.ToProto()
		}
	}

	if w.UpperBound != nil {
		upperBound = w.UpperBound.ToProto()
	}

	if w.LowerBound != nil {
		lowerBound = w.LowerBound.ToProto()
	}

	return &proto.Expression{
		RexType: &proto.Expression_WindowFunction_{
			WindowFunction: &proto.Expression_WindowFunction{
				FunctionReference: w.FuncRef,
				Arguments:         args,
				Options:           w.Options,
				OutputType:        types.TypeToProto(w.OutputType),
				Phase:             w.Phase,
				Sorts:             sorts,
				Invocation:        w.Invocation,
				Partitions:        parts,
				LowerBound:        lowerBound,
				UpperBound:        upperBound,
			},
		},
	}
}

func (w *WindowFunction) ToProtoFuncArg() *proto.FunctionArgument {
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Value{
			Value: w.ToProto(),
		},
	}
}

func (w *WindowFunction) Visit(visit VisitFunc) Expression {
	var args []types.FuncArg
	for i, arg := range w.Args {
		var after types.FuncArg
		switch t := arg.(type) {
		case Expression:
			after = visit(t)
		default:
			after = arg
		}

		if args == nil && arg != after {
			args = make([]types.FuncArg, len(w.Args))
			for j := 0; j < i; j++ {
				args[j] = w.Args[i]
			}
		}

		if args != nil {
			args[i] = after
		}
	}

	if args == nil {
		return w
	}

	out := *w
	out.Args = args
	return &out
}

type AggregateFunction struct {
	FuncRef     uint32
	ID          extensions.ID
	Declaration *extensions.AggregateFunctionVariant

	Args       []types.FuncArg
	Options    []*types.FunctionOption
	OutputType types.Type
	Phase      types.AggregationPhase
	Invocation types.AggregationInvocation
	Sorts      []SortField
}

func NewAggregateFunctionFromProto(agg *proto.AggregateFunction, baseSchema types.Type, ext extensions.Set, c *extensions.Collection) (*AggregateFunction, error) {
	if agg.OutputType == nil {
		return nil, fmt.Errorf("%w: missing output type", substraitgo.ErrInvalidExpr)
	}

	var err error
	args := make([]types.FuncArg, len(agg.Arguments))
	for i, a := range agg.Arguments {
		if args[i], err = FuncArgFromProto(a, baseSchema, ext, c); err != nil {
			return nil, err
		}
	}

	sorts := make([]SortField, len(agg.Sorts))
	for i, s := range agg.Sorts {
		if sorts[i], err = SortFieldFromProto(s, baseSchema, ext, c); err != nil {
			return nil, err
		}
	}

	var (
		id   extensions.ID
		decl *extensions.AggregateFunctionVariant
		ok   bool
	)

	if ext != nil {
		if id, ok = ext.DecodeFunc(agg.FunctionReference); !ok {
			return nil, substraitgo.ErrNotFound
		}
	}

	if c != nil {
		if decl, ok = ext.LookupAggregateFunction(agg.FunctionReference, c); !ok {
			return nil, substraitgo.ErrNotFound
		}
	}

	return &AggregateFunction{
		FuncRef:     agg.FunctionReference,
		ID:          id,
		Declaration: decl,
		Args:        args,
		Options:     agg.Options,
		OutputType:  types.TypeFromProto(agg.OutputType),
		Phase:       agg.Phase,
		Invocation:  agg.Invocation,
		Sorts:       sorts,
	}, nil
}

func (a *AggregateFunction) String() string {
	var b strings.Builder

	b.WriteString(a.ID.Name)
	b.WriteByte('(')

	for i, arg := range a.Args {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(arg.String())
	}

	b.WriteString(") => ")
	b.WriteString(a.OutputType.String())

	return b.String()
}

func (a *AggregateFunction) GetOption(name string) []string {
	for _, o := range a.Options {
		if name == o.Name {
			return o.GetPreference()
		}
	}
	return nil
}

func (a *AggregateFunction) GetType() types.Type { return a.OutputType }

func (a *AggregateFunction) ToProto() *proto.AggregateFunction {
	var (
		args  []*proto.FunctionArgument
		sorts []*proto.SortField
	)
	if len(a.Args) > 0 {
		args = make([]*proto.FunctionArgument, len(a.Args))
		for i, arg := range a.Args {
			args[i] = arg.ToProtoFuncArg()
		}
	}

	if len(a.Sorts) > 0 {
		sorts = make([]*proto.SortField, len(a.Sorts))
		for i, s := range a.Sorts {
			sorts[i] = s.ToProto()
		}
	}

	return &proto.AggregateFunction{
		FunctionReference: a.FuncRef,
		Arguments:         args,
		Options:           a.Options,
		OutputType:        types.TypeToProto(a.OutputType),
		Phase:             a.Phase,
		Sorts:             sorts,
		Invocation:        a.Invocation,
	}
}
