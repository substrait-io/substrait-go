package expr

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/proto"
	"github.com/substrait-io/substrait-go/types"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestIntervalYearToMonthToProto(t *testing.T) {
	// nullability belong to type. In type unit tests they are already tested
	// for different values so no need to test for multiple values
	nullable := true
	nullability := types.NullabilityNullable
	var oneYear int32 = 1
	var oneMonth int32 = 1

	for _, tc := range []struct {
		name               string
		literal            Literal
		expectedExpression *proto.Expression
	}{
		{"WithOnlyYear",
			IntervalYearToMonthLiteral{}.WithYear(oneYear).WithNullability(nullability),
			&proto.Expression{
				RexType: &proto.Expression_Literal_{Literal: &proto.Expression_Literal{
					LiteralType: &proto.Expression_Literal_IntervalYearToMonth_{
						IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{Years: oneYear}},
					Nullable: nullable,
				}},
			},
		},
		{"WithOnlyMonth",
			IntervalYearToMonthLiteral{}.WithMonth(oneMonth).WithNullability(nullability),
			&proto.Expression{
				RexType: &proto.Expression_Literal_{Literal: &proto.Expression_Literal{
					LiteralType: &proto.Expression_Literal_IntervalYearToMonth_{
						IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{Months: oneMonth}},
					Nullable: nullable,
				}},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			toProto := tc.literal.ToProto()
			if diff := cmp.Diff(toProto, tc.expectedExpression, protocmp.Transform()); diff != "" {
				t.Errorf("expression proto didn't match, diff:\n%v", diff)
			}
			// verify ToProtoFuncArg
			funcArgProto := &proto.FunctionArgument{
				ArgType: &proto.FunctionArgument_Value{Value: toProto},
			}
			if diff := cmp.Diff(tc.literal.ToProtoFuncArg(), funcArgProto, protocmp.Transform()); diff != "" {
				t.Errorf("expression proto didn't match, diff:\n%v", diff)
			}
		})

	}
}

func TestIntervalYearToMonthFromProto(t *testing.T) {
	nullable := true
	nullability := types.NullabilityNullable
	var oneYear int32 = 1
	var oneMonth int32 = 1
	for _, tc := range []struct {
		name            string
		inputProto      *proto.Expression_Literal
		expectedLiteral IntervalYearToMonthLiteral
	}{
		{"OnlyYearToMonth",
			&proto.Expression_Literal{
				LiteralType: &proto.Expression_Literal_IntervalYearToMonth_{
					IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{Years: oneYear, Months: oneMonth}},
				Nullable: nullable},
			IntervalYearToMonthLiteral{Years: oneYear, Months: oneMonth, Nullability: nullability},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			gotLiteral := intervalYearToMonthLiteralFromProto(tc.inputProto)
			assert.Equal(t, tc.expectedLiteral, gotLiteral)
			// verify equal method too returns true
			assert.True(t, tc.expectedLiteral.Equals(gotLiteral))
			assert.True(t, gotLiteral.IsScalar())
			// got literal after serialization is different from empty literal
			assert.False(t, IntervalYearToMonthLiteral{}.Equals(gotLiteral))
		})

	}
}
