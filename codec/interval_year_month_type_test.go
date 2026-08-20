// SPDX-License-Identifier: Apache-2.0

package codec_test

import (
	"testing"

	"github.com/substrait-io/substrait-go/v8/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func TestIntervalYearToMonthTypeToProto(t *testing.T) {
	for _, nullability := range allCodecNullability {
		assertTypeProto(t, "IntervalYearToMonthType",
			types.NewIntervalYearToMonthType().WithTypeVariationRef(0).WithNullability(nullability).(serializableType),
			&proto.Type{Kind: &proto.Type_IntervalYear_{
				IntervalYear: &proto.Type_IntervalYear{
					Nullability: proto.Type_Nullability(nullability),
				},
			}})
	}
}
