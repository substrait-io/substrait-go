// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/plan"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func TestSetPredicateOpString(t *testing.T) {
	for _, td := range []struct {
		o        plan.SetPredicateOp
		expected string
	}{
		{plan.SetPredicateOpUnspecified, "PREDICATE_OP_UNSPECIFIED"},
		{plan.SetPredicateOpExists, "PREDICATE_OP_EXISTS"},
		{plan.SetPredicateOpUnique, "PREDICATE_OP_UNIQUE"},
		{plan.SetPredicateOp(99), "99"},
	} {
		t.Run(td.expected, func(t *testing.T) {
			assert.Equal(t, td.expected, td.o.String())
			assert.Equal(t, td.expected, fmt.Sprintf("%v", td.o))
		})
	}
}

func TestSetPredicateOpMatchesProto(t *testing.T) {
	cases := []struct {
		domain plan.SetPredicateOp
		pb     proto.Expression_Subquery_SetPredicate_PredicateOp
	}{
		{plan.SetPredicateOpUnspecified, proto.Expression_Subquery_SetPredicate_PREDICATE_OP_UNSPECIFIED},
		{plan.SetPredicateOpExists, proto.Expression_Subquery_SetPredicate_PREDICATE_OP_EXISTS},
		{plan.SetPredicateOpUnique, proto.Expression_Subquery_SetPredicate_PREDICATE_OP_UNIQUE},
	}
	// fail if the spec adds a value we don't mirror
	assert.Equal(t, proto.Expression_Subquery_SetPredicate_PREDICATE_OP_UNSPECIFIED.Descriptor().Values().Len(), len(cases),
		"a proto set predicate op value is not covered")
	for _, td := range cases {
		assert.EqualValues(t, td.pb, td.domain, "wire number for %s", td.pb)
		assert.Equal(t, td.pb.String(), td.domain.String(), "enum name for %s", td.pb)
	}
}
