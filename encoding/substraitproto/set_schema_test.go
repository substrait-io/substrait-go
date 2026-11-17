// SPDX-License-Identifier: Apache-2.0

package substraitproto_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/encoding/substraitproto"
	substraitgo "github.com/substrait-io/substrait-go/v9"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/plan"
	"github.com/substrait-io/substrait-go/v9/types"
)

func setInput(fields ...types.Type) plan.Rel {
	names := make([]string, len(fields))
	for i := range names {
		names[i] = "field"
	}
	return plan.NewBuilderDefault().NamedScan([]string{"table"}, types.NamedStruct{
		Names:  names,
		Struct: types.StructType{Nullability: types.NullabilityRequired, Types: fields},
	})
}

func TestSetNullabilityRoundTrip(t *testing.T) {
	r := types.NullabilityRequired
	n := types.NullabilityNullable
	for _, tc := range []struct {
		name   string
		op     plan.SetOp
		inputs []types.Nullability
		want   types.Nullability
	}{
		{"union required nullable", plan.SetOpUnionAll, []types.Nullability{r, n}, n},
		{"union nullable required", plan.SetOpUnionAll, []types.Nullability{n, r}, n},
		{"union required control", plan.SetOpUnionDistinct, []types.Nullability{r, r}, r},
		{"minus keeps primary", plan.SetOpMinusPrimary, []types.Nullability{n, r}, n},
		{"intersection primary", plan.SetOpIntersectionPrimary, []types.Nullability{n, r, n}, n},
		{"intersection multiset", plan.SetOpIntersectionMultiset, []types.Nullability{n, r, n}, r},
	} {
		t.Run(tc.name, func(t *testing.T) {
			inputs := make([]plan.Rel, len(tc.inputs))
			for i, nullability := range tc.inputs {
				inputs[i] = setInput(&types.Int64Type{Nullability: nullability}, &types.StringType{Nullability: r})
			}
			original := plan.NewSetRel(inputs, tc.op, plan.RelCommon{}, nil)
			remapped, err := original.Remap(1, 0)
			require.NoError(t, err)
			pb := substraitproto.RelToProto(remapped)
			reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
			decoded, err := substraitproto.RelFromProto(pb, reg)
			require.NoError(t, err)
			assert.Equal(t, []types.Type{&types.StringType{Nullability: r}, &types.Int64Type{Nullability: tc.want}}, decoded.RecordType().Types())
			assert.Equal(t, []int32{1, 0}, decoded.(*plan.SetRel).OutputMapping())
			assert.Equal(t, tc.op, decoded.(*plan.SetRel).Op())
			assert.Equal(t, pb, substraitproto.RelToProto(decoded))
		})
	}
}

func TestSetRejectsDifferentFieldTypes(t *testing.T) {
	r := types.NullabilityRequired
	n := types.NullabilityNullable
	left := setInput(&types.Int64Type{Nullability: r})
	for _, right := range []plan.Rel{
		setInput(&types.Int32Type{Nullability: n}),
		setInput(),
	} {
		original := plan.NewSetRel([]plan.Rel{left, right}, plan.SetOpUnionAll, plan.RelCommon{}, nil)
		reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
		_, err := substraitproto.RelFromProto(substraitproto.RelToProto(original), reg)
		assert.ErrorIs(t, err, substraitgo.ErrInvalidRel)
	}
}
