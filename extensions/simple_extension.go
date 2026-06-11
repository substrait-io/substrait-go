// SPDX-License-Identifier: Apache-2.0

package extensions

import (
	"fmt"
	"strings"

	"github.com/goccy/go-yaml"
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

func (a *FuncParameterList) UnmarshalYAML(data []byte) error {
	return parseFuncParameterList(data, parser.TypeParser{}, a)
}

func parseFuncParameterList(data []byte, typeParser parser.TypeParser, out *FuncParameterList) error {
	var args []map[string]any
	if err := yaml.Unmarshal(data, &args); err != nil {
		return err
	}

	*out = make(FuncParameterList, len(args))
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
			(*out)[i] = EnumArg{
				Name:        name,
				Description: desc,
				Options:     values,
			}
		} else if val, ok := arg["value"]; ok {
			var constant bool
			if c, ok := arg["constant"]; ok {
				constant = c.(bool)
			}

			valueType, err := parseTypeExpressionFromYAMLValue(typeParser, val)
			if err != nil {
				return fmt.Errorf("failure reading YAML %v", err)
			}

			(*out)[i] = ValueArg{
				Name:        name,
				Description: desc,
				Value:       &valueType,
				Constant:    constant,
			}

		} else if typ, ok := arg["type"]; ok {
			typeExpr, err := parseTypeExpressionFromYAMLValue(typeParser, typ)
			if err != nil {
				return fmt.Errorf("failure reading YAML %v", err)
			}

			(*out)[i] = TypeArg{
				Name:        name,
				Description: desc,
				Type:        &typeExpr,
			}
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
	Dependencies       map[string]string   `yaml:"dependencies,omitempty"`
	Metadata           map[string]any      `yaml:"metadata,omitempty"`
	Types              []Type              `yaml:"types,omitempty"`
	TypeVariations     []TypeVariation     `yaml:"type_variations,omitempty"`
	ScalarFunctions    []ScalarFunction    `yaml:"scalar_functions,omitempty"`
	AggregateFunctions []AggregateFunction `yaml:"aggregate_functions,omitempty"`
	WindowFunctions    []WindowFunction    `yaml:"window_functions,omitempty"`
}

func (s *SimpleExtensionFile) UnmarshalYAML(data []byte) error {
	type typeDeclarations struct {
		Urn   string `yaml:"urn"`
		Types []Type `yaml:"types,omitempty"`
	}
	var declarations typeDeclarations
	if err := yaml.Unmarshal(data, &declarations); err != nil {
		return err
	}

	declaredTypes := make(map[string]struct{}, len(declarations.Types))
	for _, typ := range declarations.Types {
		declaredTypes[typ.Name] = struct{}{}
	}

	typeParser := parser.TypeParser{
		ResolveUserDefinedType: func(name string, nullability types.Nullability, parameters []types.UDTParameter) (*types.ParameterizedUserDefinedType, error) {
			if _, ok := declaredTypes[name]; !ok {
				return nil, fmt.Errorf("%w: user-defined type %q is not declared", substraitgo.ErrInvalidSimpleExtention, name)
			}
			return &types.ParameterizedUserDefinedType{
				Name:           name,
				URN:            declarations.Urn,
				Nullability:    nullability,
				TypeParameters: parameters,
			}, nil
		},
	}

	type rawFile SimpleExtensionFile
	var raw rawFile
	if err := yaml.UnmarshalWithOptions(
		data,
		&raw,
		yaml.CustomUnmarshaler[parser.TypeExpression](func(out *parser.TypeExpression, data []byte) error {
			parsed, err := parseTypeExpression(data, typeParser)
			if err != nil {
				return err
			}
			*out = parsed
			return nil
		}),
		yaml.CustomUnmarshaler[FuncParameterList](func(out *FuncParameterList, data []byte) error {
			return parseFuncParameterList(data, typeParser, out)
		}),
	); err != nil {
		return err
	}

	*s = SimpleExtensionFile(raw)
	return nil
}

func parseTypeExpressionFromYAMLValue(typeParser parser.TypeParser, value any) (parser.TypeExpression, error) {
	typeString, ok := value.(string)
	if !ok {
		return parser.TypeExpression{}, substraitgo.ErrInvalidType
	}
	return parseTypeExpression([]byte(typeString), typeParser)
}

func parseTypeExpression(data []byte, typeParser parser.TypeParser) (parser.TypeExpression, error) {
	var typeString string
	if err := yaml.Unmarshal(data, &typeString); err != nil {
		return parser.TypeExpression{}, err
	}

	typ, err := typeParser.Parse(typeString)
	if err != nil {
		return parser.TypeExpression{}, err
	}
	return parser.TypeExpression{ValueType: typ}, nil
}
