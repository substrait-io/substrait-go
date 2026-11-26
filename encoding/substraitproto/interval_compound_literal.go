// SPDX-License-Identifier: Apache-2.0

package substraitproto

import (
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func intervalCompoundLiteralFromProto(l *proto.Expression_Literal) expr.Literal {
	icLiteral := expr.IntervalCompoundLiteral{Nullability: nullabilityFromBool(l.Nullable)}
	yearToMonth := l.GetIntervalCompound().GetIntervalYearToMonth()
	if yearToMonth != nil {
		icLiteral.Years = yearToMonth.Years
		icLiteral.Months = yearToMonth.Months
	}
	dayToSecond := l.GetIntervalCompound().GetIntervalDayToSecond()
	if dayToSecond == nil {
		// no day to second part
		return icLiteral
	}
	precision, err := types.ProtoToTimePrecision(dayToSecond.Precision)
	if err != nil {
		return nil
	}
	icLiteral.Days = dayToSecond.Days
	icLiteral.Seconds = dayToSecond.Seconds
	icLiteral.SubSeconds = dayToSecond.Subseconds
	icLiteral.SubSecondPrecision = precision
	return icLiteral
}

func intervalCompoundLiteralToProto(m expr.IntervalCompoundLiteral) *proto.Expression_Literal {
	t := types.NewIntervalCompoundType().WithPrecision(m.SubSecondPrecision).WithNullability(m.Nullability)
	intrCompPB := &proto.Expression_Literal_IntervalCompound{
		IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{
			Years:  m.Years,
			Months: m.Months,
		},
		IntervalDayToSecond: &proto.Expression_Literal_IntervalDayToSecond{
			Days:       m.Days,
			Seconds:    m.Seconds,
			Precision:  m.SubSecondPrecision.ToProtoVal(),
			Subseconds: m.SubSeconds,
		},
	}

	return &proto.Expression_Literal{
		LiteralType:            &proto.Expression_Literal_IntervalCompound_{IntervalCompound: intrCompPB},
		Nullable:               t.GetNullability() == types.NullabilityNullable,
		TypeVariationReference: t.GetTypeVariationReference(),
	}
}
