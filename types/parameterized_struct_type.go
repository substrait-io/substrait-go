package types

import (
	"fmt"
	"strings"

	"github.com/substrait-io/substrait-go/types/parameter_types"
)

// ParameterizedStructType is a struct having at least one parameter of type ParameterizedAbstractType
// example: Struct<Decimal(P,S), INT8>.
// If All arguments are concrete they are represented by StructType
type ParameterizedStructType struct {
	Nullability      Nullability
	TypeVariationRef uint32
	Type             []Type
}

func (*ParameterizedStructType) isRootRef() {}
func (m *ParameterizedStructType) WithNullability(n Nullability) Type {
	m.Nullability = n
	return m
}

func (m *ParameterizedStructType) GetType() Type               { return m }
func (m *ParameterizedStructType) GetNullability() Nullability { return m.Nullability }
func (m *ParameterizedStructType) GetTypeVariationReference() uint32 {
	return m.TypeVariationRef
}
func (m *ParameterizedStructType) Equals(rhs Type) bool {
	if o, ok := rhs.(*ParameterizedStructType); ok {
		if m.Nullability != o.Nullability || len(m.Type) != len(o.Type) ||
			m.TypeVariationRef != o.TypeVariationRef {
			return false
		}
		for i := range m.Type {
			if !m.Type[i].Equals(o.Type[i]) {
				return false
			}
		}
		return true
	}
	return false
}

func (m *ParameterizedStructType) ShortString() string {
	t := StructType{}
	return t.ShortString()
}

func (m *ParameterizedStructType) String() string {
	sb := strings.Builder{}
	for i, typ := range m.Type {
		if i != 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(typ.String())
	}
	t := StructType{}
	parameterString := fmt.Sprintf("<%s>", sb.String())
	return fmt.Sprintf("%s%s%s", t.BaseString(), strNullable(m), parameterString)
}

// GetAbstractParameters returns the abstract parameter names
// this implements interface ParameterizedAbstractType
func (m *ParameterizedStructType) GetAbstractParameters() []parameter_types.AbstractParameterType {
	var abstractParams []parameter_types.AbstractParameterType
	for _, typ := range m.Type {
		if abs, ok := typ.(parameter_types.AbstractParameterType); ok {
			abstractParams = append(abstractParams, abs)
		}
	}
	return abstractParams
}

// GetAbstractParamName this implements interface AbstractParameterType
// to indicate ParameterizedStructType itself can be used as a parameter of abstract type too
func (m *ParameterizedStructType) GetAbstractParamName() string {
	return m.String()
}
