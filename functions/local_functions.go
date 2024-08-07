package functions

type localFunctionRegistryImpl struct {
	dialect Dialect

	// substrait function name to local function variants
	scalarFunctions    map[string][]*LocalScalarFunctionVariant
	aggregateFunctions map[string][]*LocalAggregateFunctionVariant
	windowFunctions    map[string][]*LocalWindowFunctionVariant

	// local function name to substrait function variants
	localScalarFunctions    map[string][]*LocalScalarFunctionVariant
	localAggregateFunctions map[string][]*LocalAggregateFunctionVariant
	localWindowFunctions    map[string][]*LocalWindowFunctionVariant
}

func newLocalFunctionRegistry(dialect Dialect, scalarFunctions map[string][]*LocalScalarFunctionVariant, aggregateFunctions map[string][]*LocalAggregateFunctionVariant, windowFunctions map[string][]*LocalWindowFunctionVariant) LocalFunctionRegistry {
	return &localFunctionRegistryImpl{
		dialect:                 dialect,
		scalarFunctions:         scalarFunctions,
		aggregateFunctions:      aggregateFunctions,
		windowFunctions:         windowFunctions,
		localScalarFunctions:    buildReverseMap(scalarFunctions),
		localAggregateFunctions: buildReverseMap(aggregateFunctions),
		localWindowFunctions:    buildReverseMap(windowFunctions),
	}
}

type withLocalName interface {
	LocalName() string
}

func buildReverseMap[T withLocalName](localFunctionVariants map[string][]T) map[string][]T {
	reverseMap := make(map[string][]T)
	for _, variants := range localFunctionVariants {
		reverseMap[variants[0].LocalName()] = variants
	}
	return reverseMap
}

func (l *localFunctionRegistryImpl) GetDialect() Dialect {
	return l.dialect
}

func (l *localFunctionRegistryImpl) GetScalarFunctionsBy(name string, kind NameKind) []*LocalScalarFunctionVariant {
	if kind == Substrait {
		return l.scalarFunctions[name]
	}
	return l.localScalarFunctions[name]
}

func (l *localFunctionRegistryImpl) GetAggregateFunctionsBy(name string, kind NameKind) []*LocalAggregateFunctionVariant {
	if kind == Substrait {
		return l.aggregateFunctions[name]
	}
	return l.aggregateFunctions[name]
}

func (l *localFunctionRegistryImpl) GetWindowFunctionsBy(name string, kind NameKind) []*LocalWindowFunctionVariant {
	if kind == Substrait {
		return l.windowFunctions[name]
	}
	return l.windowFunctions[name]
}
