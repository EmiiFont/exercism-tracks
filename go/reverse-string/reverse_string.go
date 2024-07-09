package reverse

func Reverse(input string) string {
	reversed := []rune(input)
	start := 0
	end := len(reversed) - 1
	for start <= end {
		reversed[start], reversed[end] = reversed[end], reversed[start]
		start++
		end--
	}

	return string(reversed)
}
