// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
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

func TestAggregationPhaseMatchesProto(t *testing.T) {
	cases := []struct {
		domain types.AggregationPhase
		pb     proto.AggregationPhase
	}{
		{types.AggregationPhaseUnspecified, proto.AggregationPhase_AGGREGATION_PHASE_UNSPECIFIED},
		{types.AggregationPhaseInitialToIntermediate, proto.AggregationPhase_AGGREGATION_PHASE_INITIAL_TO_INTERMEDIATE},
		{types.AggregationPhaseIntermediateToIntermediate, proto.AggregationPhase_AGGREGATION_PHASE_INTERMEDIATE_TO_INTERMEDIATE},
		{types.AggregationPhaseInitialToResult, proto.AggregationPhase_AGGREGATION_PHASE_INITIAL_TO_RESULT},
		{types.AggregationPhaseIntermediateToResult, proto.AggregationPhase_AGGREGATION_PHASE_INTERMEDIATE_TO_RESULT},
	}
	// fail if the spec adds a value we don't mirror
	assert.Equal(t, proto.AggregationPhase_AGGREGATION_PHASE_UNSPECIFIED.Descriptor().Values().Len(), len(cases),
		"a proto aggregation phase value is not covered")
	for _, td := range cases {
		assert.EqualValues(t, td.pb, td.domain, "wire number for %s", td.pb)
		assert.Equal(t, td.pb.String(), td.domain.String(), "enum name for %s", td.pb)
	}
}
