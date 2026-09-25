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
	case *expr.DynamicParameter:
		return dynamicParameterToProto(e)
	case *expr.IfThen:
		return ifThenToProto(e)
	case *expr.SwitchExpr:
		return switchExprToProto(e)
	case *expr.SingularOrList:
		return singularOrListToProto(e)
	case *expr.MultiOrList:
		return multiOrListToProto(e)
	case *expr.MapExpr:
		return mapExprToProto(e)
	case *expr.StructExpr:
		return structExprToProto(e)
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

func dynamicParameterToProto(dp *expr.DynamicParameter) *proto.Expression {
	return &proto.Expression{
		RexType: &proto.Expression_DynamicParameter{
			DynamicParameter: &proto.DynamicParameter{
				Type:               TypeToProto(dp.OutputType),
				ParameterReference: dp.ParameterReference,
			},
		},
	}
}

func ifThenToProto(ex *expr.IfThen) *proto.Expression {
	clauses := make([]*proto.Expression_IfThen_IfClause, ex.NIfs())
	for i := range clauses {
		pair := ex.IfPair(i)
		clauses[i] = &proto.Expression_IfThen_IfClause{
			If:   ExprToProto(pair.If),
			Then: ExprToProto(pair.Then),
		}
	}

	var elseClause *proto.Expression
	if e := ex.Else(); e != nil {
		elseClause = ExprToProto(e)
	}
	return &proto.Expression{
		RexType: &proto.Expression_IfThen_{
			IfThen: &proto.Expression_IfThen{
				Ifs:  clauses,
				Else: elseClause,
			},
		},
	}
}

func switchExprToProto(ex *expr.SwitchExpr) *proto.Expression {
	var elseExpr *proto.Expression
	if e := ex.Else(); e != nil {
		elseExpr = ExprToProto(e)
	}

	cases := make([]*proto.Expression_SwitchExpression_IfValue, ex.NCases())
	for i := range cases {
		c := ex.Case(i)
		cases[i] = &proto.Expression_SwitchExpression_IfValue{
			If:   LiteralToProto(c.If),
			Then: ExprToProto(c.Then),
		}
	}

	return &proto.Expression{
		RexType: &proto.Expression_SwitchExpression_{
			SwitchExpression: &proto.Expression_SwitchExpression{
				Match: ExprToProto(ex.MatchExpr()),
				Ifs:   cases,
				Else:  elseExpr,
			},
		},
	}
}

func singularOrListToProto(ex *expr.SingularOrList) *proto.Expression {
	opts := make([]*proto.Expression, len(ex.Options))
	for i, o := range ex.Options {
		opts[i] = ExprToProto(o)
	}
	return &proto.Expression{
		RexType: &proto.Expression_SingularOrList_{
			SingularOrList: &proto.Expression_SingularOrList{
				Value:   ExprToProto(ex.Value),
				Options: opts,
			},
		},
	}
}

func multiOrListToProto(ex *expr.MultiOrList) *proto.Expression {
	toSlice := func(exprs []expr.Expression) []*proto.Expression {
		out := make([]*proto.Expression, len(exprs))
		for i, e := range exprs {
			out[i] = ExprToProto(e)
		}
		return out
	}

	opts := make([]*proto.Expression_MultiOrList_Record, len(ex.Options))
	for i, o := range ex.Options {
		opts[i] = &proto.Expression_MultiOrList_Record{Fields: toSlice(o)}
	}

	return &proto.Expression{
		RexType: &proto.Expression_MultiOrList_{
			MultiOrList: &proto.Expression_MultiOrList{
				Value:   toSlice(ex.Value),
				Options: opts,
			},
		},
	}
}

func mapExprToProto(ex *expr.MapExpr) *proto.Expression {
	kvs := make([]*proto.Expression_Nested_Map_KeyValue, len(ex.KeyValues))
	for i, kv := range ex.KeyValues {
		kvs[i] = &proto.Expression_Nested_Map_KeyValue{
			Key:   ExprToProto(kv.Key),
			Value: ExprToProto(kv.Value),
		}
	}
	return &proto.Expression{
		RexType: &proto.Expression_Nested_{
			Nested: &proto.Expression_Nested{
				Nullable:               ex.Nullable,
				TypeVariationReference: ex.TypeVariationRef,
				NestedType: &proto.Expression_Nested_Map_{
					Map: &proto.Expression_Nested_Map{KeyValues: kvs},
				},
			},
		},
	}
}

func structExprToProto(ex *expr.StructExpr) *proto.Expression {
	fields := make([]*proto.Expression, len(ex.Fields))
	for i, f := range ex.Fields {
		fields[i] = ExprToProto(f)
	}
	return &proto.Expression{
		RexType: &proto.Expression_Nested_{
			Nested: &proto.Expression_Nested{
				Nullable:               ex.Nullable,
				TypeVariationReference: ex.TypeVariationRef,
				NestedType: &proto.Expression_Nested_Struct_{
					Struct: &proto.Expression_Nested_Struct{Fields: fields},
				},
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
	case *proto.Expression_DynamicParameter:
		if et.DynamicParameter == nil {
			return nil, fmt.Errorf("%w: dynamic parameter is nil", substraitgo.ErrInvalidExpr)
		}
		return &expr.DynamicParameter{
			OutputType:         TypeFromProto(et.DynamicParameter.Type),
			ParameterReference: et.DynamicParameter.ParameterReference,
		}, nil
	case *proto.Expression_IfThen_:
		elseExpr, err := ExprFromProto(et.IfThen.Else, baseSchema, reg)
		if err != nil {
			return nil, err
		}

		ifs := make([]expr.IfThenPair, len(et.IfThen.Ifs))
		for i, clause := range et.IfThen.Ifs {
			ifs[i].If, err = ExprFromProto(clause.If, baseSchema, reg)
			if err != nil {
				return nil, err
			}

			ifs[i].Then, err = ExprFromProto(clause.Then, baseSchema, reg)
			if err != nil {
				return nil, err
			}
		}

		return expr.NewIfThenFromParts(ifs, elseExpr), nil
	case *proto.Expression_SwitchExpression_:
		matched, err := ExprFromProto(et.SwitchExpression.Match, baseSchema, reg)
		if err != nil {
			return nil, err
		}

		elseExpr, err := ExprFromProto(et.SwitchExpression.Else, baseSchema, reg)
		if err != nil {
			return nil, err
		}

		ifs := make([]struct {
			If   expr.Literal
			Then expr.Expression
		}, len(et.SwitchExpression.Ifs))
		for i, clause := range et.SwitchExpression.Ifs {
			ifs[i].If = LiteralFromProto(clause.If)
			ifs[i].Then, err = ExprFromProto(clause.Then, baseSchema, reg)
			if err != nil {
				return nil, err
			}
		}

		return expr.NewSwitchExprFromParts(matched, ifs, elseExpr), nil
	case *proto.Expression_SingularOrList_:
		val, err := ExprFromProto(et.SingularOrList.Value, baseSchema, reg)
		if err != nil {
			return nil, err
		}

		opts := make([]expr.Expression, len(et.SingularOrList.Options))
		for i, o := range et.SingularOrList.Options {
			opts[i], err = ExprFromProto(o, baseSchema, reg)
			if err != nil {
				return nil, err
			}
		}

		return &expr.SingularOrList{
			Value:   val,
			Options: opts,
		}, nil
	case *proto.Expression_MultiOrList_:
		var err error
		val := make([]expr.Expression, len(et.MultiOrList.Value))
		for i, v := range et.MultiOrList.Value {
			val[i], err = ExprFromProto(v, baseSchema, reg)
			if err != nil {
				return nil, err
			}
		}

		options := make([][]expr.Expression, len(et.MultiOrList.Options))
		for i, opts := range et.MultiOrList.Options {
			options[i] = make([]expr.Expression, len(opts.Fields))
			for j, o := range opts.Fields {
				options[i][j], err = ExprFromProto(o, baseSchema, reg)
				if err != nil {
					return nil, err
				}
			}
		}

		return &expr.MultiOrList{
			Value:   val,
			Options: options,
		}, nil
	case *proto.Expression_Nested_:
		var err error
		nullable, typevar := et.Nested.Nullable, et.Nested.TypeVariationReference

		switch n := et.Nested.NestedType.(type) {
		case *proto.Expression_Nested_Map_:
			if len(n.Map.KeyValues) == 0 {
				return nil, fmt.Errorf("%w: use an empty map literal instead of NestedExpr map to preserve type info",
					substraitgo.ErrInvalidExpr)
			}

			keyValues := make([]struct{ Key, Value expr.Expression }, len(n.Map.KeyValues))
			for i, kv := range n.Map.KeyValues {
				keyValues[i].Key, err = ExprFromProto(kv.Key, baseSchema, reg)
				if err != nil {
					return nil, err
				}

				keyValues[i].Value, err = ExprFromProto(kv.Value, baseSchema, reg)
				if err != nil {
					return nil, err
				}
			}

			return &expr.MapExpr{
				Nullable:         nullable,
				TypeVariationRef: typevar,
				KeyValues:        keyValues,
			}, nil
		case *proto.Expression_Nested_Struct_:
			fields := make([]expr.Expression, len(n.Struct.Fields))
			for i, f := range n.Struct.Fields {
				fields[i], err = ExprFromProto(f, baseSchema, reg)
				if err != nil {
					return nil, err
				}
			}

			return &expr.StructExpr{
				Nullable:         nullable,
				TypeVariationRef: typevar,
				Fields:           fields,
			}, nil
		default:
			return nil, fmt.Errorf("%w: nested expression: %s",
				substraitgo.ErrInvalidExpr, n)
		}
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
