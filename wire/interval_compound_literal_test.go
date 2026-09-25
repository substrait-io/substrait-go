// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestIntervalCompoundLiteralToProto(t *testing.T) {
	var (
		yearVal          int32 = 1
		monthVal         int32 = 10
		dayVal           int32 = 100
		secondsVal       int32 = 1000
		subSecondsVal    int64 = 10000
		negativeInt32Val int32 = -5
		negativeInt64Val int64 = -100000
	)
	precisionVal := types.PrecisionNanoSeconds
	nanoSecPrecision := &proto.Expression_Literal_IntervalDayToSecond_Precision{Precision: precisionVal.ToProtoVal()}
	nullable := true
	nullability := types.NullabilityNullable

	for _, tc := range []struct {
		name                      string
		inputLiteral              expr.IntervalCompoundLiteral
		expectedExpressionLiteral *proto.Expression_Literal_IntervalCompound_
	}{
		{"WithOnlyYearAndMonth",
			expr.IntervalCompoundLiteral{Nullability: nullability, Years: yearVal, Months: monthVal},
			&proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
				IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{Years: yearVal, Months: monthVal},
			}},
		},
		{"WithOnlyYearAndMonthNegativeVal",
			expr.IntervalCompoundLiteral{Nullability: nullability, Years: negativeInt32Val, Months: negativeInt32Val},
			&proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
				IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{Years: negativeInt32Val, Months: negativeInt32Val},
			}},
		},
		{"WithOnlyDayToSecond",
			expr.IntervalCompoundLiteral{Nullability: nullability, Days: dayVal, Seconds: secondsVal, SubSeconds: subSecondsVal, SubSecondPrecision: precisionVal},
			&proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
				IntervalDayToSecond: &proto.Expression_Literal_IntervalDayToSecond{
					Days: dayVal, Seconds: secondsVal, PrecisionMode: nanoSecPrecision, Subseconds: subSecondsVal,
				},
			}},
		},
		{"WithOnlyDayToSecondNegativeVal",
			expr.IntervalCompoundLiteral{Nullability: nullability, Days: negativeInt32Val, Seconds: negativeInt32Val, SubSeconds: negativeInt64Val, SubSecondPrecision: precisionVal},
			&proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
				IntervalDayToSecond: &proto.Expression_Literal_IntervalDayToSecond{
					Days: negativeInt32Val, Seconds: negativeInt32Val, PrecisionMode: nanoSecPrecision, Subseconds: negativeInt64Val,
				},
			}},
		},
		{"WithBothYearToMonthAndDayToSecond",
			expr.IntervalCompoundLiteral{Nullability: nullability, Years: yearVal, Months: monthVal, Days: dayVal, Seconds: secondsVal, SubSeconds: subSecondsVal, SubSecondPrecision: precisionVal},
			&proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
				IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{Years: yearVal, Months: monthVal},
				IntervalDayToSecond: &proto.Expression_Literal_IntervalDayToSecond{
					Days: dayVal, Seconds: secondsVal, PrecisionMode: nanoSecPrecision, Subseconds: subSecondsVal,
				},
			}},
		},
		{"WithBothYearToMonthAndDayToSecondAllNegativeVal",
			expr.IntervalCompoundLiteral{Nullability: nullability, Years: negativeInt32Val, Months: negativeInt32Val, Days: negativeInt32Val, Seconds: negativeInt32Val, SubSeconds: negativeInt64Val, SubSecondPrecision: precisionVal},
			&proto.Expression_Literal_IntervalCompound_{IntervalCompound: &proto.Expression_Literal_IntervalCompound{
				IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{Years: negativeInt32Val, Months: negativeInt32Val},
				IntervalDayToSecond: &proto.Expression_Literal_IntervalDayToSecond{
					Days: negativeInt32Val, Seconds: negativeInt32Val, PrecisionMode: nanoSecPrecision, Subseconds: negativeInt64Val,
				},
			}},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			expected := &proto.Expression{RexType: &proto.Expression_Literal_{Literal: &proto.Expression_Literal{LiteralType: tc.expectedExpressionLiteral, Nullable: nullable}}}
			got := ExprToProto(tc.inputLiteral)
			if diff := cmp.Diff(got, expected, protocmp.Transform()); diff != "" {
				t.Errorf("proto didn't match, diff:\n%v", diff)
			}
			funcArgProto := &proto.FunctionArgument{ArgType: &proto.FunctionArgument_Value{Value: got}}
			if diff := cmp.Diff(FuncArgToProto(tc.inputLiteral), funcArgProto, protocmp.Transform()); diff != "" {
				t.Errorf("func arg proto didn't match, diff:\n%v", diff)
			}
		})
	}
}
