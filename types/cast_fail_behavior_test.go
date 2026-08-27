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

