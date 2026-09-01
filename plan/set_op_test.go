// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/plan"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func TestSetOpString(t *testing.T) {
	for _, td := range []struct {
		o        plan.SetOp
		expected string
	}{
		{plan.SetOpUnspecified, "SET_OP_UNSPECIFIED"},
		{plan.SetOpMinusPrimary, "SET_OP_MINUS_PRIMARY"},
		{plan.SetOpMinusMultiset, "SET_OP_MINUS_MULTISET"},
		{plan.SetOpIntersectionPrimary, "SET_OP_INTERSECTION_PRIMARY"},
		{plan.SetOpIntersectionMultiset, "SET_OP_INTERSECTION_MULTISET"},
		{plan.SetOpUnionDistinct, "SET_OP_UNION_DISTINCT"},
		{plan.SetOpUnionAll, "SET_OP_UNION_ALL"},
		{plan.SetOpMinusPrimaryAll, "SET_OP_MINUS_PRIMARY_ALL"},
		{plan.SetOpIntersectionMultisetAll, "SET_OP_INTERSECTION_MULTISET_ALL"},
	} {
		t.Run(td.expected, func(t *testing.T) {
			assert.Equal(t, td.expected, td.o.String())
			assert.Equal(t, td.expected, fmt.Sprintf("%v", td.o))
		})
	}
}

func TestSetOpMatchesProto(t *testing.T) {
	cases := []struct {
		domain plan.SetOp
		pb     proto.SetRel_SetOp
	}{
		{plan.SetOpUnspecified, proto.SetRel_SET_OP_UNSPECIFIED},
		{plan.SetOpMinusPrimary, proto.SetRel_SET_OP_MINUS_PRIMARY},
		{plan.SetOpMinusMultiset, proto.SetRel_SET_OP_MINUS_MULTISET},
		{plan.SetOpIntersectionPrimary, proto.SetRel_SET_OP_INTERSECTION_PRIMARY},
		{plan.SetOpIntersectionMultiset, proto.SetRel_SET_OP_INTERSECTION_MULTISET},
		{plan.SetOpUnionDistinct, proto.SetRel_SET_OP_UNION_DISTINCT},
		{plan.SetOpUnionAll, proto.SetRel_SET_OP_UNION_ALL},
		{plan.SetOpMinusPrimaryAll, proto.SetRel_SET_OP_MINUS_PRIMARY_ALL},
		{plan.SetOpIntersectionMultisetAll, proto.SetRel_SET_OP_INTERSECTION_MULTISET_ALL},
	}
	// fail if the spec adds a value we don't mirror
	assert.Equal(t, proto.SetRel_SET_OP_UNSPECIFIED.Descriptor().Values().Len(), len(cases),
		"a proto set operation value is not covered")
	for _, td := range cases {
		assert.EqualValues(t, td.pb, td.domain, "wire number for %s", td.pb)
		assert.Equal(t, td.pb.String(), td.domain.String(), "enum name for %s", td.pb)
	}
}
