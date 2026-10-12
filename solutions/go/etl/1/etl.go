package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	res := map[string]int{}

	for key, letters := range in {
		for _, val := range letters {
			res[strings.ToLower(val)] = key
		}
	}
	return res
}
