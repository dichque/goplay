package main

import "time"
import "fmt"

type myTime struct {
	time.Time //anonymous field
}

func (t myTime) first5Chars() string {
	return t.Time.String()[0:5]
}

func main() {
	aTime := myTime{time.Now()}  //since time.LocalTime returns an address, we convert it to a value with *
	fmt.Println("Full time now:", aTime.String()) //calling existing String method on anonymous Time field
	fmt.Println("First 5 chars of time:", aTime.first5Chars()) //calling myTime.first5Chars
}
