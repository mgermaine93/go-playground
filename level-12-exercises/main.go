package main

import "fmt"

type doggo struct {
	name  string
	breed string
	age   int
}

func main() {
	doggy := doggo{
		name:  "Benny",
		breed: "Beagle",
		age:   dog.HumanYearsToDogYears(5),
	}
	fmt.Printf("%v the %v is %v dog years-old.\n", doggy.name, doggy.breed, doggy.age)
}
