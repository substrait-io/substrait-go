// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v6/expr"
	"github.com/substrait-io/substrait-go/v6/literal"
	"github.com/substrait-io/substrait-go/v6/types"
)

func MustLiteral(l expr.Literal, err error) expr.Literal {
	if err != nil {
		panic(err)
	}
	return l
}

func TestLiteralToString(t *testing.T) {
	tests := []struct {
		t   expr.Literal
		exp string
	}{
		{&expr.PrimitiveLiteral[int16]{Value: 0, Type: &types.Int16Type{}}, "i16(0)"},
		{expr.NewPrimitiveLiteral[int8](0, true), "i8?(0)"},
		{expr.NewNestedLiteral(expr.ListLiteralValue{
			expr.NewNestedLiteral(expr.MapLiteralValue{
				{
					Key:   expr.NewPrimitiveLiteral("foo", false),
					Value: expr.NewFixedCharLiteral(types.FixedChar("bar"), false),
				},
				{
					Key:   expr.NewPrimitiveLiteral("baz", false),
					Value: expr.NewFixedCharLiteral(types.FixedChar("bar"), false),
				},
			}, true),
		}, true), "list?<map?<string, fixedchar<3>>>([map?<string, fixedchar<3>>([{string(foo) fixedchar<3>(bar)} {string(baz) fixedchar<3>(bar)}])])"},
		{MustLiteral(expr.NewLiteral(float32(1.5), false)), "fp32(1.5)"},
		{MustLiteral(expr.NewLiteral(&types.VarChar{Value: "foobar", Length: 7}, true)), "varchar?<7>(foobar)"},
		{expr.NewPrecisionTimeLiteral(123456, types.PrecisionSeconds, types.NullabilityNullable), "precision_time?<0>(10:17:36)"},
		{expr.NewPrecisionTimeLiteral(123456, types.PrecisionMilliSeconds, types.NullabilityNullable), "precision_time?<3>(00:02:03.456)"},
		{expr.NewPrecisionTimeLiteral(123456, types.PrecisionMicroSeconds, types.NullabilityNullable), "precision_time?<6>(00:00:00.123456)"},
		{expr.NewPrecisionTimeLiteral(123456, types.PrecisionNanoSeconds, types.NullabilityNullable), "precision_time?<9>(00:00:00.000123456)"},
		{expr.NewPrecisionTimestampLiteral(123456, types.PrecisionSeconds, types.NullabilityNullable), "precision_timestamp?<0>(1970-01-02 10:17:36)"},
		{expr.NewPrecisionTimestampLiteral(123456, types.PrecisionMilliSeconds, types.NullabilityNullable), "precision_timestamp?<3>(1970-01-01 00:02:03.456)"},
		{expr.NewPrecisionTimestampLiteral(123456, types.PrecisionMicroSeconds, types.NullabilityNullable), "precision_timestamp?<6>(1970-01-01 00:00:00.123456)"},
		{expr.NewPrecisionTimestampLiteral(123456, types.PrecisionNanoSeconds, types.NullabilityNullable), "precision_timestamp?<9>(1970-01-01 00:00:00.000123456)"},
		{expr.NewPrecisionTimestampTzLiteral(123456, types.PrecisionSeconds, types.NullabilityNullable), "precision_timestamp_tz?<0>(1970-01-02T10:17:36Z)"},
		{expr.NewPrecisionTimestampTzLiteral(123456, types.PrecisionMilliSeconds, types.NullabilityNullable), "precision_timestamp_tz?<3>(1970-01-01T00:02:03.456Z)"},
		{expr.NewPrecisionTimestampTzLiteral(123456, types.PrecisionMicroSeconds, types.NullabilityNullable), "precision_timestamp_tz?<6>(1970-01-01T00:00:00.123456Z)"},
		{expr.NewPrecisionTimestampTzLiteral(123456, types.PrecisionNanoSeconds, types.NullabilityNullable), "precision_timestamp_tz?<9>(1970-01-01T00:00:00.000123456Z)"},
		{MustLiteral(literal.NewDecimalFromString("12.345", false)), "decimal<5,3>(12.345)"},
		{MustLiteral(literal.NewDecimalFromString("-12.345", false)), "decimal<5,3>(-12.345)"},
	}

	for _, tt := range tests {
		t.Run(tt.exp, func(t *testing.T) {
			assert.Equal(t, tt.exp, tt.t.String())
		})
	}
}

func TestLiteralToValueString(t *testing.T) {
	tests := []struct {
		t   expr.Literal
		exp string
	}{
		{expr.NewNullLiteral(&types.Float32Type{}), "null"},
		{literal.NewBool(true, false), "true"},
		{literal.NewInt8(12, false), "12"},
		{expr.NewPrimitiveLiteral[int8](0, true), "0"},
		{literal.NewInt16(0, false), "0"},
		{literal.NewInt32(99, false), "99"},
		{literal.NewFloat32(99.10, false), "99.1"},
		{literal.NewFloat64(99.20, false), "99.2"},
		{literal.NewString("99.30", false), "99.30"},
		{MustLiteral(literal.NewDate(365, false)), "1971-01-01"},
		{MustLiteral(literal.NewTimeFromString("12:34:56", false)), "12:34:56"},
		{MustLiteral(literal.NewTimestampFromString("2021-03-05T12:34:56", false)), "2021-03-05 12:34:56"},
		{MustLiteral(literal.NewTimestampTZFromString("2021-03-05T12:34:56", false)), "2021-03-05T12:34:56Z"},
		// Test the first implementation.
		{MustLiteral(literal.NewIntervalYearsToMonth(5, 4, false)), "5 years, 4 months"},
		// Test the other implementation.
		{&expr.IntervalYearToMonthLiteral{Years: 7, Months: 6}, "7 years, 6 months"},
		{MustLiteral(literal.NewIntervalDaysToSecond(5, 4, 3, false)), "5 days, 4 seconds, 3 subseconds"},
		{t: &expr.IntervalCompoundLiteral{
			Years: 5, Months: 4, Days: 3,
			Seconds: 2, SubSeconds: 1, SubSecondPrecision: types.PrecisionMicroSeconds,
			Nullability: types.NullabilityRequired}, exp: "5 years, 4 months, 3 days, 2 seconds, 1 subseconds"},
		{MustLiteral(literal.NewUUIDFromBytes([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, false)),
			"01020304-0506-0708-090a-0b0c0d0e0f10"},
		{MustLiteral(literal.NewFixedChar("text", false)), "text"},
		{MustLiteral(literal.NewFixedBinary([]byte{1, 2, 3}, false)), "0x010203"},
		{MustLiteral(literal.NewVarChar("vartext", false)), "vartext"},
		{expr.NewNestedLiteral(expr.ListLiteralValue{
			expr.NewNestedLiteral(expr.MapLiteralValue{
				{
					Key:   expr.NewPrimitiveLiteral("foo", false),
					Value: expr.NewFixedCharLiteral(types.FixedChar("bar"), false),
				},
				{
					Key:   expr.NewPrimitiveLiteral("baz", false),
					Value: expr.NewFixedCharLiteral(types.FixedChar("bar"), false),
				},
			}, true),
		}, true), "[[{string(foo) fixedchar<3>(bar)} {string(baz) fixedchar<3>(bar)}]]"},
		{expr.NewNestedLiteral(expr.MapLiteralValue{
			{
				Key:   expr.NewPrimitiveLiteral("foo", false),
				Value: expr.NewFixedCharLiteral(types.FixedChar("bar"), false),
			},
			{
				Key:   expr.NewPrimitiveLiteral("baz", false),
				Value: expr.NewFixedCharLiteral(types.FixedChar("bar"), false),
			},
		}, true), "[{string(foo) fixedchar<3>(bar)} {string(baz) fixedchar<3>(bar)}]"},
		{MustLiteral(expr.NewLiteral(float32(1.5), false)), "1.5"},
		{MustLiteral(expr.NewLiteral(&types.VarChar{Value: "foobar", Length: 7}, true)), "foobar"},
		{expr.NewPrecisionTimeLiteral(123456, types.PrecisionSeconds, types.NullabilityNullable), "10:17:36"},
		{expr.NewPrecisionTimeLiteral(123456, types.PrecisionDeciSeconds, types.NullabilityNullable), "03:25:45.6"},
		{expr.NewPrecisionTimeLiteral(123456, types.PrecisionCentiSeconds, types.NullabilityNullable), "00:20:34.56"},
		{expr.NewPrecisionTimeLiteral(123456, types.PrecisionMilliSeconds, types.NullabilityNullable), "00:02:03.456"},
		{expr.NewPrecisionTimeLiteral(123456, types.PrecisionEMinus4Seconds, types.NullabilityNullable), "00:00:12.3456"},
		{expr.NewPrecisionTimeLiteral(123456, types.PrecisionEMinus5Seconds, types.NullabilityNullable), "00:00:01.23456"},
		{expr.NewPrecisionTimeLiteral(123456, types.PrecisionMicroSeconds, types.NullabilityNullable), "00:00:00.123456"},
		{expr.NewPrecisionTimeLiteral(123456, types.PrecisionEMinus7Seconds, types.NullabilityNullable), "00:00:00.0123456"},
		{expr.NewPrecisionTimeLiteral(123456, types.PrecisionEMinus8Seconds, types.NullabilityNullable), "00:00:00.00123456"},
		{expr.NewPrecisionTimeLiteral(123456, types.PrecisionNanoSeconds, types.NullabilityNullable), "00:00:00.000123456"},
		{expr.NewPrecisionTimestampLiteral(123456, types.PrecisionSeconds, types.NullabilityNullable), "1970-01-02 10:17:36"},
		{expr.NewPrecisionTimestampLiteral(123456, types.PrecisionDeciSeconds, types.NullabilityNullable), "1970-01-01 03:25:45.6"},
		{expr.NewPrecisionTimestampLiteral(123456, types.PrecisionCentiSeconds, types.NullabilityNullable), "1970-01-01 00:20:34.56"},
		{expr.NewPrecisionTimestampLiteral(123456, types.PrecisionMilliSeconds, types.NullabilityNullable), "1970-01-01 00:02:03.456"},
		{expr.NewPrecisionTimestampLiteral(123456, types.PrecisionEMinus4Seconds, types.NullabilityNullable), "1970-01-01 00:00:12.3456"},
		{expr.NewPrecisionTimestampLiteral(123456, types.PrecisionEMinus5Seconds, types.NullabilityNullable), "1970-01-01 00:00:01.23456"},
		{expr.NewPrecisionTimestampLiteral(123456, types.PrecisionMicroSeconds, types.NullabilityNullable), "1970-01-01 00:00:00.123456"},
		{expr.NewPrecisionTimestampLiteral(123456, types.PrecisionEMinus7Seconds, types.NullabilityNullable), "1970-01-01 00:00:00.0123456"},
		{expr.NewPrecisionTimestampLiteral(123456, types.PrecisionEMinus8Seconds, types.NullabilityNullable), "1970-01-01 00:00:00.00123456"},
		{expr.NewPrecisionTimestampLiteral(123456, types.PrecisionNanoSeconds, types.NullabilityNullable), "1970-01-01 00:00:00.000123456"},
		{expr.NewPrecisionTimestampTzLiteral(123456, types.PrecisionSeconds, types.NullabilityNullable), "1970-01-02T10:17:36Z"},
		{expr.NewPrecisionTimestampTzLiteral(123456, types.PrecisionDeciSeconds, types.NullabilityNullable), "1970-01-01T03:25:45.6Z"},
		{expr.NewPrecisionTimestampTzLiteral(123456, types.PrecisionCentiSeconds, types.NullabilityNullable), "1970-01-01T00:20:34.56Z"},
		{expr.NewPrecisionTimestampTzLiteral(123456, types.PrecisionMilliSeconds, types.NullabilityNullable), "1970-01-01T00:02:03.456Z"},
		{expr.NewPrecisionTimestampTzLiteral(123456, types.PrecisionEMinus4Seconds, types.NullabilityNullable), "1970-01-01T00:00:12.3456Z"},
		{expr.NewPrecisionTimestampTzLiteral(123456, types.PrecisionEMinus5Seconds, types.NullabilityNullable), "1970-01-01T00:00:01.23456Z"},
		{expr.NewPrecisionTimestampTzLiteral(123456, types.PrecisionMicroSeconds, types.NullabilityNullable), "1970-01-01T00:00:00.123456Z"},
		{expr.NewPrecisionTimestampTzLiteral(123456, types.PrecisionEMinus7Seconds, types.NullabilityNullable), "1970-01-01T00:00:00.0123456Z"},
		{expr.NewPrecisionTimestampTzLiteral(123456, types.PrecisionEMinus8Seconds, types.NullabilityNullable), "1970-01-01T00:00:00.00123456Z"},
		{expr.NewPrecisionTimestampTzLiteral(123456, types.PrecisionNanoSeconds, types.NullabilityNullable), "1970-01-01T00:00:00.000123456Z"},
		{MustLiteral(literal.NewDecimalFromString("12.345", false)), "12.345"},
		{MustLiteral(literal.NewDecimalFromString("-12.345", false)), "-12.345"},
		{MustLiteral(literal.NewList([]expr.Literal{literal.NewInt8(2, false), literal.NewInt8(4, false), literal.NewInt8(6, false)}, false)), "[2, 4, 6]"},
	}

	for _, tt := range tests {
		t.Run(tt.t.String(), func(t *testing.T) {
			assert.Equal(t, tt.exp, tt.t.ValueString())
		})
	}
}

func TestLiteralToStringBrokenDecimal(t *testing.T) {
	brokenDecimalLit, _ := literal.NewDecimalFromString("1234.56", false)
	brokenDecimalLitAsProtoLit := brokenDecimalLit.(*expr.ProtoLiteral)
	brokenDecimalLitAsProtoLit.Value = []byte{1, 2, 3}

	tests := []struct {
		t   expr.Literal
		exp string
	}{
		{brokenDecimalLit, "decimal<6,2>(expected 16 bytes, got 3)"},
	}

	for _, tt := range tests {
		t.Run(tt.exp, func(t *testing.T) {
			assert.Panics(t, func() {
				_ = tt.t.String()
			})
		})
	}
}
