package main

import "fmt"

func main() {

	favSport := "Dakar"

	switch favSport {
	case "Dakar":
		fmt.Println("Yay, lets watch!!")
	case "Cricket":
		fmt.Println("Yawn..")
	default:
		fmt.Println("Huh ?")
	}
}
