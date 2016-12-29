// Package etl implements a function for transforming the old Scrabble score table format to the new one.
package etl

import (
	"strings"
)

// Transform transforms the old Scrabble score table format to the new one.
func Transform(input map[int][]string) (output map[string]int) {
	output = make(map[string]int)
	for score, letters := range input {
		for _, letter := range letters {
			output[strings.ToLower(letter)] = score
		}
	}

	return output
}
