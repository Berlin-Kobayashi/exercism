// Package wordy implements a function for parsing and evaluating simple math word problems returning the answer as an integer.
package wordy

import (
	"regexp"
	"strconv"
	"strings"
)

// Answer parses and evaluates the given math word problem returning the answer as an integer.
// If the input has syntactical errors ok is false.
func Answer(input string) (result int, ok bool) {
	validation := regexp.MustCompile("^What is (\\-?\\d+ .+ \\-?\\d+)\\?$")
	if !validation.Match([]byte(input)) {
		return 0, false
	}

	input = validation.ReplaceAllString(input, "$1")
	words := strings.Fields(input)
	for i := 0; i < len(words); {
		word := words[i]

		if i == 0 {
			number, err := strconv.Atoi(word)
			if err != nil {
				return 0, false
			}
			result += number
			i++
		} else {
			if i+1 == len(words) {
				return 0, false
			}
			switch word {
			case "minus":
				number, err := strconv.Atoi(words[i+1])
				if err != nil {
					return 0, false
				}
				result -= number
				i += 2
			case "plus":
				number, err := strconv.Atoi(words[i+1])
				if err != nil {
					return 0, false
				}
				result += number
				i += 2
			default:
				if words[i+1] == "by" && i+2 != len(words) {
					switch word {
					case "divided":
						number, err := strconv.Atoi(words[i+2])
						if err != nil {
							return 0, false
						}
						result /= number
						i += 3
					case "multiplied":
						number, err := strconv.Atoi(words[i+2])
						if err != nil {
							return 0, false
						}
						result *= number
						i += 3
					default:
						return 0, false
					}
				} else {
					return 0, false
				}
			}
		}
	}

	return result, true
}
