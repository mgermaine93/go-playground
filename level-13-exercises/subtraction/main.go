package main

import "fmt"

func main() {
	fmt.Println("This should be 0: ", subtractOneFromOne())
}

func subtractOneFromOne() int {
	difference := 1 - 1
	return difference
}