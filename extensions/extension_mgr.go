// SPDX-License-Identifier: Apache-2.0

package extensions

import (
	substraitgo "github.com/substrait-io/substrait-go"
	"github.com/substrait-io/substrait-go/proto"
	"github.com/substrait-io/substrait-go/proto/extensions"
)

type ID struct {
	URI, Name string
}

type Set struct {
	uris map[uint32]string

	typesMap map[uint32]struct {
		id  ID
		def *Type
	}
	types map[ID]uint32

	typeVariationMap map[uint32]struct {
		id  ID
		def *TypeVariation
	}
	typeVariations map[ID]uint32

	funcMap map[uint32]struct {
		id  ID
		def FunctionVariant
	}
	funcs map[ID]uint32
}

func (e *Set) Load(uri string, f *SimpleExtensionFile) error {
	_, err := e.addURI(uri)
	if err != nil {
		return err
	}

	for i := range f.Types {
		if err := e.RegisterType(uri, &f.Types[i]); err != nil {
			return err
		}
	}

	for i := range f.TypeVariations {
		if err := e.RegisterTypeVariation(uri, &f.TypeVariations[i]); err != nil {
			return err
		}
	}

	for _, f := range f.ScalarFunctions {
		if err := e.RegisterFunction(uri, &f); err != nil {
			return err
		}
	}

	for _, f := range f.AggregateFunctions {
		if err := e.RegisterFunction(uri, &f); err != nil {
			return err
		}
	}

	for _, f := range f.WindowFunctions {
		if err := e.RegisterFunction(uri, &f); err != nil {
			return err
		}
	}

	return nil
}

func (e *Set) RegisterFunction(uri string, def Function) error {
	if def == nil {
		return substraitgo.ErrInvalidType
	}

	_, err := e.findURI(uri)
	if err != nil {
		return err
	}

	add := func(id ID, v FunctionVariant) error {
		if _, ok := e.funcs[id]; ok {
			return substraitgo.ErrKeyExists
		}

		anchor := uint32(len(e.funcMap))
		tm, ok := e.funcMap[anchor]
		if ok {
			return substraitgo.ErrKeyExists
		}

		tm.id = id
		tm.def = v
		e.funcMap[anchor] = tm
		e.funcs[id] = anchor
		return nil
	}

	variants := def.ResolveURI(uri)
	for _, v := range variants {
		id := ID{URI: uri, Name: v.CompoundName()}
		if err := add(id, v); err != nil {
			return err
		}
	}

	// if there's only one implementation of a function for a given URI
	// then we need to also allow looking up by the regular name
	// instead of the compound name: i.e. "add" in addition to "add:i8_i8"
	if len(variants) == 1 {
		id := ID{URI: uri, Name: variants[0].Name()}
		return add(id, variants[0])
	}

	return nil
}

func (e *Set) RegisterTypeVariation(uri string, def *TypeVariation) error {
	if def == nil {
		return substraitgo.ErrInvalidType
	}

	_, err := e.findURI(uri)
	if err != nil {
		return err
	}

	id := ID{URI: uri, Name: def.Name}

	if _, ok := e.typeVariations[id]; ok {
		return substraitgo.ErrKeyExists
	}

	anchor := uint32(len(e.typeVariationMap))
	tm, ok := e.typeVariationMap[anchor]
	if ok {
		return substraitgo.ErrKeyExists
	}

	tm.id = id
	tm.def = def
	e.typeVariationMap[anchor] = tm
	e.typeVariations[id] = anchor
	return nil
}

func (e *Set) RegisterType(uri string, def *Type) error {
	if def == nil {
		return substraitgo.ErrInvalidType
	}

	_, err := e.findURI(uri)
	if err != nil {
		e.addURI(uri)
	}

	id := ID{URI: uri, Name: def.Name}

	if _, ok := e.types[id]; ok {
		return substraitgo.ErrKeyExists
	}

	anchor := uint32(len(e.typesMap))
	tm, ok := e.typesMap[anchor]
	if ok {
		return substraitgo.ErrKeyExists
	}

	tm.id = id
	tm.def = def
	e.typesMap[anchor] = tm
	e.types[id] = anchor
	return nil
}

func (e *Set) DecodeTypeVariation(anchor uint32) (ID, bool) {
	t, ok := e.typeVariationMap[anchor]
	if !ok {
		return ID{}, false
	}

	return t.id, true
}

func (e *Set) DecodeFunc(anchor uint32) (ID, bool) {
	f, ok := e.funcMap[anchor]
	if !ok {
		return ID{}, false
	}

	return f.id, true
}

func (e *Set) DecodeType(anchor uint32) (ID, bool) {
	t, ok := e.typesMap[anchor]
	if !ok {
		return ID{}, false
	}

	return t.id, true
}

func (e *Set) LookupTypeVariation(anchor uint32) (*TypeVariation, bool) {
	tv, ok := e.typeVariationMap[anchor]
	if !ok {
		return nil, false
	}
	return tv.def, true
}

func (e *Set) LookupType(anchor uint32) (*Type, bool) {
	t, ok := e.typesMap[anchor]
	if !ok {
		return nil, false
	}

	return t.def, true
}

func (e *Set) LookupFunction(anchor uint32) (FunctionVariant, bool) {
	f, ok := e.funcMap[anchor]
	if !ok {
		return nil, false
	}

	return f.def, true
}

func (e *Set) encodeType(anchor uint32, id ID) {
	tm := e.typesMap[anchor]
	tm.id = id
	e.typesMap[anchor] = tm
	e.types[id] = anchor
}

func (e *Set) encodeTypeVariation(anchor uint32, id ID) {
	tm := e.typeVariationMap[anchor]
	tm.id = id
	e.typeVariationMap[anchor] = tm
	e.typeVariations[id] = anchor
}

func (e *Set) encodeFunc(anchor uint32, id ID) {
	f := e.funcMap[anchor]
	f.id = id
	e.funcMap[anchor] = f
	e.funcs[id] = anchor
}

func (e *Set) findURI(uri string) (uint32, error) {
	for k, v := range e.uris {
		if v == uri {
			return k, nil
		}
	}
	return 0, substraitgo.ErrNotFound
}

func (e *Set) addURI(uri string) (uint32, error) {
	sz := uint32(len(e.uris))
	if _, ok := e.uris[sz]; ok {
		return 0, substraitgo.ErrKeyExists
	}

	e.uris[sz] = uri
	return sz, nil
}

func GetExtensionSet(plan *proto.Plan) *Set {
	uris := make(map[uint32]string)
	for _, uri := range plan.ExtensionUris {
		uris[uri.ExtensionUriAnchor] = uri.Uri
	}

	ret := &Set{
		uris: uris,
		funcMap: make(map[uint32]struct {
			id  ID
			def FunctionVariant
		}),
		funcs: make(map[ID]uint32),
		types: make(map[ID]uint32),
		typesMap: make(map[uint32]struct {
			id  ID
			def *Type
		}),
		typeVariationMap: make(map[uint32]struct {
			id  ID
			def *TypeVariation
		}),
		typeVariations: make(map[ID]uint32),
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
