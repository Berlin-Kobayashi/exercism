// Package foodchain implements a library for creating the lyrics of the song 'I Know an Old Lady Who Swallowed a Fly' or parts of it.
package foodchain

import (
	"fmt"
	"strings"
)

const testVersion = 3

type food struct {
	Name, Comment string
}

var foods = []food{
	{"fly", "I don't know why she swallowed the fly. Perhaps she'll die."},
	{"spider", "It wriggled and jiggled and tickled inside her."},
	{"bird", "How absurd to swallow a bird!"},
	{"cat", "Imagine that, to swallow a cat!"},
	{"dog", "What a hog, to swallow a dog!"},
	{"goat", "Just opened her throat and swallowed a goat!"},
	{"cow", "I don't know how she swallowed a cow!"},
	{"horse", "She's dead, of course!"},
}

// Song returns the whole song.
func Song() string {
	return Verses(1, len(foods))
}

// Verses returns the verses from from to to inclusive.
func Verses(from, to int) string {
	verses := []string{}

	for i := from; i <= to; i++ {
		verses = append(verses, Verse(i))
	}

	return strings.Join(verses, "\n\n")
}

// Verse returns the verse for the given index.
func Verse(index int) string {
	index--

	phrases := []string{}
	phrases = append(phrases, fmt.Sprintf("I know an old lady who swallowed a %s.", foods[index].Name))
	phrases = append(phrases, foods[index].Comment)

	if index != len(foods)-1 {
		for i := index - 1; i >= 0; i-- {
			if i == 1 {
				phrases = append(phrases, fmt.Sprintf("She swallowed the %s to catch the %s that %s", foods[i+1].Name, foods[i].Name, strings.SplitAfterN(foods[i].Comment, " ", 2)[1]))
			} else {
				phrases = append(phrases, fmt.Sprintf("She swallowed the %s to catch the %s.", foods[i+1].Name, foods[i].Name))
			}

			if i == 0 {
				phrases = append(phrases, foods[i].Comment)
			}
		}
	}

	return strings.Join(phrases, "\n")
}
