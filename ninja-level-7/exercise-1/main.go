package main
import "fmt"

type person struct {
  first, last string
  age int
}

func changeMe(p *person)  {
  p.first = "Jagadish"
  p.last = "Nagarajaiah"
  p.age = 41

  // This also provides same effect of dereferencing
    // (*p).first = "Jagadish"
    // (*p).last = "Nagarajaiah"
    // (*p).age = 41
}

func main()  {
  aPerson := person {
    first: "Jaggi",
    last: "None",
    age: 0,
  }

  fmt.Println(aPerson)
  changeMe(&aPerson)
  fmt.Println(aPerson)
}
