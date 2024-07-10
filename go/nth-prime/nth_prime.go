package prime

import (
	"errors"
	"math"
)

func isPrime(number int) bool {
	if number < 2 {
		return false
	}

	divisor := math.Floor(math.Sqrt(float64(number)))

	for i := 2; i <= int(divisor); i++ {
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

	i := 0
	prime := 0
	count := 0
	for {
		i++
		if isPrime(i) {
			count++
		}
		if count == n {
			prime = i
			break
		}
	}
	return prime, nil
}
