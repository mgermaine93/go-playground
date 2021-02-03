package main

import "fmt"

func main() {

	channel := make(chan int)

	// send (go routine)
	go send(channel)

	// receive
	receive(channel)

	fmt.Println("About to exit")

}

// "Sending" because the arrow shows the int GOING TO the chan
func send(channel chan<- int) {
	channel <- 14
}

// "Receiving" because the arrow shows the channel RECEIVING the chan
func receive(channel <-chan int) {
	fmt.Println(<-channel)
}
