// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/types"
	"github.com/substrait-io/substrait-go/types/integer_parameters"
)

func TestParameterizedStructType(t *testing.T) {
	decimalType := &types.ParameterizedDecimalType{
		Precision:   integer_parameters.NewVariableIntParam("P"),
		Scale:       integer_parameters.NewVariableIntParam("S"),
		Nullability: types.NullabilityRequired,
	}
	int8Type := &types.Int8Type{Nullability: types.NullabilityNullable}
	listType := &types.ParameterizedListType{Type: decimalType, Nullability: types.NullabilityNullable}
	for _, td := range []struct {
		name                           string
		params                         []types.FuncDefArgType
		expectedNullableString         string
		expectedNullableRequiredString string
		expectedHasParameterizedParam  bool
		expectedParameterizedParams    []interface{}
	}{
		{"all parameterized param", []types.FuncDefArgType{decimalType, listType}, "struct?<decimal<P,S>, list?<decimal<P,S>>>", "struct<decimal<P,S>, list?<decimal<P,S>>>", true, []interface{}{decimalType, listType}},
		{"mix parameterized concrete param", []types.FuncDefArgType{decimalType, int8Type, listType}, "struct?<decimal<P,S>, i8?, list?<decimal<P,S>>>", "struct<decimal<P,S>, i8?, list?<decimal<P,S>>>", true, []interface{}{decimalType, listType}},
		{"all concrete param", []types.FuncDefArgType{int8Type, int8Type, int8Type}, "struct?<i8?, i8?, i8?>", "struct<i8?, i8?, i8?>", false, nil},
	} {
		t.Run(td.name, func(t *testing.T) {
			pd := &types.ParameterizedStructType{Types: td.params}
			require.Equal(t, td.expectedNullableString, pd.SetNullability(types.NullabilityNullable).String())
			require.Equal(t, td.expectedNullableRequiredString, pd.SetNullability(types.NullabilityRequired).String())
			require.Equal(t, td.expectedHasParameterizedParam, pd.HasParameterizedParam())
			require.Equal(t, td.expectedParameterizedParams, pd.GetParameterizedParams())
		})
	}
}
