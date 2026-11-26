// SPDX-License-Identifier: Apache-2.0

package substraitproto

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
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

func TestIntervalCompoundFromProto(t *testing.T) {
	precisionNanoVal := types.PrecisionNanoSeconds
	nanoSecPrecision := precisionNanoVal.ToProtoVal()
	nullable := true
	nullability := types.NullabilityNullable
	for _, tc := range []struct {
		name             string
		constructedProto *proto.Expression_Literal
		expectedLiteral  expr.IntervalCompoundLiteral
	}{
		{"NoPartsValue",
			&proto.Expression_Literal{
				LiteralType: &proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{}},
				Nullable:    nullable},
			expr.IntervalCompoundLiteral{Nullability: nullability},
		},
		{"OnlyYearAndMonth",
			&proto.Expression_Literal{
				LiteralType: &proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
					IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{Years: yearVal, Months: monthVal},
				}},
				Nullable: nullable},
			expr.IntervalCompoundLiteral{Nullability: nullability, Years: yearVal, Months: monthVal},
		},
		{"OnlyYearAndMonthNegativeVal",
			&proto.Expression_Literal{
				LiteralType: &proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
					IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{Years: negativeInt32Val, Months: negativeInt32Val},
				}},
				Nullable: nullable},
			expr.IntervalCompoundLiteral{Nullability: nullability, Years: negativeInt32Val, Months: negativeInt32Val},
		},
		{"OnlyDayToSecond",
			&proto.Expression_Literal{
				LiteralType: &proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
					IntervalDayToSecond: &proto.Expression_Literal_IntervalDayToSecond{Days: dayVal, Seconds: secondsVal,
						Precision: nanoSecPrecision, Subseconds: subSecondsVal},
				}},
				Nullable: nullable},
			expr.IntervalCompoundLiteral{Nullability: nullability, Days: dayVal, Seconds: secondsVal, SubSeconds: subSecondsVal, SubSecondPrecision: precisionNanoVal},
		},
		{"OnlyDayToSecondNegativeVal",
			&proto.Expression_Literal{
				LiteralType: &proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
					IntervalDayToSecond: &proto.Expression_Literal_IntervalDayToSecond{Days: negativeInt32Val, Seconds: negativeInt32Val,
						Precision: nanoSecPrecision, Subseconds: negativeInt64Val},
				}},
				Nullable: nullable},
			expr.IntervalCompoundLiteral{Nullability: nullability, Days: negativeInt32Val, Seconds: negativeInt32Val, SubSeconds: negativeInt64Val, SubSecondPrecision: precisionNanoVal},
		},
		{"BothYearToMonthAndDayToSecond",
			&proto.Expression_Literal{
				LiteralType: &proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
					IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{Years: yearVal, Months: monthVal},
					IntervalDayToSecond: &proto.Expression_Literal_IntervalDayToSecond{Days: dayVal, Seconds: secondsVal,
						Precision: nanoSecPrecision, Subseconds: subSecondsVal},
				}},
				Nullable: nullable},
			expr.IntervalCompoundLiteral{Nullability: nullability, Years: yearVal, Months: monthVal, Days: dayVal, Seconds: secondsVal, SubSeconds: subSecondsVal, SubSecondPrecision: precisionNanoVal},
		},
		{"BothYearToMonthAndDayToSecondAllNegVal",
			&proto.Expression_Literal{
				LiteralType: &proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
					IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{Years: negativeInt32Val, Months: negativeInt32Val},
					IntervalDayToSecond: &proto.Expression_Literal_IntervalDayToSecond{Days: negativeInt32Val, Seconds: negativeInt32Val,
						Precision: nanoSecPrecision, Subseconds: negativeInt64Val},
				}},
				Nullable: nullable},
			expr.IntervalCompoundLiteral{Nullability: nullability, Years: negativeInt32Val, Months: negativeInt32Val, Days: negativeInt32Val, Seconds: negativeInt32Val, SubSeconds: negativeInt64Val, SubSecondPrecision: precisionNanoVal},
		},
		{"DayToSecondWithSecondsPrecision",
			&proto.Expression_Literal{
				LiteralType: &proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
					IntervalDayToSecond: &proto.Expression_Literal_IntervalDayToSecond{Days: dayVal, Seconds: secondsVal},
				}},
				Nullable: nullable},
			expr.IntervalCompoundLiteral{Nullability: nullability, Days: dayVal, Seconds: secondsVal, SubSecondPrecision: types.PrecisionSeconds},
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
			assert.False(t, expr.IntervalCompoundLiteral{}.Equals(gotLiteral))
		})

	}
}

func TestIntervalCompoundFromProtoError(t *testing.T) {
	// valid precision val is [0, 9]
	for _, tc := range []struct {
		name             string
		constructedProto *proto.Expression_Literal
	}{
		{"InvalidPrecision",
			&proto.Expression_Literal{
				LiteralType: &proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
					IntervalDayToSecond: &proto.Expression_Literal_IntervalDayToSecond{Precision: 10},
				}}},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			gotLiteral := LiteralFromProto(tc.constructedProto)
			assert.Nil(t, gotLiteral)
		})

	}
}
