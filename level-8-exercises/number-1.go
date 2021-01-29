// Starting with this code, marshal the []user to JSON. There is a little bit of a curve ball in this hands-on exercise - remember to ask yourself what you need to do to EXPORT a value from a package.

package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	// These need to be capitalized (as they are now) in order to export them from a package
	First string
	Age   int
}

func main() {
	u1 := user{
		// Same here...
		First: "James",
		Age:   32,
	}

	u2 := user{
		// And here...
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		// And here...
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here

	// We know now that Marshal() takes in a VALUE of any TYPE and returns a BYTESLICE and an ERROR

	byteSlice, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	// This will not return anything unless the fields are capitalized, as shown above in the struct definition
	// What is returns is Marshalled JSON
	fmt.Println(string(byteSlice))
}
