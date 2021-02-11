// Use a loop to print out all of the years you've been alive.

package main

import "fmt"

func main() {
	birthYear := 1993
	// Current as of 2021, haha
	for birthYear <= 2021 {
		fmt.Println(birthYear)
		birthYear++
	}
}
