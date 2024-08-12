package expr

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/proto"
	"github.com/substrait-io/substrait-go/types"
	"google.golang.org/protobuf/testing/protocmp"
)

const (
	yearVal       int32 = 1
	monthVal      int32 = 10
	dayVal        int32 = 100
	secondsVal    int32 = 1000
	subSecondsVal int64 = 10000
)

func TestIntervalCompoundToProto(t *testing.T) {
	// precision and nullability belong to type. In type unit tests they are already tested
	// for different values so no need to test for multiple values
	precisonVal := types.PrecisionNanoSeconds
	nanoSecPrecision := &proto.Expression_Literal_IntervalDayToSecond_Precision{Precision: precisonVal.ToProtoVal()}
	nullable := true
	nullability := types.NullabilityNullable

	yearOption := WithIntervalCompoundYears(yearVal)
	monthOption := WithIntervalCompoundMonths(monthVal)
	dayOption := WithIntervalCompoundDays(dayVal)
	secondsOption := WithIntervalCompoundSeconds(secondsVal)
	subSecondsOption := WithIntervalCompoundSubSeconds(subSecondsVal)
	for _, tc := range []struct {
		name                      string
		options                   []intervalDatePartsOptions
		expectedExpressionLiteral *proto.Expression_Literal_IntervalCompound_
	}{
		{"WithOnlyYearAndMonth",
			[]intervalDatePartsOptions{yearOption, monthOption},
			&proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
				IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{Years: yearVal, Months: monthVal},
				IntervalDayToSecond: &proto.Expression_Literal_IntervalDayToSecond{PrecisionMode: nanoSecPrecision},
			}},
		},
		{"WithOnlyDayToSecond",
			[]intervalDatePartsOptions{dayOption, secondsOption, subSecondsOption},
			&proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
				IntervalDayToSecond: &proto.Expression_Literal_IntervalDayToSecond{
					Days:          dayVal,
					Seconds:       secondsVal,
					PrecisionMode: nanoSecPrecision,
					Subseconds:    subSecondsVal,
				},
			}},
		},
		{"WithBothYearToMonthAndDayToSecond",
			[]intervalDatePartsOptions{yearOption, monthOption, dayOption, secondsOption, subSecondsOption},
			&proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
				IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{Years: 1, Months: 10},
				IntervalDayToSecond: &proto.Expression_Literal_IntervalDayToSecond{
					Days:          dayVal,
					Seconds:       secondsVal,
					PrecisionMode: nanoSecPrecision,
					Subseconds:    subSecondsVal,
				},
			}},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			expectedProtoExpression := &proto.Expression_Literal{LiteralType: tc.expectedExpressionLiteral, Nullable: nullable}
			intervalCompoundLiteral := NewIntervalLiteralUptoSubSecondPrecision(precisonVal, nullability, tc.options...)
			if diff := cmp.Diff(intervalCompoundLiteral.ToProtoLiteral(), expectedProtoExpression, protocmp.Transform()); diff != "" {
				t.Errorf("proto didn't match, diff:\n%v", diff)
			}
		})

	}
}

func TestIntervalCompoundFromProto(t *testing.T) {
	precisionNanoVal := types.PrecisionNanoSeconds
	nanoSecPrecision := &proto.Expression_Literal_IntervalDayToSecond_Precision{Precision: precisionNanoVal.ToProtoVal()}

	var microSecondVal int32 = 70
	deprecatedMicroSecPrecision := &proto.Expression_Literal_IntervalDayToSecond_Microseconds{
		Microseconds: microSecondVal}
	nullable := true
	nullability := types.NullabilityNullable
	for _, tc := range []struct {
		name             string
		constructedProto *proto.Expression_Literal
		expectedVal      *intervalDateParts
		expectedType     types.Type
	}{
		{"NoPartsValue",
			&proto.Expression_Literal{
				LiteralType: &proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
					IntervalDayToSecond: &proto.Expression_Literal_IntervalDayToSecond{PrecisionMode: nanoSecPrecision},
				}},
				Nullable: nullable},
			&intervalDateParts{},
			types.NewIntervalCompoundType(precisionNanoVal).WithNullability(nullability),
		},
		{"OnlyYearAndMonth",
			&proto.Expression_Literal{
				LiteralType: &proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
					IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{Years: yearVal, Months: monthVal},
					IntervalDayToSecond: &proto.Expression_Literal_IntervalDayToSecond{PrecisionMode: nanoSecPrecision},
				}},
				Nullable: nullable},
			&intervalDateParts{years: yearVal, months: monthVal},
			types.NewIntervalCompoundType(precisionNanoVal).WithNullability(nullability),
		},
		{"OnlyDayToSecond",
			&proto.Expression_Literal{
				LiteralType: &proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
					IntervalDayToSecond: &proto.Expression_Literal_IntervalDayToSecond{Days: dayVal, Seconds: secondsVal,
						PrecisionMode: nanoSecPrecision, Subseconds: subSecondsVal},
				}},
				Nullable: nullable},
			&intervalDateParts{days: dayVal, seconds: secondsVal, subSeconds: subSecondsVal},
			types.NewIntervalCompoundType(precisionNanoVal).WithNullability(nullability),
		},
		{"BothYearToMonthAndDayToSecond",
			&proto.Expression_Literal{
				LiteralType: &proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
					IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{Years: yearVal, Months: monthVal},
					IntervalDayToSecond: &proto.Expression_Literal_IntervalDayToSecond{Days: dayVal, Seconds: secondsVal,
						PrecisionMode: nanoSecPrecision, Subseconds: subSecondsVal},
				}},
				Nullable: nullable},
			&intervalDateParts{years: yearVal, months: monthVal, days: dayVal, seconds: secondsVal, subSeconds: subSecondsVal},
			types.NewIntervalCompoundType(precisionNanoVal).WithNullability(nullability),
		},
		{"WithDeprecatedMicroSecondPrecision",
			&proto.Expression_Literal{
				LiteralType: &proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
					IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{Years: yearVal, Months: monthVal},
					IntervalDayToSecond: &proto.Expression_Literal_IntervalDayToSecond{Days: dayVal, Seconds: secondsVal,
						PrecisionMode: deprecatedMicroSecPrecision},
				}},
				Nullable: nullable},
			&intervalDateParts{years: yearVal, months: monthVal, days: dayVal, seconds: secondsVal, subSeconds: 70},
			types.NewIntervalCompoundType(types.PrecisionMicroSeconds).WithNullability(nullability),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			expectedLiteral := &ProtoLiteral{Value: tc.expectedVal, Type: tc.expectedType}
			gotLiteral := LiteralFromProto(tc.constructedProto)
			assert.Equal(t, expectedLiteral, gotLiteral)
		})

	}
}

func TestIntervalCompoundFromProtoError(t *testing.T) {
	// valid precision val is [0, 9]
	invalidPrecision := &proto.Expression_Literal_IntervalDayToSecond_Precision{Precision: 10}
	deprecatedMicroSecPrecision := &proto.Expression_Literal_IntervalDayToSecond_Microseconds{Microseconds: 70}
	for _, tc := range []struct {
		name             string
		constructedProto *proto.Expression_Literal
	}{
		{"NoPrecisionMode",
			&proto.Expression_Literal{
				LiteralType: &proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
					IntervalDayToSecond: &proto.Expression_Literal_IntervalDayToSecond{},
				}}},
		},
		{"NoPrecisionModeButSubsecondsPresent",
			&proto.Expression_Literal{
				LiteralType: &proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
					IntervalDayToSecond: &proto.Expression_Literal_IntervalDayToSecond{Subseconds: 123},
				}}},
		},
		{"InvalidPrecisionMode",
			&proto.Expression_Literal{
				LiteralType: &proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
					IntervalDayToSecond: &proto.Expression_Literal_IntervalDayToSecond{PrecisionMode: invalidPrecision},
				}}},
		},
		{"DeprecatedMicrosecondPrecisionWithSubsecondsSet",
			&proto.Expression_Literal{
				LiteralType: &proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
					IntervalDayToSecond: &proto.Expression_Literal_IntervalDayToSecond{PrecisionMode: deprecatedMicroSecPrecision, Subseconds: 70},
				}}},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			gotLiteral := LiteralFromProto(tc.constructedProto)
			assert.Nil(t, gotLiteral)
		})

	}
}
