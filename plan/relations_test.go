package plan

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/expr"
	"testing"
)

func TestProjectRel_Copy(t *testing.T) {
	p := &ProjectRel{
		input: &VirtualTableReadRel{values: []expr.StructLiteralValue{[]expr.Literal{&expr.PrimitiveLiteral[int64]{Value: 5}}}},
		exprs: []expr.Expression{expr.NewPrimitiveLiteral(1.0, false), expr.NewPrimitiveLiteral(2.0, false)},
	}

	newInputs := []Rel{&VirtualTableReadRel{values: []expr.StructLiteralValue{[]expr.Literal{&expr.PrimitiveLiteral[int64]{Value: 6}}}}}
	got, err := p.Copy(newInputs...)
	if !assert.NoError(t, err, fmt.Sprintf("Copy(%v)", newInputs)) {
		return
	}
	want := &ProjectRel{
		input: newInputs[0],
		exprs: p.exprs,
	}
	assert.Equalf(t, want, got, "Copy(%v)", newInputs)

	// Copy with the same inputs and rewriteFunc is a no-op.
	sameInputs := p.GetInputs()
	got1, err := p.CopyWithExpressionRewrite(func(e expr.Expression) (expr.Expression, error) {
		return e, nil
	}, sameInputs...)
	assert.NoError(t, err)
	assert.Nil(t, got1)

	// Copy with different inputs and rewriteFunc is a no-op.
	got1, err = p.CopyWithExpressionRewrite(func(e expr.Expression) (expr.Expression, error) {
		return e, nil
	}, newInputs...)
	want1 := &ProjectRel{
		input: newInputs[0],
		exprs: p.exprs,
	}
	assert.NoError(t, err)
	assert.Equal(t, want1, got1)

	// Copy with different inputs and rewriteFunc changes all expressions.
	got2, err := p.CopyWithExpressionRewrite(func(e expr.Expression) (expr.Expression, error) {
		return expr.NewPrimitiveLiteral(9.0, false), nil
	}, newInputs...)
	want2 := &ProjectRel{
		input: newInputs[0],
		exprs: []expr.Expression{expr.NewPrimitiveLiteral(9.0, false), expr.NewPrimitiveLiteral(9.0, false)},
	}
	assert.NoError(t, err)
	assert.Equal(t, want2, got2)
}

func TestJoinRel_Copy(t *testing.T) {
	p := &JoinRel{
		left:           &VirtualTableReadRel{values: []expr.StructLiteralValue{[]expr.Literal{&expr.PrimitiveLiteral[int64]{Value: 1}}}},
		right:          &VirtualTableReadRel{values: []expr.StructLiteralValue{[]expr.Literal{&expr.PrimitiveLiteral[int64]{Value: 2}}}},
		joinType:       JoinTypeInner,
		expr:           expr.NewPrimitiveLiteral(true, false),
		postJoinFilter: expr.NewPrimitiveLiteral(true, false),
	}

	newInputs := []Rel{
		&VirtualTableReadRel{values: []expr.StructLiteralValue{[]expr.Literal{&expr.PrimitiveLiteral[int64]{Value: 6}}}},
		&VirtualTableReadRel{values: []expr.StructLiteralValue{[]expr.Literal{&expr.PrimitiveLiteral[int64]{Value: 7}}}},
	}
	got, err := p.Copy(newInputs...)
	if !assert.NoError(t, err, fmt.Sprintf("Copy(%v)", newInputs)) {
		return
	}
	want := &JoinRel{
		left:           newInputs[0],
		right:          newInputs[1],
		joinType:       p.joinType,
		expr:           p.expr,
		postJoinFilter: p.postJoinFilter,
	}
	assert.Equalf(t, want, got, "Copy(%v)", newInputs)

	// Copy with the same inputs and rewriteFunc is a no-op.
	sameInputs := p.GetInputs()
	got1, err := p.CopyWithExpressionRewrite(func(e expr.Expression) (expr.Expression, error) {
		return e, nil
	}, sameInputs...)
	assert.NoError(t, err)
	assert.Nil(t, got1)

	// Copy with different inputs and rewriteFunc is a no-op.
	got1, err = p.CopyWithExpressionRewrite(func(e expr.Expression) (expr.Expression, error) {
		return e, nil
	}, newInputs...)
	want1 := &JoinRel{
		left:           newInputs[0],
		right:          newInputs[1],
		joinType:       p.joinType,
		expr:           p.expr,
		postJoinFilter: p.postJoinFilter,
	}
	assert.NoError(t, err)
	assert.Equal(t, want1, got1)

	// Copy with different inputs and rewriteFunc changes the condition.
	got2, err := p.CopyWithExpressionRewrite(func(e expr.Expression) (expr.Expression, error) {
		return expr.NewPrimitiveLiteral(false, false), nil
	}, newInputs...)
	want2 := &JoinRel{
		left:           newInputs[0],
		right:          newInputs[1],
		joinType:       p.joinType,
		expr:           expr.NewPrimitiveLiteral(false, false),
		postJoinFilter: expr.NewPrimitiveLiteral(false, false),
	}
	assert.NoError(t, err)
	assert.Equal(t, want2, got2)
}
