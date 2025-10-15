// SPDX-License-Identifier: Apache-2.0
/*
Package traverse provides a visitor framework for analyzing Substrait plans
with multiple read-only operations in a single tree traversal.

## Key Features:
- Interface-based visitor pattern
- Single tree walk for multiple analyses
- Sequential execution
- Read-only operations - NO MUTATIONS
- Zero allocations per node

## IMPORTANT: Read-Only Analysis Only
This interface is designed for read-only analysis. Modifying the tree while traversing
it is not supported and may lead to unpredictable behavior.

## Basic Usage:

	// Option 1: Simple usage with Visit (builds context automatically)
	visitor := NewMyVisitor()
	traverse.Visit(plan, visitor)

	// Option 2: Start traversal at any point in the relation tree
	visitor := NewMyVisitor()
	traverse.VisitRelation(someRelation, visitor)

## Creating Custom Visitors:

	// Visitor that only cares about relations
	type NodeCounter struct {
		count int
	}

	func (v *NodeCounter) VisitRel(rel *proto.Rel) {
		v.count++
	}

	// Visitor that only cares about expressions (e.g., finding all functions)
	type FunctionCollector struct {
		ctx *PlanContext
		functions []string
	}

	func (v *FunctionCollector) VisitExpr(expr *proto.Expression) {
		// Collect function names
	}

	// Visitor that visits both
	type FullAnalyzer struct{}

	func (v *FullAnalyzer) VisitRel(rel *proto.Rel) {
		// Analyze relations
	}

	func (v *FullAnalyzer) VisitExpr(expr *proto.Expression) {
		// Analyze expressions
	}
*/
package traverse

import (
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// Visitor is a marker interface for all visitors.
// Visitors should implement one or more of the following methods:
//   - VisitRel(*proto.Rel) - called for each relation node
//   - VisitExpr(*proto.Expression) - called for each non-literal expression
//
// The traversal framework uses type assertions to determine which methods to call.
type Visitor interface{}

// RelVisitor is implemented by visitors that want to visit relation nodes.
type RelVisitor interface {
	VisitRel(rel *proto.Rel)
}

// ExprVisitor is implemented by visitors that want to visit expression nodes.
type ExprVisitor interface {
	VisitExpr(expr *proto.Expression)
}

// Visit provides a convenient way to traverse a Substrait plan.
// It walks all relations in the plan, calling the appropriate visitor methods
// based on what interfaces the visitor implements.
func Visit(plan *proto.Plan, visitor Visitor) {
	if visitor == nil || plan == nil {
		return
	}

	// Walk all relations in the plan
	for _, rel := range plan.Relations {
		if root := rel.GetRoot(); root != nil {
			walk(root.Input, visitor)
		}
	}
}

// VisitRelation traverses a Substrait relation tree starting at the given relation,
// applying the visitor to each node. The visitor's methods are called based on which
// interfaces it implements. This function guarantees zero allocations per node visited.
//
// Note: This function visits the entire relation tree. There is currently no mechanism
// to stop traversal early if you encounter something you cannot handle.
//
// If your visitor needs plan context (for extensions, etc.), construct it
// with NewPlanContext and pass it to your visitor's constructor.
func VisitRelation(rel *proto.Rel, visitor Visitor) {
	if visitor == nil {
		return
	}
	walk(rel, visitor)
}
