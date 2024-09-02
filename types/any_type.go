package types

import (
	"fmt"

	"github.com/substrait-io/substrait-go/proto"
)

// AnyType to represent AnyType, this type is to indicate "any" type of argument
// This type is not used in function invocation. It is only used in function definition
type AnyType struct {
	Name        string
	Nullability Nullability
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
func (AnyType) GetTypeVariationReference() uint32 {
	panic("not allowed")
}
func (AnyType) Equals(rhs Type) bool {
	// equal to every other type
	return true
}

func (AnyType) ToProtoFuncArg() *proto.FunctionArgument {
	panic("not allowed")
}

func (AnyType) ToProto() *proto.Type {
	panic("not allowed")
}

func (t AnyType) ShortString() string { return t.Name }
func (t AnyType) String() string {
	return fmt.Sprintf("%s%s", t.Name, strNullable(t))
}

// Below methods are for parser Def interface

func (AnyType) Optional() bool {
	panic("not allowed")
}

func (m AnyType) ShortType() string {
	return "any"
}

func (m AnyType) Type() (Type, error) {
	return m, nil
}
