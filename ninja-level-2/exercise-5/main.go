package main

import "fmt"

var a int

func main() {
	a = 16
	fmt.Printf("%d\t%b\t%#X\n", a, a, a)

	b := a << 1
	fmt.Printf("%d\t%b\t%#X", b, b, b)
}
