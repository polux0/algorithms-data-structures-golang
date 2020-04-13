package main

func sequentialSearch(data []int, value int) bool {
	for index := range data {
		if value == data[index] {
			return true
		}
	}
	return false
}
