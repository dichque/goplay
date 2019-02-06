package main

import "fmt"

type pizza int

var b pizza
var a int

// a and b are DIFFERENT types

func main() {
	b = 42
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	a = 32
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	// a and b are DIFFERENT types
	// You cannot assign same values to each other
	// Error: cannot use b (type pizza) as type int in assignment
	// a = b

	// We are converting b to int type
	// In other language it is called casting, but not in GO
	// https://golang.org/ref/spec#Conversions
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
