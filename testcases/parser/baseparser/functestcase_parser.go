// Code generated from FuncTestCaseParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package baseparser // FuncTestCaseParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type FuncTestCaseParser struct {
	*antlr.BaseParser
}

var FuncTestCaseParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func functestcaseparserParserInit() {
	staticData := &FuncTestCaseParserParserStaticData
	staticData.LiteralNames = []string{
		"", "", "'###'", "'SUBSTRAIT_SCALAR_TEST'", "'SUBSTRAIT_AGGREGATE_TEST'",
		"'SUBSTRAIT_INCLUDE'", "", "", "'DEFINE'", "'<!ERROR>'", "'<!UNDEFINED>'",
		"'OVERFLOW'", "'ROUNDING'", "'ERROR'", "'SATURATE'", "'SILENT'", "'TIE_TO_EVEN'",
		"'NAN'", "'ACCEPT_NULLS'", "'IGNORE_NULLS'", "'NULL_HANDLING'", "'SPACES_ONLY'",
		"'TRUNCATE'", "", "", "", "", "", "", "", "", "'P'", "'T'", "'Y'", "'M'",
		"'D'", "'H'", "'S'", "'F'", "", "", "", "", "'null'", "", "", "", "",
		"'IF'", "'THEN'", "'ELSE'", "'BOOLEAN'", "'I8'", "'I16'", "'I32'", "'I64'",
		"'FP32'", "'FP64'", "'STRING'", "'BINARY'", "'TIMESTAMP'", "'TIMESTAMP_TZ'",
		"'DATE'", "'TIME'", "'INTERVAL_YEAR'", "'INTERVAL_DAY'", "'UUID'", "'DECIMAL'",
		"'PRECISION_TIMESTAMP'", "'PRECISION_TIMESTAMP_TZ'", "'FIXEDCHAR'",
		"'VARCHAR'", "'FIXEDBINARY'", "'STRUCT'", "'NSTRUCT'", "'LIST'", "'MAP'",
		"'U!'", "'BOOL'", "'STR'", "'VBIN'", "'TS'", "'TSTZ'", "'IYEAR'", "'IDAY'",
		"'DEC'", "'PTS'", "'PTSTZ'", "'FCHAR'", "'VCHAR'", "'FBIN'", "'ANY'",
		"", "'::'", "'+'", "'-'", "'*'", "'/'", "'%'", "'='", "'!='", "'>='",
		"'<='", "'>'", "'<'", "'!'", "'('", "')'", "'['", "']'", "','", "':'",
		"'?'", "'#'", "'.'", "'AND'", "'OR'", "':='",
	}
	staticData.SymbolicNames = []string{
		"", "Whitespace", "TripleHash", "SubstraitScalarTest", "SubstraitAggregateTest",
		"SubstraitInclude", "FormatVersion", "DescriptionLine", "Define", "ErrorResult",
		"UndefineResult", "Overflow", "Rounding", "Error", "Saturate", "Silent",
		"TieToEven", "NaN", "AcceptNulls", "IgnoreNulls", "NullHandling", "SpacesOnly",
		"Truncate", "IntegerLiteral", "DecimalLiteral", "FloatLiteral", "BooleanLiteral",
		"TimestampTzLiteral", "TimestampLiteral", "TimeLiteral", "DateLiteral",
		"PeriodPrefix", "TimePrefix", "YearPrefix", "MSuffix", "DaySuffix",
		"HourSuffix", "SecondSuffix", "FractionalSecondSuffix", "OAngleBracket",
		"CAngleBracket", "IntervalYearLiteral", "IntervalDayLiteral", "NullLiteral",
		"StringLiteral", "ColumnName", "LineComment", "BlockComment", "If",
		"Then", "Else", "Boolean", "I8", "I16", "I32", "I64", "FP32", "FP64",
		"String", "Binary", "Timestamp", "Timestamp_TZ", "Date", "Time", "Interval_Year",
		"Interval_Day", "UUID", "Decimal", "Precision_Timestamp", "Precision_Timestamp_TZ",
		"FixedChar", "VarChar", "FixedBinary", "Struct", "NStruct", "List",
		"Map", "UserDefined", "Bool", "Str", "VBin", "Ts", "TsTZ", "IYear",
		"IDay", "Dec", "PTs", "PTsTZ", "FChar", "VChar", "FBin", "Any", "AnyVar",
		"DoubleColon", "Plus", "Minus", "Asterisk", "ForwardSlash", "Percent",
		"Eq", "Ne", "Gte", "Lte", "Gt", "Lt", "Bang", "OParen", "CParen", "OBracket",
		"CBracket", "Comma", "Colon", "QMark", "Hash", "Dot", "And", "Or", "Assign",
		"Number", "Identifier", "Newline",
	}
	staticData.RuleNames = []string{
		"doc", "header", "version", "include", "testGroupDescription", "testCase",
		"testGroup", "arguments", "result", "argument", "aggFuncTestCase", "aggFuncCall",
		"tableData", "tableRows", "dataColumn", "columnValues", "literal", "qualifiedAggregateFuncArgs",
		"aggregateFuncArgs", "qualifiedAggregateFuncArg", "aggregateFuncArg",
		"numericLiteral", "floatLiteral", "nullArg", "intArg", "floatArg", "decimalArg",
		"booleanArg", "stringArg", "dateArg", "timeArg", "timestampArg", "timestampTzArg",
		"intervalYearArg", "intervalDayArg", "listArg", "literalList", "intervalYearLiteral",
		"intervalDayLiteral", "timeInterval", "dataType", "scalarType", "booleanType",
		"stringType", "binaryType", "timestampType", "timestampTZType", "intervalYearType",
		"intervalDayType", "fixedCharType", "varCharType", "fixedBinaryType",
		"decimalType", "precisionTimestampType", "precisionTimestampTZType",
		"listType", "parameterizedType", "numericParameter", "substraitError",
		"func_option", "option_name", "option_value", "func_options", "nonReserved",
		"identifier",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 120, 592, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57, 7, 57,
		2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7, 62, 2,
		63, 7, 63, 2, 64, 7, 64, 1, 0, 1, 0, 4, 0, 133, 8, 0, 11, 0, 12, 0, 134,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 153, 8, 3, 10, 3, 12, 3, 156, 9, 3, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 168, 8, 5,
		1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 4, 6, 175, 8, 6, 11, 6, 12, 6, 176, 1, 6,
		1, 6, 4, 6, 181, 8, 6, 11, 6, 12, 6, 182, 3, 6, 185, 8, 6, 1, 7, 1, 7,
		1, 7, 5, 7, 190, 8, 7, 10, 7, 12, 7, 193, 9, 7, 1, 8, 1, 8, 3, 8, 197,
		8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 3, 9, 212, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10,
		219, 8, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 3, 11, 241, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 5,
		12, 249, 8, 12, 10, 12, 12, 12, 252, 9, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 262, 8, 13, 10, 13, 12, 13, 265, 9,
		13, 3, 13, 267, 8, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15,
		1, 15, 1, 15, 1, 15, 5, 15, 279, 8, 15, 10, 15, 12, 15, 282, 9, 15, 3,
		15, 284, 8, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 298, 8, 16, 1, 17, 1, 17, 1, 17, 5,
		17, 303, 8, 17, 10, 17, 12, 17, 306, 9, 17, 1, 18, 1, 18, 1, 18, 5, 18,
		311, 8, 18, 10, 18, 12, 18, 314, 9, 18, 1, 19, 1, 19, 1, 19, 1, 19, 3,
		19, 320, 8, 19, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 326, 8, 20, 1, 21, 1,
		21, 1, 21, 3, 21, 331, 8, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29,
		1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 5,
		36, 391, 8, 36, 10, 36, 12, 36, 394, 9, 36, 3, 36, 396, 8, 36, 1, 36, 1,
		36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 3, 37, 406, 8, 37, 1, 37,
		1, 37, 1, 37, 3, 37, 411, 8, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1,
		38, 3, 38, 419, 8, 38, 1, 38, 1, 38, 1, 38, 3, 38, 424, 8, 38, 1, 39, 1,
		39, 1, 39, 1, 39, 3, 39, 430, 8, 39, 1, 39, 1, 39, 3, 39, 434, 8, 39, 1,
		39, 1, 39, 3, 39, 438, 8, 39, 1, 39, 1, 39, 1, 39, 1, 39, 3, 39, 444, 8,
		39, 1, 39, 1, 39, 3, 39, 448, 8, 39, 1, 39, 1, 39, 1, 39, 1, 39, 3, 39,
		454, 8, 39, 1, 39, 1, 39, 3, 39, 458, 8, 39, 1, 40, 1, 40, 3, 40, 462,
		8, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1,
		41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 3, 41, 482,
		8, 41, 1, 42, 1, 42, 1, 43, 1, 43, 1, 44, 1, 44, 1, 45, 1, 45, 1, 46, 1,
		46, 1, 47, 1, 47, 1, 48, 1, 48, 1, 49, 1, 49, 3, 49, 500, 8, 49, 1, 49,
		1, 49, 1, 49, 1, 49, 1, 50, 1, 50, 3, 50, 508, 8, 50, 1, 50, 1, 50, 1,
		50, 1, 50, 1, 51, 1, 51, 3, 51, 516, 8, 51, 1, 51, 1, 51, 1, 51, 1, 51,
		1, 52, 1, 52, 3, 52, 524, 8, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1,
		52, 3, 52, 532, 8, 52, 1, 53, 1, 53, 3, 53, 536, 8, 53, 1, 53, 1, 53, 1,
		53, 1, 53, 1, 54, 1, 54, 3, 54, 544, 8, 54, 1, 54, 1, 54, 1, 54, 1, 54,
		1, 55, 1, 55, 3, 55, 552, 8, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 56, 1,
		56, 1, 56, 1, 56, 1, 56, 1, 56, 3, 56, 564, 8, 56, 1, 57, 1, 57, 1, 58,
		1, 58, 1, 59, 1, 59, 1, 59, 1, 59, 1, 60, 1, 60, 1, 61, 1, 61, 1, 62, 1,
		62, 1, 62, 5, 62, 581, 8, 62, 10, 62, 12, 62, 584, 9, 62, 1, 63, 1, 63,
		1, 64, 1, 64, 3, 64, 590, 8, 64, 1, 64, 0, 0, 65, 0, 2, 4, 6, 8, 10, 12,
		14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48,
		50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84,
		86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112, 114, 116,
		118, 120, 122, 124, 126, 128, 0, 21, 1, 0, 3, 4, 2, 0, 17, 17, 25, 25,
		1, 0, 52, 55, 1, 0, 56, 57, 2, 0, 51, 51, 78, 78, 2, 0, 58, 58, 79, 79,
		2, 0, 59, 59, 80, 80, 2, 0, 60, 60, 81, 81, 2, 0, 61, 61, 82, 82, 2, 0,
		64, 64, 83, 83, 2, 0, 65, 65, 84, 84, 2, 0, 70, 70, 88, 88, 2, 0, 71, 71,
		89, 89, 2, 0, 72, 72, 90, 90, 2, 0, 67, 67, 85, 85, 2, 0, 68, 68, 86, 86,
		2, 0, 69, 69, 87, 87, 1, 0, 9, 10, 3, 0, 11, 12, 20, 21, 119, 119, 5, 0,
		13, 19, 22, 22, 26, 26, 43, 43, 119, 119, 2, 0, 22, 22, 115, 116, 616,
		0, 130, 1, 0, 0, 0, 2, 138, 1, 0, 0, 0, 4, 141, 1, 0, 0, 0, 6, 146, 1,
		0, 0, 0, 8, 157, 1, 0, 0, 0, 10, 159, 1, 0, 0, 0, 12, 184, 1, 0, 0, 0,
		14, 186, 1, 0, 0, 0, 16, 196, 1, 0, 0, 0, 18, 211, 1, 0, 0, 0, 20, 213,
		1, 0, 0, 0, 22, 240, 1, 0, 0, 0, 24, 242, 1, 0, 0, 0, 26, 257, 1, 0, 0,
		0, 28, 270, 1, 0, 0, 0, 30, 274, 1, 0, 0, 0, 32, 297, 1, 0, 0, 0, 34, 299,
		1, 0, 0, 0, 36, 307, 1, 0, 0, 0, 38, 319, 1, 0, 0, 0, 40, 325, 1, 0, 0,
		0, 42, 330, 1, 0, 0, 0, 44, 332, 1, 0, 0, 0, 46, 334, 1, 0, 0, 0, 48, 338,
		1, 0, 0, 0, 50, 342, 1, 0, 0, 0, 52, 346, 1, 0, 0, 0, 54, 350, 1, 0, 0,
		0, 56, 354, 1, 0, 0, 0, 58, 358, 1, 0, 0, 0, 60, 362, 1, 0, 0, 0, 62, 366,
		1, 0, 0, 0, 64, 370, 1, 0, 0, 0, 66, 374, 1, 0, 0, 0, 68, 378, 1, 0, 0,
		0, 70, 382, 1, 0, 0, 0, 72, 386, 1, 0, 0, 0, 74, 410, 1, 0, 0, 0, 76, 423,
		1, 0, 0, 0, 78, 457, 1, 0, 0, 0, 80, 461, 1, 0, 0, 0, 82, 481, 1, 0, 0,
		0, 84, 483, 1, 0, 0, 0, 86, 485, 1, 0, 0, 0, 88, 487, 1, 0, 0, 0, 90, 489,
		1, 0, 0, 0, 92, 491, 1, 0, 0, 0, 94, 493, 1, 0, 0, 0, 96, 495, 1, 0, 0,
		0, 98, 497, 1, 0, 0, 0, 100, 505, 1, 0, 0, 0, 102, 513, 1, 0, 0, 0, 104,
		521, 1, 0, 0, 0, 106, 533, 1, 0, 0, 0, 108, 541, 1, 0, 0, 0, 110, 549,
		1, 0, 0, 0, 112, 563, 1, 0, 0, 0, 114, 565, 1, 0, 0, 0, 116, 567, 1, 0,
		0, 0, 118, 569, 1, 0, 0, 0, 120, 573, 1, 0, 0, 0, 122, 575, 1, 0, 0, 0,
		124, 577, 1, 0, 0, 0, 126, 585, 1, 0, 0, 0, 128, 589, 1, 0, 0, 0, 130,
		132, 3, 2, 1, 0, 131, 133, 3, 12, 6, 0, 132, 131, 1, 0, 0, 0, 133, 134,
		1, 0, 0, 0, 134, 132, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 136, 1, 0,
		0, 0, 136, 137, 5, 0, 0, 1, 137, 1, 1, 0, 0, 0, 138, 139, 3, 4, 2, 0, 139,
		140, 3, 6, 3, 0, 140, 3, 1, 0, 0, 0, 141, 142, 5, 2, 0, 0, 142, 143, 7,
		0, 0, 0, 143, 144, 5, 111, 0, 0, 144, 145, 5, 6, 0, 0, 145, 5, 1, 0, 0,
		0, 146, 147, 5, 2, 0, 0, 147, 148, 5, 5, 0, 0, 148, 149, 5, 111, 0, 0,
		149, 154, 5, 44, 0, 0, 150, 151, 5, 110, 0, 0, 151, 153, 5, 44, 0, 0, 152,
		150, 1, 0, 0, 0, 153, 156, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 154, 155,
		1, 0, 0, 0, 155, 7, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 157, 158, 5, 7, 0,
		0, 158, 9, 1, 0, 0, 0, 159, 160, 3, 128, 64, 0, 160, 161, 5, 106, 0, 0,
		161, 162, 3, 14, 7, 0, 162, 167, 5, 107, 0, 0, 163, 164, 5, 108, 0, 0,
		164, 165, 3, 124, 62, 0, 165, 166, 5, 109, 0, 0, 166, 168, 1, 0, 0, 0,
		167, 163, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169,
		170, 5, 99, 0, 0, 170, 171, 3, 16, 8, 0, 171, 11, 1, 0, 0, 0, 172, 174,
		3, 8, 4, 0, 173, 175, 3, 10, 5, 0, 174, 173, 1, 0, 0, 0, 175, 176, 1, 0,
		0, 0, 176, 174, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 185, 1, 0, 0, 0,
		178, 180, 3, 8, 4, 0, 179, 181, 3, 20, 10, 0, 180, 179, 1, 0, 0, 0, 181,
		182, 1, 0, 0, 0, 182, 180, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 185,
		1, 0, 0, 0, 184, 172, 1, 0, 0, 0, 184, 178, 1, 0, 0, 0, 185, 13, 1, 0,
		0, 0, 186, 191, 3, 18, 9, 0, 187, 188, 5, 110, 0, 0, 188, 190, 3, 18, 9,
		0, 189, 187, 1, 0, 0, 0, 190, 193, 1, 0, 0, 0, 191, 189, 1, 0, 0, 0, 191,
		192, 1, 0, 0, 0, 192, 15, 1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 194, 197, 3,
		18, 9, 0, 195, 197, 3, 116, 58, 0, 196, 194, 1, 0, 0, 0, 196, 195, 1, 0,
		0, 0, 197, 17, 1, 0, 0, 0, 198, 212, 3, 46, 23, 0, 199, 212, 3, 48, 24,
		0, 200, 212, 3, 50, 25, 0, 201, 212, 3, 54, 27, 0, 202, 212, 3, 56, 28,
		0, 203, 212, 3, 52, 26, 0, 204, 212, 3, 58, 29, 0, 205, 212, 3, 60, 30,
		0, 206, 212, 3, 62, 31, 0, 207, 212, 3, 64, 32, 0, 208, 212, 3, 66, 33,
		0, 209, 212, 3, 68, 34, 0, 210, 212, 3, 70, 35, 0, 211, 198, 1, 0, 0, 0,
		211, 199, 1, 0, 0, 0, 211, 200, 1, 0, 0, 0, 211, 201, 1, 0, 0, 0, 211,
		202, 1, 0, 0, 0, 211, 203, 1, 0, 0, 0, 211, 204, 1, 0, 0, 0, 211, 205,
		1, 0, 0, 0, 211, 206, 1, 0, 0, 0, 211, 207, 1, 0, 0, 0, 211, 208, 1, 0,
		0, 0, 211, 209, 1, 0, 0, 0, 211, 210, 1, 0, 0, 0, 212, 19, 1, 0, 0, 0,
		213, 218, 3, 22, 11, 0, 214, 215, 5, 108, 0, 0, 215, 216, 3, 124, 62, 0,
		216, 217, 5, 109, 0, 0, 217, 219, 1, 0, 0, 0, 218, 214, 1, 0, 0, 0, 218,
		219, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220, 221, 5, 99, 0, 0, 221, 222,
		3, 16, 8, 0, 222, 21, 1, 0, 0, 0, 223, 224, 3, 24, 12, 0, 224, 225, 3,
		128, 64, 0, 225, 226, 5, 106, 0, 0, 226, 227, 3, 34, 17, 0, 227, 228, 5,
		107, 0, 0, 228, 241, 1, 0, 0, 0, 229, 230, 3, 26, 13, 0, 230, 231, 3, 128,
		64, 0, 231, 232, 5, 106, 0, 0, 232, 233, 3, 36, 18, 0, 233, 234, 5, 107,
		0, 0, 234, 241, 1, 0, 0, 0, 235, 236, 3, 128, 64, 0, 236, 237, 5, 106,
		0, 0, 237, 238, 3, 28, 14, 0, 238, 239, 5, 107, 0, 0, 239, 241, 1, 0, 0,
		0, 240, 223, 1, 0, 0, 0, 240, 229, 1, 0, 0, 0, 240, 235, 1, 0, 0, 0, 241,
		23, 1, 0, 0, 0, 242, 243, 5, 8, 0, 0, 243, 244, 5, 119, 0, 0, 244, 245,
		5, 106, 0, 0, 245, 250, 3, 80, 40, 0, 246, 247, 5, 110, 0, 0, 247, 249,
		3, 80, 40, 0, 248, 246, 1, 0, 0, 0, 249, 252, 1, 0, 0, 0, 250, 248, 1,
		0, 0, 0, 250, 251, 1, 0, 0, 0, 251, 253, 1, 0, 0, 0, 252, 250, 1, 0, 0,
		0, 253, 254, 5, 107, 0, 0, 254, 255, 5, 99, 0, 0, 255, 256, 3, 26, 13,
		0, 256, 25, 1, 0, 0, 0, 257, 266, 5, 106, 0, 0, 258, 263, 3, 30, 15, 0,
		259, 260, 5, 110, 0, 0, 260, 262, 3, 30, 15, 0, 261, 259, 1, 0, 0, 0, 262,
		265, 1, 0, 0, 0, 263, 261, 1, 0, 0, 0, 263, 264, 1, 0, 0, 0, 264, 267,
		1, 0, 0, 0, 265, 263, 1, 0, 0, 0, 266, 258, 1, 0, 0, 0, 266, 267, 1, 0,
		0, 0, 267, 268, 1, 0, 0, 0, 268, 269, 5, 107, 0, 0, 269, 27, 1, 0, 0, 0,
		270, 271, 3, 30, 15, 0, 271, 272, 5, 93, 0, 0, 272, 273, 3, 80, 40, 0,
		273, 29, 1, 0, 0, 0, 274, 283, 5, 106, 0, 0, 275, 280, 3, 32, 16, 0, 276,
		277, 5, 110, 0, 0, 277, 279, 3, 32, 16, 0, 278, 276, 1, 0, 0, 0, 279, 282,
		1, 0, 0, 0, 280, 278, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0, 281, 284, 1, 0,
		0, 0, 282, 280, 1, 0, 0, 0, 283, 275, 1, 0, 0, 0, 283, 284, 1, 0, 0, 0,
		284, 285, 1, 0, 0, 0, 285, 286, 5, 107, 0, 0, 286, 31, 1, 0, 0, 0, 287,
		298, 5, 43, 0, 0, 288, 298, 3, 42, 21, 0, 289, 298, 5, 26, 0, 0, 290, 298,
		5, 44, 0, 0, 291, 298, 5, 30, 0, 0, 292, 298, 5, 29, 0, 0, 293, 298, 5,
		28, 0, 0, 294, 298, 5, 27, 0, 0, 295, 298, 5, 41, 0, 0, 296, 298, 5, 42,
		0, 0, 297, 287, 1, 0, 0, 0, 297, 288, 1, 0, 0, 0, 297, 289, 1, 0, 0, 0,
		297, 290, 1, 0, 0, 0, 297, 291, 1, 0, 0, 0, 297, 292, 1, 0, 0, 0, 297,
		293, 1, 0, 0, 0, 297, 294, 1, 0, 0, 0, 297, 295, 1, 0, 0, 0, 297, 296,
		1, 0, 0, 0, 298, 33, 1, 0, 0, 0, 299, 304, 3, 38, 19, 0, 300, 301, 5, 110,
		0, 0, 301, 303, 3, 38, 19, 0, 302, 300, 1, 0, 0, 0, 303, 306, 1, 0, 0,
		0, 304, 302, 1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305, 35, 1, 0, 0, 0, 306,
		304, 1, 0, 0, 0, 307, 312, 3, 40, 20, 0, 308, 309, 5, 110, 0, 0, 309, 311,
		3, 40, 20, 0, 310, 308, 1, 0, 0, 0, 311, 314, 1, 0, 0, 0, 312, 310, 1,
		0, 0, 0, 312, 313, 1, 0, 0, 0, 313, 37, 1, 0, 0, 0, 314, 312, 1, 0, 0,
		0, 315, 316, 5, 119, 0, 0, 316, 317, 5, 114, 0, 0, 317, 320, 5, 45, 0,
		0, 318, 320, 3, 18, 9, 0, 319, 315, 1, 0, 0, 0, 319, 318, 1, 0, 0, 0, 320,
		39, 1, 0, 0, 0, 321, 322, 5, 45, 0, 0, 322, 323, 5, 93, 0, 0, 323, 326,
		3, 80, 40, 0, 324, 326, 3, 18, 9, 0, 325, 321, 1, 0, 0, 0, 325, 324, 1,
		0, 0, 0, 326, 41, 1, 0, 0, 0, 327, 331, 5, 24, 0, 0, 328, 331, 5, 23, 0,
		0, 329, 331, 3, 44, 22, 0, 330, 327, 1, 0, 0, 0, 330, 328, 1, 0, 0, 0,
		330, 329, 1, 0, 0, 0, 331, 43, 1, 0, 0, 0, 332, 333, 7, 1, 0, 0, 333, 45,
		1, 0, 0, 0, 334, 335, 5, 43, 0, 0, 335, 336, 5, 93, 0, 0, 336, 337, 3,
		80, 40, 0, 337, 47, 1, 0, 0, 0, 338, 339, 5, 23, 0, 0, 339, 340, 5, 93,
		0, 0, 340, 341, 7, 2, 0, 0, 341, 49, 1, 0, 0, 0, 342, 343, 3, 42, 21, 0,
		343, 344, 5, 93, 0, 0, 344, 345, 7, 3, 0, 0, 345, 51, 1, 0, 0, 0, 346,
		347, 3, 42, 21, 0, 347, 348, 5, 93, 0, 0, 348, 349, 3, 104, 52, 0, 349,
		53, 1, 0, 0, 0, 350, 351, 5, 26, 0, 0, 351, 352, 5, 93, 0, 0, 352, 353,
		3, 84, 42, 0, 353, 55, 1, 0, 0, 0, 354, 355, 5, 44, 0, 0, 355, 356, 5,
		93, 0, 0, 356, 357, 3, 86, 43, 0, 357, 57, 1, 0, 0, 0, 358, 359, 5, 30,
		0, 0, 359, 360, 5, 93, 0, 0, 360, 361, 5, 62, 0, 0, 361, 59, 1, 0, 0, 0,
		362, 363, 5, 29, 0, 0, 363, 364, 5, 93, 0, 0, 364, 365, 5, 63, 0, 0, 365,
		61, 1, 0, 0, 0, 366, 367, 5, 28, 0, 0, 367, 368, 5, 93, 0, 0, 368, 369,
		3, 90, 45, 0, 369, 63, 1, 0, 0, 0, 370, 371, 5, 27, 0, 0, 371, 372, 5,
		93, 0, 0, 372, 373, 3, 92, 46, 0, 373, 65, 1, 0, 0, 0, 374, 375, 5, 41,
		0, 0, 375, 376, 5, 93, 0, 0, 376, 377, 3, 94, 47, 0, 377, 67, 1, 0, 0,
		0, 378, 379, 5, 42, 0, 0, 379, 380, 5, 93, 0, 0, 380, 381, 3, 96, 48, 0,
		381, 69, 1, 0, 0, 0, 382, 383, 3, 72, 36, 0, 383, 384, 5, 93, 0, 0, 384,
		385, 3, 110, 55, 0, 385, 71, 1, 0, 0, 0, 386, 395, 5, 108, 0, 0, 387, 392,
		3, 32, 16, 0, 388, 389, 5, 110, 0, 0, 389, 391, 3, 32, 16, 0, 390, 388,
		1, 0, 0, 0, 391, 394, 1, 0, 0, 0, 392, 390, 1, 0, 0, 0, 392, 393, 1, 0,
		0, 0, 393, 396, 1, 0, 0, 0, 394, 392, 1, 0, 0, 0, 395, 387, 1, 0, 0, 0,
		395, 396, 1, 0, 0, 0, 396, 397, 1, 0, 0, 0, 397, 398, 5, 109, 0, 0, 398,
		73, 1, 0, 0, 0, 399, 400, 5, 31, 0, 0, 400, 401, 5, 23, 0, 0, 401, 402,
		5, 33, 0, 0, 402, 405, 1, 0, 0, 0, 403, 404, 5, 23, 0, 0, 404, 406, 5,
		34, 0, 0, 405, 403, 1, 0, 0, 0, 405, 406, 1, 0, 0, 0, 406, 411, 1, 0, 0,
		0, 407, 408, 5, 31, 0, 0, 408, 409, 5, 23, 0, 0, 409, 411, 5, 34, 0, 0,
		410, 399, 1, 0, 0, 0, 410, 407, 1, 0, 0, 0, 411, 75, 1, 0, 0, 0, 412, 413,
		5, 31, 0, 0, 413, 414, 5, 23, 0, 0, 414, 415, 5, 35, 0, 0, 415, 418, 1,
		0, 0, 0, 416, 417, 5, 32, 0, 0, 417, 419, 3, 78, 39, 0, 418, 416, 1, 0,
		0, 0, 418, 419, 1, 0, 0, 0, 419, 424, 1, 0, 0, 0, 420, 421, 5, 31, 0, 0,
		421, 422, 5, 32, 0, 0, 422, 424, 3, 78, 39, 0, 423, 412, 1, 0, 0, 0, 423,
		420, 1, 0, 0, 0, 424, 77, 1, 0, 0, 0, 425, 426, 5, 23, 0, 0, 426, 429,
		5, 36, 0, 0, 427, 428, 5, 23, 0, 0, 428, 430, 5, 34, 0, 0, 429, 427, 1,
		0, 0, 0, 429, 430, 1, 0, 0, 0, 430, 433, 1, 0, 0, 0, 431, 432, 5, 23, 0,
		0, 432, 434, 5, 37, 0, 0, 433, 431, 1, 0, 0, 0, 433, 434, 1, 0, 0, 0, 434,
		437, 1, 0, 0, 0, 435, 436, 5, 23, 0, 0, 436, 438, 5, 38, 0, 0, 437, 435,
		1, 0, 0, 0, 437, 438, 1, 0, 0, 0, 438, 458, 1, 0, 0, 0, 439, 440, 5, 23,
		0, 0, 440, 443, 5, 34, 0, 0, 441, 442, 5, 23, 0, 0, 442, 444, 5, 37, 0,
		0, 443, 441, 1, 0, 0, 0, 443, 444, 1, 0, 0, 0, 444, 447, 1, 0, 0, 0, 445,
		446, 5, 23, 0, 0, 446, 448, 5, 38, 0, 0, 447, 445, 1, 0, 0, 0, 447, 448,
		1, 0, 0, 0, 448, 458, 1, 0, 0, 0, 449, 450, 5, 23, 0, 0, 450, 453, 5, 37,
		0, 0, 451, 452, 5, 23, 0, 0, 452, 454, 5, 38, 0, 0, 453, 451, 1, 0, 0,
		0, 453, 454, 1, 0, 0, 0, 454, 458, 1, 0, 0, 0, 455, 456, 5, 23, 0, 0, 456,
		458, 5, 38, 0, 0, 457, 425, 1, 0, 0, 0, 457, 439, 1, 0, 0, 0, 457, 449,
		1, 0, 0, 0, 457, 455, 1, 0, 0, 0, 458, 79, 1, 0, 0, 0, 459, 462, 3, 82,
		41, 0, 460, 462, 3, 112, 56, 0, 461, 459, 1, 0, 0, 0, 461, 460, 1, 0, 0,
		0, 462, 81, 1, 0, 0, 0, 463, 482, 3, 84, 42, 0, 464, 482, 5, 52, 0, 0,
		465, 482, 5, 53, 0, 0, 466, 482, 5, 54, 0, 0, 467, 482, 5, 55, 0, 0, 468,
		482, 5, 56, 0, 0, 469, 482, 5, 57, 0, 0, 470, 482, 3, 86, 43, 0, 471, 482,
		3, 88, 44, 0, 472, 482, 3, 90, 45, 0, 473, 482, 3, 92, 46, 0, 474, 482,
		5, 62, 0, 0, 475, 482, 5, 63, 0, 0, 476, 482, 3, 96, 48, 0, 477, 482, 3,
		94, 47, 0, 478, 482, 5, 66, 0, 0, 479, 480, 5, 77, 0, 0, 480, 482, 5, 119,
		0, 0, 481, 463, 1, 0, 0, 0, 481, 464, 1, 0, 0, 0, 481, 465, 1, 0, 0, 0,
		481, 466, 1, 0, 0, 0, 481, 467, 1, 0, 0, 0, 481, 468, 1, 0, 0, 0, 481,
		469, 1, 0, 0, 0, 481, 470, 1, 0, 0, 0, 481, 471, 1, 0, 0, 0, 481, 472,
		1, 0, 0, 0, 481, 473, 1, 0, 0, 0, 481, 474, 1, 0, 0, 0, 481, 475, 1, 0,
		0, 0, 481, 476, 1, 0, 0, 0, 481, 477, 1, 0, 0, 0, 481, 478, 1, 0, 0, 0,
		481, 479, 1, 0, 0, 0, 482, 83, 1, 0, 0, 0, 483, 484, 7, 4, 0, 0, 484, 85,
		1, 0, 0, 0, 485, 486, 7, 5, 0, 0, 486, 87, 1, 0, 0, 0, 487, 488, 7, 6,
		0, 0, 488, 89, 1, 0, 0, 0, 489, 490, 7, 7, 0, 0, 490, 91, 1, 0, 0, 0, 491,
		492, 7, 8, 0, 0, 492, 93, 1, 0, 0, 0, 493, 494, 7, 9, 0, 0, 494, 95, 1,
		0, 0, 0, 495, 496, 7, 10, 0, 0, 496, 97, 1, 0, 0, 0, 497, 499, 7, 11, 0,
		0, 498, 500, 5, 112, 0, 0, 499, 498, 1, 0, 0, 0, 499, 500, 1, 0, 0, 0,
		500, 501, 1, 0, 0, 0, 501, 502, 5, 39, 0, 0, 502, 503, 3, 114, 57, 0, 503,
		504, 5, 40, 0, 0, 504, 99, 1, 0, 0, 0, 505, 507, 7, 12, 0, 0, 506, 508,
		5, 112, 0, 0, 507, 506, 1, 0, 0, 0, 507, 508, 1, 0, 0, 0, 508, 509, 1,
		0, 0, 0, 509, 510, 5, 39, 0, 0, 510, 511, 3, 114, 57, 0, 511, 512, 5, 40,
		0, 0, 512, 101, 1, 0, 0, 0, 513, 515, 7, 13, 0, 0, 514, 516, 5, 112, 0,
		0, 515, 514, 1, 0, 0, 0, 515, 516, 1, 0, 0, 0, 516, 517, 1, 0, 0, 0, 517,
		518, 5, 39, 0, 0, 518, 519, 3, 114, 57, 0, 519, 520, 5, 40, 0, 0, 520,
		103, 1, 0, 0, 0, 521, 523, 7, 14, 0, 0, 522, 524, 5, 112, 0, 0, 523, 522,
		1, 0, 0, 0, 523, 524, 1, 0, 0, 0, 524, 531, 1, 0, 0, 0, 525, 526, 5, 39,
		0, 0, 526, 527, 3, 114, 57, 0, 527, 528, 5, 110, 0, 0, 528, 529, 3, 114,
		57, 0, 529, 530, 5, 40, 0, 0, 530, 532, 1, 0, 0, 0, 531, 525, 1, 0, 0,
		0, 531, 532, 1, 0, 0, 0, 532, 105, 1, 0, 0, 0, 533, 535, 7, 15, 0, 0, 534,
		536, 5, 112, 0, 0, 535, 534, 1, 0, 0, 0, 535, 536, 1, 0, 0, 0, 536, 537,
		1, 0, 0, 0, 537, 538, 5, 39, 0, 0, 538, 539, 3, 114, 57, 0, 539, 540, 5,
		40, 0, 0, 540, 107, 1, 0, 0, 0, 541, 543, 7, 16, 0, 0, 542, 544, 5, 112,
		0, 0, 543, 542, 1, 0, 0, 0, 543, 544, 1, 0, 0, 0, 544, 545, 1, 0, 0, 0,
		545, 546, 5, 39, 0, 0, 546, 547, 3, 114, 57, 0, 547, 548, 5, 40, 0, 0,
		548, 109, 1, 0, 0, 0, 549, 551, 5, 75, 0, 0, 550, 552, 5, 112, 0, 0, 551,
		550, 1, 0, 0, 0, 551, 552, 1, 0, 0, 0, 552, 553, 1, 0, 0, 0, 553, 554,
		5, 39, 0, 0, 554, 555, 3, 80, 40, 0, 555, 556, 5, 40, 0, 0, 556, 111, 1,
		0, 0, 0, 557, 564, 3, 98, 49, 0, 558, 564, 3, 100, 50, 0, 559, 564, 3,
		102, 51, 0, 560, 564, 3, 104, 52, 0, 561, 564, 3, 106, 53, 0, 562, 564,
		3, 108, 54, 0, 563, 557, 1, 0, 0, 0, 563, 558, 1, 0, 0, 0, 563, 559, 1,
		0, 0, 0, 563, 560, 1, 0, 0, 0, 563, 561, 1, 0, 0, 0, 563, 562, 1, 0, 0,
		0, 564, 113, 1, 0, 0, 0, 565, 566, 5, 23, 0, 0, 566, 115, 1, 0, 0, 0, 567,
		568, 7, 17, 0, 0, 568, 117, 1, 0, 0, 0, 569, 570, 3, 120, 60, 0, 570, 571,
		5, 111, 0, 0, 571, 572, 3, 122, 61, 0, 572, 119, 1, 0, 0, 0, 573, 574,
		7, 18, 0, 0, 574, 121, 1, 0, 0, 0, 575, 576, 7, 19, 0, 0, 576, 123, 1,
		0, 0, 0, 577, 582, 3, 118, 59, 0, 578, 579, 5, 110, 0, 0, 579, 581, 3,
		118, 59, 0, 580, 578, 1, 0, 0, 0, 581, 584, 1, 0, 0, 0, 582, 580, 1, 0,
		0, 0, 582, 583, 1, 0, 0, 0, 583, 125, 1, 0, 0, 0, 584, 582, 1, 0, 0, 0,
		585, 586, 7, 20, 0, 0, 586, 127, 1, 0, 0, 0, 587, 590, 3, 126, 63, 0, 588,
		590, 5, 119, 0, 0, 589, 587, 1, 0, 0, 0, 589, 588, 1, 0, 0, 0, 590, 129,
		1, 0, 0, 0, 48, 134, 154, 167, 176, 182, 184, 191, 196, 211, 218, 240,
		250, 263, 266, 280, 283, 297, 304, 312, 319, 325, 330, 392, 395, 405, 410,
		418, 423, 429, 433, 437, 443, 447, 453, 457, 461, 481, 499, 507, 515, 523,
		531, 535, 543, 551, 563, 582, 589,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// FuncTestCaseParserInit initializes any static state used to implement FuncTestCaseParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFuncTestCaseParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FuncTestCaseParserInit() {
	staticData := &FuncTestCaseParserParserStaticData
	staticData.once.Do(functestcaseparserParserInit)
}

// NewFuncTestCaseParser produces a new parser instance for the optional input antlr.TokenStream.
func NewFuncTestCaseParser(input antlr.TokenStream) *FuncTestCaseParser {
	FuncTestCaseParserInit()
	this := new(FuncTestCaseParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &FuncTestCaseParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "FuncTestCaseParser.g4"

	return this
}

// FuncTestCaseParser tokens.
const (
	FuncTestCaseParserEOF                    = antlr.TokenEOF
	FuncTestCaseParserWhitespace             = 1
	FuncTestCaseParserTripleHash             = 2
	FuncTestCaseParserSubstraitScalarTest    = 3
	FuncTestCaseParserSubstraitAggregateTest = 4
	FuncTestCaseParserSubstraitInclude       = 5
	FuncTestCaseParserFormatVersion          = 6
	FuncTestCaseParserDescriptionLine        = 7
	FuncTestCaseParserDefine                 = 8
	FuncTestCaseParserErrorResult            = 9
	FuncTestCaseParserUndefineResult         = 10
	FuncTestCaseParserOverflow               = 11
	FuncTestCaseParserRounding               = 12
	FuncTestCaseParserError                  = 13
	FuncTestCaseParserSaturate               = 14
	FuncTestCaseParserSilent                 = 15
	FuncTestCaseParserTieToEven              = 16
	FuncTestCaseParserNaN                    = 17
	FuncTestCaseParserAcceptNulls            = 18
	FuncTestCaseParserIgnoreNulls            = 19
	FuncTestCaseParserNullHandling           = 20
	FuncTestCaseParserSpacesOnly             = 21
	FuncTestCaseParserTruncate               = 22
	FuncTestCaseParserIntegerLiteral         = 23
	FuncTestCaseParserDecimalLiteral         = 24
	FuncTestCaseParserFloatLiteral           = 25
	FuncTestCaseParserBooleanLiteral         = 26
	FuncTestCaseParserTimestampTzLiteral     = 27
	FuncTestCaseParserTimestampLiteral       = 28
	FuncTestCaseParserTimeLiteral            = 29
	FuncTestCaseParserDateLiteral            = 30
	FuncTestCaseParserPeriodPrefix           = 31
	FuncTestCaseParserTimePrefix             = 32
	FuncTestCaseParserYearPrefix             = 33
	FuncTestCaseParserMSuffix                = 34
	FuncTestCaseParserDaySuffix              = 35
	FuncTestCaseParserHourSuffix             = 36
	FuncTestCaseParserSecondSuffix           = 37
	FuncTestCaseParserFractionalSecondSuffix = 38
	FuncTestCaseParserOAngleBracket          = 39
	FuncTestCaseParserCAngleBracket          = 40
	FuncTestCaseParserIntervalYearLiteral    = 41
	FuncTestCaseParserIntervalDayLiteral     = 42
	FuncTestCaseParserNullLiteral            = 43
	FuncTestCaseParserStringLiteral          = 44
	FuncTestCaseParserColumnName             = 45
	FuncTestCaseParserLineComment            = 46
	FuncTestCaseParserBlockComment           = 47
	FuncTestCaseParserIf                     = 48
	FuncTestCaseParserThen                   = 49
	FuncTestCaseParserElse                   = 50
	FuncTestCaseParserBoolean                = 51
	FuncTestCaseParserI8                     = 52
	FuncTestCaseParserI16                    = 53
	FuncTestCaseParserI32                    = 54
	FuncTestCaseParserI64                    = 55
	FuncTestCaseParserFP32                   = 56
	FuncTestCaseParserFP64                   = 57
	FuncTestCaseParserString_                = 58
	FuncTestCaseParserBinary                 = 59
	FuncTestCaseParserTimestamp              = 60
	FuncTestCaseParserTimestamp_TZ           = 61
	FuncTestCaseParserDate                   = 62
	FuncTestCaseParserTime                   = 63
	FuncTestCaseParserInterval_Year          = 64
	FuncTestCaseParserInterval_Day           = 65
	FuncTestCaseParserUUID                   = 66
	FuncTestCaseParserDecimal                = 67
	FuncTestCaseParserPrecision_Timestamp    = 68
	FuncTestCaseParserPrecision_Timestamp_TZ = 69
	FuncTestCaseParserFixedChar              = 70
	FuncTestCaseParserVarChar                = 71
	FuncTestCaseParserFixedBinary            = 72
	FuncTestCaseParserStruct                 = 73
	FuncTestCaseParserNStruct                = 74
	FuncTestCaseParserList                   = 75
	FuncTestCaseParserMap                    = 76
	FuncTestCaseParserUserDefined            = 77
	FuncTestCaseParserBool                   = 78
	FuncTestCaseParserStr                    = 79
	FuncTestCaseParserVBin                   = 80
	FuncTestCaseParserTs                     = 81
	FuncTestCaseParserTsTZ                   = 82
	FuncTestCaseParserIYear                  = 83
	FuncTestCaseParserIDay                   = 84
	FuncTestCaseParserDec                    = 85
	FuncTestCaseParserPTs                    = 86
	FuncTestCaseParserPTsTZ                  = 87
	FuncTestCaseParserFChar                  = 88
	FuncTestCaseParserVChar                  = 89
	FuncTestCaseParserFBin                   = 90
	FuncTestCaseParserAny                    = 91
	FuncTestCaseParserAnyVar                 = 92
	FuncTestCaseParserDoubleColon            = 93
	FuncTestCaseParserPlus                   = 94
	FuncTestCaseParserMinus                  = 95
	FuncTestCaseParserAsterisk               = 96
	FuncTestCaseParserForwardSlash           = 97
	FuncTestCaseParserPercent                = 98
	FuncTestCaseParserEq                     = 99
	FuncTestCaseParserNe                     = 100
	FuncTestCaseParserGte                    = 101
	FuncTestCaseParserLte                    = 102
	FuncTestCaseParserGt                     = 103
	FuncTestCaseParserLt                     = 104
	FuncTestCaseParserBang                   = 105
	FuncTestCaseParserOParen                 = 106
	FuncTestCaseParserCParen                 = 107
	FuncTestCaseParserOBracket               = 108
	FuncTestCaseParserCBracket               = 109
	FuncTestCaseParserComma                  = 110
	FuncTestCaseParserColon                  = 111
	FuncTestCaseParserQMark                  = 112
	FuncTestCaseParserHash                   = 113
	FuncTestCaseParserDot                    = 114
	FuncTestCaseParserAnd                    = 115
	FuncTestCaseParserOr                     = 116
	FuncTestCaseParserAssign                 = 117
	FuncTestCaseParserNumber                 = 118
	FuncTestCaseParserIdentifier             = 119
	FuncTestCaseParserNewline                = 120
)

// FuncTestCaseParser rules.
const (
	FuncTestCaseParserRULE_doc                        = 0
	FuncTestCaseParserRULE_header                     = 1
	FuncTestCaseParserRULE_version                    = 2
	FuncTestCaseParserRULE_include                    = 3
	FuncTestCaseParserRULE_testGroupDescription       = 4
	FuncTestCaseParserRULE_testCase                   = 5
	FuncTestCaseParserRULE_testGroup                  = 6
	FuncTestCaseParserRULE_arguments                  = 7
	FuncTestCaseParserRULE_result                     = 8
	FuncTestCaseParserRULE_argument                   = 9
	FuncTestCaseParserRULE_aggFuncTestCase            = 10
	FuncTestCaseParserRULE_aggFuncCall                = 11
	FuncTestCaseParserRULE_tableData                  = 12
	FuncTestCaseParserRULE_tableRows                  = 13
	FuncTestCaseParserRULE_dataColumn                 = 14
	FuncTestCaseParserRULE_columnValues               = 15
	FuncTestCaseParserRULE_literal                    = 16
	FuncTestCaseParserRULE_qualifiedAggregateFuncArgs = 17
	FuncTestCaseParserRULE_aggregateFuncArgs          = 18
	FuncTestCaseParserRULE_qualifiedAggregateFuncArg  = 19
	FuncTestCaseParserRULE_aggregateFuncArg           = 20
	FuncTestCaseParserRULE_numericLiteral             = 21
	FuncTestCaseParserRULE_floatLiteral               = 22
	FuncTestCaseParserRULE_nullArg                    = 23
	FuncTestCaseParserRULE_intArg                     = 24
	FuncTestCaseParserRULE_floatArg                   = 25
	FuncTestCaseParserRULE_decimalArg                 = 26
	FuncTestCaseParserRULE_booleanArg                 = 27
	FuncTestCaseParserRULE_stringArg                  = 28
	FuncTestCaseParserRULE_dateArg                    = 29
	FuncTestCaseParserRULE_timeArg                    = 30
	FuncTestCaseParserRULE_timestampArg               = 31
	FuncTestCaseParserRULE_timestampTzArg             = 32
	FuncTestCaseParserRULE_intervalYearArg            = 33
	FuncTestCaseParserRULE_intervalDayArg             = 34
	FuncTestCaseParserRULE_listArg                    = 35
	FuncTestCaseParserRULE_literalList                = 36
	FuncTestCaseParserRULE_intervalYearLiteral        = 37
	FuncTestCaseParserRULE_intervalDayLiteral         = 38
	FuncTestCaseParserRULE_timeInterval               = 39
	FuncTestCaseParserRULE_dataType                   = 40
	FuncTestCaseParserRULE_scalarType                 = 41
	FuncTestCaseParserRULE_booleanType                = 42
	FuncTestCaseParserRULE_stringType                 = 43
	FuncTestCaseParserRULE_binaryType                 = 44
	FuncTestCaseParserRULE_timestampType              = 45
	FuncTestCaseParserRULE_timestampTZType            = 46
	FuncTestCaseParserRULE_intervalYearType           = 47
	FuncTestCaseParserRULE_intervalDayType            = 48
	FuncTestCaseParserRULE_fixedCharType              = 49
	FuncTestCaseParserRULE_varCharType                = 50
	FuncTestCaseParserRULE_fixedBinaryType            = 51
	FuncTestCaseParserRULE_decimalType                = 52
	FuncTestCaseParserRULE_precisionTimestampType     = 53
	FuncTestCaseParserRULE_precisionTimestampTZType   = 54
	FuncTestCaseParserRULE_listType                   = 55
	FuncTestCaseParserRULE_parameterizedType          = 56
	FuncTestCaseParserRULE_numericParameter           = 57
	FuncTestCaseParserRULE_substraitError             = 58
	FuncTestCaseParserRULE_func_option                = 59
	FuncTestCaseParserRULE_option_name                = 60
	FuncTestCaseParserRULE_option_value               = 61
	FuncTestCaseParserRULE_func_options               = 62
	FuncTestCaseParserRULE_nonReserved                = 63
	FuncTestCaseParserRULE_identifier                 = 64
)

// IDocContext is an interface to support dynamic dispatch.
type IDocContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Header() IHeaderContext
	EOF() antlr.TerminalNode
	AllTestGroup() []ITestGroupContext
	TestGroup(i int) ITestGroupContext

	// IsDocContext differentiates from other interfaces.
	IsDocContext()
}

type DocContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocContext() *DocContext {
	var p = new(DocContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_doc
	return p
}

func InitEmptyDocContext(p *DocContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_doc
}

func (*DocContext) IsDocContext() {}

func NewDocContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocContext {
	var p = new(DocContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_doc

	return p
}

func (s *DocContext) GetParser() antlr.Parser { return s.parser }

func (s *DocContext) Header() IHeaderContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHeaderContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHeaderContext)
}

func (s *DocContext) EOF() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserEOF, 0)
}

func (s *DocContext) AllTestGroup() []ITestGroupContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITestGroupContext); ok {
			len++
		}
	}

	tst := make([]ITestGroupContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITestGroupContext); ok {
			tst[i] = t.(ITestGroupContext)
			i++
		}
	}

	return tst
}

func (s *DocContext) TestGroup(i int) ITestGroupContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITestGroupContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITestGroupContext)
}

func (s *DocContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterDoc(s)
	}
}

func (s *DocContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitDoc(s)
	}
}

func (s *DocContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitDoc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) Doc() (localctx IDocContext) {
	localctx = NewDocContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FuncTestCaseParserRULE_doc)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)
		p.Header()
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == FuncTestCaseParserDescriptionLine {
		{
			p.SetState(131)
			p.TestGroup()
		}

		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(136)
		p.Match(FuncTestCaseParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHeaderContext is an interface to support dynamic dispatch.
type IHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Version() IVersionContext
	Include() IIncludeContext

	// IsHeaderContext differentiates from other interfaces.
	IsHeaderContext()
}

type HeaderContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeaderContext() *HeaderContext {
	var p = new(HeaderContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_header
	return p
}

func InitEmptyHeaderContext(p *HeaderContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_header
}

func (*HeaderContext) IsHeaderContext() {}

func NewHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeaderContext {
	var p = new(HeaderContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_header

	return p
}

func (s *HeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *HeaderContext) Version() IVersionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVersionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVersionContext)
}

func (s *HeaderContext) Include() IIncludeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIncludeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIncludeContext)
}

func (s *HeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HeaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterHeader(s)
	}
}

func (s *HeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitHeader(s)
	}
}

func (s *HeaderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitHeader(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) Header() (localctx IHeaderContext) {
	localctx = NewHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FuncTestCaseParserRULE_header)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Version()
	}
	{
		p.SetState(139)
		p.Include()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVersionContext is an interface to support dynamic dispatch.
type IVersionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TripleHash() antlr.TerminalNode
	Colon() antlr.TerminalNode
	FormatVersion() antlr.TerminalNode
	SubstraitScalarTest() antlr.TerminalNode
	SubstraitAggregateTest() antlr.TerminalNode

	// IsVersionContext differentiates from other interfaces.
	IsVersionContext()
}

type VersionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVersionContext() *VersionContext {
	var p = new(VersionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_version
	return p
}

func InitEmptyVersionContext(p *VersionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_version
}

func (*VersionContext) IsVersionContext() {}

func NewVersionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VersionContext {
	var p = new(VersionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_version

	return p
}

func (s *VersionContext) GetParser() antlr.Parser { return s.parser }

func (s *VersionContext) TripleHash() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserTripleHash, 0)
}

func (s *VersionContext) Colon() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserColon, 0)
}

func (s *VersionContext) FormatVersion() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserFormatVersion, 0)
}

func (s *VersionContext) SubstraitScalarTest() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserSubstraitScalarTest, 0)
}

func (s *VersionContext) SubstraitAggregateTest() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserSubstraitAggregateTest, 0)
}

func (s *VersionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VersionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VersionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterVersion(s)
	}
}

func (s *VersionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitVersion(s)
	}
}

func (s *VersionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitVersion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) Version() (localctx IVersionContext) {
	localctx = NewVersionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FuncTestCaseParserRULE_version)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(FuncTestCaseParserTripleHash)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(142)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserSubstraitScalarTest || _la == FuncTestCaseParserSubstraitAggregateTest) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(143)
		p.Match(FuncTestCaseParserColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.Match(FuncTestCaseParserFormatVersion)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIncludeContext is an interface to support dynamic dispatch.
type IIncludeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TripleHash() antlr.TerminalNode
	SubstraitInclude() antlr.TerminalNode
	Colon() antlr.TerminalNode
	AllStringLiteral() []antlr.TerminalNode
	StringLiteral(i int) antlr.TerminalNode
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsIncludeContext differentiates from other interfaces.
	IsIncludeContext()
}

type IncludeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncludeContext() *IncludeContext {
	var p = new(IncludeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_include
	return p
}

func InitEmptyIncludeContext(p *IncludeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_include
}

func (*IncludeContext) IsIncludeContext() {}

func NewIncludeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncludeContext {
	var p = new(IncludeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_include

	return p
}

func (s *IncludeContext) GetParser() antlr.Parser { return s.parser }

func (s *IncludeContext) TripleHash() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserTripleHash, 0)
}

func (s *IncludeContext) SubstraitInclude() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserSubstraitInclude, 0)
}

func (s *IncludeContext) Colon() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserColon, 0)
}

func (s *IncludeContext) AllStringLiteral() []antlr.TerminalNode {
	return s.GetTokens(FuncTestCaseParserStringLiteral)
}

func (s *IncludeContext) StringLiteral(i int) antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserStringLiteral, i)
}

func (s *IncludeContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(FuncTestCaseParserComma)
}

func (s *IncludeContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserComma, i)
}

func (s *IncludeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncludeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncludeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterInclude(s)
	}
}

func (s *IncludeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitInclude(s)
	}
}

func (s *IncludeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitInclude(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) Include() (localctx IIncludeContext) {
	localctx = NewIncludeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FuncTestCaseParserRULE_include)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(146)
		p.Match(FuncTestCaseParserTripleHash)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(147)
		p.Match(FuncTestCaseParserSubstraitInclude)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(148)
		p.Match(FuncTestCaseParserColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)
		p.Match(FuncTestCaseParserStringLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FuncTestCaseParserComma {
		{
			p.SetState(150)
			p.Match(FuncTestCaseParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(151)
			p.Match(FuncTestCaseParserStringLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITestGroupDescriptionContext is an interface to support dynamic dispatch.
type ITestGroupDescriptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DescriptionLine() antlr.TerminalNode

	// IsTestGroupDescriptionContext differentiates from other interfaces.
	IsTestGroupDescriptionContext()
}

type TestGroupDescriptionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTestGroupDescriptionContext() *TestGroupDescriptionContext {
	var p = new(TestGroupDescriptionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_testGroupDescription
	return p
}

func InitEmptyTestGroupDescriptionContext(p *TestGroupDescriptionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_testGroupDescription
}

func (*TestGroupDescriptionContext) IsTestGroupDescriptionContext() {}

func NewTestGroupDescriptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TestGroupDescriptionContext {
	var p = new(TestGroupDescriptionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_testGroupDescription

	return p
}

func (s *TestGroupDescriptionContext) GetParser() antlr.Parser { return s.parser }

func (s *TestGroupDescriptionContext) DescriptionLine() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDescriptionLine, 0)
}

func (s *TestGroupDescriptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TestGroupDescriptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TestGroupDescriptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterTestGroupDescription(s)
	}
}

func (s *TestGroupDescriptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitTestGroupDescription(s)
	}
}

func (s *TestGroupDescriptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitTestGroupDescription(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) TestGroupDescription() (localctx ITestGroupDescriptionContext) {
	localctx = NewTestGroupDescriptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FuncTestCaseParserRULE_testGroupDescription)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(FuncTestCaseParserDescriptionLine)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITestCaseContext is an interface to support dynamic dispatch.
type ITestCaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFunctionName returns the functionName rule contexts.
	GetFunctionName() IIdentifierContext

	// SetFunctionName sets the functionName rule contexts.
	SetFunctionName(IIdentifierContext)

	// Getter signatures
	OParen() antlr.TerminalNode
	Arguments() IArgumentsContext
	CParen() antlr.TerminalNode
	Eq() antlr.TerminalNode
	Result() IResultContext
	Identifier() IIdentifierContext
	OBracket() antlr.TerminalNode
	Func_options() IFunc_optionsContext
	CBracket() antlr.TerminalNode

	// IsTestCaseContext differentiates from other interfaces.
	IsTestCaseContext()
}

type TestCaseContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	functionName IIdentifierContext
}

func NewEmptyTestCaseContext() *TestCaseContext {
	var p = new(TestCaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_testCase
	return p
}

func InitEmptyTestCaseContext(p *TestCaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_testCase
}

func (*TestCaseContext) IsTestCaseContext() {}

func NewTestCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TestCaseContext {
	var p = new(TestCaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_testCase

	return p
}

func (s *TestCaseContext) GetParser() antlr.Parser { return s.parser }

func (s *TestCaseContext) GetFunctionName() IIdentifierContext { return s.functionName }

func (s *TestCaseContext) SetFunctionName(v IIdentifierContext) { s.functionName = v }

func (s *TestCaseContext) OParen() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserOParen, 0)
}

func (s *TestCaseContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *TestCaseContext) CParen() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserCParen, 0)
}

func (s *TestCaseContext) Eq() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserEq, 0)
}

func (s *TestCaseContext) Result() IResultContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResultContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResultContext)
}

func (s *TestCaseContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *TestCaseContext) OBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserOBracket, 0)
}

func (s *TestCaseContext) Func_options() IFunc_optionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_optionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_optionsContext)
}

func (s *TestCaseContext) CBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserCBracket, 0)
}

func (s *TestCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TestCaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TestCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterTestCase(s)
	}
}

func (s *TestCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitTestCase(s)
	}
}

func (s *TestCaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitTestCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) TestCase() (localctx ITestCaseContext) {
	localctx = NewTestCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FuncTestCaseParserRULE_testCase)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)

		var _x = p.Identifier()

		localctx.(*TestCaseContext).functionName = _x
	}
	{
		p.SetState(160)
		p.Match(FuncTestCaseParserOParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(161)
		p.Arguments()
	}
	{
		p.SetState(162)
		p.Match(FuncTestCaseParserCParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserOBracket {
		{
			p.SetState(163)
			p.Match(FuncTestCaseParserOBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(164)
			p.Func_options()
		}
		{
			p.SetState(165)
			p.Match(FuncTestCaseParserCBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(169)
		p.Match(FuncTestCaseParserEq)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.Result()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITestGroupContext is an interface to support dynamic dispatch.
type ITestGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTestGroupContext differentiates from other interfaces.
	IsTestGroupContext()
}

type TestGroupContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTestGroupContext() *TestGroupContext {
	var p = new(TestGroupContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_testGroup
	return p
}

func InitEmptyTestGroupContext(p *TestGroupContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_testGroup
}

func (*TestGroupContext) IsTestGroupContext() {}

func NewTestGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TestGroupContext {
	var p = new(TestGroupContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_testGroup

	return p
}

func (s *TestGroupContext) GetParser() antlr.Parser { return s.parser }

func (s *TestGroupContext) CopyAll(ctx *TestGroupContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TestGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TestGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ScalarFuncTestGroupContext struct {
	TestGroupContext
}

func NewScalarFuncTestGroupContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ScalarFuncTestGroupContext {
	var p = new(ScalarFuncTestGroupContext)

	InitEmptyTestGroupContext(&p.TestGroupContext)
	p.parser = parser
	p.CopyAll(ctx.(*TestGroupContext))

	return p
}

func (s *ScalarFuncTestGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarFuncTestGroupContext) TestGroupDescription() ITestGroupDescriptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITestGroupDescriptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITestGroupDescriptionContext)
}

func (s *ScalarFuncTestGroupContext) AllTestCase() []ITestCaseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITestCaseContext); ok {
			len++
		}
	}

	tst := make([]ITestCaseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITestCaseContext); ok {
			tst[i] = t.(ITestCaseContext)
			i++
		}
	}

	return tst
}

func (s *ScalarFuncTestGroupContext) TestCase(i int) ITestCaseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITestCaseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITestCaseContext)
}

func (s *ScalarFuncTestGroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterScalarFuncTestGroup(s)
	}
}

func (s *ScalarFuncTestGroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitScalarFuncTestGroup(s)
	}
}

func (s *ScalarFuncTestGroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitScalarFuncTestGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

type AggregateFuncTestGroupContext struct {
	TestGroupContext
}

func NewAggregateFuncTestGroupContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AggregateFuncTestGroupContext {
	var p = new(AggregateFuncTestGroupContext)

	InitEmptyTestGroupContext(&p.TestGroupContext)
	p.parser = parser
	p.CopyAll(ctx.(*TestGroupContext))

	return p
}

func (s *AggregateFuncTestGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggregateFuncTestGroupContext) TestGroupDescription() ITestGroupDescriptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITestGroupDescriptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITestGroupDescriptionContext)
}

func (s *AggregateFuncTestGroupContext) AllAggFuncTestCase() []IAggFuncTestCaseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAggFuncTestCaseContext); ok {
			len++
		}
	}

	tst := make([]IAggFuncTestCaseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAggFuncTestCaseContext); ok {
			tst[i] = t.(IAggFuncTestCaseContext)
			i++
		}
	}

	return tst
}

func (s *AggregateFuncTestGroupContext) AggFuncTestCase(i int) IAggFuncTestCaseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAggFuncTestCaseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAggFuncTestCaseContext)
}

func (s *AggregateFuncTestGroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterAggregateFuncTestGroup(s)
	}
}

func (s *AggregateFuncTestGroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitAggregateFuncTestGroup(s)
	}
}

func (s *AggregateFuncTestGroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitAggregateFuncTestGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) TestGroup() (localctx ITestGroupContext) {
	localctx = NewTestGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FuncTestCaseParserRULE_testGroup)
	var _la int

	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewScalarFuncTestGroupContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(172)
			p.TestGroupDescription()
		}
		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == FuncTestCaseParserTruncate || ((int64((_la-115)) & ^0x3f) == 0 && ((int64(1)<<(_la-115))&19) != 0) {
			{
				p.SetState(173)
				p.TestCase()
			}

			p.SetState(176)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		localctx = NewAggregateFuncTestGroupContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(178)
			p.TestGroupDescription()
		}
		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == FuncTestCaseParserDefine || _la == FuncTestCaseParserTruncate || ((int64((_la-106)) & ^0x3f) == 0 && ((int64(1)<<(_la-106))&9729) != 0) {
			{
				p.SetState(179)
				p.AggFuncTestCase()
			}

			p.SetState(182)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllArgument() []IArgumentContext
	Argument(i int) IArgumentContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_arguments
	return p
}

func InitEmptyArgumentsContext(p *ArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_arguments
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) AllArgument() []IArgumentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgumentContext); ok {
			len++
		}
	}

	tst := make([]IArgumentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgumentContext); ok {
			tst[i] = t.(IArgumentContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentsContext) Argument(i int) IArgumentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *ArgumentsContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(FuncTestCaseParserComma)
}

func (s *ArgumentsContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserComma, i)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (s *ArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FuncTestCaseParserRULE_arguments)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)
		p.Argument()
	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FuncTestCaseParserComma {
		{
			p.SetState(187)
			p.Match(FuncTestCaseParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)
			p.Argument()
		}

		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IResultContext is an interface to support dynamic dispatch.
type IResultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Argument() IArgumentContext
	SubstraitError() ISubstraitErrorContext

	// IsResultContext differentiates from other interfaces.
	IsResultContext()
}

type ResultContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResultContext() *ResultContext {
	var p = new(ResultContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_result
	return p
}

func InitEmptyResultContext(p *ResultContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_result
}

func (*ResultContext) IsResultContext() {}

func NewResultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResultContext {
	var p = new(ResultContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_result

	return p
}

func (s *ResultContext) GetParser() antlr.Parser { return s.parser }

func (s *ResultContext) Argument() IArgumentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *ResultContext) SubstraitError() ISubstraitErrorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubstraitErrorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubstraitErrorContext)
}

func (s *ResultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterResult(s)
	}
}

func (s *ResultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitResult(s)
	}
}

func (s *ResultContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitResult(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) Result() (localctx IResultContext) {
	localctx = NewResultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FuncTestCaseParserRULE_result)
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FuncTestCaseParserNaN, FuncTestCaseParserIntegerLiteral, FuncTestCaseParserDecimalLiteral, FuncTestCaseParserFloatLiteral, FuncTestCaseParserBooleanLiteral, FuncTestCaseParserTimestampTzLiteral, FuncTestCaseParserTimestampLiteral, FuncTestCaseParserTimeLiteral, FuncTestCaseParserDateLiteral, FuncTestCaseParserIntervalYearLiteral, FuncTestCaseParserIntervalDayLiteral, FuncTestCaseParserNullLiteral, FuncTestCaseParserStringLiteral, FuncTestCaseParserOBracket:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(194)
			p.Argument()
		}

	case FuncTestCaseParserErrorResult, FuncTestCaseParserUndefineResult:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(195)
			p.SubstraitError()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NullArg() INullArgContext
	IntArg() IIntArgContext
	FloatArg() IFloatArgContext
	BooleanArg() IBooleanArgContext
	StringArg() IStringArgContext
	DecimalArg() IDecimalArgContext
	DateArg() IDateArgContext
	TimeArg() ITimeArgContext
	TimestampArg() ITimestampArgContext
	TimestampTzArg() ITimestampTzArgContext
	IntervalYearArg() IIntervalYearArgContext
	IntervalDayArg() IIntervalDayArgContext
	ListArg() IListArgContext

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_argument
	return p
}

func InitEmptyArgumentContext(p *ArgumentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_argument
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) NullArg() INullArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INullArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INullArgContext)
}

func (s *ArgumentContext) IntArg() IIntArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntArgContext)
}

func (s *ArgumentContext) FloatArg() IFloatArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloatArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloatArgContext)
}

func (s *ArgumentContext) BooleanArg() IBooleanArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanArgContext)
}

func (s *ArgumentContext) StringArg() IStringArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringArgContext)
}

func (s *ArgumentContext) DecimalArg() IDecimalArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecimalArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecimalArgContext)
}

func (s *ArgumentContext) DateArg() IDateArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDateArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDateArgContext)
}

func (s *ArgumentContext) TimeArg() ITimeArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimeArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimeArgContext)
}

func (s *ArgumentContext) TimestampArg() ITimestampArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimestampArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimestampArgContext)
}

func (s *ArgumentContext) TimestampTzArg() ITimestampTzArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimestampTzArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimestampTzArgContext)
}

func (s *ArgumentContext) IntervalYearArg() IIntervalYearArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntervalYearArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntervalYearArgContext)
}

func (s *ArgumentContext) IntervalDayArg() IIntervalDayArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntervalDayArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntervalDayArgContext)
}

func (s *ArgumentContext) ListArg() IListArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListArgContext)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitArgument(s)
	}
}

func (s *ArgumentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitArgument(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) Argument() (localctx IArgumentContext) {
	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FuncTestCaseParserRULE_argument)
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(198)
			p.NullArg()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(199)
			p.IntArg()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(200)
			p.FloatArg()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(201)
			p.BooleanArg()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(202)
			p.StringArg()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(203)
			p.DecimalArg()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(204)
			p.DateArg()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(205)
			p.TimeArg()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(206)
			p.TimestampArg()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(207)
			p.TimestampTzArg()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(208)
			p.IntervalYearArg()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(209)
			p.IntervalDayArg()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(210)
			p.ListArg()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAggFuncTestCaseContext is an interface to support dynamic dispatch.
type IAggFuncTestCaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AggFuncCall() IAggFuncCallContext
	Eq() antlr.TerminalNode
	Result() IResultContext
	OBracket() antlr.TerminalNode
	Func_options() IFunc_optionsContext
	CBracket() antlr.TerminalNode

	// IsAggFuncTestCaseContext differentiates from other interfaces.
	IsAggFuncTestCaseContext()
}

type AggFuncTestCaseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAggFuncTestCaseContext() *AggFuncTestCaseContext {
	var p = new(AggFuncTestCaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_aggFuncTestCase
	return p
}

func InitEmptyAggFuncTestCaseContext(p *AggFuncTestCaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_aggFuncTestCase
}

func (*AggFuncTestCaseContext) IsAggFuncTestCaseContext() {}

func NewAggFuncTestCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggFuncTestCaseContext {
	var p = new(AggFuncTestCaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_aggFuncTestCase

	return p
}

func (s *AggFuncTestCaseContext) GetParser() antlr.Parser { return s.parser }

func (s *AggFuncTestCaseContext) AggFuncCall() IAggFuncCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAggFuncCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAggFuncCallContext)
}

func (s *AggFuncTestCaseContext) Eq() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserEq, 0)
}

func (s *AggFuncTestCaseContext) Result() IResultContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResultContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResultContext)
}

func (s *AggFuncTestCaseContext) OBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserOBracket, 0)
}

func (s *AggFuncTestCaseContext) Func_options() IFunc_optionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_optionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_optionsContext)
}

func (s *AggFuncTestCaseContext) CBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserCBracket, 0)
}

func (s *AggFuncTestCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggFuncTestCaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AggFuncTestCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterAggFuncTestCase(s)
	}
}

func (s *AggFuncTestCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitAggFuncTestCase(s)
	}
}

func (s *AggFuncTestCaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitAggFuncTestCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) AggFuncTestCase() (localctx IAggFuncTestCaseContext) {
	localctx = NewAggFuncTestCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, FuncTestCaseParserRULE_aggFuncTestCase)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		p.AggFuncCall()
	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserOBracket {
		{
			p.SetState(214)
			p.Match(FuncTestCaseParserOBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(215)
			p.Func_options()
		}
		{
			p.SetState(216)
			p.Match(FuncTestCaseParserCBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(220)
		p.Match(FuncTestCaseParserEq)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(221)
		p.Result()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAggFuncCallContext is an interface to support dynamic dispatch.
type IAggFuncCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAggFuncCallContext differentiates from other interfaces.
	IsAggFuncCallContext()
}

type AggFuncCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAggFuncCallContext() *AggFuncCallContext {
	var p = new(AggFuncCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_aggFuncCall
	return p
}

func InitEmptyAggFuncCallContext(p *AggFuncCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_aggFuncCall
}

func (*AggFuncCallContext) IsAggFuncCallContext() {}

func NewAggFuncCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggFuncCallContext {
	var p = new(AggFuncCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_aggFuncCall

	return p
}

func (s *AggFuncCallContext) GetParser() antlr.Parser { return s.parser }

func (s *AggFuncCallContext) CopyAll(ctx *AggFuncCallContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AggFuncCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggFuncCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SingleArgAggregateFuncCallContext struct {
	AggFuncCallContext
	functName IIdentifierContext
}

func NewSingleArgAggregateFuncCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleArgAggregateFuncCallContext {
	var p = new(SingleArgAggregateFuncCallContext)

	InitEmptyAggFuncCallContext(&p.AggFuncCallContext)
	p.parser = parser
	p.CopyAll(ctx.(*AggFuncCallContext))

	return p
}

func (s *SingleArgAggregateFuncCallContext) GetFunctName() IIdentifierContext { return s.functName }

func (s *SingleArgAggregateFuncCallContext) SetFunctName(v IIdentifierContext) { s.functName = v }

func (s *SingleArgAggregateFuncCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleArgAggregateFuncCallContext) OParen() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserOParen, 0)
}

func (s *SingleArgAggregateFuncCallContext) DataColumn() IDataColumnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataColumnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataColumnContext)
}

func (s *SingleArgAggregateFuncCallContext) CParen() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserCParen, 0)
}

func (s *SingleArgAggregateFuncCallContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *SingleArgAggregateFuncCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterSingleArgAggregateFuncCall(s)
	}
}

func (s *SingleArgAggregateFuncCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitSingleArgAggregateFuncCall(s)
	}
}

func (s *SingleArgAggregateFuncCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitSingleArgAggregateFuncCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type MultiArgAggregateFuncCallContext struct {
	AggFuncCallContext
	funcName IIdentifierContext
}

func NewMultiArgAggregateFuncCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiArgAggregateFuncCallContext {
	var p = new(MultiArgAggregateFuncCallContext)

	InitEmptyAggFuncCallContext(&p.AggFuncCallContext)
	p.parser = parser
	p.CopyAll(ctx.(*AggFuncCallContext))

	return p
}

func (s *MultiArgAggregateFuncCallContext) GetFuncName() IIdentifierContext { return s.funcName }

func (s *MultiArgAggregateFuncCallContext) SetFuncName(v IIdentifierContext) { s.funcName = v }

func (s *MultiArgAggregateFuncCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiArgAggregateFuncCallContext) TableData() ITableDataContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableDataContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableDataContext)
}

func (s *MultiArgAggregateFuncCallContext) OParen() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserOParen, 0)
}

func (s *MultiArgAggregateFuncCallContext) QualifiedAggregateFuncArgs() IQualifiedAggregateFuncArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQualifiedAggregateFuncArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQualifiedAggregateFuncArgsContext)
}

func (s *MultiArgAggregateFuncCallContext) CParen() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserCParen, 0)
}

func (s *MultiArgAggregateFuncCallContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *MultiArgAggregateFuncCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterMultiArgAggregateFuncCall(s)
	}
}

func (s *MultiArgAggregateFuncCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitMultiArgAggregateFuncCall(s)
	}
}

func (s *MultiArgAggregateFuncCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitMultiArgAggregateFuncCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type CompactAggregateFuncCallContext struct {
	AggFuncCallContext
	functName IIdentifierContext
}

func NewCompactAggregateFuncCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompactAggregateFuncCallContext {
	var p = new(CompactAggregateFuncCallContext)

	InitEmptyAggFuncCallContext(&p.AggFuncCallContext)
	p.parser = parser
	p.CopyAll(ctx.(*AggFuncCallContext))

	return p
}

func (s *CompactAggregateFuncCallContext) GetFunctName() IIdentifierContext { return s.functName }

func (s *CompactAggregateFuncCallContext) SetFunctName(v IIdentifierContext) { s.functName = v }

func (s *CompactAggregateFuncCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompactAggregateFuncCallContext) TableRows() ITableRowsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableRowsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableRowsContext)
}

func (s *CompactAggregateFuncCallContext) OParen() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserOParen, 0)
}

func (s *CompactAggregateFuncCallContext) AggregateFuncArgs() IAggregateFuncArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAggregateFuncArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAggregateFuncArgsContext)
}

func (s *CompactAggregateFuncCallContext) CParen() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserCParen, 0)
}

func (s *CompactAggregateFuncCallContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *CompactAggregateFuncCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterCompactAggregateFuncCall(s)
	}
}

func (s *CompactAggregateFuncCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitCompactAggregateFuncCall(s)
	}
}

func (s *CompactAggregateFuncCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitCompactAggregateFuncCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) AggFuncCall() (localctx IAggFuncCallContext) {
	localctx = NewAggFuncCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, FuncTestCaseParserRULE_aggFuncCall)
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FuncTestCaseParserDefine:
		localctx = NewMultiArgAggregateFuncCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(223)
			p.TableData()
		}
		{
			p.SetState(224)

			var _x = p.Identifier()

			localctx.(*MultiArgAggregateFuncCallContext).funcName = _x
		}
		{
			p.SetState(225)
			p.Match(FuncTestCaseParserOParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(226)
			p.QualifiedAggregateFuncArgs()
		}
		{
			p.SetState(227)
			p.Match(FuncTestCaseParserCParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserOParen:
		localctx = NewCompactAggregateFuncCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(229)
			p.TableRows()
		}
		{
			p.SetState(230)

			var _x = p.Identifier()

			localctx.(*CompactAggregateFuncCallContext).functName = _x
		}
		{
			p.SetState(231)
			p.Match(FuncTestCaseParserOParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(232)
			p.AggregateFuncArgs()
		}
		{
			p.SetState(233)
			p.Match(FuncTestCaseParserCParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserTruncate, FuncTestCaseParserAnd, FuncTestCaseParserOr, FuncTestCaseParserIdentifier:
		localctx = NewSingleArgAggregateFuncCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(235)

			var _x = p.Identifier()

			localctx.(*SingleArgAggregateFuncCallContext).functName = _x
		}
		{
			p.SetState(236)
			p.Match(FuncTestCaseParserOParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(237)
			p.DataColumn()
		}
		{
			p.SetState(238)
			p.Match(FuncTestCaseParserCParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITableDataContext is an interface to support dynamic dispatch.
type ITableDataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTableName returns the tableName token.
	GetTableName() antlr.Token

	// SetTableName sets the tableName token.
	SetTableName(antlr.Token)

	// Getter signatures
	Define() antlr.TerminalNode
	OParen() antlr.TerminalNode
	AllDataType() []IDataTypeContext
	DataType(i int) IDataTypeContext
	CParen() antlr.TerminalNode
	Eq() antlr.TerminalNode
	TableRows() ITableRowsContext
	Identifier() antlr.TerminalNode
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsTableDataContext differentiates from other interfaces.
	IsTableDataContext()
}

type TableDataContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	tableName antlr.Token
}

func NewEmptyTableDataContext() *TableDataContext {
	var p = new(TableDataContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_tableData
	return p
}

func InitEmptyTableDataContext(p *TableDataContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_tableData
}

func (*TableDataContext) IsTableDataContext() {}

func NewTableDataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableDataContext {
	var p = new(TableDataContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_tableData

	return p
}

func (s *TableDataContext) GetParser() antlr.Parser { return s.parser }

func (s *TableDataContext) GetTableName() antlr.Token { return s.tableName }

func (s *TableDataContext) SetTableName(v antlr.Token) { s.tableName = v }

func (s *TableDataContext) Define() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDefine, 0)
}

func (s *TableDataContext) OParen() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserOParen, 0)
}

func (s *TableDataContext) AllDataType() []IDataTypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDataTypeContext); ok {
			len++
		}
	}

	tst := make([]IDataTypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDataTypeContext); ok {
			tst[i] = t.(IDataTypeContext)
			i++
		}
	}

	return tst
}

func (s *TableDataContext) DataType(i int) IDataTypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataTypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataTypeContext)
}

func (s *TableDataContext) CParen() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserCParen, 0)
}

func (s *TableDataContext) Eq() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserEq, 0)
}

func (s *TableDataContext) TableRows() ITableRowsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableRowsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableRowsContext)
}

func (s *TableDataContext) Identifier() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserIdentifier, 0)
}

func (s *TableDataContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(FuncTestCaseParserComma)
}

func (s *TableDataContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserComma, i)
}

func (s *TableDataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableDataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableDataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterTableData(s)
	}
}

func (s *TableDataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitTableData(s)
	}
}

func (s *TableDataContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitTableData(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) TableData() (localctx ITableDataContext) {
	localctx = NewTableDataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, FuncTestCaseParserRULE_tableData)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.Match(FuncTestCaseParserDefine)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(243)

		var _m = p.Match(FuncTestCaseParserIdentifier)

		localctx.(*TableDataContext).tableName = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)
		p.Match(FuncTestCaseParserOParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(245)
		p.DataType()
	}
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FuncTestCaseParserComma {
		{
			p.SetState(246)
			p.Match(FuncTestCaseParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(247)
			p.DataType()
		}

		p.SetState(252)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(253)
		p.Match(FuncTestCaseParserCParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(254)
		p.Match(FuncTestCaseParserEq)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(255)
		p.TableRows()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITableRowsContext is an interface to support dynamic dispatch.
type ITableRowsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OParen() antlr.TerminalNode
	CParen() antlr.TerminalNode
	AllColumnValues() []IColumnValuesContext
	ColumnValues(i int) IColumnValuesContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsTableRowsContext differentiates from other interfaces.
	IsTableRowsContext()
}

type TableRowsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableRowsContext() *TableRowsContext {
	var p = new(TableRowsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_tableRows
	return p
}

func InitEmptyTableRowsContext(p *TableRowsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_tableRows
}

func (*TableRowsContext) IsTableRowsContext() {}

func NewTableRowsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableRowsContext {
	var p = new(TableRowsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_tableRows

	return p
}

func (s *TableRowsContext) GetParser() antlr.Parser { return s.parser }

func (s *TableRowsContext) OParen() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserOParen, 0)
}

func (s *TableRowsContext) CParen() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserCParen, 0)
}

func (s *TableRowsContext) AllColumnValues() []IColumnValuesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColumnValuesContext); ok {
			len++
		}
	}

	tst := make([]IColumnValuesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColumnValuesContext); ok {
			tst[i] = t.(IColumnValuesContext)
			i++
		}
	}

	return tst
}

func (s *TableRowsContext) ColumnValues(i int) IColumnValuesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnValuesContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnValuesContext)
}

func (s *TableRowsContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(FuncTestCaseParserComma)
}

func (s *TableRowsContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserComma, i)
}

func (s *TableRowsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableRowsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableRowsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterTableRows(s)
	}
}

func (s *TableRowsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitTableRows(s)
	}
}

func (s *TableRowsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitTableRows(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) TableRows() (localctx ITableRowsContext) {
	localctx = NewTableRowsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, FuncTestCaseParserRULE_tableRows)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(257)
		p.Match(FuncTestCaseParserOParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserOParen {
		{
			p.SetState(258)
			p.ColumnValues()
		}
		p.SetState(263)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FuncTestCaseParserComma {
			{
				p.SetState(259)
				p.Match(FuncTestCaseParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(260)
				p.ColumnValues()
			}

			p.SetState(265)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(268)
		p.Match(FuncTestCaseParserCParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDataColumnContext is an interface to support dynamic dispatch.
type IDataColumnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ColumnValues() IColumnValuesContext
	DoubleColon() antlr.TerminalNode
	DataType() IDataTypeContext

	// IsDataColumnContext differentiates from other interfaces.
	IsDataColumnContext()
}

type DataColumnContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataColumnContext() *DataColumnContext {
	var p = new(DataColumnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_dataColumn
	return p
}

func InitEmptyDataColumnContext(p *DataColumnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_dataColumn
}

func (*DataColumnContext) IsDataColumnContext() {}

func NewDataColumnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataColumnContext {
	var p = new(DataColumnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_dataColumn

	return p
}

func (s *DataColumnContext) GetParser() antlr.Parser { return s.parser }

func (s *DataColumnContext) ColumnValues() IColumnValuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnValuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnValuesContext)
}

func (s *DataColumnContext) DoubleColon() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDoubleColon, 0)
}

func (s *DataColumnContext) DataType() IDataTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataTypeContext)
}

func (s *DataColumnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataColumnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataColumnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterDataColumn(s)
	}
}

func (s *DataColumnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitDataColumn(s)
	}
}

func (s *DataColumnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitDataColumn(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) DataColumn() (localctx IDataColumnContext) {
	localctx = NewDataColumnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, FuncTestCaseParserRULE_dataColumn)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(270)
		p.ColumnValues()
	}
	{
		p.SetState(271)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(272)
		p.DataType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IColumnValuesContext is an interface to support dynamic dispatch.
type IColumnValuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OParen() antlr.TerminalNode
	CParen() antlr.TerminalNode
	AllLiteral() []ILiteralContext
	Literal(i int) ILiteralContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsColumnValuesContext differentiates from other interfaces.
	IsColumnValuesContext()
}

type ColumnValuesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnValuesContext() *ColumnValuesContext {
	var p = new(ColumnValuesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_columnValues
	return p
}

func InitEmptyColumnValuesContext(p *ColumnValuesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_columnValues
}

func (*ColumnValuesContext) IsColumnValuesContext() {}

func NewColumnValuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnValuesContext {
	var p = new(ColumnValuesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_columnValues

	return p
}

func (s *ColumnValuesContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnValuesContext) OParen() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserOParen, 0)
}

func (s *ColumnValuesContext) CParen() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserCParen, 0)
}

func (s *ColumnValuesContext) AllLiteral() []ILiteralContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILiteralContext); ok {
			len++
		}
	}

	tst := make([]ILiteralContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILiteralContext); ok {
			tst[i] = t.(ILiteralContext)
			i++
		}
	}

	return tst
}

func (s *ColumnValuesContext) Literal(i int) ILiteralContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ColumnValuesContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(FuncTestCaseParserComma)
}

func (s *ColumnValuesContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserComma, i)
}

func (s *ColumnValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnValuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnValuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterColumnValues(s)
	}
}

func (s *ColumnValuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitColumnValues(s)
	}
}

func (s *ColumnValuesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitColumnValues(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) ColumnValues() (localctx IColumnValuesContext) {
	localctx = NewColumnValuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, FuncTestCaseParserRULE_columnValues)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(274)
		p.Match(FuncTestCaseParserOParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32987488059392) != 0 {
		{
			p.SetState(275)
			p.Literal()
		}
		p.SetState(280)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FuncTestCaseParserComma {
			{
				p.SetState(276)
				p.Match(FuncTestCaseParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(277)
				p.Literal()
			}

			p.SetState(282)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(285)
		p.Match(FuncTestCaseParserCParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NullLiteral() antlr.TerminalNode
	NumericLiteral() INumericLiteralContext
	BooleanLiteral() antlr.TerminalNode
	StringLiteral() antlr.TerminalNode
	DateLiteral() antlr.TerminalNode
	TimeLiteral() antlr.TerminalNode
	TimestampLiteral() antlr.TerminalNode
	TimestampTzLiteral() antlr.TerminalNode
	IntervalYearLiteral() antlr.TerminalNode
	IntervalDayLiteral() antlr.TerminalNode

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) NullLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserNullLiteral, 0)
}

func (s *LiteralContext) NumericLiteral() INumericLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericLiteralContext)
}

func (s *LiteralContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserBooleanLiteral, 0)
}

func (s *LiteralContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserStringLiteral, 0)
}

func (s *LiteralContext) DateLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDateLiteral, 0)
}

func (s *LiteralContext) TimeLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserTimeLiteral, 0)
}

func (s *LiteralContext) TimestampLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserTimestampLiteral, 0)
}

func (s *LiteralContext) TimestampTzLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserTimestampTzLiteral, 0)
}

func (s *LiteralContext) IntervalYearLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserIntervalYearLiteral, 0)
}

func (s *LiteralContext) IntervalDayLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserIntervalDayLiteral, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, FuncTestCaseParserRULE_literal)
	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FuncTestCaseParserNullLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(287)
			p.Match(FuncTestCaseParserNullLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserNaN, FuncTestCaseParserIntegerLiteral, FuncTestCaseParserDecimalLiteral, FuncTestCaseParserFloatLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(288)
			p.NumericLiteral()
		}

	case FuncTestCaseParserBooleanLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(289)
			p.Match(FuncTestCaseParserBooleanLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserStringLiteral:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(290)
			p.Match(FuncTestCaseParserStringLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserDateLiteral:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(291)
			p.Match(FuncTestCaseParserDateLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserTimeLiteral:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(292)
			p.Match(FuncTestCaseParserTimeLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserTimestampLiteral:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(293)
			p.Match(FuncTestCaseParserTimestampLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserTimestampTzLiteral:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(294)
			p.Match(FuncTestCaseParserTimestampTzLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserIntervalYearLiteral:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(295)
			p.Match(FuncTestCaseParserIntervalYearLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserIntervalDayLiteral:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(296)
			p.Match(FuncTestCaseParserIntervalDayLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IQualifiedAggregateFuncArgsContext is an interface to support dynamic dispatch.
type IQualifiedAggregateFuncArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQualifiedAggregateFuncArg() []IQualifiedAggregateFuncArgContext
	QualifiedAggregateFuncArg(i int) IQualifiedAggregateFuncArgContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsQualifiedAggregateFuncArgsContext differentiates from other interfaces.
	IsQualifiedAggregateFuncArgsContext()
}

type QualifiedAggregateFuncArgsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualifiedAggregateFuncArgsContext() *QualifiedAggregateFuncArgsContext {
	var p = new(QualifiedAggregateFuncArgsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_qualifiedAggregateFuncArgs
	return p
}

func InitEmptyQualifiedAggregateFuncArgsContext(p *QualifiedAggregateFuncArgsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_qualifiedAggregateFuncArgs
}

func (*QualifiedAggregateFuncArgsContext) IsQualifiedAggregateFuncArgsContext() {}

func NewQualifiedAggregateFuncArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifiedAggregateFuncArgsContext {
	var p = new(QualifiedAggregateFuncArgsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_qualifiedAggregateFuncArgs

	return p
}

func (s *QualifiedAggregateFuncArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *QualifiedAggregateFuncArgsContext) AllQualifiedAggregateFuncArg() []IQualifiedAggregateFuncArgContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IQualifiedAggregateFuncArgContext); ok {
			len++
		}
	}

	tst := make([]IQualifiedAggregateFuncArgContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IQualifiedAggregateFuncArgContext); ok {
			tst[i] = t.(IQualifiedAggregateFuncArgContext)
			i++
		}
	}

	return tst
}

func (s *QualifiedAggregateFuncArgsContext) QualifiedAggregateFuncArg(i int) IQualifiedAggregateFuncArgContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQualifiedAggregateFuncArgContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQualifiedAggregateFuncArgContext)
}

func (s *QualifiedAggregateFuncArgsContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(FuncTestCaseParserComma)
}

func (s *QualifiedAggregateFuncArgsContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserComma, i)
}

func (s *QualifiedAggregateFuncArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifiedAggregateFuncArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualifiedAggregateFuncArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterQualifiedAggregateFuncArgs(s)
	}
}

func (s *QualifiedAggregateFuncArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitQualifiedAggregateFuncArgs(s)
	}
}

func (s *QualifiedAggregateFuncArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitQualifiedAggregateFuncArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) QualifiedAggregateFuncArgs() (localctx IQualifiedAggregateFuncArgsContext) {
	localctx = NewQualifiedAggregateFuncArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, FuncTestCaseParserRULE_qualifiedAggregateFuncArgs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(299)
		p.QualifiedAggregateFuncArg()
	}
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FuncTestCaseParserComma {
		{
			p.SetState(300)
			p.Match(FuncTestCaseParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(301)
			p.QualifiedAggregateFuncArg()
		}

		p.SetState(306)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAggregateFuncArgsContext is an interface to support dynamic dispatch.
type IAggregateFuncArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAggregateFuncArg() []IAggregateFuncArgContext
	AggregateFuncArg(i int) IAggregateFuncArgContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsAggregateFuncArgsContext differentiates from other interfaces.
	IsAggregateFuncArgsContext()
}

type AggregateFuncArgsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAggregateFuncArgsContext() *AggregateFuncArgsContext {
	var p = new(AggregateFuncArgsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_aggregateFuncArgs
	return p
}

func InitEmptyAggregateFuncArgsContext(p *AggregateFuncArgsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_aggregateFuncArgs
}

func (*AggregateFuncArgsContext) IsAggregateFuncArgsContext() {}

func NewAggregateFuncArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggregateFuncArgsContext {
	var p = new(AggregateFuncArgsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_aggregateFuncArgs

	return p
}

func (s *AggregateFuncArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *AggregateFuncArgsContext) AllAggregateFuncArg() []IAggregateFuncArgContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAggregateFuncArgContext); ok {
			len++
		}
	}

	tst := make([]IAggregateFuncArgContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAggregateFuncArgContext); ok {
			tst[i] = t.(IAggregateFuncArgContext)
			i++
		}
	}

	return tst
}

func (s *AggregateFuncArgsContext) AggregateFuncArg(i int) IAggregateFuncArgContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAggregateFuncArgContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAggregateFuncArgContext)
}

func (s *AggregateFuncArgsContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(FuncTestCaseParserComma)
}

func (s *AggregateFuncArgsContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserComma, i)
}

func (s *AggregateFuncArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggregateFuncArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AggregateFuncArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterAggregateFuncArgs(s)
	}
}

func (s *AggregateFuncArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitAggregateFuncArgs(s)
	}
}

func (s *AggregateFuncArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitAggregateFuncArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) AggregateFuncArgs() (localctx IAggregateFuncArgsContext) {
	localctx = NewAggregateFuncArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, FuncTestCaseParserRULE_aggregateFuncArgs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(307)
		p.AggregateFuncArg()
	}
	p.SetState(312)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FuncTestCaseParserComma {
		{
			p.SetState(308)
			p.Match(FuncTestCaseParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(309)
			p.AggregateFuncArg()
		}

		p.SetState(314)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IQualifiedAggregateFuncArgContext is an interface to support dynamic dispatch.
type IQualifiedAggregateFuncArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTableName returns the tableName token.
	GetTableName() antlr.Token

	// SetTableName sets the tableName token.
	SetTableName(antlr.Token)

	// Getter signatures
	Dot() antlr.TerminalNode
	ColumnName() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	Argument() IArgumentContext

	// IsQualifiedAggregateFuncArgContext differentiates from other interfaces.
	IsQualifiedAggregateFuncArgContext()
}

type QualifiedAggregateFuncArgContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	tableName antlr.Token
}

func NewEmptyQualifiedAggregateFuncArgContext() *QualifiedAggregateFuncArgContext {
	var p = new(QualifiedAggregateFuncArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_qualifiedAggregateFuncArg
	return p
}

func InitEmptyQualifiedAggregateFuncArgContext(p *QualifiedAggregateFuncArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_qualifiedAggregateFuncArg
}

func (*QualifiedAggregateFuncArgContext) IsQualifiedAggregateFuncArgContext() {}

func NewQualifiedAggregateFuncArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifiedAggregateFuncArgContext {
	var p = new(QualifiedAggregateFuncArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_qualifiedAggregateFuncArg

	return p
}

func (s *QualifiedAggregateFuncArgContext) GetParser() antlr.Parser { return s.parser }

func (s *QualifiedAggregateFuncArgContext) GetTableName() antlr.Token { return s.tableName }

func (s *QualifiedAggregateFuncArgContext) SetTableName(v antlr.Token) { s.tableName = v }

func (s *QualifiedAggregateFuncArgContext) Dot() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDot, 0)
}

func (s *QualifiedAggregateFuncArgContext) ColumnName() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserColumnName, 0)
}

func (s *QualifiedAggregateFuncArgContext) Identifier() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserIdentifier, 0)
}

func (s *QualifiedAggregateFuncArgContext) Argument() IArgumentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *QualifiedAggregateFuncArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifiedAggregateFuncArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualifiedAggregateFuncArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterQualifiedAggregateFuncArg(s)
	}
}

func (s *QualifiedAggregateFuncArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitQualifiedAggregateFuncArg(s)
	}
}

func (s *QualifiedAggregateFuncArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitQualifiedAggregateFuncArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) QualifiedAggregateFuncArg() (localctx IQualifiedAggregateFuncArgContext) {
	localctx = NewQualifiedAggregateFuncArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, FuncTestCaseParserRULE_qualifiedAggregateFuncArg)
	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FuncTestCaseParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(315)

			var _m = p.Match(FuncTestCaseParserIdentifier)

			localctx.(*QualifiedAggregateFuncArgContext).tableName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(316)
			p.Match(FuncTestCaseParserDot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(317)
			p.Match(FuncTestCaseParserColumnName)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserNaN, FuncTestCaseParserIntegerLiteral, FuncTestCaseParserDecimalLiteral, FuncTestCaseParserFloatLiteral, FuncTestCaseParserBooleanLiteral, FuncTestCaseParserTimestampTzLiteral, FuncTestCaseParserTimestampLiteral, FuncTestCaseParserTimeLiteral, FuncTestCaseParserDateLiteral, FuncTestCaseParserIntervalYearLiteral, FuncTestCaseParserIntervalDayLiteral, FuncTestCaseParserNullLiteral, FuncTestCaseParserStringLiteral, FuncTestCaseParserOBracket:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(318)
			p.Argument()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAggregateFuncArgContext is an interface to support dynamic dispatch.
type IAggregateFuncArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ColumnName() antlr.TerminalNode
	DoubleColon() antlr.TerminalNode
	DataType() IDataTypeContext
	Argument() IArgumentContext

	// IsAggregateFuncArgContext differentiates from other interfaces.
	IsAggregateFuncArgContext()
}

type AggregateFuncArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAggregateFuncArgContext() *AggregateFuncArgContext {
	var p = new(AggregateFuncArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_aggregateFuncArg
	return p
}

func InitEmptyAggregateFuncArgContext(p *AggregateFuncArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_aggregateFuncArg
}

func (*AggregateFuncArgContext) IsAggregateFuncArgContext() {}

func NewAggregateFuncArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggregateFuncArgContext {
	var p = new(AggregateFuncArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_aggregateFuncArg

	return p
}

func (s *AggregateFuncArgContext) GetParser() antlr.Parser { return s.parser }

func (s *AggregateFuncArgContext) ColumnName() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserColumnName, 0)
}

func (s *AggregateFuncArgContext) DoubleColon() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDoubleColon, 0)
}

func (s *AggregateFuncArgContext) DataType() IDataTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataTypeContext)
}

func (s *AggregateFuncArgContext) Argument() IArgumentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *AggregateFuncArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggregateFuncArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AggregateFuncArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterAggregateFuncArg(s)
	}
}

func (s *AggregateFuncArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitAggregateFuncArg(s)
	}
}

func (s *AggregateFuncArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitAggregateFuncArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) AggregateFuncArg() (localctx IAggregateFuncArgContext) {
	localctx = NewAggregateFuncArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, FuncTestCaseParserRULE_aggregateFuncArg)
	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FuncTestCaseParserColumnName:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(321)
			p.Match(FuncTestCaseParserColumnName)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(322)
			p.Match(FuncTestCaseParserDoubleColon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(323)
			p.DataType()
		}

	case FuncTestCaseParserNaN, FuncTestCaseParserIntegerLiteral, FuncTestCaseParserDecimalLiteral, FuncTestCaseParserFloatLiteral, FuncTestCaseParserBooleanLiteral, FuncTestCaseParserTimestampTzLiteral, FuncTestCaseParserTimestampLiteral, FuncTestCaseParserTimeLiteral, FuncTestCaseParserDateLiteral, FuncTestCaseParserIntervalYearLiteral, FuncTestCaseParserIntervalDayLiteral, FuncTestCaseParserNullLiteral, FuncTestCaseParserStringLiteral, FuncTestCaseParserOBracket:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(324)
			p.Argument()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumericLiteralContext is an interface to support dynamic dispatch.
type INumericLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DecimalLiteral() antlr.TerminalNode
	IntegerLiteral() antlr.TerminalNode
	FloatLiteral() IFloatLiteralContext

	// IsNumericLiteralContext differentiates from other interfaces.
	IsNumericLiteralContext()
}

type NumericLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericLiteralContext() *NumericLiteralContext {
	var p = new(NumericLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_numericLiteral
	return p
}

func InitEmptyNumericLiteralContext(p *NumericLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_numericLiteral
}

func (*NumericLiteralContext) IsNumericLiteralContext() {}

func NewNumericLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericLiteralContext {
	var p = new(NumericLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_numericLiteral

	return p
}

func (s *NumericLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericLiteralContext) DecimalLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDecimalLiteral, 0)
}

func (s *NumericLiteralContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserIntegerLiteral, 0)
}

func (s *NumericLiteralContext) FloatLiteral() IFloatLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloatLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloatLiteralContext)
}

func (s *NumericLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterNumericLiteral(s)
	}
}

func (s *NumericLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitNumericLiteral(s)
	}
}

func (s *NumericLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitNumericLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) NumericLiteral() (localctx INumericLiteralContext) {
	localctx = NewNumericLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, FuncTestCaseParserRULE_numericLiteral)
	p.SetState(330)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FuncTestCaseParserDecimalLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(327)
			p.Match(FuncTestCaseParserDecimalLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserIntegerLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(328)
			p.Match(FuncTestCaseParserIntegerLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserNaN, FuncTestCaseParserFloatLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(329)
			p.FloatLiteral()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFloatLiteralContext is an interface to support dynamic dispatch.
type IFloatLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FloatLiteral() antlr.TerminalNode
	NaN() antlr.TerminalNode

	// IsFloatLiteralContext differentiates from other interfaces.
	IsFloatLiteralContext()
}

type FloatLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatLiteralContext() *FloatLiteralContext {
	var p = new(FloatLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_floatLiteral
	return p
}

func InitEmptyFloatLiteralContext(p *FloatLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_floatLiteral
}

func (*FloatLiteralContext) IsFloatLiteralContext() {}

func NewFloatLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatLiteralContext {
	var p = new(FloatLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_floatLiteral

	return p
}

func (s *FloatLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatLiteralContext) FloatLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserFloatLiteral, 0)
}

func (s *FloatLiteralContext) NaN() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserNaN, 0)
}

func (s *FloatLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitFloatLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) FloatLiteral() (localctx IFloatLiteralContext) {
	localctx = NewFloatLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, FuncTestCaseParserRULE_floatLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(332)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserNaN || _la == FuncTestCaseParserFloatLiteral) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INullArgContext is an interface to support dynamic dispatch.
type INullArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NullLiteral() antlr.TerminalNode
	DoubleColon() antlr.TerminalNode
	DataType() IDataTypeContext

	// IsNullArgContext differentiates from other interfaces.
	IsNullArgContext()
}

type NullArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNullArgContext() *NullArgContext {
	var p = new(NullArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_nullArg
	return p
}

func InitEmptyNullArgContext(p *NullArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_nullArg
}

func (*NullArgContext) IsNullArgContext() {}

func NewNullArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NullArgContext {
	var p = new(NullArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_nullArg

	return p
}

func (s *NullArgContext) GetParser() antlr.Parser { return s.parser }

func (s *NullArgContext) NullLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserNullLiteral, 0)
}

func (s *NullArgContext) DoubleColon() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDoubleColon, 0)
}

func (s *NullArgContext) DataType() IDataTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataTypeContext)
}

func (s *NullArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NullArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterNullArg(s)
	}
}

func (s *NullArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitNullArg(s)
	}
}

func (s *NullArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitNullArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) NullArg() (localctx INullArgContext) {
	localctx = NewNullArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, FuncTestCaseParserRULE_nullArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(334)
		p.Match(FuncTestCaseParserNullLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(335)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(336)
		p.DataType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIntArgContext is an interface to support dynamic dispatch.
type IIntArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IntegerLiteral() antlr.TerminalNode
	DoubleColon() antlr.TerminalNode
	I8() antlr.TerminalNode
	I16() antlr.TerminalNode
	I32() antlr.TerminalNode
	I64() antlr.TerminalNode

	// IsIntArgContext differentiates from other interfaces.
	IsIntArgContext()
}

type IntArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntArgContext() *IntArgContext {
	var p = new(IntArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_intArg
	return p
}

func InitEmptyIntArgContext(p *IntArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_intArg
}

func (*IntArgContext) IsIntArgContext() {}

func NewIntArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntArgContext {
	var p = new(IntArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_intArg

	return p
}

func (s *IntArgContext) GetParser() antlr.Parser { return s.parser }

func (s *IntArgContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserIntegerLiteral, 0)
}

func (s *IntArgContext) DoubleColon() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDoubleColon, 0)
}

func (s *IntArgContext) I8() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserI8, 0)
}

func (s *IntArgContext) I16() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserI16, 0)
}

func (s *IntArgContext) I32() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserI32, 0)
}

func (s *IntArgContext) I64() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserI64, 0)
}

func (s *IntArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterIntArg(s)
	}
}

func (s *IntArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitIntArg(s)
	}
}

func (s *IntArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitIntArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) IntArg() (localctx IIntArgContext) {
	localctx = NewIntArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, FuncTestCaseParserRULE_intArg)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(338)
		p.Match(FuncTestCaseParserIntegerLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(339)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(340)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&67553994410557440) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFloatArgContext is an interface to support dynamic dispatch.
type IFloatArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NumericLiteral() INumericLiteralContext
	DoubleColon() antlr.TerminalNode
	FP32() antlr.TerminalNode
	FP64() antlr.TerminalNode

	// IsFloatArgContext differentiates from other interfaces.
	IsFloatArgContext()
}

type FloatArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatArgContext() *FloatArgContext {
	var p = new(FloatArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_floatArg
	return p
}

func InitEmptyFloatArgContext(p *FloatArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_floatArg
}

func (*FloatArgContext) IsFloatArgContext() {}

func NewFloatArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatArgContext {
	var p = new(FloatArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_floatArg

	return p
}

func (s *FloatArgContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatArgContext) NumericLiteral() INumericLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericLiteralContext)
}

func (s *FloatArgContext) DoubleColon() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDoubleColon, 0)
}

func (s *FloatArgContext) FP32() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserFP32, 0)
}

func (s *FloatArgContext) FP64() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserFP64, 0)
}

func (s *FloatArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterFloatArg(s)
	}
}

func (s *FloatArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitFloatArg(s)
	}
}

func (s *FloatArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitFloatArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) FloatArg() (localctx IFloatArgContext) {
	localctx = NewFloatArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, FuncTestCaseParserRULE_floatArg)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(342)
		p.NumericLiteral()
	}
	{
		p.SetState(343)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(344)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserFP32 || _la == FuncTestCaseParserFP64) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDecimalArgContext is an interface to support dynamic dispatch.
type IDecimalArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NumericLiteral() INumericLiteralContext
	DoubleColon() antlr.TerminalNode
	DecimalType() IDecimalTypeContext

	// IsDecimalArgContext differentiates from other interfaces.
	IsDecimalArgContext()
}

type DecimalArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecimalArgContext() *DecimalArgContext {
	var p = new(DecimalArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_decimalArg
	return p
}

func InitEmptyDecimalArgContext(p *DecimalArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_decimalArg
}

func (*DecimalArgContext) IsDecimalArgContext() {}

func NewDecimalArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecimalArgContext {
	var p = new(DecimalArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_decimalArg

	return p
}

func (s *DecimalArgContext) GetParser() antlr.Parser { return s.parser }

func (s *DecimalArgContext) NumericLiteral() INumericLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericLiteralContext)
}

func (s *DecimalArgContext) DoubleColon() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDoubleColon, 0)
}

func (s *DecimalArgContext) DecimalType() IDecimalTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecimalTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecimalTypeContext)
}

func (s *DecimalArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecimalArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterDecimalArg(s)
	}
}

func (s *DecimalArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitDecimalArg(s)
	}
}

func (s *DecimalArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitDecimalArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) DecimalArg() (localctx IDecimalArgContext) {
	localctx = NewDecimalArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, FuncTestCaseParserRULE_decimalArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(346)
		p.NumericLiteral()
	}
	{
		p.SetState(347)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(348)
		p.DecimalType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBooleanArgContext is an interface to support dynamic dispatch.
type IBooleanArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BooleanLiteral() antlr.TerminalNode
	DoubleColon() antlr.TerminalNode
	BooleanType() IBooleanTypeContext

	// IsBooleanArgContext differentiates from other interfaces.
	IsBooleanArgContext()
}

type BooleanArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanArgContext() *BooleanArgContext {
	var p = new(BooleanArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_booleanArg
	return p
}

func InitEmptyBooleanArgContext(p *BooleanArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_booleanArg
}

func (*BooleanArgContext) IsBooleanArgContext() {}

func NewBooleanArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanArgContext {
	var p = new(BooleanArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_booleanArg

	return p
}

func (s *BooleanArgContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanArgContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserBooleanLiteral, 0)
}

func (s *BooleanArgContext) DoubleColon() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDoubleColon, 0)
}

func (s *BooleanArgContext) BooleanType() IBooleanTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanTypeContext)
}

func (s *BooleanArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterBooleanArg(s)
	}
}

func (s *BooleanArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitBooleanArg(s)
	}
}

func (s *BooleanArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitBooleanArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) BooleanArg() (localctx IBooleanArgContext) {
	localctx = NewBooleanArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, FuncTestCaseParserRULE_booleanArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(350)
		p.Match(FuncTestCaseParserBooleanLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(351)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(352)
		p.BooleanType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStringArgContext is an interface to support dynamic dispatch.
type IStringArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StringLiteral() antlr.TerminalNode
	DoubleColon() antlr.TerminalNode
	StringType() IStringTypeContext

	// IsStringArgContext differentiates from other interfaces.
	IsStringArgContext()
}

type StringArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringArgContext() *StringArgContext {
	var p = new(StringArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_stringArg
	return p
}

func InitEmptyStringArgContext(p *StringArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_stringArg
}

func (*StringArgContext) IsStringArgContext() {}

func NewStringArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringArgContext {
	var p = new(StringArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_stringArg

	return p
}

func (s *StringArgContext) GetParser() antlr.Parser { return s.parser }

func (s *StringArgContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserStringLiteral, 0)
}

func (s *StringArgContext) DoubleColon() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDoubleColon, 0)
}

func (s *StringArgContext) StringType() IStringTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringTypeContext)
}

func (s *StringArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterStringArg(s)
	}
}

func (s *StringArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitStringArg(s)
	}
}

func (s *StringArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitStringArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) StringArg() (localctx IStringArgContext) {
	localctx = NewStringArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, FuncTestCaseParserRULE_stringArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(354)
		p.Match(FuncTestCaseParserStringLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(355)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(356)
		p.StringType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDateArgContext is an interface to support dynamic dispatch.
type IDateArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DateLiteral() antlr.TerminalNode
	DoubleColon() antlr.TerminalNode
	Date() antlr.TerminalNode

	// IsDateArgContext differentiates from other interfaces.
	IsDateArgContext()
}

type DateArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateArgContext() *DateArgContext {
	var p = new(DateArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_dateArg
	return p
}

func InitEmptyDateArgContext(p *DateArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_dateArg
}

func (*DateArgContext) IsDateArgContext() {}

func NewDateArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateArgContext {
	var p = new(DateArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_dateArg

	return p
}

func (s *DateArgContext) GetParser() antlr.Parser { return s.parser }

func (s *DateArgContext) DateLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDateLiteral, 0)
}

func (s *DateArgContext) DoubleColon() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDoubleColon, 0)
}

func (s *DateArgContext) Date() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDate, 0)
}

func (s *DateArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterDateArg(s)
	}
}

func (s *DateArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitDateArg(s)
	}
}

func (s *DateArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitDateArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) DateArg() (localctx IDateArgContext) {
	localctx = NewDateArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, FuncTestCaseParserRULE_dateArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(358)
		p.Match(FuncTestCaseParserDateLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(359)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(360)
		p.Match(FuncTestCaseParserDate)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITimeArgContext is an interface to support dynamic dispatch.
type ITimeArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TimeLiteral() antlr.TerminalNode
	DoubleColon() antlr.TerminalNode
	Time() antlr.TerminalNode

	// IsTimeArgContext differentiates from other interfaces.
	IsTimeArgContext()
}

type TimeArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeArgContext() *TimeArgContext {
	var p = new(TimeArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_timeArg
	return p
}

func InitEmptyTimeArgContext(p *TimeArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_timeArg
}

func (*TimeArgContext) IsTimeArgContext() {}

func NewTimeArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeArgContext {
	var p = new(TimeArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_timeArg

	return p
}

func (s *TimeArgContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeArgContext) TimeLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserTimeLiteral, 0)
}

func (s *TimeArgContext) DoubleColon() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDoubleColon, 0)
}

func (s *TimeArgContext) Time() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserTime, 0)
}

func (s *TimeArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterTimeArg(s)
	}
}

func (s *TimeArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitTimeArg(s)
	}
}

func (s *TimeArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitTimeArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) TimeArg() (localctx ITimeArgContext) {
	localctx = NewTimeArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, FuncTestCaseParserRULE_timeArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(362)
		p.Match(FuncTestCaseParserTimeLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(363)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(364)
		p.Match(FuncTestCaseParserTime)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITimestampArgContext is an interface to support dynamic dispatch.
type ITimestampArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TimestampLiteral() antlr.TerminalNode
	DoubleColon() antlr.TerminalNode
	TimestampType() ITimestampTypeContext

	// IsTimestampArgContext differentiates from other interfaces.
	IsTimestampArgContext()
}

type TimestampArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimestampArgContext() *TimestampArgContext {
	var p = new(TimestampArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_timestampArg
	return p
}

func InitEmptyTimestampArgContext(p *TimestampArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_timestampArg
}

func (*TimestampArgContext) IsTimestampArgContext() {}

func NewTimestampArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimestampArgContext {
	var p = new(TimestampArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_timestampArg

	return p
}

func (s *TimestampArgContext) GetParser() antlr.Parser { return s.parser }

func (s *TimestampArgContext) TimestampLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserTimestampLiteral, 0)
}

func (s *TimestampArgContext) DoubleColon() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDoubleColon, 0)
}

func (s *TimestampArgContext) TimestampType() ITimestampTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimestampTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimestampTypeContext)
}

func (s *TimestampArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimestampArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimestampArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterTimestampArg(s)
	}
}

func (s *TimestampArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitTimestampArg(s)
	}
}

func (s *TimestampArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitTimestampArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) TimestampArg() (localctx ITimestampArgContext) {
	localctx = NewTimestampArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, FuncTestCaseParserRULE_timestampArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(366)
		p.Match(FuncTestCaseParserTimestampLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(367)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(368)
		p.TimestampType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITimestampTzArgContext is an interface to support dynamic dispatch.
type ITimestampTzArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TimestampTzLiteral() antlr.TerminalNode
	DoubleColon() antlr.TerminalNode
	TimestampTZType() ITimestampTZTypeContext

	// IsTimestampTzArgContext differentiates from other interfaces.
	IsTimestampTzArgContext()
}

type TimestampTzArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimestampTzArgContext() *TimestampTzArgContext {
	var p = new(TimestampTzArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_timestampTzArg
	return p
}

func InitEmptyTimestampTzArgContext(p *TimestampTzArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_timestampTzArg
}

func (*TimestampTzArgContext) IsTimestampTzArgContext() {}

func NewTimestampTzArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimestampTzArgContext {
	var p = new(TimestampTzArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_timestampTzArg

	return p
}

func (s *TimestampTzArgContext) GetParser() antlr.Parser { return s.parser }

func (s *TimestampTzArgContext) TimestampTzLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserTimestampTzLiteral, 0)
}

func (s *TimestampTzArgContext) DoubleColon() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDoubleColon, 0)
}

func (s *TimestampTzArgContext) TimestampTZType() ITimestampTZTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimestampTZTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimestampTZTypeContext)
}

func (s *TimestampTzArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimestampTzArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimestampTzArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterTimestampTzArg(s)
	}
}

func (s *TimestampTzArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitTimestampTzArg(s)
	}
}

func (s *TimestampTzArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitTimestampTzArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) TimestampTzArg() (localctx ITimestampTzArgContext) {
	localctx = NewTimestampTzArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, FuncTestCaseParserRULE_timestampTzArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(370)
		p.Match(FuncTestCaseParserTimestampTzLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(371)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(372)
		p.TimestampTZType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIntervalYearArgContext is an interface to support dynamic dispatch.
type IIntervalYearArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IntervalYearLiteral() antlr.TerminalNode
	DoubleColon() antlr.TerminalNode
	IntervalYearType() IIntervalYearTypeContext

	// IsIntervalYearArgContext differentiates from other interfaces.
	IsIntervalYearArgContext()
}

type IntervalYearArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntervalYearArgContext() *IntervalYearArgContext {
	var p = new(IntervalYearArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_intervalYearArg
	return p
}

func InitEmptyIntervalYearArgContext(p *IntervalYearArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_intervalYearArg
}

func (*IntervalYearArgContext) IsIntervalYearArgContext() {}

func NewIntervalYearArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntervalYearArgContext {
	var p = new(IntervalYearArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_intervalYearArg

	return p
}

func (s *IntervalYearArgContext) GetParser() antlr.Parser { return s.parser }

func (s *IntervalYearArgContext) IntervalYearLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserIntervalYearLiteral, 0)
}

func (s *IntervalYearArgContext) DoubleColon() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDoubleColon, 0)
}

func (s *IntervalYearArgContext) IntervalYearType() IIntervalYearTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntervalYearTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntervalYearTypeContext)
}

func (s *IntervalYearArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntervalYearArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntervalYearArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterIntervalYearArg(s)
	}
}

func (s *IntervalYearArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitIntervalYearArg(s)
	}
}

func (s *IntervalYearArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitIntervalYearArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) IntervalYearArg() (localctx IIntervalYearArgContext) {
	localctx = NewIntervalYearArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, FuncTestCaseParserRULE_intervalYearArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(374)
		p.Match(FuncTestCaseParserIntervalYearLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(375)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(376)
		p.IntervalYearType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIntervalDayArgContext is an interface to support dynamic dispatch.
type IIntervalDayArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IntervalDayLiteral() antlr.TerminalNode
	DoubleColon() antlr.TerminalNode
	IntervalDayType() IIntervalDayTypeContext

	// IsIntervalDayArgContext differentiates from other interfaces.
	IsIntervalDayArgContext()
}

type IntervalDayArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntervalDayArgContext() *IntervalDayArgContext {
	var p = new(IntervalDayArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_intervalDayArg
	return p
}

func InitEmptyIntervalDayArgContext(p *IntervalDayArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_intervalDayArg
}

func (*IntervalDayArgContext) IsIntervalDayArgContext() {}

func NewIntervalDayArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntervalDayArgContext {
	var p = new(IntervalDayArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_intervalDayArg

	return p
}

func (s *IntervalDayArgContext) GetParser() antlr.Parser { return s.parser }

func (s *IntervalDayArgContext) IntervalDayLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserIntervalDayLiteral, 0)
}

func (s *IntervalDayArgContext) DoubleColon() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDoubleColon, 0)
}

func (s *IntervalDayArgContext) IntervalDayType() IIntervalDayTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntervalDayTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntervalDayTypeContext)
}

func (s *IntervalDayArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntervalDayArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntervalDayArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterIntervalDayArg(s)
	}
}

func (s *IntervalDayArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitIntervalDayArg(s)
	}
}

func (s *IntervalDayArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitIntervalDayArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) IntervalDayArg() (localctx IIntervalDayArgContext) {
	localctx = NewIntervalDayArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, FuncTestCaseParserRULE_intervalDayArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(378)
		p.Match(FuncTestCaseParserIntervalDayLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(379)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(380)
		p.IntervalDayType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListArgContext is an interface to support dynamic dispatch.
type IListArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LiteralList() ILiteralListContext
	DoubleColon() antlr.TerminalNode
	ListType() IListTypeContext

	// IsListArgContext differentiates from other interfaces.
	IsListArgContext()
}

type ListArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListArgContext() *ListArgContext {
	var p = new(ListArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_listArg
	return p
}

func InitEmptyListArgContext(p *ListArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_listArg
}

func (*ListArgContext) IsListArgContext() {}

func NewListArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListArgContext {
	var p = new(ListArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_listArg

	return p
}

func (s *ListArgContext) GetParser() antlr.Parser { return s.parser }

func (s *ListArgContext) LiteralList() ILiteralListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralListContext)
}

func (s *ListArgContext) DoubleColon() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDoubleColon, 0)
}

func (s *ListArgContext) ListType() IListTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListTypeContext)
}

func (s *ListArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterListArg(s)
	}
}

func (s *ListArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitListArg(s)
	}
}

func (s *ListArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitListArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) ListArg() (localctx IListArgContext) {
	localctx = NewListArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, FuncTestCaseParserRULE_listArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(382)
		p.LiteralList()
	}
	{
		p.SetState(383)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(384)
		p.ListType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralListContext is an interface to support dynamic dispatch.
type ILiteralListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OBracket() antlr.TerminalNode
	CBracket() antlr.TerminalNode
	AllLiteral() []ILiteralContext
	Literal(i int) ILiteralContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsLiteralListContext differentiates from other interfaces.
	IsLiteralListContext()
}

type LiteralListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralListContext() *LiteralListContext {
	var p = new(LiteralListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_literalList
	return p
}

func InitEmptyLiteralListContext(p *LiteralListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_literalList
}

func (*LiteralListContext) IsLiteralListContext() {}

func NewLiteralListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralListContext {
	var p = new(LiteralListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_literalList

	return p
}

func (s *LiteralListContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralListContext) OBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserOBracket, 0)
}

func (s *LiteralListContext) CBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserCBracket, 0)
}

func (s *LiteralListContext) AllLiteral() []ILiteralContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILiteralContext); ok {
			len++
		}
	}

	tst := make([]ILiteralContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILiteralContext); ok {
			tst[i] = t.(ILiteralContext)
			i++
		}
	}

	return tst
}

func (s *LiteralListContext) Literal(i int) ILiteralContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralListContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(FuncTestCaseParserComma)
}

func (s *LiteralListContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserComma, i)
}

func (s *LiteralListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterLiteralList(s)
	}
}

func (s *LiteralListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitLiteralList(s)
	}
}

func (s *LiteralListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitLiteralList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) LiteralList() (localctx ILiteralListContext) {
	localctx = NewLiteralListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, FuncTestCaseParserRULE_literalList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(386)
		p.Match(FuncTestCaseParserOBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(395)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32987488059392) != 0 {
		{
			p.SetState(387)
			p.Literal()
		}
		p.SetState(392)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FuncTestCaseParserComma {
			{
				p.SetState(388)
				p.Match(FuncTestCaseParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(389)
				p.Literal()
			}

			p.SetState(394)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(397)
		p.Match(FuncTestCaseParserCBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIntervalYearLiteralContext is an interface to support dynamic dispatch.
type IIntervalYearLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetYears returns the years token.
	GetYears() antlr.Token

	// GetMonths returns the months token.
	GetMonths() antlr.Token

	// SetYears sets the years token.
	SetYears(antlr.Token)

	// SetMonths sets the months token.
	SetMonths(antlr.Token)

	// Getter signatures
	PeriodPrefix() antlr.TerminalNode
	YearPrefix() antlr.TerminalNode
	AllIntegerLiteral() []antlr.TerminalNode
	IntegerLiteral(i int) antlr.TerminalNode
	MSuffix() antlr.TerminalNode

	// IsIntervalYearLiteralContext differentiates from other interfaces.
	IsIntervalYearLiteralContext()
}

type IntervalYearLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	years  antlr.Token
	months antlr.Token
}

func NewEmptyIntervalYearLiteralContext() *IntervalYearLiteralContext {
	var p = new(IntervalYearLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_intervalYearLiteral
	return p
}

func InitEmptyIntervalYearLiteralContext(p *IntervalYearLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_intervalYearLiteral
}

func (*IntervalYearLiteralContext) IsIntervalYearLiteralContext() {}

func NewIntervalYearLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntervalYearLiteralContext {
	var p = new(IntervalYearLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_intervalYearLiteral

	return p
}

func (s *IntervalYearLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *IntervalYearLiteralContext) GetYears() antlr.Token { return s.years }

func (s *IntervalYearLiteralContext) GetMonths() antlr.Token { return s.months }

func (s *IntervalYearLiteralContext) SetYears(v antlr.Token) { s.years = v }

func (s *IntervalYearLiteralContext) SetMonths(v antlr.Token) { s.months = v }

func (s *IntervalYearLiteralContext) PeriodPrefix() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserPeriodPrefix, 0)
}

func (s *IntervalYearLiteralContext) YearPrefix() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserYearPrefix, 0)
}

func (s *IntervalYearLiteralContext) AllIntegerLiteral() []antlr.TerminalNode {
	return s.GetTokens(FuncTestCaseParserIntegerLiteral)
}

func (s *IntervalYearLiteralContext) IntegerLiteral(i int) antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserIntegerLiteral, i)
}

func (s *IntervalYearLiteralContext) MSuffix() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserMSuffix, 0)
}

func (s *IntervalYearLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntervalYearLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntervalYearLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterIntervalYearLiteral(s)
	}
}

func (s *IntervalYearLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitIntervalYearLiteral(s)
	}
}

func (s *IntervalYearLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitIntervalYearLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) IntervalYearLiteral() (localctx IIntervalYearLiteralContext) {
	localctx = NewIntervalYearLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, FuncTestCaseParserRULE_intervalYearLiteral)
	var _la int

	p.SetState(410)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(399)
			p.Match(FuncTestCaseParserPeriodPrefix)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(400)

			var _m = p.Match(FuncTestCaseParserIntegerLiteral)

			localctx.(*IntervalYearLiteralContext).years = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(401)
			p.Match(FuncTestCaseParserYearPrefix)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(405)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FuncTestCaseParserIntegerLiteral {
			{
				p.SetState(403)

				var _m = p.Match(FuncTestCaseParserIntegerLiteral)

				localctx.(*IntervalYearLiteralContext).months = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(404)
				p.Match(FuncTestCaseParserMSuffix)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(407)
			p.Match(FuncTestCaseParserPeriodPrefix)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(408)

			var _m = p.Match(FuncTestCaseParserIntegerLiteral)

			localctx.(*IntervalYearLiteralContext).months = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(409)
			p.Match(FuncTestCaseParserMSuffix)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIntervalDayLiteralContext is an interface to support dynamic dispatch.
type IIntervalDayLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDays returns the days token.
	GetDays() antlr.Token

	// SetDays sets the days token.
	SetDays(antlr.Token)

	// Getter signatures
	PeriodPrefix() antlr.TerminalNode
	DaySuffix() antlr.TerminalNode
	IntegerLiteral() antlr.TerminalNode
	TimePrefix() antlr.TerminalNode
	TimeInterval() ITimeIntervalContext

	// IsIntervalDayLiteralContext differentiates from other interfaces.
	IsIntervalDayLiteralContext()
}

type IntervalDayLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	days   antlr.Token
}

func NewEmptyIntervalDayLiteralContext() *IntervalDayLiteralContext {
	var p = new(IntervalDayLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_intervalDayLiteral
	return p
}

func InitEmptyIntervalDayLiteralContext(p *IntervalDayLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_intervalDayLiteral
}

func (*IntervalDayLiteralContext) IsIntervalDayLiteralContext() {}

func NewIntervalDayLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntervalDayLiteralContext {
	var p = new(IntervalDayLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_intervalDayLiteral

	return p
}

func (s *IntervalDayLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *IntervalDayLiteralContext) GetDays() antlr.Token { return s.days }

func (s *IntervalDayLiteralContext) SetDays(v antlr.Token) { s.days = v }

func (s *IntervalDayLiteralContext) PeriodPrefix() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserPeriodPrefix, 0)
}

func (s *IntervalDayLiteralContext) DaySuffix() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDaySuffix, 0)
}

func (s *IntervalDayLiteralContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserIntegerLiteral, 0)
}

func (s *IntervalDayLiteralContext) TimePrefix() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserTimePrefix, 0)
}

func (s *IntervalDayLiteralContext) TimeInterval() ITimeIntervalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimeIntervalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimeIntervalContext)
}

func (s *IntervalDayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntervalDayLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntervalDayLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterIntervalDayLiteral(s)
	}
}

func (s *IntervalDayLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitIntervalDayLiteral(s)
	}
}

func (s *IntervalDayLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitIntervalDayLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) IntervalDayLiteral() (localctx IIntervalDayLiteralContext) {
	localctx = NewIntervalDayLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, FuncTestCaseParserRULE_intervalDayLiteral)
	var _la int

	p.SetState(423)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(412)
			p.Match(FuncTestCaseParserPeriodPrefix)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(413)

			var _m = p.Match(FuncTestCaseParserIntegerLiteral)

			localctx.(*IntervalDayLiteralContext).days = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(414)
			p.Match(FuncTestCaseParserDaySuffix)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(418)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FuncTestCaseParserTimePrefix {
			{
				p.SetState(416)
				p.Match(FuncTestCaseParserTimePrefix)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(417)
				p.TimeInterval()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(420)
			p.Match(FuncTestCaseParserPeriodPrefix)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(421)
			p.Match(FuncTestCaseParserTimePrefix)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(422)
			p.TimeInterval()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITimeIntervalContext is an interface to support dynamic dispatch.
type ITimeIntervalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetHours returns the hours token.
	GetHours() antlr.Token

	// GetMinutes returns the minutes token.
	GetMinutes() antlr.Token

	// GetSeconds returns the seconds token.
	GetSeconds() antlr.Token

	// GetFractionalSeconds returns the fractionalSeconds token.
	GetFractionalSeconds() antlr.Token

	// SetHours sets the hours token.
	SetHours(antlr.Token)

	// SetMinutes sets the minutes token.
	SetMinutes(antlr.Token)

	// SetSeconds sets the seconds token.
	SetSeconds(antlr.Token)

	// SetFractionalSeconds sets the fractionalSeconds token.
	SetFractionalSeconds(antlr.Token)

	// Getter signatures
	HourSuffix() antlr.TerminalNode
	AllIntegerLiteral() []antlr.TerminalNode
	IntegerLiteral(i int) antlr.TerminalNode
	MSuffix() antlr.TerminalNode
	SecondSuffix() antlr.TerminalNode
	FractionalSecondSuffix() antlr.TerminalNode

	// IsTimeIntervalContext differentiates from other interfaces.
	IsTimeIntervalContext()
}

type TimeIntervalContext struct {
	antlr.BaseParserRuleContext
	parser            antlr.Parser
	hours             antlr.Token
	minutes           antlr.Token
	seconds           antlr.Token
	fractionalSeconds antlr.Token
}

func NewEmptyTimeIntervalContext() *TimeIntervalContext {
	var p = new(TimeIntervalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_timeInterval
	return p
}

func InitEmptyTimeIntervalContext(p *TimeIntervalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_timeInterval
}

func (*TimeIntervalContext) IsTimeIntervalContext() {}

func NewTimeIntervalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeIntervalContext {
	var p = new(TimeIntervalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_timeInterval

	return p
}

func (s *TimeIntervalContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeIntervalContext) GetHours() antlr.Token { return s.hours }

func (s *TimeIntervalContext) GetMinutes() antlr.Token { return s.minutes }

func (s *TimeIntervalContext) GetSeconds() antlr.Token { return s.seconds }

func (s *TimeIntervalContext) GetFractionalSeconds() antlr.Token { return s.fractionalSeconds }

func (s *TimeIntervalContext) SetHours(v antlr.Token) { s.hours = v }

func (s *TimeIntervalContext) SetMinutes(v antlr.Token) { s.minutes = v }

func (s *TimeIntervalContext) SetSeconds(v antlr.Token) { s.seconds = v }

func (s *TimeIntervalContext) SetFractionalSeconds(v antlr.Token) { s.fractionalSeconds = v }

func (s *TimeIntervalContext) HourSuffix() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserHourSuffix, 0)
}

func (s *TimeIntervalContext) AllIntegerLiteral() []antlr.TerminalNode {
	return s.GetTokens(FuncTestCaseParserIntegerLiteral)
}

func (s *TimeIntervalContext) IntegerLiteral(i int) antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserIntegerLiteral, i)
}

func (s *TimeIntervalContext) MSuffix() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserMSuffix, 0)
}

func (s *TimeIntervalContext) SecondSuffix() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserSecondSuffix, 0)
}

func (s *TimeIntervalContext) FractionalSecondSuffix() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserFractionalSecondSuffix, 0)
}

func (s *TimeIntervalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeIntervalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeIntervalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterTimeInterval(s)
	}
}

func (s *TimeIntervalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitTimeInterval(s)
	}
}

func (s *TimeIntervalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitTimeInterval(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) TimeInterval() (localctx ITimeIntervalContext) {
	localctx = NewTimeIntervalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, FuncTestCaseParserRULE_timeInterval)
	var _la int

	p.SetState(457)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(425)

			var _m = p.Match(FuncTestCaseParserIntegerLiteral)

			localctx.(*TimeIntervalContext).hours = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(426)
			p.Match(FuncTestCaseParserHourSuffix)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(429)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(427)

				var _m = p.Match(FuncTestCaseParserIntegerLiteral)

				localctx.(*TimeIntervalContext).minutes = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(428)
				p.Match(FuncTestCaseParserMSuffix)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(433)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(431)

				var _m = p.Match(FuncTestCaseParserIntegerLiteral)

				localctx.(*TimeIntervalContext).seconds = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(432)
				p.Match(FuncTestCaseParserSecondSuffix)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(437)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FuncTestCaseParserIntegerLiteral {
			{
				p.SetState(435)

				var _m = p.Match(FuncTestCaseParserIntegerLiteral)

				localctx.(*TimeIntervalContext).fractionalSeconds = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(436)
				p.Match(FuncTestCaseParserFractionalSecondSuffix)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(439)

			var _m = p.Match(FuncTestCaseParserIntegerLiteral)

			localctx.(*TimeIntervalContext).minutes = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(440)
			p.Match(FuncTestCaseParserMSuffix)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(443)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(441)

				var _m = p.Match(FuncTestCaseParserIntegerLiteral)

				localctx.(*TimeIntervalContext).seconds = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(442)
				p.Match(FuncTestCaseParserSecondSuffix)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(447)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FuncTestCaseParserIntegerLiteral {
			{
				p.SetState(445)

				var _m = p.Match(FuncTestCaseParserIntegerLiteral)

				localctx.(*TimeIntervalContext).fractionalSeconds = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(446)
				p.Match(FuncTestCaseParserFractionalSecondSuffix)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(449)

			var _m = p.Match(FuncTestCaseParserIntegerLiteral)

			localctx.(*TimeIntervalContext).seconds = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(450)
			p.Match(FuncTestCaseParserSecondSuffix)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(453)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FuncTestCaseParserIntegerLiteral {
			{
				p.SetState(451)

				var _m = p.Match(FuncTestCaseParserIntegerLiteral)

				localctx.(*TimeIntervalContext).fractionalSeconds = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(452)
				p.Match(FuncTestCaseParserFractionalSecondSuffix)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(455)

			var _m = p.Match(FuncTestCaseParserIntegerLiteral)

			localctx.(*TimeIntervalContext).fractionalSeconds = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(456)
			p.Match(FuncTestCaseParserFractionalSecondSuffix)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDataTypeContext is an interface to support dynamic dispatch.
type IDataTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ScalarType() IScalarTypeContext
	ParameterizedType() IParameterizedTypeContext

	// IsDataTypeContext differentiates from other interfaces.
	IsDataTypeContext()
}

type DataTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataTypeContext() *DataTypeContext {
	var p = new(DataTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_dataType
	return p
}

func InitEmptyDataTypeContext(p *DataTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_dataType
}

func (*DataTypeContext) IsDataTypeContext() {}

func NewDataTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataTypeContext {
	var p = new(DataTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_dataType

	return p
}

func (s *DataTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DataTypeContext) ScalarType() IScalarTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalarTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalarTypeContext)
}

func (s *DataTypeContext) ParameterizedType() IParameterizedTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterizedTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterizedTypeContext)
}

func (s *DataTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterDataType(s)
	}
}

func (s *DataTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitDataType(s)
	}
}

func (s *DataTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitDataType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) DataType() (localctx IDataTypeContext) {
	localctx = NewDataTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, FuncTestCaseParserRULE_dataType)
	p.SetState(461)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FuncTestCaseParserBoolean, FuncTestCaseParserI8, FuncTestCaseParserI16, FuncTestCaseParserI32, FuncTestCaseParserI64, FuncTestCaseParserFP32, FuncTestCaseParserFP64, FuncTestCaseParserString_, FuncTestCaseParserBinary, FuncTestCaseParserTimestamp, FuncTestCaseParserTimestamp_TZ, FuncTestCaseParserDate, FuncTestCaseParserTime, FuncTestCaseParserInterval_Year, FuncTestCaseParserInterval_Day, FuncTestCaseParserUUID, FuncTestCaseParserUserDefined, FuncTestCaseParserBool, FuncTestCaseParserStr, FuncTestCaseParserVBin, FuncTestCaseParserTs, FuncTestCaseParserTsTZ, FuncTestCaseParserIYear, FuncTestCaseParserIDay:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(459)
			p.ScalarType()
		}

	case FuncTestCaseParserDecimal, FuncTestCaseParserPrecision_Timestamp, FuncTestCaseParserPrecision_Timestamp_TZ, FuncTestCaseParserFixedChar, FuncTestCaseParserVarChar, FuncTestCaseParserFixedBinary, FuncTestCaseParserDec, FuncTestCaseParserPTs, FuncTestCaseParserPTsTZ, FuncTestCaseParserFChar, FuncTestCaseParserVChar, FuncTestCaseParserFBin:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(460)
			p.ParameterizedType()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IScalarTypeContext is an interface to support dynamic dispatch.
type IScalarTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsScalarTypeContext differentiates from other interfaces.
	IsScalarTypeContext()
}

type ScalarTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScalarTypeContext() *ScalarTypeContext {
	var p = new(ScalarTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_scalarType
	return p
}

func InitEmptyScalarTypeContext(p *ScalarTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_scalarType
}

func (*ScalarTypeContext) IsScalarTypeContext() {}

func NewScalarTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScalarTypeContext {
	var p = new(ScalarTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_scalarType

	return p
}

func (s *ScalarTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ScalarTypeContext) CopyAll(ctx *ScalarTypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ScalarTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DateContext struct {
	ScalarTypeContext
}

func NewDateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DateContext {
	var p = new(DateContext)

	InitEmptyScalarTypeContext(&p.ScalarTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ScalarTypeContext))

	return p
}

func (s *DateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateContext) Date() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDate, 0)
}

func (s *DateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterDate(s)
	}
}

func (s *DateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitDate(s)
	}
}

func (s *DateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitDate(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringContext struct {
	ScalarTypeContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	InitEmptyScalarTypeContext(&p.ScalarTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ScalarTypeContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) StringType() IStringTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringTypeContext)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitString(s)
	}
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

type I64Context struct {
	ScalarTypeContext
}

func NewI64Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *I64Context {
	var p = new(I64Context)

	InitEmptyScalarTypeContext(&p.ScalarTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ScalarTypeContext))

	return p
}

func (s *I64Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *I64Context) I64() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserI64, 0)
}

func (s *I64Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterI64(s)
	}
}

func (s *I64Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitI64(s)
	}
}

func (s *I64Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitI64(s)

	default:
		return t.VisitChildren(s)
	}
}

type UserDefinedContext struct {
	ScalarTypeContext
}

func NewUserDefinedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UserDefinedContext {
	var p = new(UserDefinedContext)

	InitEmptyScalarTypeContext(&p.ScalarTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ScalarTypeContext))

	return p
}

func (s *UserDefinedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UserDefinedContext) UserDefined() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserUserDefined, 0)
}

func (s *UserDefinedContext) Identifier() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserIdentifier, 0)
}

func (s *UserDefinedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterUserDefined(s)
	}
}

func (s *UserDefinedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitUserDefined(s)
	}
}

func (s *UserDefinedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitUserDefined(s)

	default:
		return t.VisitChildren(s)
	}
}

type I32Context struct {
	ScalarTypeContext
}

func NewI32Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *I32Context {
	var p = new(I32Context)

	InitEmptyScalarTypeContext(&p.ScalarTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ScalarTypeContext))

	return p
}

func (s *I32Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *I32Context) I32() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserI32, 0)
}

func (s *I32Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterI32(s)
	}
}

func (s *I32Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitI32(s)
	}
}

func (s *I32Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitI32(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntervalYearContext struct {
	ScalarTypeContext
}

func NewIntervalYearContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntervalYearContext {
	var p = new(IntervalYearContext)

	InitEmptyScalarTypeContext(&p.ScalarTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ScalarTypeContext))

	return p
}

func (s *IntervalYearContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntervalYearContext) IntervalYearType() IIntervalYearTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntervalYearTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntervalYearTypeContext)
}

func (s *IntervalYearContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterIntervalYear(s)
	}
}

func (s *IntervalYearContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitIntervalYear(s)
	}
}

func (s *IntervalYearContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitIntervalYear(s)

	default:
		return t.VisitChildren(s)
	}
}

type UuidContext struct {
	ScalarTypeContext
}

func NewUuidContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UuidContext {
	var p = new(UuidContext)

	InitEmptyScalarTypeContext(&p.ScalarTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ScalarTypeContext))

	return p
}

func (s *UuidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UuidContext) UUID() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserUUID, 0)
}

func (s *UuidContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterUuid(s)
	}
}

func (s *UuidContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitUuid(s)
	}
}

func (s *UuidContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitUuid(s)

	default:
		return t.VisitChildren(s)
	}
}

type I8Context struct {
	ScalarTypeContext
}

func NewI8Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *I8Context {
	var p = new(I8Context)

	InitEmptyScalarTypeContext(&p.ScalarTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ScalarTypeContext))

	return p
}

func (s *I8Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *I8Context) I8() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserI8, 0)
}

func (s *I8Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterI8(s)
	}
}

func (s *I8Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitI8(s)
	}
}

func (s *I8Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitI8(s)

	default:
		return t.VisitChildren(s)
	}
}

type I16Context struct {
	ScalarTypeContext
}

func NewI16Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *I16Context {
	var p = new(I16Context)

	InitEmptyScalarTypeContext(&p.ScalarTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ScalarTypeContext))

	return p
}

func (s *I16Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *I16Context) I16() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserI16, 0)
}

func (s *I16Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterI16(s)
	}
}

func (s *I16Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitI16(s)
	}
}

func (s *I16Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitI16(s)

	default:
		return t.VisitChildren(s)
	}
}

type BooleanContext struct {
	ScalarTypeContext
}

func NewBooleanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanContext {
	var p = new(BooleanContext)

	InitEmptyScalarTypeContext(&p.ScalarTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ScalarTypeContext))

	return p
}

func (s *BooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanContext) BooleanType() IBooleanTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanTypeContext)
}

func (s *BooleanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterBoolean(s)
	}
}

func (s *BooleanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitBoolean(s)
	}
}

func (s *BooleanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitBoolean(s)

	default:
		return t.VisitChildren(s)
	}
}

type BinaryContext struct {
	ScalarTypeContext
}

func NewBinaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryContext {
	var p = new(BinaryContext)

	InitEmptyScalarTypeContext(&p.ScalarTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ScalarTypeContext))

	return p
}

func (s *BinaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryContext) BinaryType() IBinaryTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinaryTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinaryTypeContext)
}

func (s *BinaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterBinary(s)
	}
}

func (s *BinaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitBinary(s)
	}
}

func (s *BinaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitBinary(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntervalDayContext struct {
	ScalarTypeContext
}

func NewIntervalDayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntervalDayContext {
	var p = new(IntervalDayContext)

	InitEmptyScalarTypeContext(&p.ScalarTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ScalarTypeContext))

	return p
}

func (s *IntervalDayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntervalDayContext) IntervalDayType() IIntervalDayTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntervalDayTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntervalDayTypeContext)
}

func (s *IntervalDayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterIntervalDay(s)
	}
}

func (s *IntervalDayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitIntervalDay(s)
	}
}

func (s *IntervalDayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitIntervalDay(s)

	default:
		return t.VisitChildren(s)
	}
}

type Fp64Context struct {
	ScalarTypeContext
}

func NewFp64Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *Fp64Context {
	var p = new(Fp64Context)

	InitEmptyScalarTypeContext(&p.ScalarTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ScalarTypeContext))

	return p
}

func (s *Fp64Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fp64Context) FP64() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserFP64, 0)
}

func (s *Fp64Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterFp64(s)
	}
}

func (s *Fp64Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitFp64(s)
	}
}

func (s *Fp64Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitFp64(s)

	default:
		return t.VisitChildren(s)
	}
}

type Fp32Context struct {
	ScalarTypeContext
}

func NewFp32Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *Fp32Context {
	var p = new(Fp32Context)

	InitEmptyScalarTypeContext(&p.ScalarTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ScalarTypeContext))

	return p
}

func (s *Fp32Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fp32Context) FP32() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserFP32, 0)
}

func (s *Fp32Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterFp32(s)
	}
}

func (s *Fp32Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitFp32(s)
	}
}

func (s *Fp32Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitFp32(s)

	default:
		return t.VisitChildren(s)
	}
}

type TimeContext struct {
	ScalarTypeContext
}

func NewTimeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TimeContext {
	var p = new(TimeContext)

	InitEmptyScalarTypeContext(&p.ScalarTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ScalarTypeContext))

	return p
}

func (s *TimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeContext) Time() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserTime, 0)
}

func (s *TimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterTime(s)
	}
}

func (s *TimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitTime(s)
	}
}

func (s *TimeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitTime(s)

	default:
		return t.VisitChildren(s)
	}
}

type TimestampContext struct {
	ScalarTypeContext
}

func NewTimestampContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TimestampContext {
	var p = new(TimestampContext)

	InitEmptyScalarTypeContext(&p.ScalarTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ScalarTypeContext))

	return p
}

func (s *TimestampContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimestampContext) TimestampType() ITimestampTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimestampTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimestampTypeContext)
}

func (s *TimestampContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterTimestamp(s)
	}
}

func (s *TimestampContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitTimestamp(s)
	}
}

func (s *TimestampContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitTimestamp(s)

	default:
		return t.VisitChildren(s)
	}
}

type TimestampTzContext struct {
	ScalarTypeContext
}

func NewTimestampTzContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TimestampTzContext {
	var p = new(TimestampTzContext)

	InitEmptyScalarTypeContext(&p.ScalarTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ScalarTypeContext))

	return p
}

func (s *TimestampTzContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimestampTzContext) TimestampTZType() ITimestampTZTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimestampTZTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimestampTZTypeContext)
}

func (s *TimestampTzContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterTimestampTz(s)
	}
}

func (s *TimestampTzContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitTimestampTz(s)
	}
}

func (s *TimestampTzContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitTimestampTz(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) ScalarType() (localctx IScalarTypeContext) {
	localctx = NewScalarTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, FuncTestCaseParserRULE_scalarType)
	p.SetState(481)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FuncTestCaseParserBoolean, FuncTestCaseParserBool:
		localctx = NewBooleanContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(463)
			p.BooleanType()
		}

	case FuncTestCaseParserI8:
		localctx = NewI8Context(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(464)
			p.Match(FuncTestCaseParserI8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserI16:
		localctx = NewI16Context(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(465)
			p.Match(FuncTestCaseParserI16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserI32:
		localctx = NewI32Context(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(466)
			p.Match(FuncTestCaseParserI32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserI64:
		localctx = NewI64Context(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(467)
			p.Match(FuncTestCaseParserI64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserFP32:
		localctx = NewFp32Context(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(468)
			p.Match(FuncTestCaseParserFP32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserFP64:
		localctx = NewFp64Context(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(469)
			p.Match(FuncTestCaseParserFP64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserString_, FuncTestCaseParserStr:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(470)
			p.StringType()
		}

	case FuncTestCaseParserBinary, FuncTestCaseParserVBin:
		localctx = NewBinaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(471)
			p.BinaryType()
		}

	case FuncTestCaseParserTimestamp, FuncTestCaseParserTs:
		localctx = NewTimestampContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(472)
			p.TimestampType()
		}

	case FuncTestCaseParserTimestamp_TZ, FuncTestCaseParserTsTZ:
		localctx = NewTimestampTzContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(473)
			p.TimestampTZType()
		}

	case FuncTestCaseParserDate:
		localctx = NewDateContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(474)
			p.Match(FuncTestCaseParserDate)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserTime:
		localctx = NewTimeContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(475)
			p.Match(FuncTestCaseParserTime)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserInterval_Day, FuncTestCaseParserIDay:
		localctx = NewIntervalDayContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(476)
			p.IntervalDayType()
		}

	case FuncTestCaseParserInterval_Year, FuncTestCaseParserIYear:
		localctx = NewIntervalYearContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(477)
			p.IntervalYearType()
		}

	case FuncTestCaseParserUUID:
		localctx = NewUuidContext(p, localctx)
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(478)
			p.Match(FuncTestCaseParserUUID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserUserDefined:
		localctx = NewUserDefinedContext(p, localctx)
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(479)
			p.Match(FuncTestCaseParserUserDefined)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(480)
			p.Match(FuncTestCaseParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBooleanTypeContext is an interface to support dynamic dispatch.
type IBooleanTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Bool() antlr.TerminalNode
	Boolean() antlr.TerminalNode

	// IsBooleanTypeContext differentiates from other interfaces.
	IsBooleanTypeContext()
}

type BooleanTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanTypeContext() *BooleanTypeContext {
	var p = new(BooleanTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_booleanType
	return p
}

func InitEmptyBooleanTypeContext(p *BooleanTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_booleanType
}

func (*BooleanTypeContext) IsBooleanTypeContext() {}

func NewBooleanTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanTypeContext {
	var p = new(BooleanTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_booleanType

	return p
}

func (s *BooleanTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanTypeContext) Bool() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserBool, 0)
}

func (s *BooleanTypeContext) Boolean() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserBoolean, 0)
}

func (s *BooleanTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterBooleanType(s)
	}
}

func (s *BooleanTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitBooleanType(s)
	}
}

func (s *BooleanTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitBooleanType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) BooleanType() (localctx IBooleanTypeContext) {
	localctx = NewBooleanTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, FuncTestCaseParserRULE_booleanType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(483)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserBoolean || _la == FuncTestCaseParserBool) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStringTypeContext is an interface to support dynamic dispatch.
type IStringTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Str() antlr.TerminalNode
	String_() antlr.TerminalNode

	// IsStringTypeContext differentiates from other interfaces.
	IsStringTypeContext()
}

type StringTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringTypeContext() *StringTypeContext {
	var p = new(StringTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_stringType
	return p
}

func InitEmptyStringTypeContext(p *StringTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_stringType
}

func (*StringTypeContext) IsStringTypeContext() {}

func NewStringTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringTypeContext {
	var p = new(StringTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_stringType

	return p
}

func (s *StringTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *StringTypeContext) Str() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserStr, 0)
}

func (s *StringTypeContext) String_() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserString_, 0)
}

func (s *StringTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterStringType(s)
	}
}

func (s *StringTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitStringType(s)
	}
}

func (s *StringTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitStringType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) StringType() (localctx IStringTypeContext) {
	localctx = NewStringTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, FuncTestCaseParserRULE_stringType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(485)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserString_ || _la == FuncTestCaseParserStr) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBinaryTypeContext is an interface to support dynamic dispatch.
type IBinaryTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Binary() antlr.TerminalNode
	VBin() antlr.TerminalNode

	// IsBinaryTypeContext differentiates from other interfaces.
	IsBinaryTypeContext()
}

type BinaryTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryTypeContext() *BinaryTypeContext {
	var p = new(BinaryTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_binaryType
	return p
}

func InitEmptyBinaryTypeContext(p *BinaryTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_binaryType
}

func (*BinaryTypeContext) IsBinaryTypeContext() {}

func NewBinaryTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryTypeContext {
	var p = new(BinaryTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_binaryType

	return p
}

func (s *BinaryTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryTypeContext) Binary() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserBinary, 0)
}

func (s *BinaryTypeContext) VBin() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserVBin, 0)
}

func (s *BinaryTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterBinaryType(s)
	}
}

func (s *BinaryTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitBinaryType(s)
	}
}

func (s *BinaryTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitBinaryType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) BinaryType() (localctx IBinaryTypeContext) {
	localctx = NewBinaryTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, FuncTestCaseParserRULE_binaryType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(487)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserBinary || _la == FuncTestCaseParserVBin) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITimestampTypeContext is an interface to support dynamic dispatch.
type ITimestampTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ts() antlr.TerminalNode
	Timestamp() antlr.TerminalNode

	// IsTimestampTypeContext differentiates from other interfaces.
	IsTimestampTypeContext()
}

type TimestampTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimestampTypeContext() *TimestampTypeContext {
	var p = new(TimestampTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_timestampType
	return p
}

func InitEmptyTimestampTypeContext(p *TimestampTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_timestampType
}

func (*TimestampTypeContext) IsTimestampTypeContext() {}

func NewTimestampTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimestampTypeContext {
	var p = new(TimestampTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_timestampType

	return p
}

func (s *TimestampTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TimestampTypeContext) Ts() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserTs, 0)
}

func (s *TimestampTypeContext) Timestamp() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserTimestamp, 0)
}

func (s *TimestampTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimestampTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimestampTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterTimestampType(s)
	}
}

func (s *TimestampTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitTimestampType(s)
	}
}

func (s *TimestampTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitTimestampType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) TimestampType() (localctx ITimestampTypeContext) {
	localctx = NewTimestampTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, FuncTestCaseParserRULE_timestampType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(489)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserTimestamp || _la == FuncTestCaseParserTs) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITimestampTZTypeContext is an interface to support dynamic dispatch.
type ITimestampTZTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TsTZ() antlr.TerminalNode
	Timestamp_TZ() antlr.TerminalNode

	// IsTimestampTZTypeContext differentiates from other interfaces.
	IsTimestampTZTypeContext()
}

type TimestampTZTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimestampTZTypeContext() *TimestampTZTypeContext {
	var p = new(TimestampTZTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_timestampTZType
	return p
}

func InitEmptyTimestampTZTypeContext(p *TimestampTZTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_timestampTZType
}

func (*TimestampTZTypeContext) IsTimestampTZTypeContext() {}

func NewTimestampTZTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimestampTZTypeContext {
	var p = new(TimestampTZTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_timestampTZType

	return p
}

func (s *TimestampTZTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TimestampTZTypeContext) TsTZ() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserTsTZ, 0)
}

func (s *TimestampTZTypeContext) Timestamp_TZ() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserTimestamp_TZ, 0)
}

func (s *TimestampTZTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimestampTZTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimestampTZTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterTimestampTZType(s)
	}
}

func (s *TimestampTZTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitTimestampTZType(s)
	}
}

func (s *TimestampTZTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitTimestampTZType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) TimestampTZType() (localctx ITimestampTZTypeContext) {
	localctx = NewTimestampTZTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, FuncTestCaseParserRULE_timestampTZType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(491)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserTimestamp_TZ || _la == FuncTestCaseParserTsTZ) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIntervalYearTypeContext is an interface to support dynamic dispatch.
type IIntervalYearTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IYear() antlr.TerminalNode
	Interval_Year() antlr.TerminalNode

	// IsIntervalYearTypeContext differentiates from other interfaces.
	IsIntervalYearTypeContext()
}

type IntervalYearTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntervalYearTypeContext() *IntervalYearTypeContext {
	var p = new(IntervalYearTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_intervalYearType
	return p
}

func InitEmptyIntervalYearTypeContext(p *IntervalYearTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_intervalYearType
}

func (*IntervalYearTypeContext) IsIntervalYearTypeContext() {}

func NewIntervalYearTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntervalYearTypeContext {
	var p = new(IntervalYearTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_intervalYearType

	return p
}

func (s *IntervalYearTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *IntervalYearTypeContext) IYear() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserIYear, 0)
}

func (s *IntervalYearTypeContext) Interval_Year() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserInterval_Year, 0)
}

func (s *IntervalYearTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntervalYearTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntervalYearTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterIntervalYearType(s)
	}
}

func (s *IntervalYearTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitIntervalYearType(s)
	}
}

func (s *IntervalYearTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitIntervalYearType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) IntervalYearType() (localctx IIntervalYearTypeContext) {
	localctx = NewIntervalYearTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, FuncTestCaseParserRULE_intervalYearType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(493)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserInterval_Year || _la == FuncTestCaseParserIYear) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIntervalDayTypeContext is an interface to support dynamic dispatch.
type IIntervalDayTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDay() antlr.TerminalNode
	Interval_Day() antlr.TerminalNode

	// IsIntervalDayTypeContext differentiates from other interfaces.
	IsIntervalDayTypeContext()
}

type IntervalDayTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntervalDayTypeContext() *IntervalDayTypeContext {
	var p = new(IntervalDayTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_intervalDayType
	return p
}

func InitEmptyIntervalDayTypeContext(p *IntervalDayTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_intervalDayType
}

func (*IntervalDayTypeContext) IsIntervalDayTypeContext() {}

func NewIntervalDayTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntervalDayTypeContext {
	var p = new(IntervalDayTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_intervalDayType

	return p
}

func (s *IntervalDayTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *IntervalDayTypeContext) IDay() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserIDay, 0)
}

func (s *IntervalDayTypeContext) Interval_Day() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserInterval_Day, 0)
}

func (s *IntervalDayTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntervalDayTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntervalDayTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterIntervalDayType(s)
	}
}

func (s *IntervalDayTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitIntervalDayType(s)
	}
}

func (s *IntervalDayTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitIntervalDayType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) IntervalDayType() (localctx IIntervalDayTypeContext) {
	localctx = NewIntervalDayTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, FuncTestCaseParserRULE_intervalDayType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(495)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserInterval_Day || _la == FuncTestCaseParserIDay) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFixedCharTypeContext is an interface to support dynamic dispatch.
type IFixedCharTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFixedCharTypeContext differentiates from other interfaces.
	IsFixedCharTypeContext()
}

type FixedCharTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFixedCharTypeContext() *FixedCharTypeContext {
	var p = new(FixedCharTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_fixedCharType
	return p
}

func InitEmptyFixedCharTypeContext(p *FixedCharTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_fixedCharType
}

func (*FixedCharTypeContext) IsFixedCharTypeContext() {}

func NewFixedCharTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FixedCharTypeContext {
	var p = new(FixedCharTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_fixedCharType

	return p
}

func (s *FixedCharTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FixedCharTypeContext) CopyAll(ctx *FixedCharTypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FixedCharTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FixedCharTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FixedCharContext struct {
	FixedCharTypeContext
	isnull antlr.Token
	len_   INumericParameterContext
}

func NewFixedCharContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FixedCharContext {
	var p = new(FixedCharContext)

	InitEmptyFixedCharTypeContext(&p.FixedCharTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*FixedCharTypeContext))

	return p
}

func (s *FixedCharContext) GetIsnull() antlr.Token { return s.isnull }

func (s *FixedCharContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *FixedCharContext) GetLen_() INumericParameterContext { return s.len_ }

func (s *FixedCharContext) SetLen_(v INumericParameterContext) { s.len_ = v }

func (s *FixedCharContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FixedCharContext) OAngleBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserOAngleBracket, 0)
}

func (s *FixedCharContext) CAngleBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserCAngleBracket, 0)
}

func (s *FixedCharContext) FChar() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserFChar, 0)
}

func (s *FixedCharContext) FixedChar() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserFixedChar, 0)
}

func (s *FixedCharContext) NumericParameter() INumericParameterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericParameterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericParameterContext)
}

func (s *FixedCharContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *FixedCharContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterFixedChar(s)
	}
}

func (s *FixedCharContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitFixedChar(s)
	}
}

func (s *FixedCharContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitFixedChar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) FixedCharType() (localctx IFixedCharTypeContext) {
	localctx = NewFixedCharTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, FuncTestCaseParserRULE_fixedCharType)
	var _la int

	localctx = NewFixedCharContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(497)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserFixedChar || _la == FuncTestCaseParserFChar) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(499)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(498)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*FixedCharContext).isnull = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(501)
		p.Match(FuncTestCaseParserOAngleBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(502)

		var _x = p.NumericParameter()

		localctx.(*FixedCharContext).len_ = _x
	}
	{
		p.SetState(503)
		p.Match(FuncTestCaseParserCAngleBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarCharTypeContext is an interface to support dynamic dispatch.
type IVarCharTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVarCharTypeContext differentiates from other interfaces.
	IsVarCharTypeContext()
}

type VarCharTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarCharTypeContext() *VarCharTypeContext {
	var p = new(VarCharTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_varCharType
	return p
}

func InitEmptyVarCharTypeContext(p *VarCharTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_varCharType
}

func (*VarCharTypeContext) IsVarCharTypeContext() {}

func NewVarCharTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarCharTypeContext {
	var p = new(VarCharTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_varCharType

	return p
}

func (s *VarCharTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *VarCharTypeContext) CopyAll(ctx *VarCharTypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *VarCharTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarCharTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VarCharContext struct {
	VarCharTypeContext
	isnull antlr.Token
	len_   INumericParameterContext
}

func NewVarCharContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarCharContext {
	var p = new(VarCharContext)

	InitEmptyVarCharTypeContext(&p.VarCharTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarCharTypeContext))

	return p
}

func (s *VarCharContext) GetIsnull() antlr.Token { return s.isnull }

func (s *VarCharContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *VarCharContext) GetLen_() INumericParameterContext { return s.len_ }

func (s *VarCharContext) SetLen_(v INumericParameterContext) { s.len_ = v }

func (s *VarCharContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarCharContext) OAngleBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserOAngleBracket, 0)
}

func (s *VarCharContext) CAngleBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserCAngleBracket, 0)
}

func (s *VarCharContext) VChar() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserVChar, 0)
}

func (s *VarCharContext) VarChar() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserVarChar, 0)
}

func (s *VarCharContext) NumericParameter() INumericParameterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericParameterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericParameterContext)
}

func (s *VarCharContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *VarCharContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterVarChar(s)
	}
}

func (s *VarCharContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitVarChar(s)
	}
}

func (s *VarCharContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitVarChar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) VarCharType() (localctx IVarCharTypeContext) {
	localctx = NewVarCharTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, FuncTestCaseParserRULE_varCharType)
	var _la int

	localctx = NewVarCharContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(505)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserVarChar || _la == FuncTestCaseParserVChar) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(507)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(506)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*VarCharContext).isnull = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(509)
		p.Match(FuncTestCaseParserOAngleBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(510)

		var _x = p.NumericParameter()

		localctx.(*VarCharContext).len_ = _x
	}
	{
		p.SetState(511)
		p.Match(FuncTestCaseParserCAngleBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFixedBinaryTypeContext is an interface to support dynamic dispatch.
type IFixedBinaryTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFixedBinaryTypeContext differentiates from other interfaces.
	IsFixedBinaryTypeContext()
}

type FixedBinaryTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFixedBinaryTypeContext() *FixedBinaryTypeContext {
	var p = new(FixedBinaryTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_fixedBinaryType
	return p
}

func InitEmptyFixedBinaryTypeContext(p *FixedBinaryTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_fixedBinaryType
}

func (*FixedBinaryTypeContext) IsFixedBinaryTypeContext() {}

func NewFixedBinaryTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FixedBinaryTypeContext {
	var p = new(FixedBinaryTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_fixedBinaryType

	return p
}

func (s *FixedBinaryTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FixedBinaryTypeContext) CopyAll(ctx *FixedBinaryTypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FixedBinaryTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FixedBinaryTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FixedBinaryContext struct {
	FixedBinaryTypeContext
	isnull antlr.Token
	len_   INumericParameterContext
}

func NewFixedBinaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FixedBinaryContext {
	var p = new(FixedBinaryContext)

	InitEmptyFixedBinaryTypeContext(&p.FixedBinaryTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*FixedBinaryTypeContext))

	return p
}

func (s *FixedBinaryContext) GetIsnull() antlr.Token { return s.isnull }

func (s *FixedBinaryContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *FixedBinaryContext) GetLen_() INumericParameterContext { return s.len_ }

func (s *FixedBinaryContext) SetLen_(v INumericParameterContext) { s.len_ = v }

func (s *FixedBinaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FixedBinaryContext) OAngleBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserOAngleBracket, 0)
}

func (s *FixedBinaryContext) CAngleBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserCAngleBracket, 0)
}

func (s *FixedBinaryContext) FBin() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserFBin, 0)
}

func (s *FixedBinaryContext) FixedBinary() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserFixedBinary, 0)
}

func (s *FixedBinaryContext) NumericParameter() INumericParameterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericParameterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericParameterContext)
}

func (s *FixedBinaryContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *FixedBinaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterFixedBinary(s)
	}
}

func (s *FixedBinaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitFixedBinary(s)
	}
}

func (s *FixedBinaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitFixedBinary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) FixedBinaryType() (localctx IFixedBinaryTypeContext) {
	localctx = NewFixedBinaryTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, FuncTestCaseParserRULE_fixedBinaryType)
	var _la int

	localctx = NewFixedBinaryContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(513)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserFixedBinary || _la == FuncTestCaseParserFBin) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(515)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(514)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*FixedBinaryContext).isnull = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(517)
		p.Match(FuncTestCaseParserOAngleBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(518)

		var _x = p.NumericParameter()

		localctx.(*FixedBinaryContext).len_ = _x
	}
	{
		p.SetState(519)
		p.Match(FuncTestCaseParserCAngleBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDecimalTypeContext is an interface to support dynamic dispatch.
type IDecimalTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDecimalTypeContext differentiates from other interfaces.
	IsDecimalTypeContext()
}

type DecimalTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecimalTypeContext() *DecimalTypeContext {
	var p = new(DecimalTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_decimalType
	return p
}

func InitEmptyDecimalTypeContext(p *DecimalTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_decimalType
}

func (*DecimalTypeContext) IsDecimalTypeContext() {}

func NewDecimalTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecimalTypeContext {
	var p = new(DecimalTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_decimalType

	return p
}

func (s *DecimalTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DecimalTypeContext) CopyAll(ctx *DecimalTypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DecimalTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DecimalContext struct {
	DecimalTypeContext
	isnull    antlr.Token
	precision INumericParameterContext
	scale     INumericParameterContext
}

func NewDecimalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DecimalContext {
	var p = new(DecimalContext)

	InitEmptyDecimalTypeContext(&p.DecimalTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*DecimalTypeContext))

	return p
}

func (s *DecimalContext) GetIsnull() antlr.Token { return s.isnull }

func (s *DecimalContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *DecimalContext) GetPrecision() INumericParameterContext { return s.precision }

func (s *DecimalContext) GetScale() INumericParameterContext { return s.scale }

func (s *DecimalContext) SetPrecision(v INumericParameterContext) { s.precision = v }

func (s *DecimalContext) SetScale(v INumericParameterContext) { s.scale = v }

func (s *DecimalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalContext) Dec() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDec, 0)
}

func (s *DecimalContext) Decimal() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDecimal, 0)
}

func (s *DecimalContext) OAngleBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserOAngleBracket, 0)
}

func (s *DecimalContext) Comma() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserComma, 0)
}

func (s *DecimalContext) CAngleBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserCAngleBracket, 0)
}

func (s *DecimalContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *DecimalContext) AllNumericParameter() []INumericParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INumericParameterContext); ok {
			len++
		}
	}

	tst := make([]INumericParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INumericParameterContext); ok {
			tst[i] = t.(INumericParameterContext)
			i++
		}
	}

	return tst
}

func (s *DecimalContext) NumericParameter(i int) INumericParameterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericParameterContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericParameterContext)
}

func (s *DecimalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterDecimal(s)
	}
}

func (s *DecimalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitDecimal(s)
	}
}

func (s *DecimalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitDecimal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) DecimalType() (localctx IDecimalTypeContext) {
	localctx = NewDecimalTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, FuncTestCaseParserRULE_decimalType)
	var _la int

	localctx = NewDecimalContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(521)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserDecimal || _la == FuncTestCaseParserDec) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(523)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(522)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*DecimalContext).isnull = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(531)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserOAngleBracket {
		{
			p.SetState(525)
			p.Match(FuncTestCaseParserOAngleBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(526)

			var _x = p.NumericParameter()

			localctx.(*DecimalContext).precision = _x
		}
		{
			p.SetState(527)
			p.Match(FuncTestCaseParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(528)

			var _x = p.NumericParameter()

			localctx.(*DecimalContext).scale = _x
		}
		{
			p.SetState(529)
			p.Match(FuncTestCaseParserCAngleBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrecisionTimestampTypeContext is an interface to support dynamic dispatch.
type IPrecisionTimestampTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrecisionTimestampTypeContext differentiates from other interfaces.
	IsPrecisionTimestampTypeContext()
}

type PrecisionTimestampTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrecisionTimestampTypeContext() *PrecisionTimestampTypeContext {
	var p = new(PrecisionTimestampTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_precisionTimestampType
	return p
}

func InitEmptyPrecisionTimestampTypeContext(p *PrecisionTimestampTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_precisionTimestampType
}

func (*PrecisionTimestampTypeContext) IsPrecisionTimestampTypeContext() {}

func NewPrecisionTimestampTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrecisionTimestampTypeContext {
	var p = new(PrecisionTimestampTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_precisionTimestampType

	return p
}

func (s *PrecisionTimestampTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *PrecisionTimestampTypeContext) CopyAll(ctx *PrecisionTimestampTypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PrecisionTimestampTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrecisionTimestampTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PrecisionTimestampContext struct {
	PrecisionTimestampTypeContext
	isnull    antlr.Token
	precision INumericParameterContext
}

func NewPrecisionTimestampContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrecisionTimestampContext {
	var p = new(PrecisionTimestampContext)

	InitEmptyPrecisionTimestampTypeContext(&p.PrecisionTimestampTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrecisionTimestampTypeContext))

	return p
}

func (s *PrecisionTimestampContext) GetIsnull() antlr.Token { return s.isnull }

func (s *PrecisionTimestampContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *PrecisionTimestampContext) GetPrecision() INumericParameterContext { return s.precision }

func (s *PrecisionTimestampContext) SetPrecision(v INumericParameterContext) { s.precision = v }

func (s *PrecisionTimestampContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrecisionTimestampContext) OAngleBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserOAngleBracket, 0)
}

func (s *PrecisionTimestampContext) CAngleBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserCAngleBracket, 0)
}

func (s *PrecisionTimestampContext) PTs() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserPTs, 0)
}

func (s *PrecisionTimestampContext) Precision_Timestamp() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserPrecision_Timestamp, 0)
}

func (s *PrecisionTimestampContext) NumericParameter() INumericParameterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericParameterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericParameterContext)
}

func (s *PrecisionTimestampContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *PrecisionTimestampContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterPrecisionTimestamp(s)
	}
}

func (s *PrecisionTimestampContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitPrecisionTimestamp(s)
	}
}

func (s *PrecisionTimestampContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitPrecisionTimestamp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) PrecisionTimestampType() (localctx IPrecisionTimestampTypeContext) {
	localctx = NewPrecisionTimestampTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, FuncTestCaseParserRULE_precisionTimestampType)
	var _la int

	localctx = NewPrecisionTimestampContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(533)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserPrecision_Timestamp || _la == FuncTestCaseParserPTs) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(535)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(534)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*PrecisionTimestampContext).isnull = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(537)
		p.Match(FuncTestCaseParserOAngleBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(538)

		var _x = p.NumericParameter()

		localctx.(*PrecisionTimestampContext).precision = _x
	}
	{
		p.SetState(539)
		p.Match(FuncTestCaseParserCAngleBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrecisionTimestampTZTypeContext is an interface to support dynamic dispatch.
type IPrecisionTimestampTZTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrecisionTimestampTZTypeContext differentiates from other interfaces.
	IsPrecisionTimestampTZTypeContext()
}

type PrecisionTimestampTZTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrecisionTimestampTZTypeContext() *PrecisionTimestampTZTypeContext {
	var p = new(PrecisionTimestampTZTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_precisionTimestampTZType
	return p
}

func InitEmptyPrecisionTimestampTZTypeContext(p *PrecisionTimestampTZTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_precisionTimestampTZType
}

func (*PrecisionTimestampTZTypeContext) IsPrecisionTimestampTZTypeContext() {}

func NewPrecisionTimestampTZTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrecisionTimestampTZTypeContext {
	var p = new(PrecisionTimestampTZTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_precisionTimestampTZType

	return p
}

func (s *PrecisionTimestampTZTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *PrecisionTimestampTZTypeContext) CopyAll(ctx *PrecisionTimestampTZTypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PrecisionTimestampTZTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrecisionTimestampTZTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PrecisionTimestampTZContext struct {
	PrecisionTimestampTZTypeContext
	isnull    antlr.Token
	precision INumericParameterContext
}

func NewPrecisionTimestampTZContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrecisionTimestampTZContext {
	var p = new(PrecisionTimestampTZContext)

	InitEmptyPrecisionTimestampTZTypeContext(&p.PrecisionTimestampTZTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrecisionTimestampTZTypeContext))

	return p
}

func (s *PrecisionTimestampTZContext) GetIsnull() antlr.Token { return s.isnull }

func (s *PrecisionTimestampTZContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *PrecisionTimestampTZContext) GetPrecision() INumericParameterContext { return s.precision }

func (s *PrecisionTimestampTZContext) SetPrecision(v INumericParameterContext) { s.precision = v }

func (s *PrecisionTimestampTZContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrecisionTimestampTZContext) OAngleBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserOAngleBracket, 0)
}

func (s *PrecisionTimestampTZContext) CAngleBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserCAngleBracket, 0)
}

func (s *PrecisionTimestampTZContext) PTsTZ() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserPTsTZ, 0)
}

func (s *PrecisionTimestampTZContext) Precision_Timestamp_TZ() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserPrecision_Timestamp_TZ, 0)
}

func (s *PrecisionTimestampTZContext) NumericParameter() INumericParameterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericParameterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericParameterContext)
}

func (s *PrecisionTimestampTZContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *PrecisionTimestampTZContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterPrecisionTimestampTZ(s)
	}
}

func (s *PrecisionTimestampTZContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitPrecisionTimestampTZ(s)
	}
}

func (s *PrecisionTimestampTZContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitPrecisionTimestampTZ(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) PrecisionTimestampTZType() (localctx IPrecisionTimestampTZTypeContext) {
	localctx = NewPrecisionTimestampTZTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, FuncTestCaseParserRULE_precisionTimestampTZType)
	var _la int

	localctx = NewPrecisionTimestampTZContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(541)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserPrecision_Timestamp_TZ || _la == FuncTestCaseParserPTsTZ) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(543)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(542)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*PrecisionTimestampTZContext).isnull = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(545)
		p.Match(FuncTestCaseParserOAngleBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(546)

		var _x = p.NumericParameter()

		localctx.(*PrecisionTimestampTZContext).precision = _x
	}
	{
		p.SetState(547)
		p.Match(FuncTestCaseParserCAngleBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListTypeContext is an interface to support dynamic dispatch.
type IListTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsListTypeContext differentiates from other interfaces.
	IsListTypeContext()
}

type ListTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListTypeContext() *ListTypeContext {
	var p = new(ListTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_listType
	return p
}

func InitEmptyListTypeContext(p *ListTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_listType
}

func (*ListTypeContext) IsListTypeContext() {}

func NewListTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListTypeContext {
	var p = new(ListTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_listType

	return p
}

func (s *ListTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ListTypeContext) CopyAll(ctx *ListTypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ListTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ListContext struct {
	ListTypeContext
	isnull   antlr.Token
	elemType IDataTypeContext
}

func NewListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListContext {
	var p = new(ListContext)

	InitEmptyListTypeContext(&p.ListTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ListTypeContext))

	return p
}

func (s *ListContext) GetIsnull() antlr.Token { return s.isnull }

func (s *ListContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *ListContext) GetElemType() IDataTypeContext { return s.elemType }

func (s *ListContext) SetElemType(v IDataTypeContext) { s.elemType = v }

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) List() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserList, 0)
}

func (s *ListContext) OAngleBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserOAngleBracket, 0)
}

func (s *ListContext) CAngleBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserCAngleBracket, 0)
}

func (s *ListContext) DataType() IDataTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataTypeContext)
}

func (s *ListContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitList(s)
	}
}

func (s *ListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) ListType() (localctx IListTypeContext) {
	localctx = NewListTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, FuncTestCaseParserRULE_listType)
	var _la int

	localctx = NewListContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(549)
		p.Match(FuncTestCaseParserList)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(551)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(550)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*ListContext).isnull = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(553)
		p.Match(FuncTestCaseParserOAngleBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(554)

		var _x = p.DataType()

		localctx.(*ListContext).elemType = _x
	}
	{
		p.SetState(555)
		p.Match(FuncTestCaseParserCAngleBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParameterizedTypeContext is an interface to support dynamic dispatch.
type IParameterizedTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FixedCharType() IFixedCharTypeContext
	VarCharType() IVarCharTypeContext
	FixedBinaryType() IFixedBinaryTypeContext
	DecimalType() IDecimalTypeContext
	PrecisionTimestampType() IPrecisionTimestampTypeContext
	PrecisionTimestampTZType() IPrecisionTimestampTZTypeContext

	// IsParameterizedTypeContext differentiates from other interfaces.
	IsParameterizedTypeContext()
}

type ParameterizedTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterizedTypeContext() *ParameterizedTypeContext {
	var p = new(ParameterizedTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_parameterizedType
	return p
}

func InitEmptyParameterizedTypeContext(p *ParameterizedTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_parameterizedType
}

func (*ParameterizedTypeContext) IsParameterizedTypeContext() {}

func NewParameterizedTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterizedTypeContext {
	var p = new(ParameterizedTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_parameterizedType

	return p
}

func (s *ParameterizedTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterizedTypeContext) FixedCharType() IFixedCharTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFixedCharTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFixedCharTypeContext)
}

func (s *ParameterizedTypeContext) VarCharType() IVarCharTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarCharTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarCharTypeContext)
}

func (s *ParameterizedTypeContext) FixedBinaryType() IFixedBinaryTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFixedBinaryTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFixedBinaryTypeContext)
}

func (s *ParameterizedTypeContext) DecimalType() IDecimalTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecimalTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecimalTypeContext)
}

func (s *ParameterizedTypeContext) PrecisionTimestampType() IPrecisionTimestampTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrecisionTimestampTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrecisionTimestampTypeContext)
}

func (s *ParameterizedTypeContext) PrecisionTimestampTZType() IPrecisionTimestampTZTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrecisionTimestampTZTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrecisionTimestampTZTypeContext)
}

func (s *ParameterizedTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterizedTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterizedTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterParameterizedType(s)
	}
}

func (s *ParameterizedTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitParameterizedType(s)
	}
}

func (s *ParameterizedTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitParameterizedType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) ParameterizedType() (localctx IParameterizedTypeContext) {
	localctx = NewParameterizedTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, FuncTestCaseParserRULE_parameterizedType)
	p.SetState(563)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FuncTestCaseParserFixedChar, FuncTestCaseParserFChar:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(557)
			p.FixedCharType()
		}

	case FuncTestCaseParserVarChar, FuncTestCaseParserVChar:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(558)
			p.VarCharType()
		}

	case FuncTestCaseParserFixedBinary, FuncTestCaseParserFBin:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(559)
			p.FixedBinaryType()
		}

	case FuncTestCaseParserDecimal, FuncTestCaseParserDec:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(560)
			p.DecimalType()
		}

	case FuncTestCaseParserPrecision_Timestamp, FuncTestCaseParserPTs:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(561)
			p.PrecisionTimestampType()
		}

	case FuncTestCaseParserPrecision_Timestamp_TZ, FuncTestCaseParserPTsTZ:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(562)
			p.PrecisionTimestampTZType()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumericParameterContext is an interface to support dynamic dispatch.
type INumericParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsNumericParameterContext differentiates from other interfaces.
	IsNumericParameterContext()
}

type NumericParameterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericParameterContext() *NumericParameterContext {
	var p = new(NumericParameterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_numericParameter
	return p
}

func InitEmptyNumericParameterContext(p *NumericParameterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_numericParameter
}

func (*NumericParameterContext) IsNumericParameterContext() {}

func NewNumericParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericParameterContext {
	var p = new(NumericParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_numericParameter

	return p
}

func (s *NumericParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericParameterContext) CopyAll(ctx *NumericParameterContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *NumericParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IntegerLiteralContext struct {
	NumericParameterContext
}

func NewIntegerLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)

	InitEmptyNumericParameterContext(&p.NumericParameterContext)
	p.parser = parser
	p.CopyAll(ctx.(*NumericParameterContext))

	return p
}

func (s *IntegerLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerLiteralContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserIntegerLiteral, 0)
}

func (s *IntegerLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterIntegerLiteral(s)
	}
}

func (s *IntegerLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitIntegerLiteral(s)
	}
}

func (s *IntegerLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitIntegerLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) NumericParameter() (localctx INumericParameterContext) {
	localctx = NewNumericParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, FuncTestCaseParserRULE_numericParameter)
	localctx = NewIntegerLiteralContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(565)
		p.Match(FuncTestCaseParserIntegerLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISubstraitErrorContext is an interface to support dynamic dispatch.
type ISubstraitErrorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ErrorResult() antlr.TerminalNode
	UndefineResult() antlr.TerminalNode

	// IsSubstraitErrorContext differentiates from other interfaces.
	IsSubstraitErrorContext()
}

type SubstraitErrorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubstraitErrorContext() *SubstraitErrorContext {
	var p = new(SubstraitErrorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_substraitError
	return p
}

func InitEmptySubstraitErrorContext(p *SubstraitErrorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_substraitError
}

func (*SubstraitErrorContext) IsSubstraitErrorContext() {}

func NewSubstraitErrorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubstraitErrorContext {
	var p = new(SubstraitErrorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_substraitError

	return p
}

func (s *SubstraitErrorContext) GetParser() antlr.Parser { return s.parser }

func (s *SubstraitErrorContext) ErrorResult() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserErrorResult, 0)
}

func (s *SubstraitErrorContext) UndefineResult() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserUndefineResult, 0)
}

func (s *SubstraitErrorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubstraitErrorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubstraitErrorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterSubstraitError(s)
	}
}

func (s *SubstraitErrorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitSubstraitError(s)
	}
}

func (s *SubstraitErrorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitSubstraitError(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) SubstraitError() (localctx ISubstraitErrorContext) {
	localctx = NewSubstraitErrorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, FuncTestCaseParserRULE_substraitError)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(567)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserErrorResult || _la == FuncTestCaseParserUndefineResult) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunc_optionContext is an interface to support dynamic dispatch.
type IFunc_optionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Option_name() IOption_nameContext
	Colon() antlr.TerminalNode
	Option_value() IOption_valueContext

	// IsFunc_optionContext differentiates from other interfaces.
	IsFunc_optionContext()
}

type Func_optionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_optionContext() *Func_optionContext {
	var p = new(Func_optionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_func_option
	return p
}

func InitEmptyFunc_optionContext(p *Func_optionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_func_option
}

func (*Func_optionContext) IsFunc_optionContext() {}

func NewFunc_optionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_optionContext {
	var p = new(Func_optionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_func_option

	return p
}

func (s *Func_optionContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_optionContext) Option_name() IOption_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOption_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOption_nameContext)
}

func (s *Func_optionContext) Colon() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserColon, 0)
}

func (s *Func_optionContext) Option_value() IOption_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOption_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOption_valueContext)
}

func (s *Func_optionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_optionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Func_optionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterFunc_option(s)
	}
}

func (s *Func_optionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitFunc_option(s)
	}
}

func (s *Func_optionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitFunc_option(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) Func_option() (localctx IFunc_optionContext) {
	localctx = NewFunc_optionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, FuncTestCaseParserRULE_func_option)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(569)
		p.Option_name()
	}
	{
		p.SetState(570)
		p.Match(FuncTestCaseParserColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(571)
		p.Option_value()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOption_nameContext is an interface to support dynamic dispatch.
type IOption_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Overflow() antlr.TerminalNode
	Rounding() antlr.TerminalNode
	NullHandling() antlr.TerminalNode
	SpacesOnly() antlr.TerminalNode
	Identifier() antlr.TerminalNode

	// IsOption_nameContext differentiates from other interfaces.
	IsOption_nameContext()
}

type Option_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOption_nameContext() *Option_nameContext {
	var p = new(Option_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_option_name
	return p
}

func InitEmptyOption_nameContext(p *Option_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_option_name
}

func (*Option_nameContext) IsOption_nameContext() {}

func NewOption_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Option_nameContext {
	var p = new(Option_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_option_name

	return p
}

func (s *Option_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Option_nameContext) Overflow() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserOverflow, 0)
}

func (s *Option_nameContext) Rounding() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserRounding, 0)
}

func (s *Option_nameContext) NullHandling() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserNullHandling, 0)
}

func (s *Option_nameContext) SpacesOnly() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserSpacesOnly, 0)
}

func (s *Option_nameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserIdentifier, 0)
}

func (s *Option_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Option_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Option_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterOption_name(s)
	}
}

func (s *Option_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitOption_name(s)
	}
}

func (s *Option_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitOption_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) Option_name() (localctx IOption_nameContext) {
	localctx = NewOption_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, FuncTestCaseParserRULE_option_name)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(573)
		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3151872) != 0) || _la == FuncTestCaseParserIdentifier) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOption_valueContext is an interface to support dynamic dispatch.
type IOption_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Error() antlr.TerminalNode
	Saturate() antlr.TerminalNode
	Silent() antlr.TerminalNode
	TieToEven() antlr.TerminalNode
	NaN() antlr.TerminalNode
	Truncate() antlr.TerminalNode
	AcceptNulls() antlr.TerminalNode
	IgnoreNulls() antlr.TerminalNode
	BooleanLiteral() antlr.TerminalNode
	NullLiteral() antlr.TerminalNode
	Identifier() antlr.TerminalNode

	// IsOption_valueContext differentiates from other interfaces.
	IsOption_valueContext()
}

type Option_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOption_valueContext() *Option_valueContext {
	var p = new(Option_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_option_value
	return p
}

func InitEmptyOption_valueContext(p *Option_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_option_value
}

func (*Option_valueContext) IsOption_valueContext() {}

func NewOption_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Option_valueContext {
	var p = new(Option_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_option_value

	return p
}

func (s *Option_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Option_valueContext) Error() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserError, 0)
}

func (s *Option_valueContext) Saturate() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserSaturate, 0)
}

func (s *Option_valueContext) Silent() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserSilent, 0)
}

func (s *Option_valueContext) TieToEven() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserTieToEven, 0)
}

func (s *Option_valueContext) NaN() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserNaN, 0)
}

func (s *Option_valueContext) Truncate() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserTruncate, 0)
}

func (s *Option_valueContext) AcceptNulls() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserAcceptNulls, 0)
}

func (s *Option_valueContext) IgnoreNulls() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserIgnoreNulls, 0)
}

func (s *Option_valueContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserBooleanLiteral, 0)
}

func (s *Option_valueContext) NullLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserNullLiteral, 0)
}

func (s *Option_valueContext) Identifier() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserIdentifier, 0)
}

func (s *Option_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Option_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Option_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterOption_value(s)
	}
}

func (s *Option_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitOption_value(s)
	}
}

func (s *Option_valueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitOption_value(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) Option_value() (localctx IOption_valueContext) {
	localctx = NewOption_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, FuncTestCaseParserRULE_option_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(575)
		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8796165365760) != 0) || _la == FuncTestCaseParserIdentifier) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunc_optionsContext is an interface to support dynamic dispatch.
type IFunc_optionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFunc_option() []IFunc_optionContext
	Func_option(i int) IFunc_optionContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsFunc_optionsContext differentiates from other interfaces.
	IsFunc_optionsContext()
}

type Func_optionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_optionsContext() *Func_optionsContext {
	var p = new(Func_optionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_func_options
	return p
}

func InitEmptyFunc_optionsContext(p *Func_optionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_func_options
}

func (*Func_optionsContext) IsFunc_optionsContext() {}

func NewFunc_optionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_optionsContext {
	var p = new(Func_optionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_func_options

	return p
}

func (s *Func_optionsContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_optionsContext) AllFunc_option() []IFunc_optionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunc_optionContext); ok {
			len++
		}
	}

	tst := make([]IFunc_optionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunc_optionContext); ok {
			tst[i] = t.(IFunc_optionContext)
			i++
		}
	}

	return tst
}

func (s *Func_optionsContext) Func_option(i int) IFunc_optionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_optionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_optionContext)
}

func (s *Func_optionsContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(FuncTestCaseParserComma)
}

func (s *Func_optionsContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserComma, i)
}

func (s *Func_optionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_optionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Func_optionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterFunc_options(s)
	}
}

func (s *Func_optionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitFunc_options(s)
	}
}

func (s *Func_optionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitFunc_options(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) Func_options() (localctx IFunc_optionsContext) {
	localctx = NewFunc_optionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, FuncTestCaseParserRULE_func_options)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(577)
		p.Func_option()
	}
	p.SetState(582)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FuncTestCaseParserComma {
		{
			p.SetState(578)
			p.Match(FuncTestCaseParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(579)
			p.Func_option()
		}

		p.SetState(584)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INonReservedContext is an interface to support dynamic dispatch.
type INonReservedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	And() antlr.TerminalNode
	Or() antlr.TerminalNode
	Truncate() antlr.TerminalNode

	// IsNonReservedContext differentiates from other interfaces.
	IsNonReservedContext()
}

type NonReservedContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNonReservedContext() *NonReservedContext {
	var p = new(NonReservedContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_nonReserved
	return p
}

func InitEmptyNonReservedContext(p *NonReservedContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_nonReserved
}

func (*NonReservedContext) IsNonReservedContext() {}

func NewNonReservedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NonReservedContext {
	var p = new(NonReservedContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_nonReserved

	return p
}

func (s *NonReservedContext) GetParser() antlr.Parser { return s.parser }

func (s *NonReservedContext) And() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserAnd, 0)
}

func (s *NonReservedContext) Or() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserOr, 0)
}

func (s *NonReservedContext) Truncate() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserTruncate, 0)
}

func (s *NonReservedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NonReservedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NonReservedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterNonReserved(s)
	}
}

func (s *NonReservedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitNonReserved(s)
	}
}

func (s *NonReservedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitNonReserved(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) NonReserved() (localctx INonReservedContext) {
	localctx = NewNonReservedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, FuncTestCaseParserRULE_nonReserved)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(585)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserTruncate || _la == FuncTestCaseParserAnd || _la == FuncTestCaseParserOr) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NonReserved() INonReservedContext
	Identifier() antlr.TerminalNode

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) NonReserved() INonReservedContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INonReservedContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INonReservedContext)
}

func (s *IdentifierContext) Identifier() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserIdentifier, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncTestCaseParserListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, FuncTestCaseParserRULE_identifier)
	p.SetState(589)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FuncTestCaseParserTruncate, FuncTestCaseParserAnd, FuncTestCaseParserOr:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(587)
			p.NonReserved()
		}

	case FuncTestCaseParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(588)
			p.Match(FuncTestCaseParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
