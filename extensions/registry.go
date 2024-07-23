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

type Registry interface {
	Load(reader io.Reader) error
	LocalLookup(funcName string, argTypes []string) (ID, FuncKind, bool)
	CanonicalLookup(id ID, funcKind FuncKind) (*LocalFunctionVariant, bool)
	BinaryFunctions() []ID
}

func NewRegistry() Registry {
	return &registryImpl{dialect: Dialect{}}
}

type registryImpl struct {
	dialect Dialect
}

func (r *registryImpl) Load(reader io.Reader) error {
	err := r.dialect.Load(reader)
	if err != nil {
		return err
	}

	return nil
}

func (r *registryImpl) LocalLookup(funcName string, argTypes []string) (ID, FuncKind, bool) {
	localId := LocalFunctionID{Name: funcName, Signature: r.dialect.getSignature(argTypes)}
	lf, ok := r.dialect.functions[localId]
	if !ok {
		return ID{}, "", false
	}
	return lf.ID, lf.FuncKind, true
}

func (r *registryImpl) CanonicalLookup(id ID, funcKind FuncKind) (*LocalFunctionVariant, bool) {
	var (
		lf LocalFunctionVariant
		ok bool
	)
	switch funcKind {
	case ScalarFunc:
		lf, ok = r.dialect.localScalarFunctions[id]
	case AggregateFunc:
		lf, ok = r.dialect.localAggregateFunctions[id]
	case WindowFunc:
		lf, ok = r.dialect.localWindowFunctions[id]
	}

	if !ok {
		return nil, false
	}
	return &lf, true
}

func (r *registryImpl) BinaryFunctions() []ID {
	return r.dialect.binaryFunctions
}

type Dialect struct {
	file DialectFile

	typeMap map[string]string // local to substrait type

	binaryFunctions []ID

	localScalarFunctions    map[ID]LocalFunctionVariant
	localAggregateFunctions map[ID]LocalFunctionVariant
	localWindowFunctions    map[ID]LocalFunctionVariant

	functions map[LocalFunctionID]LocalFunctionVariant
}

type LocalFunctionID struct {
	Name      string
	Signature string
}

type LocalFunctionVariant struct {
	ID            ID
	FuncKind      FuncKind
	Name          string
	Signature     string
	LocalName     string
	ArgLocalTypes []string
	Infix         bool
}

func (d *Dialect) Load(reader io.Reader) error {
	dec := yaml.NewDecoder(reader)
	if err := dec.Decode(&d.file); err != nil {
		return err
	}

	d.buildTypeMap()
	d.localScalarFunctions = d.buildLocalFunctionMap(d.file.ScalarFunctions, ScalarFunc)
	d.localAggregateFunctions = d.buildLocalFunctionMap(d.file.AggregateFunctions, AggregateFunc)
	d.localWindowFunctions = d.buildLocalFunctionMap(d.file.WindowFunctions, WindowFunc)

	d.functions = make(map[LocalFunctionID]LocalFunctionVariant)
	d.updateFunctionMap(d.localScalarFunctions)
	d.updateFunctionMap(d.localAggregateFunctions)
	d.updateFunctionMap(d.localWindowFunctions)

	return nil
}

func (d *Dialect) buildLocalFunctionMap(functions []DialectFunction, funcKind FuncKind) map[ID]LocalFunctionVariant {
	funcMap := make(map[ID]LocalFunctionVariant)
	for _, f := range functions {
		uri, name := d.file.getUriAndFunctionName(&f)
		for _, kernel := range f.SupportedKernels {
			localName := f.LocalName
			if len(localName) == 0 {
				localName = name
			}
			lf := LocalFunctionVariant{
				ID:            ID{uri, name + ":" + kernel},
				FuncKind:      funcKind,
				Name:          name,
				LocalName:     localName,
				Infix:         f.Infix,
				Signature:     kernel,
				ArgLocalTypes: d.getArgLocalTypes(kernel),
			}
			funcMap[lf.ID] = lf
			if funcKind == ScalarFunc && len(lf.ArgLocalTypes) == 2 {
				d.binaryFunctions = append(d.binaryFunctions, lf.ID)
			}
		}
	}
	return funcMap
}

func (d *Dialect) updateFunctionMap(localFuncMap map[ID]LocalFunctionVariant) {
	for _, lf := range localFuncMap {
		name := lf.LocalName
		d.functions[LocalFunctionID{name, lf.Signature}] = lf
	}
}

func (d *Dialect) buildTypeMap() {
	d.typeMap = make(map[string]string)
	for k, v := range d.file.SupportedTypes {
		d.typeMap[v] = k
	}
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

func (d *Dialect) getSignature(argLocalTypes []string) string {
	if len(argLocalTypes) == 0 {
		return ""
	}

	var sb strings.Builder
	for i, p := range argLocalTypes {
		if i != 0 {
			sb.WriteByte('_')
		}
		sb.WriteString(d.typeMap[p])
	}
	return sb.String()
}

func (d *Dialect) getArgLocalTypes(kernel string) []string {
	parts := strings.Split(kernel, "_")
	argLocalTypes := make([]string, len(parts))
	for i, p := range parts {
		argLocalTypes[i] = d.typeMap[p]
	}
	return argLocalTypes
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
