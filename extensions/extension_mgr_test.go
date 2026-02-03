// SPDX-License-Identifier: Apache-2.0

package extensions_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v7/extensions"
	"github.com/substrait-io/substrait-go/v7/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	extensionspb "github.com/substrait-io/substrait-protobuf/go/substraitpb/extensions"
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
	require.NoError(t, c.Load(uri, strings.NewReader(sampleYAML)))

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

func TestCannotLoadDuplicateURI(t *testing.T) {
	const ext1 = `---
urn: extension:urn:ext1
scalar_functions:
`
	const ext2 = `---
urn: extension:urn:ext2
scalar_functions:
`

	var c extensions.Collection
	err := c.Load("http://localhost/test.yaml", strings.NewReader(ext1))
	assert.NoError(t, err)

	err = c.Load("http://localhost/test.yaml", strings.NewReader(ext2))

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "already loaded")
}

func TestCannotLoadDuplicateURN(t *testing.T) {
	const extension = `---
urn: extension:urn:format
scalar_functions:
`

	var c extensions.Collection
	err := c.Load("http://localhost/test.yaml", strings.NewReader(extension))
	assert.NoError(t, err)

	err = c.Load("http://localhost/test2.yaml", strings.NewReader(extension))

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "already loaded")
}

func TestGetExtensionSetWithURIAndURN(t *testing.T) {
	// Create a plan with 3 extension functions demonstrating different reference patterns:
	// 1. Only extensionUrnReference (URN anchor 1)
	// 2. Only extensionUriReference (URI anchor 1)
	// 3. Both extensionUrnReference and extensionUriReference (should validate consistency)
	plan := &proto.Plan{
		ExtensionUris: []*extensionspb.SimpleExtensionURI{
			{
				ExtensionUriAnchor: 1,
				Uri:                "https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml",
			},
		},
		ExtensionUrns: []*extensionspb.SimpleExtensionURN{
			{
				ExtensionUrnAnchor: 11,
				Urn:                "extension:io.substrait:functions_arithmetic",
			},
		},
		Extensions: []*extensionspb.SimpleExtensionDeclaration{
			{
				MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction_{
					ExtensionFunction: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction{
						ExtensionUrnReference: 11,
						FunctionAnchor:        10,
						Name:                  "add:i32_i32",
					},
				},
			},
			{
				MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction_{
					ExtensionFunction: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction{
						ExtensionUriReference: 1,
						FunctionAnchor:        20,
						Name:                  "multiply:i32_i32",
					},
				},
			},
			{
				MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction_{
					ExtensionFunction: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction{
						ExtensionUrnReference: 11,
						ExtensionUriReference: 1,
						FunctionAnchor:        30,
						Name:                  "divide:i32_i32",
					},
				},
			},
		},
		Relations: []*proto.PlanRel{},
	}

	collection := extensions.GetDefaultCollectionWithNoError()
	extSet, err := extensions.GetExtensionSet(plan, collection)
	require.NoError(t, err)

	// Extension 1: extensionUrnReference pointing to URN anchor
	id1, ok := extSet.DecodeFunc(10)
	require.True(t, ok)
	assert.Equal(t, "extension:io.substrait:functions_arithmetic", id1.URN)
	assert.Equal(t, "add:i32_i32", id1.Name)

	// Extension 2: extensionUriReference pointing to URI anchor
	id2, ok := extSet.DecodeFunc(20)
	require.True(t, ok)
	assert.Equal(t, "extension:io.substrait:functions_arithmetic", id2.URN)
	assert.Equal(t, "multiply:i32_i32", id2.Name)

	// Extension 3: Both extensionUrnReference and extensionUriReference  - should validate consistency
	id3, ok := extSet.DecodeFunc(30)
	require.True(t, ok)
	assert.Equal(t, "extension:io.substrait:functions_arithmetic", id3.URN)
	assert.Equal(t, "divide:i32_i32", id3.Name)
}

func TestGetExtensionSetWithUnknownURI(t *testing.T) {
	// Create a plan with a URI that doesn't have a URN mapping in the collection
	plan := &proto.Plan{
		ExtensionUris: []*extensionspb.SimpleExtensionURI{
			{
				ExtensionUriAnchor: 1,
				Uri:                "https://example.com/unknown_extension.yaml",
			},
		},
		Extensions: []*extensionspb.SimpleExtensionDeclaration{
			{
				MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction_{
					ExtensionFunction: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction{
						ExtensionUriReference: 1,
						FunctionAnchor:        10,
						Name:                  "unknown_function",
					},
				},
			},
		},
		Relations: []*proto.PlanRel{},
	}

	collection := extensions.GetDefaultCollectionWithNoError()
	extSet, err := extensions.GetExtensionSet(plan, collection)
	require.Error(t, err)
	require.Nil(t, extSet)
	assert.Contains(t, err.Error(), "extension URI not resolvable")
}

func TestGetExtensionSetWithMissingAnchor(t *testing.T) {
	// Create a plan with a reference to an anchor that doesn't exist
	plan := &proto.Plan{
		ExtensionUris: []*extensionspb.SimpleExtensionURI{
			{
				ExtensionUriAnchor: 1,
				Uri:                "https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml",
			},
		},
		Extensions: []*extensionspb.SimpleExtensionDeclaration{
			{
				MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction_{
					ExtensionFunction: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction{
						ExtensionUriReference: 99, // This anchor doesn't exist!
						FunctionAnchor:        10,
						Name:                  "unknown_function",
					},
				},
			},
		},
		Relations: []*proto.PlanRel{},
	}

	collection := extensions.GetDefaultCollectionWithNoError()
	extSet, err := extensions.GetExtensionSet(plan, collection)
	require.Error(t, err)
	require.Nil(t, extSet)
	assert.Contains(t, err.Error(), "unable to resolve extension reference")
}

// TestGetExtensionSetWithInvalidURN tests that URN references to non-existent URNs are rejected
func TestGetExtensionSetWithInvalidURN(t *testing.T) {
	// Create a plan with a URN that doesn't exist in the collection
	plan := &proto.Plan{
		ExtensionUrns: []*extensionspb.SimpleExtensionURN{
			{
				ExtensionUrnAnchor: 1,
				Urn:                "extension:nonexistent:extension",
			},
		},
		Extensions: []*extensionspb.SimpleExtensionDeclaration{
			{
				MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction_{
					ExtensionFunction: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction{
						ExtensionUrnReference: 1,
						FunctionAnchor:        10,
						Name:                  "nonexistent_function",
					},
				},
			},
		},
		Relations: []*proto.PlanRel{},
	}

	collection := extensions.GetDefaultCollectionWithNoError()
	extSet, err := extensions.GetExtensionSet(plan, collection)
	require.Error(t, err)
	require.Nil(t, extSet)
	assert.Contains(t, err.Error(), "not found")
	assert.Contains(t, err.Error(), "extension:nonexistent:extension")
}

// TestExtensionValidationEdgeCases tests various edge cases in extension validation
func TestExtensionValidationEdgeCases(t *testing.T) {
	collection := extensions.GetDefaultCollectionWithNoError()

	t.Run("empty URN string", func(t *testing.T) {
		plan := &proto.Plan{
			ExtensionUrns: []*extensionspb.SimpleExtensionURN{
				{ExtensionUrnAnchor: 1, Urn: ""}, // Empty URN
			},
			Extensions: []*extensionspb.SimpleExtensionDeclaration{
				{
					MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction_{
						ExtensionFunction: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction{
							ExtensionUrnReference: 1,
							FunctionAnchor:        10,
							Name:                  "test_function",
						},
					},
				},
			},
		}

		_, err := extensions.GetExtensionSet(plan, collection)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("zero URN reference", func(t *testing.T) {
		plan := &proto.Plan{
			Extensions: []*extensionspb.SimpleExtensionDeclaration{
				{
					MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction_{
						ExtensionFunction: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction{
							ExtensionUrnReference: 0, // Zero reference - should be treated as "no reference"
							ExtensionUriReference: 0, // Both zero
							FunctionAnchor:        10,
							Name:                  "test_function",
						},
					},
				},
			},
		}

		_, err := extensions.GetExtensionSet(plan, collection)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "unable to resolve extension reference")
	})

	t.Run("both URI and URN resolve to same extension", func(t *testing.T) {
		// This documents that when both URI and URN are provided, URN takes precedence
		plan := &proto.Plan{
			ExtensionUris: []*extensionspb.SimpleExtensionURI{
				{ExtensionUriAnchor: 1, Uri: "https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml"},
			},
			ExtensionUrns: []*extensionspb.SimpleExtensionURN{
				{ExtensionUrnAnchor: 1, Urn: "extension:io.substrait:functions_arithmetic"},
			},
			Extensions: []*extensionspb.SimpleExtensionDeclaration{
				{
					MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction_{
						ExtensionFunction: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction{
							ExtensionUriReference: 1, // Points to arithmetic functions
							ExtensionUrnReference: 1, // Also points to arithmetic functions (same extension)
							FunctionAnchor:        10,
							Name:                  "add:i32_i32",
						},
					},
				},
			},
		}

		extSet, err := extensions.GetExtensionSet(plan, collection)
		require.NoError(t, err)

		id, ok := extSet.DecodeFunc(10)
		require.True(t, ok)
		assert.Equal(t, "extension:io.substrait:functions_arithmetic", id.URN)
		assert.Equal(t, "add:i32_i32", id.Name)
	})
}

func TestToProtoPopulatesBothURNAndURI(t *testing.T) {
	c := &extensions.Collection{}
	err := c.Load("some/uri", strings.NewReader(sampleYAML))
	require.NoError(t, err)

	plan := &proto.Plan{
		ExtensionUrns: []*extensionspb.SimpleExtensionURN{
			{ExtensionUrnAnchor: 1, Urn: "extension:test:sample"},
		},
		Extensions: []*extensionspb.SimpleExtensionDeclaration{
			{
				MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction_{
					ExtensionFunction: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction{
						ExtensionUrnReference: 1,
						FunctionAnchor:        1,
						Name:                  "add:i8_i8",
					},
				},
			},
		},
	}

	extSet, err := extensions.GetExtensionSet(plan, c)
	require.NoError(t, err)

	urns, uris, decls := extSet.ToProto(c)

	expectedUrns := []*extensionspb.SimpleExtensionURN{
		{ExtensionUrnAnchor: 1, Urn: "extension:test:sample"},
	}
	expectedUris := []*extensionspb.SimpleExtensionURI{
		{ExtensionUriAnchor: 1, Uri: "some/uri"},
	}
	expectedDecls := []*extensionspb.SimpleExtensionDeclaration{
		{
			MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction_{
				ExtensionFunction: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction{
					ExtensionUrnReference: 1,
					ExtensionUriReference: 1,
					FunctionAnchor:        1,
					Name:                  "add:i8_i8",
				},
			},
		},
	}

	assert.Equal(t, expectedUrns, urns)
	assert.Equal(t, expectedUris, uris)
	assert.Equal(t, expectedDecls, decls)
}

func TestToProtoPopulatesBothURNAndURIFromURIOnly(t *testing.T) {
	c := &extensions.Collection{}
	err := c.Load("some/uri", strings.NewReader(sampleYAML))
	require.NoError(t, err)

	plan := &proto.Plan{
		ExtensionUris: []*extensionspb.SimpleExtensionURI{
			{ExtensionUriAnchor: 1, Uri: "some/uri"},
		},
		Extensions: []*extensionspb.SimpleExtensionDeclaration{
			{
				MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction_{
					ExtensionFunction: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction{
						ExtensionUriReference: 1, // Only URI reference, no URN reference
						FunctionAnchor:        1,
						Name:                  "add:i8_i8",
					},
				},
			},
		},
	}

	extSet, err := extensions.GetExtensionSet(plan, c)
	require.NoError(t, err)

	urns, uris, decls := extSet.ToProto(c)

	expectedUrns := []*extensionspb.SimpleExtensionURN{
		{ExtensionUrnAnchor: 1, Urn: "extension:test:sample"},
	}
	expectedUris := []*extensionspb.SimpleExtensionURI{
		{ExtensionUriAnchor: 1, Uri: "some/uri"},
	}
	expectedDecls := []*extensionspb.SimpleExtensionDeclaration{
		{
			MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction_{
				ExtensionFunction: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction{
					ExtensionUrnReference: 1,
					ExtensionUriReference: 1,
					FunctionAnchor:        1,
					Name:                  "add:i8_i8",
				},
			},
		},
	}

	assert.Equal(t, expectedUrns, urns)
	assert.Equal(t, expectedUris, uris)
	assert.Equal(t, expectedDecls, decls)
}

func TestResolveRefToURNAllConditions(t *testing.T) {
	// Setup a collection with known URN/URI mappings
	c := &extensions.Collection{}
	err := c.Load("some/uri", strings.NewReader(sampleYAML))
	require.NoError(t, err)

	// Note: In protobuf, omitted fields get zero values, so many "zero reference" cases
	// are actually the same. These tests cover all truly distinct resolution paths.

	t.Run("1. non-zero URN reference found and valid", func(t *testing.T) {
		plan := &proto.Plan{
			ExtensionUrns: []*extensionspb.SimpleExtensionURN{
				{ExtensionUrnAnchor: 1, Urn: "extension:test:sample"},
			},
			Extensions: []*extensionspb.SimpleExtensionDeclaration{
				{
					MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction_{
						ExtensionFunction: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction{
							ExtensionUrnReference: 1, // Non-zero URN reference
							FunctionAnchor:        1,
							Name:                  "test_function",
						},
					},
				},
			},
		}

		extSet, err := extensions.GetExtensionSet(plan, c)
		require.NoError(t, err)
		require.NotNil(t, extSet)

		// Should successfully resolve to the URN
		id, ok := extSet.DecodeFunc(1)
		require.True(t, ok)
		assert.Equal(t, "extension:test:sample", id.URN)
	})

	t.Run("2. non-zero URN reference found but invalid", func(t *testing.T) {
		plan := &proto.Plan{
			ExtensionUrns: []*extensionspb.SimpleExtensionURN{
				{ExtensionUrnAnchor: 1, Urn: "extension:nonexistent:urn"},
			},
			Extensions: []*extensionspb.SimpleExtensionDeclaration{
				{
					MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction_{
						ExtensionFunction: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction{
							ExtensionUrnReference: 1, // Non-zero URN reference to invalid URN
							FunctionAnchor:        1,
							Name:                  "test_function",
						},
					},
				},
			},
		}

		_, err := extensions.GetExtensionSet(plan, c)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "URN 'extension:nonexistent:urn' not found in extension collection")
	})

	t.Run("3. non-zero URI reference found and resolvable", func(t *testing.T) {
		plan := &proto.Plan{
			ExtensionUris: []*extensionspb.SimpleExtensionURI{
				{ExtensionUriAnchor: 1, Uri: "some/uri"},
			},
			Extensions: []*extensionspb.SimpleExtensionDeclaration{
				{
					MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction_{
						ExtensionFunction: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction{
							ExtensionUriReference: 1, // Non-zero URI reference
							FunctionAnchor:        1,
							Name:                  "test_function",
						},
					},
				},
			},
		}

		extSet, err := extensions.GetExtensionSet(plan, c)
		require.NoError(t, err)
		require.NotNil(t, extSet)

		// Should successfully resolve URI to URN
		id, ok := extSet.DecodeFunc(1)
		require.True(t, ok)
		assert.Equal(t, "extension:test:sample", id.URN)
	})

	t.Run("4. non-zero URI reference found but not resolvable", func(t *testing.T) {
		plan := &proto.Plan{
			ExtensionUris: []*extensionspb.SimpleExtensionURI{
				{ExtensionUriAnchor: 1, Uri: "unknown/uri"},
			},
			Extensions: []*extensionspb.SimpleExtensionDeclaration{
				{
					MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction_{
						ExtensionFunction: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction{
							ExtensionUriReference: 1, // Non-zero URI reference to unknown URI
							FunctionAnchor:        1,
							Name:                  "test_function",
						},
					},
				},
			},
		}

		_, err := extensions.GetExtensionSet(plan, c)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "cannot resolve URI 'unknown/uri' to URN")
	})

	t.Run("5. both zero URN and URI references found and consistent", func(t *testing.T) {
		plan := &proto.Plan{
			ExtensionUrns: []*extensionspb.SimpleExtensionURN{
				{ExtensionUrnAnchor: 0, Urn: "extension:test:sample"}, // Zero anchor
			},
			ExtensionUris: []*extensionspb.SimpleExtensionURI{
				{ExtensionUriAnchor: 0, Uri: "some/uri"}, // Zero anchor
			},
			Extensions: []*extensionspb.SimpleExtensionDeclaration{
				{
					MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction_{
						ExtensionFunction: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction{
							FunctionAnchor: 1,
							Name:           "test_function",
						},
					},
				},
			},
		}

		extSet, err := extensions.GetExtensionSet(plan, c)
		require.NoError(t, err)
		require.NotNil(t, extSet)

		// Should successfully resolve since URN and URI are consistent
		id, ok := extSet.DecodeFunc(1)
		require.True(t, ok)
		assert.Equal(t, "extension:test:sample", id.URN)
	})

	t.Run("6. both zero URN and URI references found but inconsistent", func(t *testing.T) {
		plan := &proto.Plan{
			ExtensionUrns: []*extensionspb.SimpleExtensionURN{
				{Urn: "extension:wrong:urn"}, // Zero anchor, wrong URN
			},
			ExtensionUris: []*extensionspb.SimpleExtensionURI{
				{Uri: "some/uri"}, // Zero anchor, correct URI
			},
			Extensions: []*extensionspb.SimpleExtensionDeclaration{
				{
					MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction_{
						ExtensionFunction: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction{
							// Missing URN/URI reference equivalent to 0
							FunctionAnchor: 1,
							Name:           "test_function",
						},
					},
				},
			},
		}

		_, err := extensions.GetExtensionSet(plan, c)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "URN mismatch: found URN \"extension:wrong:urn\" but expected \"extension:test:sample\" for URI \"some/uri\"")
	})

	// Note: Cases 7-10 are redundant with case 5 since omitted protobuf fields = zero values
	// The truly distinct "zero reference" cases are already covered above

	t.Run("11. neither URN nor URI reference resolvable", func(t *testing.T) {
		plan := &proto.Plan{
			Extensions: []*extensionspb.SimpleExtensionDeclaration{
				{
					MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction_{
						ExtensionFunction: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction{
							ExtensionUrnReference: 99, // Non-existent URN reference
							ExtensionUriReference: 99, // Non-existent URI reference
							FunctionAnchor:        1,
							Name:                  "test_function",
						},
					},
				},
			},
		}

		_, err := extensions.GetExtensionSet(plan, c)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "unable to resolve extension reference: neither URN reference 99 nor URI reference 99 could be resolved")
	})

	t.Run("12. zero URN and URI references but neither resolvable", func(t *testing.T) {
		plan := &proto.Plan{
			Extensions: []*extensionspb.SimpleExtensionDeclaration{
				{
					MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction_{
						ExtensionFunction: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction{
							ExtensionUrnReference: 0, // Zero URN reference (no URN at anchor 0)
							ExtensionUriReference: 0, // Zero URI reference (no URI at anchor 0)
							FunctionAnchor:        1,
							Name:                  "test_function",
						},
					},
				},
			},
		}

		_, err := extensions.GetExtensionSet(plan, c)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "unable to resolve extension reference: neither URN reference 0 nor URI reference 0 could be resolved")
	})
}

func TestStructReturnTypeWithAny(t *testing.T) {
	// Test for #182: struct return types with polymorphic (any) fields
	const structAnyYAML = `---
urn: extension:test:struct_any
scalar_functions:
  - name: "wrap_in_struct"
    description: "Wraps a value in a struct"
    impls:
      - args:
          - name: value
            value: any1
        return: struct<any1>
  - name: "make_map"
    description: "Creates a map with any key and value types"
    impls:
      - args:
          - name: key
            value: any1
          - name: value
            value: any2
        return: map<any1, any2>
`

	const uri = "http://localhost/struct_any.yaml"
	const urn = "extension:test:struct_any"

	var c extensions.Collection
	require.NoError(t, c.Load(uri, strings.NewReader(structAnyYAML)))

	t.Run("scalar function with struct<any1>", func(t *testing.T) {
		fn, ok := c.GetScalarFunc(extensions.ID{URN: urn, Name: "wrap_in_struct:any"})
		require.True(t, ok)
		require.NotNil(t, fn)

		// Test with i64
		i64Type := &types.Int64Type{Nullability: types.NullabilityRequired}
		result, err := fn.ResolveType([]types.Type{i64Type}, extensions.NewSet())
		require.NoError(t, err)

		structType, ok := result.(*types.StructType)
		require.True(t, ok)
		require.Len(t, structType.Types, 1)
		assert.Equal(t, i64Type, structType.Types[0])

		// Test with string
		stringType := &types.StringType{Nullability: types.NullabilityRequired}
		result, err = fn.ResolveType([]types.Type{stringType}, extensions.NewSet())
		require.NoError(t, err)

		structType, ok = result.(*types.StructType)
		require.True(t, ok)
		require.Len(t, structType.Types, 1)
		assert.Equal(t, stringType, structType.Types[0])
	})

	t.Run("scalar function with map<any1, any2>", func(t *testing.T) {
		fn, ok := c.GetScalarFunc(extensions.ID{URN: urn, Name: "make_map:any_any"})
		require.True(t, ok)
		require.NotNil(t, fn)

		stringType := &types.StringType{Nullability: types.NullabilityRequired}
		i64Type := &types.Int64Type{Nullability: types.NullabilityRequired}

		result, err := fn.ResolveType([]types.Type{stringType, i64Type}, extensions.NewSet())
		require.NoError(t, err)

		mapType, ok := result.(*types.MapType)
		require.True(t, ok)
		assert.Equal(t, stringType, mapType.Key)
		assert.Equal(t, i64Type, mapType.Value)
	})
}

func TestUserDefinedTypeReturnWithAny(t *testing.T) {
	// Test for #184: user-defined type return types with polymorphic (any) parameters
	const udtAnyYAML = `---
urn: extension:test:udt_any
types:
  - name: Wrapper
    structure:
      value: T
scalar_functions:
  - name: "wrap"
    description: "Wraps a value in a user-defined type"
    impls:
      - args:
          - name: value
            value: any1
        return: u!Wrapper<any1>
`

	const uri = "http://localhost/udt_any.yaml"
	const urn = "extension:test:udt_any"

	var c extensions.Collection
	require.NoError(t, c.Load(uri, strings.NewReader(udtAnyYAML)))

	fn, ok := c.GetScalarFunc(extensions.ID{URN: urn, Name: "wrap:any"})
	require.True(t, ok)
	require.NotNil(t, fn)

	// Test with i64
	i64Type := &types.Int64Type{Nullability: types.NullabilityRequired}
	result, err := fn.ResolveType([]types.Type{i64Type}, extensions.NewSet())
	require.NoError(t, err)

	udtType, ok := result.(*types.UserDefinedType)
	require.True(t, ok)
	require.Len(t, udtType.TypeParameters, 1)

	dataParam, ok := udtType.TypeParameters[0].(*types.DataTypeParameter)
	require.True(t, ok)
	assert.Equal(t, i64Type, dataParam.Type)

	// Test with string
	stringType := &types.StringType{Nullability: types.NullabilityRequired}
	result, err = fn.ResolveType([]types.Type{stringType}, extensions.NewSet())
	require.NoError(t, err)

	udtType, ok = result.(*types.UserDefinedType)
	require.True(t, ok)
	require.Len(t, udtType.TypeParameters, 1)

	dataParam, ok = udtType.TypeParameters[0].(*types.DataTypeParameter)
	require.True(t, ok)
	assert.Equal(t, stringType, dataParam.Type)
}
