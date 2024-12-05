// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/types"
	"github.com/substrait-io/substrait-go/types/integer_parameters"
)

func TestParameterizedDecimalType(t *testing.T) {
	precision_P := integer_parameters.NewVariableIntParam("P")
	scale_S := integer_parameters.NewVariableIntParam("S")
	precision_38 := integer_parameters.NewConcreteIntParam(38)
	scale_5 := integer_parameters.NewConcreteIntParam(5)
	for _, td := range []struct {
		name                           string
		precision                      integer_parameters.IntegerParameter
		scale                          integer_parameters.IntegerParameter
		args                           []interface{}
		expectedNullableString         string
		expectedNullableRequiredString string
		expectedHasParameterizedParam  bool
		expectedParameterizedParams    []interface{}
		expectedReturnType             types.Type
	}{
		{"both parameterized", precision_P, scale_S, []any{int64(30), int64(13)}, "decimal?<P,S>", "decimal<P,S>", true, []interface{}{precision_P, scale_S}, &types.DecimalType{Precision: 30, Scale: 13, Nullability: types.NullabilityRequired}},
		{"precision concrete", precision_38, scale_S, []any{int64(38), int64(6)}, "decimal?<38,S>", "decimal<38,S>", true, []interface{}{precision_38, scale_S}, &types.DecimalType{Precision: 38, Scale: 6, Nullability: types.NullabilityRequired}},
		{"scale concrete", precision_P, scale_5, []any{int64(30), int64(5)}, "decimal?<P,5>", "decimal<P,5>", true, []interface{}{precision_P, scale_5}, &types.DecimalType{Precision: 30, Scale: 5, Nullability: types.NullabilityRequired}},
		{"both concrete", precision_38, scale_5, []any{}, "decimal?<38,5>", "decimal<38,5>", false, nil, &types.DecimalType{Precision: 38, Scale: 5, Nullability: types.NullabilityRequired}},
	} {
		t.Run(td.name, func(t *testing.T) {
			pd := &types.ParameterizedDecimalType{Precision: td.precision, Scale: td.scale}
			require.Equal(t, td.expectedNullableString, pd.SetNullability(types.NullabilityNullable).String())
			require.Equal(t, types.NullabilityNullable, pd.GetNullability())
			require.Equal(t, td.expectedNullableRequiredString, pd.SetNullability(types.NullabilityRequired).String())
			require.Equal(t, types.NullabilityRequired, pd.GetNullability())
			require.Equal(t, td.expectedHasParameterizedParam, pd.HasParameterizedParam())
			require.Equal(t, td.expectedParameterizedParams, pd.GetParameterizedParams())
			require.Equal(t, "dec", pd.ShortString())
			retType, err := pd.ReturnType(nil, nil)
			if td.expectedHasParameterizedParam {
				require.Error(t, err)
				require.True(t, pd.HasParameterizedParam())
				retType, err = pd.ReturnType([]types.FuncDefArgType{pd}, []types.Type{td.expectedReturnType})
				require.NoError(t, err)
				require.Equal(t, td.expectedReturnType, retType)
			} else {
				require.Nil(t, err)
				require.Equal(t, td.expectedReturnType, retType)
			}
			resultType, err := pd.WithParameters(td.args)
			require.Nil(t, err)
			require.Equal(t, td.expectedReturnType, resultType)
		})
	}
}
