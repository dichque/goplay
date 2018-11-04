package main

import "fmt"

func main() {

	// Creating an array
	a := [5]int{}

	a[0] = 12
	a[1] = 13
	a[2] = 14
	a[3] = 15
	a[4] = 16

	for idx, value := range a {
		fmt.Println(idx, value)
	}

	fmt.Printf("Array a is of type: %T", a)
}
