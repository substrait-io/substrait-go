// SPDX-License-Identifier: Apache-2.0

package extensions_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/extensions"

	//"gopkg.in/yaml.v3"
	"github.com/goccy/go-yaml"
)

func TestUnmarshalSimpleExtension(t *testing.T) {
	var testExtType = `
types:
  - name: "null"
    structure: {}
  - name: "foo"
    structure: "bar"
  - name: interval_month
    structure:
      months: i32
  - name: baz
`

	var f extensions.SimpleExtensionFile
	require.NoError(t, yaml.Unmarshal([]byte(testExtType), &f))

	assert.Len(t, f.Types, 4)
	assert.IsType(t, (map[string]any)(nil), f.Types[0].Structure)

	assert.Equal(t, "null", f.Types[0].Name)
	assert.Len(t, f.Types[0].Structure, 0)
	assert.Equal(t, "foo", f.Types[1].Name)
	assert.Equal(t, "bar", f.Types[1].Structure)
	assert.Equal(t, "interval_month", f.Types[2].Name)
	assert.Equal(t, map[string]any{"months": "i32"}, f.Types[2].Structure)
	assert.Equal(t, "baz", f.Types[3].Name)
	assert.Nil(t, f.Types[3].Structure)
}

func TestUnmarshalSimpleExtensionScalarFunction(t *testing.T) {
	const addDef = `
scalar_functions:
  -
    name: "add"
    description: "Add two values."
    impls:
      - args:
          - name: x
            value: i8
          - name: y
            value: i8
        options:
          overflow:
            values: [ SILENT, SATURATE, ERROR ]
        return: i8
`

	var f extensions.SimpleExtensionFile
	require.NoError(t, yaml.Unmarshal([]byte(addDef), &f))

	assert.Len(t, f.ScalarFunctions, 1)
	assert.Len(t, f.ScalarFunctions[0].Impls, 1)
	assert.Len(t, f.ScalarFunctions[0].Impls[0].Args, 2)
	assert.IsType(t, extensions.ValueArg{}, f.ScalarFunctions[0].Impls[0].Args[0])
	assert.IsType(t, extensions.ValueArg{}, f.ScalarFunctions[0].Impls[0].Args[1])
	assert.Equal(t, extensions.ValueArg{Name: "x", Value: "i8"}, f.ScalarFunctions[0].Impls[0].Args[0])
	assert.Equal(t, extensions.ValueArg{Name: "y", Value: "i8"}, f.ScalarFunctions[0].Impls[0].Args[1])

	assert.Equal(t, map[string]extensions.Option{
		"overflow": {Values: []string{"SILENT", "SATURATE", "ERROR"}},
	}, f.ScalarFunctions[0].Impls[0].Options)
}

const snippetScalarArithmeticFile = `%YAML 1.2
---
scalar_functions:
- name: "multiply"
  description: "Multiply two values."
  impls:
    - args:
        - name: x
          value: i8
        - name: y
          value: i8
      options:
        overflow:
          values: [SILENT, SATURATE, ERROR]
      return: i8
    - args:
        - name: x
          value: i16
        - name: y
          value: i16
      options:
        overflow:
          values: [SILENT, SATURATE, ERROR]
      return: i16
    - args:
        - name: x
          value: i32
        - name: y
          value: i32
      options:
        overflow:
          values: [SILENT, SATURATE, ERROR]
      return: i32
    - args:
        - name: x
          value: i64
        - name: y
          value: i64
      options:
        overflow:
          values: [SILENT, SATURATE, ERROR]
      return: i64
    - args:
        - name: x
          value: fp32
        - name: y
          value: fp32
      options:
        rounding:
          values: [TIE_TO_EVEN, TIE_AWAY_FROM_ZERO, TRUNCATE, CEILING, FLOOR]
      return: fp32
    - args:
        - name: x
          value: fp64
        - name: y
          value: fp64
      options:
        rounding:
          values: [TIE_TO_EVEN, TIE_AWAY_FROM_ZERO, TRUNCATE, CEILING, FLOOR]
      return: fp64
- name: "divide"
  description: >
    Divide x by y. In the case of integer division, partial values are truncated (i.e. rounded towards 0).
    The 'on_division_by_zero' option governs behavior in cases where y is 0 and x is not 0.
    'LIMIT' means positive or negative infinity (depending on the sign of x and y).
    If x and y are both 0 or both +/-infinity, behavior will be governed by 'on_domain_error'.
  impls:
    - args:
        - name: x
          value: i8
        - name: y
          value: i8
      options:
        overflow:
          values: [SILENT, SATURATE, ERROR]
      return: i8
    - args:
        - name: x
          value: i16
        - name: y
          value: i16
      options:
        overflow:
          values: [SILENT, SATURATE, ERROR]
      return: i16
    - args:
        - name: x
          value: i32
        - name: y
          value: i32
      options:
        overflow:
          values: [SILENT, SATURATE, ERROR]
      return: i32
    - args:
        - name: x
          value: i64
        - name: y
          value: i64
      options:
        overflow:
          values: [SILENT, SATURATE, ERROR]
      return: i64
    - args:
        - name: x
          value: fp32
        - name: y
          value: fp32
      options:
        rounding:
          values: [TIE_TO_EVEN, TIE_AWAY_FROM_ZERO, TRUNCATE, CEILING, FLOOR]
        on_domain_error:
          values: [NAN, ERROR]
        on_division_by_zero:
          values: [LIMIT, NAN, ERROR]
      return: fp32
    - args:
        - name: x
          value: fp64
        - name: y
          value: fp64
      options:
        rounding:
          values: [TIE_TO_EVEN, TIE_AWAY_FROM_ZERO, TRUNCATE, CEILING, FLOOR]
        on_domain_error:
          values: [NAN, ERROR]
        on_division_by_zero:
          values: [LIMIT, NAN, ERROR]
      return: fp64
`

func TestScalarFunctionsRoundtrip(t *testing.T) {
	d := yaml.NewDecoder(strings.NewReader(snippetScalarArithmeticFile))

	var file extensions.SimpleExtensionFile
	require.NoError(t, d.Decode(&file))

	data, err := yaml.Marshal(&file)
	require.NoError(t, err)

	var exp yaml.MapItem
	var actual yaml.MapItem

	require.NoError(t, yaml.Unmarshal([]byte(snippetScalarArithmeticFile), &exp))
	require.NoError(t, yaml.Unmarshal(data, &actual))

	assert.Equal(t, exp, actual)
}
