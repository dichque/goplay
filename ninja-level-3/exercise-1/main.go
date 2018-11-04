package main

import "fmt"

func main() {

	for i := 10; i <= 100; i++ {
		fmt.Printf("Remainder when %v is divided by 4: %v\n", i, i%4)
	}
}
