package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("different size")
	}

	count := 0
	for idx, val := range a {
		if rune(b[idx]) != val {
			count++
		}
	}

	return count, nil
}
