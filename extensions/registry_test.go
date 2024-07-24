// SPDX-License-Identifier: Apache-2.0

package extensions

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const sampleDialectYAML = `---
name: duckdb
type: sql
dependencies:
  aggregate_approx: 
    https://github.com/substrait-io/substrait/blob/main/extensions/substrait/extensions/functions_aggregate_approx.yaml
  aggregate_generic: 
    https://github.com/substrait-io/substrait/blob/main/extensions/substrait/extensions/functions_aggregate_generic.yaml
  arithmetic: 
    https://github.com/substrait-io/substrait/blob/main/extensions/substrait/extensions/functions_arithmetic.yaml
  arithmetic_decimal: 
    https://github.com/substrait-io/substrait/blob/main/extensions/substrait/extensions/functions_arithmetic_decimal.yaml
  boolean: 
    https://github.com/substrait-io/substrait/blob/main/extensions/substrait/extensions/functions_boolean.yaml
  comparison: 
    https://github.com/substrait-io/substrait/blob/main/extensions/substrait/extensions/functions_comparison.yaml
  datetime: 
    https://github.com/substrait-io/substrait/blob/main/extensions/substrait/extensions/functions_datetime.yaml
  geometry: 
    https://github.com/substrait-io/substrait/blob/main/extensions/substrait/extensions/functions_geometry.yaml
  logarithmic: 
    https://github.com/substrait-io/substrait/blob/main/extensions/substrait/extensions/functions_logarithmic.yaml
  rounding: 
    https://github.com/substrait-io/substrait/blob/main/extensions/substrait/extensions/functions_rounding.yaml
  set: 
    https://github.com/substrait-io/substrait/blob/main/extensions/substrait/extensions/functions_set.yaml
  string: 
    https://github.com/substrait-io/substrait/blob/main/extensions/substrait/extensions/functions_string.yaml
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

func TestRegistryLoad(t *testing.T) {
	registry := NewRegistry()
	err := registry.Load(strings.NewReader(sampleDialectYAML))
	assert.NoError(t, err)
}

func TestRegistryLocalLookup(t *testing.T) {
	registry := NewRegistry()
	err := registry.Load(strings.NewReader(sampleDialectYAML))
	assert.NoError(t, err)

	expectedUri := "https://github.com/substrait-io/substrait/blob/main/extensions/substrait/extensions/functions_arithmetic.yaml"
	id, funcKind, ok := registry.LocalLookup("+", []string{"INTEGER", "INTEGER"})
	assert.True(t, ok)
	assert.Equal(t, "add:i32_i32", id.Name)
	assert.Equal(t, expectedUri, id.URI)
	assert.Equal(t, ScalarFunc, funcKind)

	id, funcKind, ok = registry.LocalLookup("min", []string{"INTEGER"})
	assert.True(t, ok)
	assert.Equal(t, expectedUri, id.URI)
	assert.Equal(t, "min:i32", id.Name)
	assert.Equal(t, AggregateFunc, funcKind)
}

func TestRegistryCanonicalLookup(t *testing.T) {
	registry := NewRegistry()
	err := registry.Load(strings.NewReader(sampleDialectYAML))
	assert.NoError(t, err)

	testUri := "https://github.com/substrait-io/substrait/blob/main/extensions/substrait/extensions/functions_arithmetic.yaml"
	lf, ok := registry.CanonicalLookup(ID{Name: "add:i32_i32", URI: testUri}, ScalarFunc)
	assert.True(t, ok)
	assert.Equal(t, "add", lf.Name)
	assert.Equal(t, "+", lf.LocalName)
	assert.Equal(t, "i32_i32", lf.Signature)

	lf, ok = registry.CanonicalLookup(ID{Name: "max:i32", URI: testUri}, AggregateFunc)
	assert.True(t, ok)
	assert.Equal(t, "max", lf.Name)
	assert.Equal(t, "max", lf.LocalName)
	assert.Equal(t, "i32", lf.Signature)
}

func TestRegistryBinaryFunctions(t *testing.T) {
	registry := NewRegistry()
	err := registry.Load(strings.NewReader(sampleDialectYAML))
	assert.NoError(t, err)

	testUri := "https://github.com/substrait-io/substrait/blob/main/extensions/substrait/extensions/functions_arithmetic.yaml"
	binaryFunctions := registry.BinaryFunctions()
	assert.Len(t, binaryFunctions, 12)
	assert.Contains(t, binaryFunctions, ID{Name: "add:i8_i8", URI: testUri})
	assert.Contains(t, binaryFunctions, ID{Name: "add:i16_i16", URI: testUri})
	assert.Contains(t, binaryFunctions, ID{Name: "add:fp32_fp32", URI: testUri})
	assert.Contains(t, binaryFunctions, ID{Name: "add:fp64_fp64", URI: testUri})
	assert.Contains(t, binaryFunctions, ID{Name: "subtract:i32_i32", URI: testUri})
	assert.Contains(t, binaryFunctions, ID{Name: "subtract:i64_i64", URI: testUri})
	assert.Equal(t, "add:i8_i8", binaryFunctions[0].Name)
	assert.Equal(t, "subtract:fp64_fp64", binaryFunctions[11].Name)
}
