// SPDX-License-Identifier: Apache-2.0

package extensions

import (
	"fmt"
	"strings"

	substraitgo "github.com/substrait-io/substrait-go/v8"
	"github.com/substrait-io/substrait-go/v8/types"
	"github.com/substrait-io/substrait-go/v8/types/integer_parameters"
	"github.com/substrait-io/substrait-go/v8/types/parser"
)

type FunctionVariant interface {
	Name() string
	CompoundName() string
	Description() string
	Args() FuncParameterList
	Options() map[string]Option
	URN() string
	// Metadata returns the arbitrary metadata from the extension YAML for this function.
	// Returns nil if no metadata was provided. Callers must not modify the returned map.
	Metadata() map[string]any
	// ResolveType computes the return type of a function variant, given the input argument types
	ResolveType(argTypes []types.Type, registry Set) (types.Type, error)
	Variadic() *VariadicBehavior
	// Match this function matches input arguments against this functions parameter list
	// returns (true, nil) if all input argument can type replace the function definition argument
	// returns (false, err) for invalid input argument. For e.g. if input argument nullability is not correctly
	// set this function will return error
	// returns (false, nil) valid input arguments and argument list type replace parameter list
	Match(argumentTypes []types.Type) (bool, error)
	// MatchAt this function matches input argument at position against definition of this
	// functions argument at same position
	// returns (true, nil) if input argument can type replace the function definition argument
	// returns (false, err) for invalid input argument. For e.g. if input arg position is negative or
	// argument nullability is not correctly set this function will return error
	// returns (false, nil) valid input argument type and argument can't type replace parameter at argPos
	MatchAt(typ types.Type, pos int) (bool, error)
	// MinArgumentCount returns minimum number of arguments required for this function
	MinArgumentCount() int
	// MaxArgumentCount returns minimum number of arguments accepted by this function
	MaxArgumentCount() int
}

func validateType(funcParameter FuncParameter, actual types.Type, idx int, nullHandling NullabilityHandling) (bool, error) {
	allNonNull := true
	switch p := funcParameter.(type) {
	case EnumArg:
		if actual != types.CommonEnumType {
			return allNonNull, fmt.Errorf("%w: arg #%d (%s) should be an enum",
				substraitgo.ErrInvalidType, idx, p.Name)
		}
	case ValueArg:
		if actual == nil {
			return allNonNull, fmt.Errorf("%w: arg #%d should be of type %s",
				substraitgo.ErrInvalidType, idx, p.toTypeString())
		}

		isNullable := actual.GetNullability() != types.NullabilityRequired
		if isNullable {
			allNonNull = false
		}

		if nullHandling == DiscreteNullability {
			if isNullable != (p.Value.ValueType.GetNullability() == types.NullabilityNullable) {
				return allNonNull, fmt.Errorf("%w: discrete nullability did not match for arg #%d",
					substraitgo.ErrInvalidType, idx)
			}
		}
	case TypeArg:
		return allNonNull, substraitgo.ErrNotImplemented
	}

	return allNonNull, nil
}

// EvaluateTypeExpression evaluates the function return type given the input argumentTypes
//
//	urn: the urn of the extension that defines the function. for functions that return user defined types, we assume the urn of the return type is the same as the urn of the function.
//	funcParameters: the function parameters as defined in the function signature in the extension
//	argumentTypes: the actual argument types provided to the function
//	registry: the Set of extensions to look up/add user defined types to
func EvaluateTypeExpression(urn string, nullHandling NullabilityHandling, returnTypeExpr types.FuncDefArgType,
	funcParameters FuncParameterList, variadic *VariadicBehavior, argumentTypes []types.Type, registry Set) (types.Type, error) {
	if variadic != nil {
		numVariadicArgs := len(argumentTypes) - (len(funcParameters) - 1)
		if numVariadicArgs < 0 {
			return nil, fmt.Errorf("%w: mismatch in number of arguments provided. got %d, expected at least %d",
				substraitgo.ErrInvalidExpr, len(argumentTypes), len(funcParameters)-1)
		}
		if !variadic.IsValidArgumentCount(numVariadicArgs) {
			return nil, fmt.Errorf("%w: mismatch in number of arguments provided, invalid number of variadic params. got %d total",
				substraitgo.ErrInvalidExpr, len(argumentTypes))
		}
	} else if len(funcParameters) != len(argumentTypes) {
		return nil, fmt.Errorf("%w: mismatch in number of arguments provided. got %d, expected %d",
			substraitgo.ErrInvalidExpr, len(argumentTypes), len(funcParameters))
	}

	allNonNull := true
	for i, p := range funcParameters {
		if i >= len(argumentTypes) {
			break
		}
		nonNull, err := validateType(p, argumentTypes[i], i, nullHandling)
		if err != nil {
			return nil, err
		}
		allNonNull = allNonNull && nonNull
	}

	// validate variadic argument consistency
	if variadic != nil && len(argumentTypes) > len(funcParameters) && variadic.ParameterConsistency == ConsistentParams {
		nparams := len(funcParameters)
		lastParam := funcParameters[nparams-1]
		for i, actual := range argumentTypes[nparams:] {
			nonNull, err := validateType(lastParam, actual, nparams+i, nullHandling)
			if err != nil {
				return nil, err
			}
			allNonNull = allNonNull && nonNull
		}
	}

	// validate non variadic arguments
	isMatch, err := matchArguments(nullHandling, funcParameters, variadic, argumentTypes)
	if err != nil {
		return nil, err
	}
	if !isMatch {
		return nil, fmt.Errorf("%w: argument types did not match", substraitgo.ErrInvalidType)
	}

	funcParameterTypes := make([]types.FuncDefArgType, len(funcParameters))
	for i, p := range funcParameters {
		funcParameterTypes[i] = p.GetTypeExpression()
	}
	outType, err := returnTypeExpr.ReturnType(funcParameterTypes, argumentTypes)
	if err != nil {
		return nil, err
	}

	// If the return type expression is a ParameterizedUserDefinedType, we need to
	// fill in the TypeReference since ParameterizedUserDefinedType.ReturnType()
	// doesn't have access to the registry to set it itself.
	// For other types like AnyType, the TypeReference is already correctly set.
	if udt, ok := outType.(*types.UserDefinedType); ok {
		if paramUDT, ok := returnTypeExpr.(*types.ParameterizedUserDefinedType); ok {
			udt.TypeReference = registry.GetTypeAnchor(ID{Name: paramUDT.Name, URN: urn})
		}
	}

	if nullHandling == MirrorNullability || nullHandling == "" {
		if allNonNull {
			return outType.WithNullability(types.NullabilityRequired), nil
		}
		return outType.WithNullability(types.NullabilityNullable), nil
	}

	if nullHandling == DeclaredOutputNullability {
		if anyReturn, ok := returnTypeExpr.(*types.AnyType); ok && anyReturn.Nullability == types.NullabilityRequired {
			return outType.WithNullability(types.NullabilityRequired), nil
		}
	}

	return outType, nil
}

func matchArguments(nullability NullabilityHandling, paramTypeList FuncParameterList, variadicBehavior *VariadicBehavior, actualTypes []types.Type) (bool, error) {
	if variadicBehavior == nil && len(actualTypes) != len(paramTypeList) {
		return false, nil
	} else if variadicBehavior != nil {
		numNonVariadicArgs := len(paramTypeList) - 1
		if !validateVariadicBehaviorForMatch(variadicBehavior, actualTypes[numNonVariadicArgs:]) {
			return false, nil
		}
	}
	funcDefArgList, err := getFuncDefFromArgList(paramTypeList)
	if err != nil {
		return false, nil
	}

	matcher := newArgumentMatcher(nullability, funcDefArgList, variadicBehavior)
	for argPos, actualType := range actualTypes {
		match, err := matcher.matchArgument(actualType, argPos)
		if err != nil {
			return false, err
		}
		if !match {
			return false, nil
		}
	}
	if HasSyncParams(funcDefArgList) {
		return types.AreSyncTypeParametersMatching(funcDefArgList, actualTypes), nil
	}
	return true, nil
}

func matchArgumentAt(actualType types.Type, argPos int, nullability NullabilityHandling, paramTypeList FuncParameterList, variadicBehavior *VariadicBehavior) (bool, error) {
	if argPos < 0 {
		return false, fmt.Errorf("non-zero argument position")
	}
	funcDefArgList, err := getFuncDefFromArgList(paramTypeList)
	if err != nil {
		return false, nil
	}
	return newArgumentMatcher(nullability, funcDefArgList, variadicBehavior).matchArgument(actualType, argPos)
}

// argumentMatcher holds the state needed to compare invocation argument types against a function variant signature.
type argumentMatcher struct {
	nullability      NullabilityHandling
	funcDefArgList   []types.FuncDefArgType
	variadicBehavior *VariadicBehavior
	bindings         map[string]types.Type
}

// newArgumentMatcher creates a matcher for one function variant signature.
func newArgumentMatcher(nullability NullabilityHandling, funcDefArgList []types.FuncDefArgType, variadicBehavior *VariadicBehavior) *argumentMatcher {
	return &argumentMatcher{
		nullability:      nullability,
		funcDefArgList:   funcDefArgList,
		variadicBehavior: variadicBehavior,
		bindings:         make(map[string]types.Type),
	}
}

// matchArgument checks whether an argument type matches the signature parameter at the same position.
func (m *argumentMatcher) matchArgument(actualType types.Type, argPos int) (bool, error) {
	funcDefArg, ok := m.parameterAt(argPos)
	if !ok {
		return false, nil
	}

	return m.matchTopLevel(funcDefArg, actualType)
}

func (m *argumentMatcher) matchTopLevel(paramType types.FuncDefArgType, argType types.Type) (bool, error) {
	switch m.nullability {
	case DiscreteNullability:
		return m.matchNested(paramType, argType)
	case MirrorNullability, DeclaredOutputNullability:
		return m.matchTopLevelWithoutNullability(paramType, argType)
	}
	return false, fmt.Errorf("invalid nullability type: %s", m.nullability)
}

func (m *argumentMatcher) matchTopLevelWithoutNullability(paramType types.FuncDefArgType, argType types.Type) (bool, error) {
	if anyType, ok := paramType.(*types.AnyType); ok {
		if anyType.Nullability == types.NullabilityNullable && argType.GetNullability() != types.NullabilityNullable {
			return false, nil
		}
		return m.bindAnyType(anyType, argType, true)
	}
	return m.matchType(paramType, argType, false)
}

func (m *argumentMatcher) matchNested(paramType types.FuncDefArgType, argType types.Type) (bool, error) {
	if anyType, ok := paramType.(*types.AnyType); ok {
		if anyType.Nullability == types.NullabilityNullable && argType.GetNullability() != types.NullabilityNullable {
			return false, nil
		}
		return m.bindAnyType(anyType, argType, anyType.Nullability != types.NullabilityRequired)
	}
	if paramType.GetNullability() != argType.GetNullability() {
		return false, nil
	}
	return m.matchType(paramType, argType, true)
}

func (m *argumentMatcher) matchType(paramType types.FuncDefArgType, argType types.Type, hasMatchedOuterNullability bool) (bool, error) {
	switch p := paramType.(type) {
	case *types.ParameterizedListType:
		listType, ok := argType.(*types.ListType)
		if !ok {
			return false, nil
		}
		return m.matchNested(p.Type, listType.Type)
	case *types.ParameterizedMapType:
		mapType, ok := argType.(*types.MapType)
		if !ok {
			return false, nil
		}
		if match, err := m.matchNested(p.Key, mapType.Key); !match || err != nil {
			return match, err
		}
		return m.matchNested(p.Value, mapType.Value)
	case *types.ParameterizedStructType:
		structType, ok := argType.(*types.StructType)
		if !ok || len(p.Types) != len(structType.Types) {
			return false, nil
		}
		for i, fieldType := range p.Types {
			if match, err := m.matchNested(fieldType, structType.Types[i]); !match || err != nil {
				return match, err
			}
		}
		return true, nil
	case *types.ParameterizedFuncType:
		funcType, ok := argType.(*types.FuncType)
		if !ok || len(p.Parameters) != len(funcType.ParameterTypes) {
			return false, nil
		}
		for i, parameterType := range p.Parameters {
			if match, err := m.matchNested(parameterType, funcType.ParameterTypes[i]); !match || err != nil {
				return match, err
			}
		}
		return m.matchNested(p.Return, funcType.ReturnType)
	}

	if hasMatchedOuterNullability {
		return paramType.MatchWithNullability(argType), nil
	}
	return paramType.MatchWithoutNullability(argType), nil
}

func (m *argumentMatcher) bindAnyType(anyType *types.AnyType, argType types.Type, ignoreNullability bool) (bool, error) {
	if anyType.Name == "any" {
		return true, nil
	}

	boundType := argType
	if boundType.GetNullability() == types.NullabilityUnspecified || ignoreNullability {
		boundType = boundType.WithNullability(types.NullabilityRequired)
	}

	if existingType, exists := m.bindings[anyType.Name]; exists {
		if !existingType.Equals(boundType) {
			return false, fmt.Errorf("%w: type parameter %s cannot be both %s and %s",
				substraitgo.ErrInvalidType, anyType.Name,
				existingType.ShortString(), boundType.ShortString())
		}
	} else {
		m.bindings[anyType.Name] = boundType
	}
	return true, nil
}

// parameterAt returns the signature parameter for an argument position, accounting for variadic signatures.
func (m *argumentMatcher) parameterAt(argPos int) (types.FuncDefArgType, bool) {
	if argPos < 0 {
		return nil, false
	}
	if m.variadicBehavior == nil {
		if argPos >= len(m.funcDefArgList) {
			return nil, false
		}
		return m.funcDefArgList[argPos], true
	}
	if !m.variadicBehavior.IsValidArgumentPosition(argPos) {
		return nil, false
	}
	if argPos < len(m.funcDefArgList) {
		return m.funcDefArgList[argPos], true
	}
	return m.funcDefArgList[len(m.funcDefArgList)-1], true
}

// validateVariadicBehaviorForMatch validates variadic argument counts and enforces
// ConsistentParams when required (all variadic args must have the same type).
func validateVariadicBehaviorForMatch(variadicBehavior *VariadicBehavior, actualTypes []types.Type) bool {
	if !variadicBehavior.IsValidArgumentCount(len(actualTypes)) {
		return false
	}
	// verify consistency of variadic behavior
	if variadicBehavior.ParameterConsistency == ConsistentParams {
		// all concrete types must be equal for all variable arguments
		firstVariadicArgIdx := max(variadicBehavior.Min-1, 0)
		for i := firstVariadicArgIdx; i < len(actualTypes)-1; i++ {
			if !actualTypes[i].Equals(actualTypes[i+1].WithNullability(actualTypes[i].GetNullability())) {
				return false
			}
		}
	}
	return true
}

func getFuncDefFromArgList(paramTypeList FuncParameterList) ([]types.FuncDefArgType, error) {
	var out []types.FuncDefArgType
	for argPos, param := range paramTypeList {
		switch paramType := param.(type) {
		case ValueArg:
			out = append(out, paramType.Value.ValueType)
		case EnumArg:
			out = append(out, types.CommonEnumType)
		case TypeArg:
			return nil, fmt.Errorf("%w: invalid argument at position %d for match operation", substraitgo.ErrInvalidType, argPos)
		default:
			return nil, fmt.Errorf("%w: invalid argument at position %d for match operation", substraitgo.ErrInvalidType, argPos)
		}
	}
	return out, nil
}

func parseFuncName(compoundName string) (name string, args FuncParameterList) {
	name, argsStr, _ := strings.Cut(compoundName, ":")
	if len(argsStr) == 0 {
		return name, nil
	}
	splitArgs := strings.Split(argsStr, "_")
	for _, argStr := range splitArgs {
		parsed, err := parser.ParseType(argStr)
		if err != nil {
			panic(err)
		}
		exp := ValueArg{Name: name, Value: &parser.TypeExpression{ValueType: parsed}}
		args = append(args, exp)
	}

	return name, args
}

func minArgumentCount(paramTypeList FuncParameterList, variadicBehavior *VariadicBehavior) int {
	if variadicBehavior == nil {
		return len(paramTypeList)
	}
	return len(paramTypeList) + variadicBehavior.Min
}

func maxArgumentCount(paramTypeList FuncParameterList, variadicBehavior *VariadicBehavior) int {
	if variadicBehavior == nil {
		return len(paramTypeList)
	}
	return len(paramTypeList) + variadicBehavior.Max
}

// NewScalarFuncVariant constructs a variant with the provided name and urn
// and uses the defaults for everything else.
//
// Return expressions aren't included here as using this variant to construct
// an expression requires an output type argument. This is for creating an
// on-the-fly function variant that will not be registered as an extension.
func NewScalarFuncVariant(id ID) *ScalarFunctionVariant {
	simpleName, args := parseFuncName(id.Name)
	return &ScalarFunctionVariant{
		name: simpleName,
		urn:  id.URN,
		impl: ScalarFunctionImpl{Args: args},
	}
}

// NewScalarFuncVariantWithProps is the same as NewScalarFuncVariant but allows
// setting the values for the SessionDependant, Variadic Behavior and Deterministic
// properties.
func NewScalarFuncVariantWithProps(id ID, variadic *VariadicBehavior, sessionDependant, deterministic bool) *ScalarFunctionVariant {
	simpleName, args := parseFuncName(id.Name)
	return &ScalarFunctionVariant{
		name: simpleName,
		urn:  id.URN,
		impl: ScalarFunctionImpl{
			Args:             args,
			Variadic:         variadic,
			SessionDependent: sessionDependant,
			Deterministic:    deterministic,
		},
	}
}

type ScalarFunctionVariant struct {
	name        string
	description string
	urn         string
	impl        ScalarFunctionImpl
	metadata    map[string]any
}

func (s *ScalarFunctionVariant) Name() string                     { return s.name }
func (s *ScalarFunctionVariant) Description() string              { return s.description }
func (s *ScalarFunctionVariant) Args() FuncParameterList          { return s.impl.Args }
func (s *ScalarFunctionVariant) Options() map[string]Option       { return s.impl.Options }
func (s *ScalarFunctionVariant) Variadic() *VariadicBehavior      { return s.impl.Variadic }
func (s *ScalarFunctionVariant) Deterministic() bool              { return s.impl.Deterministic }
func (s *ScalarFunctionVariant) SessionDependent() bool           { return s.impl.SessionDependent }
func (s *ScalarFunctionVariant) Nullability() NullabilityHandling { return s.impl.Nullability }
func (s *ScalarFunctionVariant) URN() string                      { return s.urn }
func (s *ScalarFunctionVariant) Metadata() map[string]any         { return s.metadata }
func (s *ScalarFunctionVariant) ResolveType(argumentTypes []types.Type, registry Set) (types.Type, error) {
	return EvaluateTypeExpression(s.urn, s.impl.Nullability, s.impl.Return.ValueType, s.impl.Args, s.impl.Variadic, argumentTypes, registry)
}
func (s *ScalarFunctionVariant) CompoundName() string {
	return s.name + ":" + s.impl.signatureKey()
}
func (s *ScalarFunctionVariant) ID() ID {
	return ID{URN: s.urn, Name: s.CompoundName()}
}

func (s *ScalarFunctionVariant) Match(argumentTypes []types.Type) (bool, error) {
	return matchArguments(s.Nullability(), s.impl.Args, s.impl.Variadic, argumentTypes)
}

func (s *ScalarFunctionVariant) MatchAt(typ types.Type, pos int) (bool, error) {
	return matchArgumentAt(typ, pos, s.Nullability(), s.impl.Args, s.impl.Variadic)
}

func (s *ScalarFunctionVariant) MinArgumentCount() int {
	return minArgumentCount(s.impl.Args, s.impl.Variadic)
}

func (s *ScalarFunctionVariant) MaxArgumentCount() int {
	return maxArgumentCount(s.impl.Args, s.impl.Variadic)
}

// NewAggFuncVariant constructs a variant with the provided name and urn
// and uses the defaults for everything else.
//
// Return expressions aren't included here as using this variant to construct
// an expression requires an output type argument. This is for creating an
// on-the-fly function variant that will not be registered as an extension.
func NewAggFuncVariant(id ID) *AggregateFunctionVariant {
	simpleName, args := parseFuncName(id.Name)
	return &AggregateFunctionVariant{
		name: simpleName,
		urn:  id.URN,
		impl: AggregateFunctionImpl{
			ScalarFunctionImpl: ScalarFunctionImpl{
				Args: args,
			},
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

		intermediate, err := parser.ParseType(opts.IntermediateOutputType)
		if err != nil {
			panic(err)
		}
		aggIntermediate.ValueType = intermediate
	}

	simpleName, args := parseFuncName(id.Name)
	return &AggregateFunctionVariant{
		name: simpleName,
		urn:  id.URN,
		impl: AggregateFunctionImpl{
			ScalarFunctionImpl: ScalarFunctionImpl{
				Args:             args,
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
	urn         string
	impl        AggregateFunctionImpl
	metadata    map[string]any
}

func (s *AggregateFunctionVariant) Name() string                     { return s.name }
func (s *AggregateFunctionVariant) Description() string              { return s.description }
func (s *AggregateFunctionVariant) Args() FuncParameterList          { return s.impl.Args }
func (s *AggregateFunctionVariant) Options() map[string]Option       { return s.impl.Options }
func (s *AggregateFunctionVariant) Variadic() *VariadicBehavior      { return s.impl.Variadic }
func (s *AggregateFunctionVariant) Deterministic() bool              { return s.impl.Deterministic }
func (s *AggregateFunctionVariant) SessionDependent() bool           { return s.impl.SessionDependent }
func (s *AggregateFunctionVariant) Nullability() NullabilityHandling { return s.impl.Nullability }
func (s *AggregateFunctionVariant) URN() string                      { return s.urn }
func (s *AggregateFunctionVariant) Metadata() map[string]any         { return s.metadata }
func (s *AggregateFunctionVariant) ResolveType(argumentTypes []types.Type, registry Set) (types.Type, error) {
	return EvaluateTypeExpression(s.urn, s.impl.Nullability, s.impl.Return.ValueType, s.impl.Args, s.impl.Variadic, argumentTypes, registry)
}
func (s *AggregateFunctionVariant) CompoundName() string {
	return s.name + ":" + s.impl.signatureKey()
}
func (s *AggregateFunctionVariant) ID() ID {
	return ID{URN: s.urn, Name: s.CompoundName()}
}
func (s *AggregateFunctionVariant) Decomposability() DecomposeType { return s.impl.Decomposable }
func (s *AggregateFunctionVariant) Intermediate() (types.FuncDefArgType, error) {
	if s.impl.Intermediate.ValueType != nil {
		return s.impl.Intermediate.ValueType, nil
	}
	return nil, fmt.Errorf("%w: bad intermediate type expression", substraitgo.ErrInvalidType)
}
func (s *AggregateFunctionVariant) Ordered() bool { return s.impl.Ordered }
func (s *AggregateFunctionVariant) MaxSet() int   { return s.impl.MaxSet }
func (s *AggregateFunctionVariant) Match(argumentTypes []types.Type) (bool, error) {
	return matchArguments(s.Nullability(), s.impl.Args, s.impl.Variadic, argumentTypes)
}
func (s *AggregateFunctionVariant) MatchAt(typ types.Type, pos int) (bool, error) {
	return matchArgumentAt(typ, pos, s.Nullability(), s.impl.Args, s.impl.Variadic)
}
func (s *AggregateFunctionVariant) MinArgumentCount() int {
	return minArgumentCount(s.impl.Args, s.impl.Variadic)
}

func (s *AggregateFunctionVariant) MaxArgumentCount() int {
	return maxArgumentCount(s.impl.Args, s.impl.Variadic)
}

type WindowFunctionVariant struct {
	name        string
	description string
	urn         string
	impl        WindowFunctionImpl
	metadata    map[string]any
}

func NewWindowFuncVariant(id ID) *WindowFunctionVariant {
	simpleName, args := parseFuncName(id.Name)
	return &WindowFunctionVariant{
		name: simpleName,
		urn:  id.URN,
		impl: WindowFunctionImpl{
			AggregateFunctionImpl: AggregateFunctionImpl{
				ScalarFunctionImpl: ScalarFunctionImpl{Args: args},
				Decomposable:       DecomposeNone,
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

		intermediate, err := parser.ParseType(opts.IntermediateOutputType)
		if err != nil {
			panic(err)
		}
		aggIntermediate.ValueType = intermediate
	}

	simpleName, args := parseFuncName(id.Name)
	return &WindowFunctionVariant{
		name: simpleName,
		urn:  id.URN,
		impl: WindowFunctionImpl{
			AggregateFunctionImpl: AggregateFunctionImpl{
				ScalarFunctionImpl: ScalarFunctionImpl{
					Args:             args,
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
func (s *WindowFunctionVariant) Args() FuncParameterList          { return s.impl.Args }
func (s *WindowFunctionVariant) Options() map[string]Option       { return s.impl.Options }
func (s *WindowFunctionVariant) Variadic() *VariadicBehavior      { return s.impl.Variadic }
func (s *WindowFunctionVariant) Deterministic() bool              { return s.impl.Deterministic }
func (s *WindowFunctionVariant) SessionDependent() bool           { return s.impl.SessionDependent }
func (s *WindowFunctionVariant) Nullability() NullabilityHandling { return s.impl.Nullability }
func (s *WindowFunctionVariant) URN() string                      { return s.urn }
func (s *WindowFunctionVariant) Metadata() map[string]any         { return s.metadata }
func (s *WindowFunctionVariant) ResolveType(argumentTypes []types.Type, registry Set) (types.Type, error) {
	return EvaluateTypeExpression(s.urn, s.impl.Nullability, s.impl.Return.ValueType, s.impl.Args, s.impl.Variadic, argumentTypes, registry)
}
func (s *WindowFunctionVariant) CompoundName() string {
	return s.name + ":" + s.impl.signatureKey()
}
func (s *WindowFunctionVariant) ID() ID {
	return ID{URN: s.urn, Name: s.CompoundName()}
}
func (s *WindowFunctionVariant) Decomposability() DecomposeType { return s.impl.Decomposable }
func (s *WindowFunctionVariant) Intermediate() (types.FuncDefArgType, error) {
	if s.impl.Intermediate.ValueType != nil {
		return s.impl.Intermediate.ValueType, nil
	}
	return nil, fmt.Errorf("%w: bad intermediate type expression", substraitgo.ErrInvalidType)
}
func (s *WindowFunctionVariant) Ordered() bool          { return s.impl.Ordered }
func (s *WindowFunctionVariant) MaxSet() int            { return s.impl.MaxSet }
func (s *WindowFunctionVariant) WindowType() WindowType { return s.impl.WindowType }
func (s *WindowFunctionVariant) Match(argumentTypes []types.Type) (bool, error) {
	return matchArguments(s.Nullability(), s.impl.Args, s.impl.Variadic, argumentTypes)
}
func (s *WindowFunctionVariant) MatchAt(typ types.Type, pos int) (bool, error) {
	return matchArgumentAt(typ, pos, s.Nullability(), s.impl.Args, s.impl.Variadic)
}

func (s *WindowFunctionVariant) MinArgumentCount() int {
	return minArgumentCount(s.impl.Args, s.impl.Variadic)
}

func (s *WindowFunctionVariant) MaxArgumentCount() int {
	return maxArgumentCount(s.impl.Args, s.impl.Variadic)
}

// HasSyncParams This API returns if params share a leaf param name
func HasSyncParams(params []types.FuncDefArgType) bool {
	// if any of the leaf parameters are same, it indicates parameters are same across parameters
	existingParamMap := make(map[string]bool)
	for _, p := range params {
		if !p.HasParameterizedParam() {
			// not a type which contains abstract parameters, so continue
			continue
		}
		// get list of parameterized parameters
		// parameterized param can be a Leaf or another type. If another type we recurse to find leaf
		abstractParams := p.GetParameterizedParams()
		var leafParams []string
		for _, abstractParam := range abstractParams {
			leafParams = append(leafParams, getLeafParameterizedParams(abstractParam)...)
		}
		// if map contains any of the leaf params, parameters are synced
		for _, leafParam := range leafParams {
			if _, ok := existingParamMap[leafParam]; ok {
				return true
			}
		}
		// add all params to map, kindly note we can't add these params
		// in previous loop to avoid having same leaf abstract type in same param
		// e.g. Decimal<P, P> has no sync param
		for _, leafParam := range leafParams {
			existingParamMap[leafParam] = true
		}
	}
	return false
}

// from a parameterized type, get the leaf parameters
// an parameterized param can be a leaf type (e.g. P) or a parameterized type (e.g. VARCHAR<L1>) itself
// if it is a leaf type, its param name is returned
// if it is parameterized type, leaf type is found recursively
func getLeafParameterizedParams(abstractTypes interface{}) []string {
	if leaf, ok := abstractTypes.(integer_parameters.IntegerParameter); ok {
		return []string{leaf.String()}
	}
	// if it is not a leaf type recurse
	if pat, ok := abstractTypes.(types.FuncDefArgType); ok {
		var outLeafParams []string
		for _, p := range pat.GetParameterizedParams() {
			childLeafParams := getLeafParameterizedParams(p)
			outLeafParams = append(outLeafParams, childLeafParams...)
		}
		return outLeafParams
	}
	// invalid type
	panic("invalid non-leaf, non-parameterized type param")
}

// ValidateConstrainedAnyTypeConsistency validates that all uses of the same AnyN parameter
// (e.g., any1, any2, etc.) resolve to the same concrete type across all arguments
func ValidateConstrainedAnyTypeConsistency(funcParameters []types.FuncDefArgType, argumentTypes []types.Type, variadicBehavior *VariadicBehavior) error {
	matcher := newArgumentMatcher(MirrorNullability, funcParameters, variadicBehavior)
	for argPos, argumentType := range argumentTypes {
		match, err := matcher.matchArgument(argumentType, argPos)
		if err != nil {
			return err
		}
		if !match {
			return fmt.Errorf("%w: argument types did not match", substraitgo.ErrInvalidType)
		}
	}
	return nil
}

// expandedParameters repeats the variadic parameter so parameter and argument slices can be compared positionally.
func (m *argumentMatcher) expandedParameters(argumentTypes []types.Type) []types.FuncDefArgType {
	if m.variadicBehavior == nil || len(argumentTypes) <= len(m.funcDefArgList) {
		return m.funcDefArgList
	}

	numVariadicArgs := len(argumentTypes) - len(m.funcDefArgList) + 1
	nonVariadicParams := m.funcDefArgList[:len(m.funcDefArgList)-1]
	variadicParam := m.funcDefArgList[len(m.funcDefArgList)-1]

	expandedFuncParameters := make([]types.FuncDefArgType, len(nonVariadicParams)+numVariadicArgs)
	copy(expandedFuncParameters, nonVariadicParams)
	for i := range numVariadicArgs {
		expandedFuncParameters[len(nonVariadicParams)+i] = variadicParam
	}
	return expandedFuncParameters
}
