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
		expectedNullableString         string
		expectedNullableRequiredString string
		expectedHasParameterizedParam  bool
		expectedParameterizedParams    []interface{}
		expectedReturnType             types.Type
	}{
		{"both parameterized", precision_P, scale_S, "decimal?<P,S>", "decimal<P,S>", true, []interface{}{precision_P, scale_S}, nil},
		{"precision concrete", precision_38, scale_S, "decimal?<38,S>", "decimal<38,S>", true, []interface{}{scale_S}, nil},
		{"scale concrete", precision_P, scale_5, "decimal?<P,5>", "decimal<P,5>", true, []interface{}{precision_P}, nil},
		{"both concrete", precision_38, scale_5, "decimal?<38,5>", "decimal<38,5>", false, nil, &types.DecimalType{Precision: 38, Scale: 5, Nullability: types.NullabilityRequired}},
	} {
		t.Run(td.name, func(t *testing.T) {
			pd := &types.ParameterizedDecimalType{Precision: td.precision, Scale: td.scale}
			require.Equal(t, td.expectedNullableString, pd.SetNullability(types.NullabilityNullable).String())
			require.Equal(t, td.expectedNullableRequiredString, pd.SetNullability(types.NullabilityRequired).String())
			require.Equal(t, td.expectedHasParameterizedParam, pd.HasParameterizedParam())
			require.Equal(t, td.expectedParameterizedParams, pd.GetParameterizedParams())
			retType, err := pd.ReturnType(nil, nil)
			if td.expectedReturnType == nil {
				require.Error(t, err)
				require.True(t, pd.HasParameterizedParam())
			} else {
				require.Nil(t, err)
				require.Equal(t, td.expectedReturnType, retType)
			}
		})
	}
}
