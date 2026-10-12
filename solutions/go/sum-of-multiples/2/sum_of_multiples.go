package summultiples

func SumMultiples(limit int, divisors ...int) int {
	res := map[int]bool{}

	sum := 0
	for _, di := range divisors {
		if di == 0 {
			continue
		}
		for i := di; i < limit; i++ {
			_, ok := res[i]
			if i%di == 0 && !ok {
				res[i] = true
				sum += i
			}
		}
	}

	return sum
}
