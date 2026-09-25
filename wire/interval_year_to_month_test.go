// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestIntervalYearToMonthFromProto(t *testing.T) {
	nullable := true
	nullability := types.NullabilityNullable
	var oneYear int32 = 1
	var oneMonth int32 = 1
	for _, tc := range []struct {
		name            string
		inputProto      *proto.Expression_Literal
		expectedLiteral expr.IntervalYearToMonthLiteral
	}{
		{"OnlyYearToMonth",
			&proto.Expression_Literal{
				LiteralType: &proto.Expression_Literal_IntervalYearToMonth_{
					IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{Years: oneYear, Months: oneMonth}},
				Nullable: nullable},
			expr.IntervalYearToMonthLiteral{Years: oneYear, Months: oneMonth, Nullability: nullability},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			gotLiteral := intervalYearToMonthLiteralFromProto(tc.inputProto)
			assert.Equal(t, tc.expectedLiteral, gotLiteral)
			// verify equal method too returns true
			assert.True(t, tc.expectedLiteral.Equals(gotLiteral))
			assert.True(t, gotLiteral.IsScalar())
			// got literal after serialization is different from empty literal
			assert.False(t, expr.IntervalYearToMonthLiteral{}.Equals(gotLiteral))
		})

	}
}

func TestIntervalYearToMonthLiteralToProto(t *testing.T) {
	nullable := true
	nullability := types.NullabilityNullable
	var oneYear int32 = 1
	var oneMonth int32 = 1

	for _, tc := range []struct {
		name               string
		literal            expr.Literal
		expectedExpression *proto.Expression
	}{
		{"WithOnlyYear",
			expr.IntervalYearToMonthLiteral{Nullability: nullability, Years: oneYear},
			&proto.Expression{RexType: &proto.Expression_Literal_{Literal: &proto.Expression_Literal{
				LiteralType: &proto.Expression_Literal_IntervalYearToMonth_{
					IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{Years: oneYear}},
				Nullable: nullable,
			}}},
		},
		{"WithOnlyMonth",
			expr.IntervalYearToMonthLiteral{Nullability: nullability, Months: oneMonth},
			&proto.Expression{RexType: &proto.Expression_Literal_{Literal: &proto.Expression_Literal{
				LiteralType: &proto.Expression_Literal_IntervalYearToMonth_{
					IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{Months: oneMonth}},
				Nullable: nullable,
			}}},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := ExprToProto(tc.literal)
			if diff := cmp.Diff(got, tc.expectedExpression, protocmp.Transform()); diff != "" {
				t.Errorf("expression proto didn't match, diff:\n%v", diff)
			}
			funcArgProto := &proto.FunctionArgument{ArgType: &proto.FunctionArgument_Value{Value: got}}
			if diff := cmp.Diff(FuncArgToProto(tc.literal), funcArgProto, protocmp.Transform()); diff != "" {
				t.Errorf("func arg proto didn't match, diff:\n%v", diff)
			}
		})
	}
}
