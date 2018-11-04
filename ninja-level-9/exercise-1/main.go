package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printOdds(a []int) {
	for num := range a {
		if num%2 != 0 {
			fmt.Println(num)
		}
	}
	wg.Done()
}

func printEvens(a []int) {
	for num := range a {
		if num%2 == 0 {
			fmt.Println(num)
		}
	}
	wg.Done()
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	wg.Add(1)
	go printOdds(numbers)
	wg.Add(1)
	go printEvens(numbers)
	wg.Wait()
}
