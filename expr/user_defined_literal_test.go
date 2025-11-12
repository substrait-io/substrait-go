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
		false, // nullable
	)

	require.NoError(t, err)
	require.NotNil(t, pointLiteral)

	protoLit := pointLiteral.(*expr.ProtoLiteral)
	udt := protoLit.GetType().(*types.UserDefinedType)
	require.Equal(t, registry.GetTypeAnchor(pointID), udt.TypeReference)
	require.Equal(t, types.NullabilityRequired, udt.Nullability)
}
