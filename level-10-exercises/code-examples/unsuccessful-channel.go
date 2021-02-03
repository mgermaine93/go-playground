// This will NOT run!
// When you send and receive on a channel, the transaction cannot occur until "send" and "receive" can happen at the same time.

package main

import "fmt"

func main() {

	channel := make(chan int)

	// "Put on the number 42"
	channel <- 42

	// "Take off the number 42"
	fmt.Println(<-channel)
}
