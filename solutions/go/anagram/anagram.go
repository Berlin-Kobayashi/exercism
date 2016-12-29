// Package anagram implements a function for finding correct anagrams.
package anagram

import (
	"strings"
)

// Detect finds the correct anagrams for the given word.
func Detect(word string, candidates []string) []string {
	word = strings.ToLower(word)

	anagrams := []string{}
	for _, candidate := range candidates {
		candidate = strings.ToLower(candidate)
		if isAnagram(word, candidate) {
			anagrams = append(anagrams, candidate)
		}
	}

	return anagrams
}

func isAnagram(word, candidate string) bool {
	if word == candidate {
		return false
	}

	runeCountMap := createRuneCountMap(word)
	for _, runeValue := range candidate {
		runeCount := runeCountMap[runeValue]

		if runeCount < 1 {
			return false
		}

		if runeCount == 1 {
			delete(runeCountMap, runeValue)
		} else {
			runeCountMap[runeValue]--
		}
	}

	return len(runeCountMap) == 0
}

func createRuneCountMap(input string) map[rune]int {
	runeCountMap := make(map[rune]int)
	for _, runeValue := range input {
		runeCountMap[runeValue]++
	}

	return runeCountMap
}
