package functions_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	substraitgo "github.com/substrait-io/substrait-go/v3"
	. "github.com/substrait-io/substrait-go/v3/functions"
	"github.com/substrait-io/substrait-go/v3/types"
)

func TestTypeRegistry(t *testing.T) {
	typeRegistry := NewTypeRegistry()
	tests := []struct {
		name string
		want types.Type
	}{
		{"i8", &types.Int8Type{Nullability: types.NullabilityRequired}},
		{"i16", &types.Int16Type{Nullability: types.NullabilityRequired}},
		{"i32", &types.Int32Type{Nullability: types.NullabilityRequired}},
		{"i64", &types.Int64Type{Nullability: types.NullabilityRequired}},
		{"fp32", &types.Float32Type{Nullability: types.NullabilityRequired}},
		{"fp64", &types.Float64Type{Nullability: types.NullabilityRequired}},
		{"string", &types.StringType{Nullability: types.NullabilityRequired}},
		{"timestamp", &types.TimestampType{Nullability: types.NullabilityRequired}},
		{"date", &types.DateType{Nullability: types.NullabilityRequired}},
		{"time", &types.TimeType{Nullability: types.NullabilityRequired}},
		{"timestamp_tz", &types.TimestampTzType{Nullability: types.NullabilityRequired}},
		{"interval_year", &types.IntervalYearType{Nullability: types.NullabilityRequired}},
		{"interval_day", &types.IntervalDayType{Nullability: types.NullabilityRequired}},
		{"uuid", &types.UUIDType{Nullability: types.NullabilityRequired}},
		{"binary", &types.BinaryType{Nullability: types.NullabilityRequired}},
		{"boolean", &types.BooleanType{Nullability: types.NullabilityRequired}},

		// short names
		{"bool", &types.BooleanType{Nullability: types.NullabilityRequired}},
		{"vbin", &types.BinaryType{Nullability: types.NullabilityRequired}},
		{"str", &types.StringType{Nullability: types.NullabilityRequired}},
		{"ts", &types.TimestampType{Nullability: types.NullabilityRequired}},
		{"tstz", &types.TimestampTzType{Nullability: types.NullabilityRequired}},
		{"iyear", &types.IntervalYearType{Nullability: types.NullabilityRequired}},
		{"iday", &types.IntervalDayType{Nullability: types.NullabilityRequired}},

		// nullable types
		{"i8?", &types.Int8Type{Nullability: types.NullabilityNullable}},
		{"timestamp_tz?", &types.TimestampTzType{Nullability: types.NullabilityNullable}},
		{"bool?", &types.BooleanType{Nullability: types.NullabilityNullable}},
		{"vbin?", &types.BinaryType{Nullability: types.NullabilityNullable}},
		{"str?", &types.StringType{Nullability: types.NullabilityNullable}},
		{"ts?", &types.TimestampType{Nullability: types.NullabilityNullable}},
		{"tstz?", &types.TimestampTzType{Nullability: types.NullabilityNullable}},
		{"iyear?", &types.IntervalYearType{Nullability: types.NullabilityNullable}},
		{"iday?", &types.IntervalDayType{Nullability: types.NullabilityNullable}},

		// parametrized types
		{"decimal<10,2>", &types.DecimalType{Precision: 10, Scale: 2, Nullability: types.NullabilityRequired}},
		{"decimal?<10,2>", &types.DecimalType{Precision: 10, Scale: 2, Nullability: types.NullabilityNullable}},
		{"decimal?<38,0>", &types.DecimalType{Precision: 38, Scale: 0, Nullability: types.NullabilityNullable}},
		{"varchar<10>", &types.VarCharType{Length: 10, Nullability: types.NullabilityRequired}},
		{"varchar?<10>", &types.VarCharType{Length: 10, Nullability: types.NullabilityNullable}},
		{"fixedchar<10>", &types.FixedCharType{Length: 10, Nullability: types.NullabilityRequired}},
		{"fixedchar?<10>", &types.FixedCharType{Length: 10, Nullability: types.NullabilityNullable}},
		{"fixedbinary<10>", &types.FixedBinaryType{Length: 10, Nullability: types.NullabilityRequired}},
		{"fixedbinary?<10>", &types.FixedBinaryType{Length: 10, Nullability: types.NullabilityNullable}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			typ, err := typeRegistry.GetTypeFromTypeString(tt.name)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, typ)
		})
	}

	negativeTests := []struct {
		name string
	}{
		{"badType"},
		{"nonexistent?"},
	}
	for _, tt := range negativeTests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := typeRegistry.GetTypeFromTypeString(tt.name)
			assert.Error(t, err)
		})
	}
}

func TestLocalTypeRegistry(t *testing.T) {
	typeRegistry := NewTypeRegistry()
	testDialect := `---
name: testSql
type: sql
dependencies:
  arithmetic: 
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml
supported_types:
  i32:
    sql_type_name: int32
    supported_as_column: true
  i64:
    sql_type_name: int64
  date:
    sql_type_name: DATE
    supported_as_column: true
  iyear:
    sql_type_name: INTERVAL
    supported_as_column: false
  ts:
    sql_type_name: TIMESTAMP
    supported_as_column: true
  dec:
    sql_type_name: NUMERIC
  vchar:
    sql_type_name: VARCHAR
  fchar:
    sql_type_name: CHAR
  fbin:
    sql_type_name: BINARY
scalar_functions:
- name: arithmetic.add
  local_name: +
  infix: true
  required_options:
    overflow: ERROR
    rounding: TIE_TO_EVEN
  supported_kernels:
  - i32_i32
  - i64_i64
`
	dialect, err := LoadDialect("testSql", strings.NewReader(testDialect))
	assert.NoError(t, err)
	localTypeRegistry, err := dialect.LocalizeTypeRegistry(typeRegistry)
	assert.NoError(t, err)

	tests := []struct {
		name      string
		localName string
		want      types.Type
		asColumn  bool
	}{
		{"i32", "int32", &types.Int32Type{Nullability: types.NullabilityRequired}, true},
		{"i64", "int64", &types.Int64Type{Nullability: types.NullabilityRequired}, true},
		{"date", "DATE", &types.DateType{Nullability: types.NullabilityRequired}, true},
		{"iyear", "INTERVAL", &types.IntervalYearType{Nullability: types.NullabilityRequired}, false},
		{"timestamp", "TIMESTAMP", &types.TimestampType{Nullability: types.NullabilityRequired}, true},
		{"dec<10,2>", "NUMERIC(10,2)", &types.DecimalType{Nullability: types.NullabilityRequired, Precision: 10, Scale: 2}, true},
		{"varchar<10>", "VARCHAR(10)", &types.VarCharType{Nullability: types.NullabilityRequired, Length: 10}, true},
		{"fixedchar<10>", "CHAR(10)", &types.FixedCharType{Nullability: types.NullabilityRequired, Length: 10}, true},
		{"fixedbinary<10>", "BINARY(10)", &types.FixedBinaryType{Nullability: types.NullabilityRequired, Length: 10}, true},

		// short names
		{"ts", "TIMESTAMP", &types.TimestampType{Nullability: types.NullabilityRequired}, true},

		// nullable types
		{"i32?", "int32", &types.Int32Type{Nullability: types.NullabilityNullable}, true},
		{"i64?", "int64", &types.Int64Type{Nullability: types.NullabilityNullable}, true},
		{"date?", "DATE", &types.DateType{Nullability: types.NullabilityNullable}, true},
		{"iyear?", "INTERVAL", &types.IntervalYearType{Nullability: types.NullabilityNullable}, false},
		{"timestamp?", "TIMESTAMP", &types.TimestampType{Nullability: types.NullabilityNullable}, true},
		{"dec?<10,2>", "NUMERIC(10,2)", &types.DecimalType{Nullability: types.NullabilityNullable, Precision: 10, Scale: 2}, true},
		{"varchar?<10>", "VARCHAR(10)", &types.VarCharType{Nullability: types.NullabilityNullable, Length: 10}, true},
		{"fixedchar?<10>", "CHAR(10)", &types.FixedCharType{Nullability: types.NullabilityNullable, Length: 10}, true},
		{"fixedbinary?<10>", "BINARY(10)", &types.FixedBinaryType{Nullability: types.NullabilityNullable, Length: 10}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			typ, err := localTypeRegistry.GetTypeFromTypeString(tt.name)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, typ)

			typ, err = localTypeRegistry.GetSubstraitTypeFromLocalType(tt.localName)
			assert.NoError(t, err)
			assert.Equal(t, tt.want.WithNullability(types.NullabilityRequired), typ)

			localType, err := localTypeRegistry.GetLocalTypeFromSubstraitType(tt.want)
			assert.NoError(t, err)
			assert.Equal(t, tt.localName, localType)

			assert.Equal(t, tt.asColumn, localTypeRegistry.IsTypeSupportedInTables(tt.want),
				"IsTypeSupportedInTables(%s) failed", tt.name)
		})
	}

	negativeTests := []struct {
		name      string
		localName string
		typ       types.Type
	}{
		{"i8", "int8", &types.Int8Type{Nullability: types.NullabilityRequired}},
		{"decimal<10>", "NUMERIC(10)", nil},
		{"decimal<4, 2, 1>", "NUMERIC(4, 2, 1)", nil},
		{"char<20,30>", "CHAR(20, 30)", nil},
		{"fixedbinary<10,20,30>", "BINARY(10, 20, 30)", nil},
		{"i64<10>", "int64<10>", nil},
		{"non_existent", "NON_EXISTENT", nil},
	}
	for _, tt := range negativeTests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := localTypeRegistry.GetTypeFromTypeString(tt.name)
			assert.Error(t, err, substraitgo.ErrNotFound)

			_, err = localTypeRegistry.GetSubstraitTypeFromLocalType(tt.localName)
			assert.Error(t, err, substraitgo.ErrNotFound)

			if tt.typ != nil {
				_, err = localTypeRegistry.GetLocalTypeFromSubstraitType(tt.typ)
				assert.Error(t, err, substraitgo.ErrNotFound)

				assert.False(t, localTypeRegistry.IsTypeSupportedInTables(tt.typ))
			}
		})
	}

}
