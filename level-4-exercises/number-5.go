// To DELETE from a slice, we use APPEND along with SLICING.  For this exercise, follow the steps below...

package main

import "fmt"

func main() {

	// Start with this slice
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// Use APPEND and SLICING to get the values [42, 43, 44, 48, 49, 50, 51].
	// Deleting from a slice is possible, too
	x = append(x[:3], x[6:]...)

	// Print the value above
	fmt.Println(x)

}
