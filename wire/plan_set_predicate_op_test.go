// SPDX-License-Identifier: Apache-2.0

package wire_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/plan"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

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
