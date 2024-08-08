// SPDX-License-Identifier: Apache-2.0

package functions

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

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
- name: arithmetic.max
  aggregate: true
  supported_kernels:
  - i32
  - i64
aggregate_functions:
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
	dialectName := t.Name()
	localRegistry := getLocalFunctionRegistry(t, dialectYaml)

	t.Run("GetDialect", func(t *testing.T) {
		dialect, err := GetDialect(dialectName)
		assert.Nil(t, err)
		assert.NotNil(t, dialect)
		assert.Equal(t, dialectName, dialect.Name())

		dialect, err = GetDialect("non-existent")
		assert.Error(t, err)
		assert.Nil(t, dialect)
	})

	t.Run("GetSupportedDialects", func(t *testing.T) {
		dialects := GetSupportedDialects()
		assert.Len(t, dialects, 1)
		assert.Contains(t, dialects, dialectName)

		err := LoadDialect("second_dialect", strings.NewReader(dialectYaml))
		assert.NoError(t, err)
		dialects = GetSupportedDialects()
		assert.Len(t, dialects, 2)
		assert.Contains(t, dialects, "second_dialect")
		assert.Contains(t, dialects, dialectName)
	})

	t.Run("LoadDialectError", func(t *testing.T) {
		err := LoadDialect(dialectName, strings.NewReader(dialectYaml))
		assert.Error(t, err)
	})

	t.Run("GetDialect", func(t *testing.T) {
		dialect, err := GetDialect(dialectName)
		assert.Nil(t, err)
		assert.NotNil(t, dialect)
		assert.Equal(t, dialectName, dialect.Name())

		dialect = localRegistry.GetDialect()
		assert.NotNil(t, dialect)
		assert.Equal(t, dialectName, dialect.Name())
	})
}

func getLocalFunctionRegistry(t *testing.T, dialectYaml string) LocalFunctionRegistry {
	err := LoadDialect(t.Name(), strings.NewReader(dialectYaml))
	assert.NoError(t, err)
	registry := GetDefaultFunctionRegistry()
	dialect, err := GetDialect(t.Name())
	assert.NoError(t, err)
	localRegistry, err := dialect.LocalizeFunctionRegistry(registry)
	return localRegistry
}

func TestScalarFunctionsLookup(t *testing.T) {
	expectedUri := "https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml"
	comparisionUri := "https://github.com/substrait-io/substrait/blob/main/extensions/functions_comparison.yaml"
	registry := GetDefaultFunctionRegistry()
	allFunctions := registry.GetAllFunctions()

	dialectYaml := `
name: test
type: sql
dependencies:
  arithmetic: 
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml
  comparison: 
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_comparison.yaml
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
- name: comparison.is_null
  local_name: IS NULL
  postfix: true
  supported_kernels:
  - any1
`
	localRegistry := getLocalFunctionRegistry(t, dialectYaml)
	tests := []struct {
		localName        string
		substraitName    string
		expectedUri      string
		expectedNames    []string
		expectedNotation FunctionNotation
		isOverflowError  bool
	}{
		{"+", "add", expectedUri, []string{"add:i32_i32", "add:i64_i64", "add:fp32_fp32", "add:fp64_fp64"}, INFIX, true},
		{"-", "subtract", expectedUri, []string{"subtract:fp32_fp32", "subtract:fp64_fp64"}, INFIX, true},
		{"IS NULL", "is_null", comparisionUri, []string{"is_null:any"}, POSTFIX, false},
	}
	for _, tt := range tests {
		t.Run(tt.substraitName, func(t *testing.T) {
			var fv []*LocalScalarFunctionVariant
			fv = localRegistry.GetScalarFunctionsBy(tt.localName, Local)

			assert.Greater(t, len(fv), 0)
			assert.Equal(t, tt.expectedUri, fv[0].URI())
			assert.Equal(t, tt.localName, fv[0].LocalName())
			assert.Equal(t, tt.expectedNotation, fv[0].Notation())
			assert.Equal(t, tt.isOverflowError, fv[0].IsOptionSupported("overflow", "ERROR"))
			assert.False(t, fv[0].IsOptionSupported("overflow", "SILENT"))
			checkCompoundNames(t, getScalarCompoundNames(fv), tt.expectedNames)

			fv = localRegistry.GetScalarFunctionsBy(tt.substraitName, Substrait)
			assert.Greater(t, len(fv), 0)
			assert.Equal(t, tt.expectedUri, fv[0].URI())
			assert.Equal(t, tt.localName, fv[0].LocalName())
			assert.Equal(t, tt.substraitName, fv[0].Name())
			checkCompoundNames(t, getScalarCompoundNames(fv), tt.expectedNames)

			winFunctions := registry.GetScalarFunctions(tt.substraitName)
			assert.Greater(t, len(winFunctions), 0)
			for _, f := range winFunctions {
				assert.Contains(t, allFunctions, f)
			}
		})
	}
}

func TestAggregateFunctionsLookup(t *testing.T) {
	expectedUri := "https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml"
	registry := GetDefaultFunctionRegistry()
	allFunctions := registry.GetAllFunctions()
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
		localName        string
		substraitName    string
		expectedUri      string
		expectedNames    []string
		expectedNotation FunctionNotation
	}{
		{"max", "max", expectedUri, []string{"max:i8", "max:i16", "max:fp32", "max:fp64"}, PREFIX},
		{"min", "min", expectedUri, []string{"min:i8", "min:i16", "min:fp32", "min:fp64"}, PREFIX},
	}
	for _, tt := range tests {
		t.Run(tt.substraitName, func(t *testing.T) {
			av := localRegistry.GetAggregateFunctionsBy(tt.localName, Local)

			assert.Greater(t, len(av), 0)
			assert.Equal(t, tt.expectedUri, av[0].URI())
			assert.Equal(t, tt.localName, av[0].LocalName())
			assert.Equal(t, tt.expectedNotation, av[0].Notation())
			assert.False(t, av[0].IsOptionSupported("overflow", "ERROR"))
			checkCompoundNames(t, getAggregateCompoundNames(av), tt.expectedNames)

			av = localRegistry.GetAggregateFunctionsBy(tt.substraitName, Substrait)
			assert.Greater(t, len(av), 0)
			assert.Equal(t, tt.expectedUri, av[0].URI())
			assert.Equal(t, tt.substraitName, av[0].LocalName())
			checkCompoundNames(t, getAggregateCompoundNames(av), tt.expectedNames)

			winFunctions := registry.GetAggregateFunctions(tt.substraitName)
			assert.Greater(t, len(winFunctions), 0)
			for _, f := range winFunctions {
				assert.Contains(t, allFunctions, f)
			}
		})
	}
}

func TestWindowFunctionsLookup(t *testing.T) {
	expectedUri := "https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml"
	registry := GetDefaultFunctionRegistry()
	allFunctions := registry.GetAllFunctions()
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
		localName        string
		substraitName    string
		expectedUri      string
		expectedNames    []string
		expectedNotation FunctionNotation
	}{
		{"ntile", "ntile", expectedUri, []string{"ntile:i32", "ntile:i64"}, PREFIX},
		{"rank", "rank", expectedUri, []string{"rank:", "rank:"}, PREFIX},
	}
	for _, tt := range tests {
		t.Run(tt.substraitName, func(t *testing.T) {
			wf := localRegistry.GetWindowFunctionsBy(tt.localName, Local)
			assert.Greater(t, len(wf), 0)
			assert.Equal(t, tt.expectedUri, wf[0].URI())
			assert.Equal(t, tt.localName, wf[0].LocalName())
			assert.Equal(t, tt.expectedNotation, wf[0].Notation())
			assert.False(t, wf[0].IsOptionSupported("overflow", "ERROR"))
			checkCompoundNames(t, getWindowCompoundNames(wf), tt.expectedNames)

			wf = localRegistry.GetWindowFunctionsBy(tt.substraitName, Substrait)
			assert.Greater(t, len(wf), 0)
			assert.Equal(t, tt.expectedUri, wf[0].URI())
			assert.Equal(t, tt.substraitName, wf[0].LocalName())
			checkCompoundNames(t, getWindowCompoundNames(wf), tt.expectedNames)

			winFunctions := registry.GetWindowFunctions(tt.substraitName)
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
