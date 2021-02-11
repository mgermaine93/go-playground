// Get this code running.

package main

import "fmt"

func main() {
	// Here is the problem...
	// "On a channel, we can send an int"
	// cs := make(chan<- int)

	// Can just make it a bidirectional channel
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	// Previously, we were trying here to receive from a channel that could only send stuff
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
