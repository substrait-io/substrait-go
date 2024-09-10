package types

import (
	"fmt"

	"github.com/substrait-io/substrait-go/types/parameter_types"
)

// ParameterizedMapType is a struct having at least one of key or value of type ParameterizedAbstractType
// If All arguments are concrete they are represented by MapType
type ParameterizedMapType struct {
	Nullability      Nullability
	TypeVariationRef uint32
	Key              Type
	Value            Type
}

func (*ParameterizedMapType) isRootRef() {}
func (m *ParameterizedMapType) WithNullability(n Nullability) Type {
	m.Nullability = n
	return m
}

func (m *ParameterizedMapType) GetType() Type               { return m }
func (m *ParameterizedMapType) GetNullability() Nullability { return m.Nullability }
func (m *ParameterizedMapType) GetTypeVariationReference() uint32 {
	return m.TypeVariationRef
}
func (m *ParameterizedMapType) Equals(rhs Type) bool {
	if o, ok := rhs.(*ParameterizedMapType); ok {
		return m.Nullability == o.Nullability && m.TypeVariationRef == o.TypeVariationRef &&
			m.Key.Equals(o.Key) && m.Value.Equals(o.Value)
	}
	return false
}

func (m *ParameterizedMapType) ShortString() string {
	t := MapType{}
	return t.ShortString()
}

func (m *ParameterizedMapType) String() string {
	t := MapType{}
	parameterString := fmt.Sprintf("<%s, %s>", m.Key.String(), m.Value.String())
	return fmt.Sprintf("%s%s%s", t.BaseString(), strNullable(m), parameterString)
}

// GetAbstractParameters returns the abstract parameter names
// this implements interface ParameterizedAbstractType
func (m *ParameterizedMapType) GetAbstractParameters() []parameter_types.AbstractParameterType {
	var abstractParams []parameter_types.AbstractParameterType
	if abs, ok := m.Key.(parameter_types.AbstractParameterType); ok {
		abstractParams = append(abstractParams, abs)
	}
	if abs, ok := m.Value.(parameter_types.AbstractParameterType); ok {
		abstractParams = append(abstractParams, abs)
	}
	return abstractParams
}

// GetAbstractParamName this implements interface AbstractParameterType
// to indicate ParameterizedStructType itself can be used as a parameter of abstract type too
func (m *ParameterizedMapType) GetAbstractParamName() string {
	return m.String()
}
