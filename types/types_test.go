// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v5/functions"
	. "github.com/substrait-io/substrait-go/v5/types"
	"github.com/substrait-io/substrait-go/v5/types/integer_parameters"
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
		{&IntervalDayType{Precision: 5}, "interval_day<5>", "iday"},
		{&UUIDType{Nullability: NullabilityNullable}, "uuid?", "uuid"},
		{&FixedBinaryType{Length: 10}, "fixedbinary<10>", "fbin"},
		{&FixedCharType{Length: 5}, "fixedchar<5>", "fchar"},
		{&VarCharType{Length: 15}, "varchar<15>", "vchar"},
		{&DecimalType{Scale: 2, Precision: 4}, "decimal<4,2>", "dec"},
		{&StructType{Nullability: NullabilityNullable, Types: []Type{
			&Int8Type{Nullability: NullabilityNullable},
			&DateType{Nullability: NullabilityRequired}, &FixedCharType{Length: 5}}},
			"struct?<i8?, date, fixedchar<5>>", "struct"},
		{&ListType{Type: &Int8Type{}}, "list<i8>", "list"},
		{&MapType{Key: &StringType{}, Value: &DecimalType{Precision: 10, Scale: 2}},
			"map<string, decimal<10,2>>", "map"},
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
				&UUIDType{Nullability: n},
				&FixedCharType{Nullability: n, Length: 25},
				&VarCharType{Nullability: n, Length: 35},
				&FixedBinaryType{Nullability: n, Length: 45},
				&IntervalDayType{Nullability: n, Precision: 5},
				&IntervalDayType{Nullability: n, Precision: 0},
				NewIntervalCompoundType().WithPrecision(PrecisionEMinus7Seconds).WithNullability(n),

				&DecimalType{Nullability: n, Precision: 34, Scale: 3},
				&PrecisionTimeType{Nullability: n, Precision: PrecisionEMinus4Seconds},
				&PrecisionTimestampType{Nullability: n, Precision: PrecisionEMinus4Seconds},
				&PrecisionTimestampTzType{PrecisionTimestampType: PrecisionTimestampType{Nullability: n, Precision: PrecisionEMinus5Seconds}},
				&MapType{Nullability: n, Key: &Int8Type{}, Value: &Int16Type{Nullability: n}},
				&ListType{Nullability: n, Type: &TimeType{Nullability: n}},
				&StructType{Nullability: n, Types: []Type{
					&TimeType{Nullability: n}, &TimestampType{Nullability: n},
					&TimestampTzType{Nullability: n}}},
				&UserDefinedType{TypeParameters: []TypeParam{&DataTypeParameter{Type: &Int32Type{}}}, Nullability: n},
			}

			for _, tt := range tests {
				t.Run(tt.String(), func(t *testing.T) {
					converted := TypeToProto(tt)
					convertedType := TypeFromProto(converted)
					assert.True(t, tt.Equals(convertedType))
				})
			}
		})
	}
}

func TestGetTypeNameToTypeMap(t *testing.T) {
	typeMap := GetTypeNameToTypeMap()
	tests := []struct {
		name            string
		typ             Type
		isSimple        bool
		isParameterized bool
		expError        bool
	}{
		{"boolean", &BooleanType{}, true, false, false},
		{"i8", &Int8Type{}, true, false, false},
		{"timestamp", &TimestampType{}, true, false, false},
		{"uuid", &UUIDType{}, true, false, false},
		{"fixedbinary", &FixedBinaryType{}, false, false, false},
		{"fixedchar", &FixedCharType{}, false, false, false},
		{"varchar", &VarCharType{}, false, false, false},
		{"decimal", &DecimalType{}, false, true, false},
		{"precision_time", &PrecisionTimeType{}, false, true, false},
		{"precision_timestamp", &PrecisionTimestampType{}, false, true, false},
		{"precision_timestamp_tz", &PrecisionTimestampTzType{}, false, true, false},

		{"unknown1", nil, true, false, true},
		{"unknown2", nil, false, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expError {
				assert.Nil(t, typeMap[tt.name])
				_, err := SimpleTypeNameToType(TypeName(tt.name))
				assert.Error(t, err)
				_, err = FixedTypeNameToType(TypeName(tt.name))
				assert.Error(t, err)
				return
			}
			assert.Equalf(t, tt.typ, typeMap[tt.name], "GetTypeNameToTypeMap()[%s] = %v, want %v", tt.name, typeMap[tt.name], tt.typ)
			if tt.isSimple {
				typ, err := SimpleTypeNameToType(TypeName(tt.name))
				assert.NoError(t, err)
				assert.Equalf(t, tt.typ, typ, "SimpleTypeNameToType(%s) = %v, want %v", tt.name, typ, tt.typ)

				parameters := typ.GetParameters()
				assert.Len(t, parameters, 0)
			} else if !tt.isParameterized {
				typ, err := FixedTypeNameToType(TypeName(tt.name))
				assert.NoError(t, err)
				assert.Equalf(t, tt.typ, typ, "FixedTypeNameToType(%s) = %v, want %v", tt.name, typ, tt.typ)
			}
		})
	}
}

func TestGetShortTypeName(t *testing.T) {
	tests := []struct {
		name     TypeName
		expShort string
	}{
		{"boolean", "bool"},
		{"i8", "i8"},
		{"timestamp", "ts"},
		{"uuid", "uuid"},
		{"binary", "vbin"},
		{"fixedbinary", "fbin"},
		{"fixedchar", "fchar"},
		{"varchar", "vchar"},
		{"string", "str"},
		{"decimal", "dec"},
		{"unknown", "unknown"},
		{"enum", "enum"},
	}
	for _, tt := range tests {
		t.Run(string(tt.name), func(t *testing.T) {
			assert.Equal(t, tt.expShort, GetShortTypeName(tt.name))
		})
	}
}

func TestFixedLenType_WithLength(t *testing.T) {
	tests := []struct {
		typeStr  string
		typ      FixedType
		length   int32
		expError bool
	}{
		{"fixedbinary", &FixedBinaryType{}, 10, false},
		{"fixedchar", &FixedCharType{}, 20, false},
		{"varchar", &VarCharType{}, 30, false},
	}
	for _, tt := range tests {
		t.Run(tt.typeStr, func(t *testing.T) {
			typ := tt.typ.WithLength(tt.length)
			if tt.expError {
				assert.Nil(t, typ)
				return
			}
			assert.Equal(t, fmt.Sprintf("%d", tt.length), typ.ParameterString())
			assert.Equal(t, tt.typeStr, typ.BaseString())
			parameters := typ.GetParameters()
			assert.Len(t, parameters, 1)
		})
	}
}

func TestMatchForBasicTypeResultMatch(t *testing.T) {
	for _, td := range []struct {
		name          string
		paramType     FuncDefArgType
		argOfSameType Type
	}{
		{"anyType", &AnyType{}, IntervalCompoundType{}.WithPrecision(PrecisionMilliSeconds)},
		{"binaryType", &BinaryType{}, &BinaryType{}},
		{"boolType", &BooleanType{}, &BooleanType{}},
		{"int8Type", &Int8Type{}, &Int8Type{}},
		{"int16Type", &Int16Type{}, &Int16Type{}},
		{"int32Type", &Int32Type{}, &Int32Type{}},
		{"int64Type", &Int64Type{}, &Int64Type{}},
		{"float32Type", &Float32Type{}, &Float32Type{}},
		{"float64Type", &Float64Type{}, &Float64Type{}},
		{"stringType", &StringType{}, &StringType{}},
		{"timestampType", &TimestampType{}, &TimestampType{}},
		{"dateType", &DateType{}, &DateType{}},
		{"timeType", &TimeType{}, &TimeType{}},
		{"timestampTzType", &TimestampTzType{}, &TimestampTzType{}},
		{"intervalYearType", &IntervalYearType{}, &IntervalYearType{}},
		{"uuidType", &UUIDType{}, &UUIDType{}},
		{"enumType", &EnumType{Options: []string{"A", "B", "C"}, Name: "ABC"}, &EnumType{Options: []string{"A", "B", "C"}, Name: "ABC"}},
	} {
		t.Run(td.name, func(t *testing.T) {
			// MatchWithNullability should match exact nullability and not match with different nullability
			assert.True(t, td.paramType.SetNullability(NullabilityNullable).MatchWithNullability(td.argOfSameType.WithNullability(NullabilityNullable)))
			assert.True(t, td.paramType.SetNullability(NullabilityRequired).MatchWithNullability(td.argOfSameType.WithNullability(NullabilityRequired)))
			assert.False(t, td.paramType.SetNullability(NullabilityRequired).MatchWithNullability(td.argOfSameType.WithNullability(NullabilityNullable)))
			assert.False(t, td.paramType.SetNullability(NullabilityRequired).MatchWithNullability(td.argOfSameType.WithNullability(NullabilityNullable)))

			// MatchWithoutNullability should match no matter what the nullability is
			assert.True(t, td.paramType.SetNullability(NullabilityNullable).MatchWithoutNullability(td.argOfSameType.WithNullability(NullabilityNullable)))
			assert.True(t, td.paramType.SetNullability(NullabilityRequired).MatchWithoutNullability(td.argOfSameType.WithNullability(NullabilityRequired)))
			assert.True(t, td.paramType.SetNullability(NullabilityRequired).MatchWithoutNullability(td.argOfSameType.WithNullability(NullabilityNullable)))
			assert.True(t, td.paramType.SetNullability(NullabilityRequired).MatchWithoutNullability(td.argOfSameType.WithNullability(NullabilityNullable)))
		})
	}
}

func TestMatchForBasicTypeResultMisMatch(t *testing.T) {
	for _, td := range []struct {
		name           string
		paramType      FuncDefArgType
		argOfOtherType Type
	}{
		{"binaryType", &BinaryType{}, &BooleanType{}},
		{"boolType", &BooleanType{}, &BinaryType{}},
		{"int8Type", &Int8Type{}, &BinaryType{}},
		{"int16Type", &Int16Type{}, &BinaryType{}},
		{"int32Type", &Int32Type{}, &BinaryType{}},
		{"int64Type", &Int64Type{}, &BinaryType{}},
		{"float32Type", &Float32Type{}, &BinaryType{}},
		{"float64Type", &Float64Type{}, &BinaryType{}},
		{"stringType", &StringType{}, &BinaryType{}},
		{"timestampType", &TimestampType{}, &BinaryType{}},
		{"dateType", &DateType{}, &BinaryType{}},
		{"timeType", &TimeType{}, &BinaryType{}},
		{"timestampTzType", &TimestampTzType{}, &BinaryType{}},
		{"intervalYearType", &IntervalYearType{}, &BinaryType{}},
		{"uuidType", &UUIDType{}, &BinaryType{}},
		{"enumType", &EnumType{Options: []string{"A", "B", "C"}, Name: "ABC"}, &BinaryType{}},
	} {
		t.Run(td.name, func(t *testing.T) {
			assert.False(t, td.paramType.MatchWithNullability(td.argOfOtherType))
			assert.False(t, td.paramType.MatchWithoutNullability(td.argOfOtherType))
		})
	}
}

func TestTypesHaveNoParameterizedParams(t *testing.T) {
	for _, td := range []struct {
		name           string
		paramType      FuncDefArgType
		argOfOtherType Type
	}{
		{"binaryType", &BinaryType{}, &BooleanType{}},
		{"boolType", &BooleanType{}, &BinaryType{}},
		{"int8Type", &Int8Type{}, &BinaryType{}},
		{"int16Type", &Int16Type{}, &BinaryType{}},
		{"int32Type", &Int32Type{}, &BinaryType{}},
		{"int64Type", &Int64Type{}, &BinaryType{}},
		{"float32Type", &Float32Type{}, &BinaryType{}},
		{"float64Type", &Float64Type{}, &BinaryType{}},
		{"stringType", &StringType{}, &BinaryType{}},
		{"timestampType", &TimestampType{}, &BinaryType{}},
		{"dateType", &DateType{}, &BinaryType{}},
		{"timeType", &TimeType{}, &BinaryType{}},
		{"timestampTzType", &TimestampTzType{}, &BinaryType{}},
		{"intervalYearType", &IntervalYearType{}, &BinaryType{}},
		{"uuidType", &UUIDType{}, &BinaryType{}},
		{"enumType", &EnumType{Options: []string{"A", "B", "C"}, Name: "ABC"}, &BinaryType{}},
	} {
		t.Run(td.name, func(t *testing.T) {
			parameters := td.paramType.GetParameterizedParams()
			assert.Nil(t, parameters)

			hasParameterizedParams := td.paramType.HasParameterizedParam()
			assert.False(t, hasParameterizedParams)
		})
	}
}

func TestMatchParameterizeConcreteTypeResultMatch(t *testing.T) {
	decimal382Concrete := &DecimalType{Precision: 38, Scale: 2}
	fixedCharLen5 := &FixedCharType{Length: 5}
	varCharLen5 := &VarCharType{Length: 5}
	fixedBinaryLen5 := &FixedBinaryType{Length: 5}
	intervalDayLen5 := &IntervalDayType{Precision: 5}

	concreteInt38 := integer_parameters.NewConcreteIntParam(38)
	concreteInt2 := integer_parameters.NewConcreteIntParam(2)
	concreteInt5 := integer_parameters.NewConcreteIntParam(5)
	for _, td := range []struct {
		name          string
		paramType     FuncDefArgType
		argOfSameType Type
		parameters    []interface{}
	}{
		{"decimalConcreteToDecimalConcrete", &ParameterizedDecimalType{Precision: concreteInt38, Scale: concreteInt2}, decimal382Concrete, []interface{}{int64(38), int64(2)}},
		{"fixedCharType", &ParameterizedFixedCharType{IntegerOption: concreteInt5}, fixedCharLen5, []interface{}{int64(5)}},
		{"varCharType", &ParameterizedVarCharType{IntegerOption: concreteInt5}, varCharLen5, []interface{}{int64(5)}},
		{"fixedBinaryType", &ParameterizedFixedBinaryType{IntegerOption: concreteInt5}, fixedBinaryLen5, []interface{}{int64(5)}},
		{"intervalDayType", &ParameterizedIntervalDayType{IntegerOption: concreteInt5}, intervalDayLen5, []interface{}{TimePrecision(5)}},
		{"listType", &ParameterizedListType{Type: &ParameterizedDecimalType{Precision: concreteInt38, Scale: concreteInt2}}, &ListType{Type: decimal382Concrete}, []interface{}{decimal382Concrete}},
		{"listTypeWithOtherListType", &ParameterizedListType{Type: &ParameterizedDecimalType{Precision: concreteInt38, Scale: concreteInt2}}, &ListType{Type: decimal382Concrete}, []interface{}{decimal382Concrete}},
		{"mapTypeKeyTypeDiffers", &ParameterizedMapType{Key: &Int32Type{}, Value: &BooleanType{}}, &MapType{Key: &Int32Type{}, Value: &BooleanType{}}, []interface{}{&Int32Type{}, &BooleanType{}}},
		{"mapTypeValueTypeDiffers", &ParameterizedMapType{Key: &Int32Type{}, Value: &BooleanType{}}, &MapType{Key: &Int32Type{}, Value: &BooleanType{}}, []interface{}{&Int32Type{}, &BooleanType{}}},
		{"structType", &ParameterizedStructType{Types: []FuncDefArgType{&BooleanType{}, &BooleanType{}}}, &StructType{Types: []Type{&BooleanType{}, &BooleanType{}}}, []interface{}{&BooleanType{}, &BooleanType{}}},
		{"userDefinedType", &ParameterizedUserDefinedType{TypeParameters: []UDTParameter{&DataTypeUDTParam{&Int32Type{}}}, Name: "udt"}, &UserDefinedType{TypeParameters: []TypeParam{&DataTypeParameter{Type: &Int32Type{}}}}, []interface{}{&DataTypeParameter{Type: &Int32Type{}}}},
		{"userDefinedType2", &ParameterizedUserDefinedType{TypeParameters: []UDTParameter{&DataTypeUDTParam{&Int32Type{}}, &DataTypeUDTParam{&Int32Type{}}}, Name: "udt"}, &UserDefinedType{TypeParameters: []TypeParam{&DataTypeParameter{Type: &Int32Type{}}, &DataTypeParameter{Type: &Int32Type{}}}}, []interface{}{&DataTypeParameter{Type: &Int32Type{}}, &DataTypeParameter{Type: &Int32Type{}}}},
		{"userDefinedType3", &ParameterizedUserDefinedType{TypeParameters: []UDTParameter{&StringUDTParam{StringVal: "L1"}}, Name: "udt"}, &UserDefinedType{TypeParameters: []TypeParam{StringParameter("L1")}}, []interface{}{StringParameter("L1")}},
		{"userDefinedType4", &ParameterizedUserDefinedType{TypeParameters: []UDTParameter{&IntegerUDTParam{Integer: 10}}, Name: "udt"}, &UserDefinedType{TypeParameters: []TypeParam{IntegerParameter(10)}}, []interface{}{IntegerParameter(10)}},
	} {
		t.Run(td.name, func(t *testing.T) {
			// MatchWithNullability should match exact nullability and not match with different nullability
			assert.True(t, td.paramType.SetNullability(NullabilityNullable).MatchWithNullability(td.argOfSameType.WithNullability(NullabilityNullable)))
			assert.True(t, td.paramType.SetNullability(NullabilityRequired).MatchWithNullability(td.argOfSameType.WithNullability(NullabilityRequired)))
			assert.False(t, td.paramType.SetNullability(NullabilityRequired).MatchWithNullability(td.argOfSameType.WithNullability(NullabilityNullable)))
			assert.False(t, td.paramType.SetNullability(NullabilityRequired).MatchWithNullability(td.argOfSameType.WithNullability(NullabilityNullable)))

			// MatchWithoutNullability should match no matter what the nullability is
			assert.True(t, td.paramType.SetNullability(NullabilityNullable).MatchWithoutNullability(td.argOfSameType.WithNullability(NullabilityNullable)))
			assert.True(t, td.paramType.SetNullability(NullabilityRequired).MatchWithoutNullability(td.argOfSameType.WithNullability(NullabilityRequired)))
			assert.True(t, td.paramType.SetNullability(NullabilityRequired).MatchWithoutNullability(td.argOfSameType.WithNullability(NullabilityNullable)))
			assert.True(t, td.paramType.SetNullability(NullabilityRequired).MatchWithoutNullability(td.argOfSameType.WithNullability(NullabilityNullable)))

			parameters := td.argOfSameType.GetParameters()
			assert.NotNil(t, parameters)
			assert.Equal(t, td.parameters, parameters)
		})
	}
}

func TestMatchParameterizeConcreteTypeResultMismatch(t *testing.T) {
	decimal380Concrete := &DecimalType{Precision: 38, Scale: 0}
	fixedCharLen6 := &FixedCharType{Length: 6}
	varCharLen6 := &VarCharType{Length: 6}
	fixedBinaryLen6 := &FixedBinaryType{Length: 6}
	intervalDayLen6 := &IntervalDayType{Precision: 6}

	concreteInt38 := integer_parameters.NewConcreteIntParam(38)
	concreteInt2 := integer_parameters.NewConcreteIntParam(2)
	concreteInt5 := integer_parameters.NewConcreteIntParam(5)
	enumOptions := []string{"A", "B", "C"}
	enumType := &EnumType{Options: enumOptions, Name: "ABC"}
	for _, td := range []struct {
		name           string
		paramType      FuncDefArgType
		argOfOtherType Type
	}{
		{"decimalConcreteToDecimalConcrete", &ParameterizedDecimalType{Precision: concreteInt38, Scale: concreteInt2}, decimal380Concrete},
		{"fixedCharType", &ParameterizedFixedCharType{IntegerOption: concreteInt5}, fixedCharLen6},
		{"varCharType", &ParameterizedVarCharType{IntegerOption: concreteInt5}, varCharLen6},
		{"fixedBinaryType", &ParameterizedFixedBinaryType{IntegerOption: concreteInt5}, fixedBinaryLen6},
		{"intervalDayType", &ParameterizedIntervalDayType{IntegerOption: concreteInt5}, intervalDayLen6},
		{"listType", &ParameterizedListType{Type: &ParameterizedDecimalType{Precision: concreteInt38, Scale: concreteInt2}}, &DecimalType{Precision: 38, Scale: 0}},
		{"listTypeWithOtherListType", &ParameterizedListType{Type: &ParameterizedDecimalType{Precision: concreteInt38, Scale: concreteInt2}}, &ListType{Type: &BooleanType{}}},
		{"mapTypeKeyTypeDiffers", &ParameterizedMapType{Key: &Int32Type{}, Value: &BooleanType{}}, &MapType{Key: &Int64Type{}, Value: &BooleanType{}}},
		{"mapTypeValueTypeDiffers", &ParameterizedMapType{Key: &Int32Type{}, Value: &BooleanType{}}, &MapType{Key: &Int32Type{}, Value: &DecimalType{Precision: 38, Scale: 2}}},
		{"structType Mismatch DifferInType", &ParameterizedStructType{Types: []FuncDefArgType{&BooleanType{}, &BooleanType{}}}, &MapType{Key: &Int64Type{}, Value: &BooleanType{}}},
		{"structType Mismatch DifferInEmbedTypeLen", &ParameterizedStructType{Types: []FuncDefArgType{&BooleanType{}, &BooleanType{}}}, &StructType{Types: []Type{&BooleanType{}}}},
		{"structType Mismatch InFirstEmbedType", &ParameterizedStructType{Types: []FuncDefArgType{&BooleanType{}, &BooleanType{}}}, &StructType{Types: []Type{&DecimalType{Precision: 38, Scale: 2}, &BooleanType{}}}},
		{"structType Mismatch InLastEmbedType", &ParameterizedStructType{Types: []FuncDefArgType{&BooleanType{}, &BooleanType{}}}, &StructType{Types: []Type{&BooleanType{}, &DecimalType{Precision: 38, Scale: 2}}}},
		{"userDefinedType", &ParameterizedUserDefinedType{TypeParameters: []UDTParameter{&DataTypeUDTParam{&Int32Type{}}}, Name: "udt"}, &BooleanType{}},
		{"userDefinedType", &ParameterizedUserDefinedType{TypeParameters: []UDTParameter{&DataTypeUDTParam{&Int32Type{}}, &DataTypeUDTParam{&Int32Type{}}}, Name: "udt"}, &StructType{Types: []Type{&BooleanType{}, &Int32Type{}}}},
		{"userDefinedType", &ParameterizedUserDefinedType{TypeParameters: []UDTParameter{&DataTypeUDTParam{&Int32Type{}}}, Name: "udt"}, &UserDefinedType{TypeParameters: []TypeParam{&DataTypeParameter{Type: &Int64Type{}}}}},
		{"userDefinedType", &ParameterizedUserDefinedType{TypeParameters: []UDTParameter{&DataTypeUDTParam{&Int32Type{}}}, Name: "udt"}, &UserDefinedType{TypeParameters: []TypeParam{IntegerParameter(10)}}},
		{"userDefinedType", &ParameterizedUserDefinedType{TypeParameters: []UDTParameter{&DataTypeUDTParam{&Int32Type{}}, &DataTypeUDTParam{&Int32Type{}}}, Name: "udt"}, &UserDefinedType{TypeParameters: []TypeParam{&DataTypeParameter{Type: &Int64Type{}}}}},
		{"userDefinedType", &ParameterizedUserDefinedType{TypeParameters: []UDTParameter{&DataTypeUDTParam{&Int32Type{}}, &DataTypeUDTParam{&Int32Type{}}}, Name: "udt"}, &UserDefinedType{TypeParameters: []TypeParam{IntegerParameter(10)}}},
		{"userDefinedType", &ParameterizedUserDefinedType{TypeParameters: []UDTParameter{&IntegerUDTParam{Integer: 10}}, Name: "udt"}, &UserDefinedType{TypeParameters: []TypeParam{&DataTypeParameter{Type: &Int64Type{}}}}},
		{"userDefinedType", &ParameterizedUserDefinedType{TypeParameters: []UDTParameter{&IntegerUDTParam{Integer: 10}}, Name: "udt"}, &UserDefinedType{TypeParameters: []TypeParam{IntegerParameter(11)}}},
		{"userDefinedType", &ParameterizedUserDefinedType{TypeParameters: []UDTParameter{&StringUDTParam{StringVal: "L1"}}, Name: "udt"}, &UserDefinedType{TypeParameters: []TypeParam{IntegerParameter(11)}}},
		{"enumType", &EnumType{Name: "ABCEnum", Options: enumOptions}, enumType},
	} {
		t.Run(td.name, func(t *testing.T) {
			assert.False(t, td.paramType.MatchWithNullability(td.argOfOtherType))
			assert.False(t, td.paramType.MatchWithoutNullability(td.argOfOtherType))
		})
	}
}

func TestMatchParameterizedNonNestedTypeResultMatch(t *testing.T) {
	intParamLen := integer_parameters.NewVariableIntParam("L1")
	argFixedCharLen := &FixedCharType{Length: 5}
	paramFixedChar := &ParameterizedFixedCharType{IntegerOption: intParamLen}
	argVarChar := &VarCharType{Length: 5}
	paramVarChar := &ParameterizedVarCharType{IntegerOption: intParamLen}
	argFixedBinaryLen := &FixedBinaryType{Length: 5}
	paramFixedBinary := &ParameterizedFixedBinaryType{IntegerOption: intParamLen}
	argPrecisionTime := &PrecisionTimeType{Precision: PrecisionEMinus4Seconds}
	paramPrecisionTime := &ParameterizedPrecisionTimeType{IntegerOption: intParamLen}
	argPrecisionTimeStamp := &PrecisionTimestampType{Precision: PrecisionEMinus5Seconds}
	paramPrecisionTimeStamp := &ParameterizedPrecisionTimestampType{IntegerOption: intParamLen}
	argPrecisionTimeStampTzType := &PrecisionTimestampTzType{PrecisionTimestampType: PrecisionTimestampType{Precision: PrecisionMicroSeconds}}
	paramPrecisionTimeStampTz := &ParameterizedPrecisionTimestampTzType{IntegerOption: intParamLen}
	argDecimalType := &DecimalType{Precision: 38, Scale: 2}
	paramDecimalType := &ParameterizedDecimalType{Precision: integer_parameters.NewVariableIntParam("P"), Scale: integer_parameters.NewVariableIntParam("S")}
	argIntervalDayType := &IntervalDayType{Precision: 5}
	paramIntervalDayType := &ParameterizedIntervalDayType{IntegerOption: intParamLen}

	enumOptions := []string{"A", "B", "C"}
	enumType := &EnumType{Name: "ABCEnum", Options: enumOptions}

	for _, td := range []struct {
		name      string
		paramType FuncDefArgType
		argType   Type
	}{
		{"fixedChar", paramFixedChar, argFixedCharLen},
		{"varChar", paramVarChar, argVarChar},
		{"fixBinary", paramFixedBinary, argFixedBinaryLen},
		{"intervalDay", paramIntervalDayType, argIntervalDayType},
		{"precisionTime", paramPrecisionTime, argPrecisionTime},
		{"precisionTimestamp", paramPrecisionTimeStamp, argPrecisionTimeStamp},
		{"precisionTimestampTz", paramPrecisionTimeStampTz, argPrecisionTimeStampTzType},
		{"decimalType", paramDecimalType, argDecimalType},
		{"enumType", &EnumType{Name: "ABCEnum", Options: enumOptions}, enumType},
	} {
		t.Run(td.name, func(t *testing.T) {
			// MatchWithNullability should match exact nullability and not match with different nullability
			assert.True(t, td.paramType.SetNullability(NullabilityNullable).MatchWithNullability(td.argType.WithNullability(NullabilityNullable)))
			assert.True(t, td.paramType.SetNullability(NullabilityRequired).MatchWithNullability(td.argType.WithNullability(NullabilityRequired)))
			assert.False(t, td.paramType.SetNullability(NullabilityRequired).MatchWithNullability(td.argType.WithNullability(NullabilityNullable)))
			assert.False(t, td.paramType.SetNullability(NullabilityRequired).MatchWithNullability(td.argType.WithNullability(NullabilityNullable)))

			// MatchWithoutNullability should match no matter what the nullability is
			assert.True(t, td.paramType.SetNullability(NullabilityNullable).MatchWithoutNullability(td.argType.WithNullability(NullabilityNullable)))
			assert.True(t, td.paramType.SetNullability(NullabilityRequired).MatchWithoutNullability(td.argType.WithNullability(NullabilityRequired)))
			assert.True(t, td.paramType.SetNullability(NullabilityRequired).MatchWithoutNullability(td.argType.WithNullability(NullabilityNullable)))
			assert.True(t, td.paramType.SetNullability(NullabilityRequired).MatchWithoutNullability(td.argType.WithNullability(NullabilityNullable)))
		})
	}
}

func TestMatchParameterizedNestedTypeResultMatch(t *testing.T) {
	variableIntP := integer_parameters.NewVariableIntParam("P")
	variableIntS := integer_parameters.NewVariableIntParam("S")

	argDecimalType := &DecimalType{Precision: 38, Scale: 2, Nullability: NullabilityNullable}
	paramDecimalType := markNullable(&ParameterizedDecimalType{Precision: variableIntP, Scale: variableIntS})
	listTypeAsNestedArg := &ListType{Type: argDecimalType, Nullability: NullabilityNullable}
	listTypeAsNestedParam := &ParameterizedListType{Type: paramDecimalType, Nullability: NullabilityNullable}

	argListType := &ListType{Type: argDecimalType}
	paramListType := &ParameterizedListType{Type: paramDecimalType}
	argMapType := &MapType{Key: argDecimalType, Value: listTypeAsNestedArg}
	paramMapType := &ParameterizedMapType{Key: paramDecimalType, Value: listTypeAsNestedParam}
	argStructType := &StructType{Types: []Type{argDecimalType, listTypeAsNestedArg}}
	paramStructType := &ParameterizedStructType{Types: []FuncDefArgType{paramDecimalType, listTypeAsNestedParam}}

	for _, td := range []struct {
		name      string
		paramType FuncDefArgType
		argType   Type
	}{
		{"list Type", paramListType, argListType},
		{"map Type", paramMapType, argMapType},
		{"struct Type", paramStructType, argStructType},
	} {
		t.Run(td.name, func(t *testing.T) {
			// MatchWithNullability should match exact nullability and not match with different nullability
			assert.True(t, td.paramType.SetNullability(NullabilityNullable).MatchWithNullability(td.argType.WithNullability(NullabilityNullable)))
			assert.True(t, td.paramType.SetNullability(NullabilityRequired).MatchWithNullability(td.argType.WithNullability(NullabilityRequired)))
			assert.False(t, td.paramType.SetNullability(NullabilityRequired).MatchWithNullability(td.argType.WithNullability(NullabilityNullable)))
			assert.False(t, td.paramType.SetNullability(NullabilityRequired).MatchWithNullability(td.argType.WithNullability(NullabilityNullable)))

			// MatchWithoutNullability should match no matter what the nullability is
			assert.True(t, td.paramType.SetNullability(NullabilityNullable).MatchWithoutNullability(td.argType.WithNullability(NullabilityNullable)))
			assert.True(t, td.paramType.SetNullability(NullabilityRequired).MatchWithoutNullability(td.argType.WithNullability(NullabilityRequired)))
			assert.True(t, td.paramType.SetNullability(NullabilityRequired).MatchWithoutNullability(td.argType.WithNullability(NullabilityNullable)))
			assert.True(t, td.paramType.SetNullability(NullabilityRequired).MatchWithoutNullability(td.argType.WithNullability(NullabilityNullable)))
		})
	}
}

func markNullable(t FuncDefArgType) FuncDefArgType {
	return t.SetNullability(NullabilityNullable)
}

func TestGetTimeValueByPrecision(t *testing.T) {
	timeStr := "2021-08-10T15:01:05.123456789Z"
	tests := []struct {
		name      string
		precision TimePrecision
		want      int64
	}{
		{"PrecisionSeconds", PrecisionSeconds, 1628607665},
		{"PrecisionDeciSeconds", PrecisionDeciSeconds, 16286076651},
		{"PrecisionCentiSeconds", PrecisionCentiSeconds, 162860766512},
		{"PrecisionMilliSeconds", PrecisionMilliSeconds, 1628607665123},
		{"PrecisionEMinus4Seconds", PrecisionEMinus4Seconds, 16286076651234},
		{"PrecisionEMinus5Seconds", PrecisionEMinus5Seconds, 162860766512345},
		{"PrecisionMicroSeconds", PrecisionMicroSeconds, 1628607665123456},
		{"PrecisionEMinus7Seconds", PrecisionEMinus7Seconds, 16286076651234567},
		{"PrecisionEMinus8Seconds", PrecisionEMinus8Seconds, 162860766512345678},
		{"PrecisionNanoSeconds", PrecisionNanoSeconds, 1628607665123456789},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm, _ := time.Parse(time.RFC3339Nano, timeStr)
			assert.Equalf(t, tt.want, GetTimeValueByPrecision(tm, tt.precision), "GetTimeValueByPrecision(%v, %v)", timeStr, tt.precision)
		})
	}
}

func TestGetSupportedTypes(t *testing.T) {
	dialect, err := functions.LoadDialect("test_dialect",
		strings.NewReader(`
name: test_dialect
type: sql
dependencies:
  arithmetic:
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml
supported_types:
  fp64:
    sql_type_name: float
  bool:
    sql_type_name: boolean
  varchar:
    sql_type_name: varchar
  date:
    sql_type_name: date
  time:
    sql_type_name: time
  pts:
    sql_type_name: timestamp
  ptstz:
    sql_type_name: timestamptz
  dec:
    sql_type_name: numeric
scalar_functions:
  - name: arithmetic.add
    local_name: '+'
    infix: true
    required_options:
      overflow: SILENT
      rounding: TIE_TO_EVEN
    supported_kernels:
      - fp64_fp64
`))
	require.NoError(t, err)
	typeRegistry, err := dialect.GetLocalTypeRegistry()
	require.NoError(t, err)
	st := typeRegistry.GetSupportedTypes()
	assert.Len(t, st, 8)
}
