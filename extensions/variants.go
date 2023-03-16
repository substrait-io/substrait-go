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
					if isNullable != t.Optional {
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
func (s *AggregateFunctionVariant) Decomposability() DecomposeType      { return s.impl.Decomposable }
func (s *AggregateFunctionVariant) Intermediate() parser.TypeExpression { return s.impl.Intermediate }

type WindowFunctionVariant struct {
	name        string
	description string
	uri         string
	impl        WindowFunctionImpl
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
