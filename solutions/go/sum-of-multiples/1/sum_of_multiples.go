package summultiples

func SumMultiples(limit int, divisors ...int) int {
	res := map[int]bool{}

	for _, di := range divisors {
		if di == 0 {
			continue
		}
		for i := di; i < limit; i++ {
			if i%di == 0 {
				res[i] = true
			}
		}
	}

	sum := 0
	for val := range res {
		sum += val
	}

	return sum
}
