package parser

import (
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/substrait-io/substrait-go/types"
	"github.com/substrait-io/substrait-go/types/integer_parameters"
	baseparser2 "github.com/substrait-io/substrait-go/types/parser/baseparser"
)

type TypeVisitor struct {
	baseparser2.SubstraitTypeVisitor
}

var _ baseparser2.SubstraitTypeVisitor = &TypeVisitor{}

func (v *TypeVisitor) Visit(tree antlr.ParseTree) interface{} {
	if tree == nil {
		return nil
	}
	return tree.Accept(v)
}

func (v *TypeVisitor) VisitStartRule(ctx *baseparser2.StartRuleContext) interface{} {
	return v.Visit(ctx.Expr())
}

func (v *TypeVisitor) VisitTypeStatement(ctx *baseparser2.TypeStatementContext) interface{} {
	return v.Visit(ctx.TypeDef())
}

func (v *TypeVisitor) VisitParenExpression(ctx *baseparser2.ParenExpressionContext) interface{} {
	return v.Visit(ctx.Expr())
}

func (v *TypeVisitor) VisitMultilineDefinition(*baseparser2.MultilineDefinitionContext) interface{} {
	panic("implement MultilineDefinition")
}

func (v *TypeVisitor) VisitTypeLiteral(ctx *baseparser2.TypeLiteralContext) interface{} {
	return v.Visit(ctx.TypeDef())
}

func (v *TypeVisitor) VisitLiteralNumber(ctx *baseparser2.LiteralNumberContext) interface{} {
	num, err := strconv.ParseInt(ctx.Number().GetText(), 10, 0)
	if err != nil {
		panic(err)
	}
	return num
}

func (v *TypeVisitor) VisitFunctionCall(*baseparser2.FunctionCallContext) interface{} {
	panic("implement FunctionCall")
}

func (v *TypeVisitor) VisitBinaryExpr(*baseparser2.BinaryExprContext) interface{} {
	panic("implement BinaryExpr")
}

func (v *TypeVisitor) VisitIfExpr(*baseparser2.IfExprContext) interface{} {
	panic("implement IfExpr")
}

func (v *TypeVisitor) VisitNotExpr(*baseparser2.NotExprContext) interface{} {
	panic("implement NotExpr")
}

func (v *TypeVisitor) VisitTernary(*baseparser2.TernaryContext) interface{} {
	panic("implement Ternary expr")
}

func (v *TypeVisitor) VisitTypeDef(ctx *baseparser2.TypeDefContext) interface{} {
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

func (v *TypeVisitor) VisitAnyType(ctx *baseparser2.AnyTypeContext) interface{} {
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

func (v *TypeVisitor) VisitBoolean(*baseparser2.BooleanContext) interface{} {
	return &types.BooleanType{}
}

func (v *TypeVisitor) VisitI8(*baseparser2.I8Context) interface{} {
	return &types.Int8Type{}
}

func (v *TypeVisitor) VisitI16(*baseparser2.I16Context) interface{} {
	return &types.Int16Type{}
}

func (v *TypeVisitor) VisitI32(*baseparser2.I32Context) interface{} {
	return &types.Int32Type{}
}

func (v *TypeVisitor) VisitI64(*baseparser2.I64Context) interface{} {
	return &types.Int64Type{}
}

func (v *TypeVisitor) VisitFp32(*baseparser2.Fp32Context) interface{} {
	return &types.Float32Type{}
}

func (v *TypeVisitor) VisitFp64(*baseparser2.Fp64Context) interface{} {
	return &types.Float64Type{}
}

func (v *TypeVisitor) VisitString(*baseparser2.StringContext) interface{} {
	return &types.StringType{}
}

func (v *TypeVisitor) VisitBinary(*baseparser2.BinaryContext) interface{} {
	return &types.BinaryType{}
}

func (v *TypeVisitor) VisitTimestamp(*baseparser2.TimestampContext) interface{} {
	return &types.TimestampType{}
}

func (v *TypeVisitor) VisitTimestampTz(*baseparser2.TimestampTzContext) interface{} {
	return &types.TimestampTzType{}
}

func (v *TypeVisitor) VisitDate(*baseparser2.DateContext) interface{} {
	return &types.DateType{}
}

func (v *TypeVisitor) VisitTime(*baseparser2.TimeContext) interface{} {
	return &types.TimeType{}
}

func (v *TypeVisitor) VisitIntervalYear(*baseparser2.IntervalYearContext) interface{} {
	return &types.IntervalYearType{}
}

func (v *TypeVisitor) VisitUuid(*baseparser2.UuidContext) interface{} {
	return &types.UUIDType{}
}

func (v *TypeVisitor) VisitUserDefined(ctx *baseparser2.UserDefinedContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}

	var params []types.UDTParameter
	for _, expr := range ctx.AllExpr() {
		paramExpr := v.Visit(expr)
		switch param := paramExpr.(type) {
		case types.FuncDefArgType:
			params = append(params, &types.DataTypeUDTParam{Type: param})
		case int64:
			params = append(params, &types.IntegerUDTParam{Integer: int32(param)})
		case types.StringParameter:
			params = append(params, &types.StringUDTParam{StringVal: string(param)})
		default:
			// TODO handle other user defined type parameters
			panic("User defined type parameter is not a FuncDefArgType/int/string " + fmt.Sprintf("%T", param))
		}
	}
	name := ctx.Identifier().GetText()
	return &types.ParameterizedUserDefinedType{Name: name, Nullability: nullability, TypeParameters: params}
}

func (v *TypeVisitor) VisitFixedChar(ctx *baseparser2.FixedCharContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}

	length := v.Visit(ctx.GetLength()).(integer_parameters.IntegerParameter)
	return &types.ParameterizedFixedCharType{IntegerOption: length, Nullability: nullability}
}

func (v *TypeVisitor) VisitVarChar(ctx *baseparser2.VarCharContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}

	length := v.Visit(ctx.GetLength()).(integer_parameters.IntegerParameter)
	return &types.ParameterizedVarCharType{IntegerOption: length, Nullability: nullability}
}

func (v *TypeVisitor) VisitFixedBinary(ctx *baseparser2.FixedBinaryContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}

	length := v.Visit(ctx.GetLength()).(integer_parameters.IntegerParameter)
	return &types.ParameterizedFixedBinaryType{IntegerOption: length, Nullability: nullability}
}

func (v *TypeVisitor) VisitDecimal(ctx *baseparser2.DecimalContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}

	precision := v.Visit(ctx.GetPrecision()).(integer_parameters.IntegerParameter)
	scale := v.Visit(ctx.GetScale()).(integer_parameters.IntegerParameter)

	return &types.ParameterizedDecimalType{Precision: precision, Scale: scale, Nullability: nullability}
}

func (v *TypeVisitor) VisitPrecisionTimestamp(ctx *baseparser2.PrecisionTimestampContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}

	length := v.Visit(ctx.GetPrecision()).(integer_parameters.IntegerParameter)
	return &types.ParameterizedPrecisionTimestampType{IntegerOption: length, Nullability: nullability}
}

func (v *TypeVisitor) VisitPrecisionTimestampTZ(ctx *baseparser2.PrecisionTimestampTZContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}

	length := v.Visit(ctx.GetPrecision()).(integer_parameters.IntegerParameter)
	return &types.ParameterizedPrecisionTimestampTzType{IntegerOption: length, Nullability: nullability}
}

func (v *TypeVisitor) VisitPrecisionIntervalDay(ctx *baseparser2.PrecisionIntervalDayContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}

	length := v.Visit(ctx.GetPrecision()).(integer_parameters.IntegerParameter)
	return &types.ParameterizedIntervalDayType{IntegerOption: length, Nullability: nullability}
}

func (v *TypeVisitor) VisitStruct(ctx *baseparser2.StructContext) interface{} {
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

func (v *TypeVisitor) VisitList(ctx *baseparser2.ListContext) interface{} {
	nullability := types.NullabilityRequired
	if ctx.GetIsnull() != nil {
		nullability = types.NullabilityNullable
	}
	elementType := v.Visit(ctx.Expr()).(types.FuncDefArgType)
	return &types.ParameterizedListType{Type: elementType, Nullability: nullability}
}

func (v *TypeVisitor) VisitNStruct(*baseparser2.NStructContext) interface{} {
	panic("implement me")
}

func (v *TypeVisitor) VisitMap(ctx *baseparser2.MapContext) interface{} {
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

func (v *TypeVisitor) VisitParameterName(ctx *baseparser2.ParameterNameContext) interface{} {
	return types.StringParameter(ctx.Identifier().GetText())
}

func (v *TypeVisitor) VisitNumericLiteral(ctx *baseparser2.NumericLiteralContext) interface{} {
	num, err := strconv.ParseInt(ctx.Number().GetText(), 10, 0)
	if err != nil {
		panic(err)
	}
	return integer_parameters.NewConcreteIntParam(int32(num))
}

func (v *TypeVisitor) VisitNumericParameterName(ctx *baseparser2.NumericParameterNameContext) interface{} {
	return integer_parameters.NewVariableIntParam(ctx.GetText())
}

func (v *TypeVisitor) VisitNumericExpression(ctx *baseparser2.NumericExpressionContext) interface{} {
	// TODO handle numeric expression
	return ctx.GetText()
}
