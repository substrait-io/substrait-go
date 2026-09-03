// SPDX-License-Identifier: Apache-2.0

package codec_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/codec"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/plan"
	"github.com/substrait-io/substrait-go/v9/types"
	substraitproto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func TestDynamicParameterPlanRoundtrip(t *testing.T) {
	for _, name := range []string{
		"dynamic_parameter_plan",
		"dynamic_parameter_filter",
	} {
		t.Run(name, func(t *testing.T) {
			planJSON, err := testdataFS.ReadFile(fmt.Sprintf("testdata/plan/%s.json", name))
			require.NoError(t, err)

			var protoPlan substraitproto.Plan
			require.NoError(t, protojson.Unmarshal(planJSON, &protoPlan))

			p, err := codec.PlanFromProto(&protoPlan, extensions.GetDefaultCollectionWithNoError())
			require.NoError(t, err)

			backToProto, err := codec.PlanToProto(p)
			require.NoError(t, err)
			assert.Truef(t, proto.Equal(&protoPlan, backToProto),
				"expected: %s\ngot: %s",
				protojson.Format(&protoPlan), protojson.Format(backToProto))
		})
	}
}

func TestDynamicParameterPlanWithoutBindings(t *testing.T) {
	b := plan.NewBuilderDefault()
	scan := b.NamedScan([]string{"test"}, baseSchema2)

	dp := &expr.DynamicParameter{
		OutputType:         &types.Int32Type{Nullability: types.NullabilityRequired},
		ParameterReference: 0,
	}

	project, err := b.Project(scan, dp)
	require.NoError(t, err)

	p, err := b.Plan(project, []string{"x", "y", "param"})
	require.NoError(t, err)

	assert.Empty(t, p.ParameterBindings())

	protoPlan, err := codec.PlanToProto(p)
	require.NoError(t, err)
	assert.Empty(t, protoPlan.ParameterBindings)
}
