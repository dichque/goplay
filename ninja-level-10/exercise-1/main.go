package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1) // Channel with a buffer of 1

	// We are able to send 43 to channel as it can hold a buffer of 1
	c <- 43

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
