package plan_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/plan"
	"github.com/substrait-io/substrait-go/v9/types"
)

var (
	v1 = expr.PrimitiveLiteral[int32]{Value: 1, Type: &types.Int32Type{Nullability: types.NullabilityRequired}}
	v2 = expr.PrimitiveLiteral[int32]{Value: 2, Type: &types.Int32Type{Nullability: types.NullabilityRequired}}
)

func buildLiteralExpressions(_ *testing.T, _ plan.Builder) []expr.VirtualTableExpressionValue {
	return []expr.VirtualTableExpressionValue{{&v1, &v2}}
}

func TestVirtualTableFromExpr(t *testing.T) {
	b := plan.NewBuilderDefault()
	values := buildLiteralExpressions(t, b)
	fieldNames := []string{"col0", "col1"}

	vt, err := b.VirtualTableFromExpr(fieldNames, values...)
	require.NoError(t, err)
	assert.Equal(t, 1, len(vt.Values()))
	assert.Equal(t, fieldNames, vt.BaseSchema().Names)
	assert.Equal(t, "struct<i32, i32>", vt.RecordType().String())

	vtRemap, err := vt.Remap(1)
	require.NoError(t, err)
	assert.Equal(t, "struct<i32>", vtRemap.RecordType().String())
}
