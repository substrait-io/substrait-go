// SPDX-License-Identifier: Apache-2.0

package expr

import (
	"github.com/substrait-io/substrait-go/v4/extensions"
	"github.com/substrait-io/substrait-go/v4/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// Resolver provides functionality to resolve extension references and handle subquery expressions.
// It combines an extensions.Set for looking up extension definitions with a Collection for extension metadata,
// and optionally includes a SubqueryResolver for handling subquery expressions during protobuf deserialization.
type Resolver struct {
	extensions.Set
	c *extensions.Collection

	SubqueryResolver
}

// SubQueryResolver converts subqueries and the Relations within from the native protobuf format into an Expression.
type SubqueryResolver interface {
	HandleSubqueryFromProto(sub *proto.Expression_Subquery, baseSchema *types.RecordType, reg Resolver) (Expression, error)
}

// NewExtensionRegistry creates a new registry.  If you have an existing plan you can use GetExtensionSet() to
// populate an extensions.Set.
func NewExtensionRegistry(extSet extensions.Set, c *extensions.Collection) Resolver {
	if c == nil {
		panic("cannot create registry with nil collection")
	}
	return Resolver{Set: extSet, c: c}
}

// NewEmptyExtensionRegistry creates an empty registry useful starting from scratch.
func NewEmptyExtensionRegistry(c *extensions.Collection) Resolver {
	return NewExtensionRegistry(extensions.NewSet(), c)
}

func (e *Resolver) LookupTypeVariation(anchor uint32) (extensions.TypeVariation, bool) {
	return e.Set.LookupTypeVariation(anchor, e.c)
}

func (e *Resolver) LookupType(anchor uint32) (extensions.Type, bool) {
	return e.Set.LookupType(anchor, e.c)
}

// LookupScalarFunction returns a ScalarFunctionVariant associated with a previously used function's anchor.
func (e *Resolver) LookupScalarFunction(anchor uint32) (*extensions.ScalarFunctionVariant, bool) {
	return e.Set.LookupScalarFunction(anchor, e.c)
}

// LookupAggregateFunction returns an AggregateFunctionVariant associated with a previously used function's anchor.
func (e *Resolver) LookupAggregateFunction(anchor uint32) (*extensions.AggregateFunctionVariant, bool) {
	return e.Set.LookupAggregateFunction(anchor, e.c)
}

// LookupWindowFunction returns a WindowFunctionVariant associated with a previously used function's anchor.
func (e *Resolver) LookupWindowFunction(anchor uint32) (*extensions.WindowFunctionVariant, bool) {
	return e.Set.LookupWindowFunction(anchor, e.c)
}
