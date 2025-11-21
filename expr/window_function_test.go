// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v7/expr"
	"github.com/substrait-io/substrait-go/v7/extensions"
	"github.com/substrait-io/substrait-go/v7/types"
)

func TestWindowFunctionBoundsType(t *testing.T) {
	col := extensions.GetDefaultCollectionWithNoError()
	reg := expr.NewEmptyExtensionRegistry(col)

	sumID := extensions.ID{URN: extensions.SubstraitDefaultURNPrefix + "functions_arithmetic", Name: "sum"}

	schema := types.NewRecordTypeFromTypes([]types.Type{
		&types.Int64Type{Nullability: types.NullabilityRequired},
	})

	b := expr.ExprBuilder{
		Reg:        reg,
		BaseSchema: schema,
	}

	tests := []struct {
		name       string
		boundsType types.BoundsType
	}{
		{"ROWS", types.BoundsTypeRows},
		{"RANGE", types.BoundsTypeRange},
		{"UNSPECIFIED", types.BoundsTypeUnspecified},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := b.WindowFunc(sumID).
				Args(b.RootRef(expr.NewStructFieldRef(0))).
				Phase(types.AggPhaseInitialToResult).
				BoundsType(tt.boundsType).
				Bounds(expr.PrecedingBound(5), expr.FollowingBound(5))

			// RANGE requires exactly one sort field
			if tt.boundsType == types.BoundsTypeRange {
				builder = builder.Sort(expr.SortField{
					Expr: expr.MustExpr(b.RootRef(expr.NewStructFieldRef(0)).Build()),
					Kind: types.SortAscNullsFirst,
				})
			}

			wf, err := builder.BuildExpr()
			require.NoError(t, err)

			// Roundtrip: serialize and deserialize
			protoExpr := wf.ToProto()
			deserialized, err := expr.ExprFromProto(protoExpr, schema, reg)
			require.NoError(t, err)

			// Verify the entire expression survived the roundtrip
			assert.Truef(t, wf.Equals(deserialized), "expected: %s\ngot: %s", wf, deserialized)
		})
	}
}

func TestWindowFunctionBoundsTypeDefault(t *testing.T) {
	col := extensions.GetDefaultCollectionWithNoError()
	reg := expr.NewEmptyExtensionRegistry(col)

	sumID := extensions.ID{URN: extensions.SubstraitDefaultURNPrefix + "functions_arithmetic", Name: "sum"}

	schema := types.NewRecordTypeFromTypes([]types.Type{
		&types.Int64Type{Nullability: types.NullabilityRequired},
	})

	b := expr.ExprBuilder{
		Reg:        reg,
		BaseSchema: schema,
	}

	// Create a window function without explicitly setting bounds type
	wf, err := b.WindowFunc(sumID).
		Args(b.RootRef(expr.NewStructFieldRef(0))).
		Phase(types.AggPhaseInitialToResult).
		BuildExpr()

	require.NoError(t, err)
	require.NotNil(t, wf)

	windowFunc, ok := wf.(*expr.WindowFunction)
	require.True(t, ok)

	// Verify the bounds type defaults to UNSPECIFIED (0)
	assert.Equal(t, types.BoundsTypeUnspecified, windowFunc.BoundsType)
}

func TestWindowFunctionRANGERequiresSingleSort(t *testing.T) {
	col := extensions.GetDefaultCollectionWithNoError()
	reg := expr.NewEmptyExtensionRegistry(col)

	sumID := extensions.ID{URN: extensions.SubstraitDefaultURNPrefix + "functions_arithmetic", Name: "sum"}

	schema := types.NewRecordTypeFromTypes([]types.Type{
		&types.Int64Type{Nullability: types.NullabilityRequired},
		&types.Int64Type{Nullability: types.NullabilityRequired},
	})

	b := expr.ExprBuilder{
		Reg:        reg,
		BaseSchema: schema,
	}

	t.Run("RANGE with no sort fields should error", func(t *testing.T) {
		// Create a window function with RANGE but no sort fields
		_, err := b.WindowFunc(sumID).
			Args(b.RootRef(expr.NewStructFieldRef(0))).
			Phase(types.AggPhaseInitialToResult).
			BoundsType(types.BoundsTypeRange).
			BuildExpr()

		require.Error(t, err)
		assert.Contains(t, err.Error(), "RANGE bounds type requires exactly one sort field")
		assert.Contains(t, err.Error(), "got 0")
	})

	t.Run("RANGE with one sort field should succeed", func(t *testing.T) {
		// Create a window function with RANGE and exactly one sort field
		wf, err := b.WindowFunc(sumID).
			Args(b.RootRef(expr.NewStructFieldRef(0))).
			Phase(types.AggPhaseInitialToResult).
			BoundsType(types.BoundsTypeRange).
			Sort(expr.SortField{
				Expr: expr.MustExpr(b.RootRef(expr.NewStructFieldRef(0)).Build()),
				Kind: types.SortAscNullsFirst,
			}).
			BuildExpr()

		require.NoError(t, err)
		require.NotNil(t, wf)

		windowFunc, ok := wf.(*expr.WindowFunction)
		require.True(t, ok)
		assert.Equal(t, types.BoundsTypeRange, windowFunc.BoundsType)
		assert.Len(t, windowFunc.Sorts, 1)
	})

	t.Run("RANGE with two sort fields should error", func(t *testing.T) {
		// Create a window function with RANGE and two sort fields
		_, err := b.WindowFunc(sumID).
			Args(b.RootRef(expr.NewStructFieldRef(0))).
			Phase(types.AggPhaseInitialToResult).
			BoundsType(types.BoundsTypeRange).
			Sort(
				expr.SortField{
					Expr: expr.MustExpr(b.RootRef(expr.NewStructFieldRef(0)).Build()),
					Kind: types.SortAscNullsFirst,
				},
				expr.SortField{
					Expr: expr.MustExpr(b.RootRef(expr.NewStructFieldRef(1)).Build()),
					Kind: types.SortDescNullsLast,
				},
			).
			BuildExpr()

		require.Error(t, err)
		assert.Contains(t, err.Error(), "RANGE bounds type requires exactly one sort field")
		assert.Contains(t, err.Error(), "got 2")
	})

	t.Run("ROWS with no sort fields should succeed", func(t *testing.T) {
		// ROWS bounds type has no restriction on number of sort fields
		wf, err := b.WindowFunc(sumID).
			Args(b.RootRef(expr.NewStructFieldRef(0))).
			Phase(types.AggPhaseInitialToResult).
			BoundsType(types.BoundsTypeRows).
			BuildExpr()

		require.NoError(t, err)
		require.NotNil(t, wf)

		windowFunc, ok := wf.(*expr.WindowFunction)
		require.True(t, ok)
		assert.Equal(t, types.BoundsTypeRows, windowFunc.BoundsType)
		assert.Len(t, windowFunc.Sorts, 0)
	})

	t.Run("ROWS with multiple sort fields should succeed", func(t *testing.T) {
		// ROWS bounds type has no restriction on number of sort fields
		wf, err := b.WindowFunc(sumID).
			Args(b.RootRef(expr.NewStructFieldRef(0))).
			Phase(types.AggPhaseInitialToResult).
			BoundsType(types.BoundsTypeRows).
			Sort(
				expr.SortField{
					Expr: expr.MustExpr(b.RootRef(expr.NewStructFieldRef(0)).Build()),
					Kind: types.SortAscNullsFirst,
				},
				expr.SortField{
					Expr: expr.MustExpr(b.RootRef(expr.NewStructFieldRef(1)).Build()),
					Kind: types.SortDescNullsLast,
				},
			).
			BuildExpr()

		require.NoError(t, err)
		require.NotNil(t, wf)

		windowFunc, ok := wf.(*expr.WindowFunction)
		require.True(t, ok)
		assert.Equal(t, types.BoundsTypeRows, windowFunc.BoundsType)
		assert.Len(t, windowFunc.Sorts, 2)
	})

	t.Run("UNSPECIFIED with multiple sort fields should succeed", func(t *testing.T) {
		// UNSPECIFIED bounds type has no restriction on number of sort fields
		wf, err := b.WindowFunc(sumID).
			Args(b.RootRef(expr.NewStructFieldRef(0))).
			Phase(types.AggPhaseInitialToResult).
			BoundsType(types.BoundsTypeUnspecified).
			Sort(
				expr.SortField{
					Expr: expr.MustExpr(b.RootRef(expr.NewStructFieldRef(0)).Build()),
					Kind: types.SortAscNullsFirst,
				},
				expr.SortField{
					Expr: expr.MustExpr(b.RootRef(expr.NewStructFieldRef(1)).Build()),
					Kind: types.SortDescNullsLast,
				},
			).
			BuildExpr()

		require.NoError(t, err)
		require.NotNil(t, wf)
	})
}
