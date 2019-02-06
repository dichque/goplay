package main

import "fmt"

type Kitchen struct {
	numOfForks int
	numOfKnives int
}

func(k Kitchen) totalForksAndKnives() int {
	return k.numOfForks + k.numOfKnives
}

type House struct {
	Kitchen //anonymous field
}

func main() {
	h := House{Kitchen{4, 4}} //the kitchen has 4 forks and 4 knives
	fmt.Println("Sum of forks and knives in house: ", h.totalForksAndKnives())  //called on House even though the method is associated with Kitchen
}
