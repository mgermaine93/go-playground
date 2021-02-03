// Show the comma-OK idiom starting with the provided code

package main

import "fmt"

func main() {
	c := make(chan int)

	// Put something on the channel
	go func() {
		c <- 14
	}()

	// Is there a value on the channel? (Yes)
	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	// What about now? (No)
	v, ok = <-c
	fmt.Println(v, ok)
}

// func main() {
// 	c := make(chan int)

// 	v, ok :=
// 	fmt.Println(v, ok)

// 	close(c)

// 	v, ok =
// 	fmt.Println(v, ok)
// }
