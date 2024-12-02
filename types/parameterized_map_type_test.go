// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/types"
	"github.com/substrait-io/substrait-go/types/integer_parameters"
)

func TestParameterizedMapType(t *testing.T) {
	decimalType := &types.ParameterizedDecimalType{
		Precision:   integer_parameters.NewVariableIntParam("P"),
		Scale:       integer_parameters.NewVariableIntParam("S"),
		Nullability: types.NullabilityRequired,
	}
	int8Type := &types.Int8Type{Nullability: types.NullabilityNullable}
	listType := &types.ParameterizedListType{Type: decimalType, Nullability: types.NullabilityNullable}
	for _, td := range []struct {
		name                           string
		Key                            types.FuncDefArgType
		Value                          types.FuncDefArgType
		expectedNullableString         string
		expectedNullableRequiredString string
		expectedHasParameterizedParam  bool
		expectedParameterizedParams    []interface{}
		expectedReturnType             types.Type
	}{
		{"parameterized kv", decimalType, listType, "map?<decimal<P,S>, list?<decimal<P,S>>>", "map<decimal<P,S>, list?<decimal<P,S>>>", true, []interface{}{decimalType, listType}, nil},
		{"concrete key", int8Type, listType, "map?<i8?, list?<decimal<P,S>>>", "map<i8?, list?<decimal<P,S>>>", true, []interface{}{listType}, nil},
		{"concrete value", decimalType, int8Type, "map?<decimal<P,S>, i8?>", "map<decimal<P,S>, i8?>", true, []interface{}{decimalType}, nil},
		{"no parameterized param", int8Type, int8Type, "map?<i8?, i8?>", "map<i8?, i8?>", false, nil, &types.MapType{Nullability: types.NullabilityRequired, Key: int8Type, Value: int8Type}},
	} {
		t.Run(td.name, func(t *testing.T) {
			pd := &types.ParameterizedMapType{Key: td.Key, Value: td.Value}
			assert.Equal(t, types.NullabilityUnspecified, pd.GetNullability())
			require.Equal(t, td.expectedNullableString, pd.SetNullability(types.NullabilityNullable).String())
			assert.Equal(t, types.NullabilityNullable, pd.GetNullability())
			require.Equal(t, td.expectedNullableRequiredString, pd.SetNullability(types.NullabilityRequired).String())
			assert.Equal(t, types.NullabilityRequired, pd.GetNullability())
			require.Equal(t, td.expectedHasParameterizedParam, pd.HasParameterizedParam())
			require.Equal(t, td.expectedParameterizedParams, pd.GetParameterizedParams())
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
