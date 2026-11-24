// SPDX-License-Identifier: Apache-2.0

package parser

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseUnsupportedTypeCallers(t *testing.T) {
	cases := []struct {
		name      string
		aggregate bool
		input     string
	}{
		{"scalar argument", false, "f1(null::TYPE) = 1::i8"},
		{"scalar result", false, "f1(1::i8) = null::TYPE"},
		{"list argument", false, "f1([]::list<TYPE>) = 1::i8"},
		{"nested list argument", false, "f1([]::list<list<TYPE>>) = 1::i8"},
		{"list result", false, "f1(1::i8) = []::list<TYPE>"},
		{"null list argument", false, "f1(null::list?<TYPE>) = 1::i8"},
		{"aggregate result", true, "f1((1, 2)::i8) = null::TYPE"},
		{"compact scalar argument", true, "((1), (2)) f1(col0::i8, null::TYPE) = 1::i8"},
		{"qualified scalar argument", true, "DEFINE t1(i8) = ((1), (2))\nf1(t1.col0, null::TYPE) = 1::i8"},
		{"compact column type", true, "((1), (2)) f1(col0::TYPE) = 1::i8"},
		{"single column type", true, "f1((1, 2)::TYPE) = 1::i8"},
		{"table column type", true, "DEFINE t1(TYPE) = ((1), (2))\nf1(t1.col0) = 1::i8"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			header := makeHeader("v1.0", "/extensions/functions_arithmetic.yaml")
			if tc.aggregate {
				header = makeAggregateTestHeader("v1.0", "/extensions/functions_arithmetic.yaml")
			}
			input := header + "# unsupported type caller\n" + tc.input
			t.Run("supported control", func(t *testing.T) {
				file, err := ParseTestCasesFromString(strings.ReplaceAll(input, "TYPE", "i8?"))
				require.NoError(t, err)
				require.Len(t, file.TestCases, 1)
				assert.Equal(t, "f1", file.TestCases[0].FuncName)
			})
			for _, typ := range []string{"u!geometry", "u!geometry?"} {
				t.Run(typ, func(t *testing.T) {
					file, err := ParseTestCasesFromString(strings.ReplaceAll(input, "TYPE", typ))
					require.Error(t, err)
					assert.Contains(t, err.Error(), "user-defined type")
					assert.NotContains(t, err.Error(), "Tree Visit panic error")
					assert.Nil(t, file)
				})
			}
		})
	}
}

func TestParseOptionalGroupDescription(t *testing.T) {
	for _, aggregate := range []bool{false, true} {
		header := makeHeader("v1.0", "/extensions/functions_arithmetic.yaml")
		input := "f1(1::i8) = 1::i8"
		if aggregate {
			header = makeAggregateTestHeader("v1.0", "/extensions/functions_arithmetic.yaml")
			input = "f1((1, 2)::i8) = 1::i8"
		}
		t.Run(input, func(t *testing.T) {
			file, err := ParseTestCasesFromString(header + input)
			require.NoError(t, err)
			require.Len(t, file.TestCases, 1)
			assert.Empty(t, file.TestCases[0].GroupDesc)
			assert.Equal(t, "f1", file.TestCases[0].FuncName)
		})
	}
}

func TestParseUnsupportedTypeWithoutGroupDescription(t *testing.T) {
	header := makeHeader("v1.0", "/extensions/functions_arithmetic.yaml")
	for _, typ := range []string{"u!geometry", "u!geometry?"} {
		t.Run(typ, func(t *testing.T) {
			file, err := ParseTestCasesFromString(header + "f1(null::" + typ + ") = 1::i8")
			require.Error(t, err)
			assert.Contains(t, err.Error(), "user-defined type")
			assert.NotContains(t, err.Error(), "Tree Visit panic error")
			assert.Nil(t, file)
		})
	}
}
