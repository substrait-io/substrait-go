package types

import (
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/substrait-io/substrait-go/parser/types/baseparser"
	"github.com/substrait-io/substrait-go/types"
	"github.com/substrait-io/substrait-go/types/integer_parameters"
)

type MyVisitor struct {
	baseparser.SubstraitTypeVisitor
}

var _ baseparser.SubstraitTypeVisitor = &MyVisitor{}

func NewMyVisitor() *MyVisitor {
	return &MyVisitor{}
}

func (v *MyVisitor) Visit(tree antlr.ParseTree) interface{} {
	if tree == nil {
		return nil
	}
	return tree.Accept(v)
}

func (v *MyVisitor) VisitStartRule(ctx *baseparser.StartRuleContext) interface{} {
	return v.Visit(ctx.Expr())
}

func (v *MyVisitor) VisitTypeStatement(ctx *baseparser.TypeStatementContext) interface{} {
	return v.Visit(ctx.TypeDef())
}

func (v *MyVisitor) VisitParenExpression(ctx *baseparser.ParenExpressionContext) interface{} {
	return v.Visit(ctx.Expr())
}

func (v *MyVisitor) VisitMultilineDefinition(*baseparser.MultilineDefinitionContext) interface{} {
	panic("implement MultilineDefinition")
}

func (v *MyVisitor) VisitTypeLiteral(ctx *baseparser.TypeLiteralContext) interface{} {
	return v.Visit(ctx.TypeDef())
}

func (v *MyVisitor) VisitLiteralNumber(ctx *baseparser.LiteralNumberContext) interface{} {
	num, err := strconv.ParseInt(ctx.Number().GetText(), 10, 0)
	if err != nil {
		panic(err)
	}
	return num
}

func (v *MyVisitor) VisitFunctionCall(*baseparser.FunctionCallContext) interface{} {
	panic("implement FunctionCall")
}

func (v *MyVisitor) VisitBinaryExpr(*baseparser.BinaryExprContext) interface{} {
	panic("implement BinaryExpr")
}

func (v *MyVisitor) VisitIfExpr(*baseparser.IfExprContext) interface{} {
	panic("implement IfExpr")
}

func (v *MyVisitor) VisitNotExpr(*baseparser.NotExprContext) interface{} {
	panic("implement NotExpr")
}

func (v *MyVisitor) VisitTernary(*baseparser.TernaryContext) interface{} {
	panic("implement Ternary expr")
}

func (v *MyVisitor) VisitTypeDef(ctx *baseparser.TypeDefContext) interface{} {
	if ctx.ParameterizedType() != nil {
		return v.Visit(ctx.ParameterizedType())
	}
	if ctx.ScalarType() != nil {
		nullability := types.NullabilityRequired
		if ctx.GetIsnull() != nil {
			nullability = types.NullabilityNullable
		}

		scalarTypeExpr := v.Visit(ctx.ScalarType())
		scalarType := scalarTypeExpr.(types.Type)
		return scalarType.WithNullability(nullability)
	}

	return v.Visit(ctx.AnyType())
}

func (v *MyVisitor) VisitAnyType(ctx *baseparser.AnyTypeContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}

	name := "any"
	if ctx.AnyVar() != nil {
		name = ctx.AnyVar().GetText()
	}
	return types.AnyType{Name: name, Nullability: nullability}
}

func (v *MyVisitor) VisitBoolean(*baseparser.BooleanContext) interface{} {
	return &types.BooleanType{}
}

func (v *MyVisitor) VisitI8(*baseparser.I8Context) interface{} {
	return &types.Int8Type{}
}

func (v *MyVisitor) VisitI16(*baseparser.I16Context) interface{} {
	return &types.Int16Type{}
}

func (v *MyVisitor) VisitI32(*baseparser.I32Context) interface{} {
	return &types.Int32Type{}
}

func (v *MyVisitor) VisitI64(*baseparser.I64Context) interface{} {
	return &types.Int64Type{}
}

func (v *MyVisitor) VisitFp32(*baseparser.Fp32Context) interface{} {
	return &types.Float32Type{}
}

func (v *MyVisitor) VisitFp64(*baseparser.Fp64Context) interface{} {
	return &types.Float64Type{}
}

func (v *MyVisitor) VisitString(*baseparser.StringContext) interface{} {
	return &types.StringType{}
}

func (v *MyVisitor) VisitBinary(*baseparser.BinaryContext) interface{} {
	return &types.BinaryType{}
}

func (v *MyVisitor) VisitTimestamp(*baseparser.TimestampContext) interface{} {
	return &types.TimestampType{}
}

func (v *MyVisitor) VisitTimestampTz(*baseparser.TimestampTzContext) interface{} {
	return &types.TimestampTzType{}
}

func (v *MyVisitor) VisitDate(*baseparser.DateContext) interface{} {
	return &types.DateType{}
}

func (v *MyVisitor) VisitTime(*baseparser.TimeContext) interface{} {
	return &types.TimeType{}
}

func (v *MyVisitor) VisitIntervalYear(*baseparser.IntervalYearContext) interface{} {
	return &types.IntervalYearType{}
}

func (v *MyVisitor) VisitUuid(*baseparser.UuidContext) interface{} {
	return &types.UUIDType{}
}

func (v *MyVisitor) VisitUserDefined(*baseparser.UserDefinedContext) interface{} {
	// TODO handle user defined type name & parameters
	return &types.UserDefinedType{}
}

func (v *MyVisitor) VisitFixedChar(ctx *baseparser.FixedCharContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}

	length := v.Visit(ctx.GetLength()).(integer_parameters.IntegerParameter)
	return &types.ParameterizedFixedCharType{IntegerOption: length, Nullability: nullability}
}

func (v *MyVisitor) VisitVarChar(ctx *baseparser.VarCharContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}

	length := v.Visit(ctx.GetLength()).(integer_parameters.IntegerParameter)
	return &types.ParameterizedVarCharType{IntegerOption: length, Nullability: nullability}
}

func (v *MyVisitor) VisitFixedBinary(ctx *baseparser.FixedBinaryContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}

	length := v.Visit(ctx.GetLength()).(integer_parameters.IntegerParameter)
	return &types.ParameterizedFixedBinaryType{IntegerOption: length, Nullability: nullability}
}

func (v *MyVisitor) VisitDecimal(ctx *baseparser.DecimalContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}

	precision := v.Visit(ctx.GetPrecision()).(integer_parameters.IntegerParameter)
	scale := v.Visit(ctx.GetScale()).(integer_parameters.IntegerParameter)

	return &types.ParameterizedDecimalType{Precision: precision, Scale: scale, Nullability: nullability}
}

func (v *MyVisitor) VisitPrecisionTimestamp(ctx *baseparser.PrecisionTimestampContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}

	length := v.Visit(ctx.GetPrecision()).(integer_parameters.IntegerParameter)
	return &types.ParameterizedPrecisionTimestampType{IntegerOption: length, Nullability: nullability}
}

func (v *MyVisitor) VisitPrecisionTimestampTZ(ctx *baseparser.PrecisionTimestampTZContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}

	length := v.Visit(ctx.GetPrecision()).(integer_parameters.IntegerParameter)
	return &types.ParameterizedPrecisionTimestampTzType{IntegerOption: length, Nullability: nullability}
}

func (v *MyVisitor) VisitPrecisionIntervalDay(ctx *baseparser.PrecisionIntervalDayContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}

	length := v.Visit(ctx.GetPrecision()).(integer_parameters.IntegerParameter)
	return &types.ParameterizedIntervalDayType{IntegerOption: length, Nullability: nullability}
}

func (v *MyVisitor) VisitStruct(ctx *baseparser.StructContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}
	var fieldTypes []types.FuncDefArgType
	for _, expr := range ctx.AllExpr() {
		fieldTypes = append(fieldTypes, v.Visit(expr).(types.FuncDefArgType))
	}
	return &types.ParameterizedStructType{Types: fieldTypes, Nullability: nullability}
}

func (v *MyVisitor) VisitList(ctx *baseparser.ListContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}
	elementType := v.Visit(ctx.Expr()).(types.FuncDefArgType)
	return &types.ParameterizedListType{Type: elementType, Nullability: nullability}
}

func (v *MyVisitor) VisitNStruct(*baseparser.NStructContext) interface{} {
	panic("implement me")
}

func (v *MyVisitor) VisitMap(ctx *baseparser.MapContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}
	keyType, keyOk := v.Visit(ctx.GetKey()).(types.FuncDefArgType)
	valueType, valueOk := v.Visit(ctx.GetValue()).(types.FuncDefArgType)
	if !keyOk || !valueOk {
		panic("Map key or value type is not a FuncDefArgType")
	}
	return &types.ParameterizedMapType{Key: keyType, Value: valueType, Nullability: nullability}
}

func (v *MyVisitor) VisitTypeParam(ctx *baseparser.TypeParamContext) interface{} {
	return types.StringParameter(ctx.Identifier().GetText())
}

func (v *MyVisitor) VisitNumericLiteral(ctx *baseparser.NumericLiteralContext) interface{} {
	num, err := strconv.ParseInt(ctx.Number().GetText(), 10, 0)
	if err != nil {
		panic(err)
	}
	return integer_parameters.NewConcreteIntParam(int32(num))
}

func (v *MyVisitor) VisitNumericParameterName(ctx *baseparser.NumericParameterNameContext) interface{} {
	return integer_parameters.NewVariableIntParam(ctx.GetText())
}

func (v *MyVisitor) VisitNumericExpression(ctx *baseparser.NumericExpressionContext) interface{} {
	// TODO handle numeric expression
	return ctx.GetText()
}
