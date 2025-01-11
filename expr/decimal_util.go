package expr

import (
	"fmt"
	"math/big"
	"regexp"
	"strings"

	"github.com/cockroachdb/apd/v3"
)

var decimalPattern = regexp.MustCompile(`^[+-]?\d*(\.\d*)?([eE][+-]?\d*)?$`)

// DecimalStringToBytes converts a decimal string to a 16-byte byte array.
// 16-byte bytes represents a little-endian 128-bit integer, to be divided by 10^Scale to get the decimal value.
// This function also returns the precision and scale of the decimal value.
// The precision is the total number of digits in the decimal value. The precision is limited to 38 digits.
// The scale is the number of digits to the right of the decimal point. The scale is limited to the precision.
func DecimalStringToBytes(decimalStr string) ([16]byte, int32, int32, error) {
	var (
		result    [16]byte
		precision int32
		scale     int32
	)

	strings.Trim(decimalStr, " ")
	if !decimalPattern.MatchString(decimalStr) {
		return result, 0, 0, fmt.Errorf("invalid decimal string")
	}

	// Parse the decimal string using apd
	dec, cond, err := apd.NewFromString(decimalStr)
	if err != nil || cond.Any() {
		return result, 0, 0, fmt.Errorf("invalid decimal string %s: %v", decimalStr, err)
	}

	if dec.Exponent > 0 {
		precision = int32(apd.NumDigits(&dec.Coeff)) + dec.Exponent
		scale = 0
	} else {
		scale = -dec.Exponent
		precision = max(int32(apd.NumDigits(&dec.Coeff)), scale+1)
	}
	if precision > 38 {
		return result, precision, scale, fmt.Errorf("number %s exceeds maximum precision of 38 (%d)", decimalStr, precision)
	}

	coefficient := dec.Coeff
	if dec.Exponent > 0 {
		// Multiply coefficient by 10^exponent.
		multiplier := apd.NewBigInt(1).Exp(apd.NewBigInt(10), apd.NewBigInt(int64(dec.Exponent)), nil)
		coefficient.Mul(&dec.Coeff, multiplier)
	}
	// Convert the coefficient to a byte array.
	byteArray := coefficient.Bytes()
	if len(byteArray) > 16 {
		return result, 0, 0, fmt.Errorf("number exceeds 16 bytes")
	}
	copy(result[16-len(byteArray):], byteArray)

	// Handle the sign by taking the two's complement for negative numbers.
	if dec.Negative {
		negate(result[:])
	}

	// Reverse the byte array to little-endian.
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result, precision, scale, nil
}

// negate flips the sign of a two-complements value by modifying it in place.
func negate(bytes []byte) {
	for i := range bytes {
		bytes[i] = ^bytes[i]
	}
	carry := byte(1)
	for i := len(bytes) - 1; i >= 0; i-- {
		bytes[i] += carry
		if bytes[i] != 0 {
			break
		}
	}
}

func decimalBytesToString(decimalBytes [16]byte, scale int32) string {
	isNegative := decimalBytes[15]&0x80 != 0

	// Reverse the byte array to big-endian.
	processingValue := make([]byte, 16)
	for i := len(processingValue) - 1; i >= 0; i = i - 1 {
		processingValue[i] = decimalBytes[15-i]
	}
	if isNegative {
		negate(processingValue[:])
	}

	// Convert into an apd.BigInt so it can handle the rendering.
	intValue := new(big.Int).SetBytes(processingValue[:])
	if isNegative {
		intValue.Neg(intValue)
	}
	apdBigInt := new(apd.BigInt).SetMathBigInt(intValue)
	return apd.NewWithBigInt(apdBigInt, -scale).String()
}
