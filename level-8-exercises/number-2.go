// Unmarshal JSON into a Go data structure

package main

import (
	"encoding/json"
	"fmt"
)

type bondCharacter struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {

	// This is the JSON data
	characterJSON := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(characterJSON)

	// Unmarshalling
	bs := []byte(characterJSON)

	var characters []bondCharacter

	newError := json.Unmarshal(bs, &characters)

	if newError != nil {
		fmt.Println(newError)
	}
	fmt.Println("Here's all of the data: ", characters)

	for index, person := range characters {
		fmt.Println("BOND CHARACTER NUMBER", index)
		fmt.Println(person.First)
		fmt.Println(person.Last)
		fmt.Println(person.Age)
		fmt.Println("THIS CHARACTER OFTEN SAYS THE FOLLOWING:")

		// Iterate through each saying for fun
		for _, saying := range person.Sayings {
			fmt.Println("\t", saying)
		}
		// fmt.Println(person.Sayings)
	}

}
