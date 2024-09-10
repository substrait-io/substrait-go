package types

import (
	"fmt"

	"github.com/substrait-io/substrait-go/types/parameter_types"
)

// parameterizedTypeSingleIntegerParam This is a generic type to represent parameterized type with a single integer parameter
type parameterizedTypeSingleIntegerParam[T VarCharType | FixedCharType | FixedBinaryType | PrecisionTimestampType | PrecisionTimestampTzType] struct {
	Nullability      Nullability
	TypeVariationRef uint32
	IntegerOption    parameter_types.LeafIntParamAbstractType
}

func (m parameterizedTypeSingleIntegerParam[T]) WithIntegerOption(integerOption parameter_types.LeafIntParamAbstractType) Type {
	m.IntegerOption = integerOption
	return m
}

func (parameterizedTypeSingleIntegerParam[T]) isRootRef() {}
func (m parameterizedTypeSingleIntegerParam[T]) WithNullability(n Nullability) Type {
	m.Nullability = n
	return m
}

func (m parameterizedTypeSingleIntegerParam[T]) GetType() Type               { return m }
func (m parameterizedTypeSingleIntegerParam[T]) GetNullability() Nullability { return m.Nullability }
func (m parameterizedTypeSingleIntegerParam[T]) GetTypeVariationReference() uint32 {
	return m.TypeVariationRef
}
func (m parameterizedTypeSingleIntegerParam[T]) Equals(rhs Type) bool {
	if o, ok := rhs.(parameterizedTypeSingleIntegerParam[T]); ok {
		return o == m
	}
	return false
}

func (m parameterizedTypeSingleIntegerParam[T]) ShortString() string {
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

func (m parameterizedTypeSingleIntegerParam[T]) String() string {
	return fmt.Sprintf("%s%s%s", m.baseString(), strNullable(m), m.parameterString())
}

func (m parameterizedTypeSingleIntegerParam[T]) parameterString() string {
	return fmt.Sprintf("<%s>", m.IntegerOption.GetAbstractParamName())
}

func (m parameterizedTypeSingleIntegerParam[T]) baseString() string {
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

// GetAbstractParameters returns the abstract parameter names
// this implements interface ParameterizedAbstractType
func (m parameterizedTypeSingleIntegerParam[T]) GetAbstractParameters() []parameter_types.AbstractParameterType {
	return []parameter_types.AbstractParameterType{m.IntegerOption}
}

// GetAbstractParamName this implements interface AbstractParameterType
// basically, this type itself can be used as a parameter of abstract type too
func (m parameterizedTypeSingleIntegerParam[T]) GetAbstractParamName() string {
	return m.String()
}
