// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/types"
)

func TestAggregationPhaseString(t *testing.T) {
	for _, td := range []struct {
		p        types.AggregationPhase
		expected string
	}{
		{types.AggregationPhaseUnspecified, "AGGREGATION_PHASE_UNSPECIFIED"},
		{types.AggregationPhaseInitialToIntermediate, "AGGREGATION_PHASE_INITIAL_TO_INTERMEDIATE"},
		{types.AggregationPhaseIntermediateToIntermediate, "AGGREGATION_PHASE_INTERMEDIATE_TO_INTERMEDIATE"},
		{types.AggregationPhaseInitialToResult, "AGGREGATION_PHASE_INITIAL_TO_RESULT"},
		{types.AggregationPhaseIntermediateToResult, "AGGREGATION_PHASE_INTERMEDIATE_TO_RESULT"},
		{types.AggregationPhase(99), "99"},
	} {
		t.Run(td.expected, func(t *testing.T) {
			assert.Equal(t, td.expected, td.p.String())
			assert.Equal(t, td.expected, fmt.Sprintf("%v", td.p))
		})
	}
}
