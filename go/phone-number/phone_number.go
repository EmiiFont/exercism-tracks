package phonenumber

import (
	"errors"
	"fmt"
)

func Number(phoneNumber string) (string, error) {
	cleaned := cleanNumber(phoneNumber)

	if len(cleaned) == 11 && cleaned[0] == '1' && isValidNumber(cleaned[1:]) {
		return cleaned[1:], nil
	}
	if len(cleaned) == 10 && isValidNumber(cleaned) {
		return cleaned, nil
	}
	return "", errors.New("error")
}

func isValidNumber(phoneNumber string) bool {
	for i := 0; i < len(phoneNumber); i++ {
		if (i == 0 || i == 3) && phoneNumber[i] < '2' {
			return false
		}
	}
	return true
}

func cleanNumber(phoneNumber string) string {
	cleaned := ""
	for _, val := range phoneNumber {
		if val >= '0' && val <= '9' {
			cleaned += string(val)
		}
	}
	return cleaned
}

func AreaCode(phoneNumber string) (string, error) {
	res, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	return res[0:3], nil
}

func Format(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("(%s) %s-%s", number[:3], number[3:6], number[6:]), nil
}
