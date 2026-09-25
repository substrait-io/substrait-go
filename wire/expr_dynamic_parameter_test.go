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
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	pb "google.golang.org/protobuf/proto"
)

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

			protoExpr := wire.ExprToProto(tt.dp)
			require.NotNil(t, protoExpr)

			fromProto, err := wire.ExprFromProto(protoExpr, nil, reg)
			require.NoError(t, err)
			assert.True(t, tt.dp.Equals(fromProto), "roundtrip should produce equal expression")

			protoRoundTrip := wire.ExprToProto(fromProto)
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

	_, err := wire.ExprFromProto(protoExpr, nil, expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError()))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "dynamic parameter is nil")
}
