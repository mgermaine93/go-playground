// Print out the even and odd numbers from 0 to 100, inclusive.

package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		// Even numbers
		if i%2 == 0 {
			result := fmt.Sprintf("This is an even number: %d", i)
			fmt.Println(result)
		}
		// Odd numbers
		if i%2 != 0 {
			result := fmt.Sprintf("This is an odd number: %d", i)
			fmt.Println(result)
		}

	}
}
