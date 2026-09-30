// SPDX-License-Identifier: Apache-2.0

package extensions_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/types"
	"github.com/substrait-io/substrait-go/v9/types/parser"
	"google.golang.org/protobuf/proto"
)

func TestEvaluateTypeExpressionNestedUDTAnchors(t *testing.T) {
	const urn = "extension:example:return_types"
	const required = types.NullabilityRequired
	const nullable = types.NullabilityNullable
	udt := func(registry extensions.Set, name string, nullability types.Nullability) *types.UserDefinedType {
		return &types.UserDefinedType{
			TypeReference: registry.GetTypeAnchor(extensions.TypeID{URN: urn, Name: name}),
			Nullability:   nullability,
		}
	}
	tests := []struct {
		name       string
		expression string
		expected   func(extensions.Set) types.Type
	}{
		{"top level", "u!item?", func(r extensions.Set) types.Type { return udt(r, "item", nullable) }},
		{"list element", "list?<u!item?>", func(r extensions.Set) types.Type {
			return &types.ListType{Nullability: nullable, Type: udt(r, "item", nullable)}
		}},
		{"map value", "map?<string, u!item?>", func(r extensions.Set) types.Type {
			return &types.MapType{Nullability: nullable, Key: &types.StringType{Nullability: required}, Value: udt(r, "item", nullable)}
		}},
		{"map key and value", "map<u!key, u!item?>", func(r extensions.Set) types.Type {
			return &types.MapType{Nullability: required, Key: udt(r, "key", required), Value: udt(r, "item", nullable)}
		}},
		{"struct fields", "struct?<u!item?, i64, u!other, u!item>", func(r extensions.Set) types.Type {
			return &types.StructType{Nullability: nullable, Types: []types.Type{
				udt(r, "item", nullable), &types.Int64Type{Nullability: required}, udt(r, "other", required), udt(r, "item", required),
			}}
		}},
		{"nested containers", "list<map?<string, struct<u!item?, list<u!other>>>>", func(r extensions.Set) types.Type {
			return &types.ListType{Nullability: required, Type: &types.MapType{
				Nullability: nullable, Key: &types.StringType{Nullability: required},
				Value: &types.StructType{Nullability: required, Types: []types.Type{
					udt(r, "item", nullable), &types.ListType{Nullability: required, Type: udt(r, "other", required)},
				}},
			}}
		}},
		{"function parameter and return", "func<list<u!other?> -> u!item?>", func(r extensions.Set) types.Type {
			return &types.FuncType{Nullability: required, ParameterTypes: []types.Type{
				&types.ListType{Nullability: required, Type: udt(r, "other", nullable)},
			}, ReturnType: udt(r, "item", nullable)}
		}},
		{"UDT parameters", "u!box?<10, label, list<u!item?>>", func(r extensions.Set) types.Type {
			box := udt(r, "box", nullable)
			box.TypeParameters = []types.TypeParam{
				types.IntegerParameter(10), types.StringParameter("label"),
				&types.DataTypeParameter{Type: &types.ListType{Nullability: required, Type: udt(r, "item", nullable)}},
			}
			return box
		}},
		{"derived UDT", "N = 10\nu!item?", func(r extensions.Set) types.Type { return udt(r, "item", nullable) }},
		{"derived container", "N = 10\nmap<string, u!item?>", func(r extensions.Set) types.Type {
			return &types.MapType{Nullability: required, Key: &types.StringType{Nullability: required}, Value: udt(r, "item", nullable)}
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			returnType, err := parser.ParseType(tt.expression)
			require.NoError(t, err)
			registry := extensions.NewSet()
			// An unresolved default reference must not accidentally identify the intended UDT.
			registry.GetTypeAnchor(extensions.TypeID{URN: "extension:example:unrelated", Name: "sentinel"})
			result, err := extensions.EvaluateTypeExpression(urn, extensions.DeclaredOutputNullability,
				returnType, nil, nil, nil, registry)
			require.NoError(t, err)
			require.Equal(t, tt.expected(registry), result)
		})
	}
}

func TestFunctionVariantsBindNestedUDTAnchors(t *testing.T) {
	const urn = "extension:example:return_types"
	const definition = `
urn: extension:example:return_types
types:
  - name: item
scalar_functions:
  - name: scalar
    impls:
      - args:
          - value: i64
        nullability: DECLARED_OUTPUT
        return: map?<string, u!item?>
aggregate_functions:
  - name: aggregate
    impls:
      - args:
          - value: i64
        nullability: DECLARED_OUTPUT
        return: map?<string, u!item?>
        decomposable: NONE
window_functions:
  - name: window
    impls:
      - args:
          - value: i64
        nullability: DECLARED_OUTPUT
        return: map?<string, u!item?>
        decomposable: NONE
        window_type: PARTITION
`
	var collection extensions.Collection
	require.NoError(t, collection.Load(strings.NewReader(definition)))
	scalar, ok := collection.GetScalarFunc(extensions.FunctionID{URN: urn, Name: "scalar:i64"})
	require.True(t, ok)
	aggregate, ok := collection.GetAggregateFunc(extensions.FunctionID{URN: urn, Name: "aggregate:i64"})
	require.True(t, ok)
	window, ok := collection.GetWindowFunc(extensions.FunctionID{URN: urn, Name: "window:i64"})
	require.True(t, ok)

	for name, resolve := range map[string]func([]types.Type, extensions.Set) (types.Type, error){
		"scalar": scalar.ResolveType, "aggregate": aggregate.ResolveType, "window": window.ResolveType,
	} {
		t.Run(name, func(t *testing.T) {
			registry := extensions.NewSet()
			registry.GetTypeAnchor(extensions.TypeID{URN: urn, Name: "sentinel"})
			result, err := resolve([]types.Type{&types.Int64Type{Nullability: types.NullabilityRequired}}, registry)
			require.NoError(t, err)
			require.Equal(t, &types.MapType{
				Nullability: types.NullabilityNullable,
				Key:         &types.StringType{Nullability: types.NullabilityRequired},
				Value: &types.UserDefinedType{
					Nullability:   types.NullabilityNullable,
					TypeReference: registry.GetTypeAnchor(extensions.TypeID{URN: urn, Name: "item"}),
				},
			}, result)
		})
	}
}

func TestEvaluateTypeExpressionPreservesArgumentUDTAnchors(t *testing.T) {
	const urn = "extension:example:return_types"
	const required = types.NullabilityRequired
	tests := []struct {
		expression string
		expected   func(types.Type, types.Type) types.Type
	}{
		{"any1", func(arg, _ types.Type) types.Type { return arg }},
		{"list<any1>", func(arg, _ types.Type) types.Type {
			return &types.ListType{Nullability: required, Type: arg}
		}},
		{"map<u!item, any1>", func(arg, declared types.Type) types.Type {
			return &types.MapType{Nullability: required, Key: declared, Value: arg}
		}},
		{"struct<any1, u!item>", func(arg, declared types.Type) types.Type {
			return &types.StructType{Nullability: required, Types: []types.Type{arg, declared}}
		}},
		{"u!item<any1>", func(arg, declared types.Type) types.Type {
			udt := *declared.(*types.UserDefinedType)
			udt.TypeParameters = []types.TypeParam{&types.DataTypeParameter{Type: arg}}
			return &udt
		}},
	}
	for _, tt := range tests {
		t.Run(tt.expression, func(t *testing.T) {
			registry := extensions.NewSet()
			registry.GetTypeAnchor(extensions.TypeID{URN: urn, Name: "sentinel"})
			// The argument's UDT has the same name but belongs to a different extension.
			foreign := &types.UserDefinedType{
				Nullability: required, TypeVariationRef: 7,
				TypeReference: registry.GetTypeAnchor(extensions.TypeID{URN: "extension:example:arguments", Name: "item"}),
			}
			arg := &types.ListType{Nullability: required, Type: foreign}
			before := proto.Clone(arg.ToProto())
			parameter, err := parser.ParseType("any1")
			require.NoError(t, err)
			returnType, err := parser.ParseType(tt.expression)
			require.NoError(t, err)
			result, err := extensions.EvaluateTypeExpression(urn, extensions.DeclaredOutputNullability,
				returnType, extensions.FuncParameterList{valArg(parameter)}, nil, []types.Type{arg}, registry)
			require.NoError(t, err)
			declared := &types.UserDefinedType{Nullability: required,
				TypeReference: registry.GetTypeAnchor(extensions.TypeID{URN: urn, Name: "item"})}
			require.Equal(t, tt.expected(arg, declared), result)
			require.True(t, proto.Equal(before, arg.ToProto()), "argument types must not be modified")
		})
	}
}
