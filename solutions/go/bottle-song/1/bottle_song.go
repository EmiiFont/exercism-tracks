package bottlesong

import (
	"fmt"
	"strings"
)

func Recite(startBottles, takeDown int) []string {
	res := []string{}

	numberIntToName := map[int]string{
		0:  "no",
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
	}

	for takeDown > 0 {
		plural := "bottles"
		singular := "bottles"
		if startBottles == 1 {
			plural = "bottles"
			singular = "bottle"
		}
		a1 := fmt.Sprintf("%s green %s hanging on the wall,", strings.Title(numberIntToName[startBottles]), singular)
		a2 := fmt.Sprint("And if one green bottle should accidentally fall,")
		a3 := fmt.Sprintf("There'll be %s green %s hanging on the wall.", numberIntToName[startBottles-1], plural)
		res = append(res, a1, a1, a2, a3)
		if takeDown > 1 {
			res = append(res, "")
		}
		startBottles--
		takeDown--
	}

	return res
}
