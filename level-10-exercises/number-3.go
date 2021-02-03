// Starting with the provided code, pull the values off the channel using a "for range" loop.

package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

// This will block until c is closed...
func receive(c <-chan int) {
	for value := range c {
		fmt.Println(value)
	}
}

// This is putting values 0-100 on channel and then returning the channel
func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		// c is closed here, resolving the problem mentioned above
		close(c)
	}()

	return c
}
