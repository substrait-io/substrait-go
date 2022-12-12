// SPDX-License-Identifier: Apache-2.0

package substraitgo

import (
	"fmt"

	"github.com/substrait-io/substrait-go/proto"
)

type RootRefType interface{}

var RootReference struct{}

type OuterReference uint32

type ReferenceSegment interface {
	GetChild() ReferenceSegment
}

type MapKeyRef struct {
	MapKey Literal
	Child  ReferenceSegment
}

func (r *MapKeyRef) GetChild() ReferenceSegment { return r.Child }
func (*MapKeyRef) isRefType()                   {}

type StructFieldRef struct {
	Field int32
	Child ReferenceSegment
}

func (r *StructFieldRef) GetChild() ReferenceSegment { return r.Child }
func (*StructFieldRef) isRefType()                   {}

type ListElementRef struct {
	Offset int32
	Child  ReferenceSegment
}

func (r *ListElementRef) GetChild() ReferenceSegment { return r.Child }
func (*ListElementRef) isRefType()                   {}

type MaskExpression proto.Expression_MaskExpression

func (*MaskExpression) isRefType() {}

type Reference interface {
	isRefType()
}

type FieldReference struct {
	Reference Reference
	Root      RootRefType
}

func ExprFromProto(e *proto.Expression) (Expression, error) {
	switch et := e.RexType.(type) {
	case *proto.Expression_Literal_:
		return LiteralFromProto(et.Literal), nil
	case *proto.Expression_Selection:
	case *proto.Expression_ScalarFunction_:
	case *proto.Expression_WindowFunction_:
	case *proto.Expression_IfThen_:
	case *proto.Expression_SwitchExpression_:
	case *proto.Expression_SingularOrList_:
	case *proto.Expression_MultiOrList_:
	case *proto.Expression_Cast_:
	case *proto.Expression_Nested_:
	case *proto.Expression_Enum_:
		return nil, fmt.Errorf("%w: deprecated", ErrNotImplemented)
	case *proto.Expression_Subquery_:
	}

	return nil, ErrNotImplemented
}

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
	FuncArg
	ToProto() *proto.Expression
	Equals(Expression) bool
}

type IfThen struct {
	IFs []struct {
		If   Expression
		Then Expression
	}
	Else Expression
}

type Cast struct {
	Type            Type
	Input           Expression
	FailureBehavior CastFailBehavior
}

type SwitchExpr struct {
	Match Expression
	IFs   []struct {
		If   Literal
		Then Expression
	}
	Else Expression
}

type SinglularOrList struct {
	Value   Expression
	Options []Expression
}

type MultiOrList struct {
	Value   Expression
	Options []struct {
		Fields []Expression
	}
}

type NestedExpr interface {
	FuncArg

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

type StructExpr struct {
	Nullable         bool
	TypeVariationRef uint32
	Fields           []Expression
}

func (s *StructExpr) IsNullable() bool      { return s.Nullable }
func (s *StructExpr) TypeVariation() uint32 { return s.TypeVariationRef }

type ListExpr struct {
	Nullable         bool
	TypeVariationRef uint32
	Values           []Expression
}

func (l *ListExpr) IsNullable() bool      { return l.Nullable }
func (l *ListExpr) TypeVariation() uint32 { return l.TypeVariationRef }
