// Package igpay implements a function for converting an english string to pig latin.
package igpay

import (
	"strings"
	"unicode"
)

// PigLatin converts the given english string to pig latin.
func PigLatin(input string) string {
	input = strings.ToLower(input)

	inputWords := strings.Fields(input)
	outputWords := make([]string, len(inputWords))
	for i, inputWord := range inputWords {
		outputWord := inputWord
		for i := range inputWord {
			if !(isConsonant(inputWord[i]) || outputWord[len(outputWord)-1] == 'q') || (isConditionalVowel(inputWord[i]) && i != len(inputWord)-1 && !isVowel(inputWord[i+1])) {
				break
			}
			outputWord = outputWord[1:] + string(inputWord[i])

		}
		outputWord += "ay"
		outputWords[i] = outputWord
	}

	return strings.Join(outputWords, " ")
}

func isVowel(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}

	return false
}

func isConsonant(b byte) bool {
	return unicode.IsLetter(rune(b)) && !isVowel(b)
}

func isConditionalVowel(r byte) bool {
	switch r {
	case 'x', 'y':
		return true
	}

	return false
}
