// Package summultiples implements a function for calculating the sum of the multiples of one or more numbers.
package summultiples

// MultipleSummer returns a function for calculating the sum of the multiples of one or more numbers up to the multiplier limit.
func MultipleSummer(divisors ...int) func(int) int {
	return func(limit int) (result int) {
		for i := 1; i < limit; i++ {
			for _, divisor := range divisors {
				if i%divisor == 0 {
					result += i
					break
				}
			}
		}
		return result
	}
}
