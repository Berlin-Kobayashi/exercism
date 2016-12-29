// Package binary implements a function for converting binary to decimal numbers.
package binary

import (
	"errors"
	"regexp"
)

const testVersion = 2

// ParseBinary converts a binary to a decimal number.
func ParseBinary(input string) (int, error) {
	binaryRegex := regexp.MustCompile("^[01]+$")

	if !binaryRegex.Match([]byte(input)) {
		return 0, errors.New("Binary number must only contain characters 0 and 1. Given: " + input)
	}

	result := 0
	for i, j := len(input)-1, 0; i >= 0; i, j = i-1, j+1 {
		if input[i] == '1' {
			result += 1 << uint(j)
		}
	}

	return result, nil
}
