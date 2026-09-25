// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"fmt"

	"github.com/substrait-io/substrait-go/v9/expr"
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
