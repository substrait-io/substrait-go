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

func TestLiteralFromProtoLiteral(t *testing.T) {
	intDayToSecVal := &proto.Expression_Literal_IntervalDayToSecond{Days: 1, Seconds: 2, PrecisionMode: &proto.Expression_Literal_IntervalDayToSecond_Precision{Precision: 5}}
	for _, tc := range []struct {
		name             string
		constructedProto *proto.Expression_Literal
		expectedLiteral  interface{}
	}{
		{"TimeType",
			&proto.Expression_Literal{LiteralType: &proto.Expression_Literal_PrecisionTime_{PrecisionTime: &proto.Expression_Literal_PrecisionTime{Precision: 4, Value: 12345678}}, Nullable: true},
			&expr.ProtoLiteral{Value: int64(12345678), Type: types.NewPrecisionTimeType(types.PrecisionEMinus4Seconds).WithNullability(types.NullabilityNullable)},
		},
		{"TimeStampType",
			&proto.Expression_Literal{LiteralType: &proto.Expression_Literal_PrecisionTimestamp_{PrecisionTimestamp: &proto.Expression_Literal_PrecisionTimestamp{Precision: 4, Value: 12345678}}, Nullable: true},
			&expr.ProtoLiteral{Value: int64(12345678), Type: types.NewPrecisionTimestampType(types.PrecisionEMinus4Seconds).WithNullability(types.NullabilityNullable)},
		},
		{"PreEpochTimeStampType",
			&proto.Expression_Literal{LiteralType: &proto.Expression_Literal_PrecisionTimestamp_{PrecisionTimestamp: &proto.Expression_Literal_PrecisionTimestamp{Precision: 9, Value: -1}}, Nullable: true},
			&expr.ProtoLiteral{Value: int64(-1), Type: types.NewPrecisionTimestampType(types.PrecisionNanoSeconds).WithNullability(types.NullabilityNullable)},
		},
		{"TimeStampTzType",
			&proto.Expression_Literal{LiteralType: &proto.Expression_Literal_PrecisionTimestampTz{PrecisionTimestampTz: &proto.Expression_Literal_PrecisionTimestamp{Precision: 9, Value: 12345678}}, Nullable: true},
			&expr.ProtoLiteral{Value: int64(12345678), Type: types.NewPrecisionTimestampTzType(types.PrecisionNanoSeconds).WithNullability(types.NullabilityNullable)},
		},
		{"PreEpochTimeStampTzType",
			&proto.Expression_Literal{LiteralType: &proto.Expression_Literal_PrecisionTimestampTz{PrecisionTimestampTz: &proto.Expression_Literal_PrecisionTimestamp{Precision: 9, Value: -1}}, Nullable: true},
			&expr.ProtoLiteral{Value: int64(-1), Type: types.NewPrecisionTimestampTzType(types.PrecisionNanoSeconds).WithNullability(types.NullabilityNullable)},
		},
		{"IntervalDayType",
			&proto.Expression_Literal{LiteralType: &proto.Expression_Literal_IntervalDayToSecond_{IntervalDayToSecond: intDayToSecVal}, Nullable: true},
			&expr.ProtoLiteral{Value: &types.IntervalDayToSecond{Days: 1, Seconds: 2, Precision: types.PrecisionEMinus5Seconds}, Type: &types.IntervalDayType{Precision: types.PrecisionEMinus5Seconds, Nullability: types.NullabilityNullable}},
		},
		{"IntervalYearToMonthType",
			&proto.Expression_Literal{LiteralType: &proto.Expression_Literal_IntervalYearToMonth_{IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{Years: 1234, Months: 5}}, Nullable: true},
			expr.IntervalYearToMonthLiteral{Years: 1234, Months: 5, Nullability: types.NullabilityNullable},
		},
		{"IntervalCompoundType",
			&proto.Expression_Literal{LiteralType: &proto.Expression_Literal_IntervalCompound_{
				IntervalCompound: &proto.Expression_Literal_IntervalCompound{
					IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{Years: 1234, Months: -5},
					IntervalDayToSecond: &proto.Expression_Literal_IntervalDayToSecond{Days: 6, Seconds: -7,
						PrecisionMode: &proto.Expression_Literal_IntervalDayToSecond_Precision{Precision: 8},
						Subseconds:    -9,
					},
				}}, Nullable: true},
			expr.IntervalCompoundLiteral{Years: 1234, Months: -5, Days: 6, Seconds: -7, SubSecondPrecision: 8, SubSeconds: -9, Nullability: types.NullabilityNullable},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			literal := LiteralFromProto(tc.constructedProto)
			assert.Equal(t, tc.expectedLiteral, literal)
		})

	}
}

func TestProtoLiteralToProtoLiteral(t *testing.T) {
	for _, tc := range []struct {
		name                      string
		constructedLiteral        *expr.ProtoLiteral
		expectedExpressionLiteral *proto.Expression_Literal
	}{
		{"TimeType",
			&expr.ProtoLiteral{Value: int64(12345678), Type: types.NewPrecisionTimeType(types.PrecisionEMinus4Seconds).WithNullability(types.NullabilityNullable)},
			&proto.Expression_Literal{LiteralType: &proto.Expression_Literal_PrecisionTime_{PrecisionTime: &proto.Expression_Literal_PrecisionTime{Precision: 4, Value: 12345678}}, Nullable: true},
		},
		{"TimeStampType",
			&expr.ProtoLiteral{Value: int64(12345678), Type: types.NewPrecisionTimestampType(types.PrecisionEMinus4Seconds).WithNullability(types.NullabilityNullable)},
			&proto.Expression_Literal{LiteralType: &proto.Expression_Literal_PrecisionTimestamp_{PrecisionTimestamp: &proto.Expression_Literal_PrecisionTimestamp{Precision: 4, Value: 12345678}}, Nullable: true},
		},
		{"TimeStampTzType",
			&expr.ProtoLiteral{Value: int64(12345678), Type: types.NewPrecisionTimestampTzType(types.PrecisionNanoSeconds).WithNullability(types.NullabilityNullable)},
			&proto.Expression_Literal{LiteralType: &proto.Expression_Literal_PrecisionTimestampTz{PrecisionTimestampTz: &proto.Expression_Literal_PrecisionTimestamp{Precision: 9, Value: 12345678}}, Nullable: true},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if diff := cmp.Diff(LiteralToProto(tc.constructedLiteral), tc.expectedExpressionLiteral, protocmp.Transform()); diff != "" {
				t.Errorf("proto didn't match, diff:\n%v", diff)
			}
		})
	}
}

func TestMapLiteralToProtoLiteral(t *testing.T) {
	mapLit := expr.NewNestedLiteral(expr.MapLiteralValue{
		{Key: expr.NewPrimitiveLiteral("foo", false), Value: expr.NewPrimitiveLiteral[int32](1, false)},
		{Key: expr.NewPrimitiveLiteral("bar", false), Value: expr.NewPrimitiveLiteral[int32](2, false)},
	}, false)

	expected := &proto.Expression_Literal{
		LiteralType: &proto.Expression_Literal_Map_{
			Map: &proto.Expression_Literal_Map{
				KeyValues: []*proto.Expression_Literal_Map_KeyValue{
					{Key: LiteralToProto(expr.NewPrimitiveLiteral("foo", false)), Value: LiteralToProto(expr.NewPrimitiveLiteral[int32](1, false))},
					{Key: LiteralToProto(expr.NewPrimitiveLiteral("bar", false)), Value: LiteralToProto(expr.NewPrimitiveLiteral[int32](2, false))},
				},
			},
		},
	}
	if diff := cmp.Diff(LiteralToProto(mapLit), expected, protocmp.Transform()); diff != "" {
		t.Errorf("proto didn't match, diff:\n%v", diff)
	}
}
