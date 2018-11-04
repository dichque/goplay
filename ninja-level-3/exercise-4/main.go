package main

import "fmt"

func main() {
	years := 1977
	for {
		if years <= 2018 {
			fmt.Println(years)
		}
		years++
	}
}
