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

func (m *ParameterizedMapType) MatchWithNullability(ot Type) bool {
	if m.Nullability != ot.GetNullability() {
		return false
	}
	if omt, ok := ot.(*MapType); ok {
		return m.Key.MatchWithNullability(omt.Key) && m.Value.MatchWithNullability(omt.Value)
	}
	return false
}

func (m *ParameterizedMapType) MatchWithoutNullability(ot Type) bool {
	if omt, ok := ot.(*MapType); ok {
		return m.Key.MatchWithoutNullability(omt.Key) && m.Value.MatchWithoutNullability(omt.Value)
	}
	return false
}

func (m *ParameterizedMapType) GetNullability() Nullability {
	return m.Nullability
}

func (m *ParameterizedMapType) ShortString() string {
	return "map"
}

func (m *ParameterizedMapType) ReturnType() (Type, error) {
	keyType, kerr := m.Key.ReturnType()
	if kerr != nil {
		return nil, fmt.Errorf("error in getting key type: %w", kerr)
	}
	valueType, verr := m.Value.ReturnType()
	if verr != nil {
		return nil, fmt.Errorf("error in getting value type: %w", kerr)
	}

	return &MapType{Nullability: m.Nullability, Key: keyType, Value: valueType}, nil
}
