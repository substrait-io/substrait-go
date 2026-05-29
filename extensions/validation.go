// SPDX-License-Identifier: Apache-2.0

package extensions

import (
	"fmt"

	substraitgo "github.com/substrait-io/substrait-go/v8"
	"github.com/substrait-io/substrait-go/v8/types"
	"github.com/substrait-io/substrait-go/v8/types/parser"
)

func (s SimpleExtensionFile) validateLocalUserDefinedTypeReferences() error {
	declaredTypes := make(map[string]struct{}, len(s.Types))
	for _, typ := range s.Types {
		declaredTypes[typ.Name] = struct{}{}
	}

	for _, fn := range s.ScalarFunctions {
		for i, impl := range fn.Impls {
			if err := impl.validateLocalUserDefinedTypeReferences(declaredTypes); err != nil {
				return fmt.Errorf("scalar function %q impl %d: %w", fn.Name, i, err)
			}
		}
	}

	for _, fn := range s.AggregateFunctions {
		for i, impl := range fn.Impls {
			if err := impl.validateLocalUserDefinedTypeReferences(declaredTypes); err != nil {
				return fmt.Errorf("aggregate function %q impl %d: %w", fn.Name, i, err)
			}
		}
	}

	for _, fn := range s.WindowFunctions {
		for i, impl := range fn.Impls {
			if err := impl.validateLocalUserDefinedTypeReferences(declaredTypes); err != nil {
				return fmt.Errorf("window function %q impl %d: %w", fn.Name, i, err)
			}
		}
	}

	return nil
}

func (s ScalarFunctionImpl) validateLocalUserDefinedTypeReferences(declaredTypes map[string]struct{}) error {
	for i, arg := range s.Args {
		if err := validateFunctionParameterLocalUserDefinedTypeReferences(declaredTypes, arg); err != nil {
			return fmt.Errorf("arg %d: %w", i, err)
		}
	}

	if s.Return != nil {
		if err := validateTypeExpressionLocalUserDefinedTypeReferences(declaredTypes, *s.Return); err != nil {
			return fmt.Errorf("return: %w", err)
		}
	}

	return nil
}

func (s AggregateFunctionImpl) validateLocalUserDefinedTypeReferences(declaredTypes map[string]struct{}) error {
	if err := s.ScalarFunctionImpl.validateLocalUserDefinedTypeReferences(declaredTypes); err != nil {
		return err
	}
	if err := validateTypeExpressionLocalUserDefinedTypeReferences(declaredTypes, s.Intermediate); err != nil {
		return fmt.Errorf("intermediate: %w", err)
	}
	return nil
}

func validateFunctionParameterLocalUserDefinedTypeReferences(declaredTypes map[string]struct{}, arg FuncParameter) error {
	switch arg := arg.(type) {
	case ValueArg:
		if arg.Value == nil {
			return nil
		}
		return validateTypeExpressionLocalUserDefinedTypeReferences(declaredTypes, *arg.Value)
	case TypeArg:
		if arg.Type == nil {
			return nil
		}
		return validateTypeExpressionLocalUserDefinedTypeReferences(declaredTypes, *arg.Type)
	default:
		return nil
	}
}

func validateTypeExpressionLocalUserDefinedTypeReferences(declaredTypes map[string]struct{}, expr parser.TypeExpression) error {
	if expr.ValueType == nil {
		return nil
	}
	return validateFuncDefArgTypeLocalUserDefinedTypeReferences(declaredTypes, expr.ValueType)
}

func validateFuncDefArgTypeLocalUserDefinedTypeReferences(declaredTypes map[string]struct{}, typ types.FuncDefArgType) error {
	switch typ := typ.(type) {
	case *types.ParameterizedUserDefinedType:
		if _, ok := declaredTypes[typ.Name]; !ok {
			return fmt.Errorf("%w: user-defined type %q is not declared", substraitgo.ErrInvalidSimpleExtention, typ.Name)
		}
		for _, param := range typ.TypeParameters {
			if err := validateUDTParameterLocalUserDefinedTypeReferences(declaredTypes, param); err != nil {
				return err
			}
		}
	case *types.ParameterizedListType:
		return validateFuncDefArgTypeLocalUserDefinedTypeReferences(declaredTypes, typ.Type)
	case *types.ParameterizedMapType:
		if err := validateFuncDefArgTypeLocalUserDefinedTypeReferences(declaredTypes, typ.Key); err != nil {
			return err
		}
		return validateFuncDefArgTypeLocalUserDefinedTypeReferences(declaredTypes, typ.Value)
	case *types.ParameterizedStructType:
		for _, fieldType := range typ.Types {
			if err := validateFuncDefArgTypeLocalUserDefinedTypeReferences(declaredTypes, fieldType); err != nil {
				return err
			}
		}
	case *types.ParameterizedFuncType:
		for _, paramType := range typ.Parameters {
			if err := validateFuncDefArgTypeLocalUserDefinedTypeReferences(declaredTypes, paramType); err != nil {
				return err
			}
		}
		return validateFuncDefArgTypeLocalUserDefinedTypeReferences(declaredTypes, typ.Return)
	}
	return nil
}

func validateUDTParameterLocalUserDefinedTypeReferences(declaredTypes map[string]struct{}, param types.UDTParameter) error {
	if dataParam, ok := param.(*types.DataTypeUDTParam); ok {
		return validateFuncDefArgTypeLocalUserDefinedTypeReferences(declaredTypes, dataParam.Type)
	}
	return nil
}
