// SPDX-License-Identifier: Apache-2.0

package substraitgo_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/substrait-io/substrait-go"
)

func TestTypeToString(t *testing.T) {
	tests := []struct {
		t   Type
		exp string
	}{
		{&BooleanType{}, "boolean"},
		{&Int8Type{}, "i8"},
		{&Int16Type{}, "i16"},
		{&Int32Type{}, "i32"},
		{&Int64Type{}, "i64"},
		{&Float32Type{}, "fp32"},
		{&Float64Type{}, "fp64"},
		{&BinaryType{}, "binary"},
		{&StringType{}, "string"},
		{&TimestampType{}, "timestamp"},
		{&DateType{}, "date"},
		{&TimeType{}, "time"},
		{&TimestampTzType{}, "timestamp_tz"},
		{&IntervalYearType{}, "interval_year"},
	}

	for _, tt := range tests {
		t.Run(tt.exp, func(t *testing.T) {
			assert.Equal(t, tt.exp, tt.t.String())
		})
	}
}

func TestLiteralToString(t *testing.T) {
	tests := []struct {
		t   Literal
		exp string
	}{
		{NewPrimitiveLiteral[int8](0, false), "i8(0)"},
	}

	for _, tt := range tests {
		t.Run(tt.exp, func(t *testing.T) {
			assert.Equal(t, tt.exp, tt.t.String())
		})
	}
}
