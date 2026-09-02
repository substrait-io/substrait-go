// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/expr"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	pbproto "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestMaskListElementMatchesDescriptor(t *testing.T) {
	want := map[protoreflect.Name]struct {
		number protoreflect.FieldNumber
		kind   protoreflect.Kind
	}{
		"field": {1, protoreflect.Int32Kind},
	}

	fields := (&proto.Expression_MaskExpression_ListSelect_ListSelectItem_ListElement{}).ProtoReflect().Descriptor().Fields()
	require.Equal(t, len(want), fields.Len(), "spec mask list element field set changed")
	require.Equal(t, len(want), reflect.TypeOf(expr.MaskListElement{}).NumField(), "expr.MaskListElement field count drifted from the spec")

	for i := 0; i < fields.Len(); i++ {
		f := fields.Get(i)
		w, ok := want[f.Name()]
		require.Truef(t, ok, "unexpected spec field %q", f.Name())
		assert.EqualValues(t, w.number, f.Number(), "%s wire number", f.Name())
		assert.Equal(t, w.kind, f.Kind(), "%s kind", f.Name())
	}
}

func TestMaskListSliceMatchesDescriptor(t *testing.T) {
	want := map[protoreflect.Name]struct {
		number protoreflect.FieldNumber
		kind   protoreflect.Kind
	}{
		"start": {1, protoreflect.Int32Kind},
		"end":   {2, protoreflect.Int32Kind},
	}

	fields := (&proto.Expression_MaskExpression_ListSelect_ListSelectItem_ListSlice{}).ProtoReflect().Descriptor().Fields()
	require.Equal(t, len(want), fields.Len(), "spec mask list slice field set changed")
	require.Equal(t, len(want), reflect.TypeOf(expr.MaskListSlice{}).NumField(), "expr.MaskListSlice field count drifted from the spec")

	for i := 0; i < fields.Len(); i++ {
		f := fields.Get(i)
		w, ok := want[f.Name()]
		require.Truef(t, ok, "unexpected spec field %q", f.Name())
		assert.EqualValues(t, w.number, f.Number(), "%s wire number", f.Name())
		assert.Equal(t, w.kind, f.Kind(), "%s kind", f.Name())
	}
}

func TestMaskListSelectRoundTrip(t *testing.T) {
	// Direct accessors + per-item ToProto.
	elem := &expr.MaskListElement{Field: 2}
	assert.EqualValues(t, 2, elem.GetField())
	assert.EqualValues(t, 2, elem.ToProto().GetItem().GetField())

	slice := &expr.MaskListSlice{Start: 1, End: 3}
	start, end := slice.GetBounds()
	assert.EqualValues(t, 1, start)
	assert.EqualValues(t, 3, end)
	assert.EqualValues(t, 1, slice.ToProto().GetSlice().GetStart())
	assert.EqualValues(t, 3, slice.ToProto().GetSlice().GetEnd())

	// A list select (element + slice) reached through the public MaskExpression
	// entry, exercising maskSelectFromProto and the domain -> proto round-trip.
	pb := &proto.Expression_MaskExpression{
		Select: &proto.Expression_MaskExpression_StructSelect{
			StructItems: []*proto.Expression_MaskExpression_StructItem{{
				Field: 0,
				Child: &proto.Expression_MaskExpression_Select{
					Type: &proto.Expression_MaskExpression_Select_List{
						List: &proto.Expression_MaskExpression_ListSelect{
							Selection: []*proto.Expression_MaskExpression_ListSelect_ListSelectItem{
								{Type: &proto.Expression_MaskExpression_ListSelect_ListSelectItem_Item{
									Item: &proto.Expression_MaskExpression_ListSelect_ListSelectItem_ListElement{Field: 2},
								}},
								{Type: &proto.Expression_MaskExpression_ListSelect_ListSelectItem_Slice{
									Slice: &proto.Expression_MaskExpression_ListSelect_ListSelectItem_ListSlice{Start: 1, End: 3},
								}},
							},
							Child: &proto.Expression_MaskExpression_Select{
								Type: &proto.Expression_MaskExpression_Select_Struct{
									Struct: &proto.Expression_MaskExpression_StructSelect{},
								},
							},
						},
					},
				},
			}},
		},
	}

	got := expr.MaskExpressionFromProto(pb).ToProto()
	assert.Truef(t, pbproto.Equal(pb, got), "mask expression round-trip mismatch:\nwant: %v\ngot:  %v", pb, got)
}
