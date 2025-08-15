package parser

import (
	"fmt"
	"strconv"
	"strings"

	substraitgo "github.com/substrait-io/substrait-go/v6"
	"github.com/substrait-io/substrait-go/v6/expr"
	"github.com/substrait-io/substrait-go/v6/extensions"
	"github.com/substrait-io/substrait-go/v6/functions"
	"github.com/substrait-io/substrait-go/v6/types"
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

func (c *CaseLiteral) String() string {
	if c.SubstraitError != nil {
		return c.SubstraitError.String()
	}
	if c.Value == nil {
		return "NULL"
	}
	return literalToString(c.Value) + "::" + c.Type.String()
}

func literalToString(literal expr.Literal) string {
	if literal == nil {
		panic("literal is nil")
	}

	switch lit := literal.(type) {
	case *expr.NullLiteral:
		return literal.ValueString()
	case types.IsoValuePrinter:
		switch literal.GetType().(type) {
		// for these types enclose in single quotes
		case *types.IntervalYearType, *types.IntervalDayType,
			*types.PrecisionTimeType, *types.PrecisionTimestampType,
			*types.PrecisionTimestampTzType, *types.TimestampType,
			*types.TimeType, *types.TimestampTzType:
			return fmt.Sprintf("'%s'", lit.IsoValueString())
		}
	}
	switch literal.GetType().(type) {
	// for these types enclose in single quotes
	case *types.StringType, *types.FixedCharType, *types.VarCharType,
		*types.FixedBinaryType, *types.BinaryType, *types.DateType:
		return fmt.Sprintf("'%s'", literal.ValueString())
	default:
		return literal.ValueString()
	}
}

func (c *CaseLiteral) AsAggregateArgumentString() string {
	if c.SubstraitError != nil {
		return c.SubstraitError.String()
	}
	if list, ok := c.Value.(*expr.ListLiteral); ok {
		var elements []string
		for _, element := range list.Value {
			elements = append(elements, literalToString(element))
		}
		return "(" + strings.Join(elements, ", ") + ")::" + c.Type.String()
	}
	return c.Value.ValueString() + "::" + c.Type.String()
}

// updateLiteralType updates the type of the literal CaseLiteral.Value to use the CaseLiteral.Type
// Parser creates a literal with a type using existing util functions.
// For ParameterizedTypes utils functions use minimum required values for the parameters.
// This function changes the type to use requested type, so that the function invocation object is created correctly.
func (c *CaseLiteral) updateLiteralType() error {
	if len(c.Type.GetParameters()) == 0 {
		return nil
	}
	switch proLit := c.Value.(type) {
	case *expr.NullLiteral:
		return nil
	case expr.WithTypeLiteral:
		lit, err := proLit.WithType(c.Type)
		if err != nil {
			return err
		}
		c.Value = lit
		return nil
	}
	return fmt.Errorf("literal type %T is not handled to update the type", c.Value)
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

func (tc *TestCase) String() string {
	switch tc.FuncType {
	case ScalarFuncType:
		return tc.getScalarTestString()
	case AggregateFuncType:
		return tc.getAggregateTestString()
	default:
		panic(fmt.Sprintf("unsupported function type: %s", tc.FuncType))
	}
}

func (tc *TestCase) getScalarTestString() string {
	var b strings.Builder
	b.WriteString(tc.FuncName)
	b.WriteByte('(')
	for i, arg := range tc.Args {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(arg.String())
	}
	b.WriteByte(')')
	b.WriteString(tc.getOptionString())
	b.WriteString(" = ")
	b.WriteString(tc.Result.String())
	return b.String()
}

func (tc *TestCase) getAggregateTestString() string {
	var b strings.Builder
	if tc.needCompactAggregateFuncCall() {
		b.WriteString(tc.getAggregateFuncTableString())
		b.WriteByte(' ')
	}

	b.WriteString(tc.FuncName)
	b.WriteByte('(')
	for i, arg := range tc.AggregateArgs {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(arg.String())
	}
	b.WriteByte(')')
	b.WriteString(tc.getOptionString())
	b.WriteString(" = ")
	b.WriteString(tc.Result.String())
	return b.String()
}

func (tc *TestCase) needCompactAggregateFuncCall() bool {
	if tc.FuncType == ScalarFuncType {
		return false
	}
	if len(tc.AggregateArgs) == 0 {
		return true
	}
	for _, arg := range tc.AggregateArgs {
		if arg.IsScalar || arg.ColumnName != "" {
			return true
		}
	}
	// common case of single column aggregate function
	return false
}

func (tc *TestCase) getAggregateFuncTableString() string {
	var b strings.Builder
	if len(tc.Columns) == 0 {
		return ""
	}
	b.WriteByte('(')
	numRows := len(tc.Columns[0])
	for i := 0; i < numRows; i++ {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteByte('(')
		for j, column := range tc.Columns {
			if j != 0 {
				b.WriteString(", ")
			}
			b.WriteString(literalToString(column[i]))
		}
		b.WriteByte(')')
	}
	b.WriteByte(')')
	return b.String()
}

func (tc *TestCase) getOptionString() string {
	if len(tc.Options) == 0 {
		return ""
	}
	var b strings.Builder
	b.WriteString(" [")
	var options []string
	for k, v := range tc.Options {
		options = append(options, fmt.Sprintf("%s:%s", k, v))
	}
	b.WriteString(strings.Join(options, ","))
	b.WriteByte(']')
	return b.String()
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

func (tc *TestCase) GetAggregateColumnsData() ([][]expr.Literal, error) {
	if tc.FuncType != AggregateFuncType {
		return nil, fmt.Errorf("expected function type %v, but got %v", AggregateFuncType, tc.FuncType)
	}

	if len(tc.Columns) > 0 {
		return tc.Columns, nil
	}

	columns := make([][]expr.Literal, len(tc.AggregateArgs))

	for colIdx, arg := range tc.AggregateArgs {
		values, ok := arg.Argument.Value.(*expr.NestedLiteral[expr.ListLiteralValue])
		if !ok {
			return nil, fmt.Errorf("column %d: expected NestedLiteral[ListLiteralValue], but got %T", colIdx, arg.Argument.Value)
		}

		columns[colIdx] = make([]expr.Literal, len(values.Value))
		copy(columns[colIdx], values.Value)
	}

	return columns, nil
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

func (a *AggregateArgument) String() string {
	if a.IsScalar {
		return a.Argument.String()
	}
	if a.ColumnName == "" {
		return a.Argument.AsAggregateArgumentString()
	}
	return a.ColumnName + "::" + a.ColumnType.String()
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

type SubstraitError struct {
	Error string
}

func (e SubstraitError) String() string {
	return "<!" + e.Error + ">"
}
