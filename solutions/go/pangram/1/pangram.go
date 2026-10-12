package pangram

import (
	"strings"
	"unicode"
)

func IsPangram(input string) bool {
	alpha := map[string]int{}

	for _, val := range input {
		if unicode.IsLetter(val) {
			alpha[strings.ToLower(string(val))] = 0
		}
	}

	return len(alpha) == 26
}
