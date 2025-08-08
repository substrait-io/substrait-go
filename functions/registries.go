package functions

import (
	"strings"

	"github.com/substrait-io/substrait-go/v5/expr"
	"github.com/substrait-io/substrait-go/v5/extensions"
	"github.com/substrait-io/substrait-go/v5/types"
)

type TypeRegistry interface {
	// GetTypeFromTypeString gets the Substrait type for a given Substrait standard typeString.
	// Returns an error if type doesn't parse. This type registry should support both standard types
	// and any available extension types.
	GetTypeFromTypeString(typeString string) (types.Type, error)
}

// LocalTypeRegistry is a registry that contains all types associated with a particular dialect.
type LocalTypeRegistry interface {
	// GetTypeFromTypeString gets the Substrait type for a given Substrait standard typeString.
	GetTypeFromTypeString(typeString string) (types.Type, error)

	// GetSubstraitTypeFromLocalType gets the Substrait type for a given local type string.
	GetSubstraitTypeFromLocalType(localType string) (types.Type, error)

	// GetLocalTypeFromSubstraitType gets the local type string for a given Substrait type.
	GetLocalTypeFromSubstraitType(typ types.Type) (string, error)

	//GetTypeClasses() []types.TypeClass // TODO

	// GetSupportedTypes returns the types supported by this dialect.
	GetSupportedTypes() map[string]types.Type

	// IsTypeSupportedInTables checks whether a particular type is supported in tables.
	// Some types (such as INTERVAL) may only be supported in literal contexts.
	IsTypeSupportedInTables(typ types.Type) bool
}

// Dialect is the entry point to understanding the mapping between Substrait and a specific target system
type Dialect interface {
	Name() string

	// LocalizeFunctionRegistry creates a function registry restricts the provided registry down to
	// the subset of functions supported by this dialect. This will return an error if there are
	// functions declared in the dialect that are not available within the provided registry.
	LocalizeFunctionRegistry(registry FunctionRegistry) (LocalFunctionRegistry, error)

	// LocalizeTypeRegistry creates a type registry restricts the provided registry down to
	// the subset of types supported by this dialect. This will return an error if there are
	// types declared in the dialect that are not available within the provided registry.
	LocalizeTypeRegistry(registry TypeRegistry) (LocalTypeRegistry, error)

	// GetLocalTypeRegistry returns the last created type registry using this dialect or constructs
	// one using LocalizeTypeRegistry and a default type registry if one hasn't yet been made.
	GetLocalTypeRegistry() (LocalTypeRegistry, error)
}

type FunctionName interface {
	functionName()
}

// LocalFunctionName is a function name localized to a specific dialect
type LocalFunctionName string

// SubstraitFunctionName is the short name of the function (excluded URI and argument types) in Substrait
type SubstraitFunctionName string

func (LocalFunctionName) functionName()     {}
func (SubstraitFunctionName) functionName() {}

type functionRegistryBase[L any, S extensions.FunctionVariant, A extensions.FunctionVariant, W extensions.FunctionVariant] interface {
	// GetScalarFunctions returns a slice of zero or more scalar function variants that match the provided name & numArgs.
	// Use SubstraitFunctionName as the name's type if you want to search by substrait name and LocalFunctionName as the name's
	// type if you want to search by the local name.
	GetScalarFunctions(name L, numArgs int) []S

	// GetAggregateFunctions returns a slice of zero or more aggregate function variants that match the provided name & numArgs.
	// Use SubstraitFunctionName as the name's type if you want to search by substrait name and LocalFunctionName as the name's
	// type if you want to search by the local name.
	GetAggregateFunctions(name L, numArgs int) []A

	// GetWindowFunctions returns a slice of zero or more window function variants that match the provided name & numArgs.
	// Use SubstraitFunctionName as the name's type if you want to search by substrait name and LocalFunctionName as the name's
	// type if you want to search by the local name.
	GetWindowFunctions(name L, numArgs int) []W

	// GetAllFunctions returns all function variants in the registry
	GetAllFunctions() []extensions.FunctionVariant
}

type FunctionRegistry interface {
	functionRegistryBase[string, *extensions.ScalarFunctionVariant, *extensions.AggregateFunctionVariant, *extensions.WindowFunctionVariant]
	GetScalarFunctionsByName(name string) []*extensions.ScalarFunctionVariant
	GetAggregateFunctionsByName(name string) []*extensions.AggregateFunctionVariant
	GetWindowFunctionsByName(name string) []*extensions.WindowFunctionVariant
}

// LocalFunctionRegistry is a collection of functions localized to a particular Dialect
type LocalFunctionRegistry interface {
	functionRegistryBase[FunctionName, *LocalScalarFunctionVariant, *LocalAggregateFunctionVariant, *LocalWindowFunctionVariant]
	GetDialect() Dialect
	GetFunctionRegistry() FunctionRegistry
	GetScalarFunctionByInvocation(scalarFuncInvocation *expr.ScalarFunction) (*LocalScalarFunctionVariant, error)
	GetAggregateFunctionByInvocation(aggregateFuncInvocation *expr.AggregateFunction) (*LocalAggregateFunctionVariant, error)
	GetWindowFunctionByInvocation(windowFuncInvocation *expr.WindowFunction) (*LocalWindowFunctionVariant, error)
}

type FunctionNotation int

const (
	INFIX FunctionNotation = iota
	PREFIX
	POSTFIX
)

type localFunctionVariant interface {
	extensions.FunctionVariant
	LocalName() string
	Notation() FunctionNotation
	IsOptionSupported(name string, value string) bool
	InvocationTypeName() string
}

type LocalFunctionVariant struct {
	localName        string
	supportedOptions map[string]extensions.Option
	notation         FunctionNotation
}

func (l *LocalFunctionVariant) LocalName() string {
	return l.localName
}

func (l *LocalFunctionVariant) Notation() FunctionNotation {
	return l.notation
}

func (l *LocalFunctionVariant) IsOptionSupported(name string, value string) bool {
	val, exists := l.supportedOptions[name]
	if !exists {
		// TODO: should this be true or false?
		return false
	}
	for _, v := range val.Values {
		if strings.EqualFold(v, value) {
			return true
		}
	}
	return false
}

// LocalScalarFunctionVariant is a ScalarFunctionVariant that also understands its context in a particular dialect
type LocalScalarFunctionVariant struct {
	extensions.ScalarFunctionVariant
	LocalFunctionVariant
}

var _ extensions.FunctionVariant = &LocalScalarFunctionVariant{}

func (l *LocalScalarFunctionVariant) InvocationTypeName() string {
	return "scalar"
}

type LocalAggregateFunctionVariant struct {
	extensions.AggregateFunctionVariant
	LocalFunctionVariant
}

var _ extensions.FunctionVariant = &LocalAggregateFunctionVariant{}

func (l *LocalAggregateFunctionVariant) InvocationTypeName() string {
	return "aggregate"
}

type LocalWindowFunctionVariant struct {
	extensions.WindowFunctionVariant
	LocalFunctionVariant
}

var _ extensions.FunctionVariant = &LocalWindowFunctionVariant{}

func (l *LocalWindowFunctionVariant) InvocationTypeName() string {
	return "window"
}

func newLocalScalarFunctionVariant(sf *extensions.ScalarFunctionVariant, dfi *dialectFunctionInfo) *LocalScalarFunctionVariant {
	return &LocalScalarFunctionVariant{
		ScalarFunctionVariant: *sf,
		LocalFunctionVariant: LocalFunctionVariant{
			localName:        dfi.LocalName,
			supportedOptions: dfi.Options,
			notation:         dfi.Notation,
		},
	}
}

func newLocalAggregateFunctionVariant(af *extensions.AggregateFunctionVariant, dfi *dialectFunctionInfo) *LocalAggregateFunctionVariant {
	return &LocalAggregateFunctionVariant{
		AggregateFunctionVariant: *af,
		LocalFunctionVariant: LocalFunctionVariant{
			localName:        dfi.LocalName,
			supportedOptions: dfi.Options,
			notation:         dfi.Notation,
		},
	}
}

func newLocalWindowFunctionVariant(wf *extensions.WindowFunctionVariant, dfi *dialectFunctionInfo) *LocalWindowFunctionVariant {
	return &LocalWindowFunctionVariant{
		WindowFunctionVariant: *wf,
		LocalFunctionVariant: LocalFunctionVariant{
			localName:        dfi.LocalName,
			supportedOptions: dfi.Options,
			notation:         dfi.Notation,
		},
	}
}

func NewExtensionAndFunctionRegistries(c *extensions.Collection) (expr.ExtensionRegistry, FunctionRegistry) {
	return expr.NewEmptyExtensionRegistry(c), NewFunctionRegistry(c)
}
