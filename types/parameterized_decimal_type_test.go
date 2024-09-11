// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/types"
	"github.com/substrait-io/substrait-go/types/leaf_parameters"
)

func TestParameterizedDecimalType(t *testing.T) {
	precision_P := leaf_parameters.NewVariableIntParam("P")
	scale_S := leaf_parameters.NewVariableIntParam("S")
	precision_38 := leaf_parameters.NewConcreteIntParam(38)
	scale_5 := leaf_parameters.NewConcreteIntParam(5)
	for _, td := range []struct {
		name                           string
		precision                      leaf_parameters.LeafParameter
		scale                          leaf_parameters.LeafParameter
		expectedNullableString         string
		expectedNullableRequiredString string
		expectedHasParameterizedParam  bool
		expectedParameterizedParams    []interface{}
	}{
		{"both parameterized", precision_P, scale_S, "decimal?<P,S>", "decimal<P,S>", true, []interface{}{precision_P, scale_S}},
		{"precision concrete", precision_38, scale_S, "decimal?<38,S>", "decimal<38,S>", true, []interface{}{scale_S}},
		{"scale concrete", precision_P, scale_5, "decimal?<P,5>", "decimal<P,5>", true, []interface{}{precision_P}},
		{"both concrete", precision_38, scale_5, "decimal?<38,5>", "decimal<38,5>", false, nil},
	} {
		t.Run(td.name, func(t *testing.T) {
			pd := &types.ParameterizedDecimalType{Precision: td.precision, Scale: td.scale}
			require.Equal(t, td.expectedNullableString, pd.SetNullability(types.NullabilityNullable).String())
			require.Equal(t, td.expectedNullableRequiredString, pd.SetNullability(types.NullabilityRequired).String())
			require.Equal(t, td.expectedHasParameterizedParam, pd.HasParameterizedParam())
			require.Equal(t, td.expectedParameterizedParams, pd.GetParameterizedParams())
		})
	}
}
