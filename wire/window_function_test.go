// SPDX-License-Identifier: Apache-2.0

package wire_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/types"
	"github.com/substrait-io/substrait-go/v9/wire"
)

func TestWindowFunctionBoundsType(t *testing.T) {
	col := extensions.GetDefaultCollectionWithNoError()
	reg := expr.NewEmptyExtensionRegistry(col)

	sumID := extensions.FunctionID{URN: extensions.SubstraitDefaultURNPrefix + "functions_arithmetic", Name: "sum"}

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
				Phase(types.AggregationPhaseInitialToResult).
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
			protoExpr := wire.ExprToProto(wf)
			deserialized, err := wire.ExprFromProto(protoExpr, schema, reg)
			require.NoError(t, err)

			// Verify the entire expression survived the roundtrip
			assert.Truef(t, wf.Equals(deserialized), "expected: %s\ngot: %s", wf, deserialized)
		})
	}
}
