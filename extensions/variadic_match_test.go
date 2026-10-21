// SPDX-License-Identifier: Apache-2.0

package extensions

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/types"
	"github.com/substrait-io/substrait-go/v9/types/parser"
)

func TestMatchVariadicArgumentCounts(t *testing.T) {
	param, err := parser.ParseType("i64")
	require.NoError(t, err)
	arg := ValueArg{Value: &parser.TypeExpression{ValueType: param}}
	params := FuncParameterList{arg, arg, arg}
	args := []types.Type{&types.Int64Type{}, &types.Int64Type{}, &types.Int64Type{}, &types.Int64Type{}}
	for count := 0; count <= len(args); count++ {
		matched, err := matchArguments(MirrorNullability, params, &VariadicBehavior{Min: 1}, args[:count])
		require.NoError(t, err)
		require.Equal(t, count >= 3, matched, "argument count %d", count)
	}
}
