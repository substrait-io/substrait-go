// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"

	"github.com/substrait-io/substrait-go/types/leaf_parameters"
)

// parameterizedTypeSingleIntegerParam This is a generic type to represent parameterized type with a single integer parameter
type parameterizedTypeSingleIntegerParam[T VarCharType | FixedCharType | FixedBinaryType | PrecisionTimestampType | PrecisionTimestampTzType] struct {
	Nullability      Nullability
	TypeVariationRef uint32
	IntegerOption    leaf_parameters.LeafParameter
}

func (m *parameterizedTypeSingleIntegerParam[T]) SetNullability(n Nullability) FuncDefArgType {
	m.Nullability = n
	return m
}

func (m *parameterizedTypeSingleIntegerParam[T]) String() string {
	return fmt.Sprintf("%s%s%s", m.baseString(), strFromNullability(m.Nullability), m.parameterString())
}

func (m *parameterizedTypeSingleIntegerParam[T]) parameterString() string {
	return fmt.Sprintf("<%s>", m.IntegerOption.String())
}

func (m *parameterizedTypeSingleIntegerParam[T]) baseString() string {
	switch any(m).(type) {
	case *ParameterizedVarCharType:
		t := VarCharType{}
		return t.BaseString()
	case *ParameterizedFixedCharType:
		t := FixedCharType{}
		return t.BaseString()
	case *ParameterizedFixedBinaryType:
		t := FixedBinaryType{}
		return t.BaseString()
	case *ParameterizedPrecisionTimestampType:
		t := PrecisionTimestampType{}
		return t.BaseString()
	case *ParameterizedPrecisionTimestampTzType:
		t := PrecisionTimestampTzType{}
		return t.BaseString()
	default:
		panic("unknown type")
	}
}

func (m *parameterizedTypeSingleIntegerParam[T]) HasParameterizedParam() bool {
	_, ok1 := m.IntegerOption.(*leaf_parameters.VariableIntParam)
	return ok1
}

func (m *parameterizedTypeSingleIntegerParam[T]) GetParameterizedParams() []interface{} {
	if !m.HasParameterizedParam() {
		return nil
	}
	return []interface{}{m.IntegerOption}
}
