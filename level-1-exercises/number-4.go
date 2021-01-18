// Create your own type.  Have the underlying type be an int.
// Create a VARIABLE of your new TYPE with the IDENTIFIER "x" using the "VAR" keyword.
// In func main:  print out the value of the variable "x", print out the type of the variable "x", assign 42 to the VARIABLE "x" using the "=" OPERATOR, print out the value of the variable "x"

package main

import "fmt"

// Create your own type, with underlying type of int
type guineaPig int

// Create a new variable of the type above
var x guineaPig

func main() {

	// Print of the value of "x"
	fmt.Println(x)

	// Print out the type of "x"
	fmt.Printf("%T\n", x)

	// Assign the value of 42 to "x"
	x = 42

	// Print out the value of "x"
	fmt.Println(x)

}
