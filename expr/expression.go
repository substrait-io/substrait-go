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
)

type ExtensionLookup interface {
	DecodeType(uint32) (extensions.ID, bool)
	DecodeFunc(uint32) (extensions.ID, bool)
	DecodeTypeVariation(uint32) (extensions.ID, bool)

	LookupScalarFunction(uint32, *extensions.Collection) (*extensions.ScalarFunctionVariant, bool)
	LookupAggregateFunction(uint32, *extensions.Collection) (*extensions.AggregateFunctionVariant, bool)
	LookupWindowFunction(uint32, *extensions.Collection) (*extensions.WindowFunctionVariant, bool)

	GetTypeAnchor(id extensions.ID) uint32
	GetFuncAnchor(id extensions.ID) uint32
	GetTypeVariationAnchor(id extensions.ID) uint32
}

func FuncArgFromProto(e *proto.FunctionArgument, baseSchema types.Type, ext ExtensionLookup, c *extensions.Collection) (types.FuncArg, error) {
	switch et := e.ArgType.(type) {
	case *proto.FunctionArgument_Enum:
		return types.Enum(et.Enum), nil
	case *proto.FunctionArgument_Type:
		return types.TypeFromProto(et.Type), nil
	case *proto.FunctionArgument_Value:
		return ExprFromProto(et.Value, baseSchema, ext, c)
	}
	return nil, substraitgo.ErrNotImplemented
}

func ExprFromProto(e *proto.Expression, baseSchema types.Type, ext ExtensionLookup, c *extensions.Collection) (Expression, error) {
	if e == nil {
		return nil, fmt.Errorf("%w: protobuf Expression is nil", substraitgo.ErrInvalidExpr)
	}

	switch et := e.RexType.(type) {
	case *proto.Expression_Literal_:
		return LiteralFromProto(et.Literal), nil
	case *proto.Expression_Selection:
		return FieldReferenceFromProto(et.Selection, baseSchema, ext, c)
	case *proto.Expression_ScalarFunction_:
		var err error
		args := make([]types.FuncArg, len(et.ScalarFunction.Arguments))
		for i, a := range et.ScalarFunction.Arguments {
			if args[i], err = FuncArgFromProto(a, baseSchema, ext, c); err != nil {
				return nil, err
			}
		}

		var (
			id   extensions.ID
			decl *extensions.ScalarFunctionVariant
			ok   bool
		)

		if ext != nil {
			if id, ok = ext.DecodeFunc(et.ScalarFunction.FunctionReference); !ok {
				return nil, substraitgo.ErrNotFound
			}
		}

		if c != nil {
			if decl, ok = ext.LookupScalarFunction(et.ScalarFunction.FunctionReference, c); !ok {
				return nil, substraitgo.ErrNotFound
			}
		}

		return &ScalarFunction{
			FuncRef:     et.ScalarFunction.FunctionReference,
			Declaration: decl,
			ID:          id,
			Args:        args,
			Options:     et.ScalarFunction.Options,
			OutputType:  types.TypeFromProto(et.ScalarFunction.OutputType),
		}, nil
	case *proto.Expression_WindowFunction_:
		var err error
		args := make([]types.FuncArg, len(et.WindowFunction.Arguments))
		for i, a := range et.WindowFunction.Arguments {
			if args[i], err = FuncArgFromProto(a, baseSchema, ext, c); err != nil {
				return nil, err
			}
		}

		parts := make([]Expression, len(et.WindowFunction.Partitions))
		for i, p := range et.WindowFunction.Partitions {
			if parts[i], err = ExprFromProto(p, baseSchema, ext, c); err != nil {
				return nil, err
			}
		}

		sorts := make([]SortField, len(et.WindowFunction.Sorts))
		for i, s := range et.WindowFunction.Sorts {
			if sorts[i], err = SortFieldFromProto(s, baseSchema, ext, c); err != nil {
				return nil, err
			}
		}

		var (
			id   extensions.ID
			decl *extensions.WindowFunctionVariant
			ok   bool
		)

		if ext != nil {
			if id, ok = ext.DecodeFunc(et.WindowFunction.FunctionReference); !ok {
				return nil, substraitgo.ErrNotFound
			}
		}

		if c != nil {
			if decl, ok = ext.LookupWindowFunction(et.WindowFunction.FunctionReference, c); !ok {
				return nil, substraitgo.ErrNotFound
			}
		}

		return &WindowFunction{
			FuncRef:     et.WindowFunction.FunctionReference,
			ID:          id,
			Declaration: decl,
			Args:        args,
			Options:     et.WindowFunction.Options,
			OutputType:  types.TypeFromProto(et.WindowFunction.OutputType),
			Phase:       et.WindowFunction.Phase,
			Invocation:  et.WindowFunction.Invocation,
			Partitions:  parts,
			Sorts:       sorts,
			LowerBound:  BoundFromProto(et.WindowFunction.LowerBound),
			UpperBound:  BoundFromProto(et.WindowFunction.UpperBound),
		}, nil
	case *proto.Expression_IfThen_:
		elseExpr, err := ExprFromProto(et.IfThen.Else, baseSchema, ext, c)
		if err != nil {
			return nil, err
		}

		ifs := make([]struct{ If, Then Expression }, len(et.IfThen.Ifs))
		for i, clause := range et.IfThen.Ifs {
			ifs[i].If, err = ExprFromProto(clause.If, baseSchema, ext, c)
			if err != nil {
				return nil, err
			}

			ifs[i].Then, err = ExprFromProto(clause.Then, baseSchema, ext, c)
			if err != nil {
				return nil, err
			}
		}

		return &IfThen{
			IFs:  ifs,
			Else: elseExpr,
		}, nil
	case *proto.Expression_SwitchExpression_:
		matched, err := ExprFromProto(et.SwitchExpression.Match, baseSchema, ext, c)
		if err != nil {
			return nil, err
		}

		elseExpr, err := ExprFromProto(et.SwitchExpression.Else, baseSchema, ext, c)
		if err != nil {
			return nil, err
		}

		ifs := make([]struct {
			If   Literal
			Then Expression
		}, len(et.SwitchExpression.Ifs))
		for i, clause := range et.SwitchExpression.Ifs {
			ifs[i].If = LiteralFromProto(clause.If)
			ifs[i].Then, err = ExprFromProto(clause.Then, baseSchema, ext, c)
			if err != nil {
				return nil, err
			}
		}

		return &SwitchExpr{
			Match: matched,
			IFs:   ifs,
			Else:  elseExpr,
		}, nil
	case *proto.Expression_SingularOrList_:
		val, err := ExprFromProto(et.SingularOrList.Value, baseSchema, ext, c)
		if err != nil {
			return nil, err
		}

		opts := make([]Expression, len(et.SingularOrList.Options))
		for i, o := range et.SingularOrList.Options {
			opts[i], err = ExprFromProto(o, baseSchema, ext, c)
			if err != nil {
				return nil, err
			}
		}

		return &SingularOrList{
			Value:   val,
			Options: opts,
		}, nil
	case *proto.Expression_MultiOrList_:
		var err error
		val := make([]Expression, len(et.MultiOrList.Value))
		for i, v := range et.MultiOrList.Value {
			val[i], err = ExprFromProto(v, baseSchema, ext, c)
			if err != nil {
				return nil, err
			}
		}

		options := make([][]Expression, len(et.MultiOrList.Options))
		for i, opts := range et.MultiOrList.Options {
			options[i] = make([]Expression, len(opts.Fields))
			for j, o := range opts.Fields {
				options[i][j], err = ExprFromProto(o, baseSchema, ext, c)
				if err != nil {
					return nil, err
				}
			}
		}

		return &MultiOrList{
			Value:   val,
			Options: options,
		}, nil
	case *proto.Expression_Cast_:
		input, err := ExprFromProto(et.Cast.Input, baseSchema, ext, c)
		if err != nil {
			return nil, err
		}

		return &Cast{
			Type:            types.TypeFromProto(et.Cast.Type),
			Input:           input,
			FailureBehavior: et.Cast.FailureBehavior,
		}, nil
	case *proto.Expression_Nested_:
		var err error
		nullable, typevar := et.Nested.Nullable, et.Nested.TypeVariationReference

		switch n := et.Nested.NestedType.(type) {
		case *proto.Expression_Nested_Map_:
			if len(n.Map.KeyValues) == 0 {
				return nil, fmt.Errorf("%w: use an empty map literal instead of NestedExpr map to preserve type info",
					substraitgo.ErrInvalidExpr)
			}

			keyValues := make([]struct{ Key, Value Expression }, len(n.Map.KeyValues))
			for i, kv := range n.Map.KeyValues {
				keyValues[i].Key, err = ExprFromProto(kv.Key, baseSchema, ext, c)
				if err != nil {
					return nil, err
				}

				keyValues[i].Value, err = ExprFromProto(kv.Value, baseSchema, ext, c)
				if err != nil {
					return nil, err
				}
			}

			return &MapExpr{
				Nullable:         nullable,
				TypeVariationRef: typevar,
				KeyValues:        keyValues,
			}, nil
		case *proto.Expression_Nested_Struct_:
			fields := make([]Expression, len(n.Struct.Fields))
			for i, f := range n.Struct.Fields {
				fields[i], err = ExprFromProto(f, baseSchema, ext, c)
				if err != nil {
					return nil, err
				}
			}

			return &StructExpr{
				Nullable:         nullable,
				TypeVariationRef: typevar,
				Fields:           fields,
			}, nil
		case *proto.Expression_Nested_List_:
			if len(n.List.Values) == 0 {
				return nil, fmt.Errorf("%w: use an empty list literal to preserve type info instead of nested expression",
					substraitgo.ErrInvalidExpr)
			}

			values := make([]Expression, len(n.List.Values))
			for i, v := range n.List.Values {
				values[i], err = ExprFromProto(v, baseSchema, ext, c)
				if err != nil {
					return nil, err
				}
			}

			return &ListExpr{
				Nullable:         nullable,
				TypeVariationRef: typevar,
				Values:           values,
			}, nil
		default:
			return nil, fmt.Errorf("%w: nested expression: %s",
				substraitgo.ErrInvalidExpr, n)
		}
	case *proto.Expression_Enum_:
		return nil, fmt.Errorf("%w: deprecated", substraitgo.ErrNotImplemented)
	case *proto.Expression_Subquery_:
	}

	return nil, fmt.Errorf("%w: ExprFromProto: %s", substraitgo.ErrNotImplemented, e)
}

type VisitFunc func(Expression) Expression

// Expression can be one of many different things as a generalized
// expression. It could be:
//
//   - A literal
//   - A Field Reference Selection
//   - A Scalar Function expression
//   - A Window Function expression
//   - An If-Then statement
//   - A Switch Expression
//   - A Singular Or List
//   - A Multiple Or List
//   - A Cast expression
//   - A Subquery
//   - A Nested expression
type Expression interface {
	// an Expression can also be a function argument
	types.FuncArg
	// an expression can also be the root of a reference
	RootRefType

	// IsBound means that this expression has a defined output
	// type and, in the case of Scalar/Window functions, has
	// properly set a Function Anchor and a Declaration defining
	// the function variant being called as registered in an
	// extension collection.
	//
	// To bind an expression, use BindExpression which will
	// recursively resolve the output types for all references and
	// functions that are not yet bound in it.
	IsBound() bool

	IsScalar() bool
	// GetType returns the output type of this expression
	GetType() types.Type
	// ToProto converts this Expression and its arguments
	// to the equivalent Protobuf objects.
	ToProto() *proto.Expression
	// Equals returns true if this expression and all of its
	// arguments and their children etc. are equal to the passed
	// in Expression.
	Equals(Expression) bool
	// Visit invokes the passed visit function for each child of the
	// expression. The visit function can return its input expression
	// as-is with no changes, or it can construct and return a
	// replacement expression. If any children have been replaced, Visit
	// will construct and return a new instance of this expression using
	// the new children. Callers can use the Visit method to traverse
	// and potentially rewrite the expression tree, in either pre or post
	// order. Here is a pre-order example:
	//
	//   func preOrderVisit(e Expression) Expression {
	//     // Replace some scalar function, leave everything else
	//     // as-is. This check is before the call to Visit, so
	//     // it's a pre-order traversal
	//     if f, ok := e.(*ScalarFunction); ok {
	//       return &ScalarFunction{
	//         ID: ExtID{URI: "some other uri", Name: "some other func"},
	//         Args: f.Args,
	//         Options: f.Options,
	//         OutputType: f.OutputType,
	//       }
	//     }
	//     return e.Visit(preOrderVisit)
	//   }
	//   newExpr := preOrderVisit(oldExpr)
	Visit(VisitFunc) Expression
}

type IfThen struct {
	IFs []struct {
		If   Expression
		Then Expression
	}
	Else Expression
}

func (ex *IfThen) String() string {
	var b strings.Builder
	b.WriteString("<IfThen>(")
	for i, clause := range ex.IFs {
		if i != 0 {
			b.WriteString(": ")
		}
		b.WriteString("(" + clause.If.String() + ") ?")
		b.WriteString(clause.Then.String())
	}
	b.WriteString(")<Else>(")
	if ex.Else != nil {
		b.WriteString(ex.Else.String())
	}
	b.WriteString(")")
	return b.String()
}

func (ex *IfThen) ToProtoFuncArg() *proto.FunctionArgument {
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Value{
			Value: ex.ToProto(),
		},
	}
}

func (ex *IfThen) isRootRef() {}

func (ex *IfThen) IsBound() bool {
	for _, clause := range ex.IFs {
		if clause.If != nil && !clause.If.IsBound() {
			return false
		}
		if clause.Then != nil && !clause.Then.IsBound() {
			return false
		}
	}

	return ex.Else == nil || ex.Else.IsBound()
}

func (ex *IfThen) IsScalar() bool {
	for _, clause := range ex.IFs {
		if clause.If != nil && !clause.If.IsScalar() {
			return false
		}
		if clause.Then != nil && !clause.Then.IsScalar() {
			return false
		}
	}
	return ex.Else == nil || ex.Else.IsScalar()
}

func (ex *IfThen) GetType() types.Type {
	return ex.Else.GetType()
}

func (ex *IfThen) ToProto() *proto.Expression {
	ifthenClauses := make([]*proto.Expression_IfThen_IfClause, len(ex.IFs))
	for i, c := range ex.IFs {
		ifthenClauses[i] = &proto.Expression_IfThen_IfClause{
			If:   c.If.ToProto(),
			Then: c.Then.ToProto(),
		}
	}

	var elseClause *proto.Expression
	if ex.Else != nil {
		elseClause = ex.Else.ToProto()
	}
	return &proto.Expression{
		RexType: &proto.Expression_IfThen_{
			IfThen: &proto.Expression_IfThen{
				Ifs:  ifthenClauses,
				Else: elseClause,
			},
		},
	}
}

func (ex *IfThen) Equals(other Expression) bool {
	rhs, ok := other.(*IfThen)
	if !ok {
		return false
	}

	switch {
	case ex.Else != nil && !ex.Else.Equals(rhs.Else):
		return false
	case ex.Else != nil && rhs.Else == nil:
		return false
	}

	return slices.EqualFunc(ex.IFs, rhs.IFs,
		func(l, r struct {
			If   Expression
			Then Expression
		}) bool {
			return l.If.Equals(r.If) && l.Then.Equals(r.Then)
		})
}

func (ex *IfThen) Visit(visit VisitFunc) Expression {
	var out *IfThen

	for i, clause := range ex.IFs {
		afterIf := visit(clause.If)
		afterThen := visit(clause.Then)

		if out == nil && (afterIf != clause.If || afterThen != clause.Then) {
			out = &IfThen{IFs: slices.Clone(ex.IFs)}
		}

		if out != nil {
			out.IFs[i].If = afterIf
			out.IFs[i].Then = afterThen
		}
	}

	afterElse := visit(ex.Else)
	if out == nil {
		if afterElse == ex.Else {
			return ex
		}

		out = &IfThen{IFs: slices.Clone(ex.IFs)}
	}

	out.Else = afterElse
	return out
}

type Cast struct {
	Type            types.Type
	Input           Expression
	FailureBehavior types.CastFailBehavior
}

func (ex *Cast) String() string {
	return fmt.Sprintf("cast(%s AS %s, fail: %s)",
		ex.Input, ex.Type, ex.FailureBehavior)
}

func (ex *Cast) ToProtoFuncArg() *proto.FunctionArgument {
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Value{Value: ex.ToProto()},
	}
}

func (ex *Cast) isRootRef() {}

func (ex *Cast) IsScalar() bool {
	return ex.Input.IsScalar()
}

func (ex *Cast) IsBound() bool {
	return ex.Input.IsBound()
}

func (ex *Cast) GetType() types.Type {
	return ex.Type
}

func (ex *Cast) ToProto() *proto.Expression {
	return &proto.Expression{
		RexType: &proto.Expression_Cast_{
			Cast: &proto.Expression_Cast{
				Type:            types.TypeToProto(ex.Type),
				Input:           ex.Input.ToProto(),
				FailureBehavior: ex.FailureBehavior,
			},
		},
	}
}

func (ex *Cast) Equals(other Expression) bool {
	rhs, ok := other.(*Cast)
	if !ok {
		return false
	}

	return ex.Type.Equals(rhs.Type) && ex.Input.Equals(rhs.Input) &&
		ex.FailureBehavior == rhs.FailureBehavior
}

func (ex *Cast) Visit(visit VisitFunc) Expression {
	return visit(ex.Input)
}

type SwitchExpr struct {
	Match Expression
	IFs   []struct {
		If   Literal
		Then Expression
	}
	Else Expression
}

func (ex *SwitchExpr) String() string {
	var b strings.Builder
	b.WriteString("CASE ")
	b.WriteString(ex.Match.String())
	b.WriteString(":")
	for _, c := range ex.IFs {
		b.WriteString("\nWHEN ")
		b.WriteString(c.If.String())
		b.WriteString(" THEN ")
		b.WriteString(c.Then.String())
		b.WriteByte(';')
	}
	b.WriteString("Else ")
	b.WriteString(ex.Else.String())
	return b.String()
}

func (ex *SwitchExpr) ToProtoFuncArg() *proto.FunctionArgument {
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Value{
			Value: ex.ToProto(),
		},
	}
}

func (ex *SwitchExpr) isRootRef() {}

func (ex *SwitchExpr) IsBound() bool {
	if !ex.Match.IsBound() {
		return false
	}

	for _, c := range ex.IFs {
		if !c.Then.IsBound() {
			return false
		}
	}

	if ex.Else != nil {
		return ex.Else.IsBound()
	}

	return true
}

func (ex *SwitchExpr) IsScalar() bool {
	if !ex.Match.IsScalar() {
		return false
	}

	for _, c := range ex.IFs {
		if !c.Then.IsScalar() {
			return false
		}
	}

	if ex.Else != nil {
		return ex.Else.IsScalar()
	}

	return true
}

func (ex *SwitchExpr) GetType() types.Type {
	var t types.Type
	for _, cond := range ex.IFs {
		t = cond.Then.GetType()
		// check if any of the branches return a nullable type
		// only return a non-nullable if *all* branches return
		// non-nullable
		if t.GetNullability() == types.NullabilityNullable {
			break
		}
	}

	// if there's no else clause, then return a nullable type
	if ex.Else == nil {
		return t.WithNullability(types.NullabilityNullable)
	}

	// if any branch returns nullable, then return the nullable type
	// we don't care if the else clause is nullable
	if t.GetNullability() == types.NullabilityNullable {
		return t
	}

	// if no branch returns nullable, just return the else clause type
	return ex.Else.GetType()
}

func (ex *SwitchExpr) ToProto() *proto.Expression {
	var elseExpr *proto.Expression
	if ex.Else != nil {
		elseExpr = ex.Else.ToProto()
	}

	cases := make([]*proto.Expression_SwitchExpression_IfValue, len(ex.IFs))
	for i, c := range ex.IFs {
		cases[i] = &proto.Expression_SwitchExpression_IfValue{
			If:   c.If.ToProtoLiteral(),
			Then: c.Then.ToProto(),
		}
	}

	return &proto.Expression{
		RexType: &proto.Expression_SwitchExpression_{
			SwitchExpression: &proto.Expression_SwitchExpression{
				Match: ex.Match.ToProto(),
				Ifs:   cases,
				Else:  elseExpr,
			},
		},
	}
}

func (ex *SwitchExpr) Equals(other Expression) bool {
	rhs, ok := other.(*SwitchExpr)
	if !ok {
		return false
	}

	switch {
	case !ex.Match.Equals(rhs.Match):
		return false
	case ex.Else != nil && !ex.Else.Equals(rhs.Else):
		return false
	case ex.Else == nil && rhs.Else != nil:
		return false
	}

	return slices.EqualFunc(ex.IFs, rhs.IFs,
		func(l, r struct {
			If   Literal
			Then Expression
		}) bool {
			return l.If.Equals(r.If) && l.Then.Equals(r.Then)
		})
}

func (ex *SwitchExpr) Visit(visit VisitFunc) Expression {
	var out *SwitchExpr
	if after := visit(ex.Match); after != ex.Match {
		out = &SwitchExpr{
			Match: after,
			IFs: make([]struct {
				If   Literal
				Then Expression
			}, len(ex.IFs)),
			Else: ex.Else,
		}
	}

	for i, c := range ex.IFs {
		afterIf := visit(c.If)
		afterThen := visit(c.Then)

		if out == nil && (afterIf != c.If || afterThen != c.Then) {
			out = &SwitchExpr{
				Match: ex.Match,
				IFs: make([]struct {
					If   Literal
					Then Expression
				}, len(ex.IFs)),
				Else: ex.Else,
			}
			copy(out.IFs, ex.IFs[:i])
		}

		if out != nil {
			out.IFs[i].If = afterIf.(Literal)
			out.IFs[i].Then = afterThen
		}
	}

	afterElse := visit(ex.Else)
	if out == nil {
		if afterElse == ex.Else {
			return ex
		}

		out = &SwitchExpr{
			Match: ex.Match,
			IFs:   slices.Clone(ex.IFs),
		}
	}

	out.Else = afterElse
	return out
}

func exprEqual(l, r Expression) bool { return l.Equals(r) }

type SingularOrList struct {
	Value   Expression
	Options []Expression
}

func (ex *SingularOrList) String() string {
	var b strings.Builder
	b.WriteString(ex.Value.String())
	b.WriteString(" IN [")
	for i, o := range ex.Options {
		if i != 0 {
			b.WriteString(",")
		}
		b.WriteString(o.String())
	}
	b.WriteString("]")
	return b.String()
}

func (ex *SingularOrList) ToProtoFuncArg() *proto.FunctionArgument {
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Value{
			Value: ex.ToProto(),
		},
	}
}

func (ex *SingularOrList) isRootRef() {}

func (ex *SingularOrList) IsBound() bool {
	if !ex.Value.IsBound() {
		return false
	}

	for _, o := range ex.Options {
		if !o.IsBound() {
			return false
		}
	}

	return true
}

func (ex *SingularOrList) IsScalar() bool {
	if !ex.Value.IsScalar() {
		return false
	}

	for _, o := range ex.Options {
		if !o.IsScalar() {
			return false
		}
	}

	return true
}

func (ex *SingularOrList) GetType() types.Type {
	return &types.BooleanType{Nullability: types.NullabilityRequired}
}

func (ex *SingularOrList) ToProto() *proto.Expression {
	opts := make([]*proto.Expression, len(ex.Options))
	for i, o := range ex.Options {
		opts[i] = o.ToProto()
	}
	return &proto.Expression{
		RexType: &proto.Expression_SingularOrList_{
			SingularOrList: &proto.Expression_SingularOrList{
				Value:   ex.Value.ToProto(),
				Options: opts,
			},
		},
	}
}

func (ex *SingularOrList) Equals(other Expression) bool {
	rhs, ok := other.(*SingularOrList)
	if !ok {
		return false
	}

	if !ex.Value.Equals(rhs.Value) {
		return false
	}

	return slices.EqualFunc(ex.Options, rhs.Options, exprEqual)
}

func (ex *SingularOrList) Visit(visit VisitFunc) Expression {
	var out *SingularOrList
	if temp := visit(ex.Value); temp != ex.Value {
		out = &SingularOrList{
			Value:   temp,
			Options: make([]Expression, len(ex.Options)),
		}
	}

	for i, o := range ex.Options {
		temp := visit(o)
		if out == nil && temp != o {
			out = &SingularOrList{Value: ex.Value, Options: make([]Expression, len(ex.Options))}
			copy(out.Options, ex.Options[:i])
		}

		if out != nil {
			out.Options[i] = temp
		}
	}

	if out == nil {
		return ex
	}

	return out
}

type MultiOrList struct {
	Value   []Expression
	Options [][]Expression
}

func (ex *MultiOrList) String() string {
	var b strings.Builder
	writeList := func(list []Expression) {
		b.WriteByte('[')
		for i, v := range list {
			if i != 0 {
				b.WriteString(", ")
			}
			b.WriteString(v.String())
		}
		b.WriteByte(']')
	}

	writeList(ex.Value)
	b.WriteString(" IN [")
	for i, opts := range ex.Options {
		if i != 0 {
			b.WriteString(", ")
		}
		writeList(opts)
	}
	b.WriteString("]")

	return b.String()
}

func (ex *MultiOrList) ToProtoFuncArg() *proto.FunctionArgument {
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Value{
			Value: ex.ToProto(),
		},
	}
}

func (ex *MultiOrList) isRootRef() {}

func (ex *MultiOrList) IsBound() bool {
	for _, v := range ex.Value {
		if !v.IsBound() {
			return false
		}
	}

	for _, opts := range ex.Options {
		for _, o := range opts {
			if !o.IsBound() {
				return false
			}
		}
	}

	return true
}

func (ex *MultiOrList) IsScalar() bool {
	for _, v := range ex.Value {
		if !v.IsScalar() {
			return false
		}
	}

	for _, opts := range ex.Options {
		for _, o := range opts {
			if !o.IsScalar() {
				return false
			}
		}
	}

	return true
}

func (ex *MultiOrList) GetType() types.Type {
	return &types.BooleanType{Nullability: types.NullabilityRequired}
}

func (ex *MultiOrList) ToProto() *proto.Expression {
	toSlice := func(exprs []Expression) (out []*proto.Expression) {
		out = make([]*proto.Expression, len(exprs))
		for i, e := range exprs {
			out[i] = e.ToProto()
		}
		return
	}

	opts := make([]*proto.Expression_MultiOrList_Record, len(ex.Options))
	for i, o := range ex.Options {
		opts[i] = &proto.Expression_MultiOrList_Record{
			Fields: toSlice(o),
		}
	}

	return &proto.Expression{
		RexType: &proto.Expression_MultiOrList_{
			MultiOrList: &proto.Expression_MultiOrList{
				Value:   toSlice(ex.Value),
				Options: opts,
			},
		},
	}
}

func (ex *MultiOrList) Equals(other Expression) bool {
	rhs, ok := other.(*MultiOrList)
	if !ok {
		return false
	}

	return slices.EqualFunc(ex.Value, rhs.Value, exprEqual) &&
		slices.EqualFunc(ex.Options, rhs.Options,
			func(l, r []Expression) bool {
				return slices.EqualFunc(l, r, exprEqual)
			})
}

func (ex *MultiOrList) Visit(visit VisitFunc) Expression {
	var out *MultiOrList
	for i, v := range ex.Value {
		after := visit(v)

		if out == nil && after != v {
			out = &MultiOrList{
				Value:   make([]Expression, len(ex.Value)),
				Options: make([][]Expression, len(ex.Options)),
			}
			copy(out.Value, ex.Value[:i])
		}

		if out != nil {
			out.Value[i] = after
		}
	}

	for i, opts := range ex.Options {
		if out != nil && len(out.Options[i]) == 0 {
			out.Options[i] = make([]Expression, len(ex.Options[i]))
		}

		for j, o := range opts {
			after := visit(o)

			if out == nil && after != o {
				out = &MultiOrList{
					Value:   slices.Clone(ex.Value),
					Options: make([][]Expression, len(ex.Options)),
				}
				for k := 0; k < i; k++ {
					out.Options[k] = slices.Clone(ex.Options[k])
				}

				for k := 0; k < j; k++ {
					out.Options[i][k] = ex.Options[i][k]
				}
			}

			if out != nil {
				out.Options[i][j] = after
			}
		}
	}

	if out == nil {
		return ex
	}

	return out
}

type NestedExpr interface {
	Expression

	IsNullable() bool
	TypeVariation() uint32
}

type MapExpr struct {
	Nullable         bool
	TypeVariationRef uint32
	KeyValues        []struct{ Key, Value Expression }
}

func (ex *MapExpr) IsNullable() bool      { return ex.Nullable }
func (ex *MapExpr) TypeVariation() uint32 { return ex.TypeVariationRef }

func (ex *MapExpr) String() string {
	var b strings.Builder
	b.WriteString("{")

	for i, kv := range ex.KeyValues {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(kv.Key.String())
		b.WriteString(" => ")
		b.WriteString(kv.Value.String())
	}

	b.WriteString("}")
	if ex.Nullable {
		b.WriteString("?")
	}
	fmt.Fprintf(&b, "(typeref=%d)", ex.TypeVariationRef)
	return b.String()
}

func (ex *MapExpr) ToProtoFuncArg() *proto.FunctionArgument {
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Value{
			Value: ex.ToProto(),
		},
	}
}

func (ex *MapExpr) isRootRef() {}

func (ex *MapExpr) IsBound() bool {
	for _, kv := range ex.KeyValues {
		if !kv.Key.IsBound() || !kv.Value.IsBound() {
			return false
		}
	}

	return true
}

func (ex *MapExpr) IsScalar() bool {
	for _, kv := range ex.KeyValues {
		if !kv.Key.IsScalar() || !kv.Value.IsScalar() {
			return false
		}
	}

	return true
}

func (ex *MapExpr) GetType() types.Type {
	// there *should* be at least one element in the keyvalues,
	// otherwise an EmptyMap literal should be used in order
	// to ensure type information is not lost
	return &types.MapType{
		Nullability:      getNullability(ex.Nullable),
		TypeVariationRef: ex.TypeVariationRef,
		Key:              ex.KeyValues[0].Key.GetType(),
		Value:            ex.KeyValues[0].Value.GetType(),
	}
}

func (ex *MapExpr) ToProto() *proto.Expression {
	kvs := make([]*proto.Expression_Nested_Map_KeyValue, len(ex.KeyValues))
	for i, kv := range ex.KeyValues {
		kvs[i] = &proto.Expression_Nested_Map_KeyValue{
			Key:   kv.Key.ToProto(),
			Value: kv.Value.ToProto(),
		}
	}
	return &proto.Expression{
		RexType: &proto.Expression_Nested_{
			Nested: &proto.Expression_Nested{
				Nullable:               ex.Nullable,
				TypeVariationReference: ex.TypeVariationRef,
				NestedType: &proto.Expression_Nested_Map_{
					Map: &proto.Expression_Nested_Map{
						KeyValues: kvs,
					},
				},
			},
		},
	}
}

func (ex *MapExpr) Equals(other Expression) bool {
	rhs, ok := other.(*MapExpr)
	if !ok {
		return false
	}

	if ex.Nullable != rhs.Nullable || ex.TypeVariationRef != rhs.TypeVariationRef {
		return false
	}

	return slices.EqualFunc(ex.KeyValues, rhs.KeyValues,
		func(l, r struct{ Key, Value Expression }) bool {
			return l.Key.Equals(r.Key) && l.Value.Equals(r.Value)
		})
}

func (ex *MapExpr) Visit(visit VisitFunc) Expression {
	var out *MapExpr
	for i, kv := range ex.KeyValues {
		afterKey := visit(kv.Key)
		afterValue := visit(kv.Value)

		if out == nil && (afterKey != kv.Key || afterValue != kv.Value) {
			out = &MapExpr{
				Nullable:         ex.Nullable,
				TypeVariationRef: ex.TypeVariationRef,
				KeyValues: make([]struct {
					Key   Expression
					Value Expression
				}, len(ex.KeyValues)),
			}
			copy(out.KeyValues, ex.KeyValues[:i])
		}

		if out != nil {
			out.KeyValues[i].Key = afterKey
			out.KeyValues[i].Value = afterValue
		}
	}

	if out != nil {
		return out
	}

	return ex
}

type StructExpr struct {
	Nullable         bool
	TypeVariationRef uint32
	Fields           []Expression
}

func (ex *StructExpr) IsNullable() bool      { return ex.Nullable }
func (ex *StructExpr) TypeVariation() uint32 { return ex.TypeVariationRef }

func (ex *StructExpr) String() string {
	var b strings.Builder
	b.WriteString("{")
	for i, f := range ex.Fields {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(f.String())
	}
	b.WriteString("}")
	if ex.Nullable {
		b.WriteString("?")
	}
	fmt.Fprintf(&b, "(typeref=%d)", ex.TypeVariationRef)
	return b.String()
}

func (ex *StructExpr) ToProtoFuncArg() *proto.FunctionArgument {
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Value{
			Value: ex.ToProto(),
		},
	}
}
func (ex *StructExpr) isRootRef() {}

func (ex *StructExpr) IsBound() bool {
	for _, f := range ex.Fields {
		if !f.IsBound() {
			return false
		}
	}
	return true
}

func (ex *StructExpr) IsScalar() bool {
	for _, f := range ex.Fields {
		if !f.IsScalar() {
			return false
		}
	}
	return true
}

func (ex *StructExpr) GetType() types.Type {
	typs := make([]types.Type, len(ex.Fields))
	for i, f := range ex.Fields {
		typs[i] = f.GetType()
	}
	return &types.StructType{
		Nullability:      getNullability(ex.Nullable),
		TypeVariationRef: ex.TypeVariationRef,
		Types:            typs,
	}
}

func (ex *StructExpr) ToProto() *proto.Expression {
	fields := make([]*proto.Expression, len(ex.Fields))
	for i, f := range ex.Fields {
		fields[i] = f.ToProto()
	}
	return &proto.Expression{
		RexType: &proto.Expression_Nested_{
			Nested: &proto.Expression_Nested{
				Nullable:               ex.Nullable,
				TypeVariationReference: ex.TypeVariationRef,
				NestedType: &proto.Expression_Nested_Struct_{
					Struct: &proto.Expression_Nested_Struct{
						Fields: fields,
					},
				},
			},
		},
	}
}

func (ex *StructExpr) Equals(other Expression) bool {
	rhs, ok := other.(*StructExpr)
	if !ok {
		return false
	}

	return ex.Nullable == rhs.Nullable &&
		ex.TypeVariationRef == rhs.TypeVariationRef &&
		slices.EqualFunc(ex.Fields, rhs.Fields, exprEqual)
}

func (ex *StructExpr) Visit(visit VisitFunc) Expression {
	var out *StructExpr
	for i, f := range ex.Fields {
		after := visit(f)
		if out == nil && after != f {
			out = &StructExpr{
				Nullable:         ex.Nullable,
				TypeVariationRef: ex.TypeVariationRef,
				Fields:           make([]Expression, len(ex.Fields)),
			}
			copy(out.Fields, ex.Fields[:i])
		}

		if out != nil {
			out.Fields[i] = after
		}
	}

	if out == nil {
		return ex
	}

	return out
}

type ListExpr struct {
	Nullable         bool
	TypeVariationRef uint32
	Values           []Expression
}

func NewListExpr(nullable bool, vals ...Expression) *ListExpr {
	return &ListExpr{
		Nullable: nullable,
		Values:   vals,
	}
}

func (ex *ListExpr) IsNullable() bool      { return ex.Nullable }
func (ex *ListExpr) TypeVariation() uint32 { return ex.TypeVariationRef }

func (ex *ListExpr) String() string {
	var b strings.Builder
	b.WriteString("[")
	for i, v := range ex.Values {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(v.String())
	}
	b.WriteString("]")
	if ex.Nullable {
		b.WriteString("?")
	}
	fmt.Fprintf(&b, "(typeref=%d)", ex.TypeVariationRef)
	return b.String()
}

func (ex *ListExpr) ToProtoFuncArg() *proto.FunctionArgument {
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Value{
			Value: ex.ToProto(),
		},
	}
}
func (ex *ListExpr) isRootRef() {}

func (ex *ListExpr) IsBound() bool {
	for _, v := range ex.Values {
		if !v.IsBound() {
			return false
		}
	}
	return true
}

func (ex *ListExpr) IsScalar() bool {
	for _, v := range ex.Values {
		if !v.IsScalar() {
			return false
		}
	}
	return true
}

func (ex *ListExpr) GetType() types.Type {
	return &types.ListType{
		Nullability:      getNullability(ex.Nullable),
		TypeVariationRef: ex.TypeVariationRef,
		// to specify an empty list, use an emptylist literal
		// otherwise you will be missing type information.
		// thus we should assume there's at least one value
		Type: ex.Values[0].GetType(),
	}
}

func (ex *ListExpr) ToProto() *proto.Expression {
	vals := make([]*proto.Expression, len(ex.Values))
	for i, v := range ex.Values {
		vals[i] = v.ToProto()
	}
	return &proto.Expression{
		RexType: &proto.Expression_Nested_{
			Nested: &proto.Expression_Nested{
				Nullable:               ex.Nullable,
				TypeVariationReference: ex.TypeVariationRef,
				NestedType: &proto.Expression_Nested_List_{
					List: &proto.Expression_Nested_List{
						Values: vals,
					},
				},
			},
		},
	}
}

func (ex *ListExpr) Equals(other Expression) bool {
	rhs, ok := other.(*ListExpr)
	if !ok {
		return false
	}

	if ex.Nullable != rhs.Nullable || ex.TypeVariationRef != rhs.TypeVariationRef {
		return false
	}

	return slices.EqualFunc(ex.Values, rhs.Values, exprEqual)
}

func (ex *ListExpr) Visit(visit VisitFunc) Expression {
	var out *ListExpr
	for i, v := range ex.Values {
		after := visit(v)

		if out == nil && after != v {
			out = &ListExpr{
				Nullable:         ex.Nullable,
				TypeVariationRef: ex.TypeVariationRef,
				Values:           make([]Expression, len(ex.Values)),
			}
			copy(out.Values, ex.Values[:i])
		}

		if out != nil {
			out.Values[i] = after
		}
	}

	if out != nil {
		return out
	}

	return ex
}

type ExpressionReference struct {
	OutputNames []string
	// only one of these should ever be set at a time
	// this is ensured by the explicit setters rather than
	// making them public
	expr    Expression
	measure *AggregateFunction
}

func NewExpressionReference(names []string, ex Expression) ExpressionReference {
	return ExpressionReference{OutputNames: names, expr: ex}
}

func NewMeasureReference(names []string, measure *AggregateFunction) ExpressionReference {
	return ExpressionReference{OutputNames: names, measure: measure}
}

func (er *ExpressionReference) ToProto() *proto.ExpressionReference {
	out := &proto.ExpressionReference{OutputNames: er.OutputNames}
	switch {
	case er.expr != nil:
		out.ExprType = &proto.ExpressionReference_Expression{
			Expression: er.expr.ToProto(),
		}
	case er.measure != nil:
		out.ExprType = &proto.ExpressionReference_Measure{
			Measure: er.measure.ToProto(),
		}
	}

	return out
}

func (er *ExpressionReference) SetExpr(ex Expression) {
	er.expr = ex
	er.measure = nil
}

func (er *ExpressionReference) SetMeasure(m *AggregateFunction) {
	er.expr = nil
	er.measure = m
}

func (er *ExpressionReference) GetExpr() Expression            { return er.expr }
func (er *ExpressionReference) GetMeasure() *AggregateFunction { return er.measure }

type Extended struct {
	Version          *types.Version
	Extensions       extensions.Set
	ReferredExpr     []ExpressionReference
	BaseSchema       types.NamedStruct
	AdvancedExts     *extensions.AdvancedExtension
	ExpectedTypeURLs []string
}

func ExtendedFromProto(ex *proto.ExtendedExpression, c *extensions.Collection) (*Extended, error) {
	var (
		base   = types.NewNamedStructFromProto(ex.BaseSchema)
		extSet = extensions.GetExtensionSet(ex)
		refs   = make([]ExpressionReference, len(ex.ReferredExpr))
	)

	for i, r := range ex.ReferredExpr {
		refs[i].OutputNames = r.OutputNames
		switch et := r.ExprType.(type) {
		case *proto.ExpressionReference_Expression:
			expr, err := ExprFromProto(et.Expression, &base.Struct, extSet, c)
			if err != nil {
				return nil, err
			}
			refs[i].SetExpr(expr)
		case *proto.ExpressionReference_Measure:
			agg, err := NewAggregateFunctionFromProto(et.Measure, &base.Struct, extSet, c)
			if err != nil {
				return nil, err
			}
			refs[i].SetMeasure(agg)
		}
	}

	return &Extended{
		Version:          ex.Version,
		Extensions:       extSet,
		ReferredExpr:     refs,
		BaseSchema:       base,
		AdvancedExts:     ex.AdvancedExtensions,
		ExpectedTypeURLs: ex.ExpectedTypeUrls,
	}, nil
}

func (ex *Extended) ToProto() *proto.ExtendedExpression {
	uris, decls := ex.Extensions.ToProto()
	refs := make([]*proto.ExpressionReference, len(ex.ReferredExpr))
	for i, ref := range ex.ReferredExpr {
		refs[i] = ref.ToProto()
	}

	return &proto.ExtendedExpression{
		Version:            ex.Version,
		ExtensionUris:      uris,
		Extensions:         decls,
		BaseSchema:         ex.BaseSchema.ToProto(),
		AdvancedExtensions: ex.AdvancedExts,
		ExpectedTypeUrls:   ex.ExpectedTypeURLs,
		ReferredExpr:       refs,
	}
}

// BindExpression recursively visits the constituent expressions looks up
// the extension ID for any scalar/window functions in the provided extension
// collection and retrieves or assigns a function anchor for them in the
// provided extension set. It will also recursively resolve the output types
// for the functions and field references in the expression with the provided
// base schema as the "root" schema to utilize.
//
// If a given ID cannot be found in the collection and the name in the ID
// does not already contain a ":", the provided arguments will be used to
// construct a compound name for the function call to attempt to lookup in
// the collection. If a function declaration still cannot be located, then
// an error will be returned.
func BindExpression(ex Expression, baseSchema types.NamedStruct, extSet ExtensionLookup, c *extensions.Collection) (out Expression, err error) {
	defer func() {
		if r := recover(); r != nil {
			switch r := r.(type) {
			case error:
				err = r
			default:
				err = fmt.Errorf("%w: %s", substraitgo.ErrInvalidExpr, r)
			}
		}
	}()

	var visitor VisitFunc

	visitor = func(e Expression) Expression {
		switch e := e.(type) {
		case *ScalarFunction:
			funcRef := extSet.GetFuncAnchor(e.ID)
			if e.FuncRef == funcRef && e.IsBound() {
				return e
			}

			out := e.Visit(visitor).(*ScalarFunction)
			if out == e {
				out = &ScalarFunction{
					FuncRef:    funcRef,
					ID:         e.ID,
					Args:       slices.Clone(e.Args),
					Options:    slices.Clone(e.Options),
					OutputType: e.OutputType,
				}
			} else {
				out.FuncRef = funcRef
			}

			argTypes := make([]types.Type, 0, len(out.Args))
			for _, arg := range out.Args {
				switch a := arg.(type) {
				case types.Enum:
					argTypes = append(argTypes, nil)
				case Expression:
					argTypes = append(argTypes, a.GetType())
				}
			}

			if variant, found := c.GetScalarFunc(out.ID); found {
				out.Declaration = variant
			} else {
				if strings.IndexByte(out.ID.Name, ':') == -1 {
					sigs := make([]string, len(argTypes))
					for i, t := range argTypes {
						if t == nil {
							// enum value
							sigs[i] = "req"
						} else if ud, ok := t.(*types.UserDefinedType); ok {
							id, found := extSet.DecodeType(ud.TypeReference)
							if !found {
								panic(fmt.Errorf("%w: could not find type for reference %d",
									substraitgo.ErrNotFound, ud.TypeReference))
							}
							sigs[i] = "u!" + id.Name
						} else {
							sigs[i] = t.ShortString()
						}
					}

					out.ID.Name += ":" + strings.Join(sigs, "_")
					if variant, found = c.GetScalarFunc(out.ID); !found {
						panic(fmt.Errorf("%w: could not find matching function for %v", substraitgo.ErrNotFound, out.ID))
					}
					out.Declaration = variant
				}
			}

			if out.OutputType == nil {
				outType, err := out.Declaration.ResolveType(argTypes)
				if err != nil {
					panic(err)
				}
				out.OutputType = outType
			}

			return out
		case *WindowFunction:
			funcRef := extSet.GetFuncAnchor(e.ID)
			if e.FuncRef == funcRef && e.IsBound() {
				return e
			}

			out := e.Visit(visitor).(*WindowFunction)
			if out == e {
				out = &WindowFunction{
					FuncRef:    funcRef,
					ID:         e.ID,
					Args:       slices.Clone(e.Args),
					Options:    slices.Clone(e.Options),
					OutputType: e.OutputType,
					Phase:      e.Phase,
					Sorts:      slices.Clone(e.Sorts),
					Invocation: e.Invocation,
					Partitions: slices.Clone(e.Partitions),
					LowerBound: e.LowerBound,
					UpperBound: e.UpperBound,
				}
			} else {
				out.FuncRef = funcRef
			}

			argTypes := make([]types.Type, 0, len(out.Args))
			for i, arg := range out.Args {
				switch a := arg.(type) {
				case Expression:
					a = a.Visit(visitor)
					out.Args[i] = a

					argTypes = append(argTypes, a.GetType())
				}
			}

			if variant, found := c.GetWindowFunc(out.ID); found {
				out.Declaration = variant
			} else {
				if strings.IndexByte(out.ID.Name, ':') == -1 {
					sigs := make([]string, len(argTypes))
					for i, t := range argTypes {
						if ud, ok := t.(*types.UserDefinedType); ok {
							id, found := extSet.DecodeType(ud.TypeReference)
							if !found {
								panic(fmt.Errorf("%w: could not find type for reference %d",
									substraitgo.ErrNotFound, ud.TypeReference))
							}
							sigs[i] = "u!" + id.Name
						} else {
							sigs[i] = t.ShortString()
						}
					}

					out.ID.Name += ":" + strings.Join(sigs, "_")
					if variant, found = c.GetWindowFunc(out.ID); !found {
						panic(fmt.Errorf("%w: could not find matching function for %v", substraitgo.ErrNotFound, out.ID))
					}
					out.Declaration = variant
				}
			}

			if out.OutputType == nil {
				outType, err := out.Declaration.ResolveType(argTypes)
				if err != nil {
					panic(err)
				}
				out.OutputType = outType
			}

			for i, p := range out.Partitions {
				out.Partitions[i] = visitor(p)
			}

			return out

		case *FieldReference:
			if e.IsBound() {
				return e
			}

			out := *e

			if rootExpr, ok := e.Root.(Expression); ok && !rootExpr.IsBound() {
				out.Root = visitor(rootExpr)
			}

			out.UpdateType(&baseSchema.Struct)
			return &out
		default:
			return e.Visit(visitor)
		}
	}

	return visitor(ex), nil
}
