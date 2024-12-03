// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
		expectedReturnType             types.Type
	}{
		{"all parameterized param", []types.FuncDefArgType{decimalType, listType}, "struct?<decimal<P,S>, list?<decimal<P,S>>>", "struct<decimal<P,S>, list?<decimal<P,S>>>", true, []interface{}{decimalType, listType}, nil},
		{"mix parameterized concrete param", []types.FuncDefArgType{decimalType, int8Type, listType}, "struct?<decimal<P,S>, i8?, list?<decimal<P,S>>>", "struct<decimal<P,S>, i8?, list?<decimal<P,S>>>", true, []interface{}{decimalType, listType}, nil},
		{"all concrete param", []types.FuncDefArgType{int8Type, int8Type, int8Type}, "struct?<i8?, i8?, i8?>", "struct<i8?, i8?, i8?>", false, nil, &types.StructType{Nullability: types.NullabilityRequired, Types: []types.Type{int8Type, int8Type, int8Type}}},
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
			}
		})
	}
}
