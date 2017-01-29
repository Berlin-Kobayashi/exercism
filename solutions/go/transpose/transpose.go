// Package transpose implements a function for transposing arrays.
// See <https://en.wikipedia.org/wiki/Transpose>
package transpose

// Transpose Transposes the given array.
func Transpose(input []string) []string {
	maxLength := getMaxLength(input)

	output := make([]string, maxLength, maxLength)
	for i := range output {
		paddingSize := 0
		for _, element := range input {
			if i >= len(element) {
				output[i] += " "
				paddingSize++
			} else {
				output[i] += string(element[i])
				paddingSize = 0
			}
		}
		output[i] = output[i][:len(output[i])-paddingSize]
	}

	return output
}

func getMaxLength(input []string) (maxLength int) {
	for i, element := range input {
		if len(element) > maxLength {
			maxLength = i
		}
	}

	return maxLength
}
