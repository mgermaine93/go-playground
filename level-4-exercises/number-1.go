// Using a composite literal, create an ARRAY which holds 5 VALUES of TYPE int.  Assign VALUES to each index position.  RANGE OVER the array and PRINT the values out.  Using format printing, print out the TYPE of the array.

package main

import "fmt"

func main() {

	// Using a composite literal, create an ARRAY which holds 5 VALUES of TYPE int.  Assign values to each index.
	x := []int{2, 4, 8, 16, 32}

	// RANGE OVER the array and PRINT the values out.
	for i, v := range x {
		fmt.Println(i, v)
	}

	// Using format printing, print out the TYPE of the array.
	fmt.Printf("%T\n", x)
}
