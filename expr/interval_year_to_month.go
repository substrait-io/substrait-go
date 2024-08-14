package expr

import (
	"fmt"

	"github.com/substrait-io/substrait-go/proto"
	"github.com/substrait-io/substrait-go/types"
)

// IntervalYearToMonthLiteral implements Literal interface for interval year to month type
type IntervalYearToMonthLiteral struct {
	years       int32
	months      int32
	nullability types.Nullability
}

func (m IntervalYearToMonthLiteral) WithYear(years int32) IntervalYearToMonthLiteral {
	return IntervalYearToMonthLiteral{
		years:       years,
		months:      m.months,
		nullability: m.nullability,
	}
}

func (m IntervalYearToMonthLiteral) WithMonth(months int32) IntervalYearToMonthLiteral {
	return IntervalYearToMonthLiteral{
		years:       m.years,
		months:      months,
		nullability: m.nullability,
	}
}

func (m IntervalYearToMonthLiteral) WithNullability(nullability types.Nullability) IntervalYearToMonthLiteral {
	return IntervalYearToMonthLiteral{
		years:       m.years,
		months:      m.months,
		nullability: nullability,
	}
}

func (m IntervalYearToMonthLiteral) getType() types.Type {
	return types.NewIntervalYearToMonthType().WithNullability(m.nullability)
}

func (m IntervalYearToMonthLiteral) ToProtoLiteral() *proto.Expression_Literal {
	t := m.getType()
	return &proto.Expression_Literal{
		LiteralType: &proto.Expression_Literal_IntervalYearToMonth_{
			IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{
				Years:  m.years,
				Months: m.months,
			},
		},
		Nullable:               t.GetNullability() == types.NullabilityNullable,
		TypeVariationReference: t.GetTypeVariationReference(),
	}
}

func (m IntervalYearToMonthLiteral) ToProto() *proto.Expression {
	return &proto.Expression{RexType: &proto.Expression_Literal_{
		Literal: m.ToProtoLiteral(),
	}}
}

func intervalYearToMonthLiteralFromProto(l *proto.Expression_Literal) Literal {
	return IntervalYearToMonthLiteral{
		years:       l.GetIntervalYearToMonth().Years,
		months:      l.GetIntervalYearToMonth().Months,
		nullability: getNullability(l.Nullable),
	}
}

func (IntervalYearToMonthLiteral) isRootRef()            {}
func (m IntervalYearToMonthLiteral) GetType() types.Type { return m.getType() }
func (m IntervalYearToMonthLiteral) String() string {
	return fmt.Sprintf("%s(years:%d,months:%d)", m.getType(), m.years, m.months)
}
func (m IntervalYearToMonthLiteral) Equals(rhs Expression) bool {
	if other, ok := rhs.(IntervalYearToMonthLiteral); ok {
		return m.getType().Equals(other.GetType()) && (m == other)
	}
	return false
}

func (m IntervalYearToMonthLiteral) ToProtoFuncArg() *proto.FunctionArgument {
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Value{Value: m.ToProto()},
	}
}

func (m IntervalYearToMonthLiteral) Visit(VisitFunc) Expression { return m }
func (IntervalYearToMonthLiteral) IsScalar() bool               { return true }
