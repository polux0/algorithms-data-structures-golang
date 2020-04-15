package main

import "fmt"

func main() {
	data := []int{1, 1, -2, 3, 4, -4, 6, -14, 8, 2}
	result := largestSubArraySum(data)
	fmt.Println(result)
}
