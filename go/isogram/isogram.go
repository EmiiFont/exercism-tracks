package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	hash := map[rune]bool{}

	for _, k := range strings.ToLower(word) {
		_, ok := hash[k]
		if ok {
			return false
		}
		if unicode.IsLetter(k) {
			hash[k] = true
		}
	}
	return true
}
