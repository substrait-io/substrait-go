package parser

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/goccy/go-yaml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v5/types"
	"github.com/substrait-io/substrait-go/v5/types/integer_parameters"
)

func TestParseSimpleTypes(t *testing.T) {
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

		{"Boolean", "boolean?", "bool", &types.BooleanType{Nullability: types.NullabilityRequired}},
		{"I8", "i8", "i8", &types.Int8Type{Nullability: types.NullabilityRequired}},
		{"I16", "i16", "i16", &types.Int16Type{Nullability: types.NullabilityRequired}},
		{"I32", "i32", "i32", &types.Int32Type{Nullability: types.NullabilityRequired}},
		{"I64", "i64", "i64", &types.Int64Type{Nullability: types.NullabilityRequired}},
		{"FP32", "fp32", "fp32", &types.Float32Type{Nullability: types.NullabilityRequired}},
		{"FP64", "fp64", "fp64", &types.Float64Type{Nullability: types.NullabilityRequired}},
		{"STRING", "string", "str", &types.StringType{Nullability: types.NullabilityRequired}},
		{"BINARY", "binary", "bin", &types.BinaryType{Nullability: types.NullabilityRequired}},
		{"TIMESTAMP", "timestamp", "ts", &types.TimestampType{Nullability: types.NullabilityRequired}},
		{"TIMESTAMP_TZ", "timestamp_tz", "tstz", &types.TimestampTzType{Nullability: types.NullabilityRequired}},
		{"DATE", "date", "date", &types.DateType{Nullability: types.NullabilityRequired}},
		{"TIME", "time", "time", &types.TimeType{Nullability: types.NullabilityRequired}},
		{"UUID", "uuid", "uuid", &types.UUIDType{Nullability: types.NullabilityRequired}},
		{"interval_year", "interval_year", "iyear", &types.IntervalYearType{Nullability: types.NullabilityRequired}},
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
		{"i8?", "i8?", "i8", &types.Int8Type{Nullability: types.NullabilityNullable}},
		{"i16?", "i16?", "i16", &types.Int16Type{Nullability: types.NullabilityNullable}},
		{"string?", "string?", "str", &types.StringType{Nullability: types.NullabilityNullable}},
		{"boolean?", "boolean?", "bool", &types.BooleanType{Nullability: types.NullabilityNullable}},

		{"fixedchar<5>", "fixedchar<5>", "fchar", &types.ParameterizedFixedCharType{IntegerOption: &concreteInt5, Nullability: types.NullabilityRequired}},
		{"decimal<10,5>", "decimal<10,5>", "dec", &types.ParameterizedDecimalType{Precision: &concreteInt10, Scale: &concreteInt5, Nullability: types.NullabilityRequired}},
		{"list<decimal<10,5>>", "list<decimal<10,5>>", "list", &types.ParameterizedListType{Type: &types.ParameterizedDecimalType{Precision: &concreteInt10, Scale: &concreteInt5, Nullability: types.NullabilityRequired}, Nullability: types.NullabilityRequired}},
		{"list?<decimal?<10,5>>", "list?<decimal?<10,5>>", "list", &types.ParameterizedListType{Type: &types.ParameterizedDecimalType{Precision: &concreteInt10, Scale: &concreteInt5, Nullability: types.NullabilityNullable}, Nullability: types.NullabilityNullable}},
		{"struct<i16?,i32>", "struct<i16?, i32>", "struct", &types.ParameterizedStructType{Types: []types.FuncDefArgType{&types.Int16Type{Nullability: types.NullabilityNullable}, &types.Int32Type{Nullability: types.NullabilityRequired}}, Nullability: types.NullabilityRequired}},
		{"map<boolean?,struct?<i16?,i32?,i64?>>", "map<boolean?, struct?<i16?, i32?, i64?>>", "map", &types.ParameterizedMapType{Key: &types.BooleanType{Nullability: types.NullabilityNullable}, Value: &types.ParameterizedStructType{Types: []types.FuncDefArgType{&types.Int16Type{Nullability: types.NullabilityNullable}, &types.Int32Type{Nullability: types.NullabilityNullable}, &types.Int64Type{Nullability: types.NullabilityNullable}}, Nullability: types.NullabilityNullable}, Nullability: types.NullabilityRequired}},
		{"map?<boolean?,struct?<i16?,i32?,i64?>>", "map?<boolean?, struct?<i16?, i32?, i64?>>", "map", &types.ParameterizedMapType{Key: &types.BooleanType{Nullability: types.NullabilityNullable}, Value: &types.ParameterizedStructType{Types: []types.FuncDefArgType{&types.Int16Type{Nullability: types.NullabilityNullable}, &types.Int32Type{Nullability: types.NullabilityNullable}, &types.Int64Type{Nullability: types.NullabilityNullable}}, Nullability: types.NullabilityNullable}, Nullability: types.NullabilityNullable}},
		{"precision_time<5>", "precision_time<5>", "pt", &types.ParameterizedPrecisionTimeType{IntegerOption: &concreteInt5, Nullability: types.NullabilityRequired}},
		{"precision_time?<5>", "precision_time?<5>", "pt", &types.ParameterizedPrecisionTimeType{IntegerOption: &concreteInt5, Nullability: types.NullabilityNullable}},
		{"precision_timestamp<5>", "precision_timestamp<5>", "pts", &types.ParameterizedPrecisionTimestampType{IntegerOption: &concreteInt5, Nullability: types.NullabilityRequired}},
		{"precision_timestamp?<5>", "precision_timestamp?<5>", "pts", &types.ParameterizedPrecisionTimestampType{IntegerOption: &concreteInt5, Nullability: types.NullabilityNullable}},
		{"precision_timestamp_tz<5>", "precision_timestamp_tz<5>", "ptstz", &types.ParameterizedPrecisionTimestampTzType{IntegerOption: &concreteInt5, Nullability: types.NullabilityRequired}},
		{"precision_timestamp_tz?<5>", "precision_timestamp_tz?<5>", "ptstz", &types.ParameterizedPrecisionTimestampTzType{IntegerOption: &concreteInt5, Nullability: types.NullabilityNullable}},
		{"interval_day<5>", "interval_day<5>", "iday", &types.ParameterizedIntervalDayType{IntegerOption: &concreteInt5, Nullability: types.NullabilityRequired}},
		{"interval_day?<5>", "interval_day?<5>", "iday", &types.ParameterizedIntervalDayType{IntegerOption: &concreteInt5, Nullability: types.NullabilityNullable}},

		{"varchar<L1>", "varchar<L1>", "vchar", &types.ParameterizedVarCharType{IntegerOption: &variableIntL1, Nullability: types.NullabilityRequired}},
		{"varchar?<L1>", "varchar?<L1>", "vchar", &types.ParameterizedVarCharType{IntegerOption: &variableIntL1, Nullability: types.NullabilityNullable}},
		{"fixedchar<L1>", "fixedchar<L1>", "fchar", &types.ParameterizedFixedCharType{IntegerOption: &variableIntL1, Nullability: types.NullabilityRequired}},
		{"fixedchar?<L1>", "fixedchar?<L1>", "fchar", &types.ParameterizedFixedCharType{IntegerOption: &variableIntL1, Nullability: types.NullabilityNullable}},
		{"fixedbinary<L1>", "fixedbinary<L1>", "fbin", &types.ParameterizedFixedBinaryType{IntegerOption: &variableIntL1, Nullability: types.NullabilityRequired}},
		{"fixedbinary?<L1>", "fixedbinary?<L1>", "fbin", &types.ParameterizedFixedBinaryType{IntegerOption: &variableIntL1, Nullability: types.NullabilityNullable}},
		{"precision_time<L1>", "precision_time<L1>", "pt", &types.ParameterizedPrecisionTimeType{IntegerOption: &variableIntL1, Nullability: types.NullabilityRequired}},
		{"precision_timestamp<L1>", "precision_timestamp<L1>", "pts", &types.ParameterizedPrecisionTimestampType{IntegerOption: &variableIntL1, Nullability: types.NullabilityRequired}},
		{"precision_timestamp_tz<L1>", "precision_timestamp_tz<L1>", "ptstz", &types.ParameterizedPrecisionTimestampTzType{IntegerOption: &variableIntL1, Nullability: types.NullabilityRequired}},
		{"decimal<P,S>", "decimal<P,S>", "dec", &types.ParameterizedDecimalType{Precision: &variableIntP, Scale: &variableIntS, Nullability: types.NullabilityRequired}},
		{"decimal<38,S>", "decimal<38,S>", "dec", &types.ParameterizedDecimalType{Precision: &concreteInt38, Scale: &variableIntS, Nullability: types.NullabilityRequired}},
		{"any", "any", "any", &types.AnyType{Name: "any", Nullability: types.NullabilityRequired}},
		{"any1?", "any1?", "any", &types.AnyType{Name: "any1", Nullability: types.NullabilityNullable}},
		{"list<decimal<P,S>>", "list<decimal<P,S>>", "list", &types.ParameterizedListType{Type: &types.ParameterizedDecimalType{Precision: &variableIntP, Scale: &variableIntS, Nullability: types.NullabilityRequired}, Nullability: types.NullabilityRequired}},
		{"struct<list?<decimal<P,S>>, i16>", "struct<list?<decimal<P,S>>, i16>", "struct", &types.ParameterizedStructType{Types: []types.FuncDefArgType{&types.ParameterizedListType{Type: &types.ParameterizedDecimalType{Precision: &variableIntP, Scale: &variableIntS, Nullability: types.NullabilityRequired}, Nullability: types.NullabilityNullable}, &types.Int16Type{Nullability: types.NullabilityRequired}}, Nullability: types.NullabilityRequired}},
		{"map<decimal<P,S>, i16>", "map<decimal<P,S>, i16>", "map", &types.ParameterizedMapType{Key: &types.ParameterizedDecimalType{Precision: &variableIntP, Scale: &variableIntS, Nullability: types.NullabilityRequired}, Value: &types.Int16Type{Nullability: types.NullabilityRequired}, Nullability: types.NullabilityRequired}},
		{"precision_timestamp_tz?<L1>", "precision_timestamp_tz?<L1>", "ptstz", &types.ParameterizedPrecisionTimestampTzType{IntegerOption: &variableIntL1, Nullability: types.NullabilityNullable}},
		{"u!test", "u!test", "u!test", &types.ParameterizedUserDefinedType{Name: "test", Nullability: types.NullabilityRequired}},
		{"u!test?", "u!test?", "u!test", &types.ParameterizedUserDefinedType{Name: "test", Nullability: types.NullabilityNullable}},
		{"u!test<10>", "u!test<10>", "u!test", &types.ParameterizedUserDefinedType{Name: "test", Nullability: types.NullabilityRequired, TypeParameters: []types.UDTParameter{&types.IntegerUDTParam{Integer: 10}}}},
		{"u!test?<10>", "u!test?<10>", "u!test", &types.ParameterizedUserDefinedType{Name: "test", Nullability: types.NullabilityNullable, TypeParameters: []types.UDTParameter{&types.IntegerUDTParam{Integer: 10}}}},
		{"u!test<L1>", "u!test<L1>", "u!test", &types.ParameterizedUserDefinedType{Name: "test", Nullability: types.NullabilityRequired, TypeParameters: []types.UDTParameter{&types.StringUDTParam{StringVal: "L1"}}}},
		{"u!test<decimal<P,S>>", "u!test<decimal<P,S>>", "u!test", &types.ParameterizedUserDefinedType{Name: "test", Nullability: types.NullabilityRequired, TypeParameters: []types.UDTParameter{&types.DataTypeUDTParam{Type: &types.ParameterizedDecimalType{Precision: &variableIntP, Scale: &variableIntS, Nullability: types.NullabilityRequired}}}}},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := ParseType(tt.input)
			require.NoError(t, err)
			require.NotNil(t, got)
			assert.Equal(t, tt.expected, got.String())
			if tt.shortName != "" {
				assert.Equal(t, tt.shortName, got.ShortString())
			}
			assert.True(t, reflect.DeepEqual(got, tt.want), "ParseSubstraitType() = %v, want %v", got, tt.want)
		})
	}
}

func TestParseErrors(t *testing.T) {
	tests := []struct {
		input string
	}{
		{"i"},
		{"i1"},
		{"decimal<38>"},
		{"decimal(38,10)"},
		{"map<i16, 42i>"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			_, err := ParseType(tt.input)
			assert.Error(t, err)
		})
	}
}

func TestTypeExpression_UnmarshalYAML(t *testing.T) {
	tests := []struct {
		input    string
		expected types.FuncDefArgType
		hasError bool
	}{
		{"i16", &types.Int16Type{Nullability: types.NullabilityRequired}, false},
		{"i16?", &types.Int16Type{Nullability: types.NullabilityNullable}, false},
		{"invalid", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			var te TypeExpression
			err := yaml.Unmarshal([]byte(tt.input), &te)
			if tt.hasError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, te.ValueType)
		})
	}
}

func TestTypeExpression_UnmarshalYAML1(t1 *testing.T) {
	concreteInt38 := integer_parameters.ConcreteIntParam(38)
	variableIntS := integer_parameters.VariableIntParam("S")
	tests := []struct {
		name      string
		ValueType types.FuncDefArgType
		hasError  bool
	}{
		{"i16", &types.Int16Type{Nullability: types.NullabilityRequired}, false},
		{"i16?", &types.Int16Type{Nullability: types.NullabilityNullable}, false},
		{"decimal<38,S>", &types.ParameterizedDecimalType{Precision: &concreteInt38, Scale: &variableIntS, Nullability: types.NullabilityRequired}, false},
		{"invalid", nil, true},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TypeExpression{}
			dummyFunc := func(arg interface{}) error {
				argPtr := reflect.ValueOf(arg)
				if argPtr.Kind() == reflect.Ptr && !argPtr.IsNil() {
					argPtr.Elem().Set(reflect.ValueOf(tt.name))
					return nil
				}
				return fmt.Errorf("expected pointer argument")
			}
			err := t.UnmarshalYAML(dummyFunc)
			if tt.hasError {
				require.Error(t1, err)
				return
			}
			require.NoError(t1, err)
			assert.Equal(t1, tt.ValueType, t.ValueType)

			out, err := t.MarshalYAML()
			require.NoError(t1, err)
			assert.Equal(t1, tt.name, out)
		})
	}
}

func TestParseOutputDerivation(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		finalType types.FuncDefArgType
	}{
		{"decimal", "P1 = 38\ndecimal<P1,5>", &types.ParameterizedDecimalType{Precision: integer_parameters.NewVariableIntParam("P1"), Scale: integer_parameters.NewConcreteIntParam(5), Nullability: types.NullabilityRequired}},
		{"+varchar", "x = (1 + 2)\nvarchar<x>", &types.ParameterizedVarCharType{IntegerOption: integer_parameters.NewVariableIntParam("x"), Nullability: types.NullabilityRequired}},
		{"+decimal", "P1 = (30 / 2)\nS1 = 3\ndecimal?<P1,S1>", &types.ParameterizedDecimalType{Precision: integer_parameters.NewVariableIntParam("P1"), Scale: integer_parameters.NewVariableIntParam("S1"), Nullability: types.NullabilityNullable}},
		{"-varchar", "L1 = 9\nL2 = 4\nL3 = (L1 - L2)\nvarchar<L3>", &types.ParameterizedVarCharType{IntegerOption: integer_parameters.NewVariableIntParam("L3"), Nullability: types.NullabilityRequired}},
		{"*varchar", "l2 = 5\nl3 = 6\nl4 = (l2 * l3)\nvarchar<l4>", &types.ParameterizedVarCharType{IntegerOption: integer_parameters.NewVariableIntParam("l4"), Nullability: types.NullabilityRequired}},
		{"max", "x = max(1, 2)\nvarchar<x>", &types.ParameterizedVarCharType{IntegerOption: integer_parameters.NewVariableIntParam("x"), Nullability: types.NullabilityRequired}},
		{"min", "L1 = min(3, 4)\nvarchar<L1>", &types.ParameterizedVarCharType{IntegerOption: integer_parameters.NewVariableIntParam("L1"), Nullability: types.NullabilityRequired}},
		{"abs", "l2 = abs(-5)\nvarchar?<l2>", &types.ParameterizedVarCharType{IntegerOption: integer_parameters.NewVariableIntParam("l2"), Nullability: types.NullabilityNullable}},
		{"if", "x = 1\ny = 2\nz = if !(x > y) then x else y\nvarchar<z>", &types.ParameterizedVarCharType{IntegerOption: integer_parameters.NewVariableIntParam("z"), Nullability: types.NullabilityRequired}},
		{"?ternary", "x = 1\ny = 2\nz = ((x < y) and (x > y)) ? (x * 3) : (y * 4)\nvarchar<z>", &types.ParameterizedVarCharType{IntegerOption: integer_parameters.NewVariableIntParam("z"), Nullability: types.NullabilityRequired}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseType(tt.input)
			require.NoError(t, err)
			derivation, ok := got.(*types.OutputDerivation)
			require.True(t, ok)
			assert.Equal(t, tt.finalType, derivation.FinalType)
			assert.Equal(t, tt.input, got.String())
		})
	}
}
