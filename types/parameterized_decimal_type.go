// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"

	"github.com/substrait-io/substrait-go/types/integer_parameters"
)

// ParameterizedDecimalType is a decimal type which to hold function arguments
// example: Decimal<P,S> or Decimal<P,0> or Decimal(10, 2)
type ParameterizedDecimalType struct {
	Nullability      Nullability
	TypeVariationRef uint32
	Precision        integer_parameters.IntegerParameter
	Scale            integer_parameters.IntegerParameter
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
	_, ok1 := m.Precision.(*integer_parameters.VariableIntParam)
	_, ok2 := m.Scale.(*integer_parameters.VariableIntParam)
	return ok1 || ok2
}

func (m *ParameterizedDecimalType) GetParameterizedParams() []interface{} {
	if !m.HasParameterizedParam() {
		return nil
	}
	var params []interface{}
	if p, ok := m.Precision.(*integer_parameters.VariableIntParam); ok {
		params = append(params, p)
	}
	if p, ok := m.Scale.(*integer_parameters.VariableIntParam); ok {
		params = append(params, p)
	}
	return params
}

func (m *ParameterizedDecimalType) MatchWithNullability(ot Type) bool {
	if m.Nullability != ot.GetNullability() {
		return false
	}
	return m.MatchWithoutNullability(ot)
}

func (m *ParameterizedDecimalType) MatchWithoutNullability(ot Type) bool {
	if odt, ok := ot.(*DecimalType); ok {
		concretePrecision := integer_parameters.NewConcreteIntParam(odt.Precision)
		concreteScale := integer_parameters.NewConcreteIntParam(odt.Scale)
		return m.Precision.IsCompatible(concretePrecision) && m.Scale.IsCompatible(concreteScale)
	}
	return false
}
