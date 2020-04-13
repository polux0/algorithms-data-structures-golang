package main

func SumArray(data []int) int {
	var sum int
	for _, value := range data {
		sum += data[value]
	}
	return sum
}
