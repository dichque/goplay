package main

import "fmt"

func main() {

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50}
	x = append(x, 52)

	fmt.Println(x)

	x = append(x, 53, 54, 55)
	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}
	// We need to unfurl y using ... operator
	x = append(x, y...)
	fmt.Println(x)
}
