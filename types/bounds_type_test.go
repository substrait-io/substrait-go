// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/types"
)

func TestBoundsTypeString(t *testing.T) {
	for _, td := range []struct {
		b        types.BoundsType
		expected string
	}{
		{types.BoundsTypeUnspecified, "BOUNDS_TYPE_UNSPECIFIED"},
		{types.BoundsTypeRows, "BOUNDS_TYPE_ROWS"},
		{types.BoundsTypeRange, "BOUNDS_TYPE_RANGE"},
		{types.BoundsType(99), "99"},
	} {
		t.Run(td.expected, func(t *testing.T) {
			assert.Equal(t, td.expected, td.b.String())
			assert.Equal(t, td.expected, fmt.Sprintf("%v", td.b))
		})
	}
}
