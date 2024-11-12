// Code generated from FuncTestCaseParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package baseparser // FuncTestCaseParser
import "github.com/antlr4-go/antlr/v4"

// FuncTestCaseParserListener is a complete listener for a parse tree produced by FuncTestCaseParser.
type FuncTestCaseParserListener interface {
	antlr.ParseTreeListener

	// EnterDoc is called when entering the doc production.
	EnterDoc(c *DocContext)

	// EnterHeader is called when entering the header production.
	EnterHeader(c *HeaderContext)

	// EnterVersion is called when entering the version production.
	EnterVersion(c *VersionContext)

	// EnterInclude is called when entering the include production.
	EnterInclude(c *IncludeContext)

	// EnterTestGroupDescription is called when entering the testGroupDescription production.
	EnterTestGroupDescription(c *TestGroupDescriptionContext)

	// EnterTestCase is called when entering the testCase production.
	EnterTestCase(c *TestCaseContext)

	// EnterScalarFuncTestGroup is called when entering the scalarFuncTestGroup production.
	EnterScalarFuncTestGroup(c *ScalarFuncTestGroupContext)

	// EnterAggregateFuncTestGroup is called when entering the aggregateFuncTestGroup production.
	EnterAggregateFuncTestGroup(c *AggregateFuncTestGroupContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterResult is called when entering the result production.
	EnterResult(c *ResultContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterAggFuncTestCase is called when entering the aggFuncTestCase production.
	EnterAggFuncTestCase(c *AggFuncTestCaseContext)

	// EnterMultiArgAggregateFuncCall is called when entering the multiArgAggregateFuncCall production.
	EnterMultiArgAggregateFuncCall(c *MultiArgAggregateFuncCallContext)

	// EnterCompactAggregateFuncCall is called when entering the compactAggregateFuncCall production.
	EnterCompactAggregateFuncCall(c *CompactAggregateFuncCallContext)

	// EnterSingleArgAggregateFuncCall is called when entering the singleArgAggregateFuncCall production.
	EnterSingleArgAggregateFuncCall(c *SingleArgAggregateFuncCallContext)

	// EnterTableData is called when entering the tableData production.
	EnterTableData(c *TableDataContext)

	// EnterTableRows is called when entering the tableRows production.
	EnterTableRows(c *TableRowsContext)

	// EnterDataColumn is called when entering the dataColumn production.
	EnterDataColumn(c *DataColumnContext)

	// EnterColumnValues is called when entering the columnValues production.
	EnterColumnValues(c *ColumnValuesContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterQualifiedAggregateFuncArgs is called when entering the qualifiedAggregateFuncArgs production.
	EnterQualifiedAggregateFuncArgs(c *QualifiedAggregateFuncArgsContext)

	// EnterAggregateFuncArgs is called when entering the aggregateFuncArgs production.
	EnterAggregateFuncArgs(c *AggregateFuncArgsContext)

	// EnterQualifiedAggregateFuncArg is called when entering the qualifiedAggregateFuncArg production.
	EnterQualifiedAggregateFuncArg(c *QualifiedAggregateFuncArgContext)

	// EnterAggregateFuncArg is called when entering the aggregateFuncArg production.
	EnterAggregateFuncArg(c *AggregateFuncArgContext)

	// EnterNumericLiteral is called when entering the numericLiteral production.
	EnterNumericLiteral(c *NumericLiteralContext)

	// EnterFloatLiteral is called when entering the floatLiteral production.
	EnterFloatLiteral(c *FloatLiteralContext)

	// EnterNullArg is called when entering the nullArg production.
	EnterNullArg(c *NullArgContext)

	// EnterIntArg is called when entering the intArg production.
	EnterIntArg(c *IntArgContext)

	// EnterFloatArg is called when entering the floatArg production.
	EnterFloatArg(c *FloatArgContext)

	// EnterDecimalArg is called when entering the decimalArg production.
	EnterDecimalArg(c *DecimalArgContext)

	// EnterBooleanArg is called when entering the booleanArg production.
	EnterBooleanArg(c *BooleanArgContext)

	// EnterStringArg is called when entering the stringArg production.
	EnterStringArg(c *StringArgContext)

	// EnterDateArg is called when entering the dateArg production.
	EnterDateArg(c *DateArgContext)

	// EnterTimeArg is called when entering the timeArg production.
	EnterTimeArg(c *TimeArgContext)

	// EnterTimestampArg is called when entering the timestampArg production.
	EnterTimestampArg(c *TimestampArgContext)

	// EnterTimestampTzArg is called when entering the timestampTzArg production.
	EnterTimestampTzArg(c *TimestampTzArgContext)

	// EnterIntervalYearArg is called when entering the intervalYearArg production.
	EnterIntervalYearArg(c *IntervalYearArgContext)

	// EnterIntervalDayArg is called when entering the intervalDayArg production.
	EnterIntervalDayArg(c *IntervalDayArgContext)

	// EnterListArg is called when entering the listArg production.
	EnterListArg(c *ListArgContext)

	// EnterLiteralList is called when entering the literalList production.
	EnterLiteralList(c *LiteralListContext)

	// EnterIntervalYearLiteral is called when entering the intervalYearLiteral production.
	EnterIntervalYearLiteral(c *IntervalYearLiteralContext)

	// EnterIntervalDayLiteral is called when entering the intervalDayLiteral production.
	EnterIntervalDayLiteral(c *IntervalDayLiteralContext)

	// EnterTimeInterval is called when entering the timeInterval production.
	EnterTimeInterval(c *TimeIntervalContext)

	// EnterDataType is called when entering the dataType production.
	EnterDataType(c *DataTypeContext)

	// EnterBoolean is called when entering the boolean production.
	EnterBoolean(c *BooleanContext)

	// EnterI8 is called when entering the i8 production.
	EnterI8(c *I8Context)

	// EnterI16 is called when entering the i16 production.
	EnterI16(c *I16Context)

	// EnterI32 is called when entering the i32 production.
	EnterI32(c *I32Context)

	// EnterI64 is called when entering the i64 production.
	EnterI64(c *I64Context)

	// EnterFp32 is called when entering the fp32 production.
	EnterFp32(c *Fp32Context)

	// EnterFp64 is called when entering the fp64 production.
	EnterFp64(c *Fp64Context)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterBinary is called when entering the binary production.
	EnterBinary(c *BinaryContext)

	// EnterTimestamp is called when entering the timestamp production.
	EnterTimestamp(c *TimestampContext)

	// EnterTimestampTz is called when entering the timestampTz production.
	EnterTimestampTz(c *TimestampTzContext)

	// EnterDate is called when entering the date production.
	EnterDate(c *DateContext)

	// EnterTime is called when entering the time production.
	EnterTime(c *TimeContext)

	// EnterIntervalDay is called when entering the intervalDay production.
	EnterIntervalDay(c *IntervalDayContext)

	// EnterIntervalYear is called when entering the intervalYear production.
	EnterIntervalYear(c *IntervalYearContext)

	// EnterUuid is called when entering the uuid production.
	EnterUuid(c *UuidContext)

	// EnterUserDefined is called when entering the userDefined production.
	EnterUserDefined(c *UserDefinedContext)

	// EnterBooleanType is called when entering the booleanType production.
	EnterBooleanType(c *BooleanTypeContext)

	// EnterStringType is called when entering the stringType production.
	EnterStringType(c *StringTypeContext)

	// EnterBinaryType is called when entering the binaryType production.
	EnterBinaryType(c *BinaryTypeContext)

	// EnterTimestampType is called when entering the timestampType production.
	EnterTimestampType(c *TimestampTypeContext)

	// EnterTimestampTZType is called when entering the timestampTZType production.
	EnterTimestampTZType(c *TimestampTZTypeContext)

	// EnterIntervalYearType is called when entering the intervalYearType production.
	EnterIntervalYearType(c *IntervalYearTypeContext)

	// EnterIntervalDayType is called when entering the intervalDayType production.
	EnterIntervalDayType(c *IntervalDayTypeContext)

	// EnterFixedChar is called when entering the fixedChar production.
	EnterFixedChar(c *FixedCharContext)

	// EnterVarChar is called when entering the varChar production.
	EnterVarChar(c *VarCharContext)

	// EnterFixedBinary is called when entering the fixedBinary production.
	EnterFixedBinary(c *FixedBinaryContext)

	// EnterDecimal is called when entering the decimal production.
	EnterDecimal(c *DecimalContext)

	// EnterPrecisionTimestamp is called when entering the precisionTimestamp production.
	EnterPrecisionTimestamp(c *PrecisionTimestampContext)

	// EnterPrecisionTimestampTZ is called when entering the precisionTimestampTZ production.
	EnterPrecisionTimestampTZ(c *PrecisionTimestampTZContext)

	// EnterList is called when entering the list production.
	EnterList(c *ListContext)

	// EnterParameterizedType is called when entering the parameterizedType production.
	EnterParameterizedType(c *ParameterizedTypeContext)

	// EnterIntegerLiteral is called when entering the integerLiteral production.
	EnterIntegerLiteral(c *IntegerLiteralContext)

	// EnterSubstraitError is called when entering the substraitError production.
	EnterSubstraitError(c *SubstraitErrorContext)

	// EnterFunc_option is called when entering the func_option production.
	EnterFunc_option(c *Func_optionContext)

	// EnterOption_name is called when entering the option_name production.
	EnterOption_name(c *Option_nameContext)

	// EnterOption_value is called when entering the option_value production.
	EnterOption_value(c *Option_valueContext)

	// EnterFunc_options is called when entering the func_options production.
	EnterFunc_options(c *Func_optionsContext)

	// EnterNonReserved is called when entering the nonReserved production.
	EnterNonReserved(c *NonReservedContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// ExitDoc is called when exiting the doc production.
	ExitDoc(c *DocContext)

	// ExitHeader is called when exiting the header production.
	ExitHeader(c *HeaderContext)

	// ExitVersion is called when exiting the version production.
	ExitVersion(c *VersionContext)

	// ExitInclude is called when exiting the include production.
	ExitInclude(c *IncludeContext)

	// ExitTestGroupDescription is called when exiting the testGroupDescription production.
	ExitTestGroupDescription(c *TestGroupDescriptionContext)

	// ExitTestCase is called when exiting the testCase production.
	ExitTestCase(c *TestCaseContext)

	// ExitScalarFuncTestGroup is called when exiting the scalarFuncTestGroup production.
	ExitScalarFuncTestGroup(c *ScalarFuncTestGroupContext)

	// ExitAggregateFuncTestGroup is called when exiting the aggregateFuncTestGroup production.
	ExitAggregateFuncTestGroup(c *AggregateFuncTestGroupContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitResult is called when exiting the result production.
	ExitResult(c *ResultContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitAggFuncTestCase is called when exiting the aggFuncTestCase production.
	ExitAggFuncTestCase(c *AggFuncTestCaseContext)

	// ExitMultiArgAggregateFuncCall is called when exiting the multiArgAggregateFuncCall production.
	ExitMultiArgAggregateFuncCall(c *MultiArgAggregateFuncCallContext)

	// ExitCompactAggregateFuncCall is called when exiting the compactAggregateFuncCall production.
	ExitCompactAggregateFuncCall(c *CompactAggregateFuncCallContext)

	// ExitSingleArgAggregateFuncCall is called when exiting the singleArgAggregateFuncCall production.
	ExitSingleArgAggregateFuncCall(c *SingleArgAggregateFuncCallContext)

	// ExitTableData is called when exiting the tableData production.
	ExitTableData(c *TableDataContext)

	// ExitTableRows is called when exiting the tableRows production.
	ExitTableRows(c *TableRowsContext)

	// ExitDataColumn is called when exiting the dataColumn production.
	ExitDataColumn(c *DataColumnContext)

	// ExitColumnValues is called when exiting the columnValues production.
	ExitColumnValues(c *ColumnValuesContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitQualifiedAggregateFuncArgs is called when exiting the qualifiedAggregateFuncArgs production.
	ExitQualifiedAggregateFuncArgs(c *QualifiedAggregateFuncArgsContext)

	// ExitAggregateFuncArgs is called when exiting the aggregateFuncArgs production.
	ExitAggregateFuncArgs(c *AggregateFuncArgsContext)

	// ExitQualifiedAggregateFuncArg is called when exiting the qualifiedAggregateFuncArg production.
	ExitQualifiedAggregateFuncArg(c *QualifiedAggregateFuncArgContext)

	// ExitAggregateFuncArg is called when exiting the aggregateFuncArg production.
	ExitAggregateFuncArg(c *AggregateFuncArgContext)

	// ExitNumericLiteral is called when exiting the numericLiteral production.
	ExitNumericLiteral(c *NumericLiteralContext)

	// ExitFloatLiteral is called when exiting the floatLiteral production.
	ExitFloatLiteral(c *FloatLiteralContext)

	// ExitNullArg is called when exiting the nullArg production.
	ExitNullArg(c *NullArgContext)

	// ExitIntArg is called when exiting the intArg production.
	ExitIntArg(c *IntArgContext)

	// ExitFloatArg is called when exiting the floatArg production.
	ExitFloatArg(c *FloatArgContext)

	// ExitDecimalArg is called when exiting the decimalArg production.
	ExitDecimalArg(c *DecimalArgContext)

	// ExitBooleanArg is called when exiting the booleanArg production.
	ExitBooleanArg(c *BooleanArgContext)

	// ExitStringArg is called when exiting the stringArg production.
	ExitStringArg(c *StringArgContext)

	// ExitDateArg is called when exiting the dateArg production.
	ExitDateArg(c *DateArgContext)

	// ExitTimeArg is called when exiting the timeArg production.
	ExitTimeArg(c *TimeArgContext)

	// ExitTimestampArg is called when exiting the timestampArg production.
	ExitTimestampArg(c *TimestampArgContext)

	// ExitTimestampTzArg is called when exiting the timestampTzArg production.
	ExitTimestampTzArg(c *TimestampTzArgContext)

	// ExitIntervalYearArg is called when exiting the intervalYearArg production.
	ExitIntervalYearArg(c *IntervalYearArgContext)

	// ExitIntervalDayArg is called when exiting the intervalDayArg production.
	ExitIntervalDayArg(c *IntervalDayArgContext)

	// ExitListArg is called when exiting the listArg production.
	ExitListArg(c *ListArgContext)

	// ExitLiteralList is called when exiting the literalList production.
	ExitLiteralList(c *LiteralListContext)

	// ExitIntervalYearLiteral is called when exiting the intervalYearLiteral production.
	ExitIntervalYearLiteral(c *IntervalYearLiteralContext)

	// ExitIntervalDayLiteral is called when exiting the intervalDayLiteral production.
	ExitIntervalDayLiteral(c *IntervalDayLiteralContext)

	// ExitTimeInterval is called when exiting the timeInterval production.
	ExitTimeInterval(c *TimeIntervalContext)

	// ExitDataType is called when exiting the dataType production.
	ExitDataType(c *DataTypeContext)

	// ExitBoolean is called when exiting the boolean production.
	ExitBoolean(c *BooleanContext)

	// ExitI8 is called when exiting the i8 production.
	ExitI8(c *I8Context)

	// ExitI16 is called when exiting the i16 production.
	ExitI16(c *I16Context)

	// ExitI32 is called when exiting the i32 production.
	ExitI32(c *I32Context)

	// ExitI64 is called when exiting the i64 production.
	ExitI64(c *I64Context)

	// ExitFp32 is called when exiting the fp32 production.
	ExitFp32(c *Fp32Context)

	// ExitFp64 is called when exiting the fp64 production.
	ExitFp64(c *Fp64Context)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitBinary is called when exiting the binary production.
	ExitBinary(c *BinaryContext)

	// ExitTimestamp is called when exiting the timestamp production.
	ExitTimestamp(c *TimestampContext)

	// ExitTimestampTz is called when exiting the timestampTz production.
	ExitTimestampTz(c *TimestampTzContext)

	// ExitDate is called when exiting the date production.
	ExitDate(c *DateContext)

	// ExitTime is called when exiting the time production.
	ExitTime(c *TimeContext)

	// ExitIntervalDay is called when exiting the intervalDay production.
	ExitIntervalDay(c *IntervalDayContext)

	// ExitIntervalYear is called when exiting the intervalYear production.
	ExitIntervalYear(c *IntervalYearContext)

	// ExitUuid is called when exiting the uuid production.
	ExitUuid(c *UuidContext)

	// ExitUserDefined is called when exiting the userDefined production.
	ExitUserDefined(c *UserDefinedContext)

	// ExitBooleanType is called when exiting the booleanType production.
	ExitBooleanType(c *BooleanTypeContext)

	// ExitStringType is called when exiting the stringType production.
	ExitStringType(c *StringTypeContext)

	// ExitBinaryType is called when exiting the binaryType production.
	ExitBinaryType(c *BinaryTypeContext)

	// ExitTimestampType is called when exiting the timestampType production.
	ExitTimestampType(c *TimestampTypeContext)

	// ExitTimestampTZType is called when exiting the timestampTZType production.
	ExitTimestampTZType(c *TimestampTZTypeContext)

	// ExitIntervalYearType is called when exiting the intervalYearType production.
	ExitIntervalYearType(c *IntervalYearTypeContext)

	// ExitIntervalDayType is called when exiting the intervalDayType production.
	ExitIntervalDayType(c *IntervalDayTypeContext)

	// ExitFixedChar is called when exiting the fixedChar production.
	ExitFixedChar(c *FixedCharContext)

	// ExitVarChar is called when exiting the varChar production.
	ExitVarChar(c *VarCharContext)

	// ExitFixedBinary is called when exiting the fixedBinary production.
	ExitFixedBinary(c *FixedBinaryContext)

	// ExitDecimal is called when exiting the decimal production.
	ExitDecimal(c *DecimalContext)

	// ExitPrecisionTimestamp is called when exiting the precisionTimestamp production.
	ExitPrecisionTimestamp(c *PrecisionTimestampContext)

	// ExitPrecisionTimestampTZ is called when exiting the precisionTimestampTZ production.
	ExitPrecisionTimestampTZ(c *PrecisionTimestampTZContext)

	// ExitList is called when exiting the list production.
	ExitList(c *ListContext)

	// ExitParameterizedType is called when exiting the parameterizedType production.
	ExitParameterizedType(c *ParameterizedTypeContext)

	// ExitIntegerLiteral is called when exiting the integerLiteral production.
	ExitIntegerLiteral(c *IntegerLiteralContext)

	// ExitSubstraitError is called when exiting the substraitError production.
	ExitSubstraitError(c *SubstraitErrorContext)

	// ExitFunc_option is called when exiting the func_option production.
	ExitFunc_option(c *Func_optionContext)

	// ExitOption_name is called when exiting the option_name production.
	ExitOption_name(c *Option_nameContext)

	// ExitOption_value is called when exiting the option_value production.
	ExitOption_value(c *Option_valueContext)

	// ExitFunc_options is called when exiting the func_options production.
	ExitFunc_options(c *Func_optionsContext)

	// ExitNonReserved is called when exiting the nonReserved production.
	ExitNonReserved(c *NonReservedContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)
}
