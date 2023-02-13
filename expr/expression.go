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

		id, ok := ext.DecodeFunc(et.ScalarFunction.FunctionReference)
		if !ok {
			return nil, substraitgo.ErrNotFound
		}

		decl, _ := ext.LookupScalarFunction(et.ScalarFunction.FunctionReference, c)
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

		id, ok := ext.DecodeFunc(et.WindowFunction.FunctionReference)
		if !ok {
			return nil, substraitgo.ErrNotFound
		}

		decl, _ := ext.LookupWindowFunction(et.WindowFunction.FunctionReference, c)
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
	case *proto.Expression_SwitchExpression_:
	case *proto.Expression_SingularOrList_:
	case *proto.Expression_MultiOrList_:
	case *proto.Expression_Cast_:
	case *proto.Expression_Nested_:
	case *proto.Expression_Enum_:
		return nil, fmt.Errorf("%w: deprecated", substraitgo.ErrNotImplemented)
	case *proto.Expression_Subquery_:
	}

	return nil, substraitgo.ErrNotImplemented
}

type VisitFunc func(Expression) Expression

// Expression can be one of many different things as a generalized
// expression. It could be:
//
//  - A literal
//  - A Field Reference Selection
//  - A Scalar Function expression
//  - A Window Function expression
//  - An If-Then statement
//  - A Switch Expression
//  - A Singular Or List
//  - A Multiple Or List
//  - A Cast expression
//  - A Subquery
//  - A Nested expression
type Expression interface {
	// an Expression can also be a function argument
	types.FuncArg
	// an expression can also be the root of a reference
	RootRefType

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
		ifthenClauses[i].If = c.If.ToProto()
		ifthenClauses[i].Then = c.Then.ToProto()
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

	if len(ex.IFs) != len(rhs.IFs) {
		return false
	}

	for i := range ex.IFs {
		if !ex.IFs[i].If.Equals(rhs.IFs[i].If) {
			return false
		}

		if !ex.IFs[i].Then.Equals(rhs.IFs[i].Then) {
			return false
		}
	}

	return ex.Else != nil && ex.Else.Equals(rhs.Else)
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
	if len(ex.IFs) > 0 {
		return ex.IFs[0].Then.GetType()
	}
	return ex.Else.GetType()
}

func (ex *SwitchExpr) ToProto() *proto.Expression {
	var elseExpr *proto.Expression
	if ex.Else != nil {
		elseExpr = ex.Else.ToProto()
	}

	cases := make([]*proto.Expression_SwitchExpression_IfValue, len(ex.IFs))
	for i, c := range ex.IFs {
		cases[i].If = c.If.ToProtoLiteral()
		cases[i].Then = c.Then.ToProto()
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
	case len(ex.IFs) != len(rhs.IFs):
		return false
	case !ex.Match.Equals(rhs.Match):
		return false
	}

	for i := range ex.IFs {
		if !ex.IFs[i].If.Equals(rhs.IFs[i].If) {
			return false
		}
		if !ex.IFs[i].Then.Equals(rhs.IFs[i].Then) {
			return false
		}
	}

	if ex.Else != nil && ex.Else.Equals(rhs.Else) {
		return true
	}

	// if rhs.Else == nil then we're equal, otherwise
	// ex.Else is nil and rhs.Else is not nil
	return rhs.Else == nil
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
			for j := 0; j < i; j++ {
				out.IFs[j] = ex.IFs[j]
			}
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
	return &types.BooleanType{Nullability: types.NullabilityNullable}
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

	if len(ex.Options) != len(rhs.Options) {
		return false
	}

	for i := range ex.Options {
		if !ex.Options[i].Equals(rhs.Options[i]) {
			return false
		}
	}

	return true
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
			for j := 0; j < len(ex.Options); j++ {
				out.Options[j] = ex.Options[j]
			}
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
	b.WriteString("IN [")
	for i, opts := range ex.Options {
		if i != 0 {
			b.WriteString(", ")
		}
		writeList(opts)
	}
	b.WriteString("]")

	return b.String()
}

func (ex *MultiOrList) ToProtoFuncArg() *proto.FunctionArgument {}
func (ex *MultiOrList) isRootRef()                              {}
func (ex *MultiOrList) IsScalar() bool                          {}
func (ex *MultiOrList) GetType() types.Type                     {}
func (ex *MultiOrList) ToProto() *proto.Expression              {}
func (ex *MultiOrList) Equals(Expression) bool                  {}
func (ex *MultiOrList) Visit(VisitFunc) Expression              {}

type NestedExpr interface {
	types.FuncArg

	IsNullable() bool
	TypeVariation() uint32
}

type MapExpr struct {
	Nullable         bool
	TypeVariationRef uint32
	KeyValues        []struct{ Key, Value Expression }
}

func (m *MapExpr) IsNullable() bool      { return m.Nullable }
func (m *MapExpr) TypeVariation() uint32 { return m.TypeVariationRef }

func (ex *MapExpr) String() string                          {}
func (ex *MapExpr) ToProtoFuncArg() *proto.FunctionArgument {}
func (ex *MapExpr) isRootRef()                              {}
func (ex *MapExpr) IsScalar() bool                          {}
func (ex *MapExpr) GetType() types.Type                     {}
func (ex *MapExpr) ToProto() *proto.Expression              {}
func (ex *MapExpr) Equals(Expression) bool                  {}
func (ex *MapExpr) Visit(VisitFunc) Expression              {}

type StructExpr struct {
	Nullable         bool
	TypeVariationRef uint32
	Fields           []Expression
}

func (s *StructExpr) IsNullable() bool      { return s.Nullable }
func (s *StructExpr) TypeVariation() uint32 { return s.TypeVariationRef }

func (ex *StructExpr) String() string                          {}
func (ex *StructExpr) ToProtoFuncArg() *proto.FunctionArgument {}
func (ex *StructExpr) isRootRef()                              {}
func (ex *StructExpr) IsScalar() bool                          {}
func (ex *StructExpr) GetType() types.Type                     {}
func (ex *StructExpr) ToProto() *proto.Expression              {}
func (ex *StructExpr) Equals(Expression) bool                  {}
func (ex *StructExpr) Visit(VisitFunc) Expression              {}

type ListExpr struct {
	Nullable         bool
	TypeVariationRef uint32
	Values           []Expression
}

func (l *ListExpr) IsNullable() bool      { return l.Nullable }
func (l *ListExpr) TypeVariation() uint32 { return l.TypeVariationRef }

func (ex *ListExpr) String() string                          {}
func (ex *ListExpr) ToProtoFuncArg() *proto.FunctionArgument {}
func (ex *ListExpr) isRootRef()                              {}
func (ex *ListExpr) IsScalar() bool                          {}
func (ex *ListExpr) GetType() types.Type                     {}
func (ex *ListExpr) ToProto() *proto.Expression              {}
func (ex *ListExpr) Equals(Expression) bool                  {}
func (ex *ListExpr) Visit(VisitFunc) Expression              {}

type ExpressionReference struct {
	OutputNames []string
	// only one of these should ever be set at a time
	// this is ensured by the explicit setters rather than
	// making them public
	expr    Expression
	measure *AggregateFunction
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
