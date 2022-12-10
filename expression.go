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
	funcArg

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

type Expression interface {
	FuncArg
	ToProto() *proto.Expression
	Equals(Expression) bool
}

type IfThen struct {
	funcArg

	IFs []struct {
		If   Expression
		Then Expression
	}
	Else Expression
}

type Cast struct {
	funcArg

	Type            Type
	Input           Expression
	FailureBehavior CastFailBehavior
}

type SwitchExpression struct {
	funcArg

	Match Expression
	IFs   []struct {
		If   Literal
		Then Expression
	}
	Else Expression
}

type SinglularOrList struct {
	funcArg

	Value   Expression
	Options []Expression
}

type MultiOrList struct {
	funcArg

	Value   Expression
	Options []struct {
		Fields []Expression
	}
}

type Nested interface {
	FuncArg

	IsNullable() bool
	TypeVariation() uint32
}

type Map struct {
	funcArg

	Nullable         bool
	TypeVariationRef uint32
	KeyValues        []struct{ Key, Value Expression }
}

func (m *Map) IsNullable() bool      { return m.Nullable }
func (m *Map) TypeVariation() uint32 { return m.TypeVariationRef }

type Struct struct {
	funcArg

	Nullable         bool
	TypeVariationRef uint32
	Fields           []Expression
}

func (s *Struct) IsNullable() bool      { return s.Nullable }
func (s *Struct) TypeVariation() uint32 { return s.TypeVariationRef }

type List struct {
	funcArg

	Nullable         bool
	TypeVariationRef uint32
	Values           []Expression
}

func (l *List) IsNullable() bool      { return l.Nullable }
func (l *List) TypeVariation() uint32 { return l.TypeVariationRef }
