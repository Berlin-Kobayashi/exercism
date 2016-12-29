// Package wordcount implements a library for determining the count of words in a string.
package wordcount

import (
	"regexp"
	"strings"
)

const testVersion = 2

// Frequency represents the frequency of words.
type Frequency map[string]int

// WordCount determines the count of words in a string.
func WordCount(phrase string) Frequency {
	frequency := make(Frequency)

	phrase = normalize(phrase)

	words := strings.Split(phrase, " ")
	for _, word := range words {
		if word != "" {
			frequency[word]++
		}
	}

	return frequency
}

func normalize(phrase string) string {
	nonWordCharacterOrSpaceRegex := regexp.MustCompile("([^\\w ]+)")
	phrase = string(nonWordCharacterOrSpaceRegex.ReplaceAll([]byte(phrase), []byte("")))
	phrase = strings.ToLower(phrase)

	return phrase
}
