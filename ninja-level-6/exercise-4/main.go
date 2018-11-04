package main

import "fmt"

type person struct {
	first, last string
	age         int
}

func (p person) speak() {
	fmt.Println(p.first, p.last, p.age)
}

func main() {
	p := person{
		first: "Jagadish",
		last:  "Nagarajaiah",
		age:   41,
	}

	p.speak()
}
