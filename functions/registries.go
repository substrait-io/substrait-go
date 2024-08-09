package functions

import (
	"github.com/substrait-io/substrait-go/extensions"
	"github.com/substrait-io/substrait-go/types"
	"strings"
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
}

type FunctionRegistry interface {
	// GetScalarFunctions returns a slice of zero or more scalar function variants that match the provided name & numArgs
	GetScalarFunctions(name string, numArgs int) []*extensions.ScalarFunctionVariant
	GetScalarFunctionsByName(name string) []*extensions.ScalarFunctionVariant

	// GetAggregateFunctions returns a slice of zero or more aggregate function variants that match the provided name & numArgs
	GetAggregateFunctions(name string, numArgs int) []*extensions.AggregateFunctionVariant
	GetAggregateFunctionsByName(name string) []*extensions.AggregateFunctionVariant

	// GetWindowFunctions returns a slice of zero or more window function variants that match the provided name & numArgs
	GetWindowFunctions(name string, numArgs int) []*extensions.WindowFunctionVariant
	GetWindowFunctionsByName(name string) []*extensions.WindowFunctionVariant

	// GetAllFunctions returns all function variants in the registry
	GetAllFunctions() []extensions.FunctionVariant
}

// NameKind is an enum that describes the kind of name being used to look up a function
type NameKind int

const (
	Substrait NameKind = iota
	Local
)

// LocalFunctionRegistry is a collection of functions localized to a particular Dialect
type LocalFunctionRegistry interface {
	// GetScalarFunctionsBy returns a slice of zero or more scalar function variants that match the given name, numArgs & kind.
	GetScalarFunctionsBy(name string, numArgs int, kind NameKind) []*LocalScalarFunctionVariant

	// GetAggregateFunctionsBy returns a slice of zero or more aggregate function variants that match the given name, numArgs & kind.
	GetAggregateFunctionsBy(name string, numArgs int, kind NameKind) []*LocalAggregateFunctionVariant

	// GetWindowFunctionsBy returns a slice of zero or more window function variants that match the given name, numArgs & kind.
	GetWindowFunctionsBy(name string, numArgs int, kind NameKind) []*LocalWindowFunctionVariant

	// GetDialect returns the dialect that this function registry is localized to
	GetDialect() Dialect
}

type FunctionNotation int

const (
	INFIX FunctionNotation = iota
	PREFIX
	POSTFIX
)

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

type LocalAggregateFunctionVariant struct {
	extensions.AggregateFunctionVariant
	LocalFunctionVariant
}

type LocalWindowFunctionVariant struct {
	extensions.WindowFunctionVariant
	LocalFunctionVariant
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
