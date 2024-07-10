package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	reverse := "zyxwvutsrqponmlkjihgfedcba"

	res := []rune{}
	count := 0
	for i := 0; i < len(s); i++ {
		val := unicode.ToLower(rune(s[i]))
		if count == 5 && i != len(s)-1 {
			res = append(res, ' ')
			count = 0
		}
		if unicode.IsSpace(val) || unicode.IsPunct(val) {
			continue
		}
		if val >= '0' && val <= '9' {
			res = append(res, val)
		}
		if val >= 'a' && val <= 'z' {
			idx := strings.IndexRune(alphabet, val)
			res = append(res, rune(reverse[idx]))
		}
		count++
	}

	return string(res)
}
