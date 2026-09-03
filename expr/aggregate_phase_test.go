// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	substraitgo "github.com/substrait-io/substrait-go/v9"
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

func TestAggregateFunctionPhases(t *testing.T) {
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
		types.AggPhaseInitialToIntermediate,
		types.AggPhaseIntermediateToIntermediate,
		types.AggPhaseInitialToResult,
		types.AggPhaseIntermediateToResult,
		types.AggPhaseUnspecified,
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
					if phase == types.AggPhaseInitialToIntermediate || phase == types.AggPhaseInitialToResult {
						input = tt.initial
					}
					var args []types.FuncArg
					if input != "" {
						args = []types.FuncArg{&expr.DynamicParameter{OutputType: phaseType(t, input)}}
					}
					expected := tt.result
					if phase == types.AggPhaseInitialToIntermediate || phase == types.AggPhaseIntermediateToIntermediate {
						expected = tt.intermediate
					}
					if window {
						fn, err := expr.NewWindowFunc(reg, id, nil, types.AggInvocationAll, phase, args...)
						require.NoError(t, err)
						assert.Equal(t, phaseType(t, expected), fn.GetType())
						assert.Equal(t, types.TypeToProto(phaseType(t, expected)), fn.ToProto().GetWindowFunction().OutputType)
						assert.Equal(t, id, fn.ID())
						assert.Equal(t, phase, fn.Phase())
					} else {
						fn, err := expr.NewAggregateFunc(reg, id, nil, types.AggInvocationAll, phase, nil, args...)
						require.NoError(t, err)
						assert.Equal(t, phaseType(t, expected), fn.GetType())
						assert.Equal(t, types.TypeToProto(phaseType(t, expected)), fn.ToProto().OutputType)
						assert.Equal(t, id, fn.ID())
						assert.Equal(t, phase, fn.Phase())
					}
				})
			}
		}
	}
}

func TestAggregatePhaseRejectsInvalidArguments(t *testing.T) {
	tests := []struct {
		name, function string
		phase          types.AggregationPhase
		args           []string
		target         error
	}{
		{"original_input", "avg:i32", types.AggPhaseIntermediateToResult, []string{"i32"}, substraitgo.ErrInvalidType},
		{"missing_state", "avg:i32", types.AggPhaseIntermediateToResult, nil, substraitgo.ErrInvalidExpr},
		{"multiple_states", "avg:i32", types.AggPhaseIntermediateToResult, []string{"struct<i64,i64>", "struct<i64,i64>"}, substraitgo.ErrInvalidExpr},
		{"ambiguous_state", "avg", types.AggPhaseIntermediateToResult, []string{"struct<i64,i64>"}, substraitgo.ErrNotFound},
		{"state_cannot_identify_sum_overload", "sum", types.AggPhaseIntermediateToResult, []string{"i64?"}, substraitgo.ErrNotFound},
		{"nondecomposable", "mode:i32", types.AggPhaseInitialToIntermediate, []string{"i32"}, substraitgo.ErrInvalidExpr},
		{"nondecomposable_unspecified", "mode:i32", types.AggPhaseUnspecified, []string{"i32"}, substraitgo.ErrInvalidExpr},
		{"invalid_phase", "avg:i32", types.AggregationPhase(99), []string{"i32"}, substraitgo.ErrInvalidExpr},
	}
	for _, tt := range tests {
		for _, window := range []bool{false, true} {
			kind := "aggregate"
			if window {
				kind = "window"
			}
			t.Run(tt.name+"/"+kind, func(t *testing.T) {
				reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
				id := extensions.FunctionID{URN: extensions.SubstraitDefaultURNPrefix + "functions_arithmetic", Name: tt.function}
				var args []types.FuncArg
				for _, arg := range tt.args {
					args = append(args, &expr.DynamicParameter{OutputType: phaseType(t, arg)})
				}
				var err error
				if window {
					_, err = expr.NewWindowFunc(reg, id, nil, types.AggInvocationAll, tt.phase, args...)
				} else {
					_, err = expr.NewAggregateFunc(reg, id, nil, types.AggInvocationAll, tt.phase, nil, args...)
				}
				require.ErrorIs(t, err, tt.target)
			})
		}
	}
}

func TestAggregatePhaseTypeParameters(t *testing.T) {
	const declarations = `
urn: extension:test:aggregate_phases
aggregate_functions:
  - name: nested
    impls:
      - args:
          - value: any1
        nullability: DECLARED_OUTPUT
        decomposable: MANY
        intermediate: struct<any1,i64>
        return: any1
  - name: mirror
    impls:
      - args:
          - value: decimal<P,S>
        nullability: MIRROR
        decomposable: MANY
        intermediate: struct<decimal<P,S>,i64>
        return: decimal<P,S>
  - name: discrete
    impls:
      - args:
          - value: decimal?<P,S>
        nullability: DISCRETE
        decomposable: MANY
        intermediate: struct<decimal<P,S>,i64>
        return: decimal?<P,S>
  - name: lost_parameter
    impls:
      - args:
          - value: decimal<P,S>
        nullability: DECLARED_OUTPUT
        decomposable: MANY
        intermediate: decimal<38,S>
        return: decimal<P,S>
  - name: derived_parameter
    impls:
      - args:
          - value: decimal<P,S>
        nullability: DECLARED_OUTPUT
        decomposable: MANY
        intermediate: |
          P = P + 1
          decimal<P,S>
        return: decimal<P,S>
`
	var col extensions.Collection
	require.NoError(t, col.Load(strings.NewReader(declarations)))
	tests := []struct {
		name, input, output string
		phase               types.AggregationPhase
	}{
		{"nested:any", "string", "struct<string,i64>", types.AggPhaseInitialToIntermediate},
		{"nested:any", "struct<string,i64>", "string", types.AggPhaseIntermediateToResult},
		{"nested:any", "struct<string,i64>", "struct<string,i64>", types.AggPhaseIntermediateToIntermediate},
		{"mirror:dec", "struct?<decimal<12,2>,i64>", "decimal?<12,2>", types.AggPhaseIntermediateToResult},
		{"mirror:dec", "struct<decimal?<12,2>,i64>", "decimal<12,2>", types.AggPhaseIntermediateToResult},
		{"discrete:dec", "struct<decimal<12,2>,i64>", "decimal?<12,2>", types.AggPhaseIntermediateToResult},
		{"discrete:dec", "struct?<decimal<12,2>,i64>", "", types.AggPhaseIntermediateToResult},
		{"discrete:dec", "struct<decimal?<12,2>,i64>", "", types.AggPhaseIntermediateToResult},
		{"lost_parameter:dec", "decimal<12,2>", "decimal<38,2>", types.AggPhaseInitialToIntermediate},
		{"lost_parameter:dec", "decimal<38,2>", "", types.AggPhaseIntermediateToResult},
		{"derived_parameter:dec", "decimal<12,2>", "decimal<13,2>", types.AggPhaseInitialToIntermediate},
		{"derived_parameter:dec", "decimal<13,2>", "", types.AggPhaseIntermediateToResult},
	}
	for _, tt := range tests {
		for _, window := range []bool{false, true} {
			kind := "aggregate"
			if window {
				kind = "window"
			}
			t.Run(tt.name+"/"+tt.phase.String()+"/"+kind, func(t *testing.T) {
				reg := expr.NewEmptyExtensionRegistry(&col)
				id := extensions.FunctionID{URN: "extension:test:aggregate_phases", Name: tt.name}
				arg := &expr.DynamicParameter{OutputType: phaseType(t, tt.input)}
				var output types.Type
				var err error
				if window {
					var fn *expr.WindowFunction
					fn, err = expr.NewWindowFunc(reg, id, nil, types.AggInvocationAll, tt.phase, arg)
					if err == nil {
						output = fn.GetType()
					}
				} else {
					var fn *expr.AggregateFunction
					fn, err = expr.NewAggregateFunc(reg, id, nil, types.AggInvocationAll, tt.phase, nil, arg)
					if err == nil {
						output = fn.GetType()
					}
				}
				if tt.output == "" {
					require.Error(t, err)
					return
				}
				require.NoError(t, err)
				assert.Equal(t, phaseType(t, tt.output), output)
			})
		}
	}
}

func TestAggregatePhaseUniqueAlias(t *testing.T) {
	const declaration = `
urn: extension:test:window_phase
window_functions:
  - name: partial
    impls:
      - args:
          - value: i32
        nullability: DECLARED_OUTPUT
        decomposable: MANY
        intermediate: i64
        return: i32
`
	var col extensions.Collection
	require.NoError(t, col.Load(strings.NewReader(declaration)))
	reg := expr.NewEmptyExtensionRegistry(&col)
	id := extensions.FunctionID{URN: "extension:test:window_phase", Name: "partial"}
	arg := &expr.DynamicParameter{OutputType: phaseType(t, "i64")}
	window, err := expr.NewWindowFunc(reg, id, nil, types.AggInvocationAll, types.AggPhaseIntermediateToResult, arg)
	require.NoError(t, err)
	assert.Equal(t, "partial:i32", window.CompoundName())
	assert.Equal(t, phaseType(t, "i32"), window.GetType())
}

func TestAggregateInitialPhaseInfersSignature(t *testing.T) {
	reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
	id := extensions.FunctionID{URN: extensions.SubstraitDefaultURNPrefix + "functions_arithmetic", Name: "avg"}
	arg := expr.NewPrimitiveLiteral(int32(7), false)
	agg, err := expr.NewAggregateFunc(reg, id, nil, types.AggInvocationAll, types.AggPhaseInitialToIntermediate, nil, arg)
	require.NoError(t, err)
	assert.Equal(t, "avg:i32", agg.CompoundName())
	assert.Equal(t, phaseType(t, "struct<i64,i64>"), agg.GetType())
	window, err := expr.NewWindowFunc(reg, id, nil, types.AggInvocationAll, types.AggPhaseInitialToIntermediate, arg)
	require.NoError(t, err)
	assert.Equal(t, "avg:i32", window.CompoundName())
	assert.Equal(t, phaseType(t, "struct<i64,i64>"), window.GetType())
}
