// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/plan"
)

func TestSetPredicateOpString(t *testing.T) {
	for _, td := range []struct {
		o        plan.SetPredicateOp
		expected string
	}{
		{plan.SetPredicateOpUnspecified, "PREDICATE_OP_UNSPECIFIED"},
		{plan.SetPredicateOpExists, "PREDICATE_OP_EXISTS"},
		{plan.SetPredicateOpUnique, "PREDICATE_OP_UNIQUE"},
		{plan.SetPredicateOp(99), "99"},
	} {
		t.Run(td.expected, func(t *testing.T) {
			assert.Equal(t, td.expected, td.o.String())
			assert.Equal(t, td.expected, fmt.Sprintf("%v", td.o))
		})
	}
}
