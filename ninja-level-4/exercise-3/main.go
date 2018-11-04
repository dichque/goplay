package main

import "fmt"

func main() {

	// Creating a slice
	intSlice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// Get following slices out from intSlice
	// {42, 43, 44, 45, 46}
	// {47, 48, 49, 50, 51}
	// {44, 45, 46, 47, 48}
	// {43, 44, 45, 46, 47}

	fmt.Println(intSlice[:5])
	fmt.Println(intSlice[5:])
	fmt.Println(intSlice[2:7])
	fmt.Println(intSlice[1:6])
}
