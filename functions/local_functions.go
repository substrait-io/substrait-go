package functions

import "github.com/substrait-io/substrait-go/extensions"

type localFunctionRegistryImpl struct {
	dialect Dialect

	// substrait function name to local function variants
	scalarFunctions    map[FunctionName][]*LocalScalarFunctionVariant
	aggregateFunctions map[FunctionName][]*LocalAggregateFunctionVariant
	windowFunctions    map[FunctionName][]*LocalWindowFunctionVariant

	allFunctions []extensions.FunctionVariant
}

func (l *localFunctionRegistryImpl) GetAllFunctions() []extensions.FunctionVariant {
	return l.allFunctions
}

func (l *localFunctionRegistryImpl) GetDialect() Dialect {
	return l.dialect
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

var _ LocalFunctionRegistry = &localFunctionRegistryImpl{}
