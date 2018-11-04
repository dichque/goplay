package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("Remainder when %v is div by 4: %v\n", i, i%4)
	}
}
