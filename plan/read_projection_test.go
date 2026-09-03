// SPDX-License-Identifier: Apache-2.0

package plan

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

func projectionStruct(fields ...int32) *proto.Expression_MaskExpression_StructSelect {
	items := make([]*proto.Expression_MaskExpression_StructItem, len(fields))
	for i, field := range fields {
		items[i] = &proto.Expression_MaskExpression_StructItem{Field: field}
	}
	return &proto.Expression_MaskExpression_StructSelect{StructItems: items}
}

func projectionRead(schema types.NamedStruct, projection *proto.Expression_MaskExpression) *proto.Rel {
	return &proto.Rel{RelType: &proto.Rel_Read{Read: &proto.ReadRel{
		Common:     &proto.RelCommon{EmitKind: &proto.RelCommon_Direct_{Direct: &proto.RelCommon_Direct{}}},
		BaseSchema: schema.ToProto(),
		Projection: projection,
		ReadType:   &proto.ReadRel_NamedTable_{NamedTable: &proto.ReadRel_NamedTable{Names: []string{"table"}}},
	}}}
}

func projectionField(field int32) *proto.Expression {
	return &proto.Expression{RexType: &proto.Expression_Selection{Selection: &proto.Expression_FieldReference{
		RootType: &proto.Expression_FieldReference_RootReference_{RootReference: &proto.Expression_FieldReference_RootReference{}},
		ReferenceType: &proto.Expression_FieldReference_DirectReference{DirectReference: &proto.Expression_ReferenceSegment{
			ReferenceType: &proto.Expression_ReferenceSegment_StructField_{StructField: &proto.Expression_ReferenceSegment_StructField{Field: field}},
		}},
	}}}
}

func TestReadProjectionPlanRoundTrip(t *testing.T) {
	schema := types.NamedStruct{
		Names: []string{"id", "included", "label"},
		Struct: types.StructType{Nullability: types.NullabilityRequired, Types: []types.Type{
			&types.Int64Type{Nullability: types.NullabilityRequired},
			&types.BooleanType{Nullability: types.NullabilityRequired},
			&types.StringType{Nullability: types.NullabilityNullable},
		}},
	}
	read := projectionRead(schema, &proto.Expression_MaskExpression{Select: projectionStruct(2), MaintainSingularStruct: true})
	// Read filters still reference the base schema, including fields removed by projection.
	read.GetRead().Filter = projectionField(1)
	read.GetRead().BestEffortFilter = projectionField(1)
	project := &proto.Rel{RelType: &proto.Rel_Project{Project: &proto.ProjectRel{
		Common:      &proto.RelCommon{EmitKind: &proto.RelCommon_Direct_{Direct: &proto.RelCommon_Direct{}}},
		Input:       read,
		Expressions: []*proto.Expression{projectionField(0)},
	}}}
	original := &proto.Plan{Version: &proto.Version{MinorNumber: 85}, Relations: []*proto.PlanRel{{
		RelType: &proto.PlanRel_Root{Root: &proto.RelRoot{Input: project, Names: []string{"label", "copy"}}},
	}}}
	wire, err := protobuf.Marshal(original)
	require.NoError(t, err)
	decoded := &proto.Plan{}
	require.NoError(t, protobuf.Unmarshal(wire, decoded))
	p, err := FromProto(decoded, extensions.GetDefaultCollectionWithNoError())
	require.NoError(t, err)
	parent := p.GetRoots()[0].Input().(*ProjectRel)
	assert.Equal(t, schema.Struct.Types[2], parent.Expressions()[0].GetType())
	assert.Equal(t, []types.Type{schema.Struct.Types[2], schema.Struct.Types[2]}, parent.RecordType().Types())
	scan := parent.Input().(*NamedTableReadRel)
	assert.Equal(t, schema, scan.BaseSchema())
	assert.Equal(t, schema.Struct.Types[1], scan.Filter().GetType())
	assert.Equal(t, schema.Struct.Types[1], scan.BestEffortFilter().GetType())
	roundTripped, err := p.ToProto()
	require.NoError(t, err)
	assert.True(t, protobuf.Equal(original, roundTripped), "original: %s\nround trip: %s", original, roundTripped)
}

func TestReadProjectionBeforeEmit(t *testing.T) {
	schema := types.NamedStruct{Struct: types.StructType{Types: []types.Type{&types.Int64Type{}, &types.StringType{}, &types.BooleanType{}}}}
	read := projectionRead(schema, &proto.Expression_MaskExpression{Select: projectionStruct(1, 2), MaintainSingularStruct: true})
	read.GetRead().Common = &proto.RelCommon{EmitKind: &proto.RelCommon_Emit_{Emit: &proto.RelCommon_Emit{OutputMapping: []int32{1, 0, 1}}}}
	r, err := RelFromProto(read, expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError()))
	require.NoError(t, err)
	assert.Equal(t, []types.Type{schema.Struct.Types[2], schema.Struct.Types[1], schema.Struct.Types[2]}, r.RecordType().Types())
	r, err = r.Remap(1)
	require.NoError(t, err)
	assert.Equal(t, []types.Type{schema.Struct.Types[1]}, r.RecordType().Types())
}

func TestReadProjectionSetAndClear(t *testing.T) {
	schema := types.NamedStruct{Struct: types.StructType{Types: []types.Type{&types.Int64Type{}, &types.StringType{}}}}
	r := NewBuilderDefault().NamedScan([]string{"table"}, schema)
	r.SetProjection(expr.MaskExpressionFromProto(&proto.Expression_MaskExpression{Select: projectionStruct(1), MaintainSingularStruct: true}))
	assert.Equal(t, []types.Type{schema.Struct.Types[1]}, r.RecordType().Types())
	r.SetProjection(expr.MaskExpressionFromProto(&proto.Expression_MaskExpression{Select: projectionStruct()}))
	assert.Zero(t, r.RecordType().FieldCount())
	r.SetProjection(nil)
	assert.Equal(t, schema.Struct.Types, r.RecordType().Types())
}

func TestReadProjectionNestedMasks(t *testing.T) {
	integer := &types.Int64Type{Nullability: types.NullabilityRequired}
	str := &types.StringType{Nullability: types.NullabilityRequired, TypeVariationRef: 7}
	boolean := &types.BooleanType{Nullability: types.NullabilityNullable}
	inner := &types.StructType{Nullability: types.NullabilityNullable, TypeVariationRef: 11, Types: []types.Type{integer, str, boolean}}
	projected := *inner
	projected.Types = []types.Type{str, boolean}
	list := &types.ListType{Nullability: types.NullabilityNullable, TypeVariationRef: 12, Type: inner}
	mapType := &types.MapType{Nullability: types.NullabilityRequired, TypeVariationRef: 13, Key: str, Value: inner}
	structMask := &proto.Expression_MaskExpression_Select{Type: &proto.Expression_MaskExpression_Select_Struct{Struct: projectionStruct(1, 2)}}
	listSelection := []*proto.Expression_MaskExpression_ListSelect_ListSelectItem{{Type: &proto.Expression_MaskExpression_ListSelect_ListSelectItem_Slice{
		Slice: &proto.Expression_MaskExpression_ListSelect_ListSelectItem_ListSlice{Start: 0, End: 5},
	}}}
	for _, tc := range []struct {
		name        string
		input, want types.Type
		mask        *proto.Expression_MaskExpression_Select
	}{
		{"struct", inner, &projected, structMask},
		{"list child", list, &types.ListType{Nullability: list.Nullability, TypeVariationRef: list.TypeVariationRef, Type: &projected},
			&proto.Expression_MaskExpression_Select{Type: &proto.Expression_MaskExpression_Select_List{List: &proto.Expression_MaskExpression_ListSelect{Selection: listSelection, Child: structMask}}}},
		{"map child", mapType, &types.MapType{Nullability: mapType.Nullability, TypeVariationRef: mapType.TypeVariationRef, Key: str, Value: &projected},
			&proto.Expression_MaskExpression_Select{Type: &proto.Expression_MaskExpression_Select_Map{Map: &proto.Expression_MaskExpression_MapSelect{Child: structMask, Select: &proto.Expression_MaskExpression_MapSelect_Expression{Expression: &proto.Expression_MaskExpression_MapSelect_MapKeyExpression{MapKeyExpression: "a.*"}}}}}},
		{"list without child", list, list,
			&proto.Expression_MaskExpression_Select{Type: &proto.Expression_MaskExpression_Select_List{List: &proto.Expression_MaskExpression_ListSelect{Selection: listSelection}}}},
		{"map without child", mapType, mapType,
			&proto.Expression_MaskExpression_Select{Type: &proto.Expression_MaskExpression_Select_Map{Map: &proto.Expression_MaskExpression_MapSelect{Select: &proto.Expression_MaskExpression_MapSelect_Key{Key: &proto.Expression_MaskExpression_MapSelect_MapKey{MapKey: "a"}}}}}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			schema := types.NamedStruct{Struct: types.StructType{Types: []types.Type{integer, tc.input}}}
			mask := &proto.Expression_MaskExpression{MaintainSingularStruct: true, Select: &proto.Expression_MaskExpression_StructSelect{StructItems: []*proto.Expression_MaskExpression_StructItem{{Field: 1, Child: tc.mask}}}}
			original := projectionRead(schema, mask)
			saved := protobuf.Clone(original)
			r, err := RelFromProto(original, expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError()))
			require.NoError(t, err)
			assert.Equal(t, []types.Type{tc.want}, r.RecordType().Types())
			assert.True(t, protobuf.Equal(saved, r.ToProto()), "read, base schema and mask must round trip unchanged")
			assert.Equal(t, schema, r.(ReadRel).BaseSchema())
		})
	}
}

func TestReadProjectionSingularStruct(t *testing.T) {
	leaf := &types.StringType{Nullability: types.NullabilityRequired, TypeVariationRef: 7}
	inner := &types.StructType{Nullability: types.NullabilityNullable, TypeVariationRef: 11, Types: []types.Type{&types.Int64Type{}, leaf}}
	for _, maintain := range []bool{false, true} {
		t.Run(map[bool]string{false: "default", true: "maintain"}[maintain], func(t *testing.T) {
			schema := types.NamedStruct{Struct: types.StructType{Types: []types.Type{&types.Int64Type{}, inner}}}
			mask := &proto.Expression_MaskExpression{MaintainSingularStruct: maintain, Select: &proto.Expression_MaskExpression_StructSelect{StructItems: []*proto.Expression_MaskExpression_StructItem{{Field: 0}, {Field: 1, Child: &proto.Expression_MaskExpression_Select{Type: &proto.Expression_MaskExpression_Select_Struct{Struct: projectionStruct(1)}}}}}}
			r, err := RelFromProto(projectionRead(schema, mask), expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError()))
			require.NoError(t, err)
			want := &types.StructType{Nullability: inner.Nullability, TypeVariationRef: inner.TypeVariationRef, Types: []types.Type{leaf}}
			assert.Equal(t, []types.Type{schema.Struct.Types[0], want}, r.RecordType().Types())
			assert.Equal(t, schema, r.(ReadRel).BaseSchema())
		})
	}
}

func TestReadProjectionInvalidField(t *testing.T) {
	schema := types.NamedStruct{Struct: types.StructType{Types: []types.Type{&types.Int64Type{}}}}
	for _, field := range []int32{-1, 1} {
		_, err := RelFromProto(projectionRead(schema, &proto.Expression_MaskExpression{Select: projectionStruct(field)}), expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError()))
		assert.ErrorIs(t, err, substraitgo.ErrInvalidRel)
	}
}

func TestReadProjectionInvalidNestedSelection(t *testing.T) {
	inner := &types.StructType{Types: []types.Type{&types.Int64Type{}}}
	invalidField := &proto.Expression_MaskExpression_Select{Type: &proto.Expression_MaskExpression_Select_Struct{Struct: projectionStruct(1)}}
	for _, tc := range []struct {
		name      string
		input     types.Type
		selection *proto.Expression_MaskExpression_Select
		message   string
	}{
		{"struct field", inner, invalidField, "field 1 out of range"},
		{"list element field", &types.ListType{Type: inner}, &proto.Expression_MaskExpression_Select{Type: &proto.Expression_MaskExpression_Select_List{List: &proto.Expression_MaskExpression_ListSelect{Child: invalidField}}}, "field 1 out of range"},
		{"map value field", &types.MapType{Key: &types.StringType{}, Value: inner}, &proto.Expression_MaskExpression_Select{Type: &proto.Expression_MaskExpression_Select_Map{Map: &proto.Expression_MaskExpression_MapSelect{
			Child: invalidField, Select: &proto.Expression_MaskExpression_MapSelect_Key{Key: &proto.Expression_MaskExpression_MapSelect_MapKey{MapKey: "name"}},
		}}}, "field 1 out of range"},
		{"struct mask on scalar", &types.Int64Type{}, invalidField, "cannot select from i64"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			schema := types.NamedStruct{Struct: types.StructType{Types: []types.Type{tc.input}}}
			projection := &proto.Expression_MaskExpression{MaintainSingularStruct: true, Select: &proto.Expression_MaskExpression_StructSelect{StructItems: []*proto.Expression_MaskExpression_StructItem{{Field: 0, Child: tc.selection}}}}
			r, err := RelFromProto(projectionRead(schema, projection), expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError()))
			assert.Nil(t, r)
			assert.ErrorIs(t, err, substraitgo.ErrInvalidRel)
			assert.ErrorContains(t, err, tc.message)
		})
	}
}
