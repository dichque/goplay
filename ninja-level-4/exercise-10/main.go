package main

import "fmt"

func main() {

	charFav := map[string][]string{
		"bond_james":      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	delete(charFav, `no_dr`)

	for artist := range charFav {
		fmt.Printf("%v likes\n", artist)
		for theirToys := range charFav[artist] {
			fmt.Println(charFav[artist][theirToys])
		}
	}
}
