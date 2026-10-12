package sieve

func Sieve(limit int) []int {
	primes := []int{}
	sieve := make([]bool, limit)

	if limit < 2 {
		return primes
	}

	for i := 2; i < limit; i++ {
		if !sieve[i] {
			primes = append(primes, i)
			for j := i; j < limit; j += i {
				sieve[j] = true
			}
		}
	}

	if isPrime(limit) {
		primes = append(primes, limit)
	}

	return primes
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
