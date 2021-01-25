// Using the previous code, store the values of type person in a map with the key of last name.  Access each value in the map.  Print out the values, ranging over the slice.

package main

import "fmt"

func main() {

	// Create the data type
	type person struct {
		firstName        string
		lastName         string
		favoriteIceCream []string
	}

	// Create the values
	frazz := person{
		firstName: "Edwin",
		lastName:  "Frazier",
		favoriteIceCream: []string{
			"Mackinac Island Fudge",
			"Moose Tracks",
			"Superman"},
	}

	plainwell := person{
		firstName: "Jane",
		lastName:  "Plainwell",
		favoriteIceCream: []string{
			"Traverse City Cherry",
			"Strawberry",
			"Mint Chocolate Chip"},
	}

	// Key is a string, person is the value
	m := map[string]person{
		frazz.lastName:     frazz,
		plainwell.lastName: plainwell,
	}

	// Range over the values.  "_" throws away the key value so you don't need to use it in the loop
	for _, value := range m {
		fmt.Println(value.firstName, value.lastName)
		for index, val := range value.favoriteIceCream {
			fmt.Println(index, val)
		}
	}
}
