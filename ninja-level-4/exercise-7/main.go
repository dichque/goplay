package main

import "fmt"

func main() {

	j := []string{"James", "Bond", "Shaken, not stirred"}
	m := []string{"Miss", "Monneypenny", "Helloooo, James."}

	bondCharacters := [][]string{j, m}

	for characters := range bondCharacters {
		fmt.Println(characters)
		for entry := range bondCharacters[characters] {
			fmt.Println(bondCharacters[characters][entry])
		}
	}
}
