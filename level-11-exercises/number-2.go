// Starting with the following code. Create a custom error message using “fmt.Errorf”.

package main

import (
	"encoding/json"
	"fmt"
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

	bs, err := toJSON(p1)
	if err != nil {
		fmt.Println("Data was unable to be converted to JSON", err)
		return
	}

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		// Once again, needs to return a byteSlice AND and error
		return []byte{}, fmt.Errorf("data was unable to be marshaled: %v", err)
	}
	return bs, nil
}
