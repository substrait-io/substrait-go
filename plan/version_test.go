// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v8/extensions"
	"github.com/substrait-io/substrait-go/v8/plan"
	substraitproto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// The version field is optional on the wire, so a plan can hold no version at all. Callers reach it
// through the plan.Version interface, which gives them no way to check for nil first.
func TestPlanWithoutAVersion(t *testing.T) {
	p, err := plan.FromProto(&substraitproto.Plan{}, extensions.GetDefaultCollectionWithNoError())
	require.NoError(t, err)

	// the interface value itself is not nil, it holds a nil *types.Version, which is the case the
	// getters have to survive
	v := p.Version()
	assert.Zero(t, v.GetMajorNumber())
	assert.Zero(t, v.GetMinorNumber())
	assert.Zero(t, v.GetPatchNumber())
	assert.Empty(t, v.GetGitHash())
	assert.Empty(t, v.GetProducer())
	assert.Equal(t, "<nil>", v.String())

	roundTrip, err := p.ToProto()
	require.NoError(t, err)
	assert.Nil(t, roundTrip.Version, "an absent version must not come back as an empty message")
}

// ToProto used to hand out &CurrentVersion itself, since types.Version was the generated message.
// It now encodes a copy, so a caller editing the returned message no longer rewrites the version
// every later plan reports.
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
	assert.Equal(t, producer, p.Version().GetProducer())
}
