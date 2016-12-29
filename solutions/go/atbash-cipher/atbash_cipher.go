// Package atbash implements a function for encoding strings with the atbash cipher.
package atbash

import (
	"bytes"
	"strings"
	"unicode"
)

// Atbash encodes the given string with the atbash cipher.
func Atbash(input string) string {
	input = strings.ToLower(input)

	var output bytes.Buffer
	i := 1
	for j, char := range input {
		if !unicode.IsSpace(char) && !unicode.IsPunct(char) {
			if char >= 'a' && char <= 'z' {
				char = 'z' - char + 'a'
			}
			output.WriteRune(char)

			if j != len(input)-1 && i%5 == 0 {
				output.WriteRune(' ')
			}
			i++
		}
	}

	return strings.TrimRight(output.String(), " ")
}
