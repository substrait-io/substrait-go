// SPDX-License-Identifier: Apache-2.0

package codec_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/codec"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/plan"
	"github.com/substrait-io/substrait-go/v9/types"
	extensionspb "github.com/substrait-io/substrait-protobuf/go/substraitpb/extensions"
	"google.golang.org/protobuf/types/known/anypb"
)

var baseSchema2 = types.NamedStruct{Names: []string{"x", "y"},
	Struct: types.StructType{
		Nullability: types.NullabilityRequired,
		Types: []types.Type{
			&types.Int32Type{Nullability: types.NullabilityRequired},
			&types.BooleanType{Nullability: types.NullabilityRequired},
		},
	}}

var planCases = []struct {
	name  string
	build func(t *testing.T) *plan.Plan
}{
	{"named_scan", func(t *testing.T) *plan.Plan {
		b := plan.NewBuilderDefault()
		root, err := b.NamedScan([]string{"test"}, baseSchema2).Remap(1, 0)
		require.NoError(t, err)
		p, err := b.Plan(root, []string{"a", "b"})
		require.NoError(t, err)
		return p
	}},
	{"fetch", func(t *testing.T) *plan.Plan {
		b := plan.NewBuilderDefault()
		root, err := b.Fetch(b.NamedScan([]string{"test"}, baseSchema2), 1, 2)
		require.NoError(t, err)
		p, err := b.Plan(root, []string{"x", "y"})
		require.NoError(t, err)
		return p
	}},
	{"parameter_bindings", func(t *testing.T) *plan.Plan {
		b := plan.NewBuilderDefault()
		dp := &expr.DynamicParameter{
			OutputType:         &types.Int32Type{Nullability: types.NullabilityRequired},
			ParameterReference: 0,
		}
		project, err := b.Project(b.NamedScan([]string{"test"}, baseSchema2), dp)
		require.NoError(t, err)
		bindings := []plan.DynamicParameterBinding{
			{ParameterAnchor: 0, Value: expr.NewPrimitiveLiteral(int32(42), false)},
		}
		p, err := b.PlanWithBindings(project, []string{"x", "y", "p"}, []string{"urn:example"}, bindings)
		require.NoError(t, err)
		return p
	}},
	{"advanced_extension", func(t *testing.T) *plan.Plan {
		b := plan.NewBuilderDefault()
		root, err := b.NamedScan([]string{"test"}, baseSchema2).Remap(0, 1)
		require.NoError(t, err)
		base, err := b.Plan(root, []string{"x", "y"})
		require.NoError(t, err)

		pb, err := codec.PlanToProto(base)
		require.NoError(t, err)
		pb.AdvancedExtensions = &extensionspb.AdvancedExtension{
			Enhancement: &anypb.Any{TypeUrl: "urn:enh", Value: []byte("e")},
		}
		p, err := codec.PlanFromProto(pb, extensions.GetDefaultCollectionWithNoError())
		require.NoError(t, err)
		return p
	}},
}

func TestPlanToProto(t *testing.T) {
	for _, tc := range planCases {
		t.Run(tc.name, func(t *testing.T) {
			p := tc.build(t)

			got, err := codec.PlanToProto(p)
			require.NoError(t, err)

			checkGolden(t, "plan", tc.name, got)

			back, err := codec.PlanFromProto(got, extensions.GetDefaultCollectionWithNoError())
			require.NoError(t, err)
			regot, err := codec.PlanToProto(back)
			require.NoError(t, err)
			if diff := protoDiff(got, regot); diff != "" {
				t.Errorf("%s round-trip not stable (-got +regot):\n%s", tc.name, diff)
			}
		})
	}
}
