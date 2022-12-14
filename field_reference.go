// SPDX-License-Identifier: Apache-2.0

package substraitgo

import (
	"fmt"

	"github.com/substrait-io/substrait-go/proto"
)

type RootRefType interface {
	isRootRef()
}

var RootReference RootRefType

type OuterReference uint32

func (OuterReference) isRootRef() {}

type ReferenceSegment interface {
	Reference
	fmt.Stringer
	GetChild() ReferenceSegment
	GetType(Type) (Type, error)
	ToProto() *proto.Expression_ReferenceSegment
}

func RefSegmentFromProto(p *proto.Expression_ReferenceSegment) ReferenceSegment {
	if p == nil {
		return nil
	}

	switch seg := p.ReferenceType.(type) {
	case *proto.Expression_ReferenceSegment_MapKey_:
		return &MapKeyRef{
			MapKey: LiteralFromProto(seg.MapKey.MapKey),
			Child:  RefSegmentFromProto(seg.MapKey.Child),
		}
	case *proto.Expression_ReferenceSegment_StructField_:
		return &StructFieldRef{
			Field: seg.StructField.Field,
			Child: RefSegmentFromProto(seg.StructField.Child),
		}
	case *proto.Expression_ReferenceSegment_ListElement_:
		return &ListElementRef{
			Offset: seg.ListElement.Offset,
			Child:  RefSegmentFromProto(seg.ListElement.Child),
		}
	}

	return nil
}

type MapKeyRef struct {
	MapKey Literal
	Child  ReferenceSegment
}

func (r *MapKeyRef) String() string {
	var c string
	if r.Child != nil {
		c = r.Child.String()
	}
	return ".[" + r.MapKey.String() + "]" + c
}

func (r *MapKeyRef) ToProto() *proto.Expression_ReferenceSegment {
	var c *proto.Expression_ReferenceSegment
	if r.Child != nil {
		c = r.Child.ToProto()
	}

	return &proto.Expression_ReferenceSegment{
		ReferenceType: &proto.Expression_ReferenceSegment_MapKey_{
			MapKey: &proto.Expression_ReferenceSegment_MapKey{
				MapKey: r.MapKey.ToProtoLiteral(),
				Child:  c,
			},
		},
	}
}

func (r *MapKeyRef) GetType(parentType Type) (Type, error) {
	mt, ok := parentType.(*MapType)
	if !ok {
		return nil, ErrInvalidType
	}

	if !r.MapKey.GetType().Equals(mt.Key) {
		return nil, ErrInvalidType
	}

	if r.Child != nil {
		return r.Child.GetType(mt.Value)
	}

	return mt.Value, nil
}
func (r *MapKeyRef) GetChild() ReferenceSegment { return r.Child }
func (*MapKeyRef) isRefType()                   {}

type StructFieldRef struct {
	Field int32
	Child ReferenceSegment
}

func (r *StructFieldRef) String() string {
	var c string
	if r.Child != nil {
		c = r.Child.String()
	}

	return fmt.Sprintf(".field(%d)%s", r.Field, c)
}

func (r *StructFieldRef) GetType(parentType Type) (Type, error) {
	st, ok := parentType.(*StructType)
	if !ok {
		return nil, ErrInvalidType
	}

	if len(st.Types) < int(r.Field) {
		return nil, ErrInvalidType
	}

	if r.Child != nil {
		return r.Child.GetType(st.Types[r.Field])
	}

	return st.Types[r.Field], nil
}

func (r *StructFieldRef) ToProto() *proto.Expression_ReferenceSegment {
	var c *proto.Expression_ReferenceSegment
	if r.Child != nil {
		c = r.Child.ToProto()
	}

	return &proto.Expression_ReferenceSegment{
		ReferenceType: &proto.Expression_ReferenceSegment_StructField_{
			StructField: &proto.Expression_ReferenceSegment_StructField{
				Field: r.Field,
				Child: c,
			},
		},
	}
}

func (r *StructFieldRef) GetChild() ReferenceSegment { return r.Child }
func (*StructFieldRef) isRefType()                   {}

type ListElementRef struct {
	Offset int32
	Child  ReferenceSegment
}

func (r *ListElementRef) String() string {
	var c string
	if r.Child != nil {
		c = r.Child.String()
	}
	return fmt.Sprintf(".[%d]%s", r.Offset, c)
}

func (r *ListElementRef) GetType(parentType Type) (Type, error) {
	lt, ok := parentType.(*ListType)
	if !ok {
		return nil, ErrInvalidType
	}

	if r.Child != nil {
		return r.Child.GetType(lt.Type)
	}
	return lt.Type, nil
}

func (r *ListElementRef) ToProto() *proto.Expression_ReferenceSegment {
	var c *proto.Expression_ReferenceSegment
	if r.Child != nil {
		c = r.Child.ToProto()
	}

	return &proto.Expression_ReferenceSegment{
		ReferenceType: &proto.Expression_ReferenceSegment_ListElement_{
			ListElement: &proto.Expression_ReferenceSegment_ListElement{
				Offset: r.Offset,
				Child:  c,
			},
		},
	}
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

	knownType Type
}

func (*FieldReference) isRootRef() {}

func (f *FieldReference) String() string {
	if f.Root == RootReference {
		var typ string
		if f.knownType != nil {
			typ = " => " + f.knownType.String()
		}
		return f.Reference.(ReferenceSegment).String() + typ
	}
	return ""
}

func (f *FieldReference) ToProtoFuncArg() *proto.FunctionArgument {
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Value{Value: f.ToProto()},
	}
}
func (f *FieldReference) ToProtoFieldRef() *proto.Expression_FieldReference {
	ret := &proto.Expression_FieldReference{}
	switch r := f.Reference.(type) {
	case ReferenceSegment:
		ret.ReferenceType = &proto.Expression_FieldReference_DirectReference{
			DirectReference: r.ToProto()}
	case *MaskExpression:
		ret.ReferenceType = &proto.Expression_FieldReference_MaskedReference{
			MaskedReference: (*proto.Expression_MaskExpression)(r),
		}
	}

	if f.Root == RootReference {
		ret.RootType = &proto.Expression_FieldReference_RootReference_{}
	} else {
		switch r := f.Root.(type) {
		case Expression:
			ret.RootType = &proto.Expression_FieldReference_Expression{
				Expression: r.ToProto(),
			}
		case OuterReference:
			ret.RootType = &proto.Expression_FieldReference_OuterReference_{
				OuterReference: &proto.Expression_FieldReference_OuterReference{
					StepsOut: uint32(r),
				},
			}
		}
	}

	return ret
}

func (f *FieldReference) ToProto() *proto.Expression {
	return &proto.Expression{
		RexType: &proto.Expression_Selection{
			Selection: f.ToProtoFieldRef(),
		},
	}
}

func (f *FieldReference) Equals(rhs Expression) bool { return false }

func (f *FieldReference) GetType() Type {
	return f.knownType
}

func FieldReferenceFromProto(p *proto.Expression_FieldReference, baseSchema Type, ext ExtensionRegistry) (*FieldReference, error) {
	var (
		ref       Reference
		root      RootRefType
		knownType Type
		err       error
	)

	switch rt := p.RootType.(type) {
	case *proto.Expression_FieldReference_Expression:
		if root, err = ExprFromProto(rt.Expression, baseSchema, ext); err != nil {
			return nil, err
		}
	case *proto.Expression_FieldReference_OuterReference_:
		root = OuterReference(rt.OuterReference.StepsOut)
	case *proto.Expression_FieldReference_RootReference_:
		root = RootReference
	}

	switch rt := p.ReferenceType.(type) {
	case *proto.Expression_FieldReference_DirectReference:
		refseg := RefSegmentFromProto(rt.DirectReference)
		if root == RootReference && baseSchema != nil {
			knownType, err = refseg.GetType(baseSchema)
			if err != nil {
				return nil, err
			}
		} else if rootExpr, ok := root.(Expression); ok {
			knownType, err = refseg.GetType(rootExpr.GetType())
			if err != nil {
				return nil, err
			}
		}

		ref = refseg

	case *proto.Expression_FieldReference_MaskedReference:
		ref = (*MaskExpression)(rt.MaskedReference)
	}

	return &FieldReference{Root: root, Reference: ref, knownType: knownType}, nil
}
