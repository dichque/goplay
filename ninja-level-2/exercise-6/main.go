package main

import "fmt"

const (
	// This works too..
	// year1 = (iota + 2019)
	// year2
	// year3
	// year4

	year1 = 2019 + iota
	year2 = 2019 + iota
	year3 = 2019 + iota
	year4 = 2019 + iota
)

func main() {
	fmt.Println(year1)
	fmt.Println(year2)
	fmt.Println(year3)
	fmt.Println(year4)
}
