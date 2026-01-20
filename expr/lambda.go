package expr

import (
	"fmt"
	"strings"

	substraitgo "github.com/substrait-io/substrait-go/v7"
	"github.com/substrait-io/substrait-go/v7/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

type Lambda struct {
	Parameters *types.StructType
	Body       Expression
}

// NewLambda constructs a new Lambda expression with validation.
func NewLambda(parameters *types.StructType, body Expression) (*Lambda, error) {
	if parameters == nil {
		return nil, fmt.Errorf("%w: lambda must have parameters struct", substraitgo.ErrInvalidExpr)
	}
	if body == nil {
		return nil, fmt.Errorf("%w: lambda must have a body expression", substraitgo.ErrInvalidExpr)
	}
	if parameters.Nullability != types.NullabilityRequired {
		return nil, fmt.Errorf("%w: lambda parameters struct must have NULLABILITY_REQUIRED", substraitgo.ErrInvalidExpr)
	}

	// Validate each parameter type is non-nil
	for i, paramType := range parameters.Types {
		if paramType == nil {
			return nil, fmt.Errorf("%w: lambda parameter %d has nil type", substraitgo.ErrInvalidExpr, i)
		}
	}

	// Resolve types for FieldReferences with LambdaParameterReference roots.
	// This is needed because when parsing from protobuf, the body is parsed
	// before we know the lambda parameters, so types can't be resolved then.
	resolvedBody := resolveLambdaParamTypes(body, parameters)

	// Create the lambda
	lambda := &Lambda{Parameters: parameters, Body: resolvedBody}

	// Validate lambda parameter references for stepsOut == 0 only
	// (outer refs can't be validated without context, so skip them at construction)
	if err := validateCurrentLambdaRefs(lambda.Body, lambda.Parameters); err != nil {
		return nil, err
	}

	return lambda, nil
}

// ValidateAllRefs validates ALL lambda parameter references in this lambda and any
// nested lambdas within its body. Call this on the outermost lambda to validate
// the entire tree - it automatically tracks outer lambda context as it recurses.
//
// For top-level lambdas, any stepsOut > 0 will error (no outer lambdas exist).
// For nested lambdas, outer references are validated against the accumulated context.
func (l *Lambda) ValidateAllRefs() error {
	return l.validateRefsRecursive(nil)
}

// validateRefsRecursive validates this lambda and recursively validates nested lambdas.
// outerParams accumulates the parameters of outer lambdas as we recurse inward.
func (l *Lambda) validateRefsRecursive(outerParams []*types.StructType) error {
	// Validate FieldReferences in this lambda's body (not nested lambdas)
	if err := validateFieldRefsInBody(l.Body, l.Parameters, outerParams); err != nil {
		return err
	}

	// Find and recursively validate any nested lambdas
	// When we enter a nested lambda, prepend our params to the outer context
	var validationErr error
	nestedOuterParams := append([]*types.StructType{l.Parameters}, outerParams...)

	l.Body.Visit(func(e Expression) Expression {
		if validationErr != nil {
			return e
		}
		if nestedLambda, ok := e.(*Lambda); ok {
			validationErr = nestedLambda.validateRefsRecursive(nestedOuterParams)
		}
		return e
	})

	// Also check if body itself is a lambda
	if nestedLambda, ok := l.Body.(*Lambda); ok {
		if err := nestedLambda.validateRefsRecursive(nestedOuterParams); err != nil {
			return err
		}
	}

	return validationErr
}

// validateCurrentLambdaRefs validates stepsOut=0 references only (used during construction).
// Outer refs (stepsOut > 0) are skipped since we don't have outer context yet.
func validateCurrentLambdaRefs(body Expression, params *types.StructType) error {
	// If body IS a nested lambda, don't validate it here - nested lambdas validate themselves
	if _, ok := body.(*Lambda); ok {
		return nil
	}

	// Check body itself
	if err := validateFieldRef(body, params, nil, true); err != nil {
		return err
	}
	// Check children (but not nested lambdas - they'll validate themselves)
	var validationErr error
	body.Visit(func(e Expression) Expression {
		if validationErr != nil {
			return e
		}
		if _, ok := e.(*Lambda); ok {
			return e // Skip nested lambdas
		}
		validationErr = validateFieldRef(e, params, nil, true)
		return e
	})
	return validationErr
}

// validateFieldRefsInBody validates all FieldReferences in a lambda body (not nested lambdas).
func validateFieldRefsInBody(body Expression, currentParams *types.StructType, outerParams []*types.StructType) error {
	// If body IS a nested lambda, don't validate it here - it's handled by validateRefsRecursive
	// Otherwise body.Visit() would visit INTO the nested lambda's body with wrong context
	if _, ok := body.(*Lambda); ok {
		return nil
	}

	// Check body itself
	if err := validateFieldRef(body, currentParams, outerParams, false); err != nil {
		return err
	}
	// Check children (but not nested lambdas - they're handled recursively)
	var validationErr error
	body.Visit(func(e Expression) Expression {
		if validationErr != nil {
			return e
		}
		if _, ok := e.(*Lambda); ok {
			return e // Skip nested lambdas - handled by validateRefsRecursive
		}
		validationErr = validateFieldRef(e, currentParams, outerParams, false)
		return e
	})
	return validationErr
}

// validateFieldRef validates a single expression if it's a FieldReference
// with LambdaParameterReference root, checking against current or outer lambda params.
// If skipOuterRefs is true, stepsOut > 0 references are skipped.
func validateFieldRef(e Expression, currentParams *types.StructType, outerParams []*types.StructType, skipOuterRefs bool) error {
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
		// stepsOut > 0: references outer lambda's parameter
		if skipOuterRefs {
			// Skip validation during construction (can't validate without context)
			return nil
		}
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
// FieldReferences that have LambdaParameterReference roots (StepsOut == 0).
// This is needed because when parsing from protobuf, the body is parsed before
// the lambda parameters are known, so types can't be resolved at parse time.
func resolveLambdaParamTypes(body Expression, params *types.StructType) Expression {
	// First, try to resolve the body itself if it's a FieldReference
	resolved := tryResolveFieldRef(body, params)

	// Then walk children via Visit to handle nested FieldReferences
	return resolved.Visit(func(e Expression) Expression {
		return tryResolveFieldRef(e, params)
	})
}

// tryResolveFieldRef attempts to resolve the type of a FieldReference with
// LambdaParameterReference root. Returns the expression unchanged if it's not
// applicable or already resolved.
func tryResolveFieldRef(e Expression, params *types.StructType) Expression {
	fieldRef, ok := e.(*FieldReference)
	if !ok {
		return e
	}

	lambdaRef, ok := fieldRef.Root.(LambdaParameterReference)
	if !ok || lambdaRef.StepsOut != 0 {
		return e // Not a reference to this lambda's parameters
	}

	// Already has a type resolved
	if fieldRef.GetType() != nil {
		return e
	}

	// Resolve the type using the lambda parameters as the base struct type
	if refSeg, ok := fieldRef.Reference.(ReferenceSegment); ok {
		resolvedType, err := refSeg.GetType(params)
		if err != nil {
			return e // Can't resolve, leave as-is
		}
		// Return new FieldReference with resolved type
		return &FieldReference{
			Root:      fieldRef.Root,
			Reference: fieldRef.Reference,
			knownType: resolvedType,
		}
	}
	return e
}

func (l *Lambda) GetParameters() *types.StructType {
	return l.Parameters
}

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
