// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func intervalYearToMonthLiteralFromProto(l *proto.Expression_Literal) expr.Literal {
	return expr.IntervalYearToMonthLiteral{
		Years:       l.GetIntervalYearToMonth().Years,
		Months:      l.GetIntervalYearToMonth().Months,
		Nullability: nullabilityFromBool(l.Nullable),
	}
}

func intervalYearToMonthLiteralToProto(m expr.IntervalYearToMonthLiteral) *proto.Expression_Literal {
	t := types.NewIntervalYearToMonthType().WithNullability(m.Nullability)
	return &proto.Expression_Literal{
		LiteralType: &proto.Expression_Literal_IntervalYearToMonth_{
			IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{
				Years:  m.Years,
				Months: m.Months,
			},
		},
		Nullable:               t.GetNullability() == types.NullabilityNullable,
		TypeVariationReference: t.GetTypeVariationReference(),
	}
}
