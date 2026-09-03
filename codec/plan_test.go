// SPDX-License-Identifier: Apache-2.0

package codec_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/codec"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/plan"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	extensionspb "github.com/substrait-io/substrait-protobuf/go/substraitpb/extensions"
	"google.golang.org/protobuf/encoding/protojson"
	protobuf "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

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
	require.NoError(t, c.Load(strings.NewReader(sampleYAML)))

	original := &proto.Plan{
		Version: &proto.Version{MinorNumber: 29},
		ExtensionUrns: []*extensionspb.SimpleExtensionURN{
			{ExtensionUrnAnchor: 1, Urn: "extension:test:sample"},
		},
		Extensions: []*extensionspb.SimpleExtensionDeclaration{
			{MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction_{
				ExtensionFunction: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction{
					ExtensionUrnReference: 1,
					FunctionAnchor:        1,
					Name:                  "add:i8_i8",
				},
			}},
		},
		Relations: []*proto.PlanRel{},
	}

	p, err := codec.PlanFromProto(original, c)
	require.NoError(t, err)

	roundTripped, err := codec.PlanToProto(p)
	require.NoError(t, err)

	assert.True(t, protobuf.Equal(original, roundTripped),
		"Plan should be equivalent after round-trip.\nOriginal:      %s\nRound-tripped: %s",
		protojson.Format(original), protojson.Format(roundTripped))
}

func TestPlanRoundTripWithSubqueries(t *testing.T) {
	c := extensions.GetDefaultCollectionWithNoError()
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
	needle := expr.NewPrimitiveLiteral(int32(1), false).ToProto()

	tests := []struct {
		name     string
		subquery *proto.Expression_Subquery
	}{
		{
			name: "ScalarSubquery",
			subquery: &proto.Expression_Subquery{
				SubqueryType: &proto.Expression_Subquery_Scalar_{
					Scalar: &proto.Expression_Subquery_Scalar{Input: scanProto},
				},
			},
		},
		{
			name: "InPredicateSubquery",
			subquery: &proto.Expression_Subquery{
				SubqueryType: &proto.Expression_Subquery_InPredicate_{
					InPredicate: &proto.Expression_Subquery_InPredicate{
						Needles:  []*proto.Expression{needle},
						Haystack: scanProto,
					},
				},
			},
		},
		{
			name: "SetPredicateSubquery",
			subquery: &proto.Expression_Subquery{
				SubqueryType: &proto.Expression_Subquery_SetPredicate_{
					SetPredicate: &proto.Expression_Subquery_SetPredicate{
						PredicateOp: proto.Expression_Subquery_SetPredicate_PREDICATE_OP_EXISTS,
						Tuples:      scanProto,
					},
				},
			},
		},
		{
			name: "SetComparisonSubquery",
			subquery: &proto.Expression_Subquery{
				SubqueryType: &proto.Expression_Subquery_SetComparison_{
					SetComparison: &proto.Expression_Subquery_SetComparison{
						ReductionOp:  proto.Expression_Subquery_SetComparison_REDUCTION_OP_ANY,
						ComparisonOp: proto.Expression_Subquery_SetComparison_COMPARISON_OP_EQ,
						Left:         needle,
						Right:        scanProto,
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			original := &proto.Plan{
				Version: &proto.Version{MinorNumber: 29},
				Relations: []*proto.PlanRel{{
					RelType: &proto.PlanRel_Root{
						Root: &proto.RelRoot{
							Names: []string{"out"},
							Input: &proto.Rel{
								RelType: &proto.Rel_Filter{
									Filter: &proto.FilterRel{
										Common:    &proto.RelCommon{EmitKind: &proto.RelCommon_Direct_{Direct: &proto.RelCommon_Direct{}}},
										Input:     scanProto,
										Condition: &proto.Expression{RexType: &proto.Expression_Subquery_{Subquery: tc.subquery}},
									},
								},
							},
						},
					},
				}},
			}
			p, err := codec.PlanFromProto(original, c)
			require.NoError(t, err)
			roundTripped, err := codec.PlanToProto(p)
			require.NoError(t, err)
			if diff := protoDiff(original, roundTripped); diff != "" {
				t.Errorf("plan round-trip mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFromProtoWithSubqueries(t *testing.T) {
	c := extensions.GetDefaultCollectionWithNoError()
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
	needle := expr.NewPrimitiveLiteral(int32(1), false).ToProto()
	tests := []struct {
		name     string
		subquery *proto.Expression_Subquery
	}{
		{
			name: "ScalarSubquery",
			subquery: &proto.Expression_Subquery{
				SubqueryType: &proto.Expression_Subquery_Scalar_{
					Scalar: &proto.Expression_Subquery_Scalar{Input: scanProto},
				},
			},
		},
		{
			name: "InPredicateSubquery",
			subquery: &proto.Expression_Subquery{
				SubqueryType: &proto.Expression_Subquery_InPredicate_{
					InPredicate: &proto.Expression_Subquery_InPredicate{
						Needles:  []*proto.Expression{needle},
						Haystack: scanProto,
					},
				},
			},
		},
		{
			name: "SetPredicateSubquery",
			subquery: &proto.Expression_Subquery{
				SubqueryType: &proto.Expression_Subquery_SetPredicate_{
					SetPredicate: &proto.Expression_Subquery_SetPredicate{
						PredicateOp: proto.Expression_Subquery_SetPredicate_PREDICATE_OP_EXISTS,
						Tuples:      scanProto,
					},
				},
			},
		},
		{
			name: "SetComparisonSubquery",
			subquery: &proto.Expression_Subquery{
				SubqueryType: &proto.Expression_Subquery_SetComparison_{
					SetComparison: &proto.Expression_Subquery_SetComparison{
						ReductionOp:  proto.Expression_Subquery_SetComparison_REDUCTION_OP_ANY,
						ComparisonOp: proto.Expression_Subquery_SetComparison_COMPARISON_OP_EQ,
						Left:         needle,
						Right:        scanProto,
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := &proto.Plan{
				Version: &proto.Version{MinorNumber: 29},
				Relations: []*proto.PlanRel{{
					RelType: &proto.PlanRel_Root{
						Root: &proto.RelRoot{
							Names: []string{"out"},
							Input: &proto.Rel{
								RelType: &proto.Rel_Filter{
									Filter: &proto.FilterRel{
										Common:    &proto.RelCommon{EmitKind: &proto.RelCommon_Direct_{Direct: &proto.RelCommon_Direct{}}},
										Input:     scanProto,
										Condition: &proto.Expression{RexType: &proto.Expression_Subquery_{Subquery: tc.subquery}},
									},
								},
							},
						},
					},
				}},
			}
			result, err := codec.PlanFromProto(p, c)
			require.NoError(t, err)
			require.NotNil(t, result)
		})
	}
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
	_, err := codec.PlanFromProto(&p, c)
	require.NoError(t, err)
}

// TestFromProtoWithDecoder is a full-plan integration test for FromProtoWithDecoder.
// It verifies decoder wiring and round-trip fidelity for all three extension rel types.
// Relation-level decoder error cases are covered by TestExtensionRelDecoder below.
func TestFromProtoWithDecoder(t *testing.T) {
	const typeURL = "type.googleapis.com/test.MyExtension"
	c := extensions.GetDefaultCollectionWithNoError()

	emit201 := &proto.RelCommon{EmitKind: &proto.RelCommon_Emit_{Emit: &proto.RelCommon_Emit{OutputMapping: []int32{2, 0, 1}}}}
	direct := &proto.RelCommon{EmitKind: &proto.RelCommon_Direct_{Direct: &proto.RelCommon_Direct{}}}
	detail := &anypb.Any{TypeUrl: typeURL, Value: []byte("irrelevant")}

	extSchema := *types.NewRecordTypeFromTypes([]types.Type{
		&types.Int64Type{Nullability: types.NullabilityRequired},
		&types.Int32Type{Nullability: types.NullabilityRequired},
		&types.StringType{Nullability: types.NullabilityRequired},
	})

	oneColInput := &proto.Rel{RelType: &proto.Rel_Read{Read: &proto.ReadRel{
		Common: direct,
		BaseSchema: &proto.NamedStruct{
			Names: []string{"a"},
			Struct: &proto.Type_Struct{
				Types: []*proto.Type{{Kind: &proto.Type_I64_{I64: &proto.Type_I64{Nullability: proto.Type_NULLABILITY_REQUIRED}}}},
			},
		},
		ReadType: &proto.ReadRel_VirtualTable_{VirtualTable: &proto.ReadRel_VirtualTable{}},
	}}}

	makePlan := func(extRel *proto.Rel) *proto.Plan {
		return &proto.Plan{Version: &proto.Version{MinorNumber: 29}, Relations: []*proto.PlanRel{{
			RelType: &proto.PlanRel_Root{Root: &proto.RelRoot{Names: []string{"c", "a", "b"}, Input: extRel}},
		}}}
	}

	successCases := []struct {
		name string
		plan *proto.Plan
	}{
		{
			name: "Single",
			plan: makePlan(&proto.Rel{RelType: &proto.Rel_ExtensionSingle{ExtensionSingle: &proto.ExtensionSingleRel{
				Common: emit201, Input: oneColInput, Detail: detail,
			}}}),
		},
		{
			name: "Leaf",
			plan: makePlan(&proto.Rel{RelType: &proto.Rel_ExtensionLeaf{ExtensionLeaf: &proto.ExtensionLeafRel{
				Common: emit201, Detail: detail,
			}}}),
		},
		{
			name: "Multi",
			plan: makePlan(&proto.Rel{RelType: &proto.Rel_ExtensionMulti{ExtensionMulti: &proto.ExtensionMultiRel{
				Common: emit201, Inputs: []*proto.Rel{oneColInput, oneColInput}, Detail: detail,
			}}}),
		},
	}

	for _, tc := range successCases {
		t.Run(tc.name, func(t *testing.T) {
			p, err := codec.PlanFromProtoWithDecoder(tc.plan, c, map[string]expr.ExtensionRelDecoder{typeURL: &customDecoder{schema: extSchema}})
			require.NoError(t, err)
			require.Len(t, p.Relations()[0].Root().RecordType().Struct.Types, 3)
		})
	}

	t.Run("decoder error propagates", func(t *testing.T) {
		plan := makePlan(&proto.Rel{RelType: &proto.Rel_ExtensionSingle{ExtensionSingle: &proto.ExtensionSingleRel{
			Common: direct, Input: oneColInput, Detail: detail,
		}}})
		_, err := codec.PlanFromProtoWithDecoder(plan, c, map[string]expr.ExtensionRelDecoder{typeURL: &errorDecoder{err: errors.New("decode failed")}})
		require.ErrorContains(t, err, "decode failed")
	})

}

// customExtDef is a test ExtensionRelDefinition that claims a fixed output schema.
type customExtDef struct {
	detail *anypb.Any
	schema types.RecordType
}

func (d *customExtDef) Schema(inputs []plan.Rel) types.RecordType  { return d.schema }
func (d *customExtDef) Build(_ []plan.Rel) *anypb.Any              { return d.detail }
func (d *customExtDef) Expressions(_ []plan.Rel) []expr.Expression { return nil }

// customDecoder returns a customExtDef with a fixed schema. The registry dispatches
// to it only for the type URL it was registered under.
type customDecoder struct {
	schema types.RecordType
}

func (cd *customDecoder) DecodeExtensionRel(detail *anypb.Any) (any, error) {
	return &customExtDef{detail: detail, schema: cd.schema}, nil
}

// errorDecoder always returns a non-nil error from DecodeExtensionRel.
type errorDecoder struct{ err error }

func (d *errorDecoder) DecodeExtensionRel(_ *anypb.Any) (any, error) { return nil, d.err }
