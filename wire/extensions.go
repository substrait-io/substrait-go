// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"fmt"
	"sort"

	substraitgo "github.com/substrait-io/substrait-go/v9"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	extensionspb "github.com/substrait-io/substrait-protobuf/go/substraitpb/extensions"
	"google.golang.org/protobuf/types/known/anypb"
)

// advancedExtensionToProto encodes a domain AdvancedExtension as its protobuf message.
func advancedExtensionToProto(a *extensions.AdvancedExtension) *extensionspb.AdvancedExtension {
	if a == nil {
		return nil
	}
	var optimizations []*anypb.Any
	for _, o := range a.Optimizations {
		optimizations = append(optimizations, (*anypb.Any)(o))
	}
	return &extensionspb.AdvancedExtension{
		Optimization: optimizations,
		Enhancement:  (*anypb.Any)(a.Enhancement),
	}
}

// advancedExtensionFromProto converts a protobuf AdvancedExtension to the domain type.
func advancedExtensionFromProto(a *extensionspb.AdvancedExtension) *extensions.AdvancedExtension {
	if a == nil {
		return nil
	}
	var optimizations []*extensions.Optimization
	for _, o := range a.Optimization {
		optimizations = append(optimizations, (*extensions.Optimization)(o))
	}
	return &extensions.AdvancedExtension{
		Optimizations: optimizations,
		Enhancement:   (*extensions.Enhancement)(a.Enhancement),
	}
}

// ExtensionsToProto returns the URN and declaration messages for a registry's
// extension set.
func ExtensionsToProto(reg expr.ExtensionRegistry) ([]*extensionspb.SimpleExtensionURN, []*extensionspb.SimpleExtensionDeclaration) {
	return ExtensionSetToProto(reg.Set)
}

// ExtensionSetToProto encodes an extension set as its protobuf URN and
// declaration messages, sorted by anchor for stable output.
func ExtensionSetToProto(s extensions.Set) ([]*extensionspb.SimpleExtensionURN, []*extensionspb.SimpleExtensionDeclaration) {
	urnMap := s.URNs()
	urnBackRef := make(map[string]uint32)

	urns := make([]*extensionspb.SimpleExtensionURN, 0, len(urnMap))
	for anchor, urn := range urnMap {
		urnBackRef[urn] = anchor
		urns = append(urns, &extensionspb.SimpleExtensionURN{
			ExtensionUrnAnchor: anchor,
			Urn:                urn,
		})
	}

	// Sort URN extensions by the anchor for consistent output
	sort.Slice(urns, func(i, j int) bool { return urns[i].ExtensionUrnAnchor < urns[j].ExtensionUrnAnchor })

	types := s.Types()
	typeVariations := s.TypeVariations()
	funcs := s.Functions()

	decls := make([]*extensionspb.SimpleExtensionDeclaration, 0, len(types)+len(typeVariations)+len(funcs))
	for id, anchor := range types {
		decls = append(decls, &extensionspb.SimpleExtensionDeclaration{
			MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionType_{
				ExtensionType: &extensionspb.SimpleExtensionDeclaration_ExtensionType{
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

	for id, anchor := range typeVariations {
		decls = append(decls, &extensionspb.SimpleExtensionDeclaration{
			MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionTypeVariation_{
				ExtensionTypeVariation: &extensionspb.SimpleExtensionDeclaration_ExtensionTypeVariation{
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
	for id, anchor := range funcs {
		decls = append(decls, &extensionspb.SimpleExtensionDeclaration{
			MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction_{
				ExtensionFunction: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction{
					ExtensionUrnReference: urnBackRef[id.URN],
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

	return urns, decls
}

// extensionCarrier is satisfied by the protobuf plan and extended-expression
// messages, which both carry a set of extension URNs and declarations.
type extensionCarrier interface {
	GetExtensionUrns() []*extensionspb.SimpleExtensionURN
	GetExtensions() []*extensionspb.SimpleExtensionDeclaration
}

// GetExtensionSet decodes the extension URNs and declarations of a plan (or
// extended expression) into an extensions.Set, resolving each declaration's URN
// reference against the collection c.
func GetExtensionSet(plan extensionCarrier, c *extensions.Collection) (extensions.Set, error) {
	urns := make(map[uint32]string)
	for _, urn := range plan.GetExtensionUrns() {
		urns[urn.ExtensionUrnAnchor] = urn.Urn
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

	types := make(map[uint32]extensions.TypeID)
	typeVariations := make(map[uint32]extensions.TypeVariationID)
	funcs := make(map[uint32]extensions.FunctionID)

	for _, ext := range plan.GetExtensions() {
		switch e := ext.MappingType.(type) {
		case *extensionspb.SimpleExtensionDeclaration_ExtensionTypeVariation_:
			etv := e.ExtensionTypeVariation
			urn, err := resolveRefToURN(etv.ExtensionUrnReference)
			if err != nil {
				return nil, err
			}
			typeVariations[etv.TypeVariationAnchor] = extensions.TypeVariationID{URN: urn, Name: etv.Name}
		case *extensionspb.SimpleExtensionDeclaration_ExtensionType_:
			et := e.ExtensionType
			urn, err := resolveRefToURN(et.ExtensionUrnReference)
			if err != nil {
				return nil, err
			}
			types[et.TypeAnchor] = extensions.TypeID{URN: urn, Name: et.Name}
		case *extensionspb.SimpleExtensionDeclaration_ExtensionFunction_:
			ef := e.ExtensionFunction
			urn, err := resolveRefToURN(ef.ExtensionUrnReference)
			if err != nil {
				return nil, err
			}
			funcs[ef.FunctionAnchor] = extensions.FunctionID{URN: urn, Name: ef.Name}
		}
	}

	return extensions.NewSetFromParts(urns, types, typeVariations, funcs), nil
}
