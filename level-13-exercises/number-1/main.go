package main

import (
	"github.com/mgermaine93/learn-to-go/level-12-exercises/dog"
	"fmt"
)

type doggo struct {
	name  string
	breed string
	age   int
}

func main() {
	benny := doggo{
		name:  "Benny",
		breed: "Beagle",
		age:   dog.HumanYearsToDogYears(5),
	}
	fmt.Printf("My favorite dog is %v the %v.  He is %v dog-years old.\n", benny.name, benny.breed, benny.age)
}
