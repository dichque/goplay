package main

import (
	"fmt"
)

func main() {
	ctrlCh := make(chan int)
	dataCh := gen(ctrlCh)

	receive(dataCh, ctrlCh)

	fmt.Println("about to exit")
}

func gen(ctrlCh chan<- int) <-chan int {
	dataCh := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			dataCh <- i
		}
		// Close data channel cleanly after job is done
		close(dataCh)
		// Send 1 to channel ctrlCh when we are done
		ctrlCh <- 1
	}()

	return dataCh
}

func receive(dataChannel <-chan int, controlChannel <-chan int) {
	for {
		select {
		case msg := <-dataChannel:
			fmt.Println(msg)
		case <-controlChannel: // If we receive 1 (true) from control channel exit
			return
		}
	}
}
