package romannumerals

import (
	"errors"
)

func Divider(rem, digit int) int {
	if rem < 5 {
		return 1
	}
	switch digit {
	case 4:
		return 1000
	case 3:
		return 100
	case 2:
		return 10
	default:
		return 5
	}
}

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input >= 4000 {
		return "", errors.New("not valid input")
	}

	romanHundreds := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	romanTens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	romanOnes := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	romanThousands := []string{"", "M", "MM", "MMM"}

	res := romanThousands[input/1000] + romanHundreds[(input%1000)/100] + romanTens[(input%100)/10] + romanOnes[input%10]

	return res, nil
}
