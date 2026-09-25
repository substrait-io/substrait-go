// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func intervalDayTypeToProto(t *types.IntervalDayType) *proto.Type {
	precision := t.Precision.ToProtoVal()
	return &proto.Type{Kind: &proto.Type_IntervalDay_{
		IntervalDay: &proto.Type_IntervalDay{
			Precision:              &precision,
			Nullability:            proto.Type_Nullability(t.Nullability),
			TypeVariationReference: t.TypeVariationRef}}}
}
