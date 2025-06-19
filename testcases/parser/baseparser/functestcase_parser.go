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
		"'PRECISION_TIME'", "'PRECISION_TIMESTAMP'", "'PRECISION_TIMESTAMP_TZ'",
		"'FIXEDCHAR'", "'VARCHAR'", "'FIXEDBINARY'", "'STRUCT'", "'NSTRUCT'",
		"'LIST'", "'MAP'", "'U!'", "'BOOL'", "'STR'", "'VBIN'", "'TS'", "'TSTZ'",
		"'IYEAR'", "'IDAY'", "'DEC'", "'PT'", "'PTS'", "'PTSTZ'", "'FCHAR'",
		"'VCHAR'", "'FBIN'", "'ANY'", "", "'::'", "'+'", "'-'", "'*'", "'/'",
		"'%'", "'='", "'!='", "'>='", "'<='", "'>'", "'<'", "'!'", "'('", "')'",
		"'['", "']'", "','", "':'", "'?'", "'#'", "'.'", "'AND'", "'OR'", "':='",
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
		"Interval_Day", "UUID", "Decimal", "Precision_Time", "Precision_Timestamp",
		"Precision_Timestamp_TZ", "FixedChar", "VarChar", "FixedBinary", "Struct",
		"NStruct", "List", "Map", "UserDefined", "Bool", "Str", "VBin", "Ts",
		"TsTZ", "IYear", "IDay", "Dec", "PT", "PTs", "PTsTZ", "FChar", "VChar",
		"FBin", "Any", "AnyVar", "DoubleColon", "Plus", "Minus", "Asterisk",
		"ForwardSlash", "Percent", "Eq", "Ne", "Gte", "Lte", "Gt", "Lt", "Bang",
		"OParen", "CParen", "OBracket", "CBracket", "Comma", "Colon", "QMark",
		"Hash", "Dot", "And", "Or", "Assign", "Number", "Identifier", "Newline",
	}
	staticData.RuleNames = []string{
		"doc", "header", "version", "include", "testGroupDescription", "testCase",
		"testGroup", "arguments", "result", "argument", "aggFuncTestCase", "aggFuncCall",
		"tableData", "tableRows", "dataColumn", "columnValues", "literal", "qualifiedAggregateFuncArgs",
		"aggregateFuncArgs", "qualifiedAggregateFuncArg", "aggregateFuncArg",
		"numericLiteral", "floatLiteral", "nullArg", "intArg", "floatArg", "decimalArg",
		"booleanArg", "stringArg", "dateArg", "timeArg", "timestampArg", "timestampTzArg",
		"intervalYearArg", "intervalDayArg", "fixedCharArg", "varCharArg", "fixedBinaryArg",
		"precisionTimeArg", "precisionTimestampArg", "precisionTimestampTZArg",
		"listArg", "literalList", "dataType", "scalarType", "booleanType", "stringType",
		"binaryType", "timestampType", "timestampTZType", "intervalYearType",
		"intervalDayType", "fixedCharType", "varCharType", "fixedBinaryType",
		"decimalType", "precisionTimeType", "precisionTimestampType", "precisionTimestampTZType",
		"listType", "parameterizedType", "numericParameter", "substraitError",
		"funcOption", "optionName", "optionValue", "funcOptions", "nonReserved",
		"identifier",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 122, 634, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
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
		63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67, 2, 68,
		7, 68, 1, 0, 1, 0, 4, 0, 141, 8, 0, 11, 0, 12, 0, 142, 1, 0, 1, 0, 1, 1,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 5, 3, 161, 8, 3, 10, 3, 12, 3, 164, 9, 3, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 176, 8, 5, 1, 5, 1, 5, 1, 5,
		1, 6, 3, 6, 182, 8, 6, 1, 6, 4, 6, 185, 8, 6, 11, 6, 12, 6, 186, 1, 6,
		3, 6, 190, 8, 6, 1, 6, 4, 6, 193, 8, 6, 11, 6, 12, 6, 194, 3, 6, 197, 8,
		6, 1, 7, 1, 7, 1, 7, 5, 7, 202, 8, 7, 10, 7, 12, 7, 205, 9, 7, 1, 8, 1,
		8, 3, 8, 209, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 230,
		8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 237, 8, 10, 1, 10, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 246, 8, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 3, 11, 254, 8, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 3, 11, 263, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 5, 12, 271, 8, 12, 10, 12, 12, 12, 274, 9, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 284, 8, 13, 10, 13, 12,
		13, 287, 9, 13, 3, 13, 289, 8, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 301, 8, 15, 10, 15, 12, 15, 304,
		9, 15, 3, 15, 306, 8, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 320, 8, 16, 1, 17, 1, 17,
		1, 17, 5, 17, 325, 8, 17, 10, 17, 12, 17, 328, 9, 17, 1, 18, 1, 18, 1,
		18, 5, 18, 333, 8, 18, 10, 18, 12, 18, 336, 9, 18, 1, 19, 1, 19, 1, 19,
		1, 19, 3, 19, 342, 8, 19, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 348, 8, 20,
		1, 21, 1, 21, 1, 21, 3, 21, 353, 8, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 365, 8, 24, 1, 25, 1, 25,
		1, 25, 1, 25, 3, 25, 371, 8, 25, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 377,
		8, 26, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 383, 8, 27, 1, 28, 1, 28, 1,
		28, 1, 28, 3, 28, 389, 8, 28, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29, 395, 8,
		29, 1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 401, 8, 30, 1, 31, 1, 31, 1, 31,
		1, 31, 3, 31, 407, 8, 31, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 413, 8, 32,
		1, 33, 1, 33, 1, 33, 1, 33, 3, 33, 419, 8, 33, 1, 34, 1, 34, 1, 34, 1,
		34, 3, 34, 425, 8, 34, 1, 35, 1, 35, 1, 35, 1, 35, 3, 35, 431, 8, 35, 1,
		36, 1, 36, 1, 36, 1, 36, 3, 36, 437, 8, 36, 1, 37, 1, 37, 1, 37, 1, 37,
		3, 37, 443, 8, 37, 1, 38, 1, 38, 1, 38, 1, 38, 3, 38, 449, 8, 38, 1, 39,
		1, 39, 1, 39, 1, 39, 3, 39, 455, 8, 39, 1, 40, 1, 40, 1, 40, 1, 40, 3,
		40, 461, 8, 40, 1, 41, 1, 41, 1, 41, 1, 41, 3, 41, 467, 8, 41, 1, 42, 1,
		42, 1, 42, 1, 42, 5, 42, 473, 8, 42, 10, 42, 12, 42, 476, 9, 42, 3, 42,
		478, 8, 42, 1, 42, 1, 42, 1, 43, 1, 43, 3, 43, 484, 8, 43, 1, 43, 3, 43,
		487, 8, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1,
		44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 3, 44, 506,
		8, 44, 1, 45, 1, 45, 1, 46, 1, 46, 1, 47, 1, 47, 1, 48, 1, 48, 1, 49, 1,
		49, 1, 50, 1, 50, 1, 51, 1, 51, 3, 51, 522, 8, 51, 1, 51, 1, 51, 1, 51,
		1, 51, 3, 51, 528, 8, 51, 1, 52, 1, 52, 3, 52, 532, 8, 52, 1, 52, 1, 52,
		1, 52, 1, 52, 1, 53, 1, 53, 3, 53, 540, 8, 53, 1, 53, 1, 53, 1, 53, 1,
		53, 1, 54, 1, 54, 3, 54, 548, 8, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 55,
		1, 55, 3, 55, 556, 8, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 3,
		55, 564, 8, 55, 1, 56, 1, 56, 3, 56, 568, 8, 56, 1, 56, 1, 56, 1, 56, 1,
		56, 1, 57, 1, 57, 3, 57, 576, 8, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 58,
		1, 58, 3, 58, 584, 8, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 59, 1, 59, 3,
		59, 592, 8, 59, 1, 59, 1, 59, 1, 59, 1, 59, 1, 60, 1, 60, 1, 60, 1, 60,
		1, 60, 1, 60, 1, 60, 1, 60, 3, 60, 606, 8, 60, 1, 61, 1, 61, 1, 62, 1,
		62, 1, 63, 1, 63, 1, 63, 1, 63, 1, 64, 1, 64, 1, 65, 1, 65, 1, 66, 1, 66,
		1, 66, 5, 66, 623, 8, 66, 10, 66, 12, 66, 626, 9, 66, 1, 67, 1, 67, 1,
		68, 1, 68, 3, 68, 632, 8, 68, 1, 68, 0, 0, 69, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50,
		52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86,
		88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112, 114, 116, 118,
		120, 122, 124, 126, 128, 130, 132, 134, 136, 0, 22, 1, 0, 3, 4, 2, 0, 17,
		17, 25, 25, 1, 0, 52, 55, 1, 0, 56, 57, 2, 0, 51, 51, 79, 79, 2, 0, 58,
		58, 80, 80, 2, 0, 59, 59, 81, 81, 2, 0, 60, 60, 82, 82, 2, 0, 61, 61, 83,
		83, 2, 0, 64, 64, 84, 84, 2, 0, 65, 65, 85, 85, 2, 0, 71, 71, 90, 90, 2,
		0, 72, 72, 91, 91, 2, 0, 73, 73, 92, 92, 2, 0, 67, 67, 86, 86, 2, 0, 68,
		68, 87, 87, 2, 0, 69, 69, 88, 88, 2, 0, 70, 70, 89, 89, 1, 0, 9, 10, 3,
		0, 11, 12, 20, 21, 121, 121, 5, 0, 13, 19, 22, 22, 26, 26, 43, 43, 121,
		121, 2, 0, 22, 22, 117, 118, 674, 0, 138, 1, 0, 0, 0, 2, 146, 1, 0, 0,
		0, 4, 149, 1, 0, 0, 0, 6, 154, 1, 0, 0, 0, 8, 165, 1, 0, 0, 0, 10, 167,
		1, 0, 0, 0, 12, 196, 1, 0, 0, 0, 14, 198, 1, 0, 0, 0, 16, 208, 1, 0, 0,
		0, 18, 229, 1, 0, 0, 0, 20, 231, 1, 0, 0, 0, 22, 262, 1, 0, 0, 0, 24, 264,
		1, 0, 0, 0, 26, 279, 1, 0, 0, 0, 28, 292, 1, 0, 0, 0, 30, 296, 1, 0, 0,
		0, 32, 319, 1, 0, 0, 0, 34, 321, 1, 0, 0, 0, 36, 329, 1, 0, 0, 0, 38, 341,
		1, 0, 0, 0, 40, 347, 1, 0, 0, 0, 42, 352, 1, 0, 0, 0, 44, 354, 1, 0, 0,
		0, 46, 356, 1, 0, 0, 0, 48, 360, 1, 0, 0, 0, 50, 366, 1, 0, 0, 0, 52, 372,
		1, 0, 0, 0, 54, 378, 1, 0, 0, 0, 56, 384, 1, 0, 0, 0, 58, 390, 1, 0, 0,
		0, 60, 396, 1, 0, 0, 0, 62, 402, 1, 0, 0, 0, 64, 408, 1, 0, 0, 0, 66, 414,
		1, 0, 0, 0, 68, 420, 1, 0, 0, 0, 70, 426, 1, 0, 0, 0, 72, 432, 1, 0, 0,
		0, 74, 438, 1, 0, 0, 0, 76, 444, 1, 0, 0, 0, 78, 450, 1, 0, 0, 0, 80, 456,
		1, 0, 0, 0, 82, 462, 1, 0, 0, 0, 84, 468, 1, 0, 0, 0, 86, 486, 1, 0, 0,
		0, 88, 505, 1, 0, 0, 0, 90, 507, 1, 0, 0, 0, 92, 509, 1, 0, 0, 0, 94, 511,
		1, 0, 0, 0, 96, 513, 1, 0, 0, 0, 98, 515, 1, 0, 0, 0, 100, 517, 1, 0, 0,
		0, 102, 519, 1, 0, 0, 0, 104, 529, 1, 0, 0, 0, 106, 537, 1, 0, 0, 0, 108,
		545, 1, 0, 0, 0, 110, 553, 1, 0, 0, 0, 112, 565, 1, 0, 0, 0, 114, 573,
		1, 0, 0, 0, 116, 581, 1, 0, 0, 0, 118, 589, 1, 0, 0, 0, 120, 605, 1, 0,
		0, 0, 122, 607, 1, 0, 0, 0, 124, 609, 1, 0, 0, 0, 126, 611, 1, 0, 0, 0,
		128, 615, 1, 0, 0, 0, 130, 617, 1, 0, 0, 0, 132, 619, 1, 0, 0, 0, 134,
		627, 1, 0, 0, 0, 136, 631, 1, 0, 0, 0, 138, 140, 3, 2, 1, 0, 139, 141,
		3, 12, 6, 0, 140, 139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 140, 1, 0,
		0, 0, 142, 143, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 145, 5, 0, 0, 1,
		145, 1, 1, 0, 0, 0, 146, 147, 3, 4, 2, 0, 147, 148, 3, 6, 3, 0, 148, 3,
		1, 0, 0, 0, 149, 150, 5, 2, 0, 0, 150, 151, 7, 0, 0, 0, 151, 152, 5, 113,
		0, 0, 152, 153, 5, 6, 0, 0, 153, 5, 1, 0, 0, 0, 154, 155, 5, 2, 0, 0, 155,
		156, 5, 5, 0, 0, 156, 157, 5, 113, 0, 0, 157, 162, 5, 44, 0, 0, 158, 159,
		5, 112, 0, 0, 159, 161, 5, 44, 0, 0, 160, 158, 1, 0, 0, 0, 161, 164, 1,
		0, 0, 0, 162, 160, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 7, 1, 0, 0, 0,
		164, 162, 1, 0, 0, 0, 165, 166, 5, 7, 0, 0, 166, 9, 1, 0, 0, 0, 167, 168,
		3, 136, 68, 0, 168, 169, 5, 108, 0, 0, 169, 170, 3, 14, 7, 0, 170, 175,
		5, 109, 0, 0, 171, 172, 5, 110, 0, 0, 172, 173, 3, 132, 66, 0, 173, 174,
		5, 111, 0, 0, 174, 176, 1, 0, 0, 0, 175, 171, 1, 0, 0, 0, 175, 176, 1,
		0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 178, 5, 101, 0, 0, 178, 179, 3, 16,
		8, 0, 179, 11, 1, 0, 0, 0, 180, 182, 3, 8, 4, 0, 181, 180, 1, 0, 0, 0,
		181, 182, 1, 0, 0, 0, 182, 184, 1, 0, 0, 0, 183, 185, 3, 10, 5, 0, 184,
		183, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 186, 187,
		1, 0, 0, 0, 187, 197, 1, 0, 0, 0, 188, 190, 3, 8, 4, 0, 189, 188, 1, 0,
		0, 0, 189, 190, 1, 0, 0, 0, 190, 192, 1, 0, 0, 0, 191, 193, 3, 20, 10,
		0, 192, 191, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 192, 1, 0, 0, 0, 194,
		195, 1, 0, 0, 0, 195, 197, 1, 0, 0, 0, 196, 181, 1, 0, 0, 0, 196, 189,
		1, 0, 0, 0, 197, 13, 1, 0, 0, 0, 198, 203, 3, 18, 9, 0, 199, 200, 5, 112,
		0, 0, 200, 202, 3, 18, 9, 0, 201, 199, 1, 0, 0, 0, 202, 205, 1, 0, 0, 0,
		203, 201, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204, 15, 1, 0, 0, 0, 205, 203,
		1, 0, 0, 0, 206, 209, 3, 18, 9, 0, 207, 209, 3, 124, 62, 0, 208, 206, 1,
		0, 0, 0, 208, 207, 1, 0, 0, 0, 209, 17, 1, 0, 0, 0, 210, 230, 3, 46, 23,
		0, 211, 230, 3, 48, 24, 0, 212, 230, 3, 50, 25, 0, 213, 230, 3, 54, 27,
		0, 214, 230, 3, 56, 28, 0, 215, 230, 3, 52, 26, 0, 216, 230, 3, 58, 29,
		0, 217, 230, 3, 60, 30, 0, 218, 230, 3, 62, 31, 0, 219, 230, 3, 64, 32,
		0, 220, 230, 3, 66, 33, 0, 221, 230, 3, 68, 34, 0, 222, 230, 3, 70, 35,
		0, 223, 230, 3, 72, 36, 0, 224, 230, 3, 74, 37, 0, 225, 230, 3, 76, 38,
		0, 226, 230, 3, 78, 39, 0, 227, 230, 3, 80, 40, 0, 228, 230, 3, 82, 41,
		0, 229, 210, 1, 0, 0, 0, 229, 211, 1, 0, 0, 0, 229, 212, 1, 0, 0, 0, 229,
		213, 1, 0, 0, 0, 229, 214, 1, 0, 0, 0, 229, 215, 1, 0, 0, 0, 229, 216,
		1, 0, 0, 0, 229, 217, 1, 0, 0, 0, 229, 218, 1, 0, 0, 0, 229, 219, 1, 0,
		0, 0, 229, 220, 1, 0, 0, 0, 229, 221, 1, 0, 0, 0, 229, 222, 1, 0, 0, 0,
		229, 223, 1, 0, 0, 0, 229, 224, 1, 0, 0, 0, 229, 225, 1, 0, 0, 0, 229,
		226, 1, 0, 0, 0, 229, 227, 1, 0, 0, 0, 229, 228, 1, 0, 0, 0, 230, 19, 1,
		0, 0, 0, 231, 236, 3, 22, 11, 0, 232, 233, 5, 110, 0, 0, 233, 234, 3, 132,
		66, 0, 234, 235, 5, 111, 0, 0, 235, 237, 1, 0, 0, 0, 236, 232, 1, 0, 0,
		0, 236, 237, 1, 0, 0, 0, 237, 238, 1, 0, 0, 0, 238, 239, 5, 101, 0, 0,
		239, 240, 3, 16, 8, 0, 240, 21, 1, 0, 0, 0, 241, 242, 3, 24, 12, 0, 242,
		243, 3, 136, 68, 0, 243, 245, 5, 108, 0, 0, 244, 246, 3, 34, 17, 0, 245,
		244, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 247, 1, 0, 0, 0, 247, 248,
		5, 109, 0, 0, 248, 263, 1, 0, 0, 0, 249, 250, 3, 26, 13, 0, 250, 251, 3,
		136, 68, 0, 251, 253, 5, 108, 0, 0, 252, 254, 3, 36, 18, 0, 253, 252, 1,
		0, 0, 0, 253, 254, 1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 256, 5, 109,
		0, 0, 256, 263, 1, 0, 0, 0, 257, 258, 3, 136, 68, 0, 258, 259, 5, 108,
		0, 0, 259, 260, 3, 28, 14, 0, 260, 261, 5, 109, 0, 0, 261, 263, 1, 0, 0,
		0, 262, 241, 1, 0, 0, 0, 262, 249, 1, 0, 0, 0, 262, 257, 1, 0, 0, 0, 263,
		23, 1, 0, 0, 0, 264, 265, 5, 8, 0, 0, 265, 266, 5, 121, 0, 0, 266, 267,
		5, 108, 0, 0, 267, 272, 3, 86, 43, 0, 268, 269, 5, 112, 0, 0, 269, 271,
		3, 86, 43, 0, 270, 268, 1, 0, 0, 0, 271, 274, 1, 0, 0, 0, 272, 270, 1,
		0, 0, 0, 272, 273, 1, 0, 0, 0, 273, 275, 1, 0, 0, 0, 274, 272, 1, 0, 0,
		0, 275, 276, 5, 109, 0, 0, 276, 277, 5, 101, 0, 0, 277, 278, 3, 26, 13,
		0, 278, 25, 1, 0, 0, 0, 279, 288, 5, 108, 0, 0, 280, 285, 3, 30, 15, 0,
		281, 282, 5, 112, 0, 0, 282, 284, 3, 30, 15, 0, 283, 281, 1, 0, 0, 0, 284,
		287, 1, 0, 0, 0, 285, 283, 1, 0, 0, 0, 285, 286, 1, 0, 0, 0, 286, 289,
		1, 0, 0, 0, 287, 285, 1, 0, 0, 0, 288, 280, 1, 0, 0, 0, 288, 289, 1, 0,
		0, 0, 289, 290, 1, 0, 0, 0, 290, 291, 5, 109, 0, 0, 291, 27, 1, 0, 0, 0,
		292, 293, 3, 30, 15, 0, 293, 294, 5, 95, 0, 0, 294, 295, 3, 86, 43, 0,
		295, 29, 1, 0, 0, 0, 296, 305, 5, 108, 0, 0, 297, 302, 3, 32, 16, 0, 298,
		299, 5, 112, 0, 0, 299, 301, 3, 32, 16, 0, 300, 298, 1, 0, 0, 0, 301, 304,
		1, 0, 0, 0, 302, 300, 1, 0, 0, 0, 302, 303, 1, 0, 0, 0, 303, 306, 1, 0,
		0, 0, 304, 302, 1, 0, 0, 0, 305, 297, 1, 0, 0, 0, 305, 306, 1, 0, 0, 0,
		306, 307, 1, 0, 0, 0, 307, 308, 5, 109, 0, 0, 308, 31, 1, 0, 0, 0, 309,
		320, 5, 43, 0, 0, 310, 320, 3, 42, 21, 0, 311, 320, 5, 26, 0, 0, 312, 320,
		5, 44, 0, 0, 313, 320, 5, 30, 0, 0, 314, 320, 5, 29, 0, 0, 315, 320, 5,
		28, 0, 0, 316, 320, 5, 27, 0, 0, 317, 320, 5, 41, 0, 0, 318, 320, 5, 42,
		0, 0, 319, 309, 1, 0, 0, 0, 319, 310, 1, 0, 0, 0, 319, 311, 1, 0, 0, 0,
		319, 312, 1, 0, 0, 0, 319, 313, 1, 0, 0, 0, 319, 314, 1, 0, 0, 0, 319,
		315, 1, 0, 0, 0, 319, 316, 1, 0, 0, 0, 319, 317, 1, 0, 0, 0, 319, 318,
		1, 0, 0, 0, 320, 33, 1, 0, 0, 0, 321, 326, 3, 38, 19, 0, 322, 323, 5, 112,
		0, 0, 323, 325, 3, 38, 19, 0, 324, 322, 1, 0, 0, 0, 325, 328, 1, 0, 0,
		0, 326, 324, 1, 0, 0, 0, 326, 327, 1, 0, 0, 0, 327, 35, 1, 0, 0, 0, 328,
		326, 1, 0, 0, 0, 329, 334, 3, 40, 20, 0, 330, 331, 5, 112, 0, 0, 331, 333,
		3, 40, 20, 0, 332, 330, 1, 0, 0, 0, 333, 336, 1, 0, 0, 0, 334, 332, 1,
		0, 0, 0, 334, 335, 1, 0, 0, 0, 335, 37, 1, 0, 0, 0, 336, 334, 1, 0, 0,
		0, 337, 338, 5, 121, 0, 0, 338, 339, 5, 116, 0, 0, 339, 342, 5, 45, 0,
		0, 340, 342, 3, 18, 9, 0, 341, 337, 1, 0, 0, 0, 341, 340, 1, 0, 0, 0, 342,
		39, 1, 0, 0, 0, 343, 344, 5, 45, 0, 0, 344, 345, 5, 95, 0, 0, 345, 348,
		3, 86, 43, 0, 346, 348, 3, 18, 9, 0, 347, 343, 1, 0, 0, 0, 347, 346, 1,
		0, 0, 0, 348, 41, 1, 0, 0, 0, 349, 353, 5, 24, 0, 0, 350, 353, 5, 23, 0,
		0, 351, 353, 3, 44, 22, 0, 352, 349, 1, 0, 0, 0, 352, 350, 1, 0, 0, 0,
		352, 351, 1, 0, 0, 0, 353, 43, 1, 0, 0, 0, 354, 355, 7, 1, 0, 0, 355, 45,
		1, 0, 0, 0, 356, 357, 5, 43, 0, 0, 357, 358, 5, 95, 0, 0, 358, 359, 3,
		86, 43, 0, 359, 47, 1, 0, 0, 0, 360, 361, 5, 23, 0, 0, 361, 362, 5, 95,
		0, 0, 362, 364, 7, 2, 0, 0, 363, 365, 5, 114, 0, 0, 364, 363, 1, 0, 0,
		0, 364, 365, 1, 0, 0, 0, 365, 49, 1, 0, 0, 0, 366, 367, 3, 42, 21, 0, 367,
		368, 5, 95, 0, 0, 368, 370, 7, 3, 0, 0, 369, 371, 5, 114, 0, 0, 370, 369,
		1, 0, 0, 0, 370, 371, 1, 0, 0, 0, 371, 51, 1, 0, 0, 0, 372, 373, 3, 42,
		21, 0, 373, 374, 5, 95, 0, 0, 374, 376, 3, 110, 55, 0, 375, 377, 5, 114,
		0, 0, 376, 375, 1, 0, 0, 0, 376, 377, 1, 0, 0, 0, 377, 53, 1, 0, 0, 0,
		378, 379, 5, 26, 0, 0, 379, 380, 5, 95, 0, 0, 380, 382, 3, 90, 45, 0, 381,
		383, 5, 114, 0, 0, 382, 381, 1, 0, 0, 0, 382, 383, 1, 0, 0, 0, 383, 55,
		1, 0, 0, 0, 384, 385, 5, 44, 0, 0, 385, 386, 5, 95, 0, 0, 386, 388, 3,
		92, 46, 0, 387, 389, 5, 114, 0, 0, 388, 387, 1, 0, 0, 0, 388, 389, 1, 0,
		0, 0, 389, 57, 1, 0, 0, 0, 390, 391, 5, 30, 0, 0, 391, 392, 5, 95, 0, 0,
		392, 394, 5, 62, 0, 0, 393, 395, 5, 114, 0, 0, 394, 393, 1, 0, 0, 0, 394,
		395, 1, 0, 0, 0, 395, 59, 1, 0, 0, 0, 396, 397, 5, 29, 0, 0, 397, 398,
		5, 95, 0, 0, 398, 400, 5, 63, 0, 0, 399, 401, 5, 114, 0, 0, 400, 399, 1,
		0, 0, 0, 400, 401, 1, 0, 0, 0, 401, 61, 1, 0, 0, 0, 402, 403, 5, 28, 0,
		0, 403, 404, 5, 95, 0, 0, 404, 406, 3, 96, 48, 0, 405, 407, 5, 114, 0,
		0, 406, 405, 1, 0, 0, 0, 406, 407, 1, 0, 0, 0, 407, 63, 1, 0, 0, 0, 408,
		409, 5, 27, 0, 0, 409, 410, 5, 95, 0, 0, 410, 412, 3, 98, 49, 0, 411, 413,
		5, 114, 0, 0, 412, 411, 1, 0, 0, 0, 412, 413, 1, 0, 0, 0, 413, 65, 1, 0,
		0, 0, 414, 415, 5, 41, 0, 0, 415, 416, 5, 95, 0, 0, 416, 418, 3, 100, 50,
		0, 417, 419, 5, 114, 0, 0, 418, 417, 1, 0, 0, 0, 418, 419, 1, 0, 0, 0,
		419, 67, 1, 0, 0, 0, 420, 421, 5, 42, 0, 0, 421, 422, 5, 95, 0, 0, 422,
		424, 3, 102, 51, 0, 423, 425, 5, 114, 0, 0, 424, 423, 1, 0, 0, 0, 424,
		425, 1, 0, 0, 0, 425, 69, 1, 0, 0, 0, 426, 427, 5, 44, 0, 0, 427, 428,
		5, 95, 0, 0, 428, 430, 3, 104, 52, 0, 429, 431, 5, 114, 0, 0, 430, 429,
		1, 0, 0, 0, 430, 431, 1, 0, 0, 0, 431, 71, 1, 0, 0, 0, 432, 433, 5, 44,
		0, 0, 433, 434, 5, 95, 0, 0, 434, 436, 3, 106, 53, 0, 435, 437, 5, 114,
		0, 0, 436, 435, 1, 0, 0, 0, 436, 437, 1, 0, 0, 0, 437, 73, 1, 0, 0, 0,
		438, 439, 5, 44, 0, 0, 439, 440, 5, 95, 0, 0, 440, 442, 3, 108, 54, 0,
		441, 443, 5, 114, 0, 0, 442, 441, 1, 0, 0, 0, 442, 443, 1, 0, 0, 0, 443,
		75, 1, 0, 0, 0, 444, 445, 5, 29, 0, 0, 445, 446, 5, 95, 0, 0, 446, 448,
		3, 112, 56, 0, 447, 449, 5, 114, 0, 0, 448, 447, 1, 0, 0, 0, 448, 449,
		1, 0, 0, 0, 449, 77, 1, 0, 0, 0, 450, 451, 5, 28, 0, 0, 451, 452, 5, 95,
		0, 0, 452, 454, 3, 114, 57, 0, 453, 455, 5, 114, 0, 0, 454, 453, 1, 0,
		0, 0, 454, 455, 1, 0, 0, 0, 455, 79, 1, 0, 0, 0, 456, 457, 5, 27, 0, 0,
		457, 458, 5, 95, 0, 0, 458, 460, 3, 116, 58, 0, 459, 461, 5, 114, 0, 0,
		460, 459, 1, 0, 0, 0, 460, 461, 1, 0, 0, 0, 461, 81, 1, 0, 0, 0, 462, 463,
		3, 84, 42, 0, 463, 464, 5, 95, 0, 0, 464, 466, 3, 118, 59, 0, 465, 467,
		5, 114, 0, 0, 466, 465, 1, 0, 0, 0, 466, 467, 1, 0, 0, 0, 467, 83, 1, 0,
		0, 0, 468, 477, 5, 110, 0, 0, 469, 474, 3, 32, 16, 0, 470, 471, 5, 112,
		0, 0, 471, 473, 3, 32, 16, 0, 472, 470, 1, 0, 0, 0, 473, 476, 1, 0, 0,
		0, 474, 472, 1, 0, 0, 0, 474, 475, 1, 0, 0, 0, 475, 478, 1, 0, 0, 0, 476,
		474, 1, 0, 0, 0, 477, 469, 1, 0, 0, 0, 477, 478, 1, 0, 0, 0, 478, 479,
		1, 0, 0, 0, 479, 480, 5, 111, 0, 0, 480, 85, 1, 0, 0, 0, 481, 483, 3, 88,
		44, 0, 482, 484, 5, 114, 0, 0, 483, 482, 1, 0, 0, 0, 483, 484, 1, 0, 0,
		0, 484, 487, 1, 0, 0, 0, 485, 487, 3, 120, 60, 0, 486, 481, 1, 0, 0, 0,
		486, 485, 1, 0, 0, 0, 487, 87, 1, 0, 0, 0, 488, 506, 3, 90, 45, 0, 489,
		506, 5, 52, 0, 0, 490, 506, 5, 53, 0, 0, 491, 506, 5, 54, 0, 0, 492, 506,
		5, 55, 0, 0, 493, 506, 5, 56, 0, 0, 494, 506, 5, 57, 0, 0, 495, 506, 3,
		92, 46, 0, 496, 506, 3, 94, 47, 0, 497, 506, 3, 96, 48, 0, 498, 506, 3,
		98, 49, 0, 499, 506, 5, 62, 0, 0, 500, 506, 5, 63, 0, 0, 501, 506, 3, 100,
		50, 0, 502, 506, 5, 66, 0, 0, 503, 504, 5, 78, 0, 0, 504, 506, 5, 121,
		0, 0, 505, 488, 1, 0, 0, 0, 505, 489, 1, 0, 0, 0, 505, 490, 1, 0, 0, 0,
		505, 491, 1, 0, 0, 0, 505, 492, 1, 0, 0, 0, 505, 493, 1, 0, 0, 0, 505,
		494, 1, 0, 0, 0, 505, 495, 1, 0, 0, 0, 505, 496, 1, 0, 0, 0, 505, 497,
		1, 0, 0, 0, 505, 498, 1, 0, 0, 0, 505, 499, 1, 0, 0, 0, 505, 500, 1, 0,
		0, 0, 505, 501, 1, 0, 0, 0, 505, 502, 1, 0, 0, 0, 505, 503, 1, 0, 0, 0,
		506, 89, 1, 0, 0, 0, 507, 508, 7, 4, 0, 0, 508, 91, 1, 0, 0, 0, 509, 510,
		7, 5, 0, 0, 510, 93, 1, 0, 0, 0, 511, 512, 7, 6, 0, 0, 512, 95, 1, 0, 0,
		0, 513, 514, 7, 7, 0, 0, 514, 97, 1, 0, 0, 0, 515, 516, 7, 8, 0, 0, 516,
		99, 1, 0, 0, 0, 517, 518, 7, 9, 0, 0, 518, 101, 1, 0, 0, 0, 519, 521, 7,
		10, 0, 0, 520, 522, 5, 114, 0, 0, 521, 520, 1, 0, 0, 0, 521, 522, 1, 0,
		0, 0, 522, 527, 1, 0, 0, 0, 523, 524, 5, 39, 0, 0, 524, 525, 3, 122, 61,
		0, 525, 526, 5, 40, 0, 0, 526, 528, 1, 0, 0, 0, 527, 523, 1, 0, 0, 0, 527,
		528, 1, 0, 0, 0, 528, 103, 1, 0, 0, 0, 529, 531, 7, 11, 0, 0, 530, 532,
		5, 114, 0, 0, 531, 530, 1, 0, 0, 0, 531, 532, 1, 0, 0, 0, 532, 533, 1,
		0, 0, 0, 533, 534, 5, 39, 0, 0, 534, 535, 3, 122, 61, 0, 535, 536, 5, 40,
		0, 0, 536, 105, 1, 0, 0, 0, 537, 539, 7, 12, 0, 0, 538, 540, 5, 114, 0,
		0, 539, 538, 1, 0, 0, 0, 539, 540, 1, 0, 0, 0, 540, 541, 1, 0, 0, 0, 541,
		542, 5, 39, 0, 0, 542, 543, 3, 122, 61, 0, 543, 544, 5, 40, 0, 0, 544,
		107, 1, 0, 0, 0, 545, 547, 7, 13, 0, 0, 546, 548, 5, 114, 0, 0, 547, 546,
		1, 0, 0, 0, 547, 548, 1, 0, 0, 0, 548, 549, 1, 0, 0, 0, 549, 550, 5, 39,
		0, 0, 550, 551, 3, 122, 61, 0, 551, 552, 5, 40, 0, 0, 552, 109, 1, 0, 0,
		0, 553, 555, 7, 14, 0, 0, 554, 556, 5, 114, 0, 0, 555, 554, 1, 0, 0, 0,
		555, 556, 1, 0, 0, 0, 556, 563, 1, 0, 0, 0, 557, 558, 5, 39, 0, 0, 558,
		559, 3, 122, 61, 0, 559, 560, 5, 112, 0, 0, 560, 561, 3, 122, 61, 0, 561,
		562, 5, 40, 0, 0, 562, 564, 1, 0, 0, 0, 563, 557, 1, 0, 0, 0, 563, 564,
		1, 0, 0, 0, 564, 111, 1, 0, 0, 0, 565, 567, 7, 15, 0, 0, 566, 568, 5, 114,
		0, 0, 567, 566, 1, 0, 0, 0, 567, 568, 1, 0, 0, 0, 568, 569, 1, 0, 0, 0,
		569, 570, 5, 39, 0, 0, 570, 571, 3, 122, 61, 0, 571, 572, 5, 40, 0, 0,
		572, 113, 1, 0, 0, 0, 573, 575, 7, 16, 0, 0, 574, 576, 5, 114, 0, 0, 575,
		574, 1, 0, 0, 0, 575, 576, 1, 0, 0, 0, 576, 577, 1, 0, 0, 0, 577, 578,
		5, 39, 0, 0, 578, 579, 3, 122, 61, 0, 579, 580, 5, 40, 0, 0, 580, 115,
		1, 0, 0, 0, 581, 583, 7, 17, 0, 0, 582, 584, 5, 114, 0, 0, 583, 582, 1,
		0, 0, 0, 583, 584, 1, 0, 0, 0, 584, 585, 1, 0, 0, 0, 585, 586, 5, 39, 0,
		0, 586, 587, 3, 122, 61, 0, 587, 588, 5, 40, 0, 0, 588, 117, 1, 0, 0, 0,
		589, 591, 5, 76, 0, 0, 590, 592, 5, 114, 0, 0, 591, 590, 1, 0, 0, 0, 591,
		592, 1, 0, 0, 0, 592, 593, 1, 0, 0, 0, 593, 594, 5, 39, 0, 0, 594, 595,
		3, 86, 43, 0, 595, 596, 5, 40, 0, 0, 596, 119, 1, 0, 0, 0, 597, 606, 3,
		104, 52, 0, 598, 606, 3, 106, 53, 0, 599, 606, 3, 108, 54, 0, 600, 606,
		3, 110, 55, 0, 601, 606, 3, 102, 51, 0, 602, 606, 3, 112, 56, 0, 603, 606,
		3, 114, 57, 0, 604, 606, 3, 116, 58, 0, 605, 597, 1, 0, 0, 0, 605, 598,
		1, 0, 0, 0, 605, 599, 1, 0, 0, 0, 605, 600, 1, 0, 0, 0, 605, 601, 1, 0,
		0, 0, 605, 602, 1, 0, 0, 0, 605, 603, 1, 0, 0, 0, 605, 604, 1, 0, 0, 0,
		606, 121, 1, 0, 0, 0, 607, 608, 5, 23, 0, 0, 608, 123, 1, 0, 0, 0, 609,
		610, 7, 18, 0, 0, 610, 125, 1, 0, 0, 0, 611, 612, 3, 128, 64, 0, 612, 613,
		5, 113, 0, 0, 613, 614, 3, 130, 65, 0, 614, 127, 1, 0, 0, 0, 615, 616,
		7, 19, 0, 0, 616, 129, 1, 0, 0, 0, 617, 618, 7, 20, 0, 0, 618, 131, 1,
		0, 0, 0, 619, 624, 3, 126, 63, 0, 620, 621, 5, 112, 0, 0, 621, 623, 3,
		126, 63, 0, 622, 620, 1, 0, 0, 0, 623, 626, 1, 0, 0, 0, 624, 622, 1, 0,
		0, 0, 624, 625, 1, 0, 0, 0, 625, 133, 1, 0, 0, 0, 626, 624, 1, 0, 0, 0,
		627, 628, 7, 21, 0, 0, 628, 135, 1, 0, 0, 0, 629, 632, 3, 134, 67, 0, 630,
		632, 5, 121, 0, 0, 631, 629, 1, 0, 0, 0, 631, 630, 1, 0, 0, 0, 632, 137,
		1, 0, 0, 0, 63, 142, 162, 175, 181, 186, 189, 194, 196, 203, 208, 229,
		236, 245, 253, 262, 272, 285, 288, 302, 305, 319, 326, 334, 341, 347, 352,
		364, 370, 376, 382, 388, 394, 400, 406, 412, 418, 424, 430, 436, 442, 448,
		454, 460, 466, 474, 477, 483, 486, 505, 521, 527, 531, 539, 547, 555, 563,
		567, 575, 583, 591, 605, 624, 631,
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
	FuncTestCaseParserPrecision_Time         = 68
	FuncTestCaseParserPrecision_Timestamp    = 69
	FuncTestCaseParserPrecision_Timestamp_TZ = 70
	FuncTestCaseParserFixedChar              = 71
	FuncTestCaseParserVarChar                = 72
	FuncTestCaseParserFixedBinary            = 73
	FuncTestCaseParserStruct                 = 74
	FuncTestCaseParserNStruct                = 75
	FuncTestCaseParserList                   = 76
	FuncTestCaseParserMap                    = 77
	FuncTestCaseParserUserDefined            = 78
	FuncTestCaseParserBool                   = 79
	FuncTestCaseParserStr                    = 80
	FuncTestCaseParserVBin                   = 81
	FuncTestCaseParserTs                     = 82
	FuncTestCaseParserTsTZ                   = 83
	FuncTestCaseParserIYear                  = 84
	FuncTestCaseParserIDay                   = 85
	FuncTestCaseParserDec                    = 86
	FuncTestCaseParserPT                     = 87
	FuncTestCaseParserPTs                    = 88
	FuncTestCaseParserPTsTZ                  = 89
	FuncTestCaseParserFChar                  = 90
	FuncTestCaseParserVChar                  = 91
	FuncTestCaseParserFBin                   = 92
	FuncTestCaseParserAny                    = 93
	FuncTestCaseParserAnyVar                 = 94
	FuncTestCaseParserDoubleColon            = 95
	FuncTestCaseParserPlus                   = 96
	FuncTestCaseParserMinus                  = 97
	FuncTestCaseParserAsterisk               = 98
	FuncTestCaseParserForwardSlash           = 99
	FuncTestCaseParserPercent                = 100
	FuncTestCaseParserEq                     = 101
	FuncTestCaseParserNe                     = 102
	FuncTestCaseParserGte                    = 103
	FuncTestCaseParserLte                    = 104
	FuncTestCaseParserGt                     = 105
	FuncTestCaseParserLt                     = 106
	FuncTestCaseParserBang                   = 107
	FuncTestCaseParserOParen                 = 108
	FuncTestCaseParserCParen                 = 109
	FuncTestCaseParserOBracket               = 110
	FuncTestCaseParserCBracket               = 111
	FuncTestCaseParserComma                  = 112
	FuncTestCaseParserColon                  = 113
	FuncTestCaseParserQMark                  = 114
	FuncTestCaseParserHash                   = 115
	FuncTestCaseParserDot                    = 116
	FuncTestCaseParserAnd                    = 117
	FuncTestCaseParserOr                     = 118
	FuncTestCaseParserAssign                 = 119
	FuncTestCaseParserNumber                 = 120
	FuncTestCaseParserIdentifier             = 121
	FuncTestCaseParserNewline                = 122
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
	FuncTestCaseParserRULE_fixedCharArg               = 35
	FuncTestCaseParserRULE_varCharArg                 = 36
	FuncTestCaseParserRULE_fixedBinaryArg             = 37
	FuncTestCaseParserRULE_precisionTimeArg           = 38
	FuncTestCaseParserRULE_precisionTimestampArg      = 39
	FuncTestCaseParserRULE_precisionTimestampTZArg    = 40
	FuncTestCaseParserRULE_listArg                    = 41
	FuncTestCaseParserRULE_literalList                = 42
	FuncTestCaseParserRULE_dataType                   = 43
	FuncTestCaseParserRULE_scalarType                 = 44
	FuncTestCaseParserRULE_booleanType                = 45
	FuncTestCaseParserRULE_stringType                 = 46
	FuncTestCaseParserRULE_binaryType                 = 47
	FuncTestCaseParserRULE_timestampType              = 48
	FuncTestCaseParserRULE_timestampTZType            = 49
	FuncTestCaseParserRULE_intervalYearType           = 50
	FuncTestCaseParserRULE_intervalDayType            = 51
	FuncTestCaseParserRULE_fixedCharType              = 52
	FuncTestCaseParserRULE_varCharType                = 53
	FuncTestCaseParserRULE_fixedBinaryType            = 54
	FuncTestCaseParserRULE_decimalType                = 55
	FuncTestCaseParserRULE_precisionTimeType          = 56
	FuncTestCaseParserRULE_precisionTimestampType     = 57
	FuncTestCaseParserRULE_precisionTimestampTZType   = 58
	FuncTestCaseParserRULE_listType                   = 59
	FuncTestCaseParserRULE_parameterizedType          = 60
	FuncTestCaseParserRULE_numericParameter           = 61
	FuncTestCaseParserRULE_substraitError             = 62
	FuncTestCaseParserRULE_funcOption                 = 63
	FuncTestCaseParserRULE_optionName                 = 64
	FuncTestCaseParserRULE_optionValue                = 65
	FuncTestCaseParserRULE_funcOptions                = 66
	FuncTestCaseParserRULE_nonReserved                = 67
	FuncTestCaseParserRULE_identifier                 = 68
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
		p.SetState(138)
		p.Header()
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4194688) != 0) || ((int64((_la-108)) & ^0x3f) == 0 && ((int64(1)<<(_la-108))&9729) != 0) {
		{
			p.SetState(139)
			p.TestGroup()
		}

		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(144)
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
		p.SetState(146)
		p.Version()
	}
	{
		p.SetState(147)
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
		p.SetState(149)
		p.Match(FuncTestCaseParserTripleHash)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(150)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserSubstraitScalarTest || _la == FuncTestCaseParserSubstraitAggregateTest) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(151)
		p.Match(FuncTestCaseParserColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(152)
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
		p.SetState(154)
		p.Match(FuncTestCaseParserTripleHash)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(155)
		p.Match(FuncTestCaseParserSubstraitInclude)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		p.Match(FuncTestCaseParserColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(157)
		p.Match(FuncTestCaseParserStringLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FuncTestCaseParserComma {
		{
			p.SetState(158)
			p.Match(FuncTestCaseParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(159)
			p.Match(FuncTestCaseParserStringLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(164)
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
		p.SetState(165)
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
	FuncOptions() IFuncOptionsContext
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

func (s *TestCaseContext) FuncOptions() IFuncOptionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncOptionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncOptionsContext)
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
		p.SetState(167)

		var _x = p.Identifier()

		localctx.(*TestCaseContext).functionName = _x
	}
	{
		p.SetState(168)
		p.Match(FuncTestCaseParserOParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
		p.Arguments()
	}
	{
		p.SetState(170)
		p.Match(FuncTestCaseParserCParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserOBracket {
		{
			p.SetState(171)
			p.Match(FuncTestCaseParserOBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(172)
			p.FuncOptions()
		}
		{
			p.SetState(173)
			p.Match(FuncTestCaseParserCBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(177)
		p.Match(FuncTestCaseParserEq)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(178)
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

	var _alt int

	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		localctx = NewScalarFuncTestGroupContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FuncTestCaseParserDescriptionLine {
			{
				p.SetState(180)
				p.TestGroupDescription()
			}

		}
		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(183)
					p.TestCase()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(186)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case 2:
		localctx = NewAggregateFuncTestGroupContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(189)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FuncTestCaseParserDescriptionLine {
			{
				p.SetState(188)
				p.TestGroupDescription()
			}

		}
		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(191)
					p.AggFuncTestCase()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(194)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
			if p.HasError() {
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
		p.SetState(198)
		p.Argument()
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FuncTestCaseParserComma {
		{
			p.SetState(199)
			p.Match(FuncTestCaseParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(200)
			p.Argument()
		}

		p.SetState(205)
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
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FuncTestCaseParserNaN, FuncTestCaseParserIntegerLiteral, FuncTestCaseParserDecimalLiteral, FuncTestCaseParserFloatLiteral, FuncTestCaseParserBooleanLiteral, FuncTestCaseParserTimestampTzLiteral, FuncTestCaseParserTimestampLiteral, FuncTestCaseParserTimeLiteral, FuncTestCaseParserDateLiteral, FuncTestCaseParserIntervalYearLiteral, FuncTestCaseParserIntervalDayLiteral, FuncTestCaseParserNullLiteral, FuncTestCaseParserStringLiteral, FuncTestCaseParserOBracket:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(206)
			p.Argument()
		}

	case FuncTestCaseParserErrorResult, FuncTestCaseParserUndefineResult:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(207)
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
	FixedCharArg() IFixedCharArgContext
	VarCharArg() IVarCharArgContext
	FixedBinaryArg() IFixedBinaryArgContext
	PrecisionTimeArg() IPrecisionTimeArgContext
	PrecisionTimestampArg() IPrecisionTimestampArgContext
	PrecisionTimestampTZArg() IPrecisionTimestampTZArgContext
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

func (s *ArgumentContext) FixedCharArg() IFixedCharArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFixedCharArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFixedCharArgContext)
}

func (s *ArgumentContext) VarCharArg() IVarCharArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarCharArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarCharArgContext)
}

func (s *ArgumentContext) FixedBinaryArg() IFixedBinaryArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFixedBinaryArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFixedBinaryArgContext)
}

func (s *ArgumentContext) PrecisionTimeArg() IPrecisionTimeArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrecisionTimeArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrecisionTimeArgContext)
}

func (s *ArgumentContext) PrecisionTimestampArg() IPrecisionTimestampArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrecisionTimestampArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrecisionTimestampArgContext)
}

func (s *ArgumentContext) PrecisionTimestampTZArg() IPrecisionTimestampTZArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrecisionTimestampTZArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrecisionTimestampTZArgContext)
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
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(210)
			p.NullArg()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(211)
			p.IntArg()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(212)
			p.FloatArg()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(213)
			p.BooleanArg()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(214)
			p.StringArg()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(215)
			p.DecimalArg()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(216)
			p.DateArg()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(217)
			p.TimeArg()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(218)
			p.TimestampArg()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(219)
			p.TimestampTzArg()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(220)
			p.IntervalYearArg()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(221)
			p.IntervalDayArg()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(222)
			p.FixedCharArg()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(223)
			p.VarCharArg()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(224)
			p.FixedBinaryArg()
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(225)
			p.PrecisionTimeArg()
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(226)
			p.PrecisionTimestampArg()
		}

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(227)
			p.PrecisionTimestampTZArg()
		}

	case 19:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(228)
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
	FuncOptions() IFuncOptionsContext
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

func (s *AggFuncTestCaseContext) FuncOptions() IFuncOptionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncOptionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncOptionsContext)
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
		p.SetState(231)
		p.AggFuncCall()
	}
	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserOBracket {
		{
			p.SetState(232)
			p.Match(FuncTestCaseParserOBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(233)
			p.FuncOptions()
		}
		{
			p.SetState(234)
			p.Match(FuncTestCaseParserCBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(238)
		p.Match(FuncTestCaseParserEq)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(239)
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
	var _la int

	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FuncTestCaseParserDefine:
		localctx = NewMultiArgAggregateFuncCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(241)
			p.TableData()
		}
		{
			p.SetState(242)

			var _x = p.Identifier()

			localctx.(*MultiArgAggregateFuncCallContext).funcName = _x
		}
		{
			p.SetState(243)
			p.Match(FuncTestCaseParserOParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(245)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32987488059392) != 0) || _la == FuncTestCaseParserOBracket || _la == FuncTestCaseParserIdentifier {
			{
				p.SetState(244)
				p.QualifiedAggregateFuncArgs()
			}

		}
		{
			p.SetState(247)
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
			p.SetState(249)
			p.TableRows()
		}
		{
			p.SetState(250)

			var _x = p.Identifier()

			localctx.(*CompactAggregateFuncCallContext).functName = _x
		}
		{
			p.SetState(251)
			p.Match(FuncTestCaseParserOParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&68171860148224) != 0) || _la == FuncTestCaseParserOBracket {
			{
				p.SetState(252)
				p.AggregateFuncArgs()
			}

		}
		{
			p.SetState(255)
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
			p.SetState(257)

			var _x = p.Identifier()

			localctx.(*SingleArgAggregateFuncCallContext).functName = _x
		}
		{
			p.SetState(258)
			p.Match(FuncTestCaseParserOParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(259)
			p.DataColumn()
		}
		{
			p.SetState(260)
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
		p.SetState(264)
		p.Match(FuncTestCaseParserDefine)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(265)

		var _m = p.Match(FuncTestCaseParserIdentifier)

		localctx.(*TableDataContext).tableName = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(266)
		p.Match(FuncTestCaseParserOParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(267)
		p.DataType()
	}
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FuncTestCaseParserComma {
		{
			p.SetState(268)
			p.Match(FuncTestCaseParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(269)
			p.DataType()
		}

		p.SetState(274)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(275)
		p.Match(FuncTestCaseParserCParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(276)
		p.Match(FuncTestCaseParserEq)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(277)
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
		p.SetState(279)
		p.Match(FuncTestCaseParserOParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserOParen {
		{
			p.SetState(280)
			p.ColumnValues()
		}
		p.SetState(285)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FuncTestCaseParserComma {
			{
				p.SetState(281)
				p.Match(FuncTestCaseParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(282)
				p.ColumnValues()
			}

			p.SetState(287)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(290)
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
		p.SetState(292)
		p.ColumnValues()
	}
	{
		p.SetState(293)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(294)
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
		p.SetState(296)
		p.Match(FuncTestCaseParserOParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32987488059392) != 0 {
		{
			p.SetState(297)
			p.Literal()
		}
		p.SetState(302)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FuncTestCaseParserComma {
			{
				p.SetState(298)
				p.Match(FuncTestCaseParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(299)
				p.Literal()
			}

			p.SetState(304)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(307)
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
	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FuncTestCaseParserNullLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(309)
			p.Match(FuncTestCaseParserNullLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserNaN, FuncTestCaseParserIntegerLiteral, FuncTestCaseParserDecimalLiteral, FuncTestCaseParserFloatLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(310)
			p.NumericLiteral()
		}

	case FuncTestCaseParserBooleanLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(311)
			p.Match(FuncTestCaseParserBooleanLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserStringLiteral:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(312)
			p.Match(FuncTestCaseParserStringLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserDateLiteral:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(313)
			p.Match(FuncTestCaseParserDateLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserTimeLiteral:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(314)
			p.Match(FuncTestCaseParserTimeLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserTimestampLiteral:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(315)
			p.Match(FuncTestCaseParserTimestampLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserTimestampTzLiteral:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(316)
			p.Match(FuncTestCaseParserTimestampTzLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserIntervalYearLiteral:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(317)
			p.Match(FuncTestCaseParserIntervalYearLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserIntervalDayLiteral:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(318)
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
		p.SetState(321)
		p.QualifiedAggregateFuncArg()
	}
	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FuncTestCaseParserComma {
		{
			p.SetState(322)
			p.Match(FuncTestCaseParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(323)
			p.QualifiedAggregateFuncArg()
		}

		p.SetState(328)
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
		p.SetState(329)
		p.AggregateFuncArg()
	}
	p.SetState(334)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FuncTestCaseParserComma {
		{
			p.SetState(330)
			p.Match(FuncTestCaseParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)
			p.AggregateFuncArg()
		}

		p.SetState(336)
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
	p.SetState(341)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FuncTestCaseParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(337)

			var _m = p.Match(FuncTestCaseParserIdentifier)

			localctx.(*QualifiedAggregateFuncArgContext).tableName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(338)
			p.Match(FuncTestCaseParserDot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(339)
			p.Match(FuncTestCaseParserColumnName)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserNaN, FuncTestCaseParserIntegerLiteral, FuncTestCaseParserDecimalLiteral, FuncTestCaseParserFloatLiteral, FuncTestCaseParserBooleanLiteral, FuncTestCaseParserTimestampTzLiteral, FuncTestCaseParserTimestampLiteral, FuncTestCaseParserTimeLiteral, FuncTestCaseParserDateLiteral, FuncTestCaseParserIntervalYearLiteral, FuncTestCaseParserIntervalDayLiteral, FuncTestCaseParserNullLiteral, FuncTestCaseParserStringLiteral, FuncTestCaseParserOBracket:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(340)
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
	p.SetState(347)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FuncTestCaseParserColumnName:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(343)
			p.Match(FuncTestCaseParserColumnName)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(344)
			p.Match(FuncTestCaseParserDoubleColon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(345)
			p.DataType()
		}

	case FuncTestCaseParserNaN, FuncTestCaseParserIntegerLiteral, FuncTestCaseParserDecimalLiteral, FuncTestCaseParserFloatLiteral, FuncTestCaseParserBooleanLiteral, FuncTestCaseParserTimestampTzLiteral, FuncTestCaseParserTimestampLiteral, FuncTestCaseParserTimeLiteral, FuncTestCaseParserDateLiteral, FuncTestCaseParserIntervalYearLiteral, FuncTestCaseParserIntervalDayLiteral, FuncTestCaseParserNullLiteral, FuncTestCaseParserStringLiteral, FuncTestCaseParserOBracket:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(346)
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
	p.SetState(352)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FuncTestCaseParserDecimalLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(349)
			p.Match(FuncTestCaseParserDecimalLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserIntegerLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(350)
			p.Match(FuncTestCaseParserIntegerLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserNaN, FuncTestCaseParserFloatLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(351)
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
		p.SetState(354)
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
		p.SetState(356)
		p.Match(FuncTestCaseParserNullLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(357)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(358)
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

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// Getter signatures
	IntegerLiteral() antlr.TerminalNode
	DoubleColon() antlr.TerminalNode
	I8() antlr.TerminalNode
	I16() antlr.TerminalNode
	I32() antlr.TerminalNode
	I64() antlr.TerminalNode
	QMark() antlr.TerminalNode

	// IsIntArgContext differentiates from other interfaces.
	IsIntArgContext()
}

type IntArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
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

func (s *IntArgContext) GetIsnull() antlr.Token { return s.isnull }

func (s *IntArgContext) SetIsnull(v antlr.Token) { s.isnull = v }

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

func (s *IntArgContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *IntArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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
		p.SetState(360)
		p.Match(FuncTestCaseParserIntegerLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(361)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(362)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&67553994410557440) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(364)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(363)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*IntArgContext).isnull = _m
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

// IFloatArgContext is an interface to support dynamic dispatch.
type IFloatArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// Getter signatures
	NumericLiteral() INumericLiteralContext
	DoubleColon() antlr.TerminalNode
	FP32() antlr.TerminalNode
	FP64() antlr.TerminalNode
	QMark() antlr.TerminalNode

	// IsFloatArgContext differentiates from other interfaces.
	IsFloatArgContext()
}

type FloatArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
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

func (s *FloatArgContext) GetIsnull() antlr.Token { return s.isnull }

func (s *FloatArgContext) SetIsnull(v antlr.Token) { s.isnull = v }

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

func (s *FloatArgContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *FloatArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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
		p.SetState(366)
		p.NumericLiteral()
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
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserFP32 || _la == FuncTestCaseParserFP64) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(370)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(369)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*FloatArgContext).isnull = _m
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

// IDecimalArgContext is an interface to support dynamic dispatch.
type IDecimalArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// Getter signatures
	NumericLiteral() INumericLiteralContext
	DoubleColon() antlr.TerminalNode
	DecimalType() IDecimalTypeContext
	QMark() antlr.TerminalNode

	// IsDecimalArgContext differentiates from other interfaces.
	IsDecimalArgContext()
}

type DecimalArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
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

func (s *DecimalArgContext) GetIsnull() antlr.Token { return s.isnull }

func (s *DecimalArgContext) SetIsnull(v antlr.Token) { s.isnull = v }

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

func (s *DecimalArgContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *DecimalArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(372)
		p.NumericLiteral()
	}
	{
		p.SetState(373)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(374)
		p.DecimalType()
	}
	p.SetState(376)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(375)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*DecimalArgContext).isnull = _m
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

// IBooleanArgContext is an interface to support dynamic dispatch.
type IBooleanArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// Getter signatures
	BooleanLiteral() antlr.TerminalNode
	DoubleColon() antlr.TerminalNode
	BooleanType() IBooleanTypeContext
	QMark() antlr.TerminalNode

	// IsBooleanArgContext differentiates from other interfaces.
	IsBooleanArgContext()
}

type BooleanArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
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

func (s *BooleanArgContext) GetIsnull() antlr.Token { return s.isnull }

func (s *BooleanArgContext) SetIsnull(v antlr.Token) { s.isnull = v }

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

func (s *BooleanArgContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *BooleanArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(378)
		p.Match(FuncTestCaseParserBooleanLiteral)
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
		p.BooleanType()
	}
	p.SetState(382)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(381)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*BooleanArgContext).isnull = _m
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

// IStringArgContext is an interface to support dynamic dispatch.
type IStringArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// Getter signatures
	StringLiteral() antlr.TerminalNode
	DoubleColon() antlr.TerminalNode
	StringType() IStringTypeContext
	QMark() antlr.TerminalNode

	// IsStringArgContext differentiates from other interfaces.
	IsStringArgContext()
}

type StringArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
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

func (s *StringArgContext) GetIsnull() antlr.Token { return s.isnull }

func (s *StringArgContext) SetIsnull(v antlr.Token) { s.isnull = v }

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

func (s *StringArgContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *StringArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(384)
		p.Match(FuncTestCaseParserStringLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(385)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(386)
		p.StringType()
	}
	p.SetState(388)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(387)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*StringArgContext).isnull = _m
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

// IDateArgContext is an interface to support dynamic dispatch.
type IDateArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// Getter signatures
	DateLiteral() antlr.TerminalNode
	DoubleColon() antlr.TerminalNode
	Date() antlr.TerminalNode
	QMark() antlr.TerminalNode

	// IsDateArgContext differentiates from other interfaces.
	IsDateArgContext()
}

type DateArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
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

func (s *DateArgContext) GetIsnull() antlr.Token { return s.isnull }

func (s *DateArgContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *DateArgContext) DateLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDateLiteral, 0)
}

func (s *DateArgContext) DoubleColon() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDoubleColon, 0)
}

func (s *DateArgContext) Date() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDate, 0)
}

func (s *DateArgContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *DateArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(390)
		p.Match(FuncTestCaseParserDateLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(391)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(392)
		p.Match(FuncTestCaseParserDate)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(394)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(393)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*DateArgContext).isnull = _m
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

// ITimeArgContext is an interface to support dynamic dispatch.
type ITimeArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// Getter signatures
	TimeLiteral() antlr.TerminalNode
	DoubleColon() antlr.TerminalNode
	Time() antlr.TerminalNode
	QMark() antlr.TerminalNode

	// IsTimeArgContext differentiates from other interfaces.
	IsTimeArgContext()
}

type TimeArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
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

func (s *TimeArgContext) GetIsnull() antlr.Token { return s.isnull }

func (s *TimeArgContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *TimeArgContext) TimeLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserTimeLiteral, 0)
}

func (s *TimeArgContext) DoubleColon() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDoubleColon, 0)
}

func (s *TimeArgContext) Time() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserTime, 0)
}

func (s *TimeArgContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *TimeArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(396)
		p.Match(FuncTestCaseParserTimeLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(397)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(398)
		p.Match(FuncTestCaseParserTime)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(400)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(399)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*TimeArgContext).isnull = _m
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

// ITimestampArgContext is an interface to support dynamic dispatch.
type ITimestampArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// Getter signatures
	TimestampLiteral() antlr.TerminalNode
	DoubleColon() antlr.TerminalNode
	TimestampType() ITimestampTypeContext
	QMark() antlr.TerminalNode

	// IsTimestampArgContext differentiates from other interfaces.
	IsTimestampArgContext()
}

type TimestampArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
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

func (s *TimestampArgContext) GetIsnull() antlr.Token { return s.isnull }

func (s *TimestampArgContext) SetIsnull(v antlr.Token) { s.isnull = v }

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

func (s *TimestampArgContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *TimestampArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimestampArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(402)
		p.Match(FuncTestCaseParserTimestampLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(403)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(404)
		p.TimestampType()
	}
	p.SetState(406)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(405)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*TimestampArgContext).isnull = _m
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

// ITimestampTzArgContext is an interface to support dynamic dispatch.
type ITimestampTzArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// Getter signatures
	TimestampTzLiteral() antlr.TerminalNode
	DoubleColon() antlr.TerminalNode
	TimestampTZType() ITimestampTZTypeContext
	QMark() antlr.TerminalNode

	// IsTimestampTzArgContext differentiates from other interfaces.
	IsTimestampTzArgContext()
}

type TimestampTzArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
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

func (s *TimestampTzArgContext) GetIsnull() antlr.Token { return s.isnull }

func (s *TimestampTzArgContext) SetIsnull(v antlr.Token) { s.isnull = v }

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

func (s *TimestampTzArgContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *TimestampTzArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimestampTzArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(408)
		p.Match(FuncTestCaseParserTimestampTzLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(409)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(410)
		p.TimestampTZType()
	}
	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(411)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*TimestampTzArgContext).isnull = _m
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

// IIntervalYearArgContext is an interface to support dynamic dispatch.
type IIntervalYearArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// Getter signatures
	IntervalYearLiteral() antlr.TerminalNode
	DoubleColon() antlr.TerminalNode
	IntervalYearType() IIntervalYearTypeContext
	QMark() antlr.TerminalNode

	// IsIntervalYearArgContext differentiates from other interfaces.
	IsIntervalYearArgContext()
}

type IntervalYearArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
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

func (s *IntervalYearArgContext) GetIsnull() antlr.Token { return s.isnull }

func (s *IntervalYearArgContext) SetIsnull(v antlr.Token) { s.isnull = v }

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

func (s *IntervalYearArgContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *IntervalYearArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntervalYearArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(414)
		p.Match(FuncTestCaseParserIntervalYearLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(415)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(416)
		p.IntervalYearType()
	}
	p.SetState(418)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(417)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*IntervalYearArgContext).isnull = _m
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

// IIntervalDayArgContext is an interface to support dynamic dispatch.
type IIntervalDayArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// Getter signatures
	IntervalDayLiteral() antlr.TerminalNode
	DoubleColon() antlr.TerminalNode
	IntervalDayType() IIntervalDayTypeContext
	QMark() antlr.TerminalNode

	// IsIntervalDayArgContext differentiates from other interfaces.
	IsIntervalDayArgContext()
}

type IntervalDayArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
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

func (s *IntervalDayArgContext) GetIsnull() antlr.Token { return s.isnull }

func (s *IntervalDayArgContext) SetIsnull(v antlr.Token) { s.isnull = v }

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

func (s *IntervalDayArgContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *IntervalDayArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntervalDayArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(420)
		p.Match(FuncTestCaseParserIntervalDayLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(421)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(422)
		p.IntervalDayType()
	}
	p.SetState(424)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(423)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*IntervalDayArgContext).isnull = _m
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

// IFixedCharArgContext is an interface to support dynamic dispatch.
type IFixedCharArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// Getter signatures
	StringLiteral() antlr.TerminalNode
	DoubleColon() antlr.TerminalNode
	FixedCharType() IFixedCharTypeContext
	QMark() antlr.TerminalNode

	// IsFixedCharArgContext differentiates from other interfaces.
	IsFixedCharArgContext()
}

type FixedCharArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
}

func NewEmptyFixedCharArgContext() *FixedCharArgContext {
	var p = new(FixedCharArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_fixedCharArg
	return p
}

func InitEmptyFixedCharArgContext(p *FixedCharArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_fixedCharArg
}

func (*FixedCharArgContext) IsFixedCharArgContext() {}

func NewFixedCharArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FixedCharArgContext {
	var p = new(FixedCharArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_fixedCharArg

	return p
}

func (s *FixedCharArgContext) GetParser() antlr.Parser { return s.parser }

func (s *FixedCharArgContext) GetIsnull() antlr.Token { return s.isnull }

func (s *FixedCharArgContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *FixedCharArgContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserStringLiteral, 0)
}

func (s *FixedCharArgContext) DoubleColon() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDoubleColon, 0)
}

func (s *FixedCharArgContext) FixedCharType() IFixedCharTypeContext {
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

func (s *FixedCharArgContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *FixedCharArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FixedCharArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FixedCharArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitFixedCharArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) FixedCharArg() (localctx IFixedCharArgContext) {
	localctx = NewFixedCharArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, FuncTestCaseParserRULE_fixedCharArg)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(426)
		p.Match(FuncTestCaseParserStringLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(427)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(428)
		p.FixedCharType()
	}
	p.SetState(430)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(429)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*FixedCharArgContext).isnull = _m
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

// IVarCharArgContext is an interface to support dynamic dispatch.
type IVarCharArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// Getter signatures
	StringLiteral() antlr.TerminalNode
	DoubleColon() antlr.TerminalNode
	VarCharType() IVarCharTypeContext
	QMark() antlr.TerminalNode

	// IsVarCharArgContext differentiates from other interfaces.
	IsVarCharArgContext()
}

type VarCharArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
}

func NewEmptyVarCharArgContext() *VarCharArgContext {
	var p = new(VarCharArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_varCharArg
	return p
}

func InitEmptyVarCharArgContext(p *VarCharArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_varCharArg
}

func (*VarCharArgContext) IsVarCharArgContext() {}

func NewVarCharArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarCharArgContext {
	var p = new(VarCharArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_varCharArg

	return p
}

func (s *VarCharArgContext) GetParser() antlr.Parser { return s.parser }

func (s *VarCharArgContext) GetIsnull() antlr.Token { return s.isnull }

func (s *VarCharArgContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *VarCharArgContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserStringLiteral, 0)
}

func (s *VarCharArgContext) DoubleColon() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDoubleColon, 0)
}

func (s *VarCharArgContext) VarCharType() IVarCharTypeContext {
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

func (s *VarCharArgContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *VarCharArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarCharArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarCharArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitVarCharArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) VarCharArg() (localctx IVarCharArgContext) {
	localctx = NewVarCharArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, FuncTestCaseParserRULE_varCharArg)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(432)
		p.Match(FuncTestCaseParserStringLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(433)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(434)
		p.VarCharType()
	}
	p.SetState(436)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(435)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*VarCharArgContext).isnull = _m
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

// IFixedBinaryArgContext is an interface to support dynamic dispatch.
type IFixedBinaryArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// Getter signatures
	StringLiteral() antlr.TerminalNode
	DoubleColon() antlr.TerminalNode
	FixedBinaryType() IFixedBinaryTypeContext
	QMark() antlr.TerminalNode

	// IsFixedBinaryArgContext differentiates from other interfaces.
	IsFixedBinaryArgContext()
}

type FixedBinaryArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
}

func NewEmptyFixedBinaryArgContext() *FixedBinaryArgContext {
	var p = new(FixedBinaryArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_fixedBinaryArg
	return p
}

func InitEmptyFixedBinaryArgContext(p *FixedBinaryArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_fixedBinaryArg
}

func (*FixedBinaryArgContext) IsFixedBinaryArgContext() {}

func NewFixedBinaryArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FixedBinaryArgContext {
	var p = new(FixedBinaryArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_fixedBinaryArg

	return p
}

func (s *FixedBinaryArgContext) GetParser() antlr.Parser { return s.parser }

func (s *FixedBinaryArgContext) GetIsnull() antlr.Token { return s.isnull }

func (s *FixedBinaryArgContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *FixedBinaryArgContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserStringLiteral, 0)
}

func (s *FixedBinaryArgContext) DoubleColon() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDoubleColon, 0)
}

func (s *FixedBinaryArgContext) FixedBinaryType() IFixedBinaryTypeContext {
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

func (s *FixedBinaryArgContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *FixedBinaryArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FixedBinaryArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FixedBinaryArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitFixedBinaryArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) FixedBinaryArg() (localctx IFixedBinaryArgContext) {
	localctx = NewFixedBinaryArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, FuncTestCaseParserRULE_fixedBinaryArg)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(438)
		p.Match(FuncTestCaseParserStringLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(439)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(440)
		p.FixedBinaryType()
	}
	p.SetState(442)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(441)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*FixedBinaryArgContext).isnull = _m
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

// IPrecisionTimeArgContext is an interface to support dynamic dispatch.
type IPrecisionTimeArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// Getter signatures
	TimeLiteral() antlr.TerminalNode
	DoubleColon() antlr.TerminalNode
	PrecisionTimeType() IPrecisionTimeTypeContext
	QMark() antlr.TerminalNode

	// IsPrecisionTimeArgContext differentiates from other interfaces.
	IsPrecisionTimeArgContext()
}

type PrecisionTimeArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
}

func NewEmptyPrecisionTimeArgContext() *PrecisionTimeArgContext {
	var p = new(PrecisionTimeArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_precisionTimeArg
	return p
}

func InitEmptyPrecisionTimeArgContext(p *PrecisionTimeArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_precisionTimeArg
}

func (*PrecisionTimeArgContext) IsPrecisionTimeArgContext() {}

func NewPrecisionTimeArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrecisionTimeArgContext {
	var p = new(PrecisionTimeArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_precisionTimeArg

	return p
}

func (s *PrecisionTimeArgContext) GetParser() antlr.Parser { return s.parser }

func (s *PrecisionTimeArgContext) GetIsnull() antlr.Token { return s.isnull }

func (s *PrecisionTimeArgContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *PrecisionTimeArgContext) TimeLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserTimeLiteral, 0)
}

func (s *PrecisionTimeArgContext) DoubleColon() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDoubleColon, 0)
}

func (s *PrecisionTimeArgContext) PrecisionTimeType() IPrecisionTimeTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrecisionTimeTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrecisionTimeTypeContext)
}

func (s *PrecisionTimeArgContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *PrecisionTimeArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrecisionTimeArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrecisionTimeArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitPrecisionTimeArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) PrecisionTimeArg() (localctx IPrecisionTimeArgContext) {
	localctx = NewPrecisionTimeArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, FuncTestCaseParserRULE_precisionTimeArg)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(444)
		p.Match(FuncTestCaseParserTimeLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(445)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(446)
		p.PrecisionTimeType()
	}
	p.SetState(448)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(447)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*PrecisionTimeArgContext).isnull = _m
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

// IPrecisionTimestampArgContext is an interface to support dynamic dispatch.
type IPrecisionTimestampArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// Getter signatures
	TimestampLiteral() antlr.TerminalNode
	DoubleColon() antlr.TerminalNode
	PrecisionTimestampType() IPrecisionTimestampTypeContext
	QMark() antlr.TerminalNode

	// IsPrecisionTimestampArgContext differentiates from other interfaces.
	IsPrecisionTimestampArgContext()
}

type PrecisionTimestampArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
}

func NewEmptyPrecisionTimestampArgContext() *PrecisionTimestampArgContext {
	var p = new(PrecisionTimestampArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_precisionTimestampArg
	return p
}

func InitEmptyPrecisionTimestampArgContext(p *PrecisionTimestampArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_precisionTimestampArg
}

func (*PrecisionTimestampArgContext) IsPrecisionTimestampArgContext() {}

func NewPrecisionTimestampArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrecisionTimestampArgContext {
	var p = new(PrecisionTimestampArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_precisionTimestampArg

	return p
}

func (s *PrecisionTimestampArgContext) GetParser() antlr.Parser { return s.parser }

func (s *PrecisionTimestampArgContext) GetIsnull() antlr.Token { return s.isnull }

func (s *PrecisionTimestampArgContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *PrecisionTimestampArgContext) TimestampLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserTimestampLiteral, 0)
}

func (s *PrecisionTimestampArgContext) DoubleColon() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDoubleColon, 0)
}

func (s *PrecisionTimestampArgContext) PrecisionTimestampType() IPrecisionTimestampTypeContext {
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

func (s *PrecisionTimestampArgContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *PrecisionTimestampArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrecisionTimestampArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrecisionTimestampArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitPrecisionTimestampArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) PrecisionTimestampArg() (localctx IPrecisionTimestampArgContext) {
	localctx = NewPrecisionTimestampArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, FuncTestCaseParserRULE_precisionTimestampArg)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(450)
		p.Match(FuncTestCaseParserTimestampLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(451)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(452)
		p.PrecisionTimestampType()
	}
	p.SetState(454)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(453)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*PrecisionTimestampArgContext).isnull = _m
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

// IPrecisionTimestampTZArgContext is an interface to support dynamic dispatch.
type IPrecisionTimestampTZArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// Getter signatures
	TimestampTzLiteral() antlr.TerminalNode
	DoubleColon() antlr.TerminalNode
	PrecisionTimestampTZType() IPrecisionTimestampTZTypeContext
	QMark() antlr.TerminalNode

	// IsPrecisionTimestampTZArgContext differentiates from other interfaces.
	IsPrecisionTimestampTZArgContext()
}

type PrecisionTimestampTZArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
}

func NewEmptyPrecisionTimestampTZArgContext() *PrecisionTimestampTZArgContext {
	var p = new(PrecisionTimestampTZArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_precisionTimestampTZArg
	return p
}

func InitEmptyPrecisionTimestampTZArgContext(p *PrecisionTimestampTZArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_precisionTimestampTZArg
}

func (*PrecisionTimestampTZArgContext) IsPrecisionTimestampTZArgContext() {}

func NewPrecisionTimestampTZArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrecisionTimestampTZArgContext {
	var p = new(PrecisionTimestampTZArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_precisionTimestampTZArg

	return p
}

func (s *PrecisionTimestampTZArgContext) GetParser() antlr.Parser { return s.parser }

func (s *PrecisionTimestampTZArgContext) GetIsnull() antlr.Token { return s.isnull }

func (s *PrecisionTimestampTZArgContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *PrecisionTimestampTZArgContext) TimestampTzLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserTimestampTzLiteral, 0)
}

func (s *PrecisionTimestampTZArgContext) DoubleColon() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDoubleColon, 0)
}

func (s *PrecisionTimestampTZArgContext) PrecisionTimestampTZType() IPrecisionTimestampTZTypeContext {
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

func (s *PrecisionTimestampTZArgContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *PrecisionTimestampTZArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrecisionTimestampTZArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrecisionTimestampTZArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitPrecisionTimestampTZArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) PrecisionTimestampTZArg() (localctx IPrecisionTimestampTZArgContext) {
	localctx = NewPrecisionTimestampTZArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, FuncTestCaseParserRULE_precisionTimestampTZArg)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(456)
		p.Match(FuncTestCaseParserTimestampTzLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(457)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(458)
		p.PrecisionTimestampTZType()
	}
	p.SetState(460)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(459)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*PrecisionTimestampTZArgContext).isnull = _m
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

// IListArgContext is an interface to support dynamic dispatch.
type IListArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// Getter signatures
	LiteralList() ILiteralListContext
	DoubleColon() antlr.TerminalNode
	ListType() IListTypeContext
	QMark() antlr.TerminalNode

	// IsListArgContext differentiates from other interfaces.
	IsListArgContext()
}

type ListArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
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

func (s *ListArgContext) GetIsnull() antlr.Token { return s.isnull }

func (s *ListArgContext) SetIsnull(v antlr.Token) { s.isnull = v }

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

func (s *ListArgContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *ListArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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
	p.EnterRule(localctx, 82, FuncTestCaseParserRULE_listArg)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(462)
		p.LiteralList()
	}
	{
		p.SetState(463)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(464)
		p.ListType()
	}
	p.SetState(466)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(465)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*ListArgContext).isnull = _m
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
	p.EnterRule(localctx, 84, FuncTestCaseParserRULE_literalList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(468)
		p.Match(FuncTestCaseParserOBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(477)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32987488059392) != 0 {
		{
			p.SetState(469)
			p.Literal()
		}
		p.SetState(474)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FuncTestCaseParserComma {
			{
				p.SetState(470)
				p.Match(FuncTestCaseParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(471)
				p.Literal()
			}

			p.SetState(476)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(479)
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

// IDataTypeContext is an interface to support dynamic dispatch.
type IDataTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// Getter signatures
	ScalarType() IScalarTypeContext
	QMark() antlr.TerminalNode
	ParameterizedType() IParameterizedTypeContext

	// IsDataTypeContext differentiates from other interfaces.
	IsDataTypeContext()
}

type DataTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
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

func (s *DataTypeContext) GetIsnull() antlr.Token { return s.isnull }

func (s *DataTypeContext) SetIsnull(v antlr.Token) { s.isnull = v }

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

func (s *DataTypeContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
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
	p.EnterRule(localctx, 86, FuncTestCaseParserRULE_dataType)
	var _la int

	p.SetState(486)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FuncTestCaseParserBoolean, FuncTestCaseParserI8, FuncTestCaseParserI16, FuncTestCaseParserI32, FuncTestCaseParserI64, FuncTestCaseParserFP32, FuncTestCaseParserFP64, FuncTestCaseParserString_, FuncTestCaseParserBinary, FuncTestCaseParserTimestamp, FuncTestCaseParserTimestamp_TZ, FuncTestCaseParserDate, FuncTestCaseParserTime, FuncTestCaseParserInterval_Year, FuncTestCaseParserUUID, FuncTestCaseParserUserDefined, FuncTestCaseParserBool, FuncTestCaseParserStr, FuncTestCaseParserVBin, FuncTestCaseParserTs, FuncTestCaseParserTsTZ, FuncTestCaseParserIYear:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(481)
			p.ScalarType()
		}
		p.SetState(483)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FuncTestCaseParserQMark {
			{
				p.SetState(482)

				var _m = p.Match(FuncTestCaseParserQMark)

				localctx.(*DataTypeContext).isnull = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case FuncTestCaseParserInterval_Day, FuncTestCaseParserDecimal, FuncTestCaseParserPrecision_Time, FuncTestCaseParserPrecision_Timestamp, FuncTestCaseParserPrecision_Timestamp_TZ, FuncTestCaseParserFixedChar, FuncTestCaseParserVarChar, FuncTestCaseParserFixedBinary, FuncTestCaseParserIDay, FuncTestCaseParserDec, FuncTestCaseParserPT, FuncTestCaseParserPTs, FuncTestCaseParserPTsTZ, FuncTestCaseParserFChar, FuncTestCaseParserVChar, FuncTestCaseParserFBin:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(485)
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

func (s *BinaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitBinary(s)

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
	p.EnterRule(localctx, 88, FuncTestCaseParserRULE_scalarType)
	p.SetState(505)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FuncTestCaseParserBoolean, FuncTestCaseParserBool:
		localctx = NewBooleanContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(488)
			p.BooleanType()
		}

	case FuncTestCaseParserI8:
		localctx = NewI8Context(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(489)
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
			p.SetState(490)
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
			p.SetState(491)
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
			p.SetState(492)
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
			p.SetState(493)
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
			p.SetState(494)
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
			p.SetState(495)
			p.StringType()
		}

	case FuncTestCaseParserBinary, FuncTestCaseParserVBin:
		localctx = NewBinaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(496)
			p.BinaryType()
		}

	case FuncTestCaseParserTimestamp, FuncTestCaseParserTs:
		localctx = NewTimestampContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(497)
			p.TimestampType()
		}

	case FuncTestCaseParserTimestamp_TZ, FuncTestCaseParserTsTZ:
		localctx = NewTimestampTzContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(498)
			p.TimestampTZType()
		}

	case FuncTestCaseParserDate:
		localctx = NewDateContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(499)
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
			p.SetState(500)
			p.Match(FuncTestCaseParserTime)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserInterval_Year, FuncTestCaseParserIYear:
		localctx = NewIntervalYearContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(501)
			p.IntervalYearType()
		}

	case FuncTestCaseParserUUID:
		localctx = NewUuidContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(502)
			p.Match(FuncTestCaseParserUUID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserUserDefined:
		localctx = NewUserDefinedContext(p, localctx)
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(503)
			p.Match(FuncTestCaseParserUserDefined)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(504)
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
	p.EnterRule(localctx, 90, FuncTestCaseParserRULE_booleanType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(507)
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
	p.EnterRule(localctx, 92, FuncTestCaseParserRULE_stringType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(509)
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
	p.EnterRule(localctx, 94, FuncTestCaseParserRULE_binaryType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(511)
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
	p.EnterRule(localctx, 96, FuncTestCaseParserRULE_timestampType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(513)
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
	p.EnterRule(localctx, 98, FuncTestCaseParserRULE_timestampTZType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(515)
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
	p.EnterRule(localctx, 100, FuncTestCaseParserRULE_intervalYearType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(517)
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

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// GetLen_ returns the len_ rule contexts.
	GetLen_() INumericParameterContext

	// SetLen_ sets the len_ rule contexts.
	SetLen_(INumericParameterContext)

	// Getter signatures
	IDay() antlr.TerminalNode
	Interval_Day() antlr.TerminalNode
	OAngleBracket() antlr.TerminalNode
	CAngleBracket() antlr.TerminalNode
	QMark() antlr.TerminalNode
	NumericParameter() INumericParameterContext

	// IsIntervalDayTypeContext differentiates from other interfaces.
	IsIntervalDayTypeContext()
}

type IntervalDayTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
	len_   INumericParameterContext
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

func (s *IntervalDayTypeContext) GetIsnull() antlr.Token { return s.isnull }

func (s *IntervalDayTypeContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *IntervalDayTypeContext) GetLen_() INumericParameterContext { return s.len_ }

func (s *IntervalDayTypeContext) SetLen_(v INumericParameterContext) { s.len_ = v }

func (s *IntervalDayTypeContext) IDay() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserIDay, 0)
}

func (s *IntervalDayTypeContext) Interval_Day() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserInterval_Day, 0)
}

func (s *IntervalDayTypeContext) OAngleBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserOAngleBracket, 0)
}

func (s *IntervalDayTypeContext) CAngleBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserCAngleBracket, 0)
}

func (s *IntervalDayTypeContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *IntervalDayTypeContext) NumericParameter() INumericParameterContext {
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

func (s *IntervalDayTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntervalDayTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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
	p.EnterRule(localctx, 102, FuncTestCaseParserRULE_intervalDayType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(519)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserInterval_Day || _la == FuncTestCaseParserIDay) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(521)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(520)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*IntervalDayTypeContext).isnull = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(527)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserOAngleBracket {
		{
			p.SetState(523)
			p.Match(FuncTestCaseParserOAngleBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(524)

			var _x = p.NumericParameter()

			localctx.(*IntervalDayTypeContext).len_ = _x
		}
		{
			p.SetState(525)
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

// IFixedCharTypeContext is an interface to support dynamic dispatch.
type IFixedCharTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// GetLen_ returns the len_ rule contexts.
	GetLen_() INumericParameterContext

	// SetLen_ sets the len_ rule contexts.
	SetLen_(INumericParameterContext)

	// Getter signatures
	OAngleBracket() antlr.TerminalNode
	CAngleBracket() antlr.TerminalNode
	FChar() antlr.TerminalNode
	FixedChar() antlr.TerminalNode
	NumericParameter() INumericParameterContext
	QMark() antlr.TerminalNode

	// IsFixedCharTypeContext differentiates from other interfaces.
	IsFixedCharTypeContext()
}

type FixedCharTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
	len_   INumericParameterContext
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

func (s *FixedCharTypeContext) GetIsnull() antlr.Token { return s.isnull }

func (s *FixedCharTypeContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *FixedCharTypeContext) GetLen_() INumericParameterContext { return s.len_ }

func (s *FixedCharTypeContext) SetLen_(v INumericParameterContext) { s.len_ = v }

func (s *FixedCharTypeContext) OAngleBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserOAngleBracket, 0)
}

func (s *FixedCharTypeContext) CAngleBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserCAngleBracket, 0)
}

func (s *FixedCharTypeContext) FChar() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserFChar, 0)
}

func (s *FixedCharTypeContext) FixedChar() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserFixedChar, 0)
}

func (s *FixedCharTypeContext) NumericParameter() INumericParameterContext {
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

func (s *FixedCharTypeContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *FixedCharTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FixedCharTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FixedCharTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitFixedCharType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) FixedCharType() (localctx IFixedCharTypeContext) {
	localctx = NewFixedCharTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, FuncTestCaseParserRULE_fixedCharType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(529)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserFixedChar || _la == FuncTestCaseParserFChar) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(531)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(530)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*FixedCharTypeContext).isnull = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(533)
		p.Match(FuncTestCaseParserOAngleBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(534)

		var _x = p.NumericParameter()

		localctx.(*FixedCharTypeContext).len_ = _x
	}
	{
		p.SetState(535)
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

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// GetLen_ returns the len_ rule contexts.
	GetLen_() INumericParameterContext

	// SetLen_ sets the len_ rule contexts.
	SetLen_(INumericParameterContext)

	// Getter signatures
	OAngleBracket() antlr.TerminalNode
	CAngleBracket() antlr.TerminalNode
	VChar() antlr.TerminalNode
	VarChar() antlr.TerminalNode
	NumericParameter() INumericParameterContext
	QMark() antlr.TerminalNode

	// IsVarCharTypeContext differentiates from other interfaces.
	IsVarCharTypeContext()
}

type VarCharTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
	len_   INumericParameterContext
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

func (s *VarCharTypeContext) GetIsnull() antlr.Token { return s.isnull }

func (s *VarCharTypeContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *VarCharTypeContext) GetLen_() INumericParameterContext { return s.len_ }

func (s *VarCharTypeContext) SetLen_(v INumericParameterContext) { s.len_ = v }

func (s *VarCharTypeContext) OAngleBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserOAngleBracket, 0)
}

func (s *VarCharTypeContext) CAngleBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserCAngleBracket, 0)
}

func (s *VarCharTypeContext) VChar() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserVChar, 0)
}

func (s *VarCharTypeContext) VarChar() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserVarChar, 0)
}

func (s *VarCharTypeContext) NumericParameter() INumericParameterContext {
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

func (s *VarCharTypeContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *VarCharTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarCharTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarCharTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitVarCharType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) VarCharType() (localctx IVarCharTypeContext) {
	localctx = NewVarCharTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, FuncTestCaseParserRULE_varCharType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(537)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserVarChar || _la == FuncTestCaseParserVChar) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(539)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(538)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*VarCharTypeContext).isnull = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(541)
		p.Match(FuncTestCaseParserOAngleBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(542)

		var _x = p.NumericParameter()

		localctx.(*VarCharTypeContext).len_ = _x
	}
	{
		p.SetState(543)
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

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// GetLen_ returns the len_ rule contexts.
	GetLen_() INumericParameterContext

	// SetLen_ sets the len_ rule contexts.
	SetLen_(INumericParameterContext)

	// Getter signatures
	OAngleBracket() antlr.TerminalNode
	CAngleBracket() antlr.TerminalNode
	FBin() antlr.TerminalNode
	FixedBinary() antlr.TerminalNode
	NumericParameter() INumericParameterContext
	QMark() antlr.TerminalNode

	// IsFixedBinaryTypeContext differentiates from other interfaces.
	IsFixedBinaryTypeContext()
}

type FixedBinaryTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
	len_   INumericParameterContext
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

func (s *FixedBinaryTypeContext) GetIsnull() antlr.Token { return s.isnull }

func (s *FixedBinaryTypeContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *FixedBinaryTypeContext) GetLen_() INumericParameterContext { return s.len_ }

func (s *FixedBinaryTypeContext) SetLen_(v INumericParameterContext) { s.len_ = v }

func (s *FixedBinaryTypeContext) OAngleBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserOAngleBracket, 0)
}

func (s *FixedBinaryTypeContext) CAngleBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserCAngleBracket, 0)
}

func (s *FixedBinaryTypeContext) FBin() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserFBin, 0)
}

func (s *FixedBinaryTypeContext) FixedBinary() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserFixedBinary, 0)
}

func (s *FixedBinaryTypeContext) NumericParameter() INumericParameterContext {
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

func (s *FixedBinaryTypeContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *FixedBinaryTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FixedBinaryTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FixedBinaryTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitFixedBinaryType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) FixedBinaryType() (localctx IFixedBinaryTypeContext) {
	localctx = NewFixedBinaryTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, FuncTestCaseParserRULE_fixedBinaryType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(545)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserFixedBinary || _la == FuncTestCaseParserFBin) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(547)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(546)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*FixedBinaryTypeContext).isnull = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(549)
		p.Match(FuncTestCaseParserOAngleBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(550)

		var _x = p.NumericParameter()

		localctx.(*FixedBinaryTypeContext).len_ = _x
	}
	{
		p.SetState(551)
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

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// GetPrecision returns the precision rule contexts.
	GetPrecision() INumericParameterContext

	// GetScale returns the scale rule contexts.
	GetScale() INumericParameterContext

	// SetPrecision sets the precision rule contexts.
	SetPrecision(INumericParameterContext)

	// SetScale sets the scale rule contexts.
	SetScale(INumericParameterContext)

	// Getter signatures
	Dec() antlr.TerminalNode
	Decimal() antlr.TerminalNode
	OAngleBracket() antlr.TerminalNode
	Comma() antlr.TerminalNode
	CAngleBracket() antlr.TerminalNode
	QMark() antlr.TerminalNode
	AllNumericParameter() []INumericParameterContext
	NumericParameter(i int) INumericParameterContext

	// IsDecimalTypeContext differentiates from other interfaces.
	IsDecimalTypeContext()
}

type DecimalTypeContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	isnull    antlr.Token
	precision INumericParameterContext
	scale     INumericParameterContext
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

func (s *DecimalTypeContext) GetIsnull() antlr.Token { return s.isnull }

func (s *DecimalTypeContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *DecimalTypeContext) GetPrecision() INumericParameterContext { return s.precision }

func (s *DecimalTypeContext) GetScale() INumericParameterContext { return s.scale }

func (s *DecimalTypeContext) SetPrecision(v INumericParameterContext) { s.precision = v }

func (s *DecimalTypeContext) SetScale(v INumericParameterContext) { s.scale = v }

func (s *DecimalTypeContext) Dec() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDec, 0)
}

func (s *DecimalTypeContext) Decimal() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDecimal, 0)
}

func (s *DecimalTypeContext) OAngleBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserOAngleBracket, 0)
}

func (s *DecimalTypeContext) Comma() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserComma, 0)
}

func (s *DecimalTypeContext) CAngleBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserCAngleBracket, 0)
}

func (s *DecimalTypeContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *DecimalTypeContext) AllNumericParameter() []INumericParameterContext {
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

func (s *DecimalTypeContext) NumericParameter(i int) INumericParameterContext {
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

func (s *DecimalTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecimalTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitDecimalType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) DecimalType() (localctx IDecimalTypeContext) {
	localctx = NewDecimalTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, FuncTestCaseParserRULE_decimalType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(553)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserDecimal || _la == FuncTestCaseParserDec) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(555)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 54, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(554)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*DecimalTypeContext).isnull = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(563)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserOAngleBracket {
		{
			p.SetState(557)
			p.Match(FuncTestCaseParserOAngleBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(558)

			var _x = p.NumericParameter()

			localctx.(*DecimalTypeContext).precision = _x
		}
		{
			p.SetState(559)
			p.Match(FuncTestCaseParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(560)

			var _x = p.NumericParameter()

			localctx.(*DecimalTypeContext).scale = _x
		}
		{
			p.SetState(561)
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

// IPrecisionTimeTypeContext is an interface to support dynamic dispatch.
type IPrecisionTimeTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// GetPrecision returns the precision rule contexts.
	GetPrecision() INumericParameterContext

	// SetPrecision sets the precision rule contexts.
	SetPrecision(INumericParameterContext)

	// Getter signatures
	OAngleBracket() antlr.TerminalNode
	CAngleBracket() antlr.TerminalNode
	PT() antlr.TerminalNode
	Precision_Time() antlr.TerminalNode
	NumericParameter() INumericParameterContext
	QMark() antlr.TerminalNode

	// IsPrecisionTimeTypeContext differentiates from other interfaces.
	IsPrecisionTimeTypeContext()
}

type PrecisionTimeTypeContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	isnull    antlr.Token
	precision INumericParameterContext
}

func NewEmptyPrecisionTimeTypeContext() *PrecisionTimeTypeContext {
	var p = new(PrecisionTimeTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_precisionTimeType
	return p
}

func InitEmptyPrecisionTimeTypeContext(p *PrecisionTimeTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_precisionTimeType
}

func (*PrecisionTimeTypeContext) IsPrecisionTimeTypeContext() {}

func NewPrecisionTimeTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrecisionTimeTypeContext {
	var p = new(PrecisionTimeTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_precisionTimeType

	return p
}

func (s *PrecisionTimeTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *PrecisionTimeTypeContext) GetIsnull() antlr.Token { return s.isnull }

func (s *PrecisionTimeTypeContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *PrecisionTimeTypeContext) GetPrecision() INumericParameterContext { return s.precision }

func (s *PrecisionTimeTypeContext) SetPrecision(v INumericParameterContext) { s.precision = v }

func (s *PrecisionTimeTypeContext) OAngleBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserOAngleBracket, 0)
}

func (s *PrecisionTimeTypeContext) CAngleBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserCAngleBracket, 0)
}

func (s *PrecisionTimeTypeContext) PT() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserPT, 0)
}

func (s *PrecisionTimeTypeContext) Precision_Time() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserPrecision_Time, 0)
}

func (s *PrecisionTimeTypeContext) NumericParameter() INumericParameterContext {
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

func (s *PrecisionTimeTypeContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *PrecisionTimeTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrecisionTimeTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrecisionTimeTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitPrecisionTimeType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) PrecisionTimeType() (localctx IPrecisionTimeTypeContext) {
	localctx = NewPrecisionTimeTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, FuncTestCaseParserRULE_precisionTimeType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(565)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserPrecision_Time || _la == FuncTestCaseParserPT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(567)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(566)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*PrecisionTimeTypeContext).isnull = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(569)
		p.Match(FuncTestCaseParserOAngleBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(570)

		var _x = p.NumericParameter()

		localctx.(*PrecisionTimeTypeContext).precision = _x
	}
	{
		p.SetState(571)
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

// IPrecisionTimestampTypeContext is an interface to support dynamic dispatch.
type IPrecisionTimestampTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// GetPrecision returns the precision rule contexts.
	GetPrecision() INumericParameterContext

	// SetPrecision sets the precision rule contexts.
	SetPrecision(INumericParameterContext)

	// Getter signatures
	OAngleBracket() antlr.TerminalNode
	CAngleBracket() antlr.TerminalNode
	PTs() antlr.TerminalNode
	Precision_Timestamp() antlr.TerminalNode
	NumericParameter() INumericParameterContext
	QMark() antlr.TerminalNode

	// IsPrecisionTimestampTypeContext differentiates from other interfaces.
	IsPrecisionTimestampTypeContext()
}

type PrecisionTimestampTypeContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	isnull    antlr.Token
	precision INumericParameterContext
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

func (s *PrecisionTimestampTypeContext) GetIsnull() antlr.Token { return s.isnull }

func (s *PrecisionTimestampTypeContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *PrecisionTimestampTypeContext) GetPrecision() INumericParameterContext { return s.precision }

func (s *PrecisionTimestampTypeContext) SetPrecision(v INumericParameterContext) { s.precision = v }

func (s *PrecisionTimestampTypeContext) OAngleBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserOAngleBracket, 0)
}

func (s *PrecisionTimestampTypeContext) CAngleBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserCAngleBracket, 0)
}

func (s *PrecisionTimestampTypeContext) PTs() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserPTs, 0)
}

func (s *PrecisionTimestampTypeContext) Precision_Timestamp() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserPrecision_Timestamp, 0)
}

func (s *PrecisionTimestampTypeContext) NumericParameter() INumericParameterContext {
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

func (s *PrecisionTimestampTypeContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *PrecisionTimestampTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrecisionTimestampTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrecisionTimestampTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitPrecisionTimestampType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) PrecisionTimestampType() (localctx IPrecisionTimestampTypeContext) {
	localctx = NewPrecisionTimestampTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, FuncTestCaseParserRULE_precisionTimestampType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(573)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserPrecision_Timestamp || _la == FuncTestCaseParserPTs) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(575)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(574)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*PrecisionTimestampTypeContext).isnull = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(577)
		p.Match(FuncTestCaseParserOAngleBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(578)

		var _x = p.NumericParameter()

		localctx.(*PrecisionTimestampTypeContext).precision = _x
	}
	{
		p.SetState(579)
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

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// GetPrecision returns the precision rule contexts.
	GetPrecision() INumericParameterContext

	// SetPrecision sets the precision rule contexts.
	SetPrecision(INumericParameterContext)

	// Getter signatures
	OAngleBracket() antlr.TerminalNode
	CAngleBracket() antlr.TerminalNode
	PTsTZ() antlr.TerminalNode
	Precision_Timestamp_TZ() antlr.TerminalNode
	NumericParameter() INumericParameterContext
	QMark() antlr.TerminalNode

	// IsPrecisionTimestampTZTypeContext differentiates from other interfaces.
	IsPrecisionTimestampTZTypeContext()
}

type PrecisionTimestampTZTypeContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	isnull    antlr.Token
	precision INumericParameterContext
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

func (s *PrecisionTimestampTZTypeContext) GetIsnull() antlr.Token { return s.isnull }

func (s *PrecisionTimestampTZTypeContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *PrecisionTimestampTZTypeContext) GetPrecision() INumericParameterContext { return s.precision }

func (s *PrecisionTimestampTZTypeContext) SetPrecision(v INumericParameterContext) { s.precision = v }

func (s *PrecisionTimestampTZTypeContext) OAngleBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserOAngleBracket, 0)
}

func (s *PrecisionTimestampTZTypeContext) CAngleBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserCAngleBracket, 0)
}

func (s *PrecisionTimestampTZTypeContext) PTsTZ() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserPTsTZ, 0)
}

func (s *PrecisionTimestampTZTypeContext) Precision_Timestamp_TZ() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserPrecision_Timestamp_TZ, 0)
}

func (s *PrecisionTimestampTZTypeContext) NumericParameter() INumericParameterContext {
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

func (s *PrecisionTimestampTZTypeContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *PrecisionTimestampTZTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrecisionTimestampTZTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrecisionTimestampTZTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitPrecisionTimestampTZType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) PrecisionTimestampTZType() (localctx IPrecisionTimestampTZTypeContext) {
	localctx = NewPrecisionTimestampTZTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, FuncTestCaseParserRULE_precisionTimestampTZType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(581)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserPrecision_Timestamp_TZ || _la == FuncTestCaseParserPTsTZ) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(583)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(582)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*PrecisionTimestampTZTypeContext).isnull = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(585)
		p.Match(FuncTestCaseParserOAngleBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(586)

		var _x = p.NumericParameter()

		localctx.(*PrecisionTimestampTZTypeContext).precision = _x
	}
	{
		p.SetState(587)
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
	p.EnterRule(localctx, 118, FuncTestCaseParserRULE_listType)
	var _la int

	localctx = NewListContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(589)
		p.Match(FuncTestCaseParserList)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(591)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(590)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*ListContext).isnull = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(593)
		p.Match(FuncTestCaseParserOAngleBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(594)

		var _x = p.DataType()

		localctx.(*ListContext).elemType = _x
	}
	{
		p.SetState(595)
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
	IntervalDayType() IIntervalDayTypeContext
	PrecisionTimeType() IPrecisionTimeTypeContext
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

func (s *ParameterizedTypeContext) IntervalDayType() IIntervalDayTypeContext {
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

func (s *ParameterizedTypeContext) PrecisionTimeType() IPrecisionTimeTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrecisionTimeTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrecisionTimeTypeContext)
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
	p.EnterRule(localctx, 120, FuncTestCaseParserRULE_parameterizedType)
	p.SetState(605)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FuncTestCaseParserFixedChar, FuncTestCaseParserFChar:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(597)
			p.FixedCharType()
		}

	case FuncTestCaseParserVarChar, FuncTestCaseParserVChar:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(598)
			p.VarCharType()
		}

	case FuncTestCaseParserFixedBinary, FuncTestCaseParserFBin:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(599)
			p.FixedBinaryType()
		}

	case FuncTestCaseParserDecimal, FuncTestCaseParserDec:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(600)
			p.DecimalType()
		}

	case FuncTestCaseParserInterval_Day, FuncTestCaseParserIDay:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(601)
			p.IntervalDayType()
		}

	case FuncTestCaseParserPrecision_Time, FuncTestCaseParserPT:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(602)
			p.PrecisionTimeType()
		}

	case FuncTestCaseParserPrecision_Timestamp, FuncTestCaseParserPTs:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(603)
			p.PrecisionTimestampType()
		}

	case FuncTestCaseParserPrecision_Timestamp_TZ, FuncTestCaseParserPTsTZ:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(604)
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
	p.EnterRule(localctx, 122, FuncTestCaseParserRULE_numericParameter)
	localctx = NewIntegerLiteralContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(607)
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
	p.EnterRule(localctx, 124, FuncTestCaseParserRULE_substraitError)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(609)
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

// IFuncOptionContext is an interface to support dynamic dispatch.
type IFuncOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OptionName() IOptionNameContext
	Colon() antlr.TerminalNode
	OptionValue() IOptionValueContext

	// IsFuncOptionContext differentiates from other interfaces.
	IsFuncOptionContext()
}

type FuncOptionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncOptionContext() *FuncOptionContext {
	var p = new(FuncOptionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_funcOption
	return p
}

func InitEmptyFuncOptionContext(p *FuncOptionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_funcOption
}

func (*FuncOptionContext) IsFuncOptionContext() {}

func NewFuncOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncOptionContext {
	var p = new(FuncOptionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_funcOption

	return p
}

func (s *FuncOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncOptionContext) OptionName() IOptionNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *FuncOptionContext) Colon() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserColon, 0)
}

func (s *FuncOptionContext) OptionValue() IOptionValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionValueContext)
}

func (s *FuncOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncOptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitFuncOption(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) FuncOption() (localctx IFuncOptionContext) {
	localctx = NewFuncOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, FuncTestCaseParserRULE_funcOption)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(611)
		p.OptionName()
	}
	{
		p.SetState(612)
		p.Match(FuncTestCaseParserColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(613)
		p.OptionValue()
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

// IOptionNameContext is an interface to support dynamic dispatch.
type IOptionNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Overflow() antlr.TerminalNode
	Rounding() antlr.TerminalNode
	NullHandling() antlr.TerminalNode
	SpacesOnly() antlr.TerminalNode
	Identifier() antlr.TerminalNode

	// IsOptionNameContext differentiates from other interfaces.
	IsOptionNameContext()
}

type OptionNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionNameContext() *OptionNameContext {
	var p = new(OptionNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_optionName
	return p
}

func InitEmptyOptionNameContext(p *OptionNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_optionName
}

func (*OptionNameContext) IsOptionNameContext() {}

func NewOptionNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionNameContext {
	var p = new(OptionNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_optionName

	return p
}

func (s *OptionNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionNameContext) Overflow() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserOverflow, 0)
}

func (s *OptionNameContext) Rounding() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserRounding, 0)
}

func (s *OptionNameContext) NullHandling() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserNullHandling, 0)
}

func (s *OptionNameContext) SpacesOnly() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserSpacesOnly, 0)
}

func (s *OptionNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserIdentifier, 0)
}

func (s *OptionNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitOptionName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) OptionName() (localctx IOptionNameContext) {
	localctx = NewOptionNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, FuncTestCaseParserRULE_optionName)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(615)
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

// IOptionValueContext is an interface to support dynamic dispatch.
type IOptionValueContext interface {
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

	// IsOptionValueContext differentiates from other interfaces.
	IsOptionValueContext()
}

type OptionValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionValueContext() *OptionValueContext {
	var p = new(OptionValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_optionValue
	return p
}

func InitEmptyOptionValueContext(p *OptionValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_optionValue
}

func (*OptionValueContext) IsOptionValueContext() {}

func NewOptionValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionValueContext {
	var p = new(OptionValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_optionValue

	return p
}

func (s *OptionValueContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionValueContext) Error() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserError, 0)
}

func (s *OptionValueContext) Saturate() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserSaturate, 0)
}

func (s *OptionValueContext) Silent() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserSilent, 0)
}

func (s *OptionValueContext) TieToEven() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserTieToEven, 0)
}

func (s *OptionValueContext) NaN() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserNaN, 0)
}

func (s *OptionValueContext) Truncate() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserTruncate, 0)
}

func (s *OptionValueContext) AcceptNulls() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserAcceptNulls, 0)
}

func (s *OptionValueContext) IgnoreNulls() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserIgnoreNulls, 0)
}

func (s *OptionValueContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserBooleanLiteral, 0)
}

func (s *OptionValueContext) NullLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserNullLiteral, 0)
}

func (s *OptionValueContext) Identifier() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserIdentifier, 0)
}

func (s *OptionValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitOptionValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) OptionValue() (localctx IOptionValueContext) {
	localctx = NewOptionValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, FuncTestCaseParserRULE_optionValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(617)
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

// IFuncOptionsContext is an interface to support dynamic dispatch.
type IFuncOptionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFuncOption() []IFuncOptionContext
	FuncOption(i int) IFuncOptionContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsFuncOptionsContext differentiates from other interfaces.
	IsFuncOptionsContext()
}

type FuncOptionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncOptionsContext() *FuncOptionsContext {
	var p = new(FuncOptionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_funcOptions
	return p
}

func InitEmptyFuncOptionsContext(p *FuncOptionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_funcOptions
}

func (*FuncOptionsContext) IsFuncOptionsContext() {}

func NewFuncOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncOptionsContext {
	var p = new(FuncOptionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_funcOptions

	return p
}

func (s *FuncOptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncOptionsContext) AllFuncOption() []IFuncOptionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncOptionContext); ok {
			len++
		}
	}

	tst := make([]IFuncOptionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncOptionContext); ok {
			tst[i] = t.(IFuncOptionContext)
			i++
		}
	}

	return tst
}

func (s *FuncOptionsContext) FuncOption(i int) IFuncOptionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncOptionContext); ok {
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

	return t.(IFuncOptionContext)
}

func (s *FuncOptionsContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(FuncTestCaseParserComma)
}

func (s *FuncOptionsContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserComma, i)
}

func (s *FuncOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncOptionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitFuncOptions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) FuncOptions() (localctx IFuncOptionsContext) {
	localctx = NewFuncOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 132, FuncTestCaseParserRULE_funcOptions)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(619)
		p.FuncOption()
	}
	p.SetState(624)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FuncTestCaseParserComma {
		{
			p.SetState(620)
			p.Match(FuncTestCaseParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(621)
			p.FuncOption()
		}

		p.SetState(626)
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
	p.EnterRule(localctx, 134, FuncTestCaseParserRULE_nonReserved)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(627)
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
	p.EnterRule(localctx, 136, FuncTestCaseParserRULE_identifier)
	p.SetState(631)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FuncTestCaseParserTruncate, FuncTestCaseParserAnd, FuncTestCaseParserOr:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(629)
			p.NonReserved()
		}

	case FuncTestCaseParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(630)
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
