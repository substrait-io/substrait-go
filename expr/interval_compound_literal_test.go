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
	yearVal          int32 = 1
	monthVal         int32 = 10
	dayVal           int32 = 100
	secondsVal       int32 = 1000
	subSecondsVal    int64 = 10000
	negativeInt32Val int32 = -5
	negativeInt64Val int64 = -100000
)

func TestDatePartBuilder(t *testing.T) {

	var origYearVal, origMonthVal, origDayVal, origSecondsVal int32
	var origSubSecondsVal int64
	origYearVal = 1
	origMonthVal = -2
	origDayVal = 3
	origSecondsVal = -4
	origNullabilityVal := types.NullabilityRequired
	origPrecisionVal := types.PrecisionEMinus7Seconds
	origSubSecondsVal = 5
	origLiteral := IntervalCompoundLiteral{Years: origYearVal, Months: origMonthVal, Days: origDayVal, Seconds: origSecondsVal, SubSeconds: origSubSecondsVal, SubSecondPrecision: origPrecisionVal, Nullability: origNullabilityVal}

	// verify that With* Method only change respective fields

	expectedUpdatedLiteral := origLiteral
	expectedUpdatedLiteral.Nullability = types.NullabilityNullable
	assert.Equal(t, origLiteral.WithNullability(types.NullabilityNullable), expectedUpdatedLiteral)

	expectedUpdatedLiteral = origLiteral
	expectedUpdatedLiteral.Years = 10
	assert.Equal(t, origLiteral.WithYears(10), expectedUpdatedLiteral)

	expectedUpdatedLiteral = origLiteral
	expectedUpdatedLiteral.Months = -20
	assert.Equal(t, origLiteral.WithMonths(-20), expectedUpdatedLiteral)

	expectedUpdatedLiteral = origLiteral
	expectedUpdatedLiteral.Days = 30
	assert.Equal(t, origLiteral.WithDays(30), expectedUpdatedLiteral)

	expectedUpdatedLiteral = origLiteral
	expectedUpdatedLiteral.Seconds = 40
	assert.Equal(t, origLiteral.WithSeconds(40), expectedUpdatedLiteral)

	expectedUpdatedLiteral = origLiteral
	expectedUpdatedLiteral.SubSecondPrecision = types.PrecisionDeciSeconds
	expectedUpdatedLiteral.SubSeconds = -50
	assert.Equal(t, origLiteral.WithSubSecond(-50, types.PrecisionDeciSeconds), expectedUpdatedLiteral)

	expectedUpdatedLiteral = origLiteral
	expectedUpdatedLiteral.SubSecondPrecision = types.PrecisionMilliSeconds
	expectedUpdatedLiteral.SubSeconds = -50
	assert.Equal(t, origLiteral.WithMilliSecond(-50), expectedUpdatedLiteral)

	expectedUpdatedLiteral = origLiteral
	expectedUpdatedLiteral.SubSecondPrecision = types.PrecisionMicroSeconds
	expectedUpdatedLiteral.SubSeconds = -50
	assert.Equal(t, origLiteral.WithMicroSecond(-50), expectedUpdatedLiteral)

	expectedUpdatedLiteral = origLiteral
	expectedUpdatedLiteral.SubSecondPrecision = types.PrecisionNanoSeconds
	expectedUpdatedLiteral.SubSeconds = -50
	assert.Equal(t, origLiteral.WithNanoSecond(-50), expectedUpdatedLiteral)
}

func TestIntervalCompoundToProto(t *testing.T) {
	// precision and nullability belong to type. In type unit tests they are already tested
	// for different values so no need to test for multiple values
	precisionVal := types.PrecisionNanoSeconds
	nanoSecPrecision := &proto.Expression_Literal_IntervalDayToSecond_Precision{Precision: precisionVal.ToProtoVal()}
	nullable := true
	nullability := types.NullabilityNullable
	literalWithNullability := IntervalCompoundLiteral{}.WithNullability(nullability)

	for _, tc := range []struct {
		name                      string
		inputLiteral              IntervalCompoundLiteral
		expectedExpressionLiteral *proto.Expression_Literal_IntervalCompound_
	}{
		{"WithOnlyYearAndMonth",
			literalWithNullability.WithYears(yearVal).WithMonths(monthVal),
			&proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
				IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{Years: yearVal, Months: monthVal},
			}},
		},
		{"WithOnlyYearAndMonthNegativeVal",
			literalWithNullability.WithYears(negativeInt32Val).WithMonths(negativeInt32Val),
			&proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
				IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{Years: negativeInt32Val, Months: negativeInt32Val},
			}},
		},
		{"WithOnlyDayToSecond",
			literalWithNullability.WithDays(dayVal).WithSeconds(secondsVal).WithSubSecond(subSecondsVal, precisionVal),
			&proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
				IntervalDayToSecond: &proto.Expression_Literal_IntervalDayToSecond{
					Days: dayVal, Seconds: secondsVal, PrecisionMode: nanoSecPrecision, Subseconds: subSecondsVal,
				},
			}},
		},
		{"WithOnlyDayToSecondNegativeVal",
			literalWithNullability.WithDays(negativeInt32Val).WithSeconds(negativeInt32Val).WithSubSecond(negativeInt64Val, precisionVal),
			&proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
				IntervalDayToSecond: &proto.Expression_Literal_IntervalDayToSecond{
					Days: negativeInt32Val, Seconds: negativeInt32Val, PrecisionMode: nanoSecPrecision, Subseconds: negativeInt64Val,
				},
			}},
		},
		{"WithBothYearToMonthAndDayToSecond",
			literalWithNullability.WithYears(yearVal).WithMonths(monthVal).WithDays(dayVal).WithSeconds(secondsVal).WithSubSecond(subSecondsVal, precisionVal),
			&proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
				IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{Years: yearVal, Months: monthVal},
				IntervalDayToSecond: &proto.Expression_Literal_IntervalDayToSecond{
					Days: dayVal, Seconds: secondsVal, PrecisionMode: nanoSecPrecision, Subseconds: subSecondsVal,
				},
			}},
		},
		{"WithBothYearToMonthAndDayToSecondAllNegativeVal",
			literalWithNullability.WithYears(negativeInt32Val).WithMonths(negativeInt32Val).WithDays(negativeInt32Val).WithSeconds(negativeInt32Val).WithSubSecond(negativeInt64Val, precisionVal),
			&proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
				IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{Years: negativeInt32Val, Months: negativeInt32Val},
				IntervalDayToSecond: &proto.Expression_Literal_IntervalDayToSecond{
					Days: negativeInt32Val, Seconds: negativeInt32Val, PrecisionMode: nanoSecPrecision, Subseconds: negativeInt64Val,
				},
			}},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			expectedProtoExpression := &proto.Expression{RexType: &proto.Expression_Literal_{Literal: &proto.Expression_Literal{LiteralType: tc.expectedExpressionLiteral, Nullable: nullable}}}
			gotExpressionProto := tc.inputLiteral.ToProto()
			assert.NotNil(t, gotExpressionProto)
			if diff := cmp.Diff(gotExpressionProto, expectedProtoExpression, protocmp.Transform()); diff != "" {
				t.Errorf("proto didn't match, diff:\n%v", diff)
			}
			// verify ToProtoFuncArg
			funcArgProto := &proto.FunctionArgument{
				ArgType: &proto.FunctionArgument_Value{Value: gotExpressionProto},
			}
			if diff := cmp.Diff(tc.inputLiteral.ToProtoFuncArg(), funcArgProto, protocmp.Transform()); diff != "" {
				t.Errorf("expression proto didn't match, diff:\n%v", diff)
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
	literalWithNullability := IntervalCompoundLiteral{}.WithNullability(nullability)
	for _, tc := range []struct {
		name             string
		constructedProto *proto.Expression_Literal
		expectedLiteral  IntervalCompoundLiteral
	}{
		{"NoPartsValue",
			&proto.Expression_Literal{
				LiteralType: &proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{}},
				Nullable:    nullable},
			literalWithNullability,
		},
		{"OnlyYearAndMonth",
			&proto.Expression_Literal{
				LiteralType: &proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
					IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{Years: yearVal, Months: monthVal},
				}},
				Nullable: nullable},
			literalWithNullability.WithYears(yearVal).WithMonths(monthVal),
		},
		{"OnlyYearAndMonthNegativeVal",
			&proto.Expression_Literal{
				LiteralType: &proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
					IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{Years: negativeInt32Val, Months: negativeInt32Val},
				}},
				Nullable: nullable},
			literalWithNullability.WithYears(negativeInt32Val).WithMonths(negativeInt32Val),
		},
		{"OnlyDayToSecond",
			&proto.Expression_Literal{
				LiteralType: &proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
					IntervalDayToSecond: &proto.Expression_Literal_IntervalDayToSecond{Days: dayVal, Seconds: secondsVal,
						PrecisionMode: nanoSecPrecision, Subseconds: subSecondsVal},
				}},
				Nullable: nullable},
			literalWithNullability.WithDays(dayVal).WithSeconds(secondsVal).WithSubSecond(subSecondsVal, precisionNanoVal),
		},
		{"OnlyDayToSecondNegativeVal",
			&proto.Expression_Literal{
				LiteralType: &proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
					IntervalDayToSecond: &proto.Expression_Literal_IntervalDayToSecond{Days: negativeInt32Val, Seconds: negativeInt32Val,
						PrecisionMode: nanoSecPrecision, Subseconds: negativeInt64Val},
				}},
				Nullable: nullable},
			literalWithNullability.WithDays(negativeInt32Val).WithSeconds(negativeInt32Val).WithSubSecond(negativeInt64Val, precisionNanoVal),
		},
		{"BothYearToMonthAndDayToSecond",
			&proto.Expression_Literal{
				LiteralType: &proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
					IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{Years: yearVal, Months: monthVal},
					IntervalDayToSecond: &proto.Expression_Literal_IntervalDayToSecond{Days: dayVal, Seconds: secondsVal,
						PrecisionMode: nanoSecPrecision, Subseconds: subSecondsVal},
				}},
				Nullable: nullable},
			literalWithNullability.WithYears(yearVal).WithMonths(monthVal).WithDays(dayVal).WithSeconds(secondsVal).WithSubSecond(subSecondsVal, precisionNanoVal),
		},
		{"BothYearToMonthAndDayToSecondAllNegVal",
			&proto.Expression_Literal{
				LiteralType: &proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
					IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{Years: negativeInt32Val, Months: negativeInt32Val},
					IntervalDayToSecond: &proto.Expression_Literal_IntervalDayToSecond{Days: negativeInt32Val, Seconds: negativeInt32Val,
						PrecisionMode: nanoSecPrecision, Subseconds: negativeInt64Val},
				}},
				Nullable: nullable},
			literalWithNullability.WithYears(negativeInt32Val).WithMonths(negativeInt32Val).WithDays(negativeInt32Val).WithSeconds(negativeInt32Val).WithSubSecond(negativeInt64Val, precisionNanoVal),
		},
		{"WithDeprecatedMicroSecondPrecision",
			&proto.Expression_Literal{
				LiteralType: &proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
					IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{Years: yearVal, Months: monthVal},
					IntervalDayToSecond: &proto.Expression_Literal_IntervalDayToSecond{Days: dayVal, Seconds: secondsVal,
						PrecisionMode: deprecatedMicroSecPrecision},
				}},
				Nullable: nullable},
			literalWithNullability.WithYears(yearVal).WithMonths(monthVal).WithDays(dayVal).WithSeconds(secondsVal).WithMicroSecond(int64(microSecondVal)),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			gotLiteral := LiteralFromProto(tc.constructedProto)
			assert.NotNil(t, gotLiteral)
			assert.Equal(t, tc.expectedLiteral, gotLiteral)
			// verify equal method too returns true
			assert.True(t, tc.expectedLiteral.Equals(gotLiteral))
			assert.True(t, gotLiteral.IsScalar())
			// got literal after serialization is different from empty literal
			assert.False(t, IntervalCompoundLiteral{}.Equals(gotLiteral))
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
