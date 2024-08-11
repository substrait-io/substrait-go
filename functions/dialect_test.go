// SPDX-License-Identifier: Apache-2.0

package functions

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/extensions"
)

var gFunctionRegistry FunctionRegistry

func TestMain(m *testing.M) {
	gFunctionRegistry = NewFunctionRegistry(&extensions.DefaultCollection)
	m.Run()
}

func TestDialectApis(t *testing.T) {
	dialectYaml := `
name: testdb
type: sql
dependencies:
  arithmetic: 
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml
  comparison: 
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_comparison.yaml
supported_types:
  i32:
    sql_type_name: TINYINT,
    supported_as_column: true
  i64:
    sql_type_name: TINYINT,
    supported_as_column: true
scalar_functions:
- name: comparison.is_null
  local_name: IS NULL
  postfix: true
  supported_kernels:
  - any1
aggregate_functions:
- name: arithmetic.max
  aggregate: true
  supported_kernels:
  - i32
  - i64
- name: arithmetic.min
  aggregate: true
  supported_kernels:
  - i32
  - i64
window_functions:
- name: arithmetic.ntile
  supported_kernels:
  - i32
- name: arithmetic.rank
  supported_kernels:
  - ""
`
	dialect, err := LoadDialect(t.Name(), strings.NewReader(dialectYaml))
	assert.Nil(t, err)
	assert.NotNil(t, dialect)
	assert.Equal(t, t.Name(), dialect.Name())
	localRegistry, err := dialect.LocalizeFunctionRegistry(gFunctionRegistry)
	assert.NoError(t, err)
	assert.Equal(t, t.Name(), localRegistry.GetDialect().Name())
}

func TestBadDialects(t *testing.T) {
	type testcase struct {
		error       string
		dialectYaml string
	}
	tests := []testcase{
		{`no supported types`, `name: testdb`},
		{`no supported types`, `name: testdb
type: sql
dependencies:
supported_types:
window_functions:
- name: arithmetic.ntile
  supported_kernels:
  - i32
`},
		{`no functions`, `name: testdb
type: sql
dependencies:
supported_types:
  i32:
   sql_type_name: TINYINT,
   supported_as_column: true
`,
		},
		{`unknown dependency`, `name: testdb
type: sql
dependencies:
supported_types:
  i32:
    sql_type_name: TINYINT,
    supported_as_column: true
window_functions:
- name: arithmetic.ntile
  supported_kernels:
  - i32
`},
		{`invalid function name`, `name: testdb
type: sql
dependencies:
supported_types:
  i32:
    sql_type_name: TINYINT,
    supported_as_column: true
aggregate_functions:
- name: max
  supported_kernels:
  - i32
`},
		{`invalid function`, `name: testdb
type: sql
dependencies:
supported_types:
  i32:
    sql_type_name: TINYINT,
    supported_as_column: true
scalar_functions:
- name: arithmetic.add
  supported_kernels:
`},
		{`unsupported type`, `name: testdb
type: sql
dependencies:
dependencies:
  arithmetic: 
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml
supported_types:
  i32:
    sql_type_name: TINYINT,
    supported_as_column: true
window_functions:
- name: arithmetic.unknown_function
  supported_kernels:
  - i99
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.error, func(t *testing.T) {
			_, err := LoadDialect(t.Name(), strings.NewReader(tt.dialectYaml))
			assert.Error(t, err)
			assert.Contains(t, err.Error(), tt.error)
		})
	}

	localizeTestcases := []testcase{
		{`no function variant found`, `name: testdb
type: sql
dependencies:
  arithmetic: 
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml
supported_types:
  i32:
    sql_type_name: TINYINT,
    supported_as_column: true
  fp32:
    sql_type_name: FLOAT,
    supported_as_column: true
scalar_functions:
- name: arithmetic.add
  supported_kernels:
    - i32_i32
    - i32_fp32
`},
		{`no function variant found`, `name: testdb
type: sql
dependencies:
  arithmetic: 
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml
supported_types:
  i32:
    sql_type_name: TINYINT,
    supported_as_column: true
aggregate_functions:
- name: arithmetic.max
  supported_kernels:
    - i32
    - str
`},
		{`no function variant found`, `name: testdb
type: sql
dependencies:
  arithmetic: 
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml
supported_types:
  i32:
    sql_type_name: TINYINT,
    supported_as_column: true
window_functions:
- name: arithmetic.ntile
  supported_kernels:
    - str
    - i64
`},
	}
	for _, tt := range localizeTestcases {
		t.Run(tt.error, func(t *testing.T) {
			dialect, err := LoadDialect(t.Name(), strings.NewReader(tt.dialectYaml))
			assert.NoError(t, err)

			_, err = dialect.LocalizeFunctionRegistry(gFunctionRegistry)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), tt.error)
		})
	}

	badTypeTestcases := []testcase{
		{`unknown type`, `name: testdb
type: sql
dependencies:
dependencies:
  arithmetic: 
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml
supported_types:
  i32:
    sql_type_name: TINYINT,
    supported_as_column: true
  myth:
    sql_type_name: TINYINT,
    supported_as_column: true
scalar_functions:
- name: arithmetic.add
  supported_kernels:
  - i32_i32
`,
		},
	}
	for _, tt := range badTypeTestcases {
		t.Run(tt.error, func(t *testing.T) {
			dialect, err := LoadDialect(t.Name(), strings.NewReader(tt.dialectYaml))
			assert.NoError(t, err)

			typeRegistry := NewTypeRegistry()
			_, err = dialect.LocalizeTypeRegistry(typeRegistry)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), tt.error)
		})
	}
}

func getLocalFunctionRegistry(t *testing.T, dialectYaml string) LocalFunctionRegistry {
	dialect, err := LoadDialect(t.Name(), strings.NewReader(dialectYaml))
	assert.NoError(t, err)
	localRegistry, err := dialect.LocalizeFunctionRegistry(gFunctionRegistry)
	assert.NoError(t, err)
	return localRegistry
}

func TestScalarFunctionsLookup(t *testing.T) {
	baseUri := "https://github.com/substrait-io/substrait/blob/main/extensions/"
	arithmeticUri := baseUri + "functions_arithmetic.yaml"
	booleanUri := baseUri + "functions_boolean.yaml"
	comparisonUri := baseUri + "functions_comparison.yaml"
	stringUri := baseUri + "functions_string.yaml"
	allFunctions := gFunctionRegistry.GetAllFunctions()

	dialectYaml := `
name: test
type: sql
dependencies:
  arithmetic: 
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml
  boolean: 
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_boolean.yaml
  comparison: 
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_comparison.yaml
  string: 
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_string.yaml
supported_types:
  i32:
    sql_type_name: INTEGER
    supported_as_column: true
  i64:
    sql_type_name: BIGINT
    supported_as_column: true
  fp32:
    sql_type_name: REAL
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
  - i32_i32
  - i64_i64
  - fp32_fp32
  - fp64_fp64
- name: arithmetic.subtract
  local_name: '-'
  infix: true
  required_options:
    overflow: ERROR
    rounding: TIE_TO_EVEN
  supported_kernels:
  - fp32_fp32
  - fp64_fp64
- name: boolean.and
  local_name: and
  infix: true
  supported_kernels:
  - bool
- name: comparison.is_null
  local_name: IS NULL
  postfix: true
  supported_kernels:
  - any1
- name: string.concat
  local_name: '||'
  required_options:
    null_handling: ACCEPT_NULLS
  infix: true
  supported_kernels:
  - vchar
  - str
  variadic: 1
`
	localRegistry := getLocalFunctionRegistry(t, dialectYaml)
	tests := []struct {
		numArgs          int
		localName        string
		substraitName    string
		expectedUri      string
		expectedNames    []string
		expectedNotation FunctionNotation
		isOverflowError  bool
	}{
		{2, "+", "add", arithmeticUri, []string{"add:i32_i32", "add:i64_i64", "add:fp32_fp32", "add:fp64_fp64"}, INFIX, true},
		{2, "-", "subtract", arithmeticUri, []string{"subtract:fp32_fp32", "subtract:fp64_fp64"}, INFIX, true},
		{1, "IS NULL", "is_null", comparisonUri, []string{"is_null:any"}, POSTFIX, false},
		{3, "and", "and", booleanUri, []string{"and:bool"}, INFIX, false},
		{2, "and", "and", booleanUri, []string{"and:bool"}, INFIX, false},
		{1, "and", "and", booleanUri, []string{"and:bool"}, INFIX, false},
		{0, "and", "and", booleanUri, []string{"and:bool"}, INFIX, false},
		{2, "||", "concat", stringUri, []string{"concat:vchar", "concat:str"}, INFIX, false},
		{3, "||", "concat", stringUri, []string{"concat:vchar", "concat:str"}, INFIX, false},
	}
	for _, tt := range tests {
		t.Run(tt.substraitName, func(t *testing.T) {
			var fv []*LocalScalarFunctionVariant
			fv = localRegistry.GetScalarFunctions(LocalFunctionName(tt.localName), tt.numArgs)

			assert.Greater(t, len(fv), 0)
			assert.Equal(t, tt.expectedUri, fv[0].URI())
			assert.Equal(t, tt.localName, fv[0].LocalName())
			assert.Equal(t, tt.expectedNotation, fv[0].Notation())
			assert.Equal(t, tt.isOverflowError, fv[0].IsOptionSupported("overflow", "ERROR"))
			assert.False(t, fv[0].IsOptionSupported("overflow", "SILENT"))
			checkCompoundNames(t, getScalarCompoundNames(fv), tt.expectedNames)

			fv = localRegistry.GetScalarFunctions(SubstraitFunctionName(tt.substraitName), tt.numArgs)
			assert.Greater(t, len(fv), 0)
			assert.Equal(t, tt.expectedUri, fv[0].URI())
			assert.Equal(t, tt.localName, fv[0].LocalName())
			assert.Equal(t, tt.substraitName, fv[0].Name())
			checkCompoundNames(t, getScalarCompoundNames(fv), tt.expectedNames)

			scalarFunctions := gFunctionRegistry.GetScalarFunctions(tt.substraitName, tt.numArgs)
			assert.Greater(t, len(scalarFunctions), 0)
			for _, f := range scalarFunctions {
				assert.Contains(t, allFunctions, f)
			}
			scalarFunctions = gFunctionRegistry.GetScalarFunctionsByName(tt.substraitName)
			assert.Greater(t, len(scalarFunctions), 0)
			for _, f := range scalarFunctions {
				assert.Contains(t, allFunctions, f)
			}
		})
	}
}

func TestAggregateFunctionsLookup(t *testing.T) {
	expectedUri := "https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml"
	allFunctions := gFunctionRegistry.GetAllFunctions()
	dialectYaml := `
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
    sql_type_name: BIGINT
    supported_as_column: true
  fp32:
    sql_type_name: REAL
    supported_as_column: true
  fp64:
    sql_type_name: DOUBLE
    supported_as_column: true
aggregate_functions:
- name: arithmetic.min
  aggregate: true
  supported_kernels:
  - i8
  - i16
  - fp32
  - fp64
- name: arithmetic.max
  aggregate: true
  supported_kernels:
  - i8
  - i16
  - fp32
  - fp64
`
	localRegistry := getLocalFunctionRegistry(t, dialectYaml)

	tests := []struct {
		numArgs          int
		localName        string
		substraitName    string
		expectedUri      string
		expectedNames    []string
		expectedNotation FunctionNotation
	}{
		{1, "max", "max", expectedUri, []string{"max:i8", "max:i16", "max:fp32", "max:fp64"}, PREFIX},
		{1, "min", "min", expectedUri, []string{"min:i8", "min:i16", "min:fp32", "min:fp64"}, PREFIX},
	}
	for _, tt := range tests {
		t.Run(tt.substraitName, func(t *testing.T) {
			av := localRegistry.GetAggregateFunctions(LocalFunctionName(tt.localName), 1)

			assert.Greater(t, len(av), 0)
			assert.Equal(t, tt.expectedUri, av[0].URI())
			assert.Equal(t, tt.localName, av[0].LocalName())
			assert.Equal(t, tt.expectedNotation, av[0].Notation())
			assert.False(t, av[0].IsOptionSupported("overflow", "ERROR"))
			checkCompoundNames(t, getAggregateCompoundNames(av), tt.expectedNames)

			av = localRegistry.GetAggregateFunctions(SubstraitFunctionName(tt.substraitName), 1)
			assert.Greater(t, len(av), 0)
			assert.Equal(t, tt.expectedUri, av[0].URI())
			assert.Equal(t, tt.substraitName, av[0].LocalName())
			checkCompoundNames(t, getAggregateCompoundNames(av), tt.expectedNames)

			winFunctions := gFunctionRegistry.GetAggregateFunctions(tt.substraitName, tt.numArgs)
			assert.Greater(t, len(winFunctions), 0)
			for _, f := range winFunctions {
				assert.Contains(t, allFunctions, f)
			}
			winFunctions = gFunctionRegistry.GetAggregateFunctionsByName(tt.substraitName)
			assert.Greater(t, len(winFunctions), 0)
			for _, f := range winFunctions {
				assert.Contains(t, allFunctions, f)
			}
		})
	}
}

func TestWindowFunctionsLookup(t *testing.T) {
	expectedUri := "https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml"
	allFunctions := gFunctionRegistry.GetAllFunctions()
	dialectYaml := `
name: test
type: sql
dependencies:
  arithmetic: 
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml
supported_types:
  i32:
    sql_type_name: INTEGER
    supported_as_column: true
  i64:
    sql_type_name: BIGINT
    supported_as_column: true
window_functions:
- name: arithmetic.ntile
  supported_kernels:
  - i32
  - i64
- name: arithmetic.rank
  supported_kernels:
  - ""
`
	localRegistry := getLocalFunctionRegistry(t, dialectYaml)
	tests := []struct {
		numArgs          int
		localName        string
		substraitName    string
		expectedUri      string
		expectedNames    []string
		expectedNotation FunctionNotation
	}{
		{1, "ntile", "ntile", expectedUri, []string{"ntile:i32", "ntile:i64"}, PREFIX},
		{0, "rank", "rank", expectedUri, []string{"rank:", "rank:"}, PREFIX},
	}
	for _, tt := range tests {
		t.Run(tt.substraitName, func(t *testing.T) {
			wf := localRegistry.GetWindowFunctions(LocalFunctionName(tt.localName), tt.numArgs)
			assert.Greater(t, len(wf), 0)
			assert.Equal(t, tt.expectedUri, wf[0].URI())
			assert.Equal(t, tt.localName, wf[0].LocalName())
			assert.Equal(t, tt.expectedNotation, wf[0].Notation())
			assert.False(t, wf[0].IsOptionSupported("overflow", "ERROR"))
			checkCompoundNames(t, getWindowCompoundNames(wf), tt.expectedNames)

			wf = localRegistry.GetWindowFunctions(SubstraitFunctionName(tt.substraitName), tt.numArgs)
			assert.Greater(t, len(wf), 0)
			assert.Equal(t, tt.expectedUri, wf[0].URI())
			assert.Equal(t, tt.substraitName, wf[0].LocalName())
			checkCompoundNames(t, getWindowCompoundNames(wf), tt.expectedNames)

			winFunctions := gFunctionRegistry.GetWindowFunctions(tt.substraitName, tt.numArgs)
			assert.Greater(t, len(winFunctions), 0)
			for _, f := range winFunctions {
				assert.Contains(t, allFunctions, f)
			}
			winFunctions = gFunctionRegistry.GetWindowFunctionsByName(tt.substraitName)
			assert.Greater(t, len(winFunctions), 0)
			for _, f := range winFunctions {
				assert.Contains(t, allFunctions, f)
			}
		})
	}
}

func getScalarCompoundNames(fv []*LocalScalarFunctionVariant) []string {
	var compoundNames []string
	for _, f := range fv {
		compoundNames = append(compoundNames, f.CompoundName())
	}
	return compoundNames
}

func getAggregateCompoundNames(av []*LocalAggregateFunctionVariant) []string {
	var compoundNames []string
	for _, f := range av {
		compoundNames = append(compoundNames, f.CompoundName())
	}
	return compoundNames
}

func getWindowCompoundNames(wv []*LocalWindowFunctionVariant) []string {
	var compoundNames []string
	for _, f := range wv {
		compoundNames = append(compoundNames, f.CompoundName())
	}
	return compoundNames
}

func checkCompoundNames(t *testing.T, compoundNames []string, expectedNames []string) {
	for _, name := range expectedNames {
		assert.Contains(t, compoundNames, name)
	}
}
