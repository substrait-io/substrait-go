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

func TestAggregateRelPlan(t *testing.T) {
	const expectedJSON = `{
		` + versionStruct + `,
		"extensionUrns": [
			{
				"extensionUrnAnchor": 1,
				"urn": "extension:io.substrait:functions_aggregate_generic"
			}
		],
		"extensions": [
			{
				"extensionFunction": {
					"extensionUrnReference": 1,
					"functionAnchor": 1,
					"name": "count:"
				}
			}
		],
		"relations": [
			{
				"root": {
					"input": {
						"aggregate": {
							"common": {"direct": {}},
							"input": {
								"read": {
									"common": {"direct": {}},
									"baseSchema": {
										"names": ["a", "b"],
										"struct": {
											"types": [
												{"string": { "nullability": "NULLABILITY_REQUIRED"}},
												{"fp32": { "nullability": "NULLABILITY_REQUIRED"}}
											],
											"nullability": "NULLABILITY_REQUIRED"
										}
									},
									"namedTable": { "names": [ "test" ]}
								}
							},
							"groupingExpressions": [
								{
									"selection": {
										"rootReference": {},
										"directReference": { "structField": { "field": 0 }}
									}
								}
							],
							"groupings": [
								{
									"expressionReferences": [
										0
									]
								}
							],
							"measures": [
								{
									"measure": {
										"functionReference": 1,
										"outputType": {
											"i64": {
												"nullability": "NULLABILITY_REQUIRED"
											}
										},
										"phase": "AGGREGATION_PHASE_INITIAL_TO_RESULT",
										"invocation": "AGGREGATION_INVOCATION_ALL"
									}
								}
							]
						}
					},
					"names": ["val", "cnt"]
				}
			}
		]
	}`

	b := plan.NewBuilderDefault()
	aggCount, err := b.AggregateFn(extensions.SubstraitDefaultURNPrefix+"functions_aggregate_generic",
		"count", nil)
	require.NoError(t, err)
	scan := b.NamedScan([]string{"test"}, baseSchema)
	root, err := b.AggregateColumns(scan, []plan.AggRelMeasure{b.Measure(aggCount, nil)}, 0)
	require.NoError(t, err)

	p, err := b.Plan(root, []string{"val", "cnt"})
	require.NoError(t, err)
	assert.Equal(t, "NSTRUCT<val: string, cnt: i64>", p.GetRoots()[0].RecordType().String())

	checkRoundTrip(t, expectedJSON, p)

	// Test with grouping expressions and references
	ref, err := b.RootFieldRef(scan, 0)
	require.NoError(t, err)
	exprs := make([]expr.Expression, 0)
	exprs = append(exprs, ref)
	root, err = b.AggregateExprs(scan, []plan.AggRelMeasure{b.Measure(aggCount, nil)}, [][]expr.Expression{exprs}...)
	require.NoError(t, err)

	p, err = b.Plan(root, []string{"val", "cnt"})
	require.NoError(t, err)
	assert.Equal(t, "NSTRUCT<val: string, cnt: i64>", p.GetRoots()[0].RecordType().String())

	checkRoundTrip(t, expectedJSON, p)
}

func TestFetchRel(t *testing.T) {
	const expectedJSON = `{
		` + versionStruct + `,
		"relations": [
			{
				"root": {
					"input": {
						"fetch": {
							"common": {"direct": {}},
							"input": {
								"read": {
									"common": {
										"direct": {}
									},
									"baseSchema": {
										"names": ["a"],
										"struct": {
											"nullability": "NULLABILITY_REQUIRED",
											"types": [
												{"string": { "nullability": "NULLABILITY_REQUIRED" }}
											]
										}
									},
									"namedTable": {
										"names": ["test"]
									}
								}
							},
							"offsetExpr": {"literal": {"i64": "100"}}
						}
					},
					"names": ["a"]
				}
			}
		]
	}`

	b := plan.NewBuilderDefault()
	scan := b.NamedScan([]string{"test"}, types.NamedStruct{
		Names: []string{"a"},
		Struct: types.StructType{
			Nullability: types.NullabilityRequired,
			Types: []types.Type{
				&types.StringType{Nullability: types.NullabilityRequired}},
		},
	})

	offsetExpr := expr.Expression(expr.NewPrimitiveLiteral(int64(100), false))
	fetch, err := b.Fetch(scan, offsetExpr, nil)
	require.NoError(t, err)

	p, err := b.Plan(fetch, []string{"a"})
	require.NoError(t, err)

	assert.Equal(t, "NSTRUCT<a: string>", p.GetRoots()[0].RecordType().String())

	checkRoundTrip(t, expectedJSON, p)

	_, err = fetch.Remap(0)
	assert.NoError(t, err)
}

func TestFilterRelation(t *testing.T) {
	const expectedJSON = `{
		` + versionStruct + `,
		"relations": [
			{
				"root": {
					"input": {
						"filter": {
							"common": {
								"direct": {}
							},
							"input": {
								"read": {
									"common": {"direct": {}},
									"baseSchema": {
										"names": ["x", "y"],
										"struct": {
											"types": [
												{"i32": { "nullability": "NULLABILITY_REQUIRED"}},
												{"bool": { "nullability": "NULLABILITY_REQUIRED"}}
											],
											"nullability": "NULLABILITY_REQUIRED"
										}
									},
									"namedTable": { "names": [ "test" ]}
								}
							},
							"condition": {
								"selection": {
									"rootReference": {},
									"directReference": { "structField": { "field": 1 }}
								}
							}
						}
					},
					"names": ["a", "b"]
				}
			}
		]
	}`

	b := plan.NewBuilderDefault()
	scan := b.NamedScan([]string{"test"}, baseSchema2)
	ref, err := b.RootFieldRef(scan, 1)
	require.NoError(t, err)

	filter, err := b.Filter(scan, ref)
	require.NoError(t, err)

	p, err := b.Plan(filter, []string{"a", "b"})
	require.NoError(t, err)

	assert.Equal(t, "NSTRUCT<a: i32, b: boolean>", p.GetRoots()[0].RecordType().String())

	checkRoundTrip(t, expectedJSON, p)

	_, err = filter.Remap(0)
	assert.NoError(t, err)
}

func TestSortRelationsCoalesce(t *testing.T) {
	const expectedJSON = `{
		` + versionStruct + `,
		"relations": [
			{
				"root": {
					"input": {
						"sort": {
							"common": {"direct": {}},
							"input": {
								"read": {
									"common": {"direct": {}},
									"baseSchema": {
										"names": ["a", "b"],
										"struct": {
											"types": [
												{"string": { "nullability": "NULLABILITY_REQUIRED"}},
												{"fp32": { "nullability": "NULLABILITY_REQUIRED"}}
											],
											"nullability": "NULLABILITY_REQUIRED"
										}
									},
									"namedTable": { "names": [ "test" ]}
								}
							},
							"sorts": [
								{
									"expr": {
										"selection": {
											"rootReference": {},
											"directReference": { "structField": { "field": 0 }}
										}
									},
									"direction": "SORT_DIRECTION_CLUSTERED"
								}
							]
						}
					},
					"names": ["a", "b"]
				}
			}
		]
	}`

	b := plan.NewBuilderDefault()
	scan := b.NamedScan([]string{"test"}, baseSchema)

	ref, err := b.RootFieldRef(scan, 0)
	require.NoError(t, err)

	sort, err := b.Sort(scan, expr.SortField{Expr: ref, Kind: types.SortClustered})
	require.NoError(t, err)

	p, err := b.Plan(sort, []string{"a", "b"})
	require.NoError(t, err)

	assert.Equal(t, "NSTRUCT<a: string, b: fp32>", p.GetRoots()[0].RecordType().String())

	checkRoundTrip(t, expectedJSON, p)
}

func TestSortRelationKeyEqual(t *testing.T) {
	const expectedJSON = `{
		` + versionStruct + `,
		"extensionUrns": [
			{
				"extensionUrnAnchor": 1,
				"urn": "extension:io.substrait:functions_comparison"
			}
		],
		"extensions": [
			{
				"extensionFunction": {
					"extensionUrnReference": 1,
					"functionAnchor": 1,
					"name": "equal"
				}
			}
		],
		"relations": [
			{
				"root": {
					"input": {
						"sort": {
							"common": {"direct": {}},
							"input": {
								"read": {
									"common": {"direct": {}},
									"baseSchema": {
										"names": ["a", "b"],
										"struct": {
											"types": [
												{"string": { "nullability": "NULLABILITY_REQUIRED"}},
												{"fp32": { "nullability": "NULLABILITY_REQUIRED"}}
											],
											"nullability": "NULLABILITY_REQUIRED"
										}
									},
									"namedTable": { "names": [ "test" ]}
								}
							},
							"sorts": [
								{
									"expr": {
										"selection": {
											"rootReference": {},
											"directReference": {"structField": {"field": 0}}
										}
									},
									"comparisonFunctionReference": 1
								}
							]
						}
					},
					"names": ["a", "b"]
				}
			}
		]
	}`

	b := plan.NewBuilderDefault()
	scan := b.NamedScan([]string{"test"}, baseSchema)

	ref, err := b.RootFieldRef(scan, 0)
	require.NoError(t, err)

	sort, err := b.Sort(scan, expr.SortField{Expr: ref, Kind: b.GetFunctionRef(extensions.SubstraitDefaultURNPrefix+"functions_comparison", "equal")})
	require.NoError(t, err)

	p, err := b.Plan(sort, []string{"a", "b"})
	require.NoError(t, err)

	checkRoundTrip(t, expectedJSON, p)
}

func TestSortRelationMultiple(t *testing.T) {
	const expectedJSON = `{
		` + versionStruct + `,
		"relations": [
			{
				"root": {
					"input": {
						"sort": {
							"common": {"direct": {}},
							"input": {
								"read": {
									"common": {"direct": {}},
									"baseSchema": {
										"names": ["a", "b"],
										"struct": {
											"types": [
												{"string": { "nullability": "NULLABILITY_REQUIRED"}},
												{"fp32": { "nullability": "NULLABILITY_REQUIRED"}}
											],
											"nullability": "NULLABILITY_REQUIRED"
										}
									},
									"namedTable": { "names": [ "test" ]}
								}
							},
							"sorts": [
								{
									"expr": {
										"selection": {
											"rootReference": {},
											"directReference": {"structField": {"field": 1}}
										}
									},
									"direction": "SORT_DIRECTION_ASC_NULLS_LAST"
								},
								{
									"expr": {
										"selection": {
											"rootReference": {},
											"directReference": {"structField": {"field": 0}}
										}
									},
									"direction": "SORT_DIRECTION_DESC_NULLS_FIRST"
								}
							]
						}
					},
					"names": ["a", "b"]
				}
			}
		]
	}`

	b := plan.NewBuilderDefault()
	scan := b.NamedScan([]string{"test"}, baseSchema)

	ref, err := b.RootFieldRef(scan, 0)
	require.NoError(t, err)

	ref1, err := b.RootFieldRef(scan, 1)
	require.NoError(t, err)

	sort, err := b.Sort(scan, expr.SortField{Expr: ref1, Kind: types.SortAscNullsLast}, expr.SortField{Expr: ref, Kind: types.SortDescNullsFirst})
	require.NoError(t, err)

	p, err := b.Plan(sort, []string{"a", "b"})
	require.NoError(t, err)

	checkRoundTrip(t, expectedJSON, p)
}

func TestProjectExpressions(t *testing.T) {
	const expectedJSON = `{
		` + versionStruct + `,
		"extensionUrns": [
			{
				"extensionUrnAnchor": 1,
				"urn": "extension:io.substrait:functions_arithmetic"
			}
			],
			"extensions": [
			{
				"extensionFunction": {
				"extensionUrnReference": 1,
				"functionAnchor": 1,
				"name": "abs:fp32"
				}
			},
			{
				"extensionFunction": {
				"extensionUrnReference": 1,
				"functionAnchor": 2,
				"name": "add:fp32_fp32"
				}
			}
			],
		"relations": [
			{
				"root": {
				"input": {
					"project": {
					"common": {
						"direct": {}
					},
					"input": {
						"read": {
						"common": {
							"direct": {}
						},
						"baseSchema": {
							"names": [
							"a",
							"b"
							],
							"struct": {
							"types": [
								{
								"string": {
									"nullability": "NULLABILITY_REQUIRED"
								}
								},
								{
								"fp32": {
									"nullability": "NULLABILITY_REQUIRED"
								}
								}
							],
							"nullability": "NULLABILITY_REQUIRED"
							}
						},
						"namedTable": {
							"names": [
							"test"
							]
						}
						}
					},
					"expressions": [
						{
						"scalarFunction": {
							"functionReference": 2,
							"arguments": [
							{
								"value": {
								"scalarFunction": {
									"functionReference": 1,
									"arguments": [
									{
										"value": {
										"selection": {
											"directReference": {
											"structField": {
												"field": 1
											}
											},
											"rootReference": {}
										}
										}
									}
									],
									"outputType": {
									"fp32": {
										"nullability": "NULLABILITY_REQUIRED"
									}
									}
								}
								}
							},
							{
								"value": {
								"selection": {
									"directReference": {
									"structField": {
										"field": 1
									}
									},
									"rootReference": {}
								}
								}
							}
							],
							"options":  [
							  {}
							],
							"outputType": {
							"fp32": {
								"nullability": "NULLABILITY_REQUIRED"
							}
							}
						}
						}
					]
					}
				},
				"names": [
					"a",
					"b",
					"c"
				]
				}
			}
			]
		}`

	arithmeticURN := extensions.SubstraitDefaultURNPrefix + "functions_arithmetic"
	b := plan.NewBuilderDefault()
	scan := b.NamedScan([]string{"test"}, baseSchema)
	ref, err := b.RootFieldRef(scan, 1)
	require.NoError(t, err)

	abs, err := b.ScalarFn(arithmeticURN, "abs", nil, ref)
	require.NoError(t, err)

	add, err := b.GetExprBuilder().ScalarFunc(
		extensions.FunctionID{URN: arithmeticURN, Name: "add"}, nil).Args(
		b.GetExprBuilder().Expression(abs),
		b.GetExprBuilder().Expression(ref)).Build()
	require.NoError(t, err)

	project, err := b.Project(scan, add)
	require.NoError(t, err)

	p, err := b.Plan(project, []string{"a", "b", "c"})
	require.NoError(t, err)

	assert.Equal(t, "NSTRUCT<a: string, b: fp32, c: fp32>", p.GetRoots()[0].RecordType().String())

	checkRoundTrip(t, expectedJSON, p)
}

func TestProjectRelation(t *testing.T) {
	const expectedJSON = `{
		` + versionStruct + `,
		"relations": [
			{
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
											"types": [
												{"string": { "nullability": "NULLABILITY_REQUIRED"}},
												{"fp32": { "nullability": "NULLABILITY_REQUIRED"}}
											],
											"nullability": "NULLABILITY_REQUIRED"
										}
									},
									"namedTable": { "names": [ "test" ]}
								}
							},
							"expressions": [
								{
									"selection": {
										"rootReference": {},
										"directReference": { "structField": { "field": 1 }}
									}
								}
							]
						}
					},
					"names": ["a", "b", "c"]
				}
			}
		]
	}`

	b := plan.NewBuilderDefault()
	scan := b.NamedScan([]string{"test"}, baseSchema)
	ref, err := b.RootFieldRef(scan, 1)
	require.NoError(t, err)

	project, err := b.Project(scan, ref)
	require.NoError(t, err)

	p, err := b.Plan(project, []string{"a", "b", "c"})
	require.NoError(t, err)

	assert.Equal(t, "NSTRUCT<a: string, b: fp32, c: fp32>", p.GetRoots()[0].RecordType().String())

	checkRoundTrip(t, expectedJSON, p)
}

func TestProjectMultipleRelation(t *testing.T) {
	const expectedJSON = `{
		` + versionStruct + `,
		"relations": [
			{
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
											"types": [
												{"string": { "nullability": "NULLABILITY_REQUIRED"}},
												{"fp32": { "nullability": "NULLABILITY_REQUIRED"}}
											],
											"nullability": "NULLABILITY_REQUIRED"
										}
									},
									"namedTable": { "names": [ "test" ]}
								}
							},
							"expressions": [
								{
									"selection": {
										"rootReference": {},
										"directReference": { "structField": { "field": 1 }}
									}
								},
								{
									"selection": {
										"rootReference": {},
										"directReference": { "structField": { "field": 0 }}
									}
								}
							]
						}
					},
					"names": ["a", "b", "c", "d"]
				}
			}
		]
	}`

	b := plan.NewBuilderDefault()
	scan := b.NamedScan([]string{"test"}, baseSchema)
	ref, err := b.RootFieldRef(scan, 1)
	require.NoError(t, err)

	ref0, err := b.RootFieldRef(scan, 0)
	require.NoError(t, err)

	project, err := b.Project(scan, ref, ref0)
	require.NoError(t, err)

	p, err := b.Plan(project, []string{"a", "b", "c", "d"})
	require.NoError(t, err)

	assert.Equal(t, "NSTRUCT<a: string, b: fp32, c: fp32, d: string>", p.GetRoots()[0].RecordType().String())

	checkRoundTrip(t, expectedJSON, p)
}

func TestSetRelations(t *testing.T) {
	const expectedJSON = `{
		` + versionStruct + `,
		"relations": [
			{
				"root": {
					"input": {
						"set": {
							"common": {"direct": {}},
							"inputs": [
								{
									"read": {
										"common": {"direct": {}},
										"baseSchema": {
											"names": ["a", "b"],
											"struct": {
												"types": [
													{"string": { "nullability": "NULLABILITY_REQUIRED"}},
													{"fp32": { "nullability": "NULLABILITY_REQUIRED"}}
												],
												"nullability": "NULLABILITY_REQUIRED"
											}
										},
										"namedTable": { "names": [ "test" ]}
									}
								},
								{
									"read": {
										"common": {"direct": {}},
										"baseSchema": {
											"names": ["c", "d"],
											"struct": {
												"types": [
													{"string": { "nullability": "NULLABILITY_REQUIRED"}},
													{"fp32": { "nullability": "NULLABILITY_REQUIRED"}}
												],
												"nullability": "NULLABILITY_REQUIRED"
											}
										},
										"virtualTable": {
											"expressions": [
												{
													"fields": [
														{"literal": { "string": "foo", "nullable": false }},
														{"literal": { "fp32": 1.5, "nullable": false }}
													]
												},
												{
													"fields": [
														{"literal": { "string": "bar", "nullable": false }},
														{"literal": { "fp32": 3.5, "nullable": false }}
													]
												}
											]
										}
									}
								},
								{
									"read": {
										"common": {"emit": {
											"outputMapping": [1, 0]
										}},
										"baseSchema": {
											"names": ["x", "y"],
											"struct": {
												"types": [
													{"fp32": { "nullability": "NULLABILITY_REQUIRED"}},
													{"string": { "nullability": "NULLABILITY_REQUIRED"}}
												],
												"nullability": "NULLABILITY_REQUIRED"
											}
										},
										"namedTable": { "names": [ "test2" ]}
									}
								}
							],
							"op": "SET_OP_UNION_ALL"
						}
					},
					"names": ["a", "b"]
				}
			}
		]
	}`

	b := plan.NewBuilderDefault()
	scan1 := b.NamedScan([]string{"test"}, baseSchema)
	scan2, err := b.NamedScan([]string{"test2"}, baseSchemaReverse).Remap(1, 0)
	require.NoError(t, err)

	virtual, err := b.VirtualTable([]string{"c", "d"},
		expr.StructLiteralValue{expr.NewPrimitiveLiteral("foo", false), expr.NewPrimitiveLiteral(float32(1.5), false)},
		expr.StructLiteralValue{expr.NewPrimitiveLiteral("bar", false), expr.NewPrimitiveLiteral(float32(3.5), false)})
	require.NoError(t, err)

	set, err := b.Set(plan.SetOpUnionAll, scan1, virtual, scan2)
	require.NoError(t, err)

	p, err := b.Plan(set, []string{"a", "b"})
	require.NoError(t, err)

	assert.Equal(t, "NSTRUCT<a: string, b: fp32>", p.GetRoots()[0].RecordType().String())

	checkRoundTrip(t, expectedJSON, p)
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
