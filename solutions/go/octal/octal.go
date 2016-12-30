// Package octal implements a function for converting octal to decimal numbers.
package octal

import (
	"errors"
	"math"
	"regexp"
	"strconv"
)

const testVersion = 1

// ParseOctal converts the given octal number to a decimal number.
// If the given string contains characters other than 0-7 an error is returned.
func ParseOctal(input string) (int64, error) {
	octalRegex := regexp.MustCompile("^[0-7]+$")

	if !octalRegex.Match([]byte(input)) {
		return 0, errors.New("Cannot convert octal number with invalid characters. Given: " + input)
	}

	var result int64
	for i, digit := range input {
		value, _ := strconv.Atoi(string(digit))
		result += int64(value) * int64(math.Pow(8, float64(len(input)-i-1)))
	}

	return result, nil
}
