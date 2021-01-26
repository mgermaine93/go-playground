// Follow the steps below...

package main

import "fmt"

func main() {

	// Call both funcs
	foo()
	bar()

	// Print out their results
	fmt.Println(foo())
	fmt.Println(bar())

}

// Create a func with the identifier "foo" that returns an int
func foo() int {
	return 14
}

// Create a func with the identifier "bar" that returns an int and a string
func bar() (string, int) {
	return "Shanahan", 14
}
