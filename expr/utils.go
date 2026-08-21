// SPDX-License-Identifier: Apache-2.0

package expr

import (
	"fmt"

	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	extensionspb "github.com/substrait-io/substrait-protobuf/go/substraitpb/extensions"
	"google.golang.org/protobuf/types/known/anypb"
)

// ExtensionRelDecoder decodes an extension relation's Any detail into an
// ExtensionRelDefinition that knows its output schema and expressions.
// It is called by RelFromProto for extension rels whose type URL has a
// registered decoder; unregistered type URLs fall back to UndecodedExtension.
//
// The return type is intentionally left as any to avoid an import cycle between
// the expr and plan packages; callers cast it to the appropriate concrete type.
type ExtensionRelDecoder interface {
	// DecodeExtensionRel decodes detail into a plan.ExtensionRelDefinition.
	// The return type is any because ExtensionRelDefinition is defined in the plan
	// package, and importing it here would create a cycle with the expr package.
	DecodeExtensionRel(detail *anypb.Any) (any, error)
}

// ExtensionRegistry provides functionality to resolve extension references and handle subquery expressions.
// It combines an extensions.Set for looking up extension definitions with a Collection for extension metadata.
type ExtensionRegistry struct {
	extensions.Set
	c *extensions.Collection

	// subqueryConverter is injected by the plan package to handle subquery expressions
	// TODO: We may want to consider refactoring to make a cleaner interface here
	subqueryConverter

	// extensionRelDecoders maps type URLs to decoders for extension relation details.
	extensionRelDecoders map[string]ExtensionRelDecoder
}

// subqueryConverter converts subqueries and the Relations within from the native
// protobuf format into an Expression.
//
// This interface is private to avoid exposing the dependency cycle - a Subquery
// contains a Plan, so the implementor of this has to exist in / import the plan
// package, which we can't do here without creating a cycle with the expr
// package.
//
// TODO: We may want to refactor this interface to be more generic or use a
// different approach to avoid the cycle.
type subqueryConverter interface {
	SubqueryFromProto(sub *proto.Expression_Subquery, baseSchema *types.RecordType, reg ExtensionRegistry) (Expression, error)
}

// SetSubqueryConverter allows the plan package to inject a subquery converter.
// This is an internal function used to break the dependency cycle between expr and plan packages.
func (e *ExtensionRegistry) SetSubqueryConverter(converter subqueryConverter) {
	e.subqueryConverter = converter
}

// SetExtensionRelDecoder registers a decoder for the given type URL.
// RelFromProto dispatches to it when it encounters an extension rel whose
// detail matches typeURL; unregistered type URLs fall back to UndecodedExtension.
// Returns an error if a decoder is already registered for typeURL.
func (e *ExtensionRegistry) SetExtensionRelDecoder(typeURL string, decoder ExtensionRelDecoder) error {
	if e.extensionRelDecoders == nil {
		e.extensionRelDecoders = make(map[string]ExtensionRelDecoder)
	}
	if _, exists := e.extensionRelDecoders[typeURL]; exists {
		return fmt.Errorf("decoder already registered for type URL %q", typeURL)
	}
	e.extensionRelDecoders[typeURL] = decoder
	return nil
}

// ExtensionRelDecoderFor returns the decoder registered for typeURL, or nil if none was set.
func (e *ExtensionRegistry) ExtensionRelDecoderFor(typeURL string) ExtensionRelDecoder {
	return e.extensionRelDecoders[typeURL]
}

// NewExtensionRegistry creates a new registry.  If you have an existing plan you can use GetExtensionSet() to
// populate an extensions.Set.
func NewExtensionRegistry(extSet extensions.Set, c *extensions.Collection) ExtensionRegistry {
	if c == nil {
		panic("cannot create registry with nil collection")
	}
	return ExtensionRegistry{Set: extSet, c: c}
}

// NewEmptyExtensionRegistry creates an empty registry useful starting from scratch.
func NewEmptyExtensionRegistry(c *extensions.Collection) ExtensionRegistry {
	return NewExtensionRegistry(extensions.NewSet(), c)
}

func (e *ExtensionRegistry) LookupTypeVariation(anchor uint32) (extensions.TypeVariation, bool) {
	return e.Set.LookupTypeVariation(anchor, e.c)
}

func (e *ExtensionRegistry) LookupType(anchor uint32) (extensions.Type, bool) {
	return e.Set.LookupType(anchor, e.c)
}

// LookupScalarFunction returns a ScalarFunctionVariant associated with a previously used function's anchor.
func (e *ExtensionRegistry) LookupScalarFunction(anchor uint32) (*extensions.ScalarFunctionVariant, bool) {
	return e.Set.LookupScalarFunction(anchor, e.c)
}

// LookupAggregateFunction returns an AggregateFunctionVariant associated with a previously used function's anchor.
func (e *ExtensionRegistry) LookupAggregateFunction(anchor uint32) (*extensions.AggregateFunctionVariant, bool) {
	return e.Set.LookupAggregateFunction(anchor, e.c)
}

// LookupWindowFunction returns a WindowFunctionVariant associated with a previously used function's anchor.
func (e *ExtensionRegistry) LookupWindowFunction(anchor uint32) (*extensions.WindowFunctionVariant, bool) {
	return e.Set.LookupWindowFunction(anchor, e.c)
}

// ExtensionsToProto returns the URNs and declarations from the extension set using the registry's collection.
func (e *ExtensionRegistry) ExtensionsToProto() ([]*extensionspb.SimpleExtensionURN, []*extensionspb.SimpleExtensionDeclaration) {
	return e.Set.ToProto(e.c)
}
