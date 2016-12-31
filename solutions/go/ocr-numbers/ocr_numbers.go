// Package ocr implements a function which can, given a 3 x 4 grid of pipes, underscores, and spaces,
// determine which numbers are represented, or whether they are garbled.
package ocr

import (
	"errors"
	"strings"
)

// Recognize returns the numbers in the given 3 x 4 grids of pipes, underscores, and spaces.
// Non recognized digits are returned as '?'.
func Recognize(input string) (numbers []string) {
	partitionedInput, err := partitionInput(input)
	if err != nil {
		return []string{"?"}
	}

	result := make([]string, len(partitionedInput))
	for i, number := range partitionedInput {
		for _, digit := range number {
			result[i] += recognizeDigit(digit)
		}
	}

	return result
}

func partitionInput(input string) (numbers [][]string, err error) {
	rows := strings.Split(input, "\n")
	numbers = [][]string{}

	for i, row := range rows {
		if len(row)%3 != 0 {
			return [][]string{}, errors.New("Cannot partition input. Non 3 characters wide digit detected")
		}

		if i == 0 || (len(rows) > 5 && i >= 4 && (i-4)%4 == 0 && i != len(rows)-1) {
			numbers = append(numbers, make([]string, len(rows[1])/3))
			for i := range numbers[0] {
				numbers[len(numbers)-1][i] = "\n"
			}
		} else {
			for j, char := range row {
				numbers[len(numbers)-1][j/3] += string(char)
				if (j+1)%3 == 0 {
					numbers[len(numbers)-1][j/3] += "\n"
				}
			}
		}
	}

	return numbers, nil
}

func recognizeDigit(input string) string {
	rows := strings.Split(input, "\n")
	switch {
	case rows[1] == " _ " && rows[2] == "| |" && rows[3] == "|_|":
		return "0"
	case rows[1] == "   " && rows[2] == "  |" && rows[3] == "  |":
		return "1"
	case rows[1] == " _ " && rows[2] == " _|" && rows[3] == "|_ ":
		return "2"
	case rows[1] == " _ " && rows[2] == " _|" && rows[3] == " _|":
		return "3"
	case rows[1] == "   " && rows[2] == "|_|" && rows[3] == "  |":
		return "4"
	case rows[1] == " _ " && rows[2] == "|_ " && rows[3] == " _|":
		return "5"
	case rows[1] == " _ " && rows[2] == "|_ " && rows[3] == "|_|":
		return "6"
	case rows[1] == " _ " && rows[2] == "  |" && rows[3] == "  |":
		return "7"
	case rows[1] == " _ " && rows[2] == "|_|" && rows[3] == "|_|":
		return "8"
	case rows[1] == " _ " && rows[2] == "|_|" && rows[3] == " _|":
		return "9"
	default:
		return "?"
	}
}
