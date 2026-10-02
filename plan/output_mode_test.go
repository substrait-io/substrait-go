// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/plan"
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
