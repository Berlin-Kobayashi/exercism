// Package romannumerals implements a function for converting integers into roman numerals.
package romannumerals

import (
	"bytes"
	"errors"
	"strconv"
)

const testVersion = 3

type romanNumeral struct {
	Notation rune
	Value    int
}

var romanNumerals = []romanNumeral{
	{'I', 1},
	{'V', 5},
	{'X', 10},
	{'L', 50},
	{'C', 100},
	{'D', 500},
	{'M', 1000},
}

// ToRomanNumeral converts the given number into a roman numeral. The number must be between 1 and 4000 otherwise an error is returned.
func ToRomanNumeral(number int) (string, error) {
	if number <= 0 || number >= 4000 {
		return "", errors.New("The romans only knew numbers between 1 and 4000. Given: " + strconv.Itoa(number))
	}

	var romanNumeralBuffer bytes.Buffer
	for i := len(romanNumerals) - 1; i >= 0; i-- {
		romanNum := romanNumerals[i]
		for number >= romanNum.Value {
			romanNumeralBuffer.WriteRune(romanNum.Notation)
			number -= romanNum.Value
		}
		subtrahendOffset := (len(romanNumerals)-i)%2 + 1
		if i >= subtrahendOffset {
			subtrahendRomanNum := romanNumerals[i-subtrahendOffset]

			difference := romanNum.Value - subtrahendRomanNum.Value
			if number >= difference {
				romanNumeralBuffer.WriteRune(subtrahendRomanNum.Notation)
				romanNumeralBuffer.WriteRune(romanNum.Notation)
				number -= difference
			}
		}
	}

	return romanNumeralBuffer.String(), nil
}
