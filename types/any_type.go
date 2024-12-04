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

func (m *AnyType) SetNullability(n Nullability) FuncDefArgType {
	m.Nullability = n
	return m
}

func (m *AnyType) String() string {
	return fmt.Sprintf("%s%s", m.Name, strFromNullability(m.Nullability))
}

func (m *AnyType) HasParameterizedParam() bool {
	// primitive type doesn't have abstract parameters
	return false
}

func (m *AnyType) GetParameterizedParams() []interface{} {
	// any type doesn't have any abstract parameters
	return nil
}

func (m *AnyType) MatchWithNullability(ot Type) bool {
	return m.Nullability == ot.GetNullability()
}

func (m *AnyType) MatchWithoutNullability(ot Type) bool {
	return true
}

func (m *AnyType) GetNullability() Nullability {
	return m.Nullability
}

func (m *AnyType) ShortString() string {
	return "any"
}

func (m *AnyType) ReturnType(funcParameters []FuncDefArgType, argumentTypes []Type) (Type, error) {
	index := -1
	for i, param := range funcParameters {
		if anyArg, ok := param.(*AnyType); ok {
			if anyArg.Name == m.Name {
				index = i
				break
			}
		}
	}
	if index == -1 {
		return nil, fmt.Errorf("no matching any type found in function parameters")
	}
	if index >= len(argumentTypes) {
		return nil, fmt.Errorf("no matching argument found for any type")
	}
	return argumentTypes[index], nil
}

func (m *AnyType) WithParameters([]interface{}) (Type, error) {
	return nil, fmt.Errorf("any type doesn't have any parameters")
}
