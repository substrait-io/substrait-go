package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPrecisionTimestampType(t *testing.T) {
	allPossibleTimePrecision := []TimePrecision{Seconds, DeciSeconds, CentiSeconds, MilliSeconds,
		EMinus4Seconds, EMinus5Seconds, MicroSeconds, EMinus7Seconds, EMinus8Seconds, NanoSeconds}
	allPossibleNullability := []Nullability{NullabilityUnspecified, NullabilityNullable, NullabilityRequired}

	for _, precision := range allPossibleTimePrecision {
		for _, nullability := range allPossibleNullability {
			expectedPrecisionTimeStampType := PrecisionTimeStampType{precision: precision, nullability: nullability}
			expectedPrecisionTimeStampTzType := PrecisionTimeStampTzType{PrecisionTimeStampType: expectedPrecisionTimeStampType}
			assert.Equal(t, expectedPrecisionTimeStampType, NewPrecisionTimestampType(precision).WithNullability(nullability))
			assert.Equal(t, expectedPrecisionTimeStampTzType, NewPrecisionTimestampTzType(precision).WithNullability(nullability))
		}
	}
}
