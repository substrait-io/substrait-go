// SPDX-License-Identifier: Apache-2.0

package functions

import (
	"fmt"
	"io"
	"strings"

	"github.com/goccy/go-yaml"
	substraitgo "github.com/substrait-io/substrait-go"
	"github.com/substrait-io/substrait-go/extensions"
)

func LoadDialect(name string, r io.Reader) (Dialect, error) {
	dialect, err := newDialect(name, r)
	if err != nil {
		return nil, err
	}
	return dialect, nil
}

type dialectFunctionInfo struct {
	ID        extensions.ID
	Name      string
	LocalName string
	Options   map[string]extensions.Option
	Notation  FunctionNotation
}

type dialectImpl struct {
	name string
	file dialectFile

	toLocalTypeMap map[string]dialectTypeInfo // substrait short type name to dialectTypeInfo

	localScalarFunctions    map[extensions.ID]*dialectFunctionInfo
	localAggregateFunctions map[extensions.ID]*dialectFunctionInfo
	localWindowFunctions    map[extensions.ID]*dialectFunctionInfo
}

func (d *dialectImpl) Name() string {
	return d.name
}

func (d *dialectImpl) LocalizeFunctionRegistry(registry FunctionRegistry) (LocalFunctionRegistry, error) {
	scalarFunctions, err := makeLocalFunctionVariantMap(d.localScalarFunctions, registry.GetScalarFunctionsByName, newLocalScalarFunctionVariant)
	if err != nil {
		return nil, err
	}
	aggregateFunctions, err := makeLocalFunctionVariantMap(d.localAggregateFunctions, registry.GetAggregateFunctionsByName, newLocalAggregateFunctionVariant)
	if err != nil {
		return nil, err
	}
	windowFunctions, err := makeLocalFunctionVariantMap(d.localWindowFunctions, registry.GetWindowFunctionsByName, newLocalWindowFunctionVariant)
	if err != nil {
		return nil, err
	}

	return &localFunctionRegistryImpl{
		dialect:                 d,
		scalarFunctions:         scalarFunctions,
		aggregateFunctions:      aggregateFunctions,
		windowFunctions:         windowFunctions,
		localScalarFunctions:    buildReverseMap(scalarFunctions),
		localAggregateFunctions: buildReverseMap(aggregateFunctions),
		localWindowFunctions:    buildReverseMap(windowFunctions),
	}, nil
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

type withID interface {
	ID() extensions.ID
}

func makeLocalFunctionVariantMap[T withID, V any](dialectFunctionInfos map[extensions.ID]*dialectFunctionInfo, getFunctionVariants func(string) []T, createLocalVariant func(T, *dialectFunctionInfo) *V) (map[string][]*V, error) {
	processedFunctions := make(map[extensions.ID]bool)
	localFunctionVariants := make(map[string][]*V)
	for _, dfi := range dialectFunctionInfos {
		if _, nameAlreadyProcessed := localFunctionVariants[dfi.Name]; nameAlreadyProcessed {
			if _, ok := processedFunctions[dfi.ID]; !ok {
				return nil, fmt.Errorf("%w: no function variant found for '%s'", substraitgo.ErrInvalidDialect, dfi.ID)
			}
			continue
		}

		localVariantArray := make([]*V, 0)
		for _, f := range getFunctionVariants(dfi.Name) {
			if dfi, ok := dialectFunctionInfos[f.ID()]; ok {
				localVariantArray = append(localVariantArray, createLocalVariant(f, dfi))
				processedFunctions[f.ID()] = true
			}
		}
		if _, ok := processedFunctions[dfi.ID]; !ok {
			return nil, fmt.Errorf("%w: no function variant found for '%s'", substraitgo.ErrInvalidDialect, dfi.ID)
		}
		if len(localVariantArray) > 0 {
			localFunctionVariants[dfi.Name] = localVariantArray
		}
	}
	return localFunctionVariants, nil
}

func (d *dialectImpl) LocalizeTypeRegistry(registry TypeRegistry) (LocalTypeRegistry, error) {
	typeInfos := make([]typeInfo, 0, len(d.toLocalTypeMap))
	for name, info := range d.toLocalTypeMap {
		// TODO use registry.GetTypeClasses
		typ, err := getTypeFromBaseTypeName(name)
		if err != nil {
			return nil, fmt.Errorf("%w: unknown type %v", substraitgo.ErrInvalidDialect, name)
		}
		typeInfos = append(typeInfos, typeInfo{typ: typ, shortName: name, localName: info.SqlTypeName, supportedAsColumn: info.SupportedAsColumn})
	}
	return NewLocalTypeRegistry(typeInfos), nil
}

func newDialect(name string, reader io.Reader) (Dialect, error) {
	dialect := &dialectImpl{name: name}
	if err := dialect.Load(reader); err != nil {
		return nil, err
	}

	return dialect, nil
}

func (d *dialectImpl) Load(reader io.Reader) error {
	dec := yaml.NewDecoder(reader)
	if err := dec.Decode(&d.file); err != nil {
		return err
	}

	err := d.file.validate()
	if err != nil {
		return err
	}
	d.toLocalTypeMap = d.file.SupportedTypes
	d.localScalarFunctions = d.buildFunctionInfoMap(d.file.ScalarFunctions)
	d.localAggregateFunctions = d.buildFunctionInfoMap(d.file.AggregateFunctions)
	d.localWindowFunctions = d.buildFunctionInfoMap(d.file.WindowFunctions)
	return nil
}

func (d *dialectImpl) buildFunctionInfoMap(functions []dialectFunction) map[extensions.ID]*dialectFunctionInfo {
	funcMap := make(map[extensions.ID]*dialectFunctionInfo)
	for _, f := range functions {
		uri, name := d.file.getUriAndFunctionName(&f)
		for _, kernel := range f.SupportedKernels {
			localName := f.LocalName
			if len(localName) == 0 {
				localName = name
			}
			id := extensions.ID{URI: uri, Name: name + ":" + kernel}
			localFunction := dialectFunctionInfo{
				ID:        id,
				Name:      name,
				LocalName: localName,
				Notation:  f.GetNotation(),
				Options:   f.getOptions(),
			}
			funcMap[id] = &localFunction
		}
	}
	return funcMap
}

type dialectFile struct {
	Name               string                     `yaml:"name"`
	Type               string                     `yaml:"type"`
	SupportedTypes     map[string]dialectTypeInfo `yaml:"supported_types,omitempty"`
	ScalarFunctions    []dialectFunction          `yaml:"scalar_functions,omitempty"`
	AggregateFunctions []dialectFunction          `yaml:"aggregate_functions,omitempty"`
	WindowFunctions    []dialectFunction          `yaml:"window_functions,omitempty"`
	Dependencies       map[string]string          `yaml:"dependencies,omitempty"`
}

type dialectTypeInfo struct {
	SqlTypeName       string `yaml:"sql_type_name"`
	SupportedAsColumn bool   `yaml:"supported_as_column"`
}

func (d *dialectFile) getUriAndFunctionName(df *dialectFunction) (string, string) {
	parts := strings.Split(df.Name, ".")
	return d.Dependencies[parts[0]], parts[1]
}

func (d *dialectFile) validate() error {
	if len(d.SupportedTypes) == 0 {
		return fmt.Errorf("%w: no supported types provided", substraitgo.ErrInvalidDialect)
	}
	if len(d.ScalarFunctions) == 0 && len(d.AggregateFunctions) == 0 && len(d.WindowFunctions) == 0 {
		return fmt.Errorf("%w: no functions provided", substraitgo.ErrInvalidDialect)
	}

	validate := func(functions []dialectFunction) error {
		for _, f := range functions {
			if err := d.validateFunction(&f); err != nil {
				return err
			}
		}
		return nil
	}

	if err := validate(d.ScalarFunctions); err != nil {
		return err
	}
	if err := validate(d.AggregateFunctions); err != nil {
		return err
	}
	return validate(d.WindowFunctions)
}

func (d *dialectFile) validateFunction(df *dialectFunction) error {
	if len(df.Name) == 0 || len(df.SupportedKernels) == 0 {
		return fmt.Errorf("%w: invalid function", substraitgo.ErrInvalidDialect)
	}
	parts := strings.Split(df.Name, ".")
	if len(parts) != 2 {
		return fmt.Errorf("%w: invalid function name '%s'", substraitgo.ErrInvalidDialect, df.Name)
	}
	if _, ok := d.Dependencies[parts[0]]; !ok {
		return fmt.Errorf("%w: unknown dependency '%s' in function", substraitgo.ErrInvalidDialect, parts[0])
	}

	if err := df.validateAndFixKernels(); err != nil {
		return err
	}
	return nil
}

type dialectFunction struct {
	Name             string            `yaml:"name"`
	LocalName        string            `yaml:"local_name"`
	Infix            bool              `yaml:"infix"`
	Postfix          bool              `yaml:"postfix"`
	RequiredOptions  map[string]string `yaml:"required_options,omitempty"`
	Aggregate        bool              `yaml:"aggregate,omitempty"`
	SupportedKernels []string          `yaml:"supported_kernels,omitempty"`
	// TODO handle variadic behavior in dialect,
	// if infix is true variadic min cannot be 1?
}

func (df *dialectFunction) validateAndFixKernels() error {
	for i, kernel := range df.SupportedKernels {
		argTypes := strings.Split(kernel, "_")
		hasAnyType := false
		for i, argType := range argTypes {
			if strings.HasPrefix(argType, "any") {
				argTypes[i] = "any"
				hasAnyType = true
			} else if argType != "" && !isSupportedType(argType) {
				return fmt.Errorf("%w: unsupported type '%s'", substraitgo.ErrInvalidDialect, argType)
			}
		}
		if hasAnyType {
			df.SupportedKernels[i] = strings.Join(argTypes, "_")
		}
	}
	return nil
}

func (df *dialectFunction) getOptions() map[string]extensions.Option {
	if len(df.RequiredOptions) == 0 {
		return nil
	}
	options := make(map[string]extensions.Option)
	for k, v := range df.RequiredOptions {
		options[k] = extensions.Option{Values: []string{v}}
	}
	return options
}

func (df *dialectFunction) GetNotation() FunctionNotation {
	if df.Infix {
		return INFIX
	}
	if df.Postfix {
		return POSTFIX
	}
	return PREFIX
}
