// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntervalDayWithNullabilityCopiesType(t *testing.T) {
	original := &IntervalDayType{
		Precision: PrecisionMicroSeconds, Nullability: NullabilityRequired, TypeVariationRef: 7,
	}
	got := original.WithNullability(NullabilityNullable)
	assert.NotSame(t, original, got)
	assert.Equal(t, NullabilityRequired, original.Nullability)
	assert.Equal(t, &IntervalDayType{
		Precision: PrecisionMicroSeconds, Nullability: NullabilityNullable, TypeVariationRef: 7,
	}, got)
}
