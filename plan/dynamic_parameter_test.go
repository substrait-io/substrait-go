// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/plan"
	"github.com/substrait-io/substrait-go/v9/types"
)

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
