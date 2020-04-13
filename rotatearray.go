package main

// [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
func rotateArray(data []int, position int) {

}
func reverseArray(data *[]int, start int, end int) []int {
	for start < end {
		&data[start], &data[end] = &data[end], &data[start]
		start++
		end--
	}
	return data
}
