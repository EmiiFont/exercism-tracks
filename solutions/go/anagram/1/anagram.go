package anagram

import (
	"strings"
	"unicode"
)

func Detect(subject string, candidates []string) []string {
	toCompare := mapOfRuneCount(subject)
	res := []string{}
	for _, val := range candidates {
		if strings.EqualFold(val, subject) {
			continue
		}
		second := mapOfRuneCount(val)
		add := true
		for key, count := range second {
			existingCount, ok := toCompare[key]
			if !ok {
				add = false
				break
			}
			if existingCount != count {
				add = false
				break
			}
		}
		if add {
			res = append(res, val)
		}
	}
	return res
}

func mapOfRuneCount(str string) map[rune]int {
	alphabet := map[rune]int{}

	for _, val := range str {
		key := unicode.ToUpper(val)
		alphabet[key] += 1
	}
	return alphabet
}
