// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"testing"

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
	}{
		{"parameterized kv", decimalType, listType, "map?<decimal<P,S>, list?<decimal<P,S>>>", "map<decimal<P,S>, list?<decimal<P,S>>>", true, []interface{}{decimalType, listType}},
		{"concrete key", int8Type, listType, "map?<i8?, list?<decimal<P,S>>>", "map<i8?, list?<decimal<P,S>>>", true, []interface{}{listType}},
		{"concrete value", decimalType, int8Type, "map?<decimal<P,S>, i8?>", "map<decimal<P,S>, i8?>", true, []interface{}{decimalType}},
		{"no parameterized param", int8Type, int8Type, "map?<i8?, i8?>", "map<i8?, i8?>", false, nil},
	} {
		t.Run(td.name, func(t *testing.T) {
			pd := &types.ParameterizedMapType{Key: td.Key, Value: td.Value}
			require.Equal(t, td.expectedNullableString, pd.SetNullability(types.NullabilityNullable).String())
			require.Equal(t, td.expectedNullableRequiredString, pd.SetNullability(types.NullabilityRequired).String())
			require.Equal(t, td.expectedHasParameterizedParam, pd.HasParameterizedParam())
			require.Equal(t, td.expectedParameterizedParams, pd.GetParameterizedParams())
		})
	}
}
