package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}

	defer foo(a...)
	bar(a)
}

func foo(num ...int) int {
	sum := 0
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}
	fmt.Println("Exiting foo now !!")
	return sum
}

func bar(num []int) int {
	sum := 0
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}
	fmt.Println("Exiting bar now !!")
	return sum
}
