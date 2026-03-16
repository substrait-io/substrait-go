// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v7/expr"
	"github.com/substrait-io/substrait-go/v7/extensions"
	"github.com/substrait-io/substrait-go/v7/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	pb "google.golang.org/protobuf/proto"
)

func TestDynamicParameterEquals(t *testing.T) {
	dp1 := &expr.DynamicParameter{
		OutputType:         &types.Int64Type{Nullability: types.NullabilityRequired},
		ParameterReference: 0,
	}

	dp2 := &expr.DynamicParameter{
		OutputType:         &types.Int64Type{Nullability: types.NullabilityRequired},
		ParameterReference: 0,
	}

	dp3 := &expr.DynamicParameter{
		OutputType:         &types.Int64Type{Nullability: types.NullabilityRequired},
		ParameterReference: 1,
	}

	dp4 := &expr.DynamicParameter{
		OutputType:         &types.Float64Type{Nullability: types.NullabilityRequired},
		ParameterReference: 0,
	}

	assert.True(t, dp1.Equals(dp2), "same type and ref should be equal")
	assert.False(t, dp1.Equals(dp3), "different ref should not be equal")
	assert.False(t, dp1.Equals(dp4), "different type should not be equal")
	assert.False(t, dp1.Equals(expr.NewPrimitiveLiteral(int64(42), false)), "different expression type should not be equal")
}

func TestDynamicParameterVisit(t *testing.T) {
	dp := &expr.DynamicParameter{
		OutputType:         &types.Int32Type{Nullability: types.NullabilityRequired},
		ParameterReference: 5,
	}

	// Visit should return the same expression since DynamicParameter has no children
	visited := dp.Visit(func(e expr.Expression) expr.Expression {
		return e
	})

	assert.Same(t, dp, visited, "Visit should return same pointer for leaf expression")
}

// TestDynamicParameterToProtoRoundtrip tests construction, interface compliance,
// and proto roundtrip for various DynamicParameter configurations.
// The $N:type String() format (e.g. "$0:i32") is an internal debugging
// representation used by this library; it is not part of the Substrait spec.
func TestDynamicParameterToProtoRoundtrip(t *testing.T) {
	tests := []struct {
		name string
		dp   *expr.DynamicParameter
	}{
		{
			name: "required i32",
			dp: &expr.DynamicParameter{
				OutputType:         &types.Int32Type{Nullability: types.NullabilityRequired},
				ParameterReference: 0,
			},
		},
		{
			name: "nullable string",
			dp: &expr.DynamicParameter{
				OutputType:         &types.StringType{Nullability: types.NullabilityNullable},
				ParameterReference: 1,
			},
		},
		{
			name: "required fp64",
			dp: &expr.DynamicParameter{
				OutputType:         &types.Float64Type{Nullability: types.NullabilityRequired},
				ParameterReference: 5,
			},
		},
		{
			name: "required boolean",
			dp: &expr.DynamicParameter{
				OutputType:         &types.BooleanType{Nullability: types.NullabilityRequired},
				ParameterReference: 10,
			},
		},
		{
			name: "nullable i64",
			dp: &expr.DynamicParameter{
				OutputType:         &types.Int64Type{Nullability: types.NullabilityNullable},
				ParameterReference: 42,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, tt.dp.IsScalar())
			assert.True(t, tt.dp.GetType().Equals(tt.dp.OutputType))

			// Proto roundtrip: the plan should equal itself after a roundtrip
			protoExpr := tt.dp.ToProto()
			require.NotNil(t, protoExpr)

			fromProto, err := expr.ExprFromProto(protoExpr, nil, expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError()))
			require.NoError(t, err)
			assert.True(t, tt.dp.Equals(fromProto), "roundtrip should produce equal expression")

			protoRoundTrip := fromProto.ToProto()
			assert.True(t, pb.Equal(protoExpr, protoRoundTrip), "proto roundtrip should be equal")
		})
	}
}

func TestDynamicParameterToProtoFuncArg(t *testing.T) {
	dp := &expr.DynamicParameter{
		OutputType:         &types.Int32Type{Nullability: types.NullabilityRequired},
		ParameterReference: 0,
	}

	arg := dp.ToProtoFuncArg()
	require.NotNil(t, arg)
	require.NotNil(t, arg.GetValue(), "should be a value argument")
	require.NotNil(t, arg.GetValue().GetDynamicParameter(), "value should be a dynamic parameter")
}

func TestDynamicParameterFromProtoNilDynamicParam(t *testing.T) {
	// Test ExprFromProto with a DynamicParameter that has nil inner
	protoExpr := &proto.Expression{
		RexType: &proto.Expression_DynamicParameter{
			DynamicParameter: nil,
		},
	}

	_, err := expr.ExprFromProto(protoExpr, nil, expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError()))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "dynamic parameter is nil")
}

func TestDynamicParameterBuilder(t *testing.T) {
	b := expr.ExprBuilder{
		Reg: expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError()),
	}

	tests := []struct {
		name   string
		build  func() (expr.Expression, error)
		expect string
		err    string
	}{
		{
			name: "basic i32",
			build: func() (expr.Expression, error) {
				return b.DynamicParam(&types.Int32Type{Nullability: types.NullabilityRequired}, 0).BuildExpr()
			},
			expect: "$0:i32",
		},
		{
			name: "nullable string param 3",
			build: func() (expr.Expression, error) {
				return b.DynamicParam(&types.StringType{Nullability: types.NullabilityNullable}, 3).BuildExpr()
			},
			expect: "$3:string?",
		},
		{
			name: "nil type should error",
			build: func() (expr.Expression, error) {
				return b.DynamicParam(nil, 0).BuildExpr()
			},
			err: "dynamic parameter must have an output type",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, err := tt.build()
			if tt.err != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expect, e.String())
				// Verify proto roundtrip
				e.ToProto()
			}
		})
	}
}

func TestDynamicParameterBuilderAsFuncArg(t *testing.T) {
	b := expr.ExprBuilder{
		Reg:        expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError()),
		BaseSchema: types.NewRecordTypeFromStruct(boringSchema.Struct),
	}

	// Use DynamicParam as a function argument via the builder
	dpBuilder := b.DynamicParam(&types.Int8Type{Nullability: types.NullabilityRequired}, 0)

	// Verify it implements FuncArgBuilder
	funcArg, err := dpBuilder.BuildFuncArg()
	require.NoError(t, err)
	assert.NotNil(t, funcArg)

	dp, ok := funcArg.(*expr.DynamicParameter)
	require.True(t, ok)
	assert.Equal(t, uint32(0), dp.ParameterReference)

	// Build as a function argument in a scalar function
	e, err := b.ScalarFunc(addID).Args(
		dpBuilder,
		b.Wrap(expr.NewLiteral(int8(5), false)),
	).BuildExpr()
	require.NoError(t, err)
	assert.Contains(t, e.String(), "$0:i8")
}

func TestDynamicParameterInProject(t *testing.T) {
	// Test using dynamic parameter in a project expression through builders

	dp := &expr.DynamicParameter{
		OutputType:         &types.Int32Type{Nullability: types.NullabilityRequired},
		ParameterReference: 0,
	}

	// Verify it can be used as a project expression
	protoExpr := dp.ToProto()
	require.NotNil(t, protoExpr)

	// Roundtrip
	fromProto, err := expr.ExprFromProto(protoExpr, nil, expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError()))
	require.NoError(t, err)
	require.IsType(t, &expr.DynamicParameter{}, fromProto)

	roundtripped := fromProto.(*expr.DynamicParameter)
	assert.Equal(t, uint32(0), roundtripped.ParameterReference)
	assert.True(t, roundtripped.GetType().Equals(&types.Int32Type{Nullability: types.NullabilityRequired}))
}

func TestDynamicParameterMultipleInExpression(t *testing.T) {
	dp0 := &expr.DynamicParameter{
		OutputType:         &types.Int32Type{Nullability: types.NullabilityRequired},
		ParameterReference: 0,
	}

	dp1 := &expr.DynamicParameter{
		OutputType:         &types.StringType{Nullability: types.NullabilityNullable},
		ParameterReference: 1,
	}

	// Both should work independently
	proto0 := dp0.ToProto()
	proto1 := dp1.ToProto()
	require.NotNil(t, proto0)
	require.NotNil(t, proto1)

	from0, err := expr.ExprFromProto(proto0, nil, expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError()))
	require.NoError(t, err)
	from1, err := expr.ExprFromProto(proto1, nil, expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError()))
	require.NoError(t, err)

	// They should not be equal to each other
	assert.False(t, from0.Equals(from1))
	// But should be equal to themselves
	assert.True(t, from0.Equals(dp0))
	assert.True(t, from1.Equals(dp1))
}
