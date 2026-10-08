// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/plan"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	gproto "google.golang.org/protobuf/proto"
)

func TestLogicalJoinPostFilterUsesDirectOutput(t *testing.T) {
	schema := types.NamedStruct{
		Names: []string{"flag", "number", "text"},
		Struct: types.StructType{Nullability: types.NullabilityRequired, Types: []types.Type{
			&types.BooleanType{Nullability: types.NullabilityNullable},
			&types.Int64Type{Nullability: types.NullabilityRequired},
			&types.StringType{Nullability: types.NullabilityRequired},
		}},
	}
	b := plan.NewBuilderDefault()
	left := b.NamedScan([]string{"left"}, schema)
	schema.Struct.Types = append([]types.Type{}, schema.Struct.Types...)
	schema.Struct.Types[0] = &types.BooleanType{Nullability: types.NullabilityRequired}
	right := b.NamedScan([]string{"right"}, schema)
	joined := plan.NewCrossRel(left, right, plan.RelCommon{}, nil)
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
		t.Run(tc.kind.String(), func(t *testing.T) {
			common := &proto.RelCommon{EmitKind: &proto.RelCommon_Emit_{
				Emit: &proto.RelCommon_Emit{OutputMapping: []int32{1}},
			}}
			// A field reference has no encoded type: the decoder must bind it
			// against the direct output, even if emit drops that field.
			predicate := ExprToProto(keyRef(t, joined, tc.field))
			wire := &proto.Rel{RelType: &proto.Rel_Join{Join: &proto.JoinRel{
				Common: common, Left: RelToProto(left), Right: RelToProto(right), Type: tc.kind,
				Expression: ExprToProto(condition), PostJoinFilter: predicate,
			}}}
			rel, err := RelFromProto(wire, joinTestRegistry())
			require.NoError(t, err)
			assert.True(t, condition.Equals(rel.(*plan.JoinRel).Expr()))
			assert.Equal(t, &types.BooleanType{Nullability: tc.want}, rel.(*plan.JoinRel).PostJoinFilter().GetType())
			encoded := RelToProto(rel)
			assert.True(t, gproto.Equal(wire, encoded))
			roundTrip, err := RelFromProto(encoded, joinTestRegistry())
			require.NoError(t, err)
			assert.Equal(t, rel.RecordType().Types(), roundTrip.RecordType().Types())
		})
	}
}
