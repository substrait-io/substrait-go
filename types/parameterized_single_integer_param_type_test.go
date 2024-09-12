// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/types"
	"github.com/substrait-io/substrait-go/types/integer_parameters"
)

func TestParameterizedSingleIntegerType(t *testing.T) {
	abstractLeafParam_L1 := integer_parameters.NewVariableIntParam("L1")
	concreteLeafParam_38 := integer_parameters.NewConcreteIntParam(38)
	for _, td := range []struct {
		name                           string
		typ                            types.FuncDefArgType
		expectedNullableString         string
		expectedNullableRequiredString string
		expectedIsParameterized        bool
		expectedAbstractParams         []interface{}
	}{
		{"nullable parameterized varchar", &types.ParameterizedVarCharType{IntegerOption: abstractLeafParam_L1}, "varchar?<L1>", "varchar<L1>", true, []interface{}{abstractLeafParam_L1}},
		{"nullable concrete varchar", &types.ParameterizedVarCharType{IntegerOption: concreteLeafParam_38}, "varchar?<38>", "varchar<38>", false, nil},
		{"nullable fixChar", &types.ParameterizedFixedCharType{IntegerOption: abstractLeafParam_L1}, "char?<L1>", "char<L1>", true, []interface{}{abstractLeafParam_L1}},
		{"nullable concrete fixChar", &types.ParameterizedFixedCharType{IntegerOption: concreteLeafParam_38}, "char?<38>", "char<38>", false, nil},
		{"nullable fixBinary", &types.ParameterizedFixedBinaryType{IntegerOption: abstractLeafParam_L1}, "fixedbinary?<L1>", "fixedbinary<L1>", true, []interface{}{abstractLeafParam_L1}},
		{"nullable concrete fixBinary", &types.ParameterizedFixedBinaryType{IntegerOption: concreteLeafParam_38}, "fixedbinary?<38>", "fixedbinary<38>", false, nil},
		{"nullable precisionTimeStamp", &types.ParameterizedPrecisionTimestampType{IntegerOption: abstractLeafParam_L1}, "precision_timestamp?<L1>", "precision_timestamp<L1>", true, []interface{}{abstractLeafParam_L1}},
		{"nullable concrete precisionTimeStamp", &types.ParameterizedPrecisionTimestampType{IntegerOption: concreteLeafParam_38}, "precision_timestamp?<38>", "precision_timestamp<38>", false, nil},
		{"nullable precisionTimeStampTz", &types.ParameterizedPrecisionTimestampTzType{IntegerOption: abstractLeafParam_L1}, "precision_timestamp_tz?<L1>", "precision_timestamp_tz<L1>", true, []interface{}{abstractLeafParam_L1}},
		{"nullable concrete precisionTimeStampTz", &types.ParameterizedPrecisionTimestampTzType{IntegerOption: concreteLeafParam_38}, "precision_timestamp_tz?<38>", "precision_timestamp_tz<38>", false, nil},
	} {
		t.Run(td.name, func(t *testing.T) {
			require.Equal(t, td.expectedNullableString, td.typ.SetNullability(types.NullabilityNullable).String())
			require.Equal(t, td.expectedNullableRequiredString, td.typ.SetNullability(types.NullabilityRequired).String())
			require.Equal(t, td.expectedIsParameterized, td.typ.HasParameterizedParam())
			require.Equal(t, td.expectedAbstractParams, td.typ.GetParameterizedParams())
		})
	}
}
