// SPDX-License-Identifier: Apache-2.0

package substraitgo_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/substrait-io/substrait-go"
)

func TestTypeToString(t *testing.T) {
	tests := []struct {
		t   Type
		exp string
	}{
		{&BooleanType{}, "boolean"},
		{&Int8Type{}, "i8"},
		{&Int16Type{}, "i16"},
		{&Int32Type{}, "i32"},
		{&Int64Type{}, "i64"},
		{&Float32Type{}, "fp32"},
		{&Float64Type{}, "fp64"},
		{&BinaryType{}, "binary"},
		{&StringType{}, "string"},
		{&TimestampType{}, "timestamp"},
		{&DateType{}, "date"},
		{&TimeType{}, "time"},
		{&TimestampTzType{}, "timestamp_tz"},
		{&IntervalYearType{}, "interval_year"},
		{&IntervalDayType{}, "interval_day"},
		{&UUIDType{Nullability: NullabilityNullable}, "uuid?"},
		{&FixedBinaryType{Length: 10}, "fixedbinary<10>"},
		{&FixedCharType{Length: 5}, "char<5>"},
		{&VarCharType{Length: 15}, "varchar<15>"},
		{&DecimalType{Scale: 2, Precision: 4}, "decimal<4,2>"},
		{&StructType{Nullability: NullabilityNullable, Types: []Type{
			&Int8Type{Nullability: NullabilityNullable},
			&DateType{Nullability: NullabilityRequired}, &FixedCharType{Length: 5}}},
			"struct<i8?, date, char<5>>?"},
		{&ListType{Type: &Int8Type{}}, "list<i8>"},
		{&MapType{Key: &StringType{}, Value: &DecimalType{Precision: 10, Scale: 2}},
			"map<string,decimal<10,2>>"},
	}

	for _, tt := range tests {
		t.Run(tt.exp, func(t *testing.T) {
			assert.Equal(t, tt.exp, tt.t.String())
		})
	}
}

func MustLiteral(l Literal, err error) Literal {
	if err != nil {
		panic(err)
	}
	return l
}

func TestLiteralToString(t *testing.T) {
	tests := []struct {
		t   Literal
		exp string
	}{
		{&PrimitiveLiteral[int16]{Value: 0, Type: &Int16Type{}}, "i16(0)"},
		{NewPrimitiveLiteral[int8](0, true), "i8?(0)"},
		{NewNestedLiteral(ListLiteralValue{
			NewNestedLiteral(MapLiteralValue{
				{
					Key:   NewPrimitiveLiteral("foo", false),
					Value: NewFixedCharLiteral(FixedChar("bar"), false),
				},
				{
					Key:   NewPrimitiveLiteral("baz", false),
					Value: NewFixedCharLiteral(FixedChar("bar"), false),
				},
			}, true),
		}, true), "list<map<string,char<3>>?>?([map<string,char<3>>?([{string(foo) char<3>(bar)} {string(baz) char<3>(bar)}])])"},
		{MustLiteral(NewLiteral(float32(1.5), false)), "fp32(1.5)"},
		{MustLiteral(NewLiteral(&VarChar{Value: "foobar", Length: 7}, true)), "varchar<7>?(foobar)"},
	}

	for _, tt := range tests {
		t.Run(tt.exp, func(t *testing.T) {
			assert.Equal(t, tt.exp, tt.t.String())
		})
	}
}
