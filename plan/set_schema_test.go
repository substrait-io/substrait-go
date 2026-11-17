// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	substraitgo "github.com/substrait-io/substrait-go/v9"
	"github.com/substrait-io/substrait-go/v9/plan"
	"github.com/substrait-io/substrait-go/v9/types"
)

func setScan(fields ...types.Type) plan.Rel {
	names := make([]string, len(fields))
	for i := range names {
		names[i] = "field"
	}
	return plan.NewBuilderDefault().NamedScan([]string{"table"}, types.NamedStruct{
		Names:  names,
		Struct: types.StructType{Nullability: types.NullabilityRequired, Types: fields},
	})
}

func TestSetOutputNullability(t *testing.T) {
	// Spec v0.87.0's three-input example covers every required/nullable combination.
	inputs := make([]plan.Rel, 3)
	for input := range inputs {
		fields := make([]types.Type, 8)
		for column := range fields {
			n := types.NullabilityRequired
			if column&(1<<(2-input)) != 0 {
				n = types.NullabilityNullable
			}
			fields[column] = &types.Int64Type{Nullability: n}
		}
		inputs[input] = setScan(fields...)
	}
	for _, tc := range []struct {
		op            plan.SetOp
		firstNullable int
	}{
		{plan.SetOpMinusPrimary, 4},
		{plan.SetOpMinusPrimaryAll, 4},
		{plan.SetOpMinusMultiset, 4},
		{plan.SetOpIntersectionPrimary, 5},
		{plan.SetOpIntersectionMultiset, 7},
		{plan.SetOpIntersectionMultisetAll, 7},
		{plan.SetOpUnionDistinct, 1},
		{plan.SetOpUnionAll, 1},
	} {
		t.Run(tc.op.String(), func(t *testing.T) {
			check := func(t *testing.T, rel *plan.SetRel) {
				want := make([]types.Type, 8)
				for i := range want {
					n := types.NullabilityRequired
					if i >= tc.firstNullable {
						n = types.NullabilityNullable
					}
					want[i] = &types.Int64Type{Nullability: n}
				}
				assert.Equal(t, want, rel.RecordType().Types())
				remapped, err := rel.Remap(7, 0, 4)
				require.NoError(t, err)
				assert.Equal(t, []types.Type{want[7], want[0], want[4]}, remapped.RecordType().Types())
			}
			t.Run("builder", func(t *testing.T) {
				rel, err := plan.NewBuilderDefault().Set(tc.op, inputs...)
				require.NoError(t, err)
				check(t, rel)
			})
			t.Run("constructor", func(t *testing.T) {
				check(t, plan.NewSetRel(inputs, tc.op, plan.RelCommon{}, nil))
			})
		})
	}
	for input, rel := range inputs {
		for column, field := range rel.RecordType().Types() {
			assert.Equal(t, column&(1<<(2-input)) != 0, field.GetNullability() == types.NullabilityNullable)
		}
	}
}

func TestSetInputTypeCompatibility(t *testing.T) {
	r := types.NullabilityRequired
	n := types.NullabilityNullable
	for _, tc := range []struct {
		name        string
		left, right []types.Type
	}{
		{"different kinds", []types.Type{&types.Int64Type{Nullability: r}}, []types.Type{&types.Int32Type{Nullability: n}}},
		{"different widths", []types.Type{&types.Int64Type{Nullability: r}}, nil},
		{"different variations", []types.Type{&types.Int64Type{Nullability: r, TypeVariationRef: 1}}, []types.Type{&types.Int64Type{Nullability: n, TypeVariationRef: 2}}},
		{"different decimal precision", []types.Type{&types.DecimalType{Precision: 10, Scale: 2, Nullability: r}}, []types.Type{&types.DecimalType{Precision: 11, Scale: 2, Nullability: n}}},
		{"different nested fields", []types.Type{&types.ListType{Type: &types.Int64Type{Nullability: r}, Nullability: r}}, []types.Type{&types.ListType{Type: &types.Int64Type{Nullability: n}, Nullability: n}}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, err := plan.NewBuilderDefault().Set(plan.SetOpUnionAll, setScan(tc.left...), setScan(tc.right...))
			assert.ErrorIs(t, err, substraitgo.ErrInvalidRel)
		})
	}
}

func TestSetPreservesNestedFields(t *testing.T) {
	child := &types.Int64Type{Nullability: types.NullabilityRequired}
	field := &types.StructType{Nullability: types.NullabilityRequired, Types: []types.Type{child}}
	rel, err := plan.NewBuilderDefault().Set(plan.SetOpUnionAll,
		setScan(field), setScan(field.WithNullability(types.NullabilityNullable)))
	require.NoError(t, err)
	assert.Equal(t, field.WithNullability(types.NullabilityNullable), rel.RecordType().Types()[0])
	assert.Equal(t, types.NullabilityRequired, field.Nullability)
	assert.Equal(t, types.NullabilityRequired, child.Nullability)
}
