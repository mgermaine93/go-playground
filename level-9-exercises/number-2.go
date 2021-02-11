// Follow the comments below...

package main

import "fmt"

// Create a a type person struct
type person struct {
	firstName string
}

// Create a type human interface
// "To implicitly implement the interface, a human must have the speak method."
type human interface {
	speak()
}

// Attach a method "speak" to type pointer to a person
func (person1 *person) speak() {
	greeting := "Hi, I'm " + person1.firstName
	fmt.Println(greeting)
}

// Create a func "saySomething"
// Have it take a human as a parameter.
// Have it call the "speak" method.
func saySomething(human1 human) {
	human1.speak()
}

func main() {
	// Show the following in your code:
	frazzCharacter := person{firstName: "Frazz"}
	// You CAN pass a value of type *person into saySomething()
	saySomething(&frazzCharacter)
	// You CANNOT pass a value of type person into saySomething()
	// As a result, the line below will be an error
	// saySomething(frazzCharacter)
}
