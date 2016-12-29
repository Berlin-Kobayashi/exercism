// Package cryptosquare implements a function for encoding strings using square code.
package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

const testVersion = 2

// Encode encodes the given string using square code.
func Encode(input string) (output string) {
	input = normalize(input)

	maxX, maxY := calculateSquareDimensions(input)
	for i := 0; i < maxX; i++ {
		for j := 0; j < maxY; j++ {
			letterIndex := i + j*maxX
			if letterIndex < len(input) {
				output += string(input[letterIndex])
			}
		}

		if i != maxX-1 {
			output += " "
		}
	}

	return output
}

func normalize(input string) (output string) {
	input = strings.ToLower(input)

	nonLetterOrDigitRegex := regexp.MustCompile("([^\\d\\pL])")
	output = string(nonLetterOrDigitRegex.ReplaceAll([]byte(input), []byte("")))

	return output
}

func calculateSquareDimensions(input string) (x, y int) {
	inputLength := len(input)
	x = int(math.Sqrt(float64(inputLength)))
	y = x

	if x*y < inputLength {
		x++
	}
	if x*y < inputLength {
		y++
	}

	return x, y
}
