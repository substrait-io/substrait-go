package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntervalDayType(t *testing.T) {
	anotherType := &FixedCharType{Length: 10, Nullability: NullabilityNullable}
	allPossibleTimePrecision := []TimePrecision{PrecisionSeconds, PrecisionDeciSeconds, PrecisionCentiSeconds, PrecisionMilliSeconds,
		PrecisionEMinus4Seconds, PrecisionEMinus5Seconds, PrecisionMicroSeconds, PrecisionEMinus7Seconds, PrecisionEMinus8Seconds, PrecisionNanoSeconds}
	allPossibleNullability := []Nullability{NullabilityUnspecified, NullabilityNullable, NullabilityRequired}

	for _, precision := range allPossibleTimePrecision {
		for _, nullability := range allPossibleNullability {
			expectedIntervalDayType := &IntervalDayType{Precision: precision, Nullability: nullability}
			expectedFormatString := fmt.Sprintf("%s<%d>", strNullable(expectedIntervalDayType), precision.ToProtoVal())

			parameters := expectedIntervalDayType.GetParameters()
			assert.Equal(t, parameters, []interface{}{precision})
			// verify IntervalDayType
			createdIntervalDayTypeIfc := (&IntervalDayType{Precision: precision}).WithNullability(nullability)
			createdIntervalDayType := createdIntervalDayTypeIfc.(*IntervalDayType)
			assert.True(t, createdIntervalDayType.Equals(expectedIntervalDayType))
			assert.Equal(t, expectedProtoValMap[precision], createdIntervalDayType.GetPrecisionProtoVal())
			assert.Equal(t, nullability, createdIntervalDayType.GetNullability())
			assert.Zero(t, createdIntervalDayType.GetTypeVariationReference())
			assert.Equal(t, fmt.Sprintf("interval_day%s", expectedFormatString), createdIntervalDayType.String())
			assert.Equal(t, "iday", createdIntervalDayType.ShortString())
			assert.Equal(t, "interval_day", createdIntervalDayType.BaseString())
			assert.Equal(t, precision, createdIntervalDayType.GetPrecision())
			expectedParameterString := fmt.Sprintf("%d", precision.ToProtoVal())
			assert.Equal(t, expectedParameterString, createdIntervalDayType.ParameterString())
			assert.False(t, createdIntervalDayTypeIfc.Equals(anotherType))
		}
	}
}
