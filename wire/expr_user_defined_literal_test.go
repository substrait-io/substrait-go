// SPDX-License-Identifier: Apache-2.0

package wire_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/literal"
	"github.com/substrait-io/substrait-go/v9/types"
	"github.com/substrait-io/substrait-go/v9/wire"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// TestUserDefinedValueOneofMatchesProto fails if the spec's UserDefined `val`
// oneof gains or loses a variant, so UserDefinedLiteralValue and its encode/decode
// helpers stay in step with the set of variants they must handle.
func TestUserDefinedValueOneofMatchesProto(t *testing.T) {
	want := map[protoreflect.Name]struct{}{
		"value":  {},
		"struct": {},
	}

	val := (&proto.Expression_Literal_UserDefined{}).ProtoReflect().Descriptor().Oneofs().ByName("val")
	require.NotNil(t, val, "UserDefined message has no val oneof")
	require.Equal(t, len(want), val.Fields().Len(), "UserDefined val oneof variant set changed")

	for i := 0; i < val.Fields().Len(); i++ {
		name := val.Fields().Get(i).Name()
		_, ok := want[name]
		require.Truef(t, ok, "unexpected val oneof variant %q", name)
	}
}

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
	err := collection.Load(strings.NewReader(testExtensionYAML))
	pointID := extensions.TypeID{
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

	protoLiteral := wire.LiteralToProto(pointLiteral)
	require.NotNil(t, protoLiteral)

	roundTripPointLiteral := wire.LiteralFromProto(protoLiteral)
	require.NotNil(t, roundTripPointLiteral)
	require.Equal(t, pointLiteral, roundTripPointLiteral)
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
// is completely user-managed and opaque - it can be any proto message. Here we use a
// simple string to demonstrate this.
func TestUserDefinedLiteralWithAnyRepresentation(t *testing.T) {
	collection := &extensions.Collection{}
	err := collection.Load(strings.NewReader(nestedTypesYAML))
	require.NoError(t, err)

	registry := expr.NewEmptyExtensionRegistry(collection)
	pointID := extensions.TypeID{URN: "extension:io.substrait:test_nested_types", Name: "point"}

	anyValue, err := anypb.New(wrapperspb.String("<Some UserDefined Data>"))
	require.NoError(t, err)

	pointLiteral := &expr.ProtoLiteral{
		Value: expr.UserDefinedValueAny{Value: anyValue},
		Type: &types.UserDefinedType{
			Nullability:    types.NullabilityRequired,
			TypeReference:  registry.GetTypeAnchor(pointID),
			TypeParameters: []types.TypeParam{},
		},
	}

	protoLiteral := wire.LiteralToProto(pointLiteral)
	require.NotNil(t, protoLiteral)

	roundTrip := wire.LiteralFromProto(protoLiteral)
	require.NotNil(t, roundTrip)
	require.Equal(t, pointLiteral, roundTrip)
}

// TestUserDefinedLiteralWithStructRepresentation verifies round-trip conversion of a simple
// user-defined type using Struct representation.
func TestUserDefinedLiteralWithStructRepresentation(t *testing.T) {
	collection := &extensions.Collection{}
	err := collection.Load(strings.NewReader(nestedTypesYAML))
	require.NoError(t, err)

	registry := expr.NewEmptyExtensionRegistry(collection)
	pointID := extensions.TypeID{URN: "extension:io.substrait:test_nested_types", Name: "point"}

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

	protoLiteral := wire.LiteralToProto(pointLiteral)
	require.NotNil(t, protoLiteral)

	roundTrip := wire.LiteralFromProto(protoLiteral)
	require.NotNil(t, roundTrip)
	require.Equal(t, pointLiteral, roundTrip)
}

// TestNestedUserDefinedLiteralWithAnyRepresentation verifies round-trip conversion of nested
// user-defined types where both outer and nested types use Any representation. The triangle UDT
// uses Any representation, and would typically encode its nested point UDTs within that Any value.
func TestNestedUserDefinedLiteralWithAnyRepresentation(t *testing.T) {
	collection := &extensions.Collection{}
	err := collection.Load(strings.NewReader(nestedTypesYAML))
	require.NoError(t, err)

	registry := expr.NewEmptyExtensionRegistry(collection)
	triangleID := extensions.TypeID{URN: "extension:io.substrait:test_nested_types", Name: "triangle"}

	anyValue, err := anypb.New(wrapperspb.String("<Some UserDefined Data>"))
	require.NoError(t, err)

	triangleLiteral := &expr.ProtoLiteral{
		Value: expr.UserDefinedValueAny{Value: anyValue},
		Type: &types.UserDefinedType{
			Nullability:    types.NullabilityRequired,
			TypeReference:  registry.GetTypeAnchor(triangleID),
			TypeParameters: []types.TypeParam{},
		},
	}

	protoLiteral := wire.LiteralToProto(triangleLiteral)
	require.NotNil(t, protoLiteral)

	roundTrip := wire.LiteralFromProto(protoLiteral)
	require.NotNil(t, roundTrip)
	require.Equal(t, triangleLiteral, roundTrip)
}

// TestNestedUserDefinedLiteralWithStructRepresentation verifies round-trip conversion of nested
// user-defined types where a triangle UDT contains three point UDTs. Both outer and nested types
// use Struct representation.
func TestNestedUserDefinedLiteralWithStructRepresentation(t *testing.T) {
	collection := &extensions.Collection{}
	err := collection.Load(strings.NewReader(nestedTypesYAML))
	require.NoError(t, err)

	registry := expr.NewEmptyExtensionRegistry(collection)
	pointID := extensions.TypeID{URN: "extension:io.substrait:test_nested_types", Name: "point"}
	triangleID := extensions.TypeID{URN: "extension:io.substrait:test_nested_types", Name: "triangle"}

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

	protoExpression := wire.LiteralToProto(triangle)
	require.NotNil(t, protoExpression)

	result := wire.LiteralFromProto(protoExpression)
	require.Equal(t, triangle, result)
}

// TestMixedRepresentationNestedUserDefinedLiteral verifies round-trip conversion of nested
// user-defined types with mixed representations. The triangle UDT uses Struct representation
// while the nested point UDTs use Any representation.
func TestMixedRepresentationNestedUserDefinedLiteral(t *testing.T) {
	collection := &extensions.Collection{}
	err := collection.Load(strings.NewReader(nestedTypesYAML))
	require.NoError(t, err)

	registry := expr.NewEmptyExtensionRegistry(collection)
	pointID := extensions.TypeID{URN: "extension:io.substrait:test_nested_types", Name: "point"}
	triangleID := extensions.TypeID{URN: "extension:io.substrait:test_nested_types", Name: "triangle"}

	// Helper function to create a point UDT with Any representation (user-managed)
	// The Any value is completely opaque - it can be any proto message.
	createPointAny := func() expr.Literal {
		anyValue, err := anypb.New(wrapperspb.String("<Some UserDefined Data>"))
		require.NoError(t, err)

		return &expr.ProtoLiteral{
			Value: expr.UserDefinedValueAny{Value: anyValue},
			Type: &types.UserDefinedType{
				Nullability:    types.NullabilityRequired,
				TypeReference:  registry.GetTypeAnchor(pointID),
				TypeParameters: []types.TypeParam{},
			},
		}
	}

	// Create triangle UDT using Struct representation, but with Any-encoded point fields
	triangle, err := literal.NewUserDefinedLiteral(
		registry.GetTypeAnchor(triangleID),
		expr.StructLiteralValue{createPointAny(), createPointAny(), createPointAny()},
		false,
		nil,
	)
	require.NoError(t, err)
	require.NotNil(t, triangle)

	protoExpression := wire.LiteralToProto(triangle)
	require.NotNil(t, protoExpression)

	result := wire.LiteralFromProto(protoExpression)
	require.Equal(t, triangle, result)
}

// TestUserDefinedLiteralNilStructDecodes verifies a UserDefined literal whose struct
// oneof is selected but nil decodes without dereferencing the absent struct.
func TestUserDefinedLiteralNilStructDecodes(t *testing.T) {
	protoLit := &proto.Expression_Literal{
		LiteralType: &proto.Expression_Literal_UserDefined_{
			UserDefined: &proto.Expression_Literal_UserDefined{
				TypeAnchorType: &proto.Expression_Literal_UserDefined_TypeReference{TypeReference: 1},
				Val:            &proto.Expression_Literal_UserDefined_Struct{Struct: nil},
			},
		},
	}
	require.NotPanics(t, func() { wire.LiteralFromProto(protoLit) })
}

// TestUserDefinedLiteralPointerValuePanics verifies that a pointer form of a value
// (which also satisfies UserDefinedLiteralValue via the value-receiver marker) fails
// loud on serialization rather than silently dropping the payload.
func TestUserDefinedLiteralPointerValueRoundtrip(t *testing.T) {
	anyValue, err := anypb.New(wrapperspb.String("data"))
	require.NoError(t, err)

	lit := &expr.ProtoLiteral{
		Value: &expr.UserDefinedValueAny{Value: anyValue},
		Type:  &types.UserDefinedType{Nullability: types.NullabilityRequired, TypeReference: 1},
	}
	protoLit := wire.LiteralToProto(lit)

	userDefined := protoLit.GetUserDefined()
	require.NotNil(t, userDefined)
	require.Equal(t, anyValue, userDefined.GetValue())
}

func TestUserDefinedLiteralPointerStructRoundtrip(t *testing.T) {
	structValue := expr.StructLiteralValue{literal.NewInt32(1, false)}
	lit := &expr.ProtoLiteral{
		Value: &expr.UserDefinedValueStruct{Value: structValue},
		Type:  &types.UserDefinedType{Nullability: types.NullabilityRequired, TypeReference: 1},
	}
	protoLit := wire.LiteralToProto(lit)

	userDefined := protoLit.GetUserDefined()
	require.NotNil(t, userDefined)
	require.Equal(t, wire.StructLiteralValueToProto(structValue), userDefined.GetStruct())
}

// TestParameterizedVectorUDTRoundtrip verifies round-trip conversion of a parameterized
// user-defined type with multiple fields of the same type parameter. Tests that type parameters
// are correctly preserved during serialization and deserialization.
func TestParameterizedVectorUDTRoundtrip(t *testing.T) {
	collection := &extensions.Collection{}
	err := collection.Load(strings.NewReader(nestedTypesYAML))
	require.NoError(t, err)

	registry := expr.NewEmptyExtensionRegistry(collection)
	vectorID := extensions.TypeID{URN: "extension:io.substrait:test_nested_types", Name: "vector"}

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

	protoExpression := wire.LiteralToProto(vectorI32)
	require.NotNil(t, protoExpression)

	result := wire.LiteralFromProto(protoExpression)
	require.Equal(t, vectorI32, result)

	// Verify type parameters are preserved
	resultProtoLit, ok := result.(*expr.ProtoLiteral)
	require.True(t, ok)
	resultUDT, ok := resultProtoLit.GetType().(*types.UserDefinedType)
	require.True(t, ok)
	require.Len(t, resultUDT.TypeParameters, 1)
}
