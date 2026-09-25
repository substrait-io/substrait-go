// SPDX-License-Identifier: Apache-2.0

package wire_test

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/plan"
	"github.com/substrait-io/substrait-go/v9/types"
	"github.com/substrait-io/substrait-go/v9/wire"
	substraitproto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const versionStruct = `"version": {
	"majorNumber": 0,
	"minorNumber": 29,
	"patchNumber": 0,
	"producer": "substrait-go"
}`

var baseSchema = types.NamedStruct{Names: []string{"a", "b"},
	Struct: types.StructType{
		Nullability: types.NullabilityRequired,
		Types: []types.Type{
			&types.StringType{Nullability: types.NullabilityRequired},
			&types.Float32Type{Nullability: types.NullabilityRequired},
		},
	}}

var baseSchema2 = types.NamedStruct{Names: []string{"x", "y"},
	Struct: types.StructType{
		Nullability: types.NullabilityRequired,
		Types: []types.Type{
			&types.Int32Type{Nullability: types.NullabilityRequired},
			&types.BooleanType{Nullability: types.NullabilityRequired},
		},
	}}

var baseSchemaReverse = types.NamedStruct{Names: []string{"x", "y"},
	Struct: types.StructType{
		Nullability: types.NullabilityRequired,
		Types: []types.Type{
			&types.Float32Type{Nullability: types.NullabilityRequired},
			&types.StringType{Nullability: types.NullabilityRequired},
		},
	}}

func checkRoundTrip(t *testing.T, expectedJSON string, p *plan.Plan) {
	t.Helper()
	protoPlan, err := wire.PlanToProto(p)
	require.NoError(t, err)

	var expectedProto substraitproto.Plan
	require.NoError(t, protojson.Unmarshal([]byte(expectedJSON), &expectedProto))

	// Equalize producer field; it may differ between golden JSON and protoPlan
	// depending on which OS (GOOS, ARCH, and the like) this test runs.
	protoPlan.Version.Producer = expectedProto.Version.Producer

	assert.Truef(t, proto.Equal(&expectedProto, protoPlan), "JSON expected: %s\ngot: %s",
		protojson.Format(&expectedProto), protojson.Format(protoPlan))

	roundTrip, err := wire.PlanFromProto(&expectedProto, extensions.GetDefaultCollectionWithNoError())
	require.NoError(t, err)

	roundTripProto, err := wire.PlanToProto(roundTrip)
	require.NoError(t, err)

	assert.Truef(t, proto.Equal(protoPlan, roundTripProto), "plan expected: %s\ngot: %s",
		protojson.Format(protoPlan), protojson.Format(roundTripProto))
}

func TestColumnlessVirtualTable(t *testing.T) {
	const expectedJSON = `{
		` + versionStruct + `,
		"relations": [
			{
				"root": {
					"input": {
						"read": {
							"common": {"direct":{}},
							"baseSchema": {
								"struct": {
									"nullability": "NULLABILITY_REQUIRED"
								}
							},
							"virtualTable": {
								"expressions": [
									{},
									{},
									{}
								]
							}
						}
					}
				}
			}
		]
	}`

	b := plan.NewBuilderDefault()

	virtual, err := b.VirtualTable(nil, make([]expr.StructLiteralValue, 3)...)
	require.NoError(t, err)

	p, err := b.Plan(virtual, []string{})
	require.NoError(t, err)

	checkRoundTrip(t, expectedJSON, p)
}

func TestEmptyVirtualTable(t *testing.T) {
	const expectedJSON = `{
		` + versionStruct + `,
		"relations": [
			{
				"root": {
					"input": {
						"read": {
							"common": {"direct":{}},
							"baseSchema": {
								"names": ["i"],
								"struct": {
									"types": [
										{"i32": {"nullability": "NULLABILITY_REQUIRED"}}
									],
									"nullability": "NULLABILITY_REQUIRED"
								}
							},
							"virtualTable": {}
						}
					},
					"names": ["i"]
				}
			}
		]
	}`

	b := plan.NewBuilderDefault()

	i32Type := types.Int32Type{Nullability: types.NullabilityRequired}
	virtual, err := b.EmptyVirtualTable([]string{"i"}, []types.Type{&i32Type})
	require.NoError(t, err)

	p, err := b.Plan(virtual, []string{"i"})
	require.NoError(t, err)

	checkRoundTrip(t, expectedJSON, p)
}

func expectedJsonWithIceberg(metadataURI string, snapshot plan.IcebergSnapshot) string {
	snapshotId, _ := snapshot.(plan.SnapshotId)
	snapshotTimestamp, _ := snapshot.(plan.SnapshotTimestamp)

	expectedJson := `{
		` + versionStruct + `,
		"relations": [
			{
				"root":  {
					"input":  {
						"read":  {
							"common":  {
								"direct":  {}
							},
							"baseSchema":  {
								"names":  [
									"a",
									"b"
								],
							  	"struct":  {
									"types":  [
								  		{
											"string":  {
											  	"nullability":  "NULLABILITY_REQUIRED"
											}
								  		},
									  	{
											"fp32":  {
										  		"nullability":  "NULLABILITY_REQUIRED"
											}
									  	}
									],
									"nullability":  "NULLABILITY_REQUIRED"
							  	}
							},
							"icebergTable":  {
								"direct":  {`
	// Add fields to icebergTable's direct node based on the snapshot type
	if snapshotId != "" {
		expectedJson += `
									"metadataUri": "` + metadataURI + `",
									"snapshotId": "` + string(snapshotId) + `"`
	} else if snapshotTimestamp != 0 {
		expectedJson += `
									"metadataUri": "` + metadataURI + `",
									"snapshotTimestamp": "` + strconv.FormatInt(int64(snapshotTimestamp), 10) + `"`
	} else {
		expectedJson += `
									"metadataUri": "` + metadataURI + `"`
	}
	// Add the rest of the JSON
	expectedJson += `			}
							}
						}
					},
					"names":  [
					  "a",
					  "b"
					]
				}
			}
		]
	}`
	return expectedJson
}

func TestIcebergTable(t *testing.T) {
	const metadataURI = "s3://bucket/path/to/metadata.json"

	for _, td := range []struct {
		name              string
		metadataURI       string
		snapshotId        plan.SnapshotId
		snapshotTimestamp plan.SnapshotTimestamp
	}{
		{"latest snapshot", metadataURI, "", 0},
		{"snapshot id", metadataURI, "SnapshotId0", 0},
		{"snapshot timestamp", metadataURI, "", 1010101},
	} {
		t.Run(td.name, func(t *testing.T) {
			b := plan.NewBuilderDefault()

			var snapshot plan.IcebergSnapshot
			if td.snapshotId != "" {
				snapshot = td.snapshotId
			} else if td.snapshotTimestamp != 0 {
				snapshot = td.snapshotTimestamp
			}

			iceberg, err := b.IcebergTableFromMetadataFile(td.metadataURI, snapshot, baseSchema)
			require.NoError(t, err)

			p, err := b.Plan(iceberg, []string{"a", "b"})
			require.NoError(t, err)

			checkRoundTrip(t, expectedJsonWithIceberg(td.metadataURI, snapshot), p)
		})
	}
}

// TestExtensionDefinition is a simple test implementation of ExtensionRelDefinition
type TestExtensionDefinition struct {
	schema types.RecordType
	detail []byte
	exprs  []expr.Expression
}

func (t *TestExtensionDefinition) Schema(inputs []plan.Rel) types.RecordType {
	return t.schema
}

func (t *TestExtensionDefinition) Build(inputs []plan.Rel) *anypb.Any {
	if t.detail == nil {
		return nil
	}
	message := &wrapperspb.StringValue{Value: string(t.detail)}
	any, _ := anypb.New(message)
	return any
}

func (t *TestExtensionDefinition) Expressions(inputs []plan.Rel) []expr.Expression {
	return t.exprs
}

func TestExtensionTable(t *testing.T) {
	const expectedJSON = `{
		` + versionStruct + `,
		"relations": [
			{
				"root": {
					"input": {
						"read": {
							"common": {"direct":{}},
							"baseSchema": {
								"names": ["a"],
								"struct": {
									"types": [
										{"i32": {"nullability": "NULLABILITY_REQUIRED"}}
									],
									"nullability": "NULLABILITY_REQUIRED"
								}
							},
							"extensionTable": {
								"detail": {
									"@type": "type.googleapis.com/google.protobuf.StringValue",
									"value": "my_custom_table"
								}
							}
						}
					},
					"names": ["a"]
				}
			}
		]
	}`

	b := plan.NewBuilderDefault()

	detail, err := anypb.New(wrapperspb.String("my_custom_table"))
	require.NoError(t, err)

	schema := types.NamedStruct{
		Names: []string{"a"},
		Struct: types.StructType{
			Nullability: types.NullabilityRequired,
			Types:       []types.Type{&types.Int32Type{Nullability: types.NullabilityRequired}},
		},
	}

	ext := b.ExtensionTable(detail, schema)
	p, err := b.Plan(ext, []string{"a"})
	require.NoError(t, err)

	checkRoundTrip(t, expectedJSON, p)
}
