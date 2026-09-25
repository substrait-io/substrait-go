// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func precisionTimeTypeToProto(t *types.PrecisionTimeType) *proto.Type {
	return &proto.Type{Kind: &proto.Type_PrecisionTime_{
		PrecisionTime: &proto.Type_PrecisionTime{
			Precision:              int32(t.Precision),
			Nullability:            proto.Type_Nullability(t.Nullability),
			TypeVariationReference: t.TypeVariationRef}}}
}
