// SPDX-License-Identifier: Apache-2.0

package extensions

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"path"
	"sort"
	"sync"

	"github.com/creasty/defaults"
	"github.com/goccy/go-yaml"
	"github.com/substrait-io/substrait"
	substraitgo "github.com/substrait-io/substrait-go/v6"
	"github.com/substrait-io/substrait-protobuf/go/substraitpb/extensions"
)

type AdvancedExtension = extensions.AdvancedExtension

const SubstraitDefaultURIPrefix = "https://github.com/substrait-io/substrait/blob/main/extensions/"

var (
	getDefaultCollectionOnce = sync.OnceValues[*Collection, error](loadDefaultCollection)
	unsupportedExtensions    = map[string]bool{
		"unknown.yaml": true,
	}
)

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
		err = collection.Load(SubstraitDefaultURIPrefix+ent.Name(), f)
		if err != nil {
			return err
		}
	}
	return nil
}

// ID is the unique identifier for a substrait object
type ID struct {
	//Deprecated, will eventually switch to only URN
	URI string
	URN string
	// Name of the object. For functions, a simple name may be used for lookups,
	// but as a unique identifier the compound name will be used
	Name string
}


// This is a temporary internal type which will be removed once the
// migration to urns is complete
type uriUrnTranslator struct {
	uriToUrn map[string]string
	urnToUri map[string]string
}

func newUriUrnTranslator() *uriUrnTranslator {
	return &uriUrnTranslator{
		uriToUrn: make(map[string]string),
		urnToUri: make(map[string]string),
	}
}

func (u *uriUrnTranslator) toUrn(uri string) (string, bool) {
	urn, ok := u.uriToUrn[uri]
	return urn, ok
}

func (u *uriUrnTranslator) toUri(urn string) (string, bool) {
	uri, ok := u.urnToUri[urn]
	return uri, ok
}

func (u *uriUrnTranslator) addMapping(uri, urn string) error {
	if uri == "" {
		return fmt.Errorf("URI cannot be empty")
	}
	if urn == "" {
		return fmt.Errorf("URN cannot be empty")
	}

	if _, exists := u.uriToUrn[uri]; exists {
		return fmt.Errorf("URI '%s' already exists in translator", uri)
	}

	if _, exists := u.urnToUri[urn]; exists {
		return fmt.Errorf("URN '%s' already exists in translator", urn)
	}

	u.uriToUrn[uri] = urn
	u.urnToUri[urn] = uri

	return nil
}

// canonicalizeID takes an ID with partial data, and returns the fullest
// representation
// E.g. if only uri is present in the argument but urn is also available,
// it will add the urn info
// This is a temporary workaround during the uri -> urn migration
func (u *uriUrnTranslator) canonicalizeID(id ID) ID {
	canonical := id

	if canonical.URN == "" && canonical.URI != "" {
		if urn, ok := u.toUrn(canonical.URI); ok {
			canonical.URN = urn
		}
	}

	if canonical.URI == "" && canonical.URN != "" {
		if uri, ok := u.toUri(canonical.URN); ok {
			canonical.URI = uri
		}
	}

	return canonical
}

type Collection struct {
	uriSet           map[string]struct{}
	uriUrnTranslator *uriUrnTranslator

	simpleNameMap map[ID]string

	scalarMap        map[ID]*ScalarFunctionVariant
	aggregateMap     map[ID]*AggregateFunctionVariant
	windowMap        map[ID]*WindowFunctionVariant
	typeMap          map[ID]Type
	typeVariationMap map[ID]TypeVariation
}

func (c *Collection) GetType(id ID) (t Type, ok bool) {
	canonical := c.uriUrnTranslator.canonicalizeID(id)
	t, ok = c.typeMap[canonical]
	return
}

func (c *Collection) GetTypeVariation(id ID) (tv TypeVariation, ok bool) {
	canonical := c.uriUrnTranslator.canonicalizeID(id)
	tv, ok = c.typeVariationMap[canonical]
	return
}

var void = struct{}{}

type variants interface {
	*ScalarFunctionVariant | *AggregateFunctionVariant | *WindowFunctionVariant
	Name() string
	CompoundName() string
	ID() ID
}

func checkMaps[T variants](id ID, m map[ID]T, simpleNames map[ID]string, translator *uriUrnTranslator) (T, bool) {
	canonical := translator.canonicalizeID(id)

	if sv, ok := m[canonical]; ok {
		return sv, true
	}

	if compound, ok := simpleNames[canonical]; ok {
		canonical.Name = compound
		if sv, ok := m[canonical]; ok {
			return sv, true
		}
	}

	return nil, false
}

type extFn[T variants] interface {
	GetVariants() []T
}

func addToMaps[T variants](fn extFn[T], m map[ID]T, simpleMap map[string]string, translator *uriUrnTranslator) {
	variants := fn.GetVariants()
	for _, v := range variants {
		// Each variant should already have its ID set correctly
		canonical := translator.canonicalizeID(v.ID())
		m[canonical] = v
	}

	if len(variants) == 1 {
		v := variants[0]
		if _, ok := simpleMap[v.Name()]; ok {
			delete(simpleMap, v.Name())
		} else {
			simpleMap[v.Name()] = v.CompoundName()
		}
	}
}

func (c *Collection) GetScalarFunc(id ID) (*ScalarFunctionVariant, bool) {
	return checkMaps(id, c.scalarMap, c.simpleNameMap, c.uriUrnTranslator)
}

func (c *Collection) GetAggregateFunc(id ID) (*AggregateFunctionVariant, bool) {
	return checkMaps(id, c.aggregateMap, c.simpleNameMap, c.uriUrnTranslator)
}

func (c *Collection) GetWindowFunc(id ID) (*WindowFunctionVariant, bool) {
	return checkMaps(id, c.windowMap, c.simpleNameMap, c.uriUrnTranslator)
}

func (c *Collection) init() {
	if c.uriSet == nil {
		c.uriSet = make(map[string]struct{})
		c.uriUrnTranslator = newUriUrnTranslator()
		c.simpleNameMap = make(map[ID]string)
		c.scalarMap = make(map[ID]*ScalarFunctionVariant)
		c.aggregateMap = make(map[ID]*AggregateFunctionVariant)
		c.windowMap = make(map[ID]*WindowFunctionVariant)
		c.typeMap = make(map[ID]Type)
		c.typeVariationMap = make(map[ID]TypeVariation)
	}
}

func (c *Collection) Load(uri string, r io.Reader) error {
	c.init()

	if c.URILoaded(uri) {
		return fmt.Errorf("%w: uri '%s' already loaded",
			substraitgo.ErrKeyExists, uri)
	}

	c.uriSet[uri] = void

	var file SimpleExtensionFile
	dec := yaml.NewDecoder(r)
	if err := dec.Decode(&file); err != nil {
		return err
	}

	if file.URN != "" {
		if err := c.uriUrnTranslator.addMapping(uri, file.URN); err != nil {
			return fmt.Errorf("failed to add URI/URN mapping: %w", err)
		}
	}

	id := ID{URI: uri, URN: file.URN}
	for _, t := range file.Types {
		id.Name = t.Name
		canonical := c.uriUrnTranslator.canonicalizeID(id)
		c.typeMap[canonical] = t
	}

	for _, t := range file.TypeVariations {
		id.Name = t.Name
		canonical := c.uriUrnTranslator.canonicalizeID(id)
		c.typeVariationMap[canonical] = t
	}

	// if there is only one implementation for a given function
	// it should be findable by its simple name in addition to the
	// compound name.
	simpleNames := make(map[string]string)

	for _, f := range file.ScalarFunctions {
		if err := defaults.Set(&f); err != nil {
			return fmt.Errorf("failure setting defaults for scalar functions: %w", err)
		}
		f.id = id // Set the ID on the function so it can create variants with the full ID
		addToMaps[*ScalarFunctionVariant](&f, c.scalarMap, simpleNames, c.uriUrnTranslator)
	}

	for _, f := range file.AggregateFunctions {
		if err := defaults.Set(&f); err != nil {
			return fmt.Errorf("failure setting defaults for aggregate functions: %w", err)
		}
		f.id = id // Set the ID on the function so it can create variants with the full ID
		addToMaps[*AggregateFunctionVariant](&f, c.aggregateMap, simpleNames, c.uriUrnTranslator)
	}

	for _, f := range file.WindowFunctions {
		if err := defaults.Set(&f); err != nil {
			return fmt.Errorf("failure setting defaults for window functions: %w", err)
		}
		f.id = id // Set the ID on the function so it can create variants with the full ID
		addToMaps[*WindowFunctionVariant](&f, c.windowMap, simpleNames, c.uriUrnTranslator)
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
			id:          id, // Set the ID on the window function
		}
		if err := defaults.Set(&wf); err != nil {
			return fmt.Errorf("failure setting defaults for window functions: %w", err)
		}
		addToMaps[*WindowFunctionVariant](&wf, c.windowMap, simpleNames, c.uriUrnTranslator)
	}

	// add simple name aliases
	for k, v := range simpleNames {
		id.Name = k
		canonical := c.uriUrnTranslator.canonicalizeID(id)
		c.simpleNameMap[canonical] = v
	}

	return nil
}

func (c *Collection) URILoaded(uri string) bool {
	_, ok := c.uriSet[uri]
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
	DecodeTypeVariation(anchor uint32) (ID, bool)
	DecodeFunc(anchor uint32) (ID, bool)
	DecodeType(anchor uint32) (ID, bool)
	LookupTypeVariation(anchor uint32, c *Collection) (TypeVariation, bool)
	LookupType(anchor uint32, c *Collection) (Type, bool)
	LookupScalarFunction(anchor uint32, c *Collection) (*ScalarFunctionVariant, bool)
	LookupAggregateFunction(anchor uint32, c *Collection) (*AggregateFunctionVariant, bool)
	LookupWindowFunction(anchor uint32, c *Collection) (*WindowFunctionVariant, bool)

	FindURI(uri string) (uint32, bool)
	GetTypeAnchor(id ID) uint32
	GetFuncAnchor(id ID) uint32
	GetTypeVariationAnchor(id ID) uint32

	ToProto() ([]*extensions.SimpleExtensionURI, []*extensions.SimpleExtensionDeclaration)
}

func NewSet() Set {
	return &set{
		uris:             make(map[uint32]string),
		funcMap:          make(map[uint32]ID),
		funcs:            make(map[ID]uint32),
		types:            make(map[ID]uint32),
		typesMap:         make(map[uint32]ID),
		typeVariationMap: make(map[uint32]ID),
		typeVariations:   make(map[ID]uint32),
	}
}

type set struct {
	uris map[uint32]string

	typesMap map[uint32]ID
	types    map[ID]uint32

	typeVariationMap map[uint32]ID
	typeVariations   map[ID]uint32

	funcMap map[uint32]ID
	funcs   map[ID]uint32
}

func (e *set) ToProto() ([]*extensions.SimpleExtensionURI, []*extensions.SimpleExtensionDeclaration) {
	backRef := make(map[string]uint32)
	uris := make([]*extensions.SimpleExtensionURI, 0, len(e.uris))
	for anchor, uri := range e.uris {
		backRef[uri] = anchor
		uris = append(uris, &extensions.SimpleExtensionURI{
			ExtensionUriAnchor: anchor,
			Uri:                uri,
		})
	}

	// Sort extensions by the anchor for consistent output
	sort.Slice(uris, func(i, j int) bool { return uris[i].ExtensionUriAnchor < uris[j].ExtensionUriAnchor })

	decls := make([]*extensions.SimpleExtensionDeclaration, 0, len(e.types)+len(e.typeVariations)+len(e.funcs))
	for id, anchor := range e.types {
		decls = append(decls, &extensions.SimpleExtensionDeclaration{
			MappingType: &extensions.SimpleExtensionDeclaration_ExtensionType_{
				ExtensionType: &extensions.SimpleExtensionDeclaration_ExtensionType{
					ExtensionUriReference: backRef[id.URI],
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
					ExtensionUriReference: backRef[id.URI],
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
					ExtensionUriReference: backRef[id.URI],
					FunctionAnchor:        anchor,
					Name:                  id.Name,
				},
			},
		})
	}

	typeVarDecls := decls[typeVarCount:]
	sort.Slice(typeVarDecls, func(i, j int) bool {
		return decls[i].GetExtensionFunction().GetFunctionAnchor() < decls[j].GetExtensionFunction().GetFunctionAnchor()
	})

	return uris, decls
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

func (e *set) DecodeTypeVariation(anchor uint32) (id ID, ok bool) {
	id, ok = e.typeVariationMap[anchor]
	return
}

func (e *set) DecodeFunc(anchor uint32) (id ID, ok bool) {
	id, ok = e.funcMap[anchor]
	return
}

func (e *set) DecodeType(anchor uint32) (id ID, ok bool) {
	id, ok = e.typesMap[anchor]
	return
}

func (e *set) GetTypeAnchor(id ID) uint32 {
	a, ok := e.types[id]
	if !ok {
		_, err := e.addOrGetURI(id.URI)
		if err != nil {
			panic(err)
		}
		a = uint32(len(e.types)) + 1
		e.encodeType(a, id)
	}
	return a
}

func (e *set) GetFuncAnchor(id ID) uint32 {
	a, ok := e.funcs[id]
	if !ok {
		_, err := e.addOrGetURI(id.URI)
		if err != nil {
			panic(err)
		}
		a = uint32(len(e.funcs)) + 1
		e.encodeFunc(a, id)
	}
	return a
}

func (e *set) GetTypeVariationAnchor(id ID) uint32 {
	a, ok := e.typeVariations[id]
	if !ok {
		_, err := e.addOrGetURI(id.URI)
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

func (e *set) encodeType(anchor uint32, id ID) {
	e.typesMap[anchor] = id
	e.types[id] = anchor
}

func (e *set) encodeTypeVariation(anchor uint32, id ID) {
	e.typeVariationMap[anchor] = id
	e.typeVariations[id] = anchor
}

func (e *set) encodeFunc(anchor uint32, id ID) {
	e.funcMap[anchor] = id
	e.funcs[id] = anchor
}

func (e *set) FindURI(uri string) (uint32, bool) {
	for k, v := range e.uris {
		if v == uri {
			return k, true
		}
	}
	return 0, false
}

func (e *set) addOrGetURI(uri string) (uint32, error) {
	for k, v := range e.uris {
		if v == uri {
			return k, nil
		}
	}
	sz := uint32(len(e.uris)) + 1
	if _, ok := e.uris[sz]; ok {
		return 0, substraitgo.ErrKeyExists
	}

	e.uris[sz] = uri
	return sz, nil
}

type TopLevel interface {
	GetExtensionUris() []*extensions.SimpleExtensionURI
	GetExtensions() []*extensions.SimpleExtensionDeclaration
}

func GetExtensionSet(plan TopLevel) Set {
	uris := make(map[uint32]string)
	for _, uri := range plan.GetExtensionUris() {
		uris[uri.ExtensionUriAnchor] = uri.Uri
	}

	ret := &set{
		uris:             uris,
		funcMap:          make(map[uint32]ID),
		funcs:            make(map[ID]uint32),
		typesMap:         make(map[uint32]ID),
		types:            make(map[ID]uint32),
		typeVariationMap: make(map[uint32]ID),
		typeVariations:   make(map[ID]uint32),
	}

	for _, ext := range plan.GetExtensions() {
		switch e := ext.MappingType.(type) {
		case *extensions.SimpleExtensionDeclaration_ExtensionTypeVariation_:
			etv := e.ExtensionTypeVariation
			ret.encodeTypeVariation(etv.TypeVariationAnchor, ID{
				URI:  uris[etv.ExtensionUriReference],
				Name: etv.Name,
			})
		case *extensions.SimpleExtensionDeclaration_ExtensionType_:
			et := e.ExtensionType
			ret.encodeType(et.TypeAnchor, ID{
				URI:  uris[et.ExtensionUriReference],
				Name: et.Name,
			})
		case *extensions.SimpleExtensionDeclaration_ExtensionFunction_:
			ef := e.ExtensionFunction
			ret.encodeFunc(ef.FunctionAnchor, ID{
				URI:  uris[ef.ExtensionUriReference],
				Name: ef.Name,
			})
		}
	}

	return ret
}
