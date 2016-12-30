// Package beer implements a library for generating the lyrics of the song '99 Bottles of Beer on the Wall'.
package beer

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

const testVersion = 1

// Song returns the whole song.
func Song() string {
	song, _ := Verses(99, 0)

	return song
}

// Verses returns the verses from lowerBound to upperBound inclusively.
// lowerBound and upperBound must be between 0 and 99 inclusively otherwise an error is returned.
func Verses(upperBound, lowerBound int) (string, error) {
	if lowerBound > upperBound {
		return "", fmt.Errorf(
			"Cannot create verses. Lower bound must bot be higher than upper bound. Lower bound was: %s Upper bound was: %s",
			strconv.Itoa(lowerBound),
			strconv.Itoa(upperBound),
		)
	}

	var verses bytes.Buffer
	for i := upperBound; i >= lowerBound; i-- {
		verse, err := Verse(i)
		if err != nil {
			return "", err
		}
		verses.WriteString(verse + "\n")
	}

	return verses.String(), nil
}

// Verse returns the nth verse.
// n must be between 0 and 99 inclusively otherwise an error is returned.
func Verse(n int) (string, error) {
	if n < 0 || n > 99 {
		return "", errors.New("Cannot create verse. Unknown verse index " + strconv.Itoa(n))
	}

	format := "%s of beer on the wall, %s of beer.\n%s, %s of beer on the wall.\n"
	uncapitalisedBottlesPhrase := bottlesPhrase(n)
	capitalizedBottlePhrase := capitalizeFirstCharacter(uncapitalisedBottlesPhrase)
	actionPhrase := actionPhrase(n)
	afterActionBottleCount := n - 1
	if afterActionBottleCount < 0 {
		afterActionBottleCount = 99
	}
	afterActionBottlePhrase := bottlesPhrase(afterActionBottleCount)

	verse := fmt.Sprintf(format, capitalizedBottlePhrase, uncapitalisedBottlesPhrase, actionPhrase, afterActionBottlePhrase)

	return verse, nil
}

func bottlesPhrase(n int) string {
	switch n {
	case 0:
		return "no more bottles"
	case 1:
		return "1 bottle"
	default:
		return strconv.Itoa(n) + " bottles"
	}
}

func capitalizeFirstCharacter(input string) string {
	runes := []rune(input)
	runes[0] = unicode.ToUpper(runes[0])

	return string(runes)
}

func actionPhrase(n int) string {
	switch n {
	case 0:
		return "Go to the store and buy some more"
	case 1:
		return "Take it down and pass it around"
	default:
		return "Take one down and pass it around"
	}
}
