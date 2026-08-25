package plan

import (
	"errors"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	substraitgo "github.com/substrait-io/substrait-go/v9"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	extensionspb "github.com/substrait-io/substrait-protobuf/go/substraitpb/extensions"
	"google.golang.org/protobuf/encoding/protojson"
	protobuf "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/anypb"
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
		Version: &proto.Version{MinorNumber: 29},
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
			p, err := FromProto(original, c)
			require.NoError(t, err)
			roundTripped, err := p.ToProto()
			require.NoError(t, err)
			// Use cmp.Diff instead of protojson for comparison: protojson output is non-deterministic.
			if diff := cmp.Diff(original, roundTripped, protocmp.Transform()); diff != "" {
				t.Errorf("plan round-trip mismatch (-want +got):\n%s", diff)
			}
		})
	}
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
			result, err := FromProto(p, c)
			require.NoError(t, err)
			require.NotNil(t, result)
		})
	}
}

func TestBuilderPlanRegistryWithSubqueries(t *testing.T) {
	b := NewBuilderDefault()
	scan := b.NamedScan([]string{"t"}, types.NamedStruct{
		Names:  []string{"col1"},
		Struct: types.StructType{Types: []types.Type{&types.Int32Type{}}},
	})
	p, err := b.Plan(scan, []string{"col1"})
	require.NoError(t, err)

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
	subqueryExprProto := &proto.Expression{
		RexType: &proto.Expression_Subquery_{
			Subquery: &proto.Expression_Subquery{
				SubqueryType: &proto.Expression_Subquery_Scalar_{
					Scalar: &proto.Expression_Subquery_Scalar{Input: scanProto},
				},
			},
		},
	}

	result, err := expr.ExprFromProto(subqueryExprProto, &types.RecordType{}, p.ExtensionRegistry())
	require.NoError(t, err)
	require.NotNil(t, result)
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
			p, err := FromProtoWithDecoder(tc.plan, c, map[string]expr.ExtensionRelDecoder{typeURL: &customDecoder{schema: extSchema}})
			require.NoError(t, err)
			require.Len(t, p.Relations()[0].Root().RecordType().Struct.Types, 3)

			roundTripped, err := p.ToProto()
			require.NoError(t, err)
			if diff := cmp.Diff(tc.plan, roundTripped, protocmp.Transform()); diff != "" {
				t.Errorf("round-trip mismatch (-want +got):\n%s", diff)
			}
		})
	}

	t.Run("decoder error propagates", func(t *testing.T) {
		plan := makePlan(&proto.Rel{RelType: &proto.Rel_ExtensionSingle{ExtensionSingle: &proto.ExtensionSingleRel{
			Common: direct, Input: oneColInput, Detail: detail,
		}}})
		_, err := FromProtoWithDecoder(plan, c, map[string]expr.ExtensionRelDecoder{typeURL: &errorDecoder{err: errors.New("decode failed")}})
		require.ErrorContains(t, err, "decode failed")
	})

}

func TestIsRecordTypeSupported(t *testing.T) {
	fixedSchema := *types.NewRecordTypeFromTypes([]types.Type{
		&types.Int64Type{Nullability: types.NullabilityRequired},
	})
	decoded := &customExtDef{schema: fixedSchema}
	undecoded := &UndecodedExtension{}

	assert.True(t, isRecordTypeSupported(&ExtensionSingleRel{definition: decoded}))
	assert.True(t, isRecordTypeSupported(&ExtensionLeafRel{definition: decoded}))
	assert.True(t, isRecordTypeSupported(&ExtensionMultiRel{definition: decoded}))

	assert.False(t, isRecordTypeSupported(&ExtensionSingleRel{definition: undecoded}))
	assert.False(t, isRecordTypeSupported(&ExtensionLeafRel{definition: undecoded}))
	assert.False(t, isRecordTypeSupported(&ExtensionMultiRel{definition: undecoded}))
}

// customExtDef is a test ExtensionRelDefinition that claims a fixed output schema.
type customExtDef struct {
	detail *anypb.Any
	schema types.RecordType
}

func (d *customExtDef) Schema(inputs []Rel) types.RecordType  { return d.schema }
func (d *customExtDef) Build(_ []Rel) *anypb.Any              { return d.detail }
func (d *customExtDef) Expressions(_ []Rel) []expr.Expression { return nil }

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

// wrongTypeDecoder returns a value that does not implement ExtensionRelDefinition.
type wrongTypeDecoder struct{}

func (d *wrongTypeDecoder) DecodeExtensionRel(_ *anypb.Any) (any, error) { return "not-a-def", nil }

// TestExtensionRelDecoder verifies the decoder hook for all three extension rel types.
// Without a decoder, UndecodedExtension returns a schema too narrow for the emit
// mapping, causing a panic on RecordType(). A registered decoder provides the correct
// wider schema.
func TestExtensionRelDecoder(t *testing.T) {
	const typeURL = "type.googleapis.com/test.MyExtension"
	detail := &anypb.Any{TypeUrl: typeURL, Value: []byte("irrelevant")}

	// extSchema is the 3-column output schema the decoder claims for the extension.
	// It is wider than oneColInput, so emit [2,0,1] would OOB without a decoder.
	extSchema := *types.NewRecordTypeFromTypes([]types.Type{
		&types.Int64Type{Nullability: types.NullabilityRequired},
		&types.Int32Type{Nullability: types.NullabilityRequired},
		&types.StringType{Nullability: types.NullabilityRequired},
	})

	// direct is a pass-through RelCommon; emit201 exercises out-of-order reordering.
	// Both are tested to cover the two common real-world mapping shapes.
	direct := &proto.RelCommon{EmitKind: &proto.RelCommon_Direct_{Direct: &proto.RelCommon_Direct{}}}
	emit201 := &proto.RelCommon{EmitKind: &proto.RelCommon_Emit_{Emit: &proto.RelCommon_Emit{OutputMapping: []int32{2, 0, 1}}}}

	// oneColInput is the single-column input fed into Single/Multi extension rels.
	// The extension decoder claims a wider schema, so OOB only occurs without it.
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

	// rels covers all three extension rel types × direct and out-of-order emit mappings.
	// Direct tests the no-remap path; Emit tests the OOB-without-decoder path.
	rels := []struct {
		name                 string
		rel                  *proto.Rel
		panicsWithoutDecoder bool // true when emit mapping exceeds UndecodedExtension schema
	}{
		{
			name: "Single/Direct",
			rel: &proto.Rel{RelType: &proto.Rel_ExtensionSingle{ExtensionSingle: &proto.ExtensionSingleRel{
				Common: direct, Input: oneColInput, Detail: detail,
			}}},
		},
		{
			name:                 "Single/Emit",
			panicsWithoutDecoder: true,
			rel: &proto.Rel{RelType: &proto.Rel_ExtensionSingle{ExtensionSingle: &proto.ExtensionSingleRel{
				Common: emit201, Input: oneColInput, Detail: detail,
			}}},
		},
		{
			name: "Leaf/Direct",
			rel: &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{ExtensionLeaf: &proto.ExtensionLeafRel{
				Common: direct, Detail: detail,
			}}},
		},
		{
			name:                 "Leaf/Emit",
			panicsWithoutDecoder: true,
			rel: &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{ExtensionLeaf: &proto.ExtensionLeafRel{
				Common: emit201, Detail: detail,
			}}},
		},
		{
			name: "Multi/Direct",
			rel: &proto.Rel{RelType: &proto.Rel_ExtensionMulti{ExtensionMulti: &proto.ExtensionMultiRel{
				Common: direct, Inputs: []*proto.Rel{oneColInput, oneColInput}, Detail: detail,
			}}},
		},
		{
			name:                 "Multi/Emit",
			panicsWithoutDecoder: true,
			rel: &proto.Rel{RelType: &proto.Rel_ExtensionMulti{ExtensionMulti: &proto.ExtensionMultiRel{
				Common: emit201, Inputs: []*proto.Rel{oneColInput, oneColInput}, Detail: detail,
			}}},
		},
	}

	for _, tc := range rels {
		t.Run(tc.name, func(t *testing.T) {
			if tc.panicsWithoutDecoder {
				// Without a decoder, UndecodedExtension returns a schema narrower than the
				// emit mapping, so RecordType() panics on OOB access.
				t.Run("without decoder panics on emit OOB", func(t *testing.T) {
					reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
					require.Panics(t, func() {
						out, err := RelFromProto(tc.rel, reg)
						require.NoError(t, err)
						_ = out.RecordType()
					})
				})

				// Registering a decoder under a different type URL leaves this detail
				// unmatched, so RelFromProto falls back to UndecodedExtension and panics on OOB.
				t.Run("decoder registered under different typeURL falls back to UndecodedExtension", func(t *testing.T) {
					reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
					require.NoError(t, reg.SetExtensionRelDecoder("type.googleapis.com/other.Type", &customDecoder{schema: extSchema}))
					require.Panics(t, func() {
						out, err := RelFromProto(tc.rel, reg)
						require.NoError(t, err)
						_ = out.RecordType()
					})
				})
			}

			// With a decoder, the custom schema is used and RecordType() returns all 3 cols.
			t.Run("with decoder uses custom schema", func(t *testing.T) {
				reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
				require.NoError(t, reg.SetExtensionRelDecoder(typeURL, &customDecoder{schema: extSchema}))
				out, err := RelFromProto(tc.rel, reg)
				require.NoError(t, err)
				require.Equal(t, int32(3), out.RecordType().FieldCount())
			})

			// A decoder error is propagated directly to the RelFromProto caller.
			t.Run("decoder returning error propagates to RelFromProto", func(t *testing.T) {
				reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
				require.NoError(t, reg.SetExtensionRelDecoder(typeURL, &errorDecoder{err: errors.New("decode failed")}))
				_, err := RelFromProto(tc.rel, reg)
				require.ErrorContains(t, err, "decode failed")
			})

			// DecodeExtensionRel returns any; returning a non-ExtensionRelDefinition
			// value is caught at cast time and surfaced as an error.
			t.Run("DecodeExtensionRel returns the wrong type errors", func(t *testing.T) {
				reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
				require.NoError(t, reg.SetExtensionRelDecoder(typeURL, &wrongTypeDecoder{}))
				_, err := RelFromProto(tc.rel, reg)
				require.ErrorContains(t, err, "does not implement ExtensionRelDefinition")
			})
		})
	}
}
