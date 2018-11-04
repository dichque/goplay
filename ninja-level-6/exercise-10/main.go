package main

import "fmt"

func foo(a string) {
	fmt.Println("Value of a inside of foo: ", a)

	// Begining of a closure block
	{
		a := "Howdy !!"
		fmt.Println("Value of a inside of foo & within closure block: ", a)
	}

}

func main() {
	foo("Hi there")
}
