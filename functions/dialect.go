// SPDX-License-Identifier: Apache-2.0

package functions

import (
	"fmt"
	"github.com/goccy/go-yaml"
	substraitgo "github.com/substrait-io/substrait-go"
	"github.com/substrait-io/substrait-go/extensions"
	"io"
	"strings"
)

var loadedDialects = make(map[string]Dialect)

func LoadDialect(name string, r io.Reader) error {
	if _, ok := loadedDialects[name]; ok {
		return fmt.Errorf("%w: dialect '%s' already loaded", substraitgo.ErrKeyExists, name)
	}

	dialect, err := newDialect(name, r)
	if err != nil {
		return err
	}
	loadedDialects[name] = dialect
	return nil
}

func GetDialect(name string) (Dialect, error) {
	if d, ok := loadedDialects[name]; ok {
		return d, nil
	}
	return nil, fmt.Errorf("%w: dialect '%s' not found", substraitgo.ErrNotFound, name)
}

func GetSupportedDialects() []string {
	var ret []string
	for k := range loadedDialects {
		ret = append(ret, k)
	}
	return ret
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

	toLocalTypeMap map[string]dialectTypeInfo // substrait type name to dialectTypeInfo

	binaryFunctions []extensions.FunctionVariant

	localScalarFunctions    map[extensions.ID]*dialectFunctionInfo
	localAggregateFunctions map[extensions.ID]*dialectFunctionInfo
	localWindowFunctions    map[extensions.ID]*dialectFunctionInfo
}

func (d *dialectImpl) Name() string {
	return d.name
}

func (d *dialectImpl) LocalizeFunctionRegistry(registry FunctionRegistry) (LocalFunctionRegistry, error) {
	scalarFunctions := makeLocalFunctionVariantMap(d.localScalarFunctions, registry.GetScalarFunctions, newLocalScalarFunctionVariant)
	aggregateFunctions := makeLocalFunctionVariantMap(d.localAggregateFunctions, registry.GetAggregateFunctions, newLocalAggregateFunctionVariant)
	windowFunctions := makeLocalFunctionVariantMap(d.localWindowFunctions, registry.GetWindowFunctions, newLocalWindowFunctionVariant)

	return newLocalFunctionRegistry(d, scalarFunctions, aggregateFunctions, windowFunctions), nil
}

type withID interface {
	ID() extensions.ID
}

func makeLocalFunctionVariantMap[T withID, V any](dialectFunctionInfos map[extensions.ID]*dialectFunctionInfo, getFunctionVariants func(string) []T, createLocalVariant func(T, *dialectFunctionInfo) *V) map[string][]*V {
	localFunctionVariants := make(map[string][]*V)
	for _, dfi := range dialectFunctionInfos {
		if _, ok := localFunctionVariants[dfi.Name]; ok {
			// skip if function name is already added
			continue
		}
		localVariantArray := make([]*V, 0)
		for _, f := range getFunctionVariants(dfi.Name) {
			if dfi, ok := dialectFunctionInfos[f.ID()]; ok {
				// add functionVariant to local registry only if it is present in the dialect
				localVariantArray = append(localVariantArray, createLocalVariant(f, dfi))
			}
		}
		if len(localVariantArray) != 0 {
			localFunctionVariants[dfi.Name] = localVariantArray
		}
	}
	return localFunctionVariants
}

func makeLocalScalarFunctionVariantMap(registry FunctionRegistry, dialectFunctionInfos map[extensions.ID]*dialectFunctionInfo, localFunctionVariants map[string][]*LocalScalarFunctionVariant) {
	for _, dfi1 := range dialectFunctionInfos {
		if _, ok := localFunctionVariants[dfi1.Name]; ok {
			continue
		}
		lsfArray := make([]*LocalScalarFunctionVariant, 0)
		for _, sf := range registry.GetScalarFunctions(dfi1.Name) {
			if dfi, ok := dialectFunctionInfos[sf.ID()]; ok {
				lsf := LocalScalarFunctionVariant{
					ScalarFunctionVariant: *sf,
					localName:             dfi.LocalName,
					supportedOptions:      dfi.Options,
					notation:              dfi.Notation,
				}
				lsfArray = append(lsfArray, &lsf)
			}
		}
		if len(lsfArray) != 0 {
			localFunctionVariants[dfi1.Name] = lsfArray
		}
	}
}

func (d *dialectImpl) LocalizeTypeRegistry(registry TypeRegistry) (LocalTypeRegistry, error) {
	typeInfos := make([]typeInfo, 0, len(d.toLocalTypeMap))
	for name, info := range d.toLocalTypeMap {
		typ, err := registry.GetTypeFromTypeString(name)
		if err != nil {
			return nil, err
		}
		typeInfos = append(typeInfos, typeInfo{typ: typ, name: name, localName: info.SqlTypeName, supportedAsColumn: info.SupportedAsColumn})
	}
	return NewLocalTypeRegistry(typeInfos), nil
}

type LocalFunctionID struct {
	Name      string
	Signature string
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

func (d *dialectImpl) buildTypeMap() {
	d.toLocalTypeMap = d.file.SupportedTypes
}

func (d *dialectImpl) GetBinaryFunctions() []extensions.FunctionVariant {
	return d.binaryFunctions
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
		return substraitgo.ErrInvalidDialect
	}
	if len(d.ScalarFunctions) == 0 && len(d.AggregateFunctions) == 0 && len(d.WindowFunctions) == 0 {
		return substraitgo.ErrInvalidDialect
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
	if len(df.Name) == 0 {
		return substraitgo.ErrInvalidDialect
	}
	parts := strings.Split(df.Name, ".")
	if len(parts) != 2 {
		return substraitgo.ErrInvalidDialect
	}
	if _, ok := d.Dependencies[parts[0]]; !ok {
		return substraitgo.ErrInvalidDialect
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
