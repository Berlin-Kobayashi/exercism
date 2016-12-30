// Package isogram implements a function for checking if a string is an isogram.
// An isogram (also known as a "nonpattern word") is a word or phrase without a repeating letter.
package isogram

import (
	"strings"
	"unicode"
)

const testVersion = 1

// IsIsogram checks if a string is an isogram.
func IsIsogram(input string) bool {
	input = strings.ToLower(input)

	letters := make(map[rune](bool))
	for _, letter := range input {
		if unicode.IsLetter(letter) {
			if letters[letter] {
				return false
			}
			letters[letter] = true
		}
	}

	return true
}
