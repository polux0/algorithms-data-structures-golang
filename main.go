package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := binarySearch(data, 8)
	fmt.Println("Element exists?", result)
}
