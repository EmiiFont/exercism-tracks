package isbn

import (
	"strings"
)

func IsValidISBN(isbn string) bool {
	res := 0
	isbn = strings.ReplaceAll(isbn, "-", "")

	if len(isbn) != 10 {
		return false
	}

	for i := len(isbn) - 1; i >= 0; i-- {
		digit := isbn[i]
		if len(isbn)-1 == i && digit == 'X' {
			res += 10
			continue
		}
		if digit < '0' || digit > '9' {
			return false
		}

		res += int(digit-'0') * (10 - i)
	}

	return res%11 == 0
}
