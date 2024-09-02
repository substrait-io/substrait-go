// SPDX-License-Identifier: Apache-2.0

package parser_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/types"
	"github.com/substrait-io/substrait-go/types/parser"
)

func TestParser(t *testing.T) {
	tests := []struct {
		expr        string
		expected    string
		shortName   string
		expectedTyp types.Type
	}{
		{"2", "2", "", nil},
		{"-2", "-2", "", nil},
		{"i16?", "i16?", "i16", &types.Int16Type{Nullability: types.NullabilityNullable}},
		{"boolean", "boolean", "bool", &types.BooleanType{Nullability: types.NullabilityRequired}},
		{"fixedchar<5>", "fixedchar<5>", "fchar", &types.FixedCharType{Length: 5}},
		{"decimal<10,5>", "decimal<10, 5>", "dec", &types.DecimalType{Precision: 10, Scale: 5}},
		{"list<decimal<10,5>>", "list<decimal<10, 5>>", "list", &types.ListType{Type: &types.DecimalType{Precision: 10, Scale: 5}, Nullability: types.NullabilityRequired}},
		{"list?<decimal?<10,5>>", "list?<decimal?<10, 5>>", "list", &types.ListType{Type: &types.DecimalType{Precision: 10, Scale: 5}, Nullability: types.NullabilityNullable}},
		{"struct<i16?,i32>", "struct<i16?, i32>", "struct", &types.StructType{Types: []types.Type{&types.Int16Type{Nullability: types.NullabilityNullable}, &types.Int32Type{Nullability: types.NullabilityRequired}}, Nullability: types.NullabilityRequired}},
		{"map<boolean?,struct?<i16?,i32?,i64?>>", "map<boolean?,struct?<i16?, i32?, i64?>>", "map", &types.MapType{Key: &types.BooleanType{Nullability: types.NullabilityNullable}, Value: &types.StructType{Types: []types.Type{&types.Int16Type{Nullability: types.NullabilityNullable}, &types.Int32Type{Nullability: types.NullabilityNullable}, &types.Int64Type{Nullability: types.NullabilityNullable}}, Nullability: types.NullabilityNullable}, Nullability: types.NullabilityRequired}},
		{"map?<boolean?,struct?<i16?,i32?,i64?>>", "map?<boolean?,struct?<i16?, i32?, i64?>>", "map", &types.MapType{Key: &types.BooleanType{Nullability: types.NullabilityNullable}, Value: &types.StructType{Types: []types.Type{&types.Int16Type{Nullability: types.NullabilityNullable}, &types.Int32Type{Nullability: types.NullabilityNullable}, &types.Int64Type{Nullability: types.NullabilityNullable}}, Nullability: types.NullabilityNullable}, Nullability: types.NullabilityNullable}},
		{"precision_timestamp<5>", "precision_timestamp<5>", "prets", &types.PrecisionTimestampType{Precision: types.PrecisionEMinus5Seconds}},
		{"precision_timestamp_tz<5>", "precision_timestamp_tz<5>", "pretstz", &types.PrecisionTimestampTzType{PrecisionTimestampType: types.PrecisionTimestampType{Precision: types.PrecisionEMinus5Seconds}}},
		{"varchar<L1>", "varchar<L1>", "vchar", types.ParameterizedVarCharType{IntegerOption: types.IntegerParam{Name: "L1"}}},
		{"fixedchar<L1>", "fixedchar<L1>", "fchar", types.ParameterizedFixedCharType{IntegerOption: types.IntegerParam{Name: "L1"}}},
		{"fixedbinary<L1>", "fixedbinary<L1>", "fbin", types.ParameterizedFixedBinaryType{IntegerOption: types.IntegerParam{Name: "L1"}}},
		{"precision_timestamp<L1>", "precision_timestamp<L1>", "prets", types.ParameterizedPrecisionTimestampType{IntegerOption: types.IntegerParam{Name: "L1"}}},
		{"precision_timestamp_tz<L1>", "precision_timestamp_tz<L1>", "pretstz", types.ParameterizedPrecisionTimestampTzType{IntegerOption: types.IntegerParam{Name: "L1"}}},
		{"decimal<P,S>", "decimal<P, S>", "dec", types.ParameterizedDecimalType{Precision: types.IntegerParam{Name: "P"}, Scale: types.IntegerParam{Name: "S"}}},
	}

	p, err := parser.New()
	require.NoError(t, err)

	for _, td := range tests {
		t.Run(td.expr, func(t *testing.T) {
			d, err := p.ParseString(td.expr)
			assert.NoError(t, err)
			assert.Equal(t, td.expected, d.Expr.String())
			if td.shortName != "" {
				assert.Equal(t, td.shortName, d.Expr.(*parser.Type).ShortType())
				typ, err := d.Expr.(*parser.Type).Type()
				assert.NoError(t, err)
				assert.Equal(t, reflect.TypeOf(td.expectedTyp), reflect.TypeOf(typ))
				assert.True(t, td.expectedTyp.Equals(typ))
			}
		})
	}
}
