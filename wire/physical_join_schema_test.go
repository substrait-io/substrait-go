// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	substraitgo "github.com/substrait-io/substrait-go/v9"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/plan"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func keyRef(t *testing.T, rel plan.Rel, idx int32) *expr.FieldReference {
	t.Helper()
	base := rel.RecordType()
	ref, err := expr.NewRootFieldRef(expr.NewStructFieldRef(idx), &base)
	require.NoError(t, err)
	return ref
}

func TestPhysicalJoinOutputSchema(t *testing.T) {
	i64 := &types.Int64Type{Nullability: types.NullabilityRequired}
	i64Null := &types.Int64Type{Nullability: types.NullabilityNullable}
	str := &types.StringType{Nullability: types.NullabilityRequired}
	strNull := &types.StringType{Nullability: types.NullabilityNullable}
	mark := &types.BooleanType{Nullability: types.NullabilityNullable}
	b := plan.NewBuilderDefault()
	left := b.NamedScan([]string{"left"}, types.NamedStruct{
		Names: []string{"l0", "l1"}, Struct: types.StructType{
			Nullability: types.NullabilityRequired, Types: []types.Type{i64, strNull},
		},
	})
	right := b.NamedScan([]string{"right"}, types.NamedStruct{
		Names: []string{"r0", "r1"}, Struct: types.StructType{
			Nullability: types.NullabilityRequired, Types: []types.Type{i64Null, str},
		},
	})
	keys := comparisonJoinKeysToProto([]*plan.ComparisonJoinKey{
		plan.NewEqualityJoinKey(keyRef(t, left, 0), keyRef(t, right, 0)),
	})

	// Name both protobuf enums explicitly: their semi/anti/single numbering is
	// different from JoinRel's, so a cast to JoinType gives the wrong schema.
	tests := []struct {
		name  string
		hash  proto.HashJoinRel_JoinType
		merge proto.MergeJoinRel_JoinType
		want  []types.Type
	}{
		{"inner", proto.HashJoinRel_JOIN_TYPE_INNER, proto.MergeJoinRel_JOIN_TYPE_INNER, []types.Type{i64, strNull, i64Null, str}},
		{"outer", proto.HashJoinRel_JOIN_TYPE_OUTER, proto.MergeJoinRel_JOIN_TYPE_OUTER, []types.Type{i64Null, strNull, i64Null, strNull}},
		{"left", proto.HashJoinRel_JOIN_TYPE_LEFT, proto.MergeJoinRel_JOIN_TYPE_LEFT, []types.Type{i64, strNull, i64Null, strNull}},
		{"right", proto.HashJoinRel_JOIN_TYPE_RIGHT, proto.MergeJoinRel_JOIN_TYPE_RIGHT, []types.Type{i64Null, strNull, i64Null, str}},
		{"left semi", proto.HashJoinRel_JOIN_TYPE_LEFT_SEMI, proto.MergeJoinRel_JOIN_TYPE_LEFT_SEMI, []types.Type{i64, strNull}},
		{"right semi", proto.HashJoinRel_JOIN_TYPE_RIGHT_SEMI, proto.MergeJoinRel_JOIN_TYPE_RIGHT_SEMI, []types.Type{i64Null, str}},
		{"left anti", proto.HashJoinRel_JOIN_TYPE_LEFT_ANTI, proto.MergeJoinRel_JOIN_TYPE_LEFT_ANTI, []types.Type{i64, strNull}},
		{"right anti", proto.HashJoinRel_JOIN_TYPE_RIGHT_ANTI, proto.MergeJoinRel_JOIN_TYPE_RIGHT_ANTI, []types.Type{i64Null, str}},
		{"left single", proto.HashJoinRel_JOIN_TYPE_LEFT_SINGLE, proto.MergeJoinRel_JOIN_TYPE_LEFT_SINGLE, []types.Type{i64, strNull, i64Null, strNull}},
		{"right single", proto.HashJoinRel_JOIN_TYPE_RIGHT_SINGLE, proto.MergeJoinRel_JOIN_TYPE_RIGHT_SINGLE, []types.Type{i64Null, strNull, i64Null, str}},
		{"left mark", proto.HashJoinRel_JOIN_TYPE_LEFT_MARK, proto.MergeJoinRel_JOIN_TYPE_LEFT_MARK, []types.Type{i64, strNull, mark}},
		{"right mark", proto.HashJoinRel_JOIN_TYPE_RIGHT_MARK, proto.MergeJoinRel_JOIN_TYPE_RIGHT_MARK, []types.Type{i64Null, str, mark}},
	}

	for _, tc := range tests {
		t.Run("hash/"+tc.name, func(t *testing.T) {
			input := &proto.Rel{RelType: &proto.Rel_HashJoin{HashJoin: &proto.HashJoinRel{
				Common: &proto.RelCommon{}, Left: RelToProto(left), Right: RelToProto(right), Keys: keys, Type: tc.hash,
			}}}
			rel, err := RelFromProto(input, joinTestRegistry())
			require.NoError(t, err)
			assert.Equal(t, tc.want, rel.RecordType().Types())
		})
		t.Run("merge/"+tc.name, func(t *testing.T) {
			input := &proto.Rel{RelType: &proto.Rel_MergeJoin{MergeJoin: &proto.MergeJoinRel{
				Common: &proto.RelCommon{}, Left: RelToProto(left), Right: RelToProto(right), Keys: keys, Type: tc.merge,
			}}}
			rel, err := RelFromProto(input, joinTestRegistry())
			require.NoError(t, err)
			assert.Equal(t, tc.want, rel.RecordType().Types())
		})
	}
}

func postJoinFilterField(index int32) *proto.Expression {
	return &proto.Expression{RexType: &proto.Expression_Selection{Selection: fieldRefProto(index)}}
}

// joinInputWithRightBoolean builds a read rel with three required columns whose
// last column is a boolean, so a post-filter can reference it.
func joinInputWithRightBoolean(name string) *plan.NamedTableReadRel {
	schema := types.NamedStruct{
		Names: []string{"x", "y", "z"},
		Struct: types.StructType{
			Nullability: types.NullabilityRequired,
			Types:       []types.Type{&types.Int64Type{}, &types.Int64Type{}, &types.BooleanType{Nullability: types.NullabilityRequired}},
		},
	}
	base := plan.NewBaseReadRel(plan.RelCommon{}, schema, nil, nil, nil, nil)
	return plan.NewNamedTableReadRel(base, []string{name}, nil)
}

func TestPhysicalJoinPostFilterUsesDirectOutput(t *testing.T) {
	left, right := createJoinInput("left"), joinInputWithRightBoolean("right")
	keys := comparisonJoinKeysToProto([]*plan.ComparisonJoinKey{
		plan.NewEqualityJoinKey(keyRef(t, left, 0), keyRef(t, right, 0)),
	})
	// Emit keeps only left field 1, but the filter reads right field 5.
	common := &proto.RelCommon{EmitKind: &proto.RelCommon_Emit_{
		Emit: &proto.RelCommon_Emit{OutputMapping: []int32{1}},
	}}
	const rightBooleanField = 5
	predicate := postJoinFilterField(rightBooleanField)

	t.Run("hash", func(t *testing.T) {
		input := &proto.Rel{RelType: &proto.Rel_HashJoin{HashJoin: &proto.HashJoinRel{
			Common: common, Left: RelToProto(left), Right: RelToProto(right), Keys: keys,
			Type: proto.HashJoinRel_JOIN_TYPE_LEFT, PostJoinFilter: predicate,
		}}}
		rel, err := RelFromProto(input, joinTestRegistry())
		require.NoError(t, err)
		filter := rel.(*plan.HashJoinRel).PostJoinFilter()
		assert.Equal(t, &types.BooleanType{Nullability: types.NullabilityNullable}, filter.GetType())
	})
	t.Run("merge", func(t *testing.T) {
		input := &proto.Rel{RelType: &proto.Rel_MergeJoin{MergeJoin: &proto.MergeJoinRel{
			Common: common, Left: RelToProto(left), Right: RelToProto(right), Keys: keys,
			Type: proto.MergeJoinRel_JOIN_TYPE_LEFT, PostJoinFilter: predicate,
		}}}
		rel, err := RelFromProto(input, joinTestRegistry())
		require.NoError(t, err)
		filter := rel.(*plan.MergeJoinRel).PostJoinFilter()
		assert.Equal(t, &types.BooleanType{Nullability: types.NullabilityNullable}, filter.GetType())
	})
}

func TestPhysicalJoinPostFilterRejectsDiscardedField(t *testing.T) {
	left, right := createJoinInput("left"), createJoinInput("right")
	keys := comparisonJoinKeysToProto([]*plan.ComparisonJoinKey{
		plan.NewEqualityJoinKey(keyRef(t, left, 0), keyRef(t, right, 0)),
	})
	const discardedRightField = 5
	predicate := postJoinFilterField(discardedRightField)
	common := &proto.RelCommon{EmitKind: &proto.RelCommon_Emit_{
		Emit: &proto.RelCommon_Emit{OutputMapping: []int32{1}},
	}}

	t.Run("hash", func(t *testing.T) {
		input := &proto.Rel{RelType: &proto.Rel_HashJoin{HashJoin: &proto.HashJoinRel{
			Common: common, Left: RelToProto(left), Right: RelToProto(right), Keys: keys,
			Type: proto.HashJoinRel_JOIN_TYPE_LEFT_SEMI, PostJoinFilter: predicate,
		}}}
		_, err := RelFromProto(input, joinTestRegistry())
		require.ErrorContains(t, err, "post join filter")
		require.ErrorIs(t, err, substraitgo.ErrInvalidType)
	})
	t.Run("merge", func(t *testing.T) {
		input := &proto.Rel{RelType: &proto.Rel_MergeJoin{MergeJoin: &proto.MergeJoinRel{
			Common: common, Left: RelToProto(left), Right: RelToProto(right), Keys: keys,
			Type: proto.MergeJoinRel_JOIN_TYPE_LEFT_SEMI, PostJoinFilter: predicate,
		}}}
		_, err := RelFromProto(input, joinTestRegistry())
		require.ErrorContains(t, err, "post join filter")
		require.ErrorIs(t, err, substraitgo.ErrInvalidType)
	})
}
