package main

import "fmt"

func main() {

	a := func() {
		fmt.Println("I am an anonymous function without a name !!")
	}

	a()
	fmt.Printf("This is variable a of type: %T", a)

}
