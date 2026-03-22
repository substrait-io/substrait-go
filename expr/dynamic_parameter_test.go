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
	i64Req := &types.Int64Type{Nullability: types.NullabilityRequired}
	fp64Req := &types.Float64Type{Nullability: types.NullabilityRequired}

	base := &expr.DynamicParameter{OutputType: i64Req, ParameterReference: 0}

	tests := []struct {
		name  string
		other expr.Expression
		want  bool
	}{
		{"same type and ref", &expr.DynamicParameter{OutputType: i64Req, ParameterReference: 0}, true},
		{"different ref", &expr.DynamicParameter{OutputType: i64Req, ParameterReference: 1}, false},
		{"different type", &expr.DynamicParameter{OutputType: fp64Req, ParameterReference: 0}, false},
		{"different expression kind", expr.NewPrimitiveLiteral(int64(42), false), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, base.Equals(tt.other))
		})
	}
}

func TestDynamicParameterVisit(t *testing.T) {
	dp := &expr.DynamicParameter{
		OutputType:         &types.Int32Type{Nullability: types.NullabilityRequired},
		ParameterReference: 5,
	}

	visited := dp.Visit(func(e expr.Expression) expr.Expression { return e })
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
		{"required i32", &expr.DynamicParameter{
			OutputType: &types.Int32Type{Nullability: types.NullabilityRequired}, ParameterReference: 0}},
		{"nullable string", &expr.DynamicParameter{
			OutputType: &types.StringType{Nullability: types.NullabilityNullable}, ParameterReference: 1}},
		{"required fp64", &expr.DynamicParameter{
			OutputType: &types.Float64Type{Nullability: types.NullabilityRequired}, ParameterReference: 5}},
		{"required boolean", &expr.DynamicParameter{
			OutputType: &types.BooleanType{Nullability: types.NullabilityRequired}, ParameterReference: 10}},
		{"nullable i64", &expr.DynamicParameter{
			OutputType: &types.Int64Type{Nullability: types.NullabilityNullable}, ParameterReference: 42}},
	}

	reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, tt.dp.IsScalar())
			assert.True(t, tt.dp.GetType().Equals(tt.dp.OutputType))

			protoExpr := tt.dp.ToProto()
			require.NotNil(t, protoExpr)

			fromProto, err := expr.ExprFromProto(protoExpr, nil, reg)
			require.NoError(t, err)
			assert.True(t, tt.dp.Equals(fromProto), "roundtrip should produce equal expression")

			protoRoundTrip := fromProto.ToProto()
			assert.True(t, pb.Equal(protoExpr, protoRoundTrip), "proto roundtrip should be equal")
		})
	}
}

func TestDynamicParameterFromProtoNilDynamicParam(t *testing.T) {
	protoExpr := &proto.Expression{
		RexType: &proto.Expression_DynamicParameter{
			DynamicParameter: nil,
		},
	}

	_, err := expr.ExprFromProto(protoExpr, nil, expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError()))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "dynamic parameter is nil")
}

func TestDynamicParameterBuilderNilType(t *testing.T) {
	b := expr.ExprBuilder{
		Reg: expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError()),
	}

	_, err := b.DynamicParam(nil, 0).BuildExpr()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "dynamic parameter must have an output type")
}

func TestDynamicParameterBuilderAsFuncArg(t *testing.T) {
	b := expr.ExprBuilder{
		Reg:        expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError()),
		BaseSchema: types.NewRecordTypeFromStruct(boringSchema.Struct),
	}

	dpBuilder := b.DynamicParam(&types.Int8Type{Nullability: types.NullabilityRequired}, 0)

	e, err := b.ScalarFunc(addID).Args(
		dpBuilder,
		b.Wrap(expr.NewLiteral(int8(5), false)),
	).BuildExpr()
	require.NoError(t, err)
	assert.Contains(t, e.String(), "$0:i8")
}

func TestDynamicParameterTypeMismatchInFunction(t *testing.T) {
	b := expr.ExprBuilder{
		Reg:        expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError()),
		BaseSchema: types.NewRecordTypeFromStruct(boringSchema.Struct),
	}

	tests := []struct {
		name   string
		funcID extensions.ID
		dpType types.Type
		lit    func() (expr.Literal, error)
	}{
		{
			name:   "i32 where i8 expected",
			funcID: extensions.ID{URN: extensions.SubstraitDefaultURNPrefix + "functions_arithmetic", Name: "add:i8_i8"},
			dpType: &types.Int32Type{Nullability: types.NullabilityRequired},
			lit:    func() (expr.Literal, error) { return expr.NewLiteral(int8(5), false) },
		},
		{
			name:   "string where numeric expected",
			funcID: addID,
			dpType: &types.StringType{Nullability: types.NullabilityRequired},
			lit:    func() (expr.Literal, error) { return expr.NewLiteral(int32(5), false) },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := b.ScalarFunc(tt.funcID).Args(
				b.DynamicParam(tt.dpType, 0),
				b.Wrap(tt.lit()),
			).BuildExpr()
			require.Error(t, err)
		})
	}
}
