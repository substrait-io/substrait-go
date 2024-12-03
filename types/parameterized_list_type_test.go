// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/types"
	"github.com/substrait-io/substrait-go/types/integer_parameters"
)

func TestParameterizedListType(t *testing.T) {
	decimalType := &types.ParameterizedDecimalType{
		Precision:   integer_parameters.NewVariableIntParam("P"),
		Scale:       integer_parameters.NewVariableIntParam("S"),
		Nullability: types.NullabilityRequired,
	}
	int8Type := &types.Int8Type{}
	dec30PS5 := &types.DecimalType{Precision: 30, Scale: 5, Nullability: types.NullabilityRequired}
	for _, td := range []struct {
		name                           string
		param                          types.FuncDefArgType
		args                           []interface{}
		expectedNullableString         string
		expectedNullableRequiredString string
		expectedHasParameterizedParam  bool
		expectedParameterizedParams    []interface{}
		expectedReturnType             types.Type
	}{
		{"parameterized param", decimalType, []any{dec30PS5}, "list?<decimal<P,S>>", "list<decimal<P,S>>", true, []interface{}{decimalType}, nil},
		{"concrete param", int8Type, []any{int8Type}, "list?<i8>", "list<i8>", false, nil, &types.ListType{Nullability: types.NullabilityRequired, Type: int8Type}},
	} {
		t.Run(td.name, func(t *testing.T) {
			pd := &types.ParameterizedListType{Type: td.param}
			assert.Equal(t, types.NullabilityUnspecified, pd.GetNullability())
			require.Equal(t, td.expectedNullableString, pd.SetNullability(types.NullabilityNullable).String())
			require.Equal(t, td.expectedNullableRequiredString, pd.SetNullability(types.NullabilityRequired).String())
			assert.Equal(t, types.NullabilityRequired, pd.GetNullability())
			require.Equal(t, td.expectedHasParameterizedParam, pd.HasParameterizedParam())
			require.Equal(t, td.expectedParameterizedParams, pd.GetParameterizedParams())
			assert.Equal(t, "list", pd.ShortString())
			retType, err := pd.ReturnType(nil, nil)
			if td.expectedReturnType == nil {
				assert.Error(t, err)
				require.True(t, pd.HasParameterizedParam())
			} else {
				require.Nil(t, err)
				require.Equal(t, td.expectedReturnType, retType)
				resultType, err := pd.WithParameters(td.args)
				require.Nil(t, err)
				require.Equal(t, td.expectedReturnType, resultType)
			}
		})
	}
}
