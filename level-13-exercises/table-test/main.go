package main

import "fmt"

func main() {
	fmt.Println("101 + 2 = ", sumUpTheseNumbers(101, 2))
	fmt.Println("98 + 53 = ", sumUpTheseNumbers(98, 53))
	fmt.Println("0 + 8 = ", sumUpTheseNumbers(0, 8))
}

func sumUpTheseNumbers(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}