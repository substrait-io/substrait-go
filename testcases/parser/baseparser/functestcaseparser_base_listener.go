// Code generated from FuncTestCaseParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package baseparser // FuncTestCaseParser
import "github.com/antlr4-go/antlr/v4"

// BaseFuncTestCaseParserListener is a complete listener for a parse tree produced by FuncTestCaseParser.
type BaseFuncTestCaseParserListener struct{}

var _ FuncTestCaseParserListener = &BaseFuncTestCaseParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFuncTestCaseParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFuncTestCaseParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFuncTestCaseParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFuncTestCaseParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDoc is called when production doc is entered.
func (s *BaseFuncTestCaseParserListener) EnterDoc(ctx *DocContext) {}

// ExitDoc is called when production doc is exited.
func (s *BaseFuncTestCaseParserListener) ExitDoc(ctx *DocContext) {}

// EnterHeader is called when production header is entered.
func (s *BaseFuncTestCaseParserListener) EnterHeader(ctx *HeaderContext) {}

// ExitHeader is called when production header is exited.
func (s *BaseFuncTestCaseParserListener) ExitHeader(ctx *HeaderContext) {}

// EnterVersion is called when production version is entered.
func (s *BaseFuncTestCaseParserListener) EnterVersion(ctx *VersionContext) {}

// ExitVersion is called when production version is exited.
func (s *BaseFuncTestCaseParserListener) ExitVersion(ctx *VersionContext) {}

// EnterInclude is called when production include is entered.
func (s *BaseFuncTestCaseParserListener) EnterInclude(ctx *IncludeContext) {}

// ExitInclude is called when production include is exited.
func (s *BaseFuncTestCaseParserListener) ExitInclude(ctx *IncludeContext) {}

// EnterTestGroupDescription is called when production testGroupDescription is entered.
func (s *BaseFuncTestCaseParserListener) EnterTestGroupDescription(ctx *TestGroupDescriptionContext) {
}

// ExitTestGroupDescription is called when production testGroupDescription is exited.
func (s *BaseFuncTestCaseParserListener) ExitTestGroupDescription(ctx *TestGroupDescriptionContext) {}

// EnterTestCase is called when production testCase is entered.
func (s *BaseFuncTestCaseParserListener) EnterTestCase(ctx *TestCaseContext) {}

// ExitTestCase is called when production testCase is exited.
func (s *BaseFuncTestCaseParserListener) ExitTestCase(ctx *TestCaseContext) {}

// EnterScalarFuncTestGroup is called when production scalarFuncTestGroup is entered.
func (s *BaseFuncTestCaseParserListener) EnterScalarFuncTestGroup(ctx *ScalarFuncTestGroupContext) {}

// ExitScalarFuncTestGroup is called when production scalarFuncTestGroup is exited.
func (s *BaseFuncTestCaseParserListener) ExitScalarFuncTestGroup(ctx *ScalarFuncTestGroupContext) {}

// EnterAggregateFuncTestGroup is called when production aggregateFuncTestGroup is entered.
func (s *BaseFuncTestCaseParserListener) EnterAggregateFuncTestGroup(ctx *AggregateFuncTestGroupContext) {
}

// ExitAggregateFuncTestGroup is called when production aggregateFuncTestGroup is exited.
func (s *BaseFuncTestCaseParserListener) ExitAggregateFuncTestGroup(ctx *AggregateFuncTestGroupContext) {
}

// EnterArguments is called when production arguments is entered.
func (s *BaseFuncTestCaseParserListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseFuncTestCaseParserListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterResult is called when production result is entered.
func (s *BaseFuncTestCaseParserListener) EnterResult(ctx *ResultContext) {}

// ExitResult is called when production result is exited.
func (s *BaseFuncTestCaseParserListener) ExitResult(ctx *ResultContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseFuncTestCaseParserListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseFuncTestCaseParserListener) ExitArgument(ctx *ArgumentContext) {}

// EnterAggFuncTestCase is called when production aggFuncTestCase is entered.
func (s *BaseFuncTestCaseParserListener) EnterAggFuncTestCase(ctx *AggFuncTestCaseContext) {}

// ExitAggFuncTestCase is called when production aggFuncTestCase is exited.
func (s *BaseFuncTestCaseParserListener) ExitAggFuncTestCase(ctx *AggFuncTestCaseContext) {}

// EnterMultiArgAggregateFuncCall is called when production multiArgAggregateFuncCall is entered.
func (s *BaseFuncTestCaseParserListener) EnterMultiArgAggregateFuncCall(ctx *MultiArgAggregateFuncCallContext) {
}

// ExitMultiArgAggregateFuncCall is called when production multiArgAggregateFuncCall is exited.
func (s *BaseFuncTestCaseParserListener) ExitMultiArgAggregateFuncCall(ctx *MultiArgAggregateFuncCallContext) {
}

// EnterCompactAggregateFuncCall is called when production compactAggregateFuncCall is entered.
func (s *BaseFuncTestCaseParserListener) EnterCompactAggregateFuncCall(ctx *CompactAggregateFuncCallContext) {
}

// ExitCompactAggregateFuncCall is called when production compactAggregateFuncCall is exited.
func (s *BaseFuncTestCaseParserListener) ExitCompactAggregateFuncCall(ctx *CompactAggregateFuncCallContext) {
}

// EnterSingleArgAggregateFuncCall is called when production singleArgAggregateFuncCall is entered.
func (s *BaseFuncTestCaseParserListener) EnterSingleArgAggregateFuncCall(ctx *SingleArgAggregateFuncCallContext) {
}

// ExitSingleArgAggregateFuncCall is called when production singleArgAggregateFuncCall is exited.
func (s *BaseFuncTestCaseParserListener) ExitSingleArgAggregateFuncCall(ctx *SingleArgAggregateFuncCallContext) {
}

// EnterTableData is called when production tableData is entered.
func (s *BaseFuncTestCaseParserListener) EnterTableData(ctx *TableDataContext) {}

// ExitTableData is called when production tableData is exited.
func (s *BaseFuncTestCaseParserListener) ExitTableData(ctx *TableDataContext) {}

// EnterTableRows is called when production tableRows is entered.
func (s *BaseFuncTestCaseParserListener) EnterTableRows(ctx *TableRowsContext) {}

// ExitTableRows is called when production tableRows is exited.
func (s *BaseFuncTestCaseParserListener) ExitTableRows(ctx *TableRowsContext) {}

// EnterDataColumn is called when production dataColumn is entered.
func (s *BaseFuncTestCaseParserListener) EnterDataColumn(ctx *DataColumnContext) {}

// ExitDataColumn is called when production dataColumn is exited.
func (s *BaseFuncTestCaseParserListener) ExitDataColumn(ctx *DataColumnContext) {}

// EnterColumnValues is called when production columnValues is entered.
func (s *BaseFuncTestCaseParserListener) EnterColumnValues(ctx *ColumnValuesContext) {}

// ExitColumnValues is called when production columnValues is exited.
func (s *BaseFuncTestCaseParserListener) ExitColumnValues(ctx *ColumnValuesContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseFuncTestCaseParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseFuncTestCaseParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterQualifiedAggregateFuncArgs is called when production qualifiedAggregateFuncArgs is entered.
func (s *BaseFuncTestCaseParserListener) EnterQualifiedAggregateFuncArgs(ctx *QualifiedAggregateFuncArgsContext) {
}

// ExitQualifiedAggregateFuncArgs is called when production qualifiedAggregateFuncArgs is exited.
func (s *BaseFuncTestCaseParserListener) ExitQualifiedAggregateFuncArgs(ctx *QualifiedAggregateFuncArgsContext) {
}

// EnterAggregateFuncArgs is called when production aggregateFuncArgs is entered.
func (s *BaseFuncTestCaseParserListener) EnterAggregateFuncArgs(ctx *AggregateFuncArgsContext) {}

// ExitAggregateFuncArgs is called when production aggregateFuncArgs is exited.
func (s *BaseFuncTestCaseParserListener) ExitAggregateFuncArgs(ctx *AggregateFuncArgsContext) {}

// EnterQualifiedAggregateFuncArg is called when production qualifiedAggregateFuncArg is entered.
func (s *BaseFuncTestCaseParserListener) EnterQualifiedAggregateFuncArg(ctx *QualifiedAggregateFuncArgContext) {
}

// ExitQualifiedAggregateFuncArg is called when production qualifiedAggregateFuncArg is exited.
func (s *BaseFuncTestCaseParserListener) ExitQualifiedAggregateFuncArg(ctx *QualifiedAggregateFuncArgContext) {
}

// EnterAggregateFuncArg is called when production aggregateFuncArg is entered.
func (s *BaseFuncTestCaseParserListener) EnterAggregateFuncArg(ctx *AggregateFuncArgContext) {}

// ExitAggregateFuncArg is called when production aggregateFuncArg is exited.
func (s *BaseFuncTestCaseParserListener) ExitAggregateFuncArg(ctx *AggregateFuncArgContext) {}

// EnterNumericLiteral is called when production numericLiteral is entered.
func (s *BaseFuncTestCaseParserListener) EnterNumericLiteral(ctx *NumericLiteralContext) {}

// ExitNumericLiteral is called when production numericLiteral is exited.
func (s *BaseFuncTestCaseParserListener) ExitNumericLiteral(ctx *NumericLiteralContext) {}

// EnterFloatLiteral is called when production floatLiteral is entered.
func (s *BaseFuncTestCaseParserListener) EnterFloatLiteral(ctx *FloatLiteralContext) {}

// ExitFloatLiteral is called when production floatLiteral is exited.
func (s *BaseFuncTestCaseParserListener) ExitFloatLiteral(ctx *FloatLiteralContext) {}

// EnterNullArg is called when production nullArg is entered.
func (s *BaseFuncTestCaseParserListener) EnterNullArg(ctx *NullArgContext) {}

// ExitNullArg is called when production nullArg is exited.
func (s *BaseFuncTestCaseParserListener) ExitNullArg(ctx *NullArgContext) {}

// EnterIntArg is called when production intArg is entered.
func (s *BaseFuncTestCaseParserListener) EnterIntArg(ctx *IntArgContext) {}

// ExitIntArg is called when production intArg is exited.
func (s *BaseFuncTestCaseParserListener) ExitIntArg(ctx *IntArgContext) {}

// EnterFloatArg is called when production floatArg is entered.
func (s *BaseFuncTestCaseParserListener) EnterFloatArg(ctx *FloatArgContext) {}

// ExitFloatArg is called when production floatArg is exited.
func (s *BaseFuncTestCaseParserListener) ExitFloatArg(ctx *FloatArgContext) {}

// EnterDecimalArg is called when production decimalArg is entered.
func (s *BaseFuncTestCaseParserListener) EnterDecimalArg(ctx *DecimalArgContext) {}

// ExitDecimalArg is called when production decimalArg is exited.
func (s *BaseFuncTestCaseParserListener) ExitDecimalArg(ctx *DecimalArgContext) {}

// EnterBooleanArg is called when production booleanArg is entered.
func (s *BaseFuncTestCaseParserListener) EnterBooleanArg(ctx *BooleanArgContext) {}

// ExitBooleanArg is called when production booleanArg is exited.
func (s *BaseFuncTestCaseParserListener) ExitBooleanArg(ctx *BooleanArgContext) {}

// EnterStringArg is called when production stringArg is entered.
func (s *BaseFuncTestCaseParserListener) EnterStringArg(ctx *StringArgContext) {}

// ExitStringArg is called when production stringArg is exited.
func (s *BaseFuncTestCaseParserListener) ExitStringArg(ctx *StringArgContext) {}

// EnterDateArg is called when production dateArg is entered.
func (s *BaseFuncTestCaseParserListener) EnterDateArg(ctx *DateArgContext) {}

// ExitDateArg is called when production dateArg is exited.
func (s *BaseFuncTestCaseParserListener) ExitDateArg(ctx *DateArgContext) {}

// EnterTimeArg is called when production timeArg is entered.
func (s *BaseFuncTestCaseParserListener) EnterTimeArg(ctx *TimeArgContext) {}

// ExitTimeArg is called when production timeArg is exited.
func (s *BaseFuncTestCaseParserListener) ExitTimeArg(ctx *TimeArgContext) {}

// EnterTimestampArg is called when production timestampArg is entered.
func (s *BaseFuncTestCaseParserListener) EnterTimestampArg(ctx *TimestampArgContext) {}

// ExitTimestampArg is called when production timestampArg is exited.
func (s *BaseFuncTestCaseParserListener) ExitTimestampArg(ctx *TimestampArgContext) {}

// EnterTimestampTzArg is called when production timestampTzArg is entered.
func (s *BaseFuncTestCaseParserListener) EnterTimestampTzArg(ctx *TimestampTzArgContext) {}

// ExitTimestampTzArg is called when production timestampTzArg is exited.
func (s *BaseFuncTestCaseParserListener) ExitTimestampTzArg(ctx *TimestampTzArgContext) {}

// EnterIntervalYearArg is called when production intervalYearArg is entered.
func (s *BaseFuncTestCaseParserListener) EnterIntervalYearArg(ctx *IntervalYearArgContext) {}

// ExitIntervalYearArg is called when production intervalYearArg is exited.
func (s *BaseFuncTestCaseParserListener) ExitIntervalYearArg(ctx *IntervalYearArgContext) {}

// EnterIntervalDayArg is called when production intervalDayArg is entered.
func (s *BaseFuncTestCaseParserListener) EnterIntervalDayArg(ctx *IntervalDayArgContext) {}

// ExitIntervalDayArg is called when production intervalDayArg is exited.
func (s *BaseFuncTestCaseParserListener) ExitIntervalDayArg(ctx *IntervalDayArgContext) {}

// EnterListArg is called when production listArg is entered.
func (s *BaseFuncTestCaseParserListener) EnterListArg(ctx *ListArgContext) {}

// ExitListArg is called when production listArg is exited.
func (s *BaseFuncTestCaseParserListener) ExitListArg(ctx *ListArgContext) {}

// EnterLiteralList is called when production literalList is entered.
func (s *BaseFuncTestCaseParserListener) EnterLiteralList(ctx *LiteralListContext) {}

// ExitLiteralList is called when production literalList is exited.
func (s *BaseFuncTestCaseParserListener) ExitLiteralList(ctx *LiteralListContext) {}

// EnterIntervalYearLiteral is called when production intervalYearLiteral is entered.
func (s *BaseFuncTestCaseParserListener) EnterIntervalYearLiteral(ctx *IntervalYearLiteralContext) {}

// ExitIntervalYearLiteral is called when production intervalYearLiteral is exited.
func (s *BaseFuncTestCaseParserListener) ExitIntervalYearLiteral(ctx *IntervalYearLiteralContext) {}

// EnterIntervalDayLiteral is called when production intervalDayLiteral is entered.
func (s *BaseFuncTestCaseParserListener) EnterIntervalDayLiteral(ctx *IntervalDayLiteralContext) {}

// ExitIntervalDayLiteral is called when production intervalDayLiteral is exited.
func (s *BaseFuncTestCaseParserListener) ExitIntervalDayLiteral(ctx *IntervalDayLiteralContext) {}

// EnterTimeInterval is called when production timeInterval is entered.
func (s *BaseFuncTestCaseParserListener) EnterTimeInterval(ctx *TimeIntervalContext) {}

// ExitTimeInterval is called when production timeInterval is exited.
func (s *BaseFuncTestCaseParserListener) ExitTimeInterval(ctx *TimeIntervalContext) {}

// EnterDataType is called when production dataType is entered.
func (s *BaseFuncTestCaseParserListener) EnterDataType(ctx *DataTypeContext) {}

// ExitDataType is called when production dataType is exited.
func (s *BaseFuncTestCaseParserListener) ExitDataType(ctx *DataTypeContext) {}

// EnterBoolean is called when production boolean is entered.
func (s *BaseFuncTestCaseParserListener) EnterBoolean(ctx *BooleanContext) {}

// ExitBoolean is called when production boolean is exited.
func (s *BaseFuncTestCaseParserListener) ExitBoolean(ctx *BooleanContext) {}

// EnterI8 is called when production i8 is entered.
func (s *BaseFuncTestCaseParserListener) EnterI8(ctx *I8Context) {}

// ExitI8 is called when production i8 is exited.
func (s *BaseFuncTestCaseParserListener) ExitI8(ctx *I8Context) {}

// EnterI16 is called when production i16 is entered.
func (s *BaseFuncTestCaseParserListener) EnterI16(ctx *I16Context) {}

// ExitI16 is called when production i16 is exited.
func (s *BaseFuncTestCaseParserListener) ExitI16(ctx *I16Context) {}

// EnterI32 is called when production i32 is entered.
func (s *BaseFuncTestCaseParserListener) EnterI32(ctx *I32Context) {}

// ExitI32 is called when production i32 is exited.
func (s *BaseFuncTestCaseParserListener) ExitI32(ctx *I32Context) {}

// EnterI64 is called when production i64 is entered.
func (s *BaseFuncTestCaseParserListener) EnterI64(ctx *I64Context) {}

// ExitI64 is called when production i64 is exited.
func (s *BaseFuncTestCaseParserListener) ExitI64(ctx *I64Context) {}

// EnterFp32 is called when production fp32 is entered.
func (s *BaseFuncTestCaseParserListener) EnterFp32(ctx *Fp32Context) {}

// ExitFp32 is called when production fp32 is exited.
func (s *BaseFuncTestCaseParserListener) ExitFp32(ctx *Fp32Context) {}

// EnterFp64 is called when production fp64 is entered.
func (s *BaseFuncTestCaseParserListener) EnterFp64(ctx *Fp64Context) {}

// ExitFp64 is called when production fp64 is exited.
func (s *BaseFuncTestCaseParserListener) ExitFp64(ctx *Fp64Context) {}

// EnterString is called when production string is entered.
func (s *BaseFuncTestCaseParserListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseFuncTestCaseParserListener) ExitString(ctx *StringContext) {}

// EnterBinary is called when production binary is entered.
func (s *BaseFuncTestCaseParserListener) EnterBinary(ctx *BinaryContext) {}

// ExitBinary is called when production binary is exited.
func (s *BaseFuncTestCaseParserListener) ExitBinary(ctx *BinaryContext) {}

// EnterTimestamp is called when production timestamp is entered.
func (s *BaseFuncTestCaseParserListener) EnterTimestamp(ctx *TimestampContext) {}

// ExitTimestamp is called when production timestamp is exited.
func (s *BaseFuncTestCaseParserListener) ExitTimestamp(ctx *TimestampContext) {}

// EnterTimestampTz is called when production timestampTz is entered.
func (s *BaseFuncTestCaseParserListener) EnterTimestampTz(ctx *TimestampTzContext) {}

// ExitTimestampTz is called when production timestampTz is exited.
func (s *BaseFuncTestCaseParserListener) ExitTimestampTz(ctx *TimestampTzContext) {}

// EnterDate is called when production date is entered.
func (s *BaseFuncTestCaseParserListener) EnterDate(ctx *DateContext) {}

// ExitDate is called when production date is exited.
func (s *BaseFuncTestCaseParserListener) ExitDate(ctx *DateContext) {}

// EnterTime is called when production time is entered.
func (s *BaseFuncTestCaseParserListener) EnterTime(ctx *TimeContext) {}

// ExitTime is called when production time is exited.
func (s *BaseFuncTestCaseParserListener) ExitTime(ctx *TimeContext) {}

// EnterIntervalDay is called when production intervalDay is entered.
func (s *BaseFuncTestCaseParserListener) EnterIntervalDay(ctx *IntervalDayContext) {}

// ExitIntervalDay is called when production intervalDay is exited.
func (s *BaseFuncTestCaseParserListener) ExitIntervalDay(ctx *IntervalDayContext) {}

// EnterIntervalYear is called when production intervalYear is entered.
func (s *BaseFuncTestCaseParserListener) EnterIntervalYear(ctx *IntervalYearContext) {}

// ExitIntervalYear is called when production intervalYear is exited.
func (s *BaseFuncTestCaseParserListener) ExitIntervalYear(ctx *IntervalYearContext) {}

// EnterUuid is called when production uuid is entered.
func (s *BaseFuncTestCaseParserListener) EnterUuid(ctx *UuidContext) {}

// ExitUuid is called when production uuid is exited.
func (s *BaseFuncTestCaseParserListener) ExitUuid(ctx *UuidContext) {}

// EnterUserDefined is called when production userDefined is entered.
func (s *BaseFuncTestCaseParserListener) EnterUserDefined(ctx *UserDefinedContext) {}

// ExitUserDefined is called when production userDefined is exited.
func (s *BaseFuncTestCaseParserListener) ExitUserDefined(ctx *UserDefinedContext) {}

// EnterBooleanType is called when production booleanType is entered.
func (s *BaseFuncTestCaseParserListener) EnterBooleanType(ctx *BooleanTypeContext) {}

// ExitBooleanType is called when production booleanType is exited.
func (s *BaseFuncTestCaseParserListener) ExitBooleanType(ctx *BooleanTypeContext) {}

// EnterStringType is called when production stringType is entered.
func (s *BaseFuncTestCaseParserListener) EnterStringType(ctx *StringTypeContext) {}

// ExitStringType is called when production stringType is exited.
func (s *BaseFuncTestCaseParserListener) ExitStringType(ctx *StringTypeContext) {}

// EnterBinaryType is called when production binaryType is entered.
func (s *BaseFuncTestCaseParserListener) EnterBinaryType(ctx *BinaryTypeContext) {}

// ExitBinaryType is called when production binaryType is exited.
func (s *BaseFuncTestCaseParserListener) ExitBinaryType(ctx *BinaryTypeContext) {}

// EnterTimestampType is called when production timestampType is entered.
func (s *BaseFuncTestCaseParserListener) EnterTimestampType(ctx *TimestampTypeContext) {}

// ExitTimestampType is called when production timestampType is exited.
func (s *BaseFuncTestCaseParserListener) ExitTimestampType(ctx *TimestampTypeContext) {}

// EnterTimestampTZType is called when production timestampTZType is entered.
func (s *BaseFuncTestCaseParserListener) EnterTimestampTZType(ctx *TimestampTZTypeContext) {}

// ExitTimestampTZType is called when production timestampTZType is exited.
func (s *BaseFuncTestCaseParserListener) ExitTimestampTZType(ctx *TimestampTZTypeContext) {}

// EnterIntervalYearType is called when production intervalYearType is entered.
func (s *BaseFuncTestCaseParserListener) EnterIntervalYearType(ctx *IntervalYearTypeContext) {}

// ExitIntervalYearType is called when production intervalYearType is exited.
func (s *BaseFuncTestCaseParserListener) ExitIntervalYearType(ctx *IntervalYearTypeContext) {}

// EnterIntervalDayType is called when production intervalDayType is entered.
func (s *BaseFuncTestCaseParserListener) EnterIntervalDayType(ctx *IntervalDayTypeContext) {}

// ExitIntervalDayType is called when production intervalDayType is exited.
func (s *BaseFuncTestCaseParserListener) ExitIntervalDayType(ctx *IntervalDayTypeContext) {}

// EnterFixedChar is called when production fixedChar is entered.
func (s *BaseFuncTestCaseParserListener) EnterFixedChar(ctx *FixedCharContext) {}

// ExitFixedChar is called when production fixedChar is exited.
func (s *BaseFuncTestCaseParserListener) ExitFixedChar(ctx *FixedCharContext) {}

// EnterVarChar is called when production varChar is entered.
func (s *BaseFuncTestCaseParserListener) EnterVarChar(ctx *VarCharContext) {}

// ExitVarChar is called when production varChar is exited.
func (s *BaseFuncTestCaseParserListener) ExitVarChar(ctx *VarCharContext) {}

// EnterFixedBinary is called when production fixedBinary is entered.
func (s *BaseFuncTestCaseParserListener) EnterFixedBinary(ctx *FixedBinaryContext) {}

// ExitFixedBinary is called when production fixedBinary is exited.
func (s *BaseFuncTestCaseParserListener) ExitFixedBinary(ctx *FixedBinaryContext) {}

// EnterDecimal is called when production decimal is entered.
func (s *BaseFuncTestCaseParserListener) EnterDecimal(ctx *DecimalContext) {}

// ExitDecimal is called when production decimal is exited.
func (s *BaseFuncTestCaseParserListener) ExitDecimal(ctx *DecimalContext) {}

// EnterPrecisionTimestamp is called when production precisionTimestamp is entered.
func (s *BaseFuncTestCaseParserListener) EnterPrecisionTimestamp(ctx *PrecisionTimestampContext) {}

// ExitPrecisionTimestamp is called when production precisionTimestamp is exited.
func (s *BaseFuncTestCaseParserListener) ExitPrecisionTimestamp(ctx *PrecisionTimestampContext) {}

// EnterPrecisionTimestampTZ is called when production precisionTimestampTZ is entered.
func (s *BaseFuncTestCaseParserListener) EnterPrecisionTimestampTZ(ctx *PrecisionTimestampTZContext) {
}

// ExitPrecisionTimestampTZ is called when production precisionTimestampTZ is exited.
func (s *BaseFuncTestCaseParserListener) ExitPrecisionTimestampTZ(ctx *PrecisionTimestampTZContext) {}

// EnterList is called when production list is entered.
func (s *BaseFuncTestCaseParserListener) EnterList(ctx *ListContext) {}

// ExitList is called when production list is exited.
func (s *BaseFuncTestCaseParserListener) ExitList(ctx *ListContext) {}

// EnterParameterizedType is called when production parameterizedType is entered.
func (s *BaseFuncTestCaseParserListener) EnterParameterizedType(ctx *ParameterizedTypeContext) {}

// ExitParameterizedType is called when production parameterizedType is exited.
func (s *BaseFuncTestCaseParserListener) ExitParameterizedType(ctx *ParameterizedTypeContext) {}

// EnterIntegerLiteral is called when production integerLiteral is entered.
func (s *BaseFuncTestCaseParserListener) EnterIntegerLiteral(ctx *IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production integerLiteral is exited.
func (s *BaseFuncTestCaseParserListener) ExitIntegerLiteral(ctx *IntegerLiteralContext) {}

// EnterSubstraitError is called when production substraitError is entered.
func (s *BaseFuncTestCaseParserListener) EnterSubstraitError(ctx *SubstraitErrorContext) {}

// ExitSubstraitError is called when production substraitError is exited.
func (s *BaseFuncTestCaseParserListener) ExitSubstraitError(ctx *SubstraitErrorContext) {}

// EnterFunc_option is called when production func_option is entered.
func (s *BaseFuncTestCaseParserListener) EnterFunc_option(ctx *Func_optionContext) {}

// ExitFunc_option is called when production func_option is exited.
func (s *BaseFuncTestCaseParserListener) ExitFunc_option(ctx *Func_optionContext) {}

// EnterOption_name is called when production option_name is entered.
func (s *BaseFuncTestCaseParserListener) EnterOption_name(ctx *Option_nameContext) {}

// ExitOption_name is called when production option_name is exited.
func (s *BaseFuncTestCaseParserListener) ExitOption_name(ctx *Option_nameContext) {}

// EnterOption_value is called when production option_value is entered.
func (s *BaseFuncTestCaseParserListener) EnterOption_value(ctx *Option_valueContext) {}

// ExitOption_value is called when production option_value is exited.
func (s *BaseFuncTestCaseParserListener) ExitOption_value(ctx *Option_valueContext) {}

// EnterFunc_options is called when production func_options is entered.
func (s *BaseFuncTestCaseParserListener) EnterFunc_options(ctx *Func_optionsContext) {}

// ExitFunc_options is called when production func_options is exited.
func (s *BaseFuncTestCaseParserListener) ExitFunc_options(ctx *Func_optionsContext) {}

// EnterNonReserved is called when production nonReserved is entered.
func (s *BaseFuncTestCaseParserListener) EnterNonReserved(ctx *NonReservedContext) {}

// ExitNonReserved is called when production nonReserved is exited.
func (s *BaseFuncTestCaseParserListener) ExitNonReserved(ctx *NonReservedContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseFuncTestCaseParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseFuncTestCaseParserListener) ExitIdentifier(ctx *IdentifierContext) {}
