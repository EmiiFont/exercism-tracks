package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("different size")
	}

	count := 0
	for idx := range a {
		if b[idx] != a[idx] {
			count++
		}
	}

	return count, nil
}
