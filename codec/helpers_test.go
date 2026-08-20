// SPDX-License-Identifier: Apache-2.0

package codec_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/substrait-io/substrait-go/v8/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/testing/protocmp"
)

var allCodecTimePrecision = []types.TimePrecision{types.PrecisionSeconds, types.PrecisionDeciSeconds, types.PrecisionCentiSeconds,
	types.PrecisionMilliSeconds, types.PrecisionEMinus4Seconds, types.PrecisionEMinus5Seconds, types.PrecisionMicroSeconds,
	types.PrecisionEMinus7Seconds, types.PrecisionEMinus8Seconds, types.PrecisionNanoSeconds}

var allCodecNullability = []types.Nullability{types.NullabilityUnspecified, types.NullabilityNullable, types.NullabilityRequired}

// serializableType is types.Type plus ToProto, which types.Type does not declare.
type serializableType interface {
	types.Type
	ToProto() *proto.Type
}

func assertTypeProto(t *testing.T, name string, toVerifyType serializableType, expectedTypeProto *proto.Type) {
	t.Helper()

	if diff := cmp.Diff(expectedTypeProto, toVerifyType.ToProto(), protocmp.Transform()); diff != "" {
		t.Errorf("%s proto didn't match, diff:\n%v", name, diff)
	}

	expectedFuncArgProto := &proto.FunctionArgument{ArgType: &proto.FunctionArgument_Type{
		Type: expectedTypeProto,
	}}
	if diff := cmp.Diff(expectedFuncArgProto, toVerifyType.ToProtoFuncArg(), protocmp.Transform()); diff != "" {
		t.Errorf("%s func arg proto didn't match, diff:\n%v", name, diff)
	}
}
