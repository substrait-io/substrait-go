// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/plan"
)

func TestSetComparisonOpString(t *testing.T) {
	for _, td := range []struct {
		o        plan.SetComparisonOp
		expected string
	}{
		{plan.SetComparisonOpUnspecified, "COMPARISON_OP_UNSPECIFIED"},
		{plan.SetComparisonOpEq, "COMPARISON_OP_EQ"},
		{plan.SetComparisonOpNe, "COMPARISON_OP_NE"},
		{plan.SetComparisonOpLt, "COMPARISON_OP_LT"},
		{plan.SetComparisonOpGt, "COMPARISON_OP_GT"},
		{plan.SetComparisonOpLe, "COMPARISON_OP_LE"},
		{plan.SetComparisonOpGe, "COMPARISON_OP_GE"},
		{plan.SetComparisonOp(99), "99"},
	} {
		t.Run(td.expected, func(t *testing.T) {
			assert.Equal(t, td.expected, td.o.String())
			assert.Equal(t, td.expected, fmt.Sprintf("%v", td.o))
		})
	}
}
