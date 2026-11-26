// SPDX-License-Identifier: Apache-2.0

package substraitproto_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/encoding/substraitproto"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/plan"
	substraitpb "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/encoding/protojson"
)

// fetchRelFromJSON parses a plan JSON string and returns the FetchRel from the first root relation.
func fetchRelFromJSON(t *testing.T, jsonStr string) *plan.FetchRel {
	t.Helper()
	var proto substraitpb.Plan
	require.NoError(t, protojson.Unmarshal([]byte(jsonStr), &proto))
	p, err := substraitproto.PlanFromProto(&proto, extensions.GetDefaultCollectionWithNoError())
	require.NoError(t, err)
	fetch, ok := p.GetRoots()[0].Input().(*plan.FetchRel)
	require.True(t, ok, "expected root relation to be a FetchRel")
	return fetch
}

// TestFetchRelMigration_Neither verifies that a FetchRel with neither offset nor count
// (returns all rows from the start) has both fields nil.
func TestFetchRelMigration_Neither(t *testing.T) {
	const jsonStr = `{
		` + versionStruct + `,
		"relations": [{
			"root": {
				"input": {
					"fetch": {
						"common": {"direct": {}},
						"input": {
							"read": {
								"common": {"direct": {}},
								"baseSchema": {
									"names": ["a"],
									"struct": {
										"nullability": "NULLABILITY_REQUIRED",
										"types": [{"i32": {"nullability": "NULLABILITY_REQUIRED"}}]
									}
								},
								"namedTable": {"names": ["test"]}
							}
						}
					}
				},
				"names": ["a"]
			}
		}]
	}`

	fetch := fetchRelFromJSON(t, jsonStr)

	assert.Nil(t, fetch.Offset())
	assert.Nil(t, fetch.Count())
}
