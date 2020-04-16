package main

import "fmt"

func util(num int, from string, to string, temp string) {
	if num < 1 {
		return
	}
	util(num-1, from, temp, to)
	fmt.Println("Move disk ", num, " from peg ", from, " to peg ", to)
	util(num-1, temp, to, from)
}
func TowersOfHanoi(num int) {
	fmt.Println("The sequence of moves involved in the Tower of Hanoi are: ")
	util(num, "A", "C", "B")
}
