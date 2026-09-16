// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/plan"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func TestSetComparisonReductionOpString(t *testing.T) {
	for _, td := range []struct {
		o        plan.SetComparisonReductionOp
		expected string
	}{
		{plan.SetComparisonReductionOpUnspecified, "REDUCTION_OP_UNSPECIFIED"},
		{plan.SetComparisonReductionOpAny, "REDUCTION_OP_ANY"},
		{plan.SetComparisonReductionOpAll, "REDUCTION_OP_ALL"},
		{plan.SetComparisonReductionOp(99), "99"},
	} {
		t.Run(td.expected, func(t *testing.T) {
			assert.Equal(t, td.expected, td.o.String())
			assert.Equal(t, td.expected, fmt.Sprintf("%v", td.o))
		})
	}
}

func TestSetComparisonReductionOpMatchesProto(t *testing.T) {
	cases := []struct {
		domain plan.SetComparisonReductionOp
		pb     proto.Expression_Subquery_SetComparison_ReductionOp
	}{
		{plan.SetComparisonReductionOpUnspecified, proto.Expression_Subquery_SetComparison_REDUCTION_OP_UNSPECIFIED},
		{plan.SetComparisonReductionOpAny, proto.Expression_Subquery_SetComparison_REDUCTION_OP_ANY},
		{plan.SetComparisonReductionOpAll, proto.Expression_Subquery_SetComparison_REDUCTION_OP_ALL},
	}
	// fail if the spec adds a value we don't mirror
	assert.Equal(t, proto.Expression_Subquery_SetComparison_REDUCTION_OP_UNSPECIFIED.Descriptor().Values().Len(), len(cases),
		"a proto set comparison reduction op value is not covered")
	for _, td := range cases {
		assert.EqualValues(t, td.pb, td.domain, "wire number for %s", td.pb)
		assert.Equal(t, td.pb.String(), td.domain.String(), "enum name for %s", td.pb)
	}
}
