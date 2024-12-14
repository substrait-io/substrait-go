package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait"
	"github.com/substrait-io/substrait-go/expr"
	"github.com/substrait-io/substrait-go/literal"
	"github.com/substrait-io/substrait-go/types"
)

func makeHeader(version, include string) string {
	return fmt.Sprintf("### SUBSTRAIT_SCALAR_TEST: %s\n### SUBSTRAIT_INCLUDE: '%s'\n\n", version, include)
}

func makeAggregateTestHeader(version, include string) string {
	return fmt.Sprintf("### SUBSTRAIT_AGGREGATE_TEST: %s\n### SUBSTRAIT_INCLUDE: '%s'\n\n", version, include)
}

func TestParseBasicExample(t *testing.T) {
	header := makeHeader("v1.0", "/extensions/functions_arithmetic.yaml")
	tests := `# 'Basic examples without any special cases'
add(120::i8, 5::i8) = 125::i8
add(100::i16, 100::i16) = 200::i16

# Overflow examples demonstrating overflow behavior
add(120::i8, 10::i8) [overflow:ERROR] = <!ERROR>
`

	testFile, err := ParseTestCasesFromString(header + tests)
	require.NoError(t, err)
	assert.Len(t, testFile.TestCases, 3)
}

func TestParseDataTimeExample(t *testing.T) {
	header := makeHeader("v1.0", "/extensions/functions_datetime.yaml")
	tests := `# timestamp examples using the timestamp type
lt('2016-12-31T13:30:15'::ts, '2017-12-31T13:30:15'::ts) = true::bool
`
	testFile, err := ParseTestCasesFromString(header + tests)
	require.NoError(t, err)
	require.NotNil(t, testFile)
	assert.Len(t, testFile.TestCases, 1)
	assert.Equal(t, "lt", testFile.TestCases[0].FuncName)

	assert.Equal(t, testFile.TestCases[0].BaseURI, "/extensions/functions_datetime.yaml")
	assert.Contains(t, testFile.TestCases[0].GroupDesc, "timestamp examples using the timestamp type")
	assert.Len(t, testFile.TestCases[0].Args, 2)
	tsLiteral, err := literal.NewTimestampFromString("2016-12-31T13:30:15")
	require.NoError(t, err)
	assert.Equal(t, tsLiteral, testFile.TestCases[0].Args[0].Value)
	tsLiteral, err = literal.NewTimestampFromString("2017-12-31T13:30:15")
	require.NoError(t, err)
	assert.Equal(t, tsLiteral, testFile.TestCases[0].Args[1].Value)
	boolLiteral := literal.NewBool(true)
	assert.Equal(t, boolLiteral, testFile.TestCases[0].Result.Value)
	assert.Equal(t, &types.BooleanType{}, testFile.TestCases[0].Result.Type)
	timestampType := &types.TimestampType{Nullability: types.NullabilityUnspecified}
	assert.Equal(t, timestampType, testFile.TestCases[0].Args[0].Type)
	assert.Equal(t, timestampType, testFile.TestCases[0].Args[1].Type)
}

func TestParseDecimalExample(t *testing.T) {
	header := makeHeader("v1.0", "extensions/functions_arithmetic_decimal.yaml")
	tests := `# basic
power(8::dec<38,0>, 2::dec<38, 0>) = 64::fp64
power(1.0::dec<38, 5>, -1.0::dec<38, 5>) = 1.0::fp64
power(-1::dec, 0.5::dec<38,1>) [complex_number_result:NAN] = nan::fp64
`

	testFile, err := ParseTestCasesFromString(header + tests)
	require.NoError(t, err)
	require.NotNil(t, testFile)
	assert.Len(t, testFile.TestCases, 3)
	assert.Equal(t, "power", testFile.TestCases[0].FuncName)
	assert.Equal(t, "power", testFile.TestCases[1].FuncName)
	assert.Equal(t, "power", testFile.TestCases[2].FuncName)
	dec8, _ := literal.NewDecimalFromString("8")
	dec2, _ := literal.NewDecimalFromString("2")
	dec1, _ := literal.NewDecimalFromString("1.0")
	decMinus1, _ := literal.NewDecimalFromString("-1")
	decMinus1Point0, _ := literal.NewDecimalFromString("-1.0")
	decPoint5, _ := literal.NewDecimalFromString("0.5")
	f6464 := literal.NewFloat64(64)
	f641 := literal.NewFloat64(1)
	assert.Equal(t, dec8, testFile.TestCases[0].Args[0].Value)
	assert.Equal(t, dec2, testFile.TestCases[0].Args[1].Value)
	assert.Equal(t, f6464, testFile.TestCases[0].Result.Value)
	assert.Equal(t, dec1, testFile.TestCases[1].Args[0].Value)
	assert.Equal(t, decMinus1Point0, testFile.TestCases[1].Args[1].Value)
	assert.Equal(t, f641, testFile.TestCases[1].Result.Value)
	assert.Equal(t, decMinus1, testFile.TestCases[2].Args[0].Value)
	assert.Equal(t, decPoint5, testFile.TestCases[2].Args[1].Value)
	assert.Equal(t, "fp64(NaN)", testFile.TestCases[2].Result.Value.String())
}

func TestParseTestWithVariousTypes(t *testing.T) {
	header := makeHeader("v1.0", "extensions/functions_arithmetic_decimal.yaml") + "# basic \n"
	tests := []struct {
		testCaseStr string
	}{
		{"f1(1::i8, 2::i16, 3::i32, 4::i64) = -7.0::fp32"},
		{"f2(1.0::fp32, 2.0::fp64) = -7.0::fp32"},
		{"f3('a'::str, 'b'::string) = 'c'::str"},
		{"f4(false::bool, true::boolean) = false::bool"},
		{"f5(1.1::dec, 2.2::decimal) = 3.3::dec"},
		{"f6(1.1::dec<38,10>, 2.2::dec<38,10>) = 3.3::dec<38,10>"},
		{"f7(1.1::dec<38,10>, 2.2::decimal<38,10>) = 3.3::decimal<38,10>"},
		{"f8('1991-01-01'::date) = '2001-01-01'::date"},
		{"f8('13:01:01.2345678'::time) = 123456::i64"},
		{"f8('13:01:01.234'::time) = 123::i32"},
		{"f8('1991-01-01T01:02:03.456'::ts, '1991-01-01T00:00:00'::timestamp) = '1991-01-01T22:33:44'::ts"},
		{"f8('1991-01-01T01:02:03.456+05:30'::tstz, '1991-01-01T00:00:00+15:30'::timestamp_tz) = 23::i32"},
		{"f9('1991-01-01'::date, 5::i64) = '1991-01-01T00:00:00+15:30'::timestamp_tz"},
		{"f10('P10Y5M'::interval_year, 5::i64) = 'P15Y5M'::interval_year"},
		{"f10('P10Y5M'::iyear, 5::i64) = 'P15Y5M'::iyear"},
		{"f11('P10DT5H6M7S'::interval_day, 5::i64) = 'P10DT10H6M7S'::interval_day"},
		{"f11('P10DT6M7S'::interval_day, 5::i64) = 'P10DT11M7S'::interval_day"},
		{"or(false::bool, null::bool) = null::bool"},
		{"f12('a'::vchar<9>, 'b'::varchar<4>) = 'c'::varchar<3>"},
		{"f8('1991-01-01T01:02:03.456'::pts<3>, '1991-01-01T00:00:00.000000'::pts<6>) = '1991-01-01T22:33:44'::pts<0>"},
		{"f8('1991-01-01T01:02:03.456+05:30'::ptstz<3>, '1991-01-01T00:00:00+15:30'::ptstz<0>) = '1991-01-01T22:33:44+15:30'::ptstz<0>"},
		//{"f12('P10DT6M7.2000S'::iday<4>, 5::i64) = 'P10DT11M7.2000S'::iday<4>"},  // TODO enable after fixing the grammar
		{"f12('P10DT6M7S'::interval_day, 5::i64) = 'P10DT11M7S'::interval_day"},
		{"concat('abcd'::varchar<9>, Null::str) [null_handling:ACCEPT_NULLS] = Null::str"},
		{"concat('abcd'::vchar<9>, 'ef'::varchar<9>) = Null::vchar<9>"},
		{"concat('abcd'::vchar<9>, 'ef'::fixedchar<9>) = Null::fchar<9>"},
		{"concat('abcd'::fbin<9>, 'ef'::fixedbinary<9>) = Null::fbin<9>"},
		{"f35('1991-01-01T01:02:03.456'::pts<3>) = '1991-01-01T01:02:30.123123'::precision_timestamp<3>"},
		{"f36('1991-01-01T01:02:03.456'::pts<3>, '1991-01-01T01:02:30.123123'::precision_timestamp<3>) = 123456::i64"},
		{"f37('1991-01-01T01:02:03.123456'::pts<6>, '1991-01-01T04:05:06.456'::precision_timestamp<6>) = 123456::i64"},
		{"f38('1991-01-01T01:02:03.456+05:30'::ptstz<3>) = '1991-01-01T00:00:00+15:30'::precision_timestamp_tz<3>"},
		{"f39('1991-01-01T01:02:03.123456+05:30'::ptstz<6>) = '1991-01-01T00:00:00+15:30'::precision_timestamp_tz<6>"},
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

	strAbc := literal.NewString("abc")
	strDef := literal.NewString("def")
	strRes := literal.NewString("abcdef")
	assert.Equal(t, strAbc, testFile.TestCases[0].Args[0].Value)
	assert.Equal(t, strDef, testFile.TestCases[0].Args[1].Value)
	assert.Equal(t, strRes, testFile.TestCases[0].Result.Value)

	strArg1 := literal.NewString("HHHelloooo")
	strArg2 := literal.NewString("Hel+")
	//listStr := &types.ListType{Type: &types.StringType{}}
	assert.Equal(t, strArg1, testFile.TestCases[1].Args[0].Value)
	assert.Equal(t, strArg2, testFile.TestCases[1].Args[1].Value)
	strRes1 := literal.NewString("HH")
	strRes2 := literal.NewString("oooo")
	result, _ := literal.NewList([]expr.Literal{strRes1, strRes2})
	assert.Equal(t, result, testFile.TestCases[1].Result.Value)

	str1 := literal.NewString("à")
	i642 := literal.NewInt64(2)
	assert.Equal(t, str1, testFile.TestCases[2].Args[0].Value)
	assert.Equal(t, i642, testFile.TestCases[2].Result.Value)

	str2 := literal.NewString("😄")
	i644 := literal.NewInt64(4)
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
	strAbc := literal.NewString("abc")
	strDef := literal.NewString("def")
	assert.Equal(t, strAbc, testFile.TestCases[0].Args[0].Value)
	assert.Equal(t, strDef, testFile.TestCases[0].Args[1].Value)
	i8List := &types.ListType{Type: &types.Int8Type{}}
	list, _ := literal.NewList([]expr.Literal{
		literal.NewInt8(1), literal.NewInt8(2), literal.NewInt8(3),
		literal.NewInt8(4), literal.NewInt8(5), literal.NewInt8(6),
	})
	assert.Equal(t, list, testFile.TestCases[0].Result.Value)
	assert.Equal(t, i8List, testFile.TestCases[0].Result.Type)
}

func TestScalarOptions(t *testing.T) {
	header := makeHeader("v1.0", "extensions/functions_string.yaml")
	tests := `# stuff
contains('😊a😊b😊😊'::str, 'A😊B'::str) [case_sensitivity:CASE_INSENSITIVE] = true::bool`

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
	header := makeAggregateTestHeader("v1.0", "extensions/functions_arithmetic.yaml")
	tests := `# basic
avg((1,2,3)::fp32) = 2::fp64
sum((9223372036854775806, 1, 1, 1, 1, 10000000000)::i64) [overflow:ERROR] = <!ERROR>`

	testFile, err := ParseTestCasesFromString(header + tests)
	require.NoError(t, err)
	require.NotNil(t, testFile)
	assert.Len(t, testFile.TestCases, 2)
	assert.Equal(t, "avg", testFile.TestCases[0].FuncName)
	assert.Contains(t, testFile.TestCases[0].GroupDesc, "basic")
	assert.Equal(t, testFile.TestCases[0].BaseURI, "extensions/functions_arithmetic.yaml")
	assert.Len(t, testFile.TestCases[0].Args, 0)
	assert.Len(t, testFile.TestCases[0].AggregateArgs, 1)
	assert.Equal(t, "fp32", testFile.TestCases[0].AggregateArgs[0].ColumnType.String())
	listType := &types.ListType{
		Type:        &types.Float32Type{Nullability: types.NullabilityRequired},
		Nullability: types.NullabilityRequired,
	}
	assert.Equal(t, newFloat32List(1, 2, 3), testFile.TestCases[0].AggregateArgs[0].Argument.Value)
	assert.Equal(t, listType, testFile.TestCases[0].AggregateArgs[0].Argument.Value.GetType())
	assert.Equal(t, "fp64", testFile.TestCases[0].Result.Type.String())
	assert.Equal(t, literal.NewFloat64(2), testFile.TestCases[0].Result.Value)

	assert.Equal(t, "sum", testFile.TestCases[1].FuncName)
	assert.Contains(t, testFile.TestCases[1].GroupDesc, "basic")
	assert.Equal(t, testFile.TestCases[1].BaseURI, "extensions/functions_arithmetic.yaml")
	assert.Len(t, testFile.TestCases[1].Args, 0)
	assert.Len(t, testFile.TestCases[1].AggregateArgs, 1)
	assert.Equal(t, "i64", testFile.TestCases[1].AggregateArgs[0].ColumnType.String())
	assert.Equal(t, newInt64List(9223372036854775806, 1, 1, 1, 1, 10000000000), testFile.TestCases[1].AggregateArgs[0].Argument.Value)
	assert.Equal(t, "ERROR", testFile.TestCases[1].Options["overflow"])
}

func newInt64List(values ...int64) interface{} {
	list, _ := literal.NewList(newInt64Values(values...))
	return list
}

func newInt64Values(values ...int64) []expr.Literal {
	literals := make([]expr.Literal, len(values))
	for i, v := range values {
		literals[i] = literal.NewInt64(v)
	}
	return literals
}

func newFloat32List(values ...float32) expr.Literal {
	list, _ := literal.NewList(newFloat32Values(values...))
	return list
}

func newFloat32Values(values ...float32) []expr.Literal {
	literals := make([]expr.Literal, len(values))
	for i, v := range values {
		literals[i] = literal.NewFloat32(v)
	}
	return literals
}

func TestParseAggregateFuncCompact(t *testing.T) {
	header := makeAggregateTestHeader("v1.0", "extensions/functions_arithmetic.yaml")
	tests := `# basic
((20, 20), (-3, -3), (1, 1), (10,10), (5,5)) corr(col0::fp32, col1::fp32) = 1::fp64
`

	testFile, err := ParseTestCasesFromString(header + tests)
	require.NoError(t, err)
	require.NotNil(t, testFile)
	assert.Len(t, testFile.TestCases, 1)
	assert.Equal(t, "corr", testFile.TestCases[0].FuncName)
	assert.Contains(t, testFile.TestCases[0].GroupDesc, "basic")
	assert.Equal(t, testFile.TestCases[0].BaseURI, "extensions/functions_arithmetic.yaml")
	assert.Len(t, testFile.TestCases[0].Args, 0)
	assert.Len(t, testFile.TestCases[0].AggregateArgs, 2)
	assert.Equal(t, newFloat32Values(20, -3, 1, 10, 5), testFile.TestCases[0].Columns[0])
	assert.Equal(t, newFloat32Values(20, -3, 1, 10, 5), testFile.TestCases[0].Columns[1])
	f32Type := &types.Float32Type{}
	args := []*AggregateArgument{
		createAggregateArg(t, "", "col0", f32Type),
		createAggregateArg(t, "", "col1", f32Type),
	}
	assert.Equal(t, args, testFile.TestCases[0].AggregateArgs)
	assert.Equal(t, "fp64", testFile.TestCases[0].Result.Type.String())
	assert.Equal(t, literal.NewFloat64(1), testFile.TestCases[0].Result.Value)
}

func createAggregateArg(t *testing.T, tableName, columnName string, columnType types.Type) *AggregateArgument {
	arg, err := newAggregateArgument(tableName, columnName, columnType)
	require.NoError(t, err)
	return arg
}

func TestParseAggregateFuncWithMultipleArgs(t *testing.T) {
	header := makeAggregateTestHeader("v1.0", "extensions/functions_arithmetic.yaml")
	tests := `# basic
DEFINE t1(fp32, fp32) = ((20, 20), (-3, -3), (1, 1), (10,10), (5,5.5))
corr(t1.col0, t1.col1) = 1::fp64
DEFINE t1(i64, fp32) = ((20, 20), (-3, -3), (1, 1), (10,10), (5,5.5))
corr(t1.col1, t1.col0) = 1::fp64
`

	testFile, err := ParseTestCasesFromString(header + tests)
	require.NoError(t, err)
	require.NotNil(t, testFile)
	assert.Len(t, testFile.TestCases, 2)
	assert.Equal(t, "corr", testFile.TestCases[0].FuncName)
	assert.Contains(t, testFile.TestCases[0].GroupDesc, "basic")
	assert.Equal(t, testFile.TestCases[0].BaseURI, "extensions/functions_arithmetic.yaml")
	assert.Len(t, testFile.TestCases[0].Args, 0)
	assert.Len(t, testFile.TestCases[0].AggregateArgs, 2)
	assert.Equal(t, newFloat32Values(20, -3, 1, 10, 5), testFile.TestCases[0].Columns[0])
	assert.Equal(t, newFloat32Values(20, -3, 1, 10, 5.5), testFile.TestCases[0].Columns[1])
	assert.Equal(t, "col0", testFile.TestCases[0].AggregateArgs[0].ColumnName)
	assert.Equal(t, "col1", testFile.TestCases[0].AggregateArgs[1].ColumnName)

	assert.Equal(t, "corr", testFile.TestCases[1].FuncName)
	assert.Contains(t, testFile.TestCases[1].GroupDesc, "basic")
	assert.Equal(t, testFile.TestCases[1].BaseURI, "extensions/functions_arithmetic.yaml")
	assert.Len(t, testFile.TestCases[1].Args, 0)
	assert.Len(t, testFile.TestCases[1].AggregateArgs, 2)
	assert.Equal(t, newInt64Values(20, -3, 1, 10, 5), testFile.TestCases[1].Columns[0])
	assert.Equal(t, newFloat32Values(20, -3, 1, 10, 5.5), testFile.TestCases[1].Columns[1])
	assert.Equal(t, "col1", testFile.TestCases[1].AggregateArgs[0].ColumnName)
	assert.Equal(t, "col0", testFile.TestCases[1].AggregateArgs[1].ColumnName)
}

func TestParseAggregateFuncWithVariousTypes(t *testing.T) {
	header := makeAggregateTestHeader("v1.0", "extensions/functions_arithmetic.yaml")
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

func TestParseAggregateFuncWithMixedArgs(t *testing.T) {
	header := makeAggregateTestHeader("v1.0", "extensions/functions_arithmetic.yaml")
	tests := `# basic
((20), (-3), (1), (10)) LIST_AGG(col0::fp32, ','::string) = 1::fp64
DEFINE t1(fp32) = ((20), (-3), (1), (10))
LIST_AGG(t1.col0, ','::string) = 1::fp64
`

	testFile, err := ParseTestCasesFromString(header + tests)
	require.NoError(t, err)
	require.NotNil(t, testFile)
	assert.Len(t, testFile.TestCases, 2)
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
	}
	for _, test := range tests {
		t.Run(test.testCaseStr, func(t *testing.T) {
			_, err := ParseTestCasesFromString(header + test.testCaseStr)
			require.Error(t, err)
			expectedErrorMsg := fmt.Sprintf("Syntax error at line 5:%d: %s", test.position, test.errorMsg)
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
	}{
		{"f1((1, 2, 3, 4)::i64) = 10::fp64"},
		{"f1((1, 2, 3, 4)::i16) = 10.0::fp32"},
		{"f1((1, 2, 3, 4)::i32) = 10::i64"},
		{"f2(1.0::fp32, 2.0::fp64) = -7.0::fp32"},
		{"f3(('a', 'b')::string) = 'c'::str"},
		{"f4((false, true)::boolean) = false::bool"},
		{"f5((1.1, 2.2)::fp32) = 3.3::fp32"},
		{"f5((1.1, 2.2)::fp64) = 3.3::fp64"},
		{"f5((1.1, 2.2)::decimal) = 3.3::dec"},
		{"f6((1.1, 2.2)::dec<38,10>) = 3.3::dec<38,10>"},
		{"f7((1.0, 2)::decimal<38,0>) = 3.0::decimal<38,0>"},
		{"f6((1.1, 2.2, null)::dec?<38,10>) = 3.3::dec<38,10>"},
		{"f8(('1991-01-01', '1991-02-02')::date) = '2001-01-01'::date"},
		{"f8(('13:01:01.2345678', '14:01:01.333')::time) = 123456::i64"},
		{"f8('13:01:01.234'::time) = 123::i32"},
		{"f8(('1991-01-01T01:02:03.456', '1991-01-01T00:00:00')::timestamp) = '1991-01-01T22:33:44'::ts"},
		{"f8(('1991-01-01T01:02:03.456+05:30', '1991-01-01T00:00:00+15:30')::tstz) = 23::i32"},
		{"f10(('P10Y5M', 'P11Y5M')::interval_year) = 'P21Y10M'::interval_year"},
		{"f10(('P10Y2M', 'P10Y7M')::iyear) = 'P20Y9M'::iyear"},
		{"f11(('P10DT5H6M7S', 'P10DT6M7S')::interval_day) = 'P20DT11H6M7S'::interval_day"},
		{"f11(('P10DT5H6M7S', 'P10DT6M7S')::iday?) = 'P20DT11H6M7S'::iday"},
		{"f11(('P10DT5H6M7S', 'P10DT6M7S')::iday?<6>) = 'P20DT11H6M7S'::iday"},
		{"((20, 20), (-3, -3), (1, 1), (10,10), (5,5)) count_star() = 1::fp64"},
		{"((20), (3), (1), (10), (5)) count_star() = 1::fp64"},
		{`DEFINE t1(fp32, fp32) = ((20, 20), (-3, -3), (1, 1), (10,10), (5,5))
count_star() = 1::fp64`},
		{"f20(('abcd', 'ef')::fchar?<9>) = Null::fchar<9>"},
		{"f20(('abcd', 'ef')::fixedchar<9>) = Null::fchar<9>"},
		{"f20(('abcd', 'ef', null)::vchar?<9>) = Null::vchar<9>"},
		{"f20(('abcd', 'ef')::varchar<9>) = Null::vchar<9>"},
		{"f20(('abcd', 'ef')::fbin<9>) = Null::fbin<9>"},
		{"f20(('abcd', 'ef')::fixedbinary?<9>) = Null::fixedbinary<9>"},
		{"f35(('1991-01-01T01:02:03.456')::pts?<3>) = '1991-01-01T01:02:30.123123'::precision_timestamp<3>"},
		{"f36(('1991-01-01T01:02:03.456', '1991-01-01T01:02:30.123123')::precision_timestamp<3>) = 123456::i64"},
		{"f37(('1991-01-01T01:02:03.123456', '1991-01-01T04:05:06.456')::precision_timestamp<6>) = 123456::i64"},
		{"f38(('1991-01-01T01:02:03.456+05:30')::ptstz?<3>) = '1991-01-01T00:00:00+15:30'::precision_timestamp_tz<3>"},
		{"f39(('1991-01-01T01:02:03.456+05:30', '1991-01-01T01:02:03.123456+05:30')::ptstz<6>) = '1991-01-01T00:00:00+15:30'::ptstz<6>"},
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
