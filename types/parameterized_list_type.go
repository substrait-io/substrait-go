// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"
)

// ParameterizedListType is a list type having parameter of ParameterizedAbstractType
// basically a list of which type is another abstract parameter
// example: List<Decimal(P,S)>. Kindly note concrete types List<Decimal(38, 0)> is not represented by this type
// Concrete type is represented by ListType
type ParameterizedListType struct {
	Nullability      Nullability
	TypeVariationRef uint32
	Type             FuncDefArgType
}

func (m *ParameterizedListType) SetNullability(n Nullability) FuncDefArgType {
	m.Nullability = n
	return m
}

func (m *ParameterizedListType) String() string {
	t := ListType{}
	parameterString := fmt.Sprintf("<%s>", m.Type)
	return fmt.Sprintf("%s%s%s", t.BaseString(), strFromNullability(m.Nullability), parameterString)
}

func (m *ParameterizedListType) HasParameterizedParam() bool {
	return m.Type.HasParameterizedParam()
}

func (m *ParameterizedListType) GetParameterizedParams() []interface{} {
	if !m.HasParameterizedParam() {
		return nil
	}
	return []interface{}{m.Type}
}
