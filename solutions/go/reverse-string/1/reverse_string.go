package reverse

import "unicode/utf8"

func Reverse(input string) string {
	reversed := ""
	for len(input) > 0 {
		lastChar, size := utf8.DecodeLastRuneInString(input)

		reversed += string(lastChar)
		input = input[:len(input)-size]
	}

	return reversed
}
