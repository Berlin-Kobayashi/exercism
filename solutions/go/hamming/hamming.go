//Package hamming implements a function which calculates the hamming distance between two strings.
package hamming

import (
	"fmt"
)

const testVersion = 5

//UnequalStringLengthsError means that two strings don't have the same length.
type UnequalStringLengthsError struct {
	A, B string
}

func (e UnequalStringLengthsError) Error() string {
	return fmt.Sprintf("String lengths are unequal. %s(%d characters) %s(%d characters)", e.A, len(e.A), e.B, len(e.B))
}

// Distance calculates the hamming distance between two strings. If the strings do not have the same length -1 is returned.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, UnequalStringLengthsError{a, b}
	}

	distance := 0
	for i := range a {
		charA := a[i]
		charB := b[i]
		if charA != charB {
			distance++
		}
	}

	return distance, nil
}
