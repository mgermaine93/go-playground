package main

import "fmt"

func main() {
	// This buffered channel will allow one channel to sit in it
	channel := make(chan int, 1)

	channel <- 42

	// Since the buffered channel can only hold one value (seen with the "1" above), this will not run because the the value below is too much for it too hold
	channel <- 14

	fmt.Println(<-channel)
}
