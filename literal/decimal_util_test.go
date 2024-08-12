package literal

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/big"
	"strings"
	"testing"
)

func Test_decimalStringToBytes(t *testing.T) {
	tests := []struct {
		input        string
		want         string
		expPrecision int32
		expScale     int32
		wantErr      bool
	}{
		{"12345", "39300000000000000000000000000000", 5, 0, false},
		{"+12345", "39300000000000000000000000000000", 5, 0, false},
		{"-12345", "c7cfffffffffffffffffffffffffffff", 5, 0, false},
		{"123.45", "39300000000000000000000000000000", 5, 2, false},
		{"-123.45", "c7cfffffffffffffffffffffffffffff", 5, 2, false},
		{"0.123", "7b000000000000000000000000000000", 4, 3, false},
		{"-0.123", "85ffffffffffffffffffffffffffffff", 4, 3, false},
		{"9223372036854775807", "ffffffffffffff7f0000000000000000", 19, 0, false},  // Max int64
		{"-9223372036854775808", "0000000000000080ffffffffffffffff", 19, 0, false}, // Min int64
		{"99999999999999999999999999999999999999", "ffffffff3f228a097ac4865aa84c3b4b", 38, 0, false},
		{"+99999999999999999999999999999999999999", "ffffffff3f228a097ac4865aa84c3b4b", 38, 0, false},
		{"-99999999999999999999999999999999999999", "01000000c0dd75f6853b79a557b3c4b4", 38, 0, false},
		{"0", "00000000000000000000000000000000", 1, 0, false},
		{"-0", "00000000000000000000000000000000", 1, 0, false},
		{"0.0", "00000000000000000000000000000000", 2, 1, false},
		{"65535", "ffff0000000000000000000000000000", 5, 0, false},
		{"-65535", "0100ffffffffffffffffffffffffffff", 5, 0, false},
		{"18446744073709551615", "ffffffffffffffff0000000000000000", 20, 0, false},  // Max uint64
		{"-18446744073709551616", "0000000000000000ffffffffffffffff", 20, 0, false}, // Min int64 - 1
		{"12345.6789", "15cd5b07000000000000000000000000", 9, 4, false},
		{"1234567890123456", "c0ba8a3cd56204000000000000000000", 16, 0, false},
		{"1234567890123456.78901234", "f2af966ca0101f9b241a000000000000", 24, 8, false},
		{"0.0012345678901234", "f22fce733a0b00000000000000000000", 17, 16, false},
		{"-0.0012345678901234", "0ed0318cc5f4ffffffffffffffffffff", 17, 16, false},
		{"123456789012345678901234567890.1234", "f2af967ed05c82de3297ff6fde3c0000", 34, 4, false},
		{"-1234567890.1234567890", "2ef5e0147356ab54ffffffffffffffff", 20, 10, false},
		{"0", "00000000000000000000000000000000", 1, 0, false},
		{"12345678901234567890123456789012345678901234", "", 44, 0, true},
		{"abc", "", 0, 0, true},
		{"12.34.56", "", 0, 0, true},
		{"199999999999999999999999999999999999999", "", 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, precision, scale, err := decimalStringToBytes(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("decimalStringToBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			assert.Len(t, got, 16)
			if !strings.EqualFold(tt.want, toHex(got[:])) {
				t.Errorf("decimalStringToBytes() got = %v, want %v", toHex(got[:]), tt.want)
			}
			if precision != tt.expPrecision {
				t.Errorf("decimalStringToBytes() precision = %v, want %v", precision, tt.expPrecision)
			}
			if scale != tt.expScale {
				t.Errorf("decimalStringToBytes() scale = %v, want %v", scale, tt.expScale)
			}
			if err != nil {
				input := tt.input
				if strings.HasPrefix(tt.input, "+") {
					input = tt.input[1:]
				}
				string := decimalBytesToString(got, precision, scale)
				assert.Equal(t, input, string)
			}
		})
	}
}

func toHex(bytes []byte) string {
	return fmt.Sprintf("%02x", bytes)
}

func decimalBytesToString(decimalBytes [16]byte, precision, scale int32) string {
	// Reverse the byte array to big-endian
	for i, j := 0, len(decimalBytes)-1; i < j; i, j = i+1, j-1 {
		decimalBytes[i], decimalBytes[j] = decimalBytes[j], decimalBytes[i]
	}

	// Convert the byte array to a big.Int
	intValue := new(big.Int).SetBytes(decimalBytes[:])

	// Convert the big.Int to a string
	decimalStr := intValue.String()

	// Insert the decimal point at the appropriate position
	if scale > 0 {
		decimalStr = decimalStr[:len(decimalStr)-int(scale)] + "." + decimalStr[len(decimalStr)-int(scale):]
	}

	// Add leading zeros for the integer part
	if precision > int32(len(decimalStr)) {
		decimalStr = strings.Repeat("0", int(precision)-len(decimalStr)) + decimalStr
	}

	// Add a leading zero for the fractional part
	if scale > 0 && len(decimalStr)-strings.Index(decimalStr, ".") == 2 {
		decimalStr = strings.Replace(decimalStr, ".", ".0", 1)
	}

	// Add a leading zero for the integer part
	if scale == 0 && len(decimalStr) == 1 && decimalStr == "0" {
		decimalStr = "0.0"
	}

	return decimalStr
}
