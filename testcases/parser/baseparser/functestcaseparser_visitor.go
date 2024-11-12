// Code generated from FuncTestCaseParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package baseparser // FuncTestCaseParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by FuncTestCaseParser.
type FuncTestCaseParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by FuncTestCaseParser#doc.
	VisitDoc(ctx *DocContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#header.
	VisitHeader(ctx *HeaderContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#version.
	VisitVersion(ctx *VersionContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#include.
	VisitInclude(ctx *IncludeContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#testGroupDescription.
	VisitTestGroupDescription(ctx *TestGroupDescriptionContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#testCase.
	VisitTestCase(ctx *TestCaseContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#scalarFuncTestGroup.
	VisitScalarFuncTestGroup(ctx *ScalarFuncTestGroupContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#aggregateFuncTestGroup.
	VisitAggregateFuncTestGroup(ctx *AggregateFuncTestGroupContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#result.
	VisitResult(ctx *ResultContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#argument.
	VisitArgument(ctx *ArgumentContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#aggFuncTestCase.
	VisitAggFuncTestCase(ctx *AggFuncTestCaseContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#multiArgAggregateFuncCall.
	VisitMultiArgAggregateFuncCall(ctx *MultiArgAggregateFuncCallContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#compactAggregateFuncCall.
	VisitCompactAggregateFuncCall(ctx *CompactAggregateFuncCallContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#singleArgAggregateFuncCall.
	VisitSingleArgAggregateFuncCall(ctx *SingleArgAggregateFuncCallContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#tableData.
	VisitTableData(ctx *TableDataContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#tableRows.
	VisitTableRows(ctx *TableRowsContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#dataColumn.
	VisitDataColumn(ctx *DataColumnContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#columnValues.
	VisitColumnValues(ctx *ColumnValuesContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#qualifiedAggregateFuncArgs.
	VisitQualifiedAggregateFuncArgs(ctx *QualifiedAggregateFuncArgsContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#aggregateFuncArgs.
	VisitAggregateFuncArgs(ctx *AggregateFuncArgsContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#qualifiedAggregateFuncArg.
	VisitQualifiedAggregateFuncArg(ctx *QualifiedAggregateFuncArgContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#aggregateFuncArg.
	VisitAggregateFuncArg(ctx *AggregateFuncArgContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#numericLiteral.
	VisitNumericLiteral(ctx *NumericLiteralContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#floatLiteral.
	VisitFloatLiteral(ctx *FloatLiteralContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#nullArg.
	VisitNullArg(ctx *NullArgContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#intArg.
	VisitIntArg(ctx *IntArgContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#floatArg.
	VisitFloatArg(ctx *FloatArgContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#decimalArg.
	VisitDecimalArg(ctx *DecimalArgContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#booleanArg.
	VisitBooleanArg(ctx *BooleanArgContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#stringArg.
	VisitStringArg(ctx *StringArgContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#dateArg.
	VisitDateArg(ctx *DateArgContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#timeArg.
	VisitTimeArg(ctx *TimeArgContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#timestampArg.
	VisitTimestampArg(ctx *TimestampArgContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#timestampTzArg.
	VisitTimestampTzArg(ctx *TimestampTzArgContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#intervalYearArg.
	VisitIntervalYearArg(ctx *IntervalYearArgContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#intervalDayArg.
	VisitIntervalDayArg(ctx *IntervalDayArgContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#listArg.
	VisitListArg(ctx *ListArgContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#literalList.
	VisitLiteralList(ctx *LiteralListContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#intervalYearLiteral.
	VisitIntervalYearLiteral(ctx *IntervalYearLiteralContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#intervalDayLiteral.
	VisitIntervalDayLiteral(ctx *IntervalDayLiteralContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#timeInterval.
	VisitTimeInterval(ctx *TimeIntervalContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#dataType.
	VisitDataType(ctx *DataTypeContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#boolean.
	VisitBoolean(ctx *BooleanContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#i8.
	VisitI8(ctx *I8Context) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#i16.
	VisitI16(ctx *I16Context) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#i32.
	VisitI32(ctx *I32Context) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#i64.
	VisitI64(ctx *I64Context) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#fp32.
	VisitFp32(ctx *Fp32Context) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#fp64.
	VisitFp64(ctx *Fp64Context) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#string.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#binary.
	VisitBinary(ctx *BinaryContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#timestamp.
	VisitTimestamp(ctx *TimestampContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#timestampTz.
	VisitTimestampTz(ctx *TimestampTzContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#date.
	VisitDate(ctx *DateContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#time.
	VisitTime(ctx *TimeContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#intervalDay.
	VisitIntervalDay(ctx *IntervalDayContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#intervalYear.
	VisitIntervalYear(ctx *IntervalYearContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#uuid.
	VisitUuid(ctx *UuidContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#userDefined.
	VisitUserDefined(ctx *UserDefinedContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#booleanType.
	VisitBooleanType(ctx *BooleanTypeContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#stringType.
	VisitStringType(ctx *StringTypeContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#binaryType.
	VisitBinaryType(ctx *BinaryTypeContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#timestampType.
	VisitTimestampType(ctx *TimestampTypeContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#timestampTZType.
	VisitTimestampTZType(ctx *TimestampTZTypeContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#intervalYearType.
	VisitIntervalYearType(ctx *IntervalYearTypeContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#intervalDayType.
	VisitIntervalDayType(ctx *IntervalDayTypeContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#fixedChar.
	VisitFixedChar(ctx *FixedCharContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#varChar.
	VisitVarChar(ctx *VarCharContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#fixedBinary.
	VisitFixedBinary(ctx *FixedBinaryContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#decimal.
	VisitDecimal(ctx *DecimalContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#precisionTimestamp.
	VisitPrecisionTimestamp(ctx *PrecisionTimestampContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#precisionTimestampTZ.
	VisitPrecisionTimestampTZ(ctx *PrecisionTimestampTZContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#list.
	VisitList(ctx *ListContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#parameterizedType.
	VisitParameterizedType(ctx *ParameterizedTypeContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#integerLiteral.
	VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#substraitError.
	VisitSubstraitError(ctx *SubstraitErrorContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#func_option.
	VisitFunc_option(ctx *Func_optionContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#option_name.
	VisitOption_name(ctx *Option_nameContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#option_value.
	VisitOption_value(ctx *Option_valueContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#func_options.
	VisitFunc_options(ctx *Func_optionsContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#nonReserved.
	VisitNonReserved(ctx *NonReservedContext) interface{}

	// Visit a parse tree produced by FuncTestCaseParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}
}
