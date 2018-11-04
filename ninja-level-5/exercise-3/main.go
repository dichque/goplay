package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	ram := truck{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
		fourWheel: true,
	}

	bmw := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: false,
	}

	fmt.Println(ram)
	fmt.Println(bmw)

	fmt.Println(ram.doors)
	fmt.Println(bmw.color)
}
