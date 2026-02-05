// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"
	"slices"
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

// unwrapAnyTypeWithName searches for AnyType in p with the specified name,
// and if found, returns argType.  If p is a composite type,
// recursively unwraps p to search for AnyType in p's parameters.
// Returns nil Type if AnyType was not found.
func unwrapAnyTypeWithName(name string, p FuncDefArgType, argType Type) (Type, error) {
	switch arg := p.(type) {
	case *AnyType:
		if arg.Name == name {
			return argType, nil
		}
	case *ParameterizedListType:
		argParams := argType.GetParameters()
		if len(argParams) != 1 || argParams[0] == nil {
			return nil, fmt.Errorf(
				"expected ListType to have non-nil 1 parameter, found %v", argParams)
		}
		return unwrapAnyTypeWithName(name, arg.Type, argParams[0].(Type))
	case *ParameterizedMapType:
		argParams := argType.GetParameters()
		if len(argParams) != 2 || argParams[0] == nil || argParams[1] == nil {
			return nil, fmt.Errorf(
				"expected MapType to have 2 non-nil parameters, found %v", argParams)
		}
		keyType, err := unwrapAnyTypeWithName(name, arg.Key, argParams[0].(Type))
		if err != nil {
			return nil, err
		}
		if keyType != nil {
			return keyType, nil
		}
		return unwrapAnyTypeWithName(name, arg.Value, argParams[1].(Type))
	case *ParameterizedStructType:
		argParams := argType.GetParameters()
		if len(argParams) != len(arg.Types) || slices.Contains(argParams, nil) {
			return nil, fmt.Errorf("expected StructType to have %d non-nil parameters, found %v",
				len(arg.Types), argParams)
		}
		for i, param := range argParams {
			pt, err := unwrapAnyTypeWithName(name, arg.Types[i], param.(Type))
			if err != nil {
				return nil, err
			}
			if pt != nil {
				return pt, nil
			}
		}
	case *ParameterizedFuncType:
		funcType, ok := argType.(*FuncType)
		if !ok {
			return nil, fmt.Errorf("expected FuncType for ParameterizedFuncType, found %T", argType)
		}
		if len(funcType.ParameterTypes) != len(arg.Parameters) {
			return nil, fmt.Errorf("expected FuncType to have %d parameters, found %d",
				len(arg.Parameters), len(funcType.ParameterTypes))
		}
		// Check function parameters
		for i, param := range arg.Parameters {
			pt, err := unwrapAnyTypeWithName(name, param, funcType.ParameterTypes[i])
			if err != nil {
				return nil, err
			}
			if pt != nil {
				return pt, nil
			}
		}
		// Check return type
		return unwrapAnyTypeWithName(name, arg.Return, funcType.ReturnType)
	}
	// Didn't find matching AnyType.
	return nil, nil
}

func (m *AnyType) ReturnType(funcParameters []FuncDefArgType, argumentTypes []Type) (Type, error) {
	// iterate through smaller of the funcParameters and argumentTypes;
	// argumentTypes may be larger than funcParameters due to variadic parameters.
	for i := 0; i < min(len(funcParameters), len(argumentTypes)); i++ {
		typ, err := unwrapAnyTypeWithName(m.Name, funcParameters[i], argumentTypes[i])
		if err != nil {
			return nil, err
		}
		if typ != nil {
			return typ, nil
		}
	}

	return nil, fmt.Errorf("no matching any type found in function parameters")
}

func (m *AnyType) WithParameters([]interface{}) (Type, error) {
	return nil, fmt.Errorf("any type doesn't have any parameters")
}
