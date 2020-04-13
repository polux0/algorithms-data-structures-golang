package main

import "fmt"

// rotate array: [1,2,3,4,5,6] by 2 -> result [3,4,5,6,1,2]
//					   -> after 1st iteration [6,5,4,3,2,1]
func main() {
	data := []int{1, 2, 3, 4, 5, 6}
	result := reverseArray(data, 0, len(data)-1)
	result1 := reverseArray(result, 3, len(result)-1)
	fmt.Println("Reversed: ", result)
	fmt.Println("Reversed again: ", result1)
}
