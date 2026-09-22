// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/plan"
	substraitproto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/encoding/protojson"
)

// fetchRelFromJSON parses a plan JSON string and returns the FetchRel from the first root relation.
func fetchRelFromJSON(t *testing.T, jsonStr string) *plan.FetchRel {
	t.Helper()
	var proto substraitproto.Plan
	require.NoError(t, protojson.Unmarshal([]byte(jsonStr), &proto))
	p, err := plan.FromProto(&proto, extensions.GetDefaultCollectionWithNoError())
	require.NoError(t, err)
	fetch, ok := p.GetRoots()[0].Input().(*plan.FetchRel)
	require.True(t, ok, "expected root relation to be a FetchRel")
	return fetch
}

// TestFetchRelMigration_LegacyOffset verifies that a FetchRel with a legacy integer offset
// is parsed into a Go struct with the correct offset expression and no count.
func TestFetchRelMigration_LegacyOffset(t *testing.T) {
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
						},
						"offset": "42"
					}
				},
				"names": ["a"]
			}
		}]
	}`

	fetch := fetchRelFromJSON(t, jsonStr)

	require.NotNil(t, fetch.Offset())
	assert.Equal(t, expr.NewPrimitiveLiteral(int64(42), false), fetch.Offset())
	assert.Nil(t, fetch.Count())
}

// TestFetchRelMigration_LegacyCount verifies that a FetchRel with a legacy integer count
// is parsed into a Go struct with no offset and the correct count expression.
func TestFetchRelMigration_LegacyCount(t *testing.T) {
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
						},
						"count": "10"
					}
				},
				"names": ["a"]
			}
		}]
	}`

	fetch := fetchRelFromJSON(t, jsonStr)

	assert.Nil(t, fetch.Offset())
	require.NotNil(t, fetch.Count())
	assert.Equal(t, expr.NewPrimitiveLiteral(int64(10), false), fetch.Count())
}

// TestFetchRelMigration_LegacyBoth verifies that a FetchRel with both legacy integer offset
// and count fields is parsed correctly.
func TestFetchRelMigration_LegacyBoth(t *testing.T) {
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
						},
						"offset": "5",
						"count": "20"
					}
				},
				"names": ["a"]
			}
		}]
	}`

	fetch := fetchRelFromJSON(t, jsonStr)

	require.NotNil(t, fetch.Offset())
	assert.Equal(t, expr.NewPrimitiveLiteral(int64(5), false), fetch.Offset())
	require.NotNil(t, fetch.Count())
	assert.Equal(t, expr.NewPrimitiveLiteral(int64(20), false), fetch.Count())
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
