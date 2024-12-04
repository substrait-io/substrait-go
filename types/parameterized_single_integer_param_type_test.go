// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/types"
	"github.com/substrait-io/substrait-go/types/integer_parameters"
)

func TestParameterizedSingleIntegerType(t *testing.T) {
	abstractLeafParam_L1 := integer_parameters.NewVariableIntParam("L1")
	concreteLeafParam_38 := integer_parameters.NewConcreteIntParam(38)
	concreteLeafParam_5 := integer_parameters.NewConcreteIntParam(5)
	for _, td := range []struct {
		name                           string
		typ                            types.FuncDefArgType
		args                           []interface{}
		expectedNullableString         string
		expectedNullableRequiredString string
		expectedShortString            string
		expectedIsParameterized        bool
		expectedAbstractParams         []interface{}
		expectedReturnType             types.Type
	}{
		{"nullable parameterized varchar", &types.ParameterizedVarCharType{IntegerOption: abstractLeafParam_L1}, []any{int64(11)}, "varchar?<L1>", "varchar<L1>", "vchar", true, []interface{}{abstractLeafParam_L1}, &types.VarCharType{Length: 11, Nullability: types.NullabilityRequired}},
		{"nullable concrete varchar", &types.ParameterizedVarCharType{IntegerOption: concreteLeafParam_38}, []any{}, "varchar?<38>", "varchar<38>", "vchar", false, nil, &types.VarCharType{Length: 38, Nullability: types.NullabilityRequired}},
		{"nullable fixChar", &types.ParameterizedFixedCharType{IntegerOption: abstractLeafParam_L1}, []any{int64(13)}, "char?<L1>", "char<L1>", "fchar", true, []interface{}{abstractLeafParam_L1}, &types.FixedCharType{Length: 13, Nullability: types.NullabilityRequired}},
		{"nullable concrete fixChar", &types.ParameterizedFixedCharType{IntegerOption: concreteLeafParam_38}, []any{}, "char?<38>", "char<38>", "fchar", false, nil, &types.FixedCharType{Length: 38, Nullability: types.NullabilityRequired}},
		{"nullable fixBinary", &types.ParameterizedFixedBinaryType{IntegerOption: abstractLeafParam_L1}, []any{int64(17)}, "fixedbinary?<L1>", "fixedbinary<L1>", "fbin", true, []interface{}{abstractLeafParam_L1}, &types.FixedBinaryType{Length: 17, Nullability: types.NullabilityRequired}},
		{"nullable concrete fixBinary", &types.ParameterizedFixedBinaryType{IntegerOption: concreteLeafParam_38}, []any{}, "fixedbinary?<38>", "fixedbinary<38>", "fbin", false, nil, &types.FixedBinaryType{Length: 38, Nullability: types.NullabilityRequired}},
		{"nullable precisionTimeStamp", &types.ParameterizedPrecisionTimestampType{IntegerOption: abstractLeafParam_L1}, []any{int64(7)}, "precision_timestamp?<L1>", "precision_timestamp<L1>", "prets", true, []interface{}{abstractLeafParam_L1}, &types.PrecisionTimestampType{Precision: 7, Nullability: types.NullabilityRequired}},
		{"nullable concrete precisionTimeStamp", &types.ParameterizedPrecisionTimestampType{IntegerOption: concreteLeafParam_38}, []any{}, "precision_timestamp?<38>", "precision_timestamp<38>", "prets", false, nil, &types.PrecisionTimestampType{Precision: 38, Nullability: types.NullabilityRequired}},
		{"nullable precisionTimeStampTz", &types.ParameterizedPrecisionTimestampTzType{IntegerOption: abstractLeafParam_L1}, []any{int64(5)}, "precision_timestamp_tz?<L1>", "precision_timestamp_tz<L1>", "pretstz", true, []interface{}{abstractLeafParam_L1}, &types.PrecisionTimestampTzType{PrecisionTimestampType: types.PrecisionTimestampType{Precision: 5, Nullability: types.NullabilityRequired}}},
		{"nullable concrete precisionTimeStampTz", &types.ParameterizedPrecisionTimestampTzType{IntegerOption: concreteLeafParam_38}, []any{}, "precision_timestamp_tz?<38>", "precision_timestamp_tz<38>", "pretstz", false, nil, &types.PrecisionTimestampTzType{PrecisionTimestampType: types.PrecisionTimestampType{Precision: 38, Nullability: types.NullabilityRequired}}},
		{"nullable interval day", &types.ParameterizedIntervalDayType{IntegerOption: abstractLeafParam_L1}, []any{int64(3)}, "interval_day?<L1>", "interval_day<L1>", "iday", true, []interface{}{abstractLeafParam_L1}, &types.IntervalDayType{Precision: 3, Nullability: types.NullabilityRequired}},
		{"nullable concrete interval day", &types.ParameterizedIntervalDayType{IntegerOption: concreteLeafParam_5}, []any{}, "interval_day?<5>", "interval_day<5>", "iday", false, nil, &types.IntervalDayType{Precision: 5, Nullability: types.NullabilityRequired}},
	} {
		t.Run(td.name, func(t *testing.T) {
			require.Equal(t, td.expectedNullableString, td.typ.SetNullability(types.NullabilityNullable).String())
			require.Equal(t, td.expectedNullableRequiredString, td.typ.SetNullability(types.NullabilityRequired).String())
			require.Equal(t, td.expectedIsParameterized, td.typ.HasParameterizedParam())
			require.Equal(t, td.expectedAbstractParams, td.typ.GetParameterizedParams())
			assert.Equal(t, td.expectedShortString, td.typ.ShortString())
			retType, err := td.typ.ReturnType(nil, nil)
			if td.expectedIsParameterized {
				require.Error(t, err)
				require.True(t, td.typ.HasParameterizedParam())
			} else {
				require.Nil(t, err)
				require.Equal(t, td.expectedReturnType, retType)
			}
			resultType, err := td.typ.WithParameters(td.args)
			require.Nil(t, err)
			require.Equal(t, td.expectedReturnType, resultType)
		})
	}
}
