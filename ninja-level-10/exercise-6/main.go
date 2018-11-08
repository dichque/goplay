package main

import "fmt"

// Creates a channel with 100 buffer to store and get data
func main() {

	ch := make(chan int, 100)

	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch)

	for n := range ch {
		fmt.Println(n)
	}
	main1()
}

// Creates a go routine to pump 100 numbers into a channel
// Uses range to listem close of channel befpre printing all the messages
func main1() {

	ch := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

}
