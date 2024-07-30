package binarysearch

func SearchInts(list []int, key int) int {
	start := 0
	end := len(list) - 1

	for start <= end {
		mid := start + (end-start)/2

		if list[mid] == key {
			return mid
		}
		if list[mid] < key {
			start = mid + 1
		} else if list[mid] > key {
			end = mid - 1
		}
	}

	return -1
}
