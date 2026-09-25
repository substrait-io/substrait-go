// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestIntervalCompoundTypeToProto(t *testing.T) {
	allPrecision := []types.TimePrecision{
		types.PrecisionSeconds, types.PrecisionDeciSeconds, types.PrecisionCentiSeconds,
		types.PrecisionMilliSeconds, types.PrecisionEMinus4Seconds, types.PrecisionEMinus5Seconds,
		types.PrecisionMicroSeconds, types.PrecisionEMinus7Seconds, types.PrecisionEMinus8Seconds,
		types.PrecisionNanoSeconds,
	}
	allNullability := []types.Nullability{types.NullabilityUnspecified, types.NullabilityNullable, types.NullabilityRequired}

	for _, precision := range allPrecision {
		for _, nullability := range allNullability {
			typ := types.NewIntervalCompoundType().WithPrecision(precision).WithTypeVariationRef(0).
				WithNullability(nullability).(types.IntervalCompoundType)
			assertIntervalCompoundTypeProto(t, precision, nullability, typ)
		}
	}
}

func assertIntervalCompoundTypeProto(t *testing.T, expectedPrecision types.TimePrecision,
	expectedNullability types.Nullability, toVerifyType types.IntervalCompoundType) {

	expectedTypeProto := &proto.Type{Kind: &proto.Type_IntervalCompound_{
		IntervalCompound: &proto.Type_IntervalCompound{
			Precision:   expectedPrecision.ToProtoVal(),
			Nullability: proto.Type_Nullability(expectedNullability),
		},
	}}
	if diff := cmp.Diff(TypeToProto(toVerifyType), expectedTypeProto, protocmp.Transform()); diff != "" {
		t.Errorf("IntervalCompoundType proto didn't match, diff:\n%v", diff)
	}

	expectedFuncArgProto := &proto.FunctionArgument{ArgType: &proto.FunctionArgument_Type{
		Type: expectedTypeProto,
	}}
	if diff := cmp.Diff(FuncArgToProto(toVerifyType), expectedFuncArgProto, protocmp.Transform()); diff != "" {
		t.Errorf("IntervalCompoundType func arg proto didn't match, diff:\n%v", diff)
	}
}
