// Package protein implements a library for converting RNA to lists of proteins.
package protein

const testVersion = 1

var codons = map[string](string){
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

// FromRNA returns the proteins corresponding to the given RNA.
func FromRNA(rna string) []string {
	proteins := []string{}
	for i := 0; i < len(rna); i += 3 {
		if protein, ok := codons[rna[i:i+3]]; ok {
			if protein == "STOP" {
				break
			}

			proteins = append(proteins, protein)
		}
	}

	return proteins
}

// FromCodon returns the protein corresponding to the given codon.
func FromCodon(codon string) string {
	return codons[codon]
}
