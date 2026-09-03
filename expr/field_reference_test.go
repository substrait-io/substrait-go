// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/expr"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	protobuf "google.golang.org/protobuf/proto"
)

func TestMaskExpressionOptionalCollectionChildRoundTrip(t *testing.T) {
	for _, tc := range []struct {
		name        string
		selectChild func(*proto.Expression_MaskExpression_Select) *proto.Expression_MaskExpression_Select
	}{
		{"list", func(child *proto.Expression_MaskExpression_Select) *proto.Expression_MaskExpression_Select {
			return &proto.Expression_MaskExpression_Select{Type: &proto.Expression_MaskExpression_Select_List{List: &proto.Expression_MaskExpression_ListSelect{
				Child: child,
				Selection: []*proto.Expression_MaskExpression_ListSelect_ListSelectItem{
					{Type: &proto.Expression_MaskExpression_ListSelect_ListSelectItem_Item{Item: &proto.Expression_MaskExpression_ListSelect_ListSelectItem_ListElement{Field: 1}}},
					{Type: &proto.Expression_MaskExpression_ListSelect_ListSelectItem_Slice{Slice: &proto.Expression_MaskExpression_ListSelect_ListSelectItem_ListSlice{Start: 3, End: 5}}},
				},
			}}}
		}},
		{"map key", func(child *proto.Expression_MaskExpression_Select) *proto.Expression_MaskExpression_Select {
			return &proto.Expression_MaskExpression_Select{Type: &proto.Expression_MaskExpression_Select_Map{Map: &proto.Expression_MaskExpression_MapSelect{
				Child: child, Select: &proto.Expression_MaskExpression_MapSelect_Key{Key: &proto.Expression_MaskExpression_MapSelect_MapKey{MapKey: "name"}},
			}}}
		}},
		{"map expression", func(child *proto.Expression_MaskExpression_Select) *proto.Expression_MaskExpression_Select {
			return &proto.Expression_MaskExpression_Select{Type: &proto.Expression_MaskExpression_Select_Map{Map: &proto.Expression_MaskExpression_MapSelect{
				Child: child, Select: &proto.Expression_MaskExpression_MapSelect_Expression{Expression: &proto.Expression_MaskExpression_MapSelect_MapKeyExpression{MapKeyExpression: "a.*"}},
			}}}
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			for _, withChild := range []bool{false, true} {
				t.Run(map[bool]string{false: "without child", true: "with child"}[withChild], func(t *testing.T) {
					var child *proto.Expression_MaskExpression_Select
					if withChild {
						child = &proto.Expression_MaskExpression_Select{Type: &proto.Expression_MaskExpression_Select_Struct{Struct: &proto.Expression_MaskExpression_StructSelect{
							StructItems: []*proto.Expression_MaskExpression_StructItem{{Field: 1}},
						}}}
					}
					original := &proto.Expression_MaskExpression{
						MaintainSingularStruct: true,
						Select: &proto.Expression_MaskExpression_StructSelect{StructItems: []*proto.Expression_MaskExpression_StructItem{
							{Field: 2, Child: tc.selectChild(child)},
						}},
					}
					wire, err := protobuf.Marshal(original)
					require.NoError(t, err)
					decoded := &proto.Expression_MaskExpression{}
					require.NoError(t, protobuf.Unmarshal(wire, decoded))
					mask := expr.MaskExpressionFromProto(decoded)
					selection := mask.Select()
					require.Len(t, selection, 1)
					collection := selection[0].Child().(interface{ Child() expr.MaskSelect })
					if withChild {
						assert.NotNil(t, collection.Child())
					} else {
						assert.Nil(t, collection.Child(), "an absent child selects the whole element or value")
					}
					assert.True(t, protobuf.Equal(original, mask.ToProto()), "mask must preserve its collection selection and optional child")
				})
			}
		})
	}
}
