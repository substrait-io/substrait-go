// SPDX-License-Identifier: Apache-2.0

package extensions

import (
	"fmt"
	"reflect"
	"strings"

	substraitgo "github.com/substrait-io/substrait-go/v8"
	"github.com/substrait-io/substrait-go/v8/types"
	"github.com/substrait-io/substrait-go/v8/types/parser"
)

type ParamType string

const (
	ParamDataType ParamType = "dataType"
	ParamBool     ParamType = "boolean"
	ParamInteger  ParamType = "integer"
	ParamEnum     ParamType = "enumeration"
	ParamString   ParamType = "string"
)

type TypeParamDef struct {
	Name        string
	Description string
	Type        ParamType
	Min         int
	Max         int
	Options     []string
	Optional    bool
}

// should be either a string or map[string]any
type TypeDef any

type Type struct {
	Name       string
	Variadic   bool
	Structure  TypeDef `yaml:",omitempty"`
	Parameters []TypeParamDef
	Metadata   map[string]any `yaml:"metadata,omitempty"`
}

type TypeVariationFunctions string

const (
	TypeVariationInheritsFuncs TypeVariationFunctions = "INHERITS"
	TypeVariationSeparateFuncs TypeVariationFunctions = "SEPARATE"
	EnumTypeString                                    = "req"
)

type TypeVariation struct {
	Name        string
	Parent      string
	Description string
	Functions   TypeVariationFunctions
}

// FuncParameter is an argument of a function in its function definition
type FuncParameter interface {
	toTypeString() string
	argumentMarker() // unexported marker method
	GetTypeExpression() types.FuncDefArgType
}

type EnumArg struct {
	Name        string `yaml:",omitempty"`
	Description string `yaml:",omitempty"`
	Options     []string
}

func (EnumArg) toTypeString() string {
	return EnumTypeString
}

func (v EnumArg) argumentMarker() {}

func (v EnumArg) GetTypeExpression() types.FuncDefArgType {
	return &types.EnumType{Name: v.Name, Options: v.Options}
}

type ValueArg struct {
	Name        string `yaml:",omitempty"`
	Description string `yaml:",omitempty"`
	Value       *parser.TypeExpression
	Constant    bool `yaml:",omitempty"`
}

func (v ValueArg) toTypeString() string {
	return v.Value.ValueType.ShortString()
}

func (v ValueArg) argumentMarker() {}

func (v ValueArg) GetTypeExpression() types.FuncDefArgType {
	return v.Value.ValueType
}

type TypeArg struct {
	Name        string `yaml:",omitempty"`
	Description string `yaml:",omitempty"`
	Type        *parser.TypeExpression
}

func (TypeArg) toTypeString() string { return "type" }

func (v TypeArg) argumentMarker() {}

func (v TypeArg) GetTypeExpression() types.FuncDefArgType {
	return v.Type.ValueType
}

type FuncParameterList []FuncParameter

func (a *FuncParameterList) UnmarshalYAML(fn func(interface{}) error) error {
	var args []map[string]any
	if err := fn(&args); err != nil {
		return err
	}

	*a = make(FuncParameterList, len(args))
	for i, arg := range args {
		var (
			name, desc string
		)
		if n, ok := arg["name"]; ok {
			name = n.(string)
		}
		if d, ok := arg["description"]; ok {
			desc = d.(string)
		}

		if opt, ok := arg["options"]; ok {
			vals := opt.([]any)
			values := make([]string, len(vals))
			for j, v := range vals {
				values[j] = v.(string)
			}
			(*a)[i] = EnumArg{
				Name:        name,
				Description: desc,
				Options:     values,
			}
		} else if val, ok := arg["value"]; ok {
			var constant bool
			if c, ok := arg["constant"]; ok {
				constant = c.(bool)
			}

			arg := ValueArg{
				Name:        name,
				Description: desc,
				Value:       new(parser.TypeExpression),
				Constant:    constant,
			}
			err := arg.Value.UnmarshalYAML(func(v any) error {
				rv := reflect.ValueOf(v)
				if rv.Type().Kind() != reflect.Ptr {
					return substraitgo.ErrInvalidType
				}
				rv.Elem().Set(reflect.ValueOf(val))
				return nil
			})
			if err != nil {
				return fmt.Errorf("failure reading YAML %v", err)
			}

			(*a)[i] = arg

		} else if typ, ok := arg["type"]; ok {
			arg := TypeArg{
				Name:        name,
				Description: desc,
				Type:        new(parser.TypeExpression),
			}
			err := arg.Type.UnmarshalYAML(func(v any) error {
				rv := reflect.ValueOf(v)
				if rv.Type().Kind() != reflect.Ptr {
					return substraitgo.ErrInvalidType
				}
				rv.Elem().Set(reflect.ValueOf(typ))
				return nil
			})
			if err != nil {
				return fmt.Errorf("failure reading YAML %v", err)
			}

			(*a)[i] = arg
		}
	}

	return nil
}

type Option struct {
	Description string `yaml:",omitempty"`
	Values      []string
}

type ParameterConsistency string

const (
	ConsistentParams   ParameterConsistency = "CONSISTENT"
	InconsistentParams ParameterConsistency = "INCONSISTENT"
)

type NullabilityHandling string

const (
	MirrorNullability         NullabilityHandling = "MIRROR"
	DeclaredOutputNullability NullabilityHandling = "DECLARED_OUTPUT"
	DiscreteNullability       NullabilityHandling = "DISCRETE"
)

type VariadicBehavior struct {
	Min                  int                  `yaml:",omitempty"`
	Max                  int                  `yaml:",omitempty"`
	ParameterConsistency ParameterConsistency `yaml:"parameterConsistency,omitempty" default:"CONSISTENT"`
}

func (v *VariadicBehavior) IsValidArgumentCount(count int) bool {
	return v != nil && count >= v.Min && (count <= v.Max || v.Max == 0)
}

func (v *VariadicBehavior) IsValidArgumentPosition(index int) bool {
	return v != nil && index >= 0 && (index <= v.Max || v.Max == 0)
}

type Function interface {
	ResolveURN(urn string) []FunctionVariant
}

type ScalarFunctionImpl struct {
	Args             FuncParameterList      `yaml:",omitempty"`
	Options          map[string]Option      `yaml:",omitempty"`
	Variadic         *VariadicBehavior      `yaml:",omitempty"`
	SessionDependent bool                   `yaml:"sessionDependent,omitempty"`
	Deterministic    bool                   `yaml:",omitempty"`
	Nullability      NullabilityHandling    `yaml:",omitempty" default:"MIRROR"`
	Return           *parser.TypeExpression `yaml:",omitempty"`
	Implementation   map[string]string      `yaml:",omitempty"`
}

func (s *ScalarFunctionImpl) signatureKey() string {
	var b strings.Builder
	for i, a := range s.Args {
		if i != 0 {
			b.WriteByte('_')
		}
		b.WriteString(a.toTypeString())
	}
	return b.String()
}

type ScalarFunction struct {
	Name        string               `yaml:",omitempty"`
	Description string               `yaml:",omitempty,flow"`
	Impls       []ScalarFunctionImpl `yaml:",omitempty"`
	Metadata    map[string]any       `yaml:"metadata,omitempty"`
}

// UnmarshalYAML decodes a ScalarFunction and applies the Substrait default
// for omitted deterministic values.
func (s *ScalarFunction) UnmarshalYAML(fn func(interface{}) error) error {
	type rawImpl struct {
		ScalarFunctionImpl `yaml:",inline"`
		Deterministic      *bool `yaml:"deterministic,omitempty"`
	}
	type rawFn struct {
		Name        string         `yaml:",omitempty"`
		Description string         `yaml:",omitempty,flow"`
		Impls       []rawImpl      `yaml:",omitempty"`
		Metadata    map[string]any `yaml:"metadata,omitempty"`
	}
	var aux rawFn
	if err := fn(&aux); err != nil {
		return err
	}
	s.Name = aux.Name
	s.Description = aux.Description
	s.Metadata = aux.Metadata
	s.Impls = make([]ScalarFunctionImpl, len(aux.Impls))
	for i, ri := range aux.Impls {
		s.Impls[i] = ri.ScalarFunctionImpl
		if ri.Deterministic != nil {
			s.Impls[i].Deterministic = *ri.Deterministic
		} else {
			s.Impls[i].Deterministic = true
		}
	}
	return nil
}

func (s *ScalarFunction) GetVariants(urn string) []*ScalarFunctionVariant {
	out := make([]*ScalarFunctionVariant, len(s.Impls))
	for i, impl := range s.Impls {
		out[i] = &ScalarFunctionVariant{
			name:        s.Name,
			description: s.Description,
			urn:         urn,
			impl:        impl,
			metadata:    s.Metadata,
		}
	}
	return out
}

func (s *ScalarFunction) ResolveURN(urn string) []FunctionVariant {
	out := make([]FunctionVariant, len(s.Impls))
	for i, impl := range s.Impls {
		out[i] = &ScalarFunctionVariant{
			name:        s.Name,
			description: s.Description,
			urn:         urn,
			impl:        impl,
			metadata:    s.Metadata,
		}
	}
	return out
}

type DecomposeType string

const (
	DecomposeNone DecomposeType = "NONE"
	DecomposeOne  DecomposeType = "ONE"
	DecomposeMany DecomposeType = "MANY"
)

type AggregateFunctionImpl struct {
	ScalarFunctionImpl `yaml:",inline"`
	Intermediate       parser.TypeExpression
	Ordered            bool
	MaxSet             int
	Decomposable       DecomposeType
}

type AggregateFunction struct {
	Name        string
	Description string
	Impls       []AggregateFunctionImpl
	Metadata    map[string]any `yaml:"metadata,omitempty"`
}

// UnmarshalYAML decodes an AggregateFunction and applies the Substrait default
// for omitted deterministic values.
func (s *AggregateFunction) UnmarshalYAML(fn func(interface{}) error) error {
	type rawImpl struct {
		AggregateFunctionImpl `yaml:",inline"`
		Deterministic         *bool `yaml:"deterministic,omitempty"`
	}
	type rawFn struct {
		Name        string         `yaml:",omitempty"`
		Description string         `yaml:",omitempty,flow"`
		Impls       []rawImpl      `yaml:",omitempty"`
		Metadata    map[string]any `yaml:"metadata,omitempty"`
	}
	var aux rawFn
	if err := fn(&aux); err != nil {
		return err
	}
	s.Name = aux.Name
	s.Description = aux.Description
	s.Metadata = aux.Metadata
	s.Impls = make([]AggregateFunctionImpl, len(aux.Impls))
	for i, ri := range aux.Impls {
		s.Impls[i] = ri.AggregateFunctionImpl
		if ri.Deterministic != nil {
			s.Impls[i].Deterministic = *ri.Deterministic
		} else {
			s.Impls[i].Deterministic = true
		}
	}
	return nil
}

func (s *AggregateFunction) GetVariants(urn string) []*AggregateFunctionVariant {
	out := make([]*AggregateFunctionVariant, len(s.Impls))
	for i, impl := range s.Impls {
		if impl.Decomposable == "" {
			impl.Decomposable = DecomposeNone
		}
		out[i] = &AggregateFunctionVariant{
			name:        s.Name,
			description: s.Description,
			urn:         urn,
			impl:        impl,
			metadata:    s.Metadata,
		}
	}
	return out
}

func (s *AggregateFunction) ResolveURN(urn string) []FunctionVariant {
	out := make([]FunctionVariant, len(s.Impls))
	for i, impl := range s.Impls {
		out[i] = &AggregateFunctionVariant{
			name:        s.Name,
			description: s.Description,
			urn:         urn,
			impl:        impl,
			metadata:    s.Metadata,
		}
	}
	return out
}

type WindowType string

const (
	StreamingWindow WindowType = "STREAMING"
	PartitionWindow WindowType = "PARTITION"
)

type WindowFunctionImpl struct {
	AggregateFunctionImpl `yaml:",inline"`
	WindowType            WindowType `yaml:"window_type"`
}

type WindowFunction struct {
	Name        string
	Description string
	Impls       []WindowFunctionImpl
	Metadata    map[string]any `yaml:"metadata,omitempty"`
}

// UnmarshalYAML decodes a WindowFunction and applies the Substrait default
// for omitted deterministic values.
func (s *WindowFunction) UnmarshalYAML(fn func(interface{}) error) error {
	type rawImpl struct {
		WindowFunctionImpl `yaml:",inline"`
		Deterministic      *bool `yaml:"deterministic,omitempty"`
	}
	type rawFn struct {
		Name        string         `yaml:",omitempty"`
		Description string         `yaml:",omitempty,flow"`
		Impls       []rawImpl      `yaml:",omitempty"`
		Metadata    map[string]any `yaml:"metadata,omitempty"`
	}
	var aux rawFn
	if err := fn(&aux); err != nil {
		return err
	}
	s.Name = aux.Name
	s.Description = aux.Description
	s.Metadata = aux.Metadata
	s.Impls = make([]WindowFunctionImpl, len(aux.Impls))
	for i, ri := range aux.Impls {
		s.Impls[i] = ri.WindowFunctionImpl
		if ri.Deterministic != nil {
			s.Impls[i].Deterministic = *ri.Deterministic
		} else {
			s.Impls[i].Deterministic = true
		}
	}
	return nil
}

func (s *WindowFunction) GetVariants(urn string) []*WindowFunctionVariant {
	out := make([]*WindowFunctionVariant, len(s.Impls))
	for i, impl := range s.Impls {
		if impl.Decomposable == "" {
			impl.Decomposable = DecomposeNone
		}
		if impl.WindowType == "" {
			impl.WindowType = PartitionWindow
		}
		out[i] = &WindowFunctionVariant{
			name:        s.Name,
			description: s.Description,
			urn:         urn,
			impl:        impl,
			metadata:    s.Metadata,
		}
	}
	return out
}

func (s *WindowFunction) ResolveURN(urn string) []FunctionVariant {
	out := make([]FunctionVariant, len(s.Impls))
	for i, impl := range s.Impls {
		out[i] = &WindowFunctionVariant{
			name:        s.Name,
			description: s.Description,
			urn:         urn,
			impl:        impl,
			metadata:    s.Metadata,
		}
	}
	return out
}

type SimpleExtensionFile struct {
	Urn                string              `yaml:"urn"`
	Metadata           map[string]any      `yaml:"metadata,omitempty"`
	Types              []Type              `yaml:"types,omitempty"`
	TypeVariations     []TypeVariation     `yaml:"type_variations,omitempty"`
	ScalarFunctions    []ScalarFunction    `yaml:"scalar_functions,omitempty"`
	AggregateFunctions []AggregateFunction `yaml:"aggregate_functions,omitempty"`
	WindowFunctions    []WindowFunction    `yaml:"window_functions,omitempty"`
}

type rawSimpleExtensionFile struct {
	Urn                string                 `yaml:"urn"`
	Metadata           map[string]any         `yaml:"metadata,omitempty"`
	Types              []Type                 `yaml:"types,omitempty"`
	TypeVariations     []TypeVariation        `yaml:"type_variations,omitempty"`
	ScalarFunctions    []rawScalarFunction    `yaml:"scalar_functions,omitempty"`
	AggregateFunctions []rawAggregateFunction `yaml:"aggregate_functions,omitempty"`
	WindowFunctions    []rawWindowFunction    `yaml:"window_functions,omitempty"`
}

type rawScalarFunction struct {
	Name        string                  `yaml:",omitempty"`
	Description string                  `yaml:",omitempty,flow"`
	Impls       []rawScalarFunctionImpl `yaml:",omitempty"`
	Metadata    map[string]any          `yaml:"metadata,omitempty"`
}

type rawScalarFunctionImpl struct {
	Args             []rawFuncParameter  `yaml:",omitempty"`
	Options          map[string]Option   `yaml:",omitempty"`
	Variadic         *VariadicBehavior   `yaml:",omitempty"`
	SessionDependent bool                `yaml:"sessionDependent,omitempty"`
	Deterministic    *bool               `yaml:"deterministic,omitempty"`
	Nullability      NullabilityHandling `yaml:",omitempty" default:"MIRROR"`
	Return           any                 `yaml:",omitempty"`
	Implementation   map[string]string   `yaml:",omitempty"`
}

type rawAggregateFunction struct {
	Name        string                     `yaml:",omitempty"`
	Description string                     `yaml:",omitempty,flow"`
	Impls       []rawAggregateFunctionImpl `yaml:",omitempty"`
	Metadata    map[string]any             `yaml:"metadata,omitempty"`
}

type rawAggregateFunctionImpl struct {
	ScalarFunctionImpl rawScalarFunctionImpl `yaml:",inline"`
	Intermediate       any
	Ordered            bool
	MaxSet             int
	Decomposable       DecomposeType
}

type rawWindowFunction struct {
	Name        string                  `yaml:",omitempty"`
	Description string                  `yaml:",omitempty,flow"`
	Impls       []rawWindowFunctionImpl `yaml:",omitempty"`
	Metadata    map[string]any          `yaml:"metadata,omitempty"`
}

type rawWindowFunctionImpl struct {
	AggregateFunctionImpl rawAggregateFunctionImpl `yaml:",inline"`
	WindowType            WindowType               `yaml:"window_type"`
}

type rawFuncParameter struct {
	Name        string   `yaml:",omitempty"`
	Description string   `yaml:",omitempty"`
	Options     []string `yaml:",omitempty"`
	Value       any
	Constant    bool `yaml:",omitempty"`
	Type        any
}

func (s *SimpleExtensionFile) UnmarshalYAML(fn func(interface{}) error) error {
	var raw rawSimpleExtensionFile
	if err := fn(&raw); err != nil {
		return err
	}

	declaredTypes := make(map[string]struct{}, len(raw.Types))
	for _, typ := range raw.Types {
		declaredTypes[typ.Name] = struct{}{}
	}

	resolveUserDefinedType := parser.WithUserDefinedTypeResolver(func(name string) (string, error) {
		if _, ok := declaredTypes[name]; ok {
			return raw.Urn, nil
		}
		return "", fmt.Errorf("%w: user-defined type %q is not declared", substraitgo.ErrInvalidSimpleExtention, name)
	})

	scalarFunctions, err := convertScalarFunctions(raw.ScalarFunctions, resolveUserDefinedType)
	if err != nil {
		return err
	}
	aggregateFunctions, err := convertAggregateFunctions(raw.AggregateFunctions, resolveUserDefinedType)
	if err != nil {
		return err
	}
	windowFunctions, err := convertWindowFunctions(raw.WindowFunctions, resolveUserDefinedType)
	if err != nil {
		return err
	}

	*s = SimpleExtensionFile{
		Urn:                raw.Urn,
		Metadata:           raw.Metadata,
		Types:              raw.Types,
		TypeVariations:     raw.TypeVariations,
		ScalarFunctions:    scalarFunctions,
		AggregateFunctions: aggregateFunctions,
		WindowFunctions:    windowFunctions,
	}
	return nil
}

func convertScalarFunctions(rawFunctions []rawScalarFunction, parseOptions ...parser.ParseOption) ([]ScalarFunction, error) {
	functions := make([]ScalarFunction, len(rawFunctions))
	for i, rawFunction := range rawFunctions {
		impls, err := convertScalarFunctionImpls(rawFunction.Impls, parseOptions...)
		if err != nil {
			return nil, fmt.Errorf("scalar function %q: %w", rawFunction.Name, err)
		}
		functions[i] = ScalarFunction{Name: rawFunction.Name, Description: rawFunction.Description, Impls: impls, Metadata: rawFunction.Metadata}
	}
	return functions, nil
}

func convertAggregateFunctions(rawFunctions []rawAggregateFunction, parseOptions ...parser.ParseOption) ([]AggregateFunction, error) {
	functions := make([]AggregateFunction, len(rawFunctions))
	for i, rawFunction := range rawFunctions {
		impls := make([]AggregateFunctionImpl, len(rawFunction.Impls))
		for j, rawImpl := range rawFunction.Impls {
			scalarImpl, err := convertScalarFunctionImpl(rawImpl.ScalarFunctionImpl, parseOptions...)
			if err != nil {
				return nil, fmt.Errorf("aggregate function %q impl %d: %w", rawFunction.Name, j, err)
			}
			intermediate, err := parseTypeExpression(rawImpl.Intermediate, parseOptions...)
			if err != nil {
				return nil, fmt.Errorf("aggregate function %q impl %d intermediate: %w", rawFunction.Name, j, err)
			}
			impls[j] = AggregateFunctionImpl{ScalarFunctionImpl: scalarImpl, Intermediate: intermediate, Ordered: rawImpl.Ordered, MaxSet: rawImpl.MaxSet, Decomposable: rawImpl.Decomposable}
		}
		functions[i] = AggregateFunction{Name: rawFunction.Name, Description: rawFunction.Description, Impls: impls, Metadata: rawFunction.Metadata}
	}
	return functions, nil
}

func convertWindowFunctions(rawFunctions []rawWindowFunction, parseOptions ...parser.ParseOption) ([]WindowFunction, error) {
	functions := make([]WindowFunction, len(rawFunctions))
	for i, rawFunction := range rawFunctions {
		impls := make([]WindowFunctionImpl, len(rawFunction.Impls))
		for j, rawImpl := range rawFunction.Impls {
			aggImpl, err := convertAggregateFunctionImpl(rawImpl.AggregateFunctionImpl, parseOptions...)
			if err != nil {
				return nil, fmt.Errorf("window function %q impl %d: %w", rawFunction.Name, j, err)
			}
			impls[j] = WindowFunctionImpl{AggregateFunctionImpl: aggImpl, WindowType: rawImpl.WindowType}
		}
		functions[i] = WindowFunction{Name: rawFunction.Name, Description: rawFunction.Description, Impls: impls, Metadata: rawFunction.Metadata}
	}
	return functions, nil
}

func convertAggregateFunctionImpl(rawImpl rawAggregateFunctionImpl, parseOptions ...parser.ParseOption) (AggregateFunctionImpl, error) {
	scalarImpl, err := convertScalarFunctionImpl(rawImpl.ScalarFunctionImpl, parseOptions...)
	if err != nil {
		return AggregateFunctionImpl{}, err
	}
	intermediate, err := parseTypeExpression(rawImpl.Intermediate, parseOptions...)
	if err != nil {
		return AggregateFunctionImpl{}, fmt.Errorf("intermediate: %w", err)
	}
	return AggregateFunctionImpl{ScalarFunctionImpl: scalarImpl, Intermediate: intermediate, Ordered: rawImpl.Ordered, MaxSet: rawImpl.MaxSet, Decomposable: rawImpl.Decomposable}, nil
}

func convertScalarFunctionImpls(rawImpls []rawScalarFunctionImpl, parseOptions ...parser.ParseOption) ([]ScalarFunctionImpl, error) {
	impls := make([]ScalarFunctionImpl, len(rawImpls))
	for i, rawImpl := range rawImpls {
		impl, err := convertScalarFunctionImpl(rawImpl, parseOptions...)
		if err != nil {
			return nil, fmt.Errorf("impl %d: %w", i, err)
		}
		impls[i] = impl
	}
	return impls, nil
}

func convertScalarFunctionImpl(rawImpl rawScalarFunctionImpl, parseOptions ...parser.ParseOption) (ScalarFunctionImpl, error) {
	args, err := convertFuncParameters(rawImpl.Args, parseOptions...)
	if err != nil {
		return ScalarFunctionImpl{}, err
	}
	returnType, err := parseTypeExpression(rawImpl.Return, parseOptions...)
	if err != nil {
		return ScalarFunctionImpl{}, fmt.Errorf("return: %w", err)
	}
	deterministic := true
	if rawImpl.Deterministic != nil {
		deterministic = *rawImpl.Deterministic
	}
	return ScalarFunctionImpl{
		Args:             args,
		Options:          rawImpl.Options,
		Variadic:         rawImpl.Variadic,
		SessionDependent: rawImpl.SessionDependent,
		Deterministic:    deterministic,
		Nullability:      rawImpl.Nullability,
		Return:           &returnType,
		Implementation:   rawImpl.Implementation,
	}, nil
}

func convertFuncParameters(rawArgs []rawFuncParameter, parseOptions ...parser.ParseOption) (FuncParameterList, error) {
	args := make(FuncParameterList, len(rawArgs))
	for i, rawArg := range rawArgs {
		arg, err := convertFuncParameter(rawArg, parseOptions...)
		if err != nil {
			return nil, fmt.Errorf("arg %d: %w", i, err)
		}
		args[i] = arg
	}
	return args, nil
}

func convertFuncParameter(rawArg rawFuncParameter, parseOptions ...parser.ParseOption) (FuncParameter, error) {
	if len(rawArg.Options) > 0 {
		return EnumArg{Name: rawArg.Name, Description: rawArg.Description, Options: rawArg.Options}, nil
	}
	if rawArg.Value != nil {
		value, err := parseTypeExpression(rawArg.Value, parseOptions...)
		if err != nil {
			return nil, err
		}
		return ValueArg{Name: rawArg.Name, Description: rawArg.Description, Value: &value, Constant: rawArg.Constant}, nil
	}
	if rawArg.Type != nil {
		typ, err := parseTypeExpression(rawArg.Type, parseOptions...)
		if err != nil {
			return nil, err
		}
		return TypeArg{Name: rawArg.Name, Description: rawArg.Description, Type: &typ}, nil
	}
	return nil, substraitgo.ErrInvalidSimpleExtention
}

func parseTypeExpression(value any, parseOptions ...parser.ParseOption) (parser.TypeExpression, error) {
	if value == nil {
		return parser.TypeExpression{}, nil
	}
	typeString, ok := value.(string)
	if !ok {
		return parser.TypeExpression{}, substraitgo.ErrInvalidType
	}
	typ, err := parser.ParseType(typeString, parseOptions...)
	if err != nil {
		return parser.TypeExpression{}, err
	}
	return parser.TypeExpression{ValueType: typ}, nil
}
