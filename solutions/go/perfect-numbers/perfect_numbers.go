// Package perfect implements a function for classifying natural numbers like the Greek mathematician Nicomachus did.
package perfect

import "errors"

const testVersion = 1

// Classification represents the classification of a natural number
type Classification int

const (
	// ClassificationDeficient classifies a number as being deficient (Sum of factors < number).
	ClassificationDeficient Classification = iota

	// ClassificationAbundant classifies a number as being abundant (Sum of factors > number).
	ClassificationAbundant Classification = iota

	// ClassificationPerfect classifies a number as being abundant (Sum of factors = number).
	ClassificationPerfect Classification = iota
)

// ErrOnlyPositive is returned when a number <= 0 is tried to being classified.
var ErrOnlyPositive = errors.New("Only positve numbers can be classified")

// Classify classifies natural numbers like the Greek mathematician Nicomachus did.
func Classify(number uint64) (Classification, error) {
	if number <= 0 {
		return 0, ErrOnlyPositive
	}

	var factorSum uint64
	for i := uint64(1); i <= number/2; i++ {
		if number%i == 0 {
			factorSum += i
		}
	}

	var classification Classification
	switch {
	case factorSum > number:
		classification = ClassificationAbundant
	case factorSum < number:
		classification = ClassificationDeficient
	case factorSum == number:
		classification = ClassificationPerfect
	}

	return classification, nil
}
