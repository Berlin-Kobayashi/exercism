// Package twelve implements a simple library for the lyrics of the 'The Twelve Days of Christmas' song.
// http://en.wikipedia.org/wiki/The_Twelve_Days_of_Christmas_(song)
package twelve

import (
	"fmt"
)

const testVersion = 1

type present struct {
	day, content string
}

var presents = []present{
	{"first", "a Partridge in a Pear Tree"},
	{"second", "two Turtle Doves"},
	{"third", "three French Hens"},
	{"fourth", "four Calling Birds"},
	{"fifth", "five Gold Rings"},
	{"sixth", "six Geese-a-Laying"},
	{"seventh", "seven Swans-a-Swimming"},
	{"eighth", "eight Maids-a-Milking"},
	{"ninth", "nine Ladies Dancing"},
	{"tenth", "ten Lords-a-Leaping"},
	{"eleventh", "eleven Pipers Piping"},
	{"twelfth", "twelve Drummers Drumming"},
}

// Song returns the lyrics of the whole song.
func Song() string {
	song := ""

	for i := 1; i <= len(presents); i++ {
		song += Verse(i) + "\n"
	}

	return song
}

// Verse returns the lyrics of the verse for the given verse number.
func Verse(verseNumber int) string {
	verse := "."

	for presentIndex := 0; presentIndex < verseNumber; presentIndex++ {
		verse = presents[presentIndex].content + verse

		if verseNumber != 1 && presentIndex == 0 {
			verse = "and " + verse
		}

		verse = ", " + verse
	}

	verseIndex := verseNumber - 1
	verse = fmt.Sprintf("On the %s day of Christmas my true love gave to me", presents[verseIndex].day) + verse

	return verse
}
