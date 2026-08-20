// SPDX-License-Identifier: Apache-2.0

package codec_test

import (
	"testing"

	"github.com/substrait-io/substrait-go/v8/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func TestPrecisionTimeTypeToProto(t *testing.T) {
	for _, precision := range allCodecTimePrecision {
		for _, nullability := range allCodecNullability {
			assertTypeProto(t, "PrecisionTimeType",
				types.NewPrecisionTimeType(precision).WithNullability(nullability).(serializableType),
				&proto.Type{Kind: &proto.Type_PrecisionTime_{
					PrecisionTime: &proto.Type_PrecisionTime{
						Precision:   precision.ToProtoVal(),
						Nullability: proto.Type_Nullability(nullability),
					},
				}})
		}
	}
}

func TestPrecisionTimestampTypeToProto(t *testing.T) {
	for _, precision := range allCodecTimePrecision {
		for _, nullability := range allCodecNullability {
			assertTypeProto(t, "PrecisionTimestampType",
				types.NewPrecisionTimestampType(precision).WithNullability(nullability).(serializableType),
				&proto.Type{Kind: &proto.Type_PrecisionTimestamp_{
					PrecisionTimestamp: &proto.Type_PrecisionTimestamp{
						Precision:   precision.ToProtoVal(),
						Nullability: proto.Type_Nullability(nullability),
					},
				}})
		}
	}
}

func TestPrecisionTimestampTzTypeToProto(t *testing.T) {
	for _, precision := range allCodecTimePrecision {
		for _, nullability := range allCodecNullability {
			assertTypeProto(t, "PrecisionTimestampTzType",
				types.NewPrecisionTimestampTzType(precision).WithNullability(nullability).(serializableType),
				&proto.Type{Kind: &proto.Type_PrecisionTimestampTz{
					PrecisionTimestampTz: &proto.Type_PrecisionTimestampTZ{
						Precision:   precision.ToProtoVal(),
						Nullability: proto.Type_Nullability(nullability),
					},
				}})
		}
	}
}
