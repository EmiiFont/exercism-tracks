package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span > len(digits) || span < 0 {
		return 0, errors.New("span must be smaller that digits length")
	}
	start := 0
	mov := 0
	var max int64

	for mov < len(digits) {
		if mov-start == (span - 1) {
			str := digits[start : mov+1]
			product := int64(1)
			for _, val := range str {
				num, err := strconv.Atoi(string(val))
				if err != nil {
					return 0, errors.New("not valid character")
				}
				product *= int64(num)
			}
			if product > max {
				max = product
			}
			start++
		} else {
			mov++
		}
	}
	return int64(max), nil
}
