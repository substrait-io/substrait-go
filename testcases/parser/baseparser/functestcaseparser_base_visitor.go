// Code generated from FuncTestCaseParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package baseparser // FuncTestCaseParser
import "github.com/antlr4-go/antlr/v4"

type BaseFuncTestCaseParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseFuncTestCaseParserVisitor) VisitDoc(ctx *DocContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitHeader(ctx *HeaderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitVersion(ctx *VersionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitInclude(ctx *IncludeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitTestGroupDescription(ctx *TestGroupDescriptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitTestCase(ctx *TestCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitScalarFuncTestGroup(ctx *ScalarFuncTestGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitAggregateFuncTestGroup(ctx *AggregateFuncTestGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitResult(ctx *ResultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitArgument(ctx *ArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitAggFuncTestCase(ctx *AggFuncTestCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitMultiArgAggregateFuncCall(ctx *MultiArgAggregateFuncCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitCompactAggregateFuncCall(ctx *CompactAggregateFuncCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitSingleArgAggregateFuncCall(ctx *SingleArgAggregateFuncCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitTableData(ctx *TableDataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitTableRows(ctx *TableRowsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitDataColumn(ctx *DataColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitColumnValues(ctx *ColumnValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitQualifiedAggregateFuncArgs(ctx *QualifiedAggregateFuncArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitAggregateFuncArgs(ctx *AggregateFuncArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitQualifiedAggregateFuncArg(ctx *QualifiedAggregateFuncArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitAggregateFuncArg(ctx *AggregateFuncArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitNumericLiteral(ctx *NumericLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitFloatLiteral(ctx *FloatLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitNullArg(ctx *NullArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitIntArg(ctx *IntArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitFloatArg(ctx *FloatArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitDecimalArg(ctx *DecimalArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitBooleanArg(ctx *BooleanArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitStringArg(ctx *StringArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitDateArg(ctx *DateArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitTimeArg(ctx *TimeArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitTimestampArg(ctx *TimestampArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitTimestampTzArg(ctx *TimestampTzArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitIntervalYearArg(ctx *IntervalYearArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitIntervalDayArg(ctx *IntervalDayArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitFixedCharArg(ctx *FixedCharArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitVarCharArg(ctx *VarCharArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitFixedBinaryArg(ctx *FixedBinaryArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitPrecisionTimeArg(ctx *PrecisionTimeArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitPrecisionTimestampArg(ctx *PrecisionTimestampArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitPrecisionTimestampTZArg(ctx *PrecisionTimestampTZArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitListArg(ctx *ListArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitLiteralList(ctx *LiteralListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitDataType(ctx *DataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitBoolean(ctx *BooleanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitI8(ctx *I8Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitI16(ctx *I16Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitI32(ctx *I32Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitI64(ctx *I64Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitFp32(ctx *Fp32Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitFp64(ctx *Fp64Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitBinary(ctx *BinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitTimestamp(ctx *TimestampContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitTimestampTz(ctx *TimestampTzContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitDate(ctx *DateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitTime(ctx *TimeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitIntervalYear(ctx *IntervalYearContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitUuid(ctx *UuidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitUserDefined(ctx *UserDefinedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitBooleanType(ctx *BooleanTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitStringType(ctx *StringTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitBinaryType(ctx *BinaryTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitTimestampType(ctx *TimestampTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitTimestampTZType(ctx *TimestampTZTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitIntervalYearType(ctx *IntervalYearTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitIntervalDayType(ctx *IntervalDayTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitFixedCharType(ctx *FixedCharTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitVarCharType(ctx *VarCharTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitFixedBinaryType(ctx *FixedBinaryTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitDecimalType(ctx *DecimalTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitPrecisionTimeType(ctx *PrecisionTimeTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitPrecisionTimestampType(ctx *PrecisionTimestampTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitPrecisionTimestampTZType(ctx *PrecisionTimestampTZTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitList(ctx *ListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitParameterizedType(ctx *ParameterizedTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitSubstraitError(ctx *SubstraitErrorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitFuncOption(ctx *FuncOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitOptionName(ctx *OptionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitOptionValue(ctx *OptionValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitFuncOptions(ctx *FuncOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitNonReserved(ctx *NonReservedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFuncTestCaseParserVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}
