// SPDX-License-Identifier: Apache-2.0

package wire_test

import (
	"bytes"
	"encoding/json"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/goccy/go-yaml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/expr"
	ext "github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/plan"
	"github.com/substrait-io/substrait-go/v9/types"
	"github.com/substrait-io/substrait-go/v9/types/parser"
	"github.com/substrait-io/substrait-go/v9/wire"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/encoding/protojson"
	pb "google.golang.org/protobuf/proto"
)

func sampleNestedExpr(reg expr.ExtensionRegistry, substraitExtURN string) expr.Expression {
	var (
		add = ext.NewScalarFuncVariant(ext.FunctionID{URN: substraitExtURN, Name: "add"})
		sub = ext.NewScalarFuncVariant(ext.FunctionID{URN: substraitExtURN, Name: "subtract"})
		mul = ext.NewScalarFuncVariant(ext.FunctionID{URN: substraitExtURN, Name: "multiply"})
	)

	baseSchema := types.NewRecordTypeFromTypes(
		[]types.Type{
			&types.BooleanType{},
			&types.Int32Type{},
			&types.Int64Type{},
			&types.Float32Type{},
		})

	// add(literal, sub(ref, mul(literal, ref)))
	exp := expr.MustExpr(expr.NewCustomScalarFunc(reg, add, &types.Float64Type{}, nil,
		expr.NewPrimitiveLiteral(float64(1.0), false),
		expr.MustExpr(expr.NewCustomScalarFunc(reg, sub, &types.Float32Type{}, nil,
			expr.MustExpr(expr.NewRootFieldRef(expr.NewStructFieldRef(3), baseSchema)),
			expr.MustExpr(expr.NewCustomScalarFunc(reg, mul, &types.Int64Type{}, nil,
				expr.NewPrimitiveLiteral(int64(2), false),
				expr.MustExpr(expr.NewFieldRef(expr.NewNestedLiteral(expr.StructLiteralValue{
					expr.NewByteSliceLiteral([]byte("baz"), true),
					expr.NewPrimitiveLiteral("foobar", false),
					expr.NewPrimitiveLiteral(int32(5), false),
				}, false), expr.NewStructFieldRef(2), nil)),
			)),
		)),
	))

	return exp
}

func TestExpressionsRoundtrip(t *testing.T) {
	const substraitExtURN = "extension:io.substrait:functions_arithmetic"
	// define extensions with no plan for now
	const planExt = `{
		"extensionUrns": [
			{
				"extensionUrnAnchor": 1,
				"urn": "` + substraitExtURN + `"
			}
		],
		"extensions": [
			{
				"extensionFunction": {
					"extensionUrnReference": 1,
					"functionAnchor": 2,
					"name": "add:fp64_fp64"
				}
			},
			{
				"extensionFunction": {
					"extensionUrnReference": 1,
					"functionAnchor": 3,
					"name": "subtract:fp32_fp32"
				}
			},
			{
				"extensionFunction": {
					"extensionUrnReference": 1,
					"functionAnchor": 4,
					"name": "multiply:i64_i64"
				}
			},
			{
				"extensionFunction": {
					"extensionUrnReference": 1,
					"functionAnchor": 5,
					"name": "ntile:"
				}
			}
		],
		"relations": []
	}`

	var (
		plan proto.Plan
	)
	if err := protojson.Unmarshal([]byte(planExt), &plan); err != nil {
		panic(err)
	}
	// get the extension set
	collection := ext.GetDefaultCollectionWithNoError()
	extSet, err := wire.GetExtensionSet(&plan, collection)
	if err != nil {
		panic(err)
	}
	reg := expr.NewExtensionRegistry(extSet, collection)
	tests := []expr.Expression{
		sampleNestedExpr(reg, substraitExtURN),
	}

	for _, exp := range tests {
		protoExpr := wire.ExprToProto(exp)
		out, err := wire.ExprFromProto(protoExpr, nil, reg)
		require.NoError(t, err)
		assert.Truef(t, exp.Equals(out), "expected: %s\ngot: %s", exp, out)
	}
}

func TestScalarFunctionMissingOutputTypeReturnsError(t *testing.T) {
	registry := expr.NewEmptyExtensionRegistry(ext.GetDefaultCollectionWithNoError())
	functionReference := registry.GetFuncAnchor(ext.FunctionID{
		URN:  "extension:io.substrait:functions_arithmetic",
		Name: "add:i64_i64",
	})

	_, err := wire.ExprFromProto(&proto.Expression{
		RexType: &proto.Expression_ScalarFunction_{ScalarFunction: &proto.Expression_ScalarFunction{
			FunctionReference: functionReference,
			// OutputType intentionally omitted.
			Arguments: []*proto.FunctionArgument{
				literalI64Arg(1),
				literalI64Arg(2),
			},
		}},
	}, nil, registry)

	require.Error(t, err)
}

func TestWindowFunctionMissingOutputTypeReturnsError(t *testing.T) {
	registry := expr.NewEmptyExtensionRegistry(ext.GetDefaultCollectionWithNoError())
	functionReference := registry.GetFuncAnchor(ext.FunctionID{
		URN:  "extension:io.substrait:functions_arithmetic",
		Name: "sum:i64",
	})

	_, err := wire.ExprFromProto(&proto.Expression{
		RexType: &proto.Expression_WindowFunction_{WindowFunction: &proto.Expression_WindowFunction{
			FunctionReference: functionReference,
			// OutputType intentionally omitted.
		}},
	}, nil, registry)

	require.Error(t, err)
}

func TestCastMissingTypeReturnsError(t *testing.T) {
	registry := expr.NewEmptyExtensionRegistry(ext.GetDefaultCollectionWithNoError())

	_, err := wire.ExprFromProto(&proto.Expression{
		RexType: &proto.Expression_Cast_{Cast: &proto.Expression_Cast{
			// Type intentionally omitted.
			Input: &proto.Expression{
				RexType: &proto.Expression_Literal_{Literal: &proto.Expression_Literal{
					LiteralType: &proto.Expression_Literal_I64{I64: 1},
				}},
			},
		}},
	}, nil, registry)

	require.Error(t, err)
}

func literalI64Arg(value int64) *proto.FunctionArgument {
	return &proto.FunctionArgument{ArgType: &proto.FunctionArgument_Value{Value: &proto.Expression{
		RexType: &proto.Expression_Literal_{Literal: &proto.Expression_Literal{
			LiteralType: &proto.Expression_Literal_I64{I64: value},
		}},
	}}}
}

func TestRoundTripUsingTestData(t *testing.T) {
	const substraitExtURN = "extension:io.substrait:functions_arithmetic"
	// define extensions with no plan for now
	const planExt = `{
		"extensionUrns": [
			{
				"extensionUrnAnchor": 1,
				"urn": "` + substraitExtURN + `"
			}
		],
		"extensions": [
			{
				"extensionFunction": {
					"extensionUrnReference": 1,
					"functionAnchor": 2,
					"name": "add:fp64_fp64"
				}
			},
			{
				"extensionFunction": {
					"extensionUrnReference": 1,
					"functionAnchor": 3,
					"name": "subtract:fp64_fp64"
				}
			},
			{
				"extensionFunction": {
					"extensionUrnReference": 1,
					"functionAnchor": 4,
					"name": "multiply:fp64_fp64"
				}
			},
			{
				"extensionFunction": {
					"extensionUrnReference": 1,
					"functionAnchor": 5,
					"name": "ntile:i32"
				}
			}
		],
		"relations": []
	}`

	var (
		plan proto.Plan
	)
	if err := protojson.Unmarshal([]byte(planExt), &plan); err != nil {
		panic(err)
	}
	// get the extension set
	collection := ext.GetDefaultCollectionWithNoError()
	extSet, err := wire.GetExtensionSet(&plan, collection)
	if err != nil {
		panic(err)
	}

	f, err := os.Open("./testdata/expressions.yaml")
	require.NoError(t, err)
	defer f.Close()

	dec := yaml.NewDecoder(f)
	var tmp map[string]any
	require.NoError(t, dec.Decode(&tmp))

	var (
		protoSchema proto.NamedStruct
	)

	raw, err := json.Marshal(tmp["baseSchema"])
	require.NoError(t, err)
	require.NoError(t, protojson.Unmarshal(raw, &protoSchema))
	baseSchema := wire.NamedStructFromProto(&protoSchema)
	reg := expr.NewExtensionRegistry(extSet, collection)
	for _, tc := range tmp["cases"].([]any) {
		tt := tc.(map[string]any)
		t.Run(tt["name"].(string), func(t *testing.T) {
			test := tt["__test"].(map[string]any)

			var buf bytes.Buffer
			enc := json.NewEncoder(&buf)
			require.NoError(t, enc.Encode(tt["expression"]))
			var ex proto.Expression
			require.NoError(t, protojson.Unmarshal(buf.Bytes(), &ex))

			e, err := wire.ExprFromProto(&ex, types.NewRecordTypeFromStruct(baseSchema.Struct), reg)
			require.NoError(t, err)

			result := wire.ExprToProto(e)
			assert.Truef(t, pb.Equal(&ex, result), "expected: %s\ngot: %s", &ex, result)

			assert.True(t, e.Equals(e))

			if typTest, ok := test["type"].(string); ok {
				exp, err := parser.ParseType(typTest)
				require.NoError(t, err)

				assert.Equal(t, exp.String(), e.GetType().String())
			}

			strvalue, ok := test["string"].(string)
			if ok {
				strvalue = strings.TrimSpace(strvalue)
				t.Run("tostring", func(t *testing.T) {
					assert.Equal(t, strvalue, e.String())
				})
			}
		})
	}
}

func TestRoundTripExtendedExpression(t *testing.T) {
	f, err := os.Open("./testdata/extended_exprs.yaml")
	require.NoError(t, err)
	defer f.Close()

	dec := yaml.NewDecoder(f)
	var tmp map[string]any
	require.NoError(t, dec.Decode(&tmp))

	for _, tc := range tmp["tests"].([]any) {
		tt := tc.(map[string]any)

		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		require.NoError(t, enc.Encode(tt))
		var ex proto.ExtendedExpression
		require.NoError(t, protojson.Unmarshal(buf.Bytes(), &ex))

		result, err := wire.ExtendedFromProto(&ex, ext.GetDefaultCollectionWithNoError())
		require.NoError(t, err)

		out := wire.ExtendedToProto(result)
		// because we read the extensions into a map, we can't guarantee
		// the order of the extensions. But we also don't care about the
		// order, so we can just sort them by functionAnchor to ensure
		// they match for pb.Equal
		sort.Slice(out.Extensions, func(i, j int) bool {
			return out.Extensions[i].GetExtensionFunction().FunctionAnchor <
				out.Extensions[j].GetExtensionFunction().FunctionAnchor
		})
		assert.Truef(t, pb.Equal(&ex, out), "expected: %s\ngot: %s", &ex, out)
	}
}

func TestCastVisit(t *testing.T) {
	var builder = plan.NewBuilderDefault()
	castExpr := expr.MustExpr(builder.GetExprBuilder().Cast(builder.GetExprBuilder().Wrap(
		expr.NewLiteral[float64](12.0, true)),
		&types.Float64Type{Nullability: types.NullabilityRequired}).FailBehavior(
		types.CastFailBehaviorThrowException).BuildExpr())

	type relationTestCase struct {
		name            string
		rewriteFunction func(rex expr.Expression) expr.Expression
		want            float64
	}
	testCases := []relationTestCase{
		{"no change", func(ex expr.Expression) expr.Expression { return ex }, 12},
		{"changed", func(ex expr.Expression) expr.Expression {
			lit, err := expr.NewLiteral[float64](16.0, true)
			require.NoError(t, err)
			return lit
		}, 16},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			visitedCastExpr := castExpr.Visit(tc.rewriteFunction)
			visitedCastProto := wire.ExprToProto(visitedCastExpr)
			assert.IsType(t, &proto.Expression_Cast_{}, visitedCastProto.GetRexType())
			assert.Equal(t, tc.want, visitedCastProto.GetCast().GetInput().GetLiteral().GetFp64())
		})
	}
}

func TestSubqueryExpressionRoundtrip(t *testing.T) {
	const substraitExtURN = "extension:io.substrait:functions_arithmetic"
	// define extensions with no plan for now
	const planExt = `{
		"extensionUrns": [
			{
				"extensionUrnAnchor": 1,
				"urn": "` + substraitExtURN + `"
			}
		],
		"extensions": [],
		"relations": []
	}`

	var planProto proto.Plan
	if err := protojson.Unmarshal([]byte(planExt), &planProto); err != nil {
		panic(err)
	}

	// get the extension set and create registry with subquery handler
	c := ext.GetDefaultCollectionWithNoError()
	extSet, err := wire.GetExtensionSet(&planProto, c)
	if err != nil {
		panic(err)
	}

	// Create extension registry with subquery handler properly
	baseReg := expr.NewExtensionRegistry(extSet, c)

	// Create a simple mock relation for subqueries - single column of int32
	mockSchema := types.NamedStruct{
		Names: []string{"col1"},
		Struct: types.StructType{
			Types: []types.Type{&types.Int32Type{}},
		},
	}
	mockRel := plan.NewBuilderDefault().NamedScan([]string{"test_table"}, mockSchema)

	// Create base schema for needle expressions
	baseSchema := types.NewRecordTypeFromTypes([]types.Type{&types.Int32Type{}, &types.StringType{}})

	tests := []struct {
		name    string
		subExpr expr.Expression
	}{
		{
			name:    "ScalarSubquery",
			subExpr: plan.NewScalarSubquery(mockRel),
		},
		{
			name: "InPredicateSubquery",
			subExpr: plan.NewInPredicateSubquery(
				[]expr.Expression{expr.NewPrimitiveLiteral(int32(42), false)},
				mockRel,
			),
		},
		{
			name: "InPredicateSubquery_MultipleNeedles",
			subExpr: func() expr.Expression {
				// Create a 2-column relation for multi-needle test
				twoColSchema := types.NamedStruct{
					Names: []string{"col1", "col2"},
					Struct: types.StructType{
						Types: []types.Type{&types.Int32Type{}, &types.StringType{}},
					},
				}
				twoColRel := plan.NewBuilderDefault().NamedScan([]string{"two_col_table"}, twoColSchema)

				return plan.NewInPredicateSubquery(
					[]expr.Expression{
						expr.NewPrimitiveLiteral(int32(42), false),
						expr.NewPrimitiveLiteral("test", false),
					},
					twoColRel,
				)
			}(),
		},
		{
			name: "SetPredicateSubquery_EXISTS",
			subExpr: plan.NewSetPredicateSubquery(
				plan.SetPredicateOpExists,
				mockRel,
			),
		},
		{
			name: "SetPredicateSubquery_UNIQUE",
			subExpr: plan.NewSetPredicateSubquery(
				plan.SetPredicateOpUnique,
				mockRel,
			),
		},
		{
			name: "SetComparisonSubquery_ANY_EQ",
			subExpr: plan.NewSetComparisonSubquery(
				plan.SetComparisonReductionOpAny,
				plan.SetComparisonOpEq,
				expr.NewPrimitiveLiteral(int32(42), false),
				mockRel,
			),
		},
		{
			name: "SetComparisonSubquery_ALL_GT",
			subExpr: plan.NewSetComparisonSubquery(
				plan.SetComparisonReductionOpAll,
				plan.SetComparisonOpGt,
				expr.NewPrimitiveLiteral(int32(100), false),
				mockRel,
			),
		},
		{
			name: "SetComparisonSubquery_ANY_NE",
			subExpr: plan.NewSetComparisonSubquery(
				plan.SetComparisonReductionOpAny,
				plan.SetComparisonOpNe,
				expr.NewPrimitiveLiteral(int32(0), false),
				mockRel,
			),
		},
		{
			name: "SetComparisonSubquery_ALL_LE",
			subExpr: plan.NewSetComparisonSubquery(
				plan.SetComparisonReductionOpAll,
				plan.SetComparisonOpLe,
				expr.NewPrimitiveLiteral(int32(50), false),
				mockRel,
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Convert expression to protobuf
			protoExpr := wire.ExprToProto(tt.subExpr)
			require.NotNil(t, protoExpr)
			require.NotNil(t, protoExpr.GetSubquery())

			// Convert back from protobuf using ExprFromProto with subquery handler
			fromProto, err := wire.ExprFromProto(protoExpr, baseSchema, baseReg)
			require.NoError(t, err)
			require.NotNil(t, fromProto)

			// Verify that we got the right type of subquery back
			switch tt.subExpr.(type) {
			case *plan.ScalarSubquery:
				assert.IsType(t, &plan.ScalarSubquery{}, fromProto)
			case *plan.InPredicateSubquery:
				assert.IsType(t, &plan.InPredicateSubquery{}, fromProto)
			case *plan.SetPredicateSubquery:
				assert.IsType(t, &plan.SetPredicateSubquery{}, fromProto)
			case *plan.SetComparisonSubquery:
				assert.IsType(t, &plan.SetComparisonSubquery{}, fromProto)
			}

			// Verify protobuf roundtrip
			roundtripProto := wire.ExprToProto(fromProto)
			assert.True(t, pb.Equal(protoExpr, roundtripProto), "protobuf roundtrip failed")

			// Verify basic properties
			assert.Equal(t, tt.subExpr.IsScalar(), fromProto.IsScalar())
			assert.True(t, tt.subExpr.GetType().Equals(fromProto.GetType()))

			// Note: We don't test Equals() here because the current implementation
			// of isRelEqual() only does pointer equality, so relations created from
			// protobuf will never be equal to the original relations, even if they
			// have identical content. This is a known limitation noted in the TODO
			// comment in plan/subquery.go
		})
	}
}
