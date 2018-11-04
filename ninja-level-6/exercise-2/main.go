package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}

	totalFoo := foo(a...)
	fmt.Println(totalFoo)

	totalBar := bar(a)
	fmt.Println(totalBar)
}

func foo(num ...int) int {
	sum := 0
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}
	return sum
}

func bar(num []int) int {
	sum := 0
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}
	return sum
}
