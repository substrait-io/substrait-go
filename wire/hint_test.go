// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/plan"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func TestHintProtoRoundTrip(t *testing.T) {
	h := &plan.Hint{
		Stats:       &plan.Stats{RowCount: 100, RecordSize: 8},
		Constraint:  &plan.RuntimeConstraint{},
		Alias:       "my_alias",
		OutputNames: []string{"a", "b"},
		SavedComputations: []*plan.SavedComputation{
			{ComputationID: 3, Type: plan.ComputationTypeHashTable},
		},
		LoadedComputations: []*plan.LoadedComputation{
			{ComputationIDReference: 3, Type: plan.ComputationTypeBloomFilter},
		},
	}

	p := hintToProto(h)
	require.NotNil(t, p)
	assert.Equal(t, "my_alias", p.Alias)
	assert.Equal(t, []string{"a", "b"}, p.OutputNames)
	require.NotNil(t, p.Stats)
	assert.Equal(t, float64(100), p.Stats.RowCount)
	require.Len(t, p.SavedComputations, 1)
	assert.Equal(t, int32(3), p.SavedComputations[0].ComputationId)
	assert.Equal(t, proto.RelCommon_Hint_COMPUTATION_TYPE_HASHTABLE, p.SavedComputations[0].Type)
	require.Len(t, p.LoadedComputations, 1)
	assert.Equal(t, proto.RelCommon_Hint_COMPUTATION_TYPE_BLOOM_FILTER, p.LoadedComputations[0].Type)

	got := hintFromProto(p)
	assert.Equal(t, h, got)
}

func TestHintProtoNil(t *testing.T) {
	assert.Nil(t, hintToProto(nil))
	assert.Nil(t, hintFromProto(nil))
}
