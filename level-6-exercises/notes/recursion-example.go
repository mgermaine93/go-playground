package main

import "fmt"

func main() {
	// fmt.Println(4 * 3 * 2 * 1)
	// n := factorial(4)
	// fmt.Println(n)
	n := loopFactorial(4)
	fmt.Println(n)
}

// Using recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	// return n * factorial(4-1) * factorial(3-1) * factorial(2-1) * 1
	return n * factorial(n-1)
}

// Using a loop (Challenge)
// Takes in an argument "n" of type "int"
// Returns a value of type "int"
func loopFactorial(n int) int {
	total := 1
	// The init can be left blank because it's not needed
	for ; n > 0; n-- {
		total *= n
	}
	// 1 * 4 * 3 * 2 * 1
	return total
}

// 4 * 3 * 2 * 1 * 0
