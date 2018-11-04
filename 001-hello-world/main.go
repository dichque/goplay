package main

import "fmt"

var hi = "Hi"

// Declare a variable  with identifier z and type int
// Assigns the value of zero: refer: https://golang.org/ref/spec#The_zero_value

var z int

// This is a STATIC language a variable is DECLARED to hold a VALUR
// of certain TYPE
var y = 42

// Raw string literal : https://golang.org/ref/spec#String_literals
var rw = `Hello world "GOlang" continued to next line....

`

func main() {

	//Declare a variable and assign a value using short declaration operator
	greeting := "Hello GOLang !!"

	n, _ := fmt.Println(greeting)

	fmt.Println(n)
	foo()
}

func foo() {
	fmt.Println(hi)
	fmt.Printf("%T\n", hi)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(rw)

	// fmt.Printf("%#x \t %b \t %x", y, y, y)
	s := fmt.Sprintf("%#x \t %b \t %x", y, y, y)
	fmt.Println(s)
}
