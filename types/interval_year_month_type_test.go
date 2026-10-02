package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIntervalYearToMonthType(t *testing.T) {
	allPossibleNullability := []Nullability{NullabilityUnspecified, NullabilityNullable, NullabilityRequired}

	for _, nullability := range allPossibleNullability {
		expectedIntervalType := IntervalYearToMonthType{nullability: nullability}

		parameters := expectedIntervalType.GetParameters()
		assert.Len(t, parameters, 0)
		// verify IntervalYearToMonthType
		createdIntervalTypeIfcType := NewIntervalYearToMonthType().WithTypeVariationRef(0).WithNullability(nullability)
		createdIntervalType := createdIntervalTypeIfcType.(IntervalYearToMonthType)
		assert.True(t, createdIntervalType.Equals(expectedIntervalType))
		assert.Equal(t, nullability, createdIntervalType.GetNullability())
		assert.Zero(t, createdIntervalTypeIfcType.GetTypeVariationReference())
		assert.Equal(t, fmt.Sprintf("interval_year%s", strNullable(expectedIntervalType)), createdIntervalType.String())
		assert.Equal(t, "iyear", createdIntervalType.ShortString())
	}
}
