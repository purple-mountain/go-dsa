package main

func bubbleSort(arr []int) []int {
	for n := len(arr); n > 1; n-- {
		for i := 0; i < n-1; i++ {
			if arr[i] > arr[i+1] {
				temp := arr[i+1]
				arr[i+1] = arr[i]
				arr[i] = temp
			}
		}
	}
	return arr
}
