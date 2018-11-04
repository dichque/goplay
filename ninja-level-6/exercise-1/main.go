package main

import "fmt"

func main() {
	fooResult := foo()
	barResult1, barResult2 := bar()

	fmt.Println(fooResult)
	fmt.Println(barResult1, barResult2)
}

func foo() int {
	return 1
}

func bar() (int, string) {
	return 1, "hello"
}
