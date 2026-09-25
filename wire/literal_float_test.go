// SPDX-License-Identifier: Apache-2.0

package wire_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/literal"
	"github.com/substrait-io/substrait-go/v9/wire"
)

func TestNewFloat32(t *testing.T) {
	tests := []struct {
		name  string
		value float32
		want  expr.Literal
	}{
		{"0", 0, expr.NewPrimitiveLiteral[float32](0, false)},
		{"1.1", 1.1, expr.NewPrimitiveLiteral[float32](1.1, false)},
		{"-1.1", -1.1, expr.NewPrimitiveLiteral[float32](-1.1, false)},
		{"NaN", float32(math.NaN()), expr.NewPrimitiveLiteral[float32](float32(math.NaN()), false)},
		{"+Inf", float32(math.Inf(1)), expr.NewPrimitiveLiteral[float32](float32(math.Inf(1)), false)},
		{"-Inf", float32(math.Inf(-1)), expr.NewPrimitiveLiteral[float32](float32(math.Inf(-1)), false)},
		{"max float32", float32(math.MaxFloat32), expr.NewPrimitiveLiteral[float32](math.MaxFloat32, false)},
		{"min float32", float32(math.SmallestNonzeroFloat32), expr.NewPrimitiveLiteral[float32](math.SmallestNonzeroFloat32, false)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := literal.NewFloat32(tt.value, false)
			if !math.IsNaN(float64(tt.value)) {
				assert.Equalf(t, tt.want, got, "literal.NewFloat32(%v)", tt.value)
			} else {
				protoExp := wire.ExprToProto(got)
				assert.True(t, math.IsNaN(float64(protoExp.GetLiteral().GetFp32())), "literal.NewFloat32(%v)", tt.value)
			}
		})
	}
}

func TestNewFloat64(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  expr.Literal
	}{
		{"0", 0, expr.NewPrimitiveLiteral[float64](0, false)},
		{"1.1", 1.1, expr.NewPrimitiveLiteral[float64](1.1, false)},
		{"-1.1", -1.1, expr.NewPrimitiveLiteral[float64](-1.1, false)},
		{"NaN", math.NaN(), expr.NewPrimitiveLiteral[float64](math.NaN(), false)},
		{"+Inf", math.Inf(1), expr.NewPrimitiveLiteral[float64](math.Inf(1), false)},
		{"-Inf", math.Inf(-1), expr.NewPrimitiveLiteral[float64](math.Inf(-1), false)},
		{"max float64", math.MaxFloat64, expr.NewPrimitiveLiteral[float64](math.MaxFloat64, false)},
		{"min float64", math.SmallestNonzeroFloat64, expr.NewPrimitiveLiteral[float64](math.SmallestNonzeroFloat64, false)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := literal.NewFloat64(tt.value, false)
			if !math.IsNaN(tt.value) {
				assert.Equalf(t, tt.want, got, "literal.NewFloat64(%v)", tt.value)
			} else {
				protoExp := wire.ExprToProto(got)
				assert.True(t, math.IsNaN(protoExp.GetLiteral().GetFp64()), "literal.NewFloat64(%v)", tt.value)
			}
		})
	}
}
