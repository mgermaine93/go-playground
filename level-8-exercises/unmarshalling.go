// Unmarshalling Data to and from JSON

package main

import (
	"encoding/json"
	"fmt"
)

type frazzCharacter struct {
	FirstName  string `json:"FirstName"`
	LastName   string `json:"LastName"`
	Occupation string `json:"Occupation"`
	Employer   string `json:"Employer"`
}

func main() {

	// Unmarshaling
	s := `[{"FirstName":"Mr.","LastName":"Burke","Occupation":"Teacher","Employer":"Bryson Elementary"},{"FirstName":"Mr.","LastName":"Spaetzle","Occupation":"Principal","Employer":"Bryson Elementary"}]`

	bs := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	// people := []person{}
	var characters []frazzCharacter // this is the same as above, but more readable

	newError := json.Unmarshal(bs, &characters)

	if newError != nil {
		fmt.Println(newError)
	}
	fmt.Println("Here's all of the data: ", characters)

	for i, v := range characters {
		fmt.Println("\n FRAZZ CHARACTER NUMBER", i)
		fmt.Println(v.FirstName)
		fmt.Println(v.LastName)
		fmt.Println(v.Occupation)
		fmt.Println(v.Employer)
	}

}
