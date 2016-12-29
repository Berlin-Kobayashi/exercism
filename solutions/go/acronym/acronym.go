//Package acronym implements a function which abbreviates strings to acronyms.
package acronym

import (
	"strings"
)

const testVersion = 2

//Abbreviate abbreviates the given string to its acronym.
func Abbreviate(input string) string {
	resultChars := make([]string, 0)
	var prevChar rune
	for _, char := range input {
		if isAcronymCharacter(char, prevChar) {
			resultChars = append(resultChars, strings.ToUpper(string(char)))
		}
		prevChar = char
	}

	return strings.Join(resultChars, "")
}

func isAcronymCharacter(char, prevChar rune) bool {
	return (isUppercase(char) && (prevChar == 0 || !isUppercase(prevChar))) || (prevChar != 0 && prevChar == ' ' || prevChar == '-')
}

func isUppercase(char rune) bool {
	return (string(char) == strings.ToUpper(string(char)) && string(char) != strings.ToLower(string(char)))
}
