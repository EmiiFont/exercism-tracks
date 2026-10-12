package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	result := initial

	for _, val := range s {
		result = fn(result, val)
	}

	return result
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	result := initial

	for i := len(s) - 1; i >= 0; i-- {
		val := s[i]
		result = fn(val, result)
	}
	return result
}

func (s IntList) Filter(fn func(int) bool) IntList {
	res := []int{}

	for _, val := range s {
		if fn(val) {
			res = append(res, val)
		}
	}
	return res
}

func (s IntList) Length() int {
	count := 0
	for range s {
		count++
	}
	return count
}

func (s IntList) Map(fn func(int) int) IntList {
	res := []int{}

	for _, val := range s {
		res = append(res, fn(val))
	}
	return res
}

func (s IntList) Reverse() IntList {
	result := []int{}

	for i := len(s) - 1; i >= 0; i-- {
		result = append(result, s[i])
	}
	return result
}

func (s IntList) Append(lst IntList) IntList {
	result := s

	result = append(result, lst...)

	return result
}

func (s IntList) Concat(lists []IntList) IntList {
	res := s

	for _, val := range lists {
		for _, vd := range val {
			res = append(res, vd)
		}
	}
	return res
}
