// Using the previous exercise, assign the following values to the three variables at package-level scope:  42 assigned to x, "James Bond" assigned to y, and true assigned to z.
// Then, in func main, use fmt.Sprintf() to print all of the VALUES to one single string.  ASSIGN the returned value of TYPE string using the short declaration operator to a variable with the IDENTIFIER "s"

package main

import "fmt"

// Package-level scope
// Declare the variables and assign values
var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	// Use fmt.Sprintf() to print all of the VALUES to one single string
	// ASSIGN the returned value of TYPE string using the short declaration operator to a variable with the IDENTIFIER "s"
	s := fmt.Sprintf("%v\n%v\n%v", x, y, z)

	// Print out the value stored by the variable
	fmt.Println(s)
}
