package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	freq := make(Frequency)
	wordPattern := `\w+('\w+)?`

	re := regexp.MustCompile(wordPattern)

	words := re.FindAllString(phrase, -1)
	for _, val := range words {
		cword := strings.ToLower(val)
		freq[cword]++
	}

	return freq
}
