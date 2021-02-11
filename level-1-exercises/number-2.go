// Use var to DECLARE three variables.  The variables should have package-level scope.  Do not assign VALUES to the variables.  Use the following IDENTIFIERS for the variables and makes sure the variables store values of the following TYPE:  identifier "x" type int, identifier "y" type string, identifier "z" type bool.
// Then, in func main, print out the values for each identifier.

package main

import "fmt"

// Package-level scope
// Declare the variables
var x int
var y string
var z bool

func main() {

	// Print out the values for each identifier
	// The values are called the "zero" values
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}
