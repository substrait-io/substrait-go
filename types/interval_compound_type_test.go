package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIntervalCompoundType(t *testing.T) {
	allPossibleTimePrecision := []TimePrecision{PrecisionSeconds, PrecisionDeciSeconds, PrecisionCentiSeconds, PrecisionMilliSeconds,
		PrecisionEMinus4Seconds, PrecisionEMinus5Seconds, PrecisionMicroSeconds, PrecisionEMinus7Seconds, PrecisionEMinus8Seconds, PrecisionNanoSeconds}
	allPossibleNullability := []Nullability{NullabilityUnspecified, NullabilityNullable, NullabilityRequired}

	for _, precision := range allPossibleTimePrecision {
		for _, nullability := range allPossibleNullability {
			expectedIntervalCompoundType := IntervalCompoundType{precision: precision, nullability: nullability}
			expectedFormatString := fmt.Sprintf("%s<%d>", strNullable(expectedIntervalCompoundType), precision.ToProtoVal())

			parameters := expectedIntervalCompoundType.GetParameters()
			assert.Equal(t, parameters, []interface{}{precision})
			// verify IntervalCompoundType
			createdIntervalCompoundTypeIfc := NewIntervalCompoundType().WithPrecision(precision).WithTypeVariationRef(0).WithNullability(nullability)
			createdIntervalCompoundType := createdIntervalCompoundTypeIfc.(IntervalCompoundType)
			assert.True(t, createdIntervalCompoundType.Equals(expectedIntervalCompoundType))
			assert.Equal(t, expectedProtoValMap[precision], createdIntervalCompoundType.GetPrecisionProtoVal())
			assert.Equal(t, nullability, createdIntervalCompoundType.GetNullability())
			assert.Zero(t, createdIntervalCompoundType.GetTypeVariationReference())
			assert.Equal(t, fmt.Sprintf("interval_compound%s", expectedFormatString), createdIntervalCompoundType.String())
			assert.Equal(t, "icompound", createdIntervalCompoundType.ShortString())
		}
	}
}
