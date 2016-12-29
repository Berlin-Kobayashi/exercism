// Package luhn implements a library for working with the Luhn algorithm.
// http://en.wikipedia.org/wiki/Luhn_algorithm
package luhn

import (
	"regexp"
	"strconv"
)

// AddCheck adds a check number to the given string to make it valid per the Luhn formula.
func AddCheck(input string) string {
	var checkNumber int

	number := removeNonDigits(input)
	if number == "" {
		checkNumber = 0
	} else {
		checksum := calculateChecksum(number, 0)
		checksumOffset := checksum % 10
		if checksumOffset == 0 {
			checkNumber = 0
		} else {
			checkNumber = 10 - checksum%10
		}
	}

	return input + strconv.Itoa(checkNumber)
}

// Valid checks if the given string is valid per the Luhn formula.
func Valid(input string) bool {
	number := removeNonDigits(input)
	if number == "" {
		return false
	}

	checksum := calculateChecksum(number, 1)

	return checksum%10 == 0
}

func removeNonDigits(input string) string {
	nonDigitRegex := regexp.MustCompile("([^\\d]+)")

	return string(nonDigitRegex.ReplaceAll([]byte(input), []byte("")))
}

func calculateChecksum(number string, offset int) int {
	var checksum int
	for i, j := len(number)-1, 0; i >= 0; i, j = i-1, j+1 {
		value, err := strconv.Atoi(string(number[i]))
		if err != nil {
			panic(err)
		}

		if j%2 == offset {
			value *= 2
			if value > 9 {
				value -= 9
			}
		}

		checksum += value
	}

	return checksum
}
