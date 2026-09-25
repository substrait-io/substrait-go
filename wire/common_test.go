// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/plan"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// TestComputationTypeMatchesProto keeps the domain enum in lockstep with the spec:
// every value maps to its protobuf counterpart, and a new spec value fails the count.
func TestComputationTypeMatchesProto(t *testing.T) {
	cases := []struct {
		domain plan.ComputationType
		pb     proto.RelCommon_Hint_ComputationType
	}{
		{plan.ComputationTypeUnspecified, proto.RelCommon_Hint_COMPUTATION_TYPE_UNSPECIFIED},
		{plan.ComputationTypeHashTable, proto.RelCommon_Hint_COMPUTATION_TYPE_HASHTABLE},
		{plan.ComputationTypeBloomFilter, proto.RelCommon_Hint_COMPUTATION_TYPE_BLOOM_FILTER},
		{plan.ComputationTypeUnknown, proto.RelCommon_Hint_COMPUTATION_TYPE_UNKNOWN},
	}
	assert.Equal(t, proto.RelCommon_Hint_COMPUTATION_TYPE_UNSPECIFIED.Descriptor().Values().Len(), len(cases),
		"the set of spec computation types changed")
	for _, c := range cases {
		assert.EqualValues(t, c.pb, c.domain, "%s wire value", c.domain)
		assert.Equal(t, c.pb.String(), c.domain.String(), "%s name", c.domain)
	}
}
