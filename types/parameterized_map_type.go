// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"
)

// ParameterizedMapType is a struct having at least one of key or value of type ParameterizedAbstractType
// If All arguments are concrete they are represented by MapType
type ParameterizedMapType struct {
	Nullability      Nullability
	TypeVariationRef uint32
	Key              FuncDefArgType
	Value            FuncDefArgType
}

func (m *ParameterizedMapType) SetNullability(n Nullability) FuncDefArgType {
	m.Nullability = n
	return m
}

func (m *ParameterizedMapType) String() string {
	t := MapType{}
	parameterString := fmt.Sprintf("<%s, %s>", m.Key.String(), m.Value.String())
	return fmt.Sprintf("%s%s%s", t.BaseString(), strFromNullability(m.Nullability), parameterString)
}

func (m *ParameterizedMapType) HasParameterizedParam() bool {
	return m.Key.HasParameterizedParam() || m.Value.HasParameterizedParam()
}

func (m *ParameterizedMapType) GetParameterizedParams() []interface{} {
	if !m.HasParameterizedParam() {
		return nil
	}
	var abstractParams []interface{}
	if m.Key.HasParameterizedParam() {
		abstractParams = append(abstractParams, m.Key)
	}
	if m.Value.HasParameterizedParam() {
		abstractParams = append(abstractParams, m.Value)
	}
	return abstractParams
}
