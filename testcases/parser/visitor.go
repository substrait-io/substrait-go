package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/substrait-io/substrait-go/v3/expr"
	"github.com/substrait-io/substrait-go/v3/literal"
	"github.com/substrait-io/substrait-go/v3/testcases/parser/baseparser"
	"github.com/substrait-io/substrait-go/v3/types"
	"github.com/substrait-io/substrait-go/v3/types/parser/util"
)

type TestCaseVisitor struct {
	baseparser.FuncTestCaseParserVisitor
	ErrorListener        util.VisitErrorListener
	literalTypeInContext types.Type
	testFuncType         TestFuncType
}

func (v *TestCaseVisitor) getLiteralTypeInContext() types.Type {
	return v.literalTypeInContext
}

func (v *TestCaseVisitor) setLiteralTypeInContext(t types.Type) {
	v.literalTypeInContext = t
}

func (v *TestCaseVisitor) clearLiteralTypeInContext() {
	v.literalTypeInContext = nil
}

var _ baseparser.FuncTestCaseParserVisitor = &TestCaseVisitor{}

func (v *TestCaseVisitor) Visit(tree antlr.ParseTree) interface{} {
	if tree == nil {
		return nil
	}
	return tree.Accept(v)
}

func (v *TestCaseVisitor) VisitDoc(ctx *baseparser.DocContext) interface{} {
	header := v.Visit(ctx.Header()).(*TestFileHeader)
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
	header := v.Visit(ctx.Version()).(*TestFileHeader)
	header.IncludedURI = v.Visit(ctx.Include()).(string)
	return header
}

func (v *TestCaseVisitor) VisitVersion(ctx *baseparser.VersionContext) interface{} {
	testFuncType := ScalarFuncType
	if ctx.SubstraitAggregateTest() != nil {
		testFuncType = AggregateFuncType
	}
	v.testFuncType = testFuncType
	return &TestFileHeader{
		Version:  ctx.FormatVersion().GetText(),
		FuncType: testFuncType,
	}
}

func (v *TestCaseVisitor) VisitInclude(ctx *baseparser.IncludeContext) interface{} {
	return getRawStringFromStringLiteral(ctx.StringLiteral(0).GetText())
}

func (v *TestCaseVisitor) VisitScalarFuncTestGroup(ctx *baseparser.ScalarFuncTestGroupContext) interface{} {
	groupDesc := v.Visit(ctx.TestGroupDescription()).(string)
	groupTestCases := make([]*TestCase, 0, len(ctx.AllTestCase()))
	if v.testFuncType != ScalarFuncType {
		v.ErrorListener.ReportVisitError(fmt.Errorf("expected %v testcase based on test file header, but got scalar function testcase", v.testFuncType))
		return groupTestCases
	}
	for _, tc := range ctx.AllTestCase() {
		testcase := v.Visit(tc).(*TestCase)
		testcase.GroupDesc = groupDesc
		testcase.FuncType = ScalarFuncType
		groupTestCases = append(groupTestCases, testcase)
	}
	return groupTestCases
}

func (v *TestCaseVisitor) VisitAggregateFuncTestGroup(ctx *baseparser.AggregateFuncTestGroupContext) interface{} {
	groupDesc := v.Visit(ctx.TestGroupDescription()).(string)
	groupTestCases := make([]*TestCase, 0, len(ctx.AllAggFuncTestCase()))
	if v.testFuncType != AggregateFuncType {
		v.ErrorListener.ReportVisitError(fmt.Errorf("expected %v testcase based on test file header, but got aggregate function testcase", v.testFuncType))
		return groupTestCases
	}
	for _, tc := range ctx.AllAggFuncTestCase() {
		testcase := v.Visit(tc).(*TestCase)
		testcase.GroupDesc = groupDesc
		testcase.FuncType = AggregateFuncType
		groupTestCases = append(groupTestCases, testcase)
	}
	return groupTestCases
}

func (v *TestCaseVisitor) VisitAggFuncTestCase(ctx *baseparser.AggFuncTestCaseContext) interface{} {
	testcase := v.Visit(ctx.AggFuncCall()).(*TestCase)
	testcase.Result = v.Visit(ctx.Result()).(*CaseLiteral)
	if ctx.FuncOptions() != nil {
		testcase.Options = v.Visit(ctx.FuncOptions()).(FuncOptions)
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
	var args []*AggregateArgument
	if ctx.AggregateFuncArgs() != nil {
		args = v.Visit(ctx.AggregateFuncArgs()).([]*AggregateArgument)
	}

	numberOfColumns := int32(len(rows[0]))
	columnTypes := make([]types.Type, numberOfColumns)
	for _, arg := range args {
		if arg.ColumnIndex >= numberOfColumns {
			v.ErrorListener.ReportVisitError(fmt.Errorf("invalid column index %d, expected less than %d", arg.ColumnIndex, len(columnTypes)))
			continue
		}
		if arg.ColumnType != nil {
			columnTypes[arg.ColumnIndex] = arg.ColumnType
		}
	}
	columns, newColumnTypes := v.getColumnsFromRows(rows, columnTypes)
	return &TestCase{
		FuncName:      ctx.Identifier().GetText(),
		Columns:       columns,
		AggregateArgs: args,
		ColumnTypes:   newColumnTypes,
	}
}

func (v *TestCaseVisitor) VisitMultiArgAggregateFuncCall(ctx *baseparser.MultiArgAggregateFuncCallContext) interface{} {
	testcase := v.Visit(ctx.TableData()).(*TestCase)
	var args []*AggregateArgument
	if ctx.QualifiedAggregateFuncArgs() != nil {
		args = v.Visit(ctx.QualifiedAggregateFuncArgs()).([]*AggregateArgument)
	}
	testcase.FuncName = ctx.Identifier().GetText()
	testcase.AggregateArgs = args
	for _, arg := range args {
		if arg.TableName != "" {
			if testcase.TableName != arg.TableName {
				err := fmt.Errorf("table name in argument %s, does not match the table name in the function call %s", arg.TableName, testcase.TableName)
				v.ErrorListener.ReportVisitError(err)
			}
		}
		if !arg.IsScalar {
			arg.ColumnType = testcase.ColumnTypes[arg.ColumnIndex]
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
			IsScalar: true,
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
	columns, newColumnTypes := v.getColumnsFromRows(rows, columnTypes)
	return &TestCase{
		Columns:     columns,
		TableName:   ctx.GetTableName().GetText(),
		ColumnTypes: newColumnTypes,
	}
}

func (v *TestCaseVisitor) getColumnsFromRows(rows [][]expr.Literal, columnTypes []types.Type) ([][]expr.Literal, []types.Type) {
	columns := make([][]expr.Literal, 0, len(columnTypes))
	retColumnTypes := make([]types.Type, 0, len(columnTypes))
	for i, columnType := range columnTypes {
		newColumnType := columnType
		column := make([]expr.Literal, 0, len(rows))
		for _, row := range rows {
			cell := v.getLiteral(row[i], columnType)
			if _, ok := cell.(*expr.NullLiteral); ok {
				newColumnType = columnType.WithNullability(types.NullabilityNullable)
				cell = expr.NewNullLiteral(newColumnType)
			}
			column = append(column, cell)
		}
		retColumnTypes = append(retColumnTypes, newColumnType)
		columns = append(columns, column)
	}
	return columns, retColumnTypes
}

func (v *TestCaseVisitor) VisitDataColumn(ctx *baseparser.DataColumnContext) interface{} {
	columnType := v.Visit(ctx.DataType()).(types.Type)
	v.setLiteralTypeInContext(columnType)
	defer v.clearLiteralTypeInContext()
	columnValues := v.Visit(ctx.ColumnValues()).([]expr.Literal)
	var err error
	var column expr.Literal
	if len(columnValues) == 0 {
		column = expr.NewEmptyListLiteral(columnType, false)
	} else {
		v.setLiteralTypeInContext(columnType)
		defer v.clearLiteralTypeInContext()
		column, err = literal.NewList(columnValues)
		if err != nil {
			v.ErrorListener.ReportVisitError(fmt.Errorf("invalid column values %v", err))
		}
	}
	for _, value := range columnValues {
		if _, ok := value.(*expr.NullLiteral); ok {
			columnType = columnType.WithNullability(types.NullabilityNullable)
			break
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
			IsScalar: true,
		}
	}
	argType := v.Visit(ctx.DataType()).(types.Type)
	arg, err := newAggregateArgument("", ctx.ColumnName().GetText(), argType)
	if err != nil {
		v.ErrorListener.ReportVisitError(fmt.Errorf("invalid aggregate func arg %v", err))
	}
	return arg
}

func (v *TestCaseVisitor) VisitFuncOptions(ctx *baseparser.FuncOptionsContext) interface{} {
	options := make(FuncOptions)
	for _, option := range ctx.AllFuncOption() {
		optionPair := v.Visit(option).([]string)
		options[optionPair[0]] = optionPair[1]
	}
	return options
}

func (v *TestCaseVisitor) VisitFuncOption(ctx *baseparser.FuncOptionContext) interface{} {
	return []string{ctx.OptionName().GetText(), ctx.OptionValue().GetText()}
}

func (v *TestCaseVisitor) VisitTestGroupDescription(ctx *baseparser.TestGroupDescriptionContext) interface{} {
	return strings.TrimSpace(strings.TrimPrefix(ctx.GetText(), "#"))
}

func (v *TestCaseVisitor) VisitTestCase(ctx *baseparser.TestCaseContext) interface{} {
	var options FuncOptions
	if ctx.FuncOptions() != nil {
		options = v.Visit(ctx.FuncOptions()).(FuncOptions)
	}

	return &TestCase{
		FuncName: ctx.Identifier().GetText(),
		Args:     v.Visit(ctx.Arguments()).([]*CaseLiteral),
		Result:   v.Visit(ctx.Result()).(*CaseLiteral),
		Options:  options,
	}
}

func (v *TestCaseVisitor) VisitArguments(ctx *baseparser.ArgumentsContext) interface{} {
	args := make([]*CaseLiteral, 0, len(ctx.AllArgument()))
	for _, argument := range ctx.AllArgument() {
		testArg := v.Visit(argument).(*CaseLiteral)
		if err := testArg.updateLiteralType(); err != nil {
			v.ErrorListener.ReportVisitError(fmt.Errorf("invalid argument %v", err))
		}
		args = append(args, testArg)
	}
	return args
}

func (v *TestCaseVisitor) VisitArgument(ctx *baseparser.ArgumentContext) interface{} {
	if ctx.NullArg() != nil {
		return v.Visit(ctx.NullArg())
	}
	if ctx.IntArg() != nil {
		return v.Visit(ctx.IntArg())
	}
	if ctx.FloatArg() != nil {
		return v.Visit(ctx.FloatArg())
	}
	if ctx.BooleanArg() != nil {
		return v.Visit(ctx.BooleanArg())
	}
	if ctx.StringArg() != nil {
		return v.Visit(ctx.StringArg())
	}
	if ctx.DecimalArg() != nil {
		return v.Visit(ctx.DecimalArg())
	}
	if ctx.DateArg() != nil {
		return v.Visit(ctx.DateArg())
	}
	if ctx.TimeArg() != nil {
		return v.Visit(ctx.TimeArg())
	}
	if ctx.TimestampArg() != nil {
		return v.Visit(ctx.TimestampArg())
	}
	if ctx.TimestampTzArg() != nil {
		return v.Visit(ctx.TimestampTzArg())
	}
	if ctx.IntervalYearArg() != nil {
		return v.Visit(ctx.IntervalYearArg())
	}
	if ctx.IntervalDayArg() != nil {
		return v.Visit(ctx.IntervalDayArg())
	}
	if ctx.FixedCharArg() != nil {
		return v.Visit(ctx.FixedCharArg())
	}
	if ctx.VarCharArg() != nil {
		return v.Visit(ctx.VarCharArg())
	}
	if ctx.FixedBinaryArg() != nil {
		return v.Visit(ctx.FixedBinaryArg())
	}
	if ctx.PrecisionTimestampArg() != nil {
		return v.Visit(ctx.PrecisionTimestampArg())
	}
	if ctx.PrecisionTimestampTZArg() != nil {
		return v.Visit(ctx.PrecisionTimestampTZArg())
	}
	if ctx.ListArg() != nil {
		return v.Visit(ctx.ListArg())
	}

	v.ErrorListener.ReportVisitError(fmt.Errorf("argument type not implemented, arg %s", ctx.GetText()))
	return &CaseLiteral{}
}

func (v *TestCaseVisitor) VisitNullArg(ctx *baseparser.NullArgContext) interface{} {
	dataType := v.Visit(ctx.DataType()).(types.Type)
	return &CaseLiteral{Value: expr.NewNullLiteral(dataType), ValueText: ctx.NullLiteral().GetText(), Type: dataType.WithNullability(types.NullabilityNullable)}
}

func (v *TestCaseVisitor) VisitBooleanArg(ctx *baseparser.BooleanArgContext) interface{} {
	value := strings.ToLower(ctx.BooleanLiteral().GetText()) == "true"
	boolLiteral := literal.NewBool(value)
	return &CaseLiteral{Value: boolLiteral, ValueText: ctx.BooleanLiteral().GetText(), Type: &types.BooleanType{Nullability: types.NullabilityRequired}}
}

func (v *TestCaseVisitor) getLiteral(value expr.Literal, elementType types.Type) expr.Literal {
	strLiteral, ok := value.(*expr.PrimitiveLiteral[string])
	if !ok {
		return value
	}
	ret := v.getLiteralFromString(strLiteral.Value, elementType)
	if ret == nil {
		return value
	}
	return ret
}

func (v *TestCaseVisitor) getLiteralFromString(value string, elementType types.Type) expr.Literal {
	switch elementType := elementType.(type) {
	case *types.Int8Type, *types.Int16Type, *types.Int32Type, *types.Int64Type:
		intLiteral, err := getIntLiteral(value, elementType)
		if err != nil {
			v.ErrorListener.ReportVisitError(fmt.Errorf("invalid int arg %v", err))
		}
		return intLiteral
	case *types.Float32Type, *types.Float64Type:
		floatLiteral, err := getFloatLiteral(value, elementType)
		if err != nil {
			v.ErrorListener.ReportVisitError(fmt.Errorf("invalid float arg %v", err))
		}
		return floatLiteral
	case *types.DecimalType:
		decimal, err := literal.NewDecimalFromString(value)
		if err != nil {
			v.ErrorListener.ReportVisitError(fmt.Errorf("invalid decimal arg %v", err))
		}
		return decimal
	default:
		return nil
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
	typ = &types.Int8Type{Nullability: types.NullabilityRequired}
	if ctx.I16() != nil {
		typ = &types.Int16Type{Nullability: types.NullabilityRequired}
	} else if ctx.I32() != nil {
		typ = &types.Int32Type{Nullability: types.NullabilityRequired}
	} else if ctx.I64() != nil {
		typ = &types.Int64Type{Nullability: types.NullabilityRequired}
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
		floatLiteral, err = getFloatLiteral(ctx.NumericLiteral().GetText(), &types.Float32Type{Nullability: types.NullabilityRequired})
	} else {
		floatLiteral, err = getFloatLiteral(ctx.NumericLiteral().GetText(), &types.Float64Type{Nullability: types.NullabilityRequired})
	}
	if err != nil {
		v.ErrorListener.ReportVisitError(fmt.Errorf("invalid float arg %v", err))
	}
	return &CaseLiteral{Value: floatLiteral, ValueText: ctx.NumericLiteral().GetText(), Type: floatLiteral.GetType()}
}

func (v *TestCaseVisitor) VisitStringArg(ctx *baseparser.StringArgContext) interface{} {
	value := literal.NewString(getRawStringFromStringLiteral(ctx.StringLiteral().GetText()))
	return &CaseLiteral{Value: value, ValueText: ctx.StringLiteral().GetText(), Type: &types.StringType{Nullability: types.NullabilityRequired}}
}

func (v *TestCaseVisitor) VisitFixedCharArg(ctx *baseparser.FixedCharArgContext) interface{} {
	value, err := literal.NewFixedChar(getRawStringFromStringLiteral(ctx.StringLiteral().GetText()))
	if err != nil {
		v.ErrorListener.ReportVisitError(fmt.Errorf("invalid fixed char arg %v", err))
	}
	argType := v.Visit(ctx.FixedCharType()).(*types.FixedCharType)
	return &CaseLiteral{Value: value, ValueText: ctx.StringLiteral().GetText(), Type: argType}
}

func (v *TestCaseVisitor) VisitVarCharArg(ctx *baseparser.VarCharArgContext) interface{} {
	value, err := literal.NewVarChar(getRawStringFromStringLiteral(ctx.StringLiteral().GetText()))
	if err != nil {
		v.ErrorListener.ReportVisitError(fmt.Errorf("invalid fixed char arg %v", err))
	}
	argType := v.Visit(ctx.VarCharType()).(*types.VarCharType)
	return &CaseLiteral{Value: value, ValueText: ctx.StringLiteral().GetText(), Type: argType}
}

func (v *TestCaseVisitor) VisitFixedBinaryArg(ctx *baseparser.FixedBinaryArgContext) interface{} {
	value, err := literal.NewFixedBinary([]byte(getRawStringFromStringLiteral(ctx.StringLiteral().GetText())))
	if err != nil {
		v.ErrorListener.ReportVisitError(fmt.Errorf("invalid fixed binary arg %v", err))
	}
	argType := v.Visit(ctx.FixedBinaryType()).(*types.FixedBinaryType)
	return &CaseLiteral{Value: value, ValueText: ctx.StringLiteral().GetText(), Type: argType}
}

func (v *TestCaseVisitor) VisitTimestampArg(ctx *baseparser.TimestampArgContext) interface{} {
	timestampStr := getRawStringFromStringLiteral(ctx.TimestampLiteral().GetText())
	value, err := literal.NewTimestampFromString(timestampStr)
	if err != nil {
		v.ErrorListener.ReportVisitError(fmt.Errorf("invalid timestampTZ arg %v", err))
	}
	return &CaseLiteral{Value: value, ValueText: ctx.TimestampLiteral().GetText(), Type: &types.TimestampType{Nullability: types.NullabilityRequired}}
}

func getRawStringFromStringLiteral(text string) string {
	return strings.Trim(text, "'")
}

func (v *TestCaseVisitor) VisitTimestampTzArg(ctx *baseparser.TimestampTzArgContext) interface{} {
	value, err := literal.NewTimestampTZFromString(getRawStringFromStringLiteral(ctx.TimestampTzLiteral().GetText()))
	if err != nil {
		v.ErrorListener.ReportVisitError(fmt.Errorf("invalid timestampTZ arg %v", err))
	}
	return &CaseLiteral{Value: value, ValueText: ctx.TimestampTzLiteral().GetText(), Type: &types.TimestampTzType{Nullability: types.NullabilityRequired}}
}

func (v *TestCaseVisitor) VisitDateArg(ctx *baseparser.DateArgContext) interface{} {
	value, err := literal.NewDateFromString(getRawStringFromStringLiteral(ctx.DateLiteral().GetText()))
	if err != nil {
		v.ErrorListener.ReportVisitError(fmt.Errorf("invalid date arg %v", err))
	}
	return &CaseLiteral{Value: value, ValueText: ctx.DateLiteral().GetText(), Type: &types.DateType{Nullability: types.NullabilityRequired}}
}

func (v *TestCaseVisitor) VisitTimeArg(ctx *baseparser.TimeArgContext) interface{} {
	value, err := literal.NewTimeFromString(getRawStringFromStringLiteral(ctx.TimeLiteral().GetText()))
	if err != nil {
		v.ErrorListener.ReportVisitError(fmt.Errorf("invalid time arg %v", err))
	}
	return &CaseLiteral{Value: value, ValueText: ctx.TimeLiteral().GetText(), Type: &types.TimeType{Nullability: types.NullabilityRequired}}
}

func (v *TestCaseVisitor) VisitIntervalYearArg(ctx *baseparser.IntervalYearArgContext) interface{} {
	interval := getRawStringFromStringLiteral(ctx.IntervalYearLiteral().GetText())
	value, err := literal.NewIntervalYearsToMonthFromString(interval)
	if err != nil {
		v.ErrorListener.ReportVisitError(fmt.Errorf("invalid interval year arg %v", err))
	}
	return &CaseLiteral{Value: value, ValueText: ctx.IntervalYearLiteral().GetText(), Type: &types.IntervalYearType{Nullability: types.NullabilityRequired}}
}

func (v *TestCaseVisitor) VisitIntervalDayArg(ctx *baseparser.IntervalDayArgContext) interface{} {
	interval := getRawStringFromStringLiteral(ctx.IntervalDayLiteral().GetText())
	value, err := literal.NewIntervalDaysToSecondFromString(interval)
	if err != nil {
		v.ErrorListener.ReportVisitError(fmt.Errorf("invalid interval day arg %v", err))
	}
	idayType := v.Visit(ctx.IntervalDayType()).(*types.IntervalDayType)
	return &CaseLiteral{Value: value, ValueText: ctx.IntervalDayLiteral().GetText(), Type: idayType}
}

func (v *TestCaseVisitor) VisitDecimalArg(ctx *baseparser.DecimalArgContext) interface{} {
	decimal, err := literal.NewDecimalFromString(ctx.NumericLiteral().GetText())
	if err != nil {
		v.ErrorListener.ReportVisitError(err)
	}
	decType := v.Visit(ctx.DecimalType()).(types.Type)
	return &CaseLiteral{Value: decimal, ValueText: ctx.NumericLiteral().GetText(), Type: decType}
}

func (v *TestCaseVisitor) VisitPrecisionTimestampArg(ctx *baseparser.PrecisionTimestampArgContext) interface{} {
	ptsType := v.Visit(ctx.PrecisionTimestampType()).(*types.PrecisionTimestampType)
	timestampStr := getRawStringFromStringLiteral(ctx.TimestampLiteral().GetText())
	value, err := literal.NewPrecisionTimestampFromString(ptsType.Precision, timestampStr)
	if err != nil {
		v.ErrorListener.ReportVisitError(fmt.Errorf("invalid precision timestamp arg %v", err))
	}
	return &CaseLiteral{Value: value, ValueText: ctx.TimestampLiteral().GetText(), Type: ptsType}
}

func (v *TestCaseVisitor) VisitPrecisionTimestampTZArg(ctx *baseparser.PrecisionTimestampTZArgContext) interface{} {
	ptszType := v.Visit(ctx.PrecisionTimestampTZType()).(*types.PrecisionTimestampTzType)
	timestampStr := getRawStringFromStringLiteral(ctx.TimestampTzLiteral().GetText())
	value, err := literal.NewPrecisionTimestampTzFromString(ptszType.Precision, timestampStr)
	if err != nil {
		v.ErrorListener.ReportVisitError(fmt.Errorf("invalid precision timestampTZ arg %v", err))
	}
	return &CaseLiteral{Value: value, ValueText: ctx.TimestampTzLiteral().GetText(), Type: ptszType}
}

func (v *TestCaseVisitor) VisitListArg(ctx *baseparser.ListArgContext) interface{} {
	listType := v.Visit(ctx.ListType()).(*types.ListType)
	v.setLiteralTypeInContext(listType.Type)
	defer v.clearLiteralTypeInContext()

	values := v.Visit(ctx.LiteralList()).([]expr.Literal)

	value, err := literal.NewList(values)
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
		if v.getLiteralTypeInContext() == nil {
			// in compactAggregateFuncCall context, the type is not set, full schema of table may not be available
			return literal.NewString(ctx.NumericLiteral().GetText())
		}
		value := v.getLiteralFromString(ctx.NumericLiteral().GetText(), v.getLiteralTypeInContext())
		if value == nil {
			v.ErrorListener.ReportVisitError(fmt.Errorf("invalid numeric arg %v", ctx.GetText()))
		}
		return value
	}

	if ctx.StringLiteral() != nil {
		valueStr := getRawStringFromStringLiteral(ctx.GetText())
		value := literal.NewString(valueStr)
		return value
	}

	if ctx.NullLiteral() != nil {
		nullType := v.getLiteralTypeInContext()
		if nullType == nil {
			// Use a dummy type for null literal. This happens in AggregateFuncCall context, where type is not set
			nullType = &types.DecimalType{Precision: 38, Scale: 0, Nullability: types.NullabilityNullable}
		}
		return expr.NewNullLiteral(nullType)
	}
	v.ErrorListener.ReportVisitError(fmt.Errorf("invalid literal arg %v", ctx.GetText()))
	return nil
}

func (v *TestCaseVisitor) VisitResult(ctx *baseparser.ResultContext) interface{} {
	if ctx.SubstraitError() != nil {
		return v.Visit(ctx.SubstraitError())
	}
	result := v.Visit(ctx.Argument()).(*CaseLiteral)
	if err := result.updateLiteralType(); err != nil {
		v.ErrorListener.ReportVisitError(fmt.Errorf("invalid result: %v", err))
	}
	return result
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
	return &types.BooleanType{Nullability: types.NullabilityRequired}
}

func (v *TestCaseVisitor) VisitI8(*baseparser.I8Context) interface{} {
	return &types.Int8Type{Nullability: types.NullabilityRequired}
}

func (v *TestCaseVisitor) VisitI16(*baseparser.I16Context) interface{} {
	return &types.Int16Type{Nullability: types.NullabilityRequired}
}

func (v *TestCaseVisitor) VisitI32(*baseparser.I32Context) interface{} {
	return &types.Int32Type{Nullability: types.NullabilityRequired}
}

func (v *TestCaseVisitor) VisitI64(*baseparser.I64Context) interface{} {
	return &types.Int64Type{Nullability: types.NullabilityRequired}
}

func (v *TestCaseVisitor) VisitFp32(*baseparser.Fp32Context) interface{} {
	return &types.Float32Type{Nullability: types.NullabilityRequired}
}

func (v *TestCaseVisitor) VisitFp64(*baseparser.Fp64Context) interface{} {
	return &types.Float64Type{Nullability: types.NullabilityRequired}
}

func (v *TestCaseVisitor) VisitString(*baseparser.StringContext) interface{} {
	return &types.StringType{Nullability: types.NullabilityRequired}
}

func (v *TestCaseVisitor) VisitBinary(*baseparser.BinaryContext) interface{} {
	return &types.BinaryType{Nullability: types.NullabilityRequired}
}

func (v *TestCaseVisitor) VisitTimestamp(*baseparser.TimestampContext) interface{} {
	return &types.TimestampType{Nullability: types.NullabilityRequired}
}

func (v *TestCaseVisitor) VisitTimestampTz(*baseparser.TimestampTzContext) interface{} {
	return &types.TimestampTzType{Nullability: types.NullabilityRequired}
}

func (v *TestCaseVisitor) VisitDate(*baseparser.DateContext) interface{} {
	return &types.DateType{Nullability: types.NullabilityRequired}
}

func (v *TestCaseVisitor) VisitTime(*baseparser.TimeContext) interface{} {
	return &types.TimeType{Nullability: types.NullabilityRequired}
}

func (v *TestCaseVisitor) VisitIntervalYear(*baseparser.IntervalYearContext) interface{} {
	return &types.IntervalYearType{Nullability: types.NullabilityRequired}
}

func (v *TestCaseVisitor) VisitUuid(*baseparser.UuidContext) interface{} {
	return &types.UUIDType{Nullability: types.NullabilityRequired}
}

func (v *TestCaseVisitor) VisitList(ctx *baseparser.ListContext) interface{} {
	elementType := v.Visit(ctx.GetElemType()).(types.Type)
	return &types.ListType{Type: elementType, Nullability: types.NullabilityRequired}
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
	if ctx.PrecisionTimestampType() != nil {
		return v.Visit(ctx.PrecisionTimestampType())
	}
	if ctx.PrecisionTimestampTZType() != nil {
		return v.Visit(ctx.PrecisionTimestampTZType())
	}
	if ctx.IntervalDayType() != nil {
		return v.Visit(ctx.IntervalDayType())
	}
	if ctx.FixedCharType() != nil {
		return v.Visit(ctx.FixedCharType())
	}
	if ctx.VarCharType() != nil {
		return v.Visit(ctx.VarCharType())
	}
	if ctx.FixedBinaryType() != nil {
		return v.Visit(ctx.FixedBinaryType())
	}

	return nil
}

func (v *TestCaseVisitor) VisitDecimalType(ctx *baseparser.DecimalTypeContext) interface{} {
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

func (v *TestCaseVisitor) VisitPrecisionTimestampType(ctx *baseparser.PrecisionTimestampTypeContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}

	length := v.Visit(ctx.GetPrecision()).(int32)
	return &types.PrecisionTimestampType{Precision: types.TimePrecision(length), Nullability: nullability}
}

func (v *TestCaseVisitor) VisitPrecisionTimestampTZType(ctx *baseparser.PrecisionTimestampTZTypeContext) interface{} {
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

func (v *TestCaseVisitor) VisitIntervalDayType(ctx *baseparser.IntervalDayTypeContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}

	var length int32
	if ctx.GetLen_() != nil {
		length = v.Visit(ctx.GetLen_()).(int32)
	}
	return &types.IntervalDayType{Precision: types.TimePrecision(length), Nullability: nullability}
}

func (v *TestCaseVisitor) VisitFixedCharType(ctx *baseparser.FixedCharTypeContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}

	length := v.Visit(ctx.GetLen_()).(int32)
	return &types.FixedCharType{Length: length, Nullability: nullability}
}

func (v *TestCaseVisitor) VisitVarCharType(ctx *baseparser.VarCharTypeContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}

	length := v.Visit(ctx.GetLen_()).(int32)
	return &types.VarCharType{Length: length, Nullability: nullability}
}

func (v *TestCaseVisitor) VisitFixedBinaryType(ctx *baseparser.FixedBinaryTypeContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}

	length := v.Visit(ctx.GetLen_()).(int32)
	return &types.FixedBinaryType{Length: length, Nullability: nullability}
}
