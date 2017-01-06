// Package matrix implements a function for finding saddle points in matrices of integers.
// A saddle point is an element which is greater than or equal to every element in its row
// and less than or equal to every element in its column.
package matrix

import (
	"errors"
	"strconv"
	"strings"
)

const testVersion = 1

// Matrix represents a matrix of integers.
type Matrix [][]int

// Pair represents an integer tuple.
type Pair struct {
	x, y int
}

// New creates a new matrix of integers from the given string.
// The rows of the given string must be seperated by a line break and the columns by a space.
// If the given string is malformed an error is returned.
func New(input string) (*Matrix, error) {
	rows := strings.Split(input, "\n")

	var columns int
	matrix := make(Matrix, len(rows))
	for i, row := range rows {
		row = strings.Trim(row, " ")
		numbers := strings.Split(row, " ")
		if i > 0 && columns != len(numbers) {
			return nil, errors.New("Invalid Matrix. Uneven columns per row")
		}
		if i == 0 {
			columns = len(numbers)
		}
		matrix[i] = make([]int, len(numbers))
		for j, number := range numbers {
			parsedNumber, ok := strconv.Atoi(number)
			if ok != nil {
				return nil, errors.New("Invalid Matrix. Matrix contains invalid integer: '" + number + "'")
			}
			matrix[i][j] = parsedNumber
		}
	}

	return &matrix, nil
}

// Saddle returns all saddle points of the matrix.
func (m *Matrix) Saddle() (saddlePoints []Pair) {
	for i, row := range *m {
		maxima := findMaxima(row)
		for _, max := range maxima {
			isSaddlePoint := true
			for _, comparisonRow := range *m {
				if row[max] > comparisonRow[max] {
					isSaddlePoint = false
					break
				}
			}
			if isSaddlePoint {
				saddlePoints = append(saddlePoints, Pair{i, max})
			}
		}
	}

	return saddlePoints
}

func findMaxima(elements []int) (result []int) {
	result = []int{}
	for i, value := range elements {
		switch {
		case len(result) == 0 || value == elements[result[0]]:
			result = append(result, i)
		case value > elements[result[0]]:
			result = []int{i}
		}
	}

	return result
}
