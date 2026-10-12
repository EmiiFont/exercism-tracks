package bottlesong

import (
	"fmt"
	"strings"
)

var numberIntToName = map[int]string{
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

func bottle(num int) string {
	if num == 1 {
		return "bottle"
	}
	return "bottles"
}

func Recite(startBottles, takeDown int) []string {
	res := []string{}
	for i := startBottles; i > startBottles-takeDown; i-- {
		a1 := fmt.Sprintf("%s green %s hanging on the wall,", strings.Title(numberIntToName[i]), bottle(i))
		a3 := fmt.Sprintf("There'll be %s green %s hanging on the wall.", numberIntToName[i-1], bottle(i-1))

		res = append(res, a1, a1, "And if one green bottle should accidentally fall,", a3, "")
	}
	return res[:len(res)-1]
}
