// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/plan"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

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
