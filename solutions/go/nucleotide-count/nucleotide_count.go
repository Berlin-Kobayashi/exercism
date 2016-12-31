// Package dna implements a function which given a DNA string, computes how many times each nucleotide occurs in the string.
package dna

import (
	"errors"
	"regexp"
)

const testVersion = 2

// DNA represents a DNA strand.
type DNA string

// Histogram represents the count of each nucleotide in a DNA string.
type Histogram map[byte]int

// Counts returns how often each nucleotide occurs in the DNA.
// If the DNA contains nucleotides other than 'A', 'C', 'G', or 'T' an error is returned.
func (dna DNA) Counts() (counts Histogram, err error) {
	validation := regexp.MustCompile("^[ACGT]*$")
	if !validation.MatchString(string(dna)) {
		return nil, errors.New("DNA contains unknown nucleotide: " + string(dna))
	}

	counts = Histogram{}
	for _, nucleotide := range []byte{'A', 'C', 'G', 'T'} {
		count, _ := dna.Count(nucleotide)
		counts[nucleotide] = count
	}

	return counts, nil
}

// Count returns how often the given nucleotide occurs in the DNA.
// If the given nucleotide is not 'A', 'C', 'G', or 'T' an error is returned.
func (dna DNA) Count(nucleotide byte) (count int, err error) {
	switch nucleotide {
	case 'A', 'C', 'G', 'T':
		for _, char := range dna {
			if nucleotide == byte(char) {
				count++
			}
		}
		return count, nil
	default:
		return 0, errors.New("Unknown nucleotide given: " + string(nucleotide))
	}
}
