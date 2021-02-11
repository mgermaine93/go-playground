// Write a program that launches 10 goroutines (each goroutine should add 10 numbers to a channel) and then pull the numbers off the channel and print them.

package main

import "fmt"

func main() {

	channel := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				channel <- j
			}
		}()
	}

	// 100 because we know how many values we'll need to print out
	for k := 0; k < 100; k++ {
		fmt.Println(k, <-channel)
	}

	fmt.Println("Exiting here...")

}
