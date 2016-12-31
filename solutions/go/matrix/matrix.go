// Package matrix implements an API for reading and manipulating matrixes of integers.
package matrix

import (
	"errors"
	"strconv"
	"strings"
)

const testVersion = 1

// Matrix represents a matrix of integers.
type Matrix [][]int

// New creates a new matrix of int from the given string.
// The rows of the given string must be seperated by a line break and the columns by a space.
// If the given string is malformed an error is returned.
func New(input string) (Matrix, error) {
	rows := strings.Split(input, "\n")

	var columns int
	matrix := make([][]int, len(rows))
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

	return matrix, nil
}

// Rows returns the rows of the Matrix.
func (m Matrix) Rows() [][]int {
	rows := make([][]int, len(m))
	for i, row := range m {
		rows[i] = make([]int, len(row))
		copy(rows[i], row)
	}

	return rows
}

// Cols returns the columns of the Matrix.
func (m Matrix) Cols() [][]int {
	var cols [][]int
	for i, row := range m {
		if i == 0 {
			cols = make([][]int, len(row))
		}
		for j, value := range row {
			cols[j] = append(cols[j], value)
		}
	}

	return cols
}

// Set sets the value in the matrix.
// If x or y are out of bounds false is returned.
func (m Matrix) Set(y, x, value int) bool {
	if y >= len(m) || y < 0 {
		return false
	}

	if x >= len(m[y]) || x < 0 {
		return false
	}

	m[y][x] = value

	return true
}
