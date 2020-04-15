package main

import "fmt"

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
