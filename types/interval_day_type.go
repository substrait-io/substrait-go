package types

import (
	"fmt"
)

// IntervalDayType this is used to represent a type of interval day.
type IntervalDayType struct {
	Precision        TimePrecision
	TypeVariationRef uint32
	Nullability      Nullability
}

func (m *IntervalDayType) GetPrecisionProtoVal() int32 {
	return m.Precision.ToProtoVal()
}

func (*IntervalDayType) isRootRef() {}
func (m *IntervalDayType) WithNullability(n Nullability) Type {
	out := *m
	out.Nullability = n
	return &out
}

func (m *IntervalDayType) GetType() Type                     { return m }
func (m *IntervalDayType) GetNullability() Nullability       { return m.Nullability }
func (m *IntervalDayType) GetTypeVariationReference() uint32 { return m.TypeVariationRef }
func (m *IntervalDayType) Equals(rhs Type) bool {
	if o, ok := rhs.(*IntervalDayType); ok {
		return *o == *m
	}
	return false
}

func (*IntervalDayType) ShortString() string { return shortTypeNames[TypeNameIntervalDay] }

func (m *IntervalDayType) String() string {
	return fmt.Sprintf("%s%s<%d>", TypeNameIntervalDay, strNullable(m),
		m.Precision.ToProtoVal())
}

func (m *IntervalDayType) ParameterString() string {
	return fmt.Sprintf("%d", m.Precision.ToProtoVal())
}

func (s *IntervalDayType) BaseString() string {
	return string(TypeNameIntervalDay)
}

func (m *IntervalDayType) GetPrecision() TimePrecision {
	return m.Precision
}

func (m *IntervalDayType) GetReturnType(length int32, nullability Nullability) Type {
	out := *m
	out.Precision = TimePrecision(length)
	out.Nullability = nullability
	return &out
}

func (m *IntervalDayType) GetParameters() []interface{} {
	return []interface{}{m.Precision}
}
