// Create your own data type "person" that will have an underlying type of "struct" so that it can store the following data: first name, last name, favorite ice cream flavors.  Create two VALUES of TYPE person.  Print out the values, ranging over the elements in the slice.

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

	fmt.Println(frazz.firstName, frazz.lastName)
	// Print the values, ranging over the elements
	for index, value := range frazz.favoriteIceCream {
		fmt.Println(index, value)
	}

	fmt.Println(plainwell.firstName, plainwell.lastName)
	for index, value := range plainwell.favoriteIceCream {
		fmt.Println(index, value)
	}

}
