package series

func All(n int, s string) []string {
	res := []string{}
	for i := 0; i < len(s); i++ {
		if i+n <= len(s) {
			res = append(res, s[i:i+n])
		}
	}
	return res
}

func UnsafeFirst(n int, s string) string {
	if n > len(s) {
		return s
	}
	return All(n, s)[0]
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return s, false
	}
	return All(n, s)[0], true
}
