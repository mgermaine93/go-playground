// Follow the steps in the comments below...

package main

import "fmt"

func main() {

	// Start with this slice
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// Append "52" to the slice
	x = append(x, 52)

	// Print out the new slice
	fmt.Println(x)

	// In ONE statement, append to that slice the following values:  53, 54, 55.
	x = append(x, 53, 54, 55)

	// Print out the new slice
	fmt.Println(x)

	// Append a new slice onto the latest one
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)

	// Print out the new slice
	fmt.Println(x)
}
