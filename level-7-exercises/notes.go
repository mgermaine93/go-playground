package main

import "fmt"

func main() {
	a := 14

	// This is the value of the variable
	fmt.Println(a)

	// This is the address in memory where the value of the variable is stored
	fmt.Println(&a)

	// This is the TYPE of the VALUE of the VARIABLE
	fmt.Printf("%T\n", a)

	// This is the TYPE of the ADDRESS to the VALUE to the VARIABLE
	// "*int" is a "POINTER to an INT"
	fmt.Printf("%T\n", &a)

	// This is the VALUE stored at the ADDRESS to the VALUE to the VARIABLE (when you have the address)
	// AKA, the value of the variable
	fmt.Println(*&a)

	x := 0

	// Prints out the POINTER/ADDRESS to "x"
	fmt.Println("x before", &x)

	// Prints out the value of "x"
	fmt.Println("x before", x)

	// Calls "foo" on the address in memory where the value of "x" is stored
	foo(&x)

	// Prints out the same POINTER/ADDRESS to "x" after "foo" is run.  The address won't change.
	fmt.Println("x after", &x)

	// Prints out the VALUE of "x" after "foo" is run.  The value DOES change.
	fmt.Println("x after", x)

}

// The value passed in is a POINTER to an INT
func foo(y *int) {

	// Prints out the ADDRESS to y, because y is a POINTER to an INT
	// It's the same as "x" above because "x" is passed in to the "foo"
	fmt.Println("y before", y)

	// Prints out the value stored at the address to "y"... same as "x" above
	fmt.Println("y before", *y)

	// REASSIGNS the value stored at address y (also x)
	*y = 43

	// The address will stay the same, since that didn't change...
	fmt.Println("y after", y)

	// ... but this does changes, since the VALUE at the address is different.
	fmt.Println("y after", *y)
}
