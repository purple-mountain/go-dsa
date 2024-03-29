package main

func binarySearch(arr []int, haystack int) bool {
	low := 0
	high := len(arr)

	for high > low {
		mid := (high + low) / 2
		if haystack == mid {
			return true
		} else if haystack > mid {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return false
}
