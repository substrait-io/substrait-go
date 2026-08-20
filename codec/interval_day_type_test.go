// SPDX-License-Identifier: Apache-2.0

package codec_test

import (
	"testing"

	"github.com/substrait-io/substrait-go/v8/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func TestIntervalDayTypeToProto(t *testing.T) {
	for _, precision := range allCodecTimePrecision {
		for _, nullability := range allCodecNullability {
			precisionProtoVal := precision.ToProtoVal()
			assertTypeProto(t, "IntervalDayType",
				(&types.IntervalDayType{Precision: precision}).WithNullability(nullability).(serializableType),
				&proto.Type{Kind: &proto.Type_IntervalDay_{
					IntervalDay: &proto.Type_IntervalDay{
						Precision:   &precisionProtoVal,
						Nullability: proto.Type_Nullability(nullability),
					},
				}})
		}
	}
}
