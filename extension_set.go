// SPDX-License-Identifier: Apache-2.0

package substraitgo

import (
	"github.com/substrait-io/substrait-go/proto"
	"github.com/substrait-io/substrait-go/proto/extensions"
)

type Plan = proto.Plan

type ExtID struct {
	URI, Name string
}

type ExtensionSet struct {
	uris           map[uint32]string
	typeVariations map[uint32]ExtID
	typeIDs        map[uint32]ExtID
	funcIDs        map[uint32]ExtID
}

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
