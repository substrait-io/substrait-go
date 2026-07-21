package plan

import "github.com/substrait-io/substrait-go/v8/expr"

// WalkExprs calls exprFn on e and every Expression reachable from it, pre-order.
// When a subquery expression is encountered, relFn is called on its embedded Rel
// and WalkRels is used to continue traversal through it.
// Either fn may be nil.
func WalkExprs(e expr.Expression, exprFn func(expr.Expression), relFn func(Rel)) {
	if exprFn != nil {
		exprFn(e)
	}
	switch sq := e.(type) {
	case *ScalarSubquery:
		WalkRels(sq.Input, relFn, exprFn)
	case *InPredicateSubquery:
		WalkRels(sq.Haystack, relFn, exprFn)
	case *SetPredicateSubquery:
		WalkRels(sq.Tuples, relFn, exprFn)
	case *SetComparisonSubquery:
		WalkRels(sq.Right, relFn, exprFn)
	}
	for _, child := range e.GetExprs() {
		WalkExprs(child, exprFn, relFn)
	}
}

// WalkRels calls relFn on rel and every Rel reachable from it, pre-order.
// Expressions in each Rel are walked via WalkExprs, crossing subquery boundaries.
// Either fn may be nil.
func WalkRels(rel Rel, relFn func(Rel), exprFn func(expr.Expression)) {
	if relFn != nil {
		relFn(rel)
	}
	for _, e := range rel.GetExprs() {
		WalkExprs(e, exprFn, relFn)
	}
	for _, input := range rel.GetInputs() {
		WalkRels(input, relFn, exprFn)
	}
}
