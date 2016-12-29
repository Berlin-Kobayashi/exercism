// Package brackets implements a function for checking if brackets are properly closed and nested.
package brackets

import (
	"errors"
	"unicode"
)

const testVersion = 4

type runeStack []rune

func (s *runeStack) Push(v rune) {
	*s = append(*s, v)
}

func (s *runeStack) Pop() (rune, error) {
	l := len(*s)

	if l == 0 {
		return 0, errors.New("Stack is empty")
	}

	res := (*s)[l-1]
	*s = (*s)[:l-1]
	return res, nil
}

// Bracket checks if brackets are properly closed and nested.
func Bracket(input string) (bool, error) {
	bracketStack := runeStack{}
	for _, runeValue := range input {
		switch {
		case unicode.Is(unicode.Ps, runeValue):
			bracketStack.Push(runeValue)
		case unicode.Is(unicode.Pe, runeValue):
			previousBracket, err := bracketStack.Pop()
			if err != nil {
				return false, nil
			}

			if !isSameBracketKind(previousBracket, runeValue) {
				return false, nil
			}
		}
	}

	return len(bracketStack) == 0, nil
}

func isSameBracketKind(openingBracket, closingBracket rune) bool {
	return openingBracket == closingBracket-1 || openingBracket == closingBracket-2
}
