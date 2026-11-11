// SPDX-License-Identifier: Apache-2.0

package substraitproto

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/plan"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	gproto "google.golang.org/protobuf/proto"
)

func TestJoinPostFilterPreservesUnresolvedOuterReference(t *testing.T) {
	b := plan.NewBuilderDefault()
	schema := types.NamedStruct{
		Names: []string{"flag"},
		Struct: types.StructType{Nullability: types.NullabilityRequired, Types: []types.Type{
			&types.BooleanType{Nullability: types.NullabilityRequired},
		}},
	}
	left := RelToProto(b.NamedScan([]string{"left"}, schema))
	right := RelToProto(b.NamedScan([]string{"right"}, schema))
	outerRef := fieldRefProto(0)
	outerRef.RootType = &proto.Expression_FieldReference_OuterReference_{
		OuterReference: &proto.Expression_FieldReference_OuterReference{StepsOut: 1},
	}
	// A decoded outer reference has no type until the enclosing scope binds it.
	projected := &proto.Rel{RelType: &proto.Rel_Project{Project: &proto.ProjectRel{
		Common: &proto.RelCommon{EmitKind: &proto.RelCommon_Emit_{
			Emit: &proto.RelCommon_Emit{OutputMapping: []int32{1}},
		}},
		Input:       right,
		Expressions: []*proto.Expression{{RexType: &proto.Expression_Selection{Selection: outerRef}}},
	}}}
	for _, kind := range []proto.JoinRel_JoinType{
		proto.JoinRel_JOIN_TYPE_LEFT, proto.JoinRel_JOIN_TYPE_LEFT_SINGLE,
		proto.JoinRel_JOIN_TYPE_RIGHT, proto.JoinRel_JOIN_TYPE_RIGHT_SINGLE,
		proto.JoinRel_JOIN_TYPE_OUTER,
	} {
		t.Run(kind.String(), func(t *testing.T) {
			leftInput, rightInput := left, projected
			if kind == proto.JoinRel_JOIN_TYPE_RIGHT || kind == proto.JoinRel_JOIN_TYPE_RIGHT_SINGLE {
				leftInput, rightInput = projected, left
			}
			field := int32(0)
			if leftInput == projected {
				field = 1
			}
			wire := &proto.Rel{RelType: &proto.Rel_Join{Join: &proto.JoinRel{
				Common: &proto.RelCommon{EmitKind: &proto.RelCommon_Direct_{Direct: &proto.RelCommon_Direct{}}},
				Left:   leftInput, Right: rightInput, Type: kind,
				Expression:     ExprToProto(expr.NewPrimitiveLiteral(true, false)),
				PostJoinFilter: postJoinFilterField(field),
			}}}
			rel, err := RelFromProto(wire, joinTestRegistry())
			require.NoError(t, err)
			fields := rel.RecordType().Types()
			assert.Nil(t, fields[1-field])
			want := types.NullabilityRequired
			if kind == proto.JoinRel_JOIN_TYPE_OUTER {
				want = types.NullabilityNullable
			}
			assert.Equal(t, &types.BooleanType{Nullability: want}, fields[field])
			assert.Equal(t, fields[field], rel.(*plan.JoinRel).PostJoinFilter().GetType())
			assert.True(t, gproto.Equal(wire, RelToProto(rel)))
		})
	}
}
