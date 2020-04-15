package main

import "fmt"

// [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
// [3, 4, 5, 6, 7, 8, 9, 10, 1, 2]
func rotateArray(data []int, position int) {
	length := len(data) - 1
	reverseArray(data, 0, position-1)
	reverseArray(data, position, length)
	reverseArray(data, 0, length)
	fmt.Println(data)
}
func reverseArray(data []int, start int, end int) []int {
	for start < end {
		data[start], data[end] = data[end], data[start]
		start++
		end--
	}
	return data
}
