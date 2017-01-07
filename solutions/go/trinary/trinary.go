// Package trinary implements a function for converting a trinary number to a decimal number.
package trinary

import (
	"fmt"
	"math"
	"strconv"
)

const testVersion = 1

// ParseTrinary converts a trinary number to a decimal number.
// If the input contains invalid characters or overflows int64 an error is returned.
func ParseTrinary(input string) (result int64, err error) {
	for i, digit := range input {
		switch digit {
		case '0', '1', '2':
			digitValue, _ := strconv.Atoi(string(digit))
			value := int64(digitValue) * pow(3, len(input)-1-i)
			if value > math.MaxInt64-result {
				return 0, fmt.Errorf("Cannot parse trinary. Input overflows int64")
			}
			result += value
		default:
			return 0, fmt.Errorf("Cannot parse trinary. Input contains invalid character %q", digit)
		}
	}

	return result, nil
}

func pow(base, exponent int) int64 {
	result := int64(1)
	for i := 0; i < exponent; i++ {
		result *= int64(base)
	}

	return result
}
