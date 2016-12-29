// Package house implements a library for generating the lyrics of the song 'This is the House that Jack Built'.
package house

// Song generates the full text of "The House That Jack Built".
func Song() (song string) {
	phrases := []struct {
		object, predicate string
	}{
		{"the house that Jack built", "lay in"},
		{"the malt", "ate"},
		{"the rat", "killed"},
		{"the cat", "worried"},
		{"the dog", "tossed"},
		{"the cow with the crumpled horn", "milked"},
		{"the maiden all forlorn", "kissed"},
		{"the man all tattered and torn", "married"},
		{"the priest all shaven and shorn", "woke"},
		{"the rooster that crowed in the morn", "kept"},
		{"the farmer sowing his corn", "belonged to"},
		{"the horse and the hound and the horn", ""},
	}
	for v := range phrases {
		prefixPhrase := ""
		relPhrases := []string{}
		suffixPhrase := ""
		for i := v; i >= 0; i-- {
			phrase := phrases[i]
			switch {
			case i == v && 0 == v:
				prefixPhrase = Embed("This is", phrase.object) + "."
				if v != len(phrases)-1 {
					prefixPhrase += "\n"
				}
			case i == v:
				prefixPhrase = Embed("This is", phrase.object) + "\n"
			case i == 0:
				suffixPhrase = Embed(Embed("that", phrase.predicate), phrase.object) + "."
				if v != len(phrases)-1 {
					suffixPhrase += "\n"
				}
			default:
				relPhrases = append(relPhrases, Embed(Embed("that", phrase.predicate), phrase.object)+"\n")
			}
		}
		verse := Verse(prefixPhrase, relPhrases, suffixPhrase)
		song += verse
		if v != len(phrases)-1 {
			song += "\n"
		}
	}

	return song
}

// Verse generates a verse with relative phrases that have a recursive structure.
func Verse(prefixPhrase string, relPhrases []string, suffixPhrase string) (verse string) {
	verse = prefixPhrase

	for _, relPhrase := range relPhrases {
		verse = Embed(verse, relPhrase)
	}

	verse = Embed(verse, suffixPhrase)

	return verse
}

// Embed embeds a phrase into another phrase.
func Embed(prefixPhrase, suffixPhrase string) (phrase string) {
	if prefixPhrase == "" || suffixPhrase == "" || prefixPhrase[len(prefixPhrase)-1] == '\n' || suffixPhrase[0] == '\n' {
		phrase = prefixPhrase + suffixPhrase
	} else {
		phrase = prefixPhrase + " " + suffixPhrase
	}

	return phrase
}
