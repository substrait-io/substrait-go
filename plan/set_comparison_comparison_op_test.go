// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/plan"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func TestSetComparisonComparisonOpString(t *testing.T) {
	for _, td := range []struct {
		o        plan.SetComparisonComparisonOp
		expected string
	}{
		{plan.SetComparisonComparisonOpUnspecified, "COMPARISON_OP_UNSPECIFIED"},
		{plan.SetComparisonComparisonOpEq, "COMPARISON_OP_EQ"},
		{plan.SetComparisonComparisonOpNe, "COMPARISON_OP_NE"},
		{plan.SetComparisonComparisonOpLt, "COMPARISON_OP_LT"},
		{plan.SetComparisonComparisonOpGt, "COMPARISON_OP_GT"},
		{plan.SetComparisonComparisonOpLe, "COMPARISON_OP_LE"},
		{plan.SetComparisonComparisonOpGe, "COMPARISON_OP_GE"},
		{plan.SetComparisonComparisonOp(99), "99"},
	} {
		t.Run(td.expected, func(t *testing.T) {
			assert.Equal(t, td.expected, td.o.String())
			assert.Equal(t, td.expected, fmt.Sprintf("%v", td.o))
		})
	}
}

func TestSetComparisonComparisonOpMatchesProto(t *testing.T) {
	cases := []struct {
		domain plan.SetComparisonComparisonOp
		pb     proto.Expression_Subquery_SetComparison_ComparisonOp
	}{
		{plan.SetComparisonComparisonOpUnspecified, proto.Expression_Subquery_SetComparison_COMPARISON_OP_UNSPECIFIED},
		{plan.SetComparisonComparisonOpEq, proto.Expression_Subquery_SetComparison_COMPARISON_OP_EQ},
		{plan.SetComparisonComparisonOpNe, proto.Expression_Subquery_SetComparison_COMPARISON_OP_NE},
		{plan.SetComparisonComparisonOpLt, proto.Expression_Subquery_SetComparison_COMPARISON_OP_LT},
		{plan.SetComparisonComparisonOpGt, proto.Expression_Subquery_SetComparison_COMPARISON_OP_GT},
		{plan.SetComparisonComparisonOpLe, proto.Expression_Subquery_SetComparison_COMPARISON_OP_LE},
		{plan.SetComparisonComparisonOpGe, proto.Expression_Subquery_SetComparison_COMPARISON_OP_GE},
	}
	// fail if the spec adds a value we don't mirror
	assert.Equal(t, proto.Expression_Subquery_SetComparison_COMPARISON_OP_UNSPECIFIED.Descriptor().Values().Len(), len(cases),
		"a proto set comparison comparison op value is not covered")
	for _, td := range cases {
		assert.EqualValues(t, td.pb, td.domain, "wire number for %s", td.pb)
		assert.Equal(t, td.pb.String(), td.domain.String(), "enum name for %s", td.pb)
	}
}
