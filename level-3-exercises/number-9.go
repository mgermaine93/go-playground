// Create a program that shows the "if" statement in action

package main

import "fmt"

func main() {

	number := 4

	if number > 4 {
		fmt.Printf("%v is greater than 4!\n", number)
	} else if number == 4 {
		fmt.Printf("%v is equal to 4!\n", number)
	} else {
		fmt.Printf("%v is less than 4!\n", number)
	}
}
