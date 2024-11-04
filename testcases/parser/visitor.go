package parser

import (
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	baseparser "github.com/substrait-io/substrait-go/testcases/parser/baseparser"
	"github.com/substrait-io/substrait-go/types"
	"github.com/substrait-io/substrait-go/types/integer_parameters"
)

type TestCaseVisitor struct {
	baseparser.FuncTestCaseParserVisitor
}

var _ baseparser.FuncTestCaseParserVisitor = &TestCaseVisitor{}

func (v *TestCaseVisitor) Visit(tree antlr.ParseTree) interface{} {
	if tree == nil {
		return nil
	}
	return tree.Accept(v)
}

func (v *TestCaseVisitor) VisitDoc(ctx *baseparser.DocContext) interface{} {
	header := v.Visit(ctx.Header()).(TestFileHeader)
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
	return TestFileHeader{
		Version:     ctx.Version().GetText(),
		IncludedURI: ctx.Include().GetText(),
	}
}

type TestGroup struct {
	Description string
	TestCases   []*TestCase
}

func (v *TestCaseVisitor) VisitTestGroup(ctx *baseparser.TestGroupContext) interface{} {
	groupDesc := v.Visit(ctx.TestGroupDescription()).(string)
	groupTestCases := make([]*TestCase, 0, len(ctx.AllTestCase()))
	for _, tc := range ctx.AllTestCase() {
		testcase := v.Visit(tc).(*TestCase)
		testcase.GroupDesc = groupDesc
		groupTestCases = append(groupTestCases, testcase)
	}
	return &TestGroup{
		Description: groupDesc,
		TestCases:   groupTestCases,
	}
}

func (v *TestCaseVisitor) VisitTestGroupDescription(ctx *baseparser.TestGroupDescriptionContext) interface{} {
	return strings.TrimPrefix(ctx.GetText(), "#")
}

type CaseLiteral struct {
	Type      types.Type
	ValueText string
	Value     any
}

func (v *TestCaseVisitor) VisitTestCase(ctx *baseparser.TestCaseContext) interface{} {
	return &TestCase{
		FuncName: ctx.Identifier().GetText(),
		Args:     v.Visit(ctx.Arguments()).([]*CaseLiteral),
		Result:   v.Visit(ctx.Result()).(*CaseLiteral),
	}
}

func (v *TestCaseVisitor) VisitArguments(ctx *baseparser.ArgumentsContext) interface{} {
	args := make([]*CaseLiteral, 0, len(ctx.AllArgument()))
	for _, expr := range ctx.AllArgument() {
		args = append(args, v.Visit(expr).(*CaseLiteral))
	}
	return args
}

func (v *TestCaseVisitor) VisitArgument(ctx *baseparser.ArgumentContext) interface{} {
	if ctx.I8Arg() != nil {
		return v.Visit(ctx.I8Arg())
	}
	if ctx.I16Arg() != nil {
		return v.Visit(ctx.I16Arg())
	}
	if ctx.I32Arg() != nil {
		return v.Visit(ctx.I32Arg())
	}
	if ctx.I64Arg() != nil {
		return v.Visit(ctx.I64Arg())
	}
	if ctx.Fp32Arg() != nil {
		return v.Visit(ctx.Fp32Arg())
	}
	if ctx.Fp64Arg() != nil {
		return v.Visit(ctx.Fp64Arg())
	}
	if ctx.StringArg() != nil {
		return v.Visit(ctx.StringArg())
	}
	if ctx.BooleanArg() != nil {
		return v.Visit(ctx.BooleanArg())
	}
	if ctx.TimestampArg() != nil {
		return v.Visit(ctx.TimestampArg())
	}
	if ctx.TimestampTzArg() != nil {
		return v.Visit(ctx.TimestampTzArg())
	}
	if ctx.DateArg() != nil {
		return v.Visit(ctx.DateArg())
	}
	if ctx.TimeArg() != nil {
		return v.Visit(ctx.TimeArg())
	}
	if ctx.IntervalYearArg() != nil {
		return v.Visit(ctx.IntervalYearArg())
	}
	if ctx.IntervalDayArg() != nil {
		return v.Visit(ctx.IntervalDayArg())
	}
	if ctx.NullArg() != nil {
		return v.Visit(ctx.NullArg())
	}
	if ctx.DecimalArg() != nil {
		return v.Visit(ctx.DecimalArg())
	}

	return CaseLiteral{}
}

func (v *TestCaseVisitor) VisitNullArg(*baseparser.NullArgContext) interface{} {
	return &CaseLiteral{}
}

func (v *TestCaseVisitor) VisitBooleanArg(ctx *baseparser.BooleanArgContext) interface{} {
	value := false
	if strings.ToLower(ctx.BooleanLiteral().GetText()) == "true" {
		value = true
	}
	return &CaseLiteral{Value: value, ValueText: ctx.BooleanLiteral().GetText(), Type: &types.BooleanType{}}
}

func (v *TestCaseVisitor) VisitI8Arg(ctx *baseparser.I8ArgContext) interface{} {
	value, err := strconv.ParseInt(ctx.IntegerLiteral().GetText(), 10, 8)
	if err != nil {
		panic(err)
	}
	return &CaseLiteral{Value: int8(value), ValueText: ctx.IntegerLiteral().GetText(), Type: &types.Int8Type{}}
}

func (v *TestCaseVisitor) VisitI16Arg(ctx *baseparser.I16ArgContext) interface{} {
	value, err := strconv.ParseInt(ctx.IntegerLiteral().GetText(), 10, 16)
	if err != nil {
		panic(err)
	}
	return &CaseLiteral{Value: int16(value), ValueText: ctx.IntegerLiteral().GetText(), Type: &types.Int16Type{}}
}

func (v *TestCaseVisitor) VisitI32Arg(ctx *baseparser.I32ArgContext) interface{} {
	value, err := strconv.ParseInt(ctx.IntegerLiteral().GetText(), 10, 32)
	if err != nil {
		panic(err)
	}
	return &CaseLiteral{Value: int32(value), ValueText: ctx.IntegerLiteral().GetText(), Type: &types.Int32Type{}}
}

func (v *TestCaseVisitor) VisitI64Arg(ctx *baseparser.I64ArgContext) interface{} {
	value, err := strconv.ParseInt(ctx.IntegerLiteral().GetText(), 10, 64)
	if err != nil {
		panic(err)
	}
	return &CaseLiteral{Value: value, ValueText: ctx.IntegerLiteral().GetText(), Type: &types.Int64Type{}}
}

func (v *TestCaseVisitor) VisitFp32Arg(ctx *baseparser.Fp32ArgContext) interface{} {
	value, err := strconv.ParseFloat(ctx.NumericLiteral().GetText(), 32)
	if err != nil {
		panic(err)
	}
	return &CaseLiteral{Value: float32(value), ValueText: ctx.NumericLiteral().GetText(), Type: &types.Float32Type{}}
}

func (v *TestCaseVisitor) VisitFp64Arg(ctx *baseparser.Fp64ArgContext) interface{} {
	value, err := strconv.ParseFloat(ctx.NumericLiteral().GetText(), 64)
	if err != nil {
		panic(err)
	}
	return &CaseLiteral{Value: value, ValueText: ctx.NumericLiteral().GetText(), Type: &types.Float64Type{}}
}

func (v *TestCaseVisitor) VisitStringArg(ctx *baseparser.StringArgContext) interface{} {
	return &CaseLiteral{Value: ctx.StringLiteral().GetText(), ValueText: ctx.StringLiteral().GetText(), Type: &types.StringType{}}
}

func (v *TestCaseVisitor) VisitTimestampArg(ctx *baseparser.TimestampArgContext) interface{} {
	return &CaseLiteral{Value: ctx.TimestampLiteral().GetText(), ValueText: ctx.TimestampLiteral().GetText(), Type: &types.TimestampType{}}
}

func (v *TestCaseVisitor) VisitTimestampTzArg(ctx *baseparser.TimestampTzArgContext) interface{} {
	return &CaseLiteral{Value: ctx.TimestampTzLiteral().GetText(), ValueText: ctx.TimestampTzLiteral().GetText(), Type: &types.TimestampTzType{}}
}

func (v *TestCaseVisitor) VisitDateArg(ctx *baseparser.DateArgContext) interface{} {
	return &CaseLiteral{Value: ctx.DateLiteral().GetText(), ValueText: ctx.DateLiteral().GetText(), Type: &types.DateType{}}
}

func (v *TestCaseVisitor) VisitTimeArg(ctx *baseparser.TimeArgContext) interface{} {
	return &CaseLiteral{Value: ctx.TimeLiteral().GetText(), ValueText: ctx.TimeLiteral().GetText(), Type: &types.TimeType{}}
}

func (v *TestCaseVisitor) VisitIntervalYearArg(ctx *baseparser.IntervalYearArgContext) interface{} {
	return &CaseLiteral{Value: ctx.IntervalYearLiteral().GetText(), ValueText: ctx.IntervalYearLiteral().GetText(), Type: &types.IntervalYearType{}}
}

func (v *TestCaseVisitor) VisitIntervalDayArg(ctx *baseparser.IntervalDayArgContext) interface{} {
	return &CaseLiteral{Value: ctx.IntervalDayLiteral().GetText(), ValueText: ctx.IntervalDayLiteral().GetText(), Type: &types.IntervalDayType{}}
}

func (v *TestCaseVisitor) VisitDecimalArg(ctx *baseparser.DecimalArgContext) interface{} {
	return &CaseLiteral{Value: v.Visit(ctx.NumericLiteral()), ValueText: ctx.NumericLiteral().GetText(), Type: &types.DecimalType{}}
}

func (v *TestCaseVisitor) VisitResult(ctx *baseparser.ResultContext) interface{} {
	return v.Visit(ctx.Argument()).(*CaseLiteral)
}

func (v *TestCaseVisitor) VisitNumericLiteral(ctx *baseparser.NumericLiteralContext) interface{} {
	if ctx.IntegerLiteral() != nil {
		value, err := strconv.ParseInt(ctx.IntegerLiteral().GetText(), 10, 64)
		if err != nil {
			panic(err)
		}
		return value
	}
	if ctx.DecimalLiteral() != nil { // TODO
		value, err := strconv.ParseFloat(ctx.DecimalLiteral().GetText(), 64)
		if err != nil {
			panic(err)
		}
		return value
	}
	return v.Visit(ctx.FloatLiteral())
}

func (v *TestCaseVisitor) VisitFloatLiteral(ctx *baseparser.FloatLiteralContext) interface{} {
	return nil // TODO
}
func (v *TestCaseVisitor) VisitBoolean(*baseparser.BooleanContext) interface{} {
	return &types.BooleanType{}
}

func (v *TestCaseVisitor) VisitI8(*baseparser.I8Context) interface{} {
	return &types.Int8Type{}
}

func (v *TestCaseVisitor) VisitI16(*baseparser.I16Context) interface{} {
	return &types.Int16Type{}
}

func (v *TestCaseVisitor) VisitI32(*baseparser.I32Context) interface{} {
	return &types.Int32Type{}
}

func (v *TestCaseVisitor) VisitI64(*baseparser.I64Context) interface{} {
	return &types.Int64Type{}
}

func (v *TestCaseVisitor) VisitFp32(*baseparser.Fp32Context) interface{} {
	return &types.Float32Type{}
}

func (v *TestCaseVisitor) VisitFp64(*baseparser.Fp64Context) interface{} {
	return &types.Float64Type{}
}

func (v *TestCaseVisitor) VisitString(*baseparser.StringContext) interface{} {
	return &types.StringType{}
}

func (v *TestCaseVisitor) VisitBinary(*baseparser.BinaryContext) interface{} {
	return &types.BinaryType{}
}

func (v *TestCaseVisitor) VisitTimestamp(*baseparser.TimestampContext) interface{} {
	return &types.TimestampType{}
}

func (v *TestCaseVisitor) VisitTimestampTz(*baseparser.TimestampTzContext) interface{} {
	return &types.TimestampTzType{}
}

func (v *TestCaseVisitor) VisitDate(*baseparser.DateContext) interface{} {
	return &types.DateType{}
}

func (v *TestCaseVisitor) VisitTime(*baseparser.TimeContext) interface{} {
	return &types.TimeType{}
}

func (v *TestCaseVisitor) VisitIntervalYear(*baseparser.IntervalYearContext) interface{} {
	return &types.IntervalYearType{}
}

func (v *TestCaseVisitor) VisitUuid(*baseparser.UuidContext) interface{} {
	return &types.UUIDType{}
}

func (v *TestCaseVisitor) VisitDecimal(ctx *baseparser.DecimalContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}

	precision := v.Visit(ctx.GetPrecision()).(integer_parameters.IntegerParameter)
	scale := v.Visit(ctx.GetScale()).(integer_parameters.IntegerParameter)

	return &types.ParameterizedDecimalType{Precision: precision, Scale: scale, Nullability: nullability}
}

func (v *TestCaseVisitor) VisitPrecisionTimestamp(ctx *baseparser.PrecisionTimestampContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}

	length := v.Visit(ctx.GetPrecision()).(integer_parameters.IntegerParameter)
	return &types.ParameterizedPrecisionTimestampType{IntegerOption: length, Nullability: nullability}
}

func (v *TestCaseVisitor) VisitPrecisionTimestampTZ(ctx *baseparser.PrecisionTimestampTZContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}

	length := v.Visit(ctx.GetPrecision()).(integer_parameters.IntegerParameter)
	return &types.ParameterizedPrecisionTimestampTzType{IntegerOption: length, Nullability: nullability}
}
