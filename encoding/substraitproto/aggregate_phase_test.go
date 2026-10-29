// SPDX-License-Identifier: Apache-2.0

package substraitproto_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/encoding/substraitproto"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/types"
	"github.com/substrait-io/substrait-go/v9/types/parser"
)

func phaseType(t *testing.T, s string) types.Type {
	t.Helper()
	parsed, err := parser.ParseType(s)
	require.NoError(t, err)
	typ, err := parsed.ReturnType(nil, nil)
	require.NoError(t, err)
	return typ
}

func TestAggregatePhaseOutputTypeToProto(t *testing.T) {
	tests := []struct {
		name, function, initial, intermediate, result string
	}{
		{"avg", "functions_arithmetic/avg:i32", "i32?", "struct<i64,i64>", "i32?"},
		{"avg_i64", "functions_arithmetic/avg:i64", "i64", "struct<i64,i64>", "i64?"},
		{"sum", "functions_arithmetic/sum:i32", "i32?", "i64?", "i64?"},
		{"count_values", "functions_aggregate_generic/count:any", "string?", "i64", "i64"},
		{"count_rows", "functions_aggregate_generic/count:", "", "i64", "i64"},
		{"any_value", "functions_aggregate_generic/any_value:any", "string?", "string?", "string?"},
		{"decimal_avg", "functions_arithmetic_decimal/avg:dec", "decimal?<12,2>", "struct<decimal<38,2>,i64>", "decimal<38,2>"},
		{"decimal_sum", "functions_arithmetic_decimal/sum:dec", "decimal?<12,2>", "decimal?<38,2>", "decimal?<38,2>"},
		{"decimal_min", "functions_arithmetic_decimal/min:dec", "decimal?<12,2>", "decimal?<12,2>", "decimal?<12,2>"},
	}
	phases := []types.AggregationPhase{
		types.AggregationPhaseInitialToIntermediate,
		types.AggregationPhaseIntermediateToIntermediate,
		types.AggregationPhaseInitialToResult,
		types.AggregationPhaseIntermediateToResult,
		types.AggregationPhaseUnspecified,
	}
	for _, tt := range tests {
		for _, phase := range phases {
			for _, window := range []bool{false, true} {
				kind := "aggregate"
				if window {
					kind = "window"
				}
				t.Run(tt.name+"/"+phase.String()+"/"+kind, func(t *testing.T) {
					reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
					urn, name, _ := strings.Cut(tt.function, "/")
					id := extensions.FunctionID{URN: extensions.SubstraitDefaultURNPrefix + urn, Name: name}
					input := tt.intermediate
					if phase == types.AggregationPhaseInitialToIntermediate || phase == types.AggregationPhaseInitialToResult {
						input = tt.initial
					}
					var args []types.FuncArg
					if input != "" {
						args = []types.FuncArg{&expr.DynamicParameter{OutputType: phaseType(t, input)}}
					}
					expected := tt.result
					if phase == types.AggregationPhaseInitialToIntermediate || phase == types.AggregationPhaseIntermediateToIntermediate {
						expected = tt.intermediate
					}
					if window {
						fn, err := expr.NewWindowFunc(reg, id, nil, types.AggregationInvocationAll, phase, args...)
						require.NoError(t, err)
						assert.Equal(t, substraitproto.TypeToProto(phaseType(t, expected)), substraitproto.ExprToProto(fn).GetWindowFunction().OutputType)
					} else {
						fn, err := expr.NewAggregateFunc(reg, id, nil, types.AggregationInvocationAll, phase, nil, args...)
						require.NoError(t, err)
						assert.Equal(t, substraitproto.TypeToProto(phaseType(t, expected)), substraitproto.AggregateFunctionToProto(fn).OutputType)
					}
				})
			}
		}
	}
}
