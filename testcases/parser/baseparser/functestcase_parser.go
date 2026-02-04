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
		"'SUBSTRAIT_INCLUDE'", "'SUBSTRAIT_DEPENDENCY'", "", "", "'DEFINE'",
		"'<!ERROR>'", "'<!UNDEFINED>'", "'OVERFLOW'", "'ROUNDING'", "'ERROR'",
		"'SATURATE'", "'SILENT'", "'TIE_TO_EVEN'", "'NAN'", "'ACCEPT_NULLS'",
		"'IGNORE_NULLS'", "'NULL_HANDLING'", "'SPACES_ONLY'", "'TRUNCATE'",
		"", "", "", "", "", "", "", "", "'P'", "'T'", "'Y'", "'M'", "'D'", "'H'",
		"'S'", "'F'", "", "", "", "", "'null'", "", "", "", "", "'IF'", "'THEN'",
		"'ELSE'", "'FUNC'", "'BOOLEAN'", "'I8'", "'I16'", "'I32'", "'I64'",
		"'FP32'", "'FP64'", "'STRING'", "'BINARY'", "'TIMESTAMP'", "'TIMESTAMP_TZ'",
		"'DATE'", "'TIME'", "'INTERVAL_YEAR'", "'INTERVAL_DAY'", "'UUID'", "'DECIMAL'",
		"'PRECISION_TIME'", "'PRECISION_TIMESTAMP'", "'PRECISION_TIMESTAMP_TZ'",
		"'FIXEDCHAR'", "'VARCHAR'", "'FIXEDBINARY'", "'STRUCT'", "'NSTRUCT'",
		"'LIST'", "'MAP'", "'U!'", "'BOOL'", "'STR'", "'VBIN'", "'TS'", "'TSTZ'",
		"'IYEAR'", "'IDAY'", "'DEC'", "'PT'", "'PTS'", "'PTSTZ'", "'FCHAR'",
		"'VCHAR'", "'FBIN'", "'ANY'", "", "'::'", "'+'", "'-'", "'*'", "'/'",
		"'%'", "'='", "'!='", "'>='", "'<='", "'>'", "'<'", "'!'", "'('", "')'",
		"'['", "']'", "','", "':'", "'?'", "'#'", "'.'", "'AND'", "'OR'", "':='",
		"'->'",
	}
	staticData.SymbolicNames = []string{
		"", "Whitespace", "TripleHash", "SubstraitScalarTest", "SubstraitAggregateTest",
		"SubstraitInclude", "SubstraitDependency", "FormatVersion", "DescriptionLine",
		"Define", "ErrorResult", "UndefineResult", "Overflow", "Rounding", "Error",
		"Saturate", "Silent", "TieToEven", "NaN", "AcceptNulls", "IgnoreNulls",
		"NullHandling", "SpacesOnly", "Truncate", "IntegerLiteral", "DecimalLiteral",
		"FloatLiteral", "BooleanLiteral", "TimestampTzLiteral", "TimestampLiteral",
		"TimeLiteral", "DateLiteral", "PeriodPrefix", "TimePrefix", "YearPrefix",
		"MSuffix", "DaySuffix", "HourSuffix", "SecondSuffix", "FractionalSecondSuffix",
		"OAngleBracket", "CAngleBracket", "IntervalYearLiteral", "IntervalDayLiteral",
		"NullLiteral", "StringLiteral", "ColumnName", "LineComment", "BlockComment",
		"If", "Then", "Else", "Func", "Boolean", "I8", "I16", "I32", "I64",
		"FP32", "FP64", "String", "Binary", "Timestamp", "Timestamp_TZ", "Date",
		"Time", "Interval_Year", "Interval_Day", "UUID", "Decimal", "Precision_Time",
		"Precision_Timestamp", "Precision_Timestamp_TZ", "FixedChar", "VarChar",
		"FixedBinary", "Struct", "NStruct", "List", "Map", "UserDefined", "Bool",
		"Str", "VBin", "Ts", "TsTZ", "IYear", "IDay", "Dec", "PT", "PTs", "PTsTZ",
		"FChar", "VChar", "FBin", "Any", "AnyVar", "DoubleColon", "Plus", "Minus",
		"Asterisk", "ForwardSlash", "Percent", "Eq", "Ne", "Gte", "Lte", "Gt",
		"Lt", "Bang", "OParen", "CParen", "OBracket", "CBracket", "Comma", "Colon",
		"QMark", "Hash", "Dot", "And", "Or", "Assign", "Arrow", "Number", "Identifier",
		"Newline",
	}
	staticData.RuleNames = []string{
		"doc", "header", "version", "include", "dependency", "testGroupDescription",
		"testCase", "testGroup", "arguments", "result", "argument", "aggFuncTestCase",
		"aggFuncCall", "tableData", "tableRows", "dataColumn", "columnValues",
		"literal", "qualifiedAggregateFuncArgs", "aggregateFuncArgs", "qualifiedAggregateFuncArg",
		"aggregateFuncArg", "numericLiteral", "floatLiteral", "nullArg", "intArg",
		"floatArg", "decimalArg", "booleanArg", "stringArg", "dateArg", "timeArg",
		"timestampArg", "timestampTzArg", "intervalYearArg", "intervalDayArg",
		"fixedCharArg", "varCharArg", "fixedBinaryArg", "precisionTimeArg",
		"precisionTimestampArg", "precisionTimestampTZArg", "listArg", "lambdaArg",
		"literalList", "literalLambda", "lambdaParameters", "lambdaBody", "dataType",
		"scalarType", "booleanType", "stringType", "binaryType", "intType",
		"floatType", "dateType", "timeType", "timestampType", "timestampTZType",
		"intervalYearType", "intervalDayType", "fixedCharType", "varCharType",
		"fixedBinaryType", "decimalType", "precisionTimeType", "precisionTimestampType",
		"precisionTimestampTZType", "listType", "funcType", "funcParameters",
		"parameterizedType", "numericParameter", "substraitError", "funcOption",
		"optionName", "optionValue", "funcOptions", "nonReserved", "identifier",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 125, 711, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
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
		7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 7, 72, 2, 73, 7,
		73, 2, 74, 7, 74, 2, 75, 7, 75, 2, 76, 7, 76, 2, 77, 7, 77, 2, 78, 7, 78,
		2, 79, 7, 79, 1, 0, 1, 0, 4, 0, 163, 8, 0, 11, 0, 12, 0, 164, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 1, 5, 1, 172, 8, 1, 10, 1, 12, 1, 175, 9, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 188, 8,
		3, 10, 3, 12, 3, 191, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 208, 8, 6, 1, 6, 1,
		6, 1, 6, 1, 7, 3, 7, 214, 8, 7, 1, 7, 4, 7, 217, 8, 7, 11, 7, 12, 7, 218,
		1, 7, 3, 7, 222, 8, 7, 1, 7, 4, 7, 225, 8, 7, 11, 7, 12, 7, 226, 3, 7,
		229, 8, 7, 1, 8, 1, 8, 1, 8, 5, 8, 234, 8, 8, 10, 8, 12, 8, 237, 9, 8,
		1, 9, 1, 9, 3, 9, 241, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 264, 8, 10, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 3, 11, 271, 8, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1,
		12, 1, 12, 3, 12, 280, 8, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		3, 12, 288, 8, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3,
		12, 297, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 305, 8,
		13, 10, 13, 12, 13, 308, 9, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14,
		1, 14, 1, 14, 5, 14, 318, 8, 14, 10, 14, 12, 14, 321, 9, 14, 3, 14, 323,
		8, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1,
		16, 5, 16, 335, 8, 16, 10, 16, 12, 16, 338, 9, 16, 3, 16, 340, 8, 16, 1,
		16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 3, 17, 354, 8, 17, 1, 18, 1, 18, 1, 18, 5, 18, 359, 8, 18, 10, 18,
		12, 18, 362, 9, 18, 1, 19, 1, 19, 1, 19, 5, 19, 367, 8, 19, 10, 19, 12,
		19, 370, 9, 19, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 376, 8, 20, 1, 21, 1,
		21, 1, 21, 1, 21, 3, 21, 382, 8, 21, 1, 22, 1, 22, 1, 22, 3, 22, 387, 8,
		22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1,
		33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1,
		39, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41,
		1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 44, 1,
		44, 1, 44, 1, 44, 5, 44, 475, 8, 44, 10, 44, 12, 44, 478, 9, 44, 3, 44,
		480, 8, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1,
		46, 1, 46, 1, 46, 1, 46, 1, 46, 4, 46, 495, 8, 46, 11, 46, 12, 46, 496,
		1, 46, 3, 46, 500, 8, 46, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 48, 1,
		48, 3, 48, 509, 8, 48, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49,
		1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 3, 49, 523, 8, 49, 1, 49, 1, 49, 1,
		49, 3, 49, 528, 8, 49, 3, 49, 530, 8, 49, 1, 50, 1, 50, 3, 50, 534, 8,
		50, 1, 51, 1, 51, 3, 51, 538, 8, 51, 1, 52, 1, 52, 3, 52, 542, 8, 52, 1,
		53, 1, 53, 3, 53, 546, 8, 53, 1, 54, 1, 54, 3, 54, 550, 8, 54, 1, 55, 1,
		55, 3, 55, 554, 8, 55, 1, 56, 1, 56, 3, 56, 558, 8, 56, 1, 57, 1, 57, 3,
		57, 562, 8, 57, 1, 58, 1, 58, 3, 58, 566, 8, 58, 1, 59, 1, 59, 3, 59, 570,
		8, 59, 1, 60, 1, 60, 3, 60, 574, 8, 60, 1, 60, 1, 60, 1, 60, 1, 60, 3,
		60, 580, 8, 60, 1, 61, 1, 61, 3, 61, 584, 8, 61, 1, 61, 1, 61, 1, 61, 1,
		61, 1, 62, 1, 62, 3, 62, 592, 8, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 63,
		1, 63, 3, 63, 600, 8, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 64, 1, 64, 3,
		64, 608, 8, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 3, 64, 616, 8,
		64, 1, 65, 1, 65, 3, 65, 620, 8, 65, 1, 65, 1, 65, 1, 65, 1, 65, 1, 66,
		1, 66, 3, 66, 628, 8, 66, 1, 66, 1, 66, 1, 66, 1, 66, 1, 67, 1, 67, 3,
		67, 636, 8, 67, 1, 67, 1, 67, 1, 67, 1, 67, 1, 68, 1, 68, 3, 68, 644, 8,
		68, 1, 68, 1, 68, 1, 68, 1, 68, 1, 69, 1, 69, 3, 69, 652, 8, 69, 1, 69,
		1, 69, 1, 69, 1, 69, 1, 69, 1, 69, 1, 70, 1, 70, 1, 70, 1, 70, 1, 70, 5,
		70, 665, 8, 70, 10, 70, 12, 70, 668, 9, 70, 1, 70, 1, 70, 3, 70, 672, 8,
		70, 1, 71, 1, 71, 1, 71, 1, 71, 1, 71, 1, 71, 1, 71, 1, 71, 1, 71, 3, 71,
		683, 8, 71, 1, 72, 1, 72, 1, 73, 1, 73, 1, 74, 1, 74, 1, 74, 1, 74, 1,
		75, 1, 75, 1, 76, 1, 76, 1, 77, 1, 77, 1, 77, 5, 77, 700, 8, 77, 10, 77,
		12, 77, 703, 9, 77, 1, 78, 1, 78, 1, 79, 1, 79, 3, 79, 709, 8, 79, 1, 79,
		0, 0, 80, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
		34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68,
		70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104,
		106, 108, 110, 112, 114, 116, 118, 120, 122, 124, 126, 128, 130, 132, 134,
		136, 138, 140, 142, 144, 146, 148, 150, 152, 154, 156, 158, 0, 22, 1, 0,
		3, 4, 2, 0, 18, 18, 26, 26, 2, 0, 53, 53, 81, 81, 2, 0, 60, 60, 82, 82,
		2, 0, 61, 61, 83, 83, 1, 0, 54, 57, 1, 0, 58, 59, 2, 0, 62, 62, 84, 84,
		2, 0, 63, 63, 85, 85, 2, 0, 66, 66, 86, 86, 2, 0, 67, 67, 87, 87, 2, 0,
		73, 73, 92, 92, 2, 0, 74, 74, 93, 93, 2, 0, 75, 75, 94, 94, 2, 0, 69, 69,
		88, 88, 2, 0, 70, 70, 89, 89, 2, 0, 71, 71, 90, 90, 2, 0, 72, 72, 91, 91,
		1, 0, 10, 11, 3, 0, 12, 13, 21, 22, 124, 124, 5, 0, 14, 20, 23, 23, 27,
		27, 44, 44, 124, 124, 2, 0, 23, 23, 119, 120, 738, 0, 160, 1, 0, 0, 0,
		2, 168, 1, 0, 0, 0, 4, 176, 1, 0, 0, 0, 6, 181, 1, 0, 0, 0, 8, 192, 1,
		0, 0, 0, 10, 197, 1, 0, 0, 0, 12, 199, 1, 0, 0, 0, 14, 228, 1, 0, 0, 0,
		16, 230, 1, 0, 0, 0, 18, 240, 1, 0, 0, 0, 20, 263, 1, 0, 0, 0, 22, 265,
		1, 0, 0, 0, 24, 296, 1, 0, 0, 0, 26, 298, 1, 0, 0, 0, 28, 313, 1, 0, 0,
		0, 30, 326, 1, 0, 0, 0, 32, 330, 1, 0, 0, 0, 34, 353, 1, 0, 0, 0, 36, 355,
		1, 0, 0, 0, 38, 363, 1, 0, 0, 0, 40, 375, 1, 0, 0, 0, 42, 381, 1, 0, 0,
		0, 44, 386, 1, 0, 0, 0, 46, 388, 1, 0, 0, 0, 48, 390, 1, 0, 0, 0, 50, 394,
		1, 0, 0, 0, 52, 398, 1, 0, 0, 0, 54, 402, 1, 0, 0, 0, 56, 406, 1, 0, 0,
		0, 58, 410, 1, 0, 0, 0, 60, 414, 1, 0, 0, 0, 62, 418, 1, 0, 0, 0, 64, 422,
		1, 0, 0, 0, 66, 426, 1, 0, 0, 0, 68, 430, 1, 0, 0, 0, 70, 434, 1, 0, 0,
		0, 72, 438, 1, 0, 0, 0, 74, 442, 1, 0, 0, 0, 76, 446, 1, 0, 0, 0, 78, 450,
		1, 0, 0, 0, 80, 454, 1, 0, 0, 0, 82, 458, 1, 0, 0, 0, 84, 462, 1, 0, 0,
		0, 86, 466, 1, 0, 0, 0, 88, 470, 1, 0, 0, 0, 90, 483, 1, 0, 0, 0, 92, 499,
		1, 0, 0, 0, 94, 501, 1, 0, 0, 0, 96, 508, 1, 0, 0, 0, 98, 529, 1, 0, 0,
		0, 100, 531, 1, 0, 0, 0, 102, 535, 1, 0, 0, 0, 104, 539, 1, 0, 0, 0, 106,
		543, 1, 0, 0, 0, 108, 547, 1, 0, 0, 0, 110, 551, 1, 0, 0, 0, 112, 555,
		1, 0, 0, 0, 114, 559, 1, 0, 0, 0, 116, 563, 1, 0, 0, 0, 118, 567, 1, 0,
		0, 0, 120, 571, 1, 0, 0, 0, 122, 581, 1, 0, 0, 0, 124, 589, 1, 0, 0, 0,
		126, 597, 1, 0, 0, 0, 128, 605, 1, 0, 0, 0, 130, 617, 1, 0, 0, 0, 132,
		625, 1, 0, 0, 0, 134, 633, 1, 0, 0, 0, 136, 641, 1, 0, 0, 0, 138, 649,
		1, 0, 0, 0, 140, 671, 1, 0, 0, 0, 142, 682, 1, 0, 0, 0, 144, 684, 1, 0,
		0, 0, 146, 686, 1, 0, 0, 0, 148, 688, 1, 0, 0, 0, 150, 692, 1, 0, 0, 0,
		152, 694, 1, 0, 0, 0, 154, 696, 1, 0, 0, 0, 156, 704, 1, 0, 0, 0, 158,
		708, 1, 0, 0, 0, 160, 162, 3, 2, 1, 0, 161, 163, 3, 14, 7, 0, 162, 161,
		1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 164, 165, 1, 0,
		0, 0, 165, 166, 1, 0, 0, 0, 166, 167, 5, 0, 0, 1, 167, 1, 1, 0, 0, 0, 168,
		169, 3, 4, 2, 0, 169, 173, 3, 6, 3, 0, 170, 172, 3, 8, 4, 0, 171, 170,
		1, 0, 0, 0, 172, 175, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 173, 174, 1, 0,
		0, 0, 174, 3, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 176, 177, 5, 2, 0, 0, 177,
		178, 7, 0, 0, 0, 178, 179, 5, 115, 0, 0, 179, 180, 5, 7, 0, 0, 180, 5,
		1, 0, 0, 0, 181, 182, 5, 2, 0, 0, 182, 183, 5, 5, 0, 0, 183, 184, 5, 115,
		0, 0, 184, 189, 5, 45, 0, 0, 185, 186, 5, 114, 0, 0, 186, 188, 5, 45, 0,
		0, 187, 185, 1, 0, 0, 0, 188, 191, 1, 0, 0, 0, 189, 187, 1, 0, 0, 0, 189,
		190, 1, 0, 0, 0, 190, 7, 1, 0, 0, 0, 191, 189, 1, 0, 0, 0, 192, 193, 5,
		2, 0, 0, 193, 194, 5, 6, 0, 0, 194, 195, 5, 115, 0, 0, 195, 196, 5, 45,
		0, 0, 196, 9, 1, 0, 0, 0, 197, 198, 5, 8, 0, 0, 198, 11, 1, 0, 0, 0, 199,
		200, 3, 158, 79, 0, 200, 201, 5, 110, 0, 0, 201, 202, 3, 16, 8, 0, 202,
		207, 5, 111, 0, 0, 203, 204, 5, 112, 0, 0, 204, 205, 3, 154, 77, 0, 205,
		206, 5, 113, 0, 0, 206, 208, 1, 0, 0, 0, 207, 203, 1, 0, 0, 0, 207, 208,
		1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209, 210, 5, 103, 0, 0, 210, 211, 3,
		18, 9, 0, 211, 13, 1, 0, 0, 0, 212, 214, 3, 10, 5, 0, 213, 212, 1, 0, 0,
		0, 213, 214, 1, 0, 0, 0, 214, 216, 1, 0, 0, 0, 215, 217, 3, 12, 6, 0, 216,
		215, 1, 0, 0, 0, 217, 218, 1, 0, 0, 0, 218, 216, 1, 0, 0, 0, 218, 219,
		1, 0, 0, 0, 219, 229, 1, 0, 0, 0, 220, 222, 3, 10, 5, 0, 221, 220, 1, 0,
		0, 0, 221, 222, 1, 0, 0, 0, 222, 224, 1, 0, 0, 0, 223, 225, 3, 22, 11,
		0, 224, 223, 1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226, 224, 1, 0, 0, 0, 226,
		227, 1, 0, 0, 0, 227, 229, 1, 0, 0, 0, 228, 213, 1, 0, 0, 0, 228, 221,
		1, 0, 0, 0, 229, 15, 1, 0, 0, 0, 230, 235, 3, 20, 10, 0, 231, 232, 5, 114,
		0, 0, 232, 234, 3, 20, 10, 0, 233, 231, 1, 0, 0, 0, 234, 237, 1, 0, 0,
		0, 235, 233, 1, 0, 0, 0, 235, 236, 1, 0, 0, 0, 236, 17, 1, 0, 0, 0, 237,
		235, 1, 0, 0, 0, 238, 241, 3, 20, 10, 0, 239, 241, 3, 146, 73, 0, 240,
		238, 1, 0, 0, 0, 240, 239, 1, 0, 0, 0, 241, 19, 1, 0, 0, 0, 242, 264, 3,
		48, 24, 0, 243, 264, 3, 50, 25, 0, 244, 264, 3, 52, 26, 0, 245, 264, 3,
		56, 28, 0, 246, 264, 3, 58, 29, 0, 247, 264, 3, 54, 27, 0, 248, 264, 3,
		60, 30, 0, 249, 264, 3, 62, 31, 0, 250, 264, 3, 64, 32, 0, 251, 264, 3,
		66, 33, 0, 252, 264, 3, 68, 34, 0, 253, 264, 3, 70, 35, 0, 254, 264, 3,
		72, 36, 0, 255, 264, 3, 74, 37, 0, 256, 264, 3, 76, 38, 0, 257, 264, 3,
		78, 39, 0, 258, 264, 3, 80, 40, 0, 259, 264, 3, 82, 41, 0, 260, 264, 3,
		84, 42, 0, 261, 264, 3, 86, 43, 0, 262, 264, 5, 124, 0, 0, 263, 242, 1,
		0, 0, 0, 263, 243, 1, 0, 0, 0, 263, 244, 1, 0, 0, 0, 263, 245, 1, 0, 0,
		0, 263, 246, 1, 0, 0, 0, 263, 247, 1, 0, 0, 0, 263, 248, 1, 0, 0, 0, 263,
		249, 1, 0, 0, 0, 263, 250, 1, 0, 0, 0, 263, 251, 1, 0, 0, 0, 263, 252,
		1, 0, 0, 0, 263, 253, 1, 0, 0, 0, 263, 254, 1, 0, 0, 0, 263, 255, 1, 0,
		0, 0, 263, 256, 1, 0, 0, 0, 263, 257, 1, 0, 0, 0, 263, 258, 1, 0, 0, 0,
		263, 259, 1, 0, 0, 0, 263, 260, 1, 0, 0, 0, 263, 261, 1, 0, 0, 0, 263,
		262, 1, 0, 0, 0, 264, 21, 1, 0, 0, 0, 265, 270, 3, 24, 12, 0, 266, 267,
		5, 112, 0, 0, 267, 268, 3, 154, 77, 0, 268, 269, 5, 113, 0, 0, 269, 271,
		1, 0, 0, 0, 270, 266, 1, 0, 0, 0, 270, 271, 1, 0, 0, 0, 271, 272, 1, 0,
		0, 0, 272, 273, 5, 103, 0, 0, 273, 274, 3, 18, 9, 0, 274, 23, 1, 0, 0,
		0, 275, 276, 3, 26, 13, 0, 276, 277, 3, 158, 79, 0, 277, 279, 5, 110, 0,
		0, 278, 280, 3, 36, 18, 0, 279, 278, 1, 0, 0, 0, 279, 280, 1, 0, 0, 0,
		280, 281, 1, 0, 0, 0, 281, 282, 5, 111, 0, 0, 282, 297, 1, 0, 0, 0, 283,
		284, 3, 28, 14, 0, 284, 285, 3, 158, 79, 0, 285, 287, 5, 110, 0, 0, 286,
		288, 3, 38, 19, 0, 287, 286, 1, 0, 0, 0, 287, 288, 1, 0, 0, 0, 288, 289,
		1, 0, 0, 0, 289, 290, 5, 111, 0, 0, 290, 297, 1, 0, 0, 0, 291, 292, 3,
		158, 79, 0, 292, 293, 5, 110, 0, 0, 293, 294, 3, 30, 15, 0, 294, 295, 5,
		111, 0, 0, 295, 297, 1, 0, 0, 0, 296, 275, 1, 0, 0, 0, 296, 283, 1, 0,
		0, 0, 296, 291, 1, 0, 0, 0, 297, 25, 1, 0, 0, 0, 298, 299, 5, 9, 0, 0,
		299, 300, 5, 124, 0, 0, 300, 301, 5, 110, 0, 0, 301, 306, 3, 96, 48, 0,
		302, 303, 5, 114, 0, 0, 303, 305, 3, 96, 48, 0, 304, 302, 1, 0, 0, 0, 305,
		308, 1, 0, 0, 0, 306, 304, 1, 0, 0, 0, 306, 307, 1, 0, 0, 0, 307, 309,
		1, 0, 0, 0, 308, 306, 1, 0, 0, 0, 309, 310, 5, 111, 0, 0, 310, 311, 5,
		103, 0, 0, 311, 312, 3, 28, 14, 0, 312, 27, 1, 0, 0, 0, 313, 322, 5, 110,
		0, 0, 314, 319, 3, 32, 16, 0, 315, 316, 5, 114, 0, 0, 316, 318, 3, 32,
		16, 0, 317, 315, 1, 0, 0, 0, 318, 321, 1, 0, 0, 0, 319, 317, 1, 0, 0, 0,
		319, 320, 1, 0, 0, 0, 320, 323, 1, 0, 0, 0, 321, 319, 1, 0, 0, 0, 322,
		314, 1, 0, 0, 0, 322, 323, 1, 0, 0, 0, 323, 324, 1, 0, 0, 0, 324, 325,
		5, 111, 0, 0, 325, 29, 1, 0, 0, 0, 326, 327, 3, 32, 16, 0, 327, 328, 5,
		97, 0, 0, 328, 329, 3, 96, 48, 0, 329, 31, 1, 0, 0, 0, 330, 339, 5, 110,
		0, 0, 331, 336, 3, 34, 17, 0, 332, 333, 5, 114, 0, 0, 333, 335, 3, 34,
		17, 0, 334, 332, 1, 0, 0, 0, 335, 338, 1, 0, 0, 0, 336, 334, 1, 0, 0, 0,
		336, 337, 1, 0, 0, 0, 337, 340, 1, 0, 0, 0, 338, 336, 1, 0, 0, 0, 339,
		331, 1, 0, 0, 0, 339, 340, 1, 0, 0, 0, 340, 341, 1, 0, 0, 0, 341, 342,
		5, 111, 0, 0, 342, 33, 1, 0, 0, 0, 343, 354, 5, 44, 0, 0, 344, 354, 3,
		44, 22, 0, 345, 354, 5, 27, 0, 0, 346, 354, 5, 45, 0, 0, 347, 354, 5, 31,
		0, 0, 348, 354, 5, 30, 0, 0, 349, 354, 5, 29, 0, 0, 350, 354, 5, 28, 0,
		0, 351, 354, 5, 42, 0, 0, 352, 354, 5, 43, 0, 0, 353, 343, 1, 0, 0, 0,
		353, 344, 1, 0, 0, 0, 353, 345, 1, 0, 0, 0, 353, 346, 1, 0, 0, 0, 353,
		347, 1, 0, 0, 0, 353, 348, 1, 0, 0, 0, 353, 349, 1, 0, 0, 0, 353, 350,
		1, 0, 0, 0, 353, 351, 1, 0, 0, 0, 353, 352, 1, 0, 0, 0, 354, 35, 1, 0,
		0, 0, 355, 360, 3, 40, 20, 0, 356, 357, 5, 114, 0, 0, 357, 359, 3, 40,
		20, 0, 358, 356, 1, 0, 0, 0, 359, 362, 1, 0, 0, 0, 360, 358, 1, 0, 0, 0,
		360, 361, 1, 0, 0, 0, 361, 37, 1, 0, 0, 0, 362, 360, 1, 0, 0, 0, 363, 368,
		3, 42, 21, 0, 364, 365, 5, 114, 0, 0, 365, 367, 3, 42, 21, 0, 366, 364,
		1, 0, 0, 0, 367, 370, 1, 0, 0, 0, 368, 366, 1, 0, 0, 0, 368, 369, 1, 0,
		0, 0, 369, 39, 1, 0, 0, 0, 370, 368, 1, 0, 0, 0, 371, 372, 5, 124, 0, 0,
		372, 373, 5, 118, 0, 0, 373, 376, 5, 46, 0, 0, 374, 376, 3, 20, 10, 0,
		375, 371, 1, 0, 0, 0, 375, 374, 1, 0, 0, 0, 376, 41, 1, 0, 0, 0, 377, 378,
		5, 46, 0, 0, 378, 379, 5, 97, 0, 0, 379, 382, 3, 96, 48, 0, 380, 382, 3,
		20, 10, 0, 381, 377, 1, 0, 0, 0, 381, 380, 1, 0, 0, 0, 382, 43, 1, 0, 0,
		0, 383, 387, 5, 25, 0, 0, 384, 387, 5, 24, 0, 0, 385, 387, 3, 46, 23, 0,
		386, 383, 1, 0, 0, 0, 386, 384, 1, 0, 0, 0, 386, 385, 1, 0, 0, 0, 387,
		45, 1, 0, 0, 0, 388, 389, 7, 1, 0, 0, 389, 47, 1, 0, 0, 0, 390, 391, 5,
		44, 0, 0, 391, 392, 5, 97, 0, 0, 392, 393, 3, 96, 48, 0, 393, 49, 1, 0,
		0, 0, 394, 395, 5, 24, 0, 0, 395, 396, 5, 97, 0, 0, 396, 397, 3, 106, 53,
		0, 397, 51, 1, 0, 0, 0, 398, 399, 3, 44, 22, 0, 399, 400, 5, 97, 0, 0,
		400, 401, 3, 108, 54, 0, 401, 53, 1, 0, 0, 0, 402, 403, 3, 44, 22, 0, 403,
		404, 5, 97, 0, 0, 404, 405, 3, 128, 64, 0, 405, 55, 1, 0, 0, 0, 406, 407,
		5, 27, 0, 0, 407, 408, 5, 97, 0, 0, 408, 409, 3, 100, 50, 0, 409, 57, 1,
		0, 0, 0, 410, 411, 5, 45, 0, 0, 411, 412, 5, 97, 0, 0, 412, 413, 3, 102,
		51, 0, 413, 59, 1, 0, 0, 0, 414, 415, 5, 31, 0, 0, 415, 416, 5, 97, 0,
		0, 416, 417, 3, 110, 55, 0, 417, 61, 1, 0, 0, 0, 418, 419, 5, 30, 0, 0,
		419, 420, 5, 97, 0, 0, 420, 421, 3, 112, 56, 0, 421, 63, 1, 0, 0, 0, 422,
		423, 5, 29, 0, 0, 423, 424, 5, 97, 0, 0, 424, 425, 3, 114, 57, 0, 425,
		65, 1, 0, 0, 0, 426, 427, 5, 28, 0, 0, 427, 428, 5, 97, 0, 0, 428, 429,
		3, 116, 58, 0, 429, 67, 1, 0, 0, 0, 430, 431, 5, 42, 0, 0, 431, 432, 5,
		97, 0, 0, 432, 433, 3, 118, 59, 0, 433, 69, 1, 0, 0, 0, 434, 435, 5, 43,
		0, 0, 435, 436, 5, 97, 0, 0, 436, 437, 3, 120, 60, 0, 437, 71, 1, 0, 0,
		0, 438, 439, 5, 45, 0, 0, 439, 440, 5, 97, 0, 0, 440, 441, 3, 122, 61,
		0, 441, 73, 1, 0, 0, 0, 442, 443, 5, 45, 0, 0, 443, 444, 5, 97, 0, 0, 444,
		445, 3, 124, 62, 0, 445, 75, 1, 0, 0, 0, 446, 447, 5, 45, 0, 0, 447, 448,
		5, 97, 0, 0, 448, 449, 3, 126, 63, 0, 449, 77, 1, 0, 0, 0, 450, 451, 5,
		30, 0, 0, 451, 452, 5, 97, 0, 0, 452, 453, 3, 130, 65, 0, 453, 79, 1, 0,
		0, 0, 454, 455, 5, 29, 0, 0, 455, 456, 5, 97, 0, 0, 456, 457, 3, 132, 66,
		0, 457, 81, 1, 0, 0, 0, 458, 459, 5, 28, 0, 0, 459, 460, 5, 97, 0, 0, 460,
		461, 3, 134, 67, 0, 461, 83, 1, 0, 0, 0, 462, 463, 3, 88, 44, 0, 463, 464,
		5, 97, 0, 0, 464, 465, 3, 136, 68, 0, 465, 85, 1, 0, 0, 0, 466, 467, 3,
		90, 45, 0, 467, 468, 5, 97, 0, 0, 468, 469, 3, 138, 69, 0, 469, 87, 1,
		0, 0, 0, 470, 479, 5, 112, 0, 0, 471, 476, 3, 34, 17, 0, 472, 473, 5, 114,
		0, 0, 473, 475, 3, 34, 17, 0, 474, 472, 1, 0, 0, 0, 475, 478, 1, 0, 0,
		0, 476, 474, 1, 0, 0, 0, 476, 477, 1, 0, 0, 0, 477, 480, 1, 0, 0, 0, 478,
		476, 1, 0, 0, 0, 479, 471, 1, 0, 0, 0, 479, 480, 1, 0, 0, 0, 480, 481,
		1, 0, 0, 0, 481, 482, 5, 113, 0, 0, 482, 89, 1, 0, 0, 0, 483, 484, 5, 110,
		0, 0, 484, 485, 3, 92, 46, 0, 485, 486, 5, 122, 0, 0, 486, 487, 3, 94,
		47, 0, 487, 488, 5, 111, 0, 0, 488, 91, 1, 0, 0, 0, 489, 500, 5, 124, 0,
		0, 490, 491, 5, 110, 0, 0, 491, 494, 5, 124, 0, 0, 492, 493, 5, 114, 0,
		0, 493, 495, 5, 124, 0, 0, 494, 492, 1, 0, 0, 0, 495, 496, 1, 0, 0, 0,
		496, 494, 1, 0, 0, 0, 496, 497, 1, 0, 0, 0, 497, 498, 1, 0, 0, 0, 498,
		500, 5, 111, 0, 0, 499, 489, 1, 0, 0, 0, 499, 490, 1, 0, 0, 0, 500, 93,
		1, 0, 0, 0, 501, 502, 3, 158, 79, 0, 502, 503, 5, 110, 0, 0, 503, 504,
		3, 16, 8, 0, 504, 505, 5, 111, 0, 0, 505, 95, 1, 0, 0, 0, 506, 509, 3,
		98, 49, 0, 507, 509, 3, 142, 71, 0, 508, 506, 1, 0, 0, 0, 508, 507, 1,
		0, 0, 0, 509, 97, 1, 0, 0, 0, 510, 530, 3, 100, 50, 0, 511, 530, 3, 106,
		53, 0, 512, 530, 3, 108, 54, 0, 513, 530, 3, 102, 51, 0, 514, 530, 3, 104,
		52, 0, 515, 530, 3, 114, 57, 0, 516, 530, 3, 116, 58, 0, 517, 530, 3, 110,
		55, 0, 518, 530, 3, 112, 56, 0, 519, 530, 3, 118, 59, 0, 520, 522, 5, 68,
		0, 0, 521, 523, 5, 116, 0, 0, 522, 521, 1, 0, 0, 0, 522, 523, 1, 0, 0,
		0, 523, 530, 1, 0, 0, 0, 524, 525, 5, 80, 0, 0, 525, 527, 5, 124, 0, 0,
		526, 528, 5, 116, 0, 0, 527, 526, 1, 0, 0, 0, 527, 528, 1, 0, 0, 0, 528,
		530, 1, 0, 0, 0, 529, 510, 1, 0, 0, 0, 529, 511, 1, 0, 0, 0, 529, 512,
		1, 0, 0, 0, 529, 513, 1, 0, 0, 0, 529, 514, 1, 0, 0, 0, 529, 515, 1, 0,
		0, 0, 529, 516, 1, 0, 0, 0, 529, 517, 1, 0, 0, 0, 529, 518, 1, 0, 0, 0,
		529, 519, 1, 0, 0, 0, 529, 520, 1, 0, 0, 0, 529, 524, 1, 0, 0, 0, 530,
		99, 1, 0, 0, 0, 531, 533, 7, 2, 0, 0, 532, 534, 5, 116, 0, 0, 533, 532,
		1, 0, 0, 0, 533, 534, 1, 0, 0, 0, 534, 101, 1, 0, 0, 0, 535, 537, 7, 3,
		0, 0, 536, 538, 5, 116, 0, 0, 537, 536, 1, 0, 0, 0, 537, 538, 1, 0, 0,
		0, 538, 103, 1, 0, 0, 0, 539, 541, 7, 4, 0, 0, 540, 542, 5, 116, 0, 0,
		541, 540, 1, 0, 0, 0, 541, 542, 1, 0, 0, 0, 542, 105, 1, 0, 0, 0, 543,
		545, 7, 5, 0, 0, 544, 546, 5, 116, 0, 0, 545, 544, 1, 0, 0, 0, 545, 546,
		1, 0, 0, 0, 546, 107, 1, 0, 0, 0, 547, 549, 7, 6, 0, 0, 548, 550, 5, 116,
		0, 0, 549, 548, 1, 0, 0, 0, 549, 550, 1, 0, 0, 0, 550, 109, 1, 0, 0, 0,
		551, 553, 5, 64, 0, 0, 552, 554, 5, 116, 0, 0, 553, 552, 1, 0, 0, 0, 553,
		554, 1, 0, 0, 0, 554, 111, 1, 0, 0, 0, 555, 557, 5, 65, 0, 0, 556, 558,
		5, 116, 0, 0, 557, 556, 1, 0, 0, 0, 557, 558, 1, 0, 0, 0, 558, 113, 1,
		0, 0, 0, 559, 561, 7, 7, 0, 0, 560, 562, 5, 116, 0, 0, 561, 560, 1, 0,
		0, 0, 561, 562, 1, 0, 0, 0, 562, 115, 1, 0, 0, 0, 563, 565, 7, 8, 0, 0,
		564, 566, 5, 116, 0, 0, 565, 564, 1, 0, 0, 0, 565, 566, 1, 0, 0, 0, 566,
		117, 1, 0, 0, 0, 567, 569, 7, 9, 0, 0, 568, 570, 5, 116, 0, 0, 569, 568,
		1, 0, 0, 0, 569, 570, 1, 0, 0, 0, 570, 119, 1, 0, 0, 0, 571, 573, 7, 10,
		0, 0, 572, 574, 5, 116, 0, 0, 573, 572, 1, 0, 0, 0, 573, 574, 1, 0, 0,
		0, 574, 579, 1, 0, 0, 0, 575, 576, 5, 40, 0, 0, 576, 577, 3, 144, 72, 0,
		577, 578, 5, 41, 0, 0, 578, 580, 1, 0, 0, 0, 579, 575, 1, 0, 0, 0, 579,
		580, 1, 0, 0, 0, 580, 121, 1, 0, 0, 0, 581, 583, 7, 11, 0, 0, 582, 584,
		5, 116, 0, 0, 583, 582, 1, 0, 0, 0, 583, 584, 1, 0, 0, 0, 584, 585, 1,
		0, 0, 0, 585, 586, 5, 40, 0, 0, 586, 587, 3, 144, 72, 0, 587, 588, 5, 41,
		0, 0, 588, 123, 1, 0, 0, 0, 589, 591, 7, 12, 0, 0, 590, 592, 5, 116, 0,
		0, 591, 590, 1, 0, 0, 0, 591, 592, 1, 0, 0, 0, 592, 593, 1, 0, 0, 0, 593,
		594, 5, 40, 0, 0, 594, 595, 3, 144, 72, 0, 595, 596, 5, 41, 0, 0, 596,
		125, 1, 0, 0, 0, 597, 599, 7, 13, 0, 0, 598, 600, 5, 116, 0, 0, 599, 598,
		1, 0, 0, 0, 599, 600, 1, 0, 0, 0, 600, 601, 1, 0, 0, 0, 601, 602, 5, 40,
		0, 0, 602, 603, 3, 144, 72, 0, 603, 604, 5, 41, 0, 0, 604, 127, 1, 0, 0,
		0, 605, 607, 7, 14, 0, 0, 606, 608, 5, 116, 0, 0, 607, 606, 1, 0, 0, 0,
		607, 608, 1, 0, 0, 0, 608, 615, 1, 0, 0, 0, 609, 610, 5, 40, 0, 0, 610,
		611, 3, 144, 72, 0, 611, 612, 5, 114, 0, 0, 612, 613, 3, 144, 72, 0, 613,
		614, 5, 41, 0, 0, 614, 616, 1, 0, 0, 0, 615, 609, 1, 0, 0, 0, 615, 616,
		1, 0, 0, 0, 616, 129, 1, 0, 0, 0, 617, 619, 7, 15, 0, 0, 618, 620, 5, 116,
		0, 0, 619, 618, 1, 0, 0, 0, 619, 620, 1, 0, 0, 0, 620, 621, 1, 0, 0, 0,
		621, 622, 5, 40, 0, 0, 622, 623, 3, 144, 72, 0, 623, 624, 5, 41, 0, 0,
		624, 131, 1, 0, 0, 0, 625, 627, 7, 16, 0, 0, 626, 628, 5, 116, 0, 0, 627,
		626, 1, 0, 0, 0, 627, 628, 1, 0, 0, 0, 628, 629, 1, 0, 0, 0, 629, 630,
		5, 40, 0, 0, 630, 631, 3, 144, 72, 0, 631, 632, 5, 41, 0, 0, 632, 133,
		1, 0, 0, 0, 633, 635, 7, 17, 0, 0, 634, 636, 5, 116, 0, 0, 635, 634, 1,
		0, 0, 0, 635, 636, 1, 0, 0, 0, 636, 637, 1, 0, 0, 0, 637, 638, 5, 40, 0,
		0, 638, 639, 3, 144, 72, 0, 639, 640, 5, 41, 0, 0, 640, 135, 1, 0, 0, 0,
		641, 643, 5, 78, 0, 0, 642, 644, 5, 116, 0, 0, 643, 642, 1, 0, 0, 0, 643,
		644, 1, 0, 0, 0, 644, 645, 1, 0, 0, 0, 645, 646, 5, 40, 0, 0, 646, 647,
		3, 96, 48, 0, 647, 648, 5, 41, 0, 0, 648, 137, 1, 0, 0, 0, 649, 651, 5,
		52, 0, 0, 650, 652, 5, 116, 0, 0, 651, 650, 1, 0, 0, 0, 651, 652, 1, 0,
		0, 0, 652, 653, 1, 0, 0, 0, 653, 654, 5, 40, 0, 0, 654, 655, 3, 140, 70,
		0, 655, 656, 5, 122, 0, 0, 656, 657, 3, 96, 48, 0, 657, 658, 5, 41, 0,
		0, 658, 139, 1, 0, 0, 0, 659, 672, 3, 96, 48, 0, 660, 661, 5, 110, 0, 0,
		661, 666, 3, 96, 48, 0, 662, 663, 5, 114, 0, 0, 663, 665, 3, 96, 48, 0,
		664, 662, 1, 0, 0, 0, 665, 668, 1, 0, 0, 0, 666, 664, 1, 0, 0, 0, 666,
		667, 1, 0, 0, 0, 667, 669, 1, 0, 0, 0, 668, 666, 1, 0, 0, 0, 669, 670,
		5, 111, 0, 0, 670, 672, 1, 0, 0, 0, 671, 659, 1, 0, 0, 0, 671, 660, 1,
		0, 0, 0, 672, 141, 1, 0, 0, 0, 673, 683, 3, 122, 61, 0, 674, 683, 3, 124,
		62, 0, 675, 683, 3, 126, 63, 0, 676, 683, 3, 128, 64, 0, 677, 683, 3, 120,
		60, 0, 678, 683, 3, 130, 65, 0, 679, 683, 3, 132, 66, 0, 680, 683, 3, 134,
		67, 0, 681, 683, 3, 138, 69, 0, 682, 673, 1, 0, 0, 0, 682, 674, 1, 0, 0,
		0, 682, 675, 1, 0, 0, 0, 682, 676, 1, 0, 0, 0, 682, 677, 1, 0, 0, 0, 682,
		678, 1, 0, 0, 0, 682, 679, 1, 0, 0, 0, 682, 680, 1, 0, 0, 0, 682, 681,
		1, 0, 0, 0, 683, 143, 1, 0, 0, 0, 684, 685, 5, 24, 0, 0, 685, 145, 1, 0,
		0, 0, 686, 687, 7, 18, 0, 0, 687, 147, 1, 0, 0, 0, 688, 689, 3, 150, 75,
		0, 689, 690, 5, 115, 0, 0, 690, 691, 3, 152, 76, 0, 691, 149, 1, 0, 0,
		0, 692, 693, 7, 19, 0, 0, 693, 151, 1, 0, 0, 0, 694, 695, 7, 20, 0, 0,
		695, 153, 1, 0, 0, 0, 696, 701, 3, 148, 74, 0, 697, 698, 5, 114, 0, 0,
		698, 700, 3, 148, 74, 0, 699, 697, 1, 0, 0, 0, 700, 703, 1, 0, 0, 0, 701,
		699, 1, 0, 0, 0, 701, 702, 1, 0, 0, 0, 702, 155, 1, 0, 0, 0, 703, 701,
		1, 0, 0, 0, 704, 705, 7, 21, 0, 0, 705, 157, 1, 0, 0, 0, 706, 709, 3, 156,
		78, 0, 707, 709, 5, 124, 0, 0, 708, 706, 1, 0, 0, 0, 708, 707, 1, 0, 0,
		0, 709, 159, 1, 0, 0, 0, 62, 164, 173, 189, 207, 213, 218, 221, 226, 228,
		235, 240, 263, 270, 279, 287, 296, 306, 319, 322, 336, 339, 353, 360, 368,
		375, 381, 386, 476, 479, 496, 499, 508, 522, 527, 529, 533, 537, 541, 545,
		549, 553, 557, 561, 565, 569, 573, 579, 583, 591, 599, 607, 615, 619, 627,
		635, 643, 651, 666, 671, 682, 701, 708,
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
	FuncTestCaseParserSubstraitDependency    = 6
	FuncTestCaseParserFormatVersion          = 7
	FuncTestCaseParserDescriptionLine        = 8
	FuncTestCaseParserDefine                 = 9
	FuncTestCaseParserErrorResult            = 10
	FuncTestCaseParserUndefineResult         = 11
	FuncTestCaseParserOverflow               = 12
	FuncTestCaseParserRounding               = 13
	FuncTestCaseParserError                  = 14
	FuncTestCaseParserSaturate               = 15
	FuncTestCaseParserSilent                 = 16
	FuncTestCaseParserTieToEven              = 17
	FuncTestCaseParserNaN                    = 18
	FuncTestCaseParserAcceptNulls            = 19
	FuncTestCaseParserIgnoreNulls            = 20
	FuncTestCaseParserNullHandling           = 21
	FuncTestCaseParserSpacesOnly             = 22
	FuncTestCaseParserTruncate               = 23
	FuncTestCaseParserIntegerLiteral         = 24
	FuncTestCaseParserDecimalLiteral         = 25
	FuncTestCaseParserFloatLiteral           = 26
	FuncTestCaseParserBooleanLiteral         = 27
	FuncTestCaseParserTimestampTzLiteral     = 28
	FuncTestCaseParserTimestampLiteral       = 29
	FuncTestCaseParserTimeLiteral            = 30
	FuncTestCaseParserDateLiteral            = 31
	FuncTestCaseParserPeriodPrefix           = 32
	FuncTestCaseParserTimePrefix             = 33
	FuncTestCaseParserYearPrefix             = 34
	FuncTestCaseParserMSuffix                = 35
	FuncTestCaseParserDaySuffix              = 36
	FuncTestCaseParserHourSuffix             = 37
	FuncTestCaseParserSecondSuffix           = 38
	FuncTestCaseParserFractionalSecondSuffix = 39
	FuncTestCaseParserOAngleBracket          = 40
	FuncTestCaseParserCAngleBracket          = 41
	FuncTestCaseParserIntervalYearLiteral    = 42
	FuncTestCaseParserIntervalDayLiteral     = 43
	FuncTestCaseParserNullLiteral            = 44
	FuncTestCaseParserStringLiteral          = 45
	FuncTestCaseParserColumnName             = 46
	FuncTestCaseParserLineComment            = 47
	FuncTestCaseParserBlockComment           = 48
	FuncTestCaseParserIf                     = 49
	FuncTestCaseParserThen                   = 50
	FuncTestCaseParserElse                   = 51
	FuncTestCaseParserFunc                   = 52
	FuncTestCaseParserBoolean                = 53
	FuncTestCaseParserI8                     = 54
	FuncTestCaseParserI16                    = 55
	FuncTestCaseParserI32                    = 56
	FuncTestCaseParserI64                    = 57
	FuncTestCaseParserFP32                   = 58
	FuncTestCaseParserFP64                   = 59
	FuncTestCaseParserString_                = 60
	FuncTestCaseParserBinary                 = 61
	FuncTestCaseParserTimestamp              = 62
	FuncTestCaseParserTimestamp_TZ           = 63
	FuncTestCaseParserDate                   = 64
	FuncTestCaseParserTime                   = 65
	FuncTestCaseParserInterval_Year          = 66
	FuncTestCaseParserInterval_Day           = 67
	FuncTestCaseParserUUID                   = 68
	FuncTestCaseParserDecimal                = 69
	FuncTestCaseParserPrecision_Time         = 70
	FuncTestCaseParserPrecision_Timestamp    = 71
	FuncTestCaseParserPrecision_Timestamp_TZ = 72
	FuncTestCaseParserFixedChar              = 73
	FuncTestCaseParserVarChar                = 74
	FuncTestCaseParserFixedBinary            = 75
	FuncTestCaseParserStruct                 = 76
	FuncTestCaseParserNStruct                = 77
	FuncTestCaseParserList                   = 78
	FuncTestCaseParserMap                    = 79
	FuncTestCaseParserUserDefined            = 80
	FuncTestCaseParserBool                   = 81
	FuncTestCaseParserStr                    = 82
	FuncTestCaseParserVBin                   = 83
	FuncTestCaseParserTs                     = 84
	FuncTestCaseParserTsTZ                   = 85
	FuncTestCaseParserIYear                  = 86
	FuncTestCaseParserIDay                   = 87
	FuncTestCaseParserDec                    = 88
	FuncTestCaseParserPT                     = 89
	FuncTestCaseParserPTs                    = 90
	FuncTestCaseParserPTsTZ                  = 91
	FuncTestCaseParserFChar                  = 92
	FuncTestCaseParserVChar                  = 93
	FuncTestCaseParserFBin                   = 94
	FuncTestCaseParserAny                    = 95
	FuncTestCaseParserAnyVar                 = 96
	FuncTestCaseParserDoubleColon            = 97
	FuncTestCaseParserPlus                   = 98
	FuncTestCaseParserMinus                  = 99
	FuncTestCaseParserAsterisk               = 100
	FuncTestCaseParserForwardSlash           = 101
	FuncTestCaseParserPercent                = 102
	FuncTestCaseParserEq                     = 103
	FuncTestCaseParserNe                     = 104
	FuncTestCaseParserGte                    = 105
	FuncTestCaseParserLte                    = 106
	FuncTestCaseParserGt                     = 107
	FuncTestCaseParserLt                     = 108
	FuncTestCaseParserBang                   = 109
	FuncTestCaseParserOParen                 = 110
	FuncTestCaseParserCParen                 = 111
	FuncTestCaseParserOBracket               = 112
	FuncTestCaseParserCBracket               = 113
	FuncTestCaseParserComma                  = 114
	FuncTestCaseParserColon                  = 115
	FuncTestCaseParserQMark                  = 116
	FuncTestCaseParserHash                   = 117
	FuncTestCaseParserDot                    = 118
	FuncTestCaseParserAnd                    = 119
	FuncTestCaseParserOr                     = 120
	FuncTestCaseParserAssign                 = 121
	FuncTestCaseParserArrow                  = 122
	FuncTestCaseParserNumber                 = 123
	FuncTestCaseParserIdentifier             = 124
	FuncTestCaseParserNewline                = 125
)

// FuncTestCaseParser rules.
const (
	FuncTestCaseParserRULE_doc                        = 0
	FuncTestCaseParserRULE_header                     = 1
	FuncTestCaseParserRULE_version                    = 2
	FuncTestCaseParserRULE_include                    = 3
	FuncTestCaseParserRULE_dependency                 = 4
	FuncTestCaseParserRULE_testGroupDescription       = 5
	FuncTestCaseParserRULE_testCase                   = 6
	FuncTestCaseParserRULE_testGroup                  = 7
	FuncTestCaseParserRULE_arguments                  = 8
	FuncTestCaseParserRULE_result                     = 9
	FuncTestCaseParserRULE_argument                   = 10
	FuncTestCaseParserRULE_aggFuncTestCase            = 11
	FuncTestCaseParserRULE_aggFuncCall                = 12
	FuncTestCaseParserRULE_tableData                  = 13
	FuncTestCaseParserRULE_tableRows                  = 14
	FuncTestCaseParserRULE_dataColumn                 = 15
	FuncTestCaseParserRULE_columnValues               = 16
	FuncTestCaseParserRULE_literal                    = 17
	FuncTestCaseParserRULE_qualifiedAggregateFuncArgs = 18
	FuncTestCaseParserRULE_aggregateFuncArgs          = 19
	FuncTestCaseParserRULE_qualifiedAggregateFuncArg  = 20
	FuncTestCaseParserRULE_aggregateFuncArg           = 21
	FuncTestCaseParserRULE_numericLiteral             = 22
	FuncTestCaseParserRULE_floatLiteral               = 23
	FuncTestCaseParserRULE_nullArg                    = 24
	FuncTestCaseParserRULE_intArg                     = 25
	FuncTestCaseParserRULE_floatArg                   = 26
	FuncTestCaseParserRULE_decimalArg                 = 27
	FuncTestCaseParserRULE_booleanArg                 = 28
	FuncTestCaseParserRULE_stringArg                  = 29
	FuncTestCaseParserRULE_dateArg                    = 30
	FuncTestCaseParserRULE_timeArg                    = 31
	FuncTestCaseParserRULE_timestampArg               = 32
	FuncTestCaseParserRULE_timestampTzArg             = 33
	FuncTestCaseParserRULE_intervalYearArg            = 34
	FuncTestCaseParserRULE_intervalDayArg             = 35
	FuncTestCaseParserRULE_fixedCharArg               = 36
	FuncTestCaseParserRULE_varCharArg                 = 37
	FuncTestCaseParserRULE_fixedBinaryArg             = 38
	FuncTestCaseParserRULE_precisionTimeArg           = 39
	FuncTestCaseParserRULE_precisionTimestampArg      = 40
	FuncTestCaseParserRULE_precisionTimestampTZArg    = 41
	FuncTestCaseParserRULE_listArg                    = 42
	FuncTestCaseParserRULE_lambdaArg                  = 43
	FuncTestCaseParserRULE_literalList                = 44
	FuncTestCaseParserRULE_literalLambda              = 45
	FuncTestCaseParserRULE_lambdaParameters           = 46
	FuncTestCaseParserRULE_lambdaBody                 = 47
	FuncTestCaseParserRULE_dataType                   = 48
	FuncTestCaseParserRULE_scalarType                 = 49
	FuncTestCaseParserRULE_booleanType                = 50
	FuncTestCaseParserRULE_stringType                 = 51
	FuncTestCaseParserRULE_binaryType                 = 52
	FuncTestCaseParserRULE_intType                    = 53
	FuncTestCaseParserRULE_floatType                  = 54
	FuncTestCaseParserRULE_dateType                   = 55
	FuncTestCaseParserRULE_timeType                   = 56
	FuncTestCaseParserRULE_timestampType              = 57
	FuncTestCaseParserRULE_timestampTZType            = 58
	FuncTestCaseParserRULE_intervalYearType           = 59
	FuncTestCaseParserRULE_intervalDayType            = 60
	FuncTestCaseParserRULE_fixedCharType              = 61
	FuncTestCaseParserRULE_varCharType                = 62
	FuncTestCaseParserRULE_fixedBinaryType            = 63
	FuncTestCaseParserRULE_decimalType                = 64
	FuncTestCaseParserRULE_precisionTimeType          = 65
	FuncTestCaseParserRULE_precisionTimestampType     = 66
	FuncTestCaseParserRULE_precisionTimestampTZType   = 67
	FuncTestCaseParserRULE_listType                   = 68
	FuncTestCaseParserRULE_funcType                   = 69
	FuncTestCaseParserRULE_funcParameters             = 70
	FuncTestCaseParserRULE_parameterizedType          = 71
	FuncTestCaseParserRULE_numericParameter           = 72
	FuncTestCaseParserRULE_substraitError             = 73
	FuncTestCaseParserRULE_funcOption                 = 74
	FuncTestCaseParserRULE_optionName                 = 75
	FuncTestCaseParserRULE_optionValue                = 76
	FuncTestCaseParserRULE_funcOptions                = 77
	FuncTestCaseParserRULE_nonReserved                = 78
	FuncTestCaseParserRULE_identifier                 = 79
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
		p.SetState(160)
		p.Header()
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8389376) != 0) || ((int64((_la-110)) & ^0x3f) == 0 && ((int64(1)<<(_la-110))&17921) != 0) {
		{
			p.SetState(161)
			p.TestGroup()
		}

		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(166)
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
	AllDependency() []IDependencyContext
	Dependency(i int) IDependencyContext

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

func (s *HeaderContext) AllDependency() []IDependencyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDependencyContext); ok {
			len++
		}
	}

	tst := make([]IDependencyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDependencyContext); ok {
			tst[i] = t.(IDependencyContext)
			i++
		}
	}

	return tst
}

func (s *HeaderContext) Dependency(i int) IDependencyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDependencyContext); ok {
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

	return t.(IDependencyContext)
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
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Version()
	}
	{
		p.SetState(169)
		p.Include()
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FuncTestCaseParserTripleHash {
		{
			p.SetState(170)
			p.Dependency()
		}

		p.SetState(175)
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
		p.SetState(176)
		p.Match(FuncTestCaseParserTripleHash)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(177)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserSubstraitScalarTest || _la == FuncTestCaseParserSubstraitAggregateTest) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(178)
		p.Match(FuncTestCaseParserColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(179)
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
		p.SetState(181)
		p.Match(FuncTestCaseParserTripleHash)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)
		p.Match(FuncTestCaseParserSubstraitInclude)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(183)
		p.Match(FuncTestCaseParserColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(184)
		p.Match(FuncTestCaseParserStringLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FuncTestCaseParserComma {
		{
			p.SetState(185)
			p.Match(FuncTestCaseParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)
			p.Match(FuncTestCaseParserStringLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(191)
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

// IDependencyContext is an interface to support dynamic dispatch.
type IDependencyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TripleHash() antlr.TerminalNode
	SubstraitDependency() antlr.TerminalNode
	Colon() antlr.TerminalNode
	StringLiteral() antlr.TerminalNode

	// IsDependencyContext differentiates from other interfaces.
	IsDependencyContext()
}

type DependencyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDependencyContext() *DependencyContext {
	var p = new(DependencyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_dependency
	return p
}

func InitEmptyDependencyContext(p *DependencyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_dependency
}

func (*DependencyContext) IsDependencyContext() {}

func NewDependencyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DependencyContext {
	var p = new(DependencyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_dependency

	return p
}

func (s *DependencyContext) GetParser() antlr.Parser { return s.parser }

func (s *DependencyContext) TripleHash() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserTripleHash, 0)
}

func (s *DependencyContext) SubstraitDependency() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserSubstraitDependency, 0)
}

func (s *DependencyContext) Colon() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserColon, 0)
}

func (s *DependencyContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserStringLiteral, 0)
}

func (s *DependencyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DependencyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DependencyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitDependency(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) Dependency() (localctx IDependencyContext) {
	localctx = NewDependencyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FuncTestCaseParserRULE_dependency)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Match(FuncTestCaseParserTripleHash)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(193)
		p.Match(FuncTestCaseParserSubstraitDependency)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(194)
		p.Match(FuncTestCaseParserColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(195)
		p.Match(FuncTestCaseParserStringLiteral)
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
	p.EnterRule(localctx, 10, FuncTestCaseParserRULE_testGroupDescription)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
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
	p.EnterRule(localctx, 12, FuncTestCaseParserRULE_testCase)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)

		var _x = p.Identifier()

		localctx.(*TestCaseContext).functionName = _x
	}
	{
		p.SetState(200)
		p.Match(FuncTestCaseParserOParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(201)
		p.Arguments()
	}
	{
		p.SetState(202)
		p.Match(FuncTestCaseParserCParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserOBracket {
		{
			p.SetState(203)
			p.Match(FuncTestCaseParserOBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(204)
			p.FuncOptions()
		}
		{
			p.SetState(205)
			p.Match(FuncTestCaseParserCBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(209)
		p.Match(FuncTestCaseParserEq)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(210)
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
	p.EnterRule(localctx, 14, FuncTestCaseParserRULE_testGroup)
	var _la int

	var _alt int

	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		localctx = NewScalarFuncTestGroupContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FuncTestCaseParserDescriptionLine {
			{
				p.SetState(212)
				p.TestGroupDescription()
			}

		}
		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(215)
					p.TestCase()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(218)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case 2:
		localctx = NewAggregateFuncTestGroupContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FuncTestCaseParserDescriptionLine {
			{
				p.SetState(220)
				p.TestGroupDescription()
			}

		}
		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(223)
					p.AggFuncTestCase()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(226)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 16, FuncTestCaseParserRULE_arguments)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.Argument()
	}
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FuncTestCaseParserComma {
		{
			p.SetState(231)
			p.Match(FuncTestCaseParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(232)
			p.Argument()
		}

		p.SetState(237)
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
	p.EnterRule(localctx, 18, FuncTestCaseParserRULE_result)
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FuncTestCaseParserNaN, FuncTestCaseParserIntegerLiteral, FuncTestCaseParserDecimalLiteral, FuncTestCaseParserFloatLiteral, FuncTestCaseParserBooleanLiteral, FuncTestCaseParserTimestampTzLiteral, FuncTestCaseParserTimestampLiteral, FuncTestCaseParserTimeLiteral, FuncTestCaseParserDateLiteral, FuncTestCaseParserIntervalYearLiteral, FuncTestCaseParserIntervalDayLiteral, FuncTestCaseParserNullLiteral, FuncTestCaseParserStringLiteral, FuncTestCaseParserOParen, FuncTestCaseParserOBracket, FuncTestCaseParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(238)
			p.Argument()
		}

	case FuncTestCaseParserErrorResult, FuncTestCaseParserUndefineResult:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(239)
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
	LambdaArg() ILambdaArgContext
	Identifier() antlr.TerminalNode

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

func (s *ArgumentContext) LambdaArg() ILambdaArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILambdaArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILambdaArgContext)
}

func (s *ArgumentContext) Identifier() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserIdentifier, 0)
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
	p.EnterRule(localctx, 20, FuncTestCaseParserRULE_argument)
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(242)
			p.NullArg()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(243)
			p.IntArg()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(244)
			p.FloatArg()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(245)
			p.BooleanArg()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(246)
			p.StringArg()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(247)
			p.DecimalArg()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(248)
			p.DateArg()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(249)
			p.TimeArg()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(250)
			p.TimestampArg()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(251)
			p.TimestampTzArg()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(252)
			p.IntervalYearArg()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(253)
			p.IntervalDayArg()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(254)
			p.FixedCharArg()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(255)
			p.VarCharArg()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(256)
			p.FixedBinaryArg()
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(257)
			p.PrecisionTimeArg()
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(258)
			p.PrecisionTimestampArg()
		}

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(259)
			p.PrecisionTimestampTZArg()
		}

	case 19:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(260)
			p.ListArg()
		}

	case 20:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(261)
			p.LambdaArg()
		}

	case 21:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(262)
			p.Match(FuncTestCaseParserIdentifier)
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
	p.EnterRule(localctx, 22, FuncTestCaseParserRULE_aggFuncTestCase)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(265)
		p.AggFuncCall()
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserOBracket {
		{
			p.SetState(266)
			p.Match(FuncTestCaseParserOBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(267)
			p.FuncOptions()
		}
		{
			p.SetState(268)
			p.Match(FuncTestCaseParserCBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(272)
		p.Match(FuncTestCaseParserEq)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(273)
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
	p.EnterRule(localctx, 24, FuncTestCaseParserRULE_aggFuncCall)
	var _la int

	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FuncTestCaseParserDefine:
		localctx = NewMultiArgAggregateFuncCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(275)
			p.TableData()
		}
		{
			p.SetState(276)

			var _x = p.Identifier()

			localctx.(*MultiArgAggregateFuncCallContext).funcName = _x
		}
		{
			p.SetState(277)
			p.Match(FuncTestCaseParserOParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(279)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&65974976118784) != 0) || ((int64((_la-110)) & ^0x3f) == 0 && ((int64(1)<<(_la-110))&16389) != 0) {
			{
				p.SetState(278)
				p.QualifiedAggregateFuncArgs()
			}

		}
		{
			p.SetState(281)
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
			p.SetState(283)
			p.TableRows()
		}
		{
			p.SetState(284)

			var _x = p.Identifier()

			localctx.(*CompactAggregateFuncCallContext).functName = _x
		}
		{
			p.SetState(285)
			p.Match(FuncTestCaseParserOParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(287)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&136343720296448) != 0) || ((int64((_la-110)) & ^0x3f) == 0 && ((int64(1)<<(_la-110))&16389) != 0) {
			{
				p.SetState(286)
				p.AggregateFuncArgs()
			}

		}
		{
			p.SetState(289)
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
			p.SetState(291)

			var _x = p.Identifier()

			localctx.(*SingleArgAggregateFuncCallContext).functName = _x
		}
		{
			p.SetState(292)
			p.Match(FuncTestCaseParserOParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(293)
			p.DataColumn()
		}
		{
			p.SetState(294)
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
	p.EnterRule(localctx, 26, FuncTestCaseParserRULE_tableData)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(298)
		p.Match(FuncTestCaseParserDefine)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(299)

		var _m = p.Match(FuncTestCaseParserIdentifier)

		localctx.(*TableDataContext).tableName = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(300)
		p.Match(FuncTestCaseParserOParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(301)
		p.DataType()
	}
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FuncTestCaseParserComma {
		{
			p.SetState(302)
			p.Match(FuncTestCaseParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(303)
			p.DataType()
		}

		p.SetState(308)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(309)
		p.Match(FuncTestCaseParserCParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(310)
		p.Match(FuncTestCaseParserEq)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(311)
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
	p.EnterRule(localctx, 28, FuncTestCaseParserRULE_tableRows)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(313)
		p.Match(FuncTestCaseParserOParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(322)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserOParen {
		{
			p.SetState(314)
			p.ColumnValues()
		}
		p.SetState(319)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FuncTestCaseParserComma {
			{
				p.SetState(315)
				p.Match(FuncTestCaseParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(316)
				p.ColumnValues()
			}

			p.SetState(321)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(324)
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
	p.EnterRule(localctx, 30, FuncTestCaseParserRULE_dataColumn)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(326)
		p.ColumnValues()
	}
	{
		p.SetState(327)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(328)
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
	p.EnterRule(localctx, 32, FuncTestCaseParserRULE_columnValues)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(330)
		p.Match(FuncTestCaseParserOParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(339)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&65974976118784) != 0 {
		{
			p.SetState(331)
			p.Literal()
		}
		p.SetState(336)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FuncTestCaseParserComma {
			{
				p.SetState(332)
				p.Match(FuncTestCaseParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(333)
				p.Literal()
			}

			p.SetState(338)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(341)
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
	p.EnterRule(localctx, 34, FuncTestCaseParserRULE_literal)
	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FuncTestCaseParserNullLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(343)
			p.Match(FuncTestCaseParserNullLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserNaN, FuncTestCaseParserIntegerLiteral, FuncTestCaseParserDecimalLiteral, FuncTestCaseParserFloatLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(344)
			p.NumericLiteral()
		}

	case FuncTestCaseParserBooleanLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(345)
			p.Match(FuncTestCaseParserBooleanLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserStringLiteral:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(346)
			p.Match(FuncTestCaseParserStringLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserDateLiteral:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(347)
			p.Match(FuncTestCaseParserDateLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserTimeLiteral:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(348)
			p.Match(FuncTestCaseParserTimeLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserTimestampLiteral:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(349)
			p.Match(FuncTestCaseParserTimestampLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserTimestampTzLiteral:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(350)
			p.Match(FuncTestCaseParserTimestampTzLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserIntervalYearLiteral:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(351)
			p.Match(FuncTestCaseParserIntervalYearLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserIntervalDayLiteral:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(352)
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
	p.EnterRule(localctx, 36, FuncTestCaseParserRULE_qualifiedAggregateFuncArgs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(355)
		p.QualifiedAggregateFuncArg()
	}
	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FuncTestCaseParserComma {
		{
			p.SetState(356)
			p.Match(FuncTestCaseParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(357)
			p.QualifiedAggregateFuncArg()
		}

		p.SetState(362)
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
	p.EnterRule(localctx, 38, FuncTestCaseParserRULE_aggregateFuncArgs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(363)
		p.AggregateFuncArg()
	}
	p.SetState(368)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FuncTestCaseParserComma {
		{
			p.SetState(364)
			p.Match(FuncTestCaseParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(365)
			p.AggregateFuncArg()
		}

		p.SetState(370)
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
	p.EnterRule(localctx, 40, FuncTestCaseParserRULE_qualifiedAggregateFuncArg)
	p.SetState(375)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(371)

			var _m = p.Match(FuncTestCaseParserIdentifier)

			localctx.(*QualifiedAggregateFuncArgContext).tableName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(372)
			p.Match(FuncTestCaseParserDot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(373)
			p.Match(FuncTestCaseParserColumnName)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(374)
			p.Argument()
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
	p.EnterRule(localctx, 42, FuncTestCaseParserRULE_aggregateFuncArg)
	p.SetState(381)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FuncTestCaseParserColumnName:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(377)
			p.Match(FuncTestCaseParserColumnName)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(378)
			p.Match(FuncTestCaseParserDoubleColon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(379)
			p.DataType()
		}

	case FuncTestCaseParserNaN, FuncTestCaseParserIntegerLiteral, FuncTestCaseParserDecimalLiteral, FuncTestCaseParserFloatLiteral, FuncTestCaseParserBooleanLiteral, FuncTestCaseParserTimestampTzLiteral, FuncTestCaseParserTimestampLiteral, FuncTestCaseParserTimeLiteral, FuncTestCaseParserDateLiteral, FuncTestCaseParserIntervalYearLiteral, FuncTestCaseParserIntervalDayLiteral, FuncTestCaseParserNullLiteral, FuncTestCaseParserStringLiteral, FuncTestCaseParserOParen, FuncTestCaseParserOBracket, FuncTestCaseParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(380)
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
	p.EnterRule(localctx, 44, FuncTestCaseParserRULE_numericLiteral)
	p.SetState(386)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FuncTestCaseParserDecimalLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(383)
			p.Match(FuncTestCaseParserDecimalLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserIntegerLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(384)
			p.Match(FuncTestCaseParserIntegerLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserNaN, FuncTestCaseParserFloatLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(385)
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
	p.EnterRule(localctx, 46, FuncTestCaseParserRULE_floatLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(388)
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
	p.EnterRule(localctx, 48, FuncTestCaseParserRULE_nullArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(390)
		p.Match(FuncTestCaseParserNullLiteral)
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
	IntType() IIntTypeContext

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

func (s *IntArgContext) IntType() IIntTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntTypeContext)
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
	p.EnterRule(localctx, 50, FuncTestCaseParserRULE_intArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(394)
		p.Match(FuncTestCaseParserIntegerLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(395)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(396)
		p.IntType()
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
	FloatType() IFloatTypeContext

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

func (s *FloatArgContext) FloatType() IFloatTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloatTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloatTypeContext)
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
	p.EnterRule(localctx, 52, FuncTestCaseParserRULE_floatArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(398)
		p.NumericLiteral()
	}
	{
		p.SetState(399)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(400)
		p.FloatType()
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
	p.EnterRule(localctx, 54, FuncTestCaseParserRULE_decimalArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(402)
		p.NumericLiteral()
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
	p.EnterRule(localctx, 56, FuncTestCaseParserRULE_booleanArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(406)
		p.Match(FuncTestCaseParserBooleanLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(407)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(408)
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
	p.EnterRule(localctx, 58, FuncTestCaseParserRULE_stringArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(410)
		p.Match(FuncTestCaseParserStringLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(411)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(412)
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
	DateType() IDateTypeContext

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

func (s *DateArgContext) DateType() IDateTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDateTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDateTypeContext)
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
	p.EnterRule(localctx, 60, FuncTestCaseParserRULE_dateArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(414)
		p.Match(FuncTestCaseParserDateLiteral)
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
		p.DateType()
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
	TimeType() ITimeTypeContext

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

func (s *TimeArgContext) TimeType() ITimeTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimeTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimeTypeContext)
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
	p.EnterRule(localctx, 62, FuncTestCaseParserRULE_timeArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(418)
		p.Match(FuncTestCaseParserTimeLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(419)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(420)
		p.TimeType()
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
	p.EnterRule(localctx, 64, FuncTestCaseParserRULE_timestampArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(422)
		p.Match(FuncTestCaseParserTimestampLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(423)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(424)
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
	p.EnterRule(localctx, 66, FuncTestCaseParserRULE_timestampTzArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(426)
		p.Match(FuncTestCaseParserTimestampTzLiteral)
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
	p.EnterRule(localctx, 68, FuncTestCaseParserRULE_intervalYearArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(430)
		p.Match(FuncTestCaseParserIntervalYearLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(431)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(432)
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
	p.EnterRule(localctx, 70, FuncTestCaseParserRULE_intervalDayArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(434)
		p.Match(FuncTestCaseParserIntervalDayLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(435)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(436)
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

// IFixedCharArgContext is an interface to support dynamic dispatch.
type IFixedCharArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StringLiteral() antlr.TerminalNode
	DoubleColon() antlr.TerminalNode
	FixedCharType() IFixedCharTypeContext

	// IsFixedCharArgContext differentiates from other interfaces.
	IsFixedCharArgContext()
}

type FixedCharArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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
	p.EnterRule(localctx, 72, FuncTestCaseParserRULE_fixedCharArg)
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
		p.FixedCharType()
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

	// Getter signatures
	StringLiteral() antlr.TerminalNode
	DoubleColon() antlr.TerminalNode
	VarCharType() IVarCharTypeContext

	// IsVarCharArgContext differentiates from other interfaces.
	IsVarCharArgContext()
}

type VarCharArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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
	p.EnterRule(localctx, 74, FuncTestCaseParserRULE_varCharArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(442)
		p.Match(FuncTestCaseParserStringLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(443)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(444)
		p.VarCharType()
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

	// Getter signatures
	StringLiteral() antlr.TerminalNode
	DoubleColon() antlr.TerminalNode
	FixedBinaryType() IFixedBinaryTypeContext

	// IsFixedBinaryArgContext differentiates from other interfaces.
	IsFixedBinaryArgContext()
}

type FixedBinaryArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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
	p.EnterRule(localctx, 76, FuncTestCaseParserRULE_fixedBinaryArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(446)
		p.Match(FuncTestCaseParserStringLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(447)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(448)
		p.FixedBinaryType()
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

	// Getter signatures
	TimeLiteral() antlr.TerminalNode
	DoubleColon() antlr.TerminalNode
	PrecisionTimeType() IPrecisionTimeTypeContext

	// IsPrecisionTimeArgContext differentiates from other interfaces.
	IsPrecisionTimeArgContext()
}

type PrecisionTimeArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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
	p.EnterRule(localctx, 78, FuncTestCaseParserRULE_precisionTimeArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(450)
		p.Match(FuncTestCaseParserTimeLiteral)
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
		p.PrecisionTimeType()
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

	// Getter signatures
	TimestampLiteral() antlr.TerminalNode
	DoubleColon() antlr.TerminalNode
	PrecisionTimestampType() IPrecisionTimestampTypeContext

	// IsPrecisionTimestampArgContext differentiates from other interfaces.
	IsPrecisionTimestampArgContext()
}

type PrecisionTimestampArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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
	p.EnterRule(localctx, 80, FuncTestCaseParserRULE_precisionTimestampArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(454)
		p.Match(FuncTestCaseParserTimestampLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(455)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(456)
		p.PrecisionTimestampType()
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

	// Getter signatures
	TimestampTzLiteral() antlr.TerminalNode
	DoubleColon() antlr.TerminalNode
	PrecisionTimestampTZType() IPrecisionTimestampTZTypeContext

	// IsPrecisionTimestampTZArgContext differentiates from other interfaces.
	IsPrecisionTimestampTZArgContext()
}

type PrecisionTimestampTZArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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
	p.EnterRule(localctx, 82, FuncTestCaseParserRULE_precisionTimestampTZArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(458)
		p.Match(FuncTestCaseParserTimestampTzLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(459)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(460)
		p.PrecisionTimestampTZType()
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
	p.EnterRule(localctx, 84, FuncTestCaseParserRULE_listArg)
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

// ILambdaArgContext is an interface to support dynamic dispatch.
type ILambdaArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LiteralLambda() ILiteralLambdaContext
	DoubleColon() antlr.TerminalNode
	FuncType() IFuncTypeContext

	// IsLambdaArgContext differentiates from other interfaces.
	IsLambdaArgContext()
}

type LambdaArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLambdaArgContext() *LambdaArgContext {
	var p = new(LambdaArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_lambdaArg
	return p
}

func InitEmptyLambdaArgContext(p *LambdaArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_lambdaArg
}

func (*LambdaArgContext) IsLambdaArgContext() {}

func NewLambdaArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LambdaArgContext {
	var p = new(LambdaArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_lambdaArg

	return p
}

func (s *LambdaArgContext) GetParser() antlr.Parser { return s.parser }

func (s *LambdaArgContext) LiteralLambda() ILiteralLambdaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralLambdaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralLambdaContext)
}

func (s *LambdaArgContext) DoubleColon() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDoubleColon, 0)
}

func (s *LambdaArgContext) FuncType() IFuncTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncTypeContext)
}

func (s *LambdaArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LambdaArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LambdaArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitLambdaArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) LambdaArg() (localctx ILambdaArgContext) {
	localctx = NewLambdaArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, FuncTestCaseParserRULE_lambdaArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(466)
		p.LiteralLambda()
	}
	{
		p.SetState(467)
		p.Match(FuncTestCaseParserDoubleColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(468)
		p.FuncType()
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
	p.EnterRule(localctx, 88, FuncTestCaseParserRULE_literalList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(470)
		p.Match(FuncTestCaseParserOBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(479)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&65974976118784) != 0 {
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

		for _la == FuncTestCaseParserComma {
			{
				p.SetState(472)
				p.Match(FuncTestCaseParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(473)
				p.Literal()
			}

			p.SetState(478)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(481)
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

// ILiteralLambdaContext is an interface to support dynamic dispatch.
type ILiteralLambdaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OParen() antlr.TerminalNode
	LambdaParameters() ILambdaParametersContext
	Arrow() antlr.TerminalNode
	LambdaBody() ILambdaBodyContext
	CParen() antlr.TerminalNode

	// IsLiteralLambdaContext differentiates from other interfaces.
	IsLiteralLambdaContext()
}

type LiteralLambdaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralLambdaContext() *LiteralLambdaContext {
	var p = new(LiteralLambdaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_literalLambda
	return p
}

func InitEmptyLiteralLambdaContext(p *LiteralLambdaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_literalLambda
}

func (*LiteralLambdaContext) IsLiteralLambdaContext() {}

func NewLiteralLambdaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralLambdaContext {
	var p = new(LiteralLambdaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_literalLambda

	return p
}

func (s *LiteralLambdaContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralLambdaContext) OParen() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserOParen, 0)
}

func (s *LiteralLambdaContext) LambdaParameters() ILambdaParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILambdaParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILambdaParametersContext)
}

func (s *LiteralLambdaContext) Arrow() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserArrow, 0)
}

func (s *LiteralLambdaContext) LambdaBody() ILambdaBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILambdaBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILambdaBodyContext)
}

func (s *LiteralLambdaContext) CParen() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserCParen, 0)
}

func (s *LiteralLambdaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralLambdaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralLambdaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitLiteralLambda(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) LiteralLambda() (localctx ILiteralLambdaContext) {
	localctx = NewLiteralLambdaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, FuncTestCaseParserRULE_literalLambda)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(483)
		p.Match(FuncTestCaseParserOParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(484)
		p.LambdaParameters()
	}
	{
		p.SetState(485)
		p.Match(FuncTestCaseParserArrow)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(486)
		p.LambdaBody()
	}
	{
		p.SetState(487)
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

// ILambdaParametersContext is an interface to support dynamic dispatch.
type ILambdaParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLambdaParametersContext differentiates from other interfaces.
	IsLambdaParametersContext()
}

type LambdaParametersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLambdaParametersContext() *LambdaParametersContext {
	var p = new(LambdaParametersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_lambdaParameters
	return p
}

func InitEmptyLambdaParametersContext(p *LambdaParametersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_lambdaParameters
}

func (*LambdaParametersContext) IsLambdaParametersContext() {}

func NewLambdaParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LambdaParametersContext {
	var p = new(LambdaParametersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_lambdaParameters

	return p
}

func (s *LambdaParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *LambdaParametersContext) CopyAll(ctx *LambdaParametersContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LambdaParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LambdaParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TupleParamsContext struct {
	LambdaParametersContext
}

func NewTupleParamsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TupleParamsContext {
	var p = new(TupleParamsContext)

	InitEmptyLambdaParametersContext(&p.LambdaParametersContext)
	p.parser = parser
	p.CopyAll(ctx.(*LambdaParametersContext))

	return p
}

func (s *TupleParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleParamsContext) OParen() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserOParen, 0)
}

func (s *TupleParamsContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(FuncTestCaseParserIdentifier)
}

func (s *TupleParamsContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserIdentifier, i)
}

func (s *TupleParamsContext) CParen() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserCParen, 0)
}

func (s *TupleParamsContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(FuncTestCaseParserComma)
}

func (s *TupleParamsContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserComma, i)
}

func (s *TupleParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitTupleParams(s)

	default:
		return t.VisitChildren(s)
	}
}

type SingleParamContext struct {
	LambdaParametersContext
}

func NewSingleParamContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleParamContext {
	var p = new(SingleParamContext)

	InitEmptyLambdaParametersContext(&p.LambdaParametersContext)
	p.parser = parser
	p.CopyAll(ctx.(*LambdaParametersContext))

	return p
}

func (s *SingleParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleParamContext) Identifier() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserIdentifier, 0)
}

func (s *SingleParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitSingleParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) LambdaParameters() (localctx ILambdaParametersContext) {
	localctx = NewLambdaParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, FuncTestCaseParserRULE_lambdaParameters)
	var _la int

	p.SetState(499)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FuncTestCaseParserIdentifier:
		localctx = NewSingleParamContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(489)
			p.Match(FuncTestCaseParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FuncTestCaseParserOParen:
		localctx = NewTupleParamsContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(490)
			p.Match(FuncTestCaseParserOParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(491)
			p.Match(FuncTestCaseParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(494)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == FuncTestCaseParserComma {
			{
				p.SetState(492)
				p.Match(FuncTestCaseParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(493)
				p.Match(FuncTestCaseParserIdentifier)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(496)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(498)
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

// ILambdaBodyContext is an interface to support dynamic dispatch.
type ILambdaBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	OParen() antlr.TerminalNode
	Arguments() IArgumentsContext
	CParen() antlr.TerminalNode

	// IsLambdaBodyContext differentiates from other interfaces.
	IsLambdaBodyContext()
}

type LambdaBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLambdaBodyContext() *LambdaBodyContext {
	var p = new(LambdaBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_lambdaBody
	return p
}

func InitEmptyLambdaBodyContext(p *LambdaBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_lambdaBody
}

func (*LambdaBodyContext) IsLambdaBodyContext() {}

func NewLambdaBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LambdaBodyContext {
	var p = new(LambdaBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_lambdaBody

	return p
}

func (s *LambdaBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *LambdaBodyContext) Identifier() IIdentifierContext {
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

func (s *LambdaBodyContext) OParen() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserOParen, 0)
}

func (s *LambdaBodyContext) Arguments() IArgumentsContext {
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

func (s *LambdaBodyContext) CParen() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserCParen, 0)
}

func (s *LambdaBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LambdaBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LambdaBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitLambdaBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) LambdaBody() (localctx ILambdaBodyContext) {
	localctx = NewLambdaBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, FuncTestCaseParserRULE_lambdaBody)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(501)
		p.Identifier()
	}
	{
		p.SetState(502)
		p.Match(FuncTestCaseParserOParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(503)
		p.Arguments()
	}
	{
		p.SetState(504)
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
	p.EnterRule(localctx, 96, FuncTestCaseParserRULE_dataType)
	p.SetState(508)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FuncTestCaseParserBoolean, FuncTestCaseParserI8, FuncTestCaseParserI16, FuncTestCaseParserI32, FuncTestCaseParserI64, FuncTestCaseParserFP32, FuncTestCaseParserFP64, FuncTestCaseParserString_, FuncTestCaseParserBinary, FuncTestCaseParserTimestamp, FuncTestCaseParserTimestamp_TZ, FuncTestCaseParserDate, FuncTestCaseParserTime, FuncTestCaseParserInterval_Year, FuncTestCaseParserUUID, FuncTestCaseParserUserDefined, FuncTestCaseParserBool, FuncTestCaseParserStr, FuncTestCaseParserVBin, FuncTestCaseParserTs, FuncTestCaseParserTsTZ, FuncTestCaseParserIYear:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(506)
			p.ScalarType()
		}

	case FuncTestCaseParserFunc, FuncTestCaseParserInterval_Day, FuncTestCaseParserDecimal, FuncTestCaseParserPrecision_Time, FuncTestCaseParserPrecision_Timestamp, FuncTestCaseParserPrecision_Timestamp_TZ, FuncTestCaseParserFixedChar, FuncTestCaseParserVarChar, FuncTestCaseParserFixedBinary, FuncTestCaseParserIDay, FuncTestCaseParserDec, FuncTestCaseParserPT, FuncTestCaseParserPTs, FuncTestCaseParserPTsTZ, FuncTestCaseParserFChar, FuncTestCaseParserVChar, FuncTestCaseParserFBin:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(507)
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

func (s *DateContext) DateType() IDateTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDateTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDateTypeContext)
}

func (s *DateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitDate(s)

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

type UserDefinedContext struct {
	ScalarTypeContext
	isnull antlr.Token
}

func NewUserDefinedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UserDefinedContext {
	var p = new(UserDefinedContext)

	InitEmptyScalarTypeContext(&p.ScalarTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ScalarTypeContext))

	return p
}

func (s *UserDefinedContext) GetIsnull() antlr.Token { return s.isnull }

func (s *UserDefinedContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *UserDefinedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UserDefinedContext) UserDefined() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserUserDefined, 0)
}

func (s *UserDefinedContext) Identifier() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserIdentifier, 0)
}

func (s *UserDefinedContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *UserDefinedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitUserDefined(s)

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

func (s *TimeContext) TimeType() ITimeTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimeTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimeTypeContext)
}

func (s *TimeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitTime(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatContext struct {
	ScalarTypeContext
}

func NewFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatContext {
	var p = new(FloatContext)

	InitEmptyScalarTypeContext(&p.ScalarTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ScalarTypeContext))

	return p
}

func (s *FloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatContext) FloatType() IFloatTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloatTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloatTypeContext)
}

func (s *FloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitFloat(s)

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
	isnull antlr.Token
}

func NewUuidContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UuidContext {
	var p = new(UuidContext)

	InitEmptyScalarTypeContext(&p.ScalarTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ScalarTypeContext))

	return p
}

func (s *UuidContext) GetIsnull() antlr.Token { return s.isnull }

func (s *UuidContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *UuidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UuidContext) UUID() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserUUID, 0)
}

func (s *UuidContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *UuidContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitUuid(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntContext struct {
	ScalarTypeContext
}

func NewIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntContext {
	var p = new(IntContext)

	InitEmptyScalarTypeContext(&p.ScalarTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ScalarTypeContext))

	return p
}

func (s *IntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntContext) IntType() IIntTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntTypeContext)
}

func (s *IntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitInt(s)

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
	p.EnterRule(localctx, 98, FuncTestCaseParserRULE_scalarType)
	var _la int

	p.SetState(529)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FuncTestCaseParserBoolean, FuncTestCaseParserBool:
		localctx = NewBooleanContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(510)
			p.BooleanType()
		}

	case FuncTestCaseParserI8, FuncTestCaseParserI16, FuncTestCaseParserI32, FuncTestCaseParserI64:
		localctx = NewIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(511)
			p.IntType()
		}

	case FuncTestCaseParserFP32, FuncTestCaseParserFP64:
		localctx = NewFloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(512)
			p.FloatType()
		}

	case FuncTestCaseParserString_, FuncTestCaseParserStr:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(513)
			p.StringType()
		}

	case FuncTestCaseParserBinary, FuncTestCaseParserVBin:
		localctx = NewBinaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(514)
			p.BinaryType()
		}

	case FuncTestCaseParserTimestamp, FuncTestCaseParserTs:
		localctx = NewTimestampContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(515)
			p.TimestampType()
		}

	case FuncTestCaseParserTimestamp_TZ, FuncTestCaseParserTsTZ:
		localctx = NewTimestampTzContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(516)
			p.TimestampTZType()
		}

	case FuncTestCaseParserDate:
		localctx = NewDateContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(517)
			p.DateType()
		}

	case FuncTestCaseParserTime:
		localctx = NewTimeContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(518)
			p.TimeType()
		}

	case FuncTestCaseParserInterval_Year, FuncTestCaseParserIYear:
		localctx = NewIntervalYearContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(519)
			p.IntervalYearType()
		}

	case FuncTestCaseParserUUID:
		localctx = NewUuidContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(520)
			p.Match(FuncTestCaseParserUUID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(522)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FuncTestCaseParserQMark {
			{
				p.SetState(521)

				var _m = p.Match(FuncTestCaseParserQMark)

				localctx.(*UuidContext).isnull = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case FuncTestCaseParserUserDefined:
		localctx = NewUserDefinedContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(524)
			p.Match(FuncTestCaseParserUserDefined)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(525)
			p.Match(FuncTestCaseParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(527)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FuncTestCaseParserQMark {
			{
				p.SetState(526)

				var _m = p.Match(FuncTestCaseParserQMark)

				localctx.(*UserDefinedContext).isnull = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
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

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// Getter signatures
	Bool() antlr.TerminalNode
	Boolean() antlr.TerminalNode
	QMark() antlr.TerminalNode

	// IsBooleanTypeContext differentiates from other interfaces.
	IsBooleanTypeContext()
}

type BooleanTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
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

func (s *BooleanTypeContext) GetIsnull() antlr.Token { return s.isnull }

func (s *BooleanTypeContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *BooleanTypeContext) Bool() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserBool, 0)
}

func (s *BooleanTypeContext) Boolean() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserBoolean, 0)
}

func (s *BooleanTypeContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
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
	p.EnterRule(localctx, 100, FuncTestCaseParserRULE_booleanType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(531)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserBoolean || _la == FuncTestCaseParserBool) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(533)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(532)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*BooleanTypeContext).isnull = _m
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

// IStringTypeContext is an interface to support dynamic dispatch.
type IStringTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// Getter signatures
	Str() antlr.TerminalNode
	String_() antlr.TerminalNode
	QMark() antlr.TerminalNode

	// IsStringTypeContext differentiates from other interfaces.
	IsStringTypeContext()
}

type StringTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
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

func (s *StringTypeContext) GetIsnull() antlr.Token { return s.isnull }

func (s *StringTypeContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *StringTypeContext) Str() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserStr, 0)
}

func (s *StringTypeContext) String_() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserString_, 0)
}

func (s *StringTypeContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
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
	p.EnterRule(localctx, 102, FuncTestCaseParserRULE_stringType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(535)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserString_ || _la == FuncTestCaseParserStr) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(537)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(536)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*StringTypeContext).isnull = _m
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

// IBinaryTypeContext is an interface to support dynamic dispatch.
type IBinaryTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// Getter signatures
	Binary() antlr.TerminalNode
	VBin() antlr.TerminalNode
	QMark() antlr.TerminalNode

	// IsBinaryTypeContext differentiates from other interfaces.
	IsBinaryTypeContext()
}

type BinaryTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
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

func (s *BinaryTypeContext) GetIsnull() antlr.Token { return s.isnull }

func (s *BinaryTypeContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *BinaryTypeContext) Binary() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserBinary, 0)
}

func (s *BinaryTypeContext) VBin() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserVBin, 0)
}

func (s *BinaryTypeContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
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
	p.EnterRule(localctx, 104, FuncTestCaseParserRULE_binaryType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(539)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserBinary || _la == FuncTestCaseParserVBin) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(541)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(540)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*BinaryTypeContext).isnull = _m
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

// IIntTypeContext is an interface to support dynamic dispatch.
type IIntTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// Getter signatures
	I8() antlr.TerminalNode
	I16() antlr.TerminalNode
	I32() antlr.TerminalNode
	I64() antlr.TerminalNode
	QMark() antlr.TerminalNode

	// IsIntTypeContext differentiates from other interfaces.
	IsIntTypeContext()
}

type IntTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
}

func NewEmptyIntTypeContext() *IntTypeContext {
	var p = new(IntTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_intType
	return p
}

func InitEmptyIntTypeContext(p *IntTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_intType
}

func (*IntTypeContext) IsIntTypeContext() {}

func NewIntTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntTypeContext {
	var p = new(IntTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_intType

	return p
}

func (s *IntTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *IntTypeContext) GetIsnull() antlr.Token { return s.isnull }

func (s *IntTypeContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *IntTypeContext) I8() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserI8, 0)
}

func (s *IntTypeContext) I16() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserI16, 0)
}

func (s *IntTypeContext) I32() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserI32, 0)
}

func (s *IntTypeContext) I64() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserI64, 0)
}

func (s *IntTypeContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *IntTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitIntType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) IntType() (localctx IIntTypeContext) {
	localctx = NewIntTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, FuncTestCaseParserRULE_intType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(543)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&270215977642229760) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(545)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(544)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*IntTypeContext).isnull = _m
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

// IFloatTypeContext is an interface to support dynamic dispatch.
type IFloatTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// Getter signatures
	FP32() antlr.TerminalNode
	FP64() antlr.TerminalNode
	QMark() antlr.TerminalNode

	// IsFloatTypeContext differentiates from other interfaces.
	IsFloatTypeContext()
}

type FloatTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
}

func NewEmptyFloatTypeContext() *FloatTypeContext {
	var p = new(FloatTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_floatType
	return p
}

func InitEmptyFloatTypeContext(p *FloatTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_floatType
}

func (*FloatTypeContext) IsFloatTypeContext() {}

func NewFloatTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatTypeContext {
	var p = new(FloatTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_floatType

	return p
}

func (s *FloatTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatTypeContext) GetIsnull() antlr.Token { return s.isnull }

func (s *FloatTypeContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *FloatTypeContext) FP32() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserFP32, 0)
}

func (s *FloatTypeContext) FP64() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserFP64, 0)
}

func (s *FloatTypeContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *FloatTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitFloatType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) FloatType() (localctx IFloatTypeContext) {
	localctx = NewFloatTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, FuncTestCaseParserRULE_floatType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(547)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserFP32 || _la == FuncTestCaseParserFP64) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(549)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(548)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*FloatTypeContext).isnull = _m
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

// IDateTypeContext is an interface to support dynamic dispatch.
type IDateTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// Getter signatures
	Date() antlr.TerminalNode
	QMark() antlr.TerminalNode

	// IsDateTypeContext differentiates from other interfaces.
	IsDateTypeContext()
}

type DateTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
}

func NewEmptyDateTypeContext() *DateTypeContext {
	var p = new(DateTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_dateType
	return p
}

func InitEmptyDateTypeContext(p *DateTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_dateType
}

func (*DateTypeContext) IsDateTypeContext() {}

func NewDateTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateTypeContext {
	var p = new(DateTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_dateType

	return p
}

func (s *DateTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DateTypeContext) GetIsnull() antlr.Token { return s.isnull }

func (s *DateTypeContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *DateTypeContext) Date() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserDate, 0)
}

func (s *DateTypeContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *DateTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitDateType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) DateType() (localctx IDateTypeContext) {
	localctx = NewDateTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, FuncTestCaseParserRULE_dateType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(551)
		p.Match(FuncTestCaseParserDate)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(553)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(552)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*DateTypeContext).isnull = _m
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

// ITimeTypeContext is an interface to support dynamic dispatch.
type ITimeTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// Getter signatures
	Time() antlr.TerminalNode
	QMark() antlr.TerminalNode

	// IsTimeTypeContext differentiates from other interfaces.
	IsTimeTypeContext()
}

type TimeTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
}

func NewEmptyTimeTypeContext() *TimeTypeContext {
	var p = new(TimeTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_timeType
	return p
}

func InitEmptyTimeTypeContext(p *TimeTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_timeType
}

func (*TimeTypeContext) IsTimeTypeContext() {}

func NewTimeTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeTypeContext {
	var p = new(TimeTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_timeType

	return p
}

func (s *TimeTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeTypeContext) GetIsnull() antlr.Token { return s.isnull }

func (s *TimeTypeContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *TimeTypeContext) Time() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserTime, 0)
}

func (s *TimeTypeContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *TimeTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitTimeType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) TimeType() (localctx ITimeTypeContext) {
	localctx = NewTimeTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, FuncTestCaseParserRULE_timeType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(555)
		p.Match(FuncTestCaseParserTime)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(557)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(556)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*TimeTypeContext).isnull = _m
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

// ITimestampTypeContext is an interface to support dynamic dispatch.
type ITimestampTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// Getter signatures
	Ts() antlr.TerminalNode
	Timestamp() antlr.TerminalNode
	QMark() antlr.TerminalNode

	// IsTimestampTypeContext differentiates from other interfaces.
	IsTimestampTypeContext()
}

type TimestampTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
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

func (s *TimestampTypeContext) GetIsnull() antlr.Token { return s.isnull }

func (s *TimestampTypeContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *TimestampTypeContext) Ts() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserTs, 0)
}

func (s *TimestampTypeContext) Timestamp() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserTimestamp, 0)
}

func (s *TimestampTypeContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
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
	p.EnterRule(localctx, 114, FuncTestCaseParserRULE_timestampType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(559)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserTimestamp || _la == FuncTestCaseParserTs) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(561)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(560)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*TimestampTypeContext).isnull = _m
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

// ITimestampTZTypeContext is an interface to support dynamic dispatch.
type ITimestampTZTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// Getter signatures
	TsTZ() antlr.TerminalNode
	Timestamp_TZ() antlr.TerminalNode
	QMark() antlr.TerminalNode

	// IsTimestampTZTypeContext differentiates from other interfaces.
	IsTimestampTZTypeContext()
}

type TimestampTZTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
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

func (s *TimestampTZTypeContext) GetIsnull() antlr.Token { return s.isnull }

func (s *TimestampTZTypeContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *TimestampTZTypeContext) TsTZ() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserTsTZ, 0)
}

func (s *TimestampTZTypeContext) Timestamp_TZ() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserTimestamp_TZ, 0)
}

func (s *TimestampTZTypeContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
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
	p.EnterRule(localctx, 116, FuncTestCaseParserRULE_timestampTZType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(563)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserTimestamp_TZ || _la == FuncTestCaseParserTsTZ) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(565)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(564)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*TimestampTZTypeContext).isnull = _m
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

// IIntervalYearTypeContext is an interface to support dynamic dispatch.
type IIntervalYearTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// Getter signatures
	IYear() antlr.TerminalNode
	Interval_Year() antlr.TerminalNode
	QMark() antlr.TerminalNode

	// IsIntervalYearTypeContext differentiates from other interfaces.
	IsIntervalYearTypeContext()
}

type IntervalYearTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
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

func (s *IntervalYearTypeContext) GetIsnull() antlr.Token { return s.isnull }

func (s *IntervalYearTypeContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *IntervalYearTypeContext) IYear() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserIYear, 0)
}

func (s *IntervalYearTypeContext) Interval_Year() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserInterval_Year, 0)
}

func (s *IntervalYearTypeContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
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
	p.EnterRule(localctx, 118, FuncTestCaseParserRULE_intervalYearType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(567)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserInterval_Year || _la == FuncTestCaseParserIYear) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(569)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(568)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*IntervalYearTypeContext).isnull = _m
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
	p.EnterRule(localctx, 120, FuncTestCaseParserRULE_intervalDayType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(571)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserInterval_Day || _la == FuncTestCaseParserIDay) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(573)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(572)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*IntervalDayTypeContext).isnull = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(579)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserOAngleBracket {
		{
			p.SetState(575)
			p.Match(FuncTestCaseParserOAngleBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(576)

			var _x = p.NumericParameter()

			localctx.(*IntervalDayTypeContext).len_ = _x
		}
		{
			p.SetState(577)
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
	p.EnterRule(localctx, 122, FuncTestCaseParserRULE_fixedCharType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(581)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserFixedChar || _la == FuncTestCaseParserFChar) {
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

			localctx.(*FixedCharTypeContext).isnull = _m
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

		localctx.(*FixedCharTypeContext).len_ = _x
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
	p.EnterRule(localctx, 124, FuncTestCaseParserRULE_varCharType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(589)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserVarChar || _la == FuncTestCaseParserVChar) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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

			localctx.(*VarCharTypeContext).isnull = _m
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

		var _x = p.NumericParameter()

		localctx.(*VarCharTypeContext).len_ = _x
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
	p.EnterRule(localctx, 126, FuncTestCaseParserRULE_fixedBinaryType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(597)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserFixedBinary || _la == FuncTestCaseParserFBin) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(599)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(598)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*FixedBinaryTypeContext).isnull = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(601)
		p.Match(FuncTestCaseParserOAngleBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(602)

		var _x = p.NumericParameter()

		localctx.(*FixedBinaryTypeContext).len_ = _x
	}
	{
		p.SetState(603)
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
	p.EnterRule(localctx, 128, FuncTestCaseParserRULE_decimalType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(605)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserDecimal || _la == FuncTestCaseParserDec) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(607)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(606)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*DecimalTypeContext).isnull = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(615)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserOAngleBracket {
		{
			p.SetState(609)
			p.Match(FuncTestCaseParserOAngleBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(610)

			var _x = p.NumericParameter()

			localctx.(*DecimalTypeContext).precision = _x
		}
		{
			p.SetState(611)
			p.Match(FuncTestCaseParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(612)

			var _x = p.NumericParameter()

			localctx.(*DecimalTypeContext).scale = _x
		}
		{
			p.SetState(613)
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
	p.EnterRule(localctx, 130, FuncTestCaseParserRULE_precisionTimeType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(617)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserPrecision_Time || _la == FuncTestCaseParserPT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(619)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(618)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*PrecisionTimeTypeContext).isnull = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(621)
		p.Match(FuncTestCaseParserOAngleBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(622)

		var _x = p.NumericParameter()

		localctx.(*PrecisionTimeTypeContext).precision = _x
	}
	{
		p.SetState(623)
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
	p.EnterRule(localctx, 132, FuncTestCaseParserRULE_precisionTimestampType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(625)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserPrecision_Timestamp || _la == FuncTestCaseParserPTs) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(627)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(626)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*PrecisionTimestampTypeContext).isnull = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(629)
		p.Match(FuncTestCaseParserOAngleBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(630)

		var _x = p.NumericParameter()

		localctx.(*PrecisionTimestampTypeContext).precision = _x
	}
	{
		p.SetState(631)
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
	p.EnterRule(localctx, 134, FuncTestCaseParserRULE_precisionTimestampTZType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(633)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncTestCaseParserPrecision_Timestamp_TZ || _la == FuncTestCaseParserPTsTZ) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(635)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(634)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*PrecisionTimestampTZTypeContext).isnull = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(637)
		p.Match(FuncTestCaseParserOAngleBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(638)

		var _x = p.NumericParameter()

		localctx.(*PrecisionTimestampTZTypeContext).precision = _x
	}
	{
		p.SetState(639)
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
	p.EnterRule(localctx, 136, FuncTestCaseParserRULE_listType)
	var _la int

	localctx = NewListContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(641)
		p.Match(FuncTestCaseParserList)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(643)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(642)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*ListContext).isnull = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(645)
		p.Match(FuncTestCaseParserOAngleBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(646)

		var _x = p.DataType()

		localctx.(*ListContext).elemType = _x
	}
	{
		p.SetState(647)
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

// IFuncTypeContext is an interface to support dynamic dispatch.
type IFuncTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// GetParams returns the params rule contexts.
	GetParams() IFuncParametersContext

	// GetReturnType returns the returnType rule contexts.
	GetReturnType() IDataTypeContext

	// SetParams sets the params rule contexts.
	SetParams(IFuncParametersContext)

	// SetReturnType sets the returnType rule contexts.
	SetReturnType(IDataTypeContext)

	// Getter signatures
	Func() antlr.TerminalNode
	OAngleBracket() antlr.TerminalNode
	Arrow() antlr.TerminalNode
	CAngleBracket() antlr.TerminalNode
	FuncParameters() IFuncParametersContext
	DataType() IDataTypeContext
	QMark() antlr.TerminalNode

	// IsFuncTypeContext differentiates from other interfaces.
	IsFuncTypeContext()
}

type FuncTypeContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	isnull     antlr.Token
	params     IFuncParametersContext
	returnType IDataTypeContext
}

func NewEmptyFuncTypeContext() *FuncTypeContext {
	var p = new(FuncTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_funcType
	return p
}

func InitEmptyFuncTypeContext(p *FuncTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_funcType
}

func (*FuncTypeContext) IsFuncTypeContext() {}

func NewFuncTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncTypeContext {
	var p = new(FuncTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_funcType

	return p
}

func (s *FuncTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncTypeContext) GetIsnull() antlr.Token { return s.isnull }

func (s *FuncTypeContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *FuncTypeContext) GetParams() IFuncParametersContext { return s.params }

func (s *FuncTypeContext) GetReturnType() IDataTypeContext { return s.returnType }

func (s *FuncTypeContext) SetParams(v IFuncParametersContext) { s.params = v }

func (s *FuncTypeContext) SetReturnType(v IDataTypeContext) { s.returnType = v }

func (s *FuncTypeContext) Func() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserFunc, 0)
}

func (s *FuncTypeContext) OAngleBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserOAngleBracket, 0)
}

func (s *FuncTypeContext) Arrow() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserArrow, 0)
}

func (s *FuncTypeContext) CAngleBracket() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserCAngleBracket, 0)
}

func (s *FuncTypeContext) FuncParameters() IFuncParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncParametersContext)
}

func (s *FuncTypeContext) DataType() IDataTypeContext {
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

func (s *FuncTypeContext) QMark() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserQMark, 0)
}

func (s *FuncTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitFuncType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) FuncType() (localctx IFuncTypeContext) {
	localctx = NewFuncTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 138, FuncTestCaseParserRULE_funcType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(649)
		p.Match(FuncTestCaseParserFunc)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(651)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncTestCaseParserQMark {
		{
			p.SetState(650)

			var _m = p.Match(FuncTestCaseParserQMark)

			localctx.(*FuncTypeContext).isnull = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(653)
		p.Match(FuncTestCaseParserOAngleBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(654)

		var _x = p.FuncParameters()

		localctx.(*FuncTypeContext).params = _x
	}
	{
		p.SetState(655)
		p.Match(FuncTestCaseParserArrow)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(656)

		var _x = p.DataType()

		localctx.(*FuncTypeContext).returnType = _x
	}
	{
		p.SetState(657)
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

// IFuncParametersContext is an interface to support dynamic dispatch.
type IFuncParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFuncParametersContext differentiates from other interfaces.
	IsFuncParametersContext()
}

type FuncParametersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncParametersContext() *FuncParametersContext {
	var p = new(FuncParametersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_funcParameters
	return p
}

func InitEmptyFuncParametersContext(p *FuncParametersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncTestCaseParserRULE_funcParameters
}

func (*FuncParametersContext) IsFuncParametersContext() {}

func NewFuncParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncParametersContext {
	var p = new(FuncParametersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncTestCaseParserRULE_funcParameters

	return p
}

func (s *FuncParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncParametersContext) CopyAll(ctx *FuncParametersContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FuncParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SingleFuncParamContext struct {
	FuncParametersContext
}

func NewSingleFuncParamContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleFuncParamContext {
	var p = new(SingleFuncParamContext)

	InitEmptyFuncParametersContext(&p.FuncParametersContext)
	p.parser = parser
	p.CopyAll(ctx.(*FuncParametersContext))

	return p
}

func (s *SingleFuncParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleFuncParamContext) DataType() IDataTypeContext {
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

func (s *SingleFuncParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitSingleFuncParam(s)

	default:
		return t.VisitChildren(s)
	}
}

type FuncParamsWithParensContext struct {
	FuncParametersContext
}

func NewFuncParamsWithParensContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncParamsWithParensContext {
	var p = new(FuncParamsWithParensContext)

	InitEmptyFuncParametersContext(&p.FuncParametersContext)
	p.parser = parser
	p.CopyAll(ctx.(*FuncParametersContext))

	return p
}

func (s *FuncParamsWithParensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncParamsWithParensContext) OParen() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserOParen, 0)
}

func (s *FuncParamsWithParensContext) AllDataType() []IDataTypeContext {
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

func (s *FuncParamsWithParensContext) DataType(i int) IDataTypeContext {
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

func (s *FuncParamsWithParensContext) CParen() antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserCParen, 0)
}

func (s *FuncParamsWithParensContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(FuncTestCaseParserComma)
}

func (s *FuncParamsWithParensContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(FuncTestCaseParserComma, i)
}

func (s *FuncParamsWithParensContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FuncTestCaseParserVisitor:
		return t.VisitFuncParamsWithParens(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FuncTestCaseParser) FuncParameters() (localctx IFuncParametersContext) {
	localctx = NewFuncParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 140, FuncTestCaseParserRULE_funcParameters)
	var _la int

	p.SetState(671)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FuncTestCaseParserFunc, FuncTestCaseParserBoolean, FuncTestCaseParserI8, FuncTestCaseParserI16, FuncTestCaseParserI32, FuncTestCaseParserI64, FuncTestCaseParserFP32, FuncTestCaseParserFP64, FuncTestCaseParserString_, FuncTestCaseParserBinary, FuncTestCaseParserTimestamp, FuncTestCaseParserTimestamp_TZ, FuncTestCaseParserDate, FuncTestCaseParserTime, FuncTestCaseParserInterval_Year, FuncTestCaseParserInterval_Day, FuncTestCaseParserUUID, FuncTestCaseParserDecimal, FuncTestCaseParserPrecision_Time, FuncTestCaseParserPrecision_Timestamp, FuncTestCaseParserPrecision_Timestamp_TZ, FuncTestCaseParserFixedChar, FuncTestCaseParserVarChar, FuncTestCaseParserFixedBinary, FuncTestCaseParserUserDefined, FuncTestCaseParserBool, FuncTestCaseParserStr, FuncTestCaseParserVBin, FuncTestCaseParserTs, FuncTestCaseParserTsTZ, FuncTestCaseParserIYear, FuncTestCaseParserIDay, FuncTestCaseParserDec, FuncTestCaseParserPT, FuncTestCaseParserPTs, FuncTestCaseParserPTsTZ, FuncTestCaseParserFChar, FuncTestCaseParserVChar, FuncTestCaseParserFBin:
		localctx = NewSingleFuncParamContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(659)
			p.DataType()
		}

	case FuncTestCaseParserOParen:
		localctx = NewFuncParamsWithParensContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(660)
			p.Match(FuncTestCaseParserOParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(661)
			p.DataType()
		}
		p.SetState(666)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FuncTestCaseParserComma {
			{
				p.SetState(662)
				p.Match(FuncTestCaseParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(663)
				p.DataType()
			}

			p.SetState(668)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(669)
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
	FuncType() IFuncTypeContext

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

func (s *ParameterizedTypeContext) FuncType() IFuncTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncTypeContext)
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
	p.EnterRule(localctx, 142, FuncTestCaseParserRULE_parameterizedType)
	p.SetState(682)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FuncTestCaseParserFixedChar, FuncTestCaseParserFChar:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(673)
			p.FixedCharType()
		}

	case FuncTestCaseParserVarChar, FuncTestCaseParserVChar:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(674)
			p.VarCharType()
		}

	case FuncTestCaseParserFixedBinary, FuncTestCaseParserFBin:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(675)
			p.FixedBinaryType()
		}

	case FuncTestCaseParserDecimal, FuncTestCaseParserDec:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(676)
			p.DecimalType()
		}

	case FuncTestCaseParserInterval_Day, FuncTestCaseParserIDay:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(677)
			p.IntervalDayType()
		}

	case FuncTestCaseParserPrecision_Time, FuncTestCaseParserPT:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(678)
			p.PrecisionTimeType()
		}

	case FuncTestCaseParserPrecision_Timestamp, FuncTestCaseParserPTs:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(679)
			p.PrecisionTimestampType()
		}

	case FuncTestCaseParserPrecision_Timestamp_TZ, FuncTestCaseParserPTsTZ:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(680)
			p.PrecisionTimestampTZType()
		}

	case FuncTestCaseParserFunc:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(681)
			p.FuncType()
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
	p.EnterRule(localctx, 144, FuncTestCaseParserRULE_numericParameter)
	localctx = NewIntegerLiteralContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(684)
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
	p.EnterRule(localctx, 146, FuncTestCaseParserRULE_substraitError)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(686)
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
	p.EnterRule(localctx, 148, FuncTestCaseParserRULE_funcOption)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(688)
		p.OptionName()
	}
	{
		p.SetState(689)
		p.Match(FuncTestCaseParserColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(690)
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
	p.EnterRule(localctx, 150, FuncTestCaseParserRULE_optionName)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(692)
		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6303744) != 0) || _la == FuncTestCaseParserIdentifier) {
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
	p.EnterRule(localctx, 152, FuncTestCaseParserRULE_optionValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(694)
		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17592330731520) != 0) || _la == FuncTestCaseParserIdentifier) {
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
	p.EnterRule(localctx, 154, FuncTestCaseParserRULE_funcOptions)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(696)
		p.FuncOption()
	}
	p.SetState(701)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FuncTestCaseParserComma {
		{
			p.SetState(697)
			p.Match(FuncTestCaseParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(698)
			p.FuncOption()
		}

		p.SetState(703)
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
	p.EnterRule(localctx, 156, FuncTestCaseParserRULE_nonReserved)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(704)
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
	p.EnterRule(localctx, 158, FuncTestCaseParserRULE_identifier)
	p.SetState(708)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FuncTestCaseParserTruncate, FuncTestCaseParserAnd, FuncTestCaseParserOr:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(706)
			p.NonReserved()
		}

	case FuncTestCaseParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(707)
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
