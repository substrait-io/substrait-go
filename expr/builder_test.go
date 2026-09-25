// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
)

func TestAny1TypeParameterConsistency(t *testing.T) {
	// equal(any1, any1) -> boolean
	reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())

	equalID := extensions.FunctionID{
		URN:  extensions.SubstraitDefaultURNPrefix + "functions_comparison",
		Name: "equal",
	}

	t.Run("non-variadic any1 - invalid", func(t *testing.T) {
		_, err := expr.NewScalarFunc(reg, equalID, nil,
			expr.NewPrimitiveLiteral(int32(5), false),
			expr.NewPrimitiveLiteral(int64(10), false))
		require.Error(t, err, "equal should reject mixed types (i32, i64)")

		_, err = expr.NewScalarFunc(reg, equalID, nil,
			expr.NewPrimitiveLiteral("hello", false),
			expr.NewPrimitiveLiteral(int32(5), false))
		require.Error(t, err, "equal should reject mixed types (string, i32)")
	})

	t.Run("non-variadic any1 - valid", func(t *testing.T) {
		_, err := expr.NewScalarFunc(reg, equalID, nil,
			expr.NewPrimitiveLiteral(int32(5), false),
			expr.NewPrimitiveLiteral(int32(10), false))
		require.NoError(t, err, "equal should accept matching types (i32, i32)")

		_, err = expr.NewScalarFunc(reg, equalID, nil,
			expr.NewPrimitiveLiteral("hello", false),
			expr.NewPrimitiveLiteral("world", false))
		require.NoError(t, err, "equal should accept matching types (string, string)")
	})

	// coalesce(any1, any1, ...) -> any1 (min: 2 args)
	coalesceID := extensions.FunctionID{
		URN:  extensions.SubstraitDefaultURNPrefix + "functions_comparison",
		Name: "coalesce",
	}

	t.Run("variadic any1 - invalid", func(t *testing.T) {
		_, err := expr.NewScalarFunc(reg, coalesceID, nil,
			expr.NewPrimitiveLiteral(int32(5), false),
			expr.NewPrimitiveLiteral(int64(10), false))
		require.Error(t, err, "coalesce should reject mixed types (i32, i64)")

		_, err = expr.NewScalarFunc(reg, coalesceID, nil,
			expr.NewPrimitiveLiteral("hello", false),
			expr.NewPrimitiveLiteral(int32(5), false))
		require.Error(t, err, "coalesce should reject mixed types (string, i32)")
	})

	t.Run("variadic any1 - valid", func(t *testing.T) {
		_, err := expr.NewScalarFunc(reg, coalesceID, nil,
			expr.NewPrimitiveLiteral(int32(1), false),
			expr.NewPrimitiveLiteral(int32(2), false),
			expr.NewPrimitiveLiteral(int32(3), false))
		require.NoError(t, err, "coalesce should accept all matching types")

		_, err = expr.NewScalarFunc(reg, coalesceID, nil,
			expr.NewPrimitiveLiteral("a", false),
			expr.NewPrimitiveLiteral("b", false),
			expr.NewPrimitiveLiteral("c", false))
		require.NoError(t, err, "coalesce should accept all matching types")
	})
}
