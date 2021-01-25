// Structs

// We create VALUES of a certain TYPE that are stored in VARIABLES
// These VARIABLES have IDENTIFIERS

package main

import "fmt"

// Structs are data structures that are composed of VALUES of different TYPES.  Also known as an aggregate/composite/complex data type.
// Don't say "object"... say "VALUE of TYPE"

// Regular struct
type person struct {
	firstName string
	lastName  string
	age       int
}

// Embedded Structs "adopts"/"inherits" the PERSON struct
type comicCharacter struct {
	person
	inspiredByRealPerson bool
}

func main() {

	// A VALUE of TYPE comicCharacter
	// PERSON is the INNER TYPE, that gets promoted to the OUTER TYPE of COMICCHARACTER
	character1 := comicCharacter{
		person: person{
			firstName: "Edwin",
			lastName:  "Frazier",
			age:       33,
		},
		inspiredByRealPerson: true,
	}

	// Regular struct
	person2 := person{
		firstName: "Jane",
		lastName:  "Plainwell",
		age:       32,
	}

	// This is an anonymous struct, because it doesn't have a name (unlike "person" and "comicCharacter" above).
	// Use them to prevent code pollution.
	// Best used when you only need a struct in one specific area.
	// Keeps code "lean".
	anonymousStruct := struct {
		firstNumber  int
		secondNumber int
		thirdNumber  int
	}{
		firstNumber:  1,
		secondNumber: 2,
		thirdNumber:  3,
	}

	fmt.Println(character1)

	// You can use dot notation to access the values within the struct
	fmt.Println(character1.firstName, character1.lastName, character1.age, character1.inspiredByRealPerson)

	fmt.Println(person2)

	fmt.Println(anonymousStruct)
}
