package expr

import (
	"github.com/substrait-io/substrait-go/proto"
	"github.com/substrait-io/substrait-go/types"
)

type intervalYearMonthVal struct {
	years  int32
	months int32
}

// NewIntervalLiteralUptoMonth creates an interval literal which allows setting only year and month
// arguments: nullable property (n), years and months
func NewIntervalLiteralUptoMonth(n types.Nullability, years int32, months int32) Literal {
	intervalCompoundType := types.NewIntervalYearToMonthType().WithNullability(n)
	intervalPartsVal := &intervalYearMonthVal{
		years:  years,
		months: months,
	}
	return &ProtoLiteral{
		Value: intervalPartsVal,
		Type:  intervalCompoundType,
	}
}

func intervalYearToMonthValToProto(idp *intervalYearMonthVal) *proto.Expression_Literal_IntervalYearToMonth {
	return &proto.Expression_Literal_IntervalYearToMonth{
		Years:  idp.years,
		Months: idp.months,
	}
}

func intervalYearToMonthFromProto(protoVal *proto.Expression_Literal_IntervalYearToMonth, nullability types.Nullability) Literal {
	return NewIntervalLiteralUptoMonth(nullability, protoVal.Years, protoVal.Months)
}
