// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"
	"reflect"

	"github.com/substrait-io/substrait-go/types/integer_parameters"
)

type singleIntegerParamType interface {
	BaseString() string
}

// parameterizedTypeSingleIntegerParam This is a generic type to represent parameterized type with a single integer parameter
type parameterizedTypeSingleIntegerParam[T singleIntegerParamType] struct {
	Nullability      Nullability
	TypeVariationRef uint32
	IntegerOption    integer_parameters.IntegerParameter
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
	var t T
	tType := reflect.TypeOf(t)
	if tType.Kind() == reflect.Ptr {
		tType = tType.Elem()
	}
	newInstance := reflect.New(tType).Interface().(T)
	return newInstance.BaseString()
}

func (m *parameterizedTypeSingleIntegerParam[T]) HasParameterizedParam() bool {
	_, ok1 := m.IntegerOption.(*integer_parameters.VariableIntParam)
	return ok1
}

func (m *parameterizedTypeSingleIntegerParam[T]) GetParameterizedParams() []interface{} {
	if !m.HasParameterizedParam() {
		return nil
	}
	return []interface{}{m.IntegerOption}
}
