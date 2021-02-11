// This will NOT run!
// When you send and receive on a channel, the transaction cannot occur until "send" and "receive" can happen at the same time.

package main

import "fmt"

// ... leaving the main Goroutine good to go
func main() {

	// This channel goes "Pyommmmmmmmm!"
	channel := make(chan int)

	// This Goroutine is off an running... ^
	go func() {
		// "Put on the number 42"
		channel <- 42
	}()

	// "Take off the number 42"
	fmt.Println(<-channel)
}
