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
		if compareCharsCount(toCompare, second) {
			res = append(res, val)
		}
	}
	return res
}

func compareCharsCount(first, second map[rune]int) bool {
	for key, count := range second {
		existingCount, ok := first[key]
		if !ok || existingCount != count {
			return false
		}
	}
	return true
}

func mapOfRuneCount(str string) map[rune]int {
	alphabet := map[rune]int{}

	for _, val := range str {
		key := unicode.ToUpper(val)
		alphabet[key]++
	}
	return alphabet
}
