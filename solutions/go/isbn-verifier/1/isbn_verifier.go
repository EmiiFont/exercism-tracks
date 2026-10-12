package isbn

import (
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	validator := 1
	res := 0
	isbn = strings.ReplaceAll(isbn, "-", "")

	if len(isbn) != 10 {
		return false
	}

	for i := len(isbn) - 1; i >= 0; i-- {
		if isbn[i] >= '0' && isbn[i] <= '9' || (isbn[i] == 'X' && len(isbn)-1 == i) {
			val := isbn[i]
			num := 0
			if val == 'X' {
				num = 10
			} else {
				converted, err := strconv.Atoi(string(val))
				if err != nil {
					return false
				}
				num = converted
			}
			res += num * validator
			validator++
		} else {
			return false
		}
	}

	return res%11 == 0
}
