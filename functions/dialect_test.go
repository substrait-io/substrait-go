// SPDX-License-Identifier: Apache-2.0

package functions

import (
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
)

const sampleDialectYAML = `---
name: duckdb
type: sql
dependencies:
  arithmetic: 
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml
  comparison: 
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_comparison.yaml
supported_types:
  i8:
    sql_type_name: TINYINT,
    supported_as_column: true
  i16:
    sql_type_name: SMALLINT
    supported_as_column: true
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
  str:
    sql_type_name: VARCHAR
    supported_as_column: true
  date:
    sql_type_name: DATE
    supported_as_column: true
  time:
    sql_type_name: TIME
    supported_as_column: true
  ts:
    sql_type_name: TIMESTAMP
    supported_as_column: true
  tstz:
    sql_type_name: TIMESTAMPTZ
    supported_as_column: true
  interval:
    sql_type_name: INTERVAL
    supported_as_column: false
  decimal:
    sql_type_name: DECIMAL
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
  - fp32_fp32
  - fp64_fp64
- name: arithmetic.subtract
  local_name: '-'
  infix: true
  required_options:
    overflow: ERROR
    rounding: TIE_TO_EVEN
  supported_kernels:
  - i8_i8
  - i16_i16
  - i32_i32
  - i64_i64
  - fp32_fp32
  - fp64_fp64
- name: comparison.is_null
  local_name: IS NULL
  postfix: true
  supported_kernels:
  - any1
aggregate_functions:
- name: arithmetic.min
  aggregate: true
  supported_kernels:
  - i8
  - i16
  - i32
  - i64
  - fp32
  - fp64
- name: arithmetic.max
  aggregate: true
  supported_kernels:
  - i8
  - i16
  - i32
  - i64
  - fp32
  - fp64
window_functions:
- name: arithmetic.ntile
  supported_kernels:
  - i32
  - i64
`

var gDialectName string
var gDialect Dialect
var gLocalFunctionRegistry LocalFunctionRegistry

func TestMain(m *testing.M) {
	gDialectName = "dialect_test"
	functionRegistry := GetDefaultFunctionRegistry()
	err := LoadDialect("dialect_test", strings.NewReader(sampleDialectYAML))
	if err != nil {
		log.Fatalf("dialect load failed %v", err)
	}
	gDialect, err = GetDialect(gDialectName)
	if err != nil {
		log.Fatalf("dialect get failed %v", err)
	}
	gLocalFunctionRegistry, err = gDialect.LocalizeFunctionRegistry(functionRegistry)
	if err != nil {
		log.Fatalf("localizeFunctionRegistry failed %v", err)
	}
	m.Run()
}

func TestDialectApis(t *testing.T) {
	dialect, err := GetDialect(gDialectName)
	assert.Nil(t, err)
	assert.NotNil(t, dialect)
	assert.Equal(t, gDialectName, dialect.Name())

	dialect, err = GetDialect("non-existent")
	assert.Error(t, err)
	assert.Nil(t, dialect)
	dialects := GetSupportedDialects()
	assert.Len(t, dialects, 1)
	assert.Contains(t, dialects, gDialectName)

	err = LoadDialect("second_dialect", strings.NewReader(sampleDialectYAML))
	assert.NoError(t, err)
	dialects = GetSupportedDialects()
	assert.Len(t, dialects, 2)
	assert.Contains(t, dialects, "second_dialect")
	assert.Contains(t, dialects, gDialectName)

	dialect, err = GetDialect(gDialectName)
	assert.Nil(t, err)
	assert.NotNil(t, dialect)

	dialect, err = GetDialect("second_dialect")
	assert.Nil(t, err)
	assert.NotNil(t, dialect)

	err = LoadDialect(gDialectName, strings.NewReader(sampleDialectYAML))
	assert.Error(t, err)
}

func TestRegistryLookupByLocal(t *testing.T) {
	expectedUri := "https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml"
	comparisionUri := "https://github.com/substrait-io/substrait/blob/main/extensions/functions_comparison.yaml"
	fv := gLocalFunctionRegistry.GetScalarFunctionsBy("+", Local)
	assert.Greater(t, len(fv), 1)
	assert.Equal(t, expectedUri, fv[0].URI())
	assert.Equal(t, "+", fv[0].LocalName())
	assert.True(t, fv[0].IsOptionSupported("overflow", "ERROR"))
	assert.Equal(t, INFIX, fv[0].Notation())
	checkCompoundNames(t, getScalarCompoundNames(fv), []string{"add:i8_i8", "add:i16_i16", "add:fp32_fp32", "add:fp64_fp64"})

	fv = gLocalFunctionRegistry.GetScalarFunctionsBy("IS NULL", Local)
	assert.Equal(t, comparisionUri, fv[0].URI())
	assert.Equal(t, "IS NULL", fv[0].LocalName())
	assert.Equal(t, POSTFIX, fv[0].Notation())
	checkCompoundNames(t, getScalarCompoundNames(fv), []string{"is_null:any"})

	av := gLocalFunctionRegistry.GetAggregateFunctionsBy("min", Local)
	assert.Equal(t, expectedUri, av[0].URI())
	assert.Equal(t, "min", av[0].LocalName())
	assert.Equal(t, PREFIX, av[0].Notation())
	assert.False(t, av[0].IsOptionSupported("overflow", "ERROR"))
	checkCompoundNames(t, getAggregateCompoundNames(av), []string{"min:i8", "min:i16", "min:fp32", "min:fp64"})

	wv := gLocalFunctionRegistry.GetWindowFunctionsBy("ntile", Local)
	assert.Equal(t, expectedUri, wv[0].URI())
	assert.Equal(t, "ntile", wv[0].LocalName())
	assert.Equal(t, PREFIX, wv[0].Notation())
	assert.False(t, wv[0].IsOptionSupported("overflow", "ERROR"))
	checkCompoundNames(t, getWindowCompoundNames(wv), []string{"ntile:i32", "ntile:i64"})
}

func TestRegistryLookupBySubstrait(t *testing.T) {
	expectedUri := "https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml"
	//b := plan.NewBuilderDefault()
	//sf, err := b.ScalarFn(testUri, "add", nil, &types.Int32Type{}, &types.Int32Type{})
	fv := gLocalFunctionRegistry.GetScalarFunctionsBy("add", Substrait)
	assert.Greater(t, len(fv), 1)
	k := 0
	assert.Equal(t, expectedUri, fv[k].URI())
	assert.Equal(t, "+", fv[k].LocalName())
	assert.True(t, fv[0].IsOptionSupported("overflow", "ERROR"))
	assert.False(t, fv[0].IsOptionSupported("overflow", "SILENT"))
	checkCompoundNames(t, getScalarCompoundNames(fv), []string{"add:i8_i8", "add:i16_i16", "add:fp32_fp32", "add:fp64_fp64"})

	av := gLocalFunctionRegistry.GetAggregateFunctionsBy("max", Substrait)
	assert.Equal(t, expectedUri, av[0].URI())
	assert.Equal(t, "max", av[0].LocalName())
	checkCompoundNames(t, getAggregateCompoundNames(av), []string{"max:i8", "max:i16", "max:fp32", "max:fp64"})

	wv := gLocalFunctionRegistry.GetWindowFunctionsBy("ntile", Substrait)
	assert.Equal(t, expectedUri, wv[0].URI())
	assert.Equal(t, "ntile", wv[0].LocalName())
	checkCompoundNames(t, getWindowCompoundNames(wv), []string{"ntile:i32", "ntile:i64"})
}

func TestFunctionRegistry(t *testing.T) {
	registry := GetDefaultFunctionRegistry()
	allFunctions := registry.GetAllFunctions()
	assert.Greater(t, len(allFunctions), 377)

	addFunctions := registry.GetScalarFunctions("add")
	assert.Greater(t, len(addFunctions), 6)
	for _, f := range addFunctions {
		assert.Contains(t, allFunctions, f)
	}

	minFunctions := registry.GetAggregateFunctions("min")
	assert.Equal(t, len(minFunctions), 6)
	for _, f := range minFunctions {
		assert.Contains(t, allFunctions, f)
	}

	isNullFunctions := registry.GetScalarFunctions("is_null")
	assert.Greater(t, len(isNullFunctions), 0)
	assert.Contains(t, allFunctions, isNullFunctions[0])

	rankFunctions := registry.GetWindowFunctions("rank")
	assert.Greater(t, len(rankFunctions), 0)
	assert.Contains(t, allFunctions, rankFunctions[0])

	rankFunctions = registry.GetWindowFunctions("ntile")
	assert.Greater(t, len(rankFunctions), 0)
	assert.Contains(t, allFunctions, rankFunctions[0])
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
