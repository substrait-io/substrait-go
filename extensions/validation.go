// SPDX-License-Identifier: Apache-2.0

package extensions

import (
	"fmt"

	substraitgo "github.com/substrait-io/substrait-go/v8"
	"github.com/substrait-io/substrait-go/v8/types"
	"github.com/substrait-io/substrait-go/v8/types/parser"
)

func validateLocalUserDefinedTypeReferences(file SimpleExtensionFile) error {
	declaredTypes := make(map[string]struct{}, len(file.Types))
	for _, typ := range file.Types {
		declaredTypes[typ.Name] = struct{}{}
	}

	for _, fn := range file.ScalarFunctions {
		for i, impl := range fn.Impls {
			if err := validateScalarFunctionImplLocalUserDefinedTypes(declaredTypes, fn.Name, i, impl); err != nil {
				return err
			}
		}
	}

	for _, fn := range file.AggregateFunctions {
		for i, impl := range fn.Impls {
			if err := validateScalarFunctionImplLocalUserDefinedTypes(declaredTypes, fn.Name, i, impl.ScalarFunctionImpl); err != nil {
				return err
			}
			if err := validateTypeExpressionLocalUserDefinedTypes(declaredTypes, impl.Intermediate); err != nil {
				return fmt.Errorf("aggregate function %q impl %d intermediate: %w", fn.Name, i, err)
			}
		}
	}

	for _, fn := range file.WindowFunctions {
		for i, impl := range fn.Impls {
			if err := validateScalarFunctionImplLocalUserDefinedTypes(declaredTypes, fn.Name, i, impl.ScalarFunctionImpl); err != nil {
				return err
			}
			if err := validateTypeExpressionLocalUserDefinedTypes(declaredTypes, impl.Intermediate); err != nil {
				return fmt.Errorf("window function %q impl %d intermediate: %w", fn.Name, i, err)
			}
		}
	}

	return nil
}

func validateScalarFunctionImplLocalUserDefinedTypes(declaredTypes map[string]struct{}, functionName string, implIndex int, impl ScalarFunctionImpl) error {
	for i, arg := range impl.Args {
		if err := validateFunctionParameterLocalUserDefinedTypes(declaredTypes, arg); err != nil {
			return fmt.Errorf("function %q impl %d arg %d: %w", functionName, implIndex, i, err)
		}
	}

	if impl.Return != nil {
		if err := validateTypeExpressionLocalUserDefinedTypes(declaredTypes, *impl.Return); err != nil {
			return fmt.Errorf("function %q impl %d return: %w", functionName, implIndex, err)
		}
	}

	return nil
}

func validateFunctionParameterLocalUserDefinedTypes(declaredTypes map[string]struct{}, arg FuncParameter) error {
	switch arg := arg.(type) {
	case ValueArg:
		return validateTypeExpressionLocalUserDefinedTypes(declaredTypes, *arg.Value)
	case TypeArg:
		return validateTypeExpressionLocalUserDefinedTypes(declaredTypes, *arg.Type)
	default:
		return nil
	}
}

func validateTypeExpressionLocalUserDefinedTypes(declaredTypes map[string]struct{}, expr parser.TypeExpression) error {
	if expr.ValueType == nil {
		return nil
	}
	return validateFuncDefArgTypeLocalUserDefinedTypes(declaredTypes, expr.ValueType)
}

func validateFuncDefArgTypeLocalUserDefinedTypes(declaredTypes map[string]struct{}, typ types.FuncDefArgType) error {
	switch typ := typ.(type) {
	case *types.ParameterizedUserDefinedType:
		if _, ok := declaredTypes[typ.Name]; !ok {
			return fmt.Errorf("%w: user-defined type %q is not declared", substraitgo.ErrInvalidSimpleExtention, typ.Name)
		}
		for _, param := range typ.TypeParameters {
			if err := validateUDTParameterLocalUserDefinedTypes(declaredTypes, param); err != nil {
				return err
			}
		}
	case *types.ParameterizedListType:
		return validateFuncDefArgTypeLocalUserDefinedTypes(declaredTypes, typ.Type)
	case *types.ParameterizedMapType:
		if err := validateFuncDefArgTypeLocalUserDefinedTypes(declaredTypes, typ.Key); err != nil {
			return err
		}
		return validateFuncDefArgTypeLocalUserDefinedTypes(declaredTypes, typ.Value)
	case *types.ParameterizedStructType:
		for _, fieldType := range typ.Types {
			if err := validateFuncDefArgTypeLocalUserDefinedTypes(declaredTypes, fieldType); err != nil {
				return err
			}
		}
	case *types.ParameterizedFuncType:
		for _, paramType := range typ.Parameters {
			if err := validateFuncDefArgTypeLocalUserDefinedTypes(declaredTypes, paramType); err != nil {
				return err
			}
		}
		return validateFuncDefArgTypeLocalUserDefinedTypes(declaredTypes, typ.Return)
	}
	return nil
}

func validateUDTParameterLocalUserDefinedTypes(declaredTypes map[string]struct{}, param types.UDTParameter) error {
	if dataParam, ok := param.(*types.DataTypeUDTParam); ok {
		return validateFuncDefArgTypeLocalUserDefinedTypes(declaredTypes, dataParam.Type)
	}
	return nil
}
