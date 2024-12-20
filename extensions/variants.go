// SPDX-License-Identifier: Apache-2.0

package extensions

import (
	"fmt"
	"strings"

	substraitgo "github.com/substrait-io/substrait-go/v3"
	"github.com/substrait-io/substrait-go/v3/types"
	"github.com/substrait-io/substrait-go/v3/types/integer_parameters"
	"github.com/substrait-io/substrait-go/v3/types/parser"
)

type FunctionVariant interface {
	Name() string
	CompoundName() string
	Description() string
	Args() FuncParameterList
	Options() map[string]Option
	URI() string
	ResolveType(argTypes []types.Type) (types.Type, error)
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
//	funcParameters: the function parameters as defined in the function signature in the extension
//	argumentTypes: the actual argument types provided to the function
func EvaluateTypeExpression(nullHandling NullabilityHandling, returnTypeExpr types.FuncDefArgType,
	funcParameters FuncParameterList, variadic *VariadicBehavior, argumentTypes []types.Type) (types.Type, error) {
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

	if nullHandling == MirrorNullability || nullHandling == "" {
		if allNonNull {
			return outType.WithNullability(types.NullabilityRequired), nil
		}
		return outType.WithNullability(types.NullabilityNullable), nil
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
	// loop over actualTypes and not params since actualTypes can be more than params
	// considering variadic type
	for argPos := range actualTypes {
		match, err1 := matchArgumentAtCommon(actualTypes[argPos], argPos, nullability, funcDefArgList, variadicBehavior)
		if err1 != nil {
			return false, err1
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
	return matchArgumentAtCommon(actualType, argPos, nullability, funcDefArgList, variadicBehavior)
}

func matchArgumentAtCommon(actualType types.Type, argPos int, nullability NullabilityHandling, funcDefArgList []types.FuncDefArgType, variadicBehavior *VariadicBehavior) (bool, error) {
	// check if argument out of range
	if variadicBehavior == nil && argPos >= len(funcDefArgList) {
		return false, nil
	} else if variadicBehavior != nil && !variadicBehavior.IsValidArgumentPosition(argPos) {
		// this argument position can't be more than the max allowed by the variadic behavior
		return false, nil
	}

	// if argPos is >= len(funcDefArgList) than last funcDefArg type should be considered for type match
	// already checked for parameter in range above (considering variadic) so no need to check again for variadic
	var funcDefArg types.FuncDefArgType
	if argPos < len(funcDefArgList) {
		funcDefArg = funcDefArgList[argPos]
	} else {
		funcDefArg = funcDefArgList[len(funcDefArgList)-1]
	}
	switch nullability {
	case DiscreteNullability:
		return funcDefArg.MatchWithNullability(actualType), nil
	case MirrorNullability, DeclaredOutputNullability:
		return funcDefArg.MatchWithoutNullability(actualType), nil
	}
	// unreachable case
	return false, fmt.Errorf("invalid nullability type: %s", nullability)
}

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

// NewScalarFuncVariant constructs a variant with the provided name and uri
// and uses the defaults for everything else.
//
// Return expressions aren't included here as using this variant to construct
// an expression requires an output type argument. This is for creating an
// on-the-fly function variant that will not be registered as an extension.
func NewScalarFuncVariant(id ID) *ScalarFunctionVariant {
	simpleName, args := parseFuncName(id.Name)
	return &ScalarFunctionVariant{
		name: simpleName,
		uri:  id.URI,
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
		uri:  id.URI,
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
	uri         string
	impl        ScalarFunctionImpl
}

func (s *ScalarFunctionVariant) Name() string                     { return s.name }
func (s *ScalarFunctionVariant) Description() string              { return s.description }
func (s *ScalarFunctionVariant) Args() FuncParameterList          { return s.impl.Args }
func (s *ScalarFunctionVariant) Options() map[string]Option       { return s.impl.Options }
func (s *ScalarFunctionVariant) Variadic() *VariadicBehavior      { return s.impl.Variadic }
func (s *ScalarFunctionVariant) Deterministic() bool              { return s.impl.Deterministic }
func (s *ScalarFunctionVariant) SessionDependent() bool           { return s.impl.SessionDependent }
func (s *ScalarFunctionVariant) Nullability() NullabilityHandling { return s.impl.Nullability }
func (s *ScalarFunctionVariant) URI() string                      { return s.uri }
func (s *ScalarFunctionVariant) ResolveType(argumentTypes []types.Type) (types.Type, error) {
	return EvaluateTypeExpression(s.impl.Nullability, s.impl.Return.ValueType, s.impl.Args, s.impl.Variadic, argumentTypes)
}
func (s *ScalarFunctionVariant) CompoundName() string {
	return s.name + ":" + s.impl.signatureKey()
}
func (s *ScalarFunctionVariant) ID() ID {
	return ID{URI: s.uri, Name: s.CompoundName()}
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

// NewAggFuncVariant constructs a variant with the provided name and uri
// and uses the defaults for everything else.
//
// Return expressions aren't included here as using this variant to construct
// an expression requires an output type argument. This is for creating an
// on-the-fly function variant that will not be registered as an extension.
func NewAggFuncVariant(id ID) *AggregateFunctionVariant {
	simpleName, args := parseFuncName(id.Name)
	return &AggregateFunctionVariant{
		name: simpleName,
		uri:  id.URI,
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
		uri:  id.URI,
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
	uri         string
	impl        AggregateFunctionImpl
}

func (s *AggregateFunctionVariant) Name() string                     { return s.name }
func (s *AggregateFunctionVariant) Description() string              { return s.description }
func (s *AggregateFunctionVariant) Args() FuncParameterList          { return s.impl.Args }
func (s *AggregateFunctionVariant) Options() map[string]Option       { return s.impl.Options }
func (s *AggregateFunctionVariant) Variadic() *VariadicBehavior      { return s.impl.Variadic }
func (s *AggregateFunctionVariant) Deterministic() bool              { return s.impl.Deterministic }
func (s *AggregateFunctionVariant) SessionDependent() bool           { return s.impl.SessionDependent }
func (s *AggregateFunctionVariant) Nullability() NullabilityHandling { return s.impl.Nullability }
func (s *AggregateFunctionVariant) URI() string                      { return s.uri }
func (s *AggregateFunctionVariant) ResolveType(argumentTypes []types.Type) (types.Type, error) {
	return EvaluateTypeExpression(s.impl.Nullability, s.impl.Return.ValueType, s.impl.Args, s.impl.Variadic, argumentTypes)
}
func (s *AggregateFunctionVariant) CompoundName() string {
	return s.name + ":" + s.impl.signatureKey()
}
func (s *AggregateFunctionVariant) ID() ID {
	return ID{URI: s.uri, Name: s.CompoundName()}
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
	uri         string
	impl        WindowFunctionImpl
}

func NewWindowFuncVariant(id ID) *WindowFunctionVariant {
	simpleName, args := parseFuncName(id.Name)
	return &WindowFunctionVariant{
		name: simpleName,
		uri:  id.URI,
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
		uri:  id.URI,
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
func (s *WindowFunctionVariant) URI() string                      { return s.uri }
func (s *WindowFunctionVariant) ResolveType(argumentTypes []types.Type) (types.Type, error) {
	return EvaluateTypeExpression(s.impl.Nullability, s.impl.Return.ValueType, s.impl.Args, s.impl.Variadic, argumentTypes)
}
func (s *WindowFunctionVariant) CompoundName() string {
	return s.name + ":" + s.impl.signatureKey()
}
func (s *WindowFunctionVariant) ID() ID {
	return ID{URI: s.uri, Name: s.CompoundName()}
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
