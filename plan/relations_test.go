package plan

import (
	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/expr"
	"testing"
)

func noOpRewrite(e expr.Expression) (expr.Expression, error) {
	return e, nil
}

func createVirtualTableReadRel(value int64) *VirtualTableReadRel {
	return &VirtualTableReadRel{values: []expr.StructLiteralValue{[]expr.Literal{&expr.PrimitiveLiteral[int64]{Value: value}}}}
}

func createPrimitiveLiteral(value float64) expr.Expression {
	return expr.NewPrimitiveLiteral(value, false)
}

func TestRelations_Copy(t *testing.T) {
	initialProjectRel := &ProjectRel{input: createVirtualTableReadRel(5), exprs: []expr.Expression{createPrimitiveLiteral(1.0), createPrimitiveLiteral(2.0)}}
	initialJoinRel := &JoinRel{left: createVirtualTableReadRel(1), right: createVirtualTableReadRel(2), joinType: JoinTypeInner, expr: expr.NewPrimitiveLiteral(true, false), postJoinFilter: expr.NewPrimitiveLiteral(true, false)}
	initialCrossRel := &CrossRel{left: createVirtualTableReadRel(1), right: createVirtualTableReadRel(2)}
	initialFilterRel := &FilterRel{input: createVirtualTableReadRel(1), cond: expr.NewPrimitiveLiteral(true, false)}
	localFileReadRel := &LocalFileReadRel{items: []FileOrFiles{{Path: "path"}},
		baseReadRel: baseReadRel{
			filter:           expr.NewPrimitiveLiteral(true, false),
			bestEffortFilter: expr.NewPrimitiveLiteral(true, false),
		},
	}
	initialSetRel := &SetRel{inputs: []Rel{createVirtualTableReadRel(1), createVirtualTableReadRel(2), createVirtualTableReadRel(3)}, op: SetOpUnionAll}
	initialExtensionLeafRel := &ExtensionLeafRel{}
	initialAggregateRel := &AggregateRel{input: createVirtualTableReadRel(1), groups: [][]expr.Expression{{createPrimitiveLiteral(1.0)}},
		measures: []AggRelMeasure{{filter: expr.NewPrimitiveLiteral(true, false)}}}

	type relationTestCase struct {
		name        string
		relation    Rel
		newInputs   []Rel
		rewriteFunc func(expr.Expression) (expr.Expression, error)
		expectedRel Rel
	}
	testCases := []relationTestCase{
		{
			name:        "ProjectRel Copy with new inputs",
			relation:    initialProjectRel,
			newInputs:   []Rel{createVirtualTableReadRel(6)},
			expectedRel: &ProjectRel{input: createVirtualTableReadRel(6), exprs: []expr.Expression{createPrimitiveLiteral(1.0), createPrimitiveLiteral(2.0)}},
		},
		{
			name:        "ProjectRel Copy with same inputs and noOpRewrite",
			relation:    initialProjectRel,
			newInputs:   []Rel{initialProjectRel.input},
			rewriteFunc: noOpRewrite,
			expectedRel: nil,
		},
		{
			name:        "ProjectRel Copy with new inputs and noOpRewrite",
			relation:    initialProjectRel,
			newInputs:   []Rel{createVirtualTableReadRel(6)},
			rewriteFunc: noOpRewrite,
			expectedRel: &ProjectRel{input: createVirtualTableReadRel(6), exprs: []expr.Expression{createPrimitiveLiteral(1.0), createPrimitiveLiteral(2.0)}},
		},
		{
			name:        "ProjectRel Copy with new inputs and rewriteFunc",
			relation:    initialProjectRel,
			newInputs:   []Rel{createVirtualTableReadRel(6)},
			rewriteFunc: func(e expr.Expression) (expr.Expression, error) { return createPrimitiveLiteral(9.0), nil },
			expectedRel: &ProjectRel{input: createVirtualTableReadRel(6), exprs: []expr.Expression{createPrimitiveLiteral(9.0), createPrimitiveLiteral(9.0)}},
		},
		{
			name:        "JoinRel Copy with new inputs",
			relation:    initialJoinRel,
			newInputs:   []Rel{createVirtualTableReadRel(6), createVirtualTableReadRel(7)},
			expectedRel: &JoinRel{left: createVirtualTableReadRel(6), right: createVirtualTableReadRel(7), joinType: JoinTypeInner, expr: expr.NewPrimitiveLiteral(true, false), postJoinFilter: expr.NewPrimitiveLiteral(true, false)},
		},
		{
			name:        "JoinRel Copy with same inputs and noOpRewrite",
			relation:    initialJoinRel,
			newInputs:   []Rel{initialJoinRel.left, initialJoinRel.right},
			rewriteFunc: noOpRewrite,
			expectedRel: nil,
		},
		{
			name:        "CrossRel Copy with new inputs",
			relation:    initialCrossRel,
			newInputs:   []Rel{createVirtualTableReadRel(6), createVirtualTableReadRel(7)},
			expectedRel: &CrossRel{left: createVirtualTableReadRel(6), right: createVirtualTableReadRel(7)},
		},
		{
			name:        "CrossRel Copy with same inputs and noOpRewrite",
			relation:    initialCrossRel,
			newInputs:   initialCrossRel.GetInputs(),
			rewriteFunc: noOpRewrite,
			expectedRel: nil,
		},
		{
			name:        "CrossRel Copy with new inputs and noOpRewrite",
			relation:    initialCrossRel,
			newInputs:   []Rel{createVirtualTableReadRel(6), createVirtualTableReadRel(7)},
			rewriteFunc: noOpRewrite,
			expectedRel: &CrossRel{left: createVirtualTableReadRel(6), right: createVirtualTableReadRel(7)},
		},
		{
			name:        "CrossRel Copy with new inputs and rewriteFunc",
			relation:    initialCrossRel,
			newInputs:   []Rel{createVirtualTableReadRel(6), createVirtualTableReadRel(7)},
			rewriteFunc: func(e expr.Expression) (expr.Expression, error) { return createPrimitiveLiteral(9.0), nil },
			expectedRel: &CrossRel{left: createVirtualTableReadRel(6), right: createVirtualTableReadRel(7)},
		},
		{
			name:        "FilterRel Copy with new inputs",
			relation:    initialFilterRel,
			newInputs:   []Rel{createVirtualTableReadRel(6)},
			expectedRel: &FilterRel{input: createVirtualTableReadRel(6), cond: expr.NewPrimitiveLiteral(true, false)},
		},
		{
			name:        "LocalFileReadRel Copy with new inputs",
			relation:    localFileReadRel,
			newInputs:   []Rel{},
			expectedRel: localFileReadRel,
		},
		{
			name:        "SetRel Copy with new inputs",
			relation:    initialSetRel,
			newInputs:   []Rel{createVirtualTableReadRel(6), createVirtualTableReadRel(7), createVirtualTableReadRel(8)},
			expectedRel: &SetRel{inputs: []Rel{createVirtualTableReadRel(6), createVirtualTableReadRel(7), createVirtualTableReadRel(8)}, op: SetOpUnionAll},
		},
		{
			name:        "ExtensionLeafRel Copy with new inputs",
			relation:    initialExtensionLeafRel,
			newInputs:   []Rel{},
			expectedRel: initialExtensionLeafRel,
		},
		{
			name:        "AggregateRel Copy with new inputs",
			relation:    initialAggregateRel,
			newInputs:   []Rel{createVirtualTableReadRel(6)},
			expectedRel: &AggregateRel{input: createVirtualTableReadRel(6), groups: initialAggregateRel.groups, measures: initialAggregateRel.measures},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var got Rel
			var err error
			if tc.rewriteFunc == nil {
				got, err = tc.relation.Copy(tc.newInputs...)
			} else {
				got, err = tc.relation.CopyWithExpressionRewrite(tc.rewriteFunc, tc.newInputs...)
			}
			if !assert.NoError(t, err) {
				return
			}
			assert.Equal(t, tc.expectedRel, got)
		})
	}
}
