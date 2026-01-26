// SPDX-License-Identifier: Apache-2.0

package expr

import (
	"fmt"

	substraitgo "github.com/substrait-io/substrait-go/v7"
	"github.com/substrait-io/substrait-go/v7/extensions"
	"github.com/substrait-io/substrait-go/v7/types"
)

// Builder is a basic interface for any type which can construct
// an expression. The `Build` method will be reserved for producing
// a concrete type while `BuildExpr` will exist for compatibility
// with this interface for ease of use. Typically it will be
// implemented as a simply a call to Build anyways.
type Builder interface {
	BuildExpr() (Expression, error)
}

// ExprBuilder is the parent context for all the expression builders.
// It maintains a pointer to an extension registry and, optionally,
// a pointer to a base input schema and a slice of lambda parameters.
// This allows less verbose expression building as it isn't necessary
// to pass these to every `New*` function to construct the expressions.
//
// This is intended to be used like:
//
//	    b := expr.ExprBuilder{
//		       Reg: ...,
//	        BaseSchema: ...,
//	    }
//	    e, err := b.ScalarFunc(fnID, options...).Args(
//	        b.RootRef(expr.NewStructFieldRef(1)),
//	        b.ScalarFunc(fn2ID, options2...).Args(
//	            b.Wrap(expr.NewLiteral(int32(5), false /* nullable type */)),
//	            b.RootRef(expr.NewStructFieldRef(2))))
//
// See the unit tests for additional examples / constructs.
type ExprBuilder struct {
	Reg           ExtensionRegistry
	BaseSchema    *types.RecordType
	lambdaContext []*types.StructType // keeps track of the lambda context for nested lambdas
}

// Literal returns a wrapped literal that can be passed as an argument
// to any of the other expression builders such as ScalarFunc.Args.
func (e *ExprBuilder) Literal(l Literal) exprWrapper {
	return exprWrapper{l, nil}
}

// Expression returns a wrapped expression that can be passed as an argument
// to any of the other expression builders such as ScalarFunc.Args.
func (e *ExprBuilder) Expression(expr Expression) exprWrapper {
	return exprWrapper{expr, nil}
}

// Wrap is like Literal or Expression but allows propagating an error
// (such as when calling expr.NewLiteral) that will bubble up when attempting
// to build an expression so it doesn't get swallowed or force a panic.
func (e *ExprBuilder) Wrap(l Literal, err error) exprWrapper {
	return exprWrapper{l, err}
}

// Enum wraps a string representing an Enum argument to a function being
// built.
func (e *ExprBuilder) Enum(val string) enumWrapper { return enumWrapper(val) }

// ScalarFunc returns a builder for the scalar function represented by the
// passed in ID and options. Use the Args method to add arguments to this
// builder. Validity of the ID, argument types and number of arguments will
// be checked at the point that the Build method is called to construct
// the final expression and will return an error if invalid.
//
// The extension registry inside of ExprBuilder will be used to resolve
// the ID, but only at the point at which Build is called. Therefore this
// can be called before actually loading the extensions as long as the
// extension identified by the ID is loaded into the registry *before*
// `Build` is called.
func (e *ExprBuilder) ScalarFunc(
	id extensions.ID, opts ...*types.FunctionOption,
) *scalarFuncBuilder {
	return &scalarFuncBuilder{
		b:    e,
		id:   id,
		opts: opts,
	}
}

// WindowFunc returns a builder for the window function represented by the
// passed in ID and options. Other properties such as Arguments,
// aggregation phase, invocation, sort fields, etc. can be then added via
// individual methods on the returned builder. Validity of the ID, argument
// types and number of arguments will be checked at the point that the
// Build method is called to construct the final expression and will return
// an error if invalid.
//
// The extension registry inside of ExprBuilder will be used to resolve
// the ID, but only at the point at which Build is called. Therefore this
// can be called before actually loading the extensions as long as the
// extension identified by the ID is loaded into the registry *before*
// `Build` is called.
func (e *ExprBuilder) WindowFunc(
	id extensions.ID, opts ...*types.FunctionOption,
) *windowFuncBuilder {
	return &windowFuncBuilder{
		b:    e,
		id:   id,
		opts: opts,
	}
}

// AggFunc returns a builder for the aggregate function represented by the
// passed in ID and options. Other properties such as Arguments,
// aggregation phase, invocation, sort fields, etc. can be then added via
// individual methods on the returned builder. Validity of the ID, argument
// types and number of arguments will be checked at the point that the
// Build method is called to construct the final expression and will return
// an error if invalid.
//
// The extension registry inside of ExprBuilder will be used to resolve
// the ID, but only at the point at which Build is called. Therefore this
// can be called before actually loading the extensions as long as the
// extension identified by the ID is loaded into the registry *before*
// `Build` is called.
func (e *ExprBuilder) AggFunc(
	id extensions.ID, opts ...*types.FunctionOption,
) *aggregateFuncBuilder {
	return &aggregateFuncBuilder{
		b:    e,
		id:   id,
		opts: opts,
	}
}

// Ref constructs a field reference with the provided root and reference
// type. When `Build` is called on the returned builder, the `BaseSchema`
// in ExprBuilder will be used to resolve the type of the expression if
// relevant (such as a StructFieldRef/ListRef/MapKeyRef).
func (e *ExprBuilder) Ref(root RootRefType, ref Reference) *fieldRefBuilder {
	return &fieldRefBuilder{
		b: e, root: root, ref: ref,
	}
}

// RootRef is a convenience method equivalent to calling ExprBuilder.Ref
// with `expr.RootReference` as the first argument.
func (e *ExprBuilder) RootRef(ref Reference) *fieldRefBuilder {
	return e.Ref(RootReference, ref)
}

// Cast returns a builder for constructing a Cast expression. The failure
// behavior can be specified by calling FailBehavior before calling Build.
func (e *ExprBuilder) Cast(from Builder, to types.Type) *castBuilder {
	return &castBuilder{
		toType: to, input: from,
	}
}

// Lambda returns a builder for constructing a Lambda expression with the
// given parameters. The body can be set via the Body method, which accepts
// any Builder (expressions, functions, or nested lambdas).
//
// When building nested lambdas (e.g., a function that takes a lambda argument
// which itself references outer lambda parameters), the ExprBuilder maintains
// a context stack that allows inner lambdas to validate stepsOut references
// against outer lambda parameters.
func (e *ExprBuilder) Lambda(params *types.StructType, body Builder) *lambdaBuilder {
	return &lambdaBuilder{
		b:      e,
		params: params,
		body:   body,
	}
}

type lambdaBuilder struct {
	b      *ExprBuilder
	params *types.StructType
	body   Builder
}

// Build constructs and validates the Lambda expression.
// During building, this lambda's parameters are pushed onto the ExprBuilder's
// context stack, allowing nested lambdas to validate stepsOut references.
func (lb *lambdaBuilder) Build() (*Lambda, error) {
	if lb.params == nil {
		return nil, fmt.Errorf("%w: lambda must have parameters struct", substraitgo.ErrInvalidExpr)
	}
	if lb.params.Nullability != types.NullabilityRequired {
		return nil, fmt.Errorf("%w: lambda parameters struct must have NULLABILITY_REQUIRED", substraitgo.ErrInvalidExpr)
	}
	for i, paramType := range lb.params.Types {
		if paramType == nil {
			return nil, fmt.Errorf("%w: lambda parameter %d has nil type", substraitgo.ErrInvalidExpr, i)
		}
	}

	if lb.body == nil {
		return nil, fmt.Errorf("%w: lambda must have a body", substraitgo.ErrInvalidExpr)
	}

	// Push this lambda's params onto context stack before building body.
	// This allows nested lambdas to validate stepsOut references against
	// outer lambda parameters.
	lb.b.lambdaContext = append(lb.b.lambdaContext, lb.params)

	bodyExpr, err := lb.body.BuildExpr()

	// Pop our params from context stack (always, even on error)
	lb.b.lambdaContext = lb.b.lambdaContext[:len(lb.b.lambdaContext)-1]

	if err != nil {
		return nil, err
	}

	// Resolve types for FieldReferences with LambdaParameterReference roots.
	resolvedBody := resolveLambdaParamTypes(bodyExpr, lb.params, lb.b.lambdaContext)

	// Validate ALL lambda parameter references (stepsOut=0 and stepsOut>0)
	if err := validateAllFieldRefs(resolvedBody, lb.params, lb.b.lambdaContext); err != nil {
		return nil, err
	}

	return &Lambda{Parameters: lb.params, Body: resolvedBody}, nil
}

// BuildExpr implements the Builder interface.
func (lb *lambdaBuilder) BuildExpr() (Expression, error) {
	return lb.Build()
}

// BuildFuncArg implements the FuncArgBuilder interface, allowing lambdas
// to be passed directly as arguments to function builders.
func (lb *lambdaBuilder) BuildFuncArg() (types.FuncArg, error) {
	return lb.Build()
}

type exprWrapper struct {
	expression Expression
	err        error
}

func (e exprWrapper) BuildFuncArg() (types.FuncArg, error) { return e.expression, e.err }
func (e exprWrapper) BuildExpr() (Expression, error)       { return e.expression, e.err }

type enumWrapper string

func (e enumWrapper) BuildFuncArg() (types.FuncArg, error) { return types.Enum(e), nil }

type FuncArgBuilder interface {
	BuildFuncArg() (types.FuncArg, error)
}

type castBuilder struct {
	toType          types.Type
	input           Builder
	failureBehavior types.CastFailBehavior
}

func (cb *castBuilder) BuildExpr() (Expression, error)       { return cb.Build() }
func (cb *castBuilder) BuildFuncArg() (types.FuncArg, error) { return cb.Build() }
func (cb *castBuilder) Build() (*Cast, error) {
	in, err := cb.input.BuildExpr()
	if err != nil {
		return nil, err
	}

	return &Cast{
		Type:            cb.toType,
		Input:           in,
		FailureBehavior: cb.failureBehavior,
	}, nil
}

// FailBehavior sets the failure behavior for the resulting Cast expression
// that is built from this builder by calling the Build method.
func (cb *castBuilder) FailBehavior(b types.CastFailBehavior) *castBuilder {
	cb.failureBehavior = b
	return cb
}

type scalarFuncBuilder struct {
	b *ExprBuilder

	id   extensions.ID
	opts []*types.FunctionOption
	args []FuncArgBuilder
}

func (sb *scalarFuncBuilder) Build() (*ScalarFunction, error) {
	var err error
	args := make([]types.FuncArg, len(sb.args))
	for i, a := range sb.args {
		if args[i], err = a.BuildFuncArg(); err != nil {
			return nil, err
		}
	}

	return NewScalarFunc(sb.b.Reg, sb.id, sb.opts, args...)
}

func (sb *scalarFuncBuilder) BuildExpr() (Expression, error) {
	return sb.Build()
}

func (sb *scalarFuncBuilder) BuildFuncArg() (types.FuncArg, error) {
	return sb.Build()
}

// Args sets the argument list for this builder. Subsequent calls to Args
// will *replace* the argument list, not append to it.
func (sb *scalarFuncBuilder) Args(args ...FuncArgBuilder) *scalarFuncBuilder {
	sb.args = args
	return sb
}

type windowFuncBuilder struct {
	b *ExprBuilder

	id   extensions.ID
	opts []*types.FunctionOption
	args []FuncArgBuilder

	phase      types.AggregationPhase
	invocation types.AggregationInvocation
	partitions []Builder
	sortList   []SortField

	boundsType             types.BoundsType
	lowerBound, upperBound Bound
}

func (wb *windowFuncBuilder) Build() (*WindowFunction, error) {
	var err error
	args := make([]types.FuncArg, len(wb.args))
	for i, a := range wb.args {
		if args[i], err = a.BuildFuncArg(); err != nil {
			return nil, err
		}
	}

	parts := make([]Expression, len(wb.partitions))
	for i, p := range wb.partitions {
		if parts[i], err = p.BuildExpr(); err != nil {
			return nil, err
		}
	}

	wf, err := NewWindowFunc(wb.b.Reg, wb.id, wb.opts, wb.invocation, wb.phase, args...)
	if err != nil {
		return nil, err
	}

	wf.Sorts, wf.BoundsType, wf.LowerBound, wf.UpperBound = wb.sortList, wb.boundsType, wb.lowerBound, wb.upperBound
	wf.Partitions = parts

	if err := wf.validate(); err != nil {
		return nil, err
	}

	return wf, nil
}

func (wb *windowFuncBuilder) BuildFuncArg() (types.FuncArg, error) {
	return wb.Build()
}

func (wb *windowFuncBuilder) BuildExpr() (Expression, error) {
	return wb.Build()
}

// Args sets the argument list for this builder. Subsequent calls to Args
// will *replace* the argument list, not append to it.
func (wb *windowFuncBuilder) Args(args ...FuncArgBuilder) *windowFuncBuilder {
	wb.args = args
	return wb
}

// Phase sets the aggregation phase for the resulting WindowFunction
// expression that will be built by this builder.
func (wb *windowFuncBuilder) Phase(p types.AggregationPhase) *windowFuncBuilder {
	wb.phase = p
	return wb
}

// Invocation will set the Aggregation Invocation property for the
// resulting WindowFunction expression that will be built by this builder.
func (wb *windowFuncBuilder) Invocation(i types.AggregationInvocation) *windowFuncBuilder {
	wb.invocation = i
	return wb
}

// Sort sets the list of sort fields for this WindowFunction. Subsequent
// calls to Sort will replace the set of sort fields, not append to it.
func (wb *windowFuncBuilder) Sort(fields ...SortField) *windowFuncBuilder {
	wb.sortList = fields
	return wb
}

// Partitions sets the list of partitions for this WindowFunction. Subsequent
// calls to Partitions will replace the set of partitions, not append to it.
// This expects to receive other Builders and will validate that they produce
// valid expressions without errors at the time that `Build` is called.
func (wb *windowFuncBuilder) Partitions(parts ...Builder) *windowFuncBuilder {
	wb.partitions = parts
	return wb
}

func (wb *windowFuncBuilder) Bounds(lower, upper Bound) *windowFuncBuilder {
	wb.lowerBound, wb.upperBound = lower, upper
	return wb
}

// BoundsType sets the bounds type for this WindowFunction which specifies
// how the window frame is interpreted (ROWS vs RANGE).
func (wb *windowFuncBuilder) BoundsType(bt types.BoundsType) *windowFuncBuilder {
	wb.boundsType = bt
	return wb
}

type aggregateFuncBuilder struct {
	b *ExprBuilder

	id   extensions.ID
	opts []*types.FunctionOption
	args []FuncArgBuilder

	phase      types.AggregationPhase
	invocation types.AggregationInvocation
	sortList   []SortField
}

func (ab *aggregateFuncBuilder) Build() (*AggregateFunction, error) {
	var err error
	args := make([]types.FuncArg, len(ab.args))
	for i, a := range ab.args {
		if args[i], err = a.BuildFuncArg(); err != nil {
			return nil, err
		}
	}

	return NewAggregateFunc(ab.b.Reg, ab.id, ab.opts, ab.invocation, ab.phase, ab.sortList, args...)
}

// Args sets the argument list for this builder. Subsequent calls to Args
// will *replace* the argument list, not append to it.
func (ab *aggregateFuncBuilder) Args(args ...FuncArgBuilder) *aggregateFuncBuilder {
	ab.args = args
	return ab
}

// Phase sets the aggregation phase for the resulting Aggregate Function
// that will be built by this builder.
func (ab *aggregateFuncBuilder) Phase(p types.AggregationPhase) *aggregateFuncBuilder {
	ab.phase = p
	return ab
}

// Invocation will set the Aggregation Invocation property for the
// resulting AggregateFunction that will be built by this builder.
func (ab *aggregateFuncBuilder) Invocation(i types.AggregationInvocation) *aggregateFuncBuilder {
	ab.invocation = i
	return ab
}

// Sort sets the list of sort fields for this AggregateFunction. Subsequent
// calls to Sort will replace the set of sort fields, not append to it.
func (ab *aggregateFuncBuilder) Sorts(fields ...SortField) *aggregateFuncBuilder {
	ab.sortList = fields
	return ab
}

type fieldRefBuilder struct {
	b *ExprBuilder

	root RootRefType
	ref  Reference
}

func (rb *fieldRefBuilder) Build() (*FieldReference, error) {
	return NewFieldRef(rb.root, rb.ref, rb.b.BaseSchema)
}

func (rb *fieldRefBuilder) BuildFuncArg() (types.FuncArg, error) {
	return rb.Build()
}

func (rb *fieldRefBuilder) BuildExpr() (Expression, error) {
	return rb.Build()
}
