package plan

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v6/expr"
	"github.com/substrait-io/substrait-go/v6/extensions"
	"github.com/substrait-io/substrait-go/v6/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	protoext "github.com/substrait-io/substrait-protobuf/go/substraitpb/extensions"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/testing/protocmp"
)

// This is a test where we will load in an extension with both a URI and URN present.
// We then make sure that a produced plan contains the appropriate data present for both
// URI and URN (e.g. correct anchors for both present) by comparing the produced plan
// with an expected output plan. This test can be dropped when the URI -> URN migration
// is complete.
func TestExtensionURNAndURIInPlanProtobuf(t *testing.T) {

	const uri = "http://localhost/test.yaml"
	const extensionYAML = `---
urn: "urn:example:test"
scalar_functions:
  - name: "add"
    impls:
      - args:
          - name: x
            value: i32
          - name: y
            value: i32
        return: i32
`

	var collection extensions.Collection
	require.NoError(t, collection.Load(uri, strings.NewReader(extensionYAML)))

	builder := NewBuilder(&collection)
	baseSchema := types.NamedStruct{
		Names:  []string{"a", "b"},
		Struct: types.StructType{Types: []types.Type{&types.Int32Type{}, &types.Int32Type{}}},
	}
	scan := builder.NamedScan([]string{"test_table"}, baseSchema)
	exprBuilder := builder.GetExprBuilder()
	exprBuilder.BaseSchema = types.NewRecordTypeFromStruct(baseSchema.Struct)

	scalarFunc := exprBuilder.ScalarFunc(extensions.ID{
		URI:  uri,
		Name: "add",
	}).Args(
		exprBuilder.RootRef(expr.NewStructFieldRef(0)),
		exprBuilder.RootRef(expr.NewStructFieldRef(1)),
	)

	addExpr, err := scalarFunc.BuildExpr()
	require.NoError(t, err)

	projected, err := builder.Project(scan, addExpr)
	require.NoError(t, err)

	planObj, err := builder.Plan(projected, []string{"a", "b", "result"})
	require.NoError(t, err)

	gotProto, err := planObj.ToProto()
	require.NoError(t, err)

	const expectedJSON = `{
		"version": {
			"minorNumber": 29,
			"producer": "substrait-go darwin/arm64"
		},
		"extensionUris": [{
			"extensionUriAnchor": 1,
			"uri": "http://localhost/test.yaml"
		}],
		"extensionUrns": [{
			"extensionUrnAnchor": 1,
			"urn": "urn:example:test"
		}],
		"extensions": [{
			"extensionFunction": {
				"extensionUriReference": 1,
				"extensionUrnReference": 1,
				"functionAnchor": 1,
				"name": "add:i32_i32"
			}
		}],
		"relations": [{
			"root": {
				"input": {
					"project": {
						"common": {"direct": {}},
						"input": {
							"read": {
								"common": {"direct": {}},
								"baseSchema": {
									"names": ["a", "b"],
									"struct": {
										"types": [{"i32": {}}, {"i32": {}}]
									}
								},
								"namedTable": {"names": ["test_table"]}
							}
						},
						"expressions": [{
							"scalarFunction": {
								"functionReference": 1,
								"arguments": [
									{
										"value": {
											"selection": {
												"directReference": {"structField": {}},
												"rootReference": {}
											}
										}
									},
									{
										"value": {
											"selection": {
												"directReference": {"structField": {"field": 1}},
												"rootReference": {}
											}
										}
									}
								],
								"outputType": {
									"i32": {
										"nullability": "NULLABILITY_NULLABLE"
									}
								}
							}
						}]
					}
				},
				"names": ["a", "b", "result"]
			}
		}]
	}`

	var expected proto.Plan
	require.NoError(t, protojson.Unmarshal([]byte(expectedJSON), &expected))

	if diff := cmp.Diff(&expected, gotProto, protocmp.Transform()); diff != "" {
		t.Errorf("Plan protobuf mismatch (-want +got):\n%s", diff)
	}

	assert.True(t, collection.URILoaded(uri))
	assert.True(t, collection.URNLoaded("urn:example:test"))
}

func TestExtensionMixedURIURNReferences(t *testing.T) {
	// Test with four extensions to ensure different URI/URN anchor values:
	// 1. URI-only extension
	// 2. URN-only extension
	// 3. Another URN-only extension (to offset anchor numbering)
	// 4. URI+URN extension

	var collection extensions.Collection

	// Extension 1: URI-only (no URN in YAML)
	const uriOnlyURI = "http://localhost/uri-only.yaml"
	const uriOnlyYAML = `---
scalar_functions:
  - name: "add_uri"
    impls:
      - args:
          - name: x
            value: i32
          - name: y
            value: i32
        return: i32
`

	// Extension 2: URN-only (loaded with Load, not LoadWithURI)
	const urnOnlyYAML = `---
urn: "urn:example:urn-only"
scalar_functions:
  - name: "add_urn"
    impls:
      - args:
          - name: x
            value: i32
          - name: y
            value: i32
        return: i32
`

	// Extension 3: Another URN-only to create different anchor numbering
	const urnOnly2YAML = `---
urn: "urn:example:urn-only-2"
scalar_functions:
  - name: "multiply_urn"
    impls:
      - args:
          - name: x
            value: i32
          - name: y
            value: i32
        return: i32
`

	// Extension 4: Both URI and URN
	const bothURI = "http://localhost/both.yaml"
	const bothYAML = `---
urn: "urn:example:both"
scalar_functions:
  - name: "add_both"
    impls:
      - args:
          - name: x
            value: i32
          - name: y
            value: i32
        return: i32
`

	require.NoError(t, collection.LoadWithoutUri(strings.NewReader(urnOnlyYAML)))
	require.NoError(t, collection.LoadWithoutUri(strings.NewReader(urnOnly2YAML)))
	require.NoError(t, collection.Load(uriOnlyURI, strings.NewReader(uriOnlyYAML)))
	require.NoError(t, collection.Load(bothURI, strings.NewReader(bothYAML)))

	builder := NewBuilder(&collection)

	baseSchema := types.NamedStruct{
		Names:  []string{"a", "b"},
		Struct: types.StructType{Types: []types.Type{&types.Int32Type{}, &types.Int32Type{}}},
	}
	scan := builder.NamedScan([]string{"test_table"}, baseSchema)

	exprBuilder := builder.GetExprBuilder()
	exprBuilder.BaseSchema = types.NewRecordTypeFromStruct(baseSchema.Struct)

	// Function 1: From URI-only extension
	addUriExpr, err := exprBuilder.ScalarFunc(extensions.ID{
		URI:  uriOnlyURI,
		Name: "add_uri",
	}).Args(
		exprBuilder.RootRef(expr.NewStructFieldRef(0)),
		exprBuilder.RootRef(expr.NewStructFieldRef(1)),
	).BuildExpr()
	require.NoError(t, err)

	// Function 2: From URN-only extension
	addUrnExpr, err := exprBuilder.ScalarFunc(extensions.ID{
		URN:  "urn:example:urn-only",
		Name: "add_urn",
	}).Args(
		exprBuilder.RootRef(expr.NewStructFieldRef(0)),
		exprBuilder.RootRef(expr.NewStructFieldRef(1)),
	).BuildExpr()
	require.NoError(t, err)

	// Function 3: From second URN-only extension
	multiplyUrnExpr, err := exprBuilder.ScalarFunc(extensions.ID{
		URN:  "urn:example:urn-only-2",
		Name: "multiply_urn",
	}).Args(
		exprBuilder.RootRef(expr.NewStructFieldRef(0)),
		exprBuilder.RootRef(expr.NewStructFieldRef(1)),
	).BuildExpr()
	require.NoError(t, err)

	// Function 4: From URI+URN extension (lookup by URI)
	addBothExpr, err := exprBuilder.ScalarFunc(extensions.ID{
		URI:  bothURI,
		Name: "add_both",
	}).Args(
		exprBuilder.RootRef(expr.NewStructFieldRef(0)),
		exprBuilder.RootRef(expr.NewStructFieldRef(1)),
	).BuildExpr()
	require.NoError(t, err)

	// Create a projection with all four function calls
	projected, err := builder.Project(scan, addUriExpr, addUrnExpr, multiplyUrnExpr, addBothExpr)
	require.NoError(t, err)

	planObj, err := builder.Plan(projected, []string{"a", "b", "result_uri", "result_urn", "result_multiply", "result_both"})
	require.NoError(t, err)

	gotProto, err := planObj.ToProto()
	require.NoError(t, err)

	const expectedJSON = `{
		"version": {
			"minorNumber": 29,
			"producer": "substrait-go darwin/arm64"
		},
		"extensionUris": [
			{
				"extensionUriAnchor": 1,
				"uri": "http://localhost/uri-only.yaml"
			},
			{
				"extensionUriAnchor": 2,
				"uri": "http://localhost/both.yaml"
			}
		],
		"extensionUrns": [
			{
				"extensionUrnAnchor": 1,
				"urn": "urn:example:urn-only"
			},
			{
				"extensionUrnAnchor": 2,
				"urn": "urn:example:urn-only-2"
			},
			{
				"extensionUrnAnchor": 3,
				"urn": "urn:example:both"
			}
		],
		"extensions": [
			{
				"extensionFunction": {
					"extensionUriReference": 1,
					"functionAnchor": 1,
					"name": "add_uri:i32_i32"
				}
			},
			{
				"extensionFunction": {
					"extensionUrnReference": 1,
					"functionAnchor": 2,
					"name": "add_urn:i32_i32"
				}
			},
			{
				"extensionFunction": {
					"extensionUrnReference": 2,
					"functionAnchor": 3,
					"name": "multiply_urn:i32_i32"
				}
			},
			{
				"extensionFunction": {
					"extensionUriReference": 2,
					"extensionUrnReference": 3,
					"functionAnchor": 4,
					"name": "add_both:i32_i32"
				}
			}
		],
		"relations": [{
			"root": {
				"input": {
					"project": {
						"common": {"direct": {}},
						"input": {
							"read": {
								"common": {"direct": {}},
								"baseSchema": {
									"names": ["a", "b"],
									"struct": {
										"types": [{"i32": {}}, {"i32": {}}]
									}
								},
								"namedTable": {"names": ["test_table"]}
							}
						},
						"expressions": [
							{
								"scalarFunction": {
									"functionReference": 1,
									"arguments": [
										{
											"value": {
												"selection": {
													"directReference": {"structField": {}},
													"rootReference": {}
												}
											}
										},
										{
											"value": {
												"selection": {
													"directReference": {"structField": {"field": 1}},
													"rootReference": {}
												}
											}
										}
									],
									"outputType": {
										"i32": {
											"nullability": "NULLABILITY_NULLABLE"
										}
									}
								}
							},
							{
								"scalarFunction": {
									"functionReference": 2,
									"arguments": [
										{
											"value": {
												"selection": {
													"directReference": {"structField": {}},
													"rootReference": {}
												}
											}
										},
										{
											"value": {
												"selection": {
													"directReference": {"structField": {"field": 1}},
													"rootReference": {}
												}
											}
										}
									],
									"outputType": {
										"i32": {
											"nullability": "NULLABILITY_NULLABLE"
										}
									}
								}
							},
							{
								"scalarFunction": {
									"functionReference": 3,
									"arguments": [
										{
											"value": {
												"selection": {
													"directReference": {"structField": {}},
													"rootReference": {}
												}
											}
										},
										{
											"value": {
												"selection": {
													"directReference": {"structField": {"field": 1}},
													"rootReference": {}
												}
											}
										}
									],
									"outputType": {
										"i32": {
											"nullability": "NULLABILITY_NULLABLE"
										}
									}
								}
							},
							{
								"scalarFunction": {
									"functionReference": 4,
									"arguments": [
										{
											"value": {
												"selection": {
													"directReference": {"structField": {}},
													"rootReference": {}
												}
											}
										},
										{
											"value": {
												"selection": {
													"directReference": {"structField": {"field": 1}},
													"rootReference": {}
												}
											}
										}
									],
									"outputType": {
										"i32": {
											"nullability": "NULLABILITY_NULLABLE"
										}
									}
								}
							}
						]
					}
				},
				"names": ["a", "b", "result_uri", "result_urn", "result_multiply", "result_both"]
			}
		}]
	}`

	var expected proto.Plan
	require.NoError(t, protojson.Unmarshal([]byte(expectedJSON), &expected))

	if diff := cmp.Diff(&expected, gotProto, protocmp.Transform()); diff != "" {
		t.Errorf("Plan protobuf mismatch (-want +got):\n%s", diff)
	}

	assert.True(t, collection.URILoaded(uriOnlyURI))
	assert.True(t, collection.URNLoaded("urn:example:urn-only"))
	assert.True(t, collection.URILoaded(bothURI))
	assert.True(t, collection.URNLoaded("urn:example:both"))
}

func TestFunctionConstructionViaURIAndURN(t *testing.T) {
	// Test that functions can be made up by both URI and URN
	var collection extensions.Collection

	// Simple test extension with both URI and URN
	testURI := "https://example.com/test.yaml"
	testYAML := `---
urn: "urn:example:test"
scalar_functions:
  - name: "add_test"
    impls:
      - args:
          - name: x
            value: i32
          - name: y
            value: i32
        return: i32
`

	err := collection.Load(testURI, strings.NewReader(testYAML))
	require.NoError(t, err)

	builder := NewBuilder(&collection)
	baseSchema := types.NamedStruct{
		Names:  []string{"a", "b"},
		Struct: types.StructType{Types: []types.Type{&types.Int32Type{}, &types.Int32Type{}}},
	}
	scan := builder.NamedScan([]string{"test_table"}, baseSchema)

	exprBuilder := builder.GetExprBuilder()
	exprBuilder.BaseSchema = types.NewRecordTypeFromStruct(baseSchema.Struct)

	// Test function construction via URI
	funcByURIExpr, err := exprBuilder.ScalarFunc(extensions.ID{
		URI:  testURI,
		Name: "add_test",
	}).Args(
		exprBuilder.RootRef(expr.NewStructFieldRef(0)),
		exprBuilder.RootRef(expr.NewStructFieldRef(1)),
	).BuildExpr()
	require.NoError(t, err)

	// Test function construction via URN
	funcByURNExpr, err := exprBuilder.ScalarFunc(extensions.ID{
		URN:  "urn:example:test",
		Name: "add_test",
	}).Args(
		exprBuilder.RootRef(expr.NewStructFieldRef(0)),
		exprBuilder.RootRef(expr.NewStructFieldRef(1)),
	).BuildExpr()
	require.NoError(t, err)

	projectedURI, err := builder.Project(scan, funcByURIExpr)
	require.NoError(t, err)
	projectedURN, err := builder.Project(scan, funcByURNExpr)
	require.NoError(t, err)

	planURI, err := builder.Plan(projectedURI, []string{"a", "b", "result"})
	require.NoError(t, err)
	planURN, err := builder.Plan(projectedURN, []string{"a", "b", "result"})
	require.NoError(t, err)

	protoURI, err := planURI.ToProto()
	require.NoError(t, err)
	protoURN, err := planURN.ToProto()
	require.NoError(t, err)

	// Both should have references to the same extension but with different anchor types
	assert.NotNil(t, protoURI)
	assert.NotNil(t, protoURN)

	// The plans should be functionally identical since they use the same function
	if diff := cmp.Diff(protoURI, protoURN, protocmp.Transform()); diff != "" {
		t.Errorf("Plans should be identical when using URI vs URN lookup (-URI +URN):\n%s", diff)
	}

	// Verify collection can find function by both URI and URN
	assert.True(t, collection.URILoaded(testURI))
	assert.True(t, collection.URNLoaded("urn:example:test"))
}

func TestMultipleFunctionsFromSameExtension(t *testing.T) {
	// Test that multiple functions from the same extension are handled correctly
	var collection extensions.Collection

	// Extension with multiple functions
	testURI := "https://example.com/multi.yaml"
	testYAML := `---
urn: "urn:example:multi"
scalar_functions:
  - name: "add_multi"
    impls:
      - args:
          - name: x
            value: i32
          - name: y
            value: i32
        return: i32
  - name: "subtract_multi"
    impls:
      - args:
          - name: x
            value: i32
          - name: y
            value: i32
        return: i32
  - name: "multiply_multi"
    impls:
      - args:
          - name: x
            value: i32
          - name: y
            value: i32
        return: i32
`

	err := collection.Load(testURI, strings.NewReader(testYAML))
	require.NoError(t, err)

	builder := NewBuilder(&collection)

	baseSchema := types.NamedStruct{
		Names:  []string{"a", "b"},
		Struct: types.StructType{Types: []types.Type{&types.Int32Type{}, &types.Int32Type{}}},
	}
	scan := builder.NamedScan([]string{"test_table"}, baseSchema)

	exprBuilder := builder.GetExprBuilder()
	exprBuilder.BaseSchema = types.NewRecordTypeFromStruct(baseSchema.Struct)

	// Create expressions for all three functions
	addExpr, err := exprBuilder.ScalarFunc(extensions.ID{
		URI:  testURI, // via URI
		Name: "add_multi",
	}).Args(
		exprBuilder.RootRef(expr.NewStructFieldRef(0)),
		exprBuilder.RootRef(expr.NewStructFieldRef(1)),
	).BuildExpr()
	require.NoError(t, err)

	subtractExpr, err := exprBuilder.ScalarFunc(extensions.ID{
		URN:  "urn:example:multi", // via URN
		Name: "subtract_multi",
	}).Args(
		exprBuilder.RootRef(expr.NewStructFieldRef(0)),
		exprBuilder.RootRef(expr.NewStructFieldRef(1)),
	).BuildExpr()
	require.NoError(t, err)

	multiplyExpr, err := exprBuilder.ScalarFunc(extensions.ID{
		URI:  testURI, // via URI
		Name: "multiply_multi",
	}).Args(
		exprBuilder.RootRef(expr.NewStructFieldRef(0)),
		exprBuilder.RootRef(expr.NewStructFieldRef(1)),
	).BuildExpr()
	require.NoError(t, err)

	projected, err := builder.Project(scan, addExpr, subtractExpr, multiplyExpr)
	require.NoError(t, err)

	planObj, err := builder.Plan(projected, []string{"a", "b", "add_result", "subtract_result", "multiply_result"})
	require.NoError(t, err)

	proto, err := planObj.ToProto()
	require.NoError(t, err)

	assert.Len(t, proto.ExtensionUris, 1, "Should have one URI reference")
	assert.Len(t, proto.ExtensionUrns, 1, "Should have one URN reference")

	functionCount := 0
	for _, ext := range proto.Extensions {
		if ext.GetExtensionFunction() != nil {
			functionCount++
		}
	}
	assert.Equal(t, 3, functionCount, "Should have exactly 3 functions")

	assert.True(t, collection.URILoaded(testURI))
	assert.True(t, collection.URNLoaded("urn:example:multi"))
}

// TestPlanFromProtoWithMissingURI tests parsing a plan that references both URI and URN
// but only has the URI available in the collection (URN is missing)
func TestPlanFromProtoWithMissingURN(t *testing.T) {
	var collection extensions.Collection

	// Load extension with URI only (no URN in YAML)
	const uri = "http://localhost/test.yaml"
	const extensionYAML = `---
scalar_functions:
  - name: "add"
    impls:
      - args:
          - name: x
            value: i32
          - name: y
            value: i32
        return: i32
`

	require.NoError(t, collection.Load(uri, strings.NewReader(extensionYAML)))

	// Create a plan protobuf that references both URI and URN but the collection only has URI
	planProto := &proto.Plan{
		Version: &proto.Version{MinorNumber: 29, Producer: "test"},
		ExtensionUris: []*protoext.SimpleExtensionURI{
			{ExtensionUriAnchor: 1, Uri: uri},
		},
		ExtensionUrns: []*protoext.SimpleExtensionURN{
			{ExtensionUrnAnchor: 1, Urn: "urn:example:missing"},
		},
		Extensions: []*protoext.SimpleExtensionDeclaration{
			{
				MappingType: &protoext.SimpleExtensionDeclaration_ExtensionFunction_{
					ExtensionFunction: &protoext.SimpleExtensionDeclaration_ExtensionFunction{
						ExtensionUriReference: 1,
						ExtensionUrnReference: 1,
						FunctionAnchor:        1,
						Name:                  "add:i32_i32",
					},
				},
			},
		},
		Relations: []*proto.PlanRel{
			{
				RelType: &proto.PlanRel_Root{
					Root: &proto.RelRoot{
						Input: &proto.Rel{
							RelType: &proto.Rel_Read{
								Read: &proto.ReadRel{
									Common: &proto.RelCommon{EmitKind: &proto.RelCommon_Direct_{Direct: &proto.RelCommon_Direct{}}},
									BaseSchema: &proto.NamedStruct{
										Names: []string{"a", "b"},
										Struct: &proto.Type_Struct{
											Types: []*proto.Type{
												{Kind: &proto.Type_I32_{I32: &proto.Type_I32{}}},
												{Kind: &proto.Type_I32_{I32: &proto.Type_I32{}}},
											},
										},
									},
									ReadType: &proto.ReadRel_NamedTable_{
										NamedTable: &proto.ReadRel_NamedTable{Names: []string{"test_table"}},
									},
								},
							},
						},
						Names: []string{"a", "b"},
					},
				},
			},
		},
	}

	parsedPlan, err := FromProto(planProto, &collection)
	require.NoError(t, err)
	require.NotNil(t, parsedPlan)

	reconstructedProto, err := parsedPlan.ToProto()
	require.NoError(t, err)

	// Expected proto should preserve both URI and URN from input (structure-preserving)
	expectedProto := &proto.Plan{
		Version: &proto.Version{MinorNumber: 29, Producer: "test"},
		ExtensionUris: []*protoext.SimpleExtensionURI{
			{ExtensionUriAnchor: 1, Uri: uri},
		},
		ExtensionUrns: []*protoext.SimpleExtensionURN{
			{ExtensionUrnAnchor: 1, Urn: "urn:example:missing"},
		},
		Extensions: []*protoext.SimpleExtensionDeclaration{
			{
				MappingType: &protoext.SimpleExtensionDeclaration_ExtensionFunction_{
					ExtensionFunction: &protoext.SimpleExtensionDeclaration_ExtensionFunction{
						ExtensionUriReference: 1,
						ExtensionUrnReference: 1,
						FunctionAnchor:        1,
						Name:                  "add:i32_i32",
					},
				},
			},
		},
		Relations: []*proto.PlanRel{
			{
				RelType: &proto.PlanRel_Root{
					Root: &proto.RelRoot{
						Input: &proto.Rel{
							RelType: &proto.Rel_Read{
								Read: &proto.ReadRel{
									Common: &proto.RelCommon{EmitKind: &proto.RelCommon_Direct_{Direct: &proto.RelCommon_Direct{}}},
									BaseSchema: &proto.NamedStruct{
										Names: []string{"a", "b"},
										Struct: &proto.Type_Struct{
											Types: []*proto.Type{
												{Kind: &proto.Type_I32_{I32: &proto.Type_I32{}}},
												{Kind: &proto.Type_I32_{I32: &proto.Type_I32{}}},
											},
										},
									},
									ReadType: &proto.ReadRel_NamedTable_{
										NamedTable: &proto.ReadRel_NamedTable{Names: []string{"test_table"}},
									},
								},
							},
						},
						Names: []string{"a", "b"},
					},
				},
			},
		},
	}

	// Compare the round-trip result with expected proto
	if diff := cmp.Diff(expectedProto, reconstructedProto, protocmp.Transform()); diff != "" {
		t.Errorf("Round-trip proto mismatch (-want +got):\n%s", diff)
	}

	assert.True(t, collection.URILoaded(uri))
	assert.False(t, collection.URNLoaded("urn:example:missing"))
}

// TestPlanFromProtoWithMissingURN tests parsing a plan that references both URI and URN
// but only has the URN available in the collection (URI is missing)
func TestPlanFromProtoWithMissingURI(t *testing.T) {
	var collection extensions.Collection

	// Load extension with URN only
	const extensionYAML = `---
urn: "urn:example:test"
scalar_functions:
  - name: "add"
    impls:
      - args:
          - name: x
            value: i32
          - name: y
            value: i32
        return: i32
`

	require.NoError(t, collection.LoadWithoutUri(strings.NewReader(extensionYAML)))

	// Create a plan protobuf that references both URI and URN but the collection only has URN
	planProto := &proto.Plan{
		Version: &proto.Version{MinorNumber: 29, Producer: "test"},
		ExtensionUris: []*protoext.SimpleExtensionURI{
			{ExtensionUriAnchor: 1, Uri: "http://localhost/missing.yaml"},
		},
		ExtensionUrns: []*protoext.SimpleExtensionURN{
			{ExtensionUrnAnchor: 1, Urn: "urn:example:test"},
		},
		Extensions: []*protoext.SimpleExtensionDeclaration{
			{
				MappingType: &protoext.SimpleExtensionDeclaration_ExtensionFunction_{
					ExtensionFunction: &protoext.SimpleExtensionDeclaration_ExtensionFunction{
						ExtensionUriReference: 1, // This URI doesn't exist in collection
						ExtensionUrnReference: 1,
						FunctionAnchor:        1,
						Name:                  "add:i32_i32",
					},
				},
			},
		},
		Relations: []*proto.PlanRel{
			{
				RelType: &proto.PlanRel_Root{
					Root: &proto.RelRoot{
						Input: &proto.Rel{
							RelType: &proto.Rel_Read{
								Read: &proto.ReadRel{
									Common: &proto.RelCommon{EmitKind: &proto.RelCommon_Direct_{Direct: &proto.RelCommon_Direct{}}},
									BaseSchema: &proto.NamedStruct{
										Names: []string{"a", "b"},
										Struct: &proto.Type_Struct{
											Types: []*proto.Type{
												{Kind: &proto.Type_I32_{I32: &proto.Type_I32{}}},
												{Kind: &proto.Type_I32_{I32: &proto.Type_I32{}}},
											},
										},
									},
									ReadType: &proto.ReadRel_NamedTable_{
										NamedTable: &proto.ReadRel_NamedTable{Names: []string{"test_table"}},
									},
								},
							},
						},
						Names: []string{"a", "b"},
					},
				},
			},
		},
	}

	parsedPlan, err := FromProto(planProto, &collection)
	require.NoError(t, err)
	require.NotNil(t, parsedPlan)

	reconstructedProto, err := parsedPlan.ToProto()
	require.NoError(t, err)

	// Expected proto should preserve both URI and URN from input (structure-preserving)
	expectedProto := &proto.Plan{
		Version: &proto.Version{MinorNumber: 29, Producer: "test"},
		ExtensionUris: []*protoext.SimpleExtensionURI{
			{ExtensionUriAnchor: 1, Uri: "http://localhost/missing.yaml"},
		},
		ExtensionUrns: []*protoext.SimpleExtensionURN{
			{ExtensionUrnAnchor: 1, Urn: "urn:example:test"},
		},
		Extensions: []*protoext.SimpleExtensionDeclaration{
			{
				MappingType: &protoext.SimpleExtensionDeclaration_ExtensionFunction_{
					ExtensionFunction: &protoext.SimpleExtensionDeclaration_ExtensionFunction{
						ExtensionUriReference: 1,
						ExtensionUrnReference: 1,
						FunctionAnchor:        1,
						Name:                  "add:i32_i32",
					},
				},
			},
		},
		Relations: []*proto.PlanRel{
			{
				RelType: &proto.PlanRel_Root{
					Root: &proto.RelRoot{
						Input: &proto.Rel{
							RelType: &proto.Rel_Read{
								Read: &proto.ReadRel{
									Common: &proto.RelCommon{EmitKind: &proto.RelCommon_Direct_{Direct: &proto.RelCommon_Direct{}}},
									BaseSchema: &proto.NamedStruct{
										Names: []string{"a", "b"},
										Struct: &proto.Type_Struct{
											Types: []*proto.Type{
												{Kind: &proto.Type_I32_{I32: &proto.Type_I32{}}},
												{Kind: &proto.Type_I32_{I32: &proto.Type_I32{}}},
											},
										},
									},
									ReadType: &proto.ReadRel_NamedTable_{
										NamedTable: &proto.ReadRel_NamedTable{Names: []string{"test_table"}},
									},
								},
							},
						},
						Names: []string{"a", "b"},
					},
				},
			},
		},
	}

	// Compare the round-trip result with expected (now fixed) behavior
	if diff := cmp.Diff(expectedProto, reconstructedProto, protocmp.Transform()); diff != "" {
		t.Errorf("Round-trip proto mismatch (-want +got):\n%s", diff)
	}

	// Verify the collection states
	assert.False(t, collection.URILoaded("http://localhost/missing.yaml"))
	assert.True(t, collection.URNLoaded("urn:example:test"))
}

func TestRelFromProto(t *testing.T) {

	registry := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
	literal5 := &proto.Expression_Literal{LiteralType: &proto.Expression_Literal_I64{I64: 5}}
	exprLiteral5 := &proto.Expression{RexType: &proto.Expression_Literal_{Literal: literal5}}

	nestedStructExpr1 := &proto.Expression_Nested_Struct{Fields: []*proto.Expression{exprLiteral5}}
	virtualTableWithExpression := &proto.ReadRel_VirtualTable_{VirtualTable: &proto.ReadRel_VirtualTable{Expressions: []*proto.Expression_Nested_Struct{nestedStructExpr1}}}
	readRelWithExpression := &proto.ReadRel{ReadType: virtualTableWithExpression}

	literalStruct := &proto.Expression_Literal_Struct{Fields: []*proto.Expression_Literal{literal5}}
	virtualTableWithLiteral := &proto.ReadRel_VirtualTable_{VirtualTable: &proto.ReadRel_VirtualTable{Values: []*proto.Expression_Literal_Struct{literalStruct}}}
	readRelWithLiteral := &proto.ReadRel{ReadType: virtualTableWithLiteral}

	for _, td := range []struct {
		name     string
		readType *proto.ReadRel
	}{
		{"virtual table with expression", readRelWithExpression},
		{"virtual table with deprecated literal", readRelWithLiteral},
	} {
		t.Run(td.name, func(t *testing.T) {
			rel := &proto.Rel{RelType: &proto.Rel_Read{Read: td.readType}}

			outRel, err := RelFromProto(rel, registry)
			require.NoError(t, err)
			gotRel := outRel.ToProto()
			gotReadRel, ok := gotRel.RelType.(*proto.Rel_Read)
			require.True(t, ok)
			gotVirtualTableReadRel, ok := gotReadRel.Read.ReadType.(*proto.ReadRel_VirtualTable_)
			require.True(t, ok)
			// in case of both deprecated or new expression, the output should be the same as the new expression
			if diff := cmp.Diff(gotVirtualTableReadRel, virtualTableWithExpression, protocmp.Transform()); diff != "" {
				t.Errorf("expression proto didn't match, diff:\n%v", diff)
			}
		})
	}

}
