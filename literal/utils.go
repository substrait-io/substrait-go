package literal

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/cockroachdb/apd/v3"
	"github.com/google/uuid"
	"github.com/substrait-io/substrait-go/v8/expr"
	"github.com/substrait-io/substrait-go/v8/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

//This package contains utility functions for creating literals

func NewBool(value bool, nullable bool) expr.Literal {
	return expr.NewPrimitiveLiteral(value, nullable)
}

func NewInt8(value int8, nullable bool) expr.Literal {
	return expr.NewPrimitiveLiteral(value, nullable)
}

func NewInt16(value int16, nullable bool) expr.Literal {
	return expr.NewPrimitiveLiteral(value, nullable)
}

func NewInt32(value int32, nullable bool) expr.Literal {
	return expr.NewPrimitiveLiteral(value, nullable)
}

func NewInt64(value int64, nullable bool) expr.Literal {
	return expr.NewPrimitiveLiteral(value, nullable)
}

func NewFloat32(value float32, nullable bool) expr.Literal {
	return expr.NewPrimitiveLiteral(value, nullable)
}

func NewFloat64(value float64, nullable bool) expr.Literal {
	return expr.NewPrimitiveLiteral(value, nullable)
}

func NewString(value string, nullable bool) expr.Literal {
	return expr.NewPrimitiveLiteral(value, nullable)
}

func NewDate(days int, nullable bool) (expr.Literal, error) {
	return expr.NewLiteral(types.Date(days), nullable)
}

func NewDateFromString(value string, nullable bool) (expr.Literal, error) {
	tm, err := time.Parse("2006-01-02", value)
	if err != nil {
		return nil, err
	}
	return NewDate(int(tm.Unix()/86400), nullable)
}

// NewTime creates a new Time literal from the given hours, minutes, seconds and microseconds.
// The total microseconds should be in the range [0, 86400_000_000) to represent a valid time within a day.
func NewTime(hours, minutes, seconds, microseconds int32, nullable bool) (expr.Literal, error) {
	duration := time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute + time.Duration(seconds)*time.Second + time.Duration(microseconds)*time.Microsecond
	micros := duration.Microseconds()
	if micros < 0 || micros >= (24*time.Hour).Microseconds() {
		return nil, fmt.Errorf("invalid time value %d:%d:%d.%d", hours, minutes, seconds, microseconds)
	}
	return expr.NewLiteral(types.Time(duration.Microseconds()), nullable)
}

// NewTimeFromMicros creates a new Time literal from the given microseconds.
func NewTimeFromMicros(micros int64, nullable bool) (expr.Literal, error) {
	if micros < 0 || micros >= (24*time.Hour).Microseconds() {
		return nil, fmt.Errorf("invalid time value %d", micros)
	}
	return expr.NewLiteral(types.Time(micros), nullable)
}

func NewTimeFromString(value string, nullable bool) (expr.Literal, error) {
	ts, err := parseTimeFromString(value)
	if err != nil {
		return nil, err
	}
	seconds := ts.Hour()*3600 + ts.Minute()*60 + ts.Second()
	micros := int64(seconds)*int64(1e6) + int64(ts.Nanosecond())/1e3
	return NewTimeFromMicros(micros, nullable)
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
func NewTimestamp(timestamp time.Time, nullable bool) (expr.Literal, error) {
	return expr.NewLiteral(types.Timestamp(timestamp.UnixMicro()), nullable)
}

func NewTimestampFromString(value string, nullable bool) (expr.Literal, error) {
	tm, err := parseTimeFromString(value)
	if err != nil {
		return nil, err
	}
	return NewTimestamp(tm, nullable)
}

func NewTimestampFromMicros(micros int64, nullable bool) (expr.Literal, error) {
	return expr.NewLiteral(types.Timestamp(micros), nullable)
}

// NewTimestampTZ creates a new TimestampTz literal from a time.Time timestamp value.
// This uses the number of microseconds elapsed since January 1, 1970 00:00:00 UTC
func NewTimestampTZ(timestamp time.Time, nullable bool) (expr.Literal, error) {
	return expr.NewLiteral(types.TimestampTz(timestamp.UnixMicro()), nullable)
}

func NewTimestampTZFromString(value string, nullable bool) (expr.Literal, error) {
	tm, err := parseTimeFromString(value)
	if err != nil {
		return nil, err
	}
	return NewTimestampTZ(tm, nullable)
}

func NewTimestampTZFromMicros(micros int64, nullable bool) (expr.Literal, error) {
	return expr.NewLiteral(types.TimestampTz(micros), nullable)
}

func NewIntervalYearsToMonthFromString(yearsToMonth string, nullable bool) (expr.Literal, error) {
	years, months, err := parseIntervalYearsToMonth(yearsToMonth)
	if err != nil {
		return nil, err
	}
	return NewIntervalYearsToMonth(years, months, nullable)
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

func NewIntervalYearsToMonth(years, months int32, nullable bool) (expr.Literal, error) {
	return expr.NewLiteral(&types.IntervalYearToMonth{Years: years, Months: months}, nullable)
}

func NewIntervalDaysToSecondFromString(daysToSecond string, nullable bool) (expr.Literal, error) {
	days, seconds, subSeconds, precision, err := parseIntervalDaysToSecond(daysToSecond)
	if err != nil {
		return nil, err
	}
	return expr.NewLiteral(&types.IntervalDayToSecond{
		Days:    days,
		Seconds: seconds,
		PrecisionMode: &proto.Expression_Literal_IntervalDayToSecond_Precision{
			Precision: precision,
		},
		Subseconds: subSeconds,
	}, nullable)
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

func NewIntervalDaysToSecond(days, seconds int32, micros int64, nullable bool) (expr.Literal, error) {
	return expr.NewLiteral(&types.IntervalDayToSecond{
		Days:    days,
		Seconds: seconds,
		PrecisionMode: &proto.Expression_Literal_IntervalDayToSecond_Precision{
			Precision: int32(types.PrecisionMicroSeconds),
		},
		Subseconds: micros,
	}, nullable)
}

func NewUUID(guid uuid.UUID, nullable bool) (expr.Literal, error) {
	bytes, err := guid.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return expr.NewLiteral[types.UUID](bytes, nullable)
}

func NewUUIDFromBytes(value []byte, nullable bool) (expr.Literal, error) {
	return expr.NewLiteral[types.UUID](value, nullable)
}

func NewFixedChar(value string, nullable bool) (expr.Literal, error) {
	return expr.NewLiteral(types.FixedChar(value), nullable)
}

func NewFixedBinary(value []byte, nullable bool) (expr.Literal, error) {
	return expr.NewLiteral[types.FixedBinary](value, nullable)
}

func NewVarChar(value string, nullable bool) (expr.Literal, error) {
	return expr.NewLiteral(&types.VarChar{Value: value, Length: uint32(len(value))}, nullable)
}

// NewDecimalFromTwosComplement create a Decimal literal from twosComplement.
// twosComplement is a little-endian twos-complement integer representation of complete value
func NewDecimalFromTwosComplement(twosComplement []byte, precision, scale int32, nullable bool) (expr.Literal, error) {
	if len(twosComplement) != 16 {
		return nil, fmt.Errorf("twosComplement must be 16 bytes")
	}
	if precision < 1 || precision > 38 {
		return nil, fmt.Errorf("precision must be in range [1, 38]")
	}
	if scale < 0 || scale > precision {
		return nil, fmt.Errorf("scale must be in range [0, precision]")
	}
	return expr.NewLiteral(&types.Decimal{Value: twosComplement, Precision: precision, Scale: scale}, nullable)

}

// NewDecimalFromString create a Decimal literal from decimal value string
func NewDecimalFromString(value string, nullable bool) (expr.Literal, error) {
	v, precision, scale, err := expr.DecimalStringToBytes(value)
	if err != nil {
		return nil, err
	}
	return expr.NewLiteral(&types.Decimal{Value: v[:16], Precision: precision, Scale: scale}, nullable)
}

func NewDecimalFromApdDecimal(value *apd.Decimal, nullable bool) (expr.Literal, error) {
	v, precision, scale, err := expr.DecimalToBytes(value)
	if err != nil {
		return nil, err
	}
	return expr.NewLiteral(&types.Decimal{Value: v[:16], Precision: precision, Scale: scale}, nullable)
}

func NewPrecisionTime(precision types.TimePrecision, value int64, nullable bool) (expr.Literal, error) {
	return expr.NewLiteral(&types.PrecisionTime{
		Precision: int32(precision),
		Value:     value,
	}, nullable)
}

var epochTM = time.Unix(0, 0).UTC()

func NewPrecisionTimeFromTime(precision types.TimePrecision, tm time.Time, nullable bool) (expr.Literal, error) {
	// truncate the date so that we only keep the duration since midnight
	tm = epochTM.Add(tm.Sub(tm.Truncate(24 * time.Hour)))
	return NewPrecisionTime(precision, types.GetTimeValueByPrecision(tm, precision), nullable)
}

func NewPrecisionTimeFromString(precision types.TimePrecision, value string, nullable bool) (expr.Literal, error) {
	tm, err := parseTimeFromString(value)
	if err != nil {
		return nil, err
	}
	return NewPrecisionTimeFromTime(precision, tm, nullable)
}

// NewPrecisionTimestampFromTime creates a new PrecisionTimestamp literal from a time.Time timestamp value with given precision.
func NewPrecisionTimestampFromTime(precision types.TimePrecision, tm time.Time, nullable bool) (expr.Literal, error) {
	return NewPrecisionTimestamp(precision, types.GetTimeValueByPrecision(tm, precision), nullable)
}

// NewPrecisionTimestamp creates a new PrecisionTimestamp literal with given precision and value.
func NewPrecisionTimestamp(precision types.TimePrecision, value int64, nullable bool) (expr.Literal, error) {
	return expr.NewLiteral(&types.PrecisionTimestamp{
		PrecisionTimestamp: &proto.Expression_Literal_PrecisionTimestamp{
			Precision: int32(precision),
			Value:     value,
		},
	}, nullable)
}

func NewPrecisionTimestampFromString(precision types.TimePrecision, value string, nullable bool) (expr.Literal, error) {
	tm, err := parseTimeFromString(value)
	if err != nil {
		return nil, err
	}
	return NewPrecisionTimestampFromTime(precision, tm, nullable)
}

// NewPrecisionTimestampTzFromTime creates a new PrecisionTimestampTz literal from a time.Time timestamp value with given precision.
func NewPrecisionTimestampTzFromTime(precision types.TimePrecision, tm time.Time, nullable bool) (expr.Literal, error) {
	return NewPrecisionTimestampTz(precision, types.GetTimeValueByPrecision(tm, precision), nullable)
}

// NewPrecisionTimestampTz creates a new PrecisionTimestampTz literal with given precision and value.
func NewPrecisionTimestampTz(precision types.TimePrecision, value int64, nullable bool) (expr.Literal, error) {
	return expr.NewLiteral(&types.PrecisionTimestampTz{
		PrecisionTimestampTz: &proto.Expression_Literal_PrecisionTimestamp{
			Precision: int32(precision),
			Value:     value,
		},
	}, nullable)
}

func NewPrecisionTimestampTzFromString(precision types.TimePrecision, value string, nullable bool) (expr.Literal, error) {
	tm, err := parseTimeFromString(value)
	if err != nil {
		return nil, err
	}
	return NewPrecisionTimestampTzFromTime(precision, tm, nullable)
}

// firstMismatch returns the index of the first nil literal or literal whose base
// type differs from the column's anchor (the first literal), or false if every
// literal is compatible. Null literals are allowed in any position when their
// typed base matches the anchor. Type parameters (such as varchar length) and
// nullability may differ within one base type; a later cast unifies those. Base
// types are compared recursively rather than by Go wrapper, which is too coarse:
// distinct types such as decimal and varchar share one wrapper, and nested types
// such as list<i8> and list<i32> must remain distinct.
func firstMismatch(column []expr.Literal) (int, bool) {
	if column[0] == nil {
		return 0, true
	}
	anchor := column[0].GetType()
	for i, l := range column[1:] {
		if l == nil {
			return i + 1, true
		}
		if !sameBaseType(anchor, l.GetType()) {
			return i + 1, true
		}
	}
	return 0, false
}

func sameBaseType(left, right types.Type) bool {
	if left == nil || right == nil {
		return false
	}
	if left.ShortString() != right.ShortString() {
		return false
	}

	switch left := left.(type) {
	case *types.ListType:
		right, ok := right.(*types.ListType)
		return ok && sameBaseType(left.Type, right.Type)
	case *types.MapType:
		right, ok := right.(*types.MapType)
		return ok && sameBaseType(left.Key, right.Key) && sameBaseType(left.Value, right.Value)
	case *types.StructType:
		right, ok := right.(*types.StructType)
		if !ok || len(left.Types) != len(right.Types) {
			return false
		}
		for i := range left.Types {
			if !sameBaseType(left.Types[i], right.Types[i]) {
				return false
			}
		}
		return true
	case *types.FuncType:
		right, ok := right.(*types.FuncType)
		if !ok || len(left.ParameterTypes) != len(right.ParameterTypes) {
			return false
		}
		if !sameBaseType(left.ReturnType, right.ReturnType) {
			return false
		}
		for i := range left.ParameterTypes {
			if !sameBaseType(left.ParameterTypes[i], right.ParameterTypes[i]) {
				return false
			}
		}
		return true
	case *types.UserDefinedType:
		right, ok := right.(*types.UserDefinedType)
		if !ok || left.TypeReference != right.TypeReference || len(left.TypeParameters) != len(right.TypeParameters) {
			return false
		}
		for i := range left.TypeParameters {
			if !sameBaseTypeParam(left.TypeParameters[i], right.TypeParameters[i]) {
				return false
			}
		}
		return true
	default:
		return true
	}
}

func sameBaseTypeParam(left, right types.TypeParam) bool {
	if left == nil || right == nil {
		return left == right
	}
	leftData, leftIsData := left.(*types.DataTypeParameter)
	rightData, rightIsData := right.(*types.DataTypeParameter)
	if leftIsData || rightIsData {
		return leftIsData && rightIsData && sameBaseType(leftData.Type, rightData.Type)
	}
	return left.Equals(right)
}

func NewList(elements []expr.Literal, nullable bool) (expr.Literal, error) {
	if len(elements) == 0 {
		return nil, fmt.Errorf("empty list literal; use NewEmptyList for an empty list")
	}
	if i, ok := firstMismatch(elements); ok {
		return nil, fmt.Errorf("element %d of list literal has invalid or different type", i)
	}
	return expr.NewLiteral[expr.ListLiteralValue](elements, nullable)
}

// NewMap creates a map literal from the given key/value entries. As in NewList,
// all keys must share one base type and all values one base type; null literals
// are allowed in either position, and type parameters (such as varchar length)
// and nullability may differ. The map's declared key/value types are taken from
// the first entry. For an empty map use NewEmptyMap, which takes the key and
// value types explicitly.
func NewMap(entries expr.MapLiteralValue, nullable bool) (expr.Literal, error) {
	if len(entries) == 0 {
		return nil, fmt.Errorf("empty map literal; use NewEmptyMap for an empty map")
	}
	keys := make([]expr.Literal, len(entries))
	values := make([]expr.Literal, len(entries))
	for i, e := range entries {
		keys[i], values[i] = e.Key, e.Value
	}
	if i, ok := firstMismatch(keys); ok {
		return nil, fmt.Errorf("key %d of map literal has invalid or different type", i)
	}
	if i, ok := firstMismatch(values); ok {
		return nil, fmt.Errorf("value %d of map literal has invalid or different type", i)
	}
	return expr.NewLiteral[expr.MapLiteralValue](entries, nullable)
}

// NewEmptyMap creates an empty map literal of the given key and value types,
// marked nullable or not.
func NewEmptyMap(keyType, valueType types.Type, nullable bool) (expr.Literal, error) {
	if keyType == nil {
		return nil, fmt.Errorf("empty map literal key type cannot be nil")
	}
	if valueType == nil {
		return nil, fmt.Errorf("empty map literal value type cannot be nil")
	}
	return expr.NewEmptyMapLiteral(keyType, valueType, nullable), nil
}

// NewEmptyList creates an empty list literal of the given element type, marked
// nullable or not.
func NewEmptyList(elementType types.Type, nullable bool) (expr.Literal, error) {
	if elementType == nil {
		return nil, fmt.Errorf("empty list literal element type cannot be nil")
	}
	return expr.NewEmptyListLiteral(elementType, nullable), nil
}

// NewUserDefinedLiteral creates a user-defined literal using the struct representation.
// The typeRef should be obtained from an extension registry's GetTypeAnchor method.
// The structValue contains the field values for the user-defined type.
// Optional type parameters can be provided for parameterized user-defined types (pass nil for none).
func NewUserDefinedLiteral(typeRef uint32, structValue expr.StructLiteralValue, nullable bool, typeParams []types.TypeParam) (expr.Literal, error) {
	structProto := structValue.ToProto()

	protoParams := make([]*proto.Type_Parameter, len(typeParams))
	for i, p := range typeParams {
		protoParams[i] = p.ToProto()
	}

	return expr.NewLiteral(
		&types.UserDefinedLiteral{
			Val:            &proto.Expression_Literal_UserDefined_Struct{Struct: structProto},
			TypeAnchorType: &proto.Expression_Literal_UserDefined_TypeReference{TypeReference: typeRef},
			TypeParameters: protoParams,
		},
		nullable,
	)
}
