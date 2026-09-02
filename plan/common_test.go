// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/plan"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// Round-trip a fully populated Hint through the proto boundary, so a field-mapping
// mistake or a dropped field in the converters fails here.
func TestHintRoundTrip(t *testing.T) {
	ae := &extensions.AdvancedExtension{
		Optimizations: []*extensions.Optimization{{TypeUrl: "opt", Value: []byte("o")}},
		Enhancement:   &extensions.Enhancement{TypeUrl: "enh", Value: []byte("e")},
	}
	h := &plan.Hint{
		Stats:              &plan.Stats{RowCount: 10, RecordSize: 20, AdvancedExtension: ae},
		Constraint:         &plan.RuntimeConstraint{AdvancedExtension: ae},
		Alias:              "myrel",
		OutputNames:        []string{"a", "b"},
		AdvancedExtension:  ae,
		SavedComputations:  []*plan.SavedComputation{{ComputationID: 1, Type: plan.ComputationTypeHashTable, AdvancedExtension: ae}},
		LoadedComputations: []*plan.LoadedComputation{{ComputationIDReference: 1, Type: plan.ComputationTypeBloomFilter, AdvancedExtension: ae}},
	}
	assert.Equal(t, h, plan.HintFromProto(plan.HintToProto(h)))

	assert.Nil(t, plan.HintToProto(nil))
	assert.Nil(t, plan.HintFromProto(nil))
}

// TestRelCommonConvertersNil verifies each Hint-tree converter maps a nil message
// to nil rather than panicking.
func TestRelCommonConvertersNil(t *testing.T) {
	assert.Nil(t, plan.HintFromProto(nil))
	assert.Nil(t, plan.HintToProto(nil))
	assert.Nil(t, plan.StatsFromProto(nil))
	assert.Nil(t, plan.StatsToProto(nil))
	assert.Nil(t, plan.RuntimeConstraintFromProto(nil))
	assert.Nil(t, plan.RuntimeConstraintToProto(nil))
	assert.Nil(t, plan.SavedComputationFromProto(nil))
	assert.Nil(t, plan.SavedComputationToProto(nil))
	assert.Nil(t, plan.LoadedComputationFromProto(nil))
	assert.Nil(t, plan.LoadedComputationToProto(nil))
}

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
