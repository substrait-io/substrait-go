// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v7/expr"
	"github.com/substrait-io/substrait-go/v7/extensions"
	"github.com/substrait-io/substrait-go/v7/plan"
	"github.com/substrait-io/substrait-go/v7/types"
	substraitproto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func TestDynamicParameterInFilterPlan(t *testing.T) {
	b := plan.NewBuilderDefault()
	scan := b.NamedScan([]string{"test"}, baseSchema2)

	dp := &expr.DynamicParameter{
		OutputType:         &types.Int32Type{Nullability: types.NullabilityRequired},
		ParameterReference: 0,
	}

	ref, err := b.RootFieldRef(scan, 0)
	require.NoError(t, err)

	gt, err := b.ScalarFn(extensions.SubstraitDefaultURNPrefix+"functions_comparison", "gt", nil, ref, dp)
	require.NoError(t, err)

	filter, err := b.Filter(scan, gt)
	require.NoError(t, err)

	bindings := []plan.DynamicParameterBinding{
		{
			ParameterAnchor: 0,
			Value:           expr.NewPrimitiveLiteral(int32(42), false),
		},
	}

	p, err := b.PlanWithBindings(filter, []string{"x", "y"}, nil, bindings)
	require.NoError(t, err)

	protoPlan, err := p.ToProto()
	require.NoError(t, err)

	roundTrip, err := plan.FromProto(protoPlan, extensions.GetDefaultCollectionWithNoError())
	require.NoError(t, err)

	roundTripProto, err := roundTrip.ToProto()
	require.NoError(t, err)

	assert.Truef(t, proto.Equal(protoPlan, roundTripProto), "plan expected: %s\ngot: %s",
		protojson.Format(protoPlan), protojson.Format(roundTripProto))
}

func TestDynamicParameterInProjectPlan(t *testing.T) {
	b := plan.NewBuilderDefault()
	scan := b.NamedScan([]string{"test"}, baseSchema2)

	dp := &expr.DynamicParameter{
		OutputType:         &types.StringType{Nullability: types.NullabilityNullable},
		ParameterReference: 0,
	}

	project, err := b.Project(scan, dp)
	require.NoError(t, err)

	bindings := []plan.DynamicParameterBinding{
		{
			ParameterAnchor: 0,
			Value:           expr.NewPrimitiveLiteral("hello", false),
		},
	}

	p, err := b.PlanWithBindings(project, []string{"x", "y", "param_val"}, nil, bindings)
	require.NoError(t, err)

	protoPlan, err := p.ToProto()
	require.NoError(t, err)

	roundTrip, err := plan.FromProto(protoPlan, extensions.GetDefaultCollectionWithNoError())
	require.NoError(t, err)

	roundTripProto, err := roundTrip.ToProto()
	require.NoError(t, err)
	assert.True(t, proto.Equal(protoPlan, roundTripProto))
}

func TestDynamicParameterMultipleBindings(t *testing.T) {
	b := plan.NewBuilderDefault()
	scan := b.NamedScan([]string{"test"}, baseSchema2)

	dp0 := &expr.DynamicParameter{
		OutputType:         &types.Int32Type{Nullability: types.NullabilityRequired},
		ParameterReference: 0,
	}

	dp1 := &expr.DynamicParameter{
		OutputType:         &types.StringType{Nullability: types.NullabilityNullable},
		ParameterReference: 1,
	}

	project, err := b.Project(scan, dp0, dp1)
	require.NoError(t, err)

	bindings := []plan.DynamicParameterBinding{
		{
			ParameterAnchor: 0,
			Value:           expr.NewPrimitiveLiteral(int32(100), false),
		},
		{
			ParameterAnchor: 1,
			Value:           expr.NewPrimitiveLiteral("world", true),
		},
	}

	p, err := b.PlanWithBindings(project, []string{"x", "y", "p0", "p1"}, nil, bindings)
	require.NoError(t, err)

	protoPlan, err := p.ToProto()
	require.NoError(t, err)

	roundTrip, err := plan.FromProto(protoPlan, extensions.GetDefaultCollectionWithNoError())
	require.NoError(t, err)

	roundTripProto, err := roundTrip.ToProto()
	require.NoError(t, err)
	assert.True(t, proto.Equal(protoPlan, roundTripProto))
}

func TestDynamicParameterPlanWithoutBindings(t *testing.T) {
	// Plan with dynamic parameters but no bindings (valid use case)
	b := plan.NewBuilderDefault()
	scan := b.NamedScan([]string{"test"}, baseSchema2)

	dp := &expr.DynamicParameter{
		OutputType:         &types.Int32Type{Nullability: types.NullabilityRequired},
		ParameterReference: 0,
	}

	project, err := b.Project(scan, dp)
	require.NoError(t, err)

	// Use regular Plan (no bindings)
	p, err := b.Plan(project, []string{"x", "y", "param"})
	require.NoError(t, err)

	// Verify no bindings
	assert.Empty(t, p.ParameterBindings())

	// Proto roundtrip should still work
	protoPlan, err := p.ToProto()
	require.NoError(t, err)
	assert.Empty(t, protoPlan.ParameterBindings)

	roundTrip, err := plan.FromProto(protoPlan, extensions.GetDefaultCollectionWithNoError())
	require.NoError(t, err)
	assert.Empty(t, roundTrip.ParameterBindings())
}

func TestDynamicParameterFromProtoJSON(t *testing.T) {
	planJSON, err := os.ReadFile("testdata/dynamic_parameter_plan.json")
	require.NoError(t, err)

	var protoPlan substraitproto.Plan
	require.NoError(t, protojson.Unmarshal(planJSON, &protoPlan))

	p, err := plan.FromProto(&protoPlan, extensions.GetDefaultCollectionWithNoError())
	require.NoError(t, err)

	backToProto, err := p.ToProto()
	require.NoError(t, err)
	assert.Truef(t, proto.Equal(&protoPlan, backToProto), "expected: %s\ngot: %s",
		protojson.Format(&protoPlan), protojson.Format(backToProto))
}

func TestDynamicParameterBuilderInPlanBuilder(t *testing.T) {
	b := plan.NewBuilderDefault()
	eb := b.GetExprBuilder()

	scan := b.NamedScan([]string{"employees"}, types.NamedStruct{
		Names: []string{"id", "salary"},
		Struct: types.StructType{
			Nullability: types.NullabilityRequired,
			Types: []types.Type{
				&types.Int64Type{Nullability: types.NullabilityRequired},
				&types.Float64Type{Nullability: types.NullabilityRequired},
			},
		},
	})

	dpExpr, err := eb.DynamicParam(
		&types.Float64Type{Nullability: types.NullabilityRequired}, 0,
	).BuildExpr()
	require.NoError(t, err)

	project, err := b.Project(scan, dpExpr)
	require.NoError(t, err)

	bindings := []plan.DynamicParameterBinding{
		{
			ParameterAnchor: 0,
			Value:           expr.NewPrimitiveLiteral(float64(50000.0), false),
		},
	}

	p, err := b.PlanWithBindings(project, []string{"id", "salary", "threshold"}, nil, bindings)
	require.NoError(t, err)

	protoPlan, err := p.ToProto()
	require.NoError(t, err)

	roundTrip, err := plan.FromProto(protoPlan, extensions.GetDefaultCollectionWithNoError())
	require.NoError(t, err)

	roundTripProto, err := roundTrip.ToProto()
	require.NoError(t, err)
	assert.True(t, proto.Equal(protoPlan, roundTripProto))
}
