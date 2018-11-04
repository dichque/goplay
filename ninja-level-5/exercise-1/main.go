package main

import "fmt"

type person struct {
	firstName  string
	lastName   string
	favFlavors []string
}

// func(r recevier) identifier (params) (return) {}
func (p person) String() string {
	return fmt.Sprintf("%s  %s ", p.firstName, p.lastName)
}

func main() {
	a := person{
		firstName:  "Jack",
		lastName:   "Hugh",
		favFlavors: []string{"chocalate", "martini", "rum&coke"},
	}

	b := person{
		firstName:  "Jill",
		lastName:   "Jen",
		favFlavors: []string{"strawberry", "vanilla", "capuccino"},
	}

	fmt.Println(a)
	fmt.Println(b)

	people := []person{a, b}

	for _, person := range people {
		fmt.Printf("%v likes: \n", person.firstName)
		for _, icecream := range person.favFlavors {
			fmt.Println("\t", icecream)
		}
	}
}
