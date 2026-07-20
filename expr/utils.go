// SPDX-License-Identifier: Apache-2.0

package expr

import (
	"github.com/substrait-io/substrait-go/v8/extensions"
	"github.com/substrait-io/substrait-go/v8/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	extensionspb "github.com/substrait-io/substrait-protobuf/go/substraitpb/extensions"
	"google.golang.org/protobuf/types/known/anypb"
)

// ExtensionRelDecoder decodes an extension relation's Any detail into an
// ExtensionRelDefinition that knows its output schema and expressions.
// It is called by RelFromProto instead of falling back to UndecodedExtension
// when a decoder is registered for the detail's type URL.
//
// The Rel type is intentionally left as an interface{} here to avoid an import
// cycle between the expr and plan packages; the plan package casts it to plan.Rel
// via the ExtensionRelDecoderFunc type alias defined there.
type ExtensionRelDecoder interface {
	// DecodeExtensionRel attempts to decode detail into an ExtensionRelDefinition.
	// Returns (nil, nil) to signal that this decoder does not handle the given detail,
	// allowing fallback to UndecodedExtension.
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

	// extensionRelDecoder is optionally set to decode extension relation details
	// into typed ExtensionRelDefinitions during RelFromProto.
	extensionRelDecoder ExtensionRelDecoder
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

// SetExtensionRelDecoder registers a decoder that RelFromProto will call for
// each ExtensionSingle/ExtensionMulti/ExtensionLeaf relation. If the decoder
// returns a non-nil value it is used as the ExtensionRelDefinition; if it
// returns (nil, nil) RelFromProto falls back to UndecodedExtension.
func (e *ExtensionRegistry) SetExtensionRelDecoder(decoder ExtensionRelDecoder) {
	e.extensionRelDecoder = decoder
}

// ExtensionRelDecoderFor returns the registered decoder, or nil if none was set.
func (e *ExtensionRegistry) ExtensionRelDecoderFor() ExtensionRelDecoder {
	return e.extensionRelDecoder
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
