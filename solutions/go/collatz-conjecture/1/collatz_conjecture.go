package collatzconjecture

import "errors"

func isEven(n int) bool {
	return n%2 == 0
}

func CollatzConjecture(n int) (int, error) {
	counter := 0
	if n <= 0 {
		return 0, errors.New("error")
	}

	for n > 1 {
		if isEven(n) {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		counter++
	}
	return counter, nil
}
