package functions

import (
	"github.com/substrait-io/substrait-go/extensions"
	"github.com/substrait-io/substrait-go/types"
	"strings"
)

type TypeRegistry interface {
	// GetTypeFromTypeString Given a Substrait standard typeString, get the Substrait type. Return
	// error if type doesn't parse. This type registry should support both standard types and any
	// available extension types.
	GetTypeFromTypeString(typeString string) (types.Type, error)

	//GetTypeClasses() []types.TypeClass // TODO define TypeClass so we can get things like DECIMAL<?,?>.
}

// LocalTypeRegistry is a registry that contains all types associated with a particular dialect.
type LocalTypeRegistry interface {
	// GetTypeFromTypeString Given a Substrait standard typeString, get the Substrait type.
	GetTypeFromTypeString(typeString string) (types.Type, error)

	// GetSubstraitTypeFromLocalType Given a local type string, get the Substrait type.
	GetSubstraitTypeFromLocalType(localType string) (types.Type, error)

	// GetLocalTypeFromSubstraitType Given a Substrait type, get the local type string.
	GetLocalTypeFromSubstraitType(typ types.Type) (string, error)

	//GetTypeClasses() []types.TypeClass // TODO

	// IsTypeSupportedInTables Whether a particular type is supported in tables. Some types (such as INTERVAL) may
	// only be supported in literal contexts.
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
	// GetScalarFunctions returns a set of zero or more function variants that match the provided name.
	GetScalarFunctions(name string) []*extensions.ScalarFunctionVariant

	// GetAggregateFunctions returns a set of zero or more function variants that match the provided name.
	GetAggregateFunctions(name string) []*extensions.AggregateFunctionVariant

	// GetWindowFunctions returns a set of zero or more function variants that match the provided name.
	GetWindowFunctions(name string) []*extensions.WindowFunctionVariant

	GetAllFunctions() []extensions.FunctionVariant
}

type NameKind int

const (
	Substrait NameKind = iota
	Local
)

// LocalFunctionRegistry is a collection of functions localized to a particular Dialect
type LocalFunctionRegistry interface {
	// GetScalarFunctionsBy returns a set of zero or more function variants that match the provided name & kind.
	GetScalarFunctionsBy(name string, kind NameKind) []*LocalScalarFunctionVariant
	GetAggregateFunctionsBy(name string, kind NameKind) []*LocalAggregateFunctionVariant
	GetWindowFunctionsBy(name string, kind NameKind) []*LocalWindowFunctionVariant
	GetDialect() Dialect
}

type FunctionNotation int

const (
	INFIX FunctionNotation = iota
	PREFIX
	POSTFIX
)

// LocalScalarFunctionVariant is a ScalarFunctionVariant that also understands its context in a particular dialect
type LocalScalarFunctionVariant struct {
	extensions.ScalarFunctionVariant
	localName        string
	supportedOptions map[string]extensions.Option
	notation         FunctionNotation
}

func newLocalScalarFunctionVariant(sf *extensions.ScalarFunctionVariant, dfi *dialectFunctionInfo) *LocalScalarFunctionVariant {
	return &LocalScalarFunctionVariant{
		ScalarFunctionVariant: *sf,
		localName:             dfi.LocalName,
		supportedOptions:      dfi.Options,
		notation:              dfi.Notation,
	}
}

func (l *LocalScalarFunctionVariant) LocalName() string {
	return l.localName
}

func (l *LocalScalarFunctionVariant) Notation() FunctionNotation {
	return l.notation
}

func (l *LocalScalarFunctionVariant) IsOptionSupported(name string, value string) bool {
	return isOptionSupported(name, value, l.supportedOptions)
}

type LocalAggregateFunctionVariant struct {
	extensions.AggregateFunctionVariant
	localName        string
	supportedOptions map[string]extensions.Option
	notation         FunctionNotation
}

func newLocalAggregateFunctionVariant(af *extensions.AggregateFunctionVariant, dfi *dialectFunctionInfo) *LocalAggregateFunctionVariant {
	return &LocalAggregateFunctionVariant{
		AggregateFunctionVariant: *af,
		localName:                dfi.LocalName,
		supportedOptions:         dfi.Options,
		notation:                 dfi.Notation,
	}
}

func (l LocalAggregateFunctionVariant) LocalName() string {
	return l.localName
}

func (l LocalAggregateFunctionVariant) Notation() FunctionNotation {
	return l.notation
}

func (l LocalAggregateFunctionVariant) IsOptionSupported(name string, value string) bool {
	return isOptionSupported(name, value, l.supportedOptions)
}

type LocalWindowFunctionVariant struct {
	extensions.WindowFunctionVariant
	localName        string
	supportedOptions map[string]extensions.Option
	notation         FunctionNotation
}

func newLocalWindowFunctionVariant(wf *extensions.WindowFunctionVariant, dfi *dialectFunctionInfo) *LocalWindowFunctionVariant {
	return &LocalWindowFunctionVariant{
		WindowFunctionVariant: *wf,
		localName:             dfi.LocalName,
		supportedOptions:      dfi.Options,
		notation:              dfi.Notation,
	}
}

func (l LocalWindowFunctionVariant) LocalName() string {
	return l.localName
}

func (l LocalWindowFunctionVariant) Notation() FunctionNotation {
	return l.notation
}

func (l LocalWindowFunctionVariant) IsOptionSupported(name string, value string) bool {
	return isOptionSupported(name, value, l.supportedOptions)
}

func isOptionSupported(name string, value string, options map[string]extensions.Option) bool {
	val, exists := options[name]
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
