// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/plan"
	substraitproto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// A plan can hold no version, in which case Version() returns nil.
func TestPlanWithoutAVersion(t *testing.T) {
	p, err := plan.FromProto(&substraitproto.Plan{}, extensions.GetDefaultCollectionWithNoError())
	require.NoError(t, err)

	assert.Nil(t, p.Version())

	roundTrip, err := p.ToProto()
	require.NoError(t, err)
	assert.Nil(t, roundTrip.Version, "an absent version must not come back as an empty message")
}

// ToProto encodes a copy of the version, so editing the returned message doesn't rewrite the
// version every later plan reports.
func TestPlanToProtoCopiesTheVersion(t *testing.T) {
	producer := plan.CurrentVersion.Producer

	b := plan.NewBuilderDefault()
	p, err := b.Plan(b.NamedScan([]string{"test"}, baseSchema), []string{"a", "b"})
	require.NoError(t, err)

	protoPlan, err := p.ToProto()
	require.NoError(t, err)
	require.Equal(t, producer, protoPlan.Version.Producer)

	protoPlan.Version.Producer = "some other producer"
	assert.Equal(t, producer, plan.CurrentVersion.Producer)
	assert.Equal(t, producer, p.Version().Producer)
}
