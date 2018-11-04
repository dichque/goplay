package main

import "fmt"

const (
	// Typed constant
	a int = 21
	// Untyped constant
	b = 12.23
)

func main() {
	fmt.Println(a, b)
	fmt.Printf("%T\t%T\n", a, b)
}
