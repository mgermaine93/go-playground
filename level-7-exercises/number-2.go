// Follow the comments below...

package main

import "fmt"

// Create a person struct
type person struct {
	firstName string
	lastName  string
	employer  string
}

func main() {

	// Create a value of type "person"
	frazzCharacter := person{
		firstName: "Edwin",
		lastName:  "Frazier",
		employer:  "Bryson Elementary",
	}

	// Print out the value
	fmt.Println("The first value of the person above is: ", frazzCharacter)

	// In func main(), call "changeMe()"
	changeMe(&frazzCharacter)

	// In func main(), print out the value
	fmt.Println("The second value of the person above is: ", frazzCharacter)

}

// Create a func called "changeMe" with *person as a parameter
func changeMe(person1 *person) {

	// Change the value stored at the *person address
	// (*person1).firstName
	person1.firstName = "Jane"
	// (*person1).lastName
	person1.lastName = "Plainwell"
	// (*person1).employer
	person1.employer = "Bryson Elementary"

}
