// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// WithNullability builds a new value, so every field it does not copy is silently dropped.
func TestWithNullabilityKeepsTypeVariationRef(t *testing.T) {
	for _, td := range []struct {
		name string
		in   Type
	}{
		{"IntervalCompoundType", NewIntervalCompoundType().WithPrecision(PrecisionSeconds).WithTypeVariationRef(7)},
		{"IntervalYearToMonthType", NewIntervalYearToMonthType().WithTypeVariationRef(7)},
		{"IntervalDayType", &IntervalDayType{Precision: PrecisionSeconds, TypeVariationRef: 7}},
		{"PrecisionTimeType", &PrecisionTimeType{Precision: PrecisionSeconds, TypeVariationRef: 7}},
		{"PrecisionTimestampType", &PrecisionTimestampType{Precision: PrecisionSeconds, TypeVariationRef: 7}},
		{"PrecisionTimestampTzType", &PrecisionTimestampTzType{
			PrecisionTimestampType: PrecisionTimestampType{Precision: PrecisionSeconds, TypeVariationRef: 7}}},
	} {
		t.Run(td.name, func(t *testing.T) {
			got := td.in.WithNullability(NullabilityRequired).GetTypeVariationReference()
			if got != 7 {
				t.Errorf("type variation reference 7 became %d", got)
			}
		})
	}
}
