package main

import (
	"fmt"
	"sort"
)

type User struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// Len, Swap and Less are methods required for sort interface, these will be called by the sort.Sort method: https://golang.org/pkg/sort/#Sort

// ByAge implements sort.Interface for []User based on
// the Age field.
type ByAge []User

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// ByLast implements sort.Interface for []User based on
// the Last field.
type ByLast []User

func (a ByLast) Len() int           { return len(a) }
func (a ByLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLast) Less(i, j int) bool { return a[i].Last < a[j].Last }

func main() {
	u1 := User{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := User{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := User{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []User{u1, u2, u3}

	// your code goes here
	for _, user := range users {
		fmt.Println(user.First, user.Last, user.Age)
		for _, s := range user.Sayings {
			fmt.Println("\t", s)
		}
	}
	fmt.Println("========================")
	sort.Sort(ByAge(users))
	for _, user := range users {
		fmt.Println(user.First, user.Last, user.Age)
		sort.Strings(user.Sayings)
		for _, s := range user.Sayings {
			fmt.Println("\t", s)
		}
	}
	fmt.Println("========================")
	sort.Sort(ByLast(users))
	for _, user := range users {
		fmt.Println(user.First, user.Last, user.Age)
		sort.Strings(user.Sayings)
		for _, s := range user.Sayings {
			fmt.Println("\t", s)
		}
	}
	fmt.Println("========================")
}
