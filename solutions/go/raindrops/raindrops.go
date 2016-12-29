//Package raindrops implements a simple function for converting numbers into raindrop strings.
package raindrops

import (
	"strconv"
)

const testVersion = 2

/*
Convert converts the given number to a raindrop string.
If the number has 3 as a factor, 'Pling' is added to the result string.
If the number has 5 as a factor, 'Plang' is added to the result string.
If the number has 7 as a factor, 'Plong' is added to the result string.
If the number does not have 3, 5, or 7 as a factor, the number is returned.
*/
func Convert(number int) string {
	result := ""
	if number%3 == 0 {
		result += "Pling"
	}
	if number%5 == 0 {
		result += "Plang"
	}
	if number%7 == 0 {
		result += "Plong"
	}

	if result == "" {
		return strconv.FormatInt(int64(number), 10)
	}

	return result
}
