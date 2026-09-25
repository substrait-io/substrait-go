// SPDX-License-Identifier: Apache-2.0

package wire_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/plan"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func TestSetComparisonOpMatchesProto(t *testing.T) {
	cases := []struct {
		domain plan.SetComparisonOp
		pb     proto.Expression_Subquery_SetComparison_ComparisonOp
	}{
		{plan.SetComparisonOpUnspecified, proto.Expression_Subquery_SetComparison_COMPARISON_OP_UNSPECIFIED},
		{plan.SetComparisonOpEq, proto.Expression_Subquery_SetComparison_COMPARISON_OP_EQ},
		{plan.SetComparisonOpNe, proto.Expression_Subquery_SetComparison_COMPARISON_OP_NE},
		{plan.SetComparisonOpLt, proto.Expression_Subquery_SetComparison_COMPARISON_OP_LT},
		{plan.SetComparisonOpGt, proto.Expression_Subquery_SetComparison_COMPARISON_OP_GT},
		{plan.SetComparisonOpLe, proto.Expression_Subquery_SetComparison_COMPARISON_OP_LE},
		{plan.SetComparisonOpGe, proto.Expression_Subquery_SetComparison_COMPARISON_OP_GE},
	}
	// fail if the spec adds a value we don't mirror
	assert.Equal(t, proto.Expression_Subquery_SetComparison_COMPARISON_OP_UNSPECIFIED.Descriptor().Values().Len(), len(cases),
		"a proto set comparison comparison op value is not covered")
	for _, td := range cases {
		assert.EqualValues(t, td.pb, td.domain, "wire number for %s", td.pb)
		assert.Equal(t, td.pb.String(), td.domain.String(), "enum name for %s", td.pb)
	}
}
