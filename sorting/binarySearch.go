package main

func binarySearch(arr []int, haystack int) (int, bool) {
	low := 0
	high := len(arr)

	for high > low {
		mid := (high + low) / 2
		if haystack == mid {
			return mid, true
		} else if haystack > mid {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return -1, false
}
