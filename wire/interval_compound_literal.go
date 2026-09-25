// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"errors"

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
	err := validateIntervalDayToSecondProto(dayToSecond)
	if err != nil {
		return nil
	}

	// get subSecond/precision value from proto. To get value it takes care of deprecated microseconds
	precision, subSeconds, err := intervalCompoundPrecisionSubSecondsFromProto(dayToSecond)
	if err != nil {
		return nil
	}
	icLiteral.Days = dayToSecond.Days
	icLiteral.Seconds = dayToSecond.Seconds
	icLiteral.SubSeconds = subSeconds
	icLiteral.SubSecondPrecision = precision
	return icLiteral
}

func validateIntervalDayToSecondProto(idts *proto.Expression_Literal_IntervalDayToSecond) error {
	if idts.PrecisionMode == nil {
		// error, precision mode must be set for intervalCompound
		return errors.New("missing precision mode for interval compound")
	}
	if _, ok := idts.PrecisionMode.(*proto.Expression_Literal_IntervalDayToSecond_Microseconds); ok {
		// if microsecond precision then subseconds must be set to zero
		if idts.Subseconds > 0 {
			return errors.New("both deprecated microseconds and subseconds can't be non zero")
		}
	}
	return nil
}

func intervalCompoundPrecisionSubSecondsFromProto(protoVal *proto.Expression_Literal_IntervalDayToSecond) (types.TimePrecision, int64, error) {
	var precisionVal int32
	var subSecondVal int64
	switch pmt := protoVal.PrecisionMode.(type) {
	case *proto.Expression_Literal_IntervalDayToSecond_Precision:
		precisionVal = pmt.Precision
		subSecondVal = protoVal.Subseconds
	case *proto.Expression_Literal_IntervalDayToSecond_Microseconds:
		// deprecated field microsecond is set, treat its value subsecond
		precisionVal = types.PrecisionMicroSeconds.ToProtoVal()
		subSecondVal = int64(pmt.Microseconds)
	}
	precision, err := types.ProtoToTimePrecision(precisionVal)
	if err != nil {
		return types.PrecisionUnknown, 0, err
	}
	return precision, subSecondVal, nil
}

func intervalCompoundLiteralToProto(m expr.IntervalCompoundLiteral) *proto.Expression_Literal {
	t := types.NewIntervalCompoundType().WithPrecision(m.SubSecondPrecision).WithNullability(m.Nullability)
	intrCompPB := &proto.Expression_Literal_IntervalCompound{}

	if m.Years != 0 || m.Months != 0 {
		intrCompPB.IntervalYearToMonth = &proto.Expression_Literal_IntervalYearToMonth{
			Years:  m.Years,
			Months: m.Months,
		}
	}

	if m.Days != 0 || m.Seconds != 0 || m.SubSeconds != 0 {
		intrCompPB.IntervalDayToSecond = &proto.Expression_Literal_IntervalDayToSecond{
			Days:          m.Days,
			Seconds:       m.Seconds,
			PrecisionMode: &proto.Expression_Literal_IntervalDayToSecond_Precision{Precision: m.SubSecondPrecision.ToProtoVal()},
			Subseconds:    m.SubSeconds,
		}
	}

	return &proto.Expression_Literal{
		LiteralType:            &proto.Expression_Literal_IntervalCompound_{IntervalCompound: intrCompPB},
		Nullable:               t.GetNullability() == types.NullabilityNullable,
		TypeVariationReference: t.GetTypeVariationReference(),
	}
}
