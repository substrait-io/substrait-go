// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/plan"
	"github.com/substrait-io/substrait-go/v9/types"
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
	proto := ExprToProto(subquery)
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
	proto := ExprToProto(validSubquery)
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
	proto := ExprToProto(subquery)
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
	proto := ExprToProto(validSubquery)
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
		plan.SetPredicateOpExists,
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
	proto := ExprToProto(subquery)
	require.NotNil(t, proto)
	require.NotNil(t, proto.GetSubquery())
	require.NotNil(t, proto.GetSubquery().GetSetPredicate())
}

func TestSetPredicateSubqueryValidConstruction(t *testing.T) {
	mockRel := createMockReadRel()

	// Test with EXISTS operation
	existsSubquery := plan.NewSetPredicateSubquery(
		plan.SetPredicateOpExists,
		mockRel,
	)
	assert.NotNil(t, existsSubquery)
	assert.Equal(t, plan.SetPredicateOpExists, existsSubquery.Operation)
	assert.Equal(t, mockRel, existsSubquery.Tuples)

	// Test with UNIQUE operation
	uniqueSubquery := plan.NewSetPredicateSubquery(
		plan.SetPredicateOpUnique,
		mockRel,
	)
	assert.NotNil(t, uniqueSubquery)
	assert.Equal(t, plan.SetPredicateOpUnique, uniqueSubquery.Operation)
	assert.Equal(t, mockRel, uniqueSubquery.Tuples)

	// Test with UNSPECIFIED operation
	unspecifiedSubquery := plan.NewSetPredicateSubquery(
		plan.SetPredicateOpUnspecified,
		mockRel,
	)
	assert.NotNil(t, unspecifiedSubquery)
	assert.Equal(t, plan.SetPredicateOpUnspecified, unspecifiedSubquery.Operation)

	// Test with nil relation
	nilRelSubquery := plan.NewSetPredicateSubquery(
		plan.SetPredicateOpExists,
		nil,
	)
	assert.NotNil(t, nilRelSubquery)
	assert.Equal(t, plan.SetPredicateOpExists, nilRelSubquery.Operation)
	assert.Nil(t, nilRelSubquery.Tuples)

	// Test protobuf conversion with valid arguments
	protoMsg := ExprToProto(existsSubquery)
	require.NotNil(t, protoMsg)
	require.NotNil(t, protoMsg.GetSubquery())
	require.NotNil(t, protoMsg.GetSubquery().GetSetPredicate())
	assert.Equal(t, proto.Expression_Subquery_SetPredicate_PREDICATE_OP_EXISTS,
		protoMsg.GetSubquery().GetSetPredicate().GetPredicateOp())
}
