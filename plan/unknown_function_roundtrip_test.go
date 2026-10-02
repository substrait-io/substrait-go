// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/plan"
	substraitproto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/encoding/protojson"
)

// A function the loaded collection does not hold is read as an on-the-fly
// variant, typed from the plan's own output_type. Its compound name is the
// variant's identity, so it has to survive a round trip whether or not the
// signature parses: "dec" does not, and rebuilding the name from the parsed
// arguments would write "no_such_function:" instead.
func TestUnknownFunctionRoundTripsUnderItsOwnName(t *testing.T) {
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

			var input substraitproto.Plan
			require.NoError(t, protojson.Unmarshal([]byte(fmt.Sprintf(template, name)), &input))

			p, err := plan.FromProto(&input, extensions.GetDefaultCollectionWithNoError())
			require.NoError(t, err)
			out, err := p.ToProto()
			require.NoError(t, err)

			require.Len(t, out.GetExtensions(), 1)
			assert.Equal(t, name, out.GetExtensions()[0].GetExtensionFunction().GetName())
		})
	}
}
