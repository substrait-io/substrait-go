// SPDX-License-Identifier: Apache-2.0

package extensions

import (
	"fmt"

	substraitgo "github.com/substrait-io/substrait-go"
	"github.com/substrait-io/substrait-go/types"
	"github.com/substrait-io/substrait-go/types/parser"
)

type FunctionVariant interface {
	Name() string
	CompoundName() string
	Description() string
	Args() ArgumentList
	Options() map[string]Option
	URI() string
	ResolveType(argTypes []types.Type) (types.Type, error)
	Variadic() *VariadicBehavior
}

func EvaluateTypeExpression(nullHandling NullabilityHandling, expr parser.TypeExpression, paramTypeList ArgumentList, actualTypes []types.Type) (types.Type, error) {
	if len(paramTypeList) != len(actualTypes) {
		return nil, fmt.Errorf("%w: mismatch in number of arguments provided. got %d, expected %d",
			substraitgo.ErrInvalidExpr, len(actualTypes), len(paramTypeList))
	}

	allNonNull := true
	for i, p := range paramTypeList {
		switch p := p.(type) {
		case EnumArg:
			if actualTypes[i] != nil {
				return nil, fmt.Errorf("%w: arg #%d (%s) should be an enum",
					substraitgo.ErrInvalidType, i, p.Name)
			}
		case ValueArg:
			if actualTypes[i] == nil {
				return nil, fmt.Errorf("%w: arg #%d should be of type %s",
					substraitgo.ErrInvalidType, i, p.toTypeString())
			}

			isNullable := actualTypes[i].GetNullability() != types.NullabilityRequired
			if isNullable {
				allNonNull = false
			}

			if nullHandling == DiscreteNullability {
				if t, ok := p.Value.Expr.(*parser.Type); ok {
					if isNullable != t.Optional() {
						return nil, fmt.Errorf("%w: discrete nullability did not match for arg #%d",
							substraitgo.ErrInvalidType, i)
					}
				} else {
					return nil, substraitgo.ErrNotImplemented
				}
			}
		case TypeArg:
			return nil, substraitgo.ErrNotImplemented
		}
	}

	var outType types.Type
	if t, ok := expr.Expr.(*parser.Type); ok {
		var err error
		outType, err = t.Type()
		if err != nil {
			return nil, err
		}
	} else {
		return nil, substraitgo.ErrNotImplemented
	}

	if nullHandling == MirrorNullability || nullHandling == "" {
		if allNonNull {
			return outType.WithNullability(types.NullabilityRequired), nil
		}
		return outType.WithNullability(types.NullabilityNullable), nil
	}

	return outType, nil
}

// NewScalarFuncVariant constructs a variant with the provided name and uri
// and uses the defaults for everything else.
//
// Return expressions aren't included here as using this variant to construct
// an expression requires an output type argument. This is for creating an
// on-the-fly function variant that will not be registered as an extension.
func NewScalarFuncVariant(id ID) *ScalarFunctionVariant {
	return &ScalarFunctionVariant{
		name: id.Name,
		uri:  id.URI,
	}
}

// NewScalarFuncVariantWithProps is the same as NewScalarFuncVariant but allows
// setting the values for the SessionDependant, Variadic Behavior and Deterministic
// properties.
func NewScalarFuncVariantWithProps(id ID, variadic *VariadicBehavior, sessionDependant, deterministic bool) *ScalarFunctionVariant {
	return &ScalarFunctionVariant{
		name: id.Name,
		uri:  id.URI,
		impl: ScalarFunctionImpl{
			Variadic:         variadic,
			SessionDependent: sessionDependant,
			Deterministic:    deterministic,
		},
	}
}

type ScalarFunctionVariant struct {
	name        string
	description string
	uri         string
	impl        ScalarFunctionImpl
}

func (s *ScalarFunctionVariant) Name() string                     { return s.name }
func (s *ScalarFunctionVariant) Description() string              { return s.description }
func (s *ScalarFunctionVariant) Args() ArgumentList               { return s.impl.Args }
func (s *ScalarFunctionVariant) Options() map[string]Option       { return s.impl.Options }
func (s *ScalarFunctionVariant) Variadic() *VariadicBehavior      { return s.impl.Variadic }
func (s *ScalarFunctionVariant) Deterministic() bool              { return s.impl.Deterministic }
func (s *ScalarFunctionVariant) SessionDependent() bool           { return s.impl.SessionDependent }
func (s *ScalarFunctionVariant) Nullability() NullabilityHandling { return s.impl.Nullability }
func (s *ScalarFunctionVariant) URI() string                      { return s.uri }
func (s *ScalarFunctionVariant) ResolveType(argumentTypes []types.Type) (types.Type, error) {
	return EvaluateTypeExpression(s.impl.Nullability, s.impl.Return, s.impl.Args, argumentTypes)
}
func (s *ScalarFunctionVariant) CompoundName() string {
	return s.name + ":" + s.impl.signatureKey()
}

// NewAggFuncVariant constructs a variant with the provided name and uri
// and uses the defaults for everything else.
//
// Return expressions aren't included here as using this variant to construct
// an expression requires an output type argument. This is for creating an
// on-the-fly function variant that will not be registered as an extension.
func NewAggFuncVariant(id ID) *AggregateFunctionVariant {
	return &AggregateFunctionVariant{
		name: id.Name,
		uri:  id.URI,
		impl: AggregateFunctionImpl{
			Decomposable: DecomposeNone,
		},
	}
}

type AggVariantOptions struct {
	Variadic         *VariadicBehavior
	SessionDependant bool
	Deterministic    bool
	Ordered          bool
	// value of 0 == unlimited
	MaxSet       uint
	Decomposable DecomposeType
	// should be a type expression
	// must not be empty if decomposable is not DecomposeNone
	IntermediateOutputType string
}

var (
	defParser, _ = parser.New()
)

func NewAggFuncVariantOpts(id ID, opts AggVariantOptions) *AggregateFunctionVariant {
	var aggIntermediate parser.TypeExpression
	if opts.Decomposable == "" {
		opts.Decomposable = DecomposeNone
	}
	if opts.Decomposable != DecomposeNone {
		if opts.IntermediateOutputType == "" {
			panic(fmt.Errorf("%w: custom Aggregate function variant %s. must provide Intermediate output type",
				substraitgo.ErrInvalidExpr, id))
		}

		intermediate, err := defParser.ParseString(opts.IntermediateOutputType)
		if err != nil {
			panic(err)
		}
		aggIntermediate = *intermediate
	}

	return &AggregateFunctionVariant{
		name: id.Name,
		uri:  id.URI,
		impl: AggregateFunctionImpl{
			ScalarFunctionImpl: ScalarFunctionImpl{
				Variadic:         opts.Variadic,
				SessionDependent: opts.SessionDependant,
				Deterministic:    opts.Deterministic,
			},
			Ordered:      opts.Ordered,
			MaxSet:       int(opts.MaxSet),
			Decomposable: opts.Decomposable,
			Intermediate: aggIntermediate,
		},
	}
}

type AggregateFunctionVariant struct {
	name        string
	description string
	uri         string
	impl        AggregateFunctionImpl
}

func (s *AggregateFunctionVariant) Name() string                     { return s.name }
func (s *AggregateFunctionVariant) Description() string              { return s.description }
func (s *AggregateFunctionVariant) Args() ArgumentList               { return s.impl.Args }
func (s *AggregateFunctionVariant) Options() map[string]Option       { return s.impl.Options }
func (s *AggregateFunctionVariant) Variadic() *VariadicBehavior      { return s.impl.Variadic }
func (s *AggregateFunctionVariant) Deterministic() bool              { return s.impl.Deterministic }
func (s *AggregateFunctionVariant) SessionDependent() bool           { return s.impl.SessionDependent }
func (s *AggregateFunctionVariant) Nullability() NullabilityHandling { return s.impl.Nullability }
func (s *AggregateFunctionVariant) URI() string                      { return s.uri }
func (s *AggregateFunctionVariant) ResolveType(argumentTypes []types.Type) (types.Type, error) {
	return EvaluateTypeExpression(s.impl.Nullability, s.impl.Return, s.impl.Args, argumentTypes)
}
func (s *AggregateFunctionVariant) CompoundName() string {
	return s.name + ":" + s.impl.signatureKey()
}
func (s *AggregateFunctionVariant) Decomposability() DecomposeType { return s.impl.Decomposable }
func (s *AggregateFunctionVariant) Intermediate() (types.Type, error) {
	if t, ok := s.impl.Intermediate.Expr.(*parser.Type); ok {
		return t.Type()
	}
	return nil, fmt.Errorf("%w: bad intermediate type expression", substraitgo.ErrInvalidType)
}
func (s *AggregateFunctionVariant) Ordered() bool { return s.impl.Ordered }
func (s *AggregateFunctionVariant) MaxSet() int   { return s.impl.MaxSet }

type WindowFunctionVariant struct {
	name        string
	description string
	uri         string
	impl        WindowFunctionImpl
}

func NewWindowFuncVariant(id ID) *WindowFunctionVariant {
	return &WindowFunctionVariant{
		name: id.Name,
		uri:  id.URI,
		impl: WindowFunctionImpl{
			AggregateFunctionImpl: AggregateFunctionImpl{
				Decomposable: DecomposeNone,
			},
			WindowType: PartitionWindow,
		},
	}
}

type WindowVariantOpts struct {
	Variadic         *VariadicBehavior
	SessionDependant bool
	Deterministic    bool
	Ordered          bool
	// value of 0 == unlimited
	MaxSet       uint
	Decomposable DecomposeType
	// should be a type expression
	// must not be empty if decomposable is not DecomposeNone
	IntermediateOutputType string
	WindowType             WindowType
}

func NewWindowFuncVariantOpts(id ID, opts WindowVariantOpts) *WindowFunctionVariant {
	var aggIntermediate parser.TypeExpression
	if opts.Decomposable == "" {
		opts.Decomposable = DecomposeNone
	}
	if opts.WindowType == "" {
		opts.WindowType = PartitionWindow
	}
	if opts.Decomposable != DecomposeNone {
		if opts.IntermediateOutputType == "" {
			panic(fmt.Errorf("%w: custom Aggregate function variant %s. must provide Intermediate output type",
				substraitgo.ErrInvalidExpr, id))
		}

		intermediate, err := defParser.ParseString(opts.IntermediateOutputType)
		if err != nil {
			panic(err)
		}
		aggIntermediate = *intermediate
	}

	return &WindowFunctionVariant{
		name: id.Name,
		uri:  id.URI,
		impl: WindowFunctionImpl{
			AggregateFunctionImpl: AggregateFunctionImpl{
				ScalarFunctionImpl: ScalarFunctionImpl{
					Variadic:         opts.Variadic,
					SessionDependent: opts.SessionDependant,
					Deterministic:    opts.Deterministic,
				},
				Ordered:      opts.Ordered,
				MaxSet:       int(opts.MaxSet),
				Decomposable: opts.Decomposable,
				Intermediate: aggIntermediate,
			},
			WindowType: opts.WindowType,
		},
	}
}

func (s *WindowFunctionVariant) Name() string                     { return s.name }
func (s *WindowFunctionVariant) Description() string              { return s.description }
func (s *WindowFunctionVariant) Args() ArgumentList               { return s.impl.Args }
func (s *WindowFunctionVariant) Options() map[string]Option       { return s.impl.Options }
func (s *WindowFunctionVariant) Variadic() *VariadicBehavior      { return s.impl.Variadic }
func (s *WindowFunctionVariant) Deterministic() bool              { return s.impl.Deterministic }
func (s *WindowFunctionVariant) SessionDependent() bool           { return s.impl.SessionDependent }
func (s *WindowFunctionVariant) Nullability() NullabilityHandling { return s.impl.Nullability }
func (s *WindowFunctionVariant) URI() string                      { return s.uri }
func (s *WindowFunctionVariant) ResolveType(argumentTypes []types.Type) (types.Type, error) {
	return EvaluateTypeExpression(s.impl.Nullability, s.impl.Return, s.impl.Args, argumentTypes)
}
func (s *WindowFunctionVariant) CompoundName() string {
	return s.name + ":" + s.impl.signatureKey()
}
func (s *WindowFunctionVariant) Decomposability() DecomposeType { return s.impl.Decomposable }
func (s *WindowFunctionVariant) Intermediate() (types.Type, error) {
	if t, ok := s.impl.Intermediate.Expr.(*parser.Type); ok {
		return t.Type()
	}
	return nil, fmt.Errorf("%w: bad intermediate type expression", substraitgo.ErrInvalidType)
}
func (s *WindowFunctionVariant) Ordered() bool          { return s.impl.Ordered }
func (s *WindowFunctionVariant) MaxSet() int            { return s.impl.MaxSet }
func (s *WindowFunctionVariant) WindowType() WindowType { return s.impl.WindowType }
