// SPDX-License-Identifier: Apache-2.0

package expr

import (
	"fmt"

	substraitgo "github.com/substrait-io/substrait-go/v9"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/types"
	"github.com/substrait-io/substrait-go/v9/types/parser"
)

type aggregateVariant interface {
	*extensions.AggregateFunctionVariant | *extensions.WindowFunctionVariant
	ResolveType([]types.Type, extensions.Set) (types.Type, error)
	Args() extensions.FuncParameterList
	Variadic() *extensions.VariadicBehavior
	URN() string
	Nullability() extensions.NullabilityHandling
	Decomposability() extensions.DecomposeType
	Intermediate() (types.FuncDefArgType, error)
	ReturnType() types.FuncDefArgType
}

func resolveAggregateVariant[T aggregateVariant](
	id extensions.FunctionID, reg ExtensionRegistry, getter func(extensions.FunctionID) (T, bool),
	phase types.AggregationPhase, args []types.FuncArg,
) (T, types.Type, error) {
	var initial, intermediateOutput bool
	switch phase {
	case types.AggPhaseInitialToResult:
		initial = true
	case types.AggPhaseInitialToIntermediate:
		initial, intermediateOutput = true, true
	case types.AggPhaseIntermediateToIntermediate:
		intermediateOutput = true
	case types.AggPhaseIntermediateToResult, types.AggPhaseUnspecified:
		// The protobuf definition specifies that UNSPECIFIED implies INTERMEDIATE_TO_RESULT.
	default:
		return nil, nil, fmt.Errorf("%w: invalid aggregation phase %d", substraitgo.ErrInvalidExpr, phase)
	}

	argTypes := functionArgumentTypes(args)
	var decl T
	if initial {
		var err error
		decl, err = lookupVariant(id, reg, getter, argTypes)
		if err != nil {
			return nil, nil, err
		}
	} else {
		var found bool
		decl, found = getter(id)
		if !found {
			// Intermediate states do not identify the original signature. For example,
			// avg:i32 and avg:i64 both consume struct<i64,i64> but return different types.
			return nil, nil, fmt.Errorf("%w: intermediate-input phase requires a matching function id (use the original compound name): %s", substraitgo.ErrNotFound, id)
		}
	}
	if decl.Decomposability() == extensions.DecomposeNone && phase != types.AggPhaseInitialToResult {
		return nil, nil, fmt.Errorf("%w: non-decomposable window or agg function '%s' must use InitialToResult phase", substraitgo.ErrInvalidExpr, id)
	}
	if initial && !intermediateOutput {
		out, err := decl.ResolveType(argTypes, reg.Set)
		return decl, out, err
	}

	intermediate, err := decl.Intermediate()
	if err != nil {
		return nil, nil, err
	}
	output := decl.ReturnType()
	if intermediateOutput {
		output = intermediate
	}
	if output == nil {
		return nil, nil, fmt.Errorf("%w: missing output type for function %s", substraitgo.ErrInvalidType, id)
	}
	parameters, variadic := decl.Args(), decl.Variadic()
	if !initial {
		if len(args) != 1 {
			return nil, nil, fmt.Errorf("%w: intermediate-input phase requires one state argument", substraitgo.ErrInvalidExpr)
		}
		if _, ok := args[0].(Expression); !ok {
			return nil, nil, fmt.Errorf("%w: intermediate state must be a value argument", substraitgo.ErrInvalidExpr)
		}
		if derivation, ok := intermediate.(*types.OutputDerivation); ok && len(derivation.Assignments) != 0 {
			return nil, nil, fmt.Errorf("%w: cannot infer original parameters from a derived intermediate type", substraitgo.ErrNotImplemented)
		}
		// Each incoming value holds the complete state, including for functions
		// whose initial signature takes zero or multiple arguments.
		parameters = extensions.FuncParameterList{extensions.ValueArg{Value: &parser.TypeExpression{ValueType: intermediate}}}
		variadic = nil
	}
	out, err := extensions.EvaluateTypeExpression(decl.URN(), decl.Nullability(), output, parameters, variadic, argTypes, reg.Set)
	if err != nil {
		return nil, nil, err
	}
	return decl, out, nil
}
