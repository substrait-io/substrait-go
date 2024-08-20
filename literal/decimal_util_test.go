package literal

import (
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/cockroachdb/apd/v3"
	"github.com/stretchr/testify/assert"
)

func Test_decimalStringToBytes(t *testing.T) {
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
		{"0.123", "7b000000000000000000000000000000", 4, 3, ""},
		{"-0.123", "85ffffffffffffffffffffffffffffff", 4, 3, ""},
		{"9223372036854775807", "ffffffffffffff7f0000000000000000", 19, 0, ""},  // Max int64
		{"-9223372036854775808", "0000000000000080ffffffffffffffff", 19, 0, ""}, // Min int64
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
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			testDecimalStringToBytes(t, tt.input, tt.hexWant, tt.expPrecision, tt.expScale, tt.expected)
		})
	}

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
			_, _, _, err := decimalStringToBytes(tt.input)
			assert.Error(t, err, "decimalStringToBytes(%v) expected error", tt.input)
		})
	}
}

func testDecimalStringToBytes(t *testing.T, input, hexWant string, expPrecision, expScale int32, expected string) {
	got, precision, scale, err := decimalStringToBytes(input)
	assert.NoError(t, err)
	assert.Len(t, got, 16)
	assert.Equal(t, hexToBytes(t, hexWant), got[:])
	assert.Equal(t, expPrecision, precision)
	assert.Equal(t, expScale, scale)
	if err == nil {
		// verify that the conversion is correct
		decStr := decimalBytesToString(got, scale)
		if expected == "" {
			expected = strings.TrimPrefix(input, "+")
		}
		assert.Equal(t, expected, decStr)
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

func decimalBytesToString(decimalBytes [16]byte, scale int32) string {
	// Reverse the byte array to big-endian
	for i, j := 0, len(decimalBytes)-1; i < j; i, j = i+1, j-1 {
		decimalBytes[i], decimalBytes[j] = decimalBytes[j], decimalBytes[i]
	}

	isNegative := decimalBytes[0]&0x80 != 0
	// compute two's complement for negative numbers
	if isNegative {
		twosComplement(decimalBytes[:])
	}

	// Convert the byte array to a big.Int
	intValue := new(big.Int).SetBytes(decimalBytes[:])
	if isNegative {
		intValue.Neg(intValue)
	}
	apdBigInt := apd.NewBigInt(0).SetMathBigInt(intValue)
	return apd.NewWithBigInt(apdBigInt, -scale).String()
}
