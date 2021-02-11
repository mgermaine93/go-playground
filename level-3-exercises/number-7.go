// Print out the remainder (modulus) for each number between 10 and 100 when it is divided by 4.

package main

import "fmt"

func main() {

	for i := 10; i <= 100; i++ {
		// %v gets the value of what is passed in after what's in the quotes
		fmt.Printf("When %v is divided by 4, the remainder is %v.\n", i, i%4)
	}
}
