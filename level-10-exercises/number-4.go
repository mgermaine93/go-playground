// Starting with the provided code, pull the values off the channel using a select statement

package main

import "fmt"

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func receive(c, q <-chan int) {
	for {
		// Asks, "Hey, which of these channels can I pull a value off of?"
		select {
		case value := <-c:
			// Get the value off the c channel
			fmt.Println("From the c channel: ", value)
		case <-q:
			return
		}
	}
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		// We're done putting values on "c" here
		close(c)
	}()

	return c
}
