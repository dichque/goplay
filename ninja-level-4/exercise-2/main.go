package main

import "fmt"

func main() {

	// Creating a slice
	intSlice := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	for idx, value := range intSlice {
		fmt.Println(idx, value)
	}

	fmt.Printf("Array a is of type: %T", intSlice)
}
