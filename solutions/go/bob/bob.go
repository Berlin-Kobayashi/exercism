//Package bob implements a functions which responds just like a lackadaisical teenager would do.
package bob

import (
	"regexp"
)

const testVersion = 2

const (
	questionResponse string = "Sure."
	yellResponse     string = "Whoa, chill out!"
	nothingResponse  string = "Fine. Be that way!"
	otherResponse    string = "Whatever."
)

//Hey responds just like a lackadaisical teenager would do.
func Hey(input string) string {
	questionRegex := regexp.MustCompile("\\?[[:space:]\\p{Zs}]*$")
	yellRegex := regexp.MustCompile("(^\\P{Ll}*\\p{Lu}+\\P{Ll}*$)")
	nothingRegex := regexp.MustCompile("^[[:space:]\\p{Zs}]*$")

	inputArray := []byte(input)
	switch {
	case yellRegex.Match(inputArray):
		return yellResponse
	case questionRegex.Match(inputArray):
		return questionResponse
	case nothingRegex.Match(inputArray):
		return nothingResponse
	default:
		return otherResponse
	}
}
