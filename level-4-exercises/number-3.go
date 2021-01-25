// Using code from #2, using SLICING to print out various SLICES of the original array.

package main

import "fmt"

func main() {

	// Using a composite literal, create a SLICE which holds 10 VALUES of TYPE int.  Assign values to each index.
	x := []int{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}

	// Get the slices
	firstSlice := x[:6]
	secondSlice := x[6:]
	thirdSlice := x[2:7]
	fourthSlice := x[1:6]

	// Print the slices
	fmt.Println(firstSlice)
	fmt.Println(secondSlice)
	fmt.Println(thirdSlice)
	fmt.Println(fourthSlice)
}
