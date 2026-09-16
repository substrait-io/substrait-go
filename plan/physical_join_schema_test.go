// SPDX-License-Identifier: Apache-2.0

package plan

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func TestPhysicalJoinOutputSchema(t *testing.T) {
	i64 := &types.Int64Type{Nullability: types.NullabilityRequired}
	i64Null := &types.Int64Type{Nullability: types.NullabilityNullable}
	str := &types.StringType{Nullability: types.NullabilityRequired}
	strNull := &types.StringType{Nullability: types.NullabilityNullable}
	mark := &types.BooleanType{Nullability: types.NullabilityNullable}
	left := &NamedTableReadRel{names: []string{"left"}, baseReadRel: baseReadRel{
		baseSchema: types.NamedStruct{Names: []string{"l0", "l1"}, Struct: types.StructType{
			Nullability: types.NullabilityRequired, Types: []types.Type{i64, strNull},
		}},
	}}
	right := &NamedTableReadRel{names: []string{"right"}, baseReadRel: baseReadRel{
		baseSchema: types.NamedStruct{Names: []string{"r0", "r1"}, Struct: types.StructType{
			Nullability: types.NullabilityRequired, Types: []types.Type{i64Null, str},
		}},
	}}
	keys := comparisonJoinKeysToProto([]*ComparisonJoinKey{
		NewEqualityJoinKey(keyRef(t, left, 0), keyRef(t, right, 0)),
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
		for _, kind := range []string{"hash", "merge"} {
			t.Run(kind+"/"+tc.name, func(t *testing.T) {
				common := &proto.RelCommon{}
				var input *proto.Rel
				if kind == "hash" {
					input = &proto.Rel{RelType: &proto.Rel_HashJoin{HashJoin: &proto.HashJoinRel{
						Common: common, Left: left.ToProto(), Right: right.ToProto(), Keys: keys, Type: tc.hash,
					}}}
				} else {
					input = &proto.Rel{RelType: &proto.Rel_MergeJoin{MergeJoin: &proto.MergeJoinRel{
						Common: common, Left: left.ToProto(), Right: right.ToProto(), Keys: keys, Type: tc.merge,
					}}}
				}
				rel, err := RelFromProto(input, joinTestRegistry())
				require.NoError(t, err)
				assert.Equal(t, tc.want, rel.RecordType().Types())
			})
		}
	}
}

func TestNullableRecordTypePreservesInputAndNestedFields(t *testing.T) {
	child := &types.Int64Type{Nullability: types.NullabilityRequired}
	nested := &types.StructType{Nullability: types.NullabilityRequired, Types: []types.Type{child}}
	input := types.NewRecordTypeFromTypes([]types.Type{nested, child})
	want := []types.Type{
		&types.StructType{Nullability: types.NullabilityNullable, Types: []types.Type{child}},
		&types.Int64Type{Nullability: types.NullabilityNullable},
	}

	output := nullableRecordType(*input)
	assert.Equal(t, want, output.Types())
	assert.Equal(t, []types.Type{nested, child}, input.Types())
	assert.Equal(t, types.NullabilityRequired, nested.Nullability)
	assert.Equal(t, types.NullabilityRequired, child.Nullability)
}

func TestPhysicalJoinPostFilterUsesDirectOutput(t *testing.T) {
	left, right := createJoinInput("left"), createJoinInput("right")
	left.baseSchema.Struct.Types[2] = &types.BooleanType{Nullability: types.NullabilityNullable}
	right.baseSchema.Struct.Types[2] = &types.BooleanType{Nullability: types.NullabilityRequired}
	keys := comparisonJoinKeysToProto([]*ComparisonJoinKey{
		NewEqualityJoinKey(keyRef(t, left, 0), keyRef(t, right, 0)),
	})
	joined := &fakeRel{outputType: left.RecordType().Concat(right.RecordType())}
	for _, tc := range []struct {
		name      string
		join      HashMergeJoinType
		field     int32
		want      types.Nullability
		wantError bool
	}{
		{"left mark", HashMergeLeftMark, 3, types.NullabilityNullable, false},
		{"right mark", HashMergeRightMark, 3, types.NullabilityNullable, false},
		{"right semi", HashMergeRightSemi, 2, types.NullabilityRequired, false},
		{"left outer", HashMergeLeft, 5, types.NullabilityNullable, false},
		{"inner", HashMergeInner, 5, types.NullabilityRequired, false},
		{"left semi discarded field", HashMergeLeftSemi, 5, types.NullabilityUnspecified, true},
	} {
		for _, kind := range []string{"hash", "merge"} {
			t.Run(kind+"/"+tc.name, func(t *testing.T) {
				// The filter's field is absent from the emitted output.
				common := &proto.RelCommon{EmitKind: &proto.RelCommon_Emit_{
					Emit: &proto.RelCommon_Emit{OutputMapping: []int32{1}},
				}}
				predicate := keyRef(t, joined, tc.field).ToProto()
				var input *proto.Rel
				if kind == "hash" {
					input = &proto.Rel{RelType: &proto.Rel_HashJoin{HashJoin: &proto.HashJoinRel{
						Common: common, Left: left.ToProto(), Right: right.ToProto(), Keys: keys,
						Type: proto.HashJoinRel_JoinType(tc.join), PostJoinFilter: predicate,
					}}}
				} else {
					input = &proto.Rel{RelType: &proto.Rel_MergeJoin{MergeJoin: &proto.MergeJoinRel{
						Common: common, Left: left.ToProto(), Right: right.ToProto(), Keys: keys,
						Type: proto.MergeJoinRel_JoinType(tc.join), PostJoinFilter: predicate,
					}}}
				}
				rel, err := RelFromProto(input, joinTestRegistry())
				if tc.wantError {
					require.ErrorContains(t, err, "post join filter")
					return
				}
				require.NoError(t, err)
				filter := rel.(interface{ PostJoinFilter() expr.Expression }).PostJoinFilter()
				assert.Equal(t, &types.BooleanType{Nullability: tc.want}, filter.GetType())
			})
		}
	}
}
