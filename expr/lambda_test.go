// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	substraitgo "github.com/substrait-io/substrait-go/v7"
	"github.com/substrait-io/substrait-go/v7/expr"
	ext "github.com/substrait-io/substrait-go/v7/extensions"
	"github.com/substrait-io/substrait-go/v7/literal"
	"github.com/substrait-io/substrait-go/v7/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/encoding/protojson"
	pb "google.golang.org/protobuf/proto"
)

// TestLambdaBuilder_ValidationErrors tests error cases in the builder.
func TestLambdaBuilder_ValidationErrors(t *testing.T) {
	b := &expr.ExprBuilder{}
	body := &expr.PrimitiveLiteral[int32]{Value: 42, Type: &types.Int32Type{Nullability: types.NullabilityRequired}}

	// Error: nil parameters
	_, err := b.Lambda(nil, b.Expression(body)).Build()
	require.ErrorIs(t, err, substraitgo.ErrInvalidExpr)
	require.Contains(t, err.Error(), "must have parameters")

	// Error: no body
	params := &types.StructType{
		Types:       []types.Type{&types.Int32Type{Nullability: types.NullabilityRequired}},
		Nullability: types.NullabilityRequired,
	}
	_, err = b.Lambda(params, nil).Build()
	require.ErrorIs(t, err, substraitgo.ErrInvalidExpr)
	require.Contains(t, err.Error(), "must have a body")

	// Error: wrong nullability on parameters struct
	badNullParams := &types.StructType{
		Types:       []types.Type{&types.Int32Type{Nullability: types.NullabilityRequired}},
		Nullability: types.NullabilityNullable, // Should be Required
	}
	_, err = b.Lambda(badNullParams, b.Expression(body)).Build()
	require.ErrorIs(t, err, substraitgo.ErrInvalidExpr)
	require.Contains(t, err.Error(), "NULLABILITY_REQUIRED")

	// Error: nil parameter type in list
	nilParamType := &types.StructType{
		Types:       []types.Type{nil},
		Nullability: types.NullabilityRequired,
	}
	_, err = b.Lambda(nilParamType, b.Expression(body)).Build()
	require.ErrorIs(t, err, substraitgo.ErrInvalidExpr)
	require.Contains(t, err.Error(), "nil type")

	// Error: outer lambda parameter out of bounds (stepsOut > 0 case)
	outerParams := &types.StructType{
		Types:       []types.Type{&types.Int32Type{Nullability: types.NullabilityRequired}},
		Nullability: types.NullabilityRequired,
	}
	innerParams := &types.StructType{
		Types:       []types.Type{&types.Float64Type{Nullability: types.NullabilityRequired}},
		Nullability: types.NullabilityRequired,
	}
	// Inner body references outer.field(5) but outer only has 1 param
	_, err = b.Lambda(outerParams,
		b.Lambda(innerParams,
			b.LambdaParamRef(expr.StructFieldRef{Field: 5}, 1), // Out of bounds!
		),
	).Build()
	require.ErrorIs(t, err, substraitgo.ErrInvalidExpr)
	require.Contains(t, err.Error(), "parameter 5")
	require.Contains(t, err.Error(), "1 steps out")
	require.Contains(t, err.Error(), "only has 1 parameters")
}

// TestLambdaProtoRoundTrip tests that lambda expressions can be round-tripped
// through protobuf serialization without losing information.
func TestLambdaProtoRoundTrip(t *testing.T) {
	// Load all JSON files from testdata/lambda/
	files, err := os.ReadDir("./testdata/lambda")
	require.NoError(t, err, "Should be able to read testdata/lambda directory")

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".json") {
			continue
		}

		t.Run(file.Name(), func(t *testing.T) {
			// Read JSON file
			data, err := os.ReadFile("./testdata/lambda/" + file.Name())
			require.NoError(t, err, "Should read JSON file")

			// Parse plan
			var plan proto.Plan
			err = protojson.Unmarshal(data, &plan)
			require.NoError(t, err, "Should parse Plan JSON")

			// Setup extension collection and registry
			collection := ext.GetDefaultCollectionWithNoError()
			extSet, err := ext.GetExtensionSet(&plan, collection)
			require.NoError(t, err, "Should get extension set from plan")
			reg := expr.NewExtensionRegistry(extSet, collection)

			// Extract expression from plan
			project := plan.Relations[0].GetRoot().GetInput().GetProject()
			require.NotNil(t, project, "Should have project relation")
			require.Len(t, project.Expressions, 1, "Should have 1 expression")
			originalExprProto := project.Expressions[0]

			// Roundtrip: proto → Go → proto
			goExpr, err := expr.ExprFromProto(originalExprProto, nil, reg)
			require.NoError(t, err, "Should convert proto to Go expression")

			resultProto := goExpr.ToProto()

			// Verify round-trip: original proto should equal result proto
			require.True(t, pb.Equal(originalExprProto, resultProto),
				"Round-trip failed for %s!\nOriginal: %v\nResult: %v",
				file.Name(), originalExprProto, resultProto)

			t.Logf("Round-trip successful for %s", file.Name())
		})
	}
}

// TestLambdaBuilder_ZeroParameters tests that lambdas with no parameters are valid.
func TestLambdaBuilder_ZeroParameters(t *testing.T) {
	// Building: () -> i32(42) : i32
	b := &expr.ExprBuilder{}
	params := &types.StructType{
		Nullability: types.NullabilityRequired,
		Types:       []types.Type{}, // No parameters
	}
	body := literal.NewInt32(42, false)

	lambda, err := b.Lambda(params, b.Expression(body)).Build()

	require.NoError(t, err, "Zero-parameter lambda should be valid")
	require.NotNil(t, lambda)
	require.Len(t, lambda.Parameters.Types, 0, "Should have zero parameters")

	funcType, ok := lambda.GetType().(*types.FuncType)
	require.True(t, ok, "Lambda type should be FuncType")
	require.Equal(t, "i32", funcType.ReturnType.ShortString(), "Return type should be i32")
	t.Logf("Zero-parameter lambda: %s", lambda.String())
}

// TestLambdaBuilder_ValidStepsOut0 tests that Build() passes for valid stepsOut=0 references.
func TestLambdaBuilder_ValidStepsOut0(t *testing.T) {
	// Building: ($0: i32) -> $0 : i32
	b := &expr.ExprBuilder{}
	// Create a parameter reference to field 0 of the lambda
	params := &types.StructType{
		Nullability: types.NullabilityRequired,
		Types:       []types.Type{&types.Int32Type{Nullability: types.NullabilityRequired}},
	}

	lambda, err := b.Lambda(params,
		b.LambdaParamRef(expr.StructFieldRef{Field: 0}, 0),
	).Build()

	require.NoError(t, err, "Build should pass for valid stepsOut=0 reference")
	require.NotNil(t, lambda)
	require.Equal(t, params, lambda.Parameters)
	require.False(t, lambda.IsScalar()) // Lambda is not scalar (type is func)
	t.Logf("Lambda built successfully: %s", lambda.String())

	// Test Equals - same lambda
	lambda2, _ := b.Lambda(params,
		b.LambdaParamRef(expr.StructFieldRef{Field: 0}, 0),
	).Build()
	require.True(t, lambda.Equals(lambda2))

	// Test Equals - different params
	// Building: ($0: i64) -> $0 : i64
	differentParams := &types.StructType{
		Nullability: types.NullabilityRequired,
		Types:       []types.Type{&types.Int64Type{Nullability: types.NullabilityRequired}},
	}

	lambda3, _ := b.Lambda(differentParams,
		b.LambdaParamRef(expr.StructFieldRef{Field: 0}, 0),
	).Build()
	require.False(t, lambda.Equals(lambda3))

	// Test Visit - body unchanged returns same lambda
	sameLambda := lambda.Visit(func(e expr.Expression) expr.Expression { return e })
	require.Equal(t, lambda, sameLambda)

	// Test Visit - body changed returns new lambda
	newBody := literal.NewInt32(99, false)
	newLambda := lambda.Visit(func(e expr.Expression) expr.Expression { return newBody })
	require.NotEqual(t, lambda, newLambda)
	require.Equal(t, newBody, newLambda.(*expr.Lambda).Body)
}

// TestLambdaBuilder_InvalidOuterRef tests that LambdaParamRef() fails for invalid outer refs.
func TestLambdaBuilder_InvalidOuterRef(t *testing.T) {
	// Building: ($0: i32) -> outer[$0] : INVALID (no outer lambda)
	b := &expr.ExprBuilder{}
	// Create a parameter reference with stepsOut=1 but no outer lambda
	params := &types.StructType{
		Nullability: types.NullabilityRequired,
		Types:       []types.Type{&types.Int32Type{Nullability: types.NullabilityRequired}},
	}

	lambda, err := b.Lambda(params,
		b.LambdaParamRef(expr.StructFieldRef{Field: 0}, 1), // stepsOut=1 but no outer lambda
	).Build()

	require.Error(t, err, "Build should fail for invalid outer reference")
	require.Nil(t, lambda)
	require.Contains(t, err.Error(), "stepsOut 1")
	require.Contains(t, err.Error(), "non-existent outer lambda")
	t.Logf("Expected error: %v", err)
}

// TestLambdaBuilder_InvalidFieldIndex tests that LambdaParamRef() fails when referencing
// a field index that is out of bounds for the lambda's parameters.
func TestLambdaBuilder_InvalidFieldIndex(t *testing.T) {
	// Building: ($0: i32) -> $5 : INVALID (only has 1 param)
	b := &expr.ExprBuilder{}
	// Lambda has 1 parameter but body references field 5
	params := &types.StructType{
		Nullability: types.NullabilityRequired,
		Types:       []types.Type{&types.Int32Type{Nullability: types.NullabilityRequired}},
	}

	lambda, err := b.Lambda(params,
		b.LambdaParamRef(expr.StructFieldRef{Field: 5}, 0), // invalid - only 1 param (index 0)
	).Build()

	require.Error(t, err, "Build should fail for out-of-bounds field index")
	require.Nil(t, lambda)
	require.Contains(t, err.Error(), "references parameter 5")
	require.Contains(t, err.Error(), "only has 1 parameters")
	t.Logf("Expected error: %v", err)
}

// TestLambdaBuilder_ValidFieldIndex tests that LambdaParamRef() passes for valid field indices.
func TestLambdaBuilder_ValidFieldIndex(t *testing.T) {
	// Building: ($0: i32, $1: i64, $2: string) -> $2 : string
	b := &expr.ExprBuilder{}
	// Lambda has 3 parameters, body references field 2 (valid)
	params := &types.StructType{
		Nullability: types.NullabilityRequired,
		Types: []types.Type{
			&types.Int32Type{Nullability: types.NullabilityRequired},
			&types.Int64Type{Nullability: types.NullabilityRequired},
			&types.StringType{Nullability: types.NullabilityRequired},
		},
	}

	// Using the builder API - should succeed
	lambda, err := b.Lambda(params,
		b.LambdaParamRef(expr.StructFieldRef{Field: 2}, 0), // valid - references 3rd param (string)
	).Build()

	require.NoError(t, err, "Build should pass for valid field index")
	require.NotNil(t, lambda)

	// Verify the resolved type is string
	funcType, ok := lambda.GetType().(*types.FuncType)
	require.True(t, ok, "Lambda type should be FuncType")
	require.Equal(t, "str", funcType.ReturnType.ShortString())
	t.Logf("Lambda built successfully: %s", lambda.String())
}

// TestLambdaBuilder_NestedLambda tests building nested lambdas with outer refs using the builder API.
func TestLambdaBuilder_NestedLambda(t *testing.T) {
	// Building: ($0: i64, $1: i64) -> (($0: i32) -> outer[$0] : i64) : func<i32 -> i64>
	b := &expr.ExprBuilder{}
	// Inner lambda references outer's parameter via stepsOut=1
	innerParams := &types.StructType{
		Nullability: types.NullabilityRequired,
		Types:       []types.Type{&types.Int32Type{Nullability: types.NullabilityRequired}},
	}

	// Outer lambda has 2 parameters
	outerParams := &types.StructType{
		Nullability: types.NullabilityRequired,
		Types: []types.Type{
			&types.Int64Type{Nullability: types.NullabilityRequired},
			&types.Int64Type{Nullability: types.NullabilityRequired},
		},
	}

	outerLambda, err := b.Lambda(outerParams,
		b.Lambda(innerParams,
			b.LambdaParamRef(expr.StructFieldRef{Field: 0}, 1), // stepsOut=1 references outer
		),
	).Build()

	require.NoError(t, err, "Build should pass - inner lambda validly references outer's parameter")
	require.NotNil(t, outerLambda)

	// Verify the structure
	innerLambda, ok := outerLambda.Body.(*expr.Lambda)
	require.True(t, ok, "Body should be a Lambda")
	require.Len(t, innerLambda.Parameters.Types, 1)

	// Test Equals - same nested lambda
	outerLambda2, _ := b.Lambda(outerParams,
		b.Lambda(innerParams,
			b.LambdaParamRef(expr.StructFieldRef{Field: 0}, 1),
		),
	).Build()
	require.True(t, outerLambda.Equals(outerLambda2))

	// Test Equals - different stepsOut in inner body
	// Building: ($0: i64, $1: i64) -> (($0: i32) -> $0 : i32) : func<i32 -> i32>
	outerLambda3, _ := b.Lambda(outerParams,
		b.Lambda(innerParams,
			b.LambdaParamRef(expr.StructFieldRef{Field: 0}, 0), // stepsOut=0 (different)
		),
	).Build()
	require.False(t, outerLambda.Equals(outerLambda3))

	t.Logf("Nested lambda built successfully: %s", outerLambda.String())
}

// TestLambdaBuilder_NestedInvalidOuterRef tests that LambdaParamRef() fails for invalid nested outer refs.
func TestLambdaBuilder_NestedInvalidOuterRef(t *testing.T) {
	// Building: ($0: i64) -> (($0: i32) -> outer.outer[$0] : INVALID (no grandparent lambda)
	b := &expr.ExprBuilder{}
	// Inner lambda references stepsOut=2, but only 1 outer lambda exists
	innerParams := &types.StructType{
		Nullability: types.NullabilityRequired,
		Types:       []types.Type{&types.Int32Type{Nullability: types.NullabilityRequired}},
	}

	// Outer lambda
	outerParams := &types.StructType{
		Nullability: types.NullabilityRequired,
		Types:       []types.Type{&types.Int64Type{Nullability: types.NullabilityRequired}},
	}

	outerLambda, err := b.Lambda(outerParams,
		b.Lambda(innerParams,
			b.LambdaParamRef(expr.StructFieldRef{Field: 0}, 2), // invalid - no grandparent
		),
	).Build()

	require.Error(t, err, "Build should fail - inner references non-existent grandparent")
	require.Nil(t, outerLambda)
	require.Contains(t, err.Error(), "stepsOut 2")
	t.Logf("Expected error: %v", err)
}

// TestLambdaBuilder_TypeResolution verifies that GetType() returns the correct type
// after building a lambda, including type resolution for LambdaParameterReferences.
func TestLambdaBuilder_TypeResolution(t *testing.T) {
	tests := []struct {
		name         string
		paramTypes   []types.Type
		fieldIndex   int32
		expectedType string
	}{
		{
			name: "reference first param (i32)",
			paramTypes: []types.Type{
				&types.Int32Type{Nullability: types.NullabilityRequired},
			},
			fieldIndex:   0,
			expectedType: "i32",
		},
		{
			name: "reference second param (i64)",
			paramTypes: []types.Type{
				&types.Int32Type{Nullability: types.NullabilityRequired},
				&types.Int64Type{Nullability: types.NullabilityRequired},
			},
			fieldIndex:   1,
			expectedType: "i64",
		},
		{
			name: "reference third param (string)",
			paramTypes: []types.Type{
				&types.Int32Type{Nullability: types.NullabilityRequired},
				&types.Int64Type{Nullability: types.NullabilityRequired},
				&types.StringType{Nullability: types.NullabilityRequired},
			},
			fieldIndex:   2,
			expectedType: "str",
		},
		{
			name: "reference float64 param",
			paramTypes: []types.Type{
				&types.Float64Type{Nullability: types.NullabilityRequired},
			},
			fieldIndex:   0,
			expectedType: "fp64",
		},
	}

	b := &expr.ExprBuilder{}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			params := &types.StructType{
				Nullability: types.NullabilityRequired,
				Types:       tc.paramTypes,
			}

			// Building: (params) -> $fieldIndex : paramType
			lambda, err := b.Lambda(params,
				b.LambdaParamRef(expr.StructFieldRef{Field: tc.fieldIndex}, 0),
			).Build()

			require.NoError(t, err, "Build should succeed")
			require.NotNil(t, lambda)

			// Verify GetType() returns a FuncType
			lambdaType := lambda.GetType()
			require.NotNil(t, lambdaType, "Lambda should have a resolved type")
			funcType, ok := lambdaType.(*types.FuncType)
			require.True(t, ok, "Lambda type should be FuncType")
			require.Equal(t, tc.expectedType, funcType.ReturnType.ShortString(),
				"Lambda return type should match referenced parameter type")

			// Also verify body's type matches
			bodyType := lambda.Body.GetType()
			require.NotNil(t, bodyType, "Body should have a resolved type")
			require.Equal(t, tc.expectedType, bodyType.ShortString(),
				"Body type should match referenced parameter type")

			t.Logf("Lambda: %s → func type: %s", lambda.String(), lambdaType.String())
		})
	}
}

// TestLambdaBuilder_OuterRefTypeResolution verifies that GetType() correctly resolves
// types for outer lambda parameter references (stepsOut > 0).
func TestLambdaBuilder_OuterRefTypeResolution(t *testing.T) {
	// Building: ($0: i32, $1: i64, $2: string) -> (($0: fp64) -> outer[$2] : string) : func<fp64 -> string>
	b := &expr.ExprBuilder{}
	outerParams := &types.StructType{
		Nullability: types.NullabilityRequired,
		Types: []types.Type{
			&types.Int32Type{Nullability: types.NullabilityRequired},
			&types.Int64Type{Nullability: types.NullabilityRequired},
			&types.StringType{Nullability: types.NullabilityRequired},
		},
	}

	innerParams := &types.StructType{
		Nullability: types.NullabilityRequired,
		Types:       []types.Type{&types.Float64Type{Nullability: types.NullabilityRequired}},
	}

	outerLambda, err := b.Lambda(outerParams,
		b.Lambda(innerParams,
			b.LambdaParamRef(expr.StructFieldRef{Field: 2}, 1), // stepsOut=1, field=2
		),
	).Build()

	require.NoError(t, err)
	require.NotNil(t, outerLambda)

	// Get inner lambda
	innerLambda, ok := outerLambda.Body.(*expr.Lambda)
	require.True(t, ok)

	// Verify inner lambda returns a FuncType with return type of string (outer.c's type)
	innerType := innerLambda.GetType()
	require.NotNil(t, innerType, "Inner lambda should have resolved type")
	funcType, ok := innerType.(*types.FuncType)
	require.True(t, ok, "Inner lambda type should be FuncType")
	require.Equal(t, "str", funcType.ReturnType.ShortString(),
		"Inner lambda return type should be string (from outer.c)")

	// Verify the body's type is also resolved
	bodyType := innerLambda.Body.GetType()
	require.NotNil(t, bodyType, "Body should have resolved type")
	require.Equal(t, "str", bodyType.ShortString(),
		"Body type should be string (outer param at field 2)")

	t.Logf("Outer lambda: %s", outerLambda.String())
	t.Logf("Inner lambda type (from outer.c): %s", funcType.ReturnType.ShortString())
}

// TestLambdaBuilder_DeeplyNestedFieldRef tests that type resolution
// works for LambdaParamRef nested inside other expressions (e.g., Cast(LambdaParamRef)).
func TestLambdaBuilder_DeeplyNestedFieldRef(t *testing.T) {
	// Building: ($0: i32) -> cast($0 as i64) : i64
	b := &expr.ExprBuilder{}

	params := &types.StructType{
		Types:       []types.Type{&types.Int32Type{Nullability: types.NullabilityRequired}},
		Nullability: types.NullabilityRequired,
	}

	lambda, err := b.Lambda(params,
		b.Cast(
			b.LambdaParamRef(expr.StructFieldRef{Field: 0}, 0),
			&types.Int64Type{Nullability: types.NullabilityRequired},
		),
	).Build()

	require.NoError(t, err, "Should build lambda with Cast(LambdaParamRef) body")
	require.NotNil(t, lambda)

	// Verify the nested FieldRef has its type resolved
	resultCast, ok := lambda.Body.(*expr.Cast)
	require.True(t, ok, "Body should be Cast")

	resultFieldRef, ok := resultCast.Input.(*expr.FieldReference)
	require.True(t, ok, "Cast input should be FieldReference")

	// This is the key assertion - type should be resolved despite being nested
	require.NotNil(t, resultFieldRef.GetType(), "Nested FieldRef should have type resolved")
	require.Equal(t, "i32", resultFieldRef.GetType().ShortString(), "Should resolve to i32")

	t.Logf("Lambda with deeply nested FieldRef: %s", lambda.String())
}

// TestLambdaBuilder_DeeplyNestedInvalidFieldRef tests that LambdaParamRef validation catches
// invalid references even when nested inside other expressions (e.g., Cast).
func TestLambdaBuilder_DeeplyNestedInvalidFieldRef(t *testing.T) {
	// Building: ($0: i32) -> cast($5 as i64) : INVALID (only has 1 param)
	b := &expr.ExprBuilder{}
	params := &types.StructType{
		Types:       []types.Type{&types.Int32Type{Nullability: types.NullabilityRequired}},
		Nullability: types.NullabilityRequired,
	}

	_, err := b.Lambda(params,
		b.Cast(
			b.LambdaParamRef(expr.StructFieldRef{Field: 5}, 0), // Invalid - only 1 param!
			&types.Int64Type{Nullability: types.NullabilityRequired},
		),
	).Build()

	require.Error(t, err, "Should fail for invalid nested LambdaParamRef")
	require.ErrorIs(t, err, substraitgo.ErrInvalidExpr)
	require.Contains(t, err.Error(), "parameter 5")

	t.Logf("Correctly caught deeply nested invalid ref: %v", err)
}

// TestNestedLambdaFromProto_OuterRefTypeResolution verifies that when parsing
// a nested lambda from proto, stepsOut > 0 references get their types resolved.
// This exercises the nested lambda handling in resolveLambdaParamTypes.
func TestLambdaBuilder_DoublyNestedFieldRef(t *testing.T) {
	// Building: ($0: i32) -> cast(cast($0 as i64) as string) : string
	b := &expr.ExprBuilder{}
	params := &types.StructType{
		Types:       []types.Type{&types.Int32Type{Nullability: types.NullabilityRequired}},
		Nullability: types.NullabilityRequired,
	}

	lambda, err := b.Lambda(params,
		b.Cast(
			b.Cast(
				b.LambdaParamRef(expr.StructFieldRef{Field: 0}, 0),
				&types.Int64Type{Nullability: types.NullabilityRequired},
			),
			&types.StringType{Nullability: types.NullabilityRequired},
		),
	).Build()

	require.NoError(t, err)
	require.NotNil(t, lambda)

	// Navigate to the deeply nested FieldRef (2 levels deep)
	resultOuter := lambda.Body.(*expr.Cast)
	resultInner := resultOuter.Input.(*expr.Cast)
	resultFieldRef := resultInner.Input.(*expr.FieldReference)

	// Verify type is resolved even at depth 2
	require.NotNil(t, resultFieldRef.GetType(), "FieldRef at depth 2 should have type resolved")
	require.Equal(t, "i32", resultFieldRef.GetType().ShortString())
}
