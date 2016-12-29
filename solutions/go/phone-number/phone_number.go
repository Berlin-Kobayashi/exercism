// Package phonenumber implements a simple library for working with phone numbers.
package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

const testVersion = 1

// Number normalizes the given phone number. If the format of the given phone number is not valid an error will be returned.
func Number(input string) (string, error) {
	nonDigitRegex := regexp.MustCompile("(\\D+)")
	input = string(nonDigitRegex.ReplaceAll([]byte(input), []byte("")))

	if (len(input) == 11 && input[0] != '1') || len(input) > 11 || len(input) < 10 {
		return "", errors.New("Phone number must be between 10 an 11 digits long. Digits given : " + strconv.Itoa(len(input)))
	}

	if len(input) == 11 {
		input = input[1:11]
	}

	return input, nil
}

// AreaCode extracts the area code from the given phone number. If the format of the given phone number is not valid an error will be returned.
func AreaCode(input string) (string, error) {
	number, err := Number(input)
	if err != nil {
		return "", err
	}

	return number[0:3], nil
}

// Format prettifies the given phone number. If the format of the given phone number is not valid an error will be returned.
func Format(input string) (string, error) {
	number, err := Number(input)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("(%s) %s-%s", number[0:3], number[3:6], number[6:]), nil
}
