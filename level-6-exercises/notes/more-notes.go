package main

import "fmt"

func main() {

	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(ii...)
	fmt.Println("all numbers", s)

	s2 := even(sum, ii...)
	fmt.Println("even numbers", s2)

	s3 := odd(sum, ii...)
	fmt.Println("odd numbers", s3)
}

// Takes a value of type slice of int, returns an int,
func sum(x ...int) int {
	fmt.Printf("%T\n", x)
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

// This is a callback
// It takes in "func(x ...int) int" as an argument, assigned to the variable "f"
// It also takes in a variadic parameter of ints
func even(f func(x ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	// This returns an int
	return f(yi...)
}

// Challenge: print all odd numbers
func odd(f func(x ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	// This returns an int
	return f(yi...)
}
