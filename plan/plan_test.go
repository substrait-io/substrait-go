package plan

import (
	"errors"
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

	i32Type := &proto.Type{Kind: &proto.Type_I32_{I32: &proto.Type_I32{Nullability: proto.Type_NULLABILITY_REQUIRED}}}
	i64Type := &proto.Type{Kind: &proto.Type_I64_{I64: &proto.Type_I64{Nullability: proto.Type_NULLABILITY_REQUIRED}}}
	strType := &proto.Type{Kind: &proto.Type_String_{String_: &proto.Type_String{Nullability: proto.Type_NULLABILITY_REQUIRED}}}

	// 3-column scan: col1 i32, col2 i64, col3 string
	scanProto := &proto.Rel{
		RelType: &proto.Rel_Read{
			Read: &proto.ReadRel{
				Common: &proto.RelCommon{EmitKind: &proto.RelCommon_Direct_{Direct: &proto.RelCommon_Direct{}}},
				BaseSchema: &proto.NamedStruct{
					Names: []string{"col1", "col2", "col3"},
					Struct: &proto.Type_Struct{
						Nullability: proto.Type_NULLABILITY_REQUIRED,
						Types:       []*proto.Type{i32Type, i64Type, strType},
					},
				},
				ReadType: &proto.ReadRel_NamedTable_{NamedTable: &proto.ReadRel_NamedTable{Names: []string{"t"}}},
			},
		},
	}
	needle := expr.NewPrimitiveLiteral(int32(1), false).ToProto()

	subqueries := []struct {
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

	commons := []struct {
		name   string
		common *proto.RelCommon
		// rootNames must match the number of output columns after the emit mapping
		rootNames []string
	}{
		{
			name:      "Direct",
			common:    &proto.RelCommon{EmitKind: &proto.RelCommon_Direct_{Direct: &proto.RelCommon_Direct{}}},
			rootNames: []string{"col1", "col2", "col3"},
		},
		{
			// out-of-order emit: col3, col1, col2 → output mapping [2, 0, 1]
			name: "Emit",
			common: &proto.RelCommon{EmitKind: &proto.RelCommon_Emit_{
				Emit: &proto.RelCommon_Emit{OutputMapping: []int32{2, 0, 1}},
			}},
			rootNames: []string{"col3", "col1", "col2"},
		},
	}

	for _, sq := range subqueries {
		for _, cm := range commons {
			t.Run(sq.name+"/"+cm.name, func(t *testing.T) {
				original := &proto.Plan{
					Relations: []*proto.PlanRel{{
						RelType: &proto.PlanRel_Root{
							Root: &proto.RelRoot{
								Names: cm.rootNames,
								Input: &proto.Rel{
									RelType: &proto.Rel_Filter{
										Filter: &proto.FilterRel{
											Common:    cm.common,
											Input:     scanProto,
											Condition: &proto.Expression{RexType: &proto.Expression_Subquery_{Subquery: sq.subquery}},
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

// TestProjectSetStyleDecoder models an extension that implements a
// set-returning function (SRF / unnest-style) relation:
//
//   - The extension detail is an anypb.Any whose Value encodes a list of SRF
//     expressions (one per output column beyond the passthrough input columns).
//   - The extension's output schema = input schema + one col per SRF expression,
//     typed from the expression's output type.
//   - A registered ExtensionRelDecoder (psStyleDecoder) unpacks the detail,
//     resolves each expression via ExprFromProto, and returns an
//     ExtensionRelDefinition with the full output schema.
//
// Without the decoder, UndecodedExtension returns only the input schema, so an
// emit mapping that references the SRF output column panics on RecordType().
// The test also verifies round-trip fidelity: ToProto() reproduces the original proto.
func TestProjectSetStyleDecoder(t *testing.T) {
	const projectSetTypeURL = "type.googleapis.com/test.ProjectSetRel"

	// inputSchema: one i64 column ("a") — the passthrough column.
	inputSchema := *types.NewRecordTypeFromTypes([]types.Type{
		&types.Int64Type{Nullability: types.NullabilityRequired},
	})

	// SRF expression: literal i32 (models an unnest() returning i32 rows).
	// The decoder will call ExprFromProto on this to learn the SRF output type.
	srfExprProto := &proto.Expression{
		RexType: &proto.Expression_Literal_{
			Literal: &proto.Expression_Literal{
				LiteralType: &proto.Expression_Literal_I32{I32: 42},
			},
		},
	}

	// Build the detail: Any{typeURL, value=marshal(Expression_Nested_Struct{srfExpr})}.
	// psStyleDecoder unmarshals this to recover the SRF expressions.
	// A real implementation would similarly marshal each SRF expression
	// into an Any-wrapped proto inside the extension's own message type.
	nestedStruct := &proto.Expression_Nested_Struct{
		Fields: []*proto.Expression{srfExprProto},
	}
	nestedBytes, err := protobuf.Marshal(nestedStruct)
	require.NoError(t, err)
	projectSetDetail := &anypb.Any{TypeUrl: projectSetTypeURL, Value: nestedBytes}

	// Build the input rel proto. directCommon is set so ToProto() round-trips exactly.
	inputRelProto := &proto.Rel{
		RelType: &proto.Rel_Read{
			Read: &proto.ReadRel{
				Common: directCommon,
				BaseSchema: &proto.NamedStruct{
					Names: []string{"a"},
					Struct: &proto.Type_Struct{
						Types: []*proto.Type{
							{Kind: &proto.Type_I64_{I64: &proto.Type_I64{Nullability: proto.Type_NULLABILITY_REQUIRED}}},
						},
					},
				},
				ReadType: &proto.ReadRel_VirtualTable_{VirtualTable: &proto.ReadRel_VirtualTable{}},
			},
		},
	}

	// ExtensionSingleRel with emit [0, 1]:
	//   col 0 = passthrough input col "a" (i64)
	//   col 1 = SRF output col (i32, from the literal expression)
	extSingleProto := &proto.ExtensionSingleRel{
		Common: emit01,
		Input:  inputRelProto,
		Detail: projectSetDetail,
	}
	rel := &proto.Rel{RelType: &proto.Rel_ExtensionSingle{ExtensionSingle: extSingleProto}}

	makeDecoder := func(reg expr.ExtensionRegistry) expr.ExtensionRelDecoder {
		return projectSetDecoderFor(projectSetTypeURL, inputSchema, reg)
	}

	t.Run("without decoder panics on emit OOB", func(t *testing.T) {
		reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
		require.Panics(t, func() {
			out, err := RelFromProto(rel, reg)
			require.NoError(t, err)
			_ = out.RecordType()
		})
	})

	t.Run("with decoder schema is input+srf", func(t *testing.T) {
		reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
		reg.SetExtensionRelDecoder(makeDecoder(reg))

		out, err := RelFromProto(rel, reg)
		require.NoError(t, err)

		got := out.RecordType()
		// emit [0,1] selects both cols from the 2-col schema — result is 2 cols.
		require.Equal(t, int32(2), got.FieldCount())
		// col 0 = i64 (passthrough), col 1 = i32 (SRF output)
		assert.IsType(t, &types.Int64Type{}, got.GetFieldRef(0))
		assert.IsType(t, &types.Int32Type{}, got.GetFieldRef(1))
	})

	t.Run("round-trip ToProto matches original", func(t *testing.T) {
		reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
		reg.SetExtensionRelDecoder(makeDecoder(reg))

		out, err := RelFromProto(rel, reg)
		require.NoError(t, err)

		got := out.ToProto()
		if diff := cmp.Diff(rel, got, protocmp.Transform()); diff != "" {
			t.Errorf("round-trip mismatch (-want +got):\n%s", diff)
		}
	})
}

// projectSetDecoderFor returns an ExtensionRelDecoder that handles the
// ProjectSet-style extension: it unpacks a nested struct of SRF expressions
// from the detail, resolves each expression's output type via ExprFromProto
// against the provided inputSchema, and returns an ExtensionRelDefinition
// whose Schema() = inputSchema + SRF output types.
//
// This mirrors the pattern used by SRF-style extension relations.
func projectSetDecoderFor(typeURL string, inputSchema types.RecordType, reg expr.ExtensionRegistry) expr.ExtensionRelDecoder {
	return &psStyleDecoder{typeURL: typeURL, inputSchema: inputSchema, reg: reg}
}

type psStyleDecoder struct {
	typeURL     string
	inputSchema types.RecordType
	reg         expr.ExtensionRegistry
}

func (d *psStyleDecoder) DecodeExtensionRel(detail *anypb.Any) (any, error) {
	if detail == nil || detail.TypeUrl != d.typeURL {
		return nil, nil
	}
	var ns proto.Expression_Nested_Struct
	if err := protobuf.Unmarshal(detail.Value, &ns); err != nil {
		return nil, err
	}
	srfFields := ns.GetFields()
	srfTypes := make([]types.Type, 0, len(srfFields))
	for _, f := range srfFields {
		stExpr, err := expr.ExprFromProto(f, &d.inputSchema, d.reg)
		if err != nil {
			return nil, err
		}
		srfTypes = append(srfTypes, stExpr.GetType())
	}
	fullSchema := d.inputSchema.Concat(*types.NewRecordTypeFromTypes(srfTypes))
	return &customExtDef{schema: fullSchema, detail: detail}, nil
}

// customExtDef is a test ExtensionRelDefinition that claims a fixed output schema.
type customExtDef struct {
	detail  *anypb.Any
	schema  types.RecordType
	typeURL string
}

func (d *customExtDef) Schema(inputs []Rel) types.RecordType { return d.schema }
func (d *customExtDef) Build(_ []Rel) *anypb.Any            { return d.detail }
func (d *customExtDef) Expressions(_ []Rel) []expr.Expression { return nil }

// customDecoder returns a customExtDef for a specific typeURL and nil for everything else.
type customDecoder struct {
	typeURL string
	schema  types.RecordType
}

func (cd *customDecoder) DecodeExtensionRel(detail *anypb.Any) (any, error) {
	if detail == nil || detail.TypeUrl != cd.typeURL {
		return nil, nil // not ours — fall back to UndecodedExtension
	}
	return &customExtDef{detail: detail, schema: cd.schema, typeURL: cd.typeURL}, nil
}

// errorDecoder always returns a non-nil error from DecodeExtensionRel.
type errorDecoder struct{ err error }

func (d *errorDecoder) DecodeExtensionRel(_ *anypb.Any) (any, error) { return nil, d.err }

// wrongTypeDecoder returns a value that does not implement ExtensionRelDefinition.
type wrongTypeDecoder struct{}

func (d *wrongTypeDecoder) DecodeExtensionRel(_ *anypb.Any) (any, error) { return "not-a-def", nil }

// extSchema2col is a 2-column RecordType used across extension decoder tests.
var extSchema2col = *types.NewRecordTypeFromTypes([]types.Type{
	&types.Int64Type{Nullability: types.NullabilityRequired},
	&types.Int32Type{Nullability: types.NullabilityRequired},
})

// emit01 is a RelCommon with output mapping [0, 1], sufficient to trigger remap OOB.
var emit01 = &proto.RelCommon{
	EmitKind: &proto.RelCommon_Emit_{
		Emit: &proto.RelCommon_Emit{OutputMapping: []int32{0, 1}},
	},
}

// directCommon is a RelCommon with no emit (direct pass-through), which is what
// substrait-go serialises back out on ToProto() when no mapping is set.
var directCommon = &proto.RelCommon{EmitKind: &proto.RelCommon_Direct_{Direct: &proto.RelCommon_Direct{}}}

// oneColInputRel is a VirtualTable ReadRel with a single i64 column.
var oneColInputRel = &proto.Rel{
	RelType: &proto.Rel_Read{
		Read: &proto.ReadRel{
			Common: directCommon,
			BaseSchema: &proto.NamedStruct{
				Names: []string{"a"},
				Struct: &proto.Type_Struct{
					Types: []*proto.Type{
						{Kind: &proto.Type_I64_{I64: &proto.Type_I64{Nullability: proto.Type_NULLABILITY_REQUIRED}}},
					},
				},
			},
			ReadType: &proto.ReadRel_VirtualTable_{VirtualTable: &proto.ReadRel_VirtualTable{}},
		},
	},
}

// TestExtensionRelDecoder_Single verifies the decoder hook for ExtensionSingleRel.
//
// ExtensionSingleRel wraps one input; UndecodedExtension.Schema returns the
// input schema unchanged, so an emit mapping that references an extra column
// added by the extension will panic on RecordType(). A registered decoder
// provides the correct wider schema, preventing the panic.
func TestExtensionRelDecoder_Single(t *testing.T) {
	const typeURL = "type.googleapis.com/test.MyExtension"
	detail := &anypb.Any{TypeUrl: typeURL, Value: []byte("irrelevant")}

	// ExtensionSingleRel: emit [0,1] — index 1 is the extension's extra column.
	rel := &proto.Rel{RelType: &proto.Rel_ExtensionSingle{ExtensionSingle: &proto.ExtensionSingleRel{
		Common: emit01,
		Input:  oneColInputRel,
		Detail: detail,
	}}}

	t.Run("without decoder panics on emit OOB", func(t *testing.T) {
		reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
		require.Panics(t, func() {
			out, err := RelFromProto(rel, reg)
			require.NoError(t, err)
			_ = out.RecordType()
		})
	})

	t.Run("with decoder uses custom schema", func(t *testing.T) {
		reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
		reg.SetExtensionRelDecoder(&customDecoder{typeURL: typeURL, schema: extSchema2col})

		out, err := RelFromProto(rel, reg)
		require.NoError(t, err)
		require.Equal(t, int32(2), out.RecordType().FieldCount())
	})

	t.Run("round-trip ToProto matches original", func(t *testing.T) {
		reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
		reg.SetExtensionRelDecoder(&customDecoder{typeURL: typeURL, schema: extSchema2col})

		out, err := RelFromProto(rel, reg)
		require.NoError(t, err)

		if diff := cmp.Diff(rel, out.ToProto(), protocmp.Transform()); diff != "" {
			t.Errorf("round-trip mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("decoder returning nil falls back to UndecodedExtension", func(t *testing.T) {
		reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
		reg.SetExtensionRelDecoder(&customDecoder{typeURL: "type.googleapis.com/other.Type"})
		require.Panics(t, func() {
			out, err := RelFromProto(rel, reg)
			require.NoError(t, err)
			_ = out.RecordType()
		})
	})

	t.Run("decoder returning error propagates to RelFromProto", func(t *testing.T) {
		reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
		sentinel := errors.New("decode failed")
		reg.SetExtensionRelDecoder(&errorDecoder{err: sentinel})
		_, err := RelFromProto(rel, reg)
		require.ErrorContains(t, err, "decode failed")
	})

	t.Run("decoder returning wrong type errors", func(t *testing.T) {
		reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
		reg.SetExtensionRelDecoder(&wrongTypeDecoder{})
		_, err := RelFromProto(rel, reg)
		require.ErrorContains(t, err, "does not implement ExtensionRelDefinition")
	})
}

// TestExtensionRelDecoder_Leaf verifies the decoder hook for ExtensionLeafRel.
//
// ExtensionLeafRel has no inputs; UndecodedExtension.Schema always returns an
// empty RecordType, so any non-empty emit mapping panics on RecordType().
func TestExtensionRelDecoder_Leaf(t *testing.T) {
	const typeURL = "type.googleapis.com/test.MyLeafExtension"
	detail := &anypb.Any{TypeUrl: typeURL, Value: []byte("irrelevant")}

	// ExtensionLeafRel: emit [0,1] — both columns come from the extension schema.
	rel := &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{ExtensionLeaf: &proto.ExtensionLeafRel{
		Common: emit01,
		Detail: detail,
	}}}

	t.Run("without decoder panics on emit OOB", func(t *testing.T) {
		reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
		require.Panics(t, func() {
			out, err := RelFromProto(rel, reg)
			require.NoError(t, err)
			_ = out.RecordType()
		})
	})

	t.Run("with decoder uses custom schema", func(t *testing.T) {
		reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
		reg.SetExtensionRelDecoder(&customDecoder{typeURL: typeURL, schema: extSchema2col})

		out, err := RelFromProto(rel, reg)
		require.NoError(t, err)
		require.Equal(t, int32(2), out.RecordType().FieldCount())
	})

	t.Run("round-trip ToProto matches original", func(t *testing.T) {
		reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
		reg.SetExtensionRelDecoder(&customDecoder{typeURL: typeURL, schema: extSchema2col})

		out, err := RelFromProto(rel, reg)
		require.NoError(t, err)

		if diff := cmp.Diff(rel, out.ToProto(), protocmp.Transform()); diff != "" {
			t.Errorf("round-trip mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("decoder returning nil falls back to UndecodedExtension", func(t *testing.T) {
		reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
		reg.SetExtensionRelDecoder(&customDecoder{typeURL: "type.googleapis.com/other.Type"})
		require.Panics(t, func() {
			out, err := RelFromProto(rel, reg)
			require.NoError(t, err)
			_ = out.RecordType()
		})
	})

	t.Run("decoder returning error propagates to RelFromProto", func(t *testing.T) {
		reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
		reg.SetExtensionRelDecoder(&errorDecoder{err: errors.New("decode failed")})
		_, err := RelFromProto(rel, reg)
		require.ErrorContains(t, err, "decode failed")
	})

	t.Run("decoder returning wrong type errors", func(t *testing.T) {
		reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
		reg.SetExtensionRelDecoder(&wrongTypeDecoder{})
		_, err := RelFromProto(rel, reg)
		require.ErrorContains(t, err, "does not implement ExtensionRelDefinition")
	})
}

// TestExtensionRelDecoder_Multi verifies the decoder hook for ExtensionMultiRel.
//
// ExtensionMultiRel has multiple inputs; UndecodedExtension.Schema returns an
// empty RecordType (len(inputs) != 1), so any non-empty emit mapping panics.
func TestExtensionRelDecoder_Multi(t *testing.T) {
	const typeURL = "type.googleapis.com/test.MyMultiExtension"
	detail := &anypb.Any{TypeUrl: typeURL, Value: []byte("irrelevant")}

	// Two identical one-column inputs; the extension combines them into 2 cols.
	rel := &proto.Rel{RelType: &proto.Rel_ExtensionMulti{ExtensionMulti: &proto.ExtensionMultiRel{
		Common:  emit01,
		Inputs:  []*proto.Rel{oneColInputRel, oneColInputRel},
		Detail:  detail,
	}}}

	t.Run("without decoder panics on emit OOB", func(t *testing.T) {
		reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
		require.Panics(t, func() {
			out, err := RelFromProto(rel, reg)
			require.NoError(t, err)
			_ = out.RecordType()
		})
	})

	t.Run("with decoder uses custom schema", func(t *testing.T) {
		reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
		reg.SetExtensionRelDecoder(&customDecoder{typeURL: typeURL, schema: extSchema2col})

		out, err := RelFromProto(rel, reg)
		require.NoError(t, err)
		require.Equal(t, int32(2), out.RecordType().FieldCount())
	})

	t.Run("round-trip ToProto matches original", func(t *testing.T) {
		reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
		reg.SetExtensionRelDecoder(&customDecoder{typeURL: typeURL, schema: extSchema2col})

		out, err := RelFromProto(rel, reg)
		require.NoError(t, err)

		if diff := cmp.Diff(rel, out.ToProto(), protocmp.Transform()); diff != "" {
			t.Errorf("round-trip mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("decoder returning nil falls back to UndecodedExtension", func(t *testing.T) {
		reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
		reg.SetExtensionRelDecoder(&customDecoder{typeURL: "type.googleapis.com/other.Type"})
		require.Panics(t, func() {
			out, err := RelFromProto(rel, reg)
			require.NoError(t, err)
			_ = out.RecordType()
		})
	})

	t.Run("decoder returning error propagates to RelFromProto", func(t *testing.T) {
		reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
		reg.SetExtensionRelDecoder(&errorDecoder{err: errors.New("decode failed")})
		_, err := RelFromProto(rel, reg)
		require.ErrorContains(t, err, "decode failed")
	})

	t.Run("decoder returning wrong type errors", func(t *testing.T) {
		reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
		reg.SetExtensionRelDecoder(&wrongTypeDecoder{})
		_, err := RelFromProto(rel, reg)
		require.ErrorContains(t, err, "does not implement ExtensionRelDefinition")
	})
}

// TestStructFieldRefGetTypeBoundsCheck verifies that StructFieldRef.GetType returns
// an error (not a panic) when the field index equals the length of the struct types.
func TestStructFieldRefGetTypeBoundsCheck(t *testing.T) {
	st := &types.StructType{
		Types: []types.Type{
			&types.Int64Type{Nullability: types.NullabilityRequired},
		},
	}
	// Field == len(Types) must return ErrInvalidType, not panic with an OOB access.
	ref := &expr.StructFieldRef{Field: 1}
	_, err := ref.GetType(st)
	require.ErrorIs(t, err, substraitgo.ErrInvalidType)
}
