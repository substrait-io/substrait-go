package functions_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v3/extensions"
	"github.com/substrait-io/substrait-go/v3/functions"
	parser2 "github.com/substrait-io/substrait-go/v3/testcases/parser"
)

func makeHeader(version, include string) string {
	return fmt.Sprintf("### SUBSTRAIT_SCALAR_TEST: %s\n### SUBSTRAIT_INCLUDE: '%s'\n\n", version, include)
}

func makeAggregateTestHeader(version, include string) string {
	return fmt.Sprintf("### SUBSTRAIT_AGGREGATE_TEST: %s\n### SUBSTRAIT_INCLUDE: '%s'\n\n", version, include)
}

var scalarFunctionDialectYaml = `
name: test
type: sql
dependencies:
  arithmetic:
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml
supported_types:
  i8:
    sql_type_name: INTEGER
    supported_as_column: true
  i16:
    sql_type_name: INTEGER
    supported_as_column: true
  i32:
    sql_type_name: INTEGER
    supported_as_column: true
  i64:
    sql_type_name: BIGINT
    supported_as_column: true
  fp64:
    sql_type_name: DOUBLE
    supported_as_column: true
  bool:
    sql_type_name: BOOLEAN
    supported_as_column: true
scalar_functions:
- name: arithmetic.add
  local_name: +
  infix: true
  required_options:
    overflow: ERROR
    rounding: TIE_TO_EVEN
  supported_kernels:
  - i8_i8
  - i16_i16
  - i32_i32
  - i64_i64
  - fp64_fp64
`

func TestGetLocalScalarFunctionByInvocation(t *testing.T) {
	header := makeHeader("v1.0", "/extensions/functions_arithmetic.yaml")
	testsStr := `# 'Basic examples without any special cases'
add(120::i8, 5::i8) = 125::i8
add(126::i16, 5::i16) = 125::i16
add(3.4e+38::fp32, 3.4e+38::fp32) = inf::fp32

# Overflow examples demonstrating overflow behavior
add(2000000000::i32, 2000000000::i32) [overflow:ERROR] = <!ERROR>
add(9223372036854775807::i64, 1::i64) [overflow:ERROR] = <!ERROR>
add(120::i8, 10::i8) [overflow:SATURATE] = 127::i8
add(120::i8, 10::i8) [overflow:SILENT] = <!UNDEFINED>

`

	testResults := []struct {
		name          string
		expectedError string
	}{
		{"add:i8_i8", ""},
		{"add:i16_i16", ""},
		{"add:fp32_f32", "function variant not found"},
		{"add:i32_i32 [overflow:ERROR]", ""},
		{"add:i64_i64 [overflow:ERROR]", ""},
		{"add:i8_i8 [overflow:SATURATE]", "unsupported option [overflow:SATURATE]"},
		{"add:i8_i8 [overflow:SILENT]", "unsupported option [overflow:SILENT]"},
	}
	localRegistry := getLocalFunctionRegistry(t, scalarFunctionDialectYaml, gFunctionRegistry)

	testFile, err := parser2.ParseTestCasesFromString(header + testsStr)
	require.NoError(t, err)
	require.NotNil(t, testFile)
	assert.Len(t, testFile.TestCases, len(testResults))
	require.GreaterOrEqual(t, len(testFile.TestCases), len(testResults))

	reg, funcRegistry := functions.NewExtensionAndFunctionRegistries(&extensions.DefaultCollection)
	for i, result := range testResults {
		tc := testFile.TestCases[i]
		t.Run(result.name, func(t *testing.T) {
			require.Equal(t, tc.FuncType, parser2.ScalarFuncType)
			invocation, err := tc.GetScalarFunctionInvocation(&reg, funcRegistry)
			require.NoError(t, err)
			require.Equal(t, tc.ID(), invocation.ID())
			localVariant, err := localRegistry.GetScalarFunctionByInvocation(invocation)
			if result.expectedError == "" {
				require.NoError(t, err)
				require.NotNil(t, localVariant)
			} else {
				require.Error(t, err)
				assert.ErrorContains(t, err, result.expectedError)
			}
		})
	}
}

var aggregateFunctionDialectYaml = `
name: test
type: sql
dependencies:
  arithmetic:
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml
  arithmetic_decimal:
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic_decimal.yaml
  comparison:
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_comparison.yaml
  string:
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_string.yaml
supported_types:
  i8:
    sql_type_name: INTEGER
    supported_as_column: true
  i16:
    sql_type_name: INTEGER
    supported_as_column: true
  i32:
    sql_type_name: INTEGER
    supported_as_column: true
  i64:
    sql_type_name: BIGINT
    supported_as_column: true
  fp64:
    sql_type_name: DOUBLE
    supported_as_column: true
  decimal:
    sql_type_name: NUMBER
    supported_as_column: true
aggregate_functions:
  - name: arithmetic.sum
    aggregate: true
    supported_kernels:
      - fp64
  - name: arithmetic.min
    aggregate: true
    supported_kernels:
      - i16
      - i32
  - name: arithmetic.avg
    aggregate: true
    supported_kernels:
      - fp64
  - name: arithmetic_decimal.min
    aggregate: true
    supported_kernels:
      - dec
  - name: arithmetic_decimal.sum
    aggregate: true
    supported_kernels:
      - dec
`

func TestGetLocalAggregateFunctionByInvocation(t *testing.T) {
	header := makeAggregateTestHeader("v1.0", "/extensions/functions_arithmetic.yaml")
	testsStr := `# basic
avg((1,2,3)::fp32) = 2::fp64
avg((1,2,3)::fp64) = 2::fp64
min((20, -3, 1, -10, 0, 5)::i8) = -10::i8
min((-32768, 32767, 20000, -30000)::i16) = -32768::i16
min((-214748648, 214748647, 21470048, 4000000)::i32) = -214748648::i32
sum((9223372036854775806, 1, 1, 1, 1, 10000000000)::i64) [overflow:ERROR] = <!ERROR>
sum((2.5000007152557373046875, 7.0000007152557373046875, 0, 7.0000007152557373046875)::fp64) = 16.500002145767212::fp64
`
	decHeader := makeAggregateTestHeader("v1.0", "/extensions/functions_arithmetic_decimal.yaml")
	decTestsStr := `# basic
min((20, -3, 1, -10, 0, 5)::dec<2, 0>) = -10::dec<2, 0>
min((-32768, 32767, 20000, -30000)::dec<5, 0>) = -32768::dec<5, 0>
sum((2.5000007152557373046875, 7.0000007152557373046875, 0, 7.0000007152557373046875)::dec<23, 22>) = 16.5000021457672119140625::dec<38, 22>
`

	testResults := []struct {
		name          string
		expectedError string
	}{
		{"avg:fp32", "function variant not found"},
		{"avg:fp64", ""},
		{"min:i8", "function variant not found"},
		{"min:i16", ""},
		{"min:i32", ""},
		{"sum:i64", "function variant not found"},
		{"sum:fp64", ""},
		{"min:dec", ""},
		{"min:dec", ""},
		{"sum:dec", ""},
	}
	localRegistry := getLocalFunctionRegistry(t, aggregateFunctionDialectYaml, gFunctionRegistry)
	testFile, err := parser2.ParseTestCasesFromString(header + testsStr)
	require.NoError(t, err)
	require.NotNil(t, testFile)
	testFile1, err := parser2.ParseTestCasesFromString(decHeader + decTestsStr)
	require.NoError(t, err)
	require.NotNil(t, testFile)
	testCases := append(testFile.TestCases, testFile1.TestCases...)
	require.GreaterOrEqual(t, len(testCases), len(testResults))

	reg, funcRegistry := functions.NewExtensionAndFunctionRegistries(&extensions.DefaultCollection)
	for i, result := range testResults {
		tc := testCases[i]
		t.Run(result.name, func(t *testing.T) {
			require.Equal(t, tc.FuncType, parser2.AggregateFuncType)
			invocation, err := tc.GetAggregateFunctionInvocation(&reg, funcRegistry)
			require.NoError(t, err)
			require.Equal(t, tc.ID(), invocation.ID())
			localVariant, err := localRegistry.GetAggregateFunctionByInvocation(invocation)
			if result.expectedError == "" {
				require.NoError(t, err)
				require.NotNil(t, localVariant)
			} else {
				require.Error(t, err)
				assert.ErrorContains(t, err, result.expectedError)
			}
		})
	}
}
