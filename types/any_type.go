package types

import (
	"fmt"
)

// AnyType to represent AnyType, this type is to indicate "any" type of argument
// This type is not used in function invocation. It is only used in function definition
type AnyType struct {
	Name             string
	TypeVariationRef uint32
	Nullability      Nullability
}

func (AnyType) isRootRef() {}
func (m AnyType) WithNullability(nullability Nullability) Type {
	m.Nullability = nullability
	return m
}
func (m AnyType) GetType() Type { return m }
func (m AnyType) GetNullability() Nullability {
	return m.Nullability
}
func (m AnyType) GetTypeVariationReference() uint32 {
	return m.TypeVariationRef
}
func (AnyType) Equals(rhs Type) bool {
	// equal to every other type
	return true
}

func (t AnyType) ShortString() string { return t.Name }
func (t AnyType) String() string {
	return fmt.Sprintf("%s%s", t.Name, strNullable(t))
}
