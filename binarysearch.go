package main

func binarySearch(data []int, value int) bool {
	low := 0
	var mid int
	high := len(data)
	for low <= high {
		mid = low + (high-low)/2
		if data[mid] == value {
			return true
		} else {
			if data[mid] < value {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return false
}
