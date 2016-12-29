// Package allergies implements a library for determining which allergies an allergy score means.
package allergies

const testVersion = 1

var allergenToBit = map[string](uint){
	"eggs":         0,
	"peanuts":      1,
	"shellfish":    2,
	"strawberries": 3,
	"tomatoes":     4,
	"chocolate":    5,
	"pollen":       6,
	"cats":         7,
}

// Allergies determines which allergies the given allergy score means.
func Allergies(score uint) []string {
	allergies := []string{}

	for allergen := range allergenToBit {
		if AllergicTo(score, allergen) {
			allergies = append(allergies, allergen)
		}
	}

	return allergies
}

// AllergicTo checks if the given allergy is included in the given score.
func AllergicTo(score uint, allergen string) bool {
	return score&(1<<allergenToBit[allergen]) != 0
}
