// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/plan"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func TestWriteOpString(t *testing.T) {
	for _, td := range []struct {
		o        plan.WriteOp
		expected string
	}{
		{plan.WriteOpUnspecified, "WRITE_OP_UNSPECIFIED"},
		{plan.WriteOpInsert, "WRITE_OP_INSERT"},
		{plan.WriteOpDelete, "WRITE_OP_DELETE"},
		{plan.WriteOpUpdate, "WRITE_OP_UPDATE"},
		{plan.WriteOpCTAS, "WRITE_OP_CTAS"},
	} {
		t.Run(td.expected, func(t *testing.T) {
			assert.Equal(t, td.expected, td.o.String())
			assert.Equal(t, td.expected, fmt.Sprintf("%v", td.o))
		})
	}
}

func TestWriteOpMatchesProto(t *testing.T) {
	cases := []struct {
		domain plan.WriteOp
		pb     proto.WriteRel_WriteOp
	}{
		{plan.WriteOpUnspecified, proto.WriteRel_WRITE_OP_UNSPECIFIED},
		{plan.WriteOpInsert, proto.WriteRel_WRITE_OP_INSERT},
		{plan.WriteOpDelete, proto.WriteRel_WRITE_OP_DELETE},
		{plan.WriteOpUpdate, proto.WriteRel_WRITE_OP_UPDATE},
		{plan.WriteOpCTAS, proto.WriteRel_WRITE_OP_CTAS},
	}
	// fail if the spec adds a value we don't mirror
	assert.Equal(t, proto.WriteRel_WRITE_OP_UNSPECIFIED.Descriptor().Values().Len(), len(cases),
		"a proto write op value is not covered")
	for _, td := range cases {
		assert.EqualValues(t, td.pb, td.domain, "wire number for %s", td.pb)
		assert.Equal(t, td.pb.String(), td.domain.String(), "enum name for %s", td.pb)
	}
}
