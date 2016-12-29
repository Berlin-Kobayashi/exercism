// Package diamond implements a function to create diamonds.
package diamond

import (
	"bytes"
	"fmt"
	"math"
	"strings"
)

const testVersion = 1

const (
	minCarat         = 'A'
	maxCarat         = 'Z'
	spacingCharacter = ' '
)

// Gen creates a diamond with the given carat weight. Carat must be between 'A' and 'Z' inclusively otherwise an error is thrown.
func Gen(carat byte) (string, error) {
	if carat < minCarat || carat > maxCarat {
		return "", fmt.Errorf("Carat must be between '%s' and '%s' inclusively. Given: %s", string(minCarat), string(maxCarat), string(carat))
	}

	triangleHeight := int(carat - minCarat)
	diamondHeight := triangleHeight * 2

	var resultBuffer bytes.Buffer
	for i := 0; i <= diamondHeight; i++ {
		spacingSize := int(math.Abs(float64(triangleHeight - i)))
		spacing := strings.Repeat(string(spacingCharacter), spacingSize)
		letter := string(minCarat + triangleHeight - spacingSize)
		if i == 0 || i == diamondHeight {
			resultBuffer.WriteString(spacing + letter + spacing + "\n")
		} else {
			padding := strings.Repeat(string(spacingCharacter), diamondHeight-spacingSize*2-1)
			resultBuffer.WriteString(spacing + letter + padding + letter + spacing + "\n")
		}
	}

	return resultBuffer.String(), nil
}
