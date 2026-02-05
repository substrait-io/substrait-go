package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/substrait-io/substrait-go/v7/expr"
	"github.com/substrait-io/substrait-go/v7/extensions"
	"github.com/substrait-io/substrait-go/v7/functions"
	"github.com/substrait-io/substrait-go/v7/literal"
	"github.com/substrait-io/substrait-go/v7/testcases/parser/baseparser"
	"github.com/substrait-io/substrait-go/v7/types"
	"github.com/substrait-io/substrait-go/v7/types/parser/util"
)

type isNull interface {
	GetIsnull() antlr.Token
}

func getNullability(ctx isNull) types.Nullability {
	if ctx.GetIsnull() != nil {
		return types.NullabilityNullable
	}
	return types.NullabilityRequired
}

type TestCaseVisitor struct {
	baseparser.FuncTestCaseParserVisitor
	ErrorListener        util.VisitErrorListener
	literalTypeInContext types.Type
	testFuncType         TestFuncType
	lambdaParams         []lambdaParamInfo
}

type lambdaParamInfo struct {
	name string
	typ  types.Type
	idx  int32 // index in the parameters struct
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
	// Dependencies are optional and not currently used
	for _, dep := range ctx.AllDependency() {
		_ = v.Visit(dep) // Visit but ignore for now
	}
	return header
}

func (v *TestCaseVisitor) VisitDependency(ctx *baseparser.DependencyContext) interface{} {
	// Dependencies are not currently used, just return empty string
	return ""
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
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("expected %v testcase based on test file header, but got scalar function testcase", v.testFuncType))
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
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("expected %v testcase based on test file header, but got aggregate function testcase", v.testFuncType))
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
			v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid column index %d, expected less than %d", arg.ColumnIndex, len(columnTypes)))
			continue
		}
		if arg.ColumnType != nil {
			columnTypes[arg.ColumnIndex] = arg.ColumnType
		}
	}
	columns, newColumnTypes := v.getColumnsFromRows(ctx, rows, columnTypes)
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
				v.ErrorListener.ReportVisitError(ctx, err)
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
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid aggregate func arg %v", err))
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
	columns, newColumnTypes := v.getColumnsFromRows(ctx, rows, columnTypes)
	return &TestCase{
		Columns:     columns,
		TableName:   ctx.GetTableName().GetText(),
		ColumnTypes: newColumnTypes,
	}
}

func (v *TestCaseVisitor) getColumnsFromRows(ctx antlr.ParserRuleContext, rows [][]expr.Literal, columnTypes []types.Type) ([][]expr.Literal, []types.Type) {
	columns := make([][]expr.Literal, 0, len(columnTypes))
	retColumnTypes := make([]types.Type, 0, len(columnTypes))
	for i, columnType := range columnTypes {
		newColumnType := columnType
		column := make([]expr.Literal, 0, len(rows))
		for _, row := range rows {
			cell := v.getLiteral(ctx, row[i], columnType)
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
		column, err = literal.NewList(columnValues, false)
		if err != nil {
			v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid column values %v", err))
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
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid aggregate func arg %v", err))
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
		result := v.Visit(argument)
		if result == nil {
			v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("argument visit returned nil"))
			args = append(args, &CaseLiteral{})
			continue
		}
		testArg, ok := result.(*CaseLiteral)
		if !ok {
			v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("argument visit returned unexpected type %T", result))
			args = append(args, &CaseLiteral{})
			continue
		}
		if err := testArg.updateLiteralType(); err != nil {
			v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid argument %v", err))
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
	if ctx.PrecisionTimeArg() != nil {
		return v.Visit(ctx.PrecisionTimeArg())
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
	if ctx.LambdaArg() != nil {
		return v.Visit(ctx.LambdaArg())
	}
	if ctx.Identifier() != nil {
		if len(v.lambdaParams) > 0 {
			return v.handleLambdaParameterRef(ctx, ctx.Identifier().GetText())
		}
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("bare identifier %s found outside lambda context", ctx.Identifier().GetText()))
		return &CaseLiteral{ValueText: ctx.Identifier().GetText()}
	}

	v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("argument type not implemented, arg %s", ctx.GetText()))
	return &CaseLiteral{}
}

func (v *TestCaseVisitor) VisitNullArg(ctx *baseparser.NullArgContext) interface{} {
	dataType := v.Visit(ctx.DataType()).(types.Type)
	return &CaseLiteral{Value: expr.NewNullLiteral(dataType), ValueText: ctx.NullLiteral().GetText(), Type: dataType.WithNullability(types.NullabilityNullable)}
}

func (v *TestCaseVisitor) VisitBooleanArg(ctx *baseparser.BooleanArgContext) interface{} {
	value := strings.ToLower(ctx.BooleanLiteral().GetText()) == "true"
	nullability := getNullability(ctx.BooleanType())
	boolLiteral := literal.NewBool(value, nullability == types.NullabilityNullable)
	return &CaseLiteral{Value: boolLiteral, ValueText: ctx.BooleanLiteral().GetText(), Type: &types.BooleanType{Nullability: nullability}}
}

func (v *TestCaseVisitor) getLiteral(ctx antlr.ParserRuleContext, value expr.Literal, elementType types.Type) expr.Literal {
	strLiteral, ok := value.(*expr.PrimitiveLiteral[string])
	if !ok {
		return value
	}
	ret := v.getLiteralFromString(ctx, strLiteral.Value, elementType)
	if ret == nil {
		return value
	}
	return ret
}

func (v *TestCaseVisitor) getLiteralFromString(ctx antlr.ParserRuleContext, value string, elementType types.Type) expr.Literal {
	switch elementType := elementType.(type) {
	case *types.Int8Type, *types.Int16Type, *types.Int32Type, *types.Int64Type:
		intLiteral, err := getIntLiteral(value, elementType)
		if err != nil {
			v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid int arg %v", err))
		}
		return intLiteral
	case *types.Float32Type, *types.Float64Type:
		floatLiteral, err := getFloatLiteral(value, elementType)
		if err != nil {
			v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid float arg %v", err))
		}
		return floatLiteral
	case *types.DecimalType:
		decimal, err := literal.NewDecimalFromString(value, false)
		if err != nil {
			v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid decimal arg %v", err))
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
	nullable := floatType.GetNullability() == types.NullabilityNullable
	switch floatType := floatType.(type) {
	case *types.Float32Type:
		return literal.NewFloat32(float32(value), nullable), nil
	case *types.Float64Type:
		return literal.NewFloat64(value, nullable), nil
	default:
		return nil, fmt.Errorf("invalid float type %v", floatType)
	}
}

func getIntLiteral(integerStr string, intType types.Type) (expr.Literal, error) {
	value, err := strconv.ParseInt(integerStr, 10, 64)
	if err != nil {
		return nil, err
	}
	nullable := intType.GetNullability() == types.NullabilityNullable
	switch intType.(type) {
	case *types.Int8Type:
		return literal.NewInt8(int8(value), nullable), nil
	case *types.Int16Type:
		return literal.NewInt16(int16(value), nullable), nil
	case *types.Int32Type:
		return literal.NewInt32(int32(value), nullable), nil
	case *types.Int64Type:
		return literal.NewInt64(value, nullable), nil
	default:
		return nil, fmt.Errorf("invalid int value %v type %v", value, intType)
	}
}

func (v *TestCaseVisitor) VisitIntArg(ctx *baseparser.IntArgContext) interface{} {
	typ := ctx.IntType().Accept(v).(types.Type)
	intLiteral, err := getIntLiteral(ctx.IntegerLiteral().GetText(), typ)
	if err != nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid int arg %v", err))
	}
	return &CaseLiteral{Value: intLiteral, ValueText: ctx.IntegerLiteral().GetText(), Type: typ}
}

func (v *TestCaseVisitor) VisitFloatArg(ctx *baseparser.FloatArgContext) interface{} {
	typ := ctx.FloatType().Accept(v).(types.Type)
	floatLiteral, err := getFloatLiteral(ctx.NumericLiteral().GetText(), typ)
	if err != nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid float arg %v", err))
	}
	return &CaseLiteral{Value: floatLiteral, ValueText: ctx.NumericLiteral().GetText(), Type: floatLiteral.GetType()}
}

func (v *TestCaseVisitor) VisitStringArg(ctx *baseparser.StringArgContext) interface{} {
	nullability := getNullability(ctx.StringType())
	value := literal.NewString(getRawStringFromStringLiteral(ctx.StringLiteral().GetText()), nullability == types.NullabilityNullable)
	return &CaseLiteral{Value: value, ValueText: ctx.StringLiteral().GetText(), Type: &types.StringType{Nullability: nullability}}
}

func (v *TestCaseVisitor) VisitFixedCharArg(ctx *baseparser.FixedCharArgContext) interface{} {
	argType := v.Visit(ctx.FixedCharType()).(*types.FixedCharType)
	value, err := literal.NewFixedChar(getRawStringFromStringLiteral(ctx.StringLiteral().GetText()), argType.GetNullability() == types.NullabilityNullable)
	if err != nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid fixed char arg %v", err))
	}
	return &CaseLiteral{Value: value, ValueText: ctx.StringLiteral().GetText(), Type: argType}
}

func (v *TestCaseVisitor) VisitVarCharArg(ctx *baseparser.VarCharArgContext) interface{} {
	argType := v.Visit(ctx.VarCharType()).(*types.VarCharType)
	value, err := literal.NewVarChar(getRawStringFromStringLiteral(ctx.StringLiteral().GetText()), argType.GetNullability() == types.NullabilityNullable)
	if err != nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid fixed char arg %v", err))
	}
	return &CaseLiteral{Value: value, ValueText: ctx.StringLiteral().GetText(), Type: argType}
}

func (v *TestCaseVisitor) VisitFixedBinaryArg(ctx *baseparser.FixedBinaryArgContext) interface{} {
	argType := v.Visit(ctx.FixedBinaryType()).(*types.FixedBinaryType)
	value, err := literal.NewFixedBinary([]byte(getRawStringFromStringLiteral(ctx.StringLiteral().GetText())), argType.GetNullability() == types.NullabilityNullable)
	if err != nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid fixed binary arg %v", err))
	}
	return &CaseLiteral{Value: value, ValueText: ctx.StringLiteral().GetText(), Type: argType}
}

func (v *TestCaseVisitor) VisitTimestampArg(ctx *baseparser.TimestampArgContext) interface{} {
	timestampStr := getRawStringFromStringLiteral(ctx.TimestampLiteral().GetText())
	nullability := getNullability(ctx.TimestampType())
	value, err := literal.NewTimestampFromString(timestampStr, nullability == types.NullabilityNullable)
	if err != nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid timestampTZ arg %v", err))
	}
	return &CaseLiteral{Value: value, ValueText: ctx.TimestampLiteral().GetText(), Type: &types.TimestampType{Nullability: nullability}}
}

func getRawStringFromStringLiteral(text string) string {
	return strings.Trim(text, "'")
}

func (v *TestCaseVisitor) VisitTimestampTzArg(ctx *baseparser.TimestampTzArgContext) interface{} {
	nullability := getNullability(ctx.TimestampTZType())
	value, err := literal.NewTimestampTZFromString(getRawStringFromStringLiteral(ctx.TimestampTzLiteral().GetText()), nullability == types.NullabilityNullable)
	if err != nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid timestampTZ arg %v", err))
	}
	return &CaseLiteral{Value: value, ValueText: ctx.TimestampTzLiteral().GetText(), Type: &types.TimestampTzType{Nullability: nullability}}
}

func (v *TestCaseVisitor) VisitDateArg(ctx *baseparser.DateArgContext) interface{} {
	nullability := getNullability(ctx.DateType())
	value, err := literal.NewDateFromString(getRawStringFromStringLiteral(ctx.DateLiteral().GetText()), nullability == types.NullabilityNullable)
	if err != nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid date arg %v", err))
	}
	return &CaseLiteral{Value: value, ValueText: ctx.DateLiteral().GetText(), Type: &types.DateType{Nullability: nullability}}
}

func (v *TestCaseVisitor) VisitTimeArg(ctx *baseparser.TimeArgContext) interface{} {
	nullability := getNullability(ctx.TimeType())
	value, err := literal.NewTimeFromString(getRawStringFromStringLiteral(ctx.TimeLiteral().GetText()), nullability == types.NullabilityNullable)
	if err != nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid time arg %v", err))
	}
	return &CaseLiteral{Value: value, ValueText: ctx.TimeLiteral().GetText(), Type: &types.TimeType{Nullability: nullability}}
}

func (v *TestCaseVisitor) VisitIntervalYearArg(ctx *baseparser.IntervalYearArgContext) interface{} {
	nullability := getNullability(ctx.IntervalYearType())
	interval := getRawStringFromStringLiteral(ctx.IntervalYearLiteral().GetText())
	value, err := literal.NewIntervalYearsToMonthFromString(interval, nullability == types.NullabilityNullable)
	if err != nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid interval year arg %v", err))
	}
	return &CaseLiteral{Value: value, ValueText: ctx.IntervalYearLiteral().GetText(), Type: &types.IntervalYearType{Nullability: nullability}}
}

func (v *TestCaseVisitor) VisitIntervalDayArg(ctx *baseparser.IntervalDayArgContext) interface{} {
	idayType := v.Visit(ctx.IntervalDayType()).(*types.IntervalDayType)
	interval := getRawStringFromStringLiteral(ctx.IntervalDayLiteral().GetText())
	value, err := literal.NewIntervalDaysToSecondFromString(interval, idayType.GetNullability() == types.NullabilityNullable)
	if err != nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid interval day arg %v", err))
	}
	return &CaseLiteral{Value: value, ValueText: ctx.IntervalDayLiteral().GetText(), Type: idayType}
}

func (v *TestCaseVisitor) VisitDecimalArg(ctx *baseparser.DecimalArgContext) interface{} {
	decType := v.Visit(ctx.DecimalType()).(types.Type)
	decimal, err := literal.NewDecimalFromString(ctx.NumericLiteral().GetText(), decType.GetNullability() == types.NullabilityNullable)
	if err != nil {
		v.ErrorListener.ReportVisitError(ctx, err)
	}
	return &CaseLiteral{Value: decimal, ValueText: ctx.NumericLiteral().GetText(), Type: decType}
}

func (v *TestCaseVisitor) VisitPrecisionTimeArg(ctx *baseparser.PrecisionTimeArgContext) interface{} {
	ptsType := v.Visit(ctx.PrecisionTimeType()).(*types.PrecisionTimeType)
	timestampStr := getRawStringFromStringLiteral(ctx.TimeLiteral().GetText())
	value, err := literal.NewPrecisionTimeFromString(ptsType.Precision, timestampStr, ptsType.GetNullability() == types.NullabilityNullable)
	if err != nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid precision time arg %v", err))
	}
	return &CaseLiteral{Value: value, ValueText: ctx.TimeLiteral().GetText(), Type: ptsType}
}

func (v *TestCaseVisitor) VisitPrecisionTimestampArg(ctx *baseparser.PrecisionTimestampArgContext) interface{} {
	ptsType := v.Visit(ctx.PrecisionTimestampType()).(*types.PrecisionTimestampType)
	timestampStr := getRawStringFromStringLiteral(ctx.TimestampLiteral().GetText())
	value, err := literal.NewPrecisionTimestampFromString(ptsType.Precision, timestampStr, ptsType.GetNullability() == types.NullabilityNullable)
	if err != nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid precision timestamp arg %v", err))
	}
	return &CaseLiteral{Value: value, ValueText: ctx.TimestampLiteral().GetText(), Type: ptsType}
}

func (v *TestCaseVisitor) VisitPrecisionTimestampTZArg(ctx *baseparser.PrecisionTimestampTZArgContext) interface{} {
	ptszType := v.Visit(ctx.PrecisionTimestampTZType()).(*types.PrecisionTimestampTzType)
	timestampStr := getRawStringFromStringLiteral(ctx.TimestampTzLiteral().GetText())
	value, err := literal.NewPrecisionTimestampTzFromString(ptszType.Precision, timestampStr, ptszType.GetNullability() == types.NullabilityNullable)
	if err != nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid precision timestampTZ arg %v", err))
	}
	return &CaseLiteral{Value: value, ValueText: ctx.TimestampTzLiteral().GetText(), Type: ptszType}
}

func (v *TestCaseVisitor) VisitListArg(ctx *baseparser.ListArgContext) interface{} {
	listType := v.Visit(ctx.ListType()).(*types.ListType)
	v.setLiteralTypeInContext(listType.Type)
	defer v.clearLiteralTypeInContext()

	values := v.Visit(ctx.LiteralList()).([]expr.Literal)

	var value expr.Literal
	var err error
	if len(values) == 0 {
		value = expr.NewEmptyListLiteral(listType.Type, false)
	} else {
		value, err = literal.NewList(values, false)
		if err != nil {
			v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid list arg %v", err))
		}
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
		return literal.NewBool(flag, false)
	}
	if ctx.DateLiteral() != nil {
		dateStr := getRawStringFromStringLiteral(ctx.DateLiteral().GetText())
		value, err := literal.NewDateFromString(dateStr, false)
		if err != nil {
			v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid date arg %v", err))
		}
		return value
	}
	if ctx.TimeLiteral() != nil {
		timeStr := getRawStringFromStringLiteral(ctx.TimeLiteral().GetText())
		value, err := literal.NewTimeFromString(timeStr, false)
		if err != nil {
			v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid time arg %v", err))
		}
		return value
	}
	if ctx.TimestampLiteral() != nil {
		timestampStr := getRawStringFromStringLiteral(ctx.TimestampLiteral().GetText())
		value, err := literal.NewTimestampFromString(timestampStr, false)
		if err != nil {
			v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid timestampTZ arg %v", err))
		}
		return value
	}
	if ctx.TimestampTzLiteral() != nil {
		timestampStr := getRawStringFromStringLiteral(ctx.TimestampTzLiteral().GetText())
		value, err := literal.NewTimestampTZFromString(timestampStr, false)
		if err != nil {
			v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid timestampTZ arg %v", err))
		}
		return value
	}
	if ctx.IntervalYearLiteral() != nil {
		iyearStr := getRawStringFromStringLiteral(ctx.IntervalYearLiteral().GetText())
		value, err := literal.NewIntervalYearsToMonthFromString(iyearStr, false)
		if err != nil {
			v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid interval year arg %v", err))
		}
		return value
	}
	if ctx.IntervalDayLiteral() != nil {
		idayStr := getRawStringFromStringLiteral(ctx.IntervalDayLiteral().GetText())
		value, err := literal.NewIntervalDaysToSecondFromString(idayStr, false)
		if err != nil {
			v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid interval day arg %v", err))
		}
		return value
	}

	if ctx.NumericLiteral() != nil {
		if v.getLiteralTypeInContext() == nil {
			// in compactAggregateFuncCall context, the type is not set, full schema of table may not be available
			return literal.NewString(ctx.NumericLiteral().GetText(), false)
		}
		value := v.getLiteralFromString(nil, ctx.NumericLiteral().GetText(), v.getLiteralTypeInContext())
		if value == nil {
			v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid numeric arg %v", ctx.GetText()))
		}
		return value
	}

	if ctx.StringLiteral() != nil {
		valueStr := getRawStringFromStringLiteral(ctx.GetText())
		value := literal.NewString(valueStr, false)
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
	v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid literal arg %v", ctx.GetText()))
	return nil
}

func (v *TestCaseVisitor) VisitResult(ctx *baseparser.ResultContext) interface{} {
	if ctx.SubstraitError() != nil {
		return v.Visit(ctx.SubstraitError())
	}
	result := v.Visit(ctx.Argument()).(*CaseLiteral)
	if err := result.updateLiteralType(); err != nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid result: %v", err))
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

func (v *TestCaseVisitor) VisitBoolean(ctx *baseparser.BooleanContext) interface{} {
	return &types.BooleanType{Nullability: getNullability(ctx.BooleanType())}
}

func (v *TestCaseVisitor) VisitInt(ctx *baseparser.IntContext) interface{} {
	return ctx.IntType().Accept(v)
}

func (v *TestCaseVisitor) VisitIntType(ctx *baseparser.IntTypeContext) interface{} {
	nullability := getNullability(ctx)
	if ctx.I8() != nil {
		return &types.Int8Type{Nullability: nullability}
	}
	if ctx.I16() != nil {
		return &types.Int16Type{Nullability: nullability}
	}
	if ctx.I32() != nil {
		return &types.Int32Type{Nullability: nullability}
	}
	if ctx.I64() != nil {
		return &types.Int64Type{Nullability: nullability}
	}
	v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid integer type %v", ctx.GetText()))
	return nil
}

func (v *TestCaseVisitor) VisitFloat(ctx *baseparser.FloatContext) interface{} {
	return ctx.FloatType().Accept(v)
}

func (v *TestCaseVisitor) VisitFloatType(ctx *baseparser.FloatTypeContext) interface{} {
	nullability := getNullability(ctx)
	if ctx.FP32() != nil {
		return &types.Float32Type{Nullability: nullability}
	} else if ctx.FP64() != nil {
		return &types.Float64Type{Nullability: nullability}
	}
	v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid float type %v", ctx.GetText()))
	return nil
}

func (v *TestCaseVisitor) VisitString(ctx *baseparser.StringContext) interface{} {
	return &types.StringType{Nullability: getNullability(ctx.StringType())}
}

func (v *TestCaseVisitor) VisitBinary(ctx *baseparser.BinaryContext) interface{} {
	return &types.BinaryType{Nullability: getNullability(ctx.BinaryType())}
}

func (v *TestCaseVisitor) VisitTimestamp(ctx *baseparser.TimestampContext) interface{} {
	return &types.TimestampType{Nullability: getNullability(ctx.TimestampType())}
}

func (v *TestCaseVisitor) VisitTimestampTz(ctx *baseparser.TimestampTzContext) interface{} {
	return &types.TimestampTzType{Nullability: getNullability(ctx.TimestampTZType())}
}

func (v *TestCaseVisitor) VisitDate(ctx *baseparser.DateContext) interface{} {
	return &types.DateType{Nullability: getNullability(ctx.DateType())}
}

func (v *TestCaseVisitor) VisitTime(ctx *baseparser.TimeContext) interface{} {
	return &types.TimeType{Nullability: getNullability(ctx.TimeType())}
}

func (v *TestCaseVisitor) VisitIntervalYear(ctx *baseparser.IntervalYearContext) interface{} {
	return &types.IntervalYearType{Nullability: getNullability(ctx.IntervalYearType())}
}

func (v *TestCaseVisitor) VisitUuid(ctx *baseparser.UuidContext) interface{} {
	return &types.UUIDType{Nullability: getNullability(ctx)}
}

func (v *TestCaseVisitor) VisitList(ctx *baseparser.ListContext) interface{} {
	elementType := v.Visit(ctx.GetElemType()).(types.Type)
	return &types.ListType{Type: elementType, Nullability: getNullability(ctx)}
}

func (v *TestCaseVisitor) VisitDataType(ctx *baseparser.DataTypeContext) interface{} {
	if ctx.ScalarType() != nil {
		return v.Visit(ctx.ScalarType())
	}
	if ctx.ParameterizedType() != nil {
		return v.Visit(ctx.ParameterizedType())
	}
	v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid data type %v", ctx.GetText()))
	return nil
}

func (v *TestCaseVisitor) VisitParameterizedType(ctx *baseparser.ParameterizedTypeContext) interface{} {
	if ctx.DecimalType() != nil {
		return v.Visit(ctx.DecimalType())
	}
	if ctx.PrecisionTimeType() != nil {
		return v.Visit(ctx.PrecisionTimeType())
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
	precision := int32(38)
	scale := int32(0)
	if ctx.GetPrecision() != nil {
		precision = v.Visit(ctx.GetPrecision()).(int32)
		scale = v.Visit(ctx.GetScale()).(int32)
	}
	return &types.DecimalType{Precision: precision, Scale: scale, Nullability: getNullability(ctx)}
}

func (v *TestCaseVisitor) VisitIntegerLiteral(ctx *baseparser.IntegerLiteralContext) interface{} {
	value, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("invalid int arg %v", err))
	}
	return int32(value)
}

func (v *TestCaseVisitor) VisitPrecisionTimeType(ctx *baseparser.PrecisionTimeTypeContext) interface{} {
	length := v.Visit(ctx.GetPrecision()).(int32)
	return &types.PrecisionTimeType{Precision: types.TimePrecision(length), Nullability: getNullability(ctx)}
}

func (v *TestCaseVisitor) VisitPrecisionTimestampType(ctx *baseparser.PrecisionTimestampTypeContext) interface{} {
	length := v.Visit(ctx.GetPrecision()).(int32)
	return &types.PrecisionTimestampType{Precision: types.TimePrecision(length), Nullability: getNullability(ctx)}
}

func (v *TestCaseVisitor) VisitPrecisionTimestampTZType(ctx *baseparser.PrecisionTimestampTZTypeContext) interface{} {
	length := v.Visit(ctx.GetPrecision()).(int32)
	return &types.PrecisionTimestampTzType{
		PrecisionTimestampType: types.PrecisionTimestampType{
			Precision:   types.TimePrecision(length),
			Nullability: getNullability(ctx),
		},
	}
}

func (v *TestCaseVisitor) VisitIntervalDayType(ctx *baseparser.IntervalDayTypeContext) interface{} {
	var length int32
	if ctx.GetLen_() != nil {
		length = v.Visit(ctx.GetLen_()).(int32)
	}
	return &types.IntervalDayType{Precision: types.TimePrecision(length), Nullability: getNullability(ctx)}
}

func (v *TestCaseVisitor) VisitFixedCharType(ctx *baseparser.FixedCharTypeContext) interface{} {
	length := v.Visit(ctx.GetLen_()).(int32)
	return &types.FixedCharType{Length: length, Nullability: getNullability(ctx)}
}

func (v *TestCaseVisitor) VisitVarCharType(ctx *baseparser.VarCharTypeContext) interface{} {
	length := v.Visit(ctx.GetLen_()).(int32)
	return &types.VarCharType{Length: length, Nullability: getNullability(ctx)}
}

func (v *TestCaseVisitor) VisitFixedBinaryType(ctx *baseparser.FixedBinaryTypeContext) interface{} {
	length := v.Visit(ctx.GetLen_()).(int32)
	return &types.FixedBinaryType{Length: length, Nullability: getNullability(ctx)}
}

func (v *TestCaseVisitor) VisitFuncType(ctx *baseparser.FuncTypeContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}

	var paramTypes []types.Type
	if ctx.GetParams() != nil {
		paramsResult := v.Visit(ctx.GetParams())
		if singleParam, ok := paramsResult.(types.Type); ok {
			paramTypes = []types.Type{singleParam}
		} else if paramSlice, ok := paramsResult.([]types.Type); ok {
			paramTypes = paramSlice
		} else {
			v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("unexpected func parameters type: %T", paramsResult))
			paramTypes = []types.Type{&types.Int32Type{}}
		}
	}

	// Parse return type
	if ctx.GetReturnType() == nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("func type missing return type"))
		return &types.FuncType{
			Nullability:    nullability,
			ParameterTypes: paramTypes,
			ReturnType:     &types.Int32Type{},
		}
	}

	returnTypeResult := v.Visit(ctx.GetReturnType())
	if returnTypeResult == nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("failed to parse func return type"))
		return &types.FuncType{
			Nullability:    nullability,
			ParameterTypes: paramTypes,
			ReturnType:     &types.Int32Type{},
		}
	}

	returnType, ok := returnTypeResult.(types.Type)
	if !ok {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("func return type is not a Type, got %T", returnTypeResult))
		return &types.FuncType{
			Nullability:    nullability,
			ParameterTypes: paramTypes,
			ReturnType:     &types.Int32Type{},
		}
	}

	return &types.FuncType{
		Nullability:    nullability,
		ParameterTypes: paramTypes,
		ReturnType:     returnType,
	}
}

// Lambda Expression Support

func (v *TestCaseVisitor) VisitLambdaArg(ctx *baseparser.LambdaArgContext) interface{} {
	if ctx.FuncType() == nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("lambda arg missing func type annotation"))
		return &CaseLiteral{ValueText: ctx.GetText()}
	}
	funcTypeResult := v.Visit(ctx.FuncType())
	if funcTypeResult == nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("failed to parse func type"))
		return &CaseLiteral{ValueText: ctx.GetText()}
	}
	funcType, ok := funcTypeResult.(types.Type)
	if !ok {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("func type is not a Type, got %T", funcTypeResult))
		return &CaseLiteral{ValueText: ctx.GetText()}
	}

	var paramTypes []types.Type
	if ft, ok := funcType.(*types.FuncType); ok {
		paramTypes = ft.ParameterTypes
	} else {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("lambda arg requires concrete FuncType, got %T", funcType))
		return &CaseLiteral{Type: funcType}
	}

	if ctx.LiteralLambda() == nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("lambda arg missing lambda expression"))
		return &CaseLiteral{Type: funcType, ValueText: ctx.GetText()}
	}

	lambdaCtx := ctx.LiteralLambda()
	if lambdaCtx.LambdaParameters() == nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("lambda missing parameters"))
		return &CaseLiteral{Type: funcType, ValueText: ctx.GetText()}
	}

	// Get parameter names
	paramInfoResult := v.Visit(lambdaCtx.LambdaParameters())
	if paramInfoResult == nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("failed to parse lambda parameters"))
		return &CaseLiteral{Type: funcType, ValueText: ctx.GetText()}
	}
	paramInfo, ok := paramInfoResult.([]lambdaParamInfo)
	if !ok {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("lambda parameters is not []lambdaParamInfo, got %T", paramInfoResult))
		return &CaseLiteral{Type: funcType, ValueText: ctx.GetText()}
	}

	// Check parameter count matches
	if len(paramInfo) != len(paramTypes) {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("lambda has %d parameters but type annotation expects %d",
			len(paramInfo), len(paramTypes)))
		return &CaseLiteral{Type: funcType, ValueText: ctx.GetText()}
	}

	// Set up parameter info with correct types BEFORE visiting the body
	for i := range paramInfo {
		paramInfo[i].typ = paramTypes[i]
	}

	// Save old params and set new ones
	oldParams := v.lambdaParams
	v.lambdaParams = paramInfo

	// Now visit the body with parameter types available
	if lambdaCtx.LambdaBody() == nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("lambda missing body"))
		v.lambdaParams = oldParams
		return &CaseLiteral{Type: funcType, ValueText: ctx.GetText()}
	}

	bodyResult := v.Visit(lambdaCtx.LambdaBody())
	if bodyResult == nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("failed to parse lambda body"))
		v.lambdaParams = oldParams
		return &CaseLiteral{Type: funcType, ValueText: ctx.GetText()}
	}

	body, ok := bodyResult.(expr.Expression)
	if !ok {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("lambda body is not an Expression, got %T", bodyResult))
		v.lambdaParams = oldParams
		return &CaseLiteral{Type: funcType, ValueText: ctx.GetText()}
	}

	// Restore old params
	v.lambdaParams = oldParams

	// Construct the lambda manually
	parameters := &types.StructType{
		Types:       paramTypes,
		Nullability: types.NullabilityRequired,
	}

	lambda := &expr.Lambda{
		Parameters: parameters,
		Body:       body,
	}

	return &CaseLiteral{
		Type:      funcType,
		Expr:      lambda,
		ValueText: ctx.GetText(),
	}
}

func (v *TestCaseVisitor) VisitLiteralLambda(ctx *baseparser.LiteralLambdaContext) interface{} {
	if ctx.LambdaParameters() == nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("lambda missing parameters"))
		return &expr.Lambda{
			Parameters: &types.StructType{Types: []types.Type{}, Nullability: types.NullabilityRequired},
			Body:       expr.NewNullLiteral(&types.Int32Type{Nullability: types.NullabilityNullable}),
		}
	}

	paramInfoResult := v.Visit(ctx.LambdaParameters())
	if paramInfoResult == nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("failed to parse lambda parameters"))
		return &expr.Lambda{
			Parameters: &types.StructType{Types: []types.Type{}, Nullability: types.NullabilityRequired},
			Body:       expr.NewNullLiteral(&types.Int32Type{Nullability: types.NullabilityNullable}),
		}
	}

	paramInfo, ok := paramInfoResult.([]lambdaParamInfo)
	if !ok {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("lambda parameters is not []lambdaParamInfo, got %T", paramInfoResult))
		return &expr.Lambda{
			Parameters: &types.StructType{Types: []types.Type{}, Nullability: types.NullabilityRequired},
			Body:       expr.NewNullLiteral(&types.Int32Type{Nullability: types.NullabilityNullable}),
		}
	}

	paramTypes := make([]types.Type, len(paramInfo))
	for i, info := range paramInfo {
		if info.typ != nil {
			paramTypes[i] = info.typ
		} else {
			paramTypes[i] = &types.Int32Type{Nullability: types.NullabilityRequired}
		}
	}

	parameters := &types.StructType{
		Types:       paramTypes,
		Nullability: types.NullabilityRequired,
	}

	oldParams := v.lambdaParams
	v.lambdaParams = paramInfo

	if ctx.LambdaBody() == nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("lambda missing body"))
		v.lambdaParams = oldParams
		return &expr.Lambda{
			Parameters: parameters,
			Body:       expr.NewNullLiteral(&types.Int32Type{Nullability: types.NullabilityNullable}),
		}
	}

	bodyResult := v.Visit(ctx.LambdaBody())
	if bodyResult == nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("failed to parse lambda body"))
		v.lambdaParams = oldParams
		return &expr.Lambda{
			Parameters: parameters,
			Body:       expr.NewNullLiteral(&types.Int32Type{Nullability: types.NullabilityNullable}),
		}
	}

	body, ok := bodyResult.(expr.Expression)
	if !ok {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("lambda body is not an Expression, got %T", bodyResult))
		v.lambdaParams = oldParams
		return &expr.Lambda{
			Parameters: parameters,
			Body:       expr.NewNullLiteral(&types.Int32Type{Nullability: types.NullabilityNullable}),
		}
	}

	v.lambdaParams = oldParams

	return &expr.Lambda{
		Parameters: parameters,
		Body:       body,
	}
}

// VisitSingleParam handles single lambda parameter: x
// Decision: Returns a slice with one element for consistency with VisitTupleParams
func (v *TestCaseVisitor) VisitSingleParam(ctx *baseparser.SingleParamContext) interface{} {
	if ctx.Identifier() == nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("single param missing identifier"))
		return []lambdaParamInfo{}
	}
	paramName := ctx.Identifier().GetText()
	// Note: We don't know the type yet - it will come from the func type annotation
	// We'll set it later in VisitLambdaArg
	return []lambdaParamInfo{
		{name: paramName, idx: 0},
	}
}

// VisitTupleParams handles multiple lambda parameters: (x, y)
// Decision: Returns a slice of parameter info, maintaining order
func (v *TestCaseVisitor) VisitTupleParams(ctx *baseparser.TupleParamsContext) interface{} {
	params := make([]lambdaParamInfo, 0, len(ctx.AllIdentifier()))
	for i, id := range ctx.AllIdentifier() {
		if id == nil {
			continue
		}
		params = append(params, lambdaParamInfo{
			name: id.GetText(),
			idx:  int32(i),
		})
	}
	return params
}

func (v *TestCaseVisitor) VisitSingleFuncParam(ctx *baseparser.SingleFuncParamContext) interface{} {
	if ctx.DataType() == nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("single func param missing data type"))
		return &types.Int32Type{}
	}
	return v.Visit(ctx.DataType())
}

func (v *TestCaseVisitor) VisitFuncParamsWithParens(ctx *baseparser.FuncParamsWithParensContext) interface{} {
	paramTypes := make([]types.Type, 0, len(ctx.AllDataType()))
	for _, dt := range ctx.AllDataType() {
		if dt == nil {
			continue
		}
		typeResult := v.Visit(dt)
		if typeResult == nil {
			continue
		}
		if t, ok := typeResult.(types.Type); ok {
			paramTypes = append(paramTypes, t)
		}
	}
	return paramTypes
}

func (v *TestCaseVisitor) VisitLambdaBody(ctx *baseparser.LambdaBodyContext) interface{} {
	if ctx.Identifier() == nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("lambda body missing function name"))
		return expr.NewNullLiteral(&types.Int32Type{Nullability: types.NullabilityNullable})
	}
	funcName := ctx.Identifier().GetText()

	var funcArgs []types.FuncArg
	if ctx.Arguments() != nil {
		argsResult := v.Visit(ctx.Arguments()).([]*CaseLiteral)
		funcArgs = make([]types.FuncArg, len(argsResult))
		for i, arg := range argsResult {
			if arg.Expr != nil {
				funcArgs[i] = arg.Expr
			} else if arg.Value != nil {
				funcArgs[i] = arg.Value
			} else {
				v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("lambda body argument %d has neither Expr nor Value", i))
				funcArgs[i] = expr.NewNullLiteral(&types.Int32Type{Nullability: types.NullabilityNullable})
			}
		}
	}

	collection := extensions.GetDefaultCollectionWithNoError()
	if collection == nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("failed to get default extension collection"))
		return expr.NewNullLiteral(&types.Int32Type{Nullability: types.NullabilityNullable})
	}
	extSet := extensions.NewSet()
	if extSet == nil {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("failed to create extension set"))
		return expr.NewNullLiteral(&types.Int32Type{Nullability: types.NullabilityNullable})
	}
	reg := expr.NewExtensionRegistry(extSet, collection)

	// Create a function registry to look up the function
	_, funcRegistry := functions.NewExtensionAndFunctionRegistries(collection)

	// Look up scalar functions with this name
	scalarFuncs := funcRegistry.GetScalarFunctions(funcName, len(funcArgs))
	if len(scalarFuncs) == 0 {
		// Function not found - return a null literal as a fallback
		return expr.NewNullLiteral(&types.Int32Type{Nullability: types.NullabilityNullable})
	}

	// Try each function variant to find one that matches
	for _, fn := range scalarFuncs {
		scalarFunc, err := expr.NewScalarFunc(reg, fn.ID(), nil, funcArgs...)
		if err == nil {
			return scalarFunc
		}
	}

	// None of the variants matched - return a null literal as a fallback
	return expr.NewNullLiteral(&types.Int32Type{Nullability: types.NullabilityNullable})
}

func (v *TestCaseVisitor) handleLambdaParameterRef(ctx antlr.ParserRuleContext, paramName string) *CaseLiteral {
	if len(v.lambdaParams) == 0 {
		v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("identifier %s found but not in lambda context", paramName))
		return &CaseLiteral{ValueText: paramName}
	}
	for i := len(v.lambdaParams) - 1; i >= 0; i-- {
		if v.lambdaParams[i].name == paramName {
			if v.lambdaParams[i].typ == nil {
				v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("lambda parameter %s has nil type", paramName))
				return &CaseLiteral{ValueText: paramName}
			}
			fieldRef, err := expr.NewRootFieldRef(
				expr.NewStructFieldRef(v.lambdaParams[i].idx),
				types.NewRecordTypeFromTypes([]types.Type{v.lambdaParams[i].typ}),
			)
			if err != nil {
				v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("error creating field ref for lambda param %s: %w", paramName, err))
				return &CaseLiteral{ValueText: paramName}
			}

			return &CaseLiteral{
				Expr:      fieldRef,
				ValueText: paramName,
				Type:      v.lambdaParams[i].typ,
			}
		}
	}

	v.ErrorListener.ReportVisitError(ctx, fmt.Errorf("unknown identifier in lambda context: %s", paramName))
	return &CaseLiteral{ValueText: paramName}
}
