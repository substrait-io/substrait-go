// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestIntervalYearToMonthTypeToProto(t *testing.T) {
	allNullability := []types.Nullability{types.NullabilityUnspecified, types.NullabilityNullable, types.NullabilityRequired}

	for _, nullability := range allNullability {
		typ := types.NewIntervalYearToMonthType().WithTypeVariationRef(0).
			WithNullability(nullability).(types.IntervalYearToMonthType)
		assertIntervalYearToMonthTypeProto(t, nullability, typ)
	}
}

func assertIntervalYearToMonthTypeProto(t *testing.T, expectedNullability types.Nullability,
	toVerifyType types.IntervalYearToMonthType) {

	expectedTypeProto := &proto.Type{Kind: &proto.Type_IntervalYear_{
		IntervalYear: &proto.Type_IntervalYear{
			Nullability: proto.Type_Nullability(expectedNullability),
		},
	}}
	if diff := cmp.Diff(TypeToProto(toVerifyType), expectedTypeProto, protocmp.Transform()); diff != "" {
		t.Errorf("IntervalYearToMonthType proto didn't match, diff:\n%v", diff)
	}

	expectedFuncArgProto := &proto.FunctionArgument{ArgType: &proto.FunctionArgument_Type{
		Type: expectedTypeProto,
	}}
	if diff := cmp.Diff(FuncArgToProto(toVerifyType), expectedFuncArgProto, protocmp.Transform()); diff != "" {
		t.Errorf("IntervalYearToMonthType func arg proto didn't match, diff:\n%v", diff)
	}
}
