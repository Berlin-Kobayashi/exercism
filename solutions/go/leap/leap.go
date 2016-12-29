/*
Package leap implements a simple library for checking if a year is a leap year.
*/
package leap

const testVersion = 2

// IsLeapYear returns true if the given year is a leap year. Otherwise false is returned.
func IsLeapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		return true
	}

	return false
}
