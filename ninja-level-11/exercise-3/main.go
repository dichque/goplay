package main

import (
	"fmt"
)

type customErr struct {
	errorMsg string
}

// Interface error requires as to have a Error method which returns a string
func (a customErr) Error() string {
	return fmt.Sprintf("Custom error handler: %v", a.errorMsg)
}

func main() {
	anError := customErr{
		errorMsg: "Shit Happens!! ",
	}

	foo(anError)
}

// Foo takes param of error type and customErr is of error type implicitly because it implements method Error()
// Refer: https://godoc.org/builtin#error
func foo(e error) {
	fmt.Println(e)
}
