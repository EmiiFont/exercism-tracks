package luhn

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

func cleanId(id string) (string, error) {
	clean := strings.TrimSpace(id)
	process := []rune(clean)
	res := ""
	for i := 0; i < len(process); i++ {
		if unicode.IsSpace(process[i]) {
			continue
		}
		if !unicode.IsDigit(process[i]) {
			return "", errors.New("invalid input")
		}
		res += string(process[i])
	}
	return res, nil
}

func Valid(id string) bool {
	cleaned, err := cleanId(id)
	if err != nil {
		return false
	}
	if len(cleaned) <= 1 {
		return false
	}

	process := []rune(cleaned)
	res := 0
	for i := len(process) - 2; i >= 0; i -= 2 {

		dig, _ := strconv.Atoi(string(process[i]))
		dig *= 2
		if dig > 9 {
			dig -= 9
		}
		normal, _ := strconv.Atoi(string(process[i+1]))
		res += (dig + normal)
	}

	if len(cleaned)%2 != 0 {
		firstDigit, _ := strconv.Atoi(string(process[0]))
		res += firstDigit
	}

	return res%10 == 0
}
