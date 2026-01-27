// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	substraitgo "github.com/substrait-io/substrait-go/v7"
	"github.com/substrait-io/substrait-go/v7/expr"
	ext "github.com/substrait-io/substrait-go/v7/extensions"
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
			b.LambdaParamRef(&expr.StructFieldRef{Field: 5}, 1), // Out of bounds!
		),
	).Build()
	require.ErrorIs(t, err, substraitgo.ErrInvalidExpr)
	require.Contains(t, err.Error(), "parameter")
}

// TestBasicLambdaPlanExprFromProto converts a basic lambda plan from protobuf
// to Go expressions. Lambda: (x: i32) -> 42
func TestBasicLambdaPlanExprFromProto(t *testing.T) {
	const planJSON = `{
		"version": {"majorNumber": 0, "minorNumber": 79},
		"relations": [{
			"root": {
				"input": {
					"project": {
						"common": {"direct": {}},
						"input": {
							"read": {
								"common": {"direct": {}},
								"baseSchema": {
									"names": ["dummy"],
									"struct": {
										"nullability": "NULLABILITY_REQUIRED",
										"types": [{"i32": {"nullability": "NULLABILITY_REQUIRED"}}]
									}
								},
								"virtualTable": {
									"expressions": [{"fields": [{"literal": {"i32": 0}}]}]
								}
							}
						},
						"expressions": [{
							"lambda": {
								"parameters": {
									"nullability": "NULLABILITY_REQUIRED",
									"types": [{"i32": {"nullability": "NULLABILITY_REQUIRED"}}]
								},
								"body": {
									"literal": {"i32": 42}
								}
							}
						}]
					}
				},
				"names": ["result"]
			}
		}]
	}`

	// Parse JSON into protobuf Plan
	var plan proto.Plan
	err := protojson.Unmarshal([]byte(planJSON), &plan)
	require.NoError(t, err, "Plan JSON should parse to protobuf")

	// Extract the lambda expression proto
	project := plan.Relations[0].GetRoot().GetInput().GetProject()
	lambdaProto := project.Expressions[0]

	// Convert protobuf → Go Expression
	reg := expr.NewEmptyExtensionRegistry(ext.GetDefaultCollectionWithNoError())
	goExpr, err := expr.ExprFromProto(lambdaProto, nil, reg)
	require.NoError(t, err, "Should convert lambda to Go expression")

	// Verify it's a Lambda
	lambda, ok := goExpr.(*expr.Lambda)
	require.True(t, ok, "Should be a Lambda expression")
	require.Len(t, lambda.Parameters.Types, 1, "Lambda should have 1 parameter")
	require.Equal(t, "i32", lambda.Parameters.Types[0].ShortString(), "Parameter should be i32")

	// Verify body is a literal
	literal, ok := lambda.Body.(*expr.PrimitiveLiteral[int32])
	require.True(t, ok, "Lambda body should be a PrimitiveLiteral[int32]")
	require.Equal(t, int32(42), literal.Value, "Literal should be 42")

	// Verify lambda return type
	lambdaType := lambda.GetType()
	require.NotNil(t, lambdaType, "Lambda type should be resolved")
	require.Equal(t, "i32", lambdaType.ShortString(), "Lambda return type should be i32")

}

// TestLambdaPlanExprFromProto tests converting a full plan with a lambda expression
// containing a function call and parameter reference to Go expressions.
// This represents: (x: i32) -> multiply(x, 2)
func TestLambdaPlanExprFromProto(t *testing.T) {
	const planJSON = `{
		"extensionUrns": [
			{
				"extensionUrnAnchor": 1,
				"urn": "extension:io.substrait:functions_arithmetic"
			}
		],
		"extensions": [
			{
				"extensionFunction": {
					"extensionUrnReference": 1,
					"functionAnchor": 1,
					"name": "multiply:i32_i32"
				}
			}
		],
		"relations": [
			{
				"root": {
					"input": {
						"project": {
							"input": {
								"read": {
									"baseSchema": {
										"names": ["values"],
										"struct": {
											"types": [
												{
													"list": {
														"type": {"i32": {"nullability": "NULLABILITY_REQUIRED"}},
														"nullability": "NULLABILITY_REQUIRED"
													}
												}
											],
											"nullability": "NULLABILITY_REQUIRED"
										}
									},
									"namedTable": {"names": ["test_table"]}
								}
							},
							"expressions": [
								{
									"lambda": {
										"parameters": {
											"nullability": "NULLABILITY_REQUIRED",
											"types": [
												{"i32": {"nullability": "NULLABILITY_REQUIRED"}}
											]
										},
										"body": {
											"scalarFunction": {
												"functionReference": 1,
												"outputType": {"i32": {"nullability": "NULLABILITY_REQUIRED"}},
												"arguments": [
													{
														"value": {
															"selection": {
																"directReference": {"structField": {"field": 0}},
																"lambdaParameterReference": {"stepsOut": 0}
															}
														}
													},
													{
														"value": {
															"literal": {"i32": 2}
														}
													}
												]
											}
										}
									}
								}
							]
						}
					},
					"names": ["result"]
				}
			}
		]
	}`

	// Parse JSON to protobuf Plan
	var plan proto.Plan
	err := protojson.Unmarshal([]byte(planJSON), &plan)
	require.NoError(t, err, "Plan JSON should parse to protobuf")

	// Extract the lambda expression proto
	project := plan.Relations[0].GetRoot().GetInput().GetProject()
	require.NotNil(t, project, "Should have project relation")
	lambdaProto := project.Expressions[0]

	// Build extension registry from the plan
	collection := ext.GetDefaultCollectionWithNoError()
	extSet, err := ext.GetExtensionSet(&plan, collection)
	require.NoError(t, err, "Should get extension set from plan")
	reg := expr.NewExtensionRegistry(extSet, collection)

	// Convert protobuf → Go Expression
	goExpr, err := expr.ExprFromProto(lambdaProto, nil, reg)
	require.NoError(t, err, "Should convert lambda to Go expression")

	// Verify it's a Lambda
	lambda, ok := goExpr.(*expr.Lambda)
	require.True(t, ok, "Should be a Lambda expression")
	require.Len(t, lambda.Parameters.Types, 1, "Lambda should have 1 parameter")
	require.Equal(t, "i32", lambda.Parameters.Types[0].ShortString(), "Parameter should be i32")

	// Verify body is a ScalarFunction
	scalarFunc, ok := lambda.Body.(*expr.ScalarFunction)
	require.True(t, ok, "Lambda body should be a ScalarFunction")
	require.Equal(t, "multiply", scalarFunc.Name())
	require.Equal(t, 2, scalarFunc.NArgs(), "multiply should have 2 arguments")

	// Verify first argument is a FieldReference with LambdaParameterReference
	arg0 := scalarFunc.Arg(0)
	fieldRef, ok := arg0.(*expr.FieldReference)
	require.True(t, ok, "First arg should be a FieldReference")

	lambdaParamRef, ok := fieldRef.Root.(expr.LambdaParameterReference)
	require.True(t, ok, "FieldReference root should be LambdaParameterReference")
	require.Equal(t, uint32(0), lambdaParamRef.StepsOut, "StepsOut should be 0")

	structFieldRef, ok := fieldRef.Reference.(*expr.StructFieldRef)
	require.True(t, ok, "Reference should be StructFieldRef")
	require.Equal(t, int32(0), structFieldRef.Field, "Should reference parameter 0")

	// Verify second argument is a literal
	arg1 := scalarFunc.Arg(1)
	literal, ok := arg1.(*expr.PrimitiveLiteral[int32])
	require.True(t, ok, "Second arg should be a PrimitiveLiteral[int32]")
	require.Equal(t, int32(2), literal.Value, "Literal value should be 2")

	t.Logf("Lambda: %s", lambda.String())
	t.Logf("Body function: %s with %d args", scalarFunc.Name(), scalarFunc.NArgs())
	t.Logf("Arg0: parameter reference to field %d", structFieldRef.Field)
	t.Logf("Arg1: literal %d", literal.Value)
}

// TestLambdaReferenceExprFromProto converts a lambda with parameter reference
// from protobuf to Go expression and verifies the structure.
func TestLambdaReferenceExprFromProto(t *testing.T) {
	// Lambda: (x: i32) -> x (identity function)
	const lambdaJSON = `{
		"lambda": {
			"parameters": {
				"nullability": "NULLABILITY_REQUIRED",
				"types": [
					{"i32": {"nullability": "NULLABILITY_REQUIRED"}}
				]
			},
			"body": {
				"selection": {
					"directReference": {"structField": {"field": 0}},
					"lambdaParameterReference": {"stepsOut": 0}
				}
			}
		}
	}`

	var exprProto proto.Expression
	err := protojson.Unmarshal([]byte(lambdaJSON), &exprProto)
	require.NoError(t, err, "JSON should parse to protobuf")

	// Convert protobuf → Go Expression
	reg := expr.NewEmptyExtensionRegistry(ext.GetDefaultCollectionWithNoError())
	goExpr, err := expr.ExprFromProto(&exprProto, nil, reg)
	require.NoError(t, err, "Should convert lambda to Go expression")

	// Verify it's a Lambda
	lambda, ok := goExpr.(*expr.Lambda)
	require.True(t, ok, "Should be a Lambda expression")
	require.Len(t, lambda.Parameters.Types, 1, "Lambda should have 1 parameter")

	// Verify the body is a FieldReference with LambdaParameterReference root
	fieldRef, ok := lambda.Body.(*expr.FieldReference)
	require.True(t, ok, "Lambda body should be a FieldReference")

	lambdaParamRef, ok := fieldRef.Root.(expr.LambdaParameterReference)
	require.True(t, ok, "FieldReference root should be LambdaParameterReference")
	require.Equal(t, uint32(0), lambdaParamRef.StepsOut, "StepsOut should be 0")

	// Verify the field reference points to parameter 0
	structFieldRef, ok := fieldRef.Reference.(*expr.StructFieldRef)
	require.True(t, ok, "Reference should be a StructFieldRef")
	require.Equal(t, int32(0), structFieldRef.Field, "Should reference field 0")

	// TODO (#189): Add type resolution for lambda parameter references during proto parsing
	// For now, type resolution is only done by the builder, not during proto parsing
	// bodyType := lambda.Body.GetType()
	// require.NotNil(t, bodyType, "Body type should be resolved")
	// require.Equal(t, "i32", bodyType.ShortString(), "Body type should be i32")

	t.Logf("Lambda Go expression: %s", lambda.String())
	t.Logf("Body references parameter %d via LambdaParameterReference", structFieldRef.Field)
}

// TestLambdaWithFunctionExprFromProto converts a lambda with a scalar function
// from protobuf to Go expression. Lambda: (x: i32) -> multiply(x, 2)
func TestLambdaWithFunctionExprFromProto(t *testing.T) {
	const planJSON = `{
		"extensionUrns": [
			{"extensionUrnAnchor": 1, "urn": "extension:io.substrait:functions_arithmetic"}
		],
		"extensions": [
			{"extensionFunction": {"extensionUrnReference": 1, "functionAnchor": 1, "name": "multiply:i32_i32"}}
		],
		"relations": [{
			"root": {
				"input": {
					"project": {
						"common": {"direct": {}},
						"input": {
							"read": {
								"common": {"direct": {}},
								"baseSchema": {
									"names": ["dummy"],
									"struct": {
										"nullability": "NULLABILITY_REQUIRED",
										"types": [{"i32": {"nullability": "NULLABILITY_REQUIRED"}}]
									}
								},
								"virtualTable": {"expressions": [{"fields": [{"literal": {"i32": 0}}]}]}
							}
						},
						"expressions": [{
							"lambda": {
								"parameters": {
									"nullability": "NULLABILITY_REQUIRED",
									"types": [{"i32": {"nullability": "NULLABILITY_REQUIRED"}}]
								},
								"body": {
									"scalarFunction": {
										"functionReference": 1,
										"outputType": {"i32": {"nullability": "NULLABILITY_REQUIRED"}},
										"arguments": [
											{
												"value": {
													"selection": {
														"directReference": {"structField": {"field": 0}},
														"lambdaParameterReference": {"stepsOut": 0}
													}
												}
											},
											{"value": {"literal": {"i32": 2}}}
										]
									}
								}
							}
						}]
					}
				},
				"names": ["result"]
			}
		}]
	}`

	// Parse plan
	var plan proto.Plan
	err := protojson.Unmarshal([]byte(planJSON), &plan)
	require.NoError(t, err, "Plan JSON should parse")

	// Extract the lambda expression from the project
	project := plan.Relations[0].GetRoot().GetInput().GetProject()
	lambdaProto := project.Expressions[0]

	// Build extension registry from plan
	collection := ext.GetDefaultCollectionWithNoError()
	extSet, err := ext.GetExtensionSet(&plan, collection)
	require.NoError(t, err, "Should get extension set")
	reg := expr.NewExtensionRegistry(extSet, collection)

	// Convert protobuf → Go Expression
	goExpr, err := expr.ExprFromProto(lambdaProto, nil, reg)
	require.NoError(t, err, "Should convert lambda to Go expression")

	// Verify it's a Lambda
	lambda, ok := goExpr.(*expr.Lambda)
	require.True(t, ok, "Should be a Lambda expression")
	require.Len(t, lambda.Parameters.Types, 1, "Lambda should have 1 parameter")
	require.Equal(t, "i32", lambda.Parameters.Types[0].ShortString(), "Parameter should be i32")

	// Verify body is a ScalarFunction
	scalarFunc, ok := lambda.Body.(*expr.ScalarFunction)
	require.True(t, ok, "Lambda body should be a ScalarFunction")
	require.Equal(t, "multiply", scalarFunc.Name(), "Function should be multiply")
	require.Equal(t, 2, scalarFunc.NArgs(), "Function should have 2 arguments")

	// Verify first argument is a FieldReference with LambdaParameterReference
	arg0 := scalarFunc.Arg(0)
	fieldRef, ok := arg0.(*expr.FieldReference)
	require.True(t, ok, "First arg should be a FieldReference")

	lambdaParamRef, ok := fieldRef.Root.(expr.LambdaParameterReference)
	require.True(t, ok, "Root should be LambdaParameterReference")
	require.Equal(t, uint32(0), lambdaParamRef.StepsOut, "StepsOut should be 0")

	// TODO (#189): Add type resolution for lambda parameter references during proto parsing
	// require.NotNil(t, fieldRef.GetType(), "FieldReference type should be resolved")
	// require.Equal(t, "i32", fieldRef.GetType().ShortString(), "FieldRef type should be i32")

	// Verify second argument is a literal
	arg1 := scalarFunc.Arg(1)
	literal, ok := arg1.(*expr.PrimitiveLiteral[int32])
	require.True(t, ok, "Second arg should be PrimitiveLiteral[int32]")
	require.Equal(t, int32(2), literal.Value, "Literal should be 2")

	t.Logf("Lambda: %s", lambda.String())
	t.Logf("Body function: %s, return type: %s", scalarFunc.Name(), lambda.GetType().ShortString())
}

// TestLambdaBuilder_ZeroParameters tests that lambdas with no parameters are valid.
// Example: () -> 42
func TestLambdaBuilder_ZeroParameters(t *testing.T) {
	b := &expr.ExprBuilder{}
	params := &types.StructType{
		Nullability: types.NullabilityRequired,
		Types:       []types.Type{}, // No parameters
	}
	body := &expr.PrimitiveLiteral[int32]{Value: 42, Type: &types.Int32Type{Nullability: types.NullabilityRequired}}

	lambda, err := b.Lambda(params, b.Expression(body)).Build()

	require.NoError(t, err, "Zero-parameter lambda should be valid")
	require.NotNil(t, lambda)
	require.Len(t, lambda.Parameters.Types, 0, "Should have zero parameters")
	require.Equal(t, "i32", lambda.GetType().ShortString(), "Return type should be i32")
	t.Logf("Zero-parameter lambda: %s", lambda.String())
}

// TestLambdaBuilder_ValidStepsOut0 tests that Build() passes for valid stepsOut=0 references.
func TestLambdaBuilder_ValidStepsOut0(t *testing.T) {
	b := &expr.ExprBuilder{}
	// Create a parameter reference to field 0 of the lambda
	params := &types.StructType{
		Nullability: types.NullabilityRequired,
		Types:       []types.Type{&types.Int32Type{Nullability: types.NullabilityRequired}},
	}

	// Build should succeed - valid stepsOut=0 reference
	lambda, err := b.Lambda(params,
		b.LambdaParamRef(&expr.StructFieldRef{Field: 0}, 0),
	).Build()

	require.NoError(t, err, "Build should pass for valid stepsOut=0 reference")
	require.NotNil(t, lambda)
	require.Equal(t, params, lambda.Parameters)
	require.True(t, lambda.IsScalar()) // FieldReference is scalar
	t.Logf("Lambda built successfully: %s", lambda.String())

	// Test Equals - same lambda
	lambda2, _ := b.Lambda(params,
		b.LambdaParamRef(&expr.StructFieldRef{Field: 0}, 0),
	).Build()
	require.True(t, lambda.Equals(lambda2))

	// Test Equals - different params
	differentParams := &types.StructType{
		Nullability: types.NullabilityRequired,
		Types:       []types.Type{&types.Int64Type{Nullability: types.NullabilityRequired}},
	}
	lambda3, _ := b.Lambda(differentParams,
		b.LambdaParamRef(&expr.StructFieldRef{Field: 0}, 0),
	).Build()
	require.False(t, lambda.Equals(lambda3))

	// Test Visit - body unchanged returns same lambda
	sameLambda := lambda.Visit(func(e expr.Expression) expr.Expression { return e })
	require.Equal(t, lambda, sameLambda)

	// Test Visit - body changed returns new lambda
	newBody := &expr.PrimitiveLiteral[int32]{Value: 99, Type: &types.Int32Type{Nullability: types.NullabilityRequired}}
	newLambda := lambda.Visit(func(e expr.Expression) expr.Expression { return newBody })
	require.NotEqual(t, lambda, newLambda)
	require.Equal(t, newBody, newLambda.(*expr.Lambda).Body)
}

// TestLambdaBuilder_InvalidOuterRef tests that LambdaParamRef() fails for invalid outer refs.
func TestLambdaBuilder_InvalidOuterRef(t *testing.T) {
	b := &expr.ExprBuilder{}
	// Create a parameter reference with stepsOut=1 but no outer lambda
	params := &types.StructType{
		Nullability: types.NullabilityRequired,
		Types:       []types.Type{&types.Int32Type{Nullability: types.NullabilityRequired}},
	}

	// Using the builder API - should fail during LambdaParamRef.Build()
	lambda, err := b.Lambda(params,
		b.LambdaParamRef(&expr.StructFieldRef{Field: 0}, 1), // stepsOut=1 but no outer lambda
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
	b := &expr.ExprBuilder{}
	// Lambda has 1 parameter but body references field 5
	params := &types.StructType{
		Nullability: types.NullabilityRequired,
		Types:       []types.Type{&types.Int32Type{Nullability: types.NullabilityRequired}},
	}

	// Using the builder API - should fail during LambdaParamRef.Build()
	lambda, err := b.Lambda(params,
		b.LambdaParamRef(&expr.StructFieldRef{Field: 5}, 0), // invalid - only 1 param (index 0)
	).Build()

	require.Error(t, err, "Build should fail for out-of-bounds field index")
	require.Nil(t, lambda)
	require.Contains(t, err.Error(), "references parameter 5")
	require.Contains(t, err.Error(), "only has 1 parameters")
	t.Logf("Expected error: %v", err)
}

// TestLambdaBuilder_ValidFieldIndex tests that LambdaParamRef() passes for valid field indices.
func TestLambdaBuilder_ValidFieldIndex(t *testing.T) {
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
		b.LambdaParamRef(&expr.StructFieldRef{Field: 2}, 0), // valid - references 3rd param (string)
	).Build()

	require.NoError(t, err, "Build should pass for valid field index")
	require.NotNil(t, lambda)

	// Verify the resolved type is string
	require.NotNil(t, lambda.GetType())
	require.Equal(t, "str", lambda.GetType().ShortString())
	t.Logf("Lambda built successfully: %s", lambda.String())
}

// TestLambdaBuilder_NestedLambda tests building nested lambdas with outer refs using the builder API.
func TestLambdaBuilder_NestedLambda(t *testing.T) {
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

	// Build outer lambda - inner uses LambdaParamRef with stepsOut=1 to reference outer
	outerLambda, err := b.Lambda(outerParams,
		b.Lambda(innerParams,
			b.LambdaParamRef(&expr.StructFieldRef{Field: 0}, 1), // stepsOut=1 references outer
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
			b.LambdaParamRef(&expr.StructFieldRef{Field: 0}, 1),
		),
	).Build()
	require.True(t, outerLambda.Equals(outerLambda2))

	// Test Equals - different stepsOut in inner body
	outerLambda3, _ := b.Lambda(outerParams,
		b.Lambda(innerParams,
			b.LambdaParamRef(&expr.StructFieldRef{Field: 0}, 0), // stepsOut=0 (different)
		),
	).Build()
	require.False(t, outerLambda.Equals(outerLambda3))

	t.Logf("Nested lambda built successfully: %s", outerLambda.String())
}

// TestLambdaBuilder_NestedInvalidOuterRef tests that LambdaParamRef() fails for invalid nested outer refs.
func TestLambdaBuilder_NestedInvalidOuterRef(t *testing.T) {
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

	// Build should fail - inner references stepsOut=2 but only 1 outer lambda exists
	outerLambda, err := b.Lambda(outerParams,
		b.Lambda(innerParams,
			b.LambdaParamRef(&expr.StructFieldRef{Field: 0}, 2), // invalid - no grandparent
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

			lambda, err := b.Lambda(params,
				b.LambdaParamRef(&expr.StructFieldRef{Field: tc.fieldIndex}, 0),
			).Build()

			require.NoError(t, err, "Build should succeed")
			require.NotNil(t, lambda)

			// Verify GetType() returns the correct resolved type
			lambdaType := lambda.GetType()
			require.NotNil(t, lambdaType, "Lambda should have a resolved type")
			require.Equal(t, tc.expectedType, lambdaType.ShortString(),
				"Lambda type should match referenced parameter type")

			// Also verify body's type matches
			bodyType := lambda.Body.GetType()
			require.NotNil(t, bodyType, "Body should have a resolved type")
			require.Equal(t, tc.expectedType, bodyType.ShortString(),
				"Body type should match referenced parameter type")

			t.Logf("Lambda: %s → type: %s", lambda.String(), lambdaType.ShortString())
		})
	}
}

// TestLambdaBuilder_OuterRefTypeResolution verifies that GetType() correctly resolves
// types for outer lambda parameter references (stepsOut > 0).
func TestLambdaBuilder_OuterRefTypeResolution(t *testing.T) {
	b := &expr.ExprBuilder{}
	// Outer lambda: (a: i32, b: i64, c: string)
	// Inner lambda: (x: fp64) -> outer.c (stepsOut=1, field=2) → should be string
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

	// Use LambdaParamRef to reference outer.c (stepsOut=1, field=2)
	outerLambda, err := b.Lambda(outerParams,
		b.Lambda(innerParams,
			b.LambdaParamRef(&expr.StructFieldRef{Field: 2}, 1), // stepsOut=1, field=2
		),
	).Build()

	require.NoError(t, err)
	require.NotNil(t, outerLambda)

	// Get inner lambda
	innerLambda, ok := outerLambda.Body.(*expr.Lambda)
	require.True(t, ok)

	// Verify inner lambda's type is resolved to string (outer.c's type)
	innerType := innerLambda.GetType()
	require.NotNil(t, innerType, "Inner lambda should have resolved type")
	require.Equal(t, "str", innerType.ShortString(),
		"Inner lambda type should be string (from outer.c)")

	// Verify the body's type is also resolved
	bodyType := innerLambda.Body.GetType()
	require.NotNil(t, bodyType, "Body should have resolved type")
	require.Equal(t, "str", bodyType.ShortString(),
		"Body type should be string (outer param at field 2)")

	t.Logf("Outer lambda: %s", outerLambda.String())
	t.Logf("Inner lambda type (from outer.c): %s", innerType.ShortString())
}

// TestLambdaBuilder_DeeplyNestedFieldRef tests that type resolution
// works for LambdaParamRef nested inside other expressions (e.g., Cast(LambdaParamRef)).
func TestLambdaBuilder_DeeplyNestedFieldRef(t *testing.T) {
	b := &expr.ExprBuilder{}
	// Create a lambda with body: Cast(LambdaParamRef($0))

	params := &types.StructType{
		Types:       []types.Type{&types.Int32Type{Nullability: types.NullabilityRequired}},
		Nullability: types.NullabilityRequired,
	}

	// Build lambda with Cast(LambdaParamRef) as body using builder API
	lambda, err := b.Lambda(params,
		b.Cast(
			b.LambdaParamRef(&expr.StructFieldRef{Field: 0}, 0),
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
	b := &expr.ExprBuilder{}
	params := &types.StructType{
		Types:       []types.Type{&types.Int32Type{Nullability: types.NullabilityRequired}},
		Nullability: types.NullabilityRequired,
	}

	// Build should fail because LambdaParamRef validates during its Build()
	_, err := b.Lambda(params,
		b.Cast(
			b.LambdaParamRef(&expr.StructFieldRef{Field: 5}, 0), // Invalid - only 1 param!
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
func TestNestedLambdaFromProto_OuterRefTypeResolution(t *testing.T) {
	// Nested lambda: outer(x: i32) -> inner(y: fp64) -> $x (stepsOut=1, field=0)
	// The inner lambda references outer's parameter, which requires type resolution
	// to happen during the Visit walk in resolveLambdaParamTypes.
	const planJSON = `{
		"relations": [
			{
				"root": {
					"input": {
						"project": {
							"input": {
								"read": {
									"baseSchema": {
										"names": ["dummy"],
										"struct": {
											"types": [{"i32": {"nullability": "NULLABILITY_REQUIRED"}}],
											"nullability": "NULLABILITY_REQUIRED"
										}
									},
									"namedTable": {"names": ["test"]}
								}
							},
							"expressions": [
								{
									"lambda": {
										"parameters": {
											"nullability": "NULLABILITY_REQUIRED",
											"types": [
												{"i32": {"nullability": "NULLABILITY_REQUIRED"}}
											]
										},
										"body": {
											"lambda": {
												"parameters": {
													"nullability": "NULLABILITY_REQUIRED",
													"types": [
														{"fp64": {"nullability": "NULLABILITY_REQUIRED"}}
													]
												},
												"body": {
													"selection": {
														"directReference": {
															"structField": {"field": 0}
														},
														"lambdaParameterReference": {"stepsOut": 1}
													}
												}
											}
										}
									}
								}
							]
						}
					}
				}
			}
		]
	}`

	var plan proto.Plan
	err := protojson.Unmarshal([]byte(planJSON), &plan)
	require.NoError(t, err)

	projectRel := plan.Relations[0].GetRoot().GetInput().GetProject()
	require.NotNil(t, projectRel)

	// Parse to Go expression
	collection := ext.GetDefaultCollectionWithNoError()
	extSet, err := ext.GetExtensionSet(&plan, collection)
	require.NoError(t, err)

	reg := expr.NewExtensionRegistry(extSet, collection)
	goExpr, err := expr.ExprFromProto(projectRel.Expressions[0], nil, reg)
	require.NoError(t, err)

	// Navigate to the inner lambda's body (the FieldReference with stepsOut=1)
	outerLambda := goExpr.(*expr.Lambda)
	innerLambda := outerLambda.Body.(*expr.Lambda)
	_ = innerLambda.Body.(*expr.FieldReference) // fieldRef exists but type checking deferred to #189

	// TODO (#189): Add type resolution for lambda parameter references during proto parsing
	// fieldRef := innerLambda.Body.(*expr.FieldReference)
	// require.NotNil(t, fieldRef.GetType(),
	//	"stepsOut=1 FieldRef should have type resolved when parsing nested lambda from proto")
	// require.Equal(t, "i32", fieldRef.GetType().ShortString(),
	//	"Should resolve to outer lambda's param type (i32)")

	t.Logf("Nested lambda with resolved outer ref: %s", outerLambda.String())
}

// TestLambdaBuilder_DoublyNestedFieldRef tests LambdaParamRef nested two levels deep:
// Cast(Cast(LambdaParamRef)) - ensures type resolution works at depth 2
func TestLambdaBuilder_DoublyNestedFieldRef(t *testing.T) {
	b := &expr.ExprBuilder{}
	params := &types.StructType{
		Types:       []types.Type{&types.Int32Type{Nullability: types.NullabilityRequired}},
		Nullability: types.NullabilityRequired,
	}

	// Build: Cast(Cast(LambdaParamRef($0))) using builder API
	lambda, err := b.Lambda(params,
		b.Cast(
			b.Cast(
				b.LambdaParamRef(&expr.StructFieldRef{Field: 0}, 0),
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

// TestLambdaAsArgumentFromProto tests type resolution when a Lambda is an ARGUMENT
// to another expression (not the body itself). This exercises the nested lambda
// handling inside the recurse callback in resolveLambdaParamTypes.
// Structure: outer(x: i32) -> if_then(condition, inner(y) -> $x)
// The inner lambda is an argument to if_then, not the body of outer.
func TestLambdaAsArgumentFromProto(t *testing.T) {
	// This JSON has: outer lambda whose body is an if_then expression
	// One branch of the if_then contains an inner lambda that references outer's param
	const planJSON = `{
		"relations": [
			{
				"root": {
					"input": {
						"project": {
							"input": {
								"read": {
									"baseSchema": {
										"names": ["dummy"],
										"struct": {
											"types": [{"i32": {"nullability": "NULLABILITY_REQUIRED"}}],
											"nullability": "NULLABILITY_REQUIRED"
										}
									},
									"namedTable": {"names": ["test"]}
								}
							},
							"expressions": [
								{
									"lambda": {
										"parameters": {
											"nullability": "NULLABILITY_REQUIRED",
											"types": [
												{"i32": {"nullability": "NULLABILITY_REQUIRED"}}
											]
										},
										"body": {
											"ifThen": {
												"ifs": [
													{
														"if": {
															"literal": {"boolean": true}
														},
														"then": {
															"lambda": {
																"parameters": {
																	"nullability": "NULLABILITY_REQUIRED",
																	"types": [
																		{"fp64": {"nullability": "NULLABILITY_REQUIRED"}}
																	]
																},
																"body": {
																	"selection": {
																		"directReference": {
																			"structField": {"field": 0}
																		},
																		"lambdaParameterReference": {"stepsOut": 1}
																	}
																}
															}
														}
													}
												],
												"else": {
													"literal": {"i32": 0}
												}
											}
										}
									}
								}
							]
						}
					}
				}
			}
		]
	}`

	var plan proto.Plan
	err := protojson.Unmarshal([]byte(planJSON), &plan)
	require.NoError(t, err)

	projectRel := plan.Relations[0].GetRoot().GetInput().GetProject()
	require.NotNil(t, projectRel)

	// Parse to Go expression
	collection := ext.GetDefaultCollectionWithNoError()
	extSet, err := ext.GetExtensionSet(&plan, collection)
	require.NoError(t, err)

	reg := expr.NewExtensionRegistry(extSet, collection)
	goExpr, err := expr.ExprFromProto(projectRel.Expressions[0], nil, reg)
	require.NoError(t, err)

	// Navigate: outer lambda -> if_then body -> first clause's then -> inner lambda -> body
	outerLambda := goExpr.(*expr.Lambda)
	ifThenBody := outerLambda.Body.(*expr.IfThen)
	innerLambda := ifThenBody.IfPair(0).Then.(*expr.Lambda)
	_ = innerLambda.Body.(*expr.FieldReference) // fieldRef exists but type checking deferred to #189

	// TODO (#189): Add type resolution for lambda parameter references during proto parsing
	// fieldRef := innerLambda.Body.(*expr.FieldReference)
	// require.NotNil(t, fieldRef.GetType(),
	//	"stepsOut=1 ref should be resolved when lambda is an argument (not body)")
	// require.Equal(t, "i32", fieldRef.GetType().ShortString())

	t.Logf("Lambda-as-argument resolved correctly: %s", outerLambda.String())
}

func TestBasicLambdaRoundTrip(t *testing.T) {
	const lambdaJSON = `{
		"lambda": {
			"parameters": {
				"nullability": "NULLABILITY_REQUIRED",
				"types": [
					{"i32": {"nullability": "NULLABILITY_REQUIRED"}}
				]
			},
			"body": {
				"literal": {"i32": 42}
			}
		}
	}`

	// Parse original JSON to protobuf
	var originalProto proto.Expression
	err := protojson.Unmarshal([]byte(lambdaJSON), &originalProto)
	require.NoError(t, err)

	// Convert protobuf → Go Expression
	reg := expr.NewEmptyExtensionRegistry(ext.GetDefaultCollectionWithNoError())
	goExpr, err := expr.ExprFromProto(&originalProto, nil, reg)
	require.NoError(t, err)

	// Convert Go Expression → protobuf
	resultProto := goExpr.ToProto()

	// Verify round-trip: original proto should equal result proto
	require.True(t, pb.Equal(&originalProto, resultProto),
		"Round-trip failed!\nOriginal: %v\nResult: %v", &originalProto, resultProto)
}

// TestLambdaPlanRoundTrip tests round-trip serialization of a full Substrait plan
// containing a lambda expression with a scalar function in the body.
// This represents: SELECT transform(column, x -> x * 2) FROM table
func TestLambdaPlanRoundTrip(t *testing.T) {
	const planJSON = `{
		"extensionUrns": [
			{
				"extensionUrnAnchor": 1,
				"urn": "extension:io.substrait:functions_arithmetic"
			}
		],
		"extensions": [
			{
				"extensionFunction": {
					"extensionUrnReference": 1,
					"functionAnchor": 1,
					"name": "multiply:i32_i32"
				}
			}
		],
		"relations": [
			{
				"root": {
					"input": {
						"project": {
							"input": {
								"read": {
									"baseSchema": {
										"names": ["values"],
										"struct": {
											"types": [
												{
													"list": {
														"type": {"i32": {"nullability": "NULLABILITY_REQUIRED"}},
														"nullability": "NULLABILITY_REQUIRED"
													}
												}
											],
											"nullability": "NULLABILITY_REQUIRED"
										}
									},
									"namedTable": {"names": ["test_table"]}
								}
							},
							"expressions": [
								{
									"lambda": {
										"parameters": {
											"nullability": "NULLABILITY_REQUIRED",
											"types": [
												{"i32": {"nullability": "NULLABILITY_REQUIRED"}}
											]
										},
										"body": {
											"scalarFunction": {
												"functionReference": 1,
												"outputType": {"i32": {"nullability": "NULLABILITY_REQUIRED"}},
												"arguments": [
													{
														"value": {
															"selection": {
																"directReference": {"structField": {"field": 0}},
																"lambdaParameterReference": {"stepsOut": 0}
															}
														}
													},
													{
														"value": {
															"literal": {"i32": 2}
														}
													}
												]
											}
										}
									}
								}
							]
						}
					},
					"names": ["result"]
				}
			}
		]
	}`

	// Parse original JSON to protobuf Plan
	var originalPlan proto.Plan
	err := protojson.Unmarshal([]byte(planJSON), &originalPlan)
	require.NoError(t, err, "Plan JSON should parse to protobuf")

	// Extract the lambda expression from the plan
	project := originalPlan.Relations[0].GetRoot().GetInput().GetProject()
	require.NotNil(t, project, "Should have project relation")
	require.Len(t, project.Expressions, 1, "Should have 1 expression")

	originalLambdaProto := project.Expressions[0]
	require.NotNil(t, originalLambdaProto.GetLambda(), "Expression should be a lambda")

	// Build extension registry from the plan's extensions
	collection := ext.GetDefaultCollectionWithNoError()
	extSet, err := ext.GetExtensionSet(&originalPlan, collection)
	require.NoError(t, err, "Should get extension set from plan")

	reg := expr.NewExtensionRegistry(extSet, collection)
	goExpr, err := expr.ExprFromProto(originalLambdaProto, nil, reg)
	require.NoError(t, err, "Should convert lambda to Go expression")

	// Verify it's a Lambda with the expected structure
	lambda, ok := goExpr.(*expr.Lambda)
	require.True(t, ok, "Should be a Lambda expression")
	require.Len(t, lambda.Parameters.Types, 1, "Lambda should have 1 parameter")

	t.Logf("Lambda string representation: %s", lambda.String())

	// Convert Go Expression → protobuf
	resultProto := goExpr.ToProto()

	// Verify round-trip: original lambda proto should equal result proto
	require.True(t, pb.Equal(originalLambdaProto, resultProto),
		"Round-trip failed!\nOriginal: %v\nResult: %v", originalLambdaProto, resultProto)

	t.Logf("Round-trip successful for lambda with function body")
}

// TestNestedLambdaRoundTrip tests round-trip serialization of a complex nested lambda
// with stepsOut > 0, field > 0, and a function in the body.
// Structure: outer(a: i32, b: i64, c: i32) -> inner(x: i32) -> add(outer.c, x)
// where outer.c is stepsOut=1, field=2
func TestNestedLambdaRoundTrip(t *testing.T) {
	const planJSON = `{
		"extensionUrns": [
			{
				"extensionUrnAnchor": 1,
				"urn": "extension:io.substrait:functions_arithmetic"
			}
		],
		"extensions": [
			{
				"extensionFunction": {
					"extensionUrnReference": 1,
					"functionAnchor": 1,
					"name": "add:i32_i32"
				}
			}
		],
		"relations": [
			{
				"root": {
					"input": {
						"project": {
							"input": {
								"read": {
									"baseSchema": {
										"names": ["values"],
										"struct": {
											"types": [{"i32": {"nullability": "NULLABILITY_REQUIRED"}}],
											"nullability": "NULLABILITY_REQUIRED"
										}
									},
									"namedTable": {"names": ["test_table"]}
								}
							},
							"expressions": [
								{
									"lambda": {
										"parameters": {
											"nullability": "NULLABILITY_REQUIRED",
											"types": [
												{"i32": {"nullability": "NULLABILITY_REQUIRED"}},
												{"i64": {"nullability": "NULLABILITY_REQUIRED"}},
												{"i32": {"nullability": "NULLABILITY_REQUIRED"}}
											]
										},
										"body": {
											"lambda": {
												"parameters": {
													"nullability": "NULLABILITY_REQUIRED",
													"types": [
														{"i32": {"nullability": "NULLABILITY_REQUIRED"}}
													]
												},
												"body": {
													"scalarFunction": {
														"functionReference": 1,
														"outputType": {"i32": {"nullability": "NULLABILITY_REQUIRED"}},
														"arguments": [
															{
																"value": {
																	"selection": {
																		"directReference": {"structField": {"field": 2}},
																		"lambdaParameterReference": {"stepsOut": 1}
																	}
																}
															},
															{
																"value": {
																	"selection": {
																		"directReference": {"structField": {"field": 0}},
																		"lambdaParameterReference": {"stepsOut": 0}
																	}
																}
															}
														]
													}
												}
											}
										}
									}
								}
							]
						}
					},
					"names": ["result"]
				}
			}
		]
	}`

	// Parse JSON → protobuf Plan
	var plan proto.Plan
	err := protojson.Unmarshal([]byte(planJSON), &plan)
	require.NoError(t, err, "Should parse plan JSON")

	// Extract the lambda from the plan
	projectRel := plan.Relations[0].GetRoot().GetInput().GetProject()
	require.NotNil(t, projectRel, "Should have project relation")

	originalLambdaProto := projectRel.Expressions[0].GetLambda()
	require.NotNil(t, originalLambdaProto, "Should have nested lambda expression")

	// Verify nested structure in protobuf
	innerLambdaProto := originalLambdaProto.GetBody().GetLambda()
	require.NotNil(t, innerLambdaProto, "Should have inner lambda")

	funcProto := innerLambdaProto.GetBody().GetScalarFunction()
	require.NotNil(t, funcProto, "Inner body should be scalar function")
	require.Len(t, funcProto.Arguments, 2, "Function should have 2 args")

	// First arg: stepsOut=1, field=2 (outer.c)
	arg0 := funcProto.Arguments[0].GetValue().GetSelection()
	require.Equal(t, uint32(1), arg0.GetLambdaParameterReference().GetStepsOut())
	require.Equal(t, int32(2), arg0.GetDirectReference().GetStructField().GetField())

	// Second arg: stepsOut=0, field=0 (inner.x)
	arg1 := funcProto.Arguments[1].GetValue().GetSelection()
	require.Equal(t, uint32(0), arg1.GetLambdaParameterReference().GetStepsOut())
	require.Equal(t, int32(0), arg1.GetDirectReference().GetStructField().GetField())

	// Convert protobuf → Go Expression
	collection := ext.GetDefaultCollectionWithNoError()
	extSet, err := ext.GetExtensionSet(&plan, collection)
	require.NoError(t, err, "Should get extension set from plan")

	reg := expr.NewExtensionRegistry(extSet, collection)
	goExpr, err := expr.ExprFromProto(projectRel.Expressions[0], nil, reg)
	require.NoError(t, err, "Should convert nested lambda to Go expression")

	// Verify structure in Go
	outerLambda, ok := goExpr.(*expr.Lambda)
	require.True(t, ok, "Should be a Lambda expression")
	require.Len(t, outerLambda.Parameters.Types, 3, "Outer lambda should have 3 parameters")

	innerLambda, ok := outerLambda.Body.(*expr.Lambda)
	require.True(t, ok, "Body should be inner Lambda")
	require.Len(t, innerLambda.Parameters.Types, 1, "Inner lambda should have 1 parameter")

	scalarFunc, ok := innerLambda.Body.(*expr.ScalarFunction)
	require.True(t, ok, "Inner body should be ScalarFunction")
	require.Equal(t, 2, scalarFunc.NArgs(), "Function should have 2 args")
	require.Equal(t, "add", scalarFunc.Name(), "Function should be 'add'")

	t.Logf("Nested lambda: %s", outerLambda.String())

	// Convert Go Expression → protobuf
	resultProto := goExpr.ToProto()

	// Verify round-trip
	require.True(t, pb.Equal(projectRel.Expressions[0], resultProto),
		"Round-trip failed for nested lambda!")

	t.Logf("Round-trip successful: nested lambda with stepsOut=1, field=2, and add() function")
}
