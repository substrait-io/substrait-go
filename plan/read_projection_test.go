// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	substraitgo "github.com/substrait-io/substrait-go/v9"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/plan"
	"github.com/substrait-io/substrait-go/v9/types"
)

func readProjection(fields ...int32) *expr.MaskExpression {
	items := make(expr.MaskStructSelect, len(fields))
	for i, field := range fields {
		items[i] = expr.NewMaskStructItem(field, nil)
	}
	return expr.NewMaskExpression(items, true)
}

func TestReadProjectionPreservesBaseSchemaAndFilters(t *testing.T) {
	b := plan.NewBuilderDefault()
	schema := types.NamedStruct{
		Names: []string{"id", "included", "label"},
		Struct: types.StructType{Nullability: types.NullabilityRequired, Types: []types.Type{
			&types.Int64Type{Nullability: types.NullabilityRequired},
			&types.BooleanType{Nullability: types.NullabilityRequired},
			&types.StringType{Nullability: types.NullabilityNullable},
		}},
	}
	filter, err := expr.NewRootFieldRef(&expr.StructFieldRef{Field: 1}, types.NewRecordTypeFromStruct(schema.Struct))
	require.NoError(t, err)
	scan := plan.NewNamedTableReadRel(plan.NewBaseReadRel(plan.RelCommon{}, schema, filter, filter, nil, nil), []string{"table"}, nil)
	require.NoError(t, scan.SetProjection(readProjection(2)))
	field, err := b.RootFieldRef(scan, 0)
	require.NoError(t, err)
	parent, err := b.Project(scan, field)
	require.NoError(t, err)
	assert.Equal(t, schema.Struct.Types[2], field.GetType())
	assert.Equal(t, []types.Type{schema.Struct.Types[2], schema.Struct.Types[2]}, parent.RecordType().Types())
	assert.Equal(t, schema, scan.BaseSchema())
	assert.Same(t, filter, scan.Filter())
	assert.Same(t, filter, scan.BestEffortFilter())
	assert.Equal(t, schema.Struct.Types[1], scan.Filter().GetType())
	_, err = b.Plan(parent, []string{"label", "copy"})
	require.NoError(t, err)
}

func TestReadProjectionBeforeEmit(t *testing.T) {
	schema := types.NamedStruct{Struct: types.StructType{Types: []types.Type{&types.Int64Type{}, &types.StringType{}, &types.BooleanType{}}}}
	scan := plan.NewBuilderDefault().NamedScan([]string{"table"}, schema)
	require.NoError(t, scan.SetProjection(readProjection(1, 2)))
	remapped, err := scan.Remap(1, 0, 1)
	require.NoError(t, err)
	assert.Equal(t, []types.Type{schema.Struct.Types[2], schema.Struct.Types[1], schema.Struct.Types[2]}, remapped.RecordType().Types())
	remapped, err = remapped.Remap(1)
	require.NoError(t, err)
	assert.Equal(t, []types.Type{schema.Struct.Types[1]}, remapped.RecordType().Types())
	for _, index := range []int32{-1, 2} {
		_, err := scan.Remap(index)
		assert.ErrorIs(t, err, substraitgo.ErrInvalidRel)
	}
}

func TestReadProjectionSetAndClear(t *testing.T) {
	schema := types.NamedStruct{Struct: types.StructType{Types: []types.Type{&types.Int64Type{}, &types.StringType{}}}}
	scan := plan.NewBuilderDefault().NamedScan([]string{"table"}, schema)
	require.NoError(t, scan.SetProjection(readProjection(1)))
	assert.Equal(t, []types.Type{schema.Struct.Types[1]}, scan.RecordType().Types())
	require.NoError(t, scan.SetProjection(readProjection()))
	assert.Zero(t, scan.RecordType().FieldCount())
	require.NoError(t, scan.SetProjection(expr.NewMaskExpression(nil, false)))
	assert.Zero(t, scan.RecordType().FieldCount())
	require.NoError(t, scan.SetProjection(nil))
	assert.Equal(t, schema.Struct.Types, scan.RecordType().Types())
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
	originalInner, originalList, originalMap := *inner, *list, *mapType
	structMask := readProjection(1, 2).Select()
	listSelection := []expr.MaskListSelectItem{&expr.MaskListSlice{Start: 0, End: 5}}
	for _, tc := range []struct {
		name        string
		input, want types.Type
		mask        expr.MaskSelect
	}{
		{"struct", inner, &projected, structMask},
		{"list child", list, &types.ListType{Nullability: list.Nullability, TypeVariationRef: list.TypeVariationRef, Type: &projected},
			expr.NewMaskListSelect(listSelection, structMask)},
		{"map child", mapType, &types.MapType{Nullability: mapType.Nullability, TypeVariationRef: mapType.TypeVariationRef, Key: str, Value: &projected},
			expr.NewMaskMapSelect(expr.MapSelectExpr, "a.*", structMask)},
		{"map all keys with child", mapType, &types.MapType{Nullability: mapType.Nullability, TypeVariationRef: mapType.TypeVariationRef, Key: str, Value: &projected},
			expr.NewMaskMapSelect(expr.MapSelectAll, "", structMask)},
		{"list without child", list, list, expr.NewMaskListSelect(listSelection, nil)},
		{"map without child", mapType, mapType, expr.NewMaskMapSelect(expr.MapSelectKey, "a", nil)},
		{"map all keys without child", mapType, mapType, expr.NewMaskMapSelect(expr.MapSelectAll, "", nil)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			schema := types.NamedStruct{Struct: types.StructType{Types: []types.Type{integer, tc.input}}}
			scan := plan.NewBuilderDefault().NamedScan([]string{"table"}, schema)
			mask := expr.NewMaskExpression(expr.MaskStructSelect{expr.NewMaskStructItem(1, tc.mask)}, true)
			require.NoError(t, scan.SetProjection(mask))
			for i := 0; i < 2; i++ {
				assert.Equal(t, []types.Type{tc.want}, scan.RecordType().Types())
				assert.Equal(t, schema, scan.BaseSchema())
				assert.Equal(t, originalInner, *inner)
				assert.Equal(t, originalList, *list)
				assert.Equal(t, originalMap, *mapType)
			}
		})
	}
}

func TestReadProjectionSingularStruct(t *testing.T) {
	leaf := &types.StringType{Nullability: types.NullabilityRequired, TypeVariationRef: 7}
	inner := &types.StructType{Nullability: types.NullabilityNullable, TypeVariationRef: 11, Types: []types.Type{&types.Int64Type{}, leaf}}
	for _, maintain := range []bool{false, true} {
		t.Run(map[bool]string{false: "default", true: "maintain"}[maintain], func(t *testing.T) {
			schema := types.NamedStruct{Struct: types.StructType{Types: []types.Type{&types.Int64Type{}, inner}}}
			scan := plan.NewBuilderDefault().NamedScan([]string{"table"}, schema)
			mask := expr.NewMaskExpression(expr.MaskStructSelect{
				expr.NewMaskStructItem(0, nil),
				expr.NewMaskStructItem(1, readProjection(1).Select()),
			}, maintain)
			require.NoError(t, scan.SetProjection(mask))
			want := &types.StructType{Nullability: inner.Nullability, TypeVariationRef: inner.TypeVariationRef, Types: []types.Type{leaf}}
			assert.Equal(t, []types.Type{schema.Struct.Types[0], want}, scan.RecordType().Types())
		})
	}
}

func TestReadProjectionRejectsInvalidSelection(t *testing.T) {
	inner := &types.StructType{Types: []types.Type{&types.Int64Type{}}}
	invalidField := readProjection(1).Select()
	for _, tc := range []struct {
		name      string
		input     types.Type
		selection expr.MaskSelect
		message   string
	}{
		{"struct field", inner, invalidField, "field 1 out of range"},
		{"list element field", &types.ListType{Type: inner}, expr.NewMaskListSelect(nil, invalidField), "field 1 out of range"},
		{"map value field", &types.MapType{Key: &types.StringType{}, Value: inner}, expr.NewMaskMapSelect(expr.MapSelectKey, "name", invalidField), "field 1 out of range"},
		{"struct mask on scalar", &types.Int64Type{}, invalidField, "cannot select from i64"},
		{"nil list selection", &types.ListType{Type: &types.Int64Type{}}, (*expr.MaskListSelect)(nil), "cannot select from"},
		{"nil map selection", &types.MapType{Key: &types.StringType{}, Value: &types.Int64Type{}}, (*expr.MaskMapSelect)(nil), "cannot select from"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			schema := types.NamedStruct{Struct: types.StructType{Types: []types.Type{tc.input}}}
			scan := plan.NewBuilderDefault().NamedScan([]string{"table"}, schema)
			previous := readProjection(0)
			require.NoError(t, scan.SetProjection(previous))
			projection := expr.NewMaskExpression(expr.MaskStructSelect{expr.NewMaskStructItem(0, tc.selection)}, true)
			err := scan.SetProjection(projection)
			assert.ErrorIs(t, err, substraitgo.ErrInvalidRel)
			assert.ErrorContains(t, err, tc.message)
			for _, field := range []int32{-1, 1} {
				assert.ErrorIs(t, scan.SetProjection(readProjection(field)), substraitgo.ErrInvalidRel)
			}
			assert.Same(t, previous, scan.Projection())
			assert.NotPanics(t, func() { assert.Equal(t, schema.Struct.Types, scan.RecordType().Types()) })
		})
	}
}

func TestReadProjectionSetPreservesEmit(t *testing.T) {
	for _, tc := range []struct {
		name       string
		types      []types.Type
		fields     []int32
		projection *expr.MaskExpression
	}{
		{"shrink", []types.Type{&types.Int64Type{}, &types.StringType{}}, []int32{1, 0}, readProjection(1)},
		{"clear", []types.Type{&types.Int64Type{}}, []int32{0, 0}, nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			schema := types.NamedStruct{Struct: types.StructType{Types: tc.types}}
			scan := plan.NewBuilderDefault().NamedScan([]string{"table"}, schema)
			previous := readProjection(tc.fields...)
			require.NoError(t, scan.SetProjection(previous))
			_, err := scan.Remap(1)
			require.NoError(t, err)
			want := scan.RecordType()
			assert.ErrorIs(t, scan.SetProjection(tc.projection), substraitgo.ErrInvalidRel)
			assert.Same(t, previous, scan.Projection())
			assert.Equal(t, []int32{1}, scan.OutputMapping())
			assert.NotPanics(t, func() { assert.Equal(t, want, scan.RecordType()) })
		})
	}
}
