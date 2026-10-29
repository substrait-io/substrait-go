// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/types"
)

func TestCustomAggregateFunctionMetadata(t *testing.T) {
	id := extensions.FunctionID{
		URN:  "extension:io.substrait:functions_arithmetic_decimal",
		Name: "sum:dec",
	}
	collection := extensions.GetDefaultCollectionWithNoError()
	reg := expr.NewEmptyExtensionRegistry(collection)
	variant, ok := collection.GetAggregateFunc(id)
	require.True(t, ok)
	fn, err := expr.NewCustomAggregateFunc(reg, variant, &types.Int64Type{}, nil,
		types.AggregationInvocationAll, types.AggregationPhaseInitialToResult, nil)
	require.NoError(t, err)
	require.Equal(t, "sum", fn.Name())
	require.Equal(t, id.Name, fn.CompoundName())
	require.Equal(t, id, fn.ID())
}
