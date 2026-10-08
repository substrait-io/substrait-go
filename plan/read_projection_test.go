// SPDX-License-Identifier: Apache-2.0

package plan

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/types"
)

func projectionTestSchema() types.NamedStruct {
	return types.NamedStruct{
		Names: []string{"c0", "c1", "c2"},
		Struct: types.StructType{Nullability: types.NullabilityRequired, Types: []types.Type{
			&types.Int64Type{Nullability: types.NullabilityRequired},
			&types.StringType{Nullability: types.NullabilityNullable},
			&types.BooleanType{Nullability: types.NullabilityNullable},
		}},
	}
}

func TestReadProjectionBeforeEmit(t *testing.T) {
	schema := projectionTestSchema()
	projection := expr.NewMaskExpression(expr.MaskStructSelect{
		expr.NewMaskStructItem(0, nil), expr.NewMaskStructItem(2, nil),
	}, true)
	read := NewNamedTableReadRel(NewBaseReadRel(
		NewRelCommon(nil, []int32{1, 0}, nil), schema, nil, nil, projection, nil,
	), []string{"t"}, nil)
	assert.Equal(t, []types.Type{schema.Struct.Types[2], schema.Struct.Types[0]}, read.RecordType().Types())
	assert.Equal(t, schema, read.BaseSchema())
}

func TestReadProjectionSetAndClear(t *testing.T) {
	schema := projectionTestSchema()
	read := NewBuilderDefault().NamedScan([]string{"t"}, schema)
	assert.Equal(t, schema.Struct.Types, read.RecordType().Types())
	for _, maintain := range []bool{false, true} {
		read.SetProjection(expr.NewMaskExpression(expr.MaskStructSelect{expr.NewMaskStructItem(2, nil)}, maintain))
		assert.Equal(t, []types.Type{schema.Struct.Types[2]}, read.RecordType().Types())
		read.SetProjection(expr.NewMaskExpression(nil, maintain))
		assert.Empty(t, read.RecordType().Types())
	}
	read.SetProjection(nil)
	assert.Equal(t, schema.Struct.Types, read.RecordType().Types())
}

func TestReadProjectionDoesNotShareFieldStorage(t *testing.T) {
	schema := projectionTestSchema()
	read := NewBuilderDefault().NamedScan([]string{"t"}, schema)
	read.SetProjection(expr.NewMaskExpression(expr.MaskStructSelect{expr.NewMaskStructItem(0, nil)}, true))
	read.RecordType().Types()[0] = schema.Struct.Types[1]
	assert.Equal(t, &types.Int64Type{Nullability: types.NullabilityRequired}, schema.Struct.Types[0])
	assert.Equal(t, []types.Type{schema.Struct.Types[0]}, read.RecordType().Types())
}
