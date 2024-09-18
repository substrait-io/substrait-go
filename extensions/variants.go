// SPDX-License-Identifier: Apache-2.0

package extensions

import (
	"fmt"
	"strings"

	substraitgo "github.com/substrait-io/substrait-go"
	"github.com/substrait-io/substrait-go/types"
	"github.com/substrait-io/substrait-go/types/integer_parameters"
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

func validateType(arg Argument, actual types.Type, idx int, nullHandling NullabilityHandling) (bool, error) {
	allNonNull := true
	switch p := arg.(type) {
	case EnumArg:
		if actual != nil {
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
			if t, ok := p.Value.Expr.(*parser.Type); ok {
				if isNullable != t.Optional() {
					return allNonNull, fmt.Errorf("%w: discrete nullability did not match for arg #%d",
						substraitgo.ErrInvalidType, idx)
				}
			} else {
				return allNonNull, substraitgo.ErrNotImplemented
			}
		}
	case TypeArg:
		return allNonNull, substraitgo.ErrNotImplemented
	}

	return allNonNull, nil
}

func EvaluateTypeExpression(nullHandling NullabilityHandling, expr parser.TypeExpression, paramTypeList ArgumentList, variadic *VariadicBehavior, actualTypes []types.Type) (types.Type, error) {
	if len(paramTypeList) != len(actualTypes) {
		if variadic == nil {
			return nil, fmt.Errorf("%w: mismatch in number of arguments provided. got %d, expected %d",
				substraitgo.ErrInvalidExpr, len(actualTypes), len(paramTypeList))
		}

		if !variadic.IsValidArgumentCount(len(actualTypes) - len(paramTypeList) - 1) {
			return nil, fmt.Errorf("%w: mismatch in number of arguments provided, invalid number of variadic params. got %d total",
				substraitgo.ErrInvalidExpr, len(actualTypes))
		}
	}

	allNonNull := true
	for i, p := range paramTypeList {
		nonNull, err := validateType(p, actualTypes[i], i, nullHandling)
		if err != nil {
			return nil, err
		}
		allNonNull = allNonNull && nonNull
	}

	// validate varidic argument consistency
	if variadic != nil && len(actualTypes) > len(paramTypeList) && variadic.ParameterConsistency == ConsistentParams {
		nparams := len(paramTypeList)
		lastParam := paramTypeList[nparams-1]
		for i, actual := range actualTypes[nparams:] {
			nonNull, err := validateType(lastParam, actual, nparams+i, nullHandling)
			if err != nil {
				return nil, err
			}
			allNonNull = allNonNull && nonNull
		}
	}

	var outType types.Type
	if t, ok := expr.Expr.(*parser.Type); ok {
		var err error
		outType, err = t.RetType()
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

func matchArguments(nullability NullabilityHandling, paramTypeList ArgumentList, variadicBehavior *VariadicBehavior, actualTypes []types.Type) (bool, error) {
	if variadicBehavior == nil && len(actualTypes) != len(paramTypeList) {
		return false, nil
	} else if variadicBehavior != nil && !validateVariadicBehaviorForMatch(variadicBehavior, actualTypes) {
		return false, nil
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
	return true, nil
}

func matchArgumentAt(actualType types.Type, argPos int, nullability NullabilityHandling, paramTypeList ArgumentList, variadicBehavior *VariadicBehavior) (bool, error) {
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
	} else if variadicBehavior != nil && !variadicBehavior.IsValidArgumentCount(argPos+1) {
		// this argument position can't be more than the max allowed by the variadic behavior
		return false, nil
	}

	if HasSyncParams(funcDefArgList) {
		return false, fmt.Errorf("%w: function has sync params", substraitgo.ErrNotImplemented)
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
		firstVariadicArgIdx := variadicBehavior.Min - 1
		for i := firstVariadicArgIdx; i < len(actualTypes)-1; i++ {
			if !actualTypes[i].Equals(actualTypes[i+1]) {
				return false
			}
		}
	}
	return true
}

func getFuncDefFromArgList(paramTypeList ArgumentList) ([]types.FuncDefArgType, error) {
	var out []types.FuncDefArgType
	for argPos, param := range paramTypeList {
		switch paramType := param.(type) {
		case ValueArg:
			funcDefArgType, err := paramType.Value.Expr.(*parser.Type).ArgType()
			if err != nil {
				return nil, err
			}
			out = append(out, funcDefArgType)
		case EnumArg:
			return nil, fmt.Errorf("%w: invalid argument at position %d for match operation", substraitgo.ErrInvalidType, argPos)
		case TypeArg:
			return nil, fmt.Errorf("%w: invalid argument at position %d for match operation", substraitgo.ErrInvalidType, argPos)
		default:
			return nil, fmt.Errorf("%w: invalid argument at position %d for match operation", substraitgo.ErrInvalidType, argPos)
		}
	}
	return out, nil
}

func parseFuncName(compoundName string) (name string, args ArgumentList) {
	name, argsStr, _ := strings.Cut(compoundName, ":")
	if len(argsStr) == 0 {
		return name, nil
	}
	splitArgs := strings.Split(argsStr, "_")
	for _, argStr := range splitArgs {
		parsed, err := defParser.ParseString(argStr)
		if err != nil {
			panic(err)
		}
		exp := ValueArg{Name: name, Value: parsed}
		args = append(args, exp)
	}

	return name, args
}

func minArgumentCount(paramTypeList ArgumentList, variadicBehavior *VariadicBehavior) int {
	if variadicBehavior == nil {
		return len(paramTypeList)
	}
	return variadicBehavior.Min
}

func maxArgumentCount(paramTypeList ArgumentList, variadicBehavior *VariadicBehavior) int {
	if variadicBehavior == nil {
		return len(paramTypeList)
	}
	return variadicBehavior.Max
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
func (s *ScalarFunctionVariant) Args() ArgumentList               { return s.impl.Args }
func (s *ScalarFunctionVariant) Options() map[string]Option       { return s.impl.Options }
func (s *ScalarFunctionVariant) Variadic() *VariadicBehavior      { return s.impl.Variadic }
func (s *ScalarFunctionVariant) Deterministic() bool              { return s.impl.Deterministic }
func (s *ScalarFunctionVariant) SessionDependent() bool           { return s.impl.SessionDependent }
func (s *ScalarFunctionVariant) Nullability() NullabilityHandling { return s.impl.Nullability }
func (s *ScalarFunctionVariant) URI() string                      { return s.uri }
func (s *ScalarFunctionVariant) ResolveType(argumentTypes []types.Type) (types.Type, error) {
	return EvaluateTypeExpression(s.impl.Nullability, s.impl.Return, s.impl.Args, s.impl.Variadic, argumentTypes)
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
func (s *AggregateFunctionVariant) Args() ArgumentList               { return s.impl.Args }
func (s *AggregateFunctionVariant) Options() map[string]Option       { return s.impl.Options }
func (s *AggregateFunctionVariant) Variadic() *VariadicBehavior      { return s.impl.Variadic }
func (s *AggregateFunctionVariant) Deterministic() bool              { return s.impl.Deterministic }
func (s *AggregateFunctionVariant) SessionDependent() bool           { return s.impl.SessionDependent }
func (s *AggregateFunctionVariant) Nullability() NullabilityHandling { return s.impl.Nullability }
func (s *AggregateFunctionVariant) URI() string                      { return s.uri }
func (s *AggregateFunctionVariant) ResolveType(argumentTypes []types.Type) (types.Type, error) {
	return EvaluateTypeExpression(s.impl.Nullability, s.impl.Return, s.impl.Args, s.impl.Variadic, argumentTypes)
}
func (s *AggregateFunctionVariant) CompoundName() string {
	return s.name + ":" + s.impl.signatureKey()
}
func (s *AggregateFunctionVariant) ID() ID {
	return ID{URI: s.uri, Name: s.CompoundName()}
}
func (s *AggregateFunctionVariant) Decomposability() DecomposeType { return s.impl.Decomposable }
func (s *AggregateFunctionVariant) Intermediate() (types.FuncDefArgType, error) {
	if t, ok := s.impl.Intermediate.Expr.(*parser.Type); ok {
		return t.ArgType()
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

		intermediate, err := defParser.ParseString(opts.IntermediateOutputType)
		if err != nil {
			panic(err)
		}
		aggIntermediate = *intermediate
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
func (s *WindowFunctionVariant) Args() ArgumentList               { return s.impl.Args }
func (s *WindowFunctionVariant) Options() map[string]Option       { return s.impl.Options }
func (s *WindowFunctionVariant) Variadic() *VariadicBehavior      { return s.impl.Variadic }
func (s *WindowFunctionVariant) Deterministic() bool              { return s.impl.Deterministic }
func (s *WindowFunctionVariant) SessionDependent() bool           { return s.impl.SessionDependent }
func (s *WindowFunctionVariant) Nullability() NullabilityHandling { return s.impl.Nullability }
func (s *WindowFunctionVariant) URI() string                      { return s.uri }
func (s *WindowFunctionVariant) ResolveType(argumentTypes []types.Type) (types.Type, error) {
	return EvaluateTypeExpression(s.impl.Nullability, s.impl.Return, s.impl.Args, s.impl.Variadic, argumentTypes)
}
func (s *WindowFunctionVariant) CompoundName() string {
	return s.name + ":" + s.impl.signatureKey()
}
func (s *WindowFunctionVariant) ID() ID {
	return ID{URI: s.uri, Name: s.CompoundName()}
}
func (s *WindowFunctionVariant) Decomposability() DecomposeType { return s.impl.Decomposable }
func (s *WindowFunctionVariant) Intermediate() (types.FuncDefArgType, error) {
	if t, ok := s.impl.Intermediate.Expr.(*parser.Type); ok {
		return t.ArgType()
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
