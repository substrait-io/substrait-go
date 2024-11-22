package literal

import (
	"fmt"
	"math"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
			got := NewBool(tt.value)
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
		name  string
		value float32
		want  expr.Literal
	}{
		{"0", 0, expr.NewPrimitiveLiteral[float32](0, false)},
		{"1.1", 1.1, expr.NewPrimitiveLiteral[float32](1.1, false)},
		{"-1.1", -1.1, expr.NewPrimitiveLiteral[float32](-1.1, false)},
		{"NaN", float32(math.NaN()), expr.NewPrimitiveLiteral[float32](float32(math.NaN()), false)},
		{"+Inf", float32(math.Inf(1)), expr.NewPrimitiveLiteral[float32](float32(math.Inf(1)), false)},
		{"-Inf", float32(math.Inf(-1)), expr.NewPrimitiveLiteral[float32](float32(math.Inf(-1)), false)},
		{"max float32", float32(math.MaxFloat32), expr.NewPrimitiveLiteral[float32](math.MaxFloat32, false)},
		{"min float32", float32(math.SmallestNonzeroFloat32), expr.NewPrimitiveLiteral[float32](math.SmallestNonzeroFloat32, false)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewFloat32(tt.value)
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
		name  string
		value float64
		want  expr.Literal
	}{
		{"0", 0, expr.NewPrimitiveLiteral[float64](0, false)},
		{"1.1", 1.1, expr.NewPrimitiveLiteral[float64](1.1, false)},
		{"-1.1", -1.1, expr.NewPrimitiveLiteral[float64](-1.1, false)},
		{"NaN", math.NaN(), expr.NewPrimitiveLiteral[float64](math.NaN(), false)},
		{"+Inf", math.Inf(1), expr.NewPrimitiveLiteral[float64](math.Inf(1), false)},
		{"-Inf", math.Inf(-1), expr.NewPrimitiveLiteral[float64](math.Inf(-1), false)},
		{"max float64", math.MaxFloat64, expr.NewPrimitiveLiteral[float64](math.MaxFloat64, false)},
		{"min float64", math.SmallestNonzeroFloat64, expr.NewPrimitiveLiteral[float64](math.SmallestNonzeroFloat64, false)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewFloat64(tt.value)
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
			got := NewInt16(tt.value)
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
			got := NewInt32(tt.value)
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
			got := NewInt64(tt.value)
			assert.Equalf(t, tt.want, got, "NewInt64(%v)", tt.value)
		})
	}
}

func TestNewInt8(t *testing.T) {
	tests := []struct {
		name  string
		value int8
		want  expr.Literal
	}{
		{"0", 0, expr.NewPrimitiveLiteral[int8](0, false)},
		{"1", 1, expr.NewPrimitiveLiteral[int8](1, false)},
		{"-1", -1, expr.NewPrimitiveLiteral[int8](-1, false)},
		{"max int8", math.MaxInt8, expr.NewPrimitiveLiteral[int8](math.MaxInt8, false)},
		{"min int8", math.MinInt8, expr.NewPrimitiveLiteral[int8](math.MinInt8, false)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewInt8(tt.value)
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

func createIntervalDaysLiteralWithNanos(days, seconds int32, nanos int64) *expr.ProtoLiteral {
	return &expr.ProtoLiteral{
		Value: &types.IntervalDayToSecond{
			Days:       days,
			Seconds:    seconds,
			Subseconds: nanos,
			PrecisionMode: &proto.Expression_Literal_IntervalDayToSecond_Precision{
				Precision: int32(types.PrecisionNanoSeconds),
			},
		},
		Type: &types.IntervalDayType{
			Nullability: proto.Type_NULLABILITY_REQUIRED,
		},
	}
}

func TestNewIntervalDaysToSecondFromString(t *testing.T) {
	tests := []struct {
		name    string
		want    expr.Literal
		wantErr assert.ErrorAssertionFunc
	}{
		{"P0D", createIntervalDaysLiteral(0, 0, 0), assert.NoError},
		{"P1D", createIntervalDaysLiteral(1, 0, 0), assert.NoError},
		{"PT2H", createIntervalDaysLiteral(0, 7200, 0), assert.NoError},
		{"PT3M", createIntervalDaysLiteral(0, 180, 0), assert.NoError},
		{"PT4S", createIntervalDaysLiteral(0, 4, 0), assert.NoError},
		{"PT0.5S", createIntervalDaysLiteral(0, 0, 500000), assert.NoError},
		{"P1DT2H3M4.5S", createIntervalDaysLiteral(1, 7384, 500000), assert.NoError},
		{"P10DT12H45M30.25S", createIntervalDaysLiteral(10, 45930, 250000), assert.NoError},
		{"P2DT3H", createIntervalDaysLiteral(2, 10800, 0), assert.NoError},
		{"P5DT10M", createIntervalDaysLiteral(5, 600, 0), assert.NoError},
		{"PT15M20S", createIntervalDaysLiteral(0, 920, 0), assert.NoError},
		{"PT1H5M10.5S", createIntervalDaysLiteral(0, 3910, 500000), assert.NoError},
		{"P3DT4H5S", createIntervalDaysLiteral(3, 14405, 0), assert.NoError},
		{"PT0.75S", createIntervalDaysLiteral(0, 0, 750000), assert.NoError},
		{"PT23H59M59.999S", createIntervalDaysLiteral(0, 86399, 999000), assert.NoError},
		{"PT0.0S", createIntervalDaysLiteral(0, 0, 0), assert.NoError},
		{"P1000D", createIntervalDaysLiteral(1000, 0, 0), assert.NoError},
		{"PT10000H", createIntervalDaysLiteral(0, 36000000, 0), assert.NoError},
		{"PT100000M", createIntervalDaysLiteral(0, 6000000, 0), assert.NoError},
		{"PT86400S", createIntervalDaysLiteral(0, 86400, 0), assert.NoError},
		{"PT1.25S", createIntervalDaysLiteral(0, 1, 250000), assert.NoError},
		{"P0DT0H0M0.0S", createIntervalDaysLiteral(0, 0, 0), assert.NoError},
		{"PT0.123S", createIntervalDaysLiteral(0, 0, 123000), assert.NoError},
		{"PT0.999S", createIntervalDaysLiteral(0, 0, 999000), assert.NoError},
		{"PT0.9999999S", createIntervalDaysLiteralWithNanos(0, 0, 999999900), assert.NoError},
		{"PT1.5S", createIntervalDaysLiteral(0, 1, 500000), assert.NoError},
		{"PT1S", createIntervalDaysLiteral(0, 1, 0), assert.NoError},
		{"P", nil, assert.Error},
		{"PT", nil, assert.Error},
		{"P5M", nil, assert.Error},
		{"PT0H0M0S0F", nil, assert.Error},
		{"P1DT2H3S5M", nil, assert.Error},
		{"P1D2H3M", nil, assert.Error},
		{"PT1H2.5M", nil, assert.Error},
		{"P10DT-5H", nil, assert.Error},
		{"P10DT3H0X", nil, assert.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewIntervalDaysToSecondFromString(tt.name)
			if !tt.wantErr(t, err, fmt.Sprintf("NewIntervalDaysToSecondFromString(%v)", tt.name)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewIntervalDaysToSecondFromString(%v)", tt.name)
		})
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

func TestNewIntervalYearsToMonthFromString(t *testing.T) {
	tests := []struct {
		name    string
		want    expr.Literal
		wantErr assert.ErrorAssertionFunc
	}{
		{"P0Y", createIntervalYearsLiteral(0, 0), assert.NoError},
		{"P2M", createIntervalYearsLiteral(0, 2), assert.NoError},
		{"P10Y", createIntervalYearsLiteral(10, 0), assert.NoError},
		{"P10M", createIntervalYearsLiteral(0, 10), assert.NoError},
		{"P8Y9M", createIntervalYearsLiteral(8, 9), assert.NoError},
		{"PY1M", nil, assert.Error},
		{"P", nil, assert.Error},
		{"PYM", nil, assert.Error},
		{"PXZ", nil, assert.Error},
		{"PXmM", nil, assert.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewIntervalYearsToMonthFromString(tt.name)
			if !tt.wantErr(t, err, fmt.Sprintf("NewIntervalYearsToMonth(%v)", tt.name)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewIntervalYearsToMonth(%v)", tt.name)
		})
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
			got := NewString(tt.value)
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

func TestNewPrecisionTimestampFromTime(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name      string
		precision types.TimePrecision
		tm        time.Time
		want      expr.Literal
	}{
		//{"zero", types.PrecisionSeconds, time.Unix(0, 0), expr.NewPrimitiveLiteral(types.Timestamp(0), false), assert.NoError},
		{"nowInSecs", types.PrecisionSeconds, now, &expr.ProtoLiteral{Value: now.Unix(), Type: &types.PrecisionTimestampType{Precision: types.PrecisionSeconds, Nullability: types.NullabilityRequired}}},
		{"nowInDeciSecs", types.PrecisionDeciSeconds, now, &expr.ProtoLiteral{Value: now.UnixMilli() / 100, Type: &types.PrecisionTimestampType{Precision: types.PrecisionDeciSeconds, Nullability: types.NullabilityRequired}}},
		{"nowInCentiSecs", types.PrecisionCentiSeconds, now, &expr.ProtoLiteral{Value: now.UnixMilli() / 10, Type: &types.PrecisionTimestampType{Precision: types.PrecisionCentiSeconds, Nullability: types.NullabilityRequired}}},
		{"nowInMilliSecs", types.PrecisionMilliSeconds, now, &expr.ProtoLiteral{Value: now.UnixMilli(), Type: &types.PrecisionTimestampType{Precision: types.PrecisionMilliSeconds, Nullability: types.NullabilityRequired}}},
		{"nowIn100MicroSecs", types.PrecisionEMinus4Seconds, now, &expr.ProtoLiteral{Value: now.UnixMicro() / 100, Type: &types.PrecisionTimestampType{Precision: types.PrecisionEMinus4Seconds, Nullability: types.NullabilityRequired}}},
		{"nowIn10MicroSecs", types.PrecisionEMinus5Seconds, now, &expr.ProtoLiteral{Value: now.UnixMicro() / 10, Type: &types.PrecisionTimestampType{Precision: types.PrecisionEMinus5Seconds, Nullability: types.NullabilityRequired}}},
		{"nowInMicros", types.PrecisionMicroSeconds, now, &expr.ProtoLiteral{Value: now.UnixMicro(), Type: &types.PrecisionTimestampType{Precision: types.PrecisionMicroSeconds, Nullability: types.NullabilityRequired}}},
		{"nowIn100NanoSecs", types.PrecisionEMinus7Seconds, now, &expr.ProtoLiteral{Value: now.UnixNano() / 100, Type: &types.PrecisionTimestampType{Precision: types.PrecisionEMinus7Seconds, Nullability: types.NullabilityRequired}}},
		{"nowIn10NanoSecs", types.PrecisionEMinus8Seconds, now, &expr.ProtoLiteral{Value: now.UnixNano() / 10, Type: &types.PrecisionTimestampType{Precision: types.PrecisionEMinus8Seconds, Nullability: types.NullabilityRequired}}},
		{"nowInNanoSecs", types.PrecisionNanoSeconds, now, &expr.ProtoLiteral{Value: now.UnixNano(), Type: &types.PrecisionTimestampType{Precision: types.PrecisionNanoSeconds, Nullability: types.NullabilityRequired}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPrecisionTimestampFromTime(tt.precision, tt.tm)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
			assert.True(t, tt.want.Equals(got))
		})
	}
}

func TestNewPrecisionTimestampTz(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name      string
		precision types.TimePrecision
		tm        time.Time
		want      expr.Literal
	}{
		//{"zero", types.PrecisionSeconds, time.Unix(0, 0), expr.NewPrimitiveLiteral(types.Timestamp(0), false), assert.NoError},
		{"nowInSecs", types.PrecisionSeconds, now, &expr.ProtoLiteral{Value: now.Unix(), Type: &types.PrecisionTimestampTzType{PrecisionTimestampType: types.PrecisionTimestampType{Precision: types.PrecisionSeconds, Nullability: types.NullabilityRequired}}}},
		{"nowInDeciSecs", types.PrecisionDeciSeconds, now, &expr.ProtoLiteral{Value: now.UnixMilli() / 100, Type: &types.PrecisionTimestampTzType{PrecisionTimestampType: types.PrecisionTimestampType{Precision: types.PrecisionDeciSeconds, Nullability: types.NullabilityRequired}}}},
		{"nowInCentiSecs", types.PrecisionCentiSeconds, now, &expr.ProtoLiteral{Value: now.UnixMilli() / 10, Type: &types.PrecisionTimestampTzType{PrecisionTimestampType: types.PrecisionTimestampType{Precision: types.PrecisionCentiSeconds, Nullability: types.NullabilityRequired}}}},
		{"nowInMilliSecs", types.PrecisionMilliSeconds, now, &expr.ProtoLiteral{Value: now.UnixMilli(), Type: &types.PrecisionTimestampTzType{PrecisionTimestampType: types.PrecisionTimestampType{Precision: types.PrecisionMilliSeconds, Nullability: types.NullabilityRequired}}}},
		{"nowIn100MicroSecs", types.PrecisionEMinus4Seconds, now, &expr.ProtoLiteral{Value: now.UnixMicro() / 100, Type: &types.PrecisionTimestampTzType{PrecisionTimestampType: types.PrecisionTimestampType{Precision: types.PrecisionEMinus4Seconds, Nullability: types.NullabilityRequired}}}},
		{"nowIn10MicroSecs", types.PrecisionEMinus5Seconds, now, &expr.ProtoLiteral{Value: now.UnixMicro() / 10, Type: &types.PrecisionTimestampTzType{PrecisionTimestampType: types.PrecisionTimestampType{Precision: types.PrecisionEMinus5Seconds, Nullability: types.NullabilityRequired}}}},
		{"nowInMicros", types.PrecisionMicroSeconds, now, &expr.ProtoLiteral{Value: now.UnixMicro(), Type: &types.PrecisionTimestampTzType{PrecisionTimestampType: types.PrecisionTimestampType{Precision: types.PrecisionMicroSeconds, Nullability: types.NullabilityRequired}}}},
		{"nowIn100NanoSecs", types.PrecisionEMinus7Seconds, now, &expr.ProtoLiteral{Value: now.UnixNano() / 100, Type: &types.PrecisionTimestampTzType{PrecisionTimestampType: types.PrecisionTimestampType{Precision: types.PrecisionEMinus7Seconds, Nullability: types.NullabilityRequired}}}},
		{"nowIn10NanoSecs", types.PrecisionEMinus8Seconds, now, &expr.ProtoLiteral{Value: now.UnixNano() / 10, Type: &types.PrecisionTimestampTzType{PrecisionTimestampType: types.PrecisionTimestampType{Precision: types.PrecisionEMinus8Seconds, Nullability: types.NullabilityRequired}}}},
		{"nowInNanoSecs", types.PrecisionNanoSeconds, now, &expr.ProtoLiteral{Value: now.UnixNano(), Type: &types.PrecisionTimestampTzType{PrecisionTimestampType: types.PrecisionTimestampType{Precision: types.PrecisionNanoSeconds, Nullability: types.NullabilityRequired}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPrecisionTimestampTzFromTime(tt.precision, tt.tm)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
			assert.True(t, tt.want.Equals(got))
		})
	}
}

func TestNewTimestampFromString(t *testing.T) {
	tests := []struct {
		name    string
		want    expr.Literal
		wantErr assert.ErrorAssertionFunc
	}{
		{"1970-01-01T00:00:00Z", expr.NewPrimitiveLiteral(types.Timestamp(0), false), assert.NoError},
		{"2016-01-02T15:04:05Z", expr.NewPrimitiveLiteral(types.Timestamp(1451747045000000), false), assert.NoError},
		{"2016-01-02 15:04:05", nil, assert.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTimestampFromString(tt.name)
			if !tt.wantErr(t, err, fmt.Sprintf("NewTimestampFromString(%v)", tt.name)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewTimestampFromString(%v)", tt.name)
		})
	}
}

func TestNewTimestampTZFromString(t *testing.T) {
	tests := []struct {
		name    string
		want    expr.Literal
		wantErr assert.ErrorAssertionFunc
	}{
		{"1970-01-01T00:00:00Z", expr.NewPrimitiveLiteral(types.TimestampTz(0), false), assert.NoError},
		{"2016-01-02T15:04:05Z", expr.NewPrimitiveLiteral(types.TimestampTz(1451747045000000), false), assert.NoError},
		{"2016-01-02T15:04:05", expr.NewPrimitiveLiteral(types.TimestampTz(1451747045000000), false), assert.NoError},
		{"2016-01-02T15:04:05+00:00", expr.NewPrimitiveLiteral(types.TimestampTz(1451747045000000), false), assert.NoError},
		{"2016-01-02T15:04:05+00:00", expr.NewPrimitiveLiteral(types.TimestampTz(1451747045000000), false), assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTimestampTZFromString(tt.name)
			if !tt.wantErr(t, err, fmt.Sprintf("NewTimestampFromString(%v)", tt.name)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewTimestampFromString(%v)", tt.name)
		})
	}
}

func TestNewList(t *testing.T) {
	i8Lit1 := NewInt8(1)
	i8Lit2 := NewInt8(2)
	i32Lit1 := NewInt32(1)
	i32Lit2 := NewInt32(2)
	listLiteral, _ := expr.NewLiteral[expr.ListLiteralValue]([]expr.Literal{i8Lit1, i8Lit2}, false)
	int8Type := &types.Int8Type{Nullability: types.NullabilityRequired}
	int32Type := &types.Int32Type{Nullability: types.NullabilityRequired}
	int8ListType := &types.ListType{Type: int8Type, Nullability: types.NullabilityRequired}
	int32ListType := &types.ListType{Type: int32Type, Nullability: types.NullabilityRequired}
	listOfListType := &types.ListType{Type: int8ListType, Nullability: types.NullabilityRequired}
	tests := []struct {
		name       string
		elements   []expr.Literal
		expSuccess bool
		litType    types.Type
		wantErr    assert.ErrorAssertionFunc
	}{
		{"empty", []expr.Literal{}, false, nil, assert.Error},
		{"i32List", []expr.Literal{i8Lit1, i32Lit2}, false, nil, assert.Error},
		{"i8ListSingle", []expr.Literal{i8Lit1}, true, int8ListType, assert.NoError},
		{"listOfListSingle", []expr.Literal{listLiteral}, true, listOfListType, assert.NoError},
		{"i8List", []expr.Literal{i8Lit1, i8Lit2}, true, int8ListType, assert.NoError},
		{"i32List", []expr.Literal{i32Lit1, i32Lit2}, true, int32ListType, assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewList(tt.elements)
			if !tt.wantErr(t, err, fmt.Sprintf("NewList(%v)", tt.elements)) {
				return
			}
			if tt.expSuccess {
				want, err := expr.NewLiteral[expr.ListLiteralValue](tt.elements, false)
				require.NoError(t, err)
				assert.Equalf(t, want, got, "NewList(%v)", tt.elements)
				assert.Equalf(t, tt.litType, got.GetType(), "NewList(%v)", tt.elements)
			}
		})
	}
}

func TestNewDateFromString(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    expr.Literal
		wantErr assert.ErrorAssertionFunc
	}{
		{"empty", "", nil, assert.Error},
		{"1970-01-01", "1970-01-01", expr.NewPrimitiveLiteral(types.Date(0), false), assert.NoError},
		{"2021-01-01", "2021-01-01", expr.NewPrimitiveLiteral(types.Date(18628), false), assert.NoError},
		{"2021-01-01T00:00:00Z", "2021-01-01T00:00:00Z", nil, assert.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDateFromString(tt.value)
			if !tt.wantErr(t, err, fmt.Sprintf("NewDateFromString(%v)", tt.value)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewDateFromString(%v)", tt.value)
		})
	}
}

func TestNewTimeFromString(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    expr.Literal
		wantErr assert.ErrorAssertionFunc
	}{
		{"empty", "", nil, assert.Error},
		{"00:00:00", "00:00:00", expr.NewPrimitiveLiteral(types.Time(0), false), assert.NoError},
		{"01:02:03", "01:02:03", expr.NewPrimitiveLiteral(types.Time(3723000000), false), assert.NoError},
		{"01:02:03.456", "01:02:03.456", expr.NewPrimitiveLiteral(types.Time(3723456000), false), assert.NoError},
		{"01:02:03.456789", "01:02:03.456789", expr.NewPrimitiveLiteral(types.Time(3723456789), false), assert.NoError},
		{"01:02:03.456789012", "01:02:03.456789012", expr.NewPrimitiveLiteral(types.Time(3723456789), false), assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTimeFromString(tt.value)
			if !tt.wantErr(t, err, fmt.Sprintf("NewTimeFromString(%v)", tt.value)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewTimeFromString(%v)", tt.value)
		})
	}
}
