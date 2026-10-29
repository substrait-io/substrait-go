// SPDX-License-Identifier: Apache-2.0

package substraitproto_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/encoding/substraitproto"
	substraitgo "github.com/substrait-io/substrait-go/v9"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestPlanRejectsUnknownFunction(t *testing.T) {
	for _, name := range []string{"no_such_function:dec_dec", "no_such_function:i64_i64"} {
		t.Run(name, func(t *testing.T) {
			const template = `{
				` + versionStruct + `,
				"extensionUrns": [{"extensionUrnAnchor": 1, "urn": "extension:io.substrait:functions_arithmetic_decimal"}],
				"extensions": [{"extensionFunction": {"functionAnchor": 1, "name": "%s", "extensionUrnReference": 1}}],
				"relations": [{"root": {"names": ["out"], "input": {"project": {
					"common": {"emit": {"outputMapping": [1]}},
					"input": {"read": {
						"baseSchema": {"names": ["c0"], "struct": {"types": [{"i64": {"nullability": "NULLABILITY_REQUIRED"}}], "nullability": "NULLABILITY_REQUIRED"}},
						"namedTable": {"names": ["t"]}}},
					"expressions": [{"scalarFunction": {
						"functionReference": 1,
						"outputType": {"i64": {"nullability": "NULLABILITY_NULLABLE"}},
						"arguments": [{"value": {"selection": {"directReference": {"structField": {"field": 0}}, "rootReference": {}}}}]}}]}}}}]
			}`

			var input proto.Plan
			require.NoError(t, protojson.Unmarshal([]byte(fmt.Sprintf(template, name)), &input))

			p, err := substraitproto.PlanFromProto(&input, extensions.GetDefaultCollectionWithNoError())
			require.ErrorIs(t, err, substraitgo.ErrNotFound)
			require.ErrorContains(t, err, name)
			require.ErrorContains(t, err, "extension:io.substrait:functions_arithmetic_decimal")
			require.Nil(t, p)
		})
	}
}

func TestFunctionDecodersRequireLoadedDefinitions(t *testing.T) {
	const urn = "extension:test:functions"
	for _, kind := range []string{"scalar", "aggregate", "window"} {
		t.Run(kind, func(t *testing.T) {
			yaml := fmt.Sprintf(`urn: %s
%s_functions:
  - name: known
    impls:
      - args:
          - value: i64
        return: i64
`, urn, kind)
			var collection extensions.Collection
			require.NoError(t, collection.Load(strings.NewReader(yaml)))
			reg := expr.NewEmptyExtensionRegistry(&collection)
			decode := func(name string) (any, error) {
				anchor := reg.GetFuncAnchor(extensions.FunctionID{URN: urn, Name: name})
				body := fmt.Sprintf(`{"functionReference": %d, "outputType": {"i64": {}}, "arguments": [{"value": {"literal": {"i64": "1"}}}]}`, anchor)
				if kind == "aggregate" {
					var input proto.AggregateFunction
					require.NoError(t, protojson.Unmarshal([]byte(body), &input))
					return substraitproto.AggregateFunctionFromProto(&input, nil, reg)
				}
				var input proto.Expression
				require.NoError(t, protojson.Unmarshal([]byte(fmt.Sprintf(`{"%sFunction": %s}`, kind, body)), &input))
				return substraitproto.ExprFromProto(&input, nil, reg)
			}
			for _, name := range []string{"no_such:dec_dec", "no_such:list", "no_such:i64", "no_such"} {
				t.Run(name, func(t *testing.T) {
					_, err := decode(name)
					require.ErrorIs(t, err, substraitgo.ErrNotFound)
					require.ErrorContains(t, err, kind+" function")
					require.ErrorContains(t, err, urn)
					require.ErrorContains(t, err, name)
				})
			}
			loaded, err := decode("known:i64")
			require.NoError(t, err)
			require.NotNil(t, loaded)
		})
	}
}
