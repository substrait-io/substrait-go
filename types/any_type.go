// SPDX-License-Identifier: Apache-2.0

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

func (t AnyType) SetNullability(n Nullability) FuncDefArgType {
	t.Nullability = n
	return t
}

func (t AnyType) String() string {
	return fmt.Sprintf("%s%s", t.Name, strFromNullability(t.Nullability))
}

func (s AnyType) HasParameterizedParam() bool {
	// primitive type doesn't have abstract parameters
	return false
}

func (s AnyType) GetParameterizedParams() []interface{} {
	// any type doesn't have any abstract parameters
	return nil
}

func (m AnyType) MatchWithNullability(ot Type) bool {
	return m.Nullability == ot.GetNullability()
}

func (m AnyType) MatchWithoutNullability(ot Type) bool {
	return true
}

func (m AnyType) GetNullability() Nullability {
	return m.Nullability
}

func (m AnyType) ShortString() string {
	return "any"
}

func (m AnyType) ReturnType([]FuncDefArgType, []Type) (Type, error) {
	return nil, nil
}
