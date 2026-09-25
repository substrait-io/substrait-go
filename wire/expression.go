// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"fmt"

	substraitgo "github.com/substrait-io/substrait-go/v9"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// ExprToProto encodes an expression as its protobuf message.
func ExprToProto(e expr.Expression) *proto.Expression {
	switch e := e.(type) {
	case *expr.Cast:
		return castToProto(e)
	case *expr.Lambda:
		return lambdaToProto(e)
	case *expr.ScalarFunction:
		return scalarFunctionToProto(e)
	case *expr.WindowFunction:
		return windowFunctionToProto(e)
	case *expr.FieldReference:
		return FieldReferenceToProto(e)
	case expr.Literal:
		return &proto.Expression{
			RexType: &proto.Expression_Literal_{Literal: LiteralToProto(e)},
		}
	default:
		panic(fmt.Sprintf("wire: unhandled expression %T", e))
	}
}

func castToProto(ex *expr.Cast) *proto.Expression {
	return &proto.Expression{
		RexType: &proto.Expression_Cast_{
			Cast: &proto.Expression_Cast{
				Type:            TypeToProto(ex.Type),
				Input:           ExprToProto(ex.Input),
				FailureBehavior: proto.Expression_Cast_FailureBehavior(ex.FailureBehavior),
			},
		},
	}
}

// VirtualTableExpressionValueToProto encodes a virtual-table row of expressions.
func VirtualTableExpressionValueToProto(s expr.VirtualTableExpressionValue) *proto.Expression_Nested_Struct {
	fields := make([]*proto.Expression, len(s))
	for i, f := range s {
		fields[i] = ExprToProto(f)
	}
	return &proto.Expression_Nested_Struct{Fields: fields}
}

// FuncArgFromProto decodes a function argument from its protobuf message.
func FuncArgFromProto(e *proto.FunctionArgument, baseSchema *types.RecordType, reg expr.ExtensionRegistry) (types.FuncArg, error) {
	switch et := e.ArgType.(type) {
	case *proto.FunctionArgument_Enum:
		return types.Enum(et.Enum), nil
	case *proto.FunctionArgument_Type:
		return TypeFromProto(et.Type), nil
	case *proto.FunctionArgument_Value:
		return ExprFromProto(et.Value, baseSchema, reg)
	}
	return nil, substraitgo.ErrNotImplemented
}

// ExprFromProto decodes an expression from its protobuf message.
func ExprFromProto(e *proto.Expression, baseSchema *types.RecordType, reg expr.ExtensionRegistry) (expr.Expression, error) {
	if e == nil {
		return nil, fmt.Errorf("%w: protobuf Expression is nil", substraitgo.ErrInvalidExpr)
	}

	switch et := e.RexType.(type) {
	case *proto.Expression_Literal_:
		return LiteralFromProto(et.Literal), nil
	case *proto.Expression_Selection:
		return FieldReferenceFromProto(et.Selection, baseSchema, reg)
	case *proto.Expression_ScalarFunction_:
		var err error
		args := make([]types.FuncArg, len(et.ScalarFunction.Arguments))
		for i, a := range et.ScalarFunction.Arguments {
			if args[i], err = FuncArgFromProto(a, baseSchema, reg); err != nil {
				return nil, err
			}
		}

		if et.ScalarFunction.OutputType == nil {
			return nil, fmt.Errorf("%w: scalar function missing output type", substraitgo.ErrInvalidExpr)
		}

		id, ok := reg.DecodeFunc(et.ScalarFunction.FunctionReference)
		if !ok {
			return nil, substraitgo.ErrNotFound
		}

		decl, ok := reg.LookupScalarFunction(et.ScalarFunction.FunctionReference)
		if !ok {
			return expr.NewCustomScalarFunc(reg, extensions.NewScalarFuncVariant(id), TypeFromProto(et.ScalarFunction.OutputType), FunctionOptionsFromProto(et.ScalarFunction.Options), args...)
		}

		return expr.NewScalarFunctionFromParts(
			et.ScalarFunction.FunctionReference,
			decl,
			args,
			FunctionOptionsFromProto(et.ScalarFunction.Options),
			TypeFromProto(et.ScalarFunction.OutputType),
		), nil
	case *proto.Expression_WindowFunction_:
		var err error
		args := make([]types.FuncArg, len(et.WindowFunction.Arguments))
		for i, a := range et.WindowFunction.Arguments {
			if args[i], err = FuncArgFromProto(a, baseSchema, reg); err != nil {
				return nil, err
			}
		}

		parts := make([]expr.Expression, len(et.WindowFunction.Partitions))
		for i, p := range et.WindowFunction.Partitions {
			if parts[i], err = ExprFromProto(p, baseSchema, reg); err != nil {
				return nil, err
			}
		}

		sorts := make([]expr.SortField, len(et.WindowFunction.Sorts))
		for i, s := range et.WindowFunction.Sorts {
			if sorts[i], err = SortFieldFromProto(s, baseSchema, reg); err != nil {
				return nil, err
			}
		}

		if et.WindowFunction.OutputType == nil {
			return nil, fmt.Errorf("%w: window function missing output type", substraitgo.ErrInvalidExpr)
		}

		id, ok := reg.DecodeFunc(et.WindowFunction.FunctionReference)
		if !ok {
			return nil, substraitgo.ErrNotFound
		}
		decl, ok := reg.LookupWindowFunction(et.WindowFunction.FunctionReference)
		if !ok {
			fn, err := expr.NewCustomWindowFunc(reg, extensions.NewWindowFuncVariant(id), TypeFromProto(et.WindowFunction.OutputType),
				FunctionOptionsFromProto(et.WindowFunction.Options), types.AggregationInvocation(et.WindowFunction.Invocation), types.AggregationPhase(et.WindowFunction.Phase), args...)
			if err != nil {
				return nil, err
			}

			fn.Partitions = parts
			fn.Sorts = sorts
			fn.LowerBound = BoundFromProto(et.WindowFunction.LowerBound)
			fn.BoundsType = types.BoundsType(et.WindowFunction.BoundsType)
			fn.UpperBound = BoundFromProto(et.WindowFunction.UpperBound)
			return fn, nil
		}

		return expr.NewWindowFunctionFromParts(
			et.WindowFunction.FunctionReference,
			decl,
			args,
			FunctionOptionsFromProto(et.WindowFunction.Options),
			TypeFromProto(et.WindowFunction.OutputType),
			types.AggregationPhase(et.WindowFunction.Phase),
			types.AggregationInvocation(et.WindowFunction.Invocation),
			sorts,
			parts,
			types.BoundsType(et.WindowFunction.BoundsType),
			BoundFromProto(et.WindowFunction.LowerBound),
			BoundFromProto(et.WindowFunction.UpperBound),
		), nil
	case *proto.Expression_Enum_:
		return nil, fmt.Errorf("%w: deprecated", substraitgo.ErrNotImplemented)
	case *proto.Expression_Cast_:
		if et.Cast.Type == nil {
			return nil, fmt.Errorf("%w: cast expression missing type", substraitgo.ErrInvalidExpr)
		}

		input, err := ExprFromProto(et.Cast.Input, baseSchema, reg)
		if err != nil {
			return nil, err
		}

		return &expr.Cast{
			Type:            TypeFromProto(et.Cast.Type),
			Input:           input,
			FailureBehavior: types.CastFailBehavior(et.Cast.FailureBehavior),
		}, nil
	case *proto.Expression_Lambda_:
		if et.Lambda.Parameters == nil {
			return nil, fmt.Errorf("%w: lambda parameters cannot be nil", substraitgo.ErrInvalidExpr)
		}
		if et.Lambda.Body == nil {
			return nil, fmt.Errorf("%w: lambda body cannot be nil", substraitgo.ErrInvalidExpr)
		}

		paramTypes := make([]types.Type, len(et.Lambda.Parameters.Types))
		for i, pt := range et.Lambda.Parameters.Types {
			paramTypes[i] = TypeFromProto(pt)
		}
		params := &types.StructType{
			Types:            paramTypes,
			TypeVariationRef: et.Lambda.Parameters.TypeVariationReference,
			Nullability:      types.Nullability(et.Lambda.Parameters.Nullability),
		}

		if params.Nullability != types.NullabilityRequired {
			return nil, fmt.Errorf("%w: lambda parameters struct must have NULLABILITY_REQUIRED", substraitgo.ErrInvalidExpr)
		}

		body, err := ExprFromProto(et.Lambda.Body, baseSchema, reg)
		if err != nil {
			return nil, err
		}

		// TODO (#189): add validation and type resolution for lambda parameter references
		return &expr.Lambda{Parameters: params, Body: body}, nil
	}
	return nil, fmt.Errorf("%w: ExprFromProto: %s", substraitgo.ErrNotImplemented, e)
}
