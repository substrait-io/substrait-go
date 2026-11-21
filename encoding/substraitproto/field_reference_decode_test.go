// SPDX-License-Identifier: Apache-2.0

package substraitproto

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	substraitgo "github.com/substrait-io/substrait-go/v9"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	protobuf "google.golang.org/protobuf/proto"
)

func maskWithChild(child *proto.Expression_MaskExpression_Select) *proto.Expression_MaskExpression {
	return &proto.Expression_MaskExpression{
		MaintainSingularStruct: true,
		Select: &proto.Expression_MaskExpression_StructSelect{
			StructItems: []*proto.Expression_MaskExpression_StructItem{{Child: child}},
		},
	}
}

func readWithMask(mask *proto.Expression_MaskExpression) *proto.Rel {
	return &proto.Rel{RelType: &proto.Rel_Read{Read: &proto.ReadRel{
		Projection: mask,
		ReadType:   &proto.ReadRel_NamedTable_{NamedTable: &proto.ReadRel_NamedTable{Names: []string{"t"}}},
	}}}
}

func TestDecodeMalformedMasks(t *testing.T) {
	listMask := func(items ...*proto.Expression_MaskExpression_ListSelect_ListSelectItem) *proto.Expression_MaskExpression {
		return maskWithChild(&proto.Expression_MaskExpression_Select{Type: &proto.Expression_MaskExpression_Select_List{
			List: &proto.Expression_MaskExpression_ListSelect{Selection: items},
		}})
	}
	mapMask := func(sel *proto.Expression_MaskExpression_MapSelect) *proto.Expression_MaskExpression {
		return maskWithChild(&proto.Expression_MaskExpression_Select{Type: &proto.Expression_MaskExpression_Select_Map{Map: sel}})
	}
	tests := []struct {
		name string
		mask *proto.Expression_MaskExpression
		want string
	}{
		{"missing root selection", &proto.Expression_MaskExpression{}, "requires a root struct selection"},
		{"nil struct item", &proto.Expression_MaskExpression{Select: &proto.Expression_MaskExpression_StructSelect{
			StructItems: []*proto.Expression_MaskExpression_StructItem{nil},
		}}, "struct item 0 is missing"},
		{"missing child type", maskWithChild(&proto.Expression_MaskExpression_Select{}), "mask selection requires a type"},
		{"nil child oneof", maskWithChild(&proto.Expression_MaskExpression_Select{
			Type: (*proto.Expression_MaskExpression_Select_Struct)(nil),
		}), "struct selection is missing"},
		{"nil struct selection", maskWithChild(&proto.Expression_MaskExpression_Select{
			Type: &proto.Expression_MaskExpression_Select_Struct{},
		}), "struct selection is missing"},
		{"nil list selection", maskWithChild(&proto.Expression_MaskExpression_Select{
			Type: &proto.Expression_MaskExpression_Select_List{},
		}), "list selection is missing"},
		{"nil map selection", mapMask(nil), "map selection is missing"},
		{"nil list item", listMask(nil), "list selection item 0 is missing"},
		{"missing list item type", listMask(&proto.Expression_MaskExpression_ListSelect_ListSelectItem{}), "list selection item requires a type"},
		{"nil list element", listMask(&proto.Expression_MaskExpression_ListSelect_ListSelectItem{
			Type: &proto.Expression_MaskExpression_ListSelect_ListSelectItem_Item{},
		}), "list element is missing"},
		{"nil list slice", listMask(&proto.Expression_MaskExpression_ListSelect_ListSelectItem{
			Type: &proto.Expression_MaskExpression_ListSelect_ListSelectItem_Slice{},
		}), "list slice is missing"},
		{"nil map key", mapMask(&proto.Expression_MaskExpression_MapSelect{
			Select: &proto.Expression_MaskExpression_MapSelect_Key{},
		}), "map key is missing"},
		{"nil map key expression", mapMask(&proto.Expression_MaskExpression_MapSelect{
			Select: &proto.Expression_MaskExpression_MapSelect_Expression{},
		}), "map key expression is missing"},
		{"missing nested list child type", maskWithChild(&proto.Expression_MaskExpression_Select{
			Type: &proto.Expression_MaskExpression_Select_List{List: &proto.Expression_MaskExpression_ListSelect{
				Child: &proto.Expression_MaskExpression_Select{},
			}},
		}), "mask selection requires a type"},
		{"missing nested struct child type", maskWithChild(&proto.Expression_MaskExpression_Select{
			Type: &proto.Expression_MaskExpression_Select_Struct{Struct: &proto.Expression_MaskExpression_StructSelect{
				StructItems: []*proto.Expression_MaskExpression_StructItem{{Child: &proto.Expression_MaskExpression_Select{}}},
			}},
		}), "mask selection requires a type"},
		{"missing nested map child type", mapMask(&proto.Expression_MaskExpression_MapSelect{
			Select: &proto.Expression_MaskExpression_MapSelect_Key{Key: &proto.Expression_MaskExpression_MapSelect_MapKey{MapKey: "k"}},
			Child:  &proto.Expression_MaskExpression_Select{},
		}), "mask selection requires a type"},
	}

	c := extensions.GetDefaultCollectionWithNoError()
	reg := expr.NewEmptyExtensionRegistry(c)
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rel := readWithMask(tc.mask)
			out, err := RelFromProto(rel, reg)
			require.ErrorIs(t, err, substraitgo.ErrInvalidExpr)
			assert.ErrorContains(t, err, tc.want)
			assert.Nil(t, out)

			outPlan, err := PlanFromProto(&proto.Plan{Relations: []*proto.PlanRel{{
				RelType: &proto.PlanRel_Rel{Rel: rel},
			}}}, c)
			require.ErrorIs(t, err, substraitgo.ErrInvalidExpr)
			assert.ErrorContains(t, err, tc.want)
			assert.Nil(t, outPlan)

			outRef, err := ExprFromProto(&proto.Expression{RexType: &proto.Expression_Selection{
				Selection: &proto.Expression_FieldReference{
					ReferenceType: &proto.Expression_FieldReference_MaskedReference{MaskedReference: tc.mask},
				},
			}}, nil, reg)
			require.ErrorIs(t, err, substraitgo.ErrInvalidExpr)
			assert.ErrorContains(t, err, tc.want)
			assert.Nil(t, outRef)
		})
	}
}

func TestDecodeMissingMaskedReference(t *testing.T) {
	reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
	for _, ref := range []*proto.Expression_FieldReference{
		{ReferenceType: &proto.Expression_FieldReference_MaskedReference{}},
		{ReferenceType: (*proto.Expression_FieldReference_MaskedReference)(nil)},
	} {
		_, err := FieldReferenceFromProto(ref, nil, reg)
		require.ErrorIs(t, err, substraitgo.ErrInvalidExpr)
	}
}

func TestMaskOptionalChildrenRoundTrip(t *testing.T) {
	i32 := &types.Int32Type{Nullability: types.NullabilityRequired}
	mapType := &types.MapType{Nullability: types.NullabilityRequired, Key: &types.StringType{Nullability: types.NullabilityRequired}, Value: i32}
	tests := []struct {
		name string
		mask *proto.Expression_MaskExpression
		typ  types.Type
	}{
		{"empty root", &proto.Expression_MaskExpression{Select: &proto.Expression_MaskExpression_StructSelect{}}, i32},
		{"struct", maskWithChild(nil), i32},
		{"list", maskWithChild(&proto.Expression_MaskExpression_Select{Type: &proto.Expression_MaskExpression_Select_List{
			List: &proto.Expression_MaskExpression_ListSelect{Selection: []*proto.Expression_MaskExpression_ListSelect_ListSelectItem{{
				Type: &proto.Expression_MaskExpression_ListSelect_ListSelectItem_Item{
					Item: &proto.Expression_MaskExpression_ListSelect_ListSelectItem_ListElement{},
				},
			}}},
		}}), &types.ListType{Nullability: types.NullabilityRequired, Type: i32}},
		{"map empty string key", maskWithChild(&proto.Expression_MaskExpression_Select{Type: &proto.Expression_MaskExpression_Select_Map{
			Map: &proto.Expression_MaskExpression_MapSelect{Select: &proto.Expression_MaskExpression_MapSelect_Key{
				Key: &proto.Expression_MaskExpression_MapSelect_MapKey{},
			}},
		}}), mapType},
		{"map key expression", maskWithChild(&proto.Expression_MaskExpression_Select{Type: &proto.Expression_MaskExpression_Select_Map{
			Map: &proto.Expression_MaskExpression_MapSelect{Select: &proto.Expression_MaskExpression_MapSelect_Expression{
				Expression: &proto.Expression_MaskExpression_MapSelect_MapKeyExpression{MapKeyExpression: "a.*"},
			}},
		}}), mapType},
		{"map all keys", maskWithChild(&proto.Expression_MaskExpression_Select{Type: &proto.Expression_MaskExpression_Select_Map{
			Map: &proto.Expression_MaskExpression_MapSelect{},
		}}), mapType},
		{"map all keys with child", maskWithChild(&proto.Expression_MaskExpression_Select{Type: &proto.Expression_MaskExpression_Select_Map{
			Map: &proto.Expression_MaskExpression_MapSelect{Child: &proto.Expression_MaskExpression_Select{
				Type: &proto.Expression_MaskExpression_Select_Struct{Struct: &proto.Expression_MaskExpression_StructSelect{
					StructItems: []*proto.Expression_MaskExpression_StructItem{{Field: 1}},
				}},
			}},
		}}), &types.MapType{Nullability: types.NullabilityRequired, Key: mapType.Key, Value: &types.StructType{
			Nullability: types.NullabilityRequired, Types: []types.Type{i32, i32},
		}}},
	}
	c := extensions.GetDefaultCollectionWithNoError()
	reg := expr.NewEmptyExtensionRegistry(c)
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := MaskExpressionToProto(MaskExpressionFromProto(tc.mask))
			require.True(t, protobuf.Equal(tc.mask, got), "mask differs after round trip: %v", got)
			rel := readWithMask(tc.mask)
			rel.GetRead().BaseSchema = &proto.NamedStruct{
				Names: []string{"col"},
				Struct: &proto.Type_Struct{
					Nullability: proto.Type_NULLABILITY_REQUIRED, Types: []*proto.Type{TypeToProto(tc.typ)},
				},
			}
			if tc.name == "map all keys with child" {
				rel.GetRead().BaseSchema.Names = []string{"col", "x", "y"}
			}
			out, err := RelFromProto(rel, reg)
			require.NoError(t, err)
			assert.True(t, protobuf.Equal(tc.mask, RelToProto(out).GetRead().Projection))

			outPlan, err := PlanFromProto(&proto.Plan{Relations: []*proto.PlanRel{{RelType: &proto.PlanRel_Rel{Rel: rel}}}}, c)
			require.NoError(t, err)
			gotPlan, err := PlanToProto(outPlan)
			require.NoError(t, err)
			assert.True(t, protobuf.Equal(tc.mask, gotPlan.Relations[0].GetRel().GetRead().Projection))

			selection := &proto.Expression{RexType: &proto.Expression_Selection{Selection: &proto.Expression_FieldReference{
				RootType:      &proto.Expression_FieldReference_RootReference_{RootReference: &proto.Expression_FieldReference_RootReference{}},
				ReferenceType: &proto.Expression_FieldReference_MaskedReference{MaskedReference: tc.mask},
			}}}
			outExpr, err := ExprFromProto(selection, nil, reg)
			require.NoError(t, err)
			assert.True(t, protobuf.Equal(selection, ExprToProto(outExpr)))
		})
	}
}
