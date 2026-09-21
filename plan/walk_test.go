// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v8/expr"
	"github.com/substrait-io/substrait-go/v8/plan"
)

// TestWalkRels verifies that WalkRels visits every Rel and Expression in the
// tree, including Rels embedded inside subquery expressions (crossing subquery
// boundaries).
//
// Plan shape:
//
//	ProjectRel  [exprs: ScalarSubquery(innerScan)]
//	  FilterRel [cond: literal(true)]
//	    outerScan
//
// innerScan is only reachable through the subquery boundary.
func TestWalkRels(t *testing.T) {
	b := plan.NewBuilderDefault()

	innerScan := b.NamedScan([]string{"inner"}, baseSchema)
	subquery := plan.NewScalarSubquery(innerScan)

	outerScan := b.NamedScan([]string{"outer"}, baseSchema2)
	filter, err := b.Filter(outerScan, expr.NewPrimitiveLiteral(true, false))
	require.NoError(t, err)
	project, err := b.Project(filter, subquery)
	require.NoError(t, err)

	var visitedRels []plan.Rel
	var visitedExprs []expr.Expression
	plan.WalkRels(project,
		func(r plan.Rel) { visitedRels = append(visitedRels, r) },
		func(e expr.Expression) { visitedExprs = append(visitedExprs, e) },
	)

	// ProjectRel, FilterRel, outerScan, innerScan (via subquery boundary)
	assert.Len(t, visitedRels, 4)
	// ScalarSubquery (project expr), literal(true) (filter cond)
	assert.Len(t, visitedExprs, 2)

	assert.IsType(t, &plan.ScalarSubquery{}, visitedExprs[0])
}
