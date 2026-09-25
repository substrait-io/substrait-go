// SPDX-License-Identifier: Apache-2.0

package wire_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

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
