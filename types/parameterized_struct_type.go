// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"
	"strings"
)

// ParameterizedStructType is a parameter type struct
// example: Struct<Decimal(P,S), INT8> or Struct<INT8>.
type ParameterizedStructType struct {
	Nullability      Nullability
	TypeVariationRef uint32
	Types            []FuncDefArgType
}

func (m *ParameterizedStructType) SetNullability(n Nullability) FuncDefArgType {
	m.Nullability = n
	return m
}

func (m *ParameterizedStructType) String() string {
	sb := strings.Builder{}
	for i, typ := range m.Types {
		if i != 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(typ.String())
	}
	t := StructType{}
	parameterString := fmt.Sprintf("<%s>", sb.String())
	return fmt.Sprintf("%s%s%s", t.BaseString(), strFromNullability(m.Nullability), parameterString)
}

func (m *ParameterizedStructType) HasParameterizedParam() bool {
	for _, typ := range m.Types {
		if typ.HasParameterizedParam() {
			return true
		}
	}
	return false
}

func (m *ParameterizedStructType) GetParameterizedParams() []interface{} {
	if !m.HasParameterizedParam() {
		return nil
	}
	var abstractParams []interface{}
	for _, typ := range m.Types {
		if typ.HasParameterizedParam() {
			abstractParams = append(abstractParams, typ)
		}
	}
	return abstractParams
}
