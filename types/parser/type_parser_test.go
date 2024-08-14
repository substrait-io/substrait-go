// SPDX-License-Identifier: Apache-2.0

package parser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/types"
	"github.com/substrait-io/substrait-go/types/parser"
)

func TestParser(t *testing.T) {
	tests := []struct {
		expr      string
		expected  string
		shortName string
		typ       types.Type
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
	}

	p, err := parser.New()
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.expr, func(t *testing.T) {
			d, err := p.ParseString(tt.expr)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, d.Expr.String())
			if tt.shortName != "" {
				assert.Equal(t, tt.shortName, d.Expr.(*parser.Type).ShortType())
				typ, err := d.Expr.(*parser.Type).Type()
				assert.NoError(t, err)
				assert.True(t, tt.typ.Equals(typ))
			}
		})
	}
}
