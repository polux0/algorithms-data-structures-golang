package main

func factorial(of int) int {
	if of <= 1 {
		return 1
	}
	return of * factorial(of-1)
}
