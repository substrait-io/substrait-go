// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/plan"
)

func TestSetComparisonReductionOpString(t *testing.T) {
	for _, td := range []struct {
		o        plan.SetComparisonReductionOp
		expected string
	}{
		{plan.SetComparisonReductionOpUnspecified, "REDUCTION_OP_UNSPECIFIED"},
		{plan.SetComparisonReductionOpAny, "REDUCTION_OP_ANY"},
		{plan.SetComparisonReductionOpAll, "REDUCTION_OP_ALL"},
		{plan.SetComparisonReductionOp(99), "99"},
	} {
		t.Run(td.expected, func(t *testing.T) {
			assert.Equal(t, td.expected, td.o.String())
			assert.Equal(t, td.expected, fmt.Sprintf("%v", td.o))
		})
	}
}
