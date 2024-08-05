// SPDX-License-Identifier: Apache-2.0

package extensions

import (
	"github.com/goccy/go-yaml"
	"io"
	"strings"
)

type FuncKind string

const (
	ScalarFunc    FuncKind = "scalar"
	AggregateFunc FuncKind = "aggregate"
	WindowFunc    FuncKind = "window"
)

type Dialect interface {
	// LocalLookup : Lookup a function by its local name and argument types
	LocalLookup(funcName string, argTypes []string) (FunctionVariant, bool)

	// GetBinaryFunctions : Get all substrait binary functions supported in this dialect
	GetBinaryFunctions() []FunctionVariant
}

type DialectFunctionBinding interface {
	GetLocalName() string
	GetName() string
	GetFuncKind() FuncKind
	GetLocalArgumentTypes() []string
	GetOptions() map[string]Option
}

type LocalFunctionVariant struct {
	ID            ID
	SubstraitFunc FunctionVariant
	FuncKind      FuncKind
	Name          string
	Signature     string
	LocalName     string
	ArgLocalTypes []string
	Infix         bool
	Options       map[string]Option
}

func (lf *LocalFunctionVariant) GetLocalName() string {
	return lf.LocalName
}

func (lf *LocalFunctionVariant) GetName() string {
	return lf.Name
}

func (lf *LocalFunctionVariant) GetFuncKind() FuncKind {
	return lf.FuncKind
}

func (lf *LocalFunctionVariant) GetLocalArgumentTypes() []string {
	return lf.ArgLocalTypes
}

func (lf *LocalFunctionVariant) GetOptions() map[string]Option {
	return lf.Options
}

type dialectImpl struct {
	file DialectFile

	fromLocalTypeMap map[string]string // local to substrait type
	toLocalTypeMap   map[string]string // local to substrait type

	binaryFunctions []FunctionVariant

	localScalarFunctions    map[ID]*LocalFunctionVariant
	localAggregateFunctions map[ID]*LocalFunctionVariant
	localWindowFunctions    map[ID]*LocalFunctionVariant

	functions map[LocalFunctionID]*LocalFunctionVariant
}

type LocalFunctionID struct {
	Name      string
	Signature string
}

func newDialect(name string, reader io.Reader) (Dialect, error) {
	dialect := &dialectImpl{}
	if err := dialect.Load(reader); err != nil {
		return nil, err
	}

	for _, lsf := range dialect.localScalarFunctions {
		lsf.SubstraitFunc.(*ScalarFunctionVariant).AddDialect(name, lsf)
	}
	for _, laf := range dialect.localAggregateFunctions {
		laf.SubstraitFunc.(*AggregateFunctionVariant).AddDialect(name, laf)
	}
	for _, lwf := range dialect.localWindowFunctions {
		lwf.SubstraitFunc.(*WindowFunctionVariant).AddDialect(name, lwf)
	}
	return dialect, nil
}

func (d *dialectImpl) Load(reader io.Reader) error {
	dec := yaml.NewDecoder(reader)
	if err := dec.Decode(&d.file); err != nil {
		return err
	}

	d.buildTypeMap()
	d.localScalarFunctions = d.buildLocalFunctionMap(d.file.ScalarFunctions, ScalarFunc)
	d.localAggregateFunctions = d.buildLocalFunctionMap(d.file.AggregateFunctions, AggregateFunc)
	d.localWindowFunctions = d.buildLocalFunctionMap(d.file.WindowFunctions, WindowFunc)

	d.functions = make(map[LocalFunctionID]*LocalFunctionVariant)
	d.updateFunctionMap(d.localScalarFunctions)
	d.updateFunctionMap(d.localAggregateFunctions)
	d.updateFunctionMap(d.localWindowFunctions)

	return nil
}

func (d *dialectImpl) buildLocalFunctionMap(functions []DialectFunction, funcKind FuncKind) map[ID]*LocalFunctionVariant {
	funcMap := make(map[ID]*LocalFunctionVariant)
	for _, f := range functions {
		uri, name := d.file.getUriAndFunctionName(&f)
		for _, kernel := range f.SupportedKernels {
			localName := f.LocalName
			if len(localName) == 0 {
				localName = name
			}
			id := ID{uri, name + ":" + kernel}
			substraitFunc, ok := getSubstraitFunc(funcKind, id)
			if !ok {
				continue
			}
			lf := LocalFunctionVariant{
				ID:            id,
				SubstraitFunc: substraitFunc,
				FuncKind:      funcKind,
				Name:          name,
				LocalName:     localName,
				Infix:         f.Infix,
				Signature:     kernel,
				ArgLocalTypes: d.getArgLocalTypes(kernel),
				Options:       f.getOptions(),
			}
			funcMap[lf.ID] = &lf
			if funcKind == ScalarFunc && len(lf.ArgLocalTypes) == 2 {
				d.binaryFunctions = append(d.binaryFunctions, lf.SubstraitFunc)
			}
		}
	}
	return funcMap
}

func (d *dialectImpl) updateFunctionMap(localFuncMap map[ID]*LocalFunctionVariant) {
	for _, lf := range localFuncMap {
		name := lf.LocalName
		d.functions[LocalFunctionID{name, lf.Signature}] = lf
	}
}

func (d *dialectImpl) buildTypeMap() {
	d.toLocalTypeMap = d.file.SupportedTypes
	d.fromLocalTypeMap = make(map[string]string)
	for k, v := range d.file.SupportedTypes {
		d.fromLocalTypeMap[v] = k
	}
}

func getSubstraitFunc(funcKind FuncKind, id ID) (FunctionVariant, bool) {
	switch funcKind {
	case ScalarFunc:
		return DefaultCollection.GetScalarFunc(id)
	case AggregateFunc:
		return DefaultCollection.GetAggregateFunc(id)
	case WindowFunc:
		return DefaultCollection.GetWindowFunc(id)
	}
	return nil, false
}

func (d *dialectImpl) LocalLookup(funcName string, argTypes []string) (FunctionVariant, bool) {
	localId := LocalFunctionID{Name: funcName, Signature: d.getSignature(argTypes)}
	lf, ok := d.functions[localId]
	if !ok {
		return nil, false
	}
	return lf.SubstraitFunc, true
}

func (d *dialectImpl) GetBinaryFunctions() []FunctionVariant {
	return d.binaryFunctions
}

func (d *dialectImpl) getSignature(argLocalTypes []string) string {
	if len(argLocalTypes) == 0 {
		return ""
	}

	var sb strings.Builder
	for i, p := range argLocalTypes {
		if i != 0 {
			sb.WriteByte('_')
		}
		sb.WriteString(d.fromLocalTypeMap[p])
	}
	return sb.String()
}

func (d *dialectImpl) getArgLocalTypes(kernel string) []string {
	parts := strings.Split(kernel, "_")
	argLocalTypes := make([]string, len(parts))
	for i, p := range parts {
		argLocalTypes[i] = d.toLocalTypeMap[p]
	}
	return argLocalTypes
}

type DialectFile struct {
	Name               string            `yaml:"name"`
	Type               string            `yaml:"type"`
	SupportedTypes     map[string]string `yaml:"supported_types,omitempty"`
	ScalarFunctions    []DialectFunction `yaml:"scalar_functions,omitempty"`
	AggregateFunctions []DialectFunction `yaml:"aggregate_functions,omitempty"`
	WindowFunctions    []DialectFunction `yaml:"window_functions,omitempty"`
	Dependencies       map[string]string `yaml:"dependencies,omitempty"`
}

func (d *DialectFile) getUriAndFunctionName(df *DialectFunction) (string, string) {
	parts := strings.Split(df.Name, ".")
	return d.Dependencies[parts[0]], parts[1]
}

type DialectFunction struct {
	Name             string            `yaml:"name"`
	LocalName        string            `yaml:"local_name"`
	Infix            bool              `yaml:"infix"`
	Postfix          bool              `yaml:"postfix"`
	RequiredOptions  map[string]string `yaml:"required_options,omitempty"`
	Aggregate        bool              `yaml:"aggregate,omitempty"`
	SupportedKernels []string          `yaml:"supported_kernels,omitempty"`
}

func (df *DialectFunction) getOptions() map[string]Option {
	if len(df.RequiredOptions) == 0 {
		return nil
	}
	options := make(map[string]Option)
	for k, v := range df.RequiredOptions {
		options[k] = Option{Values: []string{v}}
	}
	return options
}
