// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"

	"github.com/substrait-io/substrait-go/types/leaf_parameters"
)

// ParameterizedDecimalType is a decimal type which to hold function arguments
// example: Decimal<P,S> or Decimal<P,0> or Decimal(10, 2)
type ParameterizedDecimalType struct {
	Nullability      Nullability
	TypeVariationRef uint32
	Precision        leaf_parameters.LeafParameter
	Scale            leaf_parameters.LeafParameter
}

func (m *ParameterizedDecimalType) SetNullability(n Nullability) FuncDefArgType {
	m.Nullability = n
	return m
}

func (m *ParameterizedDecimalType) String() string {
	t := DecimalType{}
	parameterString := fmt.Sprintf("<%s,%s>", m.Precision.String(), m.Scale.String())
	return fmt.Sprintf("%s%s%s", t.BaseString(), strFromNullability(m.Nullability), parameterString)
}

func (m *ParameterizedDecimalType) HasParameterizedParam() bool {
	_, ok1 := m.Precision.(*leaf_parameters.VariableIntParam)
	_, ok2 := m.Scale.(*leaf_parameters.VariableIntParam)
	return ok1 || ok2
}

func (m *ParameterizedDecimalType) GetParameterizedParams() []interface{} {
	if !m.HasParameterizedParam() {
		return nil
	}
	var params []interface{}
	if p, ok := m.Precision.(*leaf_parameters.VariableIntParam); ok {
		params = append(params, p)
	}
	if p, ok := m.Scale.(*leaf_parameters.VariableIntParam); ok {
		params = append(params, p)
	}
	return params
}
