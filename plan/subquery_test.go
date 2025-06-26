// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v4/expr"
	"github.com/substrait-io/substrait-go/v4/extensions"
	"github.com/substrait-io/substrait-go/v4/plan"
	"github.com/substrait-io/substrait-go/v4/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
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

func TestScalarSubquery(t *testing.T) {
	// Create a simple mock relation that returns one column of type i32
	mockRel := createMockReadRel()

	subquery := plan.NewScalarSubquery(mockRel)

	// Test basic properties
	assert.True(t, subquery.IsScalar())
	assert.Equal(t, "scalar", subquery.GetSubqueryType())
	assert.Contains(t, subquery.String(), "SCALAR_SUBQUERY")

	// Test type inference
	expectedType := &types.Int32Type{}
	assert.True(t, expectedType.Equals(subquery.GetType()))

	// Test protobuf conversion
	proto := subquery.ToProto()
	require.NotNil(t, proto)
	require.NotNil(t, proto.GetSubquery())
	require.NotNil(t, proto.GetSubquery().GetScalar())
}

func TestScalarSubqueryValidConstruction(t *testing.T) {
	mockRel := createMockReadRel()

	// Test with valid relation
	validSubquery := plan.NewScalarSubquery(mockRel)
	assert.NotNil(t, validSubquery)
	assert.NotNil(t, validSubquery.Input)
	assert.Equal(t, mockRel, validSubquery.Input)

	// Test with nil relation - should create but with nil relation
	nilSubquery := plan.NewScalarSubquery(nil)
	assert.NotNil(t, nilSubquery)
	assert.Nil(t, nilSubquery.Input)

	// Test protobuf conversion with valid relation
	proto := validSubquery.ToProto()
	require.NotNil(t, proto)
	require.NotNil(t, proto.GetSubquery())
	require.NotNil(t, proto.GetSubquery().GetScalar())
	require.NotNil(t, proto.GetSubquery().GetScalar().GetInput())
}

func TestInPredicateSubquery(t *testing.T) {
	// Create mock expressions and relation
	needle := expr.NewPrimitiveLiteral(int32(42), false)
	mockRel := createMockReadRel()

	subquery := plan.NewInPredicateSubquery([]expr.Expression{needle}, mockRel)

	// Test basic properties
	assert.True(t, subquery.IsScalar())
	assert.Equal(t, "in_predicate", subquery.GetSubqueryType())
	assert.Contains(t, subquery.String(), "IN")

	// Test type inference - should return boolean
	expectedType := &types.BooleanType{Nullability: types.NullabilityRequired}
	assert.True(t, expectedType.Equals(subquery.GetType()))

	// Test protobuf conversion
	proto := subquery.ToProto()
	require.NotNil(t, proto)
	require.NotNil(t, proto.GetSubquery())
	require.NotNil(t, proto.GetSubquery().GetInPredicate())
}

func TestInPredicateSubqueryValidConstruction(t *testing.T) {
	needle := expr.NewPrimitiveLiteral(int32(42), false)
	mockRel := createMockReadRel()

	// Test with nil needles - should create but with nil needles
	invalidSubquery1 := plan.NewInPredicateSubquery(nil, mockRel)
	assert.NotNil(t, invalidSubquery1)
	assert.Nil(t, invalidSubquery1.Needles)

	// Test with empty needles - should create with empty slice
	invalidSubquery2 := plan.NewInPredicateSubquery([]expr.Expression{}, mockRel)
	assert.NotNil(t, invalidSubquery2)
	assert.Empty(t, invalidSubquery2.Needles)

	// Test with nil needle in slice - should create but contain nil
	invalidSubquery3 := plan.NewInPredicateSubquery([]expr.Expression{nil}, mockRel)
	assert.NotNil(t, invalidSubquery3)
	assert.Len(t, invalidSubquery3.Needles, 1)
	assert.Nil(t, invalidSubquery3.Needles[0])

	// Test with nil relation - should create but with nil haystack
	invalidSubquery4 := plan.NewInPredicateSubquery([]expr.Expression{needle}, nil)
	assert.NotNil(t, invalidSubquery4)
	assert.Nil(t, invalidSubquery4.Haystack)

	// Test with valid inputs - should be properly constructed
	validSubquery := plan.NewInPredicateSubquery([]expr.Expression{needle}, mockRel)
	assert.NotNil(t, validSubquery)
	assert.NotEmpty(t, validSubquery.Needles)
	assert.NotNil(t, validSubquery.Haystack)

	// Test protobuf conversion doesn't have UNSPECIFIED values
	proto := validSubquery.ToProto()
	require.NotNil(t, proto)
	require.NotNil(t, proto.GetSubquery())
	require.NotNil(t, proto.GetSubquery().GetInPredicate())
	require.NotNil(t, proto.GetSubquery().GetInPredicate().GetHaystack())
	require.Len(t, proto.GetSubquery().GetInPredicate().GetNeedles(), 1)
	require.NotNil(t, proto.GetSubquery().GetInPredicate().GetNeedles()[0])
}

func TestSetPredicateSubquery(t *testing.T) {
	// Create mock relation
	mockRel := createMockReadRel()

	subquery := plan.NewSetPredicateSubquery(
		proto.Expression_Subquery_SetPredicate_PREDICATE_OP_EXISTS,
		mockRel,
	)

	// Test basic properties
	assert.True(t, subquery.IsScalar())
	assert.Equal(t, "set_predicate", subquery.GetSubqueryType())
	assert.Contains(t, subquery.String(), "EXISTS")

	// Test type inference - should return boolean
	expectedType := &types.BooleanType{Nullability: types.NullabilityRequired}
	assert.True(t, expectedType.Equals(subquery.GetType()))

	// Test protobuf conversion
	proto := subquery.ToProto()
	require.NotNil(t, proto)
	require.NotNil(t, proto.GetSubquery())
	require.NotNil(t, proto.GetSubquery().GetSetPredicate())
}

func TestSetPredicateSubqueryValidConstruction(t *testing.T) {
	mockRel := createMockReadRel()

	// Test with EXISTS operation
	existsSubquery := plan.NewSetPredicateSubquery(
		proto.Expression_Subquery_SetPredicate_PREDICATE_OP_EXISTS,
		mockRel,
	)
	assert.NotNil(t, existsSubquery)
	assert.Equal(t, proto.Expression_Subquery_SetPredicate_PREDICATE_OP_EXISTS, existsSubquery.Operation)
	assert.Equal(t, mockRel, existsSubquery.Tuples)

	// Test with UNIQUE operation
	uniqueSubquery := plan.NewSetPredicateSubquery(
		proto.Expression_Subquery_SetPredicate_PREDICATE_OP_UNIQUE,
		mockRel,
	)
	assert.NotNil(t, uniqueSubquery)
	assert.Equal(t, proto.Expression_Subquery_SetPredicate_PREDICATE_OP_UNIQUE, uniqueSubquery.Operation)
	assert.Equal(t, mockRel, uniqueSubquery.Tuples)

	// Test with UNSPECIFIED operation
	unspecifiedSubquery := plan.NewSetPredicateSubquery(
		proto.Expression_Subquery_SetPredicate_PREDICATE_OP_UNSPECIFIED,
		mockRel,
	)
	assert.NotNil(t, unspecifiedSubquery)
	assert.Equal(t, proto.Expression_Subquery_SetPredicate_PREDICATE_OP_UNSPECIFIED, unspecifiedSubquery.Operation)

	// Test with nil relation
	nilRelSubquery := plan.NewSetPredicateSubquery(
		proto.Expression_Subquery_SetPredicate_PREDICATE_OP_EXISTS,
		nil,
	)
	assert.NotNil(t, nilRelSubquery)
	assert.Equal(t, proto.Expression_Subquery_SetPredicate_PREDICATE_OP_EXISTS, nilRelSubquery.Operation)
	assert.Nil(t, nilRelSubquery.Tuples)

	// Test protobuf conversion with valid arguments
	protoMsg := existsSubquery.ToProto()
	require.NotNil(t, protoMsg)
	require.NotNil(t, protoMsg.GetSubquery())
	require.NotNil(t, protoMsg.GetSubquery().GetSetPredicate())
	assert.Equal(t, proto.Expression_Subquery_SetPredicate_PREDICATE_OP_EXISTS,
		protoMsg.GetSubquery().GetSetPredicate().GetPredicateOp())
}

func TestSetComparisonSubquery(t *testing.T) {
	// Create mock expression and relation
	left := expr.NewPrimitiveLiteral(int32(42), false)
	mockRel := createMockReadRel()

	subquery := plan.NewSetComparisonSubquery(
		proto.Expression_Subquery_SetComparison_REDUCTION_OP_ANY,
		proto.Expression_Subquery_SetComparison_COMPARISON_OP_EQ,
		left,
		mockRel,
	)

	// Test basic properties
	assert.True(t, subquery.IsScalar())
	assert.Equal(t, "set_comparison", subquery.GetSubqueryType())
	assert.Contains(t, subquery.String(), "ANY")

	// Test type inference - should return boolean
	expectedType := &types.BooleanType{Nullability: types.NullabilityRequired}
	assert.True(t, expectedType.Equals(subquery.GetType()))

	// Test protobuf conversion
	proto := subquery.ToProto()
	require.NotNil(t, proto)
	require.NotNil(t, proto.GetSubquery())
	require.NotNil(t, proto.GetSubquery().GetSetComparison())
}

func TestSetComparisonSubqueryValidConstruction(t *testing.T) {
	left := expr.NewPrimitiveLiteral(int32(42), false)
	mockRel := createMockReadRel()

	// Test with ANY/EQ combination
	anyEqSubquery := plan.NewSetComparisonSubquery(
		proto.Expression_Subquery_SetComparison_REDUCTION_OP_ANY,
		proto.Expression_Subquery_SetComparison_COMPARISON_OP_EQ,
		left,
		mockRel,
	)
	assert.NotNil(t, anyEqSubquery)
	assert.Equal(t, proto.Expression_Subquery_SetComparison_REDUCTION_OP_ANY, anyEqSubquery.ReductionOp)
	assert.Equal(t, proto.Expression_Subquery_SetComparison_COMPARISON_OP_EQ, anyEqSubquery.ComparisonOp)
	assert.Equal(t, left, anyEqSubquery.Left)
	assert.Equal(t, mockRel, anyEqSubquery.Right)

	// Test with ALL/NE combination
	allNeSubquery := plan.NewSetComparisonSubquery(
		proto.Expression_Subquery_SetComparison_REDUCTION_OP_ALL,
		proto.Expression_Subquery_SetComparison_COMPARISON_OP_NE,
		left,
		mockRel,
	)
	assert.NotNil(t, allNeSubquery)
	assert.Equal(t, proto.Expression_Subquery_SetComparison_REDUCTION_OP_ALL, allNeSubquery.ReductionOp)
	assert.Equal(t, proto.Expression_Subquery_SetComparison_COMPARISON_OP_NE, allNeSubquery.ComparisonOp)

	// Test with different comparison operations
	comparisonOps := []proto.Expression_Subquery_SetComparison_ComparisonOp{
		proto.Expression_Subquery_SetComparison_COMPARISON_OP_EQ,
		proto.Expression_Subquery_SetComparison_COMPARISON_OP_NE,
		proto.Expression_Subquery_SetComparison_COMPARISON_OP_LT,
		proto.Expression_Subquery_SetComparison_COMPARISON_OP_LE,
		proto.Expression_Subquery_SetComparison_COMPARISON_OP_GT,
		proto.Expression_Subquery_SetComparison_COMPARISON_OP_GE,
	}

	for _, compOp := range comparisonOps {
		subquery := plan.NewSetComparisonSubquery(
			proto.Expression_Subquery_SetComparison_REDUCTION_OP_ANY,
			compOp,
			left,
			mockRel,
		)
		assert.NotNil(t, subquery)
		assert.Equal(t, compOp, subquery.ComparisonOp)
	}

	// Test with UNSPECIFIED operations
	unspecifiedSubquery := plan.NewSetComparisonSubquery(
		proto.Expression_Subquery_SetComparison_REDUCTION_OP_UNSPECIFIED,
		proto.Expression_Subquery_SetComparison_COMPARISON_OP_UNSPECIFIED,
		left,
		mockRel,
	)
	assert.NotNil(t, unspecifiedSubquery)
	assert.Equal(t, proto.Expression_Subquery_SetComparison_REDUCTION_OP_UNSPECIFIED, unspecifiedSubquery.ReductionOp)
	assert.Equal(t, proto.Expression_Subquery_SetComparison_COMPARISON_OP_UNSPECIFIED, unspecifiedSubquery.ComparisonOp)

	// Test with nil left expression
	nilLeftSubquery := plan.NewSetComparisonSubquery(
		proto.Expression_Subquery_SetComparison_REDUCTION_OP_ANY,
		proto.Expression_Subquery_SetComparison_COMPARISON_OP_EQ,
		nil,
		mockRel,
	)
	assert.NotNil(t, nilLeftSubquery)
	assert.Nil(t, nilLeftSubquery.Left)
	assert.Equal(t, mockRel, nilLeftSubquery.Right)

	// Test with nil right relation
	nilRightSubquery := plan.NewSetComparisonSubquery(
		proto.Expression_Subquery_SetComparison_REDUCTION_OP_ANY,
		proto.Expression_Subquery_SetComparison_COMPARISON_OP_EQ,
		left,
		nil,
	)
	assert.NotNil(t, nilRightSubquery)
	assert.Equal(t, left, nilRightSubquery.Left)
	assert.Nil(t, nilRightSubquery.Right)

	// Test with both nil
	bothNilSubquery := plan.NewSetComparisonSubquery(
		proto.Expression_Subquery_SetComparison_REDUCTION_OP_ANY,
		proto.Expression_Subquery_SetComparison_COMPARISON_OP_EQ,
		nil,
		nil,
	)
	assert.NotNil(t, bothNilSubquery)
	assert.Nil(t, bothNilSubquery.Left)
	assert.Nil(t, bothNilSubquery.Right)

	// Test protobuf conversion with valid arguments
	protoMsg := anyEqSubquery.ToProto()
	require.NotNil(t, protoMsg)
	require.NotNil(t, protoMsg.GetSubquery())
	require.NotNil(t, protoMsg.GetSubquery().GetSetComparison())
	assert.Equal(t, proto.Expression_Subquery_SetComparison_REDUCTION_OP_ANY,
		protoMsg.GetSubquery().GetSetComparison().GetReductionOp())
	assert.Equal(t, proto.Expression_Subquery_SetComparison_COMPARISON_OP_EQ,
		protoMsg.GetSubquery().GetSetComparison().GetComparisonOp())
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

func TestHandleSubqueryFromProto(t *testing.T) {
	// Create base schema for testing
	baseSchema := &types.RecordType{}

	// Create proper extension registry
	baseRegistry := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
	registry := &plan.ExpressionResolver{Resolver: baseRegistry}

	// Create a mock relation for testing
	mockRel := createMockReadRel()
	mockRelProto := mockRel.ToProto()

	t.Run("ScalarSubquery", func(t *testing.T) {
		subqueryProto := &proto.Expression_Subquery{
			SubqueryType: &proto.Expression_Subquery_Scalar_{
				Scalar: &proto.Expression_Subquery_Scalar{
					Input: mockRelProto,
				},
			},
		}

		result, err := registry.HandleSubqueryFromProto(subqueryProto, baseSchema, baseRegistry)
		require.NoError(t, err)

		scalarSubquery, ok := result.(*plan.ScalarSubquery)
		require.True(t, ok)
		assert.NotNil(t, scalarSubquery.Input)
		assert.Equal(t, "scalar", scalarSubquery.GetSubqueryType())
	})

	t.Run("InPredicateSubquery", func(t *testing.T) {
		needleProto := expr.NewPrimitiveLiteral(int32(42), false).ToProto()

		subqueryProto := &proto.Expression_Subquery{
			SubqueryType: &proto.Expression_Subquery_InPredicate_{
				InPredicate: &proto.Expression_Subquery_InPredicate{
					Needles:  []*proto.Expression{needleProto},
					Haystack: mockRelProto,
				},
			},
		}

		result, err := registry.HandleSubqueryFromProto(subqueryProto, baseSchema, baseRegistry)
		require.NoError(t, err)

		inPredicateSubquery, ok := result.(*plan.InPredicateSubquery)
		require.True(t, ok)
		assert.Len(t, inPredicateSubquery.Needles, 1)
		assert.NotNil(t, inPredicateSubquery.Haystack)
		assert.Equal(t, "in_predicate", inPredicateSubquery.GetSubqueryType())
	})

	t.Run("InPredicateSubquery_MultipleNeedles", func(t *testing.T) {
		needle1Proto := expr.NewPrimitiveLiteral(int32(42), false).ToProto()
		needle2Proto := expr.NewPrimitiveLiteral(int32(99), false).ToProto()

		subqueryProto := &proto.Expression_Subquery{
			SubqueryType: &proto.Expression_Subquery_InPredicate_{
				InPredicate: &proto.Expression_Subquery_InPredicate{
					Needles:  []*proto.Expression{needle1Proto, needle2Proto},
					Haystack: mockRelProto,
				},
			},
		}

		result, err := registry.HandleSubqueryFromProto(subqueryProto, baseSchema, baseRegistry)
		require.NoError(t, err)

		inPredicateSubquery, ok := result.(*plan.InPredicateSubquery)
		require.True(t, ok)
		assert.Len(t, inPredicateSubquery.Needles, 2)
		assert.NotNil(t, inPredicateSubquery.Haystack)
	})

	t.Run("SetPredicateSubquery_EXISTS", func(t *testing.T) {
		subqueryProto := &proto.Expression_Subquery{
			SubqueryType: &proto.Expression_Subquery_SetPredicate_{
				SetPredicate: &proto.Expression_Subquery_SetPredicate{
					PredicateOp: proto.Expression_Subquery_SetPredicate_PREDICATE_OP_EXISTS,
					Tuples:      mockRelProto,
				},
			},
		}

		result, err := registry.HandleSubqueryFromProto(subqueryProto, baseSchema, baseRegistry)
		require.NoError(t, err)

		setPredicateSubquery, ok := result.(*plan.SetPredicateSubquery)
		require.True(t, ok)
		assert.Equal(t, proto.Expression_Subquery_SetPredicate_PREDICATE_OP_EXISTS, setPredicateSubquery.Operation)
		assert.NotNil(t, setPredicateSubquery.Tuples)
		assert.Equal(t, "set_predicate", setPredicateSubquery.GetSubqueryType())
	})

	t.Run("SetPredicateSubquery_UNIQUE", func(t *testing.T) {
		subqueryProto := &proto.Expression_Subquery{
			SubqueryType: &proto.Expression_Subquery_SetPredicate_{
				SetPredicate: &proto.Expression_Subquery_SetPredicate{
					PredicateOp: proto.Expression_Subquery_SetPredicate_PREDICATE_OP_UNIQUE,
					Tuples:      mockRelProto,
				},
			},
		}

		result, err := registry.HandleSubqueryFromProto(subqueryProto, baseSchema, baseRegistry)
		require.NoError(t, err)

		setPredicateSubquery, ok := result.(*plan.SetPredicateSubquery)
		require.True(t, ok)
		assert.Equal(t, proto.Expression_Subquery_SetPredicate_PREDICATE_OP_UNIQUE, setPredicateSubquery.Operation)
		assert.NotNil(t, setPredicateSubquery.Tuples)
	})

	t.Run("SetComparisonSubquery_ANY_EQ", func(t *testing.T) {
		leftExprProto := expr.NewPrimitiveLiteral(int32(42), false).ToProto()

		subqueryProto := &proto.Expression_Subquery{
			SubqueryType: &proto.Expression_Subquery_SetComparison_{
				SetComparison: &proto.Expression_Subquery_SetComparison{
					ReductionOp:  proto.Expression_Subquery_SetComparison_REDUCTION_OP_ANY,
					ComparisonOp: proto.Expression_Subquery_SetComparison_COMPARISON_OP_EQ,
					Left:         leftExprProto,
					Right:        mockRelProto,
				},
			},
		}

		result, err := registry.HandleSubqueryFromProto(subqueryProto, baseSchema, baseRegistry)
		require.NoError(t, err)

		setComparisonSubquery, ok := result.(*plan.SetComparisonSubquery)
		require.True(t, ok)
		assert.Equal(t, proto.Expression_Subquery_SetComparison_REDUCTION_OP_ANY, setComparisonSubquery.ReductionOp)
		assert.Equal(t, proto.Expression_Subquery_SetComparison_COMPARISON_OP_EQ, setComparisonSubquery.ComparisonOp)
		assert.NotNil(t, setComparisonSubquery.Left)
		assert.NotNil(t, setComparisonSubquery.Right)
		assert.Equal(t, "set_comparison", setComparisonSubquery.GetSubqueryType())
	})

	t.Run("SetComparisonSubquery_ALL_LT", func(t *testing.T) {
		leftExprProto := expr.NewPrimitiveLiteral(int32(50), false).ToProto()

		subqueryProto := &proto.Expression_Subquery{
			SubqueryType: &proto.Expression_Subquery_SetComparison_{
				SetComparison: &proto.Expression_Subquery_SetComparison{
					ReductionOp:  proto.Expression_Subquery_SetComparison_REDUCTION_OP_ALL,
					ComparisonOp: proto.Expression_Subquery_SetComparison_COMPARISON_OP_LT,
					Left:         leftExprProto,
					Right:        mockRelProto,
				},
			},
		}

		result, err := registry.HandleSubqueryFromProto(subqueryProto, baseSchema, baseRegistry)
		require.NoError(t, err)

		setComparisonSubquery, ok := result.(*plan.SetComparisonSubquery)
		require.True(t, ok)
		assert.Equal(t, proto.Expression_Subquery_SetComparison_REDUCTION_OP_ALL, setComparisonSubquery.ReductionOp)
		assert.Equal(t, proto.Expression_Subquery_SetComparison_COMPARISON_OP_LT, setComparisonSubquery.ComparisonOp)
	})

	t.Run("UnknownSubqueryType", func(t *testing.T) {
		// Create a subquery with no subquery type set (nil SubqueryType)
		subqueryProto := &proto.Expression_Subquery{
			SubqueryType: nil,
		}

		result, err := registry.HandleSubqueryFromProto(subqueryProto, baseSchema, baseRegistry)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "unknown subquery type")
	})
}

func TestHandleSubqueryFromProtoErrors(t *testing.T) {
	// Create base schema for testing
	baseSchema := &types.RecordType{}

	// Create proper extension registry
	baseRegistry := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
	registry := &plan.ExpressionResolver{Resolver: baseRegistry}

	t.Run("ScalarSubquery_RelFromProtoError", func(t *testing.T) {
		// Create invalid relation proto that will cause RelFromProto to fail
		invalidRelProto := &proto.Rel{
			RelType: nil, // This will cause an error in RelFromProto
		}

		subqueryProto := &proto.Expression_Subquery{
			SubqueryType: &proto.Expression_Subquery_Scalar_{
				Scalar: &proto.Expression_Subquery_Scalar{
					Input: invalidRelProto,
				},
			},
		}

		result, err := registry.HandleSubqueryFromProto(subqueryProto, baseSchema, baseRegistry)
		require.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("InPredicateSubquery_NeedleExprError", func(t *testing.T) {
		mockRel := createMockReadRel()
		mockRelProto := mockRel.ToProto()

		// Create invalid expression proto that will cause ExprFromProto to fail
		invalidExprProto := &proto.Expression{
			RexType: nil, // This will cause an error in ExprFromProto
		}

		subqueryProto := &proto.Expression_Subquery{
			SubqueryType: &proto.Expression_Subquery_InPredicate_{
				InPredicate: &proto.Expression_Subquery_InPredicate{
					Needles:  []*proto.Expression{invalidExprProto},
					Haystack: mockRelProto,
				},
			},
		}

		result, err := registry.HandleSubqueryFromProto(subqueryProto, baseSchema, baseRegistry)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "error parsing needle 0 in IN predicate")
	})

	t.Run("InPredicateSubquery_HaystackRelError", func(t *testing.T) {
		needleProto := expr.NewPrimitiveLiteral(int32(42), false).ToProto()

		// Create invalid relation proto that will cause RelFromProto to fail
		invalidRelProto := &proto.Rel{
			RelType: nil, // This will cause an error in RelFromProto
		}

		subqueryProto := &proto.Expression_Subquery{
			SubqueryType: &proto.Expression_Subquery_InPredicate_{
				InPredicate: &proto.Expression_Subquery_InPredicate{
					Needles:  []*proto.Expression{needleProto},
					Haystack: invalidRelProto,
				},
			},
		}

		result, err := registry.HandleSubqueryFromProto(subqueryProto, baseSchema, baseRegistry)
		require.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("SetPredicateSubquery_TuplesRelError", func(t *testing.T) {
		// Create invalid relation proto that will cause RelFromProto to fail
		invalidRelProto := &proto.Rel{
			RelType: nil, // This will cause an error in RelFromProto
		}

		subqueryProto := &proto.Expression_Subquery{
			SubqueryType: &proto.Expression_Subquery_SetPredicate_{
				SetPredicate: &proto.Expression_Subquery_SetPredicate{
					PredicateOp: proto.Expression_Subquery_SetPredicate_PREDICATE_OP_EXISTS,
					Tuples:      invalidRelProto,
				},
			},
		}

		result, err := registry.HandleSubqueryFromProto(subqueryProto, baseSchema, baseRegistry)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "error parsing tuples in set predicate")
	})

	t.Run("SetComparisonSubquery_LeftExprError", func(t *testing.T) {
		mockRel := createMockReadRel()
		mockRelProto := mockRel.ToProto()

		// Create invalid expression proto that will cause ExprFromProto to fail
		invalidExprProto := &proto.Expression{
			RexType: nil, // This will cause an error in ExprFromProto
		}

		subqueryProto := &proto.Expression_Subquery{
			SubqueryType: &proto.Expression_Subquery_SetComparison_{
				SetComparison: &proto.Expression_Subquery_SetComparison{
					ReductionOp:  proto.Expression_Subquery_SetComparison_REDUCTION_OP_ANY,
					ComparisonOp: proto.Expression_Subquery_SetComparison_COMPARISON_OP_EQ,
					Left:         invalidExprProto,
					Right:        mockRelProto,
				},
			},
		}

		result, err := registry.HandleSubqueryFromProto(subqueryProto, baseSchema, baseRegistry)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "error parsing left expression in set comparison")
	})

	t.Run("SetComparisonSubquery_RightRelError", func(t *testing.T) {
		leftExprProto := expr.NewPrimitiveLiteral(int32(42), false).ToProto()

		// Create invalid relation proto that will cause RelFromProto to fail
		invalidRelProto := &proto.Rel{
			RelType: nil, // This will cause an error in RelFromProto
		}

		subqueryProto := &proto.Expression_Subquery{
			SubqueryType: &proto.Expression_Subquery_SetComparison_{
				SetComparison: &proto.Expression_Subquery_SetComparison{
					ReductionOp:  proto.Expression_Subquery_SetComparison_REDUCTION_OP_ANY,
					ComparisonOp: proto.Expression_Subquery_SetComparison_COMPARISON_OP_EQ,
					Left:         leftExprProto,
					Right:        invalidRelProto,
				},
			},
		}

		result, err := registry.HandleSubqueryFromProto(subqueryProto, baseSchema, baseRegistry)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "error parsing right relation in set comparison")
	})
}

func TestHandleSubqueryFromProtoEdgeCases(t *testing.T) {
	// Create base schema for testing
	baseSchema := &types.RecordType{}

	// Create proper extension registry
	baseRegistry := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
	registry := &plan.ExpressionResolver{Resolver: baseRegistry}

	// Create a mock relation for testing
	mockRel := createMockReadRel()
	mockRelProto := mockRel.ToProto()

	t.Run("InPredicateSubquery_EmptyNeedles", func(t *testing.T) {
		subqueryProto := &proto.Expression_Subquery{
			SubqueryType: &proto.Expression_Subquery_InPredicate_{
				InPredicate: &proto.Expression_Subquery_InPredicate{
					Needles:  []*proto.Expression{}, // Empty needles array
					Haystack: mockRelProto,
				},
			},
		}

		result, err := registry.HandleSubqueryFromProto(subqueryProto, baseSchema, baseRegistry)
		require.NoError(t, err)

		inPredicateSubquery, ok := result.(*plan.InPredicateSubquery)
		require.True(t, ok)
		assert.Len(t, inPredicateSubquery.Needles, 0)
		assert.NotNil(t, inPredicateSubquery.Haystack)
	})

	t.Run("InPredicateSubquery_MultipleNeedleErrors", func(t *testing.T) {
		validNeedleProto := expr.NewPrimitiveLiteral(int32(42), false).ToProto()
		invalidNeedleProto := &proto.Expression{
			RexType: nil, // This will cause an error in ExprFromProto
		}

		subqueryProto := &proto.Expression_Subquery{
			SubqueryType: &proto.Expression_Subquery_InPredicate_{
				InPredicate: &proto.Expression_Subquery_InPredicate{
					Needles:  []*proto.Expression{validNeedleProto, invalidNeedleProto},
					Haystack: mockRelProto,
				},
			},
		}

		result, err := registry.HandleSubqueryFromProto(subqueryProto, baseSchema, baseRegistry)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "error parsing needle 1 in IN predicate")
	})

	t.Run("SetComparisonSubquery_AllComparisonOperators", func(t *testing.T) {
		leftExprProto := expr.NewPrimitiveLiteral(int32(42), false).ToProto()

		comparisonOps := []proto.Expression_Subquery_SetComparison_ComparisonOp{
			proto.Expression_Subquery_SetComparison_COMPARISON_OP_EQ,
			proto.Expression_Subquery_SetComparison_COMPARISON_OP_NE,
			proto.Expression_Subquery_SetComparison_COMPARISON_OP_LT,
			proto.Expression_Subquery_SetComparison_COMPARISON_OP_LE,
			proto.Expression_Subquery_SetComparison_COMPARISON_OP_GT,
			proto.Expression_Subquery_SetComparison_COMPARISON_OP_GE,
		}

		for _, compOp := range comparisonOps {
			t.Run(compOp.String(), func(t *testing.T) {
				subqueryProto := &proto.Expression_Subquery{
					SubqueryType: &proto.Expression_Subquery_SetComparison_{
						SetComparison: &proto.Expression_Subquery_SetComparison{
							ReductionOp:  proto.Expression_Subquery_SetComparison_REDUCTION_OP_ANY,
							ComparisonOp: compOp,
							Left:         leftExprProto,
							Right:        mockRelProto,
						},
					},
				}

				result, err := registry.HandleSubqueryFromProto(subqueryProto, baseSchema, baseRegistry)
				require.NoError(t, err)

				setComparisonSubquery, ok := result.(*plan.SetComparisonSubquery)
				require.True(t, ok)
				assert.Equal(t, compOp, setComparisonSubquery.ComparisonOp)
			})
		}
	})

	t.Run("SetComparisonSubquery_AllReductionOperators", func(t *testing.T) {
		leftExprProto := expr.NewPrimitiveLiteral(int32(42), false).ToProto()

		reductionOps := []proto.Expression_Subquery_SetComparison_ReductionOp{
			proto.Expression_Subquery_SetComparison_REDUCTION_OP_ANY,
			proto.Expression_Subquery_SetComparison_REDUCTION_OP_ALL,
		}

		for _, redOp := range reductionOps {
			t.Run(redOp.String(), func(t *testing.T) {
				subqueryProto := &proto.Expression_Subquery{
					SubqueryType: &proto.Expression_Subquery_SetComparison_{
						SetComparison: &proto.Expression_Subquery_SetComparison{
							ReductionOp:  redOp,
							ComparisonOp: proto.Expression_Subquery_SetComparison_COMPARISON_OP_EQ,
							Left:         leftExprProto,
							Right:        mockRelProto,
						},
					},
				}

				result, err := registry.HandleSubqueryFromProto(subqueryProto, baseSchema, baseRegistry)
				require.NoError(t, err)

				setComparisonSubquery, ok := result.(*plan.SetComparisonSubquery)
				require.True(t, ok)
				assert.Equal(t, redOp, setComparisonSubquery.ReductionOp)
			})
		}
	})
}
