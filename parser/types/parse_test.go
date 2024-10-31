package types

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/types"
	"github.com/substrait-io/substrait-go/types/integer_parameters"
)

func TestParseSubstraitType(t *testing.T) {
	tests := []struct {
		input     string
		expected  string
		shortName string
		want      types.Type
	}{
		{"boolean", "boolean", "bool", &types.BooleanType{Nullability: types.NullabilityRequired}},
		{"i8", "i8", "i8", &types.Int8Type{Nullability: types.NullabilityRequired}},
		{"i16", "i16", "i16", &types.Int16Type{Nullability: types.NullabilityRequired}},
		{"i32", "i32", "i32", &types.Int32Type{Nullability: types.NullabilityRequired}},
		{"i64", "i64", "i64", &types.Int64Type{Nullability: types.NullabilityRequired}},
		{"fp32", "fp32", "fp32", &types.Float32Type{Nullability: types.NullabilityRequired}},
		{"fp64", "fp64", "fp64", &types.Float64Type{Nullability: types.NullabilityRequired}},
		{"string", "string", "str", &types.StringType{Nullability: types.NullabilityRequired}},
		{"binary", "binary", "bin", &types.BinaryType{Nullability: types.NullabilityRequired}},
		{"timestamp", "timestamp", "ts", &types.TimestampType{Nullability: types.NullabilityRequired}},
		{"timestamp_tz", "timestamp_tz", "tstz", &types.TimestampTzType{Nullability: types.NullabilityRequired}},
		{"date", "date", "date", &types.DateType{Nullability: types.NullabilityRequired}},
		{"time", "time", "time", &types.TimeType{Nullability: types.NullabilityRequired}},
		{"uuid", "uuid", "uuid", &types.UUIDType{Nullability: types.NullabilityRequired}},
		{"interval_year", "interval_year", "iyear", &types.IntervalYearType{Nullability: types.NullabilityRequired}},
		{"I8", "i8", "i8", &types.Int8Type{Nullability: types.NullabilityRequired}},
		{"Boolean", "boolean", "bool", &types.BooleanType{Nullability: types.NullabilityRequired}},
		{"i16?", "i16?", "i16", &types.Int16Type{Nullability: types.NullabilityNullable}},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := ParseType(tt.input)
			assert.NoError(t, err)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseSubstraitType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseFuncDefArgType(t *testing.T) {
	concreteInt5 := integer_parameters.ConcreteIntParam(5)
	concreteInt10 := integer_parameters.ConcreteIntParam(10)
	concreteInt38 := integer_parameters.ConcreteIntParam(38)
	variableIntL1 := integer_parameters.VariableIntParam("L1")
	variableIntP := integer_parameters.VariableIntParam("P")
	variableIntS := integer_parameters.VariableIntParam("S")
	tests := []struct {
		input     string
		expected  string
		shortName string
		want      types.FuncDefArgType
	}{
		{"I8", "i8", "i8", &types.Int8Type{Nullability: types.NullabilityRequired}},
		{"Boolean", "boolean", "bool", &types.BooleanType{Nullability: types.NullabilityRequired}},
		{"i8", "i8", "i8", &types.Int8Type{Nullability: types.NullabilityRequired}},
		{"string", "string", "str", &types.StringType{Nullability: types.NullabilityRequired}},
		{"i16?", "i16?", "i16", &types.Int16Type{Nullability: types.NullabilityNullable}},
		{"boolean", "boolean", "bool", &types.BooleanType{Nullability: types.NullabilityRequired}},

		{"fixedchar<5>", "char<5>", "fchar", &types.ParameterizedFixedCharType{IntegerOption: &concreteInt5, Nullability: types.NullabilityRequired}},
		{"decimal<10,5>", "decimal<10,5>", "dec", &types.ParameterizedDecimalType{Precision: &concreteInt10, Scale: &concreteInt5, Nullability: types.NullabilityRequired}},
		{"list<decimal<10,5>>", "list<decimal<10,5>>", "list", &types.ParameterizedListType{Type: &types.ParameterizedDecimalType{Precision: &concreteInt10, Scale: &concreteInt5, Nullability: types.NullabilityRequired}, Nullability: types.NullabilityRequired}},
		{"list?<decimal?<10,5>>", "list?<decimal?<10,5>>", "list", &types.ParameterizedListType{Type: &types.ParameterizedDecimalType{Precision: &concreteInt10, Scale: &concreteInt5, Nullability: types.NullabilityNullable}, Nullability: types.NullabilityNullable}},
		{"struct<i16?,i32>", "struct<i16?, i32>", "struct", &types.ParameterizedStructType{Types: []types.FuncDefArgType{&types.Int16Type{Nullability: types.NullabilityNullable}, &types.Int32Type{Nullability: types.NullabilityRequired}}, Nullability: types.NullabilityRequired}},
		{"map<boolean?,struct?<i16?,i32?,i64?>>", "map<boolean?, struct?<i16?, i32?, i64?>>", "map", &types.ParameterizedMapType{Key: &types.BooleanType{Nullability: types.NullabilityNullable}, Value: &types.ParameterizedStructType{Types: []types.FuncDefArgType{&types.Int16Type{Nullability: types.NullabilityNullable}, &types.Int32Type{Nullability: types.NullabilityNullable}, &types.Int64Type{Nullability: types.NullabilityNullable}}, Nullability: types.NullabilityNullable}, Nullability: types.NullabilityRequired}},
		{"map?<boolean?,struct?<i16?,i32?,i64?>>", "map?<boolean?, struct?<i16?, i32?, i64?>>", "map", &types.ParameterizedMapType{Key: &types.BooleanType{Nullability: types.NullabilityNullable}, Value: &types.ParameterizedStructType{Types: []types.FuncDefArgType{&types.Int16Type{Nullability: types.NullabilityNullable}, &types.Int32Type{Nullability: types.NullabilityNullable}, &types.Int64Type{Nullability: types.NullabilityNullable}}, Nullability: types.NullabilityNullable}, Nullability: types.NullabilityNullable}},
		{"precision_timestamp<5>", "precision_timestamp<5>", "prets", &types.ParameterizedPrecisionTimestampType{IntegerOption: &concreteInt5, Nullability: types.NullabilityRequired}},
		{"precision_timestamp_tz<5>", "precision_timestamp_tz<5>", "pretstz", &types.ParameterizedPrecisionTimestampTzType{IntegerOption: &concreteInt5, Nullability: types.NullabilityRequired}},

		{"varchar<L1>", "varchar<L1>", "vchar", &types.ParameterizedVarCharType{IntegerOption: &variableIntL1, Nullability: types.NullabilityRequired}},
		{"fixedchar<L1>", "char<L1>", "fchar", &types.ParameterizedFixedCharType{IntegerOption: &variableIntL1, Nullability: types.NullabilityRequired}},
		{"fixedbinary<L1>", "fixedbinary<L1>", "fbin", &types.ParameterizedFixedBinaryType{IntegerOption: &variableIntL1, Nullability: types.NullabilityRequired}},
		{"precision_timestamp<L1>", "precision_timestamp<L1>", "prets", &types.ParameterizedPrecisionTimestampType{IntegerOption: &variableIntL1, Nullability: types.NullabilityRequired}},
		{"precision_timestamp_tz<L1>", "precision_timestamp_tz<L1>", "pretstz", &types.ParameterizedPrecisionTimestampTzType{IntegerOption: &variableIntL1, Nullability: types.NullabilityRequired}},
		{"decimal<P,S>", "decimal<P,S>", "dec", &types.ParameterizedDecimalType{Precision: &variableIntP, Scale: &variableIntS, Nullability: types.NullabilityRequired}},
		{"decimal<38,S>", "decimal<38,S>", "dec", &types.ParameterizedDecimalType{Precision: &concreteInt38, Scale: &variableIntS, Nullability: types.NullabilityRequired}},
		{"any", "any", "any", types.AnyType{Name: "any", Nullability: types.NullabilityRequired}},
		{"any1?", "any1?", "any", types.AnyType{Name: "any1", Nullability: types.NullabilityNullable}},
		{"list<decimal<P,S>>", "list<decimal<P,S>>", "list", &types.ParameterizedListType{Type: &types.ParameterizedDecimalType{Precision: &variableIntP, Scale: &variableIntS, Nullability: types.NullabilityRequired}, Nullability: types.NullabilityRequired}},
		{"struct<list?<decimal<P,S>>, i16>", "struct<list?<decimal<P,S>>, i16>", "struct", &types.ParameterizedStructType{Types: []types.FuncDefArgType{&types.ParameterizedListType{Type: &types.ParameterizedDecimalType{Precision: &variableIntP, Scale: &variableIntS, Nullability: types.NullabilityRequired}, Nullability: types.NullabilityNullable}, &types.Int16Type{Nullability: types.NullabilityRequired}}, Nullability: types.NullabilityRequired}},
		{"map<decimal<P,S>, i16>", "map<decimal<P,S>, i16>", "map", &types.ParameterizedMapType{Key: &types.ParameterizedDecimalType{Precision: &variableIntP, Scale: &variableIntS, Nullability: types.NullabilityRequired}, Value: &types.Int16Type{Nullability: types.NullabilityRequired}, Nullability: types.NullabilityRequired}},
		{"precision_timestamp_tz?<L1>", "precision_timestamp_tz?<L1>", "pretstz", &types.ParameterizedPrecisionTimestampTzType{IntegerOption: &variableIntL1, Nullability: types.NullabilityNullable}},
		{"u!test", "u!test", "u!test", &types.ParameterizedUserDefinedType{Name: "test", Nullability: types.NullabilityRequired}},
		{"u!test<10>", "u!test<10>", "u!test", &types.ParameterizedUserDefinedType{Name: "test", Nullability: types.NullabilityRequired, TypeParameters: []types.UDTParameter{&types.IntegerUDTParam{Integer: 10}}}},
		{"u!test<L1>", "u!test<L1>", "u!test", &types.ParameterizedUserDefinedType{Name: "test", Nullability: types.NullabilityRequired, TypeParameters: []types.UDTParameter{&types.StringUDTParam{StringVal: "L1"}}}},
		{"u!test<decimal<P,S>>", "u!test<decimal<P,S>>", "u!test", &types.ParameterizedUserDefinedType{Name: "test", Nullability: types.NullabilityRequired, TypeParameters: []types.UDTParameter{&types.DataTypeUDTParam{Type: &types.ParameterizedDecimalType{Precision: &variableIntP, Scale: &variableIntS, Nullability: types.NullabilityRequired}}}}},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := ParseType(tt.input)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, got.String())
			if tt.shortName != "" {
				assert.Equal(t, tt.shortName, got.ShortString())
			}
			assert.True(t, reflect.DeepEqual(got, tt.want), "ParseSubstraitType() = %v, want %v", got, tt.want)
		})
	}
}
