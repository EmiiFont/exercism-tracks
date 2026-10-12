// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

func isAllCaps(s string) bool {
	cleaned := ""

	for _, f := range s {
		if unicode.IsDigit(f) || unicode.IsPunct(f) || unicode.IsSpace(f) {
			continue
		}
		cleaned += string(f)
	}

	if len(cleaned) == 0 {
		return false
	}

	for _, k := range cleaned {
		if !unicode.IsUpper(k) && unicode.IsLetter(k) {
			return false
		}
	}
	return true
}

func isQuestion(s string) bool {
	return strings.HasSuffix(s, "?")
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if remark == "" {
		return "Fine. Be that way!"
	}
	if isQuestion(remark) && !isAllCaps(remark) {
		return "Sure."
	}

	if !isQuestion(remark) && isAllCaps(remark) {
		return "Whoa, chill out!"
	}
	if isQuestion(remark) && isAllCaps(remark) {
		return "Calm down, I know what I'm doing!"
	}

	return "Whatever."
}
