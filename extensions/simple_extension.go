// SPDX-License-Identifier: Apache-2.0

package extensions

import (
	"github.com/goccy/go-yaml"
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
	Structure  TypeDef
	Parameters []TypeParamDef
}

type TypeVariationFunctions string

const (
	TypeVariationInheritsFuncs TypeVariationFunctions = "INHERITS"
	TypeVariationSeparateFuncs TypeVariationFunctions = "SEPARATE"
)

type TypeVariation struct {
	Name        string
	Parent      TypeDef
	Description string
	Functions   TypeVariationFunctions
}

type Argument interface {
	isArg()
}

type EnumArg struct {
	Name        string `yaml:",omitempty"`
	Description string `yaml:",omitempty"`
	Options     []string
}

func (EnumArg) isArg() {}

type ValueArg struct {
	Name        string `yaml:",omitempty"`
	Description string `yaml:",omitempty"`
	Value       TypeDef
	Constant    bool `yaml:",omitempty"`
}

func (ValueArg) isArg() {}

type TypeArg struct {
	Name        string `yaml:",omitempty"`
	Description string `yaml:",omitempty"`
	Type        string
}

func (TypeArg) isArg() {}

type ArgumentList []Argument

func (a *ArgumentList) UnmarshalYAML(fn func(interface{}) error) error {
	var args []yaml.MapSlice
	if err := fn(&args); err != nil {
		return err
	}

	*a = make(ArgumentList, len(args))
	for i, arg := range args {
		props := arg.ToMap()
		var (
			name, desc string
		)
		if n, ok := props["name"]; ok {
			name = n.(string)
		}
		if d, ok := props["description"]; ok {
			desc = d.(string)
		}

		if opt, ok := props["options"]; ok {
			(*a)[i] = EnumArg{
				Name:        name,
				Description: desc,
				Options:     opt.([]string),
			}
		} else if val, ok := props["value"]; ok {
			var constant bool
			if c, ok := props["constant"]; ok {
				constant = c.(bool)
			}
			(*a)[i] = ValueArg{
				Name:        name,
				Description: desc,
				Value:       val,
				Constant:    constant,
			}
		} else if typ, ok := props["type"]; ok {
			(*a)[i] = TypeArg{
				Name:        name,
				Description: desc,
				Type:        typ.(string),
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
	ParameterConsistency ParameterConsistency `yaml:"parameterConsistency,omitempty"`
}

type ScalarFunctionImpl struct {
	Args             ArgumentList        `yaml:",omitempty"`
	Options          map[string]Option   `yaml:",omitempty"`
	Variadic         VariadicBehavior    `yaml:",omitempty"`
	SessionDependent bool                `yaml:"sessionDependent,omitempty"`
	Deterministic    bool                `yaml:",omitempty"`
	Nullability      NullabilityHandling `yaml:",omitempty"`
	Return           TypeDef             `yaml:",omitempty"`
	Implementation   map[string]string   `yaml:",omitempty"`
}

type ScalarFunction struct {
	Name        string               `yaml:",omitempty"`
	Description string               `yaml:",omitempty,flow"`
	Impls       []ScalarFunctionImpl `yaml:",omitempty"`
}

type DecomposeType string

const (
	DecomposeNone DecomposeType = "NONE"
	DecomposeOne  DecomposeType = "ONE"
	DecomposeMany DecomposeType = "MANY"
)

type AggregateFunctionImpl struct {
	ScalarFunctionImpl `yaml:",inline"`
	Intermediate       TypeDef
	Ordered            bool
	MaxSet             int
	Decomposable       DecomposeType
}

type AggregateFunction struct {
	Name        string
	Description string
	Impls       []AggregateFunctionImpl
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
}

type SimpleExtensionFile struct {
	Types              []Type              `yaml:"types,omitempty"`
	TypeVariations     []TypeVariation     `yaml:"type_variations,omitempty"`
	ScalarFunctions    []ScalarFunction    `yaml:"scalar_functions,omitempty"`
	AggregateFunctions []AggregateFunction `yaml:"aggregate_functions,omitempty"`
	WindowFunctions    []WindowFunction    `yaml:"window_functions,omitempty"`
}
