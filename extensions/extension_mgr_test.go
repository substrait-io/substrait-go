// SPDX-License-Identifier: Apache-2.0

package extensions_test

import (
	"strings"
	"testing"

	"github.com/goccy/go-yaml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/extensions"
	"github.com/substrait-io/substrait-go/types"
)

const sampleYAML = `---
types:
  - name: point
    structure:
      latitude: i32
      longitude: i32
  - name: line
    structure:
      start: point
      end: point
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
  -
    name: "subtract"
    description: "Subtract one value from another."
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
      - args:
          - name: x
            value: i16
          - name: y
            value: i16
        options:
          overflow:
            values: [SILENT, SATURATE, ERROR ]
        return: i16
aggregate_functions:
  - name: "count"
    description: Count a set of values
    impls:
      - args:
          - name: x
            value: any
        options:
          overflow:
            values: [SILENT, SATURATE, ERROR]
        nullability: DECLARED_OUTPUT
        decomposable: MANY
        intermediate: i64
        return: i64
  - name: "count"
    description: "Count a set of records (not field referenced)"
    impls:
      - options:
          overflow:
            values: [SILENT, SATURATE, ERROR]
        nullability: DECLARED_OUTPUT
        decomposable: MANY
        intermediate: i64
        return: i64
`

func TestLoadExtensionSet(t *testing.T) {
	var f extensions.SimpleExtensionFile
	err := yaml.NewDecoder(strings.NewReader(sampleYAML)).Decode(&f)
	require.NoError(t, err)

	const uri = "http://localhost/sample.yaml"

	s := extensions.NewSet()
	assert.NoError(t, s.Load(uri, &f))

	ty, ok := s.LookupType(0)
	assert.True(t, ok)
	assert.NotNil(t, ty)

	assert.Equal(t, "point", ty.Name)
	assert.Equal(t, map[string]interface{}{"latitude": "i32", "longitude": "i32"}, ty.Structure)
	a, ok := s.GetTypeAnchor(extensions.ID{URI: uri, Name: "point"})
	assert.True(t, ok)
	assert.Zero(t, a)

	a, ok = s.GetFuncAnchor(extensions.ID{URI: uri, Name: "subtract"})
	assert.False(t, ok)
	assert.Zero(t, a)

	a, ok = s.GetFuncAnchor(extensions.ID{URI: uri, Name: "subtract:i16_i16"})
	assert.True(t, ok)
	assert.EqualValues(t, 2, a)

	id, ok := s.DecodeFunc(2)
	assert.True(t, ok)
	assert.Equal(t, uri, id.URI)
	assert.Equal(t, "subtract:i16_i16", id.Name)

	va, ok := s.LookupFunction(2)
	assert.True(t, ok)
	assert.IsType(t, (*extensions.ScalarFunctionVariant)(nil), va)
	assert.Equal(t, "subtract", va.Name())
	outType, err := va.ResolveType([]types.Type{})
	assert.NoError(t, err)
	assert.Equal(t, &types.Int16Type{Nullability: types.NullabilityRequired}, outType)
}
