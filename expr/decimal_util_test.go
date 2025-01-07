package expr

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v3/types"
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
			got, precision, scale, err := DecimalStringToBytes(tt.input)
			assert.NoError(t, err)
			assert.Len(t, got, 16)
			assert.Equal(t, hexToBytes(t, tt.hexWant), got[:])
			assert.Equal(t, tt.expPrecision, precision)
			assert.Equal(t, tt.expScale, scale)
			if err == nil {
				// verify that the conversion is correct
				decStr := decimalBytesToString(got, scale)
				if tt.expected == "" {
					tt.expected = strings.TrimPrefix(tt.input, "+")
				}
				assert.Equal(t, tt.expected, decStr)
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

func TestDecimalLiteralToStringBadType(t *testing.T) {
	timeLit := NewPrecisionTimestampTzLiteral(123456, types.PrecisionNanoSeconds, types.NullabilityNullable)
	timeLitAsProtoLit := timeLit.(*ProtoLiteral)
	_, err := DecimalLiteralToString(timeLitAsProtoLit)
	assert.Error(t, err)
}

func TestDecimalLiteralToStringMangledType(t *testing.T) {
	brokenLit := &ProtoLiteral{Value: "random junk", Type: &types.DecimalType{}}
	_, err := DecimalLiteralToString(brokenLit)
	assert.Error(t, err)
}
