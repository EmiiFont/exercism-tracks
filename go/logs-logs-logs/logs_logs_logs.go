package logs

import (
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, val := range log {
		switch val {
		case '❗':
			return "recommendation"
		case '🔍':
			return "search"
		case '☀':
			return "weather"
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	res := []rune(log)

	for idx, val := range res {
		if val == oldRune {
			val = newRune
		}
		res[idx] = val
	}
	return string(res)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	numberOfRunes := utf8.RuneCountInString(log)
	return numberOfRunes <= limit
}
