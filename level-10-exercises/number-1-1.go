// Get this code working using a func literal

package main

import "fmt"

func main() {
	c := make(chan int)

	// Code is blocking right here.
	go func() {
		// "Put on the number 42"
		c <- 42
	}()
	// "Take off the number 42"
	fmt.Println(<-c)
}
