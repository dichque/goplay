package main

import "fmt"

func main() {

	fruit := "Apple"

	if fruit == "Apple" {
		fmt.Println("We have apple\n")
	} else if fruit == "Orange" {
		fmt.Println("We have oranges\n")
	} else {
		fmt.Println("Don't know what we have")
	}
}
