// Package sieve implements a function for calucating primes using the Sieve of Eratosthenes.
package sieve

// Sieve calculates primes from 2 until the limt inclusively using the Sieve of Eratosthenes.
func Sieve(limit int) []int {
	nonPrimes := make([]bool, limit+1, limit+1)

	for i := 2; i <= limit/2; i++ {
		for j := i + 1; j <= limit; j++ {
			if j%i == 0 {
				nonPrimes[j] = true
			}
		}
	}

	primes := make([]int, 0, limit+1)
	for number, isNonPrime := range nonPrimes {
		if !isNonPrime && number >= 2 {
			primes = append(primes, number)
		}
	}

	return primes
}
