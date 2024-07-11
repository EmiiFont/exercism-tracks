package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number > 64 || number <= 0 {
		return 0, errors.New("not a valid chess cell square")
	}

	return uint64(math.Pow(2, float64(number-1))), nil
}

func Total() uint64 {
	var res uint64

	for i := 1; i <= 64; i++ {
		sq, _ := Square(i)
		res += sq
	}

	return res
}
