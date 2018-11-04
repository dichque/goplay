package main

import "fmt"

var (
	x1 = 42
	y1 = "James Bond"
	z1 = true
)

func main() {

	s := fmt.Sprintf("%d \t %s \t %t", x1, y1, z1)

	fmt.Println(s)
}
