// SPDX-License-Identifier: Apache-2.0

package extensions

import (
	"fmt"
	"io"

	substraitgo "github.com/substrait-io/substrait-go"
	"github.com/substrait-io/substrait-go/proto"
	"github.com/substrait-io/substrait-go/proto/extensions"
	"gopkg.in/yaml.v3"
)

type ID struct {
	URI, Name string
}

type Collection struct {
	uriSet map[string]struct{}

	simpleNameMap map[ID]string

	scalarMap        map[ID]*ScalarFunctionVariant
	aggregateMap     map[ID]*AggregateFunctionVariant
	windowMap        map[ID]*WindowFunctionVariant
	typeMap          map[ID]Type
	typeVariationMap map[ID]TypeVariation
}

func (c *Collection) GetType(id ID) (t Type, ok bool) {
	t, ok = c.typeMap[id]
	return
}

func (c *Collection) GetTypeVariation(id ID) (tv TypeVariation, ok bool) {
	tv, ok = c.typeVariationMap[id]
	return
}

var void = struct{}{}

type variants interface {
	*ScalarFunctionVariant | *AggregateFunctionVariant | *WindowFunctionVariant
	Name() string
	CompoundName() string
}

func checkMaps[T variants](id ID, m map[ID]T, simpleNames map[ID]string) (T, bool) {
	if sv, ok := m[id]; ok {
		return sv, true
	}

	if compound, ok := simpleNames[id]; ok {
		id.Name = compound
		if sv, ok := m[id]; ok {
			return sv, true
		}
	}

	return nil, false
}

type extFn[T variants] interface {
	GetVariants(uri string) []T
}

func addToMaps[T variants](id ID, fn extFn[T], m map[ID]T, simpleMap map[string]string) {
	variants := fn.GetVariants(id.URI)
	for _, v := range variants {
		id.Name = v.CompoundName()
		m[id] = v
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
	return checkMaps(id, c.scalarMap, c.simpleNameMap)
}

func (c *Collection) GetAggregateFunc(id ID) (*AggregateFunctionVariant, bool) {
	return checkMaps(id, c.aggregateMap, c.simpleNameMap)
}

func (c *Collection) GetWindowFunc(id ID) (*WindowFunctionVariant, bool) {
	return checkMaps(id, c.windowMap, c.simpleNameMap)
}

func (c *Collection) init() {
	if c.uriSet == nil {
		c.uriSet = make(map[string]struct{})
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

	id := ID{URI: uri}
	for _, t := range file.Types {
		id.Name = t.Name
		c.typeMap[id] = t
	}

	for _, t := range file.TypeVariations {
		id.Name = t.Name
		c.typeVariationMap[id] = t
	}

	// if there is only one implementation for a given function
	// it should be findable by its simple name in addition to the
	// compound name.
	simpleNames := make(map[string]string)

	for _, f := range file.ScalarFunctions {
		addToMaps[*ScalarFunctionVariant](id, &f, c.scalarMap, simpleNames)
	}

	for _, f := range file.AggregateFunctions {
		addToMaps[*AggregateFunctionVariant](id, &f, c.aggregateMap, simpleNames)
	}

	for _, f := range file.WindowFunctions {
		addToMaps[*WindowFunctionVariant](id, &f, c.windowMap, simpleNames)
	}

	// add simple name aliases
	for k, v := range simpleNames {
		id.Name = k
		c.simpleNameMap[id] = v
	}

	return nil
}

func (c *Collection) URILoaded(uri string) bool {
	_, ok := c.uriSet[uri]
	return ok
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
		e.addURI(id.URI)
		a = uint32(len(e.types))
		e.types[id] = a
		e.typesMap[a] = id
	}
	return a
}

func (e *set) GetFuncAnchor(id ID) uint32 {
	a, ok := e.funcs[id]
	if !ok {
		e.addURI(id.URI)
		a = uint32(len(e.funcs))
		e.funcs[id] = a
		e.funcMap[a] = id
	}
	return a
}

func (e *set) GetTypeVariationAnchor(id ID) uint32 {
	a, ok := e.typeVariations[id]
	if !ok {
		e.addURI(id.URI)
		a = uint32(len(e.typeVariations))
		e.typeVariations[id] = a
		e.typeVariationMap[a] = id
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

func (e *set) addURI(uri string) (uint32, error) {
	sz := uint32(len(e.uris))
	if _, ok := e.uris[sz]; ok {
		return 0, substraitgo.ErrKeyExists
	}

	e.uris[sz] = uri
	return sz, nil
}

func GetExtensionSet(plan *proto.Plan) Set {
	uris := make(map[uint32]string)
	for _, uri := range plan.ExtensionUris {
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

	for _, ext := range plan.Extensions {
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
