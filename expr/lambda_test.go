// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v7/expr"
	ext "github.com/substrait-io/substrait-go/v7/extensions"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/encoding/protojson"
	pb "google.golang.org/protobuf/proto"
)

// TestLoadLambdaPlan verifies we can parse a full Substrait plan
// containing a lambda expression from JSON into protobuf structures.
func TestBasicLambdaPlan(t *testing.T) {
	// A complete Substrait plan with a lambda expression
	// This represents: SELECT transform([1,2,3], x -> x * 2)
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

	// Parse JSON into protobuf Plan struct
	var plan proto.Plan
	err := protojson.Unmarshal([]byte(planJSON), &plan)
	require.NoError(t, err, "Plan JSON should parse to protobuf")

	// Verify we got the plan structure
	require.Len(t, plan.Relations, 1, "Should have 1 relation")

	// Navigate to the lambda expression
	root := plan.Relations[0].GetRoot()
	require.NotNil(t, root, "Should have root relation")

	project := root.GetInput().GetProject()
	require.NotNil(t, project, "Should have project relation")
	require.Len(t, project.Expressions, 1, "Should have 1 expression")

	// Verify it's a lambda
	lambda := project.Expressions[0].GetLambda()
	require.NotNil(t, lambda, "Expression should be a lambda")

	// Check lambda structure
	require.NotNil(t, lambda.Parameters, "Lambda should have parameters")
	require.Len(t, lambda.Parameters.Types, 1, "Lambda should have 1 parameter")
	require.Equal(t, proto.Type_NULLABILITY_REQUIRED, lambda.Parameters.Nullability, "Lambda parameters should have NULLABILITY_REQUIRED")
	require.NotNil(t, lambda.Parameters.Types[0].GetI32(), "Lambda parameters should have i32 type")
	require.NotNil(t, lambda.Body, "Lambda should have a body")

	t.Logf("Successfully loaded plan with lambda: %v", lambda)
}

// TestLambdaProtoFromJSON verifies that the protobuf-generated code
// can parse lambda expressions from JSON.
func TestLambdaReference(t *testing.T) {
	// A simple lambda expression: (x: i32) -> x
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

	// Inspect the parsed protobuf structure
	lambda := exprProto.GetLambda()
	require.NotNil(t, lambda, "Should have lambda expression")

	params := lambda.GetParameters()
	require.NotNil(t, params, "Should have parameters")
	require.Len(t, params.GetTypes(), 1, "Should have 1 parameter")

	body := lambda.GetBody()
	require.NotNil(t, body, "Should have body")

	// Check the body is a field reference with lambda parameter reference
	selection := body.GetSelection()
	require.NotNil(t, selection, "Body should be a selection/field reference")

	lambdaParamRef := selection.GetLambdaParameterReference()
	require.NotNil(t, lambdaParamRef, "Should have lambda parameter reference")
	require.Equal(t, uint32(0), lambdaParamRef.GetStepsOut(), "StepsOut should be 0")

	// Verify the body's field reference index is valid for the parameters
	directRef := selection.GetDirectReference()
	require.NotNil(t, directRef, "Should have direct reference")

	structField := directRef.GetStructField()
	require.NotNil(t, structField, "Should have struct field reference")

	fieldIndex := structField.GetField()
	numParams := len(params.GetTypes())
	require.Less(t, int(fieldIndex), numParams,
		"Body references parameter %d but lambda only has %d parameters", fieldIndex, numParams)

	// Verify the referenced parameter type matches what we expect
	referencedParamType := params.GetTypes()[fieldIndex]
	require.NotNil(t, referencedParamType.GetI32(),
		"Referenced parameter %d should be i32 type", fieldIndex)

	t.Logf("Successfully parsed lambda reference: %v", &exprProto)
	t.Logf("Body references parameter %d (type: i32)", fieldIndex)
}

// TestLambdaWithFunction tests proto parsing of a lambda with a scalar function
// in the body that references the lambda parameter: (x: i32) -> multiply(x, 2)
func TestLambdaWithFunction(t *testing.T) {
	const lambdaJSON = `{
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
	}`

	var exprProto proto.Expression
	err := protojson.Unmarshal([]byte(lambdaJSON), &exprProto)
	require.NoError(t, err, "JSON should parse to protobuf")

	// Verify lambda structure
	lambda := exprProto.GetLambda()
	require.NotNil(t, lambda, "Should have lambda expression")

	params := lambda.GetParameters()
	require.NotNil(t, params, "Should have parameters")
	require.Len(t, params.GetTypes(), 1, "Should have 1 parameter")
	require.NotNil(t, params.GetTypes()[0].GetI32(), "Parameter should be i32")

	// Verify body is a scalar function
	body := lambda.GetBody()
	require.NotNil(t, body, "Should have body")

	scalarFunc := body.GetScalarFunction()
	require.NotNil(t, scalarFunc, "Body should be a scalar function")
	require.Equal(t, uint32(1), scalarFunc.GetFunctionReference(), "Function reference should be 1")
	require.Len(t, scalarFunc.GetArguments(), 2, "Function should have 2 arguments")

	// Verify first argument is a lambda parameter reference
	arg0 := scalarFunc.GetArguments()[0].GetValue()
	require.NotNil(t, arg0, "First argument should have a value")

	selection := arg0.GetSelection()
	require.NotNil(t, selection, "First arg should be a selection/field reference")

	lambdaParamRef := selection.GetLambdaParameterReference()
	require.NotNil(t, lambdaParamRef, "Should have lambda parameter reference")
	require.Equal(t, uint32(0), lambdaParamRef.GetStepsOut(), "StepsOut should be 0")

	structField := selection.GetDirectReference().GetStructField()
	require.NotNil(t, structField, "Should have struct field reference")
	require.Equal(t, int32(0), structField.GetField(), "Should reference parameter 0")

	// Verify second argument is a literal
	arg1 := scalarFunc.GetArguments()[1].GetValue()
	require.NotNil(t, arg1, "Second argument should have a value")

	literal := arg1.GetLiteral()
	require.NotNil(t, literal, "Second arg should be a literal")
	require.Equal(t, int32(2), literal.GetI32(), "Literal should be 2")

	// Verify output type
	outputType := scalarFunc.GetOutputType()
	require.NotNil(t, outputType.GetI32(), "Output type should be i32")

	t.Logf("Successfully parsed lambda with function: %v", &exprProto)
	t.Logf("Function has %d arguments, first is param ref, second is literal %d",
		len(scalarFunc.GetArguments()), literal.GetI32())
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

	t.Logf("Lambda: %s", lambda.String())
	t.Logf("Body: literal %d, return type: %s", literal.Value, lambdaType.ShortString())
}

// TestLambdaPlanExprFromProto tests converting a full plan with a lambda expression
// containing a function call and parameter reference to Go expressions.
// This represents: (x: i32) -> multiply(x, 2)
func TestLambdaPlanExprFromProto(t *testing.T) {
	const planJSON = `{
		"extensionUris": [
			{
				"extensionUriAnchor": 1,
				"uri": "https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml"
			}
		],
		"extensions": [
			{
				"extensionFunction": {
					"extensionUriReference": 1,
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

	// Verify the type was resolved correctly (Option 2 implementation)
	// The body's type should now be i32, matching the lambda parameter type
	bodyType := lambda.Body.GetType()
	require.NotNil(t, bodyType, "Body type should be resolved")
	require.Equal(t, "i32", bodyType.ShortString(), "Body type should be i32")

	// Lambda.GetType() should also return i32
	lambdaType := lambda.GetType()
	require.NotNil(t, lambdaType, "Lambda type should be resolved")
	require.Equal(t, "i32", lambdaType.ShortString(), "Lambda return type should be i32")

	t.Logf("Lambda Go expression: %s", lambda.String())
	t.Logf("Body references parameter %d via LambdaParameterReference", structFieldRef.Field)
	t.Logf("Body type: %s, Lambda type: %s", bodyType.ShortString(), lambdaType.ShortString())
}

// TestLambdaWithFunctionExprFromProto converts a lambda with a scalar function
// from protobuf to Go expression. Uses a minimal plan to set up extensions.
// Lambda: (x: i32) -> multiply(x, 2)
func TestLambdaWithFunctionExprFromProto(t *testing.T) {
	// Minimal plan with just extensions (needed to resolve function reference)
	const planJSON = `{
		"extensionUris": [
			{
				"extensionUriAnchor": 1,
				"uri": "https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml"
			}
		],
		"extensions": [
			{
				"extensionFunction": {
					"extensionUriReference": 1,
					"functionAnchor": 1,
					"name": "multiply:i32_i32"
				}
			}
		],
		"relations": []
	}`

	// Lambda expression with function body
	const lambdaJSON = `{
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
	}`

	// Parse plan to get extensions
	var plan proto.Plan
	err := protojson.Unmarshal([]byte(planJSON), &plan)
	require.NoError(t, err, "Plan JSON should parse")

	// Parse lambda expression
	var exprProto proto.Expression
	err = protojson.Unmarshal([]byte(lambdaJSON), &exprProto)
	require.NoError(t, err, "Lambda JSON should parse")

	// Build extension registry from plan
	collection := ext.GetDefaultCollectionWithNoError()
	extSet, err := ext.GetExtensionSet(&plan, collection)
	require.NoError(t, err, "Should get extension set")
	reg := expr.NewExtensionRegistry(extSet, collection)

	// Convert protobuf → Go Expression
	goExpr, err := expr.ExprFromProto(&exprProto, nil, reg)
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

	// Verify the field reference type was resolved
	require.NotNil(t, fieldRef.GetType(), "FieldReference type should be resolved")
	require.Equal(t, "i32", fieldRef.GetType().ShortString(), "FieldRef type should be i32")

	// Verify second argument is a literal
	arg1 := scalarFunc.Arg(1)
	literal, ok := arg1.(*expr.PrimitiveLiteral[int32])
	require.True(t, ok, "Second arg should be PrimitiveLiteral[int32]")
	require.Equal(t, int32(2), literal.Value, "Literal should be 2")

	// Verify lambda return type
	lambdaType := lambda.GetType()
	require.NotNil(t, lambdaType, "Lambda type should be resolved")
	require.Equal(t, "i32", lambdaType.ShortString(), "Lambda return type should be i32")

	t.Logf("Lambda: %s", lambda.String())
	t.Logf("Body function: %s, return type: %s", scalarFunc.Name(), lambdaType.ShortString())
}

// TestValidateAllRefs_Valid tests that ValidateAllRefs passes for valid references.
func TestValidateAllRefs_Valid(t *testing.T) {
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
	require.NoError(t, err)

	reg := expr.NewEmptyExtensionRegistry(ext.GetDefaultCollectionWithNoError())
	goExpr, err := expr.ExprFromProto(&exprProto, nil, reg)
	require.NoError(t, err)

	lambda := goExpr.(*expr.Lambda)

	// ValidateAllRefs should pass for valid stepsOut=0 reference
	err = lambda.ValidateAllRefs()
	require.NoError(t, err, "ValidateAllRefs should pass for valid reference")

	t.Logf("Lambda validated successfully: %s", lambda.String())
}

// TestValidateAllRefs_InvalidOuterRef tests that ValidateAllRefs catches
// invalid outer lambda references (stepsOut > 0 with no outer context).
func TestValidateAllRefs_InvalidOuterRef(t *testing.T) {
	// Lambda with stepsOut=1 but no outer lambda exists
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
					"lambdaParameterReference": {"stepsOut": 1}
				}
			}
		}
	}`

	var exprProto proto.Expression
	err := protojson.Unmarshal([]byte(lambdaJSON), &exprProto)
	require.NoError(t, err)

	reg := expr.NewEmptyExtensionRegistry(ext.GetDefaultCollectionWithNoError())
	goExpr, err := expr.ExprFromProto(&exprProto, nil, reg)
	require.NoError(t, err, "ExprFromProto should succeed - outer refs are allowed")

	lambda := goExpr.(*expr.Lambda)

	// ValidateAllRefs should fail because stepsOut=1 but no outer lambdas provided
	err = lambda.ValidateAllRefs()
	require.Error(t, err, "ValidateAllRefs should fail for invalid outer reference")
	require.Contains(t, err.Error(), "stepsOut 1")
	require.Contains(t, err.Error(), "non-existent outer lambda")

	t.Logf("Expected error: %v", err)
}

// TestValidateAllRefs_ValidNestedLambda tests that ValidateAllRefs passes
// when validating from the outermost lambda (recursive validation tracks context).
func TestValidateAllRefs_ValidNestedLambda(t *testing.T) {
	// Outer lambda containing an inner lambda that references outer's parameter
	// The inner lambda references field 0 of the outer lambda (stepsOut=1)
	const nestedLambdaJSON = `{
		"lambda": {
			"parameters": {
				"nullability": "NULLABILITY_REQUIRED",
				"types": [
					{"i64": {"nullability": "NULLABILITY_REQUIRED"}},
					{"i64": {"nullability": "NULLABILITY_REQUIRED"}}
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
						"selection": {
							"directReference": {"structField": {"field": 0}},
							"lambdaParameterReference": {"stepsOut": 1}
						}
					}
				}
			}
		}
	}`

	var exprProto proto.Expression
	err := protojson.Unmarshal([]byte(nestedLambdaJSON), &exprProto)
	require.NoError(t, err)

	reg := expr.NewEmptyExtensionRegistry(ext.GetDefaultCollectionWithNoError())
	goExpr, err := expr.ExprFromProto(&exprProto, nil, reg)
	require.NoError(t, err)

	outerLambda := goExpr.(*expr.Lambda)

	// ValidateAllRefs from the outermost lambda - it recursively validates nested lambdas
	err = outerLambda.ValidateAllRefs()
	require.NoError(t, err, "ValidateAllRefs should pass - inner lambda validly references outer's parameter")

	t.Logf("Nested lambda validated successfully: %s", outerLambda.String())
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
		"extensionUris": [
			{
				"extensionUriAnchor": 1,
				"uri": "https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml"
			}
		],
		"extensions": [
			{
				"extensionFunction": {
					"extensionUriReference": 1,
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
