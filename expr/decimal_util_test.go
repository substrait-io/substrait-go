package expr

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecimalStringToBytes(t *testing.T) {
	tests := []struct {
		input        string
		hexWant      string
		expPrecision int32
		expScale     int32
		expected     string
	}{
		{"12345", "39300000000000000000000000000000", 5, 0, ""},
		{"+12345", "39300000000000000000000000000000", 5, 0, "12345"},
		{"-12345", "c7cfffffffffffffffffffffffffffff", 5, 0, ""},
		{"123.45", "39300000000000000000000000000000", 5, 2, ""},
		{"-123.45", "c7cfffffffffffffffffffffffffffff", 5, 2, ""},
		{"4.53", "c5010000000000000000000000000000", 3, 2, ""},
		{"0.5", "05000000000000000000000000000000", 2, 1, ""},
		{"0.25", "19000000000000000000000000000000", 3, 2, ""},
		{"1", "01000000000000000000000000000000", 1, 0, ""},
		{"1.000", "e8030000000000000000000000000000", 4, 3, ""},
		{"-1.0", "f6ffffffffffffffffffffffffffffff", 2, 1, ""},
		{"-1.00", "9cffffffffffffffffffffffffffffff", 3, 2, ""},
		{"-1.000", "18fcffffffffffffffffffffffffffff", 4, 3, ""},
		{"12345.6789", "15cd5b07000000000000000000000000", 9, 4, ""},
		{"12345.67890000", "5004fb711f0100000000000000000000", 13, 8, ""},
		{"0.123", "7b000000000000000000000000000000", 4, 3, ""},
		{"-0.123", "85ffffffffffffffffffffffffffffff", 4, 3, ""},
		{"9223372036854775807", "ffffffffffffff7f0000000000000000", 19, 0, ""},  // Max int64
		{"-9223372036854775808", "0000000000000080ffffffffffffffff", 19, 0, ""}, // Min int64
		{"9223372036854775807.0000", "f0d8ffffffffffff8713000000000000", 23, 4, ""},
		{"-9223372036854775808.00", "0000000000000000ceffffffffffffff", 21, 2, ""},
		{"99999999999999999999999999999999999999", "ffffffff3f228a097ac4865aa84c3b4b", 38, 0, ""},
		{"+99999999999999999999999999999999999999", "ffffffff3f228a097ac4865aa84c3b4b", 38, 0, ""},
		{"-99999999999999999999999999999999999999", "01000000c0dd75f6853b79a557b3c4b4", 38, 0, ""},
		{"0", "00000000000000000000000000000000", 1, 0, ""},
		{"-0", "00000000000000000000000000000000", 1, 0, "0"},
		{"0.0", "00000000000000000000000000000000", 2, 1, ""},
		{"65535", "ffff0000000000000000000000000000", 5, 0, ""},
		{"-65535", "0100ffffffffffffffffffffffffffff", 5, 0, ""},
		{"18446744073709551615", "ffffffffffffffff0000000000000000", 20, 0, ""},  // Max uint64
		{"-18446744073709551616", "0000000000000000ffffffffffffffff", 20, 0, ""}, // Min int64 - 1
		{"12345.6789", "15cd5b07000000000000000000000000", 9, 4, ""},
		{"1234567890123456", "c0ba8a3cd56204000000000000000000", 16, 0, ""},
		{"1234567890123456.78901234", "f2af966ca0101f9b241a000000000000", 24, 8, ""},
		{"1230000000000000", "00e012b1ad5e04000000000000000000", 16, 0, ""},
		{"0.0012345678901234", "f22fce733a0b00000000000000000000", 17, 16, ""},
		{"-0.0012345678901234", "0ed0318cc5f4ffffffffffffffffffff", 17, 16, ""},
		{"123456789012345678901234567890.1234", "f2af967ed05c82de3297ff6fde3c0000", 34, 4, ""},
		{"-1234567890.1234567890", "2ef5e0147356ab54ffffffffffffffff", 20, 10, ""},
		{"1.23e-5", "7b000000000000000000000000000000", 8, 7, "0.0000123"},
		{"1.23e15", "00e012b1ad5e04000000000000000000", 16, 0, "1230000000000000"},
		{"1.23e20", "00000c6d51c8f7aa0600000000000000", 21, 0, "123000000000000000000"},
		{"1.23e35", "00000000cebde644bc05f0425eb01700", 36, 0, "123000000000000000000000000000000000"},
		{"1.23E35", "00000000cebde644bc05f0425eb01700", 36, 0, "123000000000000000000000000000000000"},
		{"-123456789012345678901234567890.1234", "0e5069812fa37d21cd68009021c3ffff", 34, 4, ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, precision, scale, err := DecimalStringToBytes(tt.input)
			require.NoError(t, err)
			assert.Len(t, got, 16)
			assert.Equal(t, hexToBytes(t, tt.hexWant), got[:])
			assert.Equal(t, tt.expPrecision, precision)
			assert.Equal(t, tt.expScale, scale)

			// verify that the conversion is correct
			decStr := decimalBytesToString(got, scale)
			if tt.expected == "" {
				tt.expected = strings.TrimPrefix(tt.input, "+")
			}
			assert.Equal(t, tt.expected, decStr)

			// test modifyDecimalPrecisionAndScale
			targetPrecision := min(precision+2, 38)
			targetScale := scale
			if precision <= 36 {
				targetScale = min(scale+2, targetPrecision)
			}
			newBytes, newPrecision, newScale, err := modifyDecimalPrecisionAndScale(got, scale, targetPrecision, targetScale)
			require.NoError(t, err)
			assert.Equal(t, targetPrecision, newPrecision)
			decStr = decimalBytesToString(newBytes, newScale)
			if tt.expected != decStr {
				require.True(t, strings.HasPrefix(decStr, tt.expected))
				suffix := decStr[len(tt.expected):]
				assert.LessOrEqual(t, len(suffix), 3)
				assert.NotEqual(t, 1, len(suffix))
				switch len(suffix) {
				case 2:
					assert.Equal(t, "00", suffix)
				case 3:
					assert.Equal(t, ".00", suffix)
				}
			}
		})
	}
}

func TestDecimalStringToBytesErrors(t *testing.T) {
	badInputs := []struct{ input string }{
		{"12345678901234567890123456789012345678901234"},
		{"abc"},
		{"12.34.56"},
		{"199999999999999999999999999999999999999"},
		{"1.23e45"},
		{"1.23E300"},
	}
	for _, tt := range badInputs {
		t.Run(tt.input, func(t *testing.T) {
			_, _, _, err := DecimalStringToBytes(tt.input)
			assert.Error(t, err, "DecimalStringToBytes(%v) expected error", tt.input)
		})
	}
}

func hexToBytes(t *testing.T, input string) []byte {
	bytes := make([]byte, len(input)/2)
	for i := 0; i < len(input); i += 2 {
		_, err := fmt.Sscanf(input[i:i+2], "%02x", &bytes[i/2])
		assert.NoError(t, err)
	}
	return bytes
}

func TestDecimalBytesToString(t *testing.T) {
	tests := []struct {
		haveHex       string
		havePrecision int32
		haveScale     int32
		want          string
		wantErr       bool
	}{
		{"39300000000000000000000000000000", 5, 0, "12345", false},
		{"c7cfffffffffffffffffffffffffffff", 5, 0, "-12345", false},
		{"39300000000000000000000000000000", 5, 2, "123.45", false},
		{"c7cfffffffffffffffffffffffffffff", 5, 2, "-123.45", false},
		{"7b000000000000000000000000000000", 4, 3, "0.123", false},
		{"85ffffffffffffffffffffffffffffff", 4, 3, "-0.123", false},
		{"ffffffffffffff7f0000000000000000", 19, 0, "9223372036854775807", false},  // Max int64
		{"0000000000000080ffffffffffffffff", 19, 0, "-9223372036854775808", false}, // Min int64
		{"ffffffff3f228a097ac4865aa84c3b4b", 38, 0, "99999999999999999999999999999999999999", false},
		{"01000000c0dd75f6853b79a557b3c4b4", 38, 0, "-99999999999999999999999999999999999999", false},
		{"00000000000000000000000000000000", 1, 0, "0", false},
		{"00000000000000000000000000000000", 2, 1, "0.0", false},
		{"ffff0000000000000000000000000000", 5, 0, "65535", false},
		{"0100ffffffffffffffffffffffffffff", 5, 0, "-65535", false},
		{"ffffffffffffffff0000000000000000", 20, 0, "18446744073709551615", false},  // Max uint64
		{"0000000000000000ffffffffffffffff", 20, 0, "-18446744073709551616", false}, // Min int64 - 1
		{"15cd5b07000000000000000000000000", 9, 4, "12345.6789", false},
		{"c0ba8a3cd56204000000000000000000", 16, 0, "1234567890123456", false},
		{"f2af966ca0101f9b241a000000000000", 24, 8, "1234567890123456.78901234", false},
		{"00e012b1ad5e04000000000000000000", 16, 0, "1230000000000000", false},
		{"f22fce733a0b00000000000000000000", 17, 16, "0.0012345678901234", false},
		{"0ed0318cc5f4ffffffffffffffffffff", 17, 16, "-0.0012345678901234", false},
		{"f2af967ed05c82de3297ff6fde3c0000", 34, 4, "123456789012345678901234567890.1234", false},
		{"2ef5e0147356ab54ffffffffffffffff", 20, 10, "-1234567890.1234567890", false},
		{"7b000000000000000000000000000000", 8, 7, "0.0000123", false},
		{"00e012b1ad5e04000000000000000000", 16, 0, "1230000000000000", false},
		{"00000c6d51c8f7aa0600000000000000", 21, 0, "123000000000000000000", false},
		{"00000000cebde644bc05f0425eb01700", 36, 0, "123000000000000000000000000000000000", false},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			wantBytes := hexToBytes(t, tt.haveHex)
			str := decimalBytesToString([16]byte(wantBytes), tt.haveScale)
			assert.Equal(t, tt.want, str)
		})
	}
}

func TestModifyDecimalPrecisionAndScale(t *testing.T) {
	tests := []struct {
		input           string
		hexInput        string
		inputPrecision  int32
		inputScale      int32
		targetPrecision int32
		targetScale     int32
		hexWant2        string
		expected        string
		expectError     bool
	}{
		{"12345", "39300000000000000000000000000000", 5, 0, 10, 0, "", "", false},
		{"12345", "39300000000000000000000000000000", 5, 0, 20, 2, "44D61200000000000000000000000000", "12345.00", false},
		{"12345.00", "44D61200000000000000000000000000", 20, 2, 5, 2, "", "12345.00", true},
		{"12345.00", "44D61200000000000000000000000000", 20, 2, 5, 0, "39300000000000000000000000000000", "12345", false},
		{"12345.6789", "15cd5b07000000000000000000000000", 9, 4, 12, 8, "15cd5b07000000000000000000000000", "12345.67890000", true},
		{"12345.6789", "15cd5b07000000000000000000000000", 9, 4, 13, 8, "5004fb711f0100000000000000000000", "12345.67890000", false},
		{"-1.00", "9cffffffffffffffffffffffffffffff", 3, 2, 5, 3, "18fcffffffffffffffffffffffffffff", "-1.000", false},
		{"-1.0", "f6ffffffffffffffffffffffffffffff", 2, 1, 2, 0, "ffffffffffffffffffffffffffffffff", "-1", false},
		{"-1.0", "f6ffffffffffffffffffffffffffffff", 2, 1, 1, 0, "ffffffffffffffffffffffffffffffff", "-1", false},
		{"1.0", "0a000000000000000000000000000000", 2, 1, 2, 0, "01000000000000000000000000000000", "1", false},
		{"1.0", "0a000000000000000000000000000000", 2, 1, 1, 0, "01000000000000000000000000000000", "1", false},
		{"0.25", "19000000000000000000000000000000", 3, 2, 2, 2, "", "", false},
		{"0.5", "05000000000000000000000000000000", 2, 1, 1, 1, "", "", false},
		{"1.0", "0a000000000000000000000000000000", 2, 1, 40, 0, "", "", true},
		{"4.53", "c5010000000000000000000000000000", 3, 2, 1, 0, "", "", true},
		{"0.25", "19000000000000000000000000000000", 3, 2, 1, 0, "", "", true},
		{"1.0", "0a000000000000000000000000000000", 2, 1, 1, 0, "01000000000000000000000000000000", "1", false},
		{"1.000", "e8030000000000000000000000000000", 4, 3, 1, 0, "01000000000000000000000000000000", "1", false},
		{"1", "01000000000000000000000000000000", 1, 0, 3, 2, "64000000000000000000000000000000", "1.00", false},
		{"9223372036854775807", "ffffffffffffff7f0000000000000000", 19, 0, 30, 4, "f0d8ffffffffffff8713000000000000", "9223372036854775807.0000", false},
		{"-9223372036854775808", "0000000000000080ffffffffffffffff", 19, 0, 30, 2, "0000000000000000ceffffffffffffff", "-9223372036854775808.00", false},
		{"0.0000123", "7b000000000000000000000000000000", 8, 7, 10, 9, "0c300000000000000000000000000000", "0.000012300", false},
		{"1230000000000000", "00e012b1ad5e04000000000000000000", 16, 0, 20, 2, "00805f2bd9fbb4010000000000000000", "1230000000000000.00", false},
		{"123000000000000000000", "00000c6d51c8f7aa0600000000000000", 21, 0, 25, 2, "0000b098ce3fcac89a02000000000000", "123000000000000000000.00", false},
		{"123000000000000000000000000000000000", "00000000cebde644bc05f0425eb01700", 36, 0, 38, 0, "00000000cebde644bc05f0425eb01700", "123000000000000000000000000000000000", false},
		{"123000000000000000000000000000000000", "00000000cebde644bc05f0425eb01700", 36, 0, 38, 2, "00000000782422ea8a3dc225d2e44009", "123000000000000000000000000000000000.00", false},
		{"1234567890123456.78901234", "f2af966ca0101f9b241a000000000000", 24, 8, 28, 10, "88badc6aaa7e22984c360a0000000000", "1234567890123456.7890123400", false},
		{"-1234567890.1234567890", "2ef5e0147356ab54ffffffffffffffff", 20, 10, 24, 12, "f8c5df27f4c4ed12bdffffffffffffff", "-1234567890.123456789000", false},
		{"123456789012345678901234567890.1234", "f2af967ed05c82de3297ff6fde3c0000", 34, 4, 38, 6, "88badc727141eceade0fd7bfe3c61700", "123456789012345678901234567890.123400", false},
		{"-123456789012345678901234567890.1234", "0e5069812fa37d21cd68009021c3ffff", 34, 4, 38, 6, "7845238d8ebe131521f028401c39e8ff", "-123456789012345678901234567890.123400", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {

			inputBytes := [16]byte(hexToBytes(t, tt.hexInput))
			// verify that the conversion is correct
			decStr := decimalBytesToString(inputBytes, tt.inputScale)
			if tt.expected == "" {
				tt.expected = strings.TrimPrefix(tt.input, "+")
			}
			assert.Equal(t, tt.input, decStr)

			newBytes, newPrecision, newScale, err := modifyDecimalPrecisionAndScale(inputBytes, tt.inputScale, tt.targetPrecision, tt.targetScale)
			if tt.expectError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.targetPrecision, newPrecision)
			assert.Equal(t, tt.targetScale, newScale)
			if tt.hexWant2 == "" {
				tt.hexWant2 = tt.hexInput
			}
			assert.Equal(t, hexToBytes(t, tt.hexWant2), newBytes[:])
			decStr = decimalBytesToString(newBytes, newScale)
			assert.Equal(t, tt.expected, decStr)
		})
	}
}
