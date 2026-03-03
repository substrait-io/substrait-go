// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v7/expr"
	"github.com/substrait-io/substrait-go/v7/types"
)

// TestNewFieldRefFromType_LambdaParameterReference verifies that NewFieldRefFromType
// accepts LambdaParameterReference as a root type.
func TestNewFieldRefFromType_LambdaParameterReference(t *testing.T) {
	stepsOut := uint32(0)
	ref, err := expr.NewFieldRefFromType(
		expr.LambdaParameterReference{StepsOut: stepsOut},
		expr.NewStructFieldRef(0),
		&types.Int64Type{Nullability: types.NullabilityNullable},
	)
	require.NoError(t, err)
	require.NotNil(t, ref)

	// Verify the root is a LambdaParameterReference
	lambdaRef, ok := ref.Root.(expr.LambdaParameterReference)
	require.True(t, ok, "root should be LambdaParameterReference")
	require.Equal(t, stepsOut, lambdaRef.StepsOut)

	// Verify the type was correctly set
	require.True(t, ref.GetType().Equals(&types.Int64Type{Nullability: types.NullabilityNullable}),
		"type should be i64?, got %s", ref.GetType())
}

// TestNewFieldRefFromType_LambdaParameterReference_StepsOut verifies nested lambda refs.
func TestNewFieldRefFromType_LambdaParameterReference_StepsOut(t *testing.T) {
	ref, err := expr.NewFieldRefFromType(
		expr.LambdaParameterReference{StepsOut: 2},
		expr.NewStructFieldRef(1),
		&types.StringType{Nullability: types.NullabilityRequired},
	)
	require.NoError(t, err)
	require.NotNil(t, ref)

	lambdaRef, ok := ref.Root.(expr.LambdaParameterReference)
	require.True(t, ok)
	require.Equal(t, uint32(2), lambdaRef.StepsOut)
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
	require.NotNil(t, ref)

	outerRef, ok := ref.Root.(expr.OuterReference)
	require.True(t, ok, "root should be OuterReference")
	require.Equal(t, expr.OuterReference(1), outerRef)
}

// TestNewFieldRefFromType_UnknownRoot verifies that unknown root types are still rejected.
func TestNewFieldRefFromType_UnknownRoot(t *testing.T) {
	// Using nil as root (which is RootReference) with a ref should work
	ref, err := expr.NewFieldRefFromType(
		expr.RootReference,
		expr.NewStructFieldRef(0),
		&types.Int32Type{},
	)
	require.NoError(t, err)
	require.NotNil(t, ref)
}
