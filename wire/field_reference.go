// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"fmt"

	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// RefSegmentToProto encodes a reference segment as its protobuf message.
func RefSegmentToProto(r expr.ReferenceSegment) *proto.Expression_ReferenceSegment {
	switch r := r.(type) {
	case *expr.StructFieldRef:
		return structFieldRefToProto(r)
	default:
		panic(fmt.Sprintf("wire: unhandled reference segment %T", r))
	}
}

func structFieldRefToProto(r *expr.StructFieldRef) *proto.Expression_ReferenceSegment {
	var child *proto.Expression_ReferenceSegment
	if r.Child != nil {
		child = RefSegmentToProto(r.Child)
	}
	return &proto.Expression_ReferenceSegment{
		ReferenceType: &proto.Expression_ReferenceSegment_StructField_{
			StructField: &proto.Expression_ReferenceSegment_StructField{
				Field: r.Field,
				Child: child,
			},
		},
	}
}

// RefSegmentFromProto decodes a reference segment from its protobuf message.
func RefSegmentFromProto(p *proto.Expression_ReferenceSegment) expr.ReferenceSegment {
	if p == nil {
		return nil
	}

	switch seg := p.ReferenceType.(type) {
	case *proto.Expression_ReferenceSegment_StructField_:
		return &expr.StructFieldRef{
			Field: seg.StructField.Field,
			Child: RefSegmentFromProto(seg.StructField.Child),
		}
	}

	return nil
}
