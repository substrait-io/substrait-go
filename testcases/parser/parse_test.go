package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

	testFile, err := ParseTestCaseFile(header + tests)
	require.NoError(t, err)
	assert.Len(t, testFile.TestCases, 3)
}

func TestParseDataTimeExample(t *testing.T) {
	header := makeHeader("v1.0", "/extensions/functions_datetime.yaml")
	tests := `# timestamp examples using the timestamp type
lt('2016-12-31T13:30:15'::ts, '2017-12-31T13:30:15'::ts) = true::bool
`
	testFile, err := ParseTestCaseFile(header + tests)
	require.NoError(t, err)
	require.NotNil(t, testFile)
	assert.Len(t, testFile.TestCases, 1)
	assert.Equal(t, "lt", testFile.TestCases[0].FuncName)

	assert.Contains(t, testFile.TestCases[0].BaseURI, "/extensions/functions_datetime.yaml")
	assert.Contains(t, testFile.TestCases[0].GroupDesc, "timestamp examples using the timestamp type")
	assert.Len(t, testFile.TestCases[0].Args, 2)
	tsLiteral, err := literal.NewTimestampFromString("2016-12-31T13:30:15")
	require.NoError(t, err)
	assert.Equal(t, tsLiteral, testFile.TestCases[0].Args[0].Value)
	tsLiteral, err = literal.NewTimestampFromString("2017-12-31T13:30:15")
	require.NoError(t, err)
	assert.Equal(t, tsLiteral, testFile.TestCases[0].Args[1].Value)
	boolLiteral, _ := literal.NewBool(true)
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

	testFile, err := ParseTestCaseFile(header + tests)
	require.NoError(t, err)
	require.NotNil(t, testFile)
	assert.Len(t, testFile.TestCases, 3)
	assert.Equal(t, "power", testFile.TestCases[0].FuncName)
	assert.Equal(t, "power", testFile.TestCases[1].FuncName)
	assert.Equal(t, "power", testFile.TestCases[2].FuncName)
	dec8, err := literal.NewDecimalFromString("8")
	dec2, err := literal.NewDecimalFromString("2")
	dec1, err := literal.NewDecimalFromString("1.0")
	decMinus1, err := literal.NewDecimalFromString("-1")
	decMinus1Point0, err := literal.NewDecimalFromString("-1.0")
	decPoint5, err := literal.NewDecimalFromString("0.5")
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
	}
	for _, test := range tests {
		t.Run(test.testCaseStr, func(t *testing.T) {
			testFile, err := ParseTestCaseFile(header + test.testCaseStr)
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
octet_length('😄'::str) = 4::i64`

	testFile, err := ParseTestCaseFile(header + tests)
	require.NoError(t, err)
	require.NotNil(t, testFile)
	assert.Len(t, testFile.TestCases, 4)
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
	result, err := literal.NewList([]expr.Literal{strRes1, strRes2})
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

	testFile, err := ParseTestCaseFile(header + tests)
	require.NoError(t, err)
	require.NotNil(t, testFile)
	assert.Len(t, testFile.TestCases, 1)

	assert.Equal(t, "some_func", testFile.TestCases[0].FuncName)
	strAbc := literal.NewString("abc")
	strDef := literal.NewString("def")
	assert.Equal(t, strAbc, testFile.TestCases[0].Args[0].Value)
	assert.Equal(t, strDef, testFile.TestCases[0].Args[1].Value)
	i8List := &types.ListType{Type: &types.Int8Type{}}
	list, err := literal.NewList([]expr.Literal{
		literal.NewInt8(1), literal.NewInt8(2), literal.NewInt8(3),
		literal.NewInt8(4), literal.NewInt8(5), literal.NewInt8(6),
	})
	assert.Equal(t, list, testFile.TestCases[0].Result.Value)
	assert.Equal(t, i8List, testFile.TestCases[0].Result.Type)
}

func TestParseAggregateFunc(t *testing.T) {
	header := makeAggregateTestHeader("v1.0", "extensions/functions_arithmetic.yaml")
	tests := `# basic
avg((1,2,3)::fp32) = 2::fp64`

	testFile, err := ParseTestCaseFile(header + tests)
	require.NoError(t, err)
	require.NotNil(t, testFile)
	assert.Len(t, testFile.TestCases, 1)
	assert.Equal(t, "avg", testFile.TestCases[0].FuncName)
	assert.Contains(t, testFile.TestCases[0].GroupDesc, "basic")
	assert.Contains(t, testFile.TestCases[0].BaseURI, "extensions/functions_arithmetic.yaml")
	assert.Len(t, testFile.TestCases[0].Args, 0)
	assert.Len(t, testFile.TestCases[0].AggregateArgs, 1)
	assert.Equal(t, "fp32", testFile.TestCases[0].AggregateArgs[0].ColumnType.String())
	listType := &types.ListType{
		Type:        &types.Float32Type{Nullability: types.NullabilityRequired},
		Nullability: types.NullabilityRequired,
	}
	assert.Equal(t, listType, testFile.TestCases[0].AggregateArgs[0].Argument.Value.GetType())
	assert.Equal(t, "fp64", testFile.TestCases[0].Result.Type.String())
	assert.Equal(t, literal.NewFloat64(2), testFile.TestCases[0].Result.Value)
}

func TestParseAggregateFuncCompact(t *testing.T) {
	header := makeAggregateTestHeader("v1.0", "extensions/functions_arithmetic.yaml")
	tests := `# basic
((20, 20), (-3, -3), (1, 1), (10,10), (5,5)) corr(col0::fp32, col1::fp32) = 1::fp64
`

	testFile, err := ParseTestCaseFile(header + tests)
	require.NoError(t, err)
	require.NotNil(t, testFile)
	assert.Len(t, testFile.TestCases, 1)
	assert.Equal(t, "corr", testFile.TestCases[0].FuncName)
	assert.Contains(t, testFile.TestCases[0].GroupDesc, "basic")
	assert.Contains(t, testFile.TestCases[0].BaseURI, "extensions/functions_arithmetic.yaml")
	assert.Len(t, testFile.TestCases[0].Args, 0)
	assert.Len(t, testFile.TestCases[0].AggregateArgs, 2)
}

func TestParseAggregateFuncWithMultipleArgs(t *testing.T) {
	header := makeAggregateTestHeader("v1.0", "extensions/functions_arithmetic.yaml")
	tests := `# basic
DEFINE t1(fp32, fp32) = ((20, 20), (-3, -3), (1, 1), (10,10), (5,5))
corr(t1.col0, t1.col1) = 1::fp64
`

	testFile, err := ParseTestCaseFile(header + tests)
	require.NoError(t, err)
	require.NotNil(t, testFile)
	assert.Len(t, testFile.TestCases, 1)
	assert.Equal(t, "corr", testFile.TestCases[0].FuncName)
	assert.Contains(t, testFile.TestCases[0].GroupDesc, "basic")
	assert.Contains(t, testFile.TestCases[0].BaseURI, "extensions/functions_arithmetic.yaml")
	assert.Len(t, testFile.TestCases[0].Args, 0)
	assert.Len(t, testFile.TestCases[0].AggregateArgs, 2)
	// TODO add checks on row data
}
