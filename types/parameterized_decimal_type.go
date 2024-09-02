package types

import (
	"fmt"

	"github.com/substrait-io/substrait-go/proto"
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

func (ParameterizedDecimalType) ToProtoFuncArg() *proto.FunctionArgument {
	// parameterized type are never on wire so to proto is not supported
	panic("not supported")
}

func (m ParameterizedDecimalType) ShortString() string {
	t := DecimalType{}
	return t.ShortString()
}

func (m ParameterizedDecimalType) String() string {
	return fmt.Sprintf("%s%s%s", m.BaseString(), strNullable(m), m.ParameterString())
}

func (m ParameterizedDecimalType) ParameterString() string {
	return fmt.Sprintf("<%s,%s>", m.Precision.String(), m.Scale.String())
}

func (m ParameterizedDecimalType) BaseString() string {
	t := DecimalType{}
	return t.BaseString()
}
