// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"fmt"

	substraitgo "github.com/substrait-io/substrait-go/v9"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/plan"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// subqueryFromProto decodes a subquery expression from its protobuf message,
// decoding the embedded relations via RelFromProto directly.
func subqueryFromProto(sub *proto.Expression_Subquery, baseSchema *types.RecordType, reg expr.ExtensionRegistry) (expr.Expression, error) {
	switch subType := sub.SubqueryType.(type) {
	case *proto.Expression_Subquery_Scalar_:
		rel, err := RelFromProto(subType.Scalar.Input, reg)
		if err != nil {
			return nil, err
		}
		return plan.NewScalarSubquery(rel), nil

	case *proto.Expression_Subquery_InPredicate_:
		needles := make([]expr.Expression, len(subType.InPredicate.Needles))
		for i, needle := range subType.InPredicate.Needles {
			e, err := ExprFromProto(needle, baseSchema, reg)
			if err != nil {
				return nil, fmt.Errorf("error parsing needle %d in IN predicate: %w", i, err)
			}
			needles[i] = e
		}

		rel, err := RelFromProto(subType.InPredicate.Haystack, reg)
		if err != nil {
			return nil, err
		}

		return plan.NewInPredicateSubquery(needles, rel), nil

	case *proto.Expression_Subquery_SetPredicate_:
		tuples, err := RelFromProto(subType.SetPredicate.Tuples, reg)
		if err != nil {
			return nil, fmt.Errorf("error parsing tuples in set predicate: %w", err)
		}
		return plan.NewSetPredicateSubquery(plan.SetPredicateOp(subType.SetPredicate.PredicateOp), tuples), nil
	case *proto.Expression_Subquery_SetComparison_:
		left, err := ExprFromProto(subType.SetComparison.Left, baseSchema, reg)
		if err != nil {
			return nil, fmt.Errorf("error parsing left expression in set comparison: %w", err)
		}

		right, err := RelFromProto(subType.SetComparison.Right, reg)
		if err != nil {
			return nil, fmt.Errorf("error parsing right relation in set comparison: %w", err)
		}

		return plan.NewSetComparisonSubquery(
			plan.SetComparisonReductionOp(subType.SetComparison.ReductionOp),
			plan.SetComparisonOp(subType.SetComparison.ComparisonOp),
			left,
			right,
		), nil

	default:
		return nil, fmt.Errorf("%w: unknown subquery type: %T", substraitgo.ErrNotImplemented, subType)
	}
}

func scalarSubqueryToProto(s *plan.ScalarSubquery) *proto.Expression {
	return &proto.Expression{
		RexType: &proto.Expression_Subquery_{
			Subquery: &proto.Expression_Subquery{
				SubqueryType: &proto.Expression_Subquery_Scalar_{
					Scalar: &proto.Expression_Subquery_Scalar{
						Input: RelToProto(s.Input),
					},
				},
			},
		},
	}
}

func inPredicateSubqueryToProto(s *plan.InPredicateSubquery) *proto.Expression {
	needles := make([]*proto.Expression, len(s.Needles))
	for i, needle := range s.Needles {
		needles[i] = ExprToProto(needle)
	}

	return &proto.Expression{
		RexType: &proto.Expression_Subquery_{
			Subquery: &proto.Expression_Subquery{
				SubqueryType: &proto.Expression_Subquery_InPredicate_{
					InPredicate: &proto.Expression_Subquery_InPredicate{
						Needles:  needles,
						Haystack: RelToProto(s.Haystack),
					},
				},
			},
		},
	}
}

func setPredicateSubqueryToProto(s *plan.SetPredicateSubquery) *proto.Expression {
	return &proto.Expression{
		RexType: &proto.Expression_Subquery_{
			Subquery: &proto.Expression_Subquery{
				SubqueryType: &proto.Expression_Subquery_SetPredicate_{
					SetPredicate: &proto.Expression_Subquery_SetPredicate{
						PredicateOp: proto.Expression_Subquery_SetPredicate_PredicateOp(s.Operation),
						Tuples:      RelToProto(s.Tuples),
					},
				},
			},
		},
	}
}

func setComparisonSubqueryToProto(s *plan.SetComparisonSubquery) *proto.Expression {
	return &proto.Expression{
		RexType: &proto.Expression_Subquery_{
			Subquery: &proto.Expression_Subquery{
				SubqueryType: &proto.Expression_Subquery_SetComparison_{
					SetComparison: &proto.Expression_Subquery_SetComparison{
						ReductionOp:  proto.Expression_Subquery_SetComparison_ReductionOp(s.ReductionOp),
						ComparisonOp: proto.Expression_Subquery_SetComparison_ComparisonOp(s.ComparisonOp),
						Left:         ExprToProto(s.Left),
						Right:        RelToProto(s.Right),
					},
				},
			},
		},
	}
}
