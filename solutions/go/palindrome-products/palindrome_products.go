// Package palindrome implements a simple library for calculating palindromes and their factors in a certain range of factors.
package palindrome

import (
	"errors"
	"strconv"
)

const testVersion = 1

// Product represents a prodct and all of its possible two-factor factorizations in order
type Product struct {
	Product        int
	Factorizations [][2]int
}

// Products calculates the biggest and the smallest palindrome and their factors for the given range of factors inclusively.
func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	if fmin > fmax {
		return pmin, pmax, errors.New("fmin > fmax")
	}

	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			product := i * j
			if isPalindrome(product) {
				if product <= pmin.Product || pmin.Product == 0 {
					if product != pmin.Product {
						pmin.Product = product
						pmin.Factorizations = [][2]int{}
					}
					pmin.Factorizations = append(pmin.Factorizations, [2]int{i, j})
				}

				if product >= pmax.Product || pmax.Product == 0 {
					if product != pmax.Product {
						pmax.Product = product
						pmax.Factorizations = [][2]int{}
					}
					pmax.Factorizations = append(pmax.Factorizations, [2]int{i, j})
				}
			}
		}
	}

	if pmin.Product == 0 && pmax.Product == 0 {
		return pmin, pmax, errors.New("No palindromes")
	}

	return pmin, pmax, nil
}

func isPalindrome(number int) bool {
	numberString := strconv.Itoa(number)
	for i, j := 0, len(numberString)-1; i < j; i, j = i+1, j-1 {
		if numberString[i] != numberString[j] {
			return false
		}
	}

	return true
}
