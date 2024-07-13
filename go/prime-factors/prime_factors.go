package prime

func Factors(n int64) []int64 {
	if n == 1 {
		return []int64{}
	}
	prime := int64(2)
	res := []int64{}
	for n > 1 {
		if n%prime == 0 {
			n = n / prime
			res = append(res, prime)
		} else {
			prime++
		}
	}
	return res
}
