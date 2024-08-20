package literal

import (
	"fmt"
	"math"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/expr"
	"github.com/substrait-io/substrait-go/proto"
	"github.com/substrait-io/substrait-go/types"
)

func TestNewBool(t *testing.T) {
	tests := []struct {
		name    string
		value   bool
		want    expr.Literal
		wantErr assert.ErrorAssertionFunc
	}{
		{"false", false, expr.NewPrimitiveLiteral[bool](false, false), assert.NoError},
		{"true", true, expr.NewPrimitiveLiteral[bool](true, false), assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBool(tt.value)
			if !tt.wantErr(t, err, fmt.Sprintf("NewBool(%v)", tt.value)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewBool(%v)", tt.value)
		})
	}
}

func TestNewDate(t *testing.T) {
	tests := []struct {
		name    string
		days    int
		want    expr.Literal
		wantErr assert.ErrorAssertionFunc
	}{
		{"0days", 0, expr.NewPrimitiveLiteral(types.Date(0), false), assert.NoError},
		{"10000days", 10000, expr.NewPrimitiveLiteral(types.Date(10000), false), assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDate(tt.days)
			if !tt.wantErr(t, err, fmt.Sprintf("NewDate(%v)", tt.days)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewDate(%v)", tt.days)
		})
	}
}

func TestNewDecimalFromString(t *testing.T) {
	tests := []struct {
		value   string
		want    expr.Literal
		wantErr assert.ErrorAssertionFunc
	}{
		{"0", createDecimalLiteral([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 1, 0, false), assert.NoError},
		{"111111.222222", createDecimalLiteral([]byte{0xce, 0xb3, 0xbe, 0xde, 0x19, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 12, 6, false), assert.NoError},
		{"-111111.222222", createDecimalLiteral([]byte{0x32, 0x4c, 0x41, 0x21, 0xe6, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, 12, 6, false), assert.NoError},
		{"+1", createDecimalLiteral([]byte{0x1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 1, 0, false), assert.NoError},
		{"-1", createDecimalLiteral([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, 1, 0, false), assert.NoError},
		{"not a decimal", nil, assert.Error},
	}
	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			got, err := NewDecimalFromString(tt.value)
			if !tt.wantErr(t, err, fmt.Sprintf("NewDecimalFromString(%v)", tt.value)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewDecimalFromString(%v)", tt.value)
		})
	}
}

func createDecimalLiteral(value []byte, precision int32, scale int32, isNullable bool) *expr.ProtoLiteral {
	nullability := proto.Type_NULLABILITY_REQUIRED
	if isNullable {
		nullability = proto.Type_NULLABILITY_NULLABLE
	}
	return &expr.ProtoLiteral{
		Value: value[:16],
		Type: &types.DecimalType{
			Nullability: nullability,
			Precision:   precision,
			Scale:       scale,
		},
	}
}
func TestNewDecimalFromTwosComplement(t *testing.T) {
	type args struct {
		twosComplement []byte
		precision      int32
		scale          int32
	}
	tests := []struct {
		name    string
		args    args
		want    expr.Literal
		wantErr assert.ErrorAssertionFunc
	}{
		{"0", args{[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 1, 0},
			createDecimalLiteral([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 1, 0, false), assert.NoError},
		{"123", args{[]byte{0x7b, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 7, 0},
			createDecimalLiteral([]byte{0x7b, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 7, 0, false), assert.NoError},
		{"111111.222222", args{[]byte{0xce, 0xb3, 0xbe, 0xde, 0x19, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 12, 6},
			createDecimalLiteral([]byte{0xce, 0xb3, 0xbe, 0xde, 0x19, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 12, 6, false), assert.NoError},
		{"-111111.222222", args{[]byte{0x32, 0x4c, 0x41, 0x21, 0xe6, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, 12, 6},
			createDecimalLiteral([]byte{0x32, 0x4c, 0x41, 0x21, 0xe6, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, 12, 6, false), assert.NoError},
		{"precision 0 out of range", args{[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 0, 0}, nil, assert.Error},
		{"precision 40 out of range", args{[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 40, 0}, nil, assert.Error},
		{"precision -1 out of range", args{[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, -1, 0}, nil, assert.Error},
		{"scale out of range", args{[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 10, 12}, nil, assert.Error},
		{"scale out of range", args{[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 10, -1}, nil, assert.Error},
		{"invalid twosComplement", args{[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 10, 2}, nil, assert.Error},
		{"invalid twosComplement", args{[]byte{0, 0}, 10, 2}, nil, assert.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDecimalFromTwosComplement(tt.args.twosComplement, tt.args.precision, tt.args.scale)
			if !tt.wantErr(t, err, fmt.Sprintf("NewDecimalFromTwosComplement(%v, %v, %v)", tt.args.twosComplement, tt.args.precision, tt.args.scale)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewDecimalFromTwosComplement(%v, %v, %v)", tt.args.twosComplement, tt.args.precision, tt.args.scale)
		})
	}
}

func TestNewFixedBinary(t *testing.T) {
	tests := []struct {
		name    string
		value   []byte
		want    expr.Literal
		wantErr assert.ErrorAssertionFunc
	}{
		{"empty", []byte{}, expr.NewFixedBinaryLiteral([]byte{}, false), assert.NoError},
		{"1 byte", []byte{0x1}, expr.NewFixedBinaryLiteral([]byte{0x1}, false), assert.NoError},
		{"2 bytes", []byte{0x1, 0x2}, expr.NewFixedBinaryLiteral([]byte{0x1, 0x2}, false), assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFixedBinary(tt.value)
			if !tt.wantErr(t, err, fmt.Sprintf("NewFixedBinary(%v)", tt.value)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewFixedBinary(%v)", tt.value)
		})
	}
}

func TestNewFixedChar(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    expr.Literal
		wantErr assert.ErrorAssertionFunc
	}{
		{"empty", "", expr.NewFixedCharLiteral("", false), assert.NoError},
		{"1 char", "a", expr.NewFixedCharLiteral("a", false), assert.NoError},
		{"2 chars", "ab", expr.NewFixedCharLiteral("ab", false), assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFixedChar(tt.value)
			if !tt.wantErr(t, err, fmt.Sprintf("NewFixedChar(%v)", tt.value)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewFixedChar(%v)", tt.value)
		})
	}
}

func TestNewFloat32(t *testing.T) {
	tests := []struct {
		name    string
		value   float32
		want    expr.Literal
		wantErr assert.ErrorAssertionFunc
	}{
		{"0", 0, expr.NewPrimitiveLiteral[float32](0, false), assert.NoError},
		{"1.1", 1.1, expr.NewPrimitiveLiteral[float32](1.1, false), assert.NoError},
		{"-1.1", -1.1, expr.NewPrimitiveLiteral[float32](-1.1, false), assert.NoError},
		{"NaN", float32(math.NaN()), expr.NewPrimitiveLiteral[float32](float32(math.NaN()), false), assert.NoError},
		{"+Inf", float32(math.Inf(1)), expr.NewPrimitiveLiteral[float32](float32(math.Inf(1)), false), assert.NoError},
		{"-Inf", float32(math.Inf(-1)), expr.NewPrimitiveLiteral[float32](float32(math.Inf(-1)), false), assert.NoError},
		{"max float32", float32(math.MaxFloat32), expr.NewPrimitiveLiteral[float32](math.MaxFloat32, false), assert.NoError},
		{"min float32", float32(math.SmallestNonzeroFloat32), expr.NewPrimitiveLiteral[float32](math.SmallestNonzeroFloat32, false), assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFloat32(tt.value)
			if !tt.wantErr(t, err, fmt.Sprintf("NewFloat32(%v)", tt.value)) {
				return
			}
			if !math.IsNaN(float64(tt.value)) {
				assert.Equalf(t, tt.want, got, "NewFloat32(%v)", tt.value)
			} else {
				protoExp := got.ToProto()
				assert.True(t, math.IsNaN(float64(protoExp.GetLiteral().GetFp32())), "NewFloat32(%v)", tt.value)
			}
		})
	}
}

func TestNewFloat64(t *testing.T) {
	tests := []struct {
		name    string
		value   float64
		want    expr.Literal
		wantErr assert.ErrorAssertionFunc
	}{
		{"0", 0, expr.NewPrimitiveLiteral[float64](0, false), assert.NoError},
		{"1.1", 1.1, expr.NewPrimitiveLiteral[float64](1.1, false), assert.NoError},
		{"-1.1", -1.1, expr.NewPrimitiveLiteral[float64](-1.1, false), assert.NoError},
		{"NaN", math.NaN(), expr.NewPrimitiveLiteral[float64](math.NaN(), false), assert.NoError},
		{"+Inf", math.Inf(1), expr.NewPrimitiveLiteral[float64](math.Inf(1), false), assert.NoError},
		{"-Inf", math.Inf(-1), expr.NewPrimitiveLiteral[float64](math.Inf(-1), false), assert.NoError},
		{"max float64", math.MaxFloat64, expr.NewPrimitiveLiteral[float64](math.MaxFloat64, false), assert.NoError},
		{"min float64", math.SmallestNonzeroFloat64, expr.NewPrimitiveLiteral[float64](math.SmallestNonzeroFloat64, false), assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFloat64(tt.value)
			if !tt.wantErr(t, err, fmt.Sprintf("NewFloat64(%v)", tt.value)) {
				return
			}
			if !math.IsNaN(tt.value) {
				assert.Equalf(t, tt.want, got, "NewFloat64(%v)", tt.value)
			} else {
				protoExp := got.ToProto()
				assert.True(t, math.IsNaN(protoExp.GetLiteral().GetFp64()), "NewFloat64(%v)", tt.value)
			}
		})
	}
}

func TestNewInt16(t *testing.T) {
	tests := []struct {
		name    string
		value   int16
		want    expr.Literal
		wantErr assert.ErrorAssertionFunc
	}{
		{"0", 0, expr.NewPrimitiveLiteral[int16](0, false), assert.NoError},
		{"1", 1, expr.NewPrimitiveLiteral[int16](1, false), assert.NoError},
		{"-1", -1, expr.NewPrimitiveLiteral[int16](-1, false), assert.NoError},
		{"max int16", math.MaxInt16, expr.NewPrimitiveLiteral[int16](math.MaxInt16, false), assert.NoError},
		{"min int16", math.MinInt16, expr.NewPrimitiveLiteral[int16](math.MinInt16, false), assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewInt16(tt.value)
			if !tt.wantErr(t, err, fmt.Sprintf("NewInt16(%v)", tt.value)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewInt16(%v)", tt.value)
		})
	}
}

func TestNewInt32(t *testing.T) {
	tests := []struct {
		name    string
		value   int32
		want    expr.Literal
		wantErr assert.ErrorAssertionFunc
	}{
		{"0", 0, expr.NewPrimitiveLiteral[int32](0, false), assert.NoError},
		{"1", 1, expr.NewPrimitiveLiteral[int32](1, false), assert.NoError},
		{"-1", -1, expr.NewPrimitiveLiteral[int32](-1, false), assert.NoError},
		{"max int32", math.MaxInt32, expr.NewPrimitiveLiteral[int32](math.MaxInt32, false), assert.NoError},
		{"min int32", math.MinInt32, expr.NewPrimitiveLiteral[int32](math.MinInt32, false), assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewInt32(tt.value)
			if !tt.wantErr(t, err, fmt.Sprintf("NewInt32(%v)", tt.value)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewInt32(%v)", tt.value)
		})
	}
}

func TestNewInt64(t *testing.T) {
	tests := []struct {
		name    string
		value   int64
		want    expr.Literal
		wantErr assert.ErrorAssertionFunc
	}{
		{"0", 0, expr.NewPrimitiveLiteral[int64](0, false), assert.NoError},
		{"1", 1, expr.NewPrimitiveLiteral[int64](1, false), assert.NoError},
		{"-1", -1, expr.NewPrimitiveLiteral[int64](-1, false), assert.NoError},
		{"max int64", math.MaxInt64, expr.NewPrimitiveLiteral[int64](math.MaxInt64, false), assert.NoError},
		{"min int64", math.MinInt64, expr.NewPrimitiveLiteral[int64](math.MinInt64, false), assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewInt64(tt.value)
			if !tt.wantErr(t, err, fmt.Sprintf("NewInt64(%v)", tt.value)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewInt64(%v)", tt.value)
		})
	}
}

func TestNewInt8(t *testing.T) {
	tests := []struct {
		name    string
		value   int8
		want    expr.Literal
		wantErr assert.ErrorAssertionFunc
	}{
		{"0", 0, expr.NewPrimitiveLiteral[int8](0, false), assert.NoError},
		{"1", 1, expr.NewPrimitiveLiteral[int8](1, false), assert.NoError},
		{"-1", -1, expr.NewPrimitiveLiteral[int8](-1, false), assert.NoError},
		{"max int8", math.MaxInt8, expr.NewPrimitiveLiteral[int8](math.MaxInt8, false), assert.NoError},
		{"min int8", math.MinInt8, expr.NewPrimitiveLiteral[int8](math.MinInt8, false), assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewInt8(tt.value)
			if !tt.wantErr(t, err, fmt.Sprintf("NewInt8(%v)", tt.value)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewInt8(%v)", tt.value)
		})
	}
}

func TestNewIntervalDaysToSecond(t *testing.T) {
	tests := []struct {
		name    string
		days    int32
		seconds int32
		micros  int64
		want    expr.Literal
		wantErr assert.ErrorAssertionFunc
	}{
		{"0", 0, 0, 0, createIntervalDaysLiteral(0, 0, 0), assert.NoError},
		{"10 day", 10, 0, 0, createIntervalDaysLiteral(10, 0, 0), assert.NoError},
		{"20 second", 0, 20, 0, createIntervalDaysLiteral(0, 20, 0), assert.NoError},
		{"30 microsecond", 0, 0, 30, createIntervalDaysLiteral(0, 0, 30), assert.NoError},
		{"1 day, 1 second, 1 microsecond", 1, 1, 1, createIntervalDaysLiteral(1, 1, 1), assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewIntervalDaysToSecond(tt.days, tt.seconds, tt.micros)
			if !tt.wantErr(t, err, fmt.Sprintf("NewIntervalDaysToSecond(%v, %v, %v)", tt.days, tt.seconds, tt.micros)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewIntervalDaysToSecond(%v, %v, %v)", tt.days, tt.seconds, tt.micros)
		})
	}
}

func createIntervalDaysLiteral(days, seconds int32, micros int64) *expr.ProtoLiteral {
	return &expr.ProtoLiteral{
		Value: &types.IntervalDayToSecond{
			Days:       days,
			Seconds:    seconds,
			Subseconds: micros,
			PrecisionMode: &proto.Expression_Literal_IntervalDayToSecond_Precision{
				Precision: int32(types.PrecisionMicroSeconds),
			},
		},
		Type: &types.IntervalDayType{
			Nullability: proto.Type_NULLABILITY_REQUIRED,
		},
	}
}

func TestNewIntervalYearsToMonth(t *testing.T) {
	tests := []struct {
		name    string
		years   int32
		months  int32
		want    expr.Literal
		wantErr assert.ErrorAssertionFunc
	}{
		{"0", 0, 0, createIntervalYearsLiteral(0, 0), assert.NoError},
		{"10 year", 10, 0, createIntervalYearsLiteral(10, 0), assert.NoError},
		{"10 month", 0, 10, createIntervalYearsLiteral(0, 10), assert.NoError},
		{"1 year, 1 month", 1, 1, createIntervalYearsLiteral(1, 1), assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewIntervalYearsToMonth(tt.years, tt.months)
			if !tt.wantErr(t, err, fmt.Sprintf("NewIntervalYearsToMonth(%v, %v)", tt.years, tt.months)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewIntervalYearsToMonth(%v, %v)", tt.years, tt.months)
		})
	}
}

func createIntervalYearsLiteral(years, months int32) *expr.ProtoLiteral {
	return &expr.ProtoLiteral{
		Value: &types.IntervalYearToMonth{
			Years:  years,
			Months: months,
		},
		Type: &types.IntervalYearType{
			Nullability: proto.Type_NULLABILITY_REQUIRED,
		},
	}
}

func TestNewString(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    expr.Literal
		wantErr assert.ErrorAssertionFunc
	}{
		{"empty", "", expr.NewPrimitiveLiteral("", false), assert.NoError},
		{"1 char", "a", expr.NewPrimitiveLiteral("a", false), assert.NoError},
		{"3 chars", "abc", expr.NewPrimitiveLiteral("abc", false), assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewString(tt.value)
			if !tt.wantErr(t, err, fmt.Sprintf("NewString(%v)", tt.value)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewString(%v)", tt.value)
		})
	}
}

func TestNewTime(t *testing.T) {
	type testcase struct {
		name    string
		hours   int32
		minutes int32
		seconds int32
		micros  int32
	}
	tests := []testcase{
		{"zero", 0, 0, 0, 0},
		{"10:30:20.999", 10, 30, 40, 999},
		{"23:59:59.99999", 23, 59, 59, 999999},
		{"300 minutes", 0, 300, 0, 0},
		{"5000 seconds", 0, 0, 5000, 0},
		{"86399 seconds", 0, 0, 86399, 0},
		{"MaxInt32 microseconds", 0, 0, 0, math.MaxInt32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTime(tt.hours, tt.minutes, tt.seconds, tt.micros)
			assert.NoError(t, err)
			want := expr.NewPrimitiveLiteral(types.Time(getMicroSeconds(tt.hours, tt.minutes, tt.seconds, tt.micros)), false)
			assert.Equal(t, want, got)
		})
	}

	negTests := []testcase{
		{"24:01:01.123", 24, 1, 1, 123},
		{"23:60:01.123", 23, 60, 1, 123},
		{"23:59:60.123", 23, 59, 60, 123},
		{"23:59:59.1000000", 23, 59, 59, 1_000_000},
		{"23:59:59.MaxInt32", 23, 59, 59, math.MaxInt32},
		{"86400 seconds", 0, 0, 86400, 0},
		{"24 hours", 24, 0, 0, 0},
		{"-1 hour", -1, 0, 0, 0},
		{"1440 minutes", 0, 1440, 0, 0},
	}
	for _, tt := range negTests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewTime(tt.hours, tt.minutes, tt.seconds, tt.micros)
			assert.Error(t, err)
		})
	}
}

func getMicroSeconds(hours, minutes, seconds, microseconds int32) int64 {
	return (time.Duration(hours) * time.Hour).Microseconds() +
		(time.Duration(minutes) * time.Minute).Microseconds() +
		(time.Duration(seconds) * time.Second).Microseconds() +
		(time.Duration(microseconds) * time.Microsecond).Microseconds()
}

func TestNewTimeFromMicros(t *testing.T) {
	tests := []struct {
		name    string
		micros  int64
		want    expr.Literal
		wantErr assert.ErrorAssertionFunc
	}{
		{"zero", 0, expr.NewPrimitiveLiteral(types.Time(0), false), assert.NoError},
		{"now", time.Now().UnixMicro(), nil, assert.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTimeFromMicros(tt.micros)
			if !tt.wantErr(t, err, fmt.Sprintf("NewTimeFromMicros(%v)", tt.micros)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewTimeFromMicros(%v)", tt.micros)
		})
	}
}

func TestNewTimestamp(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name    string
		tm      time.Time
		want    expr.Literal
		wantErr assert.ErrorAssertionFunc
	}{
		{"zero", time.Unix(0, 0), expr.NewPrimitiveLiteral(types.Timestamp(0), false), assert.NoError},
		{"now", now, expr.NewPrimitiveLiteral(types.Timestamp(now.UnixMicro()), false), assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTimestamp(tt.tm)
			if !tt.wantErr(t, err, fmt.Sprintf("NewTimestamp(%v)", tt.tm)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewTimestamp(%v)", tt.tm)
		})
	}
}

func TestNewTimestampFromMicros(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name    string
		micros  int64
		want    expr.Literal
		wantErr assert.ErrorAssertionFunc
	}{
		{"zero", 0, expr.NewPrimitiveLiteral(types.Timestamp(0), false), assert.NoError},
		{"now", now.UnixMicro(), expr.NewPrimitiveLiteral(types.Timestamp(now.UnixMicro()), false), assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTimestampFromMicros(tt.micros)
			if !tt.wantErr(t, err, fmt.Sprintf("NewTimestampFromMicros(%v)", tt.micros)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewTimestampFromMicros(%v)", tt.micros)
		})
	}
}

func TestNewTimestampTZ(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name    string
		tm      time.Time
		want    expr.Literal
		wantErr assert.ErrorAssertionFunc
	}{
		{"zero", time.Unix(0, 0), expr.NewPrimitiveLiteral(types.TimestampTz(0), false), assert.NoError},
		{"now", now, expr.NewPrimitiveLiteral(types.TimestampTz(now.UnixMicro()), false), assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTimestampTZ(tt.tm)
			if !tt.wantErr(t, err, fmt.Sprintf("NewTimestampTZ(%v)", tt.tm)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewTimestampTZ(%v)", tt.tm)
		})
	}
}

func TestNewTimestampTZFromMicros(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name    string
		micros  int64
		want    expr.Literal
		wantErr assert.ErrorAssertionFunc
	}{
		{"zero", 0, expr.NewPrimitiveLiteral(types.TimestampTz(0), false), assert.NoError},
		{"now", now.UnixMicro(), expr.NewPrimitiveLiteral(types.TimestampTz(now.UnixMicro()), false), assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTimestampTZFromMicros(tt.micros)
			if !tt.wantErr(t, err, fmt.Sprintf("NewTimestampTZFromMicros(%v)", tt.micros)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewTimestampTZFromMicros(%v)", tt.micros)
		})
	}
}

func TestNewUUID(t *testing.T) {
	guid := uuid.New()
	tests := []struct {
		name    string
		guid    uuid.UUID
		want    expr.Literal
		wantErr assert.ErrorAssertionFunc
	}{
		{"valid", guid, expr.NewByteSliceLiteral[types.UUID](guid[:], false), assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUUID(tt.guid)
			if !tt.wantErr(t, err, fmt.Sprintf("NewUUID(%v)", tt.guid)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewUUID(%v)", tt.guid)
		})
	}
}

func TestNewUUIDFromBytes(t *testing.T) {
	tests := []struct {
		name    string
		value   []byte
		want    expr.Literal
		wantErr assert.ErrorAssertionFunc
	}{
		{"16 bytes", []byte{0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xf, 0x10},
			expr.NewByteSliceLiteral[types.UUID]([]byte{0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xf, 0x10}, false), assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUUIDFromBytes(tt.value)
			if !tt.wantErr(t, err, fmt.Sprintf("NewUUIDFromBytes(%v)", tt.value)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewUUIDFromBytes(%v)", tt.value)
		})
	}
}

func TestNewVarChar(t *testing.T) {
	longStr := strings.Repeat("a", 1000)
	tests := []struct {
		name    string
		value   string
		want    expr.Literal
		wantErr assert.ErrorAssertionFunc
	}{
		{"empty", "", createVarCharLiteral(""), assert.NoError},
		{"1 char", "a", createVarCharLiteral("a"), assert.NoError},
		{"3 chars", "abc", createVarCharLiteral("abc"), assert.NoError},
		{"long varchar", longStr, createVarCharLiteral(longStr), assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewVarChar(tt.value)
			if !tt.wantErr(t, err, fmt.Sprintf("NewVarChar(%v)", tt.value)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewVarChar(%v)", tt.value)
		})
	}
}

func createVarCharLiteral(value string) *expr.ProtoLiteral {
	return &expr.ProtoLiteral{
		// TODO check if .Value should be types.VarChar instead of string
		//Value: &types.VarChar{
		//	Value:  value,
		//	Length: uint32(len(value)),
		//},
		Value: value,
		Type: &types.VarCharType{
			Nullability: proto.Type_NULLABILITY_REQUIRED,
			Length:      int32(len(value)),
		},
	}
}
