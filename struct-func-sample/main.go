package main

import "golang.org/x/exp/errors/fmt"

type Rectangle struct {
	length, width int
}


func (r Rectangle) AreaByValue() int {
	return r.length * r.width
}


func (r *Rectangle) AreaByReference() int {
	return r.width * r.length
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

func main() {
	aRectangle := Rectangle{2, 3}
	fmt.Println("Rectangle area is: ", aRectangle.AreaByValue())
	fmt.Println("Rectangle area is: ", aRectangle.AreaByReference())
	fmt.Println("Rectangle area is: ", (&aRectangle).AreaByValue())
	fmt.Println("Rectangle area is: ", (&aRectangle).AreaByReference())

}
