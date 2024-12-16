// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v3/types"
	"github.com/substrait-io/substrait-go/v3/types/integer_parameters"
)

func TestParameterizedStructType(t *testing.T) {
	decimalType := &types.ParameterizedDecimalType{
		Precision:   integer_parameters.NewVariableIntParam("P"),
		Scale:       integer_parameters.NewVariableIntParam("S"),
		Nullability: types.NullabilityRequired,
	}
	int8Type := &types.Int8Type{Nullability: types.NullabilityNullable}
	listParameterizedType := &types.ParameterizedListType{Type: decimalType, Nullability: types.NullabilityNullable}
	dec30PS5 := &types.DecimalType{Precision: 30, Scale: 5, Nullability: types.NullabilityRequired}
	dec30PS9 := &types.DecimalType{Precision: 30, Scale: 9, Nullability: types.NullabilityRequired}
	listType := &types.ListType{Type: dec30PS9, Nullability: types.NullabilityRequired}
	for _, td := range []struct {
		name                           string
		params                         []types.FuncDefArgType
		args                           []interface{}
		expectedNullableString         string
		expectedNullableRequiredString string
		expectedHasParameterizedParam  bool
		expectedParameterizedParams    []interface{}
		expectedReturnType             types.Type
	}{
		{"all parameterized param", []types.FuncDefArgType{decimalType, listParameterizedType}, []any{dec30PS5, listType}, "struct?<decimal<P,S>, list?<decimal<P,S>>>", "struct<decimal<P,S>, list?<decimal<P,S>>>", true, []interface{}{decimalType, listParameterizedType}, nil},
		{"mix parameterized concrete param", []types.FuncDefArgType{decimalType, int8Type, listParameterizedType}, []any{dec30PS9, int8Type, listType}, "struct?<decimal<P,S>, i8?, list?<decimal<P,S>>>", "struct<decimal<P,S>, i8?, list?<decimal<P,S>>>", true, []interface{}{decimalType, listParameterizedType}, nil},
		{"all concrete param", []types.FuncDefArgType{int8Type, int8Type, int8Type}, []any{int8Type, int8Type, int8Type}, "struct?<i8?, i8?, i8?>", "struct<i8?, i8?, i8?>", false, nil, &types.StructType{Nullability: types.NullabilityRequired, Types: []types.Type{int8Type, int8Type, int8Type}}},
	} {
		t.Run(td.name, func(t *testing.T) {
			pd := &types.ParameterizedStructType{Types: td.params}
			require.Equal(t, td.expectedNullableString, pd.SetNullability(types.NullabilityNullable).String())
			require.Equal(t, td.expectedNullableRequiredString, pd.SetNullability(types.NullabilityRequired).String())
			require.Equal(t, td.expectedHasParameterizedParam, pd.HasParameterizedParam())
			require.Equal(t, td.expectedParameterizedParams, pd.GetParameterizedParams())
			assert.Equal(t, "struct", pd.ShortString())
			retType, err := pd.ReturnType(nil, nil)
			if td.expectedReturnType == nil {
				assert.Error(t, err)
				require.True(t, pd.HasParameterizedParam())
			} else {
				require.Nil(t, err)
				require.Equal(t, td.expectedReturnType, retType)
				got, err := pd.WithParameters(td.args)
				require.Nil(t, err)
				require.Equal(t, td.expectedReturnType, got)
			}
		})
	}
}
