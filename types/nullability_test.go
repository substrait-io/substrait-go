// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v8/types"
)

func TestNullabilityString(t *testing.T) {
	for _, td := range []struct {
		n        types.Nullability
		expected string
	}{
		{types.NullabilityUnspecified, "NULLABILITY_UNSPECIFIED"},
		{types.NullabilityNullable, "NULLABILITY_NULLABLE"},
		{types.NullabilityRequired, "NULLABILITY_REQUIRED"},
		// unmodelled value
		{types.Nullability(7), "7"},
	} {
		t.Run(td.expected, func(t *testing.T) {
			assert.Equal(t, td.expected, td.n.String())
			assert.Equal(t, td.expected, fmt.Sprintf("%v", td.n))
		})
	}
}
