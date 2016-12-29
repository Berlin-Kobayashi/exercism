//Package triangle implements a function to determine the kind of triangle based on its side lengths.
package triangle

import (
	"math"
	"sort"
)

const testVersion = 3

const (
	//NaT not a triangle
	NaT = iota
	//Equ equilateral
	Equ = iota
	//Iso isosceles
	Iso = iota
	//Sca scalene
	Sca = iota
)

//Kind specifies the triangle kind.
type Kind int

//KindFromSides determines the kind of triangle based on the given side lengths.
func KindFromSides(a, b, c float64) Kind {
	lengths := []float64{a, b, c}
	sort.Float64s(lengths)

	for _, length := range lengths {
		if length <= 0 || math.IsNaN(length) || math.IsInf(length, 0) {
			return NaT
		}
	}

	if lengths[2] > lengths[0]+lengths[1] {
		return NaT
	}

	if lengths[0] == lengths[1] && lengths[1] == lengths[2] {
		return Equ
	}

	if lengths[0] == lengths[1] || lengths[1] == lengths[2] {
		return Iso
	}

	return Sca
}
