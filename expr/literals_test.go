package expr_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v3/expr"
	"github.com/substrait-io/substrait-go/v3/literal"
	"github.com/substrait-io/substrait-go/v3/types"
)

func TestNewDecimalWithType(t *testing.T) {
	tests := []struct {
		name           string
		precision      int32
		scale          int32
		decType        *types.DecimalType
		expStr         string
		expectedToFail bool
	}{
		{"123.45", 5, 2, &types.DecimalType{Nullability: types.NullabilityRequired, Precision: 10, Scale: 5}, "123.45000", false},
		{"12345.678", 8, 3, &types.DecimalType{Nullability: types.NullabilityNullable, Precision: 10, Scale: 5}, "12345.67800", false},
		{"12345", 5, 0, &types.DecimalType{Nullability: types.NullabilityNullable, Precision: 3, Scale: 2}, "", true},
		{"12345.888", 8, 3, &types.DecimalType{Nullability: types.NullabilityNullable, Precision: 7, Scale: 3}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lit, err := literal.NewDecimalFromString(tt.name)
			require.NoError(t, err)
			got, err := expr.NewDecimalWithType(lit.(*expr.ProtoLiteral), tt.decType)
			if tt.expectedToFail {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.expStr, got.ValueString())
		})
	}
}

func TestNewFixedLenWithType(t *testing.T) {
	tests := []struct {
		name      string
		inputType types.Type
		wantErr   bool
	}{
		{"abc", &types.VarCharType{Length: 5, Nullability: types.NullabilityRequired}, false},
		{"abcde", &types.VarCharType{Length: 3, Nullability: types.NullabilityRequired}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input, _ := literal.NewVarChar(tt.name)
			got, err := expr.NewVarCharWithType(input.(*expr.ProtoLiteral), tt.inputType.(*types.VarCharType))
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.inputType, got.GetType())
		})
	}
}

func TestNewPrecisionTimestampWithType(t *testing.T) {
	tests := []struct {
		name           string
		inputPrecision types.TimePrecision
		inputType      *types.PrecisionTimestampType
		want           expr.Literal
		wantErr        bool
	}{
		{"1991-01-01T01:02:03.456", 3, &types.PrecisionTimestampType{Precision: 3, Nullability: types.NullabilityNullable}, nil, false},
		{"1991-01-01T01:02:03.456", 3, &types.PrecisionTimestampType{Precision: 6, Nullability: types.NullabilityNullable}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lit, err := literal.NewPrecisionTimestampFromString(tt.inputPrecision, tt.name)
			require.NoError(t, err)
			got, err := expr.NewPrecisionTimestampWithType(lit.(*expr.ProtoLiteral), tt.inputType)
			if tt.wantErr {
				require.NoError(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.inputType, got.GetType())
			assert.Equal(t, tt.name, got.(types.IsoValuePrinter).IsoValueString())
		})
	}
}

func TestNewPrecisionTimestampTzWithType(t *testing.T) {
	tests := []struct {
		name             string
		inputPrecision   types.TimePrecision
		inputType        types.PrecisionTimestampType
		expLiteralString string
		wantErr          bool
	}{
		{"1991-01-01T01:02:03.456+05:30", 3, types.PrecisionTimestampType{Precision: 3, Nullability: types.NullabilityNullable}, "1990-12-31T19:32:03.456+00:00", false},
		{"1991-01-01T01:02:03.456+05:30", 3, types.PrecisionTimestampType{Precision: 6, Nullability: types.NullabilityRequired}, "1990-12-31T19:32:03.456+00:00", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lit, err := literal.NewPrecisionTimestampTzFromString(tt.inputPrecision, tt.name)
			require.NoError(t, err)
			inputType := &types.PrecisionTimestampTzType{PrecisionTimestampType: tt.inputType}
			got, err := expr.NewPrecisionTimestampTzWithType(lit.(*expr.ProtoLiteral), inputType)
			if tt.wantErr {
				require.NoError(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, inputType, got.GetType())
			assert.Equal(t, tt.expLiteralString, got.(types.IsoValuePrinter).IsoValueString())
		})
	}
}

func TestNewIntervalDayWithType(t *testing.T) {
	tests := []struct {
		name             string
		inputType        *types.IntervalDayType
		expLiteralString string
		wantErr          bool
	}{
		{"PT23H59M59.999S", &types.IntervalDayType{Precision: 3, Nullability: types.NullabilityNullable}, "PT23H59M59.999S", false},
		{"PT23H59M59.999S", &types.IntervalDayType{Precision: 2, Nullability: types.NullabilityNullable}, "PT23H59M59.99S", false},
		{"PT23H59M59.999S", &types.IntervalDayType{Precision: 6, Nullability: types.NullabilityRequired}, "PT23H59M59.999000S", false},
		{"PT23H59M59.999S", &types.IntervalDayType{Precision: 9, Nullability: types.NullabilityRequired}, "PT23H59M59.999000000S", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lit, err := literal.NewIntervalDaysToSecondFromString(tt.name)
			require.NoError(t, err)
			got, err := expr.NewIntervalDayWithType(lit.(*expr.ProtoLiteral), tt.inputType)
			if tt.wantErr {
				require.NoError(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.inputType, got.GetType())
			assert.Equal(t, tt.expLiteralString, got.(types.IsoValuePrinter).IsoValueString())
		})
	}
}

func TestProtoLiteral_WithType(t1 *testing.T) {
	dec123, _ := literal.NewDecimalFromString("123.45")
	iday, _ := literal.NewIntervalDaysToSecondFromString("PT23H59M59.999S")
	pts, _ := literal.NewPrecisionTimestampFromString(3, "1991-01-01T01:02:03.456")
	ptstz, _ := literal.NewPrecisionTimestampTzFromString(3, "1991-01-01T01:02:03.456")
	vchar, _ := literal.NewVarChar("sun")
	tests := []struct {
		name         string
		protoLiteral *expr.ProtoLiteral
		newType      types.Type
		want         expr.Literal
		wantErr      assert.ErrorAssertionFunc
	}{
		{"Decimal", dec123.(*expr.ProtoLiteral), &types.DecimalType{Nullability: types.NullabilityNullable, Precision: 10, Scale: 5}, nil, assert.NoError},
		{"IntervalDay", iday.(*expr.ProtoLiteral), &types.IntervalDayType{Precision: 3, Nullability: types.NullabilityNullable}, nil, assert.NoError},
		{"PrecisionTimestamp", pts.(*expr.ProtoLiteral), &types.PrecisionTimestampType{Precision: 3, Nullability: types.NullabilityNullable}, nil, assert.NoError},
		{"PrecisionTimestampTz", ptstz.(*expr.ProtoLiteral), &types.PrecisionTimestampTzType{PrecisionTimestampType: types.PrecisionTimestampType{Precision: 3, Nullability: types.NullabilityNullable}}, nil, assert.NoError},
		{"VarChar", vchar.(*expr.ProtoLiteral), &types.VarCharType{Length: 3, Nullability: types.NullabilityNullable}, nil, assert.NoError},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			got, err := tt.protoLiteral.WithType(tt.newType)
			if !tt.wantErr(t1, err, fmt.Sprintf("WithType(%v)", tt.newType)) {
				return
			}
			assert.Equalf(t1, tt.newType, got.GetType(), "WithType(%v)", tt.newType)
		})
	}
}

func TestByteSliceLiteral_WithType(t1 *testing.T) {
	fbin := expr.NewByteSliceLiteral[[]byte]([]byte{0x01, 0x02, 0x03}, false)
	uuid := expr.NewByteSliceLiteral[types.UUID]([]byte{0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xf, 0x10}, false)

	list := expr.NewNestedLiteral(expr.ListLiteralValue{
		literal.NewString("sun"), literal.NewString("moon"), literal.NewString("mars"),
	}, false)
	type testCase struct {
		name    string
		t       expr.WithTypeLiteral
		newType types.Type
		want    expr.Literal
		wantErr assert.ErrorAssertionFunc
	}
	tests := []testCase{
		{"FixedBinary", fbin, &types.FixedBinaryType{Length: 3, Nullability: types.NullabilityNullable}, nil, assert.NoError},
		{"UUID", uuid, &types.UUIDType{Nullability: types.NullabilityNullable}, nil, assert.NoError},
		{"List", list.(expr.WithTypeLiteral), &types.ListType{Type: &types.StringType{Nullability: types.NullabilityNullable}, Nullability: types.NullabilityNullable}, nil, assert.NoError},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			got, err := tt.t.WithType(tt.newType)
			if !tt.wantErr(t1, err, fmt.Sprintf("WithType(%v)", tt.newType)) {
				return
			}
			assert.Equalf(t1, tt.newType, got.GetType(), "WithType(%v)", tt.newType)
		})
	}
}
