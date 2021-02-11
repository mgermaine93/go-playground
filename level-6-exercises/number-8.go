// Follow the comments below...

package main

import "fmt"

func main() {

	// Assign the returned func to a variable
	result1 := sayHello()

	// Call the returned func
	result2 := result1()

	// Print the result of the called function
	fmt.Println(result2)
}

// Create a func that returns a func, which returns a string
func sayHello() func() string {
	return func() string {
		return "Hello"
	}

}
