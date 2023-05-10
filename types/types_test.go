// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/substrait-io/substrait-go/types"
)

func TestTypeToString(t *testing.T) {
	tests := []struct {
		t        Type
		exp      string
		expShort string
	}{
		{&BooleanType{}, "boolean", "bool"},
		{&Int8Type{}, "i8", "i8"},
		{&Int16Type{}, "i16", "i16"},
		{&Int32Type{}, "i32", "i32"},
		{&Int64Type{}, "i64", "i64"},
		{&Float32Type{}, "fp32", "fp32"},
		{&Float64Type{}, "fp64", "fp64"},
		{&BinaryType{}, "binary", "vbin"},
		{&StringType{}, "string", "str"},
		{&TimestampType{}, "timestamp", "ts"},
		{&DateType{}, "date", "date"},
		{&TimeType{}, "time", "time"},
		{&TimestampTzType{}, "timestamp_tz", "tstz"},
		{&IntervalYearType{}, "interval_year", "iyear"},
		{&IntervalDayType{}, "interval_day", "iday"},
		{&UUIDType{Nullability: NullabilityNullable}, "uuid?", "uuid"},
		{&FixedBinaryType{Length: 10}, "fixedbinary<10>", "fbin"},
		{&FixedCharType{Length: 5}, "char<5>", "fchar"},
		{&VarCharType{Length: 15}, "varchar<15>", "vchar"},
		{&DecimalType{Scale: 2, Precision: 4}, "decimal<4,2>", "dec"},
		{&StructType{Nullability: NullabilityNullable, Types: []Type{
			&Int8Type{Nullability: NullabilityNullable},
			&DateType{Nullability: NullabilityRequired}, &FixedCharType{Length: 5}}},
			"struct?<i8?, date, char<5>>", "struct"},
		{&ListType{Type: &Int8Type{}}, "list<i8>", "list"},
		{&MapType{Key: &StringType{}, Value: &DecimalType{Precision: 10, Scale: 2}},
			"map<string,decimal<10,2>>", "map"},
	}

	for _, tt := range tests {
		t.Run(tt.exp, func(t *testing.T) {
			assert.Equal(t, tt.exp, tt.t.String())
			assert.Equal(t, tt.expShort, tt.t.ShortString())
		})
	}
}

func TestTypeRoundtrip(t *testing.T) {
	for _, nullable := range []bool{true, false} {
		t.Run(fmt.Sprintf("nullable=%t", nullable), func(t *testing.T) {
			n := NullabilityRequired
			if nullable {
				n = NullabilityNullable
			}

			tests := []Type{
				&BooleanType{Nullability: n},
				&Int8Type{Nullability: n},
				&Int16Type{Nullability: n},
				&Int32Type{Nullability: n},
				&Int64Type{Nullability: n},
				&Float32Type{Nullability: n},
				&Float64Type{Nullability: n},
				&StringType{Nullability: n},
				&BinaryType{Nullability: n},
				&TimeType{Nullability: n},
				&DateType{Nullability: n},
				&TimestampType{Nullability: n},
				&TimestampTzType{Nullability: n},
				&IntervalYearType{Nullability: n},
				&IntervalDayType{Nullability: n},
				&UUIDType{Nullability: n},
				&FixedCharType{Nullability: n, Length: 25},
				&VarCharType{Nullability: n, Length: 35},
				&FixedBinaryType{Nullability: n, Length: 45},
				&DecimalType{Nullability: n, Precision: 34, Scale: 3},
				&MapType{Nullability: n, Key: &Int8Type{}, Value: &Int16Type{Nullability: n}},
				&ListType{Nullability: n, Type: &TimeType{Nullability: n}},
				&StructType{Nullability: n, Types: []Type{
					&TimeType{Nullability: n}, &TimestampType{Nullability: n},
					&TimestampTzType{Nullability: n}}},
			}

			for _, tt := range tests {
				t.Run(tt.String(), func(t *testing.T) {
					converted := TypeToProto(tt)
					assert.True(t, tt.Equals(TypeFromProto(converted)))
				})
			}
		})
	}
}
