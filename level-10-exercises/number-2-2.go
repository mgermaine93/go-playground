// Get this code running.

package main

import "fmt"

func main() {

	// Once again, the issue is here.
	// cr := make(<-chan int)

	// Just make a bidirectional channel
	cr := make(chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
