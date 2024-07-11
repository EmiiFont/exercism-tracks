package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) <= 1 {
		return false
	}

	isSecond := len(id)%2 == 0

	res := 0
	for _, val := range id {
		dig, err := strconv.Atoi(string(val))
		if err != nil {
			return false
		}
		if isSecond {
			dig *= 2
			if dig > 9 {
				dig -= 9
			}
		}
		res += dig
		isSecond = !isSecond
	}

	return res%10 == 0
}
