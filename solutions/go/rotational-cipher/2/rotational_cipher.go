package rotationalcipher

import (
	"math"
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {
	runeArr := []rune{}
	for _, val := range plain {
		defaultVal := 'a'
		if !unicode.IsLetter(val) {
			runeArr = append(runeArr, val)
			continue
		}
		if unicode.IsUpper(val) {
			defaultVal = 'A'
		}
		id := math.Abs(float64(defaultVal-val)) + float64(shiftKey)
		runeArr = append(runeArr, defaultVal+rune(int(id)%26))
	}

	return string(runeArr)
}
