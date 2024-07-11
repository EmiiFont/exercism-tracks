package strain

func Keep[T any](s []T, x func(T) bool) []T {
	res := []T{}
	for _, v := range s {
		if x(v) {
			res = append(res, v)
		}
	}
	return res
}

func Discard[T any](s []T, x func(T) bool) []T {
	res := []T{}

	for _, v := range s {
		if !x(v) {
			res = append(res, v)
		}
	}
	return res
}
