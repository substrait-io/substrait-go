package functions

import (
	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/types"
	"testing"
)

func TestNewTypeRegistry(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			typ, err := typeRegistry.GetTypeFromTypeString(tt.name)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, typ)
		})
	}
}
