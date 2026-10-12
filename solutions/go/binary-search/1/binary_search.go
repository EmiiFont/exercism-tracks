package binarysearch

import (
	"math"
	"sort"
)

func SearchInts(list []int, key int) int {
	sortedList := sort.IntSlice(list)
	start := 0
	end := len(sortedList) - 1

	for start <= end {
		mid := int(math.Floor(float64(end+start) / 2))

		if sortedList[mid] == key {
			return mid
		}
		if sortedList[mid] < key {
			start = mid + 1
		} else if sortedList[mid] > key {
			end = mid - 1
		}
	}

	return -1
}
