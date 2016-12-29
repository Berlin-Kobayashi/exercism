// Package diffsquares implements a simple library for determing the difference between the square of the sums and the sum of the squares of the first n numbers.
package diffsquares

// SquareOfSums returns the square of the sums of the first n numbers.
func SquareOfSums(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result += i
	}

	return result * result
}

// SumOfSquares returns the sum of the squares of the first n numbers.
func SumOfSquares(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result += i * i
	}

	return result
}

// Difference returns the square of the sums minus the sum of the squares of the first n numbers.
func Difference(n int) int {
	result := SquareOfSums(n) - SumOfSquares(n)

	return result
}
