package expr

import (
	"errors"
	"fmt"

	"github.com/substrait-io/substrait-go/proto"
	"github.com/substrait-io/substrait-go/types"
)

// IntervalCompoundLiteral creates an interval compound literal
type IntervalCompoundLiteral struct {
	years       int32
	months      int32
	days        int32
	seconds     int32
	subSeconds  int64
	precision   types.TimePrecision
	nullability types.Nullability
}

func (m IntervalCompoundLiteral) WithYears(years int32) IntervalCompoundLiteral {
	return IntervalCompoundLiteral{
		years:       years,
		months:      m.months,
		days:        m.days,
		seconds:     m.seconds,
		subSeconds:  m.subSeconds,
		precision:   m.precision,
		nullability: m.nullability,
	}
}

func (m IntervalCompoundLiteral) WithMonths(months int32) IntervalCompoundLiteral {
	return IntervalCompoundLiteral{
		years:       m.years,
		months:      months,
		days:        m.days,
		seconds:     m.seconds,
		subSeconds:  m.subSeconds,
		precision:   m.precision,
		nullability: m.nullability,
	}
}

func (m IntervalCompoundLiteral) WithDays(days int32) IntervalCompoundLiteral {
	return IntervalCompoundLiteral{
		years:       m.years,
		months:      m.months,
		days:        days,
		seconds:     m.seconds,
		subSeconds:  m.subSeconds,
		precision:   m.precision,
		nullability: m.nullability,
	}
}

func (m IntervalCompoundLiteral) WithSeconds(seconds int32) IntervalCompoundLiteral {
	return IntervalCompoundLiteral{
		years:       m.years,
		months:      m.months,
		days:        m.days,
		seconds:     seconds,
		subSeconds:  m.subSeconds,
		precision:   m.precision,
		nullability: m.nullability,
	}
}

func (m IntervalCompoundLiteral) WithMiliSecond(milliSeconds int64) IntervalCompoundLiteral {
	return IntervalCompoundLiteral{
		years:       m.years,
		months:      m.months,
		days:        m.days,
		seconds:     m.seconds,
		subSeconds:  milliSeconds,
		precision:   types.PrecisionMilliSeconds,
		nullability: m.nullability,
	}
}

func (m IntervalCompoundLiteral) WithMicroSecond(microSeconds int64) IntervalCompoundLiteral {
	return IntervalCompoundLiteral{
		years:       m.years,
		months:      m.months,
		days:        m.days,
		seconds:     m.seconds,
		subSeconds:  microSeconds,
		precision:   types.PrecisionMicroSeconds,
		nullability: m.nullability,
	}
}

func (m IntervalCompoundLiteral) WithNanoSecond(nanoSeconds int64) IntervalCompoundLiteral {
	return IntervalCompoundLiteral{
		years:       m.years,
		months:      m.months,
		days:        m.days,
		seconds:     m.seconds,
		subSeconds:  nanoSeconds,
		precision:   types.PrecisionNanoSeconds,
		nullability: m.nullability,
	}
}

func (m IntervalCompoundLiteral) WithSubSecond(subSeconds int64, precision types.TimePrecision) IntervalCompoundLiteral {
	return IntervalCompoundLiteral{
		years:       m.years,
		months:      m.months,
		days:        m.days,
		seconds:     m.seconds,
		subSeconds:  subSeconds,
		precision:   precision,
		nullability: m.nullability,
	}
}

func (m IntervalCompoundLiteral) WithNullability(nullability types.Nullability) IntervalCompoundLiteral {
	return IntervalCompoundLiteral{
		years:       m.years,
		months:      m.months,
		days:        m.days,
		seconds:     m.seconds,
		subSeconds:  m.subSeconds,
		precision:   m.precision,
		nullability: nullability,
	}
}

func (m IntervalCompoundLiteral) getType() types.Type {
	return types.NewIntervalCompoundType().WithPrecision(m.precision).WithNullability(m.nullability)
}

func (m IntervalCompoundLiteral) ToProtoLiteral() *proto.Expression_Literal {
	t := m.getType()
	intrCompPB := &proto.Expression_Literal_IntervalCompound{}

	if m.years != 0 || m.months != 0 {
		yearToMonthProto := &proto.Expression_Literal_IntervalYearToMonth{
			Years:  m.years,
			Months: m.months,
		}
		intrCompPB.IntervalYearToMonth = yearToMonthProto
	}

	if m.days != 0 || m.seconds != 0 || m.subSeconds != 0 {
		dayToSecondProto := &proto.Expression_Literal_IntervalDayToSecond{
			Days:          m.days,
			Seconds:       m.seconds,
			PrecisionMode: &proto.Expression_Literal_IntervalDayToSecond_Precision{Precision: m.precision.ToProtoVal()},
			Subseconds:    m.subSeconds,
		}
		intrCompPB.IntervalDayToSecond = dayToSecondProto
	}

	return &proto.Expression_Literal{
		LiteralType:            &proto.Expression_Literal_IntervalCompound_{IntervalCompound: intrCompPB},
		Nullable:               t.GetNullability() == types.NullabilityNullable,
		TypeVariationReference: t.GetTypeVariationReference(),
	}
}

func (m IntervalCompoundLiteral) ToProto() *proto.Expression {
	return &proto.Expression{RexType: &proto.Expression_Literal_{
		Literal: m.ToProtoLiteral(),
	}}
}

func intervalCompoundLiteralFromProto(l *proto.Expression_Literal) Literal {
	icLiteral := IntervalCompoundLiteral{}.WithNullability(getNullability(l.Nullable))
	yearToMonth := l.GetIntervalCompound().GetIntervalYearToMonth()
	if yearToMonth != nil {
		icLiteral = icLiteral.WithYears(yearToMonth.Years).WithMonths(yearToMonth.Months)
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
	return icLiteral.WithDays(dayToSecond.Days).WithSeconds(dayToSecond.Seconds).WithSubSecond(subSeconds, precision)
}

func (IntervalCompoundLiteral) isRootRef()            {}
func (m IntervalCompoundLiteral) GetType() types.Type { return m.getType() }
func (m IntervalCompoundLiteral) String() string {
	return fmt.Sprintf("%s(years:%d,months:%d, days:%d, seconds:%d subseconds:%d)",
		m.getType(), m.years, m.months, m.days, m.seconds, m.subSeconds)
}
func (m IntervalCompoundLiteral) Equals(rhs Expression) bool {
	if other, ok := rhs.(IntervalCompoundLiteral); ok {
		return m.getType().Equals(other.GetType()) && (m == other)
	}
	return false
}

func (m IntervalCompoundLiteral) ToProtoFuncArg() *proto.FunctionArgument {
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Value{Value: m.ToProto()},
	}
}

func (m IntervalCompoundLiteral) Visit(VisitFunc) Expression { return m }
func (IntervalCompoundLiteral) IsScalar() bool               { return true }

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
