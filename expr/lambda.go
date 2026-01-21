package expr

import (
	"fmt"
	"strings"

	substraitgo "github.com/substrait-io/substrait-go/v7"
	"github.com/substrait-io/substrait-go/v7/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// Lambda represents a lambda expression with parameters and a body.
type Lambda struct {
	Parameters *types.StructType
	Body       Expression
}

// LambdaBuilder constructs Lambda expressions with full validation.
// Use NewLambdaBuilder() to create a builder, set parameters and body,
// then call Build() to get a validated Lambda.
//
// For nested lambdas, use WithNestedLambda() instead of WithBody().
type LambdaBuilder struct {
	parameters    *types.StructType
	bodyExpr      Expression
	nestedBuilder *LambdaBuilder
}

// NewLambdaBuilder creates a new LambdaBuilder.
func NewLambdaBuilder() *LambdaBuilder {
	return &LambdaBuilder{}
}

// WithParameters sets the lambda's parameters.
func (b *LambdaBuilder) WithParameters(params *types.StructType) *LambdaBuilder {
	b.parameters = params
	return b
}

// WithBody sets the lambda's body to an Expression.
// Use WithNestedLambda() instead if the body is another lambda being built.
func (b *LambdaBuilder) WithBody(body Expression) *LambdaBuilder {
	b.bodyExpr = body
	b.nestedBuilder = nil // Clear alternative
	return b
}

// WithNestedLambda sets the lambda's body to a nested LambdaBuilder.
// The nested lambda will be built with proper outer context when Build() is called.
func (b *LambdaBuilder) WithNestedLambda(nested *LambdaBuilder) *LambdaBuilder {
	b.nestedBuilder = nested
	b.bodyExpr = nil // Clear alternative
	return b
}

// Build constructs and validates the Lambda.
// If the body contains nested LambdaBuilders, they are built first.
// All lambda parameter references are validated, including outer refs (stepsOut > 0).
func (b *LambdaBuilder) Build() (*Lambda, error) {
	return b.buildWithContext(nil)
}

// buildWithContext builds the lambda with outer lambda context for validation.
// outerParams[0] = immediate parent lambda's params, outerParams[1] = grandparent, etc.
func (b *LambdaBuilder) buildWithContext(outerParams []*types.StructType) (*Lambda, error) {
	// Validate parameters
	if b.parameters == nil {
		return nil, fmt.Errorf("%w: lambda must have parameters struct", substraitgo.ErrInvalidExpr)
	}
	if b.parameters.Nullability != types.NullabilityRequired {
		return nil, fmt.Errorf("%w: lambda parameters struct must have NULLABILITY_REQUIRED", substraitgo.ErrInvalidExpr)
	}
	for i, paramType := range b.parameters.Types {
		if paramType == nil {
			return nil, fmt.Errorf("%w: lambda parameter %d has nil type", substraitgo.ErrInvalidExpr, i)
		}
	}

	// Get/build the body
	var body Expression
	switch {
	case b.nestedBuilder != nil:
		// Nested lambda: build it with this lambda as outer context
		nestedContext := append([]*types.StructType{b.parameters}, outerParams...)
		built, err := b.nestedBuilder.buildWithContext(nestedContext) // recursive
		if err != nil {
			return nil, err
		}
		body = built
	case b.bodyExpr != nil:
		body = b.bodyExpr
	default:
		return nil, fmt.Errorf("%w: lambda must have a body expression", substraitgo.ErrInvalidExpr)
	}

	// Resolve types for FieldReferences with LambdaParameterReference roots
	resolvedBody := resolveLambdaParamTypes(body, b.parameters, outerParams)

	// Validate ALL lambda parameter references (stepsOut=0 and stepsOut>0)
	if err := validateAllFieldRefs(resolvedBody, b.parameters, outerParams); err != nil {
		return nil, err
	}

	return &Lambda{Parameters: b.parameters, Body: resolvedBody}, nil
}

// validateAllFieldRefs validates ALL lambda parameter references in the body.
// Unlike construction-time validation, this validates both stepsOut=0 and stepsOut>0.
func validateAllFieldRefs(body Expression, currentParams *types.StructType, outerParams []*types.StructType) error {
	// Handle case where body IS a Lambda (e.g., outer lambda's body is inner lambda)
	if lambda, ok := body.(*Lambda); ok {
		nestedContext := append([]*types.StructType{currentParams}, outerParams...)
		return validateAllFieldRefs(lambda.Body, lambda.Parameters, nestedContext)
	}

	// Validate body itself (if it's a FieldReference)
	if err := validateFieldRef(body, currentParams, outerParams); err != nil {
		return err
	}

	// Validate children - handles case where body CONTAINS lambdas
	// (e.g., body is a ScalarFunction with a Lambda argument)
	var validationErr error
	body.Visit(func(e Expression) Expression {
		if validationErr != nil {
			return e // Already found an error, skip remaining children
		}
		// Nested lambda found in expression tree - validate with updated context
		if lambda, ok := e.(*Lambda); ok {
			nestedContext := append([]*types.StructType{currentParams}, outerParams...)
			validationErr = validateAllFieldRefs(lambda.Body, lambda.Parameters, nestedContext)
			return e
		}
		validationErr = validateFieldRef(e, currentParams, outerParams)
		return e
	})

	return validationErr
}

// validateFieldRef validates a single FieldReference with LambdaParameterReference root.
func validateFieldRef(e Expression, currentParams *types.StructType, outerParams []*types.StructType) error {
	fieldRef, ok := e.(*FieldReference)
	if !ok {
		return nil
	}

	lambdaRef, ok := fieldRef.Root.(LambdaParameterReference)
	if !ok {
		return nil
	}

	// Determine which lambda's parameters to check against
	var targetParams *types.StructType
	if lambdaRef.StepsOut == 0 {
		targetParams = currentParams
	} else {
		// stepsOut 1 = outerParams[0], stepsOut 2 = outerParams[1], etc.
		outerIndex := int(lambdaRef.StepsOut) - 1
		if outerIndex >= len(outerParams) {
			return fmt.Errorf("%w: stepsOut %d references non-existent outer lambda (only %d outer lambdas available)",
				substraitgo.ErrInvalidExpr, lambdaRef.StepsOut, len(outerParams))
		}
		targetParams = outerParams[outerIndex]
	}

	// Check if the reference is a struct field
	structRef, ok := fieldRef.Reference.(*StructFieldRef)
	if !ok {
		return nil
	}

	// Validate the field index is within bounds
	if int(structRef.Field) >= len(targetParams.Types) {
		if lambdaRef.StepsOut == 0 {
			return fmt.Errorf("%w: lambda body references parameter %d but lambda only has %d parameters",
				substraitgo.ErrInvalidExpr, structRef.Field, len(targetParams.Types))
		}
		return fmt.Errorf("%w: lambda body references outer parameter %d (stepsOut=%d) but outer lambda only has %d parameters",
			substraitgo.ErrInvalidExpr, structRef.Field, lambdaRef.StepsOut, len(targetParams.Types))
	}

	return nil
}

// resolveLambdaParamTypes walks the body expression and resolves types for any
// FieldReferences that have LambdaParameterReference roots.
// - stepsOut=0: resolves against params (this lambda's parameters)
// - stepsOut>0: resolves against outerParams (outer lambda parameters)
func resolveLambdaParamTypes(body Expression, params *types.StructType, outerParams []*types.StructType) Expression {
	// First, try to resolve the body itself if it's a FieldReference
	resolved := tryResolveFieldRef(body, params, outerParams)

	// Then walk children via Visit to handle nested FieldReferences
	return resolved.Visit(func(e Expression) Expression {
		return tryResolveFieldRef(e, params, outerParams)
	})
}

// tryResolveFieldRef attempts to resolve the type of a FieldReference with
// LambdaParameterReference root. Returns the expression unchanged if not applicable.
func tryResolveFieldRef(e Expression, currentParams *types.StructType, outerParams []*types.StructType) Expression {
	fieldRef, ok := e.(*FieldReference)
	if !ok {
		return e
	}

	lambdaRef, ok := fieldRef.Root.(LambdaParameterReference)
	if !ok {
		return e // Not a lambda parameter reference
	}

	// Already has a type resolved
	if fieldRef.GetType() != nil {
		return e
	}

	// Determine which lambda's parameters to resolve against
	var targetParams *types.StructType
	if lambdaRef.StepsOut == 0 {
		targetParams = currentParams
	} else {
		// stepsOut 1 = outerParams[0], stepsOut 2 = outerParams[1], etc.
		outerIndex := int(lambdaRef.StepsOut) - 1
		if outerIndex >= len(outerParams) {
			return e // Can't resolve without outer context
		}
		targetParams = outerParams[outerIndex]
	}

	// Guard against out-of-bounds field index before resolving type
	// (validation happens after resolution, so we need to be defensive here)
	if structRef, ok := fieldRef.Reference.(*StructFieldRef); ok {
		if int(structRef.Field) >= len(targetParams.Types) {
			return e // Out of bounds, leave unresolved (validation will catch this)
		}
	}

	// Resolve the type using the target lambda's parameters
	if refSeg, ok := fieldRef.Reference.(ReferenceSegment); ok {
		resolvedType, err := refSeg.GetType(targetParams)
		if err != nil {
			return e // Can't resolve, leave as-is
		}
		return &FieldReference{
			Root:      fieldRef.Root,
			Reference: fieldRef.Reference,
			knownType: resolvedType,
		}
	}
	return e
}

// lambdaFromProto creates a Lambda directly from protobuf without builder validation.
// This is used internally when parsing from protobuf where the structure is already valid.
// Note: Only stepsOut=0 refs are resolved; outer refs can't be resolved without context.
func lambdaFromProto(parameters *types.StructType, body Expression) *Lambda {
	resolvedBody := resolveLambdaParamTypes(body, parameters, nil)
	return &Lambda{Parameters: parameters, Body: resolvedBody}
}

// GetParameters returns the structure defining this lambda's parameters.
func (l *Lambda) GetParameters() *types.StructType {
	return l.Parameters
}

// GetBody returns the expression that forms the body of this lambda.
func (l *Lambda) GetBody() Expression {
	return l.Body
}

func (l *Lambda) String() string {
	var b strings.Builder
	b.WriteString("(")
	for i, t := range l.Parameters.Types {
		if i > 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(&b, "$%d: %s", i, t)
	}
	b.WriteString(") -> ")
	b.WriteString(l.Body.String())
	return b.String()
}

func (l *Lambda) isRootRef() {}

func (l *Lambda) IsScalar() bool {
	return l.Body.IsScalar()
}

func (l *Lambda) GetType() types.Type {
	return l.Body.GetType()
}

func (l *Lambda) Equals(other Expression) bool {
	rhs, ok := other.(*Lambda)
	if !ok {
		return false
	}
	return l.Parameters.Equals(rhs.Parameters) && l.Body.Equals(rhs.Body)
}

func (l *Lambda) ToProto() *proto.Expression {
	children := make([]*proto.Type, len(l.Parameters.Types))
	for i, c := range l.Parameters.Types {
		children[i] = types.TypeToProto(c)
	}
	paramsProto := &proto.Type_Struct{
		Types:                  children,
		TypeVariationReference: l.Parameters.TypeVariationRef,
		Nullability:            l.Parameters.Nullability,
	}

	return &proto.Expression{
		RexType: &proto.Expression_Lambda_{
			Lambda: &proto.Expression_Lambda{
				Parameters: paramsProto,
				Body:       l.Body.ToProto(),
			},
		},
	}
}

func (l *Lambda) ToProtoFuncArg() *proto.FunctionArgument {
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Value{Value: l.ToProto()},
	}
}

func (l *Lambda) Visit(visit VisitFunc) Expression {
	newBody := visit(l.Body)
	if newBody == l.Body {
		return l
	}
	return &Lambda{Parameters: l.Parameters, Body: newBody}
}
