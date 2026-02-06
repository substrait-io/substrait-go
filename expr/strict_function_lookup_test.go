// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	substraitgo "github.com/substrait-io/substrait-go/v7"
	"github.com/substrait-io/substrait-go/v7/expr"
	"github.com/substrait-io/substrait-go/v7/extensions"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/encoding/protojson"
)

const strictTestExtensionYAML = `---
urn: extension:test:strict_functions
scalar_functions:
  -
    name: "add"
    impls:
      - args:
          - name: x
            value: i32
          - name: y
            value: i32
        return: i32
aggregate_functions:
  -
    name: "sum"
    impls:
      - args:
          - name: x
            value: i64
        return: i64
window_functions:
  -
    name: "row_number"
    impls:
      - args: []
        return: i64
`

func newTestCollection(t *testing.T) *extensions.Collection {
	t.Helper()
	var c extensions.Collection
	err := c.Load("extension:test:strict_functions", strings.NewReader(strictTestExtensionYAML))
	require.NoError(t, err)
	return &c
}

func TestStrictFunctionLookup_ScalarFunction_Registered(t *testing.T) {
	c := newTestCollection(t)

	// Plan with a registered function (add:i32_i32)
	const planJSON = `{
		"extensionUrns": [{"extensionUrnAnchor": 1, "urn": "extension:test:strict_functions"}],
		"extensions": [
			{"extensionFunction": {"extensionUrnReference": 1, "functionAnchor": 1, "name": "add:i32_i32"}}
		],
		"relations": []
	}`

	var plan proto.Plan
	require.NoError(t, protojson.Unmarshal([]byte(planJSON), &plan))

	extSet, err := extensions.GetExtensionSet(&plan, c)
	require.NoError(t, err)

	// Test with strict mode enabled - should succeed for registered function
	reg := expr.NewExtensionRegistry(extSet, c).WithStrictFunctionLookup()
	assert.True(t, reg.StrictFunctionLookup())

	const scalarFuncJSON = `{
		"scalarFunction": {
			"functionReference": 1,
			"outputType": {"i32": {}},
			"arguments": [
				{"value": {"literal": {"i32": 1}}},
				{"value": {"literal": {"i32": 2}}}
			]
		}
	}`

	var exprProto proto.Expression
	require.NoError(t, protojson.Unmarshal([]byte(scalarFuncJSON), &exprProto))

	result, err := expr.ExprFromProto(&exprProto, nil, reg)
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestStrictFunctionLookup_ScalarFunction_Unregistered(t *testing.T) {
	c := newTestCollection(t)

	// Plan with an unregistered function (add:i32_string - invalid signature)
	const planJSON = `{
		"extensionUrns": [{"extensionUrnAnchor": 1, "urn": "extension:test:strict_functions"}],
		"extensions": [
			{"extensionFunction": {"extensionUrnReference": 1, "functionAnchor": 1, "name": "add:i32_string"}}
		],
		"relations": []
	}`

	var plan proto.Plan
	require.NoError(t, protojson.Unmarshal([]byte(planJSON), &plan))

	extSet, err := extensions.GetExtensionSet(&plan, c)
	require.NoError(t, err)

	const scalarFuncJSON = `{
		"scalarFunction": {
			"functionReference": 1,
			"outputType": {"i32": {}},
			"arguments": [
				{"value": {"literal": {"i32": 1}}},
				{"value": {"literal": {"string": "foo"}}}
			]
		}
	}`

	var exprProto proto.Expression
	require.NoError(t, protojson.Unmarshal([]byte(scalarFuncJSON), &exprProto))

	t.Run("strict mode returns error", func(t *testing.T) {
		reg := expr.NewExtensionRegistry(extSet, c).WithStrictFunctionLookup()

		_, err := expr.ExprFromProto(&exprProto, nil, reg)
		require.Error(t, err)
		assert.True(t, errors.Is(err, substraitgo.ErrUnregisteredFunction))
		assert.Contains(t, err.Error(), "add:i32_string")
	})

	t.Run("non-strict mode creates custom variant", func(t *testing.T) {
		reg := expr.NewExtensionRegistry(extSet, c)
		assert.False(t, reg.StrictFunctionLookup())

		result, err := expr.ExprFromProto(&exprProto, nil, reg)
		require.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestStrictFunctionLookup_AggregateFunction_Unregistered(t *testing.T) {
	c := newTestCollection(t)

	// Plan with an unregistered aggregate function
	const planJSON = `{
		"extensionUrns": [{"extensionUrnAnchor": 1, "urn": "extension:test:strict_functions"}],
		"extensions": [
			{"extensionFunction": {"extensionUrnReference": 1, "functionAnchor": 1, "name": "sum:string"}}
		],
		"relations": []
	}`

	var plan proto.Plan
	require.NoError(t, protojson.Unmarshal([]byte(planJSON), &plan))

	extSet, err := extensions.GetExtensionSet(&plan, c)
	require.NoError(t, err)

	aggFunc := &proto.AggregateFunction{
		FunctionReference: 1,
		OutputType:        &proto.Type{Kind: &proto.Type_I64_{I64: &proto.Type_I64{}}},
		Arguments: []*proto.FunctionArgument{
			{ArgType: &proto.FunctionArgument_Value{Value: &proto.Expression{
				RexType: &proto.Expression_Literal_{Literal: &proto.Expression_Literal{
					LiteralType: &proto.Expression_Literal_String_{String_: "test"},
				}},
			}}},
		},
	}

	t.Run("strict mode returns error", func(t *testing.T) {
		reg := expr.NewExtensionRegistry(extSet, c).WithStrictFunctionLookup()

		_, err := expr.NewAggregateFunctionFromProto(aggFunc, nil, reg)
		require.Error(t, err)
		assert.True(t, errors.Is(err, substraitgo.ErrUnregisteredFunction))
		assert.Contains(t, err.Error(), "sum:string")
	})

	t.Run("non-strict mode creates custom variant", func(t *testing.T) {
		reg := expr.NewExtensionRegistry(extSet, c)

		result, err := expr.NewAggregateFunctionFromProto(aggFunc, nil, reg)
		require.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestStrictFunctionLookup_WindowFunction_Unregistered(t *testing.T) {
	c := newTestCollection(t)

	// Plan with an unregistered window function
	const planJSON = `{
		"extensionUrns": [{"extensionUrnAnchor": 1, "urn": "extension:test:strict_functions"}],
		"extensions": [
			{"extensionFunction": {"extensionUrnReference": 1, "functionAnchor": 1, "name": "row_number:i32"}}
		],
		"relations": []
	}`

	var plan proto.Plan
	require.NoError(t, protojson.Unmarshal([]byte(planJSON), &plan))

	extSet, err := extensions.GetExtensionSet(&plan, c)
	require.NoError(t, err)

	const windowFuncJSON = `{
		"windowFunction": {
			"functionReference": 1,
			"outputType": {"i64": {}},
			"arguments": [
				{"value": {"literal": {"i32": 1}}}
			],
			"partitions": [],
			"sorts": [],
			"boundsType": "BOUNDS_TYPE_ROWS"
		}
	}`

	var exprProto proto.Expression
	require.NoError(t, protojson.Unmarshal([]byte(windowFuncJSON), &exprProto))

	t.Run("strict mode returns error", func(t *testing.T) {
		reg := expr.NewExtensionRegistry(extSet, c).WithStrictFunctionLookup()

		_, err := expr.ExprFromProto(&exprProto, nil, reg)
		require.Error(t, err)
		assert.True(t, errors.Is(err, substraitgo.ErrUnregisteredFunction))
		assert.Contains(t, err.Error(), "row_number:i32")
	})

	t.Run("non-strict mode creates custom variant", func(t *testing.T) {
		reg := expr.NewExtensionRegistry(extSet, c)

		result, err := expr.ExprFromProto(&exprProto, nil, reg)
		require.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestWithStrictFunctionLookup_PreservesOtherFields(t *testing.T) {
	c := newTestCollection(t)

	const planJSON = `{
		"extensionUrns": [{"extensionUrnAnchor": 1, "urn": "extension:test:strict_functions"}],
		"extensions": [
			{"extensionFunction": {"extensionUrnReference": 1, "functionAnchor": 1, "name": "add:i32_i32"}}
		],
		"relations": []
	}`

	var plan proto.Plan
	require.NoError(t, protojson.Unmarshal([]byte(planJSON), &plan))

	extSet, err := extensions.GetExtensionSet(&plan, c)
	require.NoError(t, err)

	original := expr.NewExtensionRegistry(extSet, c)
	strict := original.WithStrictFunctionLookup()

	// Verify the extension set is preserved
	id, ok := strict.DecodeFunc(1)
	require.True(t, ok)
	assert.Equal(t, "add:i32_i32", id.Name)
	assert.Equal(t, "extension:test:strict_functions", id.URN)

	// Verify strict mode is enabled on the copy
	assert.True(t, strict.StrictFunctionLookup())

	// Verify original is unchanged
	assert.False(t, original.StrictFunctionLookup())
}
