// SPDX-License-Identifier: Apache-2.0

package codec_test

import (
	"embed"
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/codec"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/literal"
	"github.com/substrait-io/substrait-go/v9/plan"
	"github.com/substrait-io/substrait-go/v9/types"
	substraitproto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// Embed test JSON files for expected output comparison.
//
//go:embed testdata/plan/*.json
var testdataFS embed.FS

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

var baseSchemaReverse = types.NamedStruct{Names: []string{"x", "y"},
	Struct: types.StructType{
		Nullability: types.NullabilityRequired,
		Types: []types.Type{
			&types.Float32Type{Nullability: types.NullabilityRequired},
			&types.StringType{Nullability: types.NullabilityRequired},
		},
	}}

// schema structures for testing purposes.
var (
	employeeSchema = types.NamedStruct{Names: []string{"employee_id", "name", "department_id", "salary", "role"},
		Struct: types.StructType{
			Nullability: types.NullabilityRequired,
			Types: []types.Type{
				&types.Int32Type{Nullability: types.NullabilityRequired},
				&types.StringType{Nullability: types.NullabilityNullable},
				&types.Int32Type{Nullability: types.NullabilityNullable},
				&types.DecimalType{Precision: 10, Scale: 2, Nullability: types.NullabilityNullable},
				&types.StringType{Nullability: types.NullabilityNullable},
			},
		}}

	employeeSalariesSchema = types.NamedStruct{Names: []string{"name", "salary"},
		Struct: types.StructType{
			Types: []types.Type{
				&types.StringType{Nullability: types.NullabilityNullable},
				&types.DecimalType{Precision: 10, Scale: 2, Nullability: types.NullabilityNullable},
			},
		}}

	employeeSchemaNullable = types.NamedStruct{Names: []string{"employee_id", "name", "department_id", "salary", "role"},
		Struct: types.StructType{
			Types: []types.Type{
				&types.Int32Type{Nullability: types.NullabilityNullable},
				&types.StringType{Nullability: types.NullabilityNullable},
				&types.Int32Type{Nullability: types.NullabilityNullable},
				&types.DecimalType{Precision: 10, Scale: 2, Nullability: types.NullabilityNullable},
				&types.StringType{Nullability: types.NullabilityNullable},
			},
		}}
)

var (
	v1 = expr.PrimitiveLiteral[int32]{Value: 1, Type: &types.Int32Type{Nullability: types.NullabilityRequired}}
	v2 = expr.PrimitiveLiteral[int32]{Value: 2, Type: &types.Int32Type{Nullability: types.NullabilityRequired}}
)

func checkRoundTrip(t *testing.T, expectedJSON string, p *plan.Plan) {
	t.Helper()
	protoPlan, err := codec.PlanToProto(p)
	require.NoError(t, err)

	var expectedProto substraitproto.Plan
	require.NoError(t, protojson.Unmarshal([]byte(expectedJSON), &expectedProto))

	// Equalize producer field; it may differ between golden JSON and protoPlan
	// depending on which OS (GOOS, ARCH, and the like) this test runs.
	protoPlan.Version.Producer = expectedProto.Version.Producer

	assert.Truef(t, proto.Equal(&expectedProto, protoPlan), "JSON expected: %s\ngot: %s",
		protojson.Format(&expectedProto), protojson.Format(protoPlan))

	roundTrip, err := codec.PlanFromProto(&expectedProto, extensions.GetDefaultCollectionWithNoError())
	require.NoError(t, err)

	roundTripProto, err := codec.PlanToProto(roundTrip)
	require.NoError(t, err)

	assert.Truef(t, proto.Equal(protoPlan, roundTripProto), "plan expected: %s\ngot: %s",
		protojson.Format(protoPlan), protojson.Format(roundTripProto))
}

// makeProjectionMaskExpr generates a MaskExpression to project or reorder columns by the given IDs.
func makeProjectionMaskExpr(columnIds []int) *expr.MaskExpression {
	structItems := make([]*substraitproto.Expression_MaskExpression_StructItem, len(columnIds))

	for index, columnId := range columnIds {
		structItems[index] = &substraitproto.Expression_MaskExpression_StructItem{
			Field: int32(columnId),
		}
	}

	return expr.MaskExpressionFromProto(
		&substraitproto.Expression_MaskExpression{
			Select: &substraitproto.Expression_MaskExpression_StructSelect{
				StructItems: structItems,
			},
			MaintainSingularStruct: true,
		},
	)
}

// makeNamedTableReadRel creates a named table read relation with the selected column IDs.
func makeNamedTableReadRel(b plan.Builder, tableNames []string, tableSchema types.NamedStruct, columnIds []int) plan.Rel {
	namedTableReadRel := b.NamedScan(tableNames, tableSchema)
	namedTableReadRel.SetProjection(makeProjectionMaskExpr(columnIds))
	return namedTableReadRel
}

// makeConditionExprForLike constructs a LIKE condition expression for the specified column and value.
func makeConditionExprForLike(t *testing.T, b plan.Builder, scan plan.Rel, colId int, valueLiteral expr.Literal) expr.Expression {
	id := extensions.FunctionID{
		URN:  "extension:io.substrait:functions_string",
		Name: "contains:str_str",
	}
	b.GetFunctionRef(id.URN, id.Name)
	colIdRef, err := b.RootFieldRef(scan, int32(colId))
	require.NoError(t, err)
	scalarExpr, err := b.ScalarFn(id.URN, id.Name, nil, colIdRef, valueLiteral)
	require.NoError(t, err)
	return scalarExpr
}

func makeFilterRel(t *testing.T, b plan.Builder, input plan.Rel, condition expr.Expression) plan.Rel {
	filterRel, err := b.Filter(input, condition)
	require.NoError(t, err)
	return filterRel
}

func makeProjectRel(t *testing.T, b plan.Builder, input plan.Rel, columnIds []int) plan.Rel {
	refs := make([]expr.Expression, len(columnIds))
	for i, c := range columnIds {
		ref, err := b.RootFieldRef(input, int32(c))
		require.NoError(t, err)
		refs[i] = ref
	}
	project, err := b.Project(input, refs...)
	require.NoError(t, err)
	return project
}

// getProjectionForTest1 returns project rel for "Select name, salary from employees"
func getProjectionForTest1(t *testing.T, b plan.Builder) plan.Rel {
	namedScanRel := makeNamedTableReadRel(b, []string{"employees"}, employeeSchema, []int{1, 3})
	return makeProjectRel(t, b, namedScanRel, []int{0, 1})
}

// getProjectionForTest2 returns project rel for "Select * from employees where role LIKE 'Engineer'"
func getProjectionForTest2(t *testing.T, b plan.Builder) plan.Rel {
	// scanRel outputs role, employee_id, name, department_id, salary
	namedScanRel := makeNamedTableReadRel(b, []string{"employees"}, employeeSchema, []int{4, 0, 1, 2, 3})

	// column 0 from the output of namedScanRel is role
	// Build the filter with condition `role LIKE 'Engineer'`
	l := literal.NewString("Engineer", false)
	roleLikeEngineer := makeConditionExprForLike(t, b, namedScanRel, 1, l)
	filterRel := makeFilterRel(t, b, namedScanRel, roleLikeEngineer)

	// projectRel output employee_id, name, department_id, salary, role
	return makeProjectRel(t, b, filterRel, []int{1, 2, 3, 4, 0})
}

// getFilterForTest1 returns filter rel for "name LIKE 'Alice'"
func getFilterForTest1(t *testing.T, b plan.Builder) plan.Rel {
	namedTableReadRel := b.NamedScan([]string{"employee_salaries"}, employeeSalariesSchema)

	// column 0 from the output of namedTableReadRel is name
	// Build the filter with condition `name LIKE 'Alice'`
	l := literal.NewString("Alice", false)
	nameLikeAlice := makeConditionExprForLike(t, b, namedTableReadRel, 0, l)
	return makeFilterRel(t, b, namedTableReadRel, nameLikeAlice)
}

// makeAddExpr constructs expression val1 + val2.
func makeAddExpr(t *testing.T, b plan.Builder, val1, val2 expr.Literal) expr.Expression {
	id := extensions.FunctionID{
		URN:  "extension:io.substrait:functions_arithmetic",
		Name: "add:i32_i32",
	}
	b.GetFunctionRef(id.URN, id.Name)
	scalarExpr, err := b.ScalarFn(id.URN, id.Name, nil, val1, val2)
	require.NoError(t, err)
	return scalarExpr
}

func buildLiteralExpressions(_ *testing.T, _ plan.Builder) []expr.VirtualTableExpressionValue {
	return []expr.VirtualTableExpressionValue{{&v1, &v2}}
}

// buildScalarAddExpression builds a scalar binary add expression
func buildScalarAddExpression(t *testing.T, b plan.Builder) []expr.VirtualTableExpressionValue {
	s1 := makeAddExpr(t, b, &v1, &v1)
	s2 := makeAddExpr(t, b, &v2, &v2)
	return []expr.VirtualTableExpressionValue{{s1, s2}}
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

func TestBasicEmitPlan(t *testing.T) {
	b := plan.NewBuilderDefault()
	root, err := b.NamedScan([]string{"test"}, baseSchema).Remap(1, 0)
	require.NoError(t, err)
	p, err := b.Plan(root, []string{"a", "b"})
	require.NoError(t, err)

	protoPlan, err := codec.PlanToProto(p)
	require.NoError(t, err)

	roundTrip, err := codec.PlanFromProto(protoPlan, extensions.GetDefaultCollectionWithNoError())
	require.NoError(t, err)

	assert.Equal(t, p, roundTrip)
	assert.Equal(t, "NSTRUCT<a: fp32, b: string>", p.GetRoots()[0].RecordType().String())
	assert.Equal(t, roundTrip.GetRoots()[0].RecordType(), p.GetRoots()[0].RecordType())
}

func TestEmitEmptyPlan(t *testing.T) {
	b := plan.NewBuilderDefault()
	root := b.NamedScan([]string{"test"}, baseSchema)
	newRoot, err := root.Remap()
	require.NoError(t, err)
	_, err = b.Plan(newRoot, []string{})
	require.NoError(t, err)

	b = plan.NewBuilderDefault()
	root = b.NamedScan([]string{"test"}, baseSchema)
	newRoot, err = root.Remap(1, 0)
	require.NoError(t, err)
	p, err := b.Plan(newRoot, []string{"a", "b"})
	require.NoError(t, err)

	assert.Equal(t, "NSTRUCT<a: fp32, b: string>", p.GetRoots()[0].RecordType().String())

	// Verify the mapping remains the same after receiving an error.
	_, err = root.Remap(-1)
	require.Error(t, err)
	assert.Equal(t, "NSTRUCT<a: fp32, b: string>", p.GetRoots()[0].RecordType().String())

	protoPlan, err := codec.PlanToProto(p)
	require.NoError(t, err)

	roundTrip, err := codec.PlanFromProto(protoPlan, extensions.GetDefaultCollectionWithNoError())
	require.NoError(t, err)

	assert.Equal(t, p, roundTrip)
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

func TestCrossRel(t *testing.T) {
	const expectedJSON = `{
		` + versionStruct + `,
		"relations": [
			{
				"root": {
					"input": {
						"cross": {
							"common": {
								"direct": {}
							},
							"left": {
								"read": {
									"common": {"direct": {}},
									"baseSchema": {
										"names": ["a", "b"],
										"struct": {
											"nullability": "NULLABILITY_REQUIRED",
											"types": [
												{ "string": { "nullability": "NULLABILITY_REQUIRED" }},
												{ "fp32": { "nullability": "NULLABILITY_REQUIRED" }}
											]
										}
									},
									"namedTable": {
										"names": [ "test" ]
									}
								}
							},
							"right": {
								"read": {
									"common": {"direct": {}},
									"baseSchema": {
										"names": ["x", "y"],
										"struct": {
											"nullability": "NULLABILITY_REQUIRED",
											"types": [
												{ "i32": { "nullability": "NULLABILITY_REQUIRED" }},
												{ "bool": { "nullability": "NULLABILITY_REQUIRED" }}
											]
										}
									},
									"namedTable": {
										"names": [ "test2" ]
									}
								}
							}
						}
					},
					"names": ["str", "fp", "i", "bool" ]
				}
			}
		]
	}`

	b := plan.NewBuilderDefault()
	left := b.NamedScan([]string{"test"}, baseSchema)
	right := b.NamedScan([]string{"test2"}, baseSchema2)

	root, err := b.Cross(left, right)
	require.NoError(t, err)

	p, err := b.Plan(root, []string{"str", "fp", "i", "bool"})
	require.NoError(t, err)

	assert.Equal(t, "NSTRUCT<str: string, fp: fp32, i: i32, bool: boolean>", p.GetRoots()[0].RecordType().String())

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
							"offset": 100,
							"count": -1
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

	fetch, err := b.Fetch(scan, 100, plan.FETCH_COUNT_ALL_RECORDS)
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

func TestJoinRelOutputRecordTypes(t *testing.T) {
	const initialJSONFmt = `{
		` + versionStruct + `,
		"relations": [
			{
				"root": {
					"input": {
						"join": {
							"common": {"direct": {}},
							"left": {
								"read": {
									"common": {"direct": {}},
									"baseSchema": {
										"names": ["a", "b"],
										"struct": {
											"nullability": "NULLABILITY_REQUIRED",
											"types": [
												{ "string": { "nullability": "NULLABILITY_REQUIRED" }},
												{ "fp32": { "nullability": "NULLABILITY_REQUIRED" }}
											]
										}
									},
									"namedTable": {
										"names": [ "test" ]
									}
								}
							},
							"right": {
								"read": {
									"common": {"direct": {}},
									"baseSchema": {
										"names": ["x", "y"],
										"struct": {
											"nullability": "NULLABILITY_REQUIRED",
											"types": [
												{ "i32": { "nullability": "NULLABILITY_REQUIRED" }},
												{ "bool": { "nullability": "NULLABILITY_REQUIRED" }}
											]
										}
									},
									"namedTable": {
										"names": [ "test2" ]
									}
								}
							},
							"expression": {
								"selection": {
									"rootReference": {},
									"directReference": { "structField": { "field": 3 }}
								}
							},
							"type": "%s"
						}
					},
					"names": %s
				}
			}
		]
	}`

	tests := []struct {
		joinString   string
		joinType     plan.JoinType
		fields       []string
		recordString string
	}{
		{"JOIN_TYPE_INNER", plan.JoinTypeInner, []string{"a", "b", "c", "d"}, "NSTRUCT<a: string, b: fp32, c: i32, d: boolean>"},
		{"JOIN_TYPE_LEFT_SEMI", plan.JoinTypeLeftSemi, []string{"a", "b"}, "NSTRUCT<a: string, b: fp32>"},
		{"JOIN_TYPE_OUTER", plan.JoinTypeOuter, []string{"a", "b", "c", "d"}, "NSTRUCT<a: string?, b: fp32?, c: i32?, d: boolean?>"},
		{"JOIN_TYPE_LEFT", plan.JoinTypeLeft, []string{"a", "b", "c", "d"}, "NSTRUCT<a: string, b: fp32, c: i32?, d: boolean?>"},
		{"JOIN_TYPE_RIGHT", plan.JoinTypeRight, []string{"a", "b", "c", "d"}, "NSTRUCT<a: string?, b: fp32?, c: i32, d: boolean>"},
		{"JOIN_TYPE_LEFT_ANTI", plan.JoinTypeLeftAnti, []string{"a", "b"}, "NSTRUCT<a: string, b: fp32>"},
		{"JOIN_TYPE_LEFT_SINGLE", plan.JoinTypeLeftSingle, []string{"a", "b", "c", "d"}, "NSTRUCT<a: string, b: fp32, c: i32?, d: boolean?>"},
	}

	for _, tt := range tests {
		t.Run(tt.joinString, func(t *testing.T) {
			b := plan.NewBuilderDefault()
			left := b.NamedScan([]string{"test"}, baseSchema)
			right := b.NamedScan([]string{"test2"}, baseSchema2)

			cond, err := b.JoinedRecordFieldRef(left, right, 3)
			require.NoError(t, err)

			join, err := b.Join(left, right, cond, tt.joinType)
			require.NoError(t, err)

			p, err := b.Plan(join, tt.fields)
			require.NoError(t, err)

			assert.Equal(t, tt.recordString, p.GetRoots()[0].RecordType().String())

			names, _ := json.Marshal(tt.fields)
			checkRoundTrip(t, fmt.Sprintf(initialJSONFmt, tt.joinString, string(names)), p)
		})
	}
}

func TestJoinAndFilterRelation(t *testing.T) {
	const expectedJSON = `{
		` + versionStruct + `,
		"relations": [
			{
				"root": {
					"input": {
						"join": {
							"common": {"direct": {}},
							"left": {
								"read": {
									"common": {"direct": {}},
									"baseSchema": {
										"names": ["a", "b"],
										"struct": {
											"nullability": "NULLABILITY_REQUIRED",
											"types": [
												{ "string": { "nullability": "NULLABILITY_REQUIRED" }},
												{ "fp32": { "nullability": "NULLABILITY_REQUIRED" }}
											]
										}
									},
									"namedTable": {
										"names": [ "test" ]
									}
								}
							},
							"right": {
								"read": {
									"common": {"direct": {}},
									"baseSchema": {
										"names": ["x", "y"],
										"struct": {
											"nullability": "NULLABILITY_REQUIRED",
											"types": [
												{ "i32": { "nullability": "NULLABILITY_REQUIRED" }},
												{ "bool": { "nullability": "NULLABILITY_REQUIRED" }}
											]
										}
									},
									"namedTable": {
										"names": [ "test2" ]
									}
								}
							},
							"expression": {
								"selection": {
									"rootReference": {},
									"directReference": { "structField": { "field": 3 }}
								}
							},
							"postJoinFilter": {
								"selection": {
									"rootReference": {},
									"directReference": { "structField": { "field": 3 }}
								}
							},
							"type": "JOIN_TYPE_INNER"
						}
					},
					"names": ["a", "b", "c", "d"]
				}
			}
		]
	}`

	b := plan.NewBuilderDefault()
	left := b.NamedScan([]string{"test"}, baseSchema)
	right := b.NamedScan([]string{"test2"}, baseSchema2)

	cond, err := b.JoinedRecordFieldRef(left, right, 3)
	require.NoError(t, err)

	join, err := b.JoinAndFilter(left, right, cond, cond, plan.JoinTypeInner)
	require.NoError(t, err)

	p, err := b.Plan(join, []string{"a", "b", "c", "d"})
	require.NoError(t, err)

	checkRoundTrip(t, expectedJSON, p)
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

func TestExtensionSingleBuilder(t *testing.T) {
	const expectedJSON = `{
		` + versionStruct + `,
		"relations": [
			{
				"root": {
					"input": {
						"extensionSingle": {
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
							"detail": {
								"@type": "type.googleapis.com/google.protobuf.StringValue",
								"value": "test-config"
							}
						}
					},
					"names": ["result"]
				}
			}
		]
	}`

	b := plan.NewBuilderDefault()
	scan := b.NamedScan([]string{"test"}, baseSchema)

	// Create custom schema for extension
	customSchema := types.StructType{
		Nullability: types.NullabilityRequired,
		Types: []types.Type{
			&types.StringType{Nullability: types.NullabilityRequired},
		},
	}

	// Create extension definition
	extensionDef := &TestExtensionDefinition{
		schema: *types.NewRecordTypeFromStruct(customSchema),
		detail: []byte("test-config"),
		exprs:  nil,
	}

	extRel, err := b.ExtensionSingle(scan, extensionDef)
	require.NoError(t, err)

	p, err := b.Plan(extRel, []string{"result"})
	require.NoError(t, err)

	assert.Equal(t, "NSTRUCT<result: string>", p.GetRoots()[0].RecordType().String())

	checkRoundTrip(t, expectedJSON, p)
}

func TestExtensionLeafBuilder(t *testing.T) {
	const expectedJSON = `{
		` + versionStruct + `,
		"relations": [
			{
				"root": {
					"input": {
						"extensionLeaf": {
							"common": {"direct": {}},
							"detail": {
								"@type": "type.googleapis.com/google.protobuf.StringValue",
								"value": "leaf-config"
							}
						}
					},
					"names": ["x", "y"]
				}
			}
		]
	}`

	b := plan.NewBuilderDefault()

	// Create custom schema for leaf extension
	customSchema := types.StructType{
		Nullability: types.NullabilityRequired,
		Types: []types.Type{
			&types.Int32Type{Nullability: types.NullabilityRequired},
			&types.BooleanType{Nullability: types.NullabilityRequired},
		},
	}

	// Create extension definition
	extensionDef := &TestExtensionDefinition{
		schema: *types.NewRecordTypeFromStruct(customSchema),
		detail: []byte("leaf-config"),
		exprs:  nil,
	}

	extRel, err := b.ExtensionLeaf(extensionDef)
	require.NoError(t, err)

	p, err := b.Plan(extRel, []string{"x", "y"})
	require.NoError(t, err)

	assert.Equal(t, "NSTRUCT<x: i32, y: boolean>", p.GetRoots()[0].RecordType().String())

	checkRoundTrip(t, expectedJSON, p)
}

func TestExtensionMultiBuilder(t *testing.T) {
	const expectedJSON = `{
		` + versionStruct + `,
		"relations": [
			{
				"root": {
					"input": {
						"extensionMulti": {
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
											"names": ["x", "y"],
											"struct": {
												"types": [
													{"i32": { "nullability": "NULLABILITY_REQUIRED"}},
													{"bool": { "nullability": "NULLABILITY_REQUIRED"}}
												],
												"nullability": "NULLABILITY_REQUIRED"
											}
										},
										"namedTable": { "names": [ "test2" ]}
									}
								}
							],
							"detail": {
								"@type": "type.googleapis.com/google.protobuf.StringValue",
								"value": "multi-config"
							}
						}
					},
					"names": ["result"]
				}
			}
		]
	}`

	b := plan.NewBuilderDefault()
	left := b.NamedScan([]string{"test"}, baseSchema)
	right := b.NamedScan([]string{"test2"}, baseSchema2)

	// Create custom schema for multi extension
	customSchema := types.StructType{
		Nullability: types.NullabilityRequired,
		Types: []types.Type{
			&types.StringType{Nullability: types.NullabilityRequired},
		},
	}

	// Create extension definition
	extensionDef := &TestExtensionDefinition{
		schema: *types.NewRecordTypeFromStruct(customSchema),
		detail: []byte("multi-config"),
		exprs:  nil,
	}

	extRel, err := b.ExtensionMulti([]plan.Rel{left, right}, extensionDef)
	require.NoError(t, err)

	p, err := b.Plan(extRel, []string{"result"})
	require.NoError(t, err)

	assert.Equal(t, "NSTRUCT<result: string>", p.GetRoots()[0].RecordType().String())

	checkRoundTrip(t, expectedJSON, p)
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
