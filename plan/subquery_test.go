// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/plan"
	"github.com/substrait-io/substrait-go/v9/types"
)

func createMockReadRel() plan.Rel {
	schema := types.NamedStruct{
		Names: []string{"col1"},
		Struct: types.StructType{
			Types: []types.Type{&types.Int32Type{}},
		},
	}
	return plan.NewBuilderDefault().NamedScan([]string{"test_table"}, schema)
}

func TestSubqueryVisit(t *testing.T) {
	// Test that Visit works correctly for InPredicateSubquery
	needle := expr.NewPrimitiveLiteral(int32(42), false)
	mockRel := createMockReadRel()

	subquery := plan.NewInPredicateSubquery([]expr.Expression{needle}, mockRel)

	// Visit function that replaces int32(42) with int32(100)
	visitFunc := func(e expr.Expression) expr.Expression {
		if lit, ok := e.(*expr.PrimitiveLiteral[int32]); ok && lit.Value == 42 {
			return expr.NewPrimitiveLiteral(int32(100), false)
		}
		return e
	}

	result := subquery.Visit(visitFunc)

	// The result should be a new InPredicateSubquery with the modified needle
	newSubquery, ok := result.(*plan.InPredicateSubquery)
	require.True(t, ok)

	// Check that the needle was changed
	newNeedle, ok := newSubquery.Needles[0].(*expr.PrimitiveLiteral[int32])
	require.True(t, ok)
	assert.Equal(t, int32(100), newNeedle.Value)
}

func TestScalarSubqueryEquals(t *testing.T) {
	mockRel1 := createMockReadRel()
	mockRel2 := createMockReadRel()

	subquery1 := plan.NewScalarSubquery(mockRel1)
	subquery2 := plan.NewScalarSubquery(mockRel1)
	subquery3 := plan.NewScalarSubquery(mockRel2)

	t.Run("SameInstance", func(t *testing.T) {
		assert.True(t, subquery1.Equals(subquery1))
	})

	t.Run("SameInput", func(t *testing.T) {
		assert.True(t, subquery1.Equals(subquery2))
	})

	t.Run("DifferentInput", func(t *testing.T) {
		assert.False(t, subquery1.Equals(subquery3))
	})

	t.Run("NilInput", func(t *testing.T) {
		nilSubquery1 := plan.NewScalarSubquery(nil)
		nilSubquery2 := plan.NewScalarSubquery(nil)
		assert.True(t, nilSubquery1.Equals(nilSubquery2))
		assert.False(t, subquery1.Equals(nilSubquery1))
		assert.False(t, nilSubquery1.Equals(subquery1))
	})

	t.Run("DifferentSubqueryType", func(t *testing.T) {
		needle := expr.NewPrimitiveLiteral(int32(42), false)
		inPredicate := plan.NewInPredicateSubquery([]expr.Expression{needle}, mockRel1)
		assert.False(t, subquery1.Equals(inPredicate))
	})

	t.Run("NilOther", func(t *testing.T) {
		assert.False(t, subquery1.Equals(nil))
	})

	t.Run("NonSubqueryExpression", func(t *testing.T) {
		literal := expr.NewPrimitiveLiteral(int32(42), false)
		assert.False(t, subquery1.Equals(literal))
	})
}

func TestInPredicateSubqueryEquals(t *testing.T) {
	mockRel1 := createMockReadRel()
	mockRel2 := createMockReadRel()

	needle1 := expr.NewPrimitiveLiteral(int32(42), false)
	needle2 := expr.NewPrimitiveLiteral(int32(99), false)
	needle3 := expr.NewPrimitiveLiteral(int32(42), false) // Same value as needle1

	subquery1 := plan.NewInPredicateSubquery([]expr.Expression{needle1}, mockRel1)
	subquery2 := plan.NewInPredicateSubquery([]expr.Expression{needle3}, mockRel1)          // Same needle value
	subquery3 := plan.NewInPredicateSubquery([]expr.Expression{needle2}, mockRel1)          // Different needle
	subquery4 := plan.NewInPredicateSubquery([]expr.Expression{needle1}, mockRel2)          // Different relation
	subquery5 := plan.NewInPredicateSubquery([]expr.Expression{needle1, needle2}, mockRel1) // More needles

	t.Run("SameInstance", func(t *testing.T) {
		assert.True(t, subquery1.Equals(subquery1))
	})

	t.Run("SameNeedleAndHaystack", func(t *testing.T) {
		assert.True(t, subquery1.Equals(subquery2))
	})

	t.Run("DifferentNeedle", func(t *testing.T) {
		assert.False(t, subquery1.Equals(subquery3))
	})

	t.Run("DifferentHaystack", func(t *testing.T) {
		assert.False(t, subquery1.Equals(subquery4))
	})

	t.Run("DifferentNeedleCount", func(t *testing.T) {
		assert.False(t, subquery1.Equals(subquery5))
	})

	t.Run("MultipleNeedles", func(t *testing.T) {
		multiNeedle1 := plan.NewInPredicateSubquery([]expr.Expression{needle1, needle2}, mockRel1)
		multiNeedle2 := plan.NewInPredicateSubquery([]expr.Expression{needle3, needle2}, mockRel1)
		multiNeedle3 := plan.NewInPredicateSubquery([]expr.Expression{needle2, needle1}, mockRel1) // Different order

		assert.True(t, multiNeedle1.Equals(multiNeedle2))
		assert.False(t, multiNeedle1.Equals(multiNeedle3)) // Order matters
	})

	t.Run("EmptyNeedles", func(t *testing.T) {
		emptyNeedles1 := plan.NewInPredicateSubquery([]expr.Expression{}, mockRel1)
		emptyNeedles2 := plan.NewInPredicateSubquery([]expr.Expression{}, mockRel1)
		assert.True(t, emptyNeedles1.Equals(emptyNeedles2))
		assert.False(t, subquery1.Equals(emptyNeedles1))
	})

	t.Run("NilNeedles", func(t *testing.T) {
		nilNeedles1 := plan.NewInPredicateSubquery(nil, mockRel1)
		nilNeedles2 := plan.NewInPredicateSubquery(nil, mockRel1)
		assert.True(t, nilNeedles1.Equals(nilNeedles2))
		assert.False(t, subquery1.Equals(nilNeedles1))
	})

	t.Run("DifferentSubqueryType", func(t *testing.T) {
		scalar := plan.NewScalarSubquery(mockRel1)
		assert.False(t, subquery1.Equals(scalar))
	})

	t.Run("NilOther", func(t *testing.T) {
		assert.False(t, subquery1.Equals(nil))
	})
}

func TestSetPredicateSubqueryEquals(t *testing.T) {
	mockRel1 := createMockReadRel()
	mockRel2 := createMockReadRel()

	existsSubquery1 := plan.NewSetPredicateSubquery(
		plan.SetPredicateOpExists,
		mockRel1,
	)
	existsSubquery2 := plan.NewSetPredicateSubquery(
		plan.SetPredicateOpExists,
		mockRel1,
	)
	uniqueSubquery := plan.NewSetPredicateSubquery(
		plan.SetPredicateOpUnique,
		mockRel1,
	)
	existsSubqueryDiffRel := plan.NewSetPredicateSubquery(
		plan.SetPredicateOpExists,
		mockRel2,
	)
	unspecifiedSubquery := plan.NewSetPredicateSubquery(
		plan.SetPredicateOpUnspecified,
		mockRel1,
	)

	t.Run("SameInstance", func(t *testing.T) {
		assert.True(t, existsSubquery1.Equals(existsSubquery1))
	})

	t.Run("SameOperationAndTuples", func(t *testing.T) {
		assert.True(t, existsSubquery1.Equals(existsSubquery2))
	})

	t.Run("DifferentOperation", func(t *testing.T) {
		assert.False(t, existsSubquery1.Equals(uniqueSubquery))
	})

	t.Run("DifferentTuples", func(t *testing.T) {
		assert.False(t, existsSubquery1.Equals(existsSubqueryDiffRel))
	})

	t.Run("UnspecifiedOperation", func(t *testing.T) {
		assert.False(t, existsSubquery1.Equals(unspecifiedSubquery))
	})

	t.Run("NilTuples", func(t *testing.T) {
		nilTuples1 := plan.NewSetPredicateSubquery(
			plan.SetPredicateOpExists,
			nil,
		)
		nilTuples2 := plan.NewSetPredicateSubquery(
			plan.SetPredicateOpExists,
			nil,
		)
		assert.True(t, nilTuples1.Equals(nilTuples2))
		assert.False(t, existsSubquery1.Equals(nilTuples1))
	})

	t.Run("DifferentSubqueryType", func(t *testing.T) {
		scalar := plan.NewScalarSubquery(mockRel1)
		assert.False(t, existsSubquery1.Equals(scalar))
	})

	t.Run("NilOther", func(t *testing.T) {
		assert.False(t, existsSubquery1.Equals(nil))
	})
}

func TestSetComparisonSubqueryEquals(t *testing.T) {
	mockRel1 := createMockReadRel()
	mockRel2 := createMockReadRel()

	left1 := expr.NewPrimitiveLiteral(int32(42), false)
	left2 := expr.NewPrimitiveLiteral(int32(42), false) // Same value
	left3 := expr.NewPrimitiveLiteral(int32(99), false) // Different value

	subquery1 := plan.NewSetComparisonSubquery(
		plan.SetComparisonReductionOpAny,
		plan.SetComparisonOpEq,
		left1,
		mockRel1,
	)
	subquery2 := plan.NewSetComparisonSubquery(
		plan.SetComparisonReductionOpAny,
		plan.SetComparisonOpEq,
		left2,
		mockRel1,
	)
	subqueryDiffReduction := plan.NewSetComparisonSubquery(
		plan.SetComparisonReductionOpAll,
		plan.SetComparisonOpEq,
		left1,
		mockRel1,
	)
	subqueryDiffComparison := plan.NewSetComparisonSubquery(
		plan.SetComparisonReductionOpAny,
		plan.SetComparisonOpNe,
		left1,
		mockRel1,
	)
	subqueryDiffLeft := plan.NewSetComparisonSubquery(
		plan.SetComparisonReductionOpAny,
		plan.SetComparisonOpEq,
		left3,
		mockRel1,
	)
	subqueryDiffRight := plan.NewSetComparisonSubquery(
		plan.SetComparisonReductionOpAny,
		plan.SetComparisonOpEq,
		left1,
		mockRel2,
	)

	t.Run("SameInstance", func(t *testing.T) {
		assert.True(t, subquery1.Equals(subquery1))
	})

	t.Run("SameAllFields", func(t *testing.T) {
		assert.True(t, subquery1.Equals(subquery2))
	})

	t.Run("DifferentReductionOp", func(t *testing.T) {
		assert.False(t, subquery1.Equals(subqueryDiffReduction))
	})

	t.Run("DifferentComparisonOp", func(t *testing.T) {
		assert.False(t, subquery1.Equals(subqueryDiffComparison))
	})

	t.Run("DifferentLeftExpression", func(t *testing.T) {
		assert.False(t, subquery1.Equals(subqueryDiffLeft))
	})

	t.Run("DifferentRightRelation", func(t *testing.T) {
		assert.False(t, subquery1.Equals(subqueryDiffRight))
	})

	t.Run("AllComparisonOperators", func(t *testing.T) {
		comparisonOps := []plan.SetComparisonOp{
			plan.SetComparisonOpEq,
			plan.SetComparisonOpNe,
			plan.SetComparisonOpLt,
			plan.SetComparisonOpLe,
			plan.SetComparisonOpGt,
			plan.SetComparisonOpGe,
		}

		for _, op := range comparisonOps {
			subquery := plan.NewSetComparisonSubquery(
				plan.SetComparisonReductionOpAny,
				op,
				left1,
				mockRel1,
			)
			sameSub := plan.NewSetComparisonSubquery(
				plan.SetComparisonReductionOpAny,
				op,
				left2,
				mockRel1,
			)
			assert.True(t, subquery.Equals(sameSub), "Failed for comparison op: %v", op)
		}
	})

	t.Run("AllReductionOperators", func(t *testing.T) {
		reductionOps := []plan.SetComparisonReductionOp{
			plan.SetComparisonReductionOpAny,
			plan.SetComparisonReductionOpAll,
		}

		for _, op := range reductionOps {
			subquery := plan.NewSetComparisonSubquery(
				op,
				plan.SetComparisonOpEq,
				left1,
				mockRel1,
			)
			sameSub := plan.NewSetComparisonSubquery(
				op,
				plan.SetComparisonOpEq,
				left2,
				mockRel1,
			)
			assert.True(t, subquery.Equals(sameSub), "Failed for reduction op: %v", op)
		}
	})

	t.Run("UnspecifiedOperations", func(t *testing.T) {
		unspecifiedSubquery := plan.NewSetComparisonSubquery(
			plan.SetComparisonReductionOpUnspecified,
			plan.SetComparisonOpUnspecified,
			left1,
			mockRel1,
		)
		assert.False(t, subquery1.Equals(unspecifiedSubquery))
	})

	t.Run("NilLeftExpression", func(t *testing.T) {
		nilLeft1 := plan.NewSetComparisonSubquery(
			plan.SetComparisonReductionOpAny,
			plan.SetComparisonOpEq,
			nil,
			mockRel1,
		)
		nilLeft2 := plan.NewSetComparisonSubquery(
			plan.SetComparisonReductionOpAny,
			plan.SetComparisonOpEq,
			nil,
			mockRel1,
		)
		assert.True(t, nilLeft1.Equals(nilLeft2))
		assert.False(t, subquery1.Equals(nilLeft1))
	})

	t.Run("NilRightRelation", func(t *testing.T) {
		nilRight1 := plan.NewSetComparisonSubquery(
			plan.SetComparisonReductionOpAny,
			plan.SetComparisonOpEq,
			left1,
			nil,
		)
		nilRight2 := plan.NewSetComparisonSubquery(
			plan.SetComparisonReductionOpAny,
			plan.SetComparisonOpEq,
			left2,
			nil,
		)
		assert.True(t, nilRight1.Equals(nilRight2))
		assert.False(t, subquery1.Equals(nilRight1))
	})

	t.Run("DifferentSubqueryType", func(t *testing.T) {
		scalar := plan.NewScalarSubquery(mockRel1)
		assert.False(t, subquery1.Equals(scalar))
	})

	t.Run("NilOther", func(t *testing.T) {
		assert.False(t, subquery1.Equals(nil))
	})
}
