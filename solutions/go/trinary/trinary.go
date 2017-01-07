// Package trinary implements a function for converting a trinary number to a decimal number.
package trinary

import "fmt"

const testVersion = 1

// ParseTrinary converts a trinary number to a decimal number.
// If the input contains invalid characters or overflows int64 an error is returned.
func ParseTrinary(input string) (result int64, err error) {
	for _, digit := range input {
		if digit < '0' || digit > '2' {
			return 0, fmt.Errorf("Cannot parse trinary. Input contains invalid character %q", digit)
		}

		digitValue := digit - '0'
		result = result*3 + int64(digitValue)

		if result < 0 {
			return 0, fmt.Errorf("Cannot parse trinary. Input overflows int64")
		}
	}

	return result, nil
}
