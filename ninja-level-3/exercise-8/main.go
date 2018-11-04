package main

import "fmt"

func main() {

	fruit := "Fig"

	switch fruit {
	case "Apple":
		fmt.Println("We have apple")
	case "Orange":
		fmt.Println("We have orange")
	default:
		fmt.Println("Don't know what we have")
	}
}
