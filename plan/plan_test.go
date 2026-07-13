package plan

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	substraitgo "github.com/substrait-io/substrait-go/v8"
	"github.com/substrait-io/substrait-go/v8/expr"
	"github.com/substrait-io/substrait-go/v8/extensions"
	"github.com/substrait-io/substrait-go/v8/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	extensionspb "github.com/substrait-io/substrait-protobuf/go/substraitpb/extensions"
	"google.golang.org/protobuf/encoding/protojson"
	protobuf "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
)

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

const sampleYAML = `---
urn: extension:test:sample
types:
  - name: point
    structure:
      latitude: i32
      longitude: i32
  - name: line
    structure:
      start: point
      end: point
scalar_functions:
  -
    name: "add"
    description: "Add two values."
    impls:
      - args:
          - name: x
            value: i8
          - name: y
            value: i8
        options:
          overflow:
            values: [ SILENT, SATURATE, ERROR ]
        return: i8`

func TestPlanRoundTripWithExtensions(t *testing.T) {
	c := &extensions.Collection{}
	err := c.Load(strings.NewReader(sampleYAML))
	require.NoError(t, err)

	original := &proto.Plan{
		ExtensionUrns: []*extensionspb.SimpleExtensionURN{
			{ExtensionUrnAnchor: 1, Urn: "extension:test:sample"},
		},
		Extensions: []*extensionspb.SimpleExtensionDeclaration{
			{
				MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction_{
					ExtensionFunction: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction{
						ExtensionUrnReference: 1,
						FunctionAnchor:        1,
						Name:                  "add:i8_i8",
					},
				},
			},
		},
		Relations: []*proto.PlanRel{},
	}

	plan, err := FromProto(original, c)
	require.NoError(t, err)

	roundTripped, err := plan.ToProto()
	require.NoError(t, err)

	assert.True(t, protobuf.Equal(original, roundTripped),
		"Plan should be equivalent after round-trip.\nOriginal:      %s\nRound-tripped: %s",
		protojson.Format(original), protojson.Format(roundTripped))
}

func TestRejectsMismatchedRootNames(t *testing.T) {
	b := NewBuilderDefault()
	scan := b.NamedScan([]string{"test"}, types.NamedStruct{
		Names: []string{"a", "b"},
		Struct: types.StructType{
			Nullability: types.NullabilityRequired,
			Types:       []types.Type{&types.Int64Type{}, &types.StringType{}},
		},
	})
	_, err := b.Plan(scan, []string{"only_one"})
	assert.ErrorIs(t, err, substraitgo.ErrInvalidRel)
	assert.ErrorContains(t, err, "1 output name(s) but the output schema requires 2")
}

func TestFromProtoWithSubqueries(t *testing.T) {
	c := extensions.GetDefaultCollectionWithNoError()

	// Build a named-scan relation proto to embed as subquery input.
	scanProto := &proto.Rel{
		RelType: &proto.Rel_Read{
			Read: &proto.ReadRel{
				Common: &proto.RelCommon{EmitKind: &proto.RelCommon_Direct_{Direct: &proto.RelCommon_Direct{}}},
				BaseSchema: &proto.NamedStruct{
					Names: []string{"col1"},
					Struct: &proto.Type_Struct{
						Nullability: proto.Type_NULLABILITY_REQUIRED,
						Types:       []*proto.Type{{Kind: &proto.Type_I32_{I32: &proto.Type_I32{Nullability: proto.Type_NULLABILITY_REQUIRED}}}},
					},
				},
				ReadType: &proto.ReadRel_NamedTable_{NamedTable: &proto.ReadRel_NamedTable{Names: []string{"t"}}},
			},
		},
	}

	makePlan := func(subquery *proto.Expression_Subquery) *proto.Plan {
		return &proto.Plan{
			Relations: []*proto.PlanRel{
				{
					RelType: &proto.PlanRel_Root{
						Root: &proto.RelRoot{
							Names: []string{"out"},
							Input: &proto.Rel{
								RelType: &proto.Rel_Filter{
									Filter: &proto.FilterRel{
										Common: &proto.RelCommon{EmitKind: &proto.RelCommon_Direct_{Direct: &proto.RelCommon_Direct{}}},
										Input:  scanProto,
										Condition: &proto.Expression{
											RexType: &proto.Expression_Subquery_{Subquery: subquery},
										},
									},
								},
							},
						},
					},
				},
			},
		}
	}

	t.Run("ScalarSubquery", func(t *testing.T) {
		p := makePlan(&proto.Expression_Subquery{
			SubqueryType: &proto.Expression_Subquery_Scalar_{
				Scalar: &proto.Expression_Subquery_Scalar{Input: scanProto},
			},
		})
		result, err := FromProto(p, c)
		require.NoError(t, err)
		require.NotNil(t, result)
	})

	t.Run("InPredicateSubquery", func(t *testing.T) {
		needle := expr.NewPrimitiveLiteral(int32(1), false).ToProto()
		p := makePlan(&proto.Expression_Subquery{
			SubqueryType: &proto.Expression_Subquery_InPredicate_{
				InPredicate: &proto.Expression_Subquery_InPredicate{
					Needles:  []*proto.Expression{needle},
					Haystack: scanProto,
				},
			},
		})
		result, err := FromProto(p, c)
		require.NoError(t, err)
		require.NotNil(t, result)
	})

	t.Run("SetPredicateSubquery", func(t *testing.T) {
		p := makePlan(&proto.Expression_Subquery{
			SubqueryType: &proto.Expression_Subquery_SetPredicate_{
				SetPredicate: &proto.Expression_Subquery_SetPredicate{
					PredicateOp: proto.Expression_Subquery_SetPredicate_PREDICATE_OP_EXISTS,
					Tuples:      scanProto,
				},
			},
		})
		result, err := FromProto(p, c)
		require.NoError(t, err)
		require.NotNil(t, result)
	})

	t.Run("SetComparisonSubquery", func(t *testing.T) {
		left := expr.NewPrimitiveLiteral(int32(1), false).ToProto()
		p := makePlan(&proto.Expression_Subquery{
			SubqueryType: &proto.Expression_Subquery_SetComparison_{
				SetComparison: &proto.Expression_Subquery_SetComparison{
					ReductionOp:  proto.Expression_Subquery_SetComparison_REDUCTION_OP_ANY,
					ComparisonOp: proto.Expression_Subquery_SetComparison_COMPARISON_OP_EQ,
					Left:         left,
					Right:        scanProto,
				},
			},
		})
		result, err := FromProto(p, c)
		require.NoError(t, err)
		require.NotNil(t, result)
	})
}

func TestFromProtoRightSemiJoinRootNames(t *testing.T) {
	// Regression: validateRootNamesFromProto must not panic on JoinRel
	// with RIGHT_SEMI (directOutputSchema panics for unsupported join types).
	planJSON := `{
		"version": { "majorNumber": 0, "minorNumber": 79 },
		"relations": [{
			"root": {
				"names": ["id"],
				"input": {
					"join": {
						"common": { "direct": {} },
						"type": "JOIN_TYPE_RIGHT_SEMI",
						"expression": {
							"literal": { "boolean": true }
						},
						"left": {
							"read": {
								"common": { "direct": {} },
								"baseSchema": {
									"names": ["id"],
									"struct": {
										"nullability": "NULLABILITY_REQUIRED",
										"types": [{"i64": {"nullability": "NULLABILITY_REQUIRED"}}]
									}
								},
								"namedTable": { "names": ["left_table"] }
							}
						},
						"right": {
							"read": {
								"common": { "direct": {} },
								"baseSchema": {
									"names": ["id"],
									"struct": {
										"nullability": "NULLABILITY_REQUIRED",
										"types": [{"i64": {"nullability": "NULLABILITY_REQUIRED"}}]
									}
								},
								"namedTable": { "names": ["right_table"] }
							}
						}
					}
				}
			}
		}]
	}`

	var p proto.Plan
	require.NoError(t, protojson.Unmarshal([]byte(planJSON), &p))

	c := extensions.GetDefaultCollectionWithNoError()
	_, err := FromProto(&p, c)
	require.NoError(t, err)
}
