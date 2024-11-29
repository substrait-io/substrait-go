// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	. "github.com/substrait-io/substrait-go/types"
	"github.com/substrait-io/substrait-go/types/integer_parameters"
)

func TestParameterizedSingleIntegerType(t *testing.T) {
	abstractLeafParam_L1 := integer_parameters.NewVariableIntParam("L1")
	concreteLeafParam_38 := integer_parameters.NewConcreteIntParam(38)
	concreteLeafParam_5 := integer_parameters.NewConcreteIntParam(5)
	for _, td := range []struct {
		name                           string
		typ                            FuncDefArgType
		expectedNullableString         string
		expectedNullableRequiredString string
		expectedShortString            string
		expectedIsParameterized        bool
		expectedAbstractParams         []interface{}
		expectedReturnType             Type
	}{
		{"nullable parameterized varchar", &ParameterizedVarCharType{IntegerOption: abstractLeafParam_L1}, "varchar?<L1>", "varchar<L1>", "vchar", true, []interface{}{abstractLeafParam_L1}, nil},
		{"nullable concrete varchar", &ParameterizedVarCharType{IntegerOption: concreteLeafParam_38}, "varchar?<38>", "varchar<38>", "vchar", false, nil, &VarCharType{Length: 38, Nullability: NullabilityRequired}},
		{"nullable fixChar", &ParameterizedFixedCharType{IntegerOption: abstractLeafParam_L1}, "char?<L1>", "char<L1>", "fchar", true, []interface{}{abstractLeafParam_L1}, nil},
		{"nullable concrete fixChar", &ParameterizedFixedCharType{IntegerOption: concreteLeafParam_38}, "char?<38>", "char<38>", "fchar", false, nil, &FixedCharType{Length: 38, Nullability: NullabilityRequired}},
		{"nullable fixBinary", &ParameterizedFixedBinaryType{IntegerOption: abstractLeafParam_L1}, "fixedbinary?<L1>", "fixedbinary<L1>", "fbin", true, []interface{}{abstractLeafParam_L1}, nil},
		{"nullable concrete fixBinary", &ParameterizedFixedBinaryType{IntegerOption: concreteLeafParam_38}, "fixedbinary?<38>", "fixedbinary<38>", "fbin", false, nil, &FixedBinaryType{Length: 38, Nullability: NullabilityRequired}},
		{"nullable precisionTimeStamp", &ParameterizedPrecisionTimestampType{IntegerOption: abstractLeafParam_L1}, "precision_timestamp?<L1>", "precision_timestamp<L1>", "prets", true, []interface{}{abstractLeafParam_L1}, nil},
		{"nullable concrete precisionTimeStamp", &ParameterizedPrecisionTimestampType{IntegerOption: concreteLeafParam_38}, "precision_timestamp?<38>", "precision_timestamp<38>", "prets", false, nil, &PrecisionTimestampType{Precision: 38, Nullability: NullabilityRequired}},
		{"nullable precisionTimeStampTz", &ParameterizedPrecisionTimestampTzType{IntegerOption: abstractLeafParam_L1}, "precision_timestamp_tz?<L1>", "precision_timestamp_tz<L1>", "pretstz", true, []interface{}{abstractLeafParam_L1}, nil},
		{"nullable concrete precisionTimeStampTz", &ParameterizedPrecisionTimestampTzType{IntegerOption: concreteLeafParam_38}, "precision_timestamp_tz?<38>", "precision_timestamp_tz<38>", "pretstz", false, nil, &PrecisionTimestampTzType{PrecisionTimestampType: PrecisionTimestampType{Precision: 38, Nullability: NullabilityRequired}}},
		{"nullable interval day", &ParameterizedIntervalDayType{IntegerOption: abstractLeafParam_L1}, "interval_day?<L1>", "interval_day<L1>", "iday", true, []interface{}{abstractLeafParam_L1}, nil},
		{"nullable concrete interval day", &ParameterizedIntervalDayType{IntegerOption: concreteLeafParam_5}, "interval_day?<5>", "interval_day<5>", "iday", false, nil, &IntervalDayType{Precision: 5, Nullability: NullabilityRequired}},
	} {
		t.Run(td.name, func(t *testing.T) {
			require.Equal(t, td.expectedNullableString, td.typ.SetNullability(NullabilityNullable).String())
			require.Equal(t, td.expectedNullableRequiredString, td.typ.SetNullability(NullabilityRequired).String())
			require.Equal(t, td.expectedIsParameterized, td.typ.HasParameterizedParam())
			require.Equal(t, td.expectedAbstractParams, td.typ.GetParameterizedParams())
			assert.Equal(t, td.expectedShortString, td.typ.ShortString())
			retType, err := td.typ.ReturnType(nil, nil)
			if td.expectedReturnType == nil {
				require.Error(t, err)
				require.True(t, td.typ.HasParameterizedParam())
			} else {
				require.Nil(t, err)
				require.Equal(t, td.expectedReturnType, retType)
			}
		})
	}
}
