// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/plan"
)

func TestSimpleComparisonTypeString(t *testing.T) {
	for _, td := range []struct {
		c        plan.SimpleComparisonType
		expected string
	}{
		{plan.SimpleComparisonTypeUnspecified, "SIMPLE_COMPARISON_TYPE_UNSPECIFIED"},
		{plan.SimpleComparisonTypeEq, "SIMPLE_COMPARISON_TYPE_EQ"},
		{plan.SimpleComparisonTypeIsNotDistinctFrom, "SIMPLE_COMPARISON_TYPE_IS_NOT_DISTINCT_FROM"},
		{plan.SimpleComparisonTypeMightEqual, "SIMPLE_COMPARISON_TYPE_MIGHT_EQUAL"},
	} {
		t.Run(td.expected, func(t *testing.T) {
			assert.Equal(t, td.expected, td.c.String())
			assert.Equal(t, td.expected, fmt.Sprintf("%v", td.c))
		})
	}
}
