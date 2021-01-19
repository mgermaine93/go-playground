// Write a program that:  assigns an int to a variable; prints that int in decimal, binary, and hex; shifts the bits of that int over 1 position to the left, and assigns that to a variable; and prints that variable in decimal, binary, and hex.

package main

import "fmt"

func main() {

	// Assign an int to a variable
	x := 14

	// Print "x" in decimal, binary, and hex
	fmt.Printf("%d\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%#x\n", x)

	// Shift the bits of that int over 1 position to the left, assigned to a new variable
	y := x << 1

	// Print that variable in decimal, binary, and hex.
	fmt.Printf("%d\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%#x\n", y)

}
