// Create a value, assign it to a variable, and print the address of that variable.

package main

import "fmt"

func main() {

	// Create a value and assign it to variable
	city := "Saginaw"
	fmt.Println("This is variable's value: ", city)

	// Print the address of the variable
	fmt.Println("This is the 'address' of the variable: ", &city)

}
