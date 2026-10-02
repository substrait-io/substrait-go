// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func intervalCompoundTypeToProto(t types.IntervalCompoundType) *proto.Type {
	return &proto.Type{Kind: &proto.Type_IntervalCompound_{
		IntervalCompound: &proto.Type_IntervalCompound{
			Precision:              t.GetPrecisionProtoVal(),
			Nullability:            proto.Type_Nullability(t.GetNullability()),
			TypeVariationReference: t.GetTypeVariationReference()}}}
}
