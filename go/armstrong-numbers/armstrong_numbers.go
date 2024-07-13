package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	if n == 0 {
		return true
	}
	numberOfDigits := 0

	for i := n; i > 0; i /= 10 {
		numberOfDigits++
	}

	str := strconv.Itoa(n)

	sum := 0
	for _, c := range str {
		digit, _ := strconv.Atoi(string(c))
		sum += int(math.Pow(float64(digit), float64(numberOfDigits)))
	}

	return sum == n
}
