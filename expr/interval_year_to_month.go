package expr

import (
	"fmt"

	"github.com/substrait-io/substrait-go/v9/types"
)

// IntervalYearToMonthLiteral implements Literal interface for interval year to month type
type IntervalYearToMonthLiteral struct {
	Years       int32
	Months      int32
	Nullability types.Nullability
}

func (m IntervalYearToMonthLiteral) getType() types.Type {
	return types.NewIntervalYearToMonthType().WithNullability(m.Nullability)
}

func (IntervalYearToMonthLiteral) isRootRef()            {}
func (m IntervalYearToMonthLiteral) GetType() types.Type { return m.getType() }
func (m IntervalYearToMonthLiteral) String() string {
	return fmt.Sprintf("%s(%s)", m.getType(), m.ValueString())
}
func (m IntervalYearToMonthLiteral) ValueString() string {
	return fmt.Sprintf("%d years, %d months", m.Years, m.Months)
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
