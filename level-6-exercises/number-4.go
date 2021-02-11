// Follow the comments below...

package main

import "fmt"

// Create a user-defined struct with the identifier "person" and the fields "first", "last", and "age".
type person struct {
	first string
	last  string
	age   int
}

// Attach a method to type "person" with the identifier "speak".
// The method should have the person say their name and age
// This is a method
func (person1 person) speak() {
	fmt.Printf("I am %v %v.  I am %v years old.\n", person1.first, person1.last, person1.age)
}

func main() {

	// Create a VALUE of TYPE person
	alexRider := person{
		"Alex",
		"Rider",
		18,
	}

	// Call the METHOD from the value of type "person"
	alexRider.speak()

}
