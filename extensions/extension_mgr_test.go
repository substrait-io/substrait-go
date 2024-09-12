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

		i8Req := &types.Int8Type{Nullability: types.NullabilityRequired}
		ty, err := add.ResolveType([]types.Type{i8Req, i8Req})
		assert.NoError(t, err)
		assert.Equal(t, i8Req, ty)
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

		i16Req := &types.Int16Type{Nullability: types.NullabilityRequired}
		ty, err := sub.ResolveType([]types.Type{i16Req, i16Req})
		assert.NoError(t, err)
		assert.Equal(t, i16Req, ty)
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
		ty, err := ct.Intermediate()
		require.NoError(t, err)
		assert.Equal(t, &types.Int64Type{Nullability: types.NullabilityRequired}, ty)
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
		assert.EqualValues(t, 1, anchor)
		nid, ok := s.DecodeFunc(1)
		assert.True(t, ok)
		assert.Equal(t, id, nid)

		id.Name = "subtract:i8_i8"
		anchor = s.GetFuncAnchor(id)
		assert.EqualValues(t, 2, anchor)

		id.Name = "point"
		anchor = s.GetTypeAnchor(id)
		assert.EqualValues(t, 1, anchor)
	})

	var c extensions.Collection
	require.NoError(t, c.Load(uri, strings.NewReader(sampleYAML)))

	t.Run("lookup from collection", func(t *testing.T) {
		fn, ok := s.LookupScalarFunction(1, &c)
		assert.True(t, ok)
		assert.NotNil(t, fn)
		assert.Equal(t, "add", fn.Name())

		fn, ok = s.LookupScalarFunction(3, &c)
		assert.False(t, ok)
		assert.Nil(t, fn)
	})
}

func TestDefaultCollection(t *testing.T) {
	type funcType int8
	const (
		scalarFunc funcType = iota
		aggFunc
		windowFunc
	)

	tests := []struct {
		typ          funcType
		uri          string
		name         string
		compoundName string
		nargs        int
		options      map[string]extensions.Option
		variadic     *extensions.VariadicBehavior
	}{
		{scalarFunc, extensions.SubstraitDefaultURIPrefix + "functions_arithmetic.yaml",
			"add", "add:i32_i32", 2, map[string]extensions.Option{"overflow": {Values: []string{"SILENT", "SATURATE", "ERROR"}}},
			nil},
		{aggFunc, extensions.SubstraitDefaultURIPrefix + "functions_arithmetic.yaml",
			"variance", "variance:fp64", 1, map[string]extensions.Option{
				"distribution": {Values: []string{"SAMPLE", "POPULATION"}},
				"rounding":     {Values: []string{"TIE_TO_EVEN", "TIE_AWAY_FROM_ZERO", "TRUNCATE", "CEILING", "FLOOR"}}},
			nil},
		{windowFunc, extensions.SubstraitDefaultURIPrefix + "functions_arithmetic.yaml",
			"dense_rank", "dense_rank:", 0, nil, nil},
		{scalarFunc, extensions.SubstraitDefaultURIPrefix + "functions_boolean.yaml",
			"or", "or:bool", 1, nil, &extensions.VariadicBehavior{Min: 0}},
		{aggFunc, extensions.SubstraitDefaultURIPrefix + "functions_boolean.yaml",
			"bool_and", "bool_and:bool", 1, nil, nil},
		{aggFunc, extensions.SubstraitDefaultURIPrefix + "functions_aggregate_approx.yaml",
			"approx_count_distinct", "approx_count_distinct:any", 1, nil, nil},
		{aggFunc, extensions.SubstraitDefaultURIPrefix + "functions_aggregate_generic.yaml",
			"count", "count:", 0, map[string]extensions.Option{"overflow": {Values: []string{"SILENT", "SATURATE", "ERROR"}}}, nil},
		{aggFunc, extensions.SubstraitDefaultURIPrefix + "functions_aggregate_generic.yaml",
			"count", "count:any", 1, map[string]extensions.Option{"overflow": {Values: []string{"SILENT", "SATURATE", "ERROR"}}}, nil},
		{scalarFunc, extensions.SubstraitDefaultURIPrefix + "functions_comparison.yaml",
			"not_equal", "not_equal:any_any", 2, nil, nil},
		{scalarFunc, extensions.SubstraitDefaultURIPrefix + "functions_comparison.yaml",
			"between", "between:any_any_any", 3, nil, nil},
		{scalarFunc, extensions.SubstraitDefaultURIPrefix + "functions_datetime.yaml",
			"add", "add:ts_iyear", 2, nil, nil},
		{scalarFunc, extensions.SubstraitDefaultURIPrefix + "functions_logarithmic.yaml",
			"ln", "ln:fp32", 1, map[string]extensions.Option{
				"rounding":        {Values: []string{"TIE_TO_EVEN", "TIE_AWAY_FROM_ZERO", "TRUNCATE", "CEILING", "FLOOR"}},
				"on_domain_error": {Values: []string{"NAN", "ERROR"}},
				"on_log_zero":     {Values: []string{"NAN", "ERROR", "MINUS_INFINITY"}},
			}, nil},
		{scalarFunc, extensions.SubstraitDefaultURIPrefix + "functions_rounding.yaml",
			"ceil", "ceil:fp64", 1, nil, nil},
		{scalarFunc, extensions.SubstraitDefaultURIPrefix + "functions_set.yaml",
			"index_in", "index_in:t_list", 2, map[string]extensions.Option{
				"nan_equality": {Values: []string{"NAN_IS_NAN", "NAN_IS_NOT_NAN"}},
			}, nil},
		{scalarFunc, extensions.SubstraitDefaultURIPrefix + "functions_string.yaml",
			"string_split", "string_split:vchar_vchar", 2, nil, nil},
		{scalarFunc, extensions.SubstraitDefaultURIPrefix + "functions_string.yaml",
			"string_split", "string_split:str_str", 2, nil, nil},
		{aggFunc, extensions.SubstraitDefaultURIPrefix + "functions_string.yaml",
			"string_agg", "string_agg:str_str", 2, nil, nil},
	}

	for _, tt := range tests {
		t.Run(tt.compoundName, func(t *testing.T) {
			var (
				variant extensions.FunctionVariant
				ok      bool

				id = extensions.ID{URI: tt.uri, Name: tt.compoundName}
			)
			switch tt.typ {
			case scalarFunc:
				variant, ok = extensions.DefaultCollection.GetScalarFunc(id)
			case aggFunc:
				variant, ok = extensions.DefaultCollection.GetAggregateFunc(id)
			case windowFunc:
				variant, ok = extensions.DefaultCollection.GetWindowFunc(id)
			}

			require.True(t, ok)
			require.NotNil(t, variant)

			assert.Equal(t, tt.name, variant.Name())
			assert.Equal(t, tt.compoundName, variant.CompoundName())
			assert.Equal(t, tt.options, variant.Options())
			assert.Equal(t, tt.uri, variant.URI())
			assert.Len(t, variant.Args(), tt.nargs)
		})
	}

	et, ok := extensions.DefaultCollection.GetType(extensions.ID{
		URI: extensions.SubstraitDefaultURIPrefix + "extension_types.yaml", Name: "point"})
	assert.True(t, ok)
	assert.Equal(t, "point", et.Name)
	assert.Equal(t, map[string]interface{}{"latitude": "i32", "longitude": "i32"}, et.Structure)
}

func TestCollection_GetAllScalarFunctions(t *testing.T) {
	scalarFunctions := extensions.DefaultCollection.GetAllScalarFunctions()
	aggregateFunctions := extensions.DefaultCollection.GetAllAggregateFunctions()
	windowFunctions := extensions.DefaultCollection.GetAllWindowFunctions()
	assert.GreaterOrEqual(t, len(scalarFunctions), 309)
	assert.GreaterOrEqual(t, len(aggregateFunctions), 62)
	assert.GreaterOrEqual(t, len(windowFunctions), 7)
	tests := []struct {
		uri         string
		signature   string
		isScalar    bool
		isAggregate bool
		isWindow    bool
	}{
		{extensions.SubstraitDefaultURIPrefix + "functions_arithmetic.yaml", "add:i32_i32", true, false, false},
		{extensions.SubstraitDefaultURIPrefix + "functions_arithmetic.yaml", "variance:fp64", false, true, false},
		{extensions.SubstraitDefaultURIPrefix + "functions_arithmetic.yaml", "dense_rank:", false, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.signature, func(t *testing.T) {
			assert.True(t, tt.isScalar || tt.isAggregate || tt.isWindow)
			c := extensions.DefaultCollection
			if tt.isScalar {
				sf, ok := c.GetScalarFunc(extensions.ID{URI: tt.uri, Name: tt.signature})
				assert.True(t, ok)
				assert.Contains(t, scalarFunctions, sf)
				// verify that default nullability is set to MIRROR
				assert.Equal(t, extensions.MirrorNullability, sf.Nullability())
			}
			if tt.isAggregate {
				af, ok := c.GetAggregateFunc(extensions.ID{URI: tt.uri, Name: tt.signature})
				assert.True(t, ok)
				assert.Contains(t, aggregateFunctions, af)
			}
			if tt.isWindow {
				wf, ok := c.GetWindowFunc(extensions.ID{URI: tt.uri, Name: tt.signature})
				assert.True(t, ok)
				assert.Contains(t, windowFunctions, wf)
			}
		})
	}
}
