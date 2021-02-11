// <-chan int === "receiving from a CHANNEL of INT"
// chan<- int === "on a channel, I'm SENDING an INT"

package main

import "fmt"

func main() {
	// This buffered channel will allow one channel to sit in it
	channel := make(chan int, 2)

	// Put values on the channel
	channel <- 42
	channel <- 43

	// Pull values off the channel
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println("----------------")
	fmt.Printf("%T\n", channel)
}
