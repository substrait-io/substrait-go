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
		kind proto.JoinRel_JoinType
		side Rel
	}{
		{proto.JoinRel_JOIN_TYPE_LEFT_MARK, left},
		{proto.JoinRel_JOIN_TYPE_RIGHT_MARK, right},
	} {
		for _, emit := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/emit=%t", tc.kind, emit), func(t *testing.T) {
				want := append([]types.Type{}, tc.side.RecordType().Types()...)
				want = append(want, mark)
				common := &proto.RelCommon{}
				if emit {
					common.EmitKind = &proto.RelCommon_Emit_{Emit: &proto.RelCommon_Emit{OutputMapping: []int32{3, 0, 3}}}
					want = []types.Type{mark, want[0], mark}
				}
				input := &proto.Rel{RelType: &proto.Rel_Join{Join: &proto.JoinRel{
					Common: common, Left: left.ToProto(), Right: right.ToProto(), Type: tc.kind,
					Expression: expr.NewPrimitiveLiteral(true, false).ToProto(),
				}}}
				names := make([]string, len(want))
				for i := range names {
					names[i] = fmt.Sprintf("c%d", i)
				}
				root := &proto.RelRoot{Input: input, Names: names}
				wire := &proto.Plan{Relations: []*proto.PlanRel{{RelType: &proto.PlanRel_Root{Root: root}}}}
				p, err := FromProto(wire, extensions.GetDefaultCollectionWithNoError())
				require.NoError(t, err)
				assert.Equal(t, want, p.GetRoots()[0].RecordType().Struct.Types)
				roundTrip, err := RelFromProto(p.GetRoots()[0].Input().ToProto(), joinTestRegistry())
				require.NoError(t, err)
				assert.Equal(t, want, roundTrip.RecordType().Types())
				root.Names = nil
				_, err = FromProto(wire, extensions.GetDefaultCollectionWithNoError())
				require.ErrorContains(t, err, "root relation has 0 output name(s)")
			})
		}
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
			assert.Equal(t, want, join.RecordType().Types())
			assert.Equal(t, []types.Type{nested, child, child}, storage)
			assert.Equal(t, types.NullabilityRequired, nested.Nullability)
			assert.Equal(t, types.NullabilityRequired, child.Nullability)
		})
	}
}

func TestLogicalJoinPostFilterUsesDirectOutput(t *testing.T) {
	left, right := createJoinInput("left"), createJoinInput("right")
	left.baseSchema.Struct.Types[0] = &types.BooleanType{Nullability: types.NullabilityNullable}
	right.baseSchema.Struct.Types[0] = &types.BooleanType{Nullability: types.NullabilityRequired}
	joined := &fakeRel{outputType: left.RecordType().Concat(right.RecordType())}
	condition := keyRef(t, joined, 3)
	for _, tc := range []struct {
		kind  proto.JoinRel_JoinType
		field int32
		want  types.Nullability
	}{
		{proto.JoinRel_JOIN_TYPE_LEFT_MARK, 3, types.NullabilityNullable},
		{proto.JoinRel_JOIN_TYPE_RIGHT_MARK, 3, types.NullabilityNullable},
		{proto.JoinRel_JOIN_TYPE_RIGHT_SEMI, 0, types.NullabilityRequired},
		{proto.JoinRel_JOIN_TYPE_OUTER, 3, types.NullabilityNullable},
		{proto.JoinRel_JOIN_TYPE_INNER, 3, types.NullabilityRequired},
	} {
		for _, emit := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/emit=%t", tc.kind, emit), func(t *testing.T) {
				common := &proto.RelCommon{}
				if emit {
					common.EmitKind = &proto.RelCommon_Emit_{Emit: &proto.RelCommon_Emit{OutputMapping: []int32{1}}}
				}
				// A field reference has no encoded type: the decoder must bind it
				// against the direct output, even if emit drops that field.
				predicate := keyRef(t, joined, tc.field).ToProto()
				wire := &proto.Rel{RelType: &proto.Rel_Join{Join: &proto.JoinRel{
					Common: common, Left: left.ToProto(), Right: right.ToProto(), Type: tc.kind,
					Expression: condition.ToProto(), PostJoinFilter: predicate,
				}}}
				rel, err := RelFromProto(wire, joinTestRegistry())
				require.NoError(t, err)
				assert.True(t, condition.Equals(rel.(*JoinRel).Expr()))
				assert.Equal(t, &types.BooleanType{Nullability: tc.want}, rel.(*JoinRel).PostJoinFilter().GetType())
				again, err := RelFromProto(rel.ToProto(), joinTestRegistry())
				require.NoError(t, err)
				assert.True(t, rel.(*JoinRel).PostJoinFilter().Equals(again.(*JoinRel).PostJoinFilter()))
			})
		}
	}
}
