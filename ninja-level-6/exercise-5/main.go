package main

import (
	"fmt"
	"math"
)

type circle struct {
	r float64
}
type square struct {
	l float64
}

type shape interface {
	area()
}

func info(a shape) {
	a.area()
}

func (a circle) area() {
	fmt.Println(math.Pi * a.r * a.r)
}

func (a square) area() {
	fmt.Println(a.l * a.l)
}

func main() {

	aSquare := square{
		l: 2.,
	}

	aCircle := circle{
		r: 4.0,
	}

	info(aCircle)
	info(aSquare)
}
