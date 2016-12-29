// Package isogram implements a function for checking if a string is an isogram.
package isogram

import (
	"regexp"
	"strings"
)

const testVersion = 1

// IsIsogram checks if a string is an isogram.
func IsIsogram(input string) bool {
	input = strings.ToLower(input)

	letters := make(map[rune](bool))
	letterRegex := regexp.MustCompile("^\\p{L}$")
	for _, letter := range input {
		if letterRegex.Match([]byte(string(letter))) {
			if letters[letter] {
				return false
			}
			letters[letter] = true
		}
	}

	return true
}
