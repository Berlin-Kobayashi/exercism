// Package say implements a function for spelling out integers in english.
package say

import "strconv"
import "strings"

var digits = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

var teens = []string{
	"ten",
	"eleven",
	"twelve",
	"thirteen",
	"fourteen",
	"fifteen",
	"sixteen",
	"seventeen",
	"eighteen",
	"nineteen",
}

var decadicNumbers = []string{
	"twenty",
	"thirty",
	"forty",
	"fifty",
	"sixty",
	"seventy",
	"eighty",
	"ninety",
}

const hundred = "hundred"

var powersOfThousand = []string{
	"",
	"thousand",
	"million",
	"billion",
	"trillion",
	"quadrillion",
	"quintillion",
}

// Say spells out integers in english.
func Say(number uint64) string {
	resultParts := []string{}

	numberString := strconv.FormatUint(number, 10)
	for i := len(numberString) - 1; i >= 0; i -= 3 {
		from := i - 2
		if from < 0 {
			from = 0
		}
		part := numberString[from : i+1]
		partValue, _ := strconv.Atoi(string(part))
		if i != len(numberString)-1 && partValue != 0 {
			resultParts = prepend(resultParts, powersOfThousand[(len(numberString)-i)/3])
		}
		twoDigitPart := part
		if len(twoDigitPart) > 2 {
			twoDigitPart = part[len(twoDigitPart)-2:]
		}

		twoDigitPartValue, _ := strconv.Atoi(string(twoDigitPart))
		if twoDigitPartValue != 0 || (i == len(numberString)-1 && i-3 < 0 && len(part) != 3) {
			switch {
			case twoDigitPartValue < 10:
				resultParts = prepend(resultParts, digits[int(twoDigitPartValue)])
			case twoDigitPartValue < 20:
				resultParts = prepend(resultParts, teens[int(twoDigitPartValue)%10])
			default:
				remainder := twoDigitPartValue % 10
				factor := twoDigitPartValue/10 - 2
				if remainder == 0 {
					resultParts = prepend(resultParts, decadicNumbers[factor])
				} else {
					resultParts = prepend(resultParts, decadicNumbers[factor]+"-"+digits[remainder])
				}
			}
		}
		if len(part) == 3 {
			hundredValue, _ := strconv.Atoi(string(part[0]))
			if hundredValue != 0 {
				resultParts = prepend(resultParts, digits[hundredValue]+" "+hundred)
			}
		}
	}

	return strings.Join(resultParts, " ")
}

func prepend(slice []string, value string) []string {
	return append([]string{value}, slice...)
}
