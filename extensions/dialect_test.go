// SPDX-License-Identifier: Apache-2.0

package extensions

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
  aggregate_approx: 
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_aggregate_approx.yaml
  aggregate_generic: 
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_aggregate_generic.yaml
  arithmetic: 
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml
  arithmetic_decimal: 
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic_decimal.yaml
  boolean: 
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_boolean.yaml
  comparison: 
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_comparison.yaml
  datetime: 
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_datetime.yaml
  geometry: 
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_geometry.yaml
  logarithmic: 
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_logarithmic.yaml
  rounding: 
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_rounding.yaml
  set: 
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_set.yaml
  string: 
    https://github.com/substrait-io/substrait/blob/main/extensions/functions_string.yaml
supported_types:
  i8: TINYINT,
  i16: SMALLINT
  i32: INTEGER
  i64: BIGINT
  fp32: REAL
  fp64: DOUBLE
  bool: BOOLEAN
  str: VARCHAR
  date: DATE
  time: TIME
  ts: TIMESTAMP
  tstz: TIMESTAMPTZ
  interval: INTERVAL
  decimal: DECIMAL
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
`

var gDialectName string
var gDialect Dialect

func TestMain(m *testing.M) {
	gDialectName = "dialect_test"
	err := GetFunctionRegistry().LoadDialect("dialect_test", strings.NewReader(sampleDialectYAML))
	if err != nil {
		log.Fatalf("dialect load failed %v", err)
	}
	gDialect, err = GetFunctionRegistry().GetDialect(gDialectName)
	if err != nil {
		log.Fatalf("dialect get failed %v", err)
	}
	m.Run()
}

func TestDialectApis(t *testing.T) {
	dialect, err := GetFunctionRegistry().GetDialect(gDialectName)
	assert.Nil(t, err)
	assert.NotNil(t, dialect)

	dialect, err = GetFunctionRegistry().GetDialect("non-existent")
	assert.Error(t, err)
	assert.Nil(t, dialect)
	dialects := DefaultCollection.GetSupportedDialects()
	assert.Len(t, dialects, 1)
	assert.Contains(t, dialects, gDialectName)

	err = GetFunctionRegistry().LoadDialect("second_dialect", strings.NewReader(sampleDialectYAML))
	assert.NoError(t, err)
	dialects = DefaultCollection.GetSupportedDialects()
	assert.Len(t, dialects, 2)
	assert.Contains(t, dialects, "second_dialect")
	assert.Contains(t, dialects, gDialectName)

	dialect, err = GetFunctionRegistry().GetDialect(gDialectName)
	assert.Nil(t, err)
	assert.NotNil(t, dialect)

	dialect, err = GetFunctionRegistry().GetDialect("second_dialect")
	assert.Nil(t, err)
	assert.NotNil(t, dialect)

	err = GetFunctionRegistry().LoadDialect(gDialectName, strings.NewReader(sampleDialectYAML))
	assert.Error(t, err)
}

func TestRegistryLocalLookup(t *testing.T) {
	expectedUri := "https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml"
	fv, ok := gDialect.LocalLookup("+", []string{"INTEGER", "INTEGER"})
	assert.True(t, ok)
	assert.Equal(t, "add:i32_i32", fv.CompoundName())
	assert.Equal(t, expectedUri, fv.URI())
	funcBinding, ok := fv.GetDialectBinding(gDialectName)
	assert.True(t, ok)
	assert.Equal(t, ScalarFunc, funcBinding.GetFuncKind())
	assert.Equal(t, "+", funcBinding.GetLocalName())
	assert.Equal(t, "add", funcBinding.GetName())
	assert.Equal(t, "INTEGER", funcBinding.GetLocalArgumentTypes()[0])
	assert.Equal(t, "INTEGER", funcBinding.GetLocalArgumentTypes()[1])
	assert.Contains(t, funcBinding.GetOptions(), "overflow")

	fv, ok = gDialect.LocalLookup("min", []string{"INTEGER"})
	assert.True(t, ok)
	assert.Equal(t, expectedUri, fv.URI())
	assert.Equal(t, "min:i32", fv.CompoundName())
	funcBinding, ok = fv.GetDialectBinding(gDialectName)
	assert.True(t, ok)
	assert.Equal(t, AggregateFunc, funcBinding.GetFuncKind())
}

func TestRegistryCanonicalLookup(t *testing.T) {
	testUri := "https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml"
	//b := plan.NewBuilderDefault()
	//sf, err := b.ScalarFn(testUri, "add", nil, &types.Int32Type{}, &types.Int32Type{})
	sf, ok := DefaultCollection.GetScalarFunc(ID{Name: "add:i32_i32", URI: testUri})
	assert.True(t, ok)
	assert.Equal(t, "add", sf.Name())
	fb, ok := sf.GetDialectBinding(gDialectName)
	assert.True(t, ok)
	assert.Equal(t, "+", fb.GetLocalName())
	assert.Equal(t, "INTEGER", fb.GetLocalArgumentTypes()[0])
	assert.Equal(t, "INTEGER", fb.GetLocalArgumentTypes()[1])

	af, ok := DefaultCollection.GetAggregateFunc(ID{Name: "max:i32", URI: testUri})
	assert.True(t, ok)
	assert.Equal(t, "max", af.Name())
	fb, ok = af.GetDialectBinding(gDialectName)
	assert.True(t, ok)
	assert.Equal(t, "max", fb.GetLocalName())
	assert.Equal(t, "INTEGER", fb.GetLocalArgumentTypes()[0])
}

func TestRegistryBinaryFunctions(t *testing.T) {
	testUri := "https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml"
	binaryFunctions := gDialect.GetBinaryFunctions()
	assert.Len(t, binaryFunctions, 12)
	var compoundNames []string
	for _, f := range binaryFunctions {
		compoundNames = append(compoundNames, f.CompoundName())
		_, ok := f.GetDialectBinding(gDialectName)
		if ok {
			assert.Equal(t, testUri, f.URI())
		}
	}
	assert.Contains(t, compoundNames, "add:i8_i8")
	assert.Contains(t, compoundNames, "add:i16_i16")
	assert.Contains(t, compoundNames, "add:fp32_fp32")
	assert.Contains(t, compoundNames, "add:fp64_fp64")
	assert.Contains(t, compoundNames, "subtract:i32_i32")
	assert.Contains(t, compoundNames, "subtract:i64_i64")
	assert.Equal(t, "add:i8_i8", binaryFunctions[0].CompoundName())
	assert.Equal(t, "subtract:fp64_fp64", binaryFunctions[11].CompoundName())

	binaryFunctions = GetFunctionRegistry().GetBinaryFunctions()
	assert.Greater(t, len(binaryFunctions), 736)
	assert.Equal(t, 2, len(binaryFunctions[0].Args()))
}
