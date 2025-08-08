package parser

import (
	"embed"
	"fmt"
	"io/fs"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait"
	"github.com/substrait-io/substrait-go/v5/expr"
	"github.com/substrait-io/substrait-go/v5/extensions"
	"github.com/substrait-io/substrait-go/v5/functions"
	"github.com/substrait-io/substrait-go/v5/literal"
	"github.com/substrait-io/substrait-go/v5/types"
)

func makeHeader(version, include string) string {
	return fmt.Sprintf("### SUBSTRAIT_SCALAR_TEST: %s\n### SUBSTRAIT_INCLUDE: '%s'\n\n", version, include)
}

func makeAggregateTestHeader(version, include string) string {
	return fmt.Sprintf("### SUBSTRAIT_AGGREGATE_TEST: %s\n### SUBSTRAIT_INCLUDE: '%s'\n\n", version, include)
}

func TestParseBasicExample(t *testing.T) {
	header := makeHeader("v1.0", "/extensions/functions_arithmetic.yaml")
	tests := `#  'Basic examples without any special cases'
add(120::i8, 5::i8) = 125::i8
add(100::i16, 100::i16) = 200::i16
add(1::i8?, 2::i8?) = 3::i8?

# Overflow examples demonstrating overflow behavior  
add(120::i8, 10::i8) [overflow:ERROR] = <!ERROR>
`

	testStrings := []string{
		"add(120::i8, 5::i8) = 125::i8",
		"add(100::i16, 100::i16) = 200::i16",
		"add(1::i8?, 2::i8?) = 3::i8?",
		"add(120::i8, 10::i8) [overflow:ERROR] = <!ERROR>",
	}
	testFile, err := ParseTestCasesFromString(header + tests)
	require.NoError(t, err)
	assert.Len(t, testFile.TestCases, 4)

	arithURI := "https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml"
	ids := []string{"add:i8_i8", "add:i16_i16", "add:i8_i8", "add:i8_i8"}
	argTypes := [][]types.Type{
		{&types.Int8Type{Nullability: types.NullabilityRequired}, &types.Int8Type{Nullability: types.NullabilityRequired}},
		{&types.Int16Type{Nullability: types.NullabilityRequired}, &types.Int16Type{Nullability: types.NullabilityRequired}},
		{&types.Int8Type{Nullability: types.NullabilityNullable}, &types.Int8Type{Nullability: types.NullabilityNullable}},
		{&types.Int8Type{Nullability: types.NullabilityRequired}, &types.Int8Type{Nullability: types.NullabilityRequired}},
	}
	reg, funcRegistry := functions.NewExtensionAndFunctionRegistries(extensions.GetDefaultCollectionWithNoError())
	basicGroupDesc := "'Basic examples without any special cases'"
	overflowGroupDesc := "Overflow examples demonstrating overflow behavior"
	groupDescs := []string{basicGroupDesc, basicGroupDesc, basicGroupDesc, overflowGroupDesc}
	for i, tc := range testFile.TestCases {
		assert.Equal(t, extensions.ID{URI: arithURI, Name: ids[i]}, tc.ID())
		scalarFunc, err1 := tc.GetScalarFunctionInvocation(&reg, funcRegistry)
		require.NoError(t, err1)
		assert.Equal(t, tc.FuncName, scalarFunc.Name())
		require.Equal(t, 2, scalarFunc.NArgs())
		assert.Equal(t, tc.Args[0].Value, scalarFunc.Arg(0))
		assert.Equal(t, tc.Args[1].Value, scalarFunc.Arg(1))
		assert.Equal(t, argTypes[i], tc.GetArgTypes())
		assert.Equal(t, ids[i], tc.CompoundFunctionName())
		assert.Equal(t, groupDescs[i], tc.GroupDesc)
		assert.Equal(t, testStrings[i], tc.String())
	}
}

func TestParseDataTimeExample(t *testing.T) {
	header := makeHeader("v1.0", "/extensions/functions_datetime.yaml")
	tests := `#  timestamp examples using the timestamp type 
lt('2016-12-31T13:30:15'::ts, '2017-12-31T13:30:15'::ts) = true::bool
`
	testFile, err := ParseTestCasesFromString(header + tests)
	require.NoError(t, err)
	require.NotNil(t, testFile)
	assert.Len(t, testFile.TestCases, 1)
	assert.Equal(t, "lt", testFile.TestCases[0].FuncName)

	assert.Equal(t, testFile.TestCases[0].BaseURI, "/extensions/functions_datetime.yaml")
	assert.Equal(t, testFile.TestCases[0].GroupDesc, "timestamp examples using the timestamp type")
	assert.Len(t, testFile.TestCases[0].Args, 2)
	tsLiteral, err := literal.NewTimestampFromString("2016-12-31T13:30:15", false)
	require.NoError(t, err)
	assert.Equal(t, tsLiteral, testFile.TestCases[0].Args[0].Value)
	tsLiteral, err = literal.NewTimestampFromString("2017-12-31T13:30:15", false)
	require.NoError(t, err)
	assert.Equal(t, tsLiteral, testFile.TestCases[0].Args[1].Value)
	boolLiteral := literal.NewBool(true, false)
	assert.Equal(t, boolLiteral, testFile.TestCases[0].Result.Value)
	assert.Equal(t, &types.BooleanType{Nullability: types.NullabilityRequired}, testFile.TestCases[0].Result.Type)
	timestampType := &types.TimestampType{Nullability: types.NullabilityRequired}
	assert.Equal(t, timestampType, testFile.TestCases[0].Args[0].Type)
	assert.Equal(t, timestampType, testFile.TestCases[0].Args[1].Type)
	assert.Equal(t, ScalarFuncType, testFile.TestCases[0].FuncType)
}

func TestParseDecimalExample(t *testing.T) {
	header := makeHeader("v1.0", "extensions/functions_arithmetic_decimal.yaml")
	tests := `# basic
power(8::dec<38,0>, 2::dec<38, 0>) = 64::fp64
power(1.0::dec<38, 5>, -1.0::dec<38, 5>) = 1.0::fp64
power(-1::dec, 0.5::dec<38,1>) [complex_number_result:NAN] = nan::fp64

add(0.5::dec<1, 1>, 0.25::dec<2, 2>) = 0.75::dec<5, 2>
`

	testFile, err := ParseTestCasesFromString(header + tests)
	require.NoError(t, err)
	require.NotNil(t, testFile)
	assert.Len(t, testFile.TestCases, 4)
	assert.Equal(t, "power", testFile.TestCases[0].FuncName)
	assert.Equal(t, "power", testFile.TestCases[1].FuncName)
	assert.Equal(t, "power", testFile.TestCases[2].FuncName)
	dec8Value, _ := literal.NewDecimalFromString("8", false)
	dec2Value, _ := literal.NewDecimalFromString("2", false)
	dec1Value, _ := literal.NewDecimalFromString("1.0", false)
	decMinus1Value, _ := literal.NewDecimalFromString("-1", false)
	decMinus1Point0Value, _ := literal.NewDecimalFromString("-1.0", false)
	decPoint5Value, _ := literal.NewDecimalFromString("0.5", false)
	dec8, _ := dec8Value.(expr.WithTypeLiteral).WithType(&types.DecimalType{Precision: 38, Scale: 0, Nullability: types.NullabilityRequired})
	dec2, _ := dec2Value.(expr.WithTypeLiteral).WithType(&types.DecimalType{Precision: 38, Scale: 0, Nullability: types.NullabilityRequired})
	dec1, _ := dec1Value.(expr.WithTypeLiteral).WithType(&types.DecimalType{Precision: 38, Scale: 5, Nullability: types.NullabilityRequired})
	decMinus1, _ := decMinus1Value.(expr.WithTypeLiteral).WithType(&types.DecimalType{Precision: 38, Scale: 0, Nullability: types.NullabilityRequired})
	decMinus1Point0, _ := decMinus1Point0Value.(expr.WithTypeLiteral).WithType(&types.DecimalType{Precision: 38, Scale: 5, Nullability: types.NullabilityRequired})
	decPoint5, _ := decPoint5Value.(expr.WithTypeLiteral).WithType(&types.DecimalType{Precision: 38, Scale: 1, Nullability: types.NullabilityRequired})
	f6464 := literal.NewFloat64(64, false)
	f641 := literal.NewFloat64(1, false)
	assert.Equal(t, dec8, testFile.TestCases[0].Args[0].Value)
	assert.Equal(t, dec2, testFile.TestCases[0].Args[1].Value)
	assert.Equal(t, f6464, testFile.TestCases[0].Result.Value)
	assert.Equal(t, dec1, testFile.TestCases[1].Args[0].Value)
	assert.Equal(t, decMinus1Point0, testFile.TestCases[1].Args[1].Value)
	assert.Equal(t, f641, testFile.TestCases[1].Result.Value)
	assert.Equal(t, decMinus1, testFile.TestCases[2].Args[0].Value)
	assert.Equal(t, decPoint5, testFile.TestCases[2].Args[1].Value)
	assert.Equal(t, "fp64(NaN)", testFile.TestCases[2].Result.Value.String())

	decPoint25Value, _ := literal.NewDecimalFromString("0.25", false)
	decPoint75Value, _ := literal.NewDecimalFromString("0.75", false)
	decPoint25, _ := decPoint25Value.(expr.WithTypeLiteral).WithType(&types.DecimalType{Precision: 2, Scale: 2, Nullability: types.NullabilityRequired})
	decPoint75, _ := decPoint75Value.(expr.WithTypeLiteral).WithType(&types.DecimalType{Precision: 5, Scale: 2, Nullability: types.NullabilityRequired})
	decPoint5, _ = decPoint5Value.(expr.WithTypeLiteral).WithType(&types.DecimalType{Precision: 1, Scale: 1, Nullability: types.NullabilityRequired})
	assert.Equal(t, decPoint5, testFile.TestCases[3].Args[0].Value)
	assert.Equal(t, decPoint25, testFile.TestCases[3].Args[1].Value)
	assert.Equal(t, decPoint75, testFile.TestCases[3].Result.Value)
}

func TestParseTestWithVariousTypes(t *testing.T) {
	header := makeHeader("v1.0", "extensions/functions_arithmetic_decimal.yaml") + "# basic \n"
	tests := []struct {
		testCaseStr string
		expTestStr  string
	}{
		{testCaseStr: "f1(1::i8, 2::i16, 3::i32, 4::i64?) = -7.0::fp32?", expTestStr: "f1(1::i8, 2::i16, 3::i32, 4::i64?) = -7::fp32?"},
		{testCaseStr: "f2(1.0::fp32, 2.0::fp64) = -7.0::fp32", expTestStr: "f2(1::fp32, 2::fp64) = -7::fp32"},
		{testCaseStr: "f3('a'::str, 'b'::string) = 'c'::str", expTestStr: "f3('a'::string, 'b'::string) = 'c'::string"},
		{testCaseStr: "f4(false::bool, true::boolean) = false::bool", expTestStr: "f4(false::boolean, true::boolean) = false::boolean"},
		{testCaseStr: "f5(1::dec, 2::decimal) = 3::dec", expTestStr: "f5(1::decimal<38,0>, 2::decimal<38,0>) = 3::decimal<38,0>"},
		{testCaseStr: "f6(1.1::dec<38,10>, 2.2::dec<38,10>) = 3.3::dec<38,10>", expTestStr: "f6(1.1000000000::decimal<38,10>, 2.2000000000::decimal<38,10>) = 3.3000000000::decimal<38,10>"},
		{testCaseStr: "f7(1.1::dec<38,1>, 2.2::decimal<38,1>) = 3.3::decimal<38,1>", expTestStr: "f7(1.1::decimal<38,1>, 2.2::decimal<38,1>) = 3.3::decimal<38,1>"},
		{testCaseStr: "f8('1991-01-01'::date) = '2001-01-01'::date"},
		{testCaseStr: "f8('13:01:01.2345678'::time) = 123456::i64", expTestStr: "f8('13:01:01.234567'::time) = 123456::i64"},
		{testCaseStr: "f8('13:01:01.234'::time?) = 123::i32?", expTestStr: "f8('13:01:01.234000'::time?) = 123::i32?"},
		{testCaseStr: "f8('1991-01-01T01:02:03.456'::ts, '1991-01-01T00:00:00'::timestamp) = '1991-01-01T22:33:44'::ts", expTestStr: "f8('1991-01-01T01:02:03.456'::timestamp, '1991-01-01T00:00:00'::timestamp) = '1991-01-01T22:33:44'::timestamp"},
		{testCaseStr: "f8('1991-01-01T01:02:03.456+05:30'::tstz, '1991-01-01T00:00:00+15:30'::timestamp_tz) = 23::i32", expTestStr: "f8('1990-12-31T19:32:03.456'::timestamp_tz, '1990-12-31T08:30:00'::timestamp_tz) = 23::i32"},
		{testCaseStr: "f9('1991-01-01'::date, 5::i64) = '1991-01-01T00:00:00+15:30'::timestamp_tz", expTestStr: "f9('1991-01-01'::date, 5::i64) = '1990-12-31T08:30:00'::timestamp_tz"},
		{testCaseStr: "f10('P10Y5M'::interval_year, 5::i64) = 'P15Y5M'::interval_year"},
		{testCaseStr: "f10('P10Y5M'::iyear, 5::i64) = 'P15Y5M'::iyear", expTestStr: "f10('P10Y5M'::interval_year, 5::i64) = 'P15Y5M'::interval_year"},
		{testCaseStr: "f11('P10DT5H6M7S'::interval_day, 5::i64) = 'P10DT10H6M7S'::interval_day", expTestStr: "f11('P10DT5H6M7S'::interval_day<0>, 5::i64) = 'P10DT10H6M7S'::interval_day<0>"},
		{testCaseStr: "f11('P10DT6M7S'::interval_day, 5::i64) = 'P10DT11M7S'::interval_day", expTestStr: "f11('P10DT6M7S'::interval_day<0>, 5::i64) = 'P10DT11M7S'::interval_day<0>"},
		{testCaseStr: "or(false::bool, null::bool) = null::bool", expTestStr: "or(false::boolean, null::boolean?) = null::boolean?"},
		{testCaseStr: "f12('a'::vchar<9>, 'b'::varchar<4>) = 'c'::varchar<3>", expTestStr: "f12('a'::varchar<9>, 'b'::varchar<4>) = 'c'::varchar<3>"},
		{testCaseStr: "f8('1991-01-01T01:02:03.456'::pts<3>, '1991-01-01T00:00:00.000000'::pts<6>) = '1991-01-01T22:33:44'::pts<0>", expTestStr: "f8('1991-01-01T01:02:03.456'::precision_timestamp<3>, '1991-01-01T00:00:00'::precision_timestamp<6>) = '1991-01-01T22:33:44'::precision_timestamp<0>"},
		{testCaseStr: "f8('1991-01-01T01:02:03.456+05:30'::ptstz<3>, '1991-01-01T00:00:00+15:30'::ptstz<0>) = '1991-01-01T22:33:44+15:30'::ptstz<0>", expTestStr: "f8('1990-12-31T19:32:03.456+00:00'::precision_timestamp_tz<3>, '1990-12-31T08:30:00.000+00:00'::precision_timestamp_tz<0>) = '1991-01-01T07:03:44.000+00:00'::precision_timestamp_tz<0>"},
		//{"f12('P10DT6M7.2000S'::iday<4>, 5::i64) = 'P10DT11M7.2000S'::iday<4>"},  // TODO enable after fixing the grammar
		{testCaseStr: "f12('P10DT6M7S'::interval_day, 5::i64) = 'P10DT11M7S'::interval_day", expTestStr: "f12('P10DT6M7S'::interval_day<0>, 5::i64) = 'P10DT11M7S'::interval_day<0>"},
		{testCaseStr: "concat('abcd'::varchar<9>, Null::str) [null_handling:ACCEPT_NULLS] = Null::str", expTestStr: "concat('abcd'::varchar<9>, null::string?) [null_handling:ACCEPT_NULLS] = null::string?"},
		{testCaseStr: "concat('abcd'::varchar<9>, null::string) [null_handling:ACCEPT_NULLS] = null::string", expTestStr: "concat('abcd'::varchar<9>, null::string?) [null_handling:ACCEPT_NULLS] = null::string?"},
		{testCaseStr: "concat('abcd'::varchar<9>, null::varchar?<9>) [null_handling:ACCEPT_NULLS] = null::varchar?<9>"},
		{testCaseStr: "concat('abcd'::vchar<9>, 'ef'::varchar<9>) = 'abcdef'::vchar<9>", expTestStr: "concat('abcd'::varchar<9>, 'ef'::varchar<9>) = 'abcdef'::varchar<9>"},
		{testCaseStr: "concat('abcd'::fchar<9>, 'ef'::fixedchar<9>) = 'abcdef'::fchar<9>", expTestStr: "concat('abcd'::fixedchar<9>, 'ef'::fixedchar<9>) = 'abcdef'::fixedchar<9>"},
		{testCaseStr: "concat('abcd'::vchar<9>, Null::varchar<9>) = Null::vchar<9>", expTestStr: "concat('abcd'::varchar<9>, null::varchar?<9>) = null::varchar?<9>"},
		{testCaseStr: "concat('abcd'::vchar<9>, Null::fixedchar<9>) = Null::fchar<9>", expTestStr: "concat('abcd'::varchar<9>, null::fixedchar?<9>) = null::fixedchar?<9>"},
		{testCaseStr: "concat('abcd'::fbin<9>, Null::fixedbinary<9>) = Null::fbin<9>", expTestStr: "concat('0x61626364'::fixedbinary<9>, null::fixedbinary?<9>) = null::fixedbinary?<9>"},
		{testCaseStr: "f35('1991-01-01T01:02:03.456'::pts<3>) = '1991-01-01T01:02:30.123123'::precision_timestamp<3>", expTestStr: "f35('1991-01-01T01:02:03.456'::precision_timestamp<3>) = '1991-01-01T01:02:30.123'::precision_timestamp<3>"},
		{testCaseStr: "f36('1991-01-01T01:02:03.456'::pts<3>, '1991-01-01T01:02:30.123123'::precision_timestamp<3>) = 123456::i64", expTestStr: "f36('1991-01-01T01:02:03.456'::precision_timestamp<3>, '1991-01-01T01:02:30.123'::precision_timestamp<3>) = 123456::i64"},
		{testCaseStr: "f37('1991-01-01T01:02:03.123456'::pts<6>, '1991-01-01T04:05:06.456'::precision_timestamp<6>) = 123456::i64", expTestStr: "f37('1991-01-01T01:02:03.123456'::precision_timestamp<6>, '1991-01-01T04:05:06.456'::precision_timestamp<6>) = 123456::i64"},
		{testCaseStr: "f38('1991-01-01T01:02:03.456+05:30'::ptstz<3>) = '1991-01-01T00:00:00+15:30'::precision_timestamp_tz<3>", expTestStr: "f38('1990-12-31T19:32:03.456+00:00'::precision_timestamp_tz<3>) = '1990-12-31T08:30:00.000+00:00'::precision_timestamp_tz<3>"},
		{testCaseStr: "f39('1991-01-01T01:02:03.123456+05:30'::ptstz<6>) = '1991-01-01T00:00:00+15:30'::precision_timestamp_tz<6>", expTestStr: "f39('1990-12-31T19:32:03.123+00:00'::precision_timestamp_tz<6>) = '1990-12-31T08:30:00.000+00:00'::precision_timestamp_tz<6>"},
	}
	for _, test := range tests {
		t.Run(test.testCaseStr, func(t *testing.T) {
			testFile, err := ParseTestCasesFromString(header + test.testCaseStr)
			require.NoError(t, err)
			require.NotNil(t, testFile)
			assert.Len(t, testFile.TestCases, 1)
			if test.expTestStr != "" {
				assert.Equal(t, test.expTestStr, testFile.TestCases[0].String())
			} else {
				assert.Equal(t, test.testCaseStr, testFile.TestCases[0].String())
			}
			for _, arg := range testFile.TestCases[0].Args {
				assert.NotNil(t, arg.Value)
				checkNullability(t, arg.Value, arg.Type)
			}
			assert.NotNil(t, testFile.TestCases[0].Result.Value)
			checkNullability(t, testFile.TestCases[0].Result.Value, testFile.TestCases[0].Result.Type)
		})
	}
}

func checkNullability(t *testing.T, lit expr.Literal, argType types.Type) {
	if _, ok := lit.(*expr.NullLiteral); !ok {
		assert.Equal(t, lit.GetType().GetNullability(), argType.GetNullability())
	} else {
		assert.Equal(t, types.NullabilityNullable, argType.GetNullability())
	}
	assert.Equal(t, argType, lit.GetType())
}

func TestParseStringTestCases(t *testing.T) {
	header := makeHeader("v1.0", "extensions/functions_arithmetic_decimal.yaml")
	tests := `# basic
concat('abc'::str, 'def'::str) = 'abcdef'::str
regexp_string_split('HHHelloooo'::str, 'Hel+'::str) = ['HH', 'oooo']::List<str>
octet_length('à'::str) = 2::i64
octet_length('😄'::str) = 4::i64
starts_with('abcd'::str, 'AB'::str) [case_sensitivity:CASE_INSENSITIVE] = true::bool`

	testFile, err := ParseTestCasesFromString(header + tests)
	require.NoError(t, err)
	require.NotNil(t, testFile)
	assert.Len(t, testFile.TestCases, 5)
	assert.Equal(t, "concat", testFile.TestCases[0].FuncName)
	assert.Equal(t, "regexp_string_split", testFile.TestCases[1].FuncName)
	assert.Equal(t, "octet_length", testFile.TestCases[2].FuncName)
	assert.Equal(t, "octet_length", testFile.TestCases[3].FuncName)

	strAbc := literal.NewString("abc", false)
	strDef := literal.NewString("def", false)
	strRes := literal.NewString("abcdef", false)
	assert.Equal(t, strAbc, testFile.TestCases[0].Args[0].Value)
	assert.Equal(t, strDef, testFile.TestCases[0].Args[1].Value)
	assert.Equal(t, strRes, testFile.TestCases[0].Result.Value)

	strArg1 := literal.NewString("HHHelloooo", false)
	strArg2 := literal.NewString("Hel+", false)
	//listStr := &types.ListType{Type: &types.StringType{}}
	assert.Equal(t, strArg1, testFile.TestCases[1].Args[0].Value)
	assert.Equal(t, strArg2, testFile.TestCases[1].Args[1].Value)
	strRes1 := literal.NewString("HH", false)
	strRes2 := literal.NewString("oooo", false)
	result, _ := literal.NewList([]expr.Literal{strRes1, strRes2}, false)
	assert.Equal(t, result, testFile.TestCases[1].Result.Value)

	str1 := literal.NewString("à", false)
	i642 := literal.NewInt64(2, false)
	assert.Equal(t, str1, testFile.TestCases[2].Args[0].Value)
	assert.Equal(t, i642, testFile.TestCases[2].Result.Value)

	str2 := literal.NewString("😄", false)
	i644 := literal.NewInt64(4, false)
	assert.Equal(t, str2, testFile.TestCases[3].Args[0].Value)
	assert.Equal(t, i644, testFile.TestCases[3].Result.Value)

}

func TestParseStringWithIntList(t *testing.T) {
	header := makeHeader("v1.0", "extensions/functions_arithmetic_decimal.yaml")
	tests := `# basic
some_func('abc'::str, 'def'::str) = [1, 2, 3, 4, 5, 6]::List<i8>`

	testFile, err := ParseTestCasesFromString(header + tests)
	require.NoError(t, err)
	require.NotNil(t, testFile)
	assert.Len(t, testFile.TestCases, 1)

	assert.Equal(t, "some_func", testFile.TestCases[0].FuncName)
	strAbc := literal.NewString("abc", false)
	strDef := literal.NewString("def", false)
	assert.Equal(t, strAbc, testFile.TestCases[0].Args[0].Value)
	assert.Equal(t, strDef, testFile.TestCases[0].Args[1].Value)
	i8List := &types.ListType{Type: &types.Int8Type{Nullability: types.NullabilityRequired}, Nullability: types.NullabilityRequired}
	list, _ := literal.NewList([]expr.Literal{
		literal.NewInt8(1, false), literal.NewInt8(2, false), literal.NewInt8(3, false),
		literal.NewInt8(4, false), literal.NewInt8(5, false), literal.NewInt8(6, false),
	}, false)
	assert.Equal(t, list, testFile.TestCases[0].Result.Value)
	assert.Equal(t, i8List, testFile.TestCases[0].Result.Type)
}

func TestScalarOptions(t *testing.T) {
	header := makeHeader("v1.0", "extensions/functions_string.yaml")
	tests := `# stuff
contains('abba'::str, 'AB'::str) [case_sensitivity:CASE_INSENSITIVE] = true::bool`

	testFile, err := ParseTestCasesFromString(header + tests)
	require.NoError(t, err)
	require.NotNil(t, testFile)
	assert.Len(t, testFile.TestCases, 1)
	assert.Len(t, testFile.TestCases[0].Options, 1)
	assert.Equal(t, "CASE_INSENSITIVE", testFile.TestCases[0].Options["case_sensitivity"])
}

func TestMultipleScalarOptions(t *testing.T) {
	header := makeHeader("v1.0", "extensions/functions_arithmetic.yaml")
	tests := `# stuff
add(2::fp64, 2::fp64) [overflow:ERROR, rounding:TIE_TO_EVEN] = 4::fp64`

	testFile, err := ParseTestCasesFromString(header + tests)
	require.NoError(t, err)
	require.NotNil(t, testFile)
	assert.Len(t, testFile.TestCases, 1)
	assert.Len(t, testFile.TestCases[0].Options, 2)
	assert.Equal(t, "ERROR", testFile.TestCases[0].Options["overflow"])
	assert.Equal(t, "TIE_TO_EVEN", testFile.TestCases[0].Options["rounding"])
}

func TestParseAggregateFunc(t *testing.T) {
	header := makeAggregateTestHeader("v1.0", "/extensions/functions_arithmetic.yaml")
	tests := `# basic
avg((1,2,3)::fp32) = 2::fp64
avg((1,2,3, NULL)::fp32?) = 2::fp64?
sum((9223372036854775806, 1, 1, 1, 1, 10000000000)::i64) [overflow:ERROR] = <!ERROR>`

	reg, funcRegistry := functions.NewExtensionAndFunctionRegistries(extensions.GetDefaultCollectionWithNoError())
	arithUri := "https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml"
	testFile, err := ParseTestCasesFromString(header + tests)
	require.NoError(t, err)
	require.NotNil(t, testFile)
	assert.Len(t, testFile.TestCases, 3)
	assert.Equal(t, "avg", testFile.TestCases[0].FuncName)
	tc := testFile.TestCases[0]
	assert.Equal(t, testFile.TestCases[0].GroupDesc, "basic")
	assert.Equal(t, testFile.TestCases[0].BaseURI, "/extensions/functions_arithmetic.yaml")
	assert.Len(t, testFile.TestCases[0].Args, 0)
	assert.Len(t, testFile.TestCases[0].AggregateArgs, 1)
	assert.Equal(t, "fp32", testFile.TestCases[0].AggregateArgs[0].ColumnType.String())
	listType := &types.ListType{
		Type:        &types.Float32Type{Nullability: types.NullabilityRequired},
		Nullability: types.NullabilityRequired,
	}

	testStrings := []string{
		"avg((1, 2, 3)::fp32) = 2::fp64",
		"avg((1, 2, 3, NULL)::fp32?) = 2::fp64?",
		"sum((9223372036854775806, 1, 1, 1, 1, 10000000000)::i64) [overflow:ERROR] = <!ERROR>",
	}
	assert.Equal(t, newFloat32List(1, 2, 3), tc.AggregateArgs[0].Argument.Value)
	assert.Equal(t, listType, tc.AggregateArgs[0].Argument.Value.GetType())
	assert.Equal(t, "fp64", tc.Result.Type.String())
	assert.Equal(t, literal.NewFloat64(2, false), tc.Result.Value)
	assert.Equal(t, AggregateFuncType, tc.FuncType)
	_, err = tc.GetScalarFunctionInvocation(nil, nil)
	require.Error(t, err)
	assert.Equal(t, extensions.ID{URI: arithUri, Name: "avg:fp32"}, tc.ID())
	assert.Equal(t, "avg:fp32", tc.CompoundFunctionName())
	aggregateFunc, err1 := tc.GetAggregateFunctionInvocation(&reg, funcRegistry)
	require.NoError(t, err1)
	assert.Equal(t, tc.FuncName, aggregateFunc.Name())
	require.Equal(t, 1, aggregateFunc.NArgs())
	aggArg, ok := aggregateFunc.Arg(0).(*expr.FieldReference)
	require.True(t, ok)
	assert.Equal(t, &types.Float32Type{Nullability: types.NullabilityRequired}, aggArg.GetType())
	assert.Equal(t, ".field(0) => fp32", aggArg.String())
	assert.Equal(t, []types.Type{&types.Float32Type{Nullability: types.NullabilityRequired}}, tc.GetArgTypes())
	assert.Equal(t, testStrings[0], tc.String())

	tc = testFile.TestCases[1]
	aggregateFunc, err1 = tc.GetAggregateFunctionInvocation(&reg, funcRegistry)
	require.NoError(t, err1)
	aggArg, ok = aggregateFunc.Arg(0).(*expr.FieldReference)
	require.True(t, ok)
	assert.Equal(t, tc.FuncName, aggregateFunc.Name())
	require.Equal(t, 1, aggregateFunc.NArgs())
	assert.Equal(t, &types.Float32Type{Nullability: types.NullabilityNullable}, aggArg.GetType())
	assert.Equal(t, []types.Type{&types.Float32Type{Nullability: types.NullabilityNullable}}, tc.GetArgTypes())
	assert.Equal(t, tc.Result.Type, &types.Float64Type{Nullability: types.NullabilityNullable})
	argValues := newFloat32Values(true, 1, 2, 3)
	argValues = append(argValues, &expr.NullLiteral{Type: &types.Float32Type{Nullability: types.NullabilityNullable}})
	argList, _ := literal.NewList(argValues, false)
	assert.Equal(t, argList, tc.AggregateArgs[0].Argument.Value)

	tc = testFile.TestCases[2]
	assert.Equal(t, "sum", tc.FuncName)
	assert.Equal(t, testFile.TestCases[1].GroupDesc, "basic")
	assert.Equal(t, testFile.TestCases[1].BaseURI, "/extensions/functions_arithmetic.yaml")
	assert.Len(t, testFile.TestCases[1].Args, 0)
	assert.Len(t, testFile.TestCases[1].AggregateArgs, 1)
	assert.Equal(t, AggregateFuncType, tc.FuncType)
	assert.Equal(t, "i64", tc.AggregateArgs[0].ColumnType.String())
	assert.Equal(t, newInt64List(9223372036854775806, 1, 1, 1, 1, 10000000000), tc.AggregateArgs[0].Argument.Value)
	assert.Equal(t, "ERROR", tc.Options["overflow"])
	assert.Equal(t, testStrings[2], tc.String())

	_, err = tc.GetScalarFunctionInvocation(nil, nil)
	require.Error(t, err)
	assert.Equal(t, extensions.ID{URI: arithUri, Name: "sum:i64"}, tc.ID())
	assert.Equal(t, "sum:i64", tc.CompoundFunctionName())
	aggregateFunc, err1 = tc.GetAggregateFunctionInvocation(&reg, funcRegistry)
	require.NoError(t, err1)
	assert.Equal(t, tc.FuncName, aggregateFunc.Name())
	require.Equal(t, 1, aggregateFunc.NArgs())
	aggArg, ok = aggregateFunc.Arg(0).(*expr.FieldReference)
	require.True(t, ok)
	assert.Equal(t, &types.Int64Type{Nullability: types.NullabilityRequired}, aggArg.GetType())
	assert.Equal(t, ".field(0) => i64", aggArg.String())
	assert.Equal(t, []types.Type{&types.Int64Type{Nullability: types.NullabilityRequired}}, tc.GetArgTypes())
}

func newInt64List(values ...int64) interface{} {
	list, _ := literal.NewList(newInt64Values(values...), false)
	return list
}

func newInt64Values(values ...int64) []expr.Literal {
	literals := make([]expr.Literal, len(values))
	for i, v := range values {
		literals[i] = literal.NewInt64(v, false)
	}
	return literals
}

func newFloat32List(values ...float32) expr.Literal {
	list, _ := literal.NewList(newFloat32Values(false, values...), false)
	return list
}

func newFloat32Values(nullable bool, values ...float32) []expr.Literal {
	literals := make([]expr.Literal, len(values))
	for i, v := range values {
		literals[i] = literal.NewFloat32(v, nullable)
	}
	return literals
}

func TestParseAggregateFuncCompact(t *testing.T) {
	header := makeAggregateTestHeader("v1.0", "/extensions/functions_arithmetic.yaml")
	tests := `# basic
((20, 20), (-3, -3), (1, 1), (10,10), (5,5)) corr(col0::fp32, col1::fp32) = 1::fp64
((20, 20), (-3, -3), (1, 1), (10,10), (5,5)) corr(col0::fp32, col1::fp32?) = 1::fp64?
`

	testString := "((20, 20), (-3, -3), (1, 1), (10, 10), (5, 5)) corr(col0::fp32, col1::fp32) = 1::fp64"
	testFile, err := ParseTestCasesFromString(header + tests)
	require.NoError(t, err)
	require.NotNil(t, testFile)
	assert.Len(t, testFile.TestCases, 2)
	tc := testFile.TestCases[0]
	assert.Equal(t, "corr", tc.FuncName)
	assert.Equal(t, tc.GroupDesc, "basic")
	assert.Equal(t, tc.BaseURI, "/extensions/functions_arithmetic.yaml")
	assert.Len(t, tc.Args, 0)
	assert.Len(t, tc.AggregateArgs, 2)
	assert.Equal(t, newFloat32Values(false, 20, -3, 1, 10, 5), tc.Columns[0])
	assert.Equal(t, newFloat32Values(false, 20, -3, 1, 10, 5), tc.Columns[1])
	f32Type := &types.Float32Type{Nullability: types.NullabilityRequired}
	args := []*AggregateArgument{
		createAggregateArg(t, "", "col0", f32Type),
		createAggregateArg(t, "", "col1", f32Type),
	}
	assert.Equal(t, args, tc.AggregateArgs)
	assert.Equal(t, "fp64", tc.Result.Type.String())
	assert.Equal(t, literal.NewFloat64(1, false), tc.Result.Value)
	assert.Equal(t, testString, tc.String())

	tc = testFile.TestCases[1]
	args[1] = createAggregateArg(t, "", "col1", f32Type.WithNullability(types.NullabilityNullable))
	assert.Equal(t, args, tc.AggregateArgs)
	assert.Equal(t, "fp64?", tc.Result.Type.String())
}

func createAggregateArg(t *testing.T, tableName, columnName string, columnType types.Type) *AggregateArgument {
	arg, err := newAggregateArgument(tableName, columnName, columnType)
	require.NoError(t, err)
	return arg
}

func TestParseAggregateFuncWithMultipleArgs(t *testing.T) {
	header := makeAggregateTestHeader("v1.0", "/extensions/functions_arithmetic.yaml")
	tests := `#  basic 
DEFINE t1(fp32, fp32) = ((20, 20), (-3, -3), (1, 1), (10,10), (5,5.5))
corr(t1.col0, t1.col1) = 1::fp64
DEFINE t1(i64, fp32) = ((20, 20), (-3, -3), (1, 1), (10,10), (5,5.5))
corr(t1.col1, t1.col0) = 1::fp64
`

	testStrings := []string{
		"((20, 20), (-3, -3), (1, 1), (10, 10), (5, 5.5)) corr(col0::fp32, col1::fp32) = 1::fp64",
		"((20, 20), (-3, -3), (1, 1), (10, 10), (5, 5.5)) corr(col1::fp32, col0::i64) = 1::fp64",
	}
	testFile, err := ParseTestCasesFromString(header + tests)
	require.NoError(t, err)
	require.NotNil(t, testFile)
	assert.Len(t, testFile.TestCases, 2)
	assert.Equal(t, "corr", testFile.TestCases[0].FuncName)
	assert.Equal(t, testFile.TestCases[0].GroupDesc, "basic")
	assert.Equal(t, testFile.TestCases[0].BaseURI, "/extensions/functions_arithmetic.yaml")
	assert.Len(t, testFile.TestCases[0].Args, 0)
	assert.Len(t, testFile.TestCases[0].AggregateArgs, 2)
	assert.Equal(t, newFloat32Values(false, 20, -3, 1, 10, 5), testFile.TestCases[0].Columns[0])
	assert.Equal(t, newFloat32Values(false, 20, -3, 1, 10, 5.5), testFile.TestCases[0].Columns[1])
	assert.Equal(t, "col0", testFile.TestCases[0].AggregateArgs[0].ColumnName)
	assert.Equal(t, "col1", testFile.TestCases[0].AggregateArgs[1].ColumnName)
	assert.Equal(t, testStrings[0], testFile.TestCases[0].String())

	assert.Equal(t, "corr", testFile.TestCases[1].FuncName)
	assert.Equal(t, testFile.TestCases[1].GroupDesc, "basic")
	assert.Equal(t, testFile.TestCases[1].BaseURI, "/extensions/functions_arithmetic.yaml")
	assert.Len(t, testFile.TestCases[1].Args, 0)
	assert.Len(t, testFile.TestCases[1].AggregateArgs, 2)
	assert.Equal(t, newInt64Values(20, -3, 1, 10, 5), testFile.TestCases[1].Columns[0])
	assert.Equal(t, newFloat32Values(false, 20, -3, 1, 10, 5.5), testFile.TestCases[1].Columns[1])
	assert.Equal(t, "col1", testFile.TestCases[1].AggregateArgs[0].ColumnName)
	assert.Equal(t, "col0", testFile.TestCases[1].AggregateArgs[1].ColumnName)
	assert.Equal(t, testStrings[1], testFile.TestCases[1].String())
}

func TestParseAggregateFuncWithVariousTypes(t *testing.T) {
	header := makeAggregateTestHeader("v1.0", "/extensions/functions_arithmetic.yaml")
	header += "# basic\n"

	tests := []struct {
		testCaseStr string
	}{
		{"avg((1,2,3)::i8) = 2::fp64"},
	}
	for _, test := range tests {
		t.Run(test.testCaseStr, func(t *testing.T) {
			testFile, err := ParseTestCasesFromString(header + test.testCaseStr)
			require.NoError(t, err)
			require.NotNil(t, testFile)
			assert.Len(t, testFile.TestCases, 1)
		})
	}
}

func TestParseAggregateFuncAllFormats(t *testing.T) {
	header := makeAggregateTestHeader("v1.0", "/extensions/functions_arithmetic.yaml")
	header += "# basic\n"

	tests := []struct {
		testCaseStr string
		wantData    [][]expr.Literal
	}{
		{"avg((1,2,3)::i64) = 2::fp64", [][]expr.Literal{newInt64Values(1, 2, 3)}},
		{"((1), (2), (3)) avg(col0::i64) = 2::fp64", [][]expr.Literal{newInt64Values(1, 2, 3)}},
		{"DEFINE t1(i64) = ((1), (2), (3))\navg(t1.col0) = 2::fp64", [][]expr.Literal{newInt64Values(1, 2, 3)}},

		// tests with empty input data
		{"avg(()::i64) = 2::fp64", [][]expr.Literal{{}}},
		{"DEFINE t1(i64) = ()\navg(t1.col0) = 2::fp64", [][]expr.Literal{{}}},

		//tests with multiple columns
		{"((20, 20), (-3, -3), (1, 1), (10,10), (5,5)) corr(col0::fp32, col1::fp32?) = 1::fp64?", [][]expr.Literal{newFloat32Values(false, 20, -3, 1, 10, 5), newFloat32Values(true, 20, -3, 1, 10, 5)}},
		{"DEFINE t1(fp32, fp32?) = ((20, 20), (-3, -3), (1, 1), (10,10), (5,5))\ncorr(t1.col0, t1.col1) = 1::fp64?", [][]expr.Literal{newFloat32Values(false, 20, -3, 1, 10, 5), newFloat32Values(true, 20, -3, 1, 10, 5)}},
	}
	for _, test := range tests {
		t.Run(test.testCaseStr, func(t *testing.T) {
			testFile, err := ParseTestCasesFromString(header + test.testCaseStr)
			require.NoError(t, err)
			require.NotNil(t, testFile)
			assert.Len(t, testFile.TestCases, 1)
			tc := testFile.TestCases[0]
			assert.Contains(t, test.testCaseStr, tc.FuncName)
			assert.Equal(t, tc.GroupDesc, "basic")
			assert.Equal(t, tc.BaseURI, "/extensions/functions_arithmetic.yaml")
			assert.Len(t, tc.Args, 0)

			// check that the types are correct
			argTypes := tc.GetArgTypes()
			assert.Len(t, argTypes, len(test.wantData))
			if len(test.wantData[0]) > 0 {
				for i, argType := range argTypes {
					assert.Equal(t, argType, test.wantData[i][0].GetType())
				}
			} else {
				// check that the type is correct for empty input data
				assert.Equal(t, &types.Int64Type{Nullability: types.NullabilityRequired}, argTypes[0])
			}

			assert.Equal(t, AggregateFuncType, tc.FuncType)
			_, err = tc.GetScalarFunctionInvocation(nil, nil)
			require.Error(t, err)

			reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
			testGetFunctionInvocation(t, tc, &reg, nil)
			data, err := tc.GetAggregateColumnsData()
			require.NoError(t, err)

			// check that the data is correct
			assert.Len(t, data, len(test.wantData))
			assert.Equal(t, test.wantData, data)
		})
	}
}

func TestBadInputsToGetAggregateColumnsData(t *testing.T) {
	tests := []struct {
		name          string
		testCase      *TestCase
		expectedError error
	}{
		{
			name:          "invalid function type",
			testCase:      &TestCase{FuncType: ScalarFuncType},
			expectedError: fmt.Errorf("expected function type %v, but got %v", AggregateFuncType, ScalarFuncType),
		},
		{
			name: "invalid argument type",
			testCase: &TestCase{
				FuncType:      AggregateFuncType,
				AggregateArgs: []*AggregateArgument{{Argument: &CaseLiteral{Value: expr.NewNullLiteral(&types.Float32Type{})}}},
			},
			expectedError: fmt.Errorf("column 0: expected NestedLiteral[ListLiteralValue], but got %T", expr.NewNullLiteral(&types.Float32Type{})),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.testCase.GetAggregateColumnsData()
			assert.Error(t, err)
			assert.Equal(t, tt.expectedError.Error(), err.Error())
		})
	}
}

func TestParseAggregateFuncWithMixedArgs(t *testing.T) {
	header := makeAggregateTestHeader("v1.0", "/extensions/functions_arithmetic.yaml")
	tests := `# basic
((20), (-3), (1), (10)) LIST_AGG(col0::fp32, ','::string) = 1::fp64
DEFINE t1(fp32) = ((20), (-3), (1), (10))
LIST_AGG(t1.col0, ','::string) = 1::fp64
`

	testString := "((20), (-3), (1), (10)) LIST_AGG(col0::fp32, ','::string) = 1::fp64"
	testFile, err := ParseTestCasesFromString(header + tests)
	require.NoError(t, err)
	require.NotNil(t, testFile)
	assert.Len(t, testFile.TestCases, 2)
	expectedArgTypes := []types.Type{&types.Float32Type{Nullability: types.NullabilityRequired}, &types.StringType{Nullability: types.NullabilityRequired}}
	for i, tc := range testFile.TestCases {
		assert.Equal(t, AggregateFuncType, tc.FuncType)
		assert.Equal(t, expectedArgTypes, tc.GetArgTypes(), "unexpected arg types in test case %d", i)
		assert.Equal(t, "LIST_AGG:fp32_str", tc.ID().Name)
		assert.Equal(t, testString, tc.String(), "unexpected string in test case %d", i)
	}

	header = makeAggregateTestHeader("v1.0", "/extensions/functions_string.yaml")
	tests = `# basic
(('ant'), ('bat'), ('cat')) string_agg(col0::str, ','::str) = 'ant,bat,cat'::str
`
	testFile, err = ParseTestCasesFromString(header + tests)
	require.NoError(t, err)
	require.NotNil(t, testFile)
	assert.Len(t, testFile.TestCases, 1)
	assert.Equal(t, AggregateFuncType, testFile.TestCases[0].FuncType)
	reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
	aggFun, err := testFile.TestCases[0].GetAggregateFunctionInvocation(&reg, nil)
	require.NoError(t, err)
	assert.Equal(t, "string_agg", aggFun.Name())
	testString = "(('ant'), ('bat'), ('cat')) string_agg(col0::string, ','::string) = 'ant,bat,cat'::string"
	assert.Equal(t, testString, testFile.TestCases[0].String())
}

func TestParseTestWithBadScalarTests(t *testing.T) {
	header := makeHeader("v1.0", "extensions/functions_arithmetic_decimal.yaml") + "# basic \n"
	tests := []struct {
		testCaseStr string
		position    int
		errorMsg    string
	}{
		{"add(-12::i8, +5::i8) = -7.0::i8", 29, "no viable alternative at input '-7.0::i8'"},
		{"add(123.5::i8, 5::i8) = 125::i8", 11, "no viable alternative at input '123.5::i8'"},
		{"add(123.5::i16, 5.5::i16) = 125::i16", 11, "no viable alternative at input '123.5::i16'"},
		{"add(123.5::i32, 5.5::i32) = 125::i32", 21, "no viable alternative at input '5.5::i32'"},
		{"add(123f::i64, 5.5::i64) = 125::i64", 7, "no viable alternative at input '123f'"},
		{"add(123::i64, 5_000::i64) = 5123::i64", 15, "no viable alternative at input '5_000'"},
		{"add(123::dec<38,10>, 5.0E::dec<38,10>) = 123::dec<38,10>", 24, "no viable alternative at input '5.0E'"},
		{"add(123::dec<38,10>, 1a.2::dec<38,10>) = 123::fp32", 22, "no viable alternative at input '1a'"},
		{"add(123::dec<38,10>, 1.2.3::dec<38,10>) = 123::fp32", 24, "no viable alternative at input '1.2.'"},
		{"add(123::dec<38,10>, +-12.3::dec<38,10>) = 123::i64", 21, "extraneous input '+'"},
		{"add(123::fp32, .5E2::fp32) = 123::fp32", 15, "extraneous input '.'"},
		{"add(123::fp32, 4.1::fp32) = ++123::fp32", 28, "extraneous input '+'"},
		{"add(123::fp32, 2.5E::fp32) = 123::fp32", 18, "no viable alternative at input '2.5E'"},
		{"add(123::fp32, 1.4E+::fp32) = 123::fp32", 18, "no viable alternative at input '1.4E'"},
		{"add(123::fp32, 3.E.5::fp32) = 123::fp32", 17, "no viable alternative at input '3.E'"},
		{"f1((1, 2, 3, 4)::i64) = 10::fp64", 0, "expected scalar testcase based on test file header, but got aggregate function testcase"},
		{"add(4.53::dec<1, 0>, 0.25::dec<2, 2>) = 0.78::dec<5, 2>", 0, "Visit error at line 5: invalid argument number 4.53"},
	}
	for _, test := range tests {
		t.Run(test.testCaseStr, func(t *testing.T) {
			_, err := ParseTestCasesFromString(header + test.testCaseStr)
			require.Error(t, err)
			expectedErrorMsg := test.errorMsg
			if test.position > 0 {
				expectedErrorMsg = fmt.Sprintf("Syntax error at line 5:%d: %s", test.position, test.errorMsg)
			}
			assert.Contains(t, err.Error(), expectedErrorMsg)
		})
	}
}

func TestParseTestWithBadAggregateTests(t *testing.T) {
	header := makeAggregateTestHeader("v1.0", "extensions/functions_arithmetic_decimal.yaml") + "# basic \n"
	tests := []struct {
		testCaseStr string
		errorMsg    string
	}{
		{"max((-12, +5)::i8) = -7.0::i8", "no viable alternative at input '-7.0::i8'"},
		{"max((-12, 'arg')::i32) = -7::i8", "invalid column values"},
		{
			`DEFINE t1(fp32, fp32) = ((20, 20), (-3, -3), (1, 1), (10,10), (5,5))
corr(t1.col0, t2.col1) = 1::fp64`,
			"table name in argument t2, does not match the table name in the function call t1",
		},
		{"((20, 20), (-3, -3), (1, 1), (10,10), (5,5)) corr(my_col::fp32, col0::fp32) = 1::fp64", "mismatched input 'my_col'"},
		{"((20, 20), (-3, -3), (1, 1), (10,10), (5,5)) corr(col0::fp32, column1::fp32) = 1::fp64", "mismatched input 'column1'"},
		{"f8('13:01:01.234'::time) = 123::i32", "expected aggregate testcase based on test file header, but got scalar function testcase"},
	}
	for _, test := range tests {
		t.Run(test.testCaseStr, func(t *testing.T) {
			_, err := ParseTestCasesFromString(header + test.testCaseStr)
			require.Error(t, err)
			assert.Contains(t, err.Error(), test.errorMsg)
		})
	}
}

func TestParseAggregateTestWithVariousTypes(t *testing.T) {
	header := makeAggregateTestHeader("v1.0", "extensions/functions_arithmetic_decimal.yaml") + "# basic \n"
	tests := []struct {
		testCaseStr string
		expTestStr  string
	}{
		{testCaseStr: "f1((1, 2, 3, 4)::i64) = 10::fp64"},
		{testCaseStr: "f1((1, 2, 3, 4)::i16) = 10.0::fp32", expTestStr: "f1((1, 2, 3, 4)::i16) = 10::fp32"},
		{testCaseStr: "f1((1, 2, 3, 4)::i32) = 10::i64"},
		{testCaseStr: "f1((1, 2, 3, 4)::i32?) = 10::i64?"},
		{testCaseStr: "f3(('a', 'b')::string) = 'c'::str", expTestStr: "f3(('a', 'b')::string) = 'c'::string"},
		{testCaseStr: "f3(('a', 'b')::string?) = 'c'::str?", expTestStr: "f3(('a', 'b')::string?) = 'c'::string?"},
		{testCaseStr: "f4((false, true)::boolean) = false::bool", expTestStr: "f4((false, true)::boolean) = false::boolean"},
		{testCaseStr: "f4((false, true)::boolean?) = false::bool?", expTestStr: "f4((false, true)::boolean?) = false::boolean?"},
		{testCaseStr: "f5((1.1, 2.2)::fp32) = 3.3::fp32"},
		{testCaseStr: "f5((1.1, 2.2)::fp64) = 3.3::fp64"},
		{testCaseStr: "f5((1.1, 2.2)::fp64?) = 3.3::fp64?"},
		{testCaseStr: "f5((1, 2)::decimal) = 3::dec", expTestStr: "f5((1, 2)::decimal<38,0>) = 3::decimal<38,0>"},
		{testCaseStr: "f5((1.1, 2.2)::dec<38,1>) = 3.3::dec<38,1>", expTestStr: "f5((1.1, 2.2)::decimal<38,1>) = 3.3::decimal<38,1>"},
		{testCaseStr: "f6((1.1, 2.2)::dec<38,10>) = 3.3::dec<38,10>", expTestStr: "f6((1.1, 2.2)::decimal<38,10>) = 3.3000000000::decimal<38,10>"},
		{testCaseStr: "f7((1.0, 2)::decimal<38,0>) = 3::decimal<38,0>"},
		{testCaseStr: "f7((1.0, 2)::decimal?<38,0>) = 3::decimal?<38,0>"},
		{testCaseStr: "f6((1.1, 2.2, null)::dec?<38,10>) = 3.3::dec<38,10>", expTestStr: "f6((1.1, 2.2, null)::decimal?<38,10>) = 3.3000000000::decimal<38,10>"},
		{testCaseStr: "f8(('1991-01-01', '1991-02-02')::date) = '2001-01-01'::date"},
		{testCaseStr: "f8(('1991-01-01', '1991-02-02')::date?) = '2001-01-01'::date?"},
		{testCaseStr: "f8(('13:01:01.2345678', '14:01:01.333')::time) = 123456::i64", expTestStr: "f8(('13:01:01.234567', '14:01:01.333000')::time) = 123456::i64"},
		{testCaseStr: "f8(('1991-01-01T01:02:03.456', '1991-01-01T00:00:00')::timestamp) = '1991-01-01T22:33:44'::ts", expTestStr: "f8(('1991-01-01T01:02:03.456', '1991-01-01T00:00:00')::timestamp) = '1991-01-01T22:33:44'::timestamp"},
		{testCaseStr: "f8(('1991-01-01T01:02:03.456+05:30', '1991-01-01T00:00:00+15:30')::tstz) = 23::i32", expTestStr: "f8(('1990-12-31T19:32:03.456', '1990-12-31T08:30:00')::timestamp_tz) = 23::i32"},
		{testCaseStr: "f10(('P10Y5M', 'P11Y5M')::interval_year) = 'P21Y10M'::interval_year"},
		{testCaseStr: "f10(('P10Y5M', null)::interval_year) = null::interval_year", expTestStr: "f10(('P10Y5M', null)::interval_year?) = null::interval_year?"},
		{testCaseStr: "f10(('P10Y2M', 'P10Y7M')::iyear) = 'P20Y9M'::iyear", expTestStr: "f10(('P10Y2M', 'P10Y7M')::interval_year) = 'P20Y9M'::interval_year"},
		{testCaseStr: "f11(('P10DT5H6M7S', 'P10DT6M7S')::interval_day) = 'P20DT11H6M7S'::interval_day", expTestStr: "f11(('P10DT5H6M7S', 'P10DT6M7S')::interval_day<0>) = 'P20DT11H6M7S'::interval_day<0>"},
		{testCaseStr: "f11(('P10DT5H6M7S', 'P10DT6M7S')::iday?) = 'P20DT11H6M7S'::iday", expTestStr: "f11(('P10DT5H6M7S', 'P10DT6M7S')::interval_day?<0>) = 'P20DT11H6M7S'::interval_day<0>"},
		{testCaseStr: "f11(('P10DT5H6M7S', 'P10DT6M7S')::iday?<6>) = 'P20DT11H6M7S'::iday", expTestStr: "f11(('P10DT5H6M7S', 'P10DT6M7S')::interval_day?<6>) = 'P20DT11H6M7S'::interval_day<0>"},
		{testCaseStr: "((20, 20), (-3, -3), (1, 1), (10,10), (5,5)) count_star() = 1::fp64", expTestStr: "(('20', '20'), ('-3', '-3'), ('1', '1'), ('10', '10'), ('5', '5')) count_star() = 1::fp64"}, // no type specified for columns
		{testCaseStr: "((20), (3), (1), (10), (5)) count_star() = 1::fp64", expTestStr: "(('20'), ('3'), ('1'), ('10'), ('5')) count_star() = 1::fp64"},                                               // no type specified for columns in the test case
		{testCaseStr: `DEFINE t1(fp32, fp32) = ((20, 20), (-3, -3), (1, 1), (10,10), (5,5))
count_star() = 1::fp64`, expTestStr: "((20, 20), (-3, -3), (1, 1), (10, 10), (5, 5)) count_star() = 1::fp64"},
		{testCaseStr: `DEFINE t1(varchar<5>) = (('cat'), ('bat'), ('rat'), (null))
count_star() = 1::fp64`, expTestStr: "(('cat'), ('bat'), ('rat'), (null)) count_star() = 1::fp64"}, // no arguments, so no type info in the output format
		{testCaseStr: `DEFINE t1(varchar<5>) = (('cat'), ('bat'), ('rat'), (null))
count(t1.col0) = 4::fp64`, expTestStr: "(('cat'), ('bat'), ('rat'), (null)) count(col0::varchar?<5>) = 4::fp64"},
		{testCaseStr: "f20(('abcd', 'ef')::fchar?<9>) = Null::fchar<9>", expTestStr: "f20(('abcd', 'ef')::fixedchar?<9>) = null::fixedchar?<9>"},
		{testCaseStr: "f20(('abcd', 'ef')::fixedchar<9>) = Null::fchar<9>", expTestStr: "f20(('abcd', 'ef')::fixedchar<9>) = null::fixedchar?<9>"},
		{testCaseStr: "f20(('abcd', null)::fixedchar<9>) = Null::fchar<9>", expTestStr: "f20(('abcd', null)::fixedchar?<9>) = null::fixedchar?<9>"},
		{testCaseStr: "f20(('abcd', 'ef', null)::vchar?<9>) = Null::vchar<9>", expTestStr: "f20(('abcd', 'ef', null)::varchar?<9>) = null::varchar?<9>"},
		{testCaseStr: "f20(('abcd', 'ef')::varchar<9>) = Null::vchar<9>", expTestStr: "f20(('abcd', 'ef')::varchar<9>) = null::varchar?<9>"},
		{testCaseStr: "f20(('abcd', 'ef')::fbin<9>) = Null::fbin<9>", expTestStr: "f20(('abcd', 'ef')::fixedbinary<9>) = null::fixedbinary?<9>"},
		{testCaseStr: "f20(('abcd', 'ef')::varchar?<9>) = 'abcdef'::varchar<9>", expTestStr: "f20(('abcd', 'ef')::varchar?<9>) = 'abcdef'::varchar<9>"},
		{testCaseStr: "f20(('abcd', null)::fixedchar?<9>) = Null::fixedchar<9>", expTestStr: "f20(('abcd', null)::fixedchar?<9>) = null::fixedchar?<9>"},
		{testCaseStr: "f20(('abcd', 'ef')::fixedbinary?<9>) = Null::fixedbinary<9>", expTestStr: "f20(('abcd', 'ef')::fixedbinary?<9>) = null::fixedbinary?<9>"},
		{testCaseStr: "f35(('1991-01-01T01:02:03.456')::pts?<3>) = '1991-01-01T01:02:30.123123'::precision_timestamp<3>",
			expTestStr: "f35(('1991-01-01T01:02:03.456')::precision_timestamp?<3>) = '1991-01-01T01:02:30.123'::precision_timestamp<3>"},
		{testCaseStr: "f36(('1991-01-01T01:02:03.456', '1991-01-01T01:02:30.123123')::precision_timestamp<3>) = 123456::i64"},
		{testCaseStr: "f37(('1991-01-01T01:02:03.123456', '1991-01-01T04:05:06.456')::precision_timestamp<6>) = 123456::i64"},
		{testCaseStr: "f38(('1991-01-01T01:02:03.456+05:30')::ptstz?<3>) = '1991-01-01T00:00:00+15:30'::precision_timestamp_tz<3>",
			expTestStr: "f38(('1990-12-31T19:32:03.456')::precision_timestamp_tz?<3>) = '1990-12-31T08:30:00.000+00:00'::precision_timestamp_tz<3>"},
		{testCaseStr: "f39(('1991-01-01T01:02:03.456+05:30', '1991-01-01T01:02:03.123456+05:30')::ptstz<6>) = '1991-01-01T00:00:00+15:30'::ptstz<6>",
			expTestStr: "f39(('1990-12-31T19:32:03.456', '1990-12-31T19:32:03.123456')::precision_timestamp_tz<6>) = '1990-12-31T08:30:00.000+00:00'::precision_timestamp_tz<6>"},
	}
	for _, test := range tests {
		t.Run(test.testCaseStr, func(t *testing.T) {
			testFile, err := ParseTestCasesFromString(header + test.testCaseStr)
			require.NoError(t, err)
			require.NotNil(t, testFile)
			assert.Len(t, testFile.TestCases, 1)
			if test.expTestStr == "" {
				test.expTestStr = test.testCaseStr
			}
			assert.Equal(t, test.expTestStr, testFile.TestCases[0].String())
		})
	}
}

func TestParseTestCaseFile(t *testing.T) {
	fs := substrait.GetSubstraitTestsFS()
	testFile, err := ParseTestCaseFileFromFS(fs, "tests/cases/arithmetic/add.test")
	require.NoError(t, err)
	require.NotNil(t, testFile)
	assert.Len(t, testFile.TestCases, 15)

	testFile, err = ParseTestCaseFileFromFS(fs, "tests/cases/arithmetic/max.test")
	require.NoError(t, err)
	require.NotNil(t, testFile)
	assert.Len(t, testFile.TestCases, 12)

	testFile, err = ParseTestCaseFileFromFS(fs, "tests/cases/arithmetic_decimal/power.test")
	require.NoError(t, err)
	require.NotNil(t, testFile)
	assert.Len(t, testFile.TestCases, 9)

	testFile, err = ParseTestCaseFileFromFS(fs, "tests/cases/datetime/lt_datetime.test")
	require.NoError(t, err)
	require.NotNil(t, testFile)
	assert.Len(t, testFile.TestCases, 13)
}

func TestLoadAllSubstraitTestFiles(t *testing.T) {
	got := substrait.GetSubstraitTestsFS()
	filePaths, err := listFiles(got, ".")
	require.NoError(t, err)
	assert.GreaterOrEqual(t, len(filePaths), 107)

	for _, filePath := range filePaths {
		t.Run(filePath, func(t *testing.T) {
			switch filePath {
			case "tests/cases/datetime/extract.test":
				// TODO deal with enum arguments in testcase
				t.Skip("Skipping extract.test")
			}

			testFile, err := ParseTestCaseFileFromFS(got, filePath)
			require.NoError(t, err)
			require.NotNil(t, testFile)
			reg, funcRegistry := functions.NewExtensionAndFunctionRegistries(extensions.GetDefaultCollectionWithNoError())
			for _, tc := range testFile.TestCases {
				testGetFunctionInvocation(t, tc, &reg, funcRegistry)
			}
		})
	}
}

func testGetFunctionInvocation(t *testing.T, tc *TestCase, reg *expr.ExtensionRegistry, registry functions.FunctionRegistry) {
	switch tc.FuncType {
	case ScalarFuncType:
		invocation, err := tc.GetScalarFunctionInvocation(reg, registry)
		require.NoError(t, err, "GetScalarFunctionInvocation failed with error in test case: %s", tc.CompoundFunctionName())
		require.Equal(t, tc.ID().URI, invocation.ID().URI)
		argTypes := invocation.GetArgTypes()
		require.Equal(t, tc.GetArgTypes(), argTypes, "unexpected arg types in test case: %s", tc.CompoundFunctionName())
	case AggregateFuncType:
		invocation, err := tc.GetAggregateFunctionInvocation(reg, registry)
		require.NoError(t, err, "GetAggregateFunctionInvocation failed with error in test case: %s", tc.CompoundFunctionName())
		require.Equal(t, tc.ID().URI, invocation.ID().URI)
		argTypes := invocation.GetArgTypes()
		require.Equal(t, tc.GetArgTypes(), argTypes, "unexpected arg types in test case: %s", tc.CompoundFunctionName())
	}
}

func listFiles(embedFs embed.FS, root string) ([]string, error) {
	var files []string
	err := fs.WalkDir(embedFs, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
