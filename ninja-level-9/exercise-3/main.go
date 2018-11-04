package main

import (
	"fmt"
	"runtime"
	"sync"
)

//GCOUNT is number of goroutines we will spawn
const GCOUNT int = 50

func main() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("GoRoutines: ", runtime.NumGoroutine())

	counter := 0
	var wg sync.WaitGroup
	wg.Add(GCOUNT)

	for i := 0; i < GCOUNT; i++ {
		go func() {
			a := counter
			// Yields to other go routines
			runtime.Gosched()
			a++
			counter = a
			wg.Done()
		}()
		fmt.Println("GoRoutines: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GoRoutines: ", runtime.NumGoroutine())
	fmt.Println(counter)
}
