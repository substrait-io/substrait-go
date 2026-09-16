// SPDX-License-Identifier: Apache-2.0

package plan

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
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
	keys := []*proto.ComparisonJoinKey{{
		Left: keyRef(t, left, 0).ToProtoFieldRef(), Right: keyRef(t, right, 0).ToProtoFieldRef(),
		Comparison: &proto.ComparisonJoinKey_ComparisonType{InnerType: &proto.ComparisonJoinKey_ComparisonType_Simple{
			Simple: proto.ComparisonJoinKey_SIMPLE_COMPARISON_TYPE_EQ,
		}},
	}}

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
			for _, emit := range []bool{false, true} {
				t.Run(fmt.Sprintf("%s/%s/emit=%t", kind, tc.name, emit), func(t *testing.T) {
					common := &proto.RelCommon{EmitKind: &proto.RelCommon_Direct_{Direct: &proto.RelCommon_Direct{}}}
					want := tc.want
					if emit {
						// Select, reorder and duplicate fields of the derived output,
						// including the mark column and the retained side of semi joins.
						last := int32(len(want) - 1)
						common.EmitKind = &proto.RelCommon_Emit_{Emit: &proto.RelCommon_Emit{OutputMapping: []int32{last, 0, last}}}
						want = []types.Type{want[last], want[0], want[last]}
					}
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
					names := make([]string, len(want))
					for i := range names {
						names[i] = fmt.Sprintf("c%d", i)
					}
					p, err := FromProto(&proto.Plan{Relations: []*proto.PlanRel{{RelType: &proto.PlanRel_Root{Root: &proto.RelRoot{
						Input: input, Names: names,
					}}}}}, extensions.GetDefaultCollectionWithNoError())
					require.NoError(t, err)
					assert.Equal(t, want, p.GetRoots()[0].RecordType().Struct.Types)
					roundTrip, err := RelFromProto(p.GetRoots()[0].Input().ToProto(), joinTestRegistry())
					require.NoError(t, err)
					assert.Equal(t, want, roundTrip.RecordType().Types())
				})
			}
		}
	}
}

func TestPhysicalJoinNestedOutputSchema(t *testing.T) {
	child := &types.Int64Type{Nullability: types.NullabilityRequired}
	nested := &types.StructType{Nullability: types.NullabilityRequired, Types: []types.Type{child}}
	// Spare capacity must not let an output append overwrite another input field.
	storage := []types.Type{nested, child, child}
	left := &fakeRel{outputType: *types.NewRecordTypeFromTypes(storage[:1])}
	right := &fakeRel{outputType: *types.NewRecordTypeFromTypes([]types.Type{child})}
	want := []types.Type{
		&types.StructType{Nullability: types.NullabilityNullable, Types: []types.Type{child}},
		&types.Int64Type{Nullability: types.NullabilityNullable},
	}
	for _, tc := range []struct {
		name string
		rel  Rel
	}{
		{"hash", &HashJoinRel{left: left, right: right, joinType: HashMergeOuter}},
		{"merge", &MergeJoinRel{left: left, right: right, joinType: HashMergeOuter}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, want, tc.rel.RecordType().Types())
			assert.Equal(t, types.NullabilityRequired, nested.Nullability)
			assert.Equal(t, types.NullabilityRequired, child.Nullability)
			assert.Equal(t, []types.Type{nested, child, child}, storage)
		})
	}
}

func TestPhysicalJoinPostFilterUsesInputSchema(t *testing.T) {
	left := createJoinInput("left")
	right := createJoinInput("right")
	right.baseSchema.Struct.Types[2] = &types.BooleanType{Nullability: types.NullabilityRequired}
	keys := comparisonJoinKeysToProto([]*ComparisonJoinKey{
		NewEqualityJoinKey(keyRef(t, left, 0), keyRef(t, right, 0)),
	})
	joined := left.RecordType().Concat(right.RecordType())
	predicate, err := expr.NewRootFieldRef(expr.NewStructFieldRef(5), &joined)
	require.NoError(t, err)
	for _, kind := range []string{"hash", "merge"} {
		for _, emit := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/emit=%t", kind, emit), func(t *testing.T) {
				common := &proto.RelCommon{}
				if emit {
					common.EmitKind = &proto.RelCommon_Emit_{Emit: &proto.RelCommon_Emit{OutputMapping: []int32{1}}}
				}
				var input *proto.Rel
				if kind == "hash" {
					input = &proto.Rel{RelType: &proto.Rel_HashJoin{HashJoin: &proto.HashJoinRel{
						Common: common, Left: left.ToProto(), Right: right.ToProto(), Keys: keys,
						Type: proto.HashJoinRel_JOIN_TYPE_LEFT_SEMI, PostJoinFilter: predicate.ToProto(),
					}}}
				} else {
					input = &proto.Rel{RelType: &proto.Rel_MergeJoin{MergeJoin: &proto.MergeJoinRel{
						Common: common, Left: left.ToProto(), Right: right.ToProto(), Keys: keys,
						Type: proto.MergeJoinRel_JOIN_TYPE_LEFT_SEMI, PostJoinFilter: predicate.ToProto(),
					}}}
				}
				rel, err := RelFromProto(input, joinTestRegistry())
				require.NoError(t, err)
				var filter expr.Expression
				if kind == "hash" {
					filter = rel.(*HashJoinRel).PostJoinFilter()
				} else {
					filter = rel.(*MergeJoinRel).PostJoinFilter()
				}
				assert.True(t, predicate.Equals(filter))
				assert.Equal(t, &types.BooleanType{Nullability: types.NullabilityRequired}, filter.GetType())
			})
		}
	}
}
