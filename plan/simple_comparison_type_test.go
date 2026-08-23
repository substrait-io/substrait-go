// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/plan"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func TestSimpleComparisonTypeString(t *testing.T) {
	for _, td := range []struct {
		c        plan.SimpleComparisonType
		expected string
	}{
		{plan.SimpleComparisonTypeUnspecified, "SIMPLE_COMPARISON_TYPE_UNSPECIFIED"},
		{plan.SimpleComparisonTypeEq, "SIMPLE_COMPARISON_TYPE_EQ"},
		{plan.SimpleComparisonTypeIsNotDistinctFrom, "SIMPLE_COMPARISON_TYPE_IS_NOT_DISTINCT_FROM"},
		{plan.SimpleComparisonTypeMightEqual, "SIMPLE_COMPARISON_TYPE_MIGHT_EQUAL"},
	} {
		t.Run(td.expected, func(t *testing.T) {
			assert.Equal(t, td.expected, td.c.String())
			assert.Equal(t, td.expected, fmt.Sprintf("%v", td.c))
		})
	}
}

func TestSimpleComparisonTypeMatchesProto(t *testing.T) {
	cases := []struct {
		domain plan.SimpleComparisonType
		pb     proto.ComparisonJoinKey_SimpleComparisonType
	}{
		{plan.SimpleComparisonTypeUnspecified, proto.ComparisonJoinKey_SIMPLE_COMPARISON_TYPE_UNSPECIFIED},
		{plan.SimpleComparisonTypeEq, proto.ComparisonJoinKey_SIMPLE_COMPARISON_TYPE_EQ},
		{plan.SimpleComparisonTypeIsNotDistinctFrom, proto.ComparisonJoinKey_SIMPLE_COMPARISON_TYPE_IS_NOT_DISTINCT_FROM},
		{plan.SimpleComparisonTypeMightEqual, proto.ComparisonJoinKey_SIMPLE_COMPARISON_TYPE_MIGHT_EQUAL},
	}
	// fail if the spec adds a value we don't mirror
	assert.Equal(t, proto.ComparisonJoinKey_SIMPLE_COMPARISON_TYPE_UNSPECIFIED.Descriptor().Values().Len(), len(cases),
		"a proto simple comparison type value is not covered")
	for _, td := range cases {
		assert.EqualValues(t, td.pb, td.domain, "wire number for %s", td.pb)
		assert.Equal(t, td.pb.String(), td.domain.String(), "enum name for %s", td.pb)
	}
}
