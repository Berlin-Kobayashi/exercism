// Package lsproduct implements a function for calculating the largest product of a substring of length n in a digit string.
package lsproduct

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

const testVersion = 4

// LargestSeriesProduct calculates the largest product of a substring of length n in the given digit string.
func LargestSeriesProduct(digits string, span int) (product int64, err error) {
	if span < 0 {
		return 0, errors.New("Span must not be negative: " + strconv.Itoa(span))
	}

	numberRegex := regexp.MustCompile("^\\d*$")
	if !numberRegex.Match([]byte(digits)) {
		return 0, errors.New("Digit string must only contain digits: " + digits)
	}

	if span > len(digits) {
		return 0, fmt.Errorf("Span must not be greater than amount of digits: %s (%d)", digits, span)
	}

	if span == 0 {
		return 1, nil
	}

	for i := range digits {
		var subProduct int64 = 1
		for j := 0; j < span; j++ {
			if i+j >= len(digits) {
				subProduct = 0
				break
			}
			value, _ := strconv.Atoi(string(digits[i+j]))
			subProduct *= int64(value)
		}
		if subProduct > product {
			product = subProduct
		}
	}

	return product, nil
}
