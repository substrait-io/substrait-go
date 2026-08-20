// SPDX-License-Identifier: Apache-2.0

package codec_test

import (
	"testing"

	"github.com/substrait-io/substrait-go/v8/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func TestIntervalCompoundTypeToProto(t *testing.T) {
	for _, precision := range allCodecTimePrecision {
		for _, nullability := range allCodecNullability {
			assertTypeProto(t, "IntervalCompoundType",
				types.NewIntervalCompoundType().WithPrecision(precision).WithTypeVariationRef(0).WithNullability(nullability).(serializableType),
				&proto.Type{Kind: &proto.Type_IntervalCompound_{
					IntervalCompound: &proto.Type_IntervalCompound{
						Precision:   precision.ToProtoVal(),
						Nullability: proto.Type_Nullability(nullability),
					},
				}})
		}
	}
}
