package main

import "fmt"

func main() {
	fmt.Println("This should be 2: ", addOneAndOne())
}

func addOneAndOne() int {
	sum := 1 + 1
	return sum
}
