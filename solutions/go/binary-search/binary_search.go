// Package binarysearch implements a library for performing binary searches.
package binarysearch

import (
	"fmt"
)

// SearchInts searches a slice of ints sorted in increasing order and returns the index of the first occurence.
// * If there are duplicate values of the key you are searching for, you can't
// just stop at the first one you find; you must find the first occurance in
// the slice.
//
// * There is no special "not found" indication.  If the search key is not
// present, SearchInts returns the index of the first value greater than the
// search key.  If the key is greater than all values in the slice, SearchInts
// returns the length of the slice.
func SearchInts(data []int, search int) int {
	offset := 0
	for len(data) > 0 {
		mid := len(data) / 2
		midValue := data[mid]
		switch {
		case search == midValue:
			for mid > 0 && data[mid-1] == search {
				mid--
			}

			return offset + mid
		case search < midValue:
			if mid == 0 {
				return offset
			}
			data = data[:mid]
		case search > midValue:
			if mid == 0 {
				return offset + 1
			}
			data = data[mid:]
			offset += mid
		}
	}

	return offset
}

// Message searches a slice of ints sorted in increasing order and returns an appropiate result message.
func Message(data []int, search int) string {
	if data == nil {
		return "slice has no values"
	}

	position := SearchInts(data, search)
	if position != len(data) && data[position] == search {
		switch position {
		case 0:
			return fmt.Sprintf("%d found at beginning of slice", search)
		case len(data) - 1:
			return fmt.Sprintf("%d found at end of slice", search)
		default:
			return fmt.Sprintf("%d found at index %d", search, position)
		}
	} else {
		switch position {
		case 0:
			return fmt.Sprintf("%d < all values", search)
		case len(data):
			return fmt.Sprintf("%d > all %d values", search, len(data))
		default:
			return fmt.Sprintf("%d > %d at index %d, < %d at index %d", search, data[position-1], position-1, data[position], position)
		}
	}
}
