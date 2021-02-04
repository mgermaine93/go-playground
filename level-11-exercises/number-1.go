// Starting with the following code, do not throw the error away using a blank identifier.  Instead, make sure the code is checking and handling the error.

// The purposes of this program is to marshal a struct to JSON and print it out.

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	// Blank identifier was here
	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatal("Whoops, an error happened: ", err)
	}
	fmt.Println(string(bs))

}
