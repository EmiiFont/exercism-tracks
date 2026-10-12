package rotationalcipher

import (
	"math"
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {
	runeArr := []rune{}
	for _, val := range plain {
		defaultVal := 'a'
		toUp := false
		if !unicode.IsLetter(val) {
			runeArr = append(runeArr, val)
			continue
		}
		if unicode.IsUpper(val) {
			defaultVal = 'A'
			toUp = true
		}
		id := math.Abs(float64(defaultVal-val)) + float64(shiftKey)
		if toUp {
			runeArr = append(runeArr, defaultVal+rune(int(id)%26))
		} else {
			runeArr = append(runeArr, defaultVal+rune(int(id)%26))
		}
	}

	return string(runeArr)
}
