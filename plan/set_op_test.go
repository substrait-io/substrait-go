// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/plan"
)

func TestSetOpString(t *testing.T) {
	for _, td := range []struct {
		o        plan.SetOp
		expected string
	}{
		{plan.SetOpUnspecified, "SET_OP_UNSPECIFIED"},
		{plan.SetOpMinusPrimary, "SET_OP_MINUS_PRIMARY"},
		{plan.SetOpMinusMultiset, "SET_OP_MINUS_MULTISET"},
		{plan.SetOpIntersectionPrimary, "SET_OP_INTERSECTION_PRIMARY"},
		{plan.SetOpIntersectionMultiset, "SET_OP_INTERSECTION_MULTISET"},
		{plan.SetOpUnionDistinct, "SET_OP_UNION_DISTINCT"},
		{plan.SetOpUnionAll, "SET_OP_UNION_ALL"},
		{plan.SetOpMinusPrimaryAll, "SET_OP_MINUS_PRIMARY_ALL"},
		{plan.SetOpIntersectionMultisetAll, "SET_OP_INTERSECTION_MULTISET_ALL"},
	} {
		t.Run(td.expected, func(t *testing.T) {
			assert.Equal(t, td.expected, td.o.String())
			assert.Equal(t, td.expected, fmt.Sprintf("%v", td.o))
		})
	}
}
