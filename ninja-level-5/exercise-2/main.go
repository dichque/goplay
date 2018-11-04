package main

import "fmt"

type person struct {
	firstName  string
	lastName   string
	favFlavors []string
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
		favFlavors: []string{"strawberry", "vanilla", "cappuccino"},
	}

	// fmt.Println(a)
	// fmt.Println(b)

	people := map[string]person{
		a.lastName: a,
		b.lastName: b,
	}

	for k, v := range people {
		fmt.Println(k, v)
		for idx, val := range v.favFlavors {
			fmt.Println(idx, val)
		}
	}
}
