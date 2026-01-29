// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v7/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func TestFuncType(t *testing.T) {
	// Create a function type: func<i32, i64 -> string>
	funcType := &types.FuncType{
		Nullability: types.NullabilityRequired,
		ParameterTypes: []types.Type{
			&types.Int32Type{Nullability: types.NullabilityRequired},
			&types.Int64Type{Nullability: types.NullabilityRequired},
		},
		ReturnType: &types.StringType{Nullability: types.NullabilityRequired},
	}

	// Test String() representation
	require.Equal(t, "func<i32, i64 -> string>", funcType.String())
	require.Equal(t, "func", funcType.ShortString())

	// Test Equals
	sameFuncType := &types.FuncType{
		Nullability: types.NullabilityRequired,
		ParameterTypes: []types.Type{
			&types.Int32Type{Nullability: types.NullabilityRequired},
			&types.Int64Type{Nullability: types.NullabilityRequired},
		},
		ReturnType: &types.StringType{Nullability: types.NullabilityRequired},
	}
	require.True(t, funcType.Equals(sameFuncType))

	differentFuncType := &types.FuncType{
		Nullability: types.NullabilityRequired,
		ParameterTypes: []types.Type{
			&types.Int32Type{Nullability: types.NullabilityRequired},
		},
		ReturnType: &types.StringType{Nullability: types.NullabilityRequired},
	}
	require.False(t, funcType.Equals(differentFuncType))

	// Test nullability variants
	nullableFuncType := funcType.WithNullability(types.NullabilityNullable)
	require.Equal(t, types.NullabilityNullable, nullableFuncType.GetNullability())
	require.Equal(t, "func?<i32, i64 -> string>", nullableFuncType.String())
}

func TestFuncTypeProtoRoundTrip(t *testing.T) {
	// Create a function type
	funcType := &types.FuncType{
		Nullability: types.NullabilityRequired,
		ParameterTypes: []types.Type{
			&types.Int32Type{Nullability: types.NullabilityRequired},
			&types.Float64Type{Nullability: types.NullabilityRequired},
		},
		ReturnType: &types.BooleanType{Nullability: types.NullabilityNullable},
	}

	// Test string representations
	require.Equal(t, "func", funcType.ShortString())
	require.Equal(t, "func<i32, fp64 -> boolean?>", funcType.String())

	// Convert to proto
	protoType := funcType.ToProto()
	require.NotNil(t, protoType)
	require.NotNil(t, protoType.GetFunc())

	// Verify proto structure
	funcProto := protoType.GetFunc()
	require.Len(t, funcProto.ParameterTypes, 2)
	require.Equal(t, proto.Type_NULLABILITY_REQUIRED, funcProto.Nullability)
	require.NotNil(t, funcProto.ReturnType)
	require.Equal(t, proto.Type_NULLABILITY_NULLABLE, funcProto.ReturnType.GetBool().Nullability)

	// Convert back from proto
	resultType := types.TypeFromProto(protoType)
	require.NotNil(t, resultType)

	resultFuncType, ok := resultType.(*types.FuncType)
	require.True(t, ok, "Should be FuncType")

	// Verify roundtrip
	require.True(t, funcType.Equals(resultFuncType))
	require.Equal(t, funcType.String(), resultFuncType.String())
	require.Equal(t, funcType.ShortString(), resultFuncType.ShortString())
}

func TestFuncTypeZeroParameters(t *testing.T) {
	// Function with no parameters: func< -> i32>
	funcType := &types.FuncType{
		Nullability:    types.NullabilityRequired,
		ParameterTypes: []types.Type{},
		ReturnType:     &types.Int32Type{Nullability: types.NullabilityRequired},
	}

	require.Equal(t, "func< -> i32>", funcType.String())
	require.Equal(t, "func", funcType.ShortString())

	// Roundtrip test
	protoType := funcType.ToProto()
	resultType := types.TypeFromProto(protoType)
	require.True(t, funcType.Equals(resultType))
}
