// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/plan"
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
