// SPDX-License-Identifier: Apache-2.0

package parser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/types/parser"
)

func TestParser(t *testing.T) {
	tests := []struct {
		expr     string
		expected string
	}{
		{"2", "2"},
		{"-2", "-2"},
		{"i16?", "i16?"},
		{"boolean", "boolean"},
		{"fixedchar<5>", "fixedchar<5>"},
		{"decimal<10,5>", "decimal<10, 5>"},
		{"list<decimal<10,5>>", "list<decimal<10, 5>>"},
		{"list?<decimal?<10,5>>", "list?<decimal?<10, 5>>"},
		{"struct<i16?,i32>", "struct<i16?, i32>"},
		{"map<boolean?,struct?<i16?,i32?,i64?>>", "map<boolean?,struct?<i16?, i32?, i64?>>"},
		{"map?<boolean?,struct?<i16?,i32?,i64?>>", "map?<boolean?,struct?<i16?, i32?, i64?>>"},
	}

	p, err := parser.New()
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.expr, func(t *testing.T) {
			d, err := p.ParseString(tt.expr)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, d.Expr.String())
		})
	}
}
