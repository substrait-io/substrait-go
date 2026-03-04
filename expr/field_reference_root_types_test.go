// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v7/expr"
	"github.com/substrait-io/substrait-go/v7/types"
)

// TestNewFieldRefFromType_LambdaParameterReference verifies that NewFieldRefFromType
// accepts LambdaParameterReference as a root type and preserves stepsOut and type.
func TestNewFieldRefFromType_LambdaParameterReference(t *testing.T) {
	ref, err := expr.NewFieldRefFromType(
		expr.LambdaParameterReference{StepsOut: 0},
		expr.NewStructFieldRef(0),
		&types.Int64Type{Nullability: types.NullabilityNullable},
	)
	require.NoError(t, err)
	require.IsType(t, expr.LambdaParameterReference{}, ref.Root)
	require.Equal(t, uint32(0), ref.Root.(expr.LambdaParameterReference).StepsOut)
	require.True(t, ref.GetType().Equals(&types.Int64Type{Nullability: types.NullabilityNullable}),
		"type should be i64?, got %s", ref.GetType())
}

// TestNewFieldRefFromType_LambdaParameterReference_StepsOut verifies nested lambda
// references with stepsOut > 0.
func TestNewFieldRefFromType_LambdaParameterReference_StepsOut(t *testing.T) {
	ref, err := expr.NewFieldRefFromType(
		expr.LambdaParameterReference{StepsOut: 2},
		expr.NewStructFieldRef(1),
		&types.StringType{Nullability: types.NullabilityRequired},
	)
	require.NoError(t, err)
	require.Equal(t, uint32(2), ref.Root.(expr.LambdaParameterReference).StepsOut)
}

// TestNewFieldRefFromType_OuterReference verifies that NewFieldRefFromType
// accepts OuterReference as a root type.
func TestNewFieldRefFromType_OuterReference(t *testing.T) {
	ref, err := expr.NewFieldRefFromType(
		expr.OuterReference(1),
		expr.NewStructFieldRef(0),
		&types.Int32Type{Nullability: types.NullabilityRequired},
	)
	require.NoError(t, err)
	require.Equal(t, expr.OuterReference(1), ref.Root)
}

// TestNewFieldRefFromType_RootReference verifies that RootReference still works.
func TestNewFieldRefFromType_RootReference(t *testing.T) {
	ref, err := expr.NewFieldRefFromType(
		expr.RootReference,
		expr.NewStructFieldRef(0),
		&types.Int32Type{},
	)
	require.NoError(t, err)
	require.NotNil(t, ref)
}

// TestNewFieldRef_LambdaParameterReference_ReturnsError verifies that NewFieldRef
// (not NewFieldRefFromType) returns a clear error when called with LambdaParameterReference,
// directing callers to use NewFieldRefFromType instead.
func TestNewFieldRef_LambdaParameterReference_ReturnsError(t *testing.T) {
	_, err := expr.NewFieldRef(
		expr.LambdaParameterReference{StepsOut: 0},
		expr.NewStructFieldRef(0),
		nil, // no base schema
	)
	require.Error(t, err)
	require.Contains(t, err.Error(), "LambdaParameterReference")
	require.Contains(t, err.Error(), "NewFieldRefFromType")
}

// TestNewFieldRef_OuterReference_ReturnsError verifies that NewFieldRef
// returns a clear error when called with OuterReference.
func TestNewFieldRef_OuterReference_ReturnsError(t *testing.T) {
	_, err := expr.NewFieldRef(
		expr.OuterReference(1),
		expr.NewStructFieldRef(0),
		nil, // no base schema
	)
	require.Error(t, err)
	require.Contains(t, err.Error(), "OuterReference")
	require.Contains(t, err.Error(), "NewFieldRefFromType")
}
