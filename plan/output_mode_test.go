// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/plan"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func TestOutputModeString(t *testing.T) {
	for _, td := range []struct {
		m        plan.OutputMode
		expected string
	}{
		{plan.OutputModeUnspecified, "OUTPUT_MODE_UNSPECIFIED"},
		{plan.OutputModeNoOutput, "OUTPUT_MODE_NO_OUTPUT"},
		{plan.OutputModeModifiedRecords, "OUTPUT_MODE_MODIFIED_RECORDS"},
	} {
		t.Run(td.expected, func(t *testing.T) {
			assert.Equal(t, td.expected, td.m.String())
			assert.Equal(t, td.expected, fmt.Sprintf("%v", td.m))
		})
	}
}

func TestOutputModeMatchesProto(t *testing.T) {
	cases := []struct {
		domain plan.OutputMode
		pb     proto.WriteRel_OutputMode
	}{
		{plan.OutputModeUnspecified, proto.WriteRel_OUTPUT_MODE_UNSPECIFIED},
		{plan.OutputModeNoOutput, proto.WriteRel_OUTPUT_MODE_NO_OUTPUT},
		{plan.OutputModeModifiedRecords, proto.WriteRel_OUTPUT_MODE_MODIFIED_RECORDS},
	}
	// fail if the spec adds a value we don't mirror
	assert.Equal(t, proto.WriteRel_OUTPUT_MODE_UNSPECIFIED.Descriptor().Values().Len(), len(cases),
		"a proto output mode value is not covered")
	for _, td := range cases {
		assert.EqualValues(t, td.pb, td.domain, "wire number for %s", td.pb)
		assert.Equal(t, td.pb.String(), td.domain.String(), "enum name for %s", td.pb)
	}
}
