// SPDX-License-Identifier: Apache-2.0

package expr

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/types"
)

func opts(pref string) []*types.FunctionOption {
	return []*types.FunctionOption{{Name: "overflow", Preference: []string{pref}}}
}

func TestScalarFunctionEqualsOptions(t *testing.T) {
	outputType := &types.Int64Type{Nullability: types.NullabilityRequired}
	left := &ScalarFunction{funcRef: 1, outputType: outputType, options: opts("ERROR")}
	right := &ScalarFunction{funcRef: 1, outputType: outputType, options: opts("ERROR")}

	assert.True(t, left.Equals(right))
}
