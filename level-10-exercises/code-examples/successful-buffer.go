package main

import "fmt"

func main() {
	// This buffered channel will allow one channel to sit in it
	channel := make(chan int, 1)

	channel <- 42
	channel <- 43

	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
