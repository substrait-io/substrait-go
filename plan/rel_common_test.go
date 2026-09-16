// SPDX-License-Identifier: Apache-2.0

package plan

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	extensionspb "github.com/substrait-io/substrait-protobuf/go/substraitpb/extensions"
	"google.golang.org/protobuf/types/known/anypb"
)

func TestRelCommonAbsentIsDirect(t *testing.T) {
	input := *types.NewRecordTypeFromTypes([]types.Type{
		&types.Int64Type{Nullability: types.NullabilityRequired},
		&types.StringType{Nullability: types.NullabilityNullable},
	})
	for _, tc := range []struct {
		name string
		wire *proto.RelCommon
		want []types.Type
	}{
		{"absent", nil, input.Types()},
		{"empty", &proto.RelCommon{}, input.Types()},
		{"direct", &proto.RelCommon{EmitKind: &proto.RelCommon_Direct_{Direct: &proto.RelCommon_Direct{}}}, input.Types()},
		{"emit", &proto.RelCommon{EmitKind: &proto.RelCommon_Emit_{Emit: &proto.RelCommon_Emit{OutputMapping: []int32{1, 0, 1}}}}, []types.Type{input.Types()[1], input.Types()[0], input.Types()[1]}},
		{"empty emit", &proto.RelCommon{EmitKind: &proto.RelCommon_Emit_{Emit: &proto.RelCommon_Emit{OutputMapping: []int32{}}}}, []types.Type{}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var common RelCommon
			// Decoding an absent common must also clear previously decoded metadata.
			common.fromProtoCommon(&proto.RelCommon{
				Hint:              &proto.RelCommon_Hint{Alias: "old"},
				AdvancedExtension: &extensionspb.AdvancedExtension{Enhancement: &anypb.Any{TypeUrl: "test"}},
				EmitKind:          &proto.RelCommon_Emit_{Emit: &proto.RelCommon_Emit{OutputMapping: []int32{1}}},
			})
			require.NotPanics(t, func() { common.fromProtoCommon(tc.wire) })
			assert.Equal(t, tc.want, common.remap(input).Types())
			assert.Nil(t, common.toProto().Hint)
			assert.Nil(t, common.toProto().AdvancedExtension)
		})
	}

	t.Run("metadata", func(t *testing.T) {
		wire := &proto.RelCommon{
			Hint:              &proto.RelCommon_Hint{Alias: "input", OutputNames: []string{"a", "b"}},
			AdvancedExtension: &extensionspb.AdvancedExtension{Enhancement: &anypb.Any{TypeUrl: "test", Value: []byte{1}}},
		}
		var common RelCommon
		common.fromProtoCommon(wire)
		assert.Equal(t, wire.Hint.Alias, common.toProto().Hint.Alias)
		assert.Equal(t, wire.Hint.OutputNames, common.toProto().Hint.OutputNames)
		assert.Equal(t, wire.AdvancedExtension.Enhancement.TypeUrl, common.toProto().AdvancedExtension.Enhancement.TypeUrl)
		assert.Equal(t, wire.AdvancedExtension.Enhancement.Value, common.toProto().AdvancedExtension.Enhancement.Value)
	})
}

func TestRelationsWithoutCommon(t *testing.T) {
	left, right := createJoinInput("left"), createJoinInput("right")
	left.baseSchema.Struct.Types = []types.Type{
		&types.Int64Type{Nullability: types.NullabilityRequired},
		&types.StringType{Nullability: types.NullabilityNullable},
	}
	left.baseSchema.Names = []string{"id", "label"}
	right.baseSchema = left.baseSchema
	keys := comparisonJoinKeysToProto([]*ComparisonJoinKey{NewEqualityJoinKey(keyRef(t, left, 0), keyRef(t, right, 0))})
	for _, tc := range []struct {
		name    string
		makeRel func(*proto.RelCommon) *proto.Rel
		want    []types.Type
	}{
		{"join", func(c *proto.RelCommon) *proto.Rel {
			return &proto.Rel{RelType: &proto.Rel_Join{Join: &proto.JoinRel{Common: c, Left: left.ToProto(), Right: right.ToProto(), Type: proto.JoinRel_JOIN_TYPE_INNER, Expression: expr.NewPrimitiveLiteral(true, false).ToProto()}}}
		}, left.RecordType().Concat(right.RecordType()).Types()},
		{"cross", func(c *proto.RelCommon) *proto.Rel {
			return &proto.Rel{RelType: &proto.Rel_Cross{Cross: &proto.CrossRel{Common: c, Left: left.ToProto(), Right: right.ToProto()}}}
		}, left.RecordType().Concat(right.RecordType()).Types()},
		{"sort", func(c *proto.RelCommon) *proto.Rel {
			return &proto.Rel{RelType: &proto.Rel_Sort{Sort: &proto.SortRel{Common: c, Input: left.ToProto(), Sorts: []*proto.SortField{{Expr: keyRef(t, left, 0).ToProto(), SortKind: &proto.SortField_Direction{Direction: proto.SortField_SORT_DIRECTION_ASC_NULLS_FIRST}}}}}}
		}, left.RecordType().Types()},
		{"aggregate", func(c *proto.RelCommon) *proto.Rel {
			return &proto.Rel{RelType: &proto.Rel_Aggregate{Aggregate: &proto.AggregateRel{Common: c, Input: left.ToProto(), GroupingExpressions: []*proto.Expression{keyRef(t, left, 0).ToProto()}, Groupings: []*proto.AggregateRel_Grouping{{ExpressionReferences: []uint32{0}}}}}}
		}, left.RecordType().Types()[:1]},
		{"set", func(c *proto.RelCommon) *proto.Rel {
			return &proto.Rel{RelType: &proto.Rel_Set{Set: &proto.SetRel{Common: c, Inputs: []*proto.Rel{left.ToProto(), right.ToProto()}, Op: proto.SetRel_SET_OP_UNION_ALL}}}
		}, left.RecordType().Types()},
		{"hash", func(c *proto.RelCommon) *proto.Rel {
			return &proto.Rel{RelType: &proto.Rel_HashJoin{HashJoin: &proto.HashJoinRel{Common: c, Left: left.ToProto(), Right: right.ToProto(), Keys: keys, Type: proto.HashJoinRel_JOIN_TYPE_INNER}}}
		}, left.RecordType().Concat(right.RecordType()).Types()},
		{"merge", func(c *proto.RelCommon) *proto.Rel {
			return &proto.Rel{RelType: &proto.Rel_MergeJoin{MergeJoin: &proto.MergeJoinRel{Common: c, Left: left.ToProto(), Right: right.ToProto(), Keys: keys, Type: proto.MergeJoinRel_JOIN_TYPE_INNER}}}
		}, left.RecordType().Concat(right.RecordType()).Types()},
	} {
		for _, mode := range []string{"absent", "empty", "direct", "emit"} {
			t.Run(tc.name+"/"+mode, func(t *testing.T) {
				var common *proto.RelCommon
				want := tc.want
				switch mode {
				case "empty":
					common = &proto.RelCommon{}
				case "direct":
					common = &proto.RelCommon{EmitKind: &proto.RelCommon_Direct_{Direct: &proto.RelCommon_Direct{}}}
				case "emit":
					last := int32(len(want) - 1)
					common = &proto.RelCommon{EmitKind: &proto.RelCommon_Emit_{Emit: &proto.RelCommon_Emit{OutputMapping: []int32{last, 0, last}}}}
					want = []types.Type{want[last], want[0], want[last]}
				}
				var rel Rel
				var err error
				require.NotPanics(t, func() { rel, err = RelFromProto(tc.makeRel(common), joinTestRegistry()) })
				require.NoError(t, err)
				assert.Equal(t, want, rel.RecordType().Types())
				again, err := RelFromProto(rel.ToProto(), joinTestRegistry())
				require.NoError(t, err)
				assert.Equal(t, want, again.RecordType().Types())
			})
		}
	}
}
