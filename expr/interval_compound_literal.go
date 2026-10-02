package expr

import (
	"fmt"

	"github.com/substrait-io/substrait-go/v9/types"
)

// IntervalCompoundLiteral creates an interval compound literal
type IntervalCompoundLiteral struct {
	Years              int32
	Months             int32
	Days               int32
	Seconds            int32
	SubSeconds         int64
	SubSecondPrecision types.TimePrecision
	Nullability        types.Nullability
}

func (m IntervalCompoundLiteral) getType() types.Type {
	return types.NewIntervalCompoundType().WithPrecision(m.SubSecondPrecision).WithNullability(m.Nullability)
}

func (IntervalCompoundLiteral) isRootRef()            {}
func (m IntervalCompoundLiteral) GetType() types.Type { return m.getType() }
func (m IntervalCompoundLiteral) String() string {
	return fmt.Sprintf("%s(%s)", m.getType(), m.ValueString())
}
func (m IntervalCompoundLiteral) ValueString() string {
	return fmt.Sprintf("%d years, %d months, %d days, %d seconds, %d subseconds",
		m.Years, m.Months, m.Days, m.Seconds, m.SubSeconds)
}
func (m IntervalCompoundLiteral) Equals(rhs Expression) bool {
	if other, ok := rhs.(IntervalCompoundLiteral); ok {
		return m.getType().Equals(other.GetType()) && (m == other)
	}
	return false
}

func (m IntervalCompoundLiteral) Visit(VisitFunc) Expression { return m }
func (IntervalCompoundLiteral) IsScalar() bool               { return true }
