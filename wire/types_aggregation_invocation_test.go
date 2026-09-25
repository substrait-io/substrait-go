// SPDX-License-Identifier: Apache-2.0

package wire_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func TestAggregationInvocationMatchesProto(t *testing.T) {
	cases := []struct {
		domain types.AggregationInvocation
		pb     proto.AggregateFunction_AggregationInvocation
	}{
		{types.AggregationInvocationUnspecified, proto.AggregateFunction_AGGREGATION_INVOCATION_UNSPECIFIED},
		{types.AggregationInvocationAll, proto.AggregateFunction_AGGREGATION_INVOCATION_ALL},
		{types.AggregationInvocationDistinct, proto.AggregateFunction_AGGREGATION_INVOCATION_DISTINCT},
	}
	// fail if the spec adds a value we don't mirror
	assert.Equal(t, proto.AggregateFunction_AGGREGATION_INVOCATION_UNSPECIFIED.Descriptor().Values().Len(), len(cases),
		"a proto aggregation invocation value is not covered")
	for _, td := range cases {
		assert.EqualValues(t, td.pb, td.domain, "wire number for %s", td.pb)
		assert.Equal(t, td.pb.String(), td.domain.String(), "enum name for %s", td.pb)
	}
}
