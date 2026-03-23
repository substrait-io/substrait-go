// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v8/expr"
	"github.com/substrait-io/substrait-go/v8/extensions"
	"github.com/substrait-io/substrait-go/v8/plan"
	"github.com/substrait-io/substrait-go/v8/types"
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
			planJSON, err := testdata.ReadFile(fmt.Sprintf("testdata/%s.json", name))
			require.NoError(t, err)

			var protoPlan substraitproto.Plan
			require.NoError(t, protojson.Unmarshal(planJSON, &protoPlan))

			p, err := plan.FromProto(&protoPlan, extensions.GetDefaultCollectionWithNoError())
			require.NoError(t, err)

			backToProto, err := p.ToProto()
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

	protoPlan, err := p.ToProto()
	require.NoError(t, err)
	assert.Empty(t, protoPlan.ParameterBindings)
}

func TestDynamicParameterBindingTypeMismatch(t *testing.T) {
	b := plan.NewBuilderDefault()
	scan := b.NamedScan([]string{"test"}, baseSchema2)

	tests := []struct {
		name      string
		dpType    types.Type
		bindValue expr.Literal
		errMsg    string
	}{
		{
			name:      "i32 param bound to string literal",
			dpType:    &types.Int32Type{Nullability: types.NullabilityRequired},
			bindValue: expr.NewPrimitiveLiteral("hello", false),
			errMsg:    "parameter binding for anchor 0 has type",
		},
		{
			name:      "string param bound to i32 literal",
			dpType:    &types.StringType{Nullability: types.NullabilityNullable},
			bindValue: expr.NewPrimitiveLiteral(int32(42), false),
			errMsg:    "parameter binding for anchor 0 has type",
		},
		{
			name:      "fp64 param bound to i64 literal",
			dpType:    &types.Float64Type{Nullability: types.NullabilityRequired},
			bindValue: expr.NewPrimitiveLiteral(int64(100), false),
			errMsg:    "parameter binding for anchor 0 has type",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dp := &expr.DynamicParameter{
				OutputType:         tt.dpType,
				ParameterReference: 0,
			}

			project, err := b.Project(scan, dp)
			require.NoError(t, err)

			bindings := []plan.DynamicParameterBinding{
				{
					ParameterAnchor: 0,
					Value:           tt.bindValue,
				},
			}

			_, err = b.PlanWithBindings(project, []string{"x", "y", "p"}, nil, bindings)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.errMsg)
		})
	}
}

func TestDynamicParameterBindingMissingAnchor(t *testing.T) {
	b := plan.NewBuilderDefault()
	scan := b.NamedScan([]string{"test"}, baseSchema2)

	dp := &expr.DynamicParameter{
		OutputType:         &types.Int32Type{Nullability: types.NullabilityRequired},
		ParameterReference: 0,
	}

	project, err := b.Project(scan, dp)
	require.NoError(t, err)

	bindings := []plan.DynamicParameterBinding{
		{
			ParameterAnchor: 99,
			Value:           expr.NewPrimitiveLiteral(int32(42), false),
		},
	}

	_, err = b.PlanWithBindings(project, []string{"x", "y", "p"}, nil, bindings)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "no DynamicParameter with that reference exists")
}

func TestDynamicParameterBindingNullabilityMismatch(t *testing.T) {
	b := plan.NewBuilderDefault()
	scan := b.NamedScan([]string{"test"}, baseSchema2)

	dp := &expr.DynamicParameter{
		OutputType:         &types.Int32Type{Nullability: types.NullabilityRequired},
		ParameterReference: 0,
	}

	project, err := b.Project(scan, dp)
	require.NoError(t, err)

	bindings := []plan.DynamicParameterBinding{
		{
			ParameterAnchor: 0,
			Value:           expr.NewPrimitiveLiteral(int32(42), true), // nullable literal
		},
	}

	p, err := b.PlanWithBindings(project, []string{"x", "y", "p"}, nil, bindings)
	require.NoError(t, err)
	assert.NotNil(t, p)
}

func TestDynamicParameterBindingInFilter(t *testing.T) {
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

	// Wrong type binding should fail
	wrongBindings := []plan.DynamicParameterBinding{
		{
			ParameterAnchor: 0,
			Value:           expr.NewPrimitiveLiteral("not-a-number", false),
		},
	}
	_, err = b.PlanWithBindings(filter, []string{"x", "y"}, nil, wrongBindings)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "parameter binding for anchor 0 has type")

	// Correct type binding should succeed
	goodBindings := []plan.DynamicParameterBinding{
		{
			ParameterAnchor: 0,
			Value:           expr.NewPrimitiveLiteral(int32(42), false),
		},
	}
	p, err := b.PlanWithBindings(filter, []string{"x", "y"}, nil, goodBindings)
	require.NoError(t, err)
	assert.NotNil(t, p)
}
