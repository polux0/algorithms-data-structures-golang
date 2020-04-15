package main

func largestSubArraySum(data []int) int {
	length := len(data)
	localMaximum := 0
	highestMaximum := 0

	for i := 0; i < length; i++ {
		localMaximum = localMaximum + data[i]
		if localMaximum < 0 {
			localMaximum = 0
		}
		if highestMaximum < localMaximum {
			highestMaximum = localMaximum
		}
	}
	return localMaximum
}
