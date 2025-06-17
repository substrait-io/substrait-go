// SPDX-License-Identifier: Apache-2.0
package traverse

import (
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// MultiVisitor combines multiple visitors into a single visitor.
// All visitors are called for each node they support.
type MultiVisitor struct {
	visitors []Visitor
}

// NewMultiVisitor creates a visitor that delegates to multiple visitors.
func NewMultiVisitor(visitors ...Visitor) *MultiVisitor {
	return &MultiVisitor{visitors: visitors}
}

// VisitRel calls VisitRel on all visitors that implement RelVisitor.
func (m *MultiVisitor) VisitRel(rel *proto.Rel) {
	for _, v := range m.visitors {
		if rv, ok := v.(RelVisitor); ok {
			rv.VisitRel(rel)
		}
	}
}

// VisitExpr calls VisitExpr on all visitors that implement ExprVisitor.
func (m *MultiVisitor) VisitExpr(expr *proto.Expression) {
	for _, v := range m.visitors {
		if ev, ok := v.(ExprVisitor); ok {
			ev.VisitExpr(expr)
		}
	}
}
