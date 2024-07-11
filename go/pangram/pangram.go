package pangram

import (
	"unicode"
)

func IsPangram(input string) bool {
	alpha := map[rune]bool{}

	for _, val := range input {
		if unicode.IsLetter(val) {
			alpha[unicode.ToLower(val)] = true
		}
	}

	return len(alpha) == 26
}
