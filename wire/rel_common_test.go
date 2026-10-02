// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func TestRelFromProtoWithoutCommon(t *testing.T) {
	input := createJoinInput("input")
	rel := &proto.Rel{RelType: &proto.Rel_Set{Set: &proto.SetRel{
		Inputs: []*proto.Rel{RelToProto(input), RelToProto(input)},
		Op:     proto.SetRel_SET_OP_UNION_ALL,
	}}}

	out, err := RelFromProto(rel, joinTestRegistry())
	require.NoError(t, err)
	assert.Equal(t, input.RecordType(), out.RecordType())
}
