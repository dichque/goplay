package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

//GCOUNT is number of goroutines we will spawn
const GCOUNT int = 100

func main() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("GoRoutines: ", runtime.NumGoroutine())

	var counter int32
	var wg sync.WaitGroup
	wg.Add(GCOUNT)

	for i := 0; i < GCOUNT; i++ {
		go func() {
			// Counter is a shared variable across all goroutines.
			atomic.AddInt32(&counter, 1)
			fmt.Println("Counter: ", atomic.LoadInt32(&counter))
			wg.Done()
		}()
		fmt.Println("GoRoutines: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GoRoutines: ", runtime.NumGoroutine())
	fmt.Println("Counter: ", counter)
}
