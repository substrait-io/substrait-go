// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v4/expr"
	"github.com/substrait-io/substrait-go/v4/extensions"
	"github.com/substrait-io/substrait-go/v4/plan"
	"github.com/substrait-io/substrait-go/v4/types"
	"github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/types/known/anypb"
)

func TestExprBuilder(t *testing.T) {
	b := expr.ExprBuilder{
		Reg:        expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError()),
		BaseSchema: types.NewRecordTypeFromStruct(boringSchema.Struct),
	}
	precomputedLiteral, _ := expr.NewLiteral(int32(3), false)
	precomputedExpression, _ := b.ScalarFunc(addID).Args(
		b.Wrap(expr.NewLiteral(int32(3), false)),
		b.Wrap(expr.NewLiteral(int32(3), false))).BuildExpr()

	tests := []struct {
		name     string
		expected string
		ex       expr.Builder
		err      string
	}{
		{"literal", "i8?(5)", b.Wrap(expr.NewLiteral(int8(5), true)), ""},
		{"preciseTimeLiteral", "precision_time?<4>(00:00:12.3456)", b.Wrap(expr.NewPrecisionTimeLiteral(123456, types.PrecisionEMinus4Seconds, types.NullabilityNullable), nil), ""},
		{"preciseTimeStampliteral", "precision_timestamp?<3>(1970-01-01 00:02:03.456)", b.Wrap(expr.NewPrecisionTimestampLiteral(123456, types.PrecisionMilliSeconds, types.NullabilityNullable), nil), ""},
		{"preciseTimeStampTzliteral", "precision_timestamp_tz?<6>(1970-01-01T00:00:00.123456Z)", b.Wrap(expr.NewPrecisionTimestampTzLiteral(123456, types.PrecisionMicroSeconds, types.NullabilityNullable), nil), ""},
		{"simple add", "add(.field(1) => i8, i8(5)) => i8?",
			b.ScalarFunc(addID).Args(
				b.RootRef(expr.NewStructFieldRef(1)),
				b.Literal(expr.NewPrimitiveLiteral(int8(5), false)),
			), ""},
		{"expect args", "",
			b.ScalarFunc(indexInID),
			"invalid expression: mismatch in number of arguments provided. got 0, expected 2"},
		{"with opt", "index_in(i32(5), list<i32>([]), {nan_equality: [NAN_IS_NAN]}) => i64?",
			b.ScalarFunc(indexInID, &types.FunctionOption{
				Name:       "nan_equality",
				Preference: []string{"NAN_IS_NAN"}}).Args(
				b.Wrap(expr.NewLiteral(int32(5), false)),
				b.Literal(expr.NewEmptyListLiteral(&types.Int32Type{}, false))), ""},
		{"with cast", "subtract(.field(3) => i32, cast(.field(6) => fp32 AS i32, fail: FAILURE_BEHAVIOR_THROW_EXCEPTION)) => i32?",
			b.ScalarFunc(subID).Args(
				b.RootRef(expr.NewStructFieldRef(3)),
				b.Cast(b.RootRef(expr.NewStructFieldRef(6)), &types.Int32Type{}).
					FailBehavior(types.BehaviorThrowException),
			), ""},
		{"expression with lit", "subtract(.field(3) => i32, i32(3)) => i32",
			b.ScalarFunc(subID).Args(b.RootRef(expr.NewStructFieldRef(3)),
				b.Expression(precomputedLiteral)), ""},
		{"expression with expr", "subtract(.field(3) => i32, add(i32(3), i32(3)) => i32) => i32",
			b.ScalarFunc(subID).Args(b.RootRef(expr.NewStructFieldRef(3)),
				b.Expression(precomputedExpression)), ""},
		{"wrap expression", "subtract(.field(3) => i32, i32(3)) => i32",
			b.ScalarFunc(subID).Args(b.RootRef(expr.NewStructFieldRef(3)),
				b.Wrap(expr.NewLiteral(int32(3), false))), ""},
		{"window func", "",
			b.WindowFunc(rankID), "invalid expression: non-decomposable window or agg function '{urn:substrait:functions_arithmetic rank}' must use InitialToResult phase"},
		{"window func", "rank(; phase: AGGREGATION_PHASE_INITIAL_TO_RESULT, invocation: AGGREGATION_INVOCATION_UNSPECIFIED) => i64?",
			b.WindowFunc(rankID).Phase(types.AggPhaseInitialToResult), ""},
		{"window func",
			"first_value(i32(3); partitions: [.field(0) => boolean]; phase: AGGREGATION_PHASE_INITIAL_TO_RESULT, invocation: AGGREGATION_INVOCATION_UNSPECIFIED) => i32",
			b.WindowFunc(firstValueID).Args(
				b.Wrap(expr.NewLiteral(int32(3), false))).
				Phase(types.AggPhaseInitialToResult).
				Partitions(b.RootRef(expr.NewStructFieldRef(0))), ""},
		{"agg as window", "sum(i32?(42); partitions: [.field(0) => boolean]; phase: AGGREGATION_PHASE_INITIAL_TO_RESULT, invocation: AGGREGATION_INVOCATION_UNSPECIFIED) => i64?",
			b.WindowFunc(sumID).Args(
				b.Wrap(expr.NewLiteral(int32(42), true))).
				Phase(types.AggPhaseInitialToResult).
				Partitions(b.RootRef(expr.NewStructFieldRef(0))), ""},
		{"nested funcs", "add(extract(YEAR, date(2000-01-01)) => i64, rank(; phase: AGGREGATION_PHASE_INITIAL_TO_RESULT, invocation: AGGREGATION_INVOCATION_ALL) => i64?) => i64?",
			b.ScalarFunc(addID).Args(
				b.ScalarFunc(extractID).Args(b.Enum("YEAR"),
					b.Wrap(expr.NewLiteral(types.Date(10957), false))),
				b.WindowFunc(rankID).Phase(types.AggPhaseInitialToResult).Invocation(types.AggInvocationAll)), ""},
		{"nested propagate error", "",
			b.ScalarFunc(addID).Args(
				b.RootRef(expr.NewListElemRef(0)),
				b.Literal(expr.NewPrimitiveLiteral(int32(5), false))), "error resolving ref type: invalid type"},
		{"window func args", "ntile(i32(5); sort: [{expr: .field(1) => i8, SORT_DIRECTION_ASC_NULLS_FIRST}]; phase: AGGREGATION_PHASE_INITIAL_TO_RESULT, invocation: AGGREGATION_INVOCATION_UNSPECIFIED) => i32?",
			b.WindowFunc(ntileID).Args(b.Wrap(expr.NewLiteral(int32(5), false))).
				Phase(types.AggPhaseInitialToResult).
				Sort(expr.SortField{
					Expr: expr.MustExpr(b.RootRef(expr.NewStructFieldRef(1)).Build()),
					Kind: types.SortAscNullsFirst}), ""},
		{"window func arg error", "",
			b.WindowFunc(ntileID).Args(b.ScalarFunc(extensions.ID{})),
			"not found: could not find matching function for id: { :}"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, err := tt.ex.BuildExpr()
			if tt.err == "" {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, e.String())
				// Also test that converting to proto does not panic.
				e.ToProto()
			} else {
				assert.EqualError(t, err, tt.err)
			}
		})
	}
}

func TestCustomTypesInFunctionOutput(t *testing.T) {
	custom := `%YAML 1.2
---
types:
  - name: custom_type1
  - name: custom_type2
  - name: custom_type3
  - name: custom_type4

scalar_functions:
  - name: custom_function
    description: "custom function that takes in and returns custom types"
    impls:
      - args:
          - name: arg1
            value: u!custom_type2
        return: u!custom_type1

aggregate_functions:
  - name: "custom_aggr"
    description: "custom aggregator that takes in and returns custom types"
    impls:
      - args:
          - name: arg1
            value: u!custom_type2
        return: u!custom_type3

window_functions:
  - name: "custom_window"
    description: "custom window function that takes in and returns custom types"
    impls:
      - args:
          - name: arg1
            value: u!custom_type2
        return: u!custom_type1
`

	customReader := strings.NewReader(custom)
	collection := extensions.Collection{}
	err := collection.Load("custom", customReader)
	require.NoError(t, err)

	planBuilder := plan.NewBuilder(&collection)

	customType1 := planBuilder.UserDefinedType("custom", "custom_type1")
	customType2 := planBuilder.UserDefinedType("custom", "custom_type2")
	customType3 := planBuilder.UserDefinedType("custom", "custom_type3")

	anyVal, err := anypb.New(expr.NewPrimitiveLiteral("foo", false).ToProto())
	require.NoError(t, err)

	customLiteral := planBuilder.GetExprBuilder().Literal(&expr.ProtoLiteral{
		Type:  &customType2,
		Value: anyVal,
	})

	// check scalar function
	scalar, err := planBuilder.GetExprBuilder().ScalarFunc(extensions.ID{
		URI:  "custom",
		Name: "custom_function",
	}).Args(
		customLiteral,
	).BuildExpr()
	require.NoError(t, err)
	scalarProto := scalar.ToProto()

	fnCall := scalarProto.GetScalarFunction()
	require.Len(t, fnCall.Arguments, 1)
	require.Equal(t, customType2.TypeReference, fnCall.Arguments[0].GetValue().GetLiteral().GetUserDefined().TypeReference)
	require.Equal(t, customType1.TypeReference, fnCall.OutputType.GetUserDefined().TypeReference)

	// check aggregate function
	aggr, err := planBuilder.GetExprBuilder().AggFunc(extensions.ID{
		URI:  "custom",
		Name: "custom_aggr",
	}).Args(
		customLiteral,
	).Build()
	require.NoError(t, err)
	aggrProto := aggr.ToProto()

	require.Len(t, aggrProto.Arguments, 1)
	require.Equal(t, customType2.TypeReference, aggrProto.Arguments[0].GetValue().GetLiteral().GetUserDefined().TypeReference)
	require.Equal(t, customType3.TypeReference, aggrProto.OutputType.GetUserDefined().TypeReference)

	// check window function
	window, err := planBuilder.GetExprBuilder().WindowFunc(extensions.ID{
		URI:  "custom",
		Name: "custom_window",
	}).Args(
		customLiteral,
	).Phase(types.AggPhaseInitialToResult).Build()
	require.NoError(t, err)
	windowProto := window.ToProto()

	windowFnCall := windowProto.GetWindowFunction()
	require.Len(t, windowFnCall.Arguments, 1)
	require.Equal(t, customType2.TypeReference, windowFnCall.Arguments[0].GetValue().GetLiteral().GetUserDefined().TypeReference)
	require.Equal(t, customType1.TypeReference, windowFnCall.OutputType.GetUserDefined().TypeReference)

	// build a full plan and check that user defined types are registered in the extensions
	table, err := planBuilder.VirtualTable([]string{"col_a", "col_b"}, []expr.Literal{expr.NewPrimitiveLiteral(int64(2), false), expr.NewPrimitiveLiteral(int64(3), false)})
	require.NoError(t, err)

	aggregated, err := planBuilder.GetRelBuilder().AggregateRel(table, []plan.AggRelMeasure{planBuilder.Measure(aggr, window)}).Build()
	require.NoError(t, err)

	project, err := planBuilder.Project(aggregated, scalar)
	require.NoError(t, err)

	p, err := planBuilder.Plan(project, []string{"output1", "output2"})
	require.NoError(t, err)

	pp, err := p.ToProto()
	require.NoError(t, err)

	// custom_type1 is referenced as an argument and return type, so should be registered in the extensions
	// custom_type2 is referenced as an argument and return type, so should be registered in the extensions
	// custom_type3 is only referenced as a return type, but should still be registered in the extensions
	// custom_type4 is not referenced in the plan at all, so not be registerd in the extensions
	typeExtensionsFound := []string{}
	for _, ext := range pp.Extensions {
		typeExt := ext.GetExtensionType()
		if typeExt == nil {
			continue
		}
		typeExtensionsFound = append(typeExtensionsFound, typeExt.GetName())
	}
	require.Equal(t, 3, len(typeExtensionsFound))
	require.Contains(t, typeExtensionsFound, "custom_type1")
	require.Contains(t, typeExtensionsFound, "custom_type2")
	require.Contains(t, typeExtensionsFound, "custom_type3")
}

func TestBoundFromProto(t *testing.T) {
	for _, tc := range []struct {
		name        string
		proto       *substraitpb.Expression_WindowFunction_Bound
		expected    expr.Bound
		expectedStr string
	}{
		{
			name: "nil",
		},
		{
			name:  "nil kind",
			proto: &substraitpb.Expression_WindowFunction_Bound{},
		},
		{
			name: "unbounded",
			proto: &substraitpb.Expression_WindowFunction_Bound{
				Kind: &substraitpb.Expression_WindowFunction_Bound_Unbounded_{},
			},
			expected:    expr.Unbounded{},
			expectedStr: "UNBOUNDED",
		},
		{
			name: "current row",
			proto: &substraitpb.Expression_WindowFunction_Bound{
				Kind: &substraitpb.Expression_WindowFunction_Bound_CurrentRow_{},
			},
			expected:    expr.CurrentRow{},
			expectedStr: "CURRENT ROW",
		},
		{
			name: "preceding 42",
			proto: &substraitpb.Expression_WindowFunction_Bound{
				Kind: &substraitpb.Expression_WindowFunction_Bound_Preceding_{
					Preceding: &substraitpb.Expression_WindowFunction_Bound_Preceding{
						Offset: 42,
					},
				},
			},
			expected:    expr.PrecedingBound(42),
			expectedStr: "42 PRECEDING",
		},
		{
			name: "following 42",
			proto: &substraitpb.Expression_WindowFunction_Bound{
				Kind: &substraitpb.Expression_WindowFunction_Bound_Following_{
					Following: &substraitpb.Expression_WindowFunction_Bound_Following{
						Offset: 42,
					},
				},
			},
			expected:    expr.FollowingBound(42),
			expectedStr: "42 FOLLOWING",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			bound := expr.BoundFromProto(tc.proto)
			require.Equal(t, tc.expected, bound)
			if bound != nil {
				require.Equal(t, tc.expectedStr, bound.String())
			}
		})
	}
}
