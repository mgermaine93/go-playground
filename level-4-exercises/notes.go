package main

import "fmt"

func main() {

	// This creates a composite literal of x, and slice containing ints
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)

	// This gets the length of x, declared above
	fmt.Println(len(x))

	// This prints out each index of slice "x" and the value assigned to each index.
	for i, v := range x {
		fmt.Println(i, v)
	}

	// This prints out the first two values in slice "x".
	// The value before the colon is included, the value after is NOT.
	fmt.Println(x[0:2])

	// This "appends" additional values to the END of "x".  The value of x changes as a result.
	x = append(x, 77, 12, 34, 56, 204)
	fmt.Println(x)

	// You can also "append" the values of one slice onto another
	y := []int{1, 2, 3, 4, 5, 6}
	x = append(x, y...)
	fmt.Println(x)

	// Deleting from a slice is possible, too
	x = append(x[:3], x[6:]...)
	fmt.Println(x)

	// If you already know what the size of your slice is going to be, use the "make" function.  (The "12" is the cap/capacity of the slice.  "10" is the length for now until more values are added.)
	z := make([]int, 10, 12)
	fmt.Println(z)

	// Slices of slices go like this
	pittsburgh := []string{"Point Breeze", "Bloomfield", "Lawrenceville", "Homewood"}
	fmt.Println(pittsburgh)
	detroit := []string{"Hamtramck", "Midtown", "Corktown", "Belle Isle"}
	fmt.Println(detroit)

	// Prints out both of the above in a single "slice of slices"
	pittsburghAndDetroit := [][]string{pittsburgh, detroit}
	fmt.Println(pittsburghAndDetroit)

	// Maps
	m := map[string]string{
		"Pittsburgh": "Pennsylvania",
		"Detroit":    "Michigan",
	}
	// Prints out the entire whole map
	fmt.Println(m)
	// Prints out the value for the "Pittsburgh" key
	fmt.Println(m["Pittsburgh"])
	// Prints out the value for the "Detroit" key
	fmt.Println(m["Detroit"])

	// Checks whether or not the value actually exists in the map... "comma ok idiom"
	v, ok := m["Phoenix"]
	// Prints out the value
	fmt.Println(v)
	// Prints out whether or not it exists in the map
	fmt.Println(ok)

	// Shorthand
	if v, ok := m["Dallas"]; ok {
		fmt.Println("THIS IS THE IF PRINT:", v)
	} else {
		fmt.Println("Not found in the map")
	}

	// Adding to a map
	m["Austin"] = "Texas"
	fmt.Println(m)

	// Print out each of the key/value pairs in the map
	for k, v := range m {
		fmt.Println(k, v)
	}

	// Delete an entry from a map
	delete(m, "Austin")
	fmt.Println(m)

	// Strangely, it is possible to delete an entry that doesn't exist in the map
	delete(m, "Aurora")
	fmt.Println(m)

}
