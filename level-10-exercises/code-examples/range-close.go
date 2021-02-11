package main

import "fmt"

func main() {

	channel := make(chan int)

	// "Sending" because the arrow shows the int GOING TO the chan
	go func() {
		for i := 0; i < 100; i++ {
			channel <- i
		}
		// This tells the loop below to not wait for any more values of "channel"
		close(channel)
	}()

	// "Receiving" because the arrow shows the channel RECEIVING the chan
	// When there are no more values left on the closed channel, it leaves the range loop
	for value := range channel {
		fmt.Println(value)
	}

	fmt.Println("About to exit.")

}
