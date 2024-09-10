package types

import (
	"fmt"

	"github.com/substrait-io/substrait-go/types/parameter_types"
)

// ParameterizedListType is a list type having parameter of ParameterizedAbstractType
// basically a list of which type is another abstract parameter
// example: List<Decimal(P,S)>. Kindly note concrete types List<Decimal(38, 0)> is not represented by this type
// Concrete type is represented by ListType
type ParameterizedListType struct {
	Nullability      Nullability
	TypeVariationRef uint32
	Type             ParameterizedAbstractType
}

func (*ParameterizedListType) isRootRef() {}
func (m *ParameterizedListType) WithNullability(n Nullability) Type {
	m.Nullability = n
	return m
}

func (m *ParameterizedListType) GetType() Type               { return m }
func (m *ParameterizedListType) GetNullability() Nullability { return m.Nullability }
func (m *ParameterizedListType) GetTypeVariationReference() uint32 {
	return m.TypeVariationRef
}
func (m *ParameterizedListType) Equals(rhs Type) bool {
	if o, ok := rhs.(*ParameterizedListType); ok {
		return m.Nullability == o.Nullability && m.TypeVariationRef == o.TypeVariationRef &&
			m.Type.Equals(o.Type)
	}
	return false
}

func (m *ParameterizedListType) ShortString() string {
	t := ListType{}
	return t.ShortString()
}

func (m *ParameterizedListType) String() string {
	t := ListType{}
	parameterString := fmt.Sprintf("<%s>", m.Type)
	return fmt.Sprintf("%s%s%s", t.BaseString(), strNullable(m), parameterString)
}

// GetAbstractParameters returns the abstract parameter names
// this implements interface ParameterizedAbstractType
func (m *ParameterizedListType) GetAbstractParameters() []parameter_types.AbstractParameterType {
	return []parameter_types.AbstractParameterType{m.Type.(parameter_types.AbstractParameterType)}
}

// GetAbstractParamName this implements interface AbstractParameterType
// to indicate ParameterizedListType itself can be used as a parameter of abstract type too
func (m *ParameterizedListType) GetAbstractParamName() string {
	return m.String()
}
