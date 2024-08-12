package expr

import (
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/proto"
	"github.com/substrait-io/substrait-go/types"
	"google.golang.org/protobuf/testing/protocmp"
	"testing"
)

func TestIntervalYearToMonthToProto(t *testing.T) {
	// nullability belong to type. In type unit tests they are already tested
	// for different values so no need to test for multiple values
	nullable := true
	nullability := types.NullabilityNullable
	var oneYear int32 = 1
	var oneMonth int32 = 1

	for _, tc := range []struct {
		name                      string
		literal                   Literal
		expectedExpressionLiteral *proto.Expression_Literal_IntervalYearToMonth_
	}{
		{"WithOnlyYear",
			NewIntervalLiteralUptoMonth(nullability, oneYear, 0),
			&proto.Expression_Literal_IntervalYearToMonth_{IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{Years: oneYear}},
		},
		{"WithOnlyMonth",
			NewIntervalLiteralUptoMonth(nullability, 0, oneMonth),
			&proto.Expression_Literal_IntervalYearToMonth_{IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{Months: oneMonth}},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			expectedProtoExpression := &proto.Expression_Literal{LiteralType: tc.expectedExpressionLiteral, Nullable: nullable}
			if diff := cmp.Diff(tc.literal.ToProtoLiteral(), expectedProtoExpression, protocmp.Transform()); diff != "" {
				t.Errorf("proto didn't match, diff:\n%v", diff)
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
		name         string
		inputProto   *proto.Expression_Literal
		expectedVal  *intervalYearMonthVal
		expectedType types.Type
	}{
		{"OnlyYearToMonth",
			&proto.Expression_Literal{
				LiteralType: &proto.Expression_Literal_IntervalYearToMonth_{
					IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{Years: oneYear, Months: oneMonth}},
				Nullable: nullable},
			&intervalYearMonthVal{years: oneYear, months: oneMonth},
			types.NewIntervalYearToMonthType().WithNullability(nullability),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			expectedLiteral := &ProtoLiteral{Value: tc.expectedVal, Type: tc.expectedType}
			gotLiteral := LiteralFromProto(tc.inputProto)
			assert.Equal(t, expectedLiteral, gotLiteral)
		})

	}
}
