package prime

import (
	"errors"
	"math"
)

func isPrime(number int) bool {
	if number%2 == 0 {
		return false
	}

	divisor := math.Floor(math.Sqrt(float64(number)))

	for i := 3; i <= int(divisor); i += 2 {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("value not accepted")
	}

	if n == 1 {
		return 2, nil
	}

	i := 2
	prime := 2
	count := 1
	for count < n {
		if isPrime(i) {
			count++
		}
		if count == n {
			prime = i
		}
		i++
	}
	return prime, nil
}
