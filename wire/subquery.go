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
