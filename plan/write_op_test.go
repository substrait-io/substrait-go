// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/plan"
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
