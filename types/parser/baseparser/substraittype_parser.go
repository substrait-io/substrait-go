// Code generated from SubstraitType.g4 by ANTLR 4.13.2. DO NOT EDIT.

package baseparser // SubstraitType
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

type SubstraitTypeParser struct {
	*antlr.BaseParser
}

var SubstraitTypeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func substraittypeParserInit() {
	staticData := &SubstraitTypeParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "'IF'", "'THEN'", "'ELSE'", "'FUNC'", "'BOOLEAN'", "'I8'",
		"'I16'", "'I32'", "'I64'", "'FP32'", "'FP64'", "'STRING'", "'BINARY'",
		"'TIMESTAMP'", "'TIMESTAMP_TZ'", "'DATE'", "'TIME'", "'INTERVAL_YEAR'",
		"'INTERVAL_DAY'", "'UUID'", "'DECIMAL'", "'PRECISION_TIME'", "'PRECISION_TIMESTAMP'",
		"'PRECISION_TIMESTAMP_TZ'", "'FIXEDCHAR'", "'VARCHAR'", "'FIXEDBINARY'",
		"'STRUCT'", "'NSTRUCT'", "'LIST'", "'MAP'", "'U!'", "'BOOL'", "'STR'",
		"'VBIN'", "'TS'", "'TSTZ'", "'IYEAR'", "'IDAY'", "'DEC'", "'PT'", "'PTS'",
		"'PTSTZ'", "'FCHAR'", "'VCHAR'", "'FBIN'", "'ANY'", "", "'::'", "'+'",
		"'-'", "'*'", "'/'", "'%'", "'='", "'!='", "'>='", "'<='", "'>'", "'<'",
		"'!'", "", "", "'('", "')'", "'['", "']'", "','", "':'", "'?'", "'#'",
		"'.'", "'AND'", "'OR'", "':='", "'->'",
	}
	staticData.SymbolicNames = []string{
		"", "LineComment", "BlockComment", "Whitespace", "If", "Then", "Else",
		"Func", "Boolean", "I8", "I16", "I32", "I64", "FP32", "FP64", "String",
		"Binary", "Timestamp", "Timestamp_TZ", "Date", "Time", "Interval_Year",
		"Interval_Day", "UUID", "Decimal", "Precision_Time", "Precision_Timestamp",
		"Precision_Timestamp_TZ", "FixedChar", "VarChar", "FixedBinary", "Struct",
		"NStruct", "List", "Map", "UserDefined", "Bool", "Str", "VBin", "Ts",
		"TsTZ", "IYear", "IDay", "Dec", "PT", "PTs", "PTsTZ", "FChar", "VChar",
		"FBin", "Any", "AnyVar", "DoubleColon", "Plus", "Minus", "Asterisk",
		"ForwardSlash", "Percent", "Eq", "Ne", "Gte", "Lte", "Gt", "Lt", "Bang",
		"OAngleBracket", "CAngleBracket", "OParen", "CParen", "OBracket", "CBracket",
		"Comma", "Colon", "QMark", "Hash", "Dot", "And", "Or", "Assign", "Arrow",
		"Number", "Identifier", "Newline",
	}
	staticData.RuleNames = []string{
		"startRule", "typeStatement", "scalarType", "parameterizedType", "funcParams",
		"numericParameter", "anyType", "typeDef", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 82, 302, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 40, 8, 2, 1, 3, 1, 3, 3, 3, 44,
		8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 52, 8, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 3, 3, 60, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		3, 3, 68, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 78,
		8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 86, 8, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 3, 3, 94, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		3, 3, 102, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 110, 8, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 5, 3, 116, 8, 3, 10, 3, 12, 3, 119, 9, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 3, 3, 125, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5,
		3, 133, 8, 3, 10, 3, 12, 3, 136, 9, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 142,
		8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 150, 8, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 160, 8, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 171, 8, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 5, 3, 177, 8, 3, 10, 3, 12, 3, 180, 9, 3, 1, 3, 1, 3, 3, 3, 184,
		8, 3, 3, 3, 186, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 193, 8, 4, 10,
		4, 12, 4, 196, 9, 4, 1, 4, 1, 4, 3, 4, 200, 8, 4, 1, 5, 1, 5, 1, 5, 3,
		5, 205, 8, 5, 1, 6, 1, 6, 3, 6, 209, 8, 6, 1, 6, 1, 6, 3, 6, 213, 8, 6,
		3, 6, 215, 8, 6, 1, 7, 1, 7, 3, 7, 219, 8, 7, 1, 7, 1, 7, 3, 7, 223, 8,
		7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 4, 8, 234, 8,
		8, 11, 8, 12, 8, 235, 1, 8, 1, 8, 1, 8, 1, 8, 4, 8, 242, 8, 8, 11, 8, 12,
		8, 243, 5, 8, 246, 8, 8, 10, 8, 12, 8, 249, 9, 8, 1, 8, 1, 8, 5, 8, 253,
		8, 8, 10, 8, 12, 8, 256, 9, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 262, 8, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 269, 8, 8, 10, 8, 12, 8, 272, 9, 8,
		3, 8, 274, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 3, 8, 286, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 5, 8, 297, 8, 8, 10, 8, 12, 8, 300, 9, 8, 1, 8, 0, 1, 16, 9, 0, 2,
		4, 6, 8, 10, 12, 14, 16, 0, 1, 3, 0, 53, 56, 58, 63, 76, 77, 363, 0, 18,
		1, 0, 0, 0, 2, 21, 1, 0, 0, 0, 4, 39, 1, 0, 0, 0, 6, 185, 1, 0, 0, 0, 8,
		199, 1, 0, 0, 0, 10, 204, 1, 0, 0, 0, 12, 214, 1, 0, 0, 0, 14, 222, 1,
		0, 0, 0, 16, 285, 1, 0, 0, 0, 18, 19, 3, 16, 8, 0, 19, 20, 5, 0, 0, 1,
		20, 1, 1, 0, 0, 0, 21, 22, 3, 14, 7, 0, 22, 23, 5, 0, 0, 1, 23, 3, 1, 0,
		0, 0, 24, 40, 5, 8, 0, 0, 25, 40, 5, 9, 0, 0, 26, 40, 5, 10, 0, 0, 27,
		40, 5, 11, 0, 0, 28, 40, 5, 12, 0, 0, 29, 40, 5, 13, 0, 0, 30, 40, 5, 14,
		0, 0, 31, 40, 5, 15, 0, 0, 32, 40, 5, 16, 0, 0, 33, 40, 5, 17, 0, 0, 34,
		40, 5, 18, 0, 0, 35, 40, 5, 19, 0, 0, 36, 40, 5, 20, 0, 0, 37, 40, 5, 21,
		0, 0, 38, 40, 5, 23, 0, 0, 39, 24, 1, 0, 0, 0, 39, 25, 1, 0, 0, 0, 39,
		26, 1, 0, 0, 0, 39, 27, 1, 0, 0, 0, 39, 28, 1, 0, 0, 0, 39, 29, 1, 0, 0,
		0, 39, 30, 1, 0, 0, 0, 39, 31, 1, 0, 0, 0, 39, 32, 1, 0, 0, 0, 39, 33,
		1, 0, 0, 0, 39, 34, 1, 0, 0, 0, 39, 35, 1, 0, 0, 0, 39, 36, 1, 0, 0, 0,
		39, 37, 1, 0, 0, 0, 39, 38, 1, 0, 0, 0, 40, 5, 1, 0, 0, 0, 41, 43, 5, 28,
		0, 0, 42, 44, 5, 73, 0, 0, 43, 42, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44,
		45, 1, 0, 0, 0, 45, 46, 5, 63, 0, 0, 46, 47, 3, 10, 5, 0, 47, 48, 5, 62,
		0, 0, 48, 186, 1, 0, 0, 0, 49, 51, 5, 29, 0, 0, 50, 52, 5, 73, 0, 0, 51,
		50, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 54, 5, 63,
		0, 0, 54, 55, 3, 10, 5, 0, 55, 56, 5, 62, 0, 0, 56, 186, 1, 0, 0, 0, 57,
		59, 5, 30, 0, 0, 58, 60, 5, 73, 0, 0, 59, 58, 1, 0, 0, 0, 59, 60, 1, 0,
		0, 0, 60, 61, 1, 0, 0, 0, 61, 62, 5, 63, 0, 0, 62, 63, 3, 10, 5, 0, 63,
		64, 5, 62, 0, 0, 64, 186, 1, 0, 0, 0, 65, 67, 5, 24, 0, 0, 66, 68, 5, 73,
		0, 0, 67, 66, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 70,
		5, 63, 0, 0, 70, 71, 3, 10, 5, 0, 71, 72, 5, 71, 0, 0, 72, 73, 3, 10, 5,
		0, 73, 74, 5, 62, 0, 0, 74, 186, 1, 0, 0, 0, 75, 77, 5, 22, 0, 0, 76, 78,
		5, 73, 0, 0, 77, 76, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0,
		79, 80, 5, 63, 0, 0, 80, 81, 3, 10, 5, 0, 81, 82, 5, 62, 0, 0, 82, 186,
		1, 0, 0, 0, 83, 85, 5, 25, 0, 0, 84, 86, 5, 73, 0, 0, 85, 84, 1, 0, 0,
		0, 85, 86, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 88, 5, 63, 0, 0, 88, 89,
		3, 10, 5, 0, 89, 90, 5, 62, 0, 0, 90, 186, 1, 0, 0, 0, 91, 93, 5, 26, 0,
		0, 92, 94, 5, 73, 0, 0, 93, 92, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 95,
		1, 0, 0, 0, 95, 96, 5, 63, 0, 0, 96, 97, 3, 10, 5, 0, 97, 98, 5, 62, 0,
		0, 98, 186, 1, 0, 0, 0, 99, 101, 5, 27, 0, 0, 100, 102, 5, 73, 0, 0, 101,
		100, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 104,
		5, 63, 0, 0, 104, 105, 3, 10, 5, 0, 105, 106, 5, 62, 0, 0, 106, 186, 1,
		0, 0, 0, 107, 109, 5, 31, 0, 0, 108, 110, 5, 73, 0, 0, 109, 108, 1, 0,
		0, 0, 109, 110, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 112, 5, 63, 0, 0,
		112, 117, 3, 16, 8, 0, 113, 114, 5, 71, 0, 0, 114, 116, 3, 16, 8, 0, 115,
		113, 1, 0, 0, 0, 116, 119, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 117, 118,
		1, 0, 0, 0, 118, 120, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 120, 121, 5, 62,
		0, 0, 121, 186, 1, 0, 0, 0, 122, 124, 5, 32, 0, 0, 123, 125, 5, 73, 0,
		0, 124, 123, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126,
		127, 5, 63, 0, 0, 127, 128, 5, 81, 0, 0, 128, 134, 3, 16, 8, 0, 129, 130,
		5, 71, 0, 0, 130, 131, 5, 81, 0, 0, 131, 133, 3, 16, 8, 0, 132, 129, 1,
		0, 0, 0, 133, 136, 1, 0, 0, 0, 134, 132, 1, 0, 0, 0, 134, 135, 1, 0, 0,
		0, 135, 137, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 137, 138, 5, 62, 0, 0, 138,
		186, 1, 0, 0, 0, 139, 141, 5, 33, 0, 0, 140, 142, 5, 73, 0, 0, 141, 140,
		1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 144, 5, 63,
		0, 0, 144, 145, 3, 16, 8, 0, 145, 146, 5, 62, 0, 0, 146, 186, 1, 0, 0,
		0, 147, 149, 5, 34, 0, 0, 148, 150, 5, 73, 0, 0, 149, 148, 1, 0, 0, 0,
		149, 150, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 152, 5, 63, 0, 0, 152,
		153, 3, 16, 8, 0, 153, 154, 5, 71, 0, 0, 154, 155, 3, 16, 8, 0, 155, 156,
		5, 62, 0, 0, 156, 186, 1, 0, 0, 0, 157, 159, 5, 7, 0, 0, 158, 160, 5, 73,
		0, 0, 159, 158, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0,
		161, 162, 5, 63, 0, 0, 162, 163, 3, 8, 4, 0, 163, 164, 5, 79, 0, 0, 164,
		165, 3, 16, 8, 0, 165, 166, 5, 62, 0, 0, 166, 186, 1, 0, 0, 0, 167, 168,
		5, 35, 0, 0, 168, 170, 5, 81, 0, 0, 169, 171, 5, 73, 0, 0, 170, 169, 1,
		0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 183, 1, 0, 0, 0, 172, 173, 5, 63, 0,
		0, 173, 178, 3, 16, 8, 0, 174, 175, 5, 71, 0, 0, 175, 177, 3, 16, 8, 0,
		176, 174, 1, 0, 0, 0, 177, 180, 1, 0, 0, 0, 178, 176, 1, 0, 0, 0, 178,
		179, 1, 0, 0, 0, 179, 181, 1, 0, 0, 0, 180, 178, 1, 0, 0, 0, 181, 182,
		5, 62, 0, 0, 182, 184, 1, 0, 0, 0, 183, 172, 1, 0, 0, 0, 183, 184, 1, 0,
		0, 0, 184, 186, 1, 0, 0, 0, 185, 41, 1, 0, 0, 0, 185, 49, 1, 0, 0, 0, 185,
		57, 1, 0, 0, 0, 185, 65, 1, 0, 0, 0, 185, 75, 1, 0, 0, 0, 185, 83, 1, 0,
		0, 0, 185, 91, 1, 0, 0, 0, 185, 99, 1, 0, 0, 0, 185, 107, 1, 0, 0, 0, 185,
		122, 1, 0, 0, 0, 185, 139, 1, 0, 0, 0, 185, 147, 1, 0, 0, 0, 185, 157,
		1, 0, 0, 0, 185, 167, 1, 0, 0, 0, 186, 7, 1, 0, 0, 0, 187, 200, 3, 16,
		8, 0, 188, 189, 5, 67, 0, 0, 189, 194, 3, 16, 8, 0, 190, 191, 5, 71, 0,
		0, 191, 193, 3, 16, 8, 0, 192, 190, 1, 0, 0, 0, 193, 196, 1, 0, 0, 0, 194,
		192, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 197, 1, 0, 0, 0, 196, 194,
		1, 0, 0, 0, 197, 198, 5, 68, 0, 0, 198, 200, 1, 0, 0, 0, 199, 187, 1, 0,
		0, 0, 199, 188, 1, 0, 0, 0, 200, 9, 1, 0, 0, 0, 201, 205, 5, 80, 0, 0,
		202, 205, 5, 81, 0, 0, 203, 205, 3, 16, 8, 0, 204, 201, 1, 0, 0, 0, 204,
		202, 1, 0, 0, 0, 204, 203, 1, 0, 0, 0, 205, 11, 1, 0, 0, 0, 206, 208, 5,
		50, 0, 0, 207, 209, 5, 73, 0, 0, 208, 207, 1, 0, 0, 0, 208, 209, 1, 0,
		0, 0, 209, 215, 1, 0, 0, 0, 210, 212, 5, 51, 0, 0, 211, 213, 5, 73, 0,
		0, 212, 211, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 215, 1, 0, 0, 0, 214,
		206, 1, 0, 0, 0, 214, 210, 1, 0, 0, 0, 215, 13, 1, 0, 0, 0, 216, 218, 3,
		4, 2, 0, 217, 219, 5, 73, 0, 0, 218, 217, 1, 0, 0, 0, 218, 219, 1, 0, 0,
		0, 219, 223, 1, 0, 0, 0, 220, 223, 3, 6, 3, 0, 221, 223, 3, 12, 6, 0, 222,
		216, 1, 0, 0, 0, 222, 220, 1, 0, 0, 0, 222, 221, 1, 0, 0, 0, 223, 15, 1,
		0, 0, 0, 224, 225, 6, 8, -1, 0, 225, 226, 5, 67, 0, 0, 226, 227, 3, 16,
		8, 0, 227, 228, 5, 68, 0, 0, 228, 286, 1, 0, 0, 0, 229, 230, 5, 81, 0,
		0, 230, 231, 5, 58, 0, 0, 231, 233, 3, 16, 8, 0, 232, 234, 5, 82, 0, 0,
		233, 232, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235, 233, 1, 0, 0, 0, 235,
		236, 1, 0, 0, 0, 236, 247, 1, 0, 0, 0, 237, 238, 5, 81, 0, 0, 238, 239,
		5, 58, 0, 0, 239, 241, 3, 16, 8, 0, 240, 242, 5, 82, 0, 0, 241, 240, 1,
		0, 0, 0, 242, 243, 1, 0, 0, 0, 243, 241, 1, 0, 0, 0, 243, 244, 1, 0, 0,
		0, 244, 246, 1, 0, 0, 0, 245, 237, 1, 0, 0, 0, 246, 249, 1, 0, 0, 0, 247,
		245, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 250, 1, 0, 0, 0, 249, 247,
		1, 0, 0, 0, 250, 254, 3, 14, 7, 0, 251, 253, 5, 82, 0, 0, 252, 251, 1,
		0, 0, 0, 253, 256, 1, 0, 0, 0, 254, 252, 1, 0, 0, 0, 254, 255, 1, 0, 0,
		0, 255, 286, 1, 0, 0, 0, 256, 254, 1, 0, 0, 0, 257, 286, 3, 14, 7, 0, 258,
		286, 5, 80, 0, 0, 259, 261, 5, 81, 0, 0, 260, 262, 5, 73, 0, 0, 261, 260,
		1, 0, 0, 0, 261, 262, 1, 0, 0, 0, 262, 286, 1, 0, 0, 0, 263, 264, 5, 81,
		0, 0, 264, 273, 5, 67, 0, 0, 265, 270, 3, 16, 8, 0, 266, 267, 5, 71, 0,
		0, 267, 269, 3, 16, 8, 0, 268, 266, 1, 0, 0, 0, 269, 272, 1, 0, 0, 0, 270,
		268, 1, 0, 0, 0, 270, 271, 1, 0, 0, 0, 271, 274, 1, 0, 0, 0, 272, 270,
		1, 0, 0, 0, 273, 265, 1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 275, 1, 0,
		0, 0, 275, 286, 5, 68, 0, 0, 276, 277, 5, 4, 0, 0, 277, 278, 3, 16, 8,
		0, 278, 279, 5, 5, 0, 0, 279, 280, 3, 16, 8, 0, 280, 281, 5, 6, 0, 0, 281,
		282, 3, 16, 8, 3, 282, 286, 1, 0, 0, 0, 283, 284, 5, 64, 0, 0, 284, 286,
		3, 16, 8, 2, 285, 224, 1, 0, 0, 0, 285, 229, 1, 0, 0, 0, 285, 257, 1, 0,
		0, 0, 285, 258, 1, 0, 0, 0, 285, 259, 1, 0, 0, 0, 285, 263, 1, 0, 0, 0,
		285, 276, 1, 0, 0, 0, 285, 283, 1, 0, 0, 0, 286, 298, 1, 0, 0, 0, 287,
		288, 10, 4, 0, 0, 288, 289, 7, 0, 0, 0, 289, 297, 3, 16, 8, 5, 290, 291,
		10, 1, 0, 0, 291, 292, 5, 73, 0, 0, 292, 293, 3, 16, 8, 0, 293, 294, 5,
		72, 0, 0, 294, 295, 3, 16, 8, 2, 295, 297, 1, 0, 0, 0, 296, 287, 1, 0,
		0, 0, 296, 290, 1, 0, 0, 0, 297, 300, 1, 0, 0, 0, 298, 296, 1, 0, 0, 0,
		298, 299, 1, 0, 0, 0, 299, 17, 1, 0, 0, 0, 300, 298, 1, 0, 0, 0, 38, 39,
		43, 51, 59, 67, 77, 85, 93, 101, 109, 117, 124, 134, 141, 149, 159, 170,
		178, 183, 185, 194, 199, 204, 208, 212, 214, 218, 222, 235, 243, 247, 254,
		261, 270, 273, 285, 296, 298,
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

// SubstraitTypeParserInit initializes any static state used to implement SubstraitTypeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSubstraitTypeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SubstraitTypeParserInit() {
	staticData := &SubstraitTypeParserStaticData
	staticData.once.Do(substraittypeParserInit)
}

// NewSubstraitTypeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSubstraitTypeParser(input antlr.TokenStream) *SubstraitTypeParser {
	SubstraitTypeParserInit()
	this := new(SubstraitTypeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SubstraitTypeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "SubstraitType.g4"

	return this
}

// SubstraitTypeParser tokens.
const (
	SubstraitTypeParserEOF                    = antlr.TokenEOF
	SubstraitTypeParserLineComment            = 1
	SubstraitTypeParserBlockComment           = 2
	SubstraitTypeParserWhitespace             = 3
	SubstraitTypeParserIf                     = 4
	SubstraitTypeParserThen                   = 5
	SubstraitTypeParserElse                   = 6
	SubstraitTypeParserFunc                   = 7
	SubstraitTypeParserBoolean                = 8
	SubstraitTypeParserI8                     = 9
	SubstraitTypeParserI16                    = 10
	SubstraitTypeParserI32                    = 11
	SubstraitTypeParserI64                    = 12
	SubstraitTypeParserFP32                   = 13
	SubstraitTypeParserFP64                   = 14
	SubstraitTypeParserString_                = 15
	SubstraitTypeParserBinary                 = 16
	SubstraitTypeParserTimestamp              = 17
	SubstraitTypeParserTimestamp_TZ           = 18
	SubstraitTypeParserDate                   = 19
	SubstraitTypeParserTime                   = 20
	SubstraitTypeParserInterval_Year          = 21
	SubstraitTypeParserInterval_Day           = 22
	SubstraitTypeParserUUID                   = 23
	SubstraitTypeParserDecimal                = 24
	SubstraitTypeParserPrecision_Time         = 25
	SubstraitTypeParserPrecision_Timestamp    = 26
	SubstraitTypeParserPrecision_Timestamp_TZ = 27
	SubstraitTypeParserFixedChar              = 28
	SubstraitTypeParserVarChar                = 29
	SubstraitTypeParserFixedBinary            = 30
	SubstraitTypeParserStruct                 = 31
	SubstraitTypeParserNStruct                = 32
	SubstraitTypeParserList                   = 33
	SubstraitTypeParserMap                    = 34
	SubstraitTypeParserUserDefined            = 35
	SubstraitTypeParserBool                   = 36
	SubstraitTypeParserStr                    = 37
	SubstraitTypeParserVBin                   = 38
	SubstraitTypeParserTs                     = 39
	SubstraitTypeParserTsTZ                   = 40
	SubstraitTypeParserIYear                  = 41
	SubstraitTypeParserIDay                   = 42
	SubstraitTypeParserDec                    = 43
	SubstraitTypeParserPT                     = 44
	SubstraitTypeParserPTs                    = 45
	SubstraitTypeParserPTsTZ                  = 46
	SubstraitTypeParserFChar                  = 47
	SubstraitTypeParserVChar                  = 48
	SubstraitTypeParserFBin                   = 49
	SubstraitTypeParserAny                    = 50
	SubstraitTypeParserAnyVar                 = 51
	SubstraitTypeParserDoubleColon            = 52
	SubstraitTypeParserPlus                   = 53
	SubstraitTypeParserMinus                  = 54
	SubstraitTypeParserAsterisk               = 55
	SubstraitTypeParserForwardSlash           = 56
	SubstraitTypeParserPercent                = 57
	SubstraitTypeParserEq                     = 58
	SubstraitTypeParserNe                     = 59
	SubstraitTypeParserGte                    = 60
	SubstraitTypeParserLte                    = 61
	SubstraitTypeParserGt                     = 62
	SubstraitTypeParserLt                     = 63
	SubstraitTypeParserBang                   = 64
	SubstraitTypeParserOAngleBracket          = 65
	SubstraitTypeParserCAngleBracket          = 66
	SubstraitTypeParserOParen                 = 67
	SubstraitTypeParserCParen                 = 68
	SubstraitTypeParserOBracket               = 69
	SubstraitTypeParserCBracket               = 70
	SubstraitTypeParserComma                  = 71
	SubstraitTypeParserColon                  = 72
	SubstraitTypeParserQMark                  = 73
	SubstraitTypeParserHash                   = 74
	SubstraitTypeParserDot                    = 75
	SubstraitTypeParserAnd                    = 76
	SubstraitTypeParserOr                     = 77
	SubstraitTypeParserAssign                 = 78
	SubstraitTypeParserArrow                  = 79
	SubstraitTypeParserNumber                 = 80
	SubstraitTypeParserIdentifier             = 81
	SubstraitTypeParserNewline                = 82
)

// SubstraitTypeParser rules.
const (
	SubstraitTypeParserRULE_startRule         = 0
	SubstraitTypeParserRULE_typeStatement     = 1
	SubstraitTypeParserRULE_scalarType        = 2
	SubstraitTypeParserRULE_parameterizedType = 3
	SubstraitTypeParserRULE_funcParams        = 4
	SubstraitTypeParserRULE_numericParameter  = 5
	SubstraitTypeParserRULE_anyType           = 6
	SubstraitTypeParserRULE_typeDef           = 7
	SubstraitTypeParserRULE_expr              = 8
)

// IStartRuleContext is an interface to support dynamic dispatch.
type IStartRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	EOF() antlr.TerminalNode

	// IsStartRuleContext differentiates from other interfaces.
	IsStartRuleContext()
}

type StartRuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartRuleContext() *StartRuleContext {
	var p = new(StartRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SubstraitTypeParserRULE_startRule
	return p
}

func InitEmptyStartRuleContext(p *StartRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SubstraitTypeParserRULE_startRule
}

func (*StartRuleContext) IsStartRuleContext() {}

func NewStartRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartRuleContext {
	var p = new(StartRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SubstraitTypeParserRULE_startRule

	return p
}

func (s *StartRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *StartRuleContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StartRuleContext) EOF() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserEOF, 0)
}

func (s *StartRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterStartRule(s)
	}
}

func (s *StartRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitStartRule(s)
	}
}

func (s *StartRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitStartRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SubstraitTypeParser) StartRule() (localctx IStartRuleContext) {
	localctx = NewStartRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SubstraitTypeParserRULE_startRule)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(18)
		p.expr(0)
	}
	{
		p.SetState(19)
		p.Match(SubstraitTypeParserEOF)
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

// ITypeStatementContext is an interface to support dynamic dispatch.
type ITypeStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TypeDef() ITypeDefContext
	EOF() antlr.TerminalNode

	// IsTypeStatementContext differentiates from other interfaces.
	IsTypeStatementContext()
}

type TypeStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeStatementContext() *TypeStatementContext {
	var p = new(TypeStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SubstraitTypeParserRULE_typeStatement
	return p
}

func InitEmptyTypeStatementContext(p *TypeStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SubstraitTypeParserRULE_typeStatement
}

func (*TypeStatementContext) IsTypeStatementContext() {}

func NewTypeStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeStatementContext {
	var p = new(TypeStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SubstraitTypeParserRULE_typeStatement

	return p
}

func (s *TypeStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeStatementContext) TypeDef() ITypeDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDefContext)
}

func (s *TypeStatementContext) EOF() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserEOF, 0)
}

func (s *TypeStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterTypeStatement(s)
	}
}

func (s *TypeStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitTypeStatement(s)
	}
}

func (s *TypeStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitTypeStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SubstraitTypeParser) TypeStatement() (localctx ITypeStatementContext) {
	localctx = NewTypeStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SubstraitTypeParserRULE_typeStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(21)
		p.TypeDef()
	}
	{
		p.SetState(22)
		p.Match(SubstraitTypeParserEOF)
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
	p.RuleIndex = SubstraitTypeParserRULE_scalarType
	return p
}

func InitEmptyScalarTypeContext(p *ScalarTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SubstraitTypeParserRULE_scalarType
}

func (*ScalarTypeContext) IsScalarTypeContext() {}

func NewScalarTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScalarTypeContext {
	var p = new(ScalarTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SubstraitTypeParserRULE_scalarType

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
	return s.GetToken(SubstraitTypeParserDate, 0)
}

func (s *DateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterDate(s)
	}
}

func (s *DateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitDate(s)
	}
}

func (s *DateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
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

func (s *StringContext) String_() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserString_, 0)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitString(s)
	}
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
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
	return s.GetToken(SubstraitTypeParserI64, 0)
}

func (s *I64Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterI64(s)
	}
}

func (s *I64Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitI64(s)
	}
}

func (s *I64Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitI64(s)

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
	return s.GetToken(SubstraitTypeParserI32, 0)
}

func (s *I32Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterI32(s)
	}
}

func (s *I32Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitI32(s)
	}
}

func (s *I32Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
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

func (s *IntervalYearContext) Interval_Year() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserInterval_Year, 0)
}

func (s *IntervalYearContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterIntervalYear(s)
	}
}

func (s *IntervalYearContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitIntervalYear(s)
	}
}

func (s *IntervalYearContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
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
	return s.GetToken(SubstraitTypeParserUUID, 0)
}

func (s *UuidContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterUuid(s)
	}
}

func (s *UuidContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitUuid(s)
	}
}

func (s *UuidContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
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
	return s.GetToken(SubstraitTypeParserI8, 0)
}

func (s *I8Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterI8(s)
	}
}

func (s *I8Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitI8(s)
	}
}

func (s *I8Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
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
	return s.GetToken(SubstraitTypeParserI16, 0)
}

func (s *I16Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterI16(s)
	}
}

func (s *I16Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitI16(s)
	}
}

func (s *I16Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
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

func (s *BooleanContext) Boolean() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserBoolean, 0)
}

func (s *BooleanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterBoolean(s)
	}
}

func (s *BooleanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitBoolean(s)
	}
}

func (s *BooleanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
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

func (s *BinaryContext) Binary() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserBinary, 0)
}

func (s *BinaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterBinary(s)
	}
}

func (s *BinaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitBinary(s)
	}
}

func (s *BinaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
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
	return s.GetToken(SubstraitTypeParserFP64, 0)
}

func (s *Fp64Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterFp64(s)
	}
}

func (s *Fp64Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitFp64(s)
	}
}

func (s *Fp64Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
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
	return s.GetToken(SubstraitTypeParserFP32, 0)
}

func (s *Fp32Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterFp32(s)
	}
}

func (s *Fp32Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitFp32(s)
	}
}

func (s *Fp32Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
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
	return s.GetToken(SubstraitTypeParserTime, 0)
}

func (s *TimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterTime(s)
	}
}

func (s *TimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitTime(s)
	}
}

func (s *TimeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
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

func (s *TimestampContext) Timestamp() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserTimestamp, 0)
}

func (s *TimestampContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterTimestamp(s)
	}
}

func (s *TimestampContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitTimestamp(s)
	}
}

func (s *TimestampContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
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

func (s *TimestampTzContext) Timestamp_TZ() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserTimestamp_TZ, 0)
}

func (s *TimestampTzContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterTimestampTz(s)
	}
}

func (s *TimestampTzContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitTimestampTz(s)
	}
}

func (s *TimestampTzContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitTimestampTz(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SubstraitTypeParser) ScalarType() (localctx IScalarTypeContext) {
	localctx = NewScalarTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SubstraitTypeParserRULE_scalarType)
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SubstraitTypeParserBoolean:
		localctx = NewBooleanContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(24)
			p.Match(SubstraitTypeParserBoolean)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SubstraitTypeParserI8:
		localctx = NewI8Context(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(25)
			p.Match(SubstraitTypeParserI8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SubstraitTypeParserI16:
		localctx = NewI16Context(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(26)
			p.Match(SubstraitTypeParserI16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SubstraitTypeParserI32:
		localctx = NewI32Context(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(27)
			p.Match(SubstraitTypeParserI32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SubstraitTypeParserI64:
		localctx = NewI64Context(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(28)
			p.Match(SubstraitTypeParserI64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SubstraitTypeParserFP32:
		localctx = NewFp32Context(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(29)
			p.Match(SubstraitTypeParserFP32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SubstraitTypeParserFP64:
		localctx = NewFp64Context(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(30)
			p.Match(SubstraitTypeParserFP64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SubstraitTypeParserString_:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(31)
			p.Match(SubstraitTypeParserString_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SubstraitTypeParserBinary:
		localctx = NewBinaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(32)
			p.Match(SubstraitTypeParserBinary)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SubstraitTypeParserTimestamp:
		localctx = NewTimestampContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(33)
			p.Match(SubstraitTypeParserTimestamp)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SubstraitTypeParserTimestamp_TZ:
		localctx = NewTimestampTzContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(34)
			p.Match(SubstraitTypeParserTimestamp_TZ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SubstraitTypeParserDate:
		localctx = NewDateContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(35)
			p.Match(SubstraitTypeParserDate)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SubstraitTypeParserTime:
		localctx = NewTimeContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(36)
			p.Match(SubstraitTypeParserTime)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SubstraitTypeParserInterval_Year:
		localctx = NewIntervalYearContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(37)
			p.Match(SubstraitTypeParserInterval_Year)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SubstraitTypeParserUUID:
		localctx = NewUuidContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(38)
			p.Match(SubstraitTypeParserUUID)
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
	p.RuleIndex = SubstraitTypeParserRULE_parameterizedType
	return p
}

func InitEmptyParameterizedTypeContext(p *ParameterizedTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SubstraitTypeParserRULE_parameterizedType
}

func (*ParameterizedTypeContext) IsParameterizedTypeContext() {}

func NewParameterizedTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterizedTypeContext {
	var p = new(ParameterizedTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SubstraitTypeParserRULE_parameterizedType

	return p
}

func (s *ParameterizedTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterizedTypeContext) CopyAll(ctx *ParameterizedTypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ParameterizedTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterizedTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructContext struct {
	ParameterizedTypeContext
	isnull antlr.Token
}

func NewStructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructContext {
	var p = new(StructContext)

	InitEmptyParameterizedTypeContext(&p.ParameterizedTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParameterizedTypeContext))

	return p
}

func (s *StructContext) GetIsnull() antlr.Token { return s.isnull }

func (s *StructContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *StructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructContext) Struct() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserStruct, 0)
}

func (s *StructContext) Lt() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserLt, 0)
}

func (s *StructContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *StructContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *StructContext) Gt() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserGt, 0)
}

func (s *StructContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(SubstraitTypeParserComma)
}

func (s *StructContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserComma, i)
}

func (s *StructContext) QMark() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserQMark, 0)
}

func (s *StructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterStruct(s)
	}
}

func (s *StructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitStruct(s)
	}
}

func (s *StructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrecisionTimestampTZContext struct {
	ParameterizedTypeContext
	isnull    antlr.Token
	precision INumericParameterContext
}

func NewPrecisionTimestampTZContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrecisionTimestampTZContext {
	var p = new(PrecisionTimestampTZContext)

	InitEmptyParameterizedTypeContext(&p.ParameterizedTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParameterizedTypeContext))

	return p
}

func (s *PrecisionTimestampTZContext) GetIsnull() antlr.Token { return s.isnull }

func (s *PrecisionTimestampTZContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *PrecisionTimestampTZContext) GetPrecision() INumericParameterContext { return s.precision }

func (s *PrecisionTimestampTZContext) SetPrecision(v INumericParameterContext) { s.precision = v }

func (s *PrecisionTimestampTZContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrecisionTimestampTZContext) Precision_Timestamp_TZ() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserPrecision_Timestamp_TZ, 0)
}

func (s *PrecisionTimestampTZContext) Lt() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserLt, 0)
}

func (s *PrecisionTimestampTZContext) Gt() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserGt, 0)
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
	return s.GetToken(SubstraitTypeParserQMark, 0)
}

func (s *PrecisionTimestampTZContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterPrecisionTimestampTZ(s)
	}
}

func (s *PrecisionTimestampTZContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitPrecisionTimestampTZ(s)
	}
}

func (s *PrecisionTimestampTZContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitPrecisionTimestampTZ(s)

	default:
		return t.VisitChildren(s)
	}
}

type NStructContext struct {
	ParameterizedTypeContext
	isnull antlr.Token
}

func NewNStructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NStructContext {
	var p = new(NStructContext)

	InitEmptyParameterizedTypeContext(&p.ParameterizedTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParameterizedTypeContext))

	return p
}

func (s *NStructContext) GetIsnull() antlr.Token { return s.isnull }

func (s *NStructContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *NStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NStructContext) NStruct() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserNStruct, 0)
}

func (s *NStructContext) Lt() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserLt, 0)
}

func (s *NStructContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(SubstraitTypeParserIdentifier)
}

func (s *NStructContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserIdentifier, i)
}

func (s *NStructContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *NStructContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *NStructContext) Gt() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserGt, 0)
}

func (s *NStructContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(SubstraitTypeParserComma)
}

func (s *NStructContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserComma, i)
}

func (s *NStructContext) QMark() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserQMark, 0)
}

func (s *NStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterNStruct(s)
	}
}

func (s *NStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitNStruct(s)
	}
}

func (s *NStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitNStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

type FixedBinaryContext struct {
	ParameterizedTypeContext
	isnull antlr.Token
	length INumericParameterContext
}

func NewFixedBinaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FixedBinaryContext {
	var p = new(FixedBinaryContext)

	InitEmptyParameterizedTypeContext(&p.ParameterizedTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParameterizedTypeContext))

	return p
}

func (s *FixedBinaryContext) GetIsnull() antlr.Token { return s.isnull }

func (s *FixedBinaryContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *FixedBinaryContext) GetLength() INumericParameterContext { return s.length }

func (s *FixedBinaryContext) SetLength(v INumericParameterContext) { s.length = v }

func (s *FixedBinaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FixedBinaryContext) FixedBinary() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserFixedBinary, 0)
}

func (s *FixedBinaryContext) Lt() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserLt, 0)
}

func (s *FixedBinaryContext) Gt() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserGt, 0)
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
	return s.GetToken(SubstraitTypeParserQMark, 0)
}

func (s *FixedBinaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterFixedBinary(s)
	}
}

func (s *FixedBinaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitFixedBinary(s)
	}
}

func (s *FixedBinaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitFixedBinary(s)

	default:
		return t.VisitChildren(s)
	}
}

type UserDefinedContext struct {
	ParameterizedTypeContext
	isnull antlr.Token
}

func NewUserDefinedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UserDefinedContext {
	var p = new(UserDefinedContext)

	InitEmptyParameterizedTypeContext(&p.ParameterizedTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParameterizedTypeContext))

	return p
}

func (s *UserDefinedContext) GetIsnull() antlr.Token { return s.isnull }

func (s *UserDefinedContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *UserDefinedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UserDefinedContext) UserDefined() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserUserDefined, 0)
}

func (s *UserDefinedContext) Identifier() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserIdentifier, 0)
}

func (s *UserDefinedContext) Lt() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserLt, 0)
}

func (s *UserDefinedContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *UserDefinedContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *UserDefinedContext) Gt() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserGt, 0)
}

func (s *UserDefinedContext) QMark() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserQMark, 0)
}

func (s *UserDefinedContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(SubstraitTypeParserComma)
}

func (s *UserDefinedContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserComma, i)
}

func (s *UserDefinedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterUserDefined(s)
	}
}

func (s *UserDefinedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitUserDefined(s)
	}
}

func (s *UserDefinedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitUserDefined(s)

	default:
		return t.VisitChildren(s)
	}
}

type FixedCharContext struct {
	ParameterizedTypeContext
	isnull antlr.Token
	length INumericParameterContext
}

func NewFixedCharContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FixedCharContext {
	var p = new(FixedCharContext)

	InitEmptyParameterizedTypeContext(&p.ParameterizedTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParameterizedTypeContext))

	return p
}

func (s *FixedCharContext) GetIsnull() antlr.Token { return s.isnull }

func (s *FixedCharContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *FixedCharContext) GetLength() INumericParameterContext { return s.length }

func (s *FixedCharContext) SetLength(v INumericParameterContext) { s.length = v }

func (s *FixedCharContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FixedCharContext) FixedChar() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserFixedChar, 0)
}

func (s *FixedCharContext) Lt() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserLt, 0)
}

func (s *FixedCharContext) Gt() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserGt, 0)
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
	return s.GetToken(SubstraitTypeParserQMark, 0)
}

func (s *FixedCharContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterFixedChar(s)
	}
}

func (s *FixedCharContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitFixedChar(s)
	}
}

func (s *FixedCharContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitFixedChar(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListContext struct {
	ParameterizedTypeContext
	isnull antlr.Token
}

func NewListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListContext {
	var p = new(ListContext)

	InitEmptyParameterizedTypeContext(&p.ParameterizedTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParameterizedTypeContext))

	return p
}

func (s *ListContext) GetIsnull() antlr.Token { return s.isnull }

func (s *ListContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) List() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserList, 0)
}

func (s *ListContext) Lt() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserLt, 0)
}

func (s *ListContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ListContext) Gt() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserGt, 0)
}

func (s *ListContext) QMark() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserQMark, 0)
}

func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitList(s)
	}
}

func (s *ListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitList(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrecisionIntervalDayContext struct {
	ParameterizedTypeContext
	isnull    antlr.Token
	precision INumericParameterContext
}

func NewPrecisionIntervalDayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrecisionIntervalDayContext {
	var p = new(PrecisionIntervalDayContext)

	InitEmptyParameterizedTypeContext(&p.ParameterizedTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParameterizedTypeContext))

	return p
}

func (s *PrecisionIntervalDayContext) GetIsnull() antlr.Token { return s.isnull }

func (s *PrecisionIntervalDayContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *PrecisionIntervalDayContext) GetPrecision() INumericParameterContext { return s.precision }

func (s *PrecisionIntervalDayContext) SetPrecision(v INumericParameterContext) { s.precision = v }

func (s *PrecisionIntervalDayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrecisionIntervalDayContext) Interval_Day() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserInterval_Day, 0)
}

func (s *PrecisionIntervalDayContext) Lt() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserLt, 0)
}

func (s *PrecisionIntervalDayContext) Gt() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserGt, 0)
}

func (s *PrecisionIntervalDayContext) NumericParameter() INumericParameterContext {
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

func (s *PrecisionIntervalDayContext) QMark() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserQMark, 0)
}

func (s *PrecisionIntervalDayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterPrecisionIntervalDay(s)
	}
}

func (s *PrecisionIntervalDayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitPrecisionIntervalDay(s)
	}
}

func (s *PrecisionIntervalDayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitPrecisionIntervalDay(s)

	default:
		return t.VisitChildren(s)
	}
}

type FuncContext struct {
	ParameterizedTypeContext
	isnull     antlr.Token
	params     IFuncParamsContext
	returnType IExprContext
}

func NewFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncContext {
	var p = new(FuncContext)

	InitEmptyParameterizedTypeContext(&p.ParameterizedTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParameterizedTypeContext))

	return p
}

func (s *FuncContext) GetIsnull() antlr.Token { return s.isnull }

func (s *FuncContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *FuncContext) GetParams() IFuncParamsContext { return s.params }

func (s *FuncContext) GetReturnType() IExprContext { return s.returnType }

func (s *FuncContext) SetParams(v IFuncParamsContext) { s.params = v }

func (s *FuncContext) SetReturnType(v IExprContext) { s.returnType = v }

func (s *FuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncContext) Func() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserFunc, 0)
}

func (s *FuncContext) Lt() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserLt, 0)
}

func (s *FuncContext) Arrow() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserArrow, 0)
}

func (s *FuncContext) Gt() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserGt, 0)
}

func (s *FuncContext) FuncParams() IFuncParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncParamsContext)
}

func (s *FuncContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FuncContext) QMark() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserQMark, 0)
}

func (s *FuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterFunc(s)
	}
}

func (s *FuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitFunc(s)
	}
}

func (s *FuncContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitFunc(s)

	default:
		return t.VisitChildren(s)
	}
}

type VarCharContext struct {
	ParameterizedTypeContext
	isnull antlr.Token
	length INumericParameterContext
}

func NewVarCharContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarCharContext {
	var p = new(VarCharContext)

	InitEmptyParameterizedTypeContext(&p.ParameterizedTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParameterizedTypeContext))

	return p
}

func (s *VarCharContext) GetIsnull() antlr.Token { return s.isnull }

func (s *VarCharContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *VarCharContext) GetLength() INumericParameterContext { return s.length }

func (s *VarCharContext) SetLength(v INumericParameterContext) { s.length = v }

func (s *VarCharContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarCharContext) VarChar() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserVarChar, 0)
}

func (s *VarCharContext) Lt() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserLt, 0)
}

func (s *VarCharContext) Gt() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserGt, 0)
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
	return s.GetToken(SubstraitTypeParserQMark, 0)
}

func (s *VarCharContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterVarChar(s)
	}
}

func (s *VarCharContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitVarChar(s)
	}
}

func (s *VarCharContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitVarChar(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrecisionTimestampContext struct {
	ParameterizedTypeContext
	isnull    antlr.Token
	precision INumericParameterContext
}

func NewPrecisionTimestampContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrecisionTimestampContext {
	var p = new(PrecisionTimestampContext)

	InitEmptyParameterizedTypeContext(&p.ParameterizedTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParameterizedTypeContext))

	return p
}

func (s *PrecisionTimestampContext) GetIsnull() antlr.Token { return s.isnull }

func (s *PrecisionTimestampContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *PrecisionTimestampContext) GetPrecision() INumericParameterContext { return s.precision }

func (s *PrecisionTimestampContext) SetPrecision(v INumericParameterContext) { s.precision = v }

func (s *PrecisionTimestampContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrecisionTimestampContext) Precision_Timestamp() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserPrecision_Timestamp, 0)
}

func (s *PrecisionTimestampContext) Lt() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserLt, 0)
}

func (s *PrecisionTimestampContext) Gt() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserGt, 0)
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
	return s.GetToken(SubstraitTypeParserQMark, 0)
}

func (s *PrecisionTimestampContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterPrecisionTimestamp(s)
	}
}

func (s *PrecisionTimestampContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitPrecisionTimestamp(s)
	}
}

func (s *PrecisionTimestampContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitPrecisionTimestamp(s)

	default:
		return t.VisitChildren(s)
	}
}

type DecimalContext struct {
	ParameterizedTypeContext
	isnull    antlr.Token
	precision INumericParameterContext
	scale     INumericParameterContext
}

func NewDecimalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DecimalContext {
	var p = new(DecimalContext)

	InitEmptyParameterizedTypeContext(&p.ParameterizedTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParameterizedTypeContext))

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

func (s *DecimalContext) Decimal() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserDecimal, 0)
}

func (s *DecimalContext) Lt() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserLt, 0)
}

func (s *DecimalContext) Comma() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserComma, 0)
}

func (s *DecimalContext) Gt() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserGt, 0)
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

func (s *DecimalContext) QMark() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserQMark, 0)
}

func (s *DecimalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterDecimal(s)
	}
}

func (s *DecimalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitDecimal(s)
	}
}

func (s *DecimalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitDecimal(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrecisionTimeContext struct {
	ParameterizedTypeContext
	isnull    antlr.Token
	precision INumericParameterContext
}

func NewPrecisionTimeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrecisionTimeContext {
	var p = new(PrecisionTimeContext)

	InitEmptyParameterizedTypeContext(&p.ParameterizedTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParameterizedTypeContext))

	return p
}

func (s *PrecisionTimeContext) GetIsnull() antlr.Token { return s.isnull }

func (s *PrecisionTimeContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *PrecisionTimeContext) GetPrecision() INumericParameterContext { return s.precision }

func (s *PrecisionTimeContext) SetPrecision(v INumericParameterContext) { s.precision = v }

func (s *PrecisionTimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrecisionTimeContext) Precision_Time() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserPrecision_Time, 0)
}

func (s *PrecisionTimeContext) Lt() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserLt, 0)
}

func (s *PrecisionTimeContext) Gt() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserGt, 0)
}

func (s *PrecisionTimeContext) NumericParameter() INumericParameterContext {
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

func (s *PrecisionTimeContext) QMark() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserQMark, 0)
}

func (s *PrecisionTimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterPrecisionTime(s)
	}
}

func (s *PrecisionTimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitPrecisionTime(s)
	}
}

func (s *PrecisionTimeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitPrecisionTime(s)

	default:
		return t.VisitChildren(s)
	}
}

type MapContext struct {
	ParameterizedTypeContext
	isnull antlr.Token
	key    IExprContext
	value  IExprContext
}

func NewMapContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MapContext {
	var p = new(MapContext)

	InitEmptyParameterizedTypeContext(&p.ParameterizedTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParameterizedTypeContext))

	return p
}

func (s *MapContext) GetIsnull() antlr.Token { return s.isnull }

func (s *MapContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *MapContext) GetKey() IExprContext { return s.key }

func (s *MapContext) GetValue() IExprContext { return s.value }

func (s *MapContext) SetKey(v IExprContext) { s.key = v }

func (s *MapContext) SetValue(v IExprContext) { s.value = v }

func (s *MapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapContext) Map() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserMap, 0)
}

func (s *MapContext) Lt() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserLt, 0)
}

func (s *MapContext) Comma() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserComma, 0)
}

func (s *MapContext) Gt() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserGt, 0)
}

func (s *MapContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *MapContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *MapContext) QMark() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserQMark, 0)
}

func (s *MapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterMap(s)
	}
}

func (s *MapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitMap(s)
	}
}

func (s *MapContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitMap(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SubstraitTypeParser) ParameterizedType() (localctx IParameterizedTypeContext) {
	localctx = NewParameterizedTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SubstraitTypeParserRULE_parameterizedType)
	var _la int

	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SubstraitTypeParserFixedChar:
		localctx = NewFixedCharContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(41)
			p.Match(SubstraitTypeParserFixedChar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SubstraitTypeParserQMark {
			{
				p.SetState(42)

				var _m = p.Match(SubstraitTypeParserQMark)

				localctx.(*FixedCharContext).isnull = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(45)
			p.Match(SubstraitTypeParserLt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(46)

			var _x = p.NumericParameter()

			localctx.(*FixedCharContext).length = _x
		}
		{
			p.SetState(47)
			p.Match(SubstraitTypeParserGt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SubstraitTypeParserVarChar:
		localctx = NewVarCharContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(49)
			p.Match(SubstraitTypeParserVarChar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SubstraitTypeParserQMark {
			{
				p.SetState(50)

				var _m = p.Match(SubstraitTypeParserQMark)

				localctx.(*VarCharContext).isnull = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(53)
			p.Match(SubstraitTypeParserLt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(54)

			var _x = p.NumericParameter()

			localctx.(*VarCharContext).length = _x
		}
		{
			p.SetState(55)
			p.Match(SubstraitTypeParserGt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SubstraitTypeParserFixedBinary:
		localctx = NewFixedBinaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(57)
			p.Match(SubstraitTypeParserFixedBinary)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SubstraitTypeParserQMark {
			{
				p.SetState(58)

				var _m = p.Match(SubstraitTypeParserQMark)

				localctx.(*FixedBinaryContext).isnull = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(61)
			p.Match(SubstraitTypeParserLt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(62)

			var _x = p.NumericParameter()

			localctx.(*FixedBinaryContext).length = _x
		}
		{
			p.SetState(63)
			p.Match(SubstraitTypeParserGt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SubstraitTypeParserDecimal:
		localctx = NewDecimalContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(65)
			p.Match(SubstraitTypeParserDecimal)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SubstraitTypeParserQMark {
			{
				p.SetState(66)

				var _m = p.Match(SubstraitTypeParserQMark)

				localctx.(*DecimalContext).isnull = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(69)
			p.Match(SubstraitTypeParserLt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)

			var _x = p.NumericParameter()

			localctx.(*DecimalContext).precision = _x
		}
		{
			p.SetState(71)
			p.Match(SubstraitTypeParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(72)

			var _x = p.NumericParameter()

			localctx.(*DecimalContext).scale = _x
		}
		{
			p.SetState(73)
			p.Match(SubstraitTypeParserGt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SubstraitTypeParserInterval_Day:
		localctx = NewPrecisionIntervalDayContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(75)
			p.Match(SubstraitTypeParserInterval_Day)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SubstraitTypeParserQMark {
			{
				p.SetState(76)

				var _m = p.Match(SubstraitTypeParserQMark)

				localctx.(*PrecisionIntervalDayContext).isnull = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(79)
			p.Match(SubstraitTypeParserLt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(80)

			var _x = p.NumericParameter()

			localctx.(*PrecisionIntervalDayContext).precision = _x
		}
		{
			p.SetState(81)
			p.Match(SubstraitTypeParserGt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SubstraitTypeParserPrecision_Time:
		localctx = NewPrecisionTimeContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(83)
			p.Match(SubstraitTypeParserPrecision_Time)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SubstraitTypeParserQMark {
			{
				p.SetState(84)

				var _m = p.Match(SubstraitTypeParserQMark)

				localctx.(*PrecisionTimeContext).isnull = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(87)
			p.Match(SubstraitTypeParserLt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(88)

			var _x = p.NumericParameter()

			localctx.(*PrecisionTimeContext).precision = _x
		}
		{
			p.SetState(89)
			p.Match(SubstraitTypeParserGt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SubstraitTypeParserPrecision_Timestamp:
		localctx = NewPrecisionTimestampContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(91)
			p.Match(SubstraitTypeParserPrecision_Timestamp)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SubstraitTypeParserQMark {
			{
				p.SetState(92)

				var _m = p.Match(SubstraitTypeParserQMark)

				localctx.(*PrecisionTimestampContext).isnull = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(95)
			p.Match(SubstraitTypeParserLt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(96)

			var _x = p.NumericParameter()

			localctx.(*PrecisionTimestampContext).precision = _x
		}
		{
			p.SetState(97)
			p.Match(SubstraitTypeParserGt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SubstraitTypeParserPrecision_Timestamp_TZ:
		localctx = NewPrecisionTimestampTZContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(99)
			p.Match(SubstraitTypeParserPrecision_Timestamp_TZ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SubstraitTypeParserQMark {
			{
				p.SetState(100)

				var _m = p.Match(SubstraitTypeParserQMark)

				localctx.(*PrecisionTimestampTZContext).isnull = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(103)
			p.Match(SubstraitTypeParserLt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(104)

			var _x = p.NumericParameter()

			localctx.(*PrecisionTimestampTZContext).precision = _x
		}
		{
			p.SetState(105)
			p.Match(SubstraitTypeParserGt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SubstraitTypeParserStruct:
		localctx = NewStructContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(107)
			p.Match(SubstraitTypeParserStruct)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SubstraitTypeParserQMark {
			{
				p.SetState(108)

				var _m = p.Match(SubstraitTypeParserQMark)

				localctx.(*StructContext).isnull = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(111)
			p.Match(SubstraitTypeParserLt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(112)
			p.expr(0)
		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SubstraitTypeParserComma {
			{
				p.SetState(113)
				p.Match(SubstraitTypeParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(114)
				p.expr(0)
			}

			p.SetState(119)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(120)
			p.Match(SubstraitTypeParserGt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SubstraitTypeParserNStruct:
		localctx = NewNStructContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(122)
			p.Match(SubstraitTypeParserNStruct)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SubstraitTypeParserQMark {
			{
				p.SetState(123)

				var _m = p.Match(SubstraitTypeParserQMark)

				localctx.(*NStructContext).isnull = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(126)
			p.Match(SubstraitTypeParserLt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(127)
			p.Match(SubstraitTypeParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.expr(0)
		}
		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SubstraitTypeParserComma {
			{
				p.SetState(129)
				p.Match(SubstraitTypeParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(130)
				p.Match(SubstraitTypeParserIdentifier)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(131)
				p.expr(0)
			}

			p.SetState(136)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(137)
			p.Match(SubstraitTypeParserGt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SubstraitTypeParserList:
		localctx = NewListContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(139)
			p.Match(SubstraitTypeParserList)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SubstraitTypeParserQMark {
			{
				p.SetState(140)

				var _m = p.Match(SubstraitTypeParserQMark)

				localctx.(*ListContext).isnull = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(143)
			p.Match(SubstraitTypeParserLt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(144)
			p.expr(0)
		}
		{
			p.SetState(145)
			p.Match(SubstraitTypeParserGt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SubstraitTypeParserMap:
		localctx = NewMapContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(147)
			p.Match(SubstraitTypeParserMap)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SubstraitTypeParserQMark {
			{
				p.SetState(148)

				var _m = p.Match(SubstraitTypeParserQMark)

				localctx.(*MapContext).isnull = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(151)
			p.Match(SubstraitTypeParserLt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(152)

			var _x = p.expr(0)

			localctx.(*MapContext).key = _x
		}
		{
			p.SetState(153)
			p.Match(SubstraitTypeParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)

			var _x = p.expr(0)

			localctx.(*MapContext).value = _x
		}
		{
			p.SetState(155)
			p.Match(SubstraitTypeParserGt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SubstraitTypeParserFunc:
		localctx = NewFuncContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(157)
			p.Match(SubstraitTypeParserFunc)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SubstraitTypeParserQMark {
			{
				p.SetState(158)

				var _m = p.Match(SubstraitTypeParserQMark)

				localctx.(*FuncContext).isnull = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(161)
			p.Match(SubstraitTypeParserLt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(162)

			var _x = p.FuncParams()

			localctx.(*FuncContext).params = _x
		}
		{
			p.SetState(163)
			p.Match(SubstraitTypeParserArrow)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(164)

			var _x = p.expr(0)

			localctx.(*FuncContext).returnType = _x
		}
		{
			p.SetState(165)
			p.Match(SubstraitTypeParserGt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SubstraitTypeParserUserDefined:
		localctx = NewUserDefinedContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(167)
			p.Match(SubstraitTypeParserUserDefined)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(168)
			p.Match(SubstraitTypeParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(170)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(169)

				var _m = p.Match(SubstraitTypeParserQMark)

				localctx.(*UserDefinedContext).isnull = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(183)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(172)
				p.Match(SubstraitTypeParserLt)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(173)
				p.expr(0)
			}
			p.SetState(178)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == SubstraitTypeParserComma {
				{
					p.SetState(174)
					p.Match(SubstraitTypeParserComma)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(175)
					p.expr(0)
				}

				p.SetState(180)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(181)
				p.Match(SubstraitTypeParserGt)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
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

// IFuncParamsContext is an interface to support dynamic dispatch.
type IFuncParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFuncParamsContext differentiates from other interfaces.
	IsFuncParamsContext()
}

type FuncParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncParamsContext() *FuncParamsContext {
	var p = new(FuncParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SubstraitTypeParserRULE_funcParams
	return p
}

func InitEmptyFuncParamsContext(p *FuncParamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SubstraitTypeParserRULE_funcParams
}

func (*FuncParamsContext) IsFuncParamsContext() {}

func NewFuncParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncParamsContext {
	var p = new(FuncParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SubstraitTypeParserRULE_funcParams

	return p
}

func (s *FuncParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncParamsContext) CopyAll(ctx *FuncParamsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FuncParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SingleFuncParamContext struct {
	FuncParamsContext
}

func NewSingleFuncParamContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleFuncParamContext {
	var p = new(SingleFuncParamContext)

	InitEmptyFuncParamsContext(&p.FuncParamsContext)
	p.parser = parser
	p.CopyAll(ctx.(*FuncParamsContext))

	return p
}

func (s *SingleFuncParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleFuncParamContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SingleFuncParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterSingleFuncParam(s)
	}
}

func (s *SingleFuncParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitSingleFuncParam(s)
	}
}

func (s *SingleFuncParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitSingleFuncParam(s)

	default:
		return t.VisitChildren(s)
	}
}

type FuncParamsWithParensContext struct {
	FuncParamsContext
}

func NewFuncParamsWithParensContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncParamsWithParensContext {
	var p = new(FuncParamsWithParensContext)

	InitEmptyFuncParamsContext(&p.FuncParamsContext)
	p.parser = parser
	p.CopyAll(ctx.(*FuncParamsContext))

	return p
}

func (s *FuncParamsWithParensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncParamsWithParensContext) OParen() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserOParen, 0)
}

func (s *FuncParamsWithParensContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *FuncParamsWithParensContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *FuncParamsWithParensContext) CParen() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserCParen, 0)
}

func (s *FuncParamsWithParensContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(SubstraitTypeParserComma)
}

func (s *FuncParamsWithParensContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserComma, i)
}

func (s *FuncParamsWithParensContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterFuncParamsWithParens(s)
	}
}

func (s *FuncParamsWithParensContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitFuncParamsWithParens(s)
	}
}

func (s *FuncParamsWithParensContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitFuncParamsWithParens(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SubstraitTypeParser) FuncParams() (localctx IFuncParamsContext) {
	localctx = NewFuncParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SubstraitTypeParserRULE_funcParams)
	var _la int

	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSingleFuncParamContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(187)
			p.expr(0)
		}

	case 2:
		localctx = NewFuncParamsWithParensContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(188)
			p.Match(SubstraitTypeParserOParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(189)
			p.expr(0)
		}
		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SubstraitTypeParserComma {
			{
				p.SetState(190)
				p.Match(SubstraitTypeParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(191)
				p.expr(0)
			}

			p.SetState(196)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(197)
			p.Match(SubstraitTypeParserCParen)
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
	p.RuleIndex = SubstraitTypeParserRULE_numericParameter
	return p
}

func InitEmptyNumericParameterContext(p *NumericParameterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SubstraitTypeParserRULE_numericParameter
}

func (*NumericParameterContext) IsNumericParameterContext() {}

func NewNumericParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericParameterContext {
	var p = new(NumericParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SubstraitTypeParserRULE_numericParameter

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

type NumericParameterNameContext struct {
	NumericParameterContext
}

func NewNumericParameterNameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumericParameterNameContext {
	var p = new(NumericParameterNameContext)

	InitEmptyNumericParameterContext(&p.NumericParameterContext)
	p.parser = parser
	p.CopyAll(ctx.(*NumericParameterContext))

	return p
}

func (s *NumericParameterNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericParameterNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserIdentifier, 0)
}

func (s *NumericParameterNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterNumericParameterName(s)
	}
}

func (s *NumericParameterNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitNumericParameterName(s)
	}
}

func (s *NumericParameterNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitNumericParameterName(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumericLiteralContext struct {
	NumericParameterContext
}

func NewNumericLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumericLiteralContext {
	var p = new(NumericLiteralContext)

	InitEmptyNumericParameterContext(&p.NumericParameterContext)
	p.parser = parser
	p.CopyAll(ctx.(*NumericParameterContext))

	return p
}

func (s *NumericLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericLiteralContext) Number() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserNumber, 0)
}

func (s *NumericLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterNumericLiteral(s)
	}
}

func (s *NumericLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitNumericLiteral(s)
	}
}

func (s *NumericLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitNumericLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumericExpressionContext struct {
	NumericParameterContext
}

func NewNumericExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumericExpressionContext {
	var p = new(NumericExpressionContext)

	InitEmptyNumericParameterContext(&p.NumericParameterContext)
	p.parser = parser
	p.CopyAll(ctx.(*NumericParameterContext))

	return p
}

func (s *NumericExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericExpressionContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NumericExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterNumericExpression(s)
	}
}

func (s *NumericExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitNumericExpression(s)
	}
}

func (s *NumericExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitNumericExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SubstraitTypeParser) NumericParameter() (localctx INumericParameterContext) {
	localctx = NewNumericParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SubstraitTypeParserRULE_numericParameter)
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNumericLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(201)
			p.Match(SubstraitTypeParserNumber)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewNumericParameterNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(202)
			p.Match(SubstraitTypeParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewNumericExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(203)
			p.expr(0)
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

// IAnyTypeContext is an interface to support dynamic dispatch.
type IAnyTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIsnull returns the isnull token.
	GetIsnull() antlr.Token

	// SetIsnull sets the isnull token.
	SetIsnull(antlr.Token)

	// Getter signatures
	Any() antlr.TerminalNode
	QMark() antlr.TerminalNode
	AnyVar() antlr.TerminalNode

	// IsAnyTypeContext differentiates from other interfaces.
	IsAnyTypeContext()
}

type AnyTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
}

func NewEmptyAnyTypeContext() *AnyTypeContext {
	var p = new(AnyTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SubstraitTypeParserRULE_anyType
	return p
}

func InitEmptyAnyTypeContext(p *AnyTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SubstraitTypeParserRULE_anyType
}

func (*AnyTypeContext) IsAnyTypeContext() {}

func NewAnyTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnyTypeContext {
	var p = new(AnyTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SubstraitTypeParserRULE_anyType

	return p
}

func (s *AnyTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *AnyTypeContext) GetIsnull() antlr.Token { return s.isnull }

func (s *AnyTypeContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *AnyTypeContext) Any() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserAny, 0)
}

func (s *AnyTypeContext) QMark() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserQMark, 0)
}

func (s *AnyTypeContext) AnyVar() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserAnyVar, 0)
}

func (s *AnyTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnyTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnyTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterAnyType(s)
	}
}

func (s *AnyTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitAnyType(s)
	}
}

func (s *AnyTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitAnyType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SubstraitTypeParser) AnyType() (localctx IAnyTypeContext) {
	localctx = NewAnyTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SubstraitTypeParserRULE_anyType)
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SubstraitTypeParserAny:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(206)
			p.Match(SubstraitTypeParserAny)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(208)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(207)

				var _m = p.Match(SubstraitTypeParserQMark)

				localctx.(*AnyTypeContext).isnull = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case SubstraitTypeParserAnyVar:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(210)
			p.Match(SubstraitTypeParserAnyVar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(212)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(211)

				var _m = p.Match(SubstraitTypeParserQMark)

				localctx.(*AnyTypeContext).isnull = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
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

// ITypeDefContext is an interface to support dynamic dispatch.
type ITypeDefContext interface {
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
	AnyType() IAnyTypeContext

	// IsTypeDefContext differentiates from other interfaces.
	IsTypeDefContext()
}

type TypeDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	isnull antlr.Token
}

func NewEmptyTypeDefContext() *TypeDefContext {
	var p = new(TypeDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SubstraitTypeParserRULE_typeDef
	return p
}

func InitEmptyTypeDefContext(p *TypeDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SubstraitTypeParserRULE_typeDef
}

func (*TypeDefContext) IsTypeDefContext() {}

func NewTypeDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDefContext {
	var p = new(TypeDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SubstraitTypeParserRULE_typeDef

	return p
}

func (s *TypeDefContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDefContext) GetIsnull() antlr.Token { return s.isnull }

func (s *TypeDefContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *TypeDefContext) ScalarType() IScalarTypeContext {
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

func (s *TypeDefContext) QMark() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserQMark, 0)
}

func (s *TypeDefContext) ParameterizedType() IParameterizedTypeContext {
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

func (s *TypeDefContext) AnyType() IAnyTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyTypeContext)
}

func (s *TypeDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterTypeDef(s)
	}
}

func (s *TypeDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitTypeDef(s)
	}
}

func (s *TypeDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitTypeDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SubstraitTypeParser) TypeDef() (localctx ITypeDefContext) {
	localctx = NewTypeDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SubstraitTypeParserRULE_typeDef)
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SubstraitTypeParserBoolean, SubstraitTypeParserI8, SubstraitTypeParserI16, SubstraitTypeParserI32, SubstraitTypeParserI64, SubstraitTypeParserFP32, SubstraitTypeParserFP64, SubstraitTypeParserString_, SubstraitTypeParserBinary, SubstraitTypeParserTimestamp, SubstraitTypeParserTimestamp_TZ, SubstraitTypeParserDate, SubstraitTypeParserTime, SubstraitTypeParserInterval_Year, SubstraitTypeParserUUID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(216)
			p.ScalarType()
		}
		p.SetState(218)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(217)

				var _m = p.Match(SubstraitTypeParserQMark)

				localctx.(*TypeDefContext).isnull = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case SubstraitTypeParserFunc, SubstraitTypeParserInterval_Day, SubstraitTypeParserDecimal, SubstraitTypeParserPrecision_Time, SubstraitTypeParserPrecision_Timestamp, SubstraitTypeParserPrecision_Timestamp_TZ, SubstraitTypeParserFixedChar, SubstraitTypeParserVarChar, SubstraitTypeParserFixedBinary, SubstraitTypeParserStruct, SubstraitTypeParserNStruct, SubstraitTypeParserList, SubstraitTypeParserMap, SubstraitTypeParserUserDefined:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(220)
			p.ParameterizedType()
		}

	case SubstraitTypeParserAny, SubstraitTypeParserAnyVar:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(221)
			p.AnyType()
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

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SubstraitTypeParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SubstraitTypeParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SubstraitTypeParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IfExprContext struct {
	ExprContext
	ifExpr   IExprContext
	thenExpr IExprContext
	elseExpr IExprContext
}

func NewIfExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfExprContext {
	var p = new(IfExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IfExprContext) GetIfExpr() IExprContext { return s.ifExpr }

func (s *IfExprContext) GetThenExpr() IExprContext { return s.thenExpr }

func (s *IfExprContext) GetElseExpr() IExprContext { return s.elseExpr }

func (s *IfExprContext) SetIfExpr(v IExprContext) { s.ifExpr = v }

func (s *IfExprContext) SetThenExpr(v IExprContext) { s.thenExpr = v }

func (s *IfExprContext) SetElseExpr(v IExprContext) { s.elseExpr = v }

func (s *IfExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfExprContext) If() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserIf, 0)
}

func (s *IfExprContext) Then() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserThen, 0)
}

func (s *IfExprContext) Else() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserElse, 0)
}

func (s *IfExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *IfExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *IfExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterIfExpr(s)
	}
}

func (s *IfExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitIfExpr(s)
	}
}

func (s *IfExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitIfExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeLiteralContext struct {
	ExprContext
}

func NewTypeLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeLiteralContext {
	var p = new(TypeLiteralContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *TypeLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeLiteralContext) TypeDef() ITypeDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDefContext)
}

func (s *TypeLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterTypeLiteral(s)
	}
}

func (s *TypeLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitTypeLiteral(s)
	}
}

func (s *TypeLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitTypeLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type MultilineDefinitionContext struct {
	ExprContext
	finalType ITypeDefContext
}

func NewMultilineDefinitionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultilineDefinitionContext {
	var p = new(MultilineDefinitionContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MultilineDefinitionContext) GetFinalType() ITypeDefContext { return s.finalType }

func (s *MultilineDefinitionContext) SetFinalType(v ITypeDefContext) { s.finalType = v }

func (s *MultilineDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultilineDefinitionContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(SubstraitTypeParserIdentifier)
}

func (s *MultilineDefinitionContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserIdentifier, i)
}

func (s *MultilineDefinitionContext) AllEq() []antlr.TerminalNode {
	return s.GetTokens(SubstraitTypeParserEq)
}

func (s *MultilineDefinitionContext) Eq(i int) antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserEq, i)
}

func (s *MultilineDefinitionContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *MultilineDefinitionContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *MultilineDefinitionContext) TypeDef() ITypeDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDefContext)
}

func (s *MultilineDefinitionContext) AllNewline() []antlr.TerminalNode {
	return s.GetTokens(SubstraitTypeParserNewline)
}

func (s *MultilineDefinitionContext) Newline(i int) antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserNewline, i)
}

func (s *MultilineDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterMultilineDefinition(s)
	}
}

func (s *MultilineDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitMultilineDefinition(s)
	}
}

func (s *MultilineDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitMultilineDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

type TernaryContext struct {
	ExprContext
	ifExpr   IExprContext
	thenExpr IExprContext
	elseExpr IExprContext
}

func NewTernaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TernaryContext {
	var p = new(TernaryContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *TernaryContext) GetIfExpr() IExprContext { return s.ifExpr }

func (s *TernaryContext) GetThenExpr() IExprContext { return s.thenExpr }

func (s *TernaryContext) GetElseExpr() IExprContext { return s.elseExpr }

func (s *TernaryContext) SetIfExpr(v IExprContext) { s.ifExpr = v }

func (s *TernaryContext) SetThenExpr(v IExprContext) { s.thenExpr = v }

func (s *TernaryContext) SetElseExpr(v IExprContext) { s.elseExpr = v }

func (s *TernaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TernaryContext) QMark() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserQMark, 0)
}

func (s *TernaryContext) Colon() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserColon, 0)
}

func (s *TernaryContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *TernaryContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *TernaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterTernary(s)
	}
}

func (s *TernaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitTernary(s)
	}
}

func (s *TernaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitTernary(s)

	default:
		return t.VisitChildren(s)
	}
}

type BinaryExprContext struct {
	ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewBinaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryExprContext {
	var p = new(BinaryExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BinaryExprContext) GetOp() antlr.Token { return s.op }

func (s *BinaryExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *BinaryExprContext) GetLeft() IExprContext { return s.left }

func (s *BinaryExprContext) GetRight() IExprContext { return s.right }

func (s *BinaryExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *BinaryExprContext) SetRight(v IExprContext) { s.right = v }

func (s *BinaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *BinaryExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *BinaryExprContext) And() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserAnd, 0)
}

func (s *BinaryExprContext) Or() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserOr, 0)
}

func (s *BinaryExprContext) Plus() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserPlus, 0)
}

func (s *BinaryExprContext) Minus() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserMinus, 0)
}

func (s *BinaryExprContext) Lt() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserLt, 0)
}

func (s *BinaryExprContext) Gt() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserGt, 0)
}

func (s *BinaryExprContext) Eq() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserEq, 0)
}

func (s *BinaryExprContext) Ne() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserNe, 0)
}

func (s *BinaryExprContext) Lte() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserLte, 0)
}

func (s *BinaryExprContext) Gte() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserGte, 0)
}

func (s *BinaryExprContext) Asterisk() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserAsterisk, 0)
}

func (s *BinaryExprContext) ForwardSlash() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserForwardSlash, 0)
}

func (s *BinaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterBinaryExpr(s)
	}
}

func (s *BinaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitBinaryExpr(s)
	}
}

func (s *BinaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitBinaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExpressionContext struct {
	ExprContext
}

func NewParenExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExpressionContext {
	var p = new(ParenExpressionContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParenExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExpressionContext) OParen() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserOParen, 0)
}

func (s *ParenExpressionContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParenExpressionContext) CParen() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserCParen, 0)
}

func (s *ParenExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterParenExpression(s)
	}
}

func (s *ParenExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitParenExpression(s)
	}
}

func (s *ParenExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitParenExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParameterNameContext struct {
	ExprContext
	isnull antlr.Token
}

func NewParameterNameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParameterNameContext {
	var p = new(ParameterNameContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParameterNameContext) GetIsnull() antlr.Token { return s.isnull }

func (s *ParameterNameContext) SetIsnull(v antlr.Token) { s.isnull = v }

func (s *ParameterNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserIdentifier, 0)
}

func (s *ParameterNameContext) QMark() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserQMark, 0)
}

func (s *ParameterNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterParameterName(s)
	}
}

func (s *ParameterNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitParameterName(s)
	}
}

func (s *ParameterNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitParameterName(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionCallContext struct {
	ExprContext
}

func NewFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallContext {
	var p = new(FunctionCallContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) Identifier() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserIdentifier, 0)
}

func (s *FunctionCallContext) OParen() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserOParen, 0)
}

func (s *FunctionCallContext) CParen() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserCParen, 0)
}

func (s *FunctionCallContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *FunctionCallContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *FunctionCallContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(SubstraitTypeParserComma)
}

func (s *FunctionCallContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserComma, i)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotExprContext struct {
	ExprContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotExprContext) Bang() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserBang, 0)
}

func (s *NotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterNotExpr(s)
	}
}

func (s *NotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitNotExpr(s)
	}
}

func (s *NotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitNotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiteralNumberContext struct {
	ExprContext
}

func NewLiteralNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralNumberContext {
	var p = new(LiteralNumberContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *LiteralNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralNumberContext) Number() antlr.TerminalNode {
	return s.GetToken(SubstraitTypeParserNumber, 0)
}

func (s *LiteralNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.EnterLiteralNumber(s)
	}
}

func (s *LiteralNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstraitTypeListener); ok {
		listenerT.ExitLiteralNumber(s)
	}
}

func (s *LiteralNumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SubstraitTypeVisitor:
		return t.VisitLiteralNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SubstraitTypeParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *SubstraitTypeParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 16
	p.EnterRecursionRule(localctx, 16, SubstraitTypeParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParenExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(225)
			p.Match(SubstraitTypeParserOParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(226)
			p.expr(0)
		}
		{
			p.SetState(227)
			p.Match(SubstraitTypeParserCParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewMultilineDefinitionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(229)
			p.Match(SubstraitTypeParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(230)
			p.Match(SubstraitTypeParserEq)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(231)
			p.expr(0)
		}
		p.SetState(233)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SubstraitTypeParserNewline {
			{
				p.SetState(232)
				p.Match(SubstraitTypeParserNewline)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(235)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(247)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SubstraitTypeParserIdentifier {
			{
				p.SetState(237)
				p.Match(SubstraitTypeParserIdentifier)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(238)
				p.Match(SubstraitTypeParserEq)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(239)
				p.expr(0)
			}
			p.SetState(241)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == SubstraitTypeParserNewline {
				{
					p.SetState(240)
					p.Match(SubstraitTypeParserNewline)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(243)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

			p.SetState(249)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(250)

			var _x = p.TypeDef()

			localctx.(*MultilineDefinitionContext).finalType = _x
		}
		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(251)
					p.Match(SubstraitTypeParserNewline)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(256)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case 3:
		localctx = NewTypeLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(257)
			p.TypeDef()
		}

	case 4:
		localctx = NewLiteralNumberContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(258)
			p.Match(SubstraitTypeParserNumber)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewParameterNameContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(259)
			p.Match(SubstraitTypeParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(261)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(260)

				var _m = p.Match(SubstraitTypeParserQMark)

				localctx.(*ParameterNameContext).isnull = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 6:
		localctx = NewFunctionCallContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(263)
			p.Match(SubstraitTypeParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(264)
			p.Match(SubstraitTypeParserOParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3377768440004496) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&196617) != 0) {
			{
				p.SetState(265)
				p.expr(0)
			}
			p.SetState(270)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == SubstraitTypeParserComma {
				{
					p.SetState(266)
					p.Match(SubstraitTypeParserComma)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(267)
					p.expr(0)
				}

				p.SetState(272)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(275)
			p.Match(SubstraitTypeParserCParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewIfExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(276)
			p.Match(SubstraitTypeParserIf)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(277)

			var _x = p.expr(0)

			localctx.(*IfExprContext).ifExpr = _x
		}
		{
			p.SetState(278)
			p.Match(SubstraitTypeParserThen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(279)

			var _x = p.expr(0)

			localctx.(*IfExprContext).thenExpr = _x
		}
		{
			p.SetState(280)
			p.Match(SubstraitTypeParserElse)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(281)

			var _x = p.expr(3)

			localctx.(*IfExprContext).elseExpr = _x
		}

	case 8:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(283)
			p.Match(SubstraitTypeParserBang)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(284)
			p.expr(2)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(296)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinaryExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SubstraitTypeParserRULE_expr)
				p.SetState(287)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(288)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64((_la-53)) & ^0x3f) == 0 && ((int64(1)<<(_la-53))&25167855) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(289)

					var _x = p.expr(5)

					localctx.(*BinaryExprContext).right = _x
				}

			case 2:
				localctx = NewTernaryContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*TernaryContext).ifExpr = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SubstraitTypeParserRULE_expr)
				p.SetState(290)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(291)
					p.Match(SubstraitTypeParserQMark)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(292)

					var _x = p.expr(0)

					localctx.(*TernaryContext).thenExpr = _x
				}
				{
					p.SetState(293)
					p.Match(SubstraitTypeParserColon)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(294)

					var _x = p.expr(2)

					localctx.(*TernaryContext).elseExpr = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(300)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
		if p.HasError() {
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *SubstraitTypeParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 8:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SubstraitTypeParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
