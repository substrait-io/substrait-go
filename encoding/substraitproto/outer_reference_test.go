// SPDX-License-Identifier: Apache-2.0

package substraitproto

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	substraitgo "github.com/substrait-io/substrait-go/v9"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/testing/protocmp"
)

func outerFieldRefProto(outer *proto.Expression_FieldReference_OuterReference) *proto.Expression_FieldReference {
	return &proto.Expression_FieldReference{
		ReferenceType: &proto.Expression_FieldReference_DirectReference{
			DirectReference: &proto.Expression_ReferenceSegment{
				ReferenceType: &proto.Expression_ReferenceSegment_StructField_{
					StructField: &proto.Expression_ReferenceSegment_StructField{Field: 0}}}},
		RootType: &proto.Expression_FieldReference_OuterReference_{OuterReference: outer},
	}
}

func TestOuterReferenceStepsOutRoundTrip(t *testing.T) {
	reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
	want := outerFieldRefProto(&proto.Expression_FieldReference_OuterReference{
		OuterReferenceType: &proto.Expression_FieldReference_OuterReference_StepsOut{StepsOut: 2},
	})

	ref, err := FieldReferenceFromProto(want, nil, reg)
	require.NoError(t, err)
	assert.Equal(t, expr.OuterReference(2), ref.Root)

	if diff := cmp.Diff(want, fieldReferenceRefToProto(ref), protocmp.Transform()); diff != "" {
		t.Errorf("outer reference did not round trip, diff:\n%v", diff)
	}
}

func TestOuterReferenceByRelReferenceNotImplemented(t *testing.T) {
	reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
	p := outerFieldRefProto(&proto.Expression_FieldReference_OuterReference{
		OuterReferenceType: &proto.Expression_FieldReference_OuterReference_RelReference{RelReference: 1},
	})

	_, err := FieldReferenceFromProto(p, nil, reg)
	assert.ErrorIs(t, err, substraitgo.ErrNotImplemented)
}
