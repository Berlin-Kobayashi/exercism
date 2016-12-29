// Package strand implements a function for converting DNA strands into RNA strands.
package strand

import (
	"bytes"
)

const testVersion = 3

var dnaNucleotideToRNANucleotide = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA converts the given DNA strand into an RNA strand.
func ToRNA(dna string) string {
	var buffer bytes.Buffer

	for _, nucleotide := range dna {
		buffer.WriteRune(dnaNucleotideToRNANucleotide[nucleotide])
	}

	return buffer.String()
}
