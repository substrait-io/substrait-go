// SPDX-License-Identifier: Apache-2.0

package extensions_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v6/extensions"
	"github.com/substrait-io/substrait-go/v6/types"
)

const sampleYAML = `---
urn: extension:test:sample
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
	const urn = "extension:test:sample"

	var c extensions.Collection
	require.NoError(t, c.Load(uri, strings.NewReader(sampleYAML)))

	t.Run("check types", func(t *testing.T) {
		id := extensions.ID{URN: urn}
		id.Name = "point"
		ty, ok := c.GetType(id)
		assert.True(t, ok)
		assert.Equal(t, "point", ty.Name)
		assert.Equal(t, map[string]interface{}{"latitude": "i32", "longitude": "i32"}, ty.Structure)
	})

	t.Run("simple and compound func signature", func(t *testing.T) {
		add, ok := c.GetScalarFunc(extensions.ID{URN: urn, Name: "add"})
		assert.True(t, ok)
		addCompound, ok := c.GetScalarFunc(extensions.ID{URN: urn, Name: "add:i8_i8"})
		assert.True(t, ok)
		assert.Same(t, add, addCompound)

		assert.Equal(t, "add", add.Name())
		assert.Equal(t, "add:i8_i8", add.CompoundName())
		assert.Equal(t, "Add two values.", add.Description())
		assert.Equal(t, urn, add.URN())
		assert.Equal(t, map[string]extensions.Option{"overflow": {
			Values: []string{"SILENT", "SATURATE", "ERROR"},
		}}, add.Options())

		i8Req := &types.Int8Type{Nullability: types.NullabilityRequired}
		ty, err := add.ResolveType([]types.Type{i8Req, i8Req}, extensions.NewSet())
		assert.NoError(t, err)
		assert.Equal(t, i8Req, ty)
	})

	t.Run("multiple impls need compound", func(t *testing.T) {
		sub, ok := c.GetScalarFunc(extensions.ID{URN: urn, Name: "subtract"})
		assert.Nil(t, sub)
		assert.False(t, ok)

		sub, ok = c.GetScalarFunc(extensions.ID{URN: urn, Name: "subtract:i16_i16"})
		assert.True(t, ok)
		assert.NotNil(t, sub)

		assert.Equal(t, "subtract", sub.Name())
		assert.Equal(t, "subtract:i16_i16", sub.CompoundName())

		i16Req := &types.Int16Type{Nullability: types.NullabilityRequired}
		ty, err := sub.ResolveType([]types.Type{i16Req, i16Req}, extensions.NewSet())
		assert.NoError(t, err)
		assert.Equal(t, i16Req, ty)
	})

	t.Run("same fn name different args", func(t *testing.T) {
		ct, ok := c.GetAggregateFunc(extensions.ID{URN: urn, Name: "count:"})
		assert.True(t, ok)
		assert.NotNil(t, ct)

		ctArgs, ok := c.GetAggregateFunc(extensions.ID{URN: urn, Name: "count:any"})
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
	const urn = "extension:test:sample"

	s := extensions.NewSet()
	_, ok := s.DecodeFunc(0)
	assert.False(t, ok)
	_, ok = s.DecodeType(0)
	assert.False(t, ok)
	_, ok = s.DecodeTypeVariation(0)
	assert.False(t, ok)

	_, ok = s.FindURN(urn)
	assert.False(t, ok)

	t.Run("add anchors", func(t *testing.T) {
		id := extensions.ID{URN: urn}
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
		urn          string
		name         string
		compoundName string
		nargs        int
		options      map[string]extensions.Option
		variadic     *extensions.VariadicBehavior
	}{
		{scalarFunc, extensions.SubstraitDefaultURNPrefix + "functions_arithmetic",
			"add", "add:i32_i32", 2, map[string]extensions.Option{"overflow": {Values: []string{"SILENT", "SATURATE", "ERROR"}}},
			nil},
		{aggFunc, extensions.SubstraitDefaultURNPrefix + "functions_arithmetic",
			"variance", "variance:fp64", 1, map[string]extensions.Option{
				"distribution": {Values: []string{"SAMPLE", "POPULATION"}},
				"rounding":     {Values: []string{"TIE_TO_EVEN", "TIE_AWAY_FROM_ZERO", "TRUNCATE", "CEILING", "FLOOR"}}},
			nil},
		{windowFunc, extensions.SubstraitDefaultURNPrefix + "functions_arithmetic",
			"dense_rank", "dense_rank:", 0, nil, nil},
		{scalarFunc, extensions.SubstraitDefaultURNPrefix + "functions_boolean",
			"or", "or:bool", 1, nil, &extensions.VariadicBehavior{Min: 0}},
		{aggFunc, extensions.SubstraitDefaultURNPrefix + "functions_boolean",
			"bool_and", "bool_and:bool", 1, nil, nil},
		{aggFunc, extensions.SubstraitDefaultURNPrefix + "functions_aggregate_approx",
			"approx_count_distinct", "approx_count_distinct:any", 1, nil, nil},
		{aggFunc, extensions.SubstraitDefaultURNPrefix + "functions_aggregate_generic",
			"count", "count:", 0, map[string]extensions.Option{"overflow": {Values: []string{"SILENT", "SATURATE", "ERROR"}}}, nil},
		{aggFunc, extensions.SubstraitDefaultURNPrefix + "functions_aggregate_generic",
			"count", "count:any", 1, map[string]extensions.Option{"overflow": {Values: []string{"SILENT", "SATURATE", "ERROR"}}}, nil},
		{scalarFunc, extensions.SubstraitDefaultURNPrefix + "functions_comparison",
			"not_equal", "not_equal:any_any", 2, nil, nil},
		{scalarFunc, extensions.SubstraitDefaultURNPrefix + "functions_comparison",
			"between", "between:any_any_any", 3, nil, nil},
		{scalarFunc, extensions.SubstraitDefaultURNPrefix + "functions_datetime",
			"add", "add:ts_iyear", 2, nil, nil},
		{scalarFunc, extensions.SubstraitDefaultURNPrefix + "functions_logarithmic",
			"ln", "ln:fp32", 1, map[string]extensions.Option{
				"rounding":        {Values: []string{"TIE_TO_EVEN", "TIE_AWAY_FROM_ZERO", "TRUNCATE", "CEILING", "FLOOR"}},
				"on_domain_error": {Values: []string{"NAN", "NULL", "ERROR"}},
				"on_log_zero":     {Values: []string{"NAN", "ERROR", "MINUS_INFINITY"}},
			}, nil},
		{scalarFunc, extensions.SubstraitDefaultURNPrefix + "functions_rounding",
			"ceil", "ceil:fp64", 1, nil, nil},
		{scalarFunc, extensions.SubstraitDefaultURNPrefix + "functions_set",
			"index_in", "index_in:any_list", 2, map[string]extensions.Option{
				"nan_equality": {Values: []string{"NAN_IS_NAN", "NAN_IS_NOT_NAN"}},
			}, nil},
		{scalarFunc, extensions.SubstraitDefaultURNPrefix + "functions_string",
			"string_split", "string_split:vchar_vchar", 2, nil, nil},
		{scalarFunc, extensions.SubstraitDefaultURNPrefix + "functions_string",
			"string_split", "string_split:str_str", 2, nil, nil},
		{aggFunc, extensions.SubstraitDefaultURNPrefix + "functions_string",
			"string_agg", "string_agg:str_str", 2, nil, nil},
	}

	for _, tt := range tests {
		t.Run(tt.compoundName, func(t *testing.T) {
			var (
				variant extensions.FunctionVariant
				ok      bool

				id = extensions.ID{URN: tt.urn, Name: tt.compoundName}
			)
			switch tt.typ {
			case scalarFunc:
				variant, ok = extensions.GetDefaultCollectionWithNoError().GetScalarFunc(id)
			case aggFunc:
				variant, ok = extensions.GetDefaultCollectionWithNoError().GetAggregateFunc(id)
			case windowFunc:
				variant, ok = extensions.GetDefaultCollectionWithNoError().GetWindowFunc(id)
			}

			require.True(t, ok)
			require.NotNil(t, variant)

			assert.Equal(t, tt.name, variant.Name())
			assert.Equal(t, tt.compoundName, variant.CompoundName())
			assert.Equal(t, tt.options, variant.Options())
			assert.Equal(t, tt.urn, variant.URN())
			assert.Len(t, variant.Args(), tt.nargs)
		})
	}

	et, ok := extensions.GetDefaultCollectionWithNoError().GetType(extensions.ID{
		URN: extensions.SubstraitDefaultURNPrefix + "extension_types", Name: "point"})
	assert.True(t, ok)
	assert.Equal(t, "point", et.Name)
	assert.Equal(t, map[string]interface{}{"latitude": "i32", "longitude": "i32"}, et.Structure)
}

func TestCollection_GetAllScalarFunctions(t *testing.T) {
	defaultExtensions := extensions.GetDefaultCollectionWithNoError()
	scalarFunctions := defaultExtensions.GetAllScalarFunctions()
	aggregateFunctions := defaultExtensions.GetAllAggregateFunctions()
	windowFunctions := defaultExtensions.GetAllWindowFunctions()
	assert.GreaterOrEqual(t, len(scalarFunctions), 309)
	assert.GreaterOrEqual(t, len(aggregateFunctions), 62)
	assert.GreaterOrEqual(t, len(windowFunctions), len(aggregateFunctions),
		"Should have at least as many window functions as aggregate functions due to aggregate function addition")
	tests := []struct {
		urn         string
		signature   string
		isScalar    bool
		isAggregate bool
		isWindow    bool
	}{
		{extensions.SubstraitDefaultURNPrefix + "functions_arithmetic", "add:i32_i32", true, false, false},
		{extensions.SubstraitDefaultURNPrefix + "functions_arithmetic", "variance:fp64", false, true, true},
		{extensions.SubstraitDefaultURNPrefix + "functions_arithmetic", "dense_rank:", false, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.signature, func(t *testing.T) {
			assert.True(t, tt.isScalar || tt.isAggregate || tt.isWindow)
			if tt.isScalar {
				sf, ok := defaultExtensions.GetScalarFunc(extensions.ID{URN: tt.urn, Name: tt.signature})
				assert.True(t, ok)
				assert.Contains(t, scalarFunctions, sf)
				// verify that default nullability is set to MIRROR
				assert.Equal(t, extensions.MirrorNullability, sf.Nullability())
			}
			if tt.isAggregate {
				af, ok := defaultExtensions.GetAggregateFunc(extensions.ID{URN: tt.urn, Name: tt.signature})
				assert.True(t, ok)
				assert.Contains(t, aggregateFunctions, af)
			}
			if tt.isWindow {
				wf, ok := defaultExtensions.GetWindowFunc(extensions.ID{URN: tt.urn, Name: tt.signature})
				assert.True(t, ok)
				assert.Contains(t, windowFunctions, wf)
			}
		})
	}
}

func TestAggregateToWindow(t *testing.T) {
	const uri = "http://localhost/sample.yaml"
	const urn = "extension:test:sample"

	var c extensions.Collection
	require.NoError(t, c.Load(urn, strings.NewReader(sampleYAML)))

	t.Run("aggregate functions available as window functions", func(t *testing.T) {
		// Test that the count function (with args) is available as both aggregate and window function
		aggFunc, ok := c.GetAggregateFunc(extensions.ID{URN: urn, Name: "count:any"})
		require.True(t, ok)
		require.NotNil(t, aggFunc)

		winFunc, ok := c.GetWindowFunc(extensions.ID{URN: urn, Name: "count:any"})
		require.True(t, ok)
		require.NotNil(t, winFunc)

		// Test that the count function (without args) is available as both aggregate and window function
		aggFuncNoArgs, ok := c.GetAggregateFunc(extensions.ID{URN: urn, Name: "count:"})
		require.True(t, ok)
		require.NotNil(t, aggFuncNoArgs)

		winFuncNoArgs, ok := c.GetWindowFunc(extensions.ID{URN: urn, Name: "count:"})
		require.True(t, ok)
		require.NotNil(t, winFuncNoArgs)
	})

	t.Run("window functions preserve aggregate properties", func(t *testing.T) {
		aggFunc, ok := c.GetAggregateFunc(extensions.ID{URN: urn, Name: "count:any"})
		require.True(t, ok)

		winFunc, ok := c.GetWindowFunc(extensions.ID{URN: urn, Name: "count:any"})
		require.True(t, ok)

		// Check that basic properties are preserved
		assert.Equal(t, aggFunc.Name(), winFunc.Name())
		assert.Equal(t, aggFunc.CompoundName(), winFunc.CompoundName())
		assert.Equal(t, aggFunc.Description(), winFunc.Description())
		assert.Equal(t, aggFunc.URN(), winFunc.URN())
		assert.Equal(t, aggFunc.Args(), winFunc.Args())
		assert.Equal(t, aggFunc.Options(), winFunc.Options())
		assert.Equal(t, aggFunc.Variadic(), winFunc.Variadic())
		assert.Equal(t, aggFunc.Deterministic(), winFunc.Deterministic())
		assert.Equal(t, aggFunc.SessionDependent(), winFunc.SessionDependent())
		assert.Equal(t, aggFunc.Nullability(), winFunc.Nullability())

		// Check that aggregate-specific properties are preserved
		assert.Equal(t, aggFunc.Decomposability(), winFunc.Decomposability())
		assert.Equal(t, aggFunc.Ordered(), winFunc.Ordered())
		assert.Equal(t, aggFunc.MaxSet(), winFunc.MaxSet())

		// Check intermediate type
		aggIntermediate, err := aggFunc.Intermediate()
		require.NoError(t, err)
		winIntermediate, err := winFunc.Intermediate()
		require.NoError(t, err)
		assert.Equal(t, aggIntermediate, winIntermediate)
	})

	t.Run("aggregate functions used as window functions have streaming window type", func(t *testing.T) {
		winFunc, ok := c.GetWindowFunc(extensions.ID{URN: urn, Name: "count:any"})
		require.True(t, ok)

		// Check that the window type is STREAMING
		assert.Equal(t, extensions.StreamingWindow, winFunc.WindowType())
	})

	t.Run("type resolution works the same", func(t *testing.T) {
		aggFunc, ok := c.GetAggregateFunc(extensions.ID{URN: urn, Name: "count:any"})
		require.True(t, ok)

		winFunc, ok := c.GetWindowFunc(extensions.ID{URN: urn, Name: "count:any"})
		require.True(t, ok)

		// Test type resolution with the same arguments
		// Use a concrete type for testing resolution
		i32Type := &types.Int32Type{Nullability: types.NullabilityRequired}
		aggType, err := aggFunc.ResolveType([]types.Type{i32Type}, extensions.NewSet())
		require.NoError(t, err)

		winType, err := winFunc.ResolveType([]types.Type{i32Type}, extensions.NewSet())
		require.NoError(t, err)

		assert.Equal(t, aggType, winType)
	})
}

func TestAggregateToWindowWithDefaultCollection(t *testing.T) {
	defaultExtensions := extensions.GetDefaultCollectionWithNoError()

	// Test cases for known aggregate functions that should be added as window functions
	testCases := []struct {
		urn          string
		functionName string
		description  string
	}{
		{
			urn:          extensions.SubstraitDefaultURNPrefix + "functions_aggregate_generic",
			functionName: "count:",
			description:  "count function without arguments",
		},
		{
			urn:          extensions.SubstraitDefaultURNPrefix + "functions_aggregate_generic",
			functionName: "count:any",
			description:  "count function with any argument",
		},
		{
			urn:          extensions.SubstraitDefaultURNPrefix + "functions_arithmetic",
			functionName: "variance:fp64",
			description:  "variance function",
		},
		{
			urn:          extensions.SubstraitDefaultURNPrefix + "functions_boolean",
			functionName: "bool_and:bool",
			description:  "bool_and function",
		},
		{
			urn:          extensions.SubstraitDefaultURNPrefix + "functions_aggregate_approx",
			functionName: "approx_count_distinct:any",
			description:  "approx_count_distinct function",
		},
		{
			urn:          extensions.SubstraitDefaultURNPrefix + "functions_string",
			functionName: "string_agg:str_str",
			description:  "string_agg function",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			id := extensions.ID{URN: tc.urn, Name: tc.functionName}

			// Verify the aggregate function exists
			aggFunc, ok := defaultExtensions.GetAggregateFunc(id)
			require.True(t, ok, "aggregate function %s should exist", tc.functionName)
			require.NotNil(t, aggFunc)

			// Verify the window function exists (added from aggregate)
			winFunc, ok := defaultExtensions.GetWindowFunc(id)
			require.True(t, ok, "window function %s should exist (added from aggregate)", tc.functionName)
			require.NotNil(t, winFunc)
		})
	}
}

func TestLoadExtensionWithoutURN(t *testing.T) {
	const extensionWithoutURN = `---
scalar_functions:
`

	var c extensions.Collection
	err := c.Load("http://localhost/test.yaml", strings.NewReader(extensionWithoutURN))

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "missing urn")
}

func TestLoadExtensionWithInvalidURN(t *testing.T) {
	const extensionWithInvalidURN = `---
urn: invalid:urn:format
scalar_functions:
`

	var c extensions.Collection
	err := c.Load("http://localhost/test.yaml", strings.NewReader(extensionWithInvalidURN))

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid urn")
}
