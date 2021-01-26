// Use the "defer" keyword to show that a deferred func runs after the func containing it exits.

package main

import "fmt"

func main() {
	defer deferred()
	fmt.Println("Deferred will appear after this, just watch!")
}

func deferred() {
	fmt.Println("Hello, Deferred")
}
