package encode

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func RunLengthEncode(input string) string {
	if len(input) == 0 {
		return ""
	}

	count := 0
	res := ""
	for i := 0; i < len(input); {
		count = 1
		for i+1 < len(input) && input[i] == input[i+1] {
			i += 1
			count += 1
		}

		if count == 1 {
			res += string(input[i])
		} else {
			res += fmt.Sprintf("%d%s", count, string(input[i]))
		}
		i += 1
	}

	return res
}

func RunLengthDecode(input string) string {
	if len(input) == 0 {
		return ""
	}
	res := ""

	num := ""
	for i := 0; i < len(input); i++ {
		if rune(input[i]) >= '0' && rune(input[i]) <= '9' {
			num += string(input[i])
		}
		if num == "" && !unicode.IsNumber(rune(input[i])) {
			res += string(input[i])
		}
		if num != "" && !unicode.IsNumber(rune(input[i])) {
			count, _ := strconv.Atoi(num)
			res += strings.Repeat(string(input[i]), count)
			num = ""
		}
	}
	if res == "" {
		return input
	}
	return res
}
