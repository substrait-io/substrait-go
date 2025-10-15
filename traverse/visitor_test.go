// SPDX-License-Identifier: Apache-2.0
package traverse

import (
	"testing"

	"github.com/stretchr/testify/require"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// Tests

// nodeTracker is a test visitor that tracks both relations and expressions
type nodeTracker struct {
	rels  []string
	exprs []string
}

// Implement RelVisitor interface
func (n *nodeTracker) VisitRel(rel *proto.Rel) {
	n.rels = append(n.rels, "rel")
}

// Implement ExprVisitor interface
func (n *nodeTracker) VisitExpr(expr *proto.Expression) {
	n.exprs = append(n.exprs, "expr")
}

// relOnlyTracker only implements RelVisitor
type relOnlyTracker struct {
	count int
}

func (r *relOnlyTracker) VisitRel(rel *proto.Rel) {
	r.count++
}

// exprOnlyTracker only implements ExprVisitor
type exprOnlyTracker struct {
	functions []uint32
}

func (e *exprOnlyTracker) VisitExpr(expr *proto.Expression) {
	if sf := expr.GetScalarFunction(); sf != nil {
		e.functions = append(e.functions, sf.FunctionReference)
	}
}

// relTypeCounter is a test visitor that counts different relation types
type relTypeCounter struct {
	reads    int
	projects int
	joins    int
}

func (r *relTypeCounter) VisitRel(rel *proto.Rel) {
	switch rel.RelType.(type) {
	case *proto.Rel_Read:
		r.reads++
	case *proto.Rel_Project:
		r.projects++
	case *proto.Rel_Join:
		r.joins++
	}
}

// allNodeTracker is a test visitor that tracks all node types
type allNodeTracker struct {
	relCount  int
	exprCount int
}

func (a *allNodeTracker) VisitRel(rel *proto.Rel) {
	a.relCount++
}

func (a *allNodeTracker) VisitExpr(expr *proto.Expression) {
	a.exprCount++
}

func TestVisitorFramework(t *testing.T) {
	t.Run("Basic Visitor Interface", func(t *testing.T) {
		// Create a simple test plan programmatically
		join := &proto.Rel{
			RelType: &proto.Rel_Join{
				Join: &proto.JoinRel{
					Left: &proto.Rel{
						RelType: &proto.Rel_Read{
							Read: &proto.ReadRel{},
						},
					},
					Right: &proto.Rel{
						RelType: &proto.Rel_Read{
							Read: &proto.ReadRel{},
						},
					},
				},
			},
		}

		// Test basic visitor functionality for data collection
		nodeCountVisitor := newStatefulVisitor(0, func(count int, rel *proto.Rel) int {
			return count + 1
		})

		joinCountVisitor := newStatefulVisitor(0, func(count int, rel *proto.Rel) int {
			if _, isJoin := rel.RelType.(*proto.Rel_Join); isJoin {
				return count + 1
			}
			return count
		})

		VisitRelation(join, NewMultiVisitor(nodeCountVisitor, joinCountVisitor))

		require.Equal(t, 3, nodeCountVisitor.Result()) // join + 2 reads
		require.Equal(t, 1, joinCountVisitor.Result()) // 1 join
	})

	t.Run("Multiple Visitors Single Walk", func(t *testing.T) {
		// Create a simple test plan programmatically
		join := &proto.Rel{
			RelType: &proto.Rel_Join{
				Join: &proto.JoinRel{
					Left: &proto.Rel{
						RelType: &proto.Rel_Read{
							Read: &proto.ReadRel{},
						},
					},
					Right: &proto.Rel{
						RelType: &proto.Rel_Project{
							Project: &proto.ProjectRel{
								Input: &proto.Rel{
									RelType: &proto.Rel_Read{
										Read: &proto.ReadRel{},
									},
								},
							},
						},
					},
				},
			},
		}

		// Test multiple visitors in single walk
		nodeCountVisitor := newStatefulVisitor(0, func(count int, rel *proto.Rel) int {
			return count + 1
		})

		projectCountVisitor := newStatefulVisitor(0, func(count int, rel *proto.Rel) int {
			if _, isProject := rel.RelType.(*proto.Rel_Project); isProject {
				return count + 1
			}
			return count
		})

		VisitRelation(join, NewMultiVisitor(nodeCountVisitor, projectCountVisitor))

		require.Equal(t, 4, nodeCountVisitor.Result())    // join + project + 2 reads
		require.Equal(t, 1, projectCountVisitor.Result()) // 1 project
	})

	t.Run("Custom Visitor Types", func(t *testing.T) {
		// Create a simple test plan
		readNode := &proto.Rel{
			RelType: &proto.Rel_Read{
				Read: &proto.ReadRel{},
			},
		}

		// Test string visitor
		nodeTypeVisitor := newStatefulVisitor("", func(acc string, rel *proto.Rel) string {
			var nodeType string
			switch rel.RelType.(type) {
			case *proto.Rel_Join:
				nodeType = "join"
			case *proto.Rel_Read:
				nodeType = "read"
			case *proto.Rel_Project:
				nodeType = "project"
			default:
				nodeType = "other"
			}

			if acc == "" {
				return nodeType
			}
			return acc + "," + nodeType
		})

		// Test slice visitor
		nodeListVisitor := newStatefulVisitor([]string{}, func(nodes []string, rel *proto.Rel) []string {
			switch rel.RelType.(type) {
			case *proto.Rel_Read:
				return append(nodes, "read")
			case *proto.Rel_Join:
				return append(nodes, "join")
			}
			return nodes
		})

		VisitRelation(readNode, NewMultiVisitor(nodeTypeVisitor, nodeListVisitor))

		require.Equal(t, "read", nodeTypeVisitor.Result())
		require.Equal(t, []string{"read"}, nodeListVisitor.Result())
	})

	t.Run("Empty Visitor List", func(t *testing.T) {
		readNode := &proto.Rel{
			RelType: &proto.Rel_Read{
				Read: &proto.ReadRel{},
			},
		}

		// Should not panic with nil visitor
		VisitRelation(readNode, nil)
	})

	t.Run("Nil Root Handling", func(t *testing.T) {
		nodeCountVisitor := newStatefulVisitor(0, func(count int, rel *proto.Rel) int {
			return count + 1
		})

		// Should not panic with nil root
		VisitRelation(nil, nodeCountVisitor)

		// Should remain at initial state
		require.Equal(t, 0, nodeCountVisitor.Result())
	})

	t.Run("StatefulVisitor Interface", func(t *testing.T) {
		// Create a stateful visitor and verify it implements Visitor interface
		nodeCountVisitor := newStatefulVisitor(0, func(count int, rel *proto.Rel) int {
			return count + 1
		})

		// Test that it implements Visitor interface
		var _ Visitor = nodeCountVisitor

		// Test initial state
		require.Equal(t, 0, nodeCountVisitor.Result())

		// Test visitation
		readNode := &proto.Rel{
			RelType: &proto.Rel_Read{
				Read: &proto.ReadRel{},
			},
		}

		nodeCountVisitor.VisitRel(readNode)
		require.Equal(t, 1, nodeCountVisitor.Result())
	})

	t.Run("Expression-only Visitor", func(t *testing.T) {
		// Create a test plan with expressions
		project := &proto.Rel{
			RelType: &proto.Rel_Project{
				Project: &proto.ProjectRel{
					Input: &proto.Rel{
						RelType: &proto.Rel_Read{
							Read: &proto.ReadRel{},
						},
					},
					Expressions: []*proto.Expression{
						{
							RexType: &proto.Expression_ScalarFunction_{
								ScalarFunction: &proto.Expression_ScalarFunction{
									FunctionReference: 42,
								},
							},
						},
						{
							RexType: &proto.Expression_ScalarFunction_{
								ScalarFunction: &proto.Expression_ScalarFunction{
									FunctionReference: 43,
								},
							},
						},
					},
				},
			},
		}

		// Visitor that only cares about expressions
		exprTracker := &exprOnlyTracker{}
		VisitRelation(project, exprTracker)

		require.Equal(t, []uint32{42, 43}, exprTracker.functions)
	})

	t.Run("Relation-only Visitor", func(t *testing.T) {
		// Create a test plan
		project := &proto.Rel{
			RelType: &proto.Rel_Project{
				Project: &proto.ProjectRel{
					Input: &proto.Rel{
						RelType: &proto.Rel_Read{
							Read: &proto.ReadRel{},
						},
					},
					Expressions: []*proto.Expression{
						{
							RexType: &proto.Expression_Selection{
								Selection: &proto.Expression_FieldReference{},
							},
						},
					},
				},
			},
		}

		// Visitor that only cares about relations
		relTracker := &relOnlyTracker{}
		VisitRelation(project, relTracker)

		require.Equal(t, 2, relTracker.count) // project + read
	})

	t.Run("Idiomatic Custom Visitor", func(t *testing.T) {
		// This is the idiomatic Go way: create a type that implements the interfaces
		tracker := &nodeTracker{}

		// Create a test plan
		project := &proto.Rel{
			RelType: &proto.Rel_Project{
				Project: &proto.ProjectRel{
					Input: &proto.Rel{
						RelType: &proto.Rel_Read{
							Read: &proto.ReadRel{},
						},
					},
					Expressions: []*proto.Expression{
						{
							RexType: &proto.Expression_Selection{
								Selection: &proto.Expression_FieldReference{},
							},
						},
					},
				},
			},
		}

		VisitRelation(project, tracker)

		require.Equal(t, 2, len(tracker.rels))  // project + read
		require.Equal(t, 1, len(tracker.exprs)) // selection expression
	})

	t.Run("Direct Interface Implementation", func(t *testing.T) {
		// Custom visitor that doesn't use StatefulVisitor
		// Create a test plan
		join := &proto.Rel{
			RelType: &proto.Rel_Join{
				Join: &proto.JoinRel{
					Left: &proto.Rel{
						RelType: &proto.Rel_Read{
							Read: &proto.ReadRel{},
						},
					},
					Right: &proto.Rel{
						RelType: &proto.Rel_Project{
							Project: &proto.ProjectRel{
								Input: &proto.Rel{
									RelType: &proto.Rel_Read{
										Read: &proto.ReadRel{},
									},
								},
							},
						},
					},
				},
			},
		}

		counter := &relTypeCounter{}
		VisitRelation(join, counter)

		require.Equal(t, 2, counter.reads)
		require.Equal(t, 1, counter.projects)
		require.Equal(t, 1, counter.joins)
	})

	t.Run("Mixed Visitor Types", func(t *testing.T) {
		// Mix of StatefulVisitor and direct interface implementation
		nodeCounter := newStatefulVisitor(0, func(count int, rel *proto.Rel) int {
			return count + 1
		})

		tracker := &nodeTracker{}

		// Create a test plan
		project := &proto.Rel{
			RelType: &proto.Rel_Project{
				Project: &proto.ProjectRel{
					Input: &proto.Rel{
						RelType: &proto.Rel_Read{
							Read: &proto.ReadRel{},
						},
					},
					Expressions: []*proto.Expression{
						{
							RexType: &proto.Expression_Selection{
								Selection: &proto.Expression_FieldReference{},
							},
						},
					},
				},
			},
		}

		// Both visitors work together
		VisitRelation(project, NewMultiVisitor(nodeCounter, tracker))

		require.Equal(t, 2, nodeCounter.Result()) // 2 relations
		require.Equal(t, 2, len(tracker.rels))    // project + read
		require.Equal(t, 1, len(tracker.exprs))   // selection expression
	})

	t.Run("Aggregate GroupingExpressions Traversal", func(t *testing.T) {
		// Create an aggregate with grouping expressions at the relation level
		groupExpr1 := &proto.Expression{
			RexType: &proto.Expression_Selection{
				Selection: &proto.Expression_FieldReference{},
			},
		}
		groupExpr2 := &proto.Expression{
			RexType: &proto.Expression_ScalarFunction_{
				ScalarFunction: &proto.Expression_ScalarFunction{
					FunctionReference: 1,
				},
			},
		}

		aggregate := &proto.Rel{
			RelType: &proto.Rel_Aggregate{
				Aggregate: &proto.AggregateRel{
					Input: &proto.Rel{RelType: &proto.Rel_Read{Read: &proto.ReadRel{}}},
					// Modern way: expressions at relation level, referenced by index
					GroupingExpressions: []*proto.Expression{groupExpr1, groupExpr2},
					Groupings: []*proto.AggregateRel_Grouping{
						{
							// Reference expressions by index
							ExpressionReferences: []uint32{0, 1},
						},
					},
					Measures: []*proto.AggregateRel_Measure{
						{
							Measure: &proto.AggregateFunction{
								FunctionReference: 10, // SUM
							},
						},
					},
				},
			},
		}

		// Count all expressions visited
		tracker := &allNodeTracker{}

		VisitRelation(aggregate, tracker)

		// Should visit both grouping expressions
		require.Equal(t, 2, tracker.exprCount, "Should visit both grouping expressions")
		require.Equal(t, 2, tracker.relCount, "Should visit aggregate and read relations")
	})

	t.Run("Visit Plan Convenience Function", func(t *testing.T) {
		// Create a plan with multiple relations
		plan := &proto.Plan{
			Relations: []*proto.PlanRel{
				{
					RelType: &proto.PlanRel_Root{
						Root: &proto.RelRoot{
							Input: &proto.Rel{
								RelType: &proto.Rel_Read{
									Read: &proto.ReadRel{},
								},
							},
						},
					},
				},
				{
					RelType: &proto.PlanRel_Root{
						Root: &proto.RelRoot{
							Input: &proto.Rel{
								RelType: &proto.Rel_Project{
									Project: &proto.ProjectRel{
										Input: &proto.Rel{
											RelType: &proto.Rel_Read{
												Read: &proto.ReadRel{},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}

		counter := &relOnlyTracker{}
		Visit(plan, counter)

		require.Equal(t, 3, counter.count) // 2 reads + 1 project
	})

	t.Run("StatefulVisitor with Expression Support", func(t *testing.T) {
		// Create a plan with both relations and expressions
		project := &proto.Rel{
			RelType: &proto.Rel_Project{
				Project: &proto.ProjectRel{
					Input: &proto.Rel{
						RelType: &proto.Rel_Read{
							Read: &proto.ReadRel{},
						},
					},
					Expressions: []*proto.Expression{
						{
							RexType: &proto.Expression_ScalarFunction_{
								ScalarFunction: &proto.Expression_ScalarFunction{
									FunctionReference: 1,
								},
							},
						},
						{
							RexType: &proto.Expression_ScalarFunction_{
								ScalarFunction: &proto.Expression_ScalarFunction{
									FunctionReference: 2,
								},
							},
						},
					},
				},
			},
		}

		// Create a visitor that tracks both relations and expressions
		type Stats struct {
			RelCount     int
			ExprCount    int
			FunctionRefs []uint32
		}

		visitor := newStatefulVisitorWithExpr(
			Stats{},
			func(s Stats, rel *proto.Rel) Stats {
				s.RelCount++
				return s
			},
			func(s Stats, expr *proto.Expression) Stats {
				s.ExprCount++
				if sf := expr.GetScalarFunction(); sf != nil {
					s.FunctionRefs = append(s.FunctionRefs, sf.FunctionReference)
				}
				return s
			},
		)

		VisitRelation(project, visitor)

		result := visitor.Result()
		require.Equal(t, 2, result.RelCount)                  // project + read
		require.Equal(t, 2, result.ExprCount)                 // 2 scalar functions
		require.Equal(t, []uint32{1, 2}, result.FunctionRefs) // function references
	})
}

// statefulVisitor is a test helper that accumulates state as it visits nodes.
// It's generic over the state type S.
type statefulVisitor[S any] struct {
	state     S
	visitRel  func(S, *proto.Rel) S
	visitExpr func(S, *proto.Expression) S
}

// newStatefulVisitor creates a visitor that only visits relations
func newStatefulVisitor[S any](initial S, visit func(S, *proto.Rel) S) *statefulVisitor[S] {
	return &statefulVisitor[S]{
		state:    initial,
		visitRel: visit,
	}
}

// newStatefulVisitorWithExpr creates a visitor that visits both relations and expressions
func newStatefulVisitorWithExpr[S any](
	initial S,
	visitRel func(S, *proto.Rel) S,
	visitExpr func(S, *proto.Expression) S,
) *statefulVisitor[S] {
	return &statefulVisitor[S]{
		state:     initial,
		visitRel:  visitRel,
		visitExpr: visitExpr,
	}
}

func (v *statefulVisitor[S]) VisitRel(rel *proto.Rel) {
	if v.visitRel != nil {
		v.state = v.visitRel(v.state, rel)
	}
}

func (v *statefulVisitor[S]) VisitExpr(expr *proto.Expression) {
	if v.visitExpr != nil {
		v.state = v.visitExpr(v.state, expr)
	}
}

func (v *statefulVisitor[S]) Result() S {
	return v.state
}
