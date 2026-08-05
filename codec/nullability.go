// SPDX-License-Identifier: Apache-2.0

package codec

import (
	"github.com/substrait-io/substrait-go/v8/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// NullabilityToProto encodes a domain nullability as its protobuf enum. The domain type carries the
// spec's own wire numbers over the same int32 range, so the mapping is total.
func NullabilityToProto(n types.Nullability) proto.Type_Nullability {
	return proto.Type_Nullability(n)
}

// NullabilityFromProto decodes a protobuf nullability enum. Protobuf enums are open, so a number
// the spec has not defined yet is carried through rather than rejected or defaulted.
func NullabilityFromProto(p proto.Type_Nullability) types.Nullability {
	return types.Nullability(p)
}
