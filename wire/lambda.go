// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"github.com/substrait-io/substrait-go/v9/expr"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func lambdaToProto(l *expr.Lambda) *proto.Expression {
	children := make([]*proto.Type, len(l.Parameters.Types))
	for i, c := range l.Parameters.Types {
		children[i] = TypeToProto(c)
	}
	params := &proto.Type_Struct{
		Types:                  children,
		TypeVariationReference: l.Parameters.TypeVariationRef,
		Nullability:            proto.Type_Nullability(l.Parameters.Nullability),
	}

	return &proto.Expression{
		RexType: &proto.Expression_Lambda_{
			Lambda: &proto.Expression_Lambda{
				Parameters: params,
				Body:       ExprToProto(l.Body),
			},
		},
	}
}
