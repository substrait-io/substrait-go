// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	substraitgo "github.com/substrait-io/substrait-go/v9"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/plan"
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
  - name: generate_pairs
    impls:
      - args:
          - value: i64
        schema: struct<i64, i64>
`

func newTableFunctionBuilder(t *testing.T) plan.Builder {
	t.Helper()
	c := &extensions.Collection{}
	require.NoError(t, c.Load(strings.NewReader(tableFunctionsYAML)))
	return plan.NewBuilder(c)
}

func i64Literal(v int64) expr.Literal {
	return expr.NewPrimitiveLiteral(v, false)
}

func TestTableFunctionRel(t *testing.T) {
	b := newTableFunctionBuilder(t)
	i64 := &types.Int64Type{Nullability: types.NullabilityRequired}

	fn, err := b.TableFn(tableFunctionsURN, "generate_series", nil, nil, i64Literal(1), i64Literal(100), i64Literal(1))
	require.NoError(t, err)
	rel, err := b.TableFunction(fn)
	require.NoError(t, err)

	assert.Same(t, fn, rel.Function())
	assert.Empty(t, rel.GetInputs())
	assert.Equal(t, *types.NewRecordTypeFromTypes([]types.Type{i64}), rel.RecordType())

	p, err := b.Plan(rel, []string{"value"})
	require.NoError(t, err)
	require.Len(t, p.GetRoots(), 1)
	assert.Equal(t, []string{"value"}, p.GetRoots()[0].RecordType().Names)

	reg := p.ExtensionRegistry()
	_, ok := reg.LookupTableFunction(fn.FuncRef())
	assert.True(t, ok, "the plan's extension set includes the table function")

	_, err = b.Plan(rel, []string{"value", "extra"})
	assert.ErrorIs(t, err, substraitgo.ErrInvalidRel)
}

func TestTableFunctionRelRemap(t *testing.T) {
	b := newTableFunctionBuilder(t)
	fn, err := b.TableFn(tableFunctionsURN, "generate_pairs", nil, nil, i64Literal(3))
	require.NoError(t, err)
	rel, err := b.TableFunction(fn)
	require.NoError(t, err)

	remapped, err := rel.Remap(1)
	require.NoError(t, err)
	assert.Equal(t, []int32{1}, remapped.OutputMapping())
	assert.Equal(t, int32(1), remapped.RecordType().FieldCount())
	assert.Nil(t, rel.OutputMapping(), "Remap leaves the original relation unchanged")
	assert.Equal(t, int32(2), rel.RecordType().FieldCount())

	_, err = rel.Remap(2)
	assert.ErrorIs(t, err, substraitgo.ErrInvalidRel)
}

func TestTableFunctionRelCopy(t *testing.T) {
	b := newTableFunctionBuilder(t)
	fn, err := b.TableFn(tableFunctionsURN, "generate_series", nil, nil, i64Literal(1), i64Literal(10), i64Literal(1))
	require.NoError(t, err)
	rel, err := b.TableFunction(fn)
	require.NoError(t, err)

	copied, err := rel.Copy()
	require.NoError(t, err)
	assert.NotSame(t, rel, copied)
	assert.Same(t, fn, copied.(*plan.TableFunctionRel).Function())

	_, err = rel.Copy(rel)
	assert.ErrorIs(t, err, substraitgo.ErrInvalidInputCount)
	_, err = rel.CopyWithExpressionRewrite(func(e expr.Expression) (expr.Expression, error) { return e, nil }, rel)
	assert.ErrorIs(t, err, substraitgo.ErrInvalidInputCount)

	unchanged, err := rel.CopyWithExpressionRewrite(func(e expr.Expression) (expr.Expression, error) { return e, nil })
	require.NoError(t, err)
	assert.Same(t, rel, unchanged)

	rewritten, err := rel.CopyWithExpressionRewrite(func(e expr.Expression) (expr.Expression, error) {
		if e.Equals(i64Literal(10)) {
			return i64Literal(20), nil
		}
		return e, nil
	})
	require.NoError(t, err)
	assert.NotSame(t, rel, rewritten)
	assert.Equal(t, "i64(20)", rewritten.(*plan.TableFunctionRel).Function().Arg(1).String())
	assert.Equal(t, "i64(10)", rel.Function().Arg(1).String())

	_, err = rel.CopyWithExpressionRewrite(func(expr.Expression) (expr.Expression, error) {
		return nil, substraitgo.ErrNotImplemented
	})
	assert.ErrorIs(t, err, substraitgo.ErrNotImplemented)
}

func TestTableFunctionRelAdvancedExtension(t *testing.T) {
	b := newTableFunctionBuilder(t)
	fn, err := b.TableFn(tableFunctionsURN, "generate_pairs", nil, nil, i64Literal(3))
	require.NoError(t, err)
	rel, err := b.TableFunction(fn)
	require.NoError(t, err)

	ext := &extensions.AdvancedExtension{}
	assert.Nil(t, rel.SetAdvancedExtension(ext))
	assert.Same(t, ext, rel.GetAdvancedExtension())
}

func TestTableFunctionRelErrors(t *testing.T) {
	b := newTableFunctionBuilder(t)

	_, err := b.TableFunction(nil)
	assert.ErrorIs(t, err, substraitgo.ErrInvalidRel)

	nullable := &types.StructType{Nullability: types.NullabilityNullable, Types: []types.Type{&types.Int64Type{}}}
	fn := expr.NewTableFunctionFromParts(1, nil, nil, nil, nullable)
	_, err = b.TableFunction(fn)
	assert.ErrorIs(t, err, substraitgo.ErrInvalidRel)
	assert.ErrorContains(t, err, "table function schema must be a non-nullable struct")

	_, err = b.TableFn(tableFunctionsURN, "generate_series", nil, nil, i64Literal(1))
	assert.ErrorIs(t, err, substraitgo.ErrInvalidExpr)
	assert.ErrorContains(t, err, "got 1, expected 3")

	_, err = b.TableFn(tableFunctionsURN, "unknown", nil, nil, i64Literal(1))
	assert.ErrorIs(t, err, substraitgo.ErrNotFound)
}
