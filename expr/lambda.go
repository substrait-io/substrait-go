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

	// Recursively validate all descendants - handles deeply nested expressions
	// (e.g., Cast(FieldRef), ScalarFunction(Cast(FieldRef)))
	var validationErr error
	var recurse func(e Expression) Expression
	recurse = func(e Expression) Expression {
		if validationErr != nil {
			return e // Already found an error, skip remaining
		}
		// Nested lambda found - validate with updated context
		if lambda, ok := e.(*Lambda); ok {
			nestedContext := append([]*types.StructType{currentParams}, outerParams...)
			validationErr = validateAllFieldRefs(lambda.Body, lambda.Parameters, nestedContext)
			return e
		}
		// Validate this node
		validationErr = validateFieldRef(e, currentParams, outerParams)
		if validationErr != nil {
			return e
		}
		// Recurse into this node's children
		return e.Visit(recurse)
	}
	body.Visit(recurse)

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

	// Lambda parameters are a struct, so the first reference segment must be StructFieldRef
	structRef, ok := fieldRef.Reference.(*StructFieldRef)
	if !ok {
		return fmt.Errorf("%w: lambda parameter reference must use StructFieldRef, got %T",
			substraitgo.ErrInvalidExpr, fieldRef.Reference)
	}

	// Validate the field index is within bounds
	if int(structRef.Field) >= len(targetParams.Types) {
		if lambdaRef.StepsOut == 0 {
			return fmt.Errorf("%w: lambda body references parameter %d but lambda only has %d parameters",
				substraitgo.ErrInvalidExpr, structRef.Field, len(targetParams.Types))
		}
		return fmt.Errorf("%w: lambda body references outer parameter %d (stepsOut=%d) but outer lambda only has %d parameters",
			substraitgo.ErrInvalidExpr, structRef.Field, lambdaRef.StepsOut, len(targetParams.Types))
	} else if int(structRef.Field) < 0 {
		return fmt.Errorf("%w: lambda body references negative field index %d",
			substraitgo.ErrInvalidExpr, structRef.Field)
	}

	return nil
}

// resolveLambdaParamTypes walks the body expression and resolves types for any
// FieldReferences that have LambdaParameterReference roots.
// - stepsOut=0: resolves against params (this lambda's parameters)
// - stepsOut>0: resolves against outerParams (outer lambda parameters)
// For nested lambdas, recursively resolves with updated outer context.
// Recurses into all descendants (e.g., Cast(FieldRef), ScalarFunction(Cast(FieldRef))).
func resolveLambdaParamTypes(body Expression, params *types.StructType, outerParams []*types.StructType) Expression {
	// Handle if body itself is a nested lambda - must check BEFORE Visit
	// because Lambda.Visit would bypass our nested lambda context handling
	if nestedLambda, ok := body.(*Lambda); ok {
		nestedContext := append([]*types.StructType{params}, outerParams...)
		resolvedNestedBody := resolveLambdaParamTypes(nestedLambda.Body, nestedLambda.Parameters, nestedContext)
		if resolvedNestedBody != nestedLambda.Body {
			return &Lambda{
				Parameters: nestedLambda.Parameters,
				Body:       resolvedNestedBody,
			}
		}
		return body
	}

	// Try to resolve the body itself if it's a FieldReference
	resolved := tryResolveFieldRef(body, params, outerParams)

	// Recursively walk all descendants via Visit
	var recurse func(e Expression) Expression
	recurse = func(e Expression) Expression {
		// Handle nested lambdas found in expression tree (e.g., as function args)
		if nestedLambda, ok := e.(*Lambda); ok {
			nestedContext := append([]*types.StructType{params}, outerParams...)
			resolvedNestedBody := resolveLambdaParamTypes(nestedLambda.Body, nestedLambda.Parameters, nestedContext)
			if resolvedNestedBody != nestedLambda.Body {
				return &Lambda{
					Parameters: nestedLambda.Parameters,
					Body:       resolvedNestedBody,
				}
			}
			return e
		}
		// Try to resolve this node, then recurse into its children
		resolvedNode := tryResolveFieldRef(e, params, outerParams)
		return resolvedNode.Visit(recurse)
	}
	return resolved.Visit(recurse)
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
		if int(structRef.Field) >= len(targetParams.Types) || int(structRef.Field) < 0 {
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
