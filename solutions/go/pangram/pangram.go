//Package pangram implements a function which checks if a string is a pangram.
package pangram

import (
	"strings"
)

const testVersion = 1

//IsPangram returns true if the given string is a pangram. Otherwise false is returned.
func IsPangram(input string) bool {
	letterSet := make(map[rune]bool)

	input = strings.ToLower(input)
	for _, char := range input {
		if char >= 'a' && char <= 'z' {
			letterSet[char] = true
		}
	}

	return len(letterSet) == 26
}
