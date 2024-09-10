package types

import (
	"fmt"

	"github.com/substrait-io/substrait-go/types/parameter_types"
)

// ParameterizedDecimalType is a decimal type with at least one of precision and scale parameters of string type
// example: Decimal<P,S> or Decimal<P,0>.
// Note concrete types e.g. Decimal(10, 2) are not represented by this type
// Concrete type is represented by DecimalType
type ParameterizedDecimalType struct {
	Nullability      Nullability
	TypeVariationRef uint32
	Precision        parameter_types.LeafParameter
	Scale            parameter_types.LeafParameter
}

func (*ParameterizedDecimalType) isRootRef() {}
func (m *ParameterizedDecimalType) WithNullability(n Nullability) Type {
	m.Nullability = n
	return m
}

func (m *ParameterizedDecimalType) GetType() Type               { return m }
func (m *ParameterizedDecimalType) GetNullability() Nullability { return m.Nullability }
func (m *ParameterizedDecimalType) GetTypeVariationReference() uint32 {
	return m.TypeVariationRef
}
func (m *ParameterizedDecimalType) Equals(rhs Type) bool {
	if o, ok := rhs.(*ParameterizedDecimalType); ok {
		return *o == *m
	}
	return false
}

func (m *ParameterizedDecimalType) ShortString() string {
	t := DecimalType{}
	return t.ShortString()
}

func (m *ParameterizedDecimalType) String() string {
	t := DecimalType{}
	parameterString := fmt.Sprintf("<%s,%s>", m.Precision.String(), m.Scale.String())
	return fmt.Sprintf("%s%s%s", t.BaseString(), strNullable(m), parameterString)
}

// GetAbstractParameters returns the abstract parameter names
// this implements interface ParameterizedAbstractType
func (m *ParameterizedDecimalType) GetAbstractParameters() []parameter_types.AbstractParameterType {
	var params []parameter_types.AbstractParameterType
	if p, ok := m.Precision.(parameter_types.AbstractParameterType); ok {
		params = append(params, p)
	}
	if p, ok := m.Scale.(parameter_types.AbstractParameterType); ok {
		params = append(params, p)
	}
	return params
}

// GetAbstractParamName this implements interface AbstractParameterType
// to indicate ParameterizedDecimalType itself can be used as a parameter of abstract type too
func (m *ParameterizedDecimalType) GetAbstractParamName() string {
	return m.String()
}
