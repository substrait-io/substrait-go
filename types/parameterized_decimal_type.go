package types

import (
	"fmt"
)

// ParameterizedDecimalType is a decimal type with precision and scale parameters of string type
// example: Decimal(P,S). Kindly note concrete types Decimal(10, 2) are not represented by this type
// Concrete type is represented by DecimalType
type ParameterizedDecimalType struct {
	Nullability      Nullability
	TypeVariationRef uint32
	Precision        IntegerParam
	Scale            IntegerParam
}

func (ParameterizedDecimalType) isRootRef() {}
func (m ParameterizedDecimalType) WithNullability(n Nullability) Type {
	m.Nullability = n
	return m
}

func (m ParameterizedDecimalType) GetType() Type               { return m }
func (m ParameterizedDecimalType) GetNullability() Nullability { return m.Nullability }
func (m ParameterizedDecimalType) GetTypeVariationReference() uint32 {
	return m.TypeVariationRef
}
func (m ParameterizedDecimalType) Equals(rhs Type) bool {
	if o, ok := rhs.(ParameterizedDecimalType); ok {
		return o == m
	}
	return false
}

func (m ParameterizedDecimalType) ShortString() string {
	t := DecimalType{}
	return t.ShortString()
}

func (m ParameterizedDecimalType) String() string {
	t := DecimalType{}
	parameterString := fmt.Sprintf("<%s,%s>", m.Precision.String(), m.Scale.String())
	return fmt.Sprintf("%s%s%s", t.BaseString(), strNullable(m), parameterString)
}
