// SPDX-License-Identifier: Apache-2.0

package extensions

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"path"
	"regexp"
	"sort"
	"sync"

	"github.com/creasty/defaults"
	"github.com/goccy/go-yaml"
	"github.com/substrait-io/substrait"
	substraitgo "github.com/substrait-io/substrait-go/v8"
	"github.com/substrait-io/substrait-protobuf/go/substraitpb/extensions"
)

type AdvancedExtension = extensions.AdvancedExtension

const SubstraitDefaultURNPrefix = "extension:io.substrait:"

var (
	getDefaultCollectionOnce = sync.OnceValues(loadDefaultCollection)
	unsupportedExtensions    = map[string]bool{
		"unknown.yaml": true,
	}
)

var (
	urnPattern             = regexp.MustCompile(`^extension:[^:]+:[^:]+$`)
	dependencyAliasPattern = regexp.MustCompile(`^[A-Za-z_$][A-Za-z0-9_$]*$`)
)

// validateURN validates that a URN follows the format "extension:<owner>:<id>"
func validateUrn(urn string) bool {
	return urnPattern.MatchString(urn)
}

func validateDependencyAlias(alias string) bool {
	return dependencyAliasPattern.MatchString(alias)
}

// GetDefaultCollectionWithNoError returns a Collection that is loaded with the default Substrait extension definitions.
// This version is provided for the ease of use of legacy code. Please use GetDefaultCollection instead.
func GetDefaultCollectionWithNoError() *Collection {
	c, err := GetDefaultCollection()
	if err != nil {
		panic(err)
	}
	return c
}

// GetDefaultCollection returns a Collection that is loaded with the default Substrait extension definitions.
func GetDefaultCollection() (*Collection, error) {
	return getDefaultCollectionOnce()
}

func loadDefaultCollection() (*Collection, error) {
	substraitFS := substrait.GetSubstraitExtensionsFS()
	entries, err := substraitFS.ReadDir("extensions")
	if err != nil {
		return nil, err
	}

	var defaultCollection Collection
	for _, ent := range entries {
		err2 := loadExtensionFile(&defaultCollection, substraitFS, ent)
		if err2 != nil {
			return nil, err2
		}
	}
	return &defaultCollection, nil
}

func loadExtensionFile(collection *Collection, substraitFS embed.FS, ent fs.DirEntry) error {
	f, err := substraitFS.Open(path.Join("extensions/", ent.Name()))
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()
	fileStat, err := f.Stat()
	if err != nil {
		return err
	}
	fileName := path.Base(fileStat.Name())
	if _, ok := unsupportedExtensions[fileName]; !ok {
		err = collection.Load(f)
		if err != nil {
			return err
		}
	}
	return nil
}

type Collection struct {
	urnSet map[string]struct{}

	scalarMap        map[FunctionID]*ScalarFunctionVariant
	aggregateMap     map[FunctionID]*AggregateFunctionVariant
	windowMap        map[FunctionID]*WindowFunctionVariant
	typeMap          map[TypeID]Type
	typeVariationMap map[TypeVariationID]TypeVariation
	fileMetadataMap  map[string]map[string]any // keyed by URN
}

func (c *Collection) GetType(id TypeID) (t Type, ok bool) {
	t, ok = c.typeMap[id]
	return
}

// GetFileMetadata returns the top-level metadata from the extension file
// identified by its URN. Returns nil if no metadata was provided or the URN is not loaded.
func (c *Collection) GetFileMetadata(urn string) map[string]any {
	return c.fileMetadataMap[urn]
}

func (c *Collection) GetTypeVariation(id TypeVariationID) (tv TypeVariation, ok bool) {
	tv, ok = c.typeVariationMap[id]
	return
}

var void = struct{}{}

type variants interface {
	*ScalarFunctionVariant | *AggregateFunctionVariant | *WindowFunctionVariant
	Name() string
	Signature() string
}

type extFn[T variants] interface {
	GetVariants(urn string) []T
}

func addToMaps[T variants](id FunctionID, fn extFn[T], m map[FunctionID]T) {
	for _, v := range fn.GetVariants(id.URN) {
		id.Signature = v.Signature()
		m[id] = v
	}
}

func (c *Collection) GetScalarFunc(id FunctionID) (*ScalarFunctionVariant, bool) {
	fn, ok := c.scalarMap[id]
	return fn, ok
}

func (c *Collection) GetAggregateFunc(id FunctionID) (*AggregateFunctionVariant, bool) {
	fn, ok := c.aggregateMap[id]
	return fn, ok
}

func (c *Collection) GetWindowFunc(id FunctionID) (*WindowFunctionVariant, bool) {
	fn, ok := c.windowMap[id]
	return fn, ok
}

// IsRegisteredFunction reports whether id resolves to a registered function variant.
func (c *Collection) IsRegisteredFunction(id FunctionID) bool {
	if _, ok := c.scalarMap[id]; ok {
		return true
	}
	if _, ok := c.aggregateMap[id]; ok {
		return true
	}
	_, ok := c.windowMap[id]
	return ok
}

func (c *Collection) init() {
	if c.urnSet == nil {
		c.urnSet = make(map[string]struct{})
		c.scalarMap = make(map[FunctionID]*ScalarFunctionVariant)
		c.aggregateMap = make(map[FunctionID]*AggregateFunctionVariant)
		c.windowMap = make(map[FunctionID]*WindowFunctionVariant)
		c.typeMap = make(map[TypeID]Type)
		c.typeVariationMap = make(map[TypeVariationID]TypeVariation)
		c.fileMetadataMap = make(map[string]map[string]any)
	}
}

func (c *Collection) Load(r io.Reader) error {
	c.init()

	var file SimpleExtensionFile
	dec := yaml.NewDecoder(r)
	if err := dec.Decode(&file); err != nil {
		return err
	}

	urn := file.Urn
	if urn == "" {
		return fmt.Errorf("%w: missing urn", substraitgo.ErrInvalidSimpleExtention)
	}
	if !validateUrn(urn) {
		return fmt.Errorf("%w: invalid urn, expected format is \"extension:<owner>:<id>\", got: %s", substraitgo.ErrInvalidSimpleExtention, urn)
	}
	if c.URNLoaded(urn) {
		return fmt.Errorf("%w:  urn %s already loaded", substraitgo.ErrKeyExists, urn)
	}

	for alias, dependencyURN := range file.Dependencies {
		if !validateDependencyAlias(alias) {
			return fmt.Errorf("%w: invalid dependency alias %q", substraitgo.ErrInvalidSimpleExtention, alias)
		}
		if !c.URNLoaded(dependencyURN) {
			return fmt.Errorf("%w: dependency urn %q for alias %q is not loaded", substraitgo.ErrNotFound, dependencyURN, alias)
		}
	}

	if err := file.validateUserDefinedTypeReferences(); err != nil {
		return err
	}

	c.urnSet[urn] = void

	if file.Metadata != nil {
		c.fileMetadataMap[urn] = file.Metadata
	}

	typeID := TypeID{URN: urn}
	for _, t := range file.Types {
		typeID.Name = t.Name
		c.typeMap[typeID] = t
	}

	typeVariationID := TypeVariationID{URN: urn}
	for _, t := range file.TypeVariations {
		typeVariationID.Name = t.Name
		c.typeVariationMap[typeVariationID] = t
	}

	id := FunctionID{URN: urn}

	for _, f := range file.ScalarFunctions {
		if err := defaults.Set(&f); err != nil {
			return fmt.Errorf("failure setting defaults for scalar functions: %w", err)
		}
		addToMaps[*ScalarFunctionVariant](id, &f, c.scalarMap)
	}

	for _, f := range file.AggregateFunctions {
		if err := defaults.Set(&f); err != nil {
			return fmt.Errorf("failure setting defaults for aggregate functions: %w", err)
		}
		addToMaps[*AggregateFunctionVariant](id, &f, c.aggregateMap)
	}

	for _, f := range file.WindowFunctions {
		if err := defaults.Set(&f); err != nil {
			return fmt.Errorf("failure setting defaults for window functions: %w", err)
		}
		addToMaps[*WindowFunctionVariant](id, &f, c.windowMap)
	}

	// Aggregate functions can be used as Window Functions
	for _, f := range file.AggregateFunctions {
		// Convert each aggregate implementation to a window implementation
		windowImpls := make([]WindowFunctionImpl, len(f.Impls))
		for i, aggImpl := range f.Impls {
			windowImpls[i] = WindowFunctionImpl{
				AggregateFunctionImpl: aggImpl,
				WindowType:            StreamingWindow, // Set window type to STREAMING
			}
		}

		wf := WindowFunction{
			Name:        f.Name,
			Description: f.Description,
			Impls:       windowImpls,
			Metadata:    f.Metadata,
		}
		if err := defaults.Set(&wf); err != nil {
			return fmt.Errorf("failure setting defaults for window functions: %w", err)
		}
		addToMaps[*WindowFunctionVariant](id, &wf, c.windowMap)
	}

	return nil
}

func (c *Collection) URNLoaded(urn string) bool {
	_, ok := c.urnSet[urn]
	return ok
}

func (c *Collection) GetAllScalarFunctions() []*ScalarFunctionVariant {
	return getValues(c.scalarMap)
}

func (c *Collection) GetAllAggregateFunctions() []*AggregateFunctionVariant {
	return getValues(c.aggregateMap)
}

func (c *Collection) GetAllWindowFunctions() []*WindowFunctionVariant {
	return getValues(c.windowMap)
}

func getValues[M ~map[K]V, K comparable, V any](m M) []V {
	result := make([]V, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

type Set interface {
	DecodeTypeVariation(anchor uint32) (TypeVariationID, bool)
	DecodeFunc(anchor uint32) (FunctionID, bool)
	DecodeType(anchor uint32) (TypeID, bool)
	LookupTypeVariation(anchor uint32, c *Collection) (TypeVariation, bool)
	LookupType(anchor uint32, c *Collection) (Type, bool)
	LookupScalarFunction(anchor uint32, c *Collection) (*ScalarFunctionVariant, bool)
	LookupAggregateFunction(anchor uint32, c *Collection) (*AggregateFunctionVariant, bool)
	LookupWindowFunction(anchor uint32, c *Collection) (*WindowFunctionVariant, bool)

	FindURN(urn string) (uint32, bool)
	GetTypeAnchor(id TypeID) uint32
	GetFuncAnchor(id FunctionID) uint32
	GetTypeVariationAnchor(id TypeVariationID) uint32

	ToProto(c *Collection) ([]*extensions.SimpleExtensionURN, []*extensions.SimpleExtensionDeclaration)
}

func NewSet() Set {
	return &set{
		urns:             make(map[uint32]string),
		funcMap:          make(map[uint32]FunctionID),
		funcs:            make(map[FunctionID]uint32),
		types:            make(map[TypeID]uint32),
		typesMap:         make(map[uint32]TypeID),
		typeVariationMap: make(map[uint32]TypeVariationID),
		typeVariations:   make(map[TypeVariationID]uint32),
	}
}

type set struct {
	urns map[uint32]string

	typesMap map[uint32]TypeID
	types    map[TypeID]uint32

	typeVariationMap map[uint32]TypeVariationID
	typeVariations   map[TypeVariationID]uint32

	funcMap map[uint32]FunctionID
	funcs   map[FunctionID]uint32
}

func (e *set) ToProto(c *Collection) ([]*extensions.SimpleExtensionURN, []*extensions.SimpleExtensionDeclaration) {
	urnBackRef := make(map[string]uint32)

	urns := make([]*extensions.SimpleExtensionURN, 0, len(e.urns))
	for anchor, urn := range e.urns {
		urnBackRef[urn] = anchor
		urns = append(urns, &extensions.SimpleExtensionURN{
			ExtensionUrnAnchor: anchor,
			Urn:                urn,
		})
	}

	// Sort URN extensions by the anchor for consistent output
	sort.Slice(urns, func(i, j int) bool { return urns[i].ExtensionUrnAnchor < urns[j].ExtensionUrnAnchor })

	decls := make([]*extensions.SimpleExtensionDeclaration, 0, len(e.types)+len(e.typeVariations)+len(e.funcs))
	for id, anchor := range e.types {
		decls = append(decls, &extensions.SimpleExtensionDeclaration{
			MappingType: &extensions.SimpleExtensionDeclaration_ExtensionType_{
				ExtensionType: &extensions.SimpleExtensionDeclaration_ExtensionType{
					ExtensionUrnReference: urnBackRef[id.URN],
					TypeAnchor:            anchor,
					Name:                  id.Name,
				},
			},
		})
	}

	sort.Slice(decls, func(i, j int) bool {
		return decls[i].GetExtensionType().TypeAnchor < decls[j].GetExtensionType().TypeAnchor
	})
	typesCount := len(decls)

	for id, anchor := range e.typeVariations {
		decls = append(decls, &extensions.SimpleExtensionDeclaration{
			MappingType: &extensions.SimpleExtensionDeclaration_ExtensionTypeVariation_{
				ExtensionTypeVariation: &extensions.SimpleExtensionDeclaration_ExtensionTypeVariation{
					ExtensionUrnReference: urnBackRef[id.URN],
					TypeVariationAnchor:   anchor,
					Name:                  id.Name,
				},
			},
		})
	}

	typeDecls := decls[typesCount:]
	sort.Slice(typeDecls, func(i, j int) bool {
		return decls[i].GetExtensionTypeVariation().TypeVariationAnchor < decls[j].GetExtensionTypeVariation().TypeVariationAnchor
	})

	typeVarCount := len(decls)
	for id, anchor := range e.funcs {
		decls = append(decls, &extensions.SimpleExtensionDeclaration{
			MappingType: &extensions.SimpleExtensionDeclaration_ExtensionFunction_{
				ExtensionFunction: &extensions.SimpleExtensionDeclaration_ExtensionFunction{
					ExtensionUrnReference: urnBackRef[id.URN],
					FunctionAnchor:        anchor,
					Name:                  id.Signature,
				},
			},
		})
	}

	typeVarDecls := decls[typeVarCount:]
	sort.Slice(typeVarDecls, func(i, j int) bool {
		return decls[i].GetExtensionFunction().GetFunctionAnchor() < decls[j].GetExtensionFunction().GetFunctionAnchor()
	})

	return urns, decls
}

func (e *set) LookupWindowFunction(anchor uint32, c *Collection) (sv *WindowFunctionVariant, ok bool) {
	id, ok := e.funcMap[anchor]
	if !ok {
		return nil, false
	}

	return c.GetWindowFunc(id)
}

func (e *set) LookupAggregateFunction(anchor uint32, c *Collection) (sv *AggregateFunctionVariant, ok bool) {
	id, ok := e.funcMap[anchor]
	if !ok {
		return nil, false
	}

	return c.GetAggregateFunc(id)
}

func (e *set) LookupScalarFunction(anchor uint32, c *Collection) (sv *ScalarFunctionVariant, ok bool) {
	id, ok := e.funcMap[anchor]
	if !ok {
		return nil, false
	}

	return c.GetScalarFunc(id)
}

func (e *set) LookupType(anchor uint32, c *Collection) (tv Type, ok bool) {
	id, ok := e.typesMap[anchor]
	if !ok {
		return tv, false
	}

	tv, ok = c.typeMap[id]
	return
}

func (e *set) LookupTypeVariation(anchor uint32, c *Collection) (tv TypeVariation, ok bool) {
	id, ok := e.typeVariationMap[anchor]
	if !ok {
		return tv, false
	}

	tv, ok = c.typeVariationMap[id]
	return
}

func (e *set) DecodeTypeVariation(anchor uint32) (id TypeVariationID, ok bool) {
	id, ok = e.typeVariationMap[anchor]
	return
}

func (e *set) DecodeFunc(anchor uint32) (id FunctionID, ok bool) {
	id, ok = e.funcMap[anchor]
	return
}

func (e *set) DecodeType(anchor uint32) (id TypeID, ok bool) {
	id, ok = e.typesMap[anchor]
	return
}

func (e *set) GetTypeAnchor(id TypeID) uint32 {
	a, ok := e.types[id]
	if !ok {
		_, err := e.addOrGetURN(id.URN)
		if err != nil {
			panic(err)
		}
		a = uint32(len(e.types)) + 1
		e.encodeType(a, id)
	}
	return a
}

func (e *set) GetFuncAnchor(id FunctionID) uint32 {
	a, ok := e.funcs[id]
	if !ok {
		_, err := e.addOrGetURN(id.URN)
		if err != nil {
			panic(err)
		}
		a = uint32(len(e.funcs)) + 1
		e.encodeFunc(a, id)
	}
	return a
}

func (e *set) GetTypeVariationAnchor(id TypeVariationID) uint32 {
	a, ok := e.typeVariations[id]
	if !ok {
		_, err := e.addOrGetURN(id.URN)
		if err != nil {
			panic(err)
		}
		// add 1 to the length to avoid an anchor of 0
		// so that it's easier to tell when there is no
		// type variation.
		a = uint32(len(e.typeVariations)) + 1
		e.encodeTypeVariation(a, id)
	}
	return a
}

func (e *set) encodeType(anchor uint32, id TypeID) {
	e.typesMap[anchor] = id
	e.types[id] = anchor
}

func (e *set) encodeTypeVariation(anchor uint32, id TypeVariationID) {
	e.typeVariationMap[anchor] = id
	e.typeVariations[id] = anchor
}

func (e *set) encodeFunc(anchor uint32, id FunctionID) {
	e.funcMap[anchor] = id
	e.funcs[id] = anchor
}

func (e *set) FindURN(urn string) (uint32, bool) {
	for k, v := range e.urns {
		if v == urn {
			return k, true
		}
	}
	return 0, false
}

func (e *set) addOrGetURN(urn string) (uint32, error) {
	for k, v := range e.urns {
		if v == urn {
			return k, nil
		}
	}
	sz := uint32(len(e.urns)) + 1
	if _, ok := e.urns[sz]; ok {
		return 0, fmt.Errorf("%w: URN anchor %d already exists", substraitgo.ErrKeyExists, sz)
	}

	e.urns[sz] = urn
	return sz, nil
}

type TopLevel interface {
	GetExtensionUrns() []*extensions.SimpleExtensionURN
	GetExtensions() []*extensions.SimpleExtensionDeclaration
}

func GetExtensionSet(plan TopLevel, c *Collection) (Set, error) {
	urns := make(map[uint32]string)
	for _, urn := range plan.GetExtensionUrns() {
		urns[urn.ExtensionUrnAnchor] = urn.Urn
	}

	ret := &set{
		urns:             urns,
		funcMap:          make(map[uint32]FunctionID),
		funcs:            make(map[FunctionID]uint32),
		typesMap:         make(map[uint32]TypeID),
		types:            make(map[TypeID]uint32),
		typeVariationMap: make(map[uint32]TypeVariationID),
		typeVariations:   make(map[TypeVariationID]uint32),
	}

	resolveRefToURN := func(urnRef uint32) (string, error) {
		urn, urnOk := urns[urnRef]
		if !urnOk {
			return "", fmt.Errorf("unable to resolve extension reference: URN reference %d could not be resolved", urnRef)
		}
		// Validate that the URN exists in the Collection
		if !c.URNLoaded(urn) {
			return "", fmt.Errorf("%w: URN '%s' not found in extension collection", substraitgo.ErrNotFound, urn)
		}
		return urn, nil
	}

	for _, ext := range plan.GetExtensions() {
		switch e := ext.MappingType.(type) {
		case *extensions.SimpleExtensionDeclaration_ExtensionTypeVariation_:
			etv := e.ExtensionTypeVariation
			urn, err := resolveRefToURN(etv.ExtensionUrnReference)
			if err != nil {
				return nil, err
			}
			ret.encodeTypeVariation(etv.TypeVariationAnchor, TypeVariationID{
				URN:  urn,
				Name: etv.Name,
			})
		case *extensions.SimpleExtensionDeclaration_ExtensionType_:
			et := e.ExtensionType
			urn, err := resolveRefToURN(et.ExtensionUrnReference)
			if err != nil {
				return nil, err
			}
			ret.encodeType(et.TypeAnchor, TypeID{
				URN:  urn,
				Name: et.Name,
			})
		case *extensions.SimpleExtensionDeclaration_ExtensionFunction_:
			ef := e.ExtensionFunction
			urn, err := resolveRefToURN(ef.ExtensionUrnReference)
			if err != nil {
				return nil, err
			}
			ret.encodeFunc(ef.FunctionAnchor, FunctionID{
				URN:       urn,
				Signature: ef.Name,
			})
		}
	}

	return ret, nil
}
