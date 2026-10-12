package pangram

import (
	"strings"
	"unicode"
)

func IsPangram(input string) bool {
	alpha := map[string]bool{}

	for _, val := range input {
		if unicode.IsLetter(val) {
			alpha[strings.ToLower(string(val))] = true
		}
	}

	return len(alpha) == 26
}
