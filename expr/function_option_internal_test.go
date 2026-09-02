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

// TestScalarFunctionEqualsOptions checks that ScalarFunction.Equals compares
// FunctionOptions by name and preference, treating nil elements as equal.
func TestScalarFunctionEqualsOptions(t *testing.T) {
	ot := &types.Int64Type{Nullability: types.NullabilityRequired}
	base := &ScalarFunction{funcRef: 1, outputType: ot, options: opts("ERROR")}
	same := &ScalarFunction{funcRef: 1, outputType: ot, options: opts("ERROR")}
	diffPref := &ScalarFunction{funcRef: 1, outputType: ot, options: opts("SILENT")}
	noOpts := &ScalarFunction{funcRef: 1, outputType: ot}

	assert.True(t, base.Equals(same))
	assert.False(t, base.Equals(diffPref))
	assert.False(t, base.Equals(noOpts))

	// nil option elements must compare equal without dereferencing.
	nilA := &ScalarFunction{funcRef: 1, outputType: ot, options: []*types.FunctionOption{nil}}
	nilB := &ScalarFunction{funcRef: 1, outputType: ot, options: []*types.FunctionOption{nil}}
	assert.True(t, nilA.Equals(nilB))
	assert.False(t, nilA.Equals(base))
}

// TestFunctionGetOption covers GetOption hit/miss on both scalar and aggregate.
func TestFunctionGetOption(t *testing.T) {
	o := opts("ERROR")
	s := &ScalarFunction{options: o}
	assert.Equal(t, []string{"ERROR"}, s.GetOption("overflow"))
	assert.Nil(t, s.GetOption("missing"))

	a := &AggregateFunction{options: o}
	assert.Equal(t, []string{"ERROR"}, a.GetOption("overflow"))
	assert.Nil(t, a.GetOption("missing"))
}
