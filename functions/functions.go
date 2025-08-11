package functions

import (
	"github.com/substrait-io/substrait-go/v5/extensions"
)

type functionRegistryImpl struct {
	scalarFunctions    map[string][]*extensions.ScalarFunctionVariant
	aggregateFunctions map[string][]*extensions.AggregateFunctionVariant
	windowFunctions    map[string][]*extensions.WindowFunctionVariant
	allFunctions       []extensions.FunctionVariant
}

func getOrEmpty[K comparable, V any](key K, m map[K][]V) []V {
	if value, exists := m[key]; exists {
		return value
	}

	return make([]V, 0)
}

var _ FunctionRegistry = &functionRegistryImpl{}

func NewFunctionRegistry(collection *extensions.Collection) FunctionRegistry {
	scalarFunctions := make(map[string][]*extensions.ScalarFunctionVariant)
	aggregateFunctions := make(map[string][]*extensions.AggregateFunctionVariant)
	windowFunctions := make(map[string][]*extensions.WindowFunctionVariant)
	allFunctions := make([]extensions.FunctionVariant, 0)

	processFunctions(collection.GetAllScalarFunctions(), scalarFunctions, &allFunctions)
	processFunctions(collection.GetAllAggregateFunctions(), aggregateFunctions, &allFunctions)
	processFunctions(collection.GetAllWindowFunctions(), windowFunctions, &allFunctions)

	return &functionRegistryImpl{
		scalarFunctions:    scalarFunctions,
		aggregateFunctions: aggregateFunctions,
		windowFunctions:    windowFunctions,
		allFunctions:       allFunctions,
	}
}

func processFunctions[T extensions.FunctionVariant](functions []T, funcMap map[string][]T, allFunctions *[]extensions.FunctionVariant) {
	for _, f := range functions {
		name := f.Name()
		if _, ok := funcMap[name]; !ok {
			funcMap[name] = make([]T, 0)
		}
		funcMap[name] = append(funcMap[name], f)
		*allFunctions = append(*allFunctions, f)
	}
}

func (f *functionRegistryImpl) GetAllFunctions() []extensions.FunctionVariant {
	return f.allFunctions
}

func (f *functionRegistryImpl) GetScalarFunctionsByName(name string) []*extensions.ScalarFunctionVariant {
	return getOrEmpty(name, f.scalarFunctions)
}

func (f *functionRegistryImpl) GetAggregateFunctionsByName(name string) []*extensions.AggregateFunctionVariant {
	return getOrEmpty(name, f.aggregateFunctions)
}

func (f *functionRegistryImpl) GetWindowFunctionsByName(name string) []*extensions.WindowFunctionVariant {
	return getOrEmpty(name, f.windowFunctions)
}

func (f *functionRegistryImpl) GetScalarFunctions(name string, numArgs int) []*extensions.ScalarFunctionVariant {
	return getFunctionVariantsByCount(f.GetScalarFunctionsByName(name), numArgs)
}

func (f *functionRegistryImpl) GetAggregateFunctions(name string, numArgs int) []*extensions.AggregateFunctionVariant {
	return getFunctionVariantsByCount(f.GetAggregateFunctionsByName(name), numArgs)
}

func (f *functionRegistryImpl) GetWindowFunctions(name string, numArgs int) []*extensions.WindowFunctionVariant {
	return getFunctionVariantsByCount(f.GetWindowFunctionsByName(name), numArgs)
}

func getFunctionVariantsByCount[T extensions.FunctionVariant](functions []T, numArgs int) []T {
	ret := make([]T, 0)
	for _, f := range functions {
		if len(f.Args()) == numArgs || f.Variadic().IsValidArgumentCount(numArgs) {
			ret = append(ret, f)
		}
	}
	return ret
}

var _ FunctionRegistry = &functionRegistryImpl{}
