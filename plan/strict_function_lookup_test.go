// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	substraitgo "github.com/substrait-io/substrait-go/v7"
	"github.com/substrait-io/substrait-go/v7/extensions"
	"github.com/substrait-io/substrait-go/v7/plan"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/encoding/protojson"
)

const arithmeticURN = extensions.SubstraitDefaultURNPrefix + "functions_arithmetic"

func TestStrictFunctionLookup_RegisteredFunction(t *testing.T) {
	// Plan with a registered function (add:i32_i32 exists in default extensions)
	const planJSON = `{
		"extensionUrns": [{"extensionUrnAnchor": 1, "urn": "` + arithmeticURN + `"}],
		"extensions": [
			{"extensionFunction": {"extensionUrnReference": 1, "functionAnchor": 1, "name": "add:i32_i32"}}
		],
		"relations": [{
			"root": {
				"input": {
					"project": {
						"input": {"read": {"baseSchema": {"struct": {"types": [{"i32": {}}]}}, "virtualTable": {"values": []}}},
						"expressions": [{
							"scalarFunction": {
								"functionReference": 1,
								"outputType": {"i32": {}},
								"arguments": [
									{"value": {"literal": {"i32": 1}}},
									{"value": {"literal": {"i32": 2}}}
								]
							}
						}]
					}
				}
			}
		}]
	}`

	var protoPlan proto.Plan
	require.NoError(t, protojson.Unmarshal([]byte(planJSON), &protoPlan))

	c := extensions.GetDefaultCollectionWithNoError().WithStrictFunctionLookup()
	p, err := plan.FromProto(&protoPlan, c)
	require.NoError(t, err)
	assert.NotNil(t, p)
}

func TestStrictFunctionLookup_UnregisteredScalarFunction(t *testing.T) {
	// Plan with an unregistered function (add:i32_string doesn't exist)
	const planJSON = `{
		"extensionUrns": [{"extensionUrnAnchor": 1, "urn": "` + arithmeticURN + `"}],
		"extensions": [
			{"extensionFunction": {"extensionUrnReference": 1, "functionAnchor": 1, "name": "add:i32_string"}}
		],
		"relations": [{
			"root": {
				"input": {
					"project": {
						"input": {"read": {"baseSchema": {"struct": {"types": [{"i32": {}}]}}, "virtualTable": {"values": []}}},
						"expressions": [{
							"scalarFunction": {
								"functionReference": 1,
								"outputType": {"i32": {}},
								"arguments": [
									{"value": {"literal": {"i32": 1}}},
									{"value": {"literal": {"string": "foo"}}}
								]
							}
						}]
					}
				}
			}
		}]
	}`

	var protoPlan proto.Plan
	require.NoError(t, protojson.Unmarshal([]byte(planJSON), &protoPlan))

	t.Run("strict mode returns error", func(t *testing.T) {
		c := extensions.GetDefaultCollectionWithNoError().WithStrictFunctionLookup()
		_, err := plan.FromProto(&protoPlan, c)
		require.Error(t, err)
		assert.True(t, errors.Is(err, substraitgo.ErrUnregisteredFunction))
		assert.Contains(t, err.Error(), "add:i32_string")
	})

	t.Run("non-strict mode succeeds", func(t *testing.T) {
		c := extensions.GetDefaultCollectionWithNoError()
		p, err := plan.FromProto(&protoPlan, c)
		require.NoError(t, err)
		assert.NotNil(t, p)
	})
}

func TestStrictFunctionLookup_UnregisteredWindowFunction(t *testing.T) {
	// Plan with an unregistered window function (row_number:i32 doesn't exist - row_number takes no args)
	const planJSON = `{
		"extensionUrns": [{"extensionUrnAnchor": 1, "urn": "` + arithmeticURN + `"}],
		"extensions": [
			{"extensionFunction": {"extensionUrnReference": 1, "functionAnchor": 1, "name": "row_number:i32"}}
		],
		"relations": [{
			"root": {
				"input": {
					"project": {
						"input": {"read": {"baseSchema": {"struct": {"types": [{"i32": {}}]}}, "virtualTable": {"values": []}}},
						"expressions": [{
							"windowFunction": {
								"functionReference": 1,
								"outputType": {"i64": {}},
								"arguments": [{"value": {"literal": {"i32": 1}}}],
								"boundsType": "BOUNDS_TYPE_ROWS"
							}
						}]
					}
				}
			}
		}]
	}`

	var protoPlan proto.Plan
	require.NoError(t, protojson.Unmarshal([]byte(planJSON), &protoPlan))

	t.Run("strict mode returns error", func(t *testing.T) {
		c := extensions.GetDefaultCollectionWithNoError().WithStrictFunctionLookup()
		_, err := plan.FromProto(&protoPlan, c)
		require.Error(t, err)
		assert.True(t, errors.Is(err, substraitgo.ErrUnregisteredFunction))
		assert.Contains(t, err.Error(), "row_number:i32")
	})

	t.Run("non-strict mode succeeds", func(t *testing.T) {
		c := extensions.GetDefaultCollectionWithNoError()
		p, err := plan.FromProto(&protoPlan, c)
		require.NoError(t, err)
		assert.NotNil(t, p)
	})
}

func TestWithStrictFunctionLookup_DoesNotMutateOriginal(t *testing.T) {
	c := extensions.GetDefaultCollectionWithNoError()
	strict := c.WithStrictFunctionLookup()

	assert.False(t, c.StrictFunctionLookup())
	assert.True(t, strict.StrictFunctionLookup())
}
