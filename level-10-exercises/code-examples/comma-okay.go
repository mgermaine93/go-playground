package main

import "fmt"

func main() {

	channel := make(chan int)
	go func() {
		channel <- 14
	}()

	value, ok := <-channel

	fmt.Println(value, ok)

}
