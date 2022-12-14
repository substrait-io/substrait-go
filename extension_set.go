// SPDX-License-Identifier: Apache-2.0

package substraitgo

import (
	"github.com/substrait-io/substrait-go/proto"
	"github.com/substrait-io/substrait-go/proto/extensions"
)

type Plan = proto.Plan

type ExtensionRegistry interface {
	DecodeType(anchor uint32) (ExtID, bool)
	DecodeFunc(anchor uint32) (ExtID, bool)
	DecodeTypeVariation(anchor uint32) (ExtID, bool)
}

// ExtID represents a specific extension type by its URI and Name.
type ExtID struct {
	URI, Name string
}

// ExtensionSet is a specific set of SimpleExtension declarations,
// mapping anchor values to URIs and Names.
type ExtensionSet struct {
	uris           map[uint32]string
	typeVariations map[uint32]ExtID
	typeIDs        map[uint32]ExtID
	funcIDs        map[uint32]ExtID
}

func (e *ExtensionSet) GetTypeIDs() map[uint32]ExtID {
	return e.typeIDs
}

func (e *ExtensionSet) DecodeType(anchor uint32) (v ExtID, ok bool) {
	v, ok = e.typeIDs[anchor]
	return
}

func (e *ExtensionSet) DecodeFunc(anchor uint32) (v ExtID, ok bool) {
	v, ok = e.funcIDs[anchor]
	return
}

func (e *ExtensionSet) DecodeTypeVariation(anchor uint32) (v ExtID, ok bool) {
	v, ok = e.typeVariations[anchor]
	return
}

// GetExtensionSet processes and returns the set of all extensions which
// are depended on by the given Plan.
func GetExtensionSet(plan *Plan) *ExtensionSet {
	uris := make(map[uint32]string)
	for _, uri := range plan.ExtensionUris {
		uris[uri.ExtensionUriAnchor] = uri.Uri
	}

	var (
		typeVariations = make(map[uint32]ExtID)
		typeIDs        = make(map[uint32]ExtID)
		funcIDs        = make(map[uint32]ExtID)
	)
	for _, ext := range plan.Extensions {
		switch e := ext.MappingType.(type) {
		case *extensions.SimpleExtensionDeclaration_ExtensionTypeVariation_:
			etv := e.ExtensionTypeVariation
			typeVariations[etv.TypeVariationAnchor] = ExtID{
				URI:  uris[etv.ExtensionUriReference],
				Name: etv.Name,
			}
		case *extensions.SimpleExtensionDeclaration_ExtensionType_:
			et := e.ExtensionType
			typeIDs[et.TypeAnchor] = ExtID{
				URI:  uris[et.ExtensionUriReference],
				Name: et.Name,
			}
		case *extensions.SimpleExtensionDeclaration_ExtensionFunction_:
			ef := e.ExtensionFunction
			funcIDs[ef.FunctionAnchor] = ExtID{
				URI:  uris[ef.ExtensionUriReference],
				Name: ef.Name,
			}
		}
	}

	return &ExtensionSet{
		uris, typeVariations, typeIDs, funcIDs,
	}
}
