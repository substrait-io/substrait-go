// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/types"
)

func TestAggregationInvocationString(t *testing.T) {
	for _, td := range []struct {
		i        types.AggregationInvocation
		expected string
	}{
		{types.AggregationInvocationUnspecified, "AGGREGATION_INVOCATION_UNSPECIFIED"},
		{types.AggregationInvocationAll, "AGGREGATION_INVOCATION_ALL"},
		{types.AggregationInvocationDistinct, "AGGREGATION_INVOCATION_DISTINCT"},
		{types.AggregationInvocation(99), "99"},
	} {
		t.Run(td.expected, func(t *testing.T) {
			assert.Equal(t, td.expected, td.i.String())
			assert.Equal(t, td.expected, fmt.Sprintf("%v", td.i))
		})
	}
}
