// SPDX-License-Identifier: Apache-2.0

package extensions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/extensions"
	"github.com/substrait-io/substrait-go/types"
	"github.com/substrait-io/substrait-go/types/integer_parameters"
	"github.com/substrait-io/substrait-go/types/parser"
)

func TestEvaluateTypeExpression(t *testing.T) {
	var (
		p, _          = parser.New()
		i64Null, _    = p.ParseString("i64?")
		i64NonNull, _ = p.ParseString("i64")
		strNull, _    = p.ParseString("string?")
	)

	tests := []struct {
		name     string
		nulls    extensions.NullabilityHandling
		ret      parser.TypeExpression
		extArgs  extensions.ArgumentList
		args     []types.Type
		expected types.Type
		err      string
	}{
		{"defaults", "", *i64NonNull, extensions.ArgumentList{
			extensions.ValueArg{Value: i64Null}},
			[]types.Type{&types.Int64Type{Nullability: types.NullabilityNullable}},
			&types.Int64Type{Nullability: types.NullabilityNullable}, ""},
		{"arg mismatch", "", *strNull, extensions.ArgumentList{extensions.ValueArg{Value: strNull}},
			[]types.Type{}, nil, "invalid expression: mismatch in number of arguments provided. got 0, expected 1"},
		{"missing enum arg", "", *i64Null, extensions.ArgumentList{
			extensions.ValueArg{Value: i64NonNull}, extensions.EnumArg{Name: "foo"}},
			[]types.Type{&types.Int64Type{}, &types.Int64Type{}}, nil, "invalid type: arg #1 (foo) should be an enum"},
		{"discrete null handling", extensions.DiscreteNullability, *strNull, extensions.ArgumentList{
			extensions.ValueArg{Value: strNull}},
			[]types.Type{&types.StringType{Nullability: types.NullabilityRequired}},
			nil, "invalid type: discrete nullability did not match for arg #0"},
		{"mirror", extensions.MirrorNullability, *strNull, extensions.ArgumentList{
			extensions.ValueArg{Value: i64NonNull}, extensions.ValueArg{Value: i64Null}},
			[]types.Type{
				&types.Int64Type{Nullability: types.NullabilityRequired},
				&types.Int64Type{Nullability: types.NullabilityRequired}},
			&types.StringType{Nullability: types.NullabilityRequired}, ""},
		{"declared output", extensions.DeclaredOutputNullability, *strNull, extensions.ArgumentList{
			extensions.ValueArg{Value: strNull}},
			[]types.Type{&types.StringType{Nullability: types.NullabilityRequired}},
			&types.StringType{Nullability: types.NullabilityNullable}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := extensions.EvaluateTypeExpression(tt.nulls, tt.ret, tt.extArgs, nil, tt.args)
			if tt.err == "" {
				assert.NoError(t, err)
				assert.Truef(t, tt.expected.Equals(result), "expected: %s\ngot: %s", tt.expected, result)
			} else {
				assert.EqualError(t, err, tt.err)
			}
		})
	}
}

func TestHasSyncParams(t *testing.T) {

	apt_P := integer_parameters.NewVariableIntParam("P")
	apt_Q := integer_parameters.NewVariableIntParam("Q")
	cpt_38 := integer_parameters.NewConcreteIntParam(38)

	fct_P := &types.ParameterizedFixedCharType{IntegerOption: apt_P}
	fct_Q := &types.ParameterizedFixedCharType{IntegerOption: apt_Q}
	decimal_PQ := &types.ParameterizedDecimalType{Precision: apt_P, Scale: apt_Q}
	decimal_38_Q := &types.ParameterizedDecimalType{Precision: cpt_38, Scale: apt_Q}
	list_decimal_38_Q := &types.ParameterizedListType{Type: decimal_38_Q}
	map_fctQ_decimal38Q := &types.ParameterizedMapType{Key: fct_Q, Value: decimal_38_Q}
	struct_fctQ_ListDecimal38Q := &types.ParameterizedStructType{Types: []types.FuncDefArgType{fct_Q, list_decimal_38_Q}}
	for _, td := range []struct {
		name                  string
		params                []types.FuncDefArgType
		expectedHasSyncParams bool
	}{
		{"No Abstract Type", []types.FuncDefArgType{&types.Int64Type{}}, false},
		{"No Sync Param P, Q", []types.FuncDefArgType{fct_P, fct_Q}, false},
		{"Sync Params P, P", []types.FuncDefArgType{fct_P, fct_P}, true},
		{"Sync Params P, <P, Q>", []types.FuncDefArgType{fct_P, decimal_PQ}, true},
		{"No Sync Params P, <38, Q>", []types.FuncDefArgType{fct_P, decimal_38_Q}, false},
		{"Sync Params P, List<Decimal<P, Q>>", []types.FuncDefArgType{fct_P, list_decimal_38_Q}, false},
		{"No Sync Params fct<P>, Map<fct<Q>, decimal<38,Q>>", []types.FuncDefArgType{fct_P, map_fctQ_decimal38Q}, false},
		{"Sync Params fct<Q>, Map<fct<Q>, decimal<38,Q>>", []types.FuncDefArgType{fct_Q, map_fctQ_decimal38Q}, true},
		{"No Sync Params fct<P>, struct<fct<Q>, list<38,Q>>", []types.FuncDefArgType{fct_P, struct_fctQ_ListDecimal38Q}, false},
		{"Sync Params fct<Q>, struct<fct<Q>, list<38,Q>>", []types.FuncDefArgType{fct_Q, struct_fctQ_ListDecimal38Q}, true},
	} {
		t.Run(td.name, func(t *testing.T) {
			if td.expectedHasSyncParams {
				require.True(t, extensions.HasSyncParams(td.params))
			} else {
				require.False(t, extensions.HasSyncParams(td.params))
			}
		})
	}
}
