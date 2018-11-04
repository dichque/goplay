package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.google.com/")
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, resp.Body)
}
