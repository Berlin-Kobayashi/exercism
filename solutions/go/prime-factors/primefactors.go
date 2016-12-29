// Package prime implements a function for calculating the prime factors of a number.
package prime

const testVersion = 2

// Factors calculates the prime factors of a number.
func Factors(number int64) []int64 {
	primeFactors := []int64{}
	for i := int64(2); i <= number; i++ {
		for number%i == 0 {
			primeFactors = append(primeFactors, i)
			number /= i
		}
	}

	return primeFactors
}
