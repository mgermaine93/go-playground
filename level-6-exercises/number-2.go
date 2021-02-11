// Follow the steps below...

package main

import "fmt"

func main() {

	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Variadic parameters need to be unfurled!
	outputOfFoo := foo(input...)
	fmt.Println(outputOfFoo)

	// Slices do not need to be unfurled
	outputOfBar := bar(input)
	fmt.Println(outputOfBar)
}

// Create a func with the identifier "foo" that takes in a variadic parameter of type "int"; pass in a value of type []int into your func, making sure to unfurl the []int; and return the sum of all values of type int passed in.
func foo(input ...int) int {
	total := 0
	for _, v := range input {
		total += v
	}
	return total
}

// Create a func with the identifier "bar" that takes in a parameter of type []int and returns the sum of all values of type int passed in.
func bar(input []int) int {
	total := 0
	for _, v := range input {
		total += v
	}
	return total
}
