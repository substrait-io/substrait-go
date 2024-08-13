package literal

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/substrait-io/substrait-go/expr"
	"github.com/substrait-io/substrait-go/types"
)

func NewBool(value bool) (expr.Literal, error) {
	return expr.NewPrimitiveLiteral[bool](value, false), nil
}

func NewInt8(value int8) (expr.Literal, error) {
	return expr.NewPrimitiveLiteral[int8](value, false), nil
}

func NewInt16(value int16) (expr.Literal, error) {
	return expr.NewPrimitiveLiteral[int16](value, false), nil
}

func NewInt32(value int32) (expr.Literal, error) {
	return expr.NewPrimitiveLiteral[int32](value, false), nil
}

func NewInt64(value int64) (expr.Literal, error) {
	return expr.NewPrimitiveLiteral[int64](value, false), nil
}

func NewFloat32(value float32) (expr.Literal, error) {
	return expr.NewPrimitiveLiteral[float32](value, false), nil
}

func NewFloat64(value float64) (expr.Literal, error) {
	return expr.NewPrimitiveLiteral[float64](value, false), nil
}

func NewString(value string) (expr.Literal, error) {
	return expr.NewPrimitiveLiteral[string](value, false), nil
}

func NewDate(days int) (expr.Literal, error) {
	return expr.NewLiteral[types.Date](types.Date(days), false)
}

// NewTime creates a new Time literal from a time.Time value.
// This uses the number of microseconds elapsed since the start of the day.
func NewTime(tm time.Time) (expr.Literal, error) {
	startOfTheDay := time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, tm.Location())
	return expr.NewLiteral[types.Time](types.Time(tm.Sub(startOfTheDay).Microseconds()), false)
}

func NewTimeFromMicros(micros int64) (expr.Literal, error) {
	if micros < 0 || micros >= (24*time.Hour).Microseconds() {
		return nil, fmt.Errorf("invalid time value %d", micros)
	}
	return expr.NewLiteral[types.Time](types.Time(micros), false)
}

// NewTimestamp creates a new Timestamp literal from a time.Time timestamp value.
// This uses the number of microseconds elapsed since January 1, 1970 00:00:00 UTC
func NewTimestamp(timestamp time.Time) (expr.Literal, error) {
	return expr.NewLiteral[types.Timestamp](types.Timestamp(timestamp.UnixMicro()), false)
}

func NewTimestampFromMicros(micros int64) (expr.Literal, error) {
	return expr.NewLiteral[types.Timestamp](types.Timestamp(micros), false)
}

// NewTimestampTZ creates a new TimestampTz literal from a time.Time timestamp value.
// This uses the number of microseconds elapsed since January 1, 1970 00:00:00 UTC
func NewTimestampTZ(timestamp time.Time) (expr.Literal, error) {
	return expr.NewLiteral[types.TimestampTz](types.TimestampTz(timestamp.UnixMicro()), false)
}

func NewTimestampTZFromMicros(micros int64) (expr.Literal, error) {
	return expr.NewLiteral[types.TimestampTz](types.TimestampTz(micros), false)
}

func NewIntervalYearsToMonth(years, months int32) (expr.Literal, error) {
	return expr.NewLiteral[*types.IntervalYearToMonth](&types.IntervalYearToMonth{Years: years, Months: months}, false)
}

func NewIntervalDaysToSecond(days, seconds, micros int32) (expr.Literal, error) {
	return expr.NewLiteral[*types.IntervalDayToSecond](&types.IntervalDayToSecond{Days: days, Seconds: seconds, Microseconds: micros}, false)
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
