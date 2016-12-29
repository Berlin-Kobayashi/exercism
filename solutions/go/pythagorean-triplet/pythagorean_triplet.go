// Package pythagorean implements a library for working with Pythagorean triplets.
package pythagorean

import (
	"math"
)

// Triplet represents a triplet of three integers.
type Triplet [3]int

// GetSum returns the sum of the triplet.
func (triplet Triplet) GetSum() int {
	return triplet[0] + triplet[1] + triplet[2]
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to perimeter.
func Sum(perimeter int) []Triplet {
	results := []Triplet{}
	all := Range(1, perimeter)
	for _, triplet := range all {
		if triplet.GetSum() == perimeter {
			results = append(results, triplet)
		}
	}

	return results
}

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) []Triplet {
	triplets := []Triplet{}
	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			c := math.Sqrt(float64(a*a) + float64(b*b))
			if c == math.Floor(c) && c <= float64(max) {
				triplets = append(triplets, Triplet{a, b, int(c)})
			}
		}
	}

	return triplets
}
