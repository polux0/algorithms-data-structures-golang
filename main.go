package main

// rotate array: [1,2,3,4,5,6,7,8,9,10] by 2 -> result [3,4,5,6,7,8,9,10,1,2]
func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// should be [10, 9, 8, 7, 6, 5, 4, 3, 2, 1]
	// if we reverse 1/2, should be [5, 4, 3, 2, 1, 6, 7, 8, 9, 10]
	rotateArray(data, 2)
}
