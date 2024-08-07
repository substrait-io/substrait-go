package functions

import (
	"github.com/substrait-io/substrait-go/extensions"
)

var DefaultRegistry FunctionRegistry

func GetDefaultFunctionRegistry() FunctionRegistry {
	if DefaultRegistry == nil {
		DefaultRegistry = NewFunctionRegistry(&extensions.DefaultCollection)
	}
	return DefaultRegistry
}

type functionRegistryImpl struct {
	collection         *extensions.Collection
	scalarFunctions    map[string][]*extensions.ScalarFunctionVariant
	aggregateFunctions map[string][]*extensions.AggregateFunctionVariant
	windowFunctions    map[string][]*extensions.WindowFunctionVariant
	allFunctions       []extensions.FunctionVariant
	dialects           map[string]Dialect
}

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

func processFunctions[T extensions.FunctionVariant](functions []T, funcMap map[string][]T, allFuncs *[]extensions.FunctionVariant) {
	for _, f := range functions {
		name := f.Name()
		if _, ok := funcMap[name]; !ok {
			funcMap[name] = make([]T, 0)
		}
		funcMap[name] = append(funcMap[name], f)
		*allFuncs = append(*allFuncs, f)
	}
}

func (f *functionRegistryImpl) GetAllFunctions() []extensions.FunctionVariant {
	return f.allFunctions
}

func (f *functionRegistryImpl) GetScalarFunctions(name string) []*extensions.ScalarFunctionVariant {
	return f.scalarFunctions[name]
}

func (f *functionRegistryImpl) GetAggregateFunctions(name string) []*extensions.AggregateFunctionVariant {
	return f.aggregateFunctions[name]
}

func (f *functionRegistryImpl) GetWindowFunctions(name string) []*extensions.WindowFunctionVariant {
	return f.windowFunctions[name]
}
