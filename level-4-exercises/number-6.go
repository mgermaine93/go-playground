// Create a slice to store the names of all the states in the USA.  Find the length of the slice.  Find the capacity of the slice.  Print out all of the values of the slice, along with their index positions in the slice, without using the "range" clause.

package main

import "fmt"

func main() {

	usaStates := make([]string, 50, 50)
	usaStates = []string{"Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Delaware", "Florida", "Georgia", "Hawaii", "Idaho", "Illinois", "Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana", "Maine", "Maryland", "Massachusetts", "Michigan", "Minnesota", "Mississippi", "Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey", "New Mexico", "New York", "North Carolina", "North Dakota", "Ohio", "Oklahoma", "Oregon", "Pennsylvania", "Rhode Island", "South Carolina", "South Dakota", "Tennessee", "Texas", "Utah", "Vermont", "Virginia", "Washington", "West Virginia", "Wisconsin", "Wyoming"}

	// Length
	fmt.Println(len(usaStates))

	// Capacity
	fmt.Println(cap(usaStates))

	// Print out the indices and the values
	for i := 0; i < len(usaStates); i++ {
		fmt.Println(i, usaStates[i])
	}
}
