// Write a program that puts 100 numbers on a channel, pulls the numbers off the channel, and prints them.

package main

import "fmt"

func main() {
	channel := zeroToOneHundred()
	receiveAndPrint(channel)

	fmt.Println("about to exit")
}

func receiveAndPrint(channel <-chan int) {
	for value := range channel {
		fmt.Println(value)
	}
}

// Put 0-100 on a channel and then return the channel
func zeroToOneHundred() <-chan int {
	channel := make(chan int)

	go func() {
		for i := 0; i < 101; i++ {
			channel <- i
		}
		// channel is closed here
		close(channel)
	}()

	return channel
}
