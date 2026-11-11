// SPDX-License-Identifier: Apache-2.0

package plan

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/types"
)

func TestMarkJoinOutputSchema(t *testing.T) {
	left, right := createJoinInput("left"), createJoinInput("right")
	left.baseSchema.Struct.Types = []types.Type{
		&types.Int64Type{Nullability: types.NullabilityRequired},
		&types.Int64Type{Nullability: types.NullabilityNullable},
		&types.StringType{Nullability: types.NullabilityRequired},
	}
	right.baseSchema.Struct.Types = append([]types.Type{}, left.baseSchema.Struct.Types...)
	right.baseSchema.Struct.Types[0] = &types.StringType{Nullability: types.NullabilityNullable}
	mark := &types.BooleanType{Nullability: types.NullabilityNullable}
	for _, tc := range []struct {
		kind JoinType
		side Rel
	}{
		{JoinTypeLeftMark, left},
		{JoinTypeRightMark, right},
	} {
		t.Run(tc.kind.String(), func(t *testing.T) {
			want := append([]types.Type{}, tc.side.RecordType().Types()...)
			want = append(want, mark)
			rel := NewJoinRel(left, right, tc.kind, expr.NewPrimitiveLiteral(true, false), nil, RelCommon{}, nil)
			assert.Equal(t, want, rel.RecordType().Types())

			t.Run("emit", func(t *testing.T) {
				mapping := []int32{3, 0, 1}
				emitted := NewJoinRel(left, right, tc.kind, expr.NewPrimitiveLiteral(true, false), nil,
					NewRelCommon(nil, mapping, nil), nil)
				assert.Equal(t, []types.Type{mark, want[0], want[1]}, emitted.RecordType().Types())
			})
		})
	}
}

func TestMarkJoinPreservesInputStorage(t *testing.T) {
	child := &types.Int64Type{Nullability: types.NullabilityRequired}
	nested := &types.StructType{Nullability: types.NullabilityRequired, Types: []types.Type{child}}
	for _, kind := range []JoinType{JoinTypeLeftMark, JoinTypeRightMark} {
		t.Run(kind.String(), func(t *testing.T) {
			storage := []types.Type{nested, child, child}
			input := &fakeRel{outputType: *types.NewRecordTypeFromTypes(storage[:1])}
			join := &JoinRel{left: input, right: input, joinType: kind}
			want := []types.Type{nested, &types.BooleanType{Nullability: types.NullabilityNullable}}
			assert.Equal(t, want, join.RecordType().Types())
			assert.Equal(t, []types.Type{nested, child, child}, storage)
			assert.Equal(t, types.NullabilityRequired, nested.Nullability)
			assert.Equal(t, types.NullabilityRequired, child.Nullability)
		})
	}
}
