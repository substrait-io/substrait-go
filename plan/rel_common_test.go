// SPDX-License-Identifier: Apache-2.0

package plan

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func TestRelFromProtoWithoutCommon(t *testing.T) {
	input := createJoinInput("input")
	wire := &proto.Rel{RelType: &proto.Rel_Set{Set: &proto.SetRel{
		Inputs: []*proto.Rel{input.ToProto(), input.ToProto()},
		Op:     proto.SetRel_SET_OP_UNION_ALL,
	}}}

	rel, err := RelFromProto(wire, joinTestRegistry())
	require.NoError(t, err)
	assert.Equal(t, input.RecordType(), rel.RecordType())
}
