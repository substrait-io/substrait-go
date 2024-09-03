package types

import (
	"fmt"
)

// IntegerParam represents a single integer parameter for a parameterized type
// Example: VARCHAR(L1) -> L1 is the integer parameter
type IntegerParam struct {
	Name string
}

func (m IntegerParam) Equals(o IntegerParam) bool {
	return m == o
}

func (m IntegerParam) String() string {
	return m.Name
}

// ParameterizedTypeSingleIntegerParam This is a generic type to represent parameterized type with a single integer parameter
type ParameterizedTypeSingleIntegerParam[T VarCharType | FixedCharType | FixedBinaryType | PrecisionTimestampType | PrecisionTimestampTzType] struct {
	Nullability      Nullability
	TypeVariationRef uint32
	IntegerOption    IntegerParam
}

func (m ParameterizedTypeSingleIntegerParam[T]) WithIntegerOption(integerOption IntegerParam) Type {
	m.IntegerOption = integerOption
	return m
}

func (ParameterizedTypeSingleIntegerParam[T]) isRootRef() {}
func (m ParameterizedTypeSingleIntegerParam[T]) WithNullability(n Nullability) Type {
	m.Nullability = n
	return m
}

func (m ParameterizedTypeSingleIntegerParam[T]) GetType() Type               { return m }
func (m ParameterizedTypeSingleIntegerParam[T]) GetNullability() Nullability { return m.Nullability }
func (m ParameterizedTypeSingleIntegerParam[T]) GetTypeVariationReference() uint32 {
	return m.TypeVariationRef
}
func (m ParameterizedTypeSingleIntegerParam[T]) Equals(rhs Type) bool {
	if o, ok := rhs.(ParameterizedTypeSingleIntegerParam[T]); ok {
		return o == m
	}
	return false
}

func (m ParameterizedTypeSingleIntegerParam[T]) ShortString() string {
	switch any(m).(type) {
	case ParameterizedVarCharType:
		t := VarCharType{}
		return t.ShortString()
	case ParameterizedFixedCharType:
		t := FixedCharType{}
		return t.ShortString()
	case ParameterizedFixedBinaryType:
		t := FixedBinaryType{}
		return t.ShortString()
	case ParameterizedPrecisionTimestampType:
		t := PrecisionTimestampType{}
		return t.ShortString()
	case ParameterizedPrecisionTimestampTzType:
		t := PrecisionTimestampTzType{}
		return t.ShortString()
	default:
		panic("unknown type")
	}
}

func (m ParameterizedTypeSingleIntegerParam[T]) String() string {
	return fmt.Sprintf("%s%s%s", m.baseString(), strNullable(m), m.parameterString())
}

func (m ParameterizedTypeSingleIntegerParam[T]) parameterString() string {
	return fmt.Sprintf("<%s>", m.IntegerOption.String())
}

func (m ParameterizedTypeSingleIntegerParam[T]) baseString() string {
	switch any(m).(type) {
	case ParameterizedVarCharType:
		t := VarCharType{}
		return t.BaseString()
	case ParameterizedFixedCharType:
		t := FixedCharType{}
		return t.BaseString()
	case ParameterizedFixedBinaryType:
		t := FixedBinaryType{}
		return t.BaseString()
	case ParameterizedPrecisionTimestampType:
		t := PrecisionTimestampType{}
		return t.BaseString()
	case ParameterizedPrecisionTimestampTzType:
		t := PrecisionTimestampTzType{}
		return t.BaseString()
	default:
		panic("unknown type")
	}
}
