// SPDX-License-Identifier: Apache-2.0

package expr

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/types"
)

// TestWindowFunctionEqualsNilOption checks that WindowFunction.Equals treats a nil
// FunctionOption element as equal without dereferencing it.
func TestWindowFunctionEqualsNilOption(t *testing.T) {
	ot := &types.Int64Type{Nullability: types.NullabilityRequired}
	a := &WindowFunction{funcRef: 1, outputType: ot, options: []*types.FunctionOption{nil}}
	b := &WindowFunction{funcRef: 1, outputType: ot, options: []*types.FunctionOption{nil}}

	assert.True(t, a.Equals(b))
}
