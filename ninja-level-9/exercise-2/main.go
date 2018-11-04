package main

import "fmt"

// Person has a first and lastname
type Person struct {
	first, last string
}

// Speak method says Howdy
func (p *Person) Speak() {
	fmt.Printf("Howdy from %v !!", p.first)
}

// Human type should Speak
type Human interface {
	Speak()
}

func saySomething(a Human) {
	a.Speak()
}

func main() {
	p1 := Person{
		first: "Jack",
		last:  "Daniels",
	}

	saySomething(&p1)

	// This will not work
	// saySomething(p1)
}
