// Package slice implements a library for getting all substrings of the length n from a string.
package slice

// All returns all substrings of the length n.
func All(n int, input string) (subStrings []string) {
	for i := range input {
		if len(input)-i >= n {
			subStrings = append(subStrings, input[i:i+n])
		}
	}

	return subStrings
}

// UnsafeFirst returns the first substring of the length n. The behaviour of this function is not defined when asking for a substring longer than the input string.
func UnsafeFirst(n int, input string) string {
	return input[:n]
}
