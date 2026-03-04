// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v7/expr"
	"github.com/substrait-io/substrait-go/v7/types"
)

func TestNewFieldRefFromType_RootTypes(t *testing.T) {
	tests := []struct {
		name     string
		root     expr.RootRefType
		datatype types.Type
		verify   func(t *testing.T, ref *expr.FieldReference, err error)
	}{
		{
			name:     "LambdaParameterReference_StepsOut0",
			root:     expr.LambdaParameterReference{StepsOut: 0},
			datatype: &types.Int64Type{Nullability: types.NullabilityNullable},
			verify: func(t *testing.T, ref *expr.FieldReference, err error) {
				require.NoError(t, err)
				require.IsType(t, expr.LambdaParameterReference{}, ref.Root)
				require.Equal(t, uint32(0), ref.Root.(expr.LambdaParameterReference).StepsOut)
				require.True(t, ref.GetType().Equals(&types.Int64Type{Nullability: types.NullabilityNullable}),
					"type should be i64?, got %s", ref.GetType())
			},
		},
		{
			name:     "LambdaParameterReference_StepsOut2",
			root:     expr.LambdaParameterReference{StepsOut: 2},
			datatype: &types.StringType{Nullability: types.NullabilityRequired},
			verify: func(t *testing.T, ref *expr.FieldReference, err error) {
				require.NoError(t, err)
				require.Equal(t, uint32(2), ref.Root.(expr.LambdaParameterReference).StepsOut)
			},
		},
		{
			name:     "OuterReference",
			root:     expr.OuterReference(1),
			datatype: &types.Int32Type{Nullability: types.NullabilityRequired},
			verify: func(t *testing.T, ref *expr.FieldReference, err error) {
				require.NoError(t, err)
				require.Equal(t, expr.OuterReference(1), ref.Root)
			},
		},
		{
			name:     "RootReference",
			root:     expr.RootReference,
			datatype: &types.Int32Type{},
			verify: func(t *testing.T, ref *expr.FieldReference, err error) {
				require.NoError(t, err)
				require.NotNil(t, ref)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ref, err := expr.NewFieldRefFromType(tt.root, expr.NewStructFieldRef(0), tt.datatype)
			tt.verify(t, ref, err)
		})
	}
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
