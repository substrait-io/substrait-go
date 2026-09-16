// SPDX-License-Identifier: Apache-2.0

package plan

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func TestSetRelWithoutCommon(t *testing.T) {
	input := createJoinInput("input")
	for _, tc := range []struct {
		name   string
		common *proto.RelCommon
	}{
		{"absent", nil},
		{"empty", &proto.RelCommon{}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			wire := &proto.Rel{RelType: &proto.Rel_Set{Set: &proto.SetRel{
				Common: tc.common,
				Inputs: []*proto.Rel{input.ToProto(), input.ToProto()},
				Op:     proto.SetRel_SET_OP_UNION_ALL,
			}}}
			var rel Rel
			var err error
			require.NotPanics(t, func() { rel, err = RelFromProto(wire, joinTestRegistry()) })
			require.NoError(t, err)
			assert.Equal(t, input.RecordType(), rel.RecordType())
		})
	}
}
