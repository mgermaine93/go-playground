// Marshalling Data to and from JSON
// Cool website == rawgit.com
// Another cool website == jsontogo.com
package main

import (
	"encoding/json"
	"fmt"
)

type frazzCharacter struct {
	FirstName  string
	LastName   string
	Occupation string
	Employer   string
}

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {

	burke := frazzCharacter{
		FirstName:  "Mr.",
		LastName:   "Burke",
		Occupation: "Teacher",
		Employer:   "Bryson Elementary",
	}

	spaetzle := frazzCharacter{
		FirstName:  "Mr.",
		LastName:   "Spaetzle",
		Occupation: "Principal",
		Employer:   "Bryson Elementary",
	}

	characters := []frazzCharacter{burke, spaetzle}

	fmt.Println(characters)

	// From GoDoc:  func Marshal(v interface{}) ([]byte, error)
	// We know now that Marshal() takes in a VALUE of any TYPE and returns a BYTESLICE and an ERROR

	byteSlice, err := json.Marshal(characters)
	if err != nil {
		fmt.Println(err)
	}
	// This will not return anything unless the fields are capitalized, as shown above in the struct definition
	// What is returns is JSON
	fmt.Println(string(byteSlice))

	// Unmarshaling
	s := `[{"FirstName":"Mr.","LastName":"Burke","Occupation":"Teacher","Employer":"Bryson Elementary"},{"FirstName":"Mr.","LastName":"Spaetzle","Occupation":"Principal","Employer":"Bryson Elementary"}]`

	bs := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	// people := []person{}
	var people []person // this is the same as above, but more readable

	newError := json.Unmarshal(bs, &people)

	if newError != nil {
		fmt.Println(newError)
	}
	fmt.Println("Here's all of the data: ", people)

	for i, v := range people {
		fmt.Println("\n PERSON NUMBER", i)
		fmt.Println(v.First)
		fmt.Println(v.Last)
		fmt.Print(v.Age)
	}

}
