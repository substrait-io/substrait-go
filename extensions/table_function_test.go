// SPDX-License-Identifier: Apache-2.0

package extensions_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	substraitgo "github.com/substrait-io/substrait-go/v9"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/types"
)

const tableFunctionsURN = "extension:test:table_functions"

const tableFunctionsYAML = `%YAML 1.2
---
urn: extension:test:table_functions
table_functions:
  - name: generate_series
    description: Generates a series of integer values from start to stop.
    metadata:
      origin: functions_table.yaml
    impls:
      - args:
          - name: start
            value: i64
          - name: stop
            value: i64
          - name: step
            value: i64
        schema: struct<i64>
      - args:
          - name: start
            value: i32
          - name: stop
            value: i32
          - name: step
            value: i32
        schema: struct<i32>
  - name: repeat_value
    impls:
      - args:
          - name: value
            value: any1
          - name: count
            value: i64
        schema: struct<any1>
  - name: parse_records
    impls:
      - args:
          - name: payload
            value: string
  - name: split_rows
    impls:
      - args:
          - name: separator
            value: string
          - name: input
            value: string
        options:
          empty:
            values: [KEEP, DROP]
        variadic:
          min: 1
        schema: struct<string>
`

func loadTableFunctions(t *testing.T) *extensions.Collection {
	t.Helper()
	var c extensions.Collection
	require.NoError(t, c.Load(strings.NewReader(tableFunctionsYAML)))
	return &c
}

func tableFunctionID(name string) extensions.FunctionID {
	return extensions.FunctionID{URN: tableFunctionsURN, Name: name}
}

func TestLoadTableFunctions(t *testing.T) {
	c := loadTableFunctions(t)

	assert.Len(t, c.GetAllTableFunctions(), 5)

	series, ok := c.GetTableFunc(tableFunctionID("generate_series:i64_i64_i64"))
	require.True(t, ok)
	assert.Equal(t, "generate_series", series.Name())
	assert.Equal(t, "generate_series:i64_i64_i64", series.CompoundName())
	assert.Equal(t, tableFunctionID("generate_series:i64_i64_i64"), series.ID())
	assert.Equal(t, tableFunctionsURN, series.URN())
	assert.Equal(t, "Generates a series of integer values from start to stop.", series.Description())
	assert.Equal(t, map[string]any{"origin": "functions_table.yaml"}, series.Metadata())
	assert.Equal(t, "struct<i64>", series.Schema().String())
	require.Len(t, series.Args(), 3)
	assert.Equal(t, "step", series.Args()[2].(extensions.ValueArg).Name)
	assert.Nil(t, series.Variadic())
	assert.Equal(t, 3, series.MinArgumentCount())
	assert.Equal(t, 3, series.MaxArgumentCount())

	series32, ok := c.GetTableFunc(tableFunctionID("generate_series:i32_i32_i32"))
	require.True(t, ok)
	assert.Equal(t, "struct<i32>", series32.Schema().String())

	_, ok = c.GetTableFunc(tableFunctionID("generate_series"))
	assert.False(t, ok, "a name with several implementations needs its compound name")

	parse, ok := c.GetTableFunc(tableFunctionID("parse_records"))
	require.True(t, ok, "a name with one implementation is found by its simple name")
	assert.Equal(t, "parse_records:str", parse.CompoundName())
	assert.Nil(t, parse.Schema())

	split, ok := c.GetTableFunc(tableFunctionID("split_rows:str_str"))
	require.True(t, ok)
	require.NotNil(t, split.Variadic())
	assert.Equal(t, extensions.ConsistentParams, split.Variadic().ParameterConsistency)
	assert.Equal(t, []string{"KEEP", "DROP"}, split.Options()["empty"].Values)

	_, ok = c.GetScalarFunc(tableFunctionID("generate_series:i64_i64_i64"))
	assert.False(t, ok, "table functions are not scalar functions")
}

func TestTableFunctionResolveType(t *testing.T) {
	c := loadTableFunctions(t)
	reg := extensions.NewSet()
	i64 := &types.Int64Type{Nullability: types.NullabilityRequired}

	series, ok := c.GetTableFunc(tableFunctionID("generate_series:i64_i64_i64"))
	require.True(t, ok)
	schema, err := series.ResolveType([]types.Type{i64, i64, i64}, reg)
	require.NoError(t, err)
	assert.Equal(t, &types.StructType{Nullability: types.NullabilityRequired, Types: []types.Type{i64}}, schema)

	nullableI64 := &types.Int64Type{Nullability: types.NullabilityNullable}
	schema, err = series.ResolveType([]types.Type{nullableI64, i64, i64}, reg)
	require.NoError(t, err)
	assert.Equal(t, types.NullabilityRequired, schema.GetNullability(),
		"table functions follow DECLARED_OUTPUT nullability")

	repeat, ok := c.GetTableFunc(tableFunctionID("repeat_value"))
	require.True(t, ok)
	str := &types.StringType{Nullability: types.NullabilityRequired}
	schema, err = repeat.ResolveType([]types.Type{str, i64}, reg)
	require.NoError(t, err)
	assert.Equal(t, &types.StructType{Nullability: types.NullabilityRequired, Types: []types.Type{str}}, schema)

	parse, ok := c.GetTableFunc(tableFunctionID("parse_records"))
	require.True(t, ok)
	_, err = parse.ResolveType([]types.Type{str}, reg)
	assert.ErrorIs(t, err, substraitgo.ErrInvalidType)
	assert.ErrorContains(t, err, "parse_records:str does not declare a schema")
}

func TestTableFunctionMatch(t *testing.T) {
	c := loadTableFunctions(t)
	i32 := &types.Int32Type{Nullability: types.NullabilityRequired}
	i64 := &types.Int64Type{Nullability: types.NullabilityRequired}
	str := &types.StringType{Nullability: types.NullabilityNullable}

	series, ok := c.GetTableFunc(tableFunctionID("generate_series:i64_i64_i64"))
	require.True(t, ok)
	match, err := series.Match([]types.Type{i64, i64, i64})
	require.NoError(t, err)
	assert.True(t, match)
	match, err = series.Match([]types.Type{i32, i32, i32})
	require.NoError(t, err)
	assert.False(t, match)

	split, ok := c.GetTableFunc(tableFunctionID("split_rows:str_str"))
	require.True(t, ok)
	for _, args := range [][]types.Type{{str, str}, {str, str, str}} {
		match, err = split.Match(args)
		require.NoError(t, err)
		assert.True(t, match)
	}
	for _, args := range [][]types.Type{nil, {str}} {
		match, err = split.Match(args)
		require.NoError(t, err)
		assert.False(t, match, "%d argument(s) is fewer than a separator and one input", len(args))
	}
}

func TestLoadExtensionRejectsInvalidTableFunctionSchema(t *testing.T) {
	for _, schema := range []string{"i64", "struct?<i64>", "list<i64>"} {
		t.Run(schema, func(t *testing.T) {
			yaml := `---
urn: extension:test:invalid_table_function
table_functions:
  - name: bad
    impls:
      - schema: "` + schema + `"
`
			var c extensions.Collection
			err := c.Load(strings.NewReader(yaml))
			assert.ErrorIs(t, err, substraitgo.ErrInvalidSimpleExtention)
			assert.ErrorContains(t, err, `table function "bad" schema must be a non-nullable struct`)
			assert.False(t, c.URNLoaded("extension:test:invalid_table_function"))
		})
	}
}

func TestLoadExtensionRejectsUndeclaredTypeInTableFunction(t *testing.T) {
	for name, impl := range map[string]string{
		"argument": "      - args:\n          - value: u!point\n",
		"schema":   "      - schema: struct<u!point>\n",
	} {
		t.Run(name, func(t *testing.T) {
			yaml := "---\nurn: extension:test:table_function_udt\ntable_functions:\n  - name: points\n    impls:\n" + impl
			var c extensions.Collection
			err := c.Load(strings.NewReader(yaml))
			assert.ErrorIs(t, err, substraitgo.ErrInvalidSimpleExtention)
			assert.ErrorContains(t, err, `user-defined type "point" is not declared`)
		})
	}
}

func TestNewTableFuncVariant(t *testing.T) {
	v := extensions.NewTableFuncVariant(tableFunctionID("generate_series:i64_i64_i64"))
	assert.Equal(t, "generate_series", v.Name())
	assert.Equal(t, "generate_series:i64_i64_i64", v.CompoundName())
	assert.Len(t, v.Args(), 3)
	assert.Nil(t, v.Schema())
}
