// Create and use an anonymous struct

package main

import "fmt"

func main() {

	// Create the anonymous struct
	frazz := struct {
		// Map
		name    map[string]string
		hobbies []string
	}{
		// Map
		name: map[string]string{
			"firstName": "Edwin",
			"lastName":  "Frazier",
		},
		// Slice
		hobbies: []string{
			"Running",
			"Cycling",
			"Swimming",
			"Basketball",
			"Playing guitar",
		},
	}

	// Use it (just print it out)
	fmt.Println(frazz)

}
