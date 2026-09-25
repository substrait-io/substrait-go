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
	}
	panic("unimplemented mask select type")
}
