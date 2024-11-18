package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/substrait-io/substrait-go/expr"
	"github.com/substrait-io/substrait-go/literal"
	"github.com/substrait-io/substrait-go/testcases/parser/baseparser"
	"github.com/substrait-io/substrait-go/types"
	"github.com/substrait-io/substrait-go/types/parser/util"
)

type TestCaseVisitor struct {
	baseparser.FuncTestCaseParserVisitor
	ErrorListener util.VisitErrorListener
}

var _ baseparser.FuncTestCaseParserVisitor = &TestCaseVisitor{}

func (v *TestCaseVisitor) Visit(tree antlr.ParseTree) interface{} {
	if tree == nil {
		return nil
	}
	return tree.Accept(v)
}

func (v *TestCaseVisitor) VisitDoc(ctx *baseparser.DocContext) interface{} {
	header := v.Visit(ctx.Header()).(TestFileHeader)
	testcases := make([]*TestCase, 0, len(ctx.AllTestGroup()))
	for _, testGroup := range ctx.AllTestGroup() {
		groupTestCases := v.Visit(testGroup).([]*TestCase)
		for _, testcase := range groupTestCases {
			testcase.BaseURI = header.IncludedURI
		}
		testcases = append(testcases, groupTestCases...)
	}
	return &TestFile{
		Header:    header,
		TestCases: testcases,
	}
}

func (v *TestCaseVisitor) VisitHeader(ctx *baseparser.HeaderContext) interface{} {
	return TestFileHeader{
		Version:     ctx.Version().GetText(),
		IncludedURI: ctx.Include().GetText(),
	}
}

type TestGroup struct {
	Description string
	TestCases   []*TestCase
}

func (v *TestCaseVisitor) VisitScalarFuncTestGroup(ctx *baseparser.ScalarFuncTestGroupContext) interface{} {
	groupDesc := v.Visit(ctx.TestGroupDescription()).(string)
	groupTestCases := make([]*TestCase, 0, len(ctx.AllTestCase()))
	for _, tc := range ctx.AllTestCase() {
		testcase := v.Visit(tc).(*TestCase)
		testcase.GroupDesc = groupDesc
		groupTestCases = append(groupTestCases, testcase)
	}
	return groupTestCases
}

func (v *TestCaseVisitor) VisitAggregateFuncTestGroup(ctx *baseparser.AggregateFuncTestGroupContext) interface{} {
	groupDesc := v.Visit(ctx.TestGroupDescription()).(string)
	groupTestCases := make([]*TestCase, 0, len(ctx.AllAggFuncTestCase()))
	for _, tc := range ctx.AllAggFuncTestCase() {
		testcase := v.Visit(tc).(*TestCase)
		testcase.GroupDesc = groupDesc
		groupTestCases = append(groupTestCases, testcase)
	}
	return groupTestCases
}

func (v *TestCaseVisitor) VisitAggFuncTestCase(ctx *baseparser.AggFuncTestCaseContext) interface{} {
	testcase := v.Visit(ctx.AggFuncCall()).(*TestCase)
	testcase.Result = v.Visit(ctx.Result()).(*CaseLiteral)
	if ctx.Func_options() != nil {
		testcase.Options = v.Visit(ctx.Func_options()).(FuncOptions)
	}
	return testcase
}

func (v *TestCaseVisitor) VisitSingleArgAggregateFuncCall(ctx *baseparser.SingleArgAggregateFuncCallContext) interface{} {
	arg := v.Visit(ctx.DataColumn()).(*CaseLiteral)
	return &TestCase{
		FuncName:      ctx.Identifier().GetText(),
		AggregateArgs: []*AggregateArgument{{Argument: arg, ColumnType: arg.Type}},
		Result:        &CaseLiteral{SubstraitError: &SubstraitError{Error: "uninitialized"}},
	}
}

func (v *TestCaseVisitor) VisitCompactAggregateFuncCall(ctx *baseparser.CompactAggregateFuncCallContext) interface{} {
	rows := v.Visit(ctx.TableRows()).([][]expr.Literal)
	args := v.Visit(ctx.AggregateFuncArgs()).([]*AggregateArgument)

	numberOfColumns := len(rows[0])
	columnTypes := make([]types.Type, numberOfColumns)
	for _, arg := range args {
		if arg.ColumnIndex >= numberOfColumns {
			v.ErrorListener.ReportVisitError(fmt.Errorf("invalid column index %d, expected less than %d", arg.ColumnIndex, len(columnTypes)))
			continue
		}
		columnTypes[arg.ColumnIndex] = arg.ColumnType
	}
	columns := v.getColumnsFromRows(rows, columnTypes)
	return &TestCase{
		FuncName:      ctx.Identifier().GetText(),
		Columns:       columns,
		AggregateArgs: args,
	}
}

func (v *TestCaseVisitor) VisitMultiArgAggregateFuncCall(ctx *baseparser.MultiArgAggregateFuncCallContext) interface{} {
	testcase := v.Visit(ctx.TableData()).(*TestCase)
	args := v.Visit(ctx.QualifiedAggregateFuncArgs()).([]*AggregateArgument)
	testcase.FuncName = ctx.Identifier().GetText()
	testcase.AggregateArgs = args
	for _, arg := range args {
		if arg.TableName != "" {
			if testcase.TableName != arg.TableName {
				err := fmt.Errorf("table name in argument %s, does not match the table name in the function call %s", arg.TableName, testcase.TableName)
				v.ErrorListener.ReportVisitError(err)
			}
		}
	}
	return testcase
}

func (v *TestCaseVisitor) VisitQualifiedAggregateFuncArgs(ctx *baseparser.QualifiedAggregateFuncArgsContext) interface{} {
	args := make([]*AggregateArgument, 0, len(ctx.AllQualifiedAggregateFuncArg()))
	for _, arg := range ctx.AllQualifiedAggregateFuncArg() {
		args = append(args, v.Visit(arg).(*AggregateArgument))
	}
	return args
}

func (v *TestCaseVisitor) VisitQualifiedAggregateFuncArg(ctx *baseparser.QualifiedAggregateFuncArgContext) interface{} {
	if ctx.Argument() != nil {
		return &AggregateArgument{
			Argument: v.Visit(ctx.Argument()).(*CaseLiteral),
		}
	}
	arg, err := newAggregateArgument(ctx.Identifier().GetText(), ctx.ColumnName().GetText(), nil)
	if err != nil {
		v.ErrorListener.ReportVisitError(fmt.Errorf("invalid aggregate func arg %v", err))
	}
	return arg
}

func (v *TestCaseVisitor) VisitTableData(ctx *baseparser.TableDataContext) interface{} {
	columnTypes := make([]types.Type, 0, len(ctx.AllDataType()))
	rows := v.Visit(ctx.TableRows()).([][]expr.Literal)
	for _, dataType := range ctx.AllDataType() {
		columnType := v.Visit(dataType).(types.Type)
		columnTypes = append(columnTypes, columnType)
	}
	columns := v.getColumnsFromRows(rows, columnTypes)
	return &TestCase{
		Columns:     columns,
		TableName:   ctx.GetTableName().GetText(),
		ColumnTypes: columnTypes,
	}
}

func (v *TestCaseVisitor) getColumnsFromRows(rows [][]expr.Literal, columnTypes []types.Type) [][]expr.Literal {
	columns := make([][]expr.Literal, 0, len(columnTypes))
	for i, columnType := range columnTypes {
		column := make([]expr.Literal, 0, len(rows))
		for _, row := range rows {
			column = append(column, v.getLiteral(row[i], columnType))
		}
		columns = append(columns, column)
	}
	return columns
}

func (v *TestCaseVisitor) VisitDataColumn(ctx *baseparser.DataColumnContext) interface{} {
	columnType := v.Visit(ctx.DataType()).(types.Type)
	columnValues := v.Visit(ctx.ColumnValues()).([]expr.Literal)
	var err error
	var column expr.Literal
	if len(columnValues) == 0 {
		column = expr.NewEmptyListLiteral(columnType, false)
	} else {
		column, err = v.getListLiteral(columnValues, columnType)
		if err != nil {
			v.ErrorListener.ReportVisitError(fmt.Errorf("invalid column values %v", err))
		}
	}
	return &CaseLiteral{
		Type:  columnType,
		Value: column,
	}
}

func (v *TestCaseVisitor) VisitTableRows(ctx *baseparser.TableRowsContext) interface{} {
	rows := make([][]expr.Literal, 0, len(ctx.AllColumnValues()))
	for _, row := range ctx.AllColumnValues() {
		rows = append(rows, v.Visit(row).([]expr.Literal))
	}
	return rows
}

func (v *TestCaseVisitor) VisitColumnValues(ctx *baseparser.ColumnValuesContext) interface{} {
	values := make([]expr.Literal, 0, len(ctx.AllLiteral()))
	for _, literalValue := range ctx.AllLiteral() {
		values = append(values, v.Visit(literalValue).(expr.Literal))
	}
	return values
}

func (v *TestCaseVisitor) VisitAggregateFuncArgs(ctx *baseparser.AggregateFuncArgsContext) interface{} {
	args := make([]*AggregateArgument, 0, len(ctx.AllAggregateFuncArg()))
	for _, arg := range ctx.AllAggregateFuncArg() {
		args = append(args, v.Visit(arg).(*AggregateArgument))
	}
	return args
}

func (v *TestCaseVisitor) VisitAggregateFuncArg(ctx *baseparser.AggregateFuncArgContext) interface{} {
	if ctx.Argument() != nil {
		return &AggregateArgument{
			Argument: v.Visit(ctx.Argument()).(*CaseLiteral),
		}
	}
	argType := v.Visit(ctx.DataType()).(types.Type)
	arg, err := newAggregateArgument("", ctx.ColumnName().GetText(), argType)
	if err != nil {
		v.ErrorListener.ReportVisitError(fmt.Errorf("invalid aggregate func arg %v", err))
	}
	return arg
}

func (v *TestCaseVisitor) VisitFunc_options(ctx *baseparser.Func_optionsContext) interface{} {
	options := make(FuncOptions)
	for _, option := range ctx.AllFunc_option() {
		optionPair := v.Visit(option).([]string)
		options[optionPair[0]] = optionPair[1]
	}
	return options
}

func (v *TestCaseVisitor) VisitFunc_option(ctx *baseparser.Func_optionContext) interface{} {
	return []string{ctx.Option_name().GetText(), ctx.Option_value().GetText()}
}

func (v *TestCaseVisitor) VisitTestGroupDescription(ctx *baseparser.TestGroupDescriptionContext) interface{} {
	return strings.TrimPrefix(ctx.GetText(), "#")
}

func (v *TestCaseVisitor) VisitTestCase(ctx *baseparser.TestCaseContext) interface{} {
	return &TestCase{
		FuncName: ctx.Identifier().GetText(),
		Args:     v.Visit(ctx.Arguments()).([]*CaseLiteral),
		Result:   v.Visit(ctx.Result()).(*CaseLiteral),
	}
}

func (v *TestCaseVisitor) VisitArguments(ctx *baseparser.ArgumentsContext) interface{} {
	args := make([]*CaseLiteral, 0, len(ctx.AllArgument()))
	for _, argument := range ctx.AllArgument() {
		args = append(args, v.Visit(argument).(*CaseLiteral))
	}
	return args
}

func (v *TestCaseVisitor) VisitArgument(ctx *baseparser.ArgumentContext) interface{} {
	if ctx.IntArg() != nil {
		return v.Visit(ctx.IntArg())
	}
	if ctx.FloatArg() != nil {
		return v.Visit(ctx.FloatArg())
	}
	if ctx.StringArg() != nil {
		return v.Visit(ctx.StringArg())
	}
	if ctx.BooleanArg() != nil {
		return v.Visit(ctx.BooleanArg())
	}
	if ctx.TimestampArg() != nil {
		return v.Visit(ctx.TimestampArg())
	}
	if ctx.TimestampTzArg() != nil {
		return v.Visit(ctx.TimestampTzArg())
	}
	if ctx.DateArg() != nil {
		return v.Visit(ctx.DateArg())
	}
	if ctx.TimeArg() != nil {
		return v.Visit(ctx.TimeArg())
	}
	if ctx.IntervalYearArg() != nil {
		return v.Visit(ctx.IntervalYearArg())
	}
	if ctx.IntervalDayArg() != nil {
		return v.Visit(ctx.IntervalDayArg())
	}
	if ctx.NullArg() != nil {
		return v.Visit(ctx.NullArg())
	}
	if ctx.DecimalArg() != nil {
		return v.Visit(ctx.DecimalArg())
	}
	if ctx.ListArg() != nil {
		return v.Visit(ctx.ListArg())
	}
	return &CaseLiteral{}
}

func (v *TestCaseVisitor) VisitNullArg(*baseparser.NullArgContext) interface{} {
	return &CaseLiteral{}
}

func (v *TestCaseVisitor) VisitBooleanArg(ctx *baseparser.BooleanArgContext) interface{} {
	value := strings.ToLower(ctx.BooleanLiteral().GetText()) == "true"
	boolLiteral := literal.NewBool(value)
	return &CaseLiteral{Value: boolLiteral, ValueText: ctx.BooleanLiteral().GetText(), Type: &types.BooleanType{}}
}

func (v *TestCaseVisitor) getListLiteral(values []expr.Literal, elementType types.Type) (expr.Literal, error) {
	var err error
	var elements []expr.Literal
	switch elementType := elementType.(type) {
	case *types.Int8Type, *types.Int16Type, *types.Int32Type, *types.Int64Type:
		elements, err = getIntLiterals(values, elementType)
		if err != nil {
			v.ErrorListener.ReportVisitError(fmt.Errorf("invalid list arg %v", err))
		}
	case *types.Float32Type, *types.Float64Type:
		elements, err = getFloatLiterals(values, elementType)
		if err != nil {
			v.ErrorListener.ReportVisitError(fmt.Errorf("invalid list arg %v", err))
		}
	case *types.DecimalType:
		elements, err = getDecimalLiterals(values)
		if err != nil {
			v.ErrorListener.ReportVisitError(fmt.Errorf("invalid list arg %v", err))
		}
	default:
		elements = values
	}
	value, err := literal.NewList(elements)
	return value, err
}

func getDecimalLiterals(values []expr.Literal) ([]expr.Literal, error) {
	var elements []expr.Literal
	for _, value := range values {
		decimal, err := literal.NewDecimalFromString(value.(*expr.PrimitiveLiteral[string]).Value)
		if err != nil {
			return nil, err
		}
		elements = append(elements, decimal)
	}
	return elements, nil
}

func getIntLiterals(strLiterals []expr.Literal, intType types.Type) ([]expr.Literal, error) {
	var elements []expr.Literal
	for _, strLiteral := range strLiterals {
		integerStr := strLiteral.(*expr.PrimitiveLiteral[string]).Value
		element, err := getIntLiteral(integerStr, intType)
		if err != nil {
			return nil, err
		}
		elements = append(elements, element)
	}
	return elements, nil
}

func getFloatLiterals(strLiterals []expr.Literal, floatType types.Type) ([]expr.Literal, error) {
	var elements []expr.Literal
	for _, strLiteral := range strLiterals {
		floatStr := strLiteral.(*expr.PrimitiveLiteral[string]).Value
		value, err2 := getFloatLiteral(floatStr, floatType)
		if err2 != nil {
			return elements, err2
		}
		elements = append(elements, value)
	}
	return elements, nil
}

func (v *TestCaseVisitor) getLiteral(value expr.Literal, elementType types.Type) expr.Literal {
	switch elementType := elementType.(type) {
	case *types.Int8Type, *types.Int16Type, *types.Int32Type, *types.Int64Type:
		intLiteral, err := getIntLiteral(value.(*expr.PrimitiveLiteral[string]).Value, elementType)
		if err != nil {
			v.ErrorListener.ReportVisitError(fmt.Errorf("invalid int arg %v", err))
		}
		return intLiteral
	case *types.Float32Type, *types.Float64Type:
		floatLiteral, err := getFloatLiteral(value.(*expr.PrimitiveLiteral[string]).Value, elementType)
		if err != nil {
			v.ErrorListener.ReportVisitError(fmt.Errorf("invalid float arg %v", err))
		}
		return floatLiteral
	case *types.DecimalType:
		decimal, err := literal.NewDecimalFromString(value.(*expr.PrimitiveLiteral[string]).Value)
		if err != nil {
			v.ErrorListener.ReportVisitError(fmt.Errorf("invalid decimal arg %v", err))
		}
		return decimal
	default:
		return value
	}
}

func getFloatLiteral(floatStr string, floatType types.Type) (expr.Literal, error) {
	value, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		return nil, err
	}
	switch floatType := floatType.(type) {
	case *types.Float32Type:
		return literal.NewFloat32(float32(value)), nil
	case *types.Float64Type:
		return literal.NewFloat64(value), nil
	default:
		return nil, fmt.Errorf("invalid float type %v", floatType)
	}
}

func getIntLiteral(integerStr string, intType types.Type) (expr.Literal, error) {
	value, err := strconv.ParseInt(integerStr, 10, 64)
	if err != nil {
		return nil, err
	}
	switch intType.(type) {
	case *types.Int8Type:
		return literal.NewInt8(int8(value)), nil
	case *types.Int16Type:
		return literal.NewInt16(int16(value)), nil
	case *types.Int32Type:
		return literal.NewInt32(int32(value)), nil
	case *types.Int64Type:
		return literal.NewInt64(value), nil
	default:
		return nil, fmt.Errorf("invalid int value %v type %v", value, intType)
	}
}

func (v *TestCaseVisitor) VisitIntArg(ctx *baseparser.IntArgContext) interface{} {
	var typ types.Type
	typ = &types.Int8Type{}
	if ctx.I16() != nil {
		typ = &types.Int16Type{}
	} else if ctx.I32() != nil {
		typ = &types.Int32Type{}
	} else if ctx.I64() != nil {
		typ = &types.Int64Type{}
	}
	intLiteral, err := getIntLiteral(ctx.IntegerLiteral().GetText(), typ)
	if err != nil {
		v.ErrorListener.ReportVisitError(fmt.Errorf("invalid int arg %v", err))
	}
	return &CaseLiteral{Value: intLiteral, ValueText: ctx.IntegerLiteral().GetText(), Type: typ}
}

func (v *TestCaseVisitor) VisitFloatArg(ctx *baseparser.FloatArgContext) interface{} {
	var floatLiteral expr.Literal
	var err error
	if ctx.FP32() != nil {
		floatLiteral, err = getFloatLiteral(ctx.NumericLiteral().GetText(), &types.Float32Type{})
	} else {
		floatLiteral, err = getFloatLiteral(ctx.NumericLiteral().GetText(), &types.Float64Type{})
	}
	if err != nil {
		v.ErrorListener.ReportVisitError(fmt.Errorf("invalid float arg %v", err))
	}
	return &CaseLiteral{Value: floatLiteral, ValueText: ctx.NumericLiteral().GetText(), Type: floatLiteral.GetType()}
}

func (v *TestCaseVisitor) VisitStringArg(ctx *baseparser.StringArgContext) interface{} {
	value := literal.NewString(getRawStringFromStringLiteral(ctx.StringLiteral().GetText()))
	return &CaseLiteral{Value: value, ValueText: ctx.StringLiteral().GetText(), Type: &types.StringType{}}
}

func (v *TestCaseVisitor) VisitTimestampArg(ctx *baseparser.TimestampArgContext) interface{} {
	timestampStr := getRawStringFromStringLiteral(ctx.TimestampLiteral().GetText())
	value, err := literal.NewTimestampFromString(timestampStr)
	if err != nil {
		v.ErrorListener.ReportVisitError(fmt.Errorf("invalid timestampTZ arg %v", err))
	}
	return &CaseLiteral{Value: value, ValueText: ctx.TimestampLiteral().GetText(), Type: &types.TimestampType{}}
}

func getRawStringFromStringLiteral(text string) string {
	return strings.Trim(text, "'")
}

func (v *TestCaseVisitor) VisitTimestampTzArg(ctx *baseparser.TimestampTzArgContext) interface{} {
	value, err := literal.NewTimestampTZFromString(getRawStringFromStringLiteral(ctx.TimestampTzLiteral().GetText()))
	if err != nil {
		v.ErrorListener.ReportVisitError(fmt.Errorf("invalid timestampTZ arg %v", err))
	}
	return &CaseLiteral{Value: value, ValueText: ctx.TimestampTzLiteral().GetText(), Type: &types.TimestampTzType{}}
}

func (v *TestCaseVisitor) VisitDateArg(ctx *baseparser.DateArgContext) interface{} {
	value, err := literal.NewDateFromString(getRawStringFromStringLiteral(ctx.DateLiteral().GetText()))
	if err != nil {
		v.ErrorListener.ReportVisitError(fmt.Errorf("invalid date arg %v", err))
	}
	return &CaseLiteral{Value: value, ValueText: ctx.DateLiteral().GetText(), Type: &types.DateType{}}
}

func (v *TestCaseVisitor) VisitTimeArg(ctx *baseparser.TimeArgContext) interface{} {
	value, err := literal.NewTimeFromString(getRawStringFromStringLiteral(ctx.TimeLiteral().GetText()))
	if err != nil {
		v.ErrorListener.ReportVisitError(fmt.Errorf("invalid time arg %v", err))
	}
	return &CaseLiteral{Value: value, ValueText: ctx.TimeLiteral().GetText(), Type: &types.TimeType{}}
}

func (v *TestCaseVisitor) VisitIntervalYearArg(ctx *baseparser.IntervalYearArgContext) interface{} {
	interval := getRawStringFromStringLiteral(ctx.IntervalYearLiteral().GetText())
	value, err := literal.NewIntervalYearsToMonthFromString(interval)
	if err != nil {
		v.ErrorListener.ReportVisitError(fmt.Errorf("invalid interval year arg %v", err))
	}
	return &CaseLiteral{Value: value, ValueText: ctx.IntervalYearLiteral().GetText(), Type: &types.IntervalYearType{}}
}

func (v *TestCaseVisitor) VisitIntervalDayArg(ctx *baseparser.IntervalDayArgContext) interface{} {
	interval := getRawStringFromStringLiteral(ctx.IntervalDayLiteral().GetText())
	value, err := literal.NewIntervalDaysToSecondFromString(interval)
	if err != nil {
		v.ErrorListener.ReportVisitError(fmt.Errorf("invalid interval day arg %v", err))
	}
	return &CaseLiteral{Value: value, ValueText: ctx.IntervalDayLiteral().GetText(), Type: &types.IntervalDayType{}}
}

func (v *TestCaseVisitor) VisitDecimalArg(ctx *baseparser.DecimalArgContext) interface{} {
	decimal, err := literal.NewDecimalFromString(ctx.NumericLiteral().GetText())
	if err != nil {
		v.ErrorListener.ReportVisitError(err)
	}
	decType := v.Visit(ctx.DecimalType()).(types.Type)
	return &CaseLiteral{Value: decimal, ValueText: ctx.NumericLiteral().GetText(), Type: decType}
}

func (v *TestCaseVisitor) VisitListArg(ctx *baseparser.ListArgContext) interface{} {
	listType := v.Visit(ctx.ListType()).(*types.ListType)
	values := v.Visit(ctx.LiteralList()).([]expr.Literal)

	value, err := v.getListLiteral(values, listType.Type)
	if err != nil {
		v.ErrorListener.ReportVisitError(fmt.Errorf("invalid list arg %v", err))
	}
	return &CaseLiteral{Value: value, Type: listType}
}

func (v *TestCaseVisitor) VisitLiteralList(ctx *baseparser.LiteralListContext) interface{} {
	literals := make([]expr.Literal, 0, len(ctx.AllLiteral()))
	for _, literalCtx := range ctx.AllLiteral() {
		literals = append(literals, v.Visit(literalCtx).(expr.Literal))
	}
	return literals
}

func (v *TestCaseVisitor) VisitLiteral(ctx *baseparser.LiteralContext) interface{} {
	if ctx.BooleanLiteral() != nil {
		flag := strings.ToLower(ctx.BooleanLiteral().GetText()) == "true"
		return literal.NewBool(flag)
	}
	if ctx.DateLiteral() != nil {
		dateStr := getRawStringFromStringLiteral(ctx.DateLiteral().GetText())
		value, err := literal.NewDateFromString(dateStr)
		if err != nil {
			v.ErrorListener.ReportVisitError(fmt.Errorf("invalid date arg %v", err))
		}
		return value
	}
	if ctx.TimeLiteral() != nil {
		timeStr := getRawStringFromStringLiteral(ctx.TimeLiteral().GetText())
		value, err := literal.NewTimeFromString(timeStr)
		if err != nil {
			v.ErrorListener.ReportVisitError(fmt.Errorf("invalid time arg %v", err))
		}
		return value
	}
	if ctx.TimestampLiteral() != nil {
		timestampStr := getRawStringFromStringLiteral(ctx.TimestampLiteral().GetText())
		value, err := literal.NewTimestampFromString(timestampStr)
		if err != nil {
			v.ErrorListener.ReportVisitError(fmt.Errorf("invalid timestampTZ arg %v", err))
		}
		return value
	}
	if ctx.TimestampTzLiteral() != nil {
		timestampStr := getRawStringFromStringLiteral(ctx.TimestampTzLiteral().GetText())
		value, err := literal.NewTimestampTZFromString(timestampStr)
		if err != nil {
			v.ErrorListener.ReportVisitError(fmt.Errorf("invalid timestampTZ arg %v", err))
		}
		return value
	}
	if ctx.IntervalYearLiteral() != nil {
		iyearStr := getRawStringFromStringLiteral(ctx.IntervalYearLiteral().GetText())
		value, err := literal.NewIntervalYearsToMonthFromString(iyearStr)
		if err != nil {
			v.ErrorListener.ReportVisitError(fmt.Errorf("invalid interval year arg %v", err))
		}
		return value
	}
	if ctx.IntervalDayLiteral() != nil {
		idayStr := getRawStringFromStringLiteral(ctx.IntervalDayLiteral().GetText())
		value, err := literal.NewIntervalDaysToSecondFromString(idayStr)
		if err != nil {
			v.ErrorListener.ReportVisitError(fmt.Errorf("invalid interval day arg %v", err))
		}
		return value
	}

	if ctx.NumericLiteral() != nil {
		value := literal.NewString(ctx.NumericLiteral().GetText())
		return value
	}

	if ctx.StringLiteral() != nil {
		valueStr := getRawStringFromStringLiteral(ctx.GetText())
		value := literal.NewString(valueStr)
		return value
	}

	if ctx.NullLiteral() != nil {
		return nil
	}
	v.ErrorListener.ReportVisitError(fmt.Errorf("invalid literal arg %v", ctx.GetText()))
	return nil
}

func (v *TestCaseVisitor) VisitResult(ctx *baseparser.ResultContext) interface{} {
	if ctx.SubstraitError() != nil {
		return v.Visit(ctx.SubstraitError())
	}
	return v.Visit(ctx.Argument()).(*CaseLiteral)
}

type SubstraitError struct {
	Error string
}

func (v *TestCaseVisitor) VisitSubstraitError(ctx *baseparser.SubstraitErrorContext) interface{} {
	err := &SubstraitError{Error: "UNKNOWN"}
	if ctx.ErrorResult() != nil {
		err.Error = "ERROR"
	} else if ctx.UndefineResult() != nil {
		err.Error = "UNDEFINED"
	}
	return &CaseLiteral{
		Type:           nil,
		ValueText:      ctx.GetText(),
		SubstraitError: err,
	}
}

func (v *TestCaseVisitor) VisitBoolean(*baseparser.BooleanContext) interface{} {
	return &types.BooleanType{}
}

func (v *TestCaseVisitor) VisitI8(*baseparser.I8Context) interface{} {
	return &types.Int8Type{}
}

func (v *TestCaseVisitor) VisitI16(*baseparser.I16Context) interface{} {
	return &types.Int16Type{}
}

func (v *TestCaseVisitor) VisitI32(*baseparser.I32Context) interface{} {
	return &types.Int32Type{}
}

func (v *TestCaseVisitor) VisitI64(*baseparser.I64Context) interface{} {
	return &types.Int64Type{}
}

func (v *TestCaseVisitor) VisitFp32(*baseparser.Fp32Context) interface{} {
	return &types.Float32Type{}
}

func (v *TestCaseVisitor) VisitFp64(*baseparser.Fp64Context) interface{} {
	return &types.Float64Type{}
}

func (v *TestCaseVisitor) VisitString(*baseparser.StringContext) interface{} {
	return &types.StringType{}
}

func (v *TestCaseVisitor) VisitBinary(*baseparser.BinaryContext) interface{} {
	return &types.BinaryType{}
}

func (v *TestCaseVisitor) VisitTimestamp(*baseparser.TimestampContext) interface{} {
	return &types.TimestampType{}
}

func (v *TestCaseVisitor) VisitTimestampTz(*baseparser.TimestampTzContext) interface{} {
	return &types.TimestampTzType{}
}

func (v *TestCaseVisitor) VisitDate(*baseparser.DateContext) interface{} {
	return &types.DateType{}
}

func (v *TestCaseVisitor) VisitTime(*baseparser.TimeContext) interface{} {
	return &types.TimeType{}
}

func (v *TestCaseVisitor) VisitIntervalYear(*baseparser.IntervalYearContext) interface{} {
	return &types.IntervalYearType{}
}

func (v *TestCaseVisitor) VisitIntervalDay(*baseparser.IntervalDayContext) interface{} {
	return &types.IntervalDayType{}
}

func (v *TestCaseVisitor) VisitUuid(*baseparser.UuidContext) interface{} {
	return &types.UUIDType{}
}

func (v *TestCaseVisitor) VisitList(ctx *baseparser.ListContext) interface{} {
	elementType := v.Visit(ctx.GetElemType()).(types.Type)
	return &types.ListType{Type: elementType}
}

func (v *TestCaseVisitor) VisitDataType(ctx *baseparser.DataTypeContext) interface{} {
	if ctx.ScalarType() != nil {
		return v.Visit(ctx.ScalarType())
	}
	return v.Visit(ctx.ParameterizedType())
}

func (v *TestCaseVisitor) VisitParameterizedType(ctx *baseparser.ParameterizedTypeContext) interface{} {
	if ctx.DecimalType() != nil {
		return v.Visit(ctx.DecimalType())
	}
	return nil
}

func (v *TestCaseVisitor) VisitDecimal(ctx *baseparser.DecimalContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}

	precision := int32(38)
	scale := int32(0)
	if ctx.GetPrecision() != nil {
		precision = v.Visit(ctx.GetPrecision()).(int32)
		scale = v.Visit(ctx.GetScale()).(int32)
	}
	return &types.DecimalType{Precision: precision, Scale: scale, Nullability: nullability}
}

func (v *TestCaseVisitor) VisitIntegerLiteral(ctx *baseparser.IntegerLiteralContext) interface{} {
	value, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		v.ErrorListener.ReportVisitError(fmt.Errorf("invalid int arg %v", err))
	}
	return int32(value)
}

func (v *TestCaseVisitor) VisitPrecisionTimestamp(ctx *baseparser.PrecisionTimestampContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}

	length := v.Visit(ctx.GetPrecision()).(int32)
	return &types.PrecisionTimestampType{Precision: types.TimePrecision(length), Nullability: nullability}
}

func (v *TestCaseVisitor) VisitPrecisionTimestampTZ(ctx *baseparser.PrecisionTimestampTZContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}

	length := v.Visit(ctx.GetPrecision()).(int32)
	return &types.PrecisionTimestampTzType{
		PrecisionTimestampType: types.PrecisionTimestampType{
			Precision:   types.TimePrecision(length),
			Nullability: nullability,
		},
	}
}

func (v *TestCaseVisitor) VisitFixedChar(ctx *baseparser.FixedCharContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}

	length := v.Visit(ctx.GetLen_()).(int32)
	return &types.FixedCharType{Length: length, Nullability: nullability}
}

func (v *TestCaseVisitor) VisitVarChar(ctx *baseparser.VarCharContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}

	length := v.Visit(ctx.GetLen_()).(int32)
	return &types.VarCharType{Length: length, Nullability: nullability}
}

func (v *TestCaseVisitor) VisitFixedBinary(ctx *baseparser.FixedBinaryContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}

	length := v.Visit(ctx.GetLen_()).(int32)
	return &types.FixedBinaryType{Length: length, Nullability: nullability}
}
