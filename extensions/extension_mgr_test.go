// SPDX-License-Identifier: Apache-2.0

package extensions_test

import (
	"strings"
	"testing"

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
            value: any1
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

func TestLoadExtensionCollection(t *testing.T) {
	const uri = "http://localhost/sample.yaml"

	var c extensions.Collection
	require.NoError(t, c.Load(uri, strings.NewReader(sampleYAML)))

	t.Run("check types", func(t *testing.T) {
		id := extensions.ID{URI: uri}
		id.Name = "point"
		ty, ok := c.GetType(id)
		assert.True(t, ok)
		assert.Equal(t, "point", ty.Name)
		assert.Equal(t, map[string]interface{}{"latitude": "i32", "longitude": "i32"}, ty.Structure)
	})

	t.Run("simple and compound func signature", func(t *testing.T) {
		add, ok := c.GetScalarFunc(extensions.ID{URI: uri, Name: "add"})
		assert.True(t, ok)
		addCompound, ok := c.GetScalarFunc(extensions.ID{URI: uri, Name: "add:i8_i8"})
		assert.True(t, ok)
		assert.Same(t, add, addCompound)

		assert.Equal(t, "add", add.Name())
		assert.Equal(t, "add:i8_i8", add.CompoundName())
		assert.Equal(t, "Add two values.", add.Description())
		assert.Equal(t, uri, add.URI())
		assert.Equal(t, map[string]extensions.Option{"overflow": {
			Values: []string{"SILENT", "SATURATE", "ERROR"},
		}}, add.Options())

		ty, err := add.ResolveType(nil)
		assert.NoError(t, err)
		assert.Equal(t, &types.Int8Type{Nullability: types.NullabilityRequired}, ty)
	})

	t.Run("multiple impls need compound", func(t *testing.T) {
		sub, ok := c.GetScalarFunc(extensions.ID{URI: uri, Name: "subtract"})
		assert.Nil(t, sub)
		assert.False(t, ok)

		sub, ok = c.GetScalarFunc(extensions.ID{URI: uri, Name: "subtract:i16_i16"})
		assert.True(t, ok)
		assert.NotNil(t, sub)

		assert.Equal(t, "subtract", sub.Name())
		assert.Equal(t, "subtract:i16_i16", sub.CompoundName())

		ty, err := sub.ResolveType(nil)
		assert.NoError(t, err)
		assert.Equal(t, &types.Int16Type{Nullability: types.NullabilityRequired}, ty)
	})

	t.Run("same fn name different args", func(t *testing.T) {
		ct, ok := c.GetAggregateFunc(extensions.ID{URI: uri, Name: "count:"})
		assert.True(t, ok)
		assert.NotNil(t, ct)

		ctArgs, ok := c.GetAggregateFunc(extensions.ID{URI: uri, Name: "count:any"})
		assert.True(t, ok)
		assert.NotNil(t, ctArgs)

		assert.Equal(t, "Count a set of records (not field referenced)", ct.Description())
		assert.Equal(t, "Count a set of values", ctArgs.Description())
		assert.Equal(t, extensions.DecomposeMany, ct.Decomposability())
		assert.Equal(t, "i64", ct.Intermediate().String())
	})
}

func TestExtensionSet(t *testing.T) {
	const uri = "http://localhost/sample.yaml"

	s := extensions.NewSet()
	_, ok := s.DecodeFunc(0)
	assert.False(t, ok)
	_, ok = s.DecodeType(0)
	assert.False(t, ok)
	_, ok = s.DecodeTypeVariation(0)
	assert.False(t, ok)

	_, ok = s.FindURI(uri)
	assert.False(t, ok)

	t.Run("add anchors", func(t *testing.T) {
		id := extensions.ID{URI: uri}
		id.Name = "add"

		anchor := s.GetFuncAnchor(id)
		assert.Zero(t, anchor)
		nid, ok := s.DecodeFunc(0)
		assert.True(t, ok)
		assert.Equal(t, id, nid)

		id.Name = "subtract:i8_i8"
		anchor = s.GetFuncAnchor(id)
		assert.EqualValues(t, 1, anchor)

		id.Name = "point"
		anchor = s.GetTypeAnchor(id)
		assert.Zero(t, anchor)
	})

	var c extensions.Collection
	require.NoError(t, c.Load(uri, strings.NewReader(sampleYAML)))

	t.Run("lookup from collection", func(t *testing.T) {
		fn, ok := s.LookupScalarFunction(0, &c)
		assert.True(t, ok)
		assert.NotNil(t, fn)
		assert.Equal(t, "add", fn.Name())

		fn, ok = s.LookupScalarFunction(2, &c)
		assert.False(t, ok)
		assert.Nil(t, fn)
	})
}
