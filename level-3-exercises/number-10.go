// Create a program that uses a switch statement with no switch expression specified

package main

import "fmt"

func main() {

	number := 4

	switch {
	case number > 4:
		fmt.Printf("%v is greater than 4!\n", number)
	case number == 4:
		fmt.Printf("%v is equal to 4!\n", number)
	case number < 4:
		fmt.Printf("%v is less than 4!\n", number)
	}
}
