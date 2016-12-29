// Package prime implements a function for calculating the nth prime.
package prime

// Nth calculates the nth prime. If n is < 1 ok is false.
func Nth(n int) (result int, ok bool) {
	if n < 1 {
		return 0, false
	}

	primes := []int{}
	for number, primeCount := 2, 0; primeCount < n; number++ {
		isPrime := true
		for _, prime := range primes {
			if number%prime == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, number)
			primeCount++
		}
	}

	return primes[len(primes)-1], true
}
