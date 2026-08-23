// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func TestCastFailBehaviorString(t *testing.T) {
	for _, td := range []struct {
		b        types.CastFailBehavior
		expected string
	}{
		{types.CastFailBehaviorUnspecified, "FAILURE_BEHAVIOR_UNSPECIFIED"},
		{types.CastFailBehaviorReturnNull, "FAILURE_BEHAVIOR_RETURN_NULL"},
		{types.CastFailBehaviorThrowException, "FAILURE_BEHAVIOR_THROW_EXCEPTION"},
	} {
		t.Run(td.expected, func(t *testing.T) {
			assert.Equal(t, td.expected, td.b.String())
			assert.Equal(t, td.expected, fmt.Sprintf("%v", td.b))
		})
	}
}

func TestCastFailBehaviorMatchesProto(t *testing.T) {
	cases := []struct {
		domain types.CastFailBehavior
		pb     proto.Expression_Cast_FailureBehavior
	}{
		{types.CastFailBehaviorUnspecified, proto.Expression_Cast_FAILURE_BEHAVIOR_UNSPECIFIED},
		{types.CastFailBehaviorReturnNull, proto.Expression_Cast_FAILURE_BEHAVIOR_RETURN_NULL},
		{types.CastFailBehaviorThrowException, proto.Expression_Cast_FAILURE_BEHAVIOR_THROW_EXCEPTION},
	}
	// fail if the spec adds a value we don't mirror
	assert.Equal(t, proto.Expression_Cast_FAILURE_BEHAVIOR_UNSPECIFIED.Descriptor().Values().Len(), len(cases),
		"a proto cast failure behavior value is not covered")
	for _, td := range cases {
		assert.EqualValues(t, td.pb, td.domain, "wire number for %s", td.pb)
		assert.Equal(t, td.pb.String(), td.domain.String(), "enum name for %s", td.pb)
	}
}
