// Write a program that prints a number in decimal, binary, and hex

package main

import "fmt"

func main() {
	x := 14

	// Normal int
	fmt.Println(x)

	// Decimal
	fmt.Printf("%d\n", x)

	// Binary
	fmt.Printf("%b\n", x)

	// Hex
	fmt.Printf("%#x\n", x)
}
