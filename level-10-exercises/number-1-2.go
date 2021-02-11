// Get this code working using a buffered channel

package main

import "fmt"

func main() {

	// Have only one channel sit on it
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
