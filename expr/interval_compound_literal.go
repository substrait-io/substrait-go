package expr

import (
	"errors"
	"github.com/substrait-io/substrait-go/proto"
	"github.com/substrait-io/substrait-go/types"
)

type intervalDateParts struct {
	years      int32
	months     int32
	days       int32
	seconds    int32
	subSeconds int64
}

type intervalDatePartsOptions func(parts *intervalDateParts)

func WithIntervalCompoundYears(years int32) func(*intervalDateParts) {
	return func(idp *intervalDateParts) {
		idp.years = years
	}
}

func WithIntervalCompoundMonths(months int32) func(*intervalDateParts) {
	return func(idp *intervalDateParts) {
		idp.months = months
	}
}

func WithIntervalCompoundDays(days int32) func(*intervalDateParts) {
	return func(idp *intervalDateParts) {
		idp.days = days
	}
}

func WithIntervalCompoundSeconds(seconds int32) func(*intervalDateParts) {
	return func(idp *intervalDateParts) {
		idp.seconds = seconds
	}
}

func WithIntervalCompoundSubSeconds(subSeconds int64) func(*intervalDateParts) {
	return func(idp *intervalDateParts) {
		idp.subSeconds = subSeconds
	}
}

// NewIntervalLiteralUptoSubSecondPrecision creates an interval literal which allows upto subsecond precision
// arguments: precision and nullable property (n)
// datePartsOptions is options to set value parts (month, year, day, seconds, subseconds).
// If multiple options of same types (e.g. multiple second options) are provided only value of last part is considered
func NewIntervalLiteralUptoSubSecondPrecision(precision types.TimePrecision, n types.Nullability, datePartsOptions ...intervalDatePartsOptions) Literal {
	intervalCompoundType := types.NewIntervalCompoundType(precision).WithNullability(n)
	intervalPartsVal := newIntervalPartsValInternal(datePartsOptions...)
	return &ProtoLiteral{
		Value: intervalPartsVal,
		Type:  intervalCompoundType,
	}
}

func newIntervalPartsValInternal(datePartsOptions ...intervalDatePartsOptions) *intervalDateParts {
	intervalPartsVal := &intervalDateParts{}
	for _, datePart := range datePartsOptions {
		datePart(intervalPartsVal)
	}
	return intervalPartsVal
}

func intervalPartsValToProto(idp *intervalDateParts, ict *types.IntervalCompoundType) *proto.Expression_Literal_IntervalCompound {
	intrCompPB := &proto.Expression_Literal_IntervalCompound{}

	if idp.years > 0 || idp.months > 0 {
		yearToMonthProto := &proto.Expression_Literal_IntervalYearToMonth{
			Years:  idp.years,
			Months: idp.months,
		}
		intrCompPB.IntervalYearToMonth = yearToMonthProto
	}
	if idp.days == 0 && idp.seconds == 0 && idp.subSeconds == 0 {
		// all parts are >= month granularity so no need to set daytosecond component
		return intrCompPB
	}

	dayToSecondProto := &proto.Expression_Literal_IntervalDayToSecond{
		Days:          idp.days,
		Seconds:       idp.seconds,
		PrecisionMode: &proto.Expression_Literal_IntervalDayToSecond_Precision{Precision: ict.GetPrecisionProtoVal()},
		Subseconds:    idp.subSeconds,
	}
	intrCompPB.IntervalDayToSecond = dayToSecondProto
	return intrCompPB
}

func intervalCompoundLiteralFromProto(protoVal *proto.Expression_Literal_IntervalCompound, nullability types.Nullability) Literal {
	var datePartsOptions []intervalDatePartsOptions
	if protoVal.IntervalYearToMonth != nil {
		if protoVal.IntervalYearToMonth.Years > 0 {
			datePartsOptions = append(datePartsOptions, WithIntervalCompoundYears(protoVal.IntervalYearToMonth.Years))
		}
		if protoVal.IntervalYearToMonth.Months > 0 {
			datePartsOptions = append(datePartsOptions, WithIntervalCompoundMonths(protoVal.IntervalYearToMonth.Months))
		}
	}
	idts := protoVal.IntervalDayToSecond
	err := validateIntervalDayToSecondProto(idts)
	if err != nil {
		return nil
	}
	precision, err := intervalCompoundPrecisionFromProto(idts)
	if err != nil {
		return nil
	}
	if idts.Days > 0 {
		datePartsOptions = append(datePartsOptions, WithIntervalCompoundDays(protoVal.IntervalDayToSecond.Days))
	}
	if idts.Seconds > 0 {
		datePartsOptions = append(datePartsOptions, WithIntervalCompoundSeconds(protoVal.IntervalDayToSecond.Seconds))
	}
	if idts.Subseconds > 0 {
		datePartsOptions = append(datePartsOptions, WithIntervalCompoundSubSeconds(protoVal.IntervalDayToSecond.Subseconds))
	} else if val, ok := idts.PrecisionMode.(*proto.Expression_Literal_IntervalDayToSecond_Microseconds); ok {
		// deprecated field microsecond is set, set its value as subsecond
		datePartsOptions = append(datePartsOptions, WithIntervalCompoundSubSeconds(int64(val.Microseconds)))
	}
	return NewIntervalLiteralUptoSubSecondPrecision(precision, nullability, datePartsOptions...)
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

func intervalCompoundPrecisionFromProto(protoVal *proto.Expression_Literal_IntervalDayToSecond) (types.TimePrecision, error) {
	var precisionVal int32
	switch pmt := protoVal.PrecisionMode.(type) {
	case *proto.Expression_Literal_IntervalDayToSecond_Precision:
		precisionVal = pmt.Precision
	case *proto.Expression_Literal_IntervalDayToSecond_Microseconds:
		precisionVal = types.PrecisionMicroSeconds.ToProtoVal()
	}
	precision, err := types.ProtoToTimePrecision(precisionVal)
	if err != nil {
		return types.PrecisionUnknown, err
	}
	return precision, nil
}
