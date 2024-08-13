package literal

import (
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_decimalStringToBytes(t *testing.T) {
	tests := []struct {
		input        string
		hexWant      string
		expPrecision int32
		expScale     int32
	}{
		{"12345", "39300000000000000000000000000000", 5, 0},
		{"+12345", "39300000000000000000000000000000", 5, 0},
		{"-12345", "c7cfffffffffffffffffffffffffffff", 5, 0},
		{"123.45", "39300000000000000000000000000000", 5, 2},
		{"-123.45", "c7cfffffffffffffffffffffffffffff", 5, 2},
		{"0.123", "7b000000000000000000000000000000", 4, 3},
		{"-0.123", "85ffffffffffffffffffffffffffffff", 4, 3},
		{"9223372036854775807", "ffffffffffffff7f0000000000000000", 19, 0},  // Max int64
		{"-9223372036854775808", "0000000000000080ffffffffffffffff", 19, 0}, // Min int64
		{"99999999999999999999999999999999999999", "ffffffff3f228a097ac4865aa84c3b4b", 38, 0},
		{"+99999999999999999999999999999999999999", "ffffffff3f228a097ac4865aa84c3b4b", 38, 0},
		{"-99999999999999999999999999999999999999", "01000000c0dd75f6853b79a557b3c4b4", 38, 0},
		{"0", "00000000000000000000000000000000", 1, 0},
		{"-0", "00000000000000000000000000000000", 1, 0},
		{"0.0", "00000000000000000000000000000000", 2, 1},
		{"65535", "ffff0000000000000000000000000000", 5, 0},
		{"-65535", "0100ffffffffffffffffffffffffffff", 5, 0},
		{"18446744073709551615", "ffffffffffffffff0000000000000000", 20, 0},  // Max uint64
		{"-18446744073709551616", "0000000000000000ffffffffffffffff", 20, 0}, // Min int64 - 1
		{"12345.6789", "15cd5b07000000000000000000000000", 9, 4},
		{"1234567890123456", "c0ba8a3cd56204000000000000000000", 16, 0},
		{"1234567890123456.78901234", "f2af966ca0101f9b241a000000000000", 24, 8},
		{"0.0012345678901234", "f22fce733a0b00000000000000000000", 17, 16},
		{"-0.0012345678901234", "0ed0318cc5f4ffffffffffffffffffff", 17, 16},
		{"123456789012345678901234567890.1234", "f2af967ed05c82de3297ff6fde3c0000", 34, 4},
		{"-1234567890.1234567890", "2ef5e0147356ab54ffffffffffffffff", 20, 10},
		{"0", "00000000000000000000000000000000", 1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			testDecimalStringToBytes(t, tt.input, tt.hexWant, tt.expPrecision, tt.expScale)
		})
	}

	badInputs := []string{
		"12345678901234567890123456789012345678901234",
		"abc",
		"12.34.56",
		"199999999999999999999999999999999999999",
		"1.23e20",
	}
	for _, input := range badInputs {
		t.Run(input, func(t *testing.T) {
			_, _, _, err := decimalStringToBytes(input)
			assert.Error(t, err, "decimalStringToBytes(%v) expected error", input)
		})
	}
}

func testDecimalStringToBytes(t *testing.T, input, hexWant string, expPrecision, expScale int32) {
	got, precision, scale, err := decimalStringToBytes(input)
	assert.NoError(t, err)
	assert.Len(t, got, 16)
	assert.Equal(t, hexToBytes(hexWant), got[:])
	assert.Equal(t, expPrecision, precision)
	assert.Equal(t, expScale, scale)
	if err != nil {
		// verify that the conversion is correct
		string := decimalBytesToString(got, precision, scale)
		strings.TrimPrefix(input, "+")
		assert.Equal(t, input, string)
	}
}

func hexToBytes(input string) []byte {
	bytes := make([]byte, len(input)/2)
	for i := 0; i < len(input); i += 2 {
		fmt.Sscanf(input[i:i+2], "%02x", &bytes[i/2])
	}
	return bytes
}

func decimalBytesToString(decimalBytes [16]byte, precision, scale int32) string {
	// Reverse the byte array to big-endian
	for i, j := 0, len(decimalBytes)-1; i < j; i, j = i+1, j-1 {
		decimalBytes[i], decimalBytes[j] = decimalBytes[j], decimalBytes[i]
	}

	// Convert the byte array to a big.Int
	intValue := new(big.Int).SetBytes(decimalBytes[:])

	return big.NewFloat(0).SetInt(intValue).SetPrec(uint(precision)).String()
}
