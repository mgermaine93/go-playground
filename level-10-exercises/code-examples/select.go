package main

import "fmt"

func main() {

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(even, odd, quit)

	// receive
	receive(even, odd, quit)

	fmt.Println("About to exit")

}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

func receive(e, o, q <-chan int) {
	// Do this until something else happens (i.e., the return)
	for {
		// Asks, "Hey, which of these channels can I pull a value off of?"
		select {
		case value := <-e:
			// Get the value off the even channel
			fmt.Println("From the even channel: ", value)
		case value := <-o:
			// Get the value off the odd channel
			fmt.Println("From the odd channel: ", value)
		case value := <-q:
			// Get the value off the even channel
			fmt.Println("From the quit channel: ", value)
			return
		}
		return
	}
}
