package main

import "fmt"

func aFunc() func() string {
	return func() string {
		msg := "I am a function"
		fmt.Println(msg)
		return msg
	}
}

func main() {
	f := aFunc()
	fmt.Println(f())
}
