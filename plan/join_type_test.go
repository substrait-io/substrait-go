// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/plan"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func TestJoinTypeString(t *testing.T) {
	for _, td := range []struct {
		j        plan.JoinType
		expected string
	}{
		{plan.JoinTypeUnspecified, "JOIN_TYPE_UNSPECIFIED"},
		{plan.JoinTypeInner, "JOIN_TYPE_INNER"},
		{plan.JoinTypeOuter, "JOIN_TYPE_OUTER"},
		{plan.JoinTypeLeft, "JOIN_TYPE_LEFT"},
		{plan.JoinTypeRight, "JOIN_TYPE_RIGHT"},
		{plan.JoinTypeLeftSemi, "JOIN_TYPE_LEFT_SEMI"},
		{plan.JoinTypeLeftAnti, "JOIN_TYPE_LEFT_ANTI"},
		{plan.JoinTypeLeftSingle, "JOIN_TYPE_LEFT_SINGLE"},
		{plan.JoinTypeRightSemi, "JOIN_TYPE_RIGHT_SEMI"},
		{plan.JoinTypeRightAnti, "JOIN_TYPE_RIGHT_ANTI"},
		{plan.JoinTypeRightSingle, "JOIN_TYPE_RIGHT_SINGLE"},
		{plan.JoinTypeLeftMark, "JOIN_TYPE_LEFT_MARK"},
		{plan.JoinTypeRightMark, "JOIN_TYPE_RIGHT_MARK"},
	} {
		t.Run(td.expected, func(t *testing.T) {
			assert.Equal(t, td.expected, td.j.String())
			assert.Equal(t, td.expected, fmt.Sprintf("%v", td.j))
		})
	}
}

func TestJoinTypeMatchesProto(t *testing.T) {
	cases := []struct {
		domain plan.JoinType
		pb     proto.JoinRel_JoinType
	}{
		{plan.JoinTypeUnspecified, proto.JoinRel_JOIN_TYPE_UNSPECIFIED},
		{plan.JoinTypeInner, proto.JoinRel_JOIN_TYPE_INNER},
		{plan.JoinTypeOuter, proto.JoinRel_JOIN_TYPE_OUTER},
		{plan.JoinTypeLeft, proto.JoinRel_JOIN_TYPE_LEFT},
		{plan.JoinTypeRight, proto.JoinRel_JOIN_TYPE_RIGHT},
		{plan.JoinTypeLeftSemi, proto.JoinRel_JOIN_TYPE_LEFT_SEMI},
		{plan.JoinTypeLeftAnti, proto.JoinRel_JOIN_TYPE_LEFT_ANTI},
		{plan.JoinTypeLeftSingle, proto.JoinRel_JOIN_TYPE_LEFT_SINGLE},
		{plan.JoinTypeRightSemi, proto.JoinRel_JOIN_TYPE_RIGHT_SEMI},
		{plan.JoinTypeRightAnti, proto.JoinRel_JOIN_TYPE_RIGHT_ANTI},
		{plan.JoinTypeRightSingle, proto.JoinRel_JOIN_TYPE_RIGHT_SINGLE},
		{plan.JoinTypeLeftMark, proto.JoinRel_JOIN_TYPE_LEFT_MARK},
		{plan.JoinTypeRightMark, proto.JoinRel_JOIN_TYPE_RIGHT_MARK},
	}
	// fail if the spec adds a value we don't mirror
	assert.Equal(t, proto.JoinRel_JOIN_TYPE_UNSPECIFIED.Descriptor().Values().Len(), len(cases),
		"a proto join type value is not covered")
	for _, td := range cases {
		assert.EqualValues(t, td.pb, td.domain, "wire number for %s", td.pb)
		assert.Equal(t, td.pb.String(), td.domain.String(), "enum name for %s", td.pb)
	}
}
