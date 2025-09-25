package functions

import (
	"fmt"

	"github.com/substrait-io/substrait-go/v7/expr"
	"github.com/substrait-io/substrait-go/v7/extensions"
)

type localFunctionRegistryImpl struct {
	dialect Dialect

	// substrait function name to local function variants
	scalarFunctions    map[FunctionName][]*LocalScalarFunctionVariant
	aggregateFunctions map[FunctionName][]*LocalAggregateFunctionVariant
	windowFunctions    map[FunctionName][]*LocalWindowFunctionVariant

	allFunctions []extensions.FunctionVariant

	idToLocalFunctionMap map[extensions.ID]localFunctionVariant
	localTypeRegistry    LocalTypeRegistry
	funcRegistry         FunctionRegistry
}

func makeLocalFunctionVariantsMap(functions []extensions.FunctionVariant) map[extensions.ID]localFunctionVariant {
	localFunctionVariants := make(map[extensions.ID]localFunctionVariant)
	for _, f := range functions {
		switch variant := f.(type) {
		case *LocalScalarFunctionVariant:
			localFunctionVariants[variant.ID()] = variant
		case *LocalAggregateFunctionVariant:
			localFunctionVariants[variant.ID()] = variant
		case *LocalWindowFunctionVariant:
			localFunctionVariants[variant.ID()] = variant
		}
	}
	return localFunctionVariants
}

func (l *localFunctionRegistryImpl) GetAllFunctions() []extensions.FunctionVariant {
	return l.allFunctions
}

func (l *localFunctionRegistryImpl) GetDialect() Dialect {
	return l.dialect
}

func (l *localFunctionRegistryImpl) GetFunctionRegistry() FunctionRegistry {
	return l.funcRegistry
}

func (l *localFunctionRegistryImpl) GetScalarFunctions(name FunctionName, numArgs int) []*LocalScalarFunctionVariant {
	return getFunctionVariantsByCount(getOrEmpty(name, l.scalarFunctions), numArgs)
}

func (l *localFunctionRegistryImpl) GetAggregateFunctions(name FunctionName, numArgs int) []*LocalAggregateFunctionVariant {
	return getFunctionVariantsByCount(getOrEmpty(name, l.aggregateFunctions), numArgs)
}

func (l *localFunctionRegistryImpl) GetWindowFunctions(name FunctionName, numArgs int) []*LocalWindowFunctionVariant {
	return getFunctionVariantsByCount(getOrEmpty(name, l.windowFunctions), numArgs)
}

func (l *localFunctionRegistryImpl) GetScalarFunctionByInvocation(scalarFuncInvocation *expr.ScalarFunction) (*LocalScalarFunctionVariant, error) {
	return getFunctionVariantByInvocation[*LocalScalarFunctionVariant](scalarFuncInvocation, l)
}

func (l *localFunctionRegistryImpl) GetAggregateFunctionByInvocation(aggregateFuncInvocation *expr.AggregateFunction) (*LocalAggregateFunctionVariant, error) {
	return getFunctionVariantByInvocation[*LocalAggregateFunctionVariant](aggregateFuncInvocation, l)
}

func (l *localFunctionRegistryImpl) GetWindowFunctionByInvocation(windowFuncInvocation *expr.WindowFunction) (*LocalWindowFunctionVariant, error) {
	return getFunctionVariantByInvocation[*LocalWindowFunctionVariant](windowFuncInvocation, l)
}

func getFunctionVariantByInvocation[V localFunctionVariant](invocation expr.FunctionInvocation, registry *localFunctionRegistryImpl) (V, error) {
	var zeroV V
	f, ok := registry.idToLocalFunctionMap[invocation.ID()]
	if !ok {
		return zeroV, fmt.Errorf("function variant not found for function: %s", invocation.ID())
	}
	argTypes := invocation.GetArgTypes()
	for i, argType := range argTypes {
		_, err := registry.localTypeRegistry.GetLocalTypeFromSubstraitType(argType)
		if err != nil {
			return zeroV, fmt.Errorf("unsupported substrait type: %v as argument %d in %s", argType, i, invocation.CompoundName())
		}
	}
	for _, option := range invocation.GetOptions() {
		for _, value := range option.Preference {
			if !f.IsOptionSupported(option.Name, value) {
				return zeroV, fmt.Errorf("unsupported option [%s:%s] in function %s", option.Name, value, invocation.CompoundName())
			}
		}
	}
	return f.(V), nil
}

var _ LocalFunctionRegistry = &localFunctionRegistryImpl{}
