package parser

import (
	"fmt"
	"strconv"
	"strings"

	substraitgo "github.com/substrait-io/substrait-go/v3"
	"github.com/substrait-io/substrait-go/v3/expr"
	"github.com/substrait-io/substrait-go/v3/extensions"
	"github.com/substrait-io/substrait-go/v3/functions"
	"github.com/substrait-io/substrait-go/v3/types"
)

type TestFuncType string

const (
	ScalarFuncType    TestFuncType = "scalar"
	AggregateFuncType TestFuncType = "aggregate"
	WindowFuncType    TestFuncType = "window"
)

type CaseLiteral struct {
	Type           types.Type
	ValueText      string
	Value          expr.Literal
	SubstraitError *SubstraitError
}

type TestFileHeader struct {
	Version     string
	FuncType    TestFuncType
	IncludedURI string
}

type TestCase struct {
	BaseURI       string
	GroupDesc     string
	FuncName      string
	Args          []*CaseLiteral
	AggregateArgs []*AggregateArgument
	Result        *CaseLiteral
	Options       FuncOptions
	Columns       [][]expr.Literal
	TableName     string
	ColumnTypes   []types.Type
	FuncType      TestFuncType
}

func (tc *TestCase) GetFunctionOptions() []*types.FunctionOption {
	if len(tc.Options) == 0 {
		return nil
	}
	funcOptions := make([]*types.FunctionOption, 0)
	for key, value := range tc.Options {
		funcOptions = append(funcOptions, &types.FunctionOption{
			Name:       key,
			Preference: []string{value},
		})
	}
	return funcOptions
}

func (tc *TestCase) getScalarFuncArgTypes() []types.Type {
	argTypes := make([]types.Type, len(tc.Args))
	for i, arg := range tc.Args {
		argTypes[i] = arg.Type
	}
	return argTypes
}

func (tc *TestCase) getAggregateFuncArgTypes() []types.Type {
	argTypes := make([]types.Type, len(tc.AggregateArgs))
	for i, arg := range tc.AggregateArgs {
		if arg.IsScalar {
			argTypes[i] = arg.Argument.Type
			continue
		}
		argTypes[i] = arg.ColumnType
	}
	return argTypes
}

func (tc *TestCase) getAggregateFuncTableSchema() []types.Type {
	schemaTypes := make([]types.Type, len(tc.AggregateArgs))
	for i, arg := range tc.AggregateArgs {
		if !arg.IsScalar {
			schemaTypes[i] = arg.ColumnType
		}
	}
	return schemaTypes
}

func (tc *TestCase) GetArgTypes() []types.Type {
	switch tc.FuncType {
	case ScalarFuncType:
		return tc.getScalarFuncArgTypes()
	case AggregateFuncType:
		return tc.getAggregateFuncArgTypes()
	default:
		panic(fmt.Sprintf("unsupported function type: %s", tc.FuncType))
	}
}

func (tc *TestCase) scalarSignatureKey() string {
	var b strings.Builder
	for i, a := range tc.Args {
		if i != 0 {
			b.WriteByte('_')
		}
		b.WriteString(a.Type.ShortString())
	}
	return b.String()
}

func (tc *TestCase) aggregateSignatureKey() string {
	var b strings.Builder
	for i, a := range tc.AggregateArgs {
		if i != 0 {
			b.WriteByte('_')
		}
		b.WriteString(a.GetType().ShortString())
	}
	return b.String()
}

func (tc *TestCase) signatureKey() string {
	switch tc.FuncType {
	case ScalarFuncType:
		return tc.scalarSignatureKey()
	case AggregateFuncType:
		return tc.aggregateSignatureKey()
	default:
		panic(fmt.Sprintf("unsupported function type: %s", tc.FuncType))
	}
}

func (tc *TestCase) CompoundFunctionName() string {
	return tc.FuncName + ":" + tc.signatureKey()
}

func (tc *TestCase) ID() extensions.ID {
	baseURI := tc.BaseURI
	if strings.HasPrefix(baseURI, "/") {
		baseURI = "https://github.com/substrait-io/substrait/blob/main" + tc.BaseURI
	}
	return extensions.ID{
		URI:  baseURI,
		Name: tc.CompoundFunctionName(),
	}
}

func (tc *TestCase) GetScalarFunctionInvocation(reg *expr.ExtensionRegistry, funcRegistry functions.FunctionRegistry) (*expr.ScalarFunction, error) {
	if tc.FuncType != ScalarFuncType {
		return nil, fmt.Errorf("not a scalar function testcase")
	}
	id := tc.ID()
	args := make([]types.FuncArg, len(tc.Args))
	for i, arg := range tc.Args {
		args[i] = arg.Value
	}

	invocation, err := expr.NewScalarFunc(*reg, id, tc.GetFunctionOptions(), args...)
	if err == nil {
		return invocation, nil
	}

	// exact match not found, try to find a function that matches with function parameter type "any"
	funcVariants := funcRegistry.GetScalarFunctions(tc.FuncName, len(args))
	for _, function := range funcVariants {
		isMatch, err1 := function.Match(tc.GetArgTypes())
		if err1 == nil && isMatch && function.ID().URI == id.URI {
			return expr.NewScalarFunc(*reg, function.ID(), tc.GetFunctionOptions(), args...)
		}
	}
	return nil, fmt.Errorf("%w: no matching function found  or %s", substraitgo.ErrNotFound, id)
}

func (tc *TestCase) GetAggregateFunctionInvocation(reg *expr.ExtensionRegistry, funcRegistry functions.FunctionRegistry) (*expr.AggregateFunction, error) {
	if tc.FuncType != AggregateFuncType {
		return nil, fmt.Errorf("not an aggregate function testcase")
	}
	id := tc.ID()
	args := make([]types.FuncArg, len(tc.AggregateArgs))
	baseSchema := types.NewRecordTypeFromTypes(tc.getAggregateFuncTableSchema())
	for i, arg := range tc.AggregateArgs {
		if arg.IsScalar {
			args[i] = arg.Argument.Value
			continue
		}

		fieldRef, err := expr.NewFieldRef(expr.RootReference, expr.NewStructFieldRef(arg.ColumnIndex), baseSchema)
		if err != nil {
			return nil, err
		}
		args[i] = fieldRef
	}

	invocation, err := expr.NewAggregateFunc(*reg, id, tc.GetFunctionOptions(),
		types.AggInvocationAll, types.AggPhaseInitialToResult, nil, args...)
	if err == nil {
		return invocation, nil
	}

	funcVariants := funcRegistry.GetAggregateFunctions(tc.FuncName, len(args))
	for _, function := range funcVariants {
		isMatch, err := function.Match(tc.GetArgTypes())
		if err == nil && isMatch && function.ID().URI == id.URI {
			return expr.NewAggregateFunc(*reg, function.ID(), tc.GetFunctionOptions(),
				types.AggInvocationAll, types.AggPhaseInitialToResult, nil, args...)
		}
	}
	return nil, fmt.Errorf("%w: no matching function found  or %s", substraitgo.ErrNotFound, id)
}

type TestGroup struct {
	Description string
	TestCases   []*TestCase
}

type TestFile struct {
	Header    *TestFileHeader
	TestCases []*TestCase
}

type FuncOptions map[string]string

type AggregateArgument struct {
	Argument    *CaseLiteral // This is used to store either a ScalarArgument or a ColumnArgument as List in the Value
	TableName   string
	ColumnName  string
	ColumnType  types.Type
	ColumnIndex int32
	IsScalar    bool
}

func (a *AggregateArgument) GetType() types.Type {
	if a.IsScalar {
		return a.Argument.Type
	}
	return a.ColumnType
}

func newAggregateArgument(tableName string, columnName string, columnType types.Type) (*AggregateArgument, error) {
	index, err := strconv.ParseInt(columnName[3:], 10, 64)
	if err != nil {
		return nil, err
	}
	if index < 0 {
		return nil, fmt.Errorf("column index must be greater than or equal to 0")
	}
	return &AggregateArgument{
		TableName:   tableName,
		ColumnName:  columnName,
		ColumnType:  columnType,
		ColumnIndex: int32(index),
	}, nil
}

type CompactAggregateFuncCall struct {
	FuncName      string
	Rows          [][]expr.Literal
	AggregateArgs []*AggregateArgument
}
