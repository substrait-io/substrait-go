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

func scalarFunctionToProto(s *expr.ScalarFunction) *proto.Expression {
	args := make([]*proto.FunctionArgument, s.NArgs())
	for i := range args {
		args[i] = FuncArgToProto(s.Arg(i))
	}

	return &proto.Expression{
		RexType: &proto.Expression_ScalarFunction_{
			ScalarFunction: &proto.Expression_ScalarFunction{
				FunctionReference: s.FuncRef(),
				Options:           FunctionOptionsToProto(s.GetOptions()),
				OutputType:        TypeToProto(s.GetType()),
				Arguments:         args,
			},
		},
	}
}

func windowFunctionToProto(w *expr.WindowFunction) *proto.Expression {
	var (
		args       []*proto.FunctionArgument
		sorts      []*proto.SortField
		parts      []*proto.Expression
		upperBound *proto.Expression_WindowFunction_Bound
		lowerBound *proto.Expression_WindowFunction_Bound
	)

	if w.NArgs() > 0 {
		args = make([]*proto.FunctionArgument, w.NArgs())
		for i := range args {
			args[i] = FuncArgToProto(w.Arg(i))
		}
	}

	if len(w.Sorts) > 0 {
		sorts = make([]*proto.SortField, len(w.Sorts))
		for i, s := range w.Sorts {
			sorts[i] = SortFieldToProto(&s)
		}
	}

	if len(w.Partitions) > 0 {
		parts = make([]*proto.Expression, len(w.Partitions))
		for i, p := range w.Partitions {
			parts[i] = ExprToProto(p)
		}
	}

	if w.UpperBound != nil {
		upperBound = BoundToProto(w.UpperBound)
	}

	if w.LowerBound != nil {
		lowerBound = BoundToProto(w.LowerBound)
	}

	return &proto.Expression{
		RexType: &proto.Expression_WindowFunction_{
			WindowFunction: &proto.Expression_WindowFunction{
				FunctionReference: w.FuncRef(),
				Arguments:         args,
				Options:           FunctionOptionsToProto(w.GetOptions()),
				OutputType:        TypeToProto(w.GetType()),
				Phase:             proto.AggregationPhase(w.Phase()),
				Sorts:             sorts,
				Invocation:        proto.AggregateFunction_AggregationInvocation(w.Invocation()),
				Partitions:        parts,
				BoundsType:        proto.Expression_WindowFunction_BoundsType(w.BoundsType),
				LowerBound:        lowerBound,
				UpperBound:        upperBound,
			},
		},
	}
}

// AggregateFunctionToProto encodes an aggregate function as its protobuf message.
func AggregateFunctionToProto(a *expr.AggregateFunction) *proto.AggregateFunction {
	var (
		args  []*proto.FunctionArgument
		sorts []*proto.SortField
	)
	if a.NArgs() > 0 {
		args = make([]*proto.FunctionArgument, a.NArgs())
		for i := range args {
			args[i] = FuncArgToProto(a.Arg(i))
		}
	}

	if len(a.Sorts) > 0 {
		sorts = make([]*proto.SortField, len(a.Sorts))
		for i, s := range a.Sorts {
			sorts[i] = SortFieldToProto(&s)
		}
	}

	return &proto.AggregateFunction{
		FunctionReference: a.FuncRef(),
		Arguments:         args,
		Options:           FunctionOptionsToProto(a.GetOptions()),
		OutputType:        TypeToProto(a.GetType()),
		Phase:             proto.AggregationPhase(a.Phase()),
		Sorts:             sorts,
		Invocation:        proto.AggregateFunction_AggregationInvocation(a.Invocation()),
	}
}

// SortFieldToProto encodes a sort field as its protobuf message.
func SortFieldToProto(s *expr.SortField) *proto.SortField {
	ret := &proto.SortField{Expr: ExprToProto(s.Expr)}
	switch k := s.Kind.(type) {
	case types.SortDirection:
		ret.SortKind = &proto.SortField_Direction{
			Direction: proto.SortField_SortDirection(k)}
	case types.FunctionRef:
		ret.SortKind = &proto.SortField_ComparisonFunctionReference{
			ComparisonFunctionReference: uint32(k)}
	}
	return ret
}

// BoundToProto encodes a window-function bound as its protobuf message.
func BoundToProto(b expr.Bound) *proto.Expression_WindowFunction_Bound {
	switch b := b.(type) {
	case expr.PrecedingBound:
		return &proto.Expression_WindowFunction_Bound{
			Kind: &proto.Expression_WindowFunction_Bound_Preceding_{
				Preceding: &proto.Expression_WindowFunction_Bound_Preceding{Offset: int64(b)},
			},
		}
	case expr.FollowingBound:
		return &proto.Expression_WindowFunction_Bound{
			Kind: &proto.Expression_WindowFunction_Bound_Following_{
				Following: &proto.Expression_WindowFunction_Bound_Following{Offset: int64(b)},
			},
		}
	case expr.CurrentRow:
		return &proto.Expression_WindowFunction_Bound{
			Kind: &proto.Expression_WindowFunction_Bound_CurrentRow_{
				CurrentRow: &proto.Expression_WindowFunction_Bound_CurrentRow{},
			},
		}
	case expr.Unbounded:
		return &proto.Expression_WindowFunction_Bound{
			Kind: &proto.Expression_WindowFunction_Bound_Unbounded_{
				Unbounded: &proto.Expression_WindowFunction_Bound_Unbounded{},
			},
		}
	default:
		panic(fmt.Sprintf("wire: unhandled bound %T", b))
	}
}

// SortFieldFromProto decodes a sort field from its protobuf message.
func SortFieldFromProto(
	f *proto.SortField, baseSchema *types.RecordType, reg expr.ExtensionRegistry,
) (sf expr.SortField, err error) {
	sf.Expr, err = ExprFromProto(f.Expr, baseSchema, reg)
	if err != nil {
		return
	}

	switch k := f.SortKind.(type) {
	case *proto.SortField_Direction:
		sf.Kind = types.SortDirection(k.Direction)
	case *proto.SortField_ComparisonFunctionReference:
		sf.Kind = types.FunctionRef(k.ComparisonFunctionReference)
	default:
		err = substraitgo.ErrNotImplemented
	}
	return
}

// BoundFromProto decodes a window function bound from its protobuf message.
func BoundFromProto(b *proto.Expression_WindowFunction_Bound) expr.Bound {
	if b == nil {
		return nil
	}

	switch t := b.Kind.(type) {
	case *proto.Expression_WindowFunction_Bound_Preceding_:
		return expr.PrecedingBound(t.Preceding.Offset)
	case *proto.Expression_WindowFunction_Bound_CurrentRow_:
		return expr.CurrentRow{}
	case *proto.Expression_WindowFunction_Bound_Following_:
		return expr.FollowingBound(t.Following.Offset)
	case *proto.Expression_WindowFunction_Bound_Unbounded_:
		return expr.Unbounded{}
	}

	// bound is optional
	return nil
}

// AggregateFunctionFromProto decodes an aggregate function from its protobuf message.
func AggregateFunctionFromProto(
	agg *proto.AggregateFunction, baseSchema *types.RecordType, reg expr.ExtensionRegistry,
) (*expr.AggregateFunction, error) {
	if agg.OutputType == nil {
		return nil, fmt.Errorf("%w: missing output type", substraitgo.ErrInvalidExpr)
	}

	var err error
	args := make([]types.FuncArg, len(agg.Arguments))
	for i, a := range agg.Arguments {
		if args[i], err = FuncArgFromProto(a, baseSchema, reg); err != nil {
			return nil, err
		}
	}

	sorts := make([]expr.SortField, len(agg.Sorts))
	for i, s := range agg.Sorts {
		if sorts[i], err = SortFieldFromProto(s, baseSchema, reg); err != nil {
			return nil, err
		}
	}

	id, ok := reg.DecodeFunc(agg.FunctionReference)
	if !ok {
		return nil, substraitgo.ErrNotFound
	}
	decl, ok := reg.LookupAggregateFunction(agg.FunctionReference)
	if !ok {
		return expr.NewCustomAggregateFunc(reg, extensions.NewAggFuncVariant(id), TypeFromProto(agg.OutputType), FunctionOptionsFromProto(agg.Options), types.AggregationInvocation(agg.Invocation), types.AggregationPhase(agg.Phase), sorts, args...)
	}

	return expr.NewAggregateFunctionFromParts(
		agg.FunctionReference,
		decl,
		args,
		FunctionOptionsFromProto(agg.Options),
		TypeFromProto(agg.OutputType),
		types.AggregationPhase(agg.Phase),
		types.AggregationInvocation(agg.Invocation),
		sorts,
	), nil
}
