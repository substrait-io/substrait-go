package literal

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/substrait-io/substrait-go/v3/expr"
	"github.com/substrait-io/substrait-go/v3/proto"
	"github.com/substrait-io/substrait-go/v3/types"
)

func NewBool(value bool) expr.Literal {
	return expr.NewPrimitiveLiteral[bool](value, false)
}

func NewInt8(value int8) expr.Literal {
	return expr.NewPrimitiveLiteral[int8](value, false)
}

func NewInt16(value int16) expr.Literal {
	return expr.NewPrimitiveLiteral[int16](value, false)
}

func NewInt32(value int32) expr.Literal {
	return expr.NewPrimitiveLiteral[int32](value, false)
}

func NewInt64(value int64) expr.Literal {
	return expr.NewPrimitiveLiteral[int64](value, false)
}

func NewFloat32(value float32) expr.Literal {
	return expr.NewPrimitiveLiteral[float32](value, false)
}

func NewFloat64(value float64) expr.Literal {
	return expr.NewPrimitiveLiteral[float64](value, false)
}

func NewString(value string) expr.Literal {
	return expr.NewPrimitiveLiteral[string](value, false)
}

func NewDate(days int) (expr.Literal, error) {
	return expr.NewLiteral[types.Date](types.Date(days), false)
}

func NewDateFromString(value string) (expr.Literal, error) {
	tm, err := time.Parse("2006-01-02", value)
	if err != nil {
		return nil, err
	}
	return NewDate(int(tm.Unix() / 86400))
}

// NewTime creates a new Time literal from the given hours, minutes, seconds and microseconds.
// The total microseconds should be in the range [0, 86400_000_000) to represent a valid time within a day.
func NewTime(hours, minutes, seconds, microseconds int32) (expr.Literal, error) {
	duration := time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute + time.Duration(seconds)*time.Second + time.Duration(microseconds)*time.Microsecond
	micros := duration.Microseconds()
	if micros < 0 || micros >= (24*time.Hour).Microseconds() {
		return nil, fmt.Errorf("invalid time value %d:%d:%d.%d", hours, minutes, seconds, microseconds)
	}
	return expr.NewLiteral[types.Time](types.Time(duration.Microseconds()), false)
}

// NewTimeFromMicros creates a new Time literal from the given microseconds.
func NewTimeFromMicros(micros int64) (expr.Literal, error) {
	if micros < 0 || micros >= (24*time.Hour).Microseconds() {
		return nil, fmt.Errorf("invalid time value %d", micros)
	}
	return expr.NewLiteral[types.Time](types.Time(micros), false)
}

func NewTimeFromString(value string) (expr.Literal, error) {
	ts, err := parseTimeFromString(value)
	if err != nil {
		return nil, err
	}
	seconds := ts.Hour()*3600 + ts.Minute()*60 + ts.Second()
	micros := int64(seconds)*int64(1e6) + int64(ts.Nanosecond())/1e3
	return NewTimeFromMicros(micros)
}

func parseTimeFromString(value string) (time.Time, error) {
	if t, err := time.Parse(time.RFC3339, value); err == nil {
		return t, nil
	}

	layoutWithoutOffset := "2006-01-02T15:04:05"
	if t, err := time.Parse(layoutWithoutOffset, value); err == nil {
		return t, nil
	}
	timeOnlyInMicros := "15:04:05.9999999"
	if t, err := time.Parse(timeOnlyInMicros, value); err == nil {
		return t, nil
	}
	timeOnlyInMillis := "15:04:05.999"
	return time.Parse(timeOnlyInMillis, value)
}

// NewTimestamp creates a new Timestamp literal from a time.Time timestamp value.
// This uses the number of microseconds elapsed since January 1, 1970 00:00:00 UTC
func NewTimestamp(timestamp time.Time) (expr.Literal, error) {
	return expr.NewLiteral[types.Timestamp](types.Timestamp(timestamp.UnixMicro()), false)
}

func NewTimestampFromString(value string) (expr.Literal, error) {
	tm, err := parseTimeFromString(value)
	if err != nil {
		return nil, err
	}
	return NewTimestamp(tm)
}

func NewTimestampFromMicros(micros int64) (expr.Literal, error) {
	return expr.NewLiteral[types.Timestamp](types.Timestamp(micros), false)
}

// NewTimestampTZ creates a new TimestampTz literal from a time.Time timestamp value.
// This uses the number of microseconds elapsed since January 1, 1970 00:00:00 UTC
func NewTimestampTZ(timestamp time.Time) (expr.Literal, error) {
	return expr.NewLiteral[types.TimestampTz](types.TimestampTz(timestamp.UnixMicro()), false)
}

func NewTimestampTZFromString(value string) (expr.Literal, error) {
	tm, err := parseTimeFromString(value)
	if err != nil {
		return nil, err
	}
	return NewTimestampTZ(tm)
}

func NewTimestampTZFromMicros(micros int64) (expr.Literal, error) {
	return expr.NewLiteral[types.TimestampTz](types.TimestampTz(micros), false)
}

func NewIntervalYearsToMonthFromString(yearsToMonth string) (expr.Literal, error) {
	years, months, err := parseIntervalYearsToMonth(yearsToMonth)
	if err != nil {
		return nil, err
	}
	return NewIntervalYearsToMonth(years, months)
}

func parseIntervalYearsToMonth(interval string) (int32, int32, error) {
	if len(interval) < 3 || interval[0] != 'P' {
		return 0, 0, fmt.Errorf("invalid interval format: %s", interval)
	}
	interval = interval[1:]
	yIndex := -1
	mIndex := -1
	for i, c := range interval {
		if c == 'Y' {
			yIndex = i
		} else if c == 'M' {
			mIndex = i
		}
	}
	if yIndex == -1 && mIndex == -1 {
		return 0, 0, fmt.Errorf("invalid interval format: %s", interval)
	}
	var months, years int
	var err error
	if yIndex != -1 {
		years, err = strconv.Atoi(interval[:yIndex])
		if err != nil {
			return 0, 0, err
		}
		interval = interval[yIndex+1:]
		mIndex -= yIndex + 1
	}
	if mIndex > 0 {
		months, err = strconv.Atoi(interval[:mIndex])
		if err != nil {
			return 0, 0, err
		}
	}
	return int32(years), int32(months), nil
}

func NewIntervalYearsToMonth(years, months int32) (expr.Literal, error) {
	return expr.NewLiteral[*types.IntervalYearToMonth](&types.IntervalYearToMonth{Years: years, Months: months}, false)
}

func NewIntervalDaysToSecondFromString(daysToSecond string) (expr.Literal, error) {
	days, seconds, subSeconds, precision, err := parseIntervalDaysToSecond(daysToSecond)
	if err != nil {
		return nil, err
	}
	return expr.NewLiteral[*types.IntervalDayToSecond](&types.IntervalDayToSecond{
		Days:    days,
		Seconds: seconds,
		PrecisionMode: &proto.Expression_Literal_IntervalDayToSecond_Precision{
			Precision: precision,
		},
		Subseconds: subSeconds,
	}, false)
}

func parseIntervalDaysToSecond(interval string) (int32, int32, int64, int32, error) {
	if len(interval) < 3 || interval[0] != 'P' {
		return 0, 0, 0, 0, fmt.Errorf("invalid interval format: %s", interval)
	}

	// Parse interval of format P[n]DT[n]H[n]M[n]S
	// Ex: 3DT4H5M6.789S
	regex := `^P(?:(\d+)D)?(?:T(?:(\d+)H)?(?:(\d+)M)?(?:(\d*(\.\d+)?)S)?)?$`
	r := regexp.MustCompile(regex)

	// Find matches
	matches := r.FindStringSubmatch(interval)
	if matches == nil {
		return 0, 0, 0, 0, fmt.Errorf("invalid interval format: %s", interval)
	}

	// Parse each component
	var err error
	var days, hours, minutes, seconds int
	var secFloat float64
	var subSeconds int64
	if matches[1] != "" {
		days, err = strconv.Atoi(matches[1])
		if err != nil {
			return 0, 0, 0, 0, fmt.Errorf("invalid day value: %v", err)
		}
	}
	if matches[2] != "" {
		hours, err = strconv.Atoi(matches[2])
		if err != nil {
			return 0, 0, 0, 0, fmt.Errorf("invalid hour value: %v", err)
		}
	}
	if matches[3] != "" {
		minutes, err = strconv.Atoi(matches[3])
		if err != nil {
			return 0, 0, 0, 0, fmt.Errorf("invalid minute value: %v", err)
		}
	}
	if matches[4] != "" {
		secFloat, err = strconv.ParseFloat(matches[4], 64)
		if err != nil {
			return 0, 0, 0, 0, fmt.Errorf("invalid fractional second value: %v", err)
		}
	}

	seconds = int(secFloat)
	secFloat -= float64(seconds)
	seconds += hours*3600 + minutes*60
	nanoSeconds := int64(secFloat * 1e9)
	subSeconds = int64(secFloat * 1e6)
	precision := int32(types.PrecisionMicroSeconds)
	if nanoSeconds > subSeconds*1e3 {
		subSeconds = nanoSeconds
		precision = int32(types.PrecisionNanoSeconds)
	}
	return int32(days), int32(seconds), subSeconds, precision, nil
}

func NewIntervalDaysToSecond(days, seconds int32, micros int64) (expr.Literal, error) {
	return expr.NewLiteral[*types.IntervalDayToSecond](&types.IntervalDayToSecond{
		Days:    days,
		Seconds: seconds,
		PrecisionMode: &proto.Expression_Literal_IntervalDayToSecond_Precision{
			Precision: int32(types.PrecisionMicroSeconds),
		},
		Subseconds: micros,
	}, false)
}

func NewUUID(guid uuid.UUID) (expr.Literal, error) {
	bytes, err := guid.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return expr.NewLiteral[types.UUID](bytes, false)
}

func NewUUIDFromBytes(value []byte) (expr.Literal, error) {
	return expr.NewLiteral[types.UUID](value, false)
}

func NewFixedChar(value string) (expr.Literal, error) {
	return expr.NewLiteral[types.FixedChar](types.FixedChar(value), false)
}

func NewFixedBinary(value []byte) (expr.Literal, error) {
	return expr.NewLiteral[types.FixedBinary](value, false)
}

func NewVarChar(value string) (expr.Literal, error) {
	return expr.NewLiteral[*types.VarChar](&types.VarChar{Value: value, Length: uint32(len(value))}, false)
}

// NewDecimalFromTwosComplement create a Decimal literal from twosComplement.
// twosComplement is a little-endian twos-complement integer representation of complete value
func NewDecimalFromTwosComplement(twosComplement []byte, precision, scale int32) (expr.Literal, error) {
	if len(twosComplement) != 16 {
		return nil, fmt.Errorf("twosComplement must be 16 bytes")
	}
	if precision < 1 || precision > 38 {
		return nil, fmt.Errorf("precision must be in range [1, 38]")
	}
	if scale < 0 || scale > precision {
		return nil, fmt.Errorf("scale must be in range [0, precision]")
	}
	return expr.NewLiteral[*types.Decimal](&types.Decimal{Value: twosComplement, Precision: precision, Scale: scale}, false)

}

// NewDecimalFromString create a Decimal literal from decimal value string
func NewDecimalFromString(value string) (expr.Literal, error) {
	v, precision, scale, err := decimalStringToBytes(value)
	if err != nil {
		return nil, err
	}
	return expr.NewLiteral[*types.Decimal](&types.Decimal{Value: v[:16], Precision: precision, Scale: scale}, false)
}

// NewPrecisionTimestampFromTime creates a new PrecisionTimestamp literal from a time.Time timestamp value with given precision.
func NewPrecisionTimestampFromTime(precision types.TimePrecision, tm time.Time) (expr.Literal, error) {
	return NewPrecisionTimestamp(precision, getTimeValueByPrecision(tm, precision))
}

// NewPrecisionTimestamp creates a new PrecisionTimestamp literal with given precision and value.
func NewPrecisionTimestamp(precision types.TimePrecision, value int64) (expr.Literal, error) {
	return expr.NewLiteral[*types.PrecisionTimestamp](&types.PrecisionTimestamp{
		PrecisionTimestamp: &proto.Expression_Literal_PrecisionTimestamp{
			Precision: int32(precision),
			Value:     value,
		},
	}, false)
}

// NewPrecisionTimestampTzFromTime creates a new PrecisionTimestampTz literal from a time.Time timestamp value with given precision.
func NewPrecisionTimestampTzFromTime(precision types.TimePrecision, tm time.Time) (expr.Literal, error) {
	return NewPrecisionTimestampTz(precision, getTimeValueByPrecision(tm, precision))
}

// NewPrecisionTimestampTz creates a new PrecisionTimestampTz literal with given precision and value.
func NewPrecisionTimestampTz(precision types.TimePrecision, value int64) (expr.Literal, error) {
	return expr.NewLiteral[*types.PrecisionTimestampTz](&types.PrecisionTimestampTz{
		PrecisionTimestampTz: &proto.Expression_Literal_PrecisionTimestamp{
			Precision: int32(precision),
			Value:     value,
		},
	}, false)
}

func getTimeValueByPrecision(tm time.Time, precision types.TimePrecision) int64 {
	switch precision {
	case types.PrecisionSeconds:
		return tm.Unix()
	case types.PrecisionDeciSeconds:
		return tm.UnixMilli() / 100
	case types.PrecisionCentiSeconds:
		return tm.UnixMilli() / 10
	case types.PrecisionMilliSeconds:
		return tm.UnixMilli()
	case types.PrecisionEMinus4Seconds:
		return tm.UnixMicro() / 100
	case types.PrecisionEMinus5Seconds:
		return tm.UnixMicro() / 10
	case types.PrecisionMicroSeconds:
		return tm.UnixMicro()
	case types.PrecisionEMinus7Seconds:
		return tm.UnixNano() / 100
	case types.PrecisionEMinus8Seconds:
		return tm.UnixNano() / 10
	case types.PrecisionNanoSeconds:
		return tm.UnixNano()
	default:
		panic(fmt.Sprintf("unknown TimePrecision %v", precision))
	}
}

func NewList(elements []expr.Literal) (expr.Literal, error) {
	if len(elements) == 0 {
		return nil, fmt.Errorf("empty list literal")
	}
	anchorType := reflect.TypeOf(elements[0])
	for i, e := range elements {
		currentType := reflect.TypeOf(e)
		if currentType != anchorType {
			nullLiteralType := reflect.TypeOf((*expr.NullLiteral)(nil))
			if currentType == nullLiteralType {
				continue
			}
			if anchorType == nullLiteralType {
				anchorType = currentType
				continue
			}
			return nil, fmt.Errorf("element %d of list literal has different type", i)
		}
	}
	return expr.NewLiteral[expr.ListLiteralValue](elements, false)
}
