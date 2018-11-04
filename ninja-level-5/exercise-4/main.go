package main

import "fmt"

func main() {

	fruit := struct {
		name     string
		color    string
		size     int
		location []string
		eaters   map[string]int
	}{
		name:  "Jack Fruit",
		color: "Green",
		size:  11,
		location: []string{
			"kerala",
			"karnataka",
			"tamilnadu"},
		eaters: map[string]int{
			"Jag":  32,
			"Jill": 21,
			"Jen":  11,
		},
	}

	fmt.Println(fruit.name)
	fmt.Println(fruit.color)
	fmt.Println(fruit.size)
	fmt.Println(fruit.location)
	fmt.Println(fruit.eaters)

	for _, l := range fruit.location {
		fmt.Println(l)
	}

	for p, age := range fruit.eaters {
		fmt.Println(p, age)
	}
}
