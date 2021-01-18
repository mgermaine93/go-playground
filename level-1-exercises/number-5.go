// Build off of problem #4...

package main

import "fmt"

// Create your own type, with underlying type of int
type guineaPig int

// Create a new variable of the type above
var x guineaPig

// Create a new variable with the identifier "y".  Type should be the underlying type above
var y int

func main() {

	// Print of the value of "x"
	fmt.Println(x)

	// Print out the type of "x"
	fmt.Printf("%T\n", x)

	// Assign the value of 14 to "x"
	x = 14

	// Print out the value of "x"
	fmt.Println(x)

	// Use CONVERSION to convert the TYPE of the VALUE stored in "x" to the UNDERLYING TYPE and ASSIGN that value to "y"
	y = int(x)

	// Print out the value stored in "y"
	fmt.Println(y)

	// Format print out the type of the value stored in "y"
	fmt.Printf("%T\n", y)

}
