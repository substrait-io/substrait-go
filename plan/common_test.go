// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/plan"
)

func TestComputationTypeString(t *testing.T) {
	for _, td := range []struct {
		c        plan.ComputationType
		expected string
	}{
		{plan.ComputationTypeUnspecified, "COMPUTATION_TYPE_UNSPECIFIED"},
		{plan.ComputationTypeHashTable, "COMPUTATION_TYPE_HASHTABLE"},
		{plan.ComputationTypeBloomFilter, "COMPUTATION_TYPE_BLOOM_FILTER"},
		{plan.ComputationTypeUnknown, "COMPUTATION_TYPE_UNKNOWN"},
		{plan.ComputationType(7), "7"},
	} {
		t.Run(td.expected, func(t *testing.T) {
			assert.Equal(t, td.expected, td.c.String())
		})
	}
}
