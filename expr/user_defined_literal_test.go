// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v7/expr"
	"github.com/substrait-io/substrait-go/v7/extensions"
	"github.com/substrait-io/substrait-go/v7/literal"
	"github.com/substrait-io/substrait-go/v7/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/structpb"
)

// Test extension YAML with point type definition
const testExtensionYAML = `---
urn: extension:test:point_type
types:
  - name: point
    structure:
      latitude: i32
      longitude: i32
`

// TestUserDefinedLiteralRoundTrip tests that a user-defined literal (point type)
// can be round-tripped
func TestUserDefinedLiteralRoundTrip(t *testing.T) {
	collection := &extensions.Collection{}
	err := collection.Load("test/uri", strings.NewReader(testExtensionYAML))
	pointID := extensions.ID{
		URN:  "extension:test:point_type",
		Name: "point",
	}
	require.NoError(t, err)

	registry := expr.NewEmptyExtensionRegistry(collection)

	pointLiteral, err := literal.NewUserDefinedLiteral(
		registry.GetTypeAnchor(pointID),
		expr.StructLiteralValue{
			literal.NewInt32(10, false),
			literal.NewInt32(20, false),
		},
		false,
		nil,
	)
	require.NoError(t, err)
	require.NotNil(t, pointLiteral)

	protoLiteral := pointLiteral.ToProtoLiteral()
	require.NotNil(t, protoLiteral)

	roundTripPointLiteral := expr.LiteralFromProto(protoLiteral)
	require.NotNil(t, roundTripPointLiteral)
	require.Equal(t, pointLiteral, roundTripPointLiteral)
}

// TestNewUserDefinedLiteralHelper demonstrates the simplified API for creating user-defined literals
func TestNewUserDefinedLiteralHelper(t *testing.T) {
	collection := &extensions.Collection{}
	err := collection.Load("test/uri", strings.NewReader(testExtensionYAML))
	require.NoError(t, err)

	registry := expr.NewEmptyExtensionRegistry(collection)
	pointID := extensions.ID{URN: "extension:test:point_type", Name: "point"}

	pointLiteral, err := literal.NewUserDefinedLiteral(
		registry.GetTypeAnchor(pointID),
		expr.StructLiteralValue{
			literal.NewInt32(42, false), // latitude
			literal.NewInt32(73, false), // longitude
		},
		false,
		nil,
	)

	require.NoError(t, err)
	require.NotNil(t, pointLiteral)

	protoLit := pointLiteral.(*expr.ProtoLiteral)
	udt := protoLit.GetType().(*types.UserDefinedType)
	require.Equal(t, registry.GetTypeAnchor(pointID), udt.TypeReference)
	require.Equal(t, types.NullabilityRequired, udt.Nullability)
}

// Test extension YAML with a parameterized generic_box<T> type
const parameterizedExtensionYAML = `---
urn: extension:test:generic
types:
  - name: generic_box
    parameters:
      - name: T
        type: dataType
    structure:
      value: T
`

// TestNewUserDefinedLiteralWithTypeParameters demonstrates creating a user-defined literal
// with type parameters (e.g., generic_box<i32>)
func TestNewUserDefinedLiteralWithTypeParameters(t *testing.T) {
	collection := &extensions.Collection{}
	err := collection.Load("test/uri", strings.NewReader(parameterizedExtensionYAML))
	require.NoError(t, err)

	registry := expr.NewEmptyExtensionRegistry(collection)
	boxID := extensions.ID{URN: "extension:test:generic", Name: "generic_box"}

	boxLiteral, err := literal.NewUserDefinedLiteral(
		registry.GetTypeAnchor(boxID),
		expr.StructLiteralValue{
			literal.NewInt32(100, false),
		},
		false, // nullable
		[]types.TypeParam{&types.DataTypeParameter{Type: &types.Int32Type{}}},
	)

	require.NoError(t, err)
	require.NotNil(t, boxLiteral)

	protoLit := boxLiteral.(*expr.ProtoLiteral)
	udt := protoLit.GetType().(*types.UserDefinedType)
	require.Equal(t, registry.GetTypeAnchor(boxID), udt.TypeReference)
	require.Len(t, udt.TypeParameters, 1)

	// Test roundtrip
	protoLiteral := boxLiteral.ToProtoLiteral()
	roundTripLiteral := expr.LiteralFromProto(protoLiteral)
	require.Equal(t, boxLiteral, roundTripLiteral)
}

// Extension YAML defining nested types (point, triangle) and parameterized type (vector)
const nestedTypesYAML = `---
urn: extension:io.substrait:test_nested_types
types:
  - name: point
    structure:
      latitude: i32
      longitude: i32
  - name: triangle
    structure:
      p1: point
      p2: point
      p3: point
  - name: vector
    parameters:
      - name: T
        type: dataType
    structure:
      x: T
      y: T
      z: T
`

// TestUserDefinedLiteralWithAnyRepresentation verifies round-trip conversion of a simple
// user-defined type using Any representation. With Any representation, the literal value
// is user-managed and can be any proto message (here we use google.protobuf.Struct).
func TestUserDefinedLiteralWithAnyRepresentation(t *testing.T) {
	collection := &extensions.Collection{}
	err := collection.Load("test/uri", strings.NewReader(nestedTypesYAML))
	require.NoError(t, err)

	registry := expr.NewEmptyExtensionRegistry(collection)
	pointID := extensions.ID{URN: "extension:io.substrait:test_nested_types", Name: "point"}

	// Create a user-managed struct representing a point with latitude=42, longitude=100
	// Using google.protobuf.Struct to demonstrate that Any can contain any proto message
	pointStruct, err := structpb.NewStruct(map[string]interface{}{
		"latitude":  float64(42),
		"longitude": float64(100),
	})
	require.NoError(t, err)

	anyValue, err := anypb.New(pointStruct)
	require.NoError(t, err)

	pointLiteral := &expr.ProtoLiteral{
		Value: &proto.Expression_Literal_UserDefined_Value{Value: anyValue},
		Type: &types.UserDefinedType{
			Nullability:    types.NullabilityRequired,
			TypeReference:  registry.GetTypeAnchor(pointID),
			TypeParameters: []types.TypeParam{},
		},
	}

	protoLiteral := pointLiteral.ToProtoLiteral()
	require.NotNil(t, protoLiteral)

	roundTrip := expr.LiteralFromProto(protoLiteral)
	require.NotNil(t, roundTrip)
	require.Equal(t, pointLiteral, roundTrip)
}

// TestUserDefinedLiteralWithStructRepresentation verifies round-trip conversion of a simple
// user-defined type using Struct representation.
func TestUserDefinedLiteralWithStructRepresentation(t *testing.T) {
	collection := &extensions.Collection{}
	err := collection.Load("test/uri", strings.NewReader(nestedTypesYAML))
	require.NoError(t, err)

	registry := expr.NewEmptyExtensionRegistry(collection)
	pointID := extensions.ID{URN: "extension:io.substrait:test_nested_types", Name: "point"}

	pointLiteral, err := literal.NewUserDefinedLiteral(
		registry.GetTypeAnchor(pointID),
		expr.StructLiteralValue{
			literal.NewInt32(42, false),
			literal.NewInt32(100, false),
		},
		false,
		nil,
	)
	require.NoError(t, err)
	require.NotNil(t, pointLiteral)

	protoLiteral := pointLiteral.ToProtoLiteral()
	require.NotNil(t, protoLiteral)

	roundTrip := expr.LiteralFromProto(protoLiteral)
	require.NotNil(t, roundTrip)
	require.Equal(t, pointLiteral, roundTrip)
}

// TestNestedUserDefinedLiteralWithStructRepresentation verifies round-trip conversion of nested
// user-defined types where a triangle UDT contains three point UDTs. Both outer and nested types
// use Struct representation.
func TestNestedUserDefinedLiteralWithStructRepresentation(t *testing.T) {
	collection := &extensions.Collection{}
	err := collection.Load("test/uri", strings.NewReader(nestedTypesYAML))
	require.NoError(t, err)

	registry := expr.NewEmptyExtensionRegistry(collection)
	pointID := extensions.ID{URN: "extension:io.substrait:test_nested_types", Name: "point"}
	triangleID := extensions.ID{URN: "extension:io.substrait:test_nested_types", Name: "triangle"}

	p1, err := literal.NewUserDefinedLiteral(
		registry.GetTypeAnchor(pointID),
		expr.StructLiteralValue{
			literal.NewInt32(0, false),
			literal.NewInt32(0, false),
		},
		false,
		nil,
	)
	require.NoError(t, err)

	p2, err := literal.NewUserDefinedLiteral(
		registry.GetTypeAnchor(pointID),
		expr.StructLiteralValue{
			literal.NewInt32(10, false),
			literal.NewInt32(0, false),
		},
		false,
		nil,
	)
	require.NoError(t, err)

	p3, err := literal.NewUserDefinedLiteral(
		registry.GetTypeAnchor(pointID),
		expr.StructLiteralValue{
			literal.NewInt32(5, false),
			literal.NewInt32(10, false),
		},
		false,
		nil,
	)
	require.NoError(t, err)

	triangle, err := literal.NewUserDefinedLiteral(
		registry.GetTypeAnchor(triangleID),
		expr.StructLiteralValue{p1, p2, p3},
		false,
		nil,
	)
	require.NoError(t, err)
	require.NotNil(t, triangle)

	protoExpression := triangle.ToProtoLiteral()
	require.NotNil(t, protoExpression)

	result := expr.LiteralFromProto(protoExpression)
	require.Equal(t, triangle, result)
}


// TestMixedRepresentationNestedUserDefinedLiteral verifies round-trip conversion of nested
// user-defined types with mixed representations. The triangle UDT uses Struct representation
// while the nested point UDTs use Any representation.
func TestMixedRepresentationNestedUserDefinedLiteral(t *testing.T) {
	collection := &extensions.Collection{}
	err := collection.Load("test/uri", strings.NewReader(nestedTypesYAML))
	require.NoError(t, err)

	registry := expr.NewEmptyExtensionRegistry(collection)
	pointID := extensions.ID{URN: "extension:io.substrait:test_nested_types", Name: "point"}
	triangleID := extensions.ID{URN: "extension:io.substrait:test_nested_types", Name: "triangle"}

	// Helper function to create a point UDT with Any representation (user-managed)
	createPointAny := func(lat, lon int32) expr.Literal {
		pointStruct, err := structpb.NewStruct(map[string]interface{}{
			"latitude":  float64(lat),
			"longitude": float64(lon),
		})
		require.NoError(t, err)

		anyValue, err := anypb.New(pointStruct)
		require.NoError(t, err)

		return &expr.ProtoLiteral{
			Value: &proto.Expression_Literal_UserDefined_Value{Value: anyValue},
			Type: &types.UserDefinedType{
				Nullability:    types.NullabilityRequired,
				TypeReference:  registry.GetTypeAnchor(pointID),
				TypeParameters: []types.TypeParam{},
			},
		}
	}

	p1 := createPointAny(0, 0)
	p2 := createPointAny(10, 0)
	p3 := createPointAny(5, 10)

	// Create triangle UDT using Struct representation, but with Any-encoded point fields
	triangle, err := literal.NewUserDefinedLiteral(
		registry.GetTypeAnchor(triangleID),
		expr.StructLiteralValue{p1, p2, p3},
		false,
		nil,
	)
	require.NoError(t, err)
	require.NotNil(t, triangle)

	// Round-trip test
	protoExpression := triangle.ToProtoLiteral()
	require.NotNil(t, protoExpression)

	result := expr.LiteralFromProto(protoExpression)
	require.Equal(t, triangle, result)
}

// TestParameterizedVectorUDTRoundtrip verifies round-trip conversion of a parameterized
// user-defined type with multiple fields of the same type parameter. Tests that type parameters
// are correctly preserved during serialization and deserialization.
func TestParameterizedVectorUDTRoundtrip(t *testing.T) {
	collection := &extensions.Collection{}
	err := collection.Load("test/uri", strings.NewReader(nestedTypesYAML))
	require.NoError(t, err)

	registry := expr.NewEmptyExtensionRegistry(collection)
	vectorID := extensions.ID{URN: "extension:io.substrait:test_nested_types", Name: "vector"}

	// Create a vector<i32> instance with fields (x: 1, y: 2, z: 3)
	vectorI32, err := literal.NewUserDefinedLiteral(
		registry.GetTypeAnchor(vectorID),
		expr.StructLiteralValue{
			literal.NewInt32(1, false),
			literal.NewInt32(2, false),
			literal.NewInt32(3, false),
		},
		false,
		[]types.TypeParam{
			&types.DataTypeParameter{Type: &types.Int32Type{Nullability: types.NullabilityRequired}},
		},
	)
	require.NoError(t, err)
	require.NotNil(t, vectorI32)

	protoExpression := vectorI32.ToProtoLiteral()
	require.NotNil(t, protoExpression)

	result := expr.LiteralFromProto(protoExpression)
	require.Equal(t, vectorI32, result)

	// Verify type parameters are preserved
	resultProtoLit, ok := result.(*expr.ProtoLiteral)
	require.True(t, ok)
	resultUDT, ok := resultProtoLit.GetType().(*types.UserDefinedType)
	require.True(t, ok)
	require.Len(t, resultUDT.TypeParameters, 1)
}
