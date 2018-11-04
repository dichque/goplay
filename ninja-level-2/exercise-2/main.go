package main

import "fmt"

func main() {
	a := (42 == 42)
	b := (31 <= 41)
	c := (32 >= 21)
	d := (5 != 12)
	e := (34 > 53)
	f := (36 < 23)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
