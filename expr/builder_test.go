// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v3/expr"
	"github.com/substrait-io/substrait-go/v3/extensions"
	"github.com/substrait-io/substrait-go/v3/types"
)

func TestExprBuilder(t *testing.T) {
	b := expr.ExprBuilder{
		Reg:        expr.NewEmptyExtensionRegistry(&extensions.DefaultCollection),
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
		{"preciseTimeStampliteral", "precisiontimestamp?<3>(123456)", b.Wrap(expr.NewPrecisionTimestampLiteral(123456, types.PrecisionMilliSeconds, types.NullabilityNullable), nil), ""},
		{"preciseTimeStampTzliteral", "precisiontimestamptz?<6>(123456)", b.Wrap(expr.NewPrecisionTimestampTzLiteral(123456, types.PrecisionMicroSeconds, types.NullabilityNullable), nil), ""},
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
			b.WindowFunc(rankID), "invalid expression: non-decomposable window or agg function '{https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml rank}' must use InitialToResult phase"},
		{"window func", "rank(; phase: AGGREGATION_PHASE_INITIAL_TO_RESULT, invocation: AGGREGATION_INVOCATION_UNSPECIFIED) => i64?",
			b.WindowFunc(rankID).Phase(types.AggPhaseInitialToResult), ""},
		{"nested funcs", "add(extract(YEAR, date(10957)) => i64, rank(; phase: AGGREGATION_PHASE_INITIAL_TO_RESULT, invocation: AGGREGATION_INVOCATION_ALL) => i64?) => i64?",
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
