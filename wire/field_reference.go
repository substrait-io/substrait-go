// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"fmt"

	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// RefSegmentToProto encodes a reference segment as its protobuf message.
func RefSegmentToProto(r expr.ReferenceSegment) *proto.Expression_ReferenceSegment {
	switch r := r.(type) {
	case *expr.MapKeyRef:
		return mapKeyRefToProto(r)
	case *expr.StructFieldRef:
		return structFieldRefToProto(r)
	case *expr.ListElementRef:
		return listElementRefToProto(r)
	default:
		panic(fmt.Sprintf("wire: unhandled reference segment %T", r))
	}
}

func mapKeyRefToProto(r *expr.MapKeyRef) *proto.Expression_ReferenceSegment {
	var child *proto.Expression_ReferenceSegment
	if r.Child != nil {
		child = RefSegmentToProto(r.Child)
	}
	return &proto.Expression_ReferenceSegment{
		ReferenceType: &proto.Expression_ReferenceSegment_MapKey_{
			MapKey: &proto.Expression_ReferenceSegment_MapKey{
				MapKey: LiteralToProto(r.MapKey),
				Child:  child,
			},
		},
	}
}

func structFieldRefToProto(r *expr.StructFieldRef) *proto.Expression_ReferenceSegment {
	var child *proto.Expression_ReferenceSegment
	if r.Child != nil {
		child = RefSegmentToProto(r.Child)
	}
	return &proto.Expression_ReferenceSegment{
		ReferenceType: &proto.Expression_ReferenceSegment_StructField_{
			StructField: &proto.Expression_ReferenceSegment_StructField{
				Field: r.Field,
				Child: child,
			},
		},
	}
}

func listElementRefToProto(r *expr.ListElementRef) *proto.Expression_ReferenceSegment {
	var child *proto.Expression_ReferenceSegment
	if r.Child != nil {
		child = RefSegmentToProto(r.Child)
	}
	return &proto.Expression_ReferenceSegment{
		ReferenceType: &proto.Expression_ReferenceSegment_ListElement_{
			ListElement: &proto.Expression_ReferenceSegment_ListElement{
				Offset: r.Offset,
				Child:  child,
			},
		},
	}
}

// MaskExpressionToProto encodes a mask expression as its protobuf message.
func MaskExpressionToProto(e *expr.MaskExpression) *proto.Expression_MaskExpression {
	return &proto.Expression_MaskExpression{
		Select:                 maskStructSelectToProto(e.Select()),
		MaintainSingularStruct: e.MaintainSingularStruct(),
	}
}

func maskStructSelectToProto(m expr.MaskStructSelect) *proto.Expression_MaskExpression_StructSelect {
	items := make([]*proto.Expression_MaskExpression_StructItem, len(m))
	for i := range m {
		items[i] = maskStructItemToProto(&m[i])
	}
	return &proto.Expression_MaskExpression_StructSelect{StructItems: items}
}

func maskSelectToProto(s expr.MaskSelect) *proto.Expression_MaskExpression_Select {
	switch s := s.(type) {
	case expr.MaskStructSelect:
		return &proto.Expression_MaskExpression_Select{
			Type: &proto.Expression_MaskExpression_Select_Struct{Struct: maskStructSelectToProto(s)},
		}
	case *expr.MaskListSelect:
		return maskListSelectToProto(s)
	case *expr.MaskMapSelect:
		return maskMapSelectToProto(s)
	default:
		panic(fmt.Sprintf("wire: unhandled mask selection %T", s))
	}
}

func maskStructItemToProto(m *expr.MaskStructItem) *proto.Expression_MaskExpression_StructItem {
	var child *proto.Expression_MaskExpression_Select
	if c := m.Child(); c != nil {
		child = maskSelectToProto(c)
	}
	return &proto.Expression_MaskExpression_StructItem{
		Field: m.Field(),
		Child: child,
	}
}

func maskListSelectToProto(m *expr.MaskListSelect) *proto.Expression_MaskExpression_Select {
	sel := m.Selection()
	items := make([]*proto.Expression_MaskExpression_ListSelect_ListSelectItem, len(sel))
	for i, s := range sel {
		items[i] = maskListSelectItemToProto(s)
	}
	return &proto.Expression_MaskExpression_Select{
		Type: &proto.Expression_MaskExpression_Select_List{
			List: &proto.Expression_MaskExpression_ListSelect{
				Selection: items,
				Child:     maskSelectToProto(m.Child()),
			},
		},
	}
}

func maskListSelectItemToProto(s expr.MaskListSelectItem) *proto.Expression_MaskExpression_ListSelect_ListSelectItem {
	switch s := s.(type) {
	case *expr.MaskListElement:
		return &proto.Expression_MaskExpression_ListSelect_ListSelectItem{
			Type: &proto.Expression_MaskExpression_ListSelect_ListSelectItem_Item{
				Item: &proto.Expression_MaskExpression_ListSelect_ListSelectItem_ListElement{
					Field: s.GetField(),
				},
			},
		}
	case *expr.MaskListSlice:
		start, end := s.GetBounds()
		return &proto.Expression_MaskExpression_ListSelect_ListSelectItem{
			Type: &proto.Expression_MaskExpression_ListSelect_ListSelectItem_Slice{
				Slice: &proto.Expression_MaskExpression_ListSelect_ListSelectItem_ListSlice{
					Start: start,
					End:   end,
				},
			},
		}
	default:
		panic(fmt.Sprintf("wire: unhandled mask selection %T", s))
	}
}

func maskMapSelectToProto(m *expr.MaskMapSelect) *proto.Expression_MaskExpression_Select {
	mapSelect := &proto.Expression_MaskExpression_Select_Map{
		Map: &proto.Expression_MaskExpression_MapSelect{
			Child: maskSelectToProto(m.Child()),
		},
	}

	if m.KeyKind() == expr.MapSelectKey {
		mapSelect.Map.Select = &proto.Expression_MaskExpression_MapSelect_Key{
			Key: &proto.Expression_MaskExpression_MapSelect_MapKey{MapKey: m.Key()},
		}
	} else {
		mapSelect.Map.Select = &proto.Expression_MaskExpression_MapSelect_Expression{
			Expression: &proto.Expression_MaskExpression_MapSelect_MapKeyExpression{MapKeyExpression: m.Key()},
		}
	}

	return &proto.Expression_MaskExpression_Select{Type: mapSelect}
}

// FieldReferenceToProto encodes a field reference as its protobuf message.
func FieldReferenceToProto(f *expr.FieldReference) *proto.Expression {
	return &proto.Expression{
		RexType: &proto.Expression_Selection{Selection: fieldReferenceRefToProto(f)},
	}
}

func fieldReferenceRefToProto(f *expr.FieldReference) *proto.Expression_FieldReference {
	ret := &proto.Expression_FieldReference{}
	switch r := f.Reference.(type) {
	case expr.ReferenceSegment:
		ret.ReferenceType = &proto.Expression_FieldReference_DirectReference{
			DirectReference: RefSegmentToProto(r),
		}
	case *expr.MaskExpression:
		ret.ReferenceType = &proto.Expression_FieldReference_MaskedReference{
			MaskedReference: MaskExpressionToProto(r),
		}
	}

	if f.Root != expr.RootReference {
		switch r := f.Root.(type) {
		case expr.Expression:
			ret.RootType = &proto.Expression_FieldReference_Expression{
				Expression: ExprToProto(r),
			}
		case expr.OuterReference:
			ret.RootType = &proto.Expression_FieldReference_OuterReference_{
				OuterReference: &proto.Expression_FieldReference_OuterReference{
					StepsOut: uint32(r),
				},
			}
		case expr.LambdaParameterReference:
			ret.RootType = &proto.Expression_FieldReference_LambdaParameterReference_{
				LambdaParameterReference: &proto.Expression_FieldReference_LambdaParameterReference{
					StepsOut: r.StepsOut,
				},
			}
		}
	} else {
		ret.RootType = &proto.Expression_FieldReference_RootReference_{
			RootReference: &proto.Expression_FieldReference_RootReference{},
		}
	}

	return ret
}

// RefSegmentFromProto decodes a reference segment from its protobuf message.
func RefSegmentFromProto(p *proto.Expression_ReferenceSegment) expr.ReferenceSegment {
	if p == nil {
		return nil
	}

	switch seg := p.ReferenceType.(type) {
	case *proto.Expression_ReferenceSegment_MapKey_:
		return &expr.MapKeyRef{
			MapKey: LiteralFromProto(seg.MapKey.MapKey),
			Child:  RefSegmentFromProto(seg.MapKey.Child),
		}
	case *proto.Expression_ReferenceSegment_StructField_:
		return &expr.StructFieldRef{
			Field: seg.StructField.Field,
			Child: RefSegmentFromProto(seg.StructField.Child),
		}
	case *proto.Expression_ReferenceSegment_ListElement_:
		return &expr.ListElementRef{
			Offset: seg.ListElement.Offset,
			Child:  RefSegmentFromProto(seg.ListElement.Child),
		}
	}

	return nil
}

// MaskExpressionFromProto decodes a mask expression from its protobuf message.
func MaskExpressionFromProto(p *proto.Expression_MaskExpression) *expr.MaskExpression {
	sel := make(expr.MaskStructSelect, len(p.Select.StructItems))
	for i, item := range p.Select.StructItems {
		var child expr.MaskSelect
		if item.Child != nil {
			child = maskSelectFromProto(item.Child)
		}
		sel[i] = expr.NewMaskStructItem(item.Field, child)
	}
	return expr.NewMaskExpression(sel, p.MaintainSingularStruct)
}

func maskSelectFromProto(p *proto.Expression_MaskExpression_Select) expr.MaskSelect {
	switch s := p.Type.(type) {
	case *proto.Expression_MaskExpression_Select_Struct:
		items := make(expr.MaskStructSelect, len(s.Struct.StructItems))
		for i, item := range s.Struct.StructItems {
			var child expr.MaskSelect
			if item.Child != nil {
				child = maskSelectFromProto(item.Child)
			}
			items[i] = expr.NewMaskStructItem(item.Field, child)
		}
		return items
	case *proto.Expression_MaskExpression_Select_List:
		selection := make([]expr.MaskListSelectItem, len(s.List.Selection))
		for i, sel := range s.List.Selection {
			switch s := sel.Type.(type) {
			case *proto.Expression_MaskExpression_ListSelect_ListSelectItem_Item:
				selection[i] = &expr.MaskListElement{Field: s.Item.Field}
			case *proto.Expression_MaskExpression_ListSelect_ListSelectItem_Slice:
				selection[i] = &expr.MaskListSlice{Start: s.Slice.Start, End: s.Slice.End}
			}
		}
		return expr.NewMaskListSelect(selection, maskSelectFromProto(s.List.Child))
	case *proto.Expression_MaskExpression_Select_Map:
		var child expr.MaskSelect
		if s.Map.Child != nil {
			child = maskSelectFromProto(s.Map.Child)
		}

		switch sk := s.Map.Select.(type) {
		case *proto.Expression_MaskExpression_MapSelect_Expression:
			return expr.NewMaskMapSelect(expr.MapSelectExpr, sk.Expression.MapKeyExpression, child)
		case *proto.Expression_MaskExpression_MapSelect_Key:
			return expr.NewMaskMapSelect(expr.MapSelectKey, sk.Key.MapKey, child)
		}
		return expr.NewMaskMapSelect(expr.MapSelectKey, "", child)
	}
	panic("unimplemented mask select type")
}

// FieldReferenceFromProto decodes a field reference from its protobuf message.
func FieldReferenceFromProto(p *proto.Expression_FieldReference, baseSchema *types.RecordType, reg expr.ExtensionRegistry) (*expr.FieldReference, error) {
	var (
		ref       expr.Reference
		root      expr.RootRefType
		knownType types.Type
		err       error
	)

	switch rt := p.RootType.(type) {
	case *proto.Expression_FieldReference_Expression:
		if root, err = ExprFromProto(rt.Expression, baseSchema, reg); err != nil {
			return nil, err
		}
	case *proto.Expression_FieldReference_OuterReference_:
		root = expr.OuterReference(rt.OuterReference.StepsOut)
	case *proto.Expression_FieldReference_RootReference_:
		root = expr.RootReference
	case *proto.Expression_FieldReference_LambdaParameterReference_:
		root = expr.LambdaParameterReference{StepsOut: rt.LambdaParameterReference.StepsOut}
	}

	switch rt := p.ReferenceType.(type) {
	case *proto.Expression_FieldReference_DirectReference:
		refseg := RefSegmentFromProto(rt.DirectReference)
		if root == expr.RootReference && baseSchema != nil {
			baseType := baseSchema.AsStructType()
			knownType, err = refseg.GetType(baseType)
			if err != nil {
				return nil, err
			}
		} else if rootExpr, ok := root.(expr.Expression); ok {
			knownType, err = refseg.GetType(rootExpr.GetType())
			if err != nil {
				return nil, err
			}
		}

		ref = refseg

	case *proto.Expression_FieldReference_MaskedReference:
		ref = MaskExpressionFromProto(rt.MaskedReference)
	}

	return expr.NewFieldReference(root, ref, knownType), nil
}
