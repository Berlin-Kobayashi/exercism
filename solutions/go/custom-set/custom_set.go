// Package stringset implement a set as a collection of unique string values.
package stringset

import (
	"fmt"
	"reflect"
	"strings"
)

const testVersion = 3

// Set represents a set of unique strings.
type Set map[string]bool

// New creates an empty Set.
func New() Set {
	return make(Set)
}

// NewFromSlice creates a Set from the contents of the slice.
// If a string exists multiple times in the slice it will only be once in the Set.
func NewFromSlice(slice []string) Set {
	set := make(Set)
	for _, element := range slice {
		set.Add(element)
	}
	return set
}

// Add adds the element to the Set.
func (s Set) Add(element string) {
	s[element] = true
}

// Delete deletes the element from the Set.
func (s Set) Delete(element string) {
	delete(s, element)
}

// Has checks if the element is in the Set.
func (s Set) Has(element string) bool {
	return s[element]
}

// IsEmpty checks if the Set is empty.
func (s Set) IsEmpty() bool {
	return s.Len() == 0
}

// Len returns the length of the Set.
func (s Set) Len() int {
	return len(s)
}

// Slice converts the Set to a slice.
func (s Set) Slice() []string {
	slice := make([]string, s.Len())
	i := 0
	for element := range s {
		slice[i] = element
		i++
	}

	return slice
}

func (s Set) String() string {
	quotedStrings := s.Slice()
	for i, element := range quotedStrings {
		quotedStrings[i] = fmt.Sprintf("%q", element)
	}

	return fmt.Sprintf("{%s}", strings.Join(quotedStrings, ", "))
}

// Equal checks if the Sets contain the same elements. Order does not matter.
func Equal(s1, s2 Set) bool {
	return reflect.DeepEqual(s1, s2)
}

// Subset checks if all elements of subset are contained by set.
func Subset(subset, set Set) bool {
	for element := range subset {
		if !set[element] {
			return false
		}
	}

	return true
}

// Disjoint checks if no value is contained by both Sets.
func Disjoint(s1, s2 Set) bool {
	for element := range s1 {
		if s2[element] {
			return false
		}
	}

	return true
}

// Intersection returns a Set representing the intersection of the given Sets.
func Intersection(s1, s2 Set) Set {
	intersection := New()
	for element := range s1 {
		if s2[element] {
			intersection.Add(element)
		}
	}

	return intersection
}

// Union returns a Set representing the union of the given Sets.
func Union(s1, s2 Set) Set {
	return NewFromSlice(append(s1.Slice(), s2.Slice()...))
}

// Difference returns a Set representing the difference of s2 from s1.
func Difference(s1, s2 Set) Set {
	difference := New()
	for element := range s1 {
		if !s2[element] {
			difference.Add(element)
		}
	}

	return difference
}

// SymmetricDifference returns a Set representing the difference of the given Sets.
func SymmetricDifference(s1, s2 Set) Set {
	return Union(Difference(s1, s2), Difference(s2, s1))
}
