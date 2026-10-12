package resistorcolorduo

import "strconv"

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	resistorColorMap := map[string]int{
		"black":  0,
		"brown":  1,
		"red":    2,
		"orange": 3,
		"yellow": 4,
		"green":  5,
		"blue":   6,
		"violet": 7,
		"grey":   8,
		"white":  9,
	}

	res := ""
	for i := 0; i < 2; i++ {
		val, ok := resistorColorMap[colors[i]]
		if ok {
			res += strconv.Itoa(val)
		}
	}

	result, _ := strconv.Atoi(res)
	return result
}
