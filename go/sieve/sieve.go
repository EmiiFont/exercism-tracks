package sieve

func Sieve(limit int) []int {
	primes := []int{}
	sieve := map[int]bool{}

	if limit < 2 {
		return primes
	}

	for i := 2; i <= limit; i++ {
		if !sieve[i] {
			primes = append(primes, i)
			for j := i; j <= limit; j += i {
				sieve[j] = true
			}
		}
	}

	return primes
}
