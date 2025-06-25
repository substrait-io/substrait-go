// SPDX-License-Identifier: Apache-2.0

package expr

import (
	"github.com/substrait-io/substrait-go/v4/extensions"
	"github.com/substrait-io/substrait-go/v4/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

type SubqueryHandler interface {
	HandleSubqueryFromProto(sub *proto.Expression_Subquery, baseSchema *types.RecordType, reg ExtensionRegistry) (Expression, error)
}

type ExtensionRegistry struct {
	extensions.Set
	c *extensions.Collection

	subqueryHandler SubqueryHandler
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

// WithSubqueryHandler returns a new ExtensionRegistry with the provided subquery handler set.
// This is primarily used for testing purposes.
func (e ExtensionRegistry) WithSubqueryHandler(handler SubqueryHandler) ExtensionRegistry {
	e.subqueryHandler = handler
	return e
}
