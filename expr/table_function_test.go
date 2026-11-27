// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	substraitgo "github.com/substrait-io/substrait-go/v9"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/types"
)

const tableFunctionsURN = "extension:test:table_functions"

const tableFunctionsYAML = `---
urn: extension:test:table_functions
table_functions:
  - name: generate_series
    impls:
      - args:
          - value: i64
          - value: i64
          - value: i64
        schema: struct<i64>
      - args:
          - value: i32
          - value: i32
          - value: i32
        schema: struct<i32>
  - name: parse_records
    impls:
      - args:
          - value: string
`

func newTableFunctionRegistry(t *testing.T) expr.ExtensionRegistry {
	t.Helper()
	c := &extensions.Collection{}
	require.NoError(t, c.Load(strings.NewReader(tableFunctionsYAML)))
	return expr.NewEmptyExtensionRegistry(c)
}

func requiredStruct(fields ...types.Type) *types.StructType {
	return &types.StructType{Nullability: types.NullabilityRequired, Types: fields}
}

var (
	requiredI64    = &types.Int64Type{Nullability: types.NullabilityRequired}
	requiredI32    = &types.Int32Type{Nullability: types.NullabilityRequired}
	requiredString = &types.StringType{Nullability: types.NullabilityRequired}
)

func seriesArgs[T int32 | int64](start, stop, step T) []types.FuncArg {
	return []types.FuncArg{
		expr.NewPrimitiveLiteral(start, false),
		expr.NewPrimitiveLiteral(stop, false),
		expr.NewPrimitiveLiteral(step, false),
	}
}

func TestNewTableFuncDerivesSchema(t *testing.T) {
	reg := newTableFunctionRegistry(t)
	id := extensions.FunctionID{URN: tableFunctionsURN, Name: "generate_series"}

	fn, err := expr.NewTableFunc(reg, id, nil, nil, seriesArgs[int64](1, 10, 1)...)
	require.NoError(t, err)
	assert.Equal(t, "generate_series:i64_i64_i64", fn.CompoundName())
	assert.Equal(t, requiredStruct(requiredI64), fn.Schema())
	assert.Equal(t, 3, fn.NArgs())
	assert.Equal(t, []types.Type{requiredI64, requiredI64, requiredI64}, fn.GetArgTypes())
	assert.Equal(t, "generate_series(i64(1), i64(10), i64(1)) => struct<i64>", fn.String())

	fn32, err := expr.NewTableFunc(reg, id, nil, nil, seriesArgs[int32](1, 10, 1)...)
	require.NoError(t, err)
	assert.Equal(t, "generate_series:i32_i32_i32", fn32.CompoundName())
	assert.Equal(t, requiredStruct(requiredI32), fn32.Schema())
	assert.NotEqual(t, fn.FuncRef(), fn32.FuncRef())

	decl, ok := reg.LookupTableFunction(fn.FuncRef())
	require.True(t, ok)
	assert.Equal(t, fn.ID(), decl.ID())
}

func TestNewTableFuncWithSchema(t *testing.T) {
	reg := newTableFunctionRegistry(t)
	series := extensions.FunctionID{URN: tableFunctionsURN, Name: "generate_series:i64_i64_i64"}
	parse := extensions.FunctionID{URN: tableFunctionsURN, Name: "parse_records"}
	payload := expr.NewPrimitiveLiteral("1,a;2,b", false)
	parsed := requiredStruct(requiredI64, requiredString)

	t.Run("matches the declared schema", func(t *testing.T) {
		fn, err := expr.NewTableFunc(reg, series, requiredStruct(requiredI64), nil, seriesArgs[int64](1, 10, 1)...)
		require.NoError(t, err)
		assert.Equal(t, requiredStruct(requiredI64), fn.Schema())
	})

	t.Run("producer-supplied schema", func(t *testing.T) {
		opts := []*types.FunctionOption{{Name: "header", Preference: []string{"NONE"}}}
		fn, err := expr.NewTableFunc(reg, parse, parsed, opts, payload)
		require.NoError(t, err)
		assert.Equal(t, "parse_records:str", fn.CompoundName())
		assert.Equal(t, parsed, fn.Schema())
		assert.Equal(t, []string{"NONE"}, fn.GetOption("header"))
		assert.Equal(t, "parse_records(string(1,a;2,b), {header: [NONE]}) => struct<i64, string>", fn.String())
	})

	for name, tc := range map[string]struct {
		id      extensions.FunctionID
		schema  *types.StructType
		args    []types.FuncArg
		wantErr error
		msg     string
	}{
		"schema differs from the declared schema": {
			id: series, schema: requiredStruct(requiredI32), args: seriesArgs[int64](1, 10, 1),
			wantErr: substraitgo.ErrInvalidType, msg: "schema struct<i32> does not match the declared schema struct<i64>",
		},
		"no schema declared or provided": {
			id: parse, args: []types.FuncArg{payload},
			wantErr: substraitgo.ErrInvalidType, msg: "parse_records:str does not declare a schema, one must be provided",
		},
		"arguments don't match a producer-resolved function": {
			id: parse, schema: parsed, args: []types.FuncArg{expr.NewPrimitiveLiteral(int64(1), false)},
			wantErr: substraitgo.ErrInvalidType, msg: "argument types did not match table function parse_records:str",
		},
		"nullable schema": {
			id: parse, schema: &types.StructType{Nullability: types.NullabilityNullable, Types: parsed.Types}, args: []types.FuncArg{payload},
			wantErr: substraitgo.ErrInvalidType, msg: "schema must be a non-nullable struct",
		},
		"unknown function": {
			id: extensions.FunctionID{URN: tableFunctionsURN, Name: "unknown"}, args: seriesArgs[int64](1, 10, 1),
			wantErr: substraitgo.ErrNotFound,
		},
	} {
		t.Run(name, func(t *testing.T) {
			_, err := expr.NewTableFunc(reg, tc.id, tc.schema, nil, tc.args...)
			assert.ErrorIs(t, err, tc.wantErr)
			assert.ErrorContains(t, err, tc.msg)
		})
	}
}

func TestNewCustomTableFunc(t *testing.T) {
	reg := newTableFunctionRegistry(t)
	v := extensions.NewTableFuncVariant(extensions.FunctionID{URN: "extension:test:custom", Name: "read_rows:i64"})
	payload := expr.NewPrimitiveLiteral(int64(10), false)

	fn, err := expr.NewCustomTableFunc(reg, v, requiredStruct(requiredString), nil, payload)
	require.NoError(t, err)
	assert.Equal(t, "read_rows:i64", fn.CompoundName())
	assert.Equal(t, reg.GetFuncAnchor(v.ID()), fn.FuncRef())

	_, err = expr.NewCustomTableFunc(reg, v, nil, nil, payload)
	assert.ErrorIs(t, err, substraitgo.ErrInvalidType)
}

func TestTableFunctionCloneAndEquals(t *testing.T) {
	reg := newTableFunctionRegistry(t)
	id := extensions.FunctionID{URN: tableFunctionsURN, Name: "generate_series"}
	fn, err := expr.NewTableFunc(reg, id, nil, nil, seriesArgs[int64](1, 10, 1)...)
	require.NoError(t, err)

	clone := fn.Clone()
	assert.True(t, fn.Equals(clone))

	clone.SetArg(1, expr.NewPrimitiveLiteral(int64(20), false))
	assert.False(t, fn.Equals(clone))
	assert.Equal(t, "i64(10)", fn.Arg(1).String(), "SetArg on a clone leaves the original unchanged")

	fromParts := expr.NewTableFunctionFromParts(fn.FuncRef(), nil, seriesArgs[int64](1, 10, 1), nil, requiredStruct(requiredI64))
	assert.True(t, fn.Equals(fromParts))
	assert.False(t, fn.Equals(nil))
}
