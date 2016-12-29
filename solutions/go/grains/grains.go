// Package grains calculates the ammount of grains on each field of a chessboard if their ammount is doubled on each square.
package grains

import (
	"errors"
	"strconv"
)

const testVersion = 1

// Total calculates the total ammount of grains on a chessboard.
func Total() (totalGrains uint64) {
	for i := 1; i <= 64; i++ {
		squareGrains, _ := Square(i)
		totalGrains += squareGrains
	}

	return totalGrains
}

// Square calculates the ammount of grains on a specific square.
func Square(index int) (uint64, error) {
	if index <= 0 || index > 64 {
		return 0, errors.New("Square index must be between 1 and 64. Given: " + strconv.Itoa(index))
	}

	return 1 << (uint(index) - 1), nil
}
