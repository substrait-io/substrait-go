package literal

import (
	"fmt"
	"math/big"
	"strings"
)

func decimalStringToBytes(decimalStr string) ([16]byte, int32, int32, error) {
	var result [16]byte

	// Split the string into integer and fractional parts
	parts := strings.Split(decimalStr, ".")
	if len(parts) > 2 {
		return result, 0, 0, fmt.Errorf("invalid decimal string")
	}
	precision := int32(len(parts[0])) // Precision starts with the length of the integer part
	scale := int32(0)

	isNeg := strings.HasPrefix(decimalStr, "-")
	if strings.HasPrefix(decimalStr, "-") || strings.HasPrefix(decimalStr, "+") {
		precision-- // Exclude the sign from the precision
	}
	// If there's a fractional part, adjust precision and scale
	if len(parts) > 1 {
		scale = int32(len(parts[1]))
		precision += scale
		decimalStr = parts[0] + parts[1] // Concatenate parts without the decimal point
	}

	// Parse the concatenated string to a big.Int
	intValue, success := new(big.Int).SetString(decimalStr, 10)
	if !success {
		return result, 0, 0, fmt.Errorf("invalid decimal string")
	}

	// Convert the big.Int to a byte array
	byteArray := intValue.Bytes()

	// Ensure the byte array fits within 16 bytes
	if len(byteArray) > 16 {
		return result, precision, scale, fmt.Errorf("number exceeds 16 bytes")
	}

	// Copy the bytes to the fixed 16-byte array
	copy(result[16-len(byteArray):], byteArray)

	if isNeg {
		// Negate the bytes to get the two's complement
		for i := range result {
			result[i] = ^result[i]
		}
		// Add 1 to the two's complement to get the negative value
		carry := byte(1)
		for i := len(result) - 1; i >= 0; i-- {
			result[i] += carry
			if result[i] == 0 {
				carry = 1
			} else {
				break
			}
		}
	}
	// Reverse the byte array to little-endian
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	if precision > 38 {
		return result, precision, scale, fmt.Errorf("number exceeds maximum precision of 38")
	}
	return result, precision, scale, nil
}
